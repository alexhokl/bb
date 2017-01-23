package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
)

// ApproveRequest makes an API call to approve the specified pull request
func (client *Client) ApproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error {
	path := fmt.Sprintf("%s/%d/approve", getBasePath(repo), id)
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
