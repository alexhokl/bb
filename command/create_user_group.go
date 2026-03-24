package command

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/
// Note: workspace user groups are managed via the legacy 1.0 API which is not
// included in the swagger client.

type createUserGroupOptions struct {
	workspace string
	name      string
}

var createUserGroupOpts createUserGroupOptions

func (opts *createUserGroupOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

var createUserGroupCmd = &cobra.Command{
	Use:     "user-group",
	Aliases: []string{"user-groups"},
	Short:   "Create a user group in a workspace",
	RunE:    runCreateUserGroup,
}

func init() {
	createCmd.AddCommand(createUserGroupCmd)
	flags := createUserGroupCmd.Flags()
	flags.StringVarP(&createUserGroupOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&createUserGroupOpts.name, "name", "n", "", "Name of the user group")
	_ = createUserGroupCmd.MarkFlagRequired("workspace")
	_ = createUserGroupCmd.MarkFlagRequired("name")
}

func runCreateUserGroup(_ *cobra.Command, _ []string) error {
	if err := createUserGroupOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	formData := url.Values{}
	formData.Set("name", createUserGroupOpts.name)

	apiURL := fmt.Sprintf("https://api.bitbucket.org/1.0/groups/%s/", createUserGroupOpts.workspace)
	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call groups API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("groups API returned status %d: %s", resp.StatusCode, string(body))
	}

	var group models.UserGroup
	if err := json.Unmarshal(body, &group); err != nil {
		return fmt.Errorf("failed to parse group response: %w", err)
	}

	fmt.Printf("User group created successfully\n")
	models.PrintUserGroup(&group)
	return nil
}
