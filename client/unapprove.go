package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
)

// UnapproveRequest makes an API call to remove approval of the specified pull request
func (client *Client) UnapproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error {
	path := fmt.Sprintf("%s/%d/approve", getBasePath(repo), id)
	req := newRequest(cred, "DELETE", path)
	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		msg := getErrorResponseMessage(resp)
		return errors.New(msg)
	}
	return nil
}
