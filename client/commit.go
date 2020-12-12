package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alexhokl/go-bb-pr/models"
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
		listResponse, errParse := parseCommitList(resp)
		if errParse != nil {
			return nil, errParse
		}
		if list == nil {
			list = listResponse.Items
		} else {
			for _, r := range listResponse.Items {
				list = append(list, r)
			}
		}
		path = listResponse.Next
	}

	return list, nil
}

func parseCommitList(resp *http.Response) (*commitListResponse, error) {
	var jsonObj commitListResponse
	err := json.NewDecoder(resp.Body).Decode(&jsonObj)
	if err != nil {
		return nil, err
	}
	return &jsonObj, nil
}
