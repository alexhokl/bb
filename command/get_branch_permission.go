package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type getBranchPermissionOptions struct {
	workspace string
	slug      string
	id        int32
}

var getBranchPermissionOpts getBranchPermissionOptions

func (opts *getBranchPermissionOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("id is required")
	}
	return nil
}

var getBranchPermissionCmd = &cobra.Command{
	Use:   "branch-permission",
	Short: "Get details of a branch permission (restriction) rule",
	RunE:  runGetBranchPermission,
}

func init() {
	getCmd.AddCommand(getBranchPermissionCmd)
	flags := getBranchPermissionCmd.Flags()
	flags.StringVarP(&getBranchPermissionOpts.workspace, "workspace", "w", "", "Workspace slug (defaults to the current repository's workspace)")
	flags.StringVarP(&getBranchPermissionOpts.slug, "slug", "s", "", "Repository slug (defaults to the current repository)")
	flags.Int32Var(&getBranchPermissionOpts.id, "id", 0, "Branch restriction rule id")
	_ = getBranchPermissionCmd.MarkFlagRequired("id")
}

func runGetBranchPermission(_ *cobra.Command, _ []string) error {
	if err := getBranchPermissionOpts.validate(); err != nil {
		return err
	}

	workspace, slug, err := resolveWorkspaceAndSlug(getBranchPermissionOpts.workspace, getBranchPermissionOpts.slug)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/branch-restrictions/%d",
		workspace,
		slug,
		getBranchPermissionOpts.id,
	)

	resp, body, err := callBranchPermissionsAPI(http.MethodGet, apiURL, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("branch restrictions API returned status %d: %s", resp.StatusCode, string(body))
	}

	var p branchPermission
	if err := json.Unmarshal(body, &p); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}
	printBranchPermission(&p)

	return nil
}
