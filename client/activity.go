package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alexhokl/go-bb-pr/models"
)

// ActivityRequest makes API call(s) to retrieve activities of a pull request
func (client *Client) ActivityRequest(cred *models.UserCredential, repo *models.Repository, id int) (*models.PullRequestActivityList, error) {
	var list *models.PullRequestActivityList

	path := fmt.Sprintf("%s/%d/activity", getBasePath(repo), id)

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
		var events *models.PullRequestActivityList
		events, errParse := parseActivities(resp)
		if errParse != nil {
			return nil, errParse
		}
		if list == nil {
			list = events
		} else {
			for _, e := range events.Items {
				list.Items = append(list.Items, e)
			}
		}
		path = events.Next
	}
	return list, nil
}

func parseActivities(resp *http.Response) (*models.PullRequestActivityList, error) {
	var jsonObj models.PullRequestActivityList
	err := json.NewDecoder(resp.Body).Decode(&jsonObj)
	if err != nil {
		return nil, err
	}
	return &jsonObj, nil
}
