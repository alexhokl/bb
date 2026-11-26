package command

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/alexhokl/helper/datetime"
	"github.com/spf13/cobra"

	"github.com/alexhokl/bb/internal/pipelineapi"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-pipelines-get

const maxPipelineListLimit = 100

type listPipelinesOptions struct {
	branch string
	status string
	limit  int
	format string
}

var listPipelinesOpts listPipelinesOptions

func (opts *listPipelinesOptions) validate() error {
	if opts.limit < 1 || opts.limit > maxPipelineListLimit {
		return fmt.Errorf("limit must be between 1 and %d", maxPipelineListLimit)
	}
	return validatePipelineFormat(opts.format)
}

var listPipelinesCmd = &cobra.Command{
	Use:     "pipelines",
	Aliases: []string{"pipeline", "pl"},
	Short:   "List recent pipelines of the current repository",
	RunE:    runListPipelines,
}

func init() {
	listCmd.AddCommand(listPipelinesCmd)
	flags := listPipelinesCmd.Flags()
	flags.StringVarP(&listPipelinesOpts.branch, "branch", "b", "", "Only list pipelines of this branch")
	flags.StringVarP(&listPipelinesOpts.status, "status", "s", "", "Only list pipelines with this status (e.g. SUCCESSFUL, FAILED, STOPPED, IN_PROGRESS, PENDING)")
	flags.IntVarP(&listPipelinesOpts.limit, "limit", "n", 10, fmt.Sprintf("Number of pipelines to list (max %d)", maxPipelineListLimit))
	flags.StringVarP(&listPipelinesOpts.format, "format", "f", pipelineFormatText, "Output format (text or json)")
}

func runListPipelines(_ *cobra.Command, _ []string) error {
	if err := listPipelinesOpts.validate(); err != nil {
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

	pipelines, err := client.List(repo.Org, repo.Name, pipelineapi.ListOptions{
		Branch: listPipelinesOpts.branch,
		Status: strings.ToUpper(listPipelinesOpts.status),
		Limit:  listPipelinesOpts.limit,
	})
	if err != nil {
		return fmt.Errorf("failed to list pipelines: %w", err)
	}

	if listPipelinesOpts.format == pipelineFormatJSON {
		if pipelines == nil {
			pipelines = []pipelineapi.Pipeline{}
		}
		output, err := json.MarshalIndent(pipelines, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal pipelines: %w", err)
		}
		fmt.Println(string(output))
		return nil
	}

	if len(pipelines) == 0 {
		fmt.Println("There are no pipelines.")
		return nil
	}

	now := time.Now()
	for i := range pipelines {
		fmt.Println(formatPipelineLine(&pipelines[i], now))
	}
	return nil
}

// formatPipelineLine returns a one-line summary of a pipeline.
func formatPipelineLine(p *pipelineapi.Pipeline, now time.Time) string {
	created := "-"
	if p.CreatedOn != nil {
		created = datetime.GetLocalDateTimeString(p.CreatedOn)
	}
	return fmt.Sprintf(
		"#%d %s %s %s %s %s %s %s",
		p.BuildNumber,
		colorStatus(p.Status()),
		formatTarget(p.Target),
		shortHash(p.Target),
		pipelineTrigger(p),
		pipelineCreator(p),
		created,
		formatDuration(pipelineDuration(p, now)),
	)
}
