package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/helper/jsonhelper"
)

type commitListResponse struct {
	Items []models.CommitInfo `json:"values"`
	Next  string              `json:"next"`
}

// ListCommits makes an API call to retrieve a list of commits of a pull request
func (client *Client) ListCommits(cred *models.UserCredential, repo *models.Repository, pullRequestID int) ([]models.CommitInfo, error) {
	var list []models.CommitInfo
	path := fmt.Sprintf("%s/%d/commits", getBasePath(repo), pullRequestID)

	for path != "" {
		req := newRequest(cred, "GET", path)
		resp, err := client.do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			msg := getErrorResponseMessage(resp)
			return nil, errors.New(msg)
		}
		var listResponse commitListResponse
		errParse := jsonhelper.ParseJSONReader(resp.Body, &listResponse)
		if errParse != nil {
			return nil, errParse
		}
		if list == nil {
			list = listResponse.Items
		} else {
			list = append(list, listResponse.Items...)
		}
		path = listResponse.Next
	}

	return list, nil
}
