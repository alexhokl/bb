package command

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type updateProjectRemovePermissionOptions struct {
	workspace  string
	projectKey string
	groupSlug  string
}

var updateProjectRemovePermissionOpts updateProjectRemovePermissionOptions

func (opts *updateProjectRemovePermissionOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.projectKey == "" {
		return fmt.Errorf("project-key is required")
	}
	if opts.groupSlug == "" {
		return fmt.Errorf("group-slug is required")
	}
	return nil
}

var updateProjectRemovePermissionCmd = &cobra.Command{
	Use:   "remove-permission",
	Short: "Remove a user group permission from a project",
	RunE:  runUpdateProjectRemovePermission,
}

func init() {
	updateProjectCmd.AddCommand(updateProjectRemovePermissionCmd)

	flags := updateProjectRemovePermissionCmd.Flags()
	flags.StringVarP(&updateProjectRemovePermissionOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&updateProjectRemovePermissionOpts.projectKey, "project-key", "k", "", "Project key")
	flags.StringVarP(&updateProjectRemovePermissionOpts.groupSlug, "group-slug", "g", "", "User group slug")

	_ = updateProjectRemovePermissionCmd.MarkFlagRequired("workspace")
	_ = updateProjectRemovePermissionCmd.MarkFlagRequired("project-key")
	_ = updateProjectRemovePermissionCmd.MarkFlagRequired("group-slug")
}

func runUpdateProjectRemovePermission(_ *cobra.Command, _ []string) error {
	if err := updateProjectRemovePermissionOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/workspaces/%s/projects/%s/permissions-config/groups/%s",
		updateProjectRemovePermissionOpts.workspace,
		updateProjectRemovePermissionOpts.projectKey,
		updateProjectRemovePermissionOpts.groupSlug,
	)

	req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call permissions API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}
		return fmt.Errorf("permissions API returned status %d: %s", resp.StatusCode, string(respBody))
	}

	fmt.Printf("Permission removed: %s from project %s\n",
		updateProjectRemovePermissionOpts.groupSlug,
		updateProjectRemovePermissionOpts.projectKey,
	)
	return nil
}
