package command

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get
// Note: workspace user groups are managed via the legacy 1.0 API which is not
// included in the swagger client.

type listUserGroupOptions struct {
	workspace string
}

var listUserGroupOpts listUserGroupOptions

func (opts *listUserGroupOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	return nil
}

var listUserGroupCmd = &cobra.Command{
	Use:     "user-groups",
	Aliases: []string{"user-group"},
	Short:   "List user groups of a workspace",
	RunE:    runListUserGroups,
}

func init() {
	listCmd.AddCommand(listUserGroupCmd)
	flags := listUserGroupCmd.Flags()
	flags.StringVarP(&listUserGroupOpts.workspace, "workspace", "w", "", "Workspace slug")
	_ = listUserGroupCmd.MarkFlagRequired("workspace")
}

func runListUserGroups(_ *cobra.Command, _ []string) error {
	if err := listUserGroupOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.bitbucket.org/1.0/groups/%s/", listUserGroupOpts.workspace)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call groups API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("groups API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var groups []models.UserGroup
	if err := json.Unmarshal(body, &groups); err != nil {
		return fmt.Errorf("failed to parse groups response: %w", err)
	}

	if len(groups) == 0 {
		fmt.Println("There are no user groups in this workspace.")
		return nil
	}

	for i := range groups {
		models.PrintUserGroup(&groups[i])
	}
	return nil
}
