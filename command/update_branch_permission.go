package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type updateBranchPermissionOptions struct {
	workspace  string
	slug       string
	id         int32
	kind       string
	pattern    string
	branchType string
	value      int32
	users      string
	groups     string
}

var updateBranchPermissionOpts updateBranchPermissionOptions

func (opts *updateBranchPermissionOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("id is required")
	}
	if opts.kind != "" && !isValidBranchPermissionKind(opts.kind) {
		return fmt.Errorf("kind must be one of: %v", validBranchPermissionKinds)
	}
	if opts.pattern != "" && opts.branchType != "" {
		return fmt.Errorf("only one of pattern or branch-type may be specified")
	}
	return nil
}

func (opts *updateBranchPermissionOptions) toBranchPermission() *branchPermission {
	p := &branchPermission{
		Type_: "branchrestriction",
	}
	if opts.kind != "" {
		p.Kind = opts.kind
	}
	if opts.pattern != "" {
		p.BranchMatchKind = "glob"
		p.Pattern = opts.pattern
	} else if opts.branchType != "" {
		p.BranchMatchKind = "branching_model"
		p.BranchType = opts.branchType
	}
	if opts.value > 0 {
		v := opts.value
		p.Value = &v
	}
	if opts.users != "" {
		p.Users = parseAccountList(opts.users)
	}
	if opts.groups != "" {
		p.Groups = parseGroupList(opts.groups)
	}
	return p
}

var updateBranchPermissionCmd = &cobra.Command{
	Use:   "branch-permission",
	Short: "Update a branch permission (restriction) rule",
	RunE:  runUpdateBranchPermission,
}

func init() {
	updateCmd.AddCommand(updateBranchPermissionCmd)
	flags := updateBranchPermissionCmd.Flags()
	flags.StringVarP(&updateBranchPermissionOpts.workspace, "workspace", "w", "", "Workspace slug (defaults to the current repository's workspace)")
	flags.StringVarP(&updateBranchPermissionOpts.slug, "slug", "s", "", "Repository slug (defaults to the current repository)")
	flags.Int32Var(&updateBranchPermissionOpts.id, "id", 0, "Branch restriction rule id")
	flags.StringVar(&updateBranchPermissionOpts.kind, "kind", "", fmt.Sprintf("Kind of restriction (%v)", validBranchPermissionKinds))
	flags.StringVar(&updateBranchPermissionOpts.pattern, "pattern", "", "Glob pattern of branches this rule applies to")
	flags.StringVar(&updateBranchPermissionOpts.branchType, "branch-type", "", "Branching model branch type this rule applies to (production, development, bugfix, release, feature, hotfix)")
	flags.Int32Var(&updateBranchPermissionOpts.value, "value", 0, "Value of the restriction (e.g. minimum number of approvals)")
	flags.StringVar(&updateBranchPermissionOpts.users, "users", "", "Comma-separated list of user UUIDs exempted from the restriction")
	flags.StringVar(&updateBranchPermissionOpts.groups, "groups", "", "Comma-separated list of group slugs exempted from the restriction")
	_ = updateBranchPermissionCmd.MarkFlagRequired("id")
}

func runUpdateBranchPermission(_ *cobra.Command, _ []string) error {
	if err := updateBranchPermissionOpts.validate(); err != nil {
		return err
	}

	workspace, slug, err := resolveWorkspaceAndSlug(updateBranchPermissionOpts.workspace, updateBranchPermissionOpts.slug)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/branch-restrictions/%d",
		workspace,
		slug,
		updateBranchPermissionOpts.id,
	)

	reqBody, err := marshalBranchPermission(updateBranchPermissionOpts.toBranchPermission())
	if err != nil {
		return err
	}

	resp, body, err := callBranchPermissionsAPI(http.MethodPut, apiURL, reqBody)
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
	fmt.Printf("Branch permission rule %d updated successfully\n", p.ID)

	return nil
}
