package command

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/alexhokl/helper/datetime"
	"github.com/alexhokl/helper/git"
	"github.com/spf13/cobra"

	"github.com/alexhokl/bb/internal/pipelineapi"
	"github.com/alexhokl/bb/models"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-pipelines-pipeline-uuid-get

type getPipelineOptions struct {
	buildNumber int32
	format      string
	wait        bool
	interval    int
}

var getPipelineOpts getPipelineOptions

func (opts *getPipelineOptions) validate() error {
	if opts.buildNumber < 0 {
		return fmt.Errorf("invalid build number")
	}
	if opts.wait && opts.interval < 1 {
		return fmt.Errorf("interval must be at least 1 second")
	}
	return validatePipelineFormat(opts.format)
}

// pipelineJSON is the JSON output of a pipeline together with its steps.
type pipelineJSON struct {
	pipelineapi.Pipeline
	Steps []pipelineapi.Step `json:"steps"`
}

var getPipelineCmd = &cobra.Command{
	Use:     "pipeline",
	Aliases: []string{"pl"},
	Short:   "Show status and steps of a pipeline (latest of the current branch by default)",
	RunE:    runGetPipeline,
}

func init() {
	getCmd.AddCommand(getPipelineCmd)
	flags := getPipelineCmd.Flags()
	flags.Int32VarP(&getPipelineOpts.buildNumber, "build-number", "n", 0, "Build number of the pipeline (default: latest pipeline of the current branch)")
	flags.StringVarP(&getPipelineOpts.format, "format", "f", pipelineFormatText, "Output format (text or json)")
	flags.BoolVarP(&getPipelineOpts.wait, "wait", "w", false, "Wait until the pipeline completes; exit with an error unless it is successful")
	flags.IntVar(&getPipelineOpts.interval, "interval", 10, "Polling interval in seconds when waiting")
}

func runGetPipeline(_ *cobra.Command, _ []string) error {
	if err := getPipelineOpts.validate(); err != nil {
		return err
	}

	repo, err := getRepositoryInfoFromCurrentPath()
	if err != nil {
		return err
	}
	client, err := newPipelineClient()
	if err != nil {
		return err
	}

	pipeline, err := findPipeline(client, repo, getPipelineOpts.buildNumber)
	if err != nil {
		return err
	}

	if getPipelineOpts.wait {
		lastStatus := ""
		for !pipeline.IsCompleted() {
			if status := pipeline.Status(); status != lastStatus && getPipelineOpts.format == pipelineFormatText {
				fmt.Printf("Pipeline #%d is %s; waiting...\n", pipeline.BuildNumber, colorStatus(status))
				lastStatus = status
			}
			time.Sleep(time.Duration(getPipelineOpts.interval) * time.Second)
			pipeline, err = client.Get(repo.Org, repo.Name, pipeline.UUID)
			if err != nil {
				return fmt.Errorf("failed to get pipeline: %w", err)
			}
		}
	}

	steps, err := client.Steps(repo.Org, repo.Name, pipeline.UUID)
	if err != nil {
		return fmt.Errorf("failed to get steps of pipeline #%d: %w", pipeline.BuildNumber, err)
	}

	if getPipelineOpts.format == pipelineFormatJSON {
		if steps == nil {
			steps = []pipelineapi.Step{}
		}
		output, err := json.MarshalIndent(pipelineJSON{Pipeline: *pipeline, Steps: steps}, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal pipeline: %w", err)
		}
		fmt.Println(string(output))
	} else {
		printPipeline(repo, pipeline, steps, time.Now())
	}

	if getPipelineOpts.wait && pipeline.Status() != pipelineapi.ResultSuccessful {
		return fmt.Errorf("pipeline #%d completed with result %s", pipeline.BuildNumber, pipeline.Status())
	}
	return nil
}

// findPipeline returns the pipeline with the build number, or the latest
// pipeline of the current branch when buildNumber is 0.
func findPipeline(client *pipelineapi.Client, repo *models.Repository, buildNumber int32) (*pipelineapi.Pipeline, error) {
	if buildNumber > 0 {
		p, err := client.Get(repo.Org, repo.Name, strconv.Itoa(int(buildNumber)))
		if err != nil {
			return nil, fmt.Errorf("failed to get pipeline #%d: %w", buildNumber, err)
		}
		return p, nil
	}

	branch, err := git.GetCurrentBranchName()
	if err != nil {
		return nil, fmt.Errorf("failed to get current branch: %w", err)
	}
	pipelines, err := client.List(repo.Org, repo.Name, pipelineapi.ListOptions{Branch: branch, Limit: 1})
	if err != nil {
		return nil, fmt.Errorf("failed to list pipelines: %w", err)
	}
	if len(pipelines) == 0 {
		return nil, fmt.Errorf("no pipelines found for branch %s", branch)
	}
	return &pipelines[0], nil
}

func printPipeline(repo *models.Repository, p *pipelineapi.Pipeline, steps []pipelineapi.Step, now time.Time) {
	fmt.Printf("Pipeline #%d %s\n", p.BuildNumber, colorStatus(p.Status()))
	fmt.Printf("  Target:    %s\n", formatTarget(p.Target))
	fmt.Printf("  Commit:    %s\n", shortHash(p.Target))
	fmt.Printf("  Trigger:   %s\n", pipelineTrigger(p))
	fmt.Printf("  Creator:   %s\n", pipelineCreator(p))
	if p.CreatedOn != nil {
		fmt.Printf("  Created:   %s\n", datetime.GetLocalDateTimeString(p.CreatedOn))
	}
	if p.CompletedOn != nil {
		fmt.Printf("  Completed: %s\n", datetime.GetLocalDateTimeString(p.CompletedOn))
	}
	fmt.Printf("  Duration:  %s\n", formatDuration(pipelineDuration(p, now)))
	fmt.Printf("  URL:       %s\n", pipelineWebURL(repo.Org, repo.Name, p.BuildNumber))

	if len(steps) == 0 {
		return
	}
	fmt.Println("Steps:")
	for i := range steps {
		s := &steps[i]
		name := s.Name
		if name == "" {
			name = "(unnamed)"
		}
		fmt.Printf("  %s %s %s\n", colorStatus(s.Status()), name, formatDuration(stepDuration(s, now)))
	}
}
