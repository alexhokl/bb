package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/alexhokl/go-bb-pr/models"
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

func newRequest(cred *models.UserCredential, verb string, path string) *http.Request {
	req, _ := http.NewRequest(verb, path, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cred.AccessToken))
	return req
}

func newPostRequest(cred *models.UserCredential, path string, data interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	jsonStr := string(buf.Bytes())
	replacedStr := strings.Replace(jsonStr, "'", "", -1)
	req, _ := http.NewRequest("POST", path, bytes.NewBufferString(replacedStr))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cred.AccessToken))
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func (client *Client) do(req *http.Request) (*http.Response, error) {
	res, err := client.client.Do(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode == http.StatusUnauthorized {
		return res, fmt.Errorf("Please run command login before continue on")
	}
	if res.StatusCode == http.StatusForbidden {
		return res, fmt.Errorf("You are not authorized to perform this action")
	}
	return res, err
}

func dumpResponse(resp *http.Response) error {
	_, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func getErrorResponseMessage(resp *http.Response) string {
	return fmt.Sprintf(
		"Failed response (status code: %d): %s", resp.StatusCode, resp.Status)
}

func parse(resp *http.Response) (*models.PullRequestDetail, error) {
	var jsonObj models.PullRequestDetail
	err := json.NewDecoder(resp.Body).Decode(&jsonObj)
	if err != nil {
		return nil, err
	}
	return &jsonObj, nil
}
