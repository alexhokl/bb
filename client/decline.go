package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/bb/models"
)

// DeclineRequest makes an API call to decline the specified pull request
func (client *Client) DeclineRequest(cred *models.UserCredential, repo *models.Repository, id int) error {
	path := fmt.Sprintf("%s/%d/decline", getBasePath(repo), id)
	req := newRequest(cred, "POST", path)
	resp, err := client.do(req)
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
