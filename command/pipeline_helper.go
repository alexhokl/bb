package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/alexhokl/helper/authhelper"
	"github.com/fatih/color"

	"github.com/alexhokl/bb/internal/pipelineapi"
)

const (
	pipelineFormatText = "text"
	pipelineFormatJSON = "json"
)

var validPipelineFormats = []string{pipelineFormatText, pipelineFormatJSON}

func validatePipelineFormat(format string) error {
	if format != pipelineFormatText && format != pipelineFormatJSON {
		return fmt.Errorf("format must be one of: %s", strings.Join(validPipelineFormats, ", "))
	}
	return nil
}

// newPipelineClient returns a Pipelines API client authenticated with the
// saved token.
func newPipelineClient() (*pipelineapi.Client, error) {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return nil, err
	}
	return pipelineapi.New(pipelineapi.DefaultBaseURL, savedToken.AccessToken), nil
}

// colorStatus returns the status coloured by its meaning.
func colorStatus(status string) string {
	var c *color.Color
	switch status {
	case "SUCCESSFUL":
		c = color.New(color.FgGreen)
	case "FAILED", "ERROR", "EXPIRED":
		c = color.New(color.FgRed)
	case "STOPPED", "NOT_RUN":
		c = color.New(color.FgHiBlack)
	case "":
		return "UNKNOWN"
	default:
		c = color.New(color.FgYellow)
	}
	return c.Sprint(status)
}

// formatDuration formats a number of seconds as e.g. "1h2m3s", "4m5s" or "6s".
func formatDuration(seconds int32) string {
	if seconds <= 0 {
		return "-"
	}
	return (time.Duration(seconds) * time.Second).String()
}

// pipelineDuration returns the duration of a pipeline; for a pipeline still
// running, it is the time elapsed since it was created.
func pipelineDuration(p *pipelineapi.Pipeline, now time.Time) int32 {
	if p.DurationInSeconds > 0 || p.IsCompleted() || p.CreatedOn == nil {
		return p.DurationInSeconds
	}
	return int32(now.Sub(*p.CreatedOn).Seconds())
}

// stepDuration returns the duration of a step; for a step still running, it
// is the time elapsed since it started.
func stepDuration(s *pipelineapi.Step, now time.Time) int32 {
	if s.DurationInSeconds > 0 || s.StartedOn == nil || s.CompletedOn != nil {
		return s.DurationInSeconds
	}
	return int32(now.Sub(*s.StartedOn).Seconds())
}

// shortHash returns the first 7 characters of a commit hash.
func shortHash(t *pipelineapi.Target) string {
	if t == nil || t.Commit == nil || t.Commit.Hash == "" {
		return "-"
	}
	if len(t.Commit.Hash) > 7 {
		return t.Commit.Hash[:7]
	}
	return t.Commit.Hash
}

// formatTarget describes what a pipeline built.
func formatTarget(t *pipelineapi.Target) string {
	if t == nil {
		return "-"
	}
	var s string
	switch {
	case t.Source != "" || t.Destination != "":
		s = fmt.Sprintf("%s -> %s", t.Source, t.Destination)
		if t.PullRequest != nil && t.PullRequest.ID > 0 {
			s = fmt.Sprintf("PR #%d %s", t.PullRequest.ID, s)
		}
	case t.RefName != "":
		s = t.RefName
		if t.RefType != "" && t.RefType != "branch" {
			s = fmt.Sprintf("%s %s", t.RefType, t.RefName)
		}
	default:
		s = "-"
	}
	if t.Selector != nil && t.Selector.Type == "custom" && t.Selector.Pattern != "" {
		s = fmt.Sprintf("%s (custom: %s)", s, t.Selector.Pattern)
	}
	return s
}

func pipelineTrigger(p *pipelineapi.Pipeline) string {
	if p.Trigger == nil || p.Trigger.Name == "" {
		return "-"
	}
	return p.Trigger.Name
}

func pipelineCreator(p *pipelineapi.Pipeline) string {
	if p.Creator == nil || p.Creator.DisplayName == "" {
		return "-"
	}
	return p.Creator.DisplayName
}

func pipelineWebURL(workspace, repo string, buildNumber int32) string {
	return fmt.Sprintf("https://bitbucket.org/%s/%s/pipelines/results/%d", workspace, repo, buildNumber)
}
