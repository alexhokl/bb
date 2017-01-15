package client

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alexhokl/go-bb-pr/models"
)

type APIClient interface {
	ListRequests(cred *models.UserCredential, repo *models.Repository) (*models.PullRequestList, error)
	GetRequest(cred *models.UserCredential, repo *models.Repository, id int) (*models.PullRequestDetail, error)
	ApproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	UnapproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	DeclineRequest(cred *models.UserCredential, repo *models.Repository, id int) error
}

type Client struct {
	client *http.Client
}

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
	req.SetBasicAuth(cred.Username, cred.Password)
	return req
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
