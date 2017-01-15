package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexhokl/go-bb-pr/models"
)

type pullRequestListResponse struct {
	Items []models.PullRequestInfo `json:"values"`
	Next  string                   `json:"next"`
}

// ListRequests makes an API call to retrieve a list of pull requests
func (client *Client) ListRequests(cred *models.UserCredential, repo *models.Repository) ([]models.PullRequestInfo, error) {
	var list []models.PullRequestInfo
	path := getBasePath(repo)

	for path != "" {
		req := newRequest(cred, "GET", path)
		resp, err := client.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			msg := getErrorResponseMessage(resp)
			return nil, errors.New(msg)
		}
		listResponse, errParse := parseList(resp)
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

func parseList(resp *http.Response) (*pullRequestListResponse, error) {
	var jsonObj pullRequestListResponse
	err := json.NewDecoder(resp.Body).Decode(&jsonObj)
	if err != nil {
		return nil, err
	}
	return &jsonObj, nil
}
