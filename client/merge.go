package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
)

// MergeRequest makes API call to merge a pull request
func (client *Client) MergeRequest(cred *models.UserCredential, repo *models.Repository, id int) error {
	path := fmt.Sprintf("%s/%d/merge", getBasePath(repo), id)
	req := newRequest(cred, "POST", path)
	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		msg := getErrorResponseMessage(resp)
		return errors.New(msg)
	}
	return nil
}
