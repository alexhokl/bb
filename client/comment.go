package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/bb/models"
)

// AddComment makes an API call to add a comment to the specified pull request
func (client *Client) AddComment(cred *models.UserCredential, repo *models.Repository, id int, comment string) error {
	path := fmt.Sprintf("%s/%d/comments", getBasePath(repo), id)

	data := &models.CommentRequest{
		Content: models.RawContent{
			Raw: comment,
		},
	}

	req, errReq := newPostRequest(cred, path, data)
	if errReq != nil {
		return errReq
	}
	resp, err := client.do(req)
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
