package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type updateProjectAddPermissionOptions struct {
	workspace  string
	projectKey string
	groupSlug  string
	permission string
}

var updateProjectAddPermissionOpts updateProjectAddPermissionOptions

var validProjectPermissions = []string{"read", "write", "create-repo", "admin"}

func (opts *updateProjectAddPermissionOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.projectKey == "" {
		return fmt.Errorf("project-key is required")
	}
	if opts.groupSlug == "" {
		return fmt.Errorf("group-slug is required")
	}
	for _, v := range validProjectPermissions {
		if opts.permission == v {
			return nil
		}
	}
	return fmt.Errorf("permission must be one of: read, write, create-repo, admin")
}

var updateProjectAddPermissionCmd = &cobra.Command{
	Use:   "add-permission",
	Short: "Grant a user group permission to a project",
	RunE:  runUpdateProjectAddPermission,
}

func init() {
	updateProjectCmd.AddCommand(updateProjectAddPermissionCmd)

	flags := updateProjectAddPermissionCmd.Flags()
	flags.StringVarP(&updateProjectAddPermissionOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&updateProjectAddPermissionOpts.projectKey, "project-key", "k", "", "Project key")
	flags.StringVarP(&updateProjectAddPermissionOpts.groupSlug, "group-slug", "g", "", "User group slug")
	flags.StringVarP(&updateProjectAddPermissionOpts.permission, "permission", "p", "", "Permission level (read, write, create-repo, admin)")

	_ = updateProjectAddPermissionCmd.MarkFlagRequired("workspace")
	_ = updateProjectAddPermissionCmd.MarkFlagRequired("project-key")
	_ = updateProjectAddPermissionCmd.MarkFlagRequired("group-slug")
	_ = updateProjectAddPermissionCmd.MarkFlagRequired("permission")
}

func runUpdateProjectAddPermission(_ *cobra.Command, _ []string) error {
	if err := updateProjectAddPermissionOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	reqBody, err := json.Marshal(map[string]string{
		"permission": updateProjectAddPermissionOpts.permission,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	apiURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/workspaces/%s/projects/%s/permissions-config/groups/%s",
		updateProjectAddPermissionOpts.workspace,
		updateProjectAddPermissionOpts.projectKey,
		updateProjectAddPermissionOpts.groupSlug,
	)

	req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call permissions API: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("permissions API returned status %d: %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		Permission string `json:"permission"`
		Group      struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"group"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	fmt.Printf("Permission granted: %s -> %s (%s)\n", result.Group.Name, result.Permission, result.Group.Slug)
	return nil
}
