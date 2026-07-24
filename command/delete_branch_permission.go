package command

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type deleteBranchPermissionOptions struct {
	workspace string
	slug      string
	id        int32
}

var deleteBranchPermissionOpts deleteBranchPermissionOptions

func (opts *deleteBranchPermissionOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("id is required")
	}
	return nil
}

var deleteBranchPermissionCmd = &cobra.Command{
	Use:   "branch-permission",
	Short: "Delete a branch permission (restriction) rule",
	RunE:  runDeleteBranchPermission,
}

func init() {
	deleteCmd.AddCommand(deleteBranchPermissionCmd)
	flags := deleteBranchPermissionCmd.Flags()
	flags.StringVarP(&deleteBranchPermissionOpts.workspace, "workspace", "w", "", "Workspace slug (defaults to the current repository's workspace)")
	flags.StringVarP(&deleteBranchPermissionOpts.slug, "slug", "s", "", "Repository slug (defaults to the current repository)")
	flags.Int32Var(&deleteBranchPermissionOpts.id, "id", 0, "Branch restriction rule id")
	_ = deleteBranchPermissionCmd.MarkFlagRequired("id")
}

func runDeleteBranchPermission(_ *cobra.Command, _ []string) error {
	if err := deleteBranchPermissionOpts.validate(); err != nil {
		return err
	}

	workspace, slug, err := resolveWorkspaceAndSlug(deleteBranchPermissionOpts.workspace, deleteBranchPermissionOpts.slug)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/branch-restrictions/%d",
		workspace,
		slug,
		deleteBranchPermissionOpts.id,
	)

	resp, body, err := callBranchPermissionsAPI(http.MethodDelete, apiURL, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("branch restrictions API returned status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println("Branch permission rule deleted successfully")
	return nil
}
