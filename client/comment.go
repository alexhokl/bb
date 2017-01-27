package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
)

// AddComment makes an API call to add a comment to the specified pull request
func (client *Client) AddComment(cred *models.UserCredential, repo *models.Repository, id int, comment string) error {
	path := fmt.Sprintf("%s/%d/comments", getVersion1BasePath(repo), id)

	data := map[string]string{
		"content": comment,
	}
	req, errReq := newPostURLDataRequest(cred, path, data)
	if errReq != nil {
		return errReq
	}
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
