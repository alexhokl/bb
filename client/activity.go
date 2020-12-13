package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
	"github.com/alexhokl/helper/jsonhelper"
)

type pullRequestActivityListResponse struct {
	Items []models.PullRequestActivity `json:"values"`
	Next  string                       `json:"next"`
}

// ActivityRequest makes API call(s) to retrieve activities of the specified pull request
func (client *Client) ActivityRequest(cred *models.UserCredential, repo *models.Repository, id int) ([]models.PullRequestActivity, error) {
	var list []models.PullRequestActivity

	path := fmt.Sprintf("%s/%d/activity", getBasePath(repo), id)

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

		var listResponse pullRequestActivityListResponse
		errParse := jsonhelper.ParseJSONReader(resp.Body, &listResponse)
		if errParse != nil {
			return nil, errParse
		}
		if list == nil {
			list = listResponse.Items
		} else {
			for _, e := range listResponse.Items {
				list = append(list, e)
			}
		}
		path = listResponse.Next
	}

	// getting around a bug of the API
	updatedList := list[:len(list)-1]

	return updatedList, nil
}
