package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/helper/httphelper"
)

// APIClient interface
type APIClient interface {
	ListRequests(cred *models.UserCredential, repo *models.Repository) ([]models.PullRequestInfo, error)
	GetRequest(cred *models.UserCredential, repo *models.Repository, id int) (*models.PullRequestDetail, error)
	ApproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	UnapproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	DeclineRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	MergeRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	ActivityRequest(cred *models.UserCredential, repo *models.Repository, id int) ([]models.PullRequestActivity, error)
	CreateRequest(cred *models.UserCredential, repo *models.Repository, req *models.PullRequestCreateRequest) (*models.PullRequestDetail, error)
	AddComment(cred *models.UserCredential, repo *models.Repository, id int, comment string) error
	ListCommits(cred *models.UserCredential, repo *models.Repository, pullRequestID int) ([]models.CommitInfo, error)
	AddJiraLabels(cred *models.UserCredential, repo *models.Repository, jiraID string, labels ...string) error
}

// Client struct
type Client struct {
	client *http.Client
}

// NewClient creates a new Client instance
func NewClient() *Client {
	client := &http.Client{}

	return &Client{
		client: client,
	}
}

func getBasePath(repo *models.Repository) string {
	return fmt.Sprintf(
		"https://bitbucket.org/api/2.0/repositories/%s/%s/pullrequests",
		repo.Org,
		repo.Name)
}

func getJiraAPIPath(repo *models.Repository, endpoint string) string {
	return fmt.Sprintf(
		"https://%s.atlassian.net/rest/api/3/%s",
		repo.Org,
		endpoint,
	)
}

func newRequest(cred *models.UserCredential, verb string, path string) *http.Request {
	req, _ := http.NewRequest(verb, path, nil)
	httphelper.SetBearerTokenHeader(req, cred.AccessToken)
	return req
}

func newPostRequest(cred *models.UserCredential, path string, data interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	jsonStr := buf.String()
	replacedStr := strings.ReplaceAll(jsonStr, "'", "")
	req, _ := http.NewRequest("POST", path, bytes.NewBufferString(replacedStr))
	httphelper.SetBearerTokenHeader(req, cred.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func newJiraPostRequest(cred *models.UserCredential, path string, data interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	jsonStr := buf.String()
	replacedStr := strings.ReplaceAll(jsonStr, "'", "")
	req, _ := http.NewRequest("POST", path, bytes.NewBufferString(replacedStr))
	req.SetBasicAuth(cred.JiraEmailAddress, cred.JiraAPIKey)
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func newJiraPutRequest(cred *models.UserCredential, path string, data interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	jsonStr := buf.String()
	replacedStr := strings.ReplaceAll(jsonStr, "'", "")
	req, _ := http.NewRequest("PUT", path, bytes.NewBufferString(replacedStr))
	req.SetBasicAuth(cred.JiraEmailAddress, cred.JiraAPIKey)
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func (client *Client) do(req *http.Request) (*http.Response, error) {
	res, err := client.client.Do(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode == http.StatusUnauthorized {
		return res, fmt.Errorf("please run command login before continue on")
	}
	if res.StatusCode == http.StatusForbidden {
		return res, fmt.Errorf("you are not authorized to perform this action")
	}
	return res, err
}

func getErrorResponseMessage(resp *http.Response) string {
	return fmt.Sprintf(
		"Failed response (status code: %d): %s", resp.StatusCode, resp.Status)
}
