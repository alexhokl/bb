package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexhokl/go-bb-pr/models"
)

func (client *Client) ListRequests(cred *models.UserCredential, repo *models.Repository) (*models.PullRequestList, error) {
	path := getBasePath(repo)
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
	return parseList(resp)
}

func parseList(resp *http.Response) (*models.PullRequestList, error) {
	var jsonObj models.PullRequestList
	err := json.NewDecoder(resp.Body).Decode(&jsonObj)
	if err != nil {
		return nil, err
	}
	return &jsonObj, nil
}
