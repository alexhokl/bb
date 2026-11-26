package pipelineapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const pipelinesPage = `{
  "values": [
    {
      "uuid": "{p1}",
      "build_number": 42,
      "creator": {"display_name": "Alex"},
      "target": {
        "type": "pipeline_ref_target",
        "ref_type": "branch",
        "ref_name": "master",
        "commit": {"type": "commit", "hash": "0123456789abcdef"},
        "selector": {"type": "branches", "pattern": "master"}
      },
      "trigger": {"type": "pipeline_trigger_push", "name": "PUSH"},
      "state": {"type": "pipeline_state_completed", "name": "COMPLETED",
                "result": {"type": "pipeline_state_completed_failed", "name": "FAILED"}},
      "created_on": "2026-11-20T10:00:00.000000Z",
      "completed_on": "2026-11-20T10:05:00.000000Z",
      "duration_in_seconds": 300
    },
    {
      "uuid": "{p2}",
      "build_number": 43,
      "target": {
        "type": "pipeline_pullrequest_target",
        "source": "feature",
        "destination": "master",
        "pullrequest": {"id": 7}
      },
      "state": {"type": "pipeline_state_in_progress", "name": "IN_PROGRESS",
                "stage": {"type": "pipeline_state_in_progress_running", "name": "RUNNING"}}
    }
  ]
}`

func TestList(t *testing.T) {
	var gotPath, gotQuery, gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
		gotAuth = r.Header.Get("Authorization")
		_, _ = fmt.Fprint(w, pipelinesPage)
	}))
	defer server.Close()

	client := New(server.URL, "token")
	pipelines, err := client.List("ws", "repo", ListOptions{Branch: "master", Status: "FAILED", Limit: 5})
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if gotPath != "/repositories/ws/repo/pipelines/" {
		t.Errorf("path = %q", gotPath)
	}
	for _, want := range []string{"sort=-created_on", "target.branch=master", "status=FAILED", "pagelen=5"} {
		if !strings.Contains(gotQuery, want) {
			t.Errorf("query %q does not contain %q", gotQuery, want)
		}
	}
	if gotAuth != "Bearer token" {
		t.Errorf("Authorization = %q", gotAuth)
	}

	if len(pipelines) != 2 {
		t.Fatalf("len(pipelines) = %d, want 2", len(pipelines))
	}
	p := pipelines[0]
	if p.BuildNumber != 42 || p.Status() != "FAILED" || !p.IsCompleted() {
		t.Errorf("pipeline 0 = %+v, status %s", p, p.Status())
	}
	if p.Target.RefName != "master" || p.Target.Commit.Hash != "0123456789abcdef" {
		t.Errorf("target = %+v", p.Target)
	}
	if p.CreatedOn == nil || p.CompletedOn == nil || p.DurationInSeconds != 300 {
		t.Errorf("times not decoded: %+v", p)
	}
	p = pipelines[1]
	if p.Status() != "RUNNING" || p.IsCompleted() {
		t.Errorf("pipeline 1 status = %s", p.Status())
	}
	if p.Target.Source != "feature" || p.Target.PullRequest.ID != 7 {
		t.Errorf("PR target = %+v", p.Target)
	}
}

func TestList_TruncatesToLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, pipelinesPage)
	}))
	defer server.Close()

	pipelines, err := New(server.URL, "t").List("ws", "repo", ListOptions{Limit: 1})
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(pipelines) != 1 {
		t.Errorf("len(pipelines) = %d, want 1", len(pipelines))
	}
}

func TestGet(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		_, _ = fmt.Fprint(w, `{"uuid":"{p1}","build_number":42,"state":{"name":"PENDING"}}`)
	}))
	defer server.Close()

	p, err := New(server.URL, "t").Get("ws", "repo", "42")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if gotPath != "/repositories/ws/repo/pipelines/42" {
		t.Errorf("path = %q", gotPath)
	}
	if p.UUID != "{p1}" || p.Status() != "PENDING" {
		t.Errorf("pipeline = %+v", p)
	}
}

func TestGet_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	if _, err := New(server.URL, "t").Get("ws", "repo", "1"); err == nil {
		t.Error("Get() error = nil, want error")
	}
}

func TestSteps_FollowsPagination(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") == "2" {
			_, _ = fmt.Fprint(w, `{"values":[{"uuid":"{s2}","name":"Deploy","state":{"name":"PENDING"}}]}`)
			return
		}
		_, _ = fmt.Fprintf(w, `{"values":[{"uuid":"{s1}","name":"Build","state":{"name":"COMPLETED","result":{"name":"SUCCESSFUL"}},"duration_in_seconds":65}],"next":"%s/repositories/ws/repo/pipelines/p1/steps/?page=2"}`, server.URL)
	}))
	defer server.Close()

	steps, err := New(server.URL, "t").Steps("ws", "repo", "p1")
	if err != nil {
		t.Fatalf("Steps() error = %v", err)
	}
	if len(steps) != 2 {
		t.Fatalf("len(steps) = %d, want 2", len(steps))
	}
	if steps[0].Name != "Build" || steps[0].Status() != "SUCCESSFUL" || steps[0].DurationInSeconds != 65 {
		t.Errorf("step 0 = %+v", steps[0])
	}
	if steps[1].Status() != "PENDING" {
		t.Errorf("step 1 status = %s", steps[1].Status())
	}
}

func TestStateStatus(t *testing.T) {
	tests := []struct {
		name  string
		state *State
		want  string
	}{
		{"nil", nil, ""},
		{"pending", &State{Name: "PENDING"}, "PENDING"},
		{"running", &State{Name: "IN_PROGRESS", Stage: &Named{Name: "RUNNING"}}, "RUNNING"},
		{"in progress without stage", &State{Name: "IN_PROGRESS"}, "IN_PROGRESS"},
		{"completed", &State{Name: "COMPLETED", Result: &Named{Name: "STOPPED"}}, "STOPPED"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.state.Status(); got != tt.want {
				t.Errorf("Status() = %q, want %q", got, tt.want)
			}
		})
	}
}
