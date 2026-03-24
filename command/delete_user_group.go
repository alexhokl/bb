package command

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/
// Note: workspace user groups are managed via the legacy 1.0 API which is not
// included in the swagger client.

type deleteUserGroupOptions struct {
	workspace string
	slug      string
}

var deleteUserGroupOpts deleteUserGroupOptions

func (opts *deleteUserGroupOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.slug == "" {
		return fmt.Errorf("slug is required")
	}
	return nil
}

var deleteUserGroupCmd = &cobra.Command{
	Use:     "user-group",
	Aliases: []string{"user-groups"},
	Short:   "Delete a user group in a workspace",
	RunE:    runDeleteUserGroup,
}

func init() {
	deleteCmd.AddCommand(deleteUserGroupCmd)
	flags := deleteUserGroupCmd.Flags()
	flags.StringVarP(&deleteUserGroupOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&deleteUserGroupOpts.slug, "slug", "s", "", "Slug of the user group")
	_ = deleteUserGroupCmd.MarkFlagRequired("workspace")
	_ = deleteUserGroupCmd.MarkFlagRequired("slug")
}

func runDeleteUserGroup(_ *cobra.Command, _ []string) error {
	if err := deleteUserGroupOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf("https://api.bitbucket.org/1.0/groups/%s/%s/", deleteUserGroupOpts.workspace, deleteUserGroupOpts.slug)
	req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call groups API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("groups API returned status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println("User group deleted successfully")
	return nil
}
