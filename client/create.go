package client

import (
	"errors"

	"github.com/alexhokl/go-bb-pr/models"
)

// CreateRequest makes an API call to create the specified pull requests
func (client *Client) CreateRequest(cred *models.UserCredential, repo *models.Repository, pr *models.PullRequestCreateRequest) (*models.PullRequestDetail, error) {
	path := getBasePath(repo)

	req, errReq := newPostRequest(cred, path, pr)
	if errReq != nil {
		return nil, errReq
	}
	resp, err := client.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		msg := getErrorResponseMessage(resp)
		return nil, errors.New(msg)
	}

	return parse(resp)
}
