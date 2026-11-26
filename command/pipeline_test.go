package command

import (
	"strings"
	"testing"
	"time"

	"github.com/alexhokl/bb/internal/pipelineapi"
)

func TestListPipelinesOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    listPipelinesOptions
		wantErr bool
	}{
		{"valid", listPipelinesOptions{limit: 10, format: "text"}, false},
		{"json", listPipelinesOptions{limit: 100, format: "json"}, false},
		{"zero limit", listPipelinesOptions{limit: 0, format: "text"}, true},
		{"limit too large", listPipelinesOptions{limit: 101, format: "text"}, true},
		{"bad format", listPipelinesOptions{limit: 10, format: "xml"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.opts.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPipelineOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    getPipelineOptions
		wantErr bool
	}{
		{"latest", getPipelineOptions{format: "text", interval: 10}, false},
		{"build number", getPipelineOptions{buildNumber: 3, format: "json", interval: 10}, false},
		{"negative build number", getPipelineOptions{buildNumber: -1, format: "text", interval: 10}, true},
		{"bad format", getPipelineOptions{format: "yaml", interval: 10}, true},
		{"wait with zero interval", getPipelineOptions{format: "text", wait: true, interval: 0}, true},
		{"zero interval without wait", getPipelineOptions{format: "text", interval: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.opts.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFormatTarget(t *testing.T) {
	tests := []struct {
		name   string
		target *pipelineapi.Target
		want   string
	}{
		{"nil", nil, "-"},
		{"branch", &pipelineapi.Target{RefType: "branch", RefName: "master"}, "master"},
		{"tag", &pipelineapi.Target{RefType: "tag", RefName: "v1.0.0"}, "tag v1.0.0"},
		{"pull request", &pipelineapi.Target{Source: "feature", Destination: "master", PullRequest: &pipelineapi.PullRequest{ID: 7}}, "PR #7 feature -> master"},
		{"custom", &pipelineapi.Target{RefName: "master", Selector: &pipelineapi.Selector{Type: "custom", Pattern: "deploy"}}, "master (custom: deploy)"},
		{"commit only", &pipelineapi.Target{Commit: &pipelineapi.Commit{Hash: "abc"}}, "-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTarget(tt.target); got != tt.want {
				t.Errorf("formatTarget() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestShortHash(t *testing.T) {
	if got := shortHash(nil); got != "-" {
		t.Errorf("shortHash(nil) = %q", got)
	}
	if got := shortHash(&pipelineapi.Target{Commit: &pipelineapi.Commit{Hash: "0123456789"}}); got != "0123456" {
		t.Errorf("shortHash() = %q", got)
	}
	if got := shortHash(&pipelineapi.Target{Commit: &pipelineapi.Commit{Hash: "abc"}}); got != "abc" {
		t.Errorf("shortHash() = %q", got)
	}
}

func TestFormatDuration(t *testing.T) {
	tests := map[int32]string{0: "-", -1: "-", 6: "6s", 245: "4m5s", 3723: "1h2m3s"}
	for in, want := range tests {
		if got := formatDuration(in); got != want {
			t.Errorf("formatDuration(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestPipelineDuration(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 10, 0, 0, time.UTC)
	created := now.Add(-90 * time.Second)

	completed := &pipelineapi.Pipeline{DurationInSeconds: 300, State: &pipelineapi.State{Name: "COMPLETED"}}
	if got := pipelineDuration(completed, now); got != 300 {
		t.Errorf("completed duration = %d", got)
	}
	running := &pipelineapi.Pipeline{CreatedOn: &created, State: &pipelineapi.State{Name: "IN_PROGRESS"}}
	if got := pipelineDuration(running, now); got != 90 {
		t.Errorf("running duration = %d", got)
	}
	step := &pipelineapi.Step{StartedOn: &created}
	if got := stepDuration(step, now); got != 90 {
		t.Errorf("running step duration = %d", got)
	}
	if got := stepDuration(&pipelineapi.Step{}, now); got != 0 {
		t.Errorf("pending step duration = %d", got)
	}
}

func TestFormatPipelineLine(t *testing.T) {
	p := &pipelineapi.Pipeline{
		BuildNumber:       42,
		Creator:           &pipelineapi.Account{DisplayName: "Alex"},
		Target:            &pipelineapi.Target{RefType: "branch", RefName: "master", Commit: &pipelineapi.Commit{Hash: "0123456789"}},
		Trigger:           &pipelineapi.Named{Name: "PUSH"},
		State:             &pipelineapi.State{Name: "COMPLETED", Result: &pipelineapi.Named{Name: "SUCCESSFUL"}},
		DurationInSeconds: 65,
	}
	got := formatPipelineLine(p, time.Now())
	for _, want := range []string{"#42", "SUCCESSFUL", "master", "0123456", "PUSH", "Alex", "1m5s"} {
		if !strings.Contains(got, want) {
			t.Errorf("formatPipelineLine() = %q, missing %q", got, want)
		}
	}
}
