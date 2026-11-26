// Package pipelineapi provides a minimal hand-written client for the
// Bitbucket Cloud Pipelines endpoints. The auto-generated swagger client
// cannot be used for these because it does not decode the polymorphic
// state and target objects (only their "type" survives) and it does not
// accept the query parameters needed for sorting and filtering.
package pipelineapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// DefaultBaseURL is the base URL of the Bitbucket Cloud API.
const DefaultBaseURL = "https://api.bitbucket.org/2.0"

// State names of a pipeline or a step.
const (
	StateCompleted  = "COMPLETED"
	StateInProgress = "IN_PROGRESS"
	StatePending    = "PENDING"

	ResultSuccessful = "SUCCESSFUL"
)

// Named is an object identified by its name, such as a result or a stage.
type Named struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

// State is the state of a pipeline or a step. Result is set when the state
// is COMPLETED and Stage may be set when it is IN_PROGRESS.
type State struct {
	Type   string `json:"type,omitempty"`
	Name   string `json:"name,omitempty"`
	Result *Named `json:"result,omitempty"`
	Stage  *Named `json:"stage,omitempty"`
}

// Status returns the most specific status of the state: the result of a
// completed state, the stage of an in-progress state, or the state name.
func (s *State) Status() string {
	if s == nil {
		return ""
	}
	if s.Result != nil && s.Result.Name != "" {
		return s.Result.Name
	}
	if s.Stage != nil && s.Stage.Name != "" {
		return s.Stage.Name
	}
	return s.Name
}

// Commit is a commit referred to by a pipeline target.
type Commit struct {
	Hash string `json:"hash,omitempty"`
}

// Selector is the part of bitbucket-pipelines.yml that was run.
type Selector struct {
	Type    string `json:"type,omitempty"`
	Pattern string `json:"pattern,omitempty"`
}

// PullRequest is a pull request referred to by a pipeline target.
type PullRequest struct {
	ID int32 `json:"id,omitempty"`
}

// Target is what a pipeline built: a ref (branch or tag), a commit (custom
// pipeline) or a pull request.
type Target struct {
	Type        string       `json:"type,omitempty"`
	RefType     string       `json:"ref_type,omitempty"`
	RefName     string       `json:"ref_name,omitempty"`
	Source      string       `json:"source,omitempty"`
	Destination string       `json:"destination,omitempty"`
	Commit      *Commit      `json:"commit,omitempty"`
	Selector    *Selector    `json:"selector,omitempty"`
	PullRequest *PullRequest `json:"pullrequest,omitempty"`
}

// Account is the Bitbucket account that created a pipeline.
type Account struct {
	DisplayName string `json:"display_name,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

// Pipeline is a run of Bitbucket Pipelines.
type Pipeline struct {
	UUID              string     `json:"uuid,omitempty"`
	BuildNumber       int32      `json:"build_number,omitempty"`
	Creator           *Account   `json:"creator,omitempty"`
	Target            *Target    `json:"target,omitempty"`
	Trigger           *Named     `json:"trigger,omitempty"`
	State             *State     `json:"state,omitempty"`
	CreatedOn         *time.Time `json:"created_on,omitempty"`
	CompletedOn       *time.Time `json:"completed_on,omitempty"`
	DurationInSeconds int32      `json:"duration_in_seconds,omitempty"`
	BuildSecondsUsed  int32      `json:"build_seconds_used,omitempty"`
}

// Status returns the most specific status of the pipeline.
func (p *Pipeline) Status() string {
	return p.State.Status()
}

// IsCompleted reports whether the pipeline has finished.
func (p *Pipeline) IsCompleted() bool {
	return p.State != nil && p.State.Name == StateCompleted
}

// Step is a step of a pipeline.
type Step struct {
	UUID              string     `json:"uuid,omitempty"`
	Name              string     `json:"name,omitempty"`
	State             *State     `json:"state,omitempty"`
	StartedOn         *time.Time `json:"started_on,omitempty"`
	CompletedOn       *time.Time `json:"completed_on,omitempty"`
	DurationInSeconds int32      `json:"duration_in_seconds,omitempty"`
}

// Status returns the most specific status of the step.
func (s *Step) Status() string {
	return s.State.Status()
}

type paginated[T any] struct {
	Values []T    `json:"values,omitempty"`
	Next   string `json:"next,omitempty"`
}

// ListOptions filters the pipelines returned by List.
type ListOptions struct {
	// Branch limits the result to pipelines of this branch.
	Branch string
	// Status limits the result to pipelines with this status.
	Status string
	// Limit is the maximum number of pipelines returned (1 to 100).
	Limit int
}

// Client calls the Bitbucket Pipelines API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
}

// New returns a Client that authenticates with the given OAuth access token
// against baseURL (normally DefaultBaseURL).
func New(baseURL, token string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    baseURL,
		token:      token,
	}
}

func (c *Client) repositoryURL(workspace, repo string) string {
	return fmt.Sprintf("%s/repositories/%s/%s/pipelines", c.baseURL, url.PathEscape(workspace), url.PathEscape(repo))
}

// List returns the most recent pipelines of a repository, newest first.
func (c *Client) List(workspace, repo string, opts ListOptions) ([]Pipeline, error) {
	query := url.Values{}
	query.Set("sort", "-created_on")
	if opts.Limit > 0 {
		query.Set("pagelen", strconv.Itoa(opts.Limit))
	}
	if opts.Branch != "" {
		query.Set("target.branch", opts.Branch)
	}
	if opts.Status != "" {
		query.Set("status", opts.Status)
	}

	var page paginated[Pipeline]
	if err := c.get(c.repositoryURL(workspace, repo)+"/?"+query.Encode(), &page); err != nil {
		return nil, err
	}
	if opts.Limit > 0 && len(page.Values) > opts.Limit {
		return page.Values[:opts.Limit], nil
	}
	return page.Values, nil
}

// Get returns a pipeline identified by its UUID (with curly braces) or its
// build number.
func (c *Client) Get(workspace, repo, idOrBuildNumber string) (*Pipeline, error) {
	var p Pipeline
	if err := c.get(c.repositoryURL(workspace, repo)+"/"+url.PathEscape(idOrBuildNumber), &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Steps returns all steps of a pipeline, following pagination.
func (c *Client) Steps(workspace, repo, pipelineUUID string) ([]Step, error) {
	next := c.repositoryURL(workspace, repo) + "/" + url.PathEscape(pipelineUUID) + "/steps/?pagelen=100"
	var steps []Step
	for next != "" {
		var page paginated[Step]
		if err := c.get(next, &page); err != nil {
			return nil, err
		}
		steps = append(steps, page.Values...)
		next = page.Next
	}
	return steps, nil
}

func (c *Client) get(requestURL string, out any) error {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("not found (status %s)", resp.Status)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	return nil
}
