package client

import (
	"fmt"

	"github.com/alexhokl/bb/models"
)

// AddJiraLabels adds labels to the specified JIRA issues
func (client *Client) AddJiraLabels(cred *models.UserCredential, repo *models.Repository, jiraID string, labels ...string) error {
	if !cred.HasJiraCredentials() {
		return fmt.Errorf("JIRA credentials has not been configured")
	}
	path := getJiraAPIPath(repo, fmt.Sprintf("issue/%s", jiraID))

	var labelRequests []models.LabelReqeuest
	for _, s := range labels {
		labelRequests = append(labelRequests, models.LabelReqeuest{Add: s})
	}

	data := &models.IssueRequest{
		Update: models.UpdateIssue{
			Labels: labelRequests,
		},
	}
	req, err := newJiraPutRequest(cred, path, data)
	if err != nil {
		return err
	}
	resp, err := client.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		msg := getErrorResponseMessage(resp)
		return fmt.Errorf(msg)
	}
	return nil
}
