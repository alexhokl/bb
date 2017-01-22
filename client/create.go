package client

import (
	"errors"

	"github.com/alexhokl/go-bb-pr/models"
)

// CreateRequest makes an API call to create a pull requests
func (client *Client) CreateRequest(cred *models.UserCredential, repo *models.Repository, pr *models.PullRequestCreateRequest) error {
	path := getBasePath(repo)

	req, errReq := newPostRequest(cred, path, pr)
	if errReq != nil {
		return errReq
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		msg := getErrorResponseMessage(resp)
		return errors.New(msg)
	}

	return nil
}
