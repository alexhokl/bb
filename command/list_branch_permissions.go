package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type listBranchPermissionsOptions struct {
	workspace string
	slug      string
	kind      string
	pattern   string
}

var listBranchPermissionsOpts listBranchPermissionsOptions

func (opts *listBranchPermissionsOptions) validate() error {
	if opts.kind != "" && !isValidBranchPermissionKind(opts.kind) {
		return fmt.Errorf("kind must be one of: %v", validBranchPermissionKinds)
	}
	return nil
}

var listBranchPermissionsCmd = &cobra.Command{
	Use:   "branch-permissions",
	Short: "List branch permission (restriction) rules of a repository",
	RunE:  runListBranchPermissions,
}

func init() {
	listCmd.AddCommand(listBranchPermissionsCmd)
	flags := listBranchPermissionsCmd.Flags()
	flags.StringVarP(&listBranchPermissionsOpts.workspace, "workspace", "w", "", "Workspace slug (defaults to the current repository's workspace)")
	flags.StringVarP(&listBranchPermissionsOpts.slug, "slug", "s", "", "Repository slug (defaults to the current repository)")
	flags.StringVar(&listBranchPermissionsOpts.kind, "kind", "", "Filter branch restrictions of this kind")
	flags.StringVar(&listBranchPermissionsOpts.pattern, "pattern", "", "Filter branch restrictions applied to branches of this pattern")
}

func runListBranchPermissions(_ *cobra.Command, _ []string) error {
	if err := listBranchPermissionsOpts.validate(); err != nil {
		return err
	}

	workspace, slug, err := resolveWorkspaceAndSlug(listBranchPermissionsOpts.workspace, listBranchPermissionsOpts.slug)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/branch-restrictions",
		workspace,
		slug,
	)

	query := ""
	if listBranchPermissionsOpts.kind != "" {
		query += "kind=" + listBranchPermissionsOpts.kind
	}
	if listBranchPermissionsOpts.pattern != "" {
		if query != "" {
			query += "&"
		}
		query += "pattern=" + listBranchPermissionsOpts.pattern
	}
	if query != "" {
		apiURL += "?" + query
	}

	var permissions []branchPermission
	for apiURL != "" {
		resp, body, err := callBranchPermissionsAPI(http.MethodGet, apiURL, nil)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("branch restrictions API returned status %d: %s", resp.StatusCode, string(body))
		}

		var page paginatedBranchPermissions
		if err := json.Unmarshal(body, &page); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
		permissions = append(permissions, page.Values...)
		apiURL = page.Next
	}

	if len(permissions) == 0 {
		fmt.Println("There are no branch permission rules.")
		return nil
	}

	printBranchPermissionsTable(permissions)

	return nil
}
