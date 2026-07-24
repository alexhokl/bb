package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type createBranchPermissionOptions struct {
	workspace  string
	slug       string
	kind       string
	pattern    string
	branchType string
	value      int32
	users      string
	groups     string
}

var createBranchPermissionOpts createBranchPermissionOptions

func (opts *createBranchPermissionOptions) validate() error {
	if opts.kind == "" {
		return fmt.Errorf("kind is required")
	}
	if !isValidBranchPermissionKind(opts.kind) {
		return fmt.Errorf("kind must be one of: %v", validBranchPermissionKinds)
	}
	if opts.pattern == "" && opts.branchType == "" {
		return fmt.Errorf("either pattern or branch-type is required")
	}
	if opts.pattern != "" && opts.branchType != "" {
		return fmt.Errorf("only one of pattern or branch-type may be specified")
	}
	return nil
}

func (opts *createBranchPermissionOptions) toBranchPermission() *branchPermission {
	p := &branchPermission{
		Type_: "branchrestriction",
		Kind:  opts.kind,
	}
	if opts.pattern != "" {
		p.BranchMatchKind = "glob"
		p.Pattern = opts.pattern
	} else {
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

var createBranchPermissionCmd = &cobra.Command{
	Use:   "branch-permission",
	Short: "Create a branch permission (restriction) rule",
	RunE:  runCreateBranchPermission,
}

func init() {
	createCmd.AddCommand(createBranchPermissionCmd)
	flags := createBranchPermissionCmd.Flags()
	flags.StringVarP(&createBranchPermissionOpts.workspace, "workspace", "w", "", "Workspace slug (defaults to the current repository's workspace)")
	flags.StringVarP(&createBranchPermissionOpts.slug, "slug", "s", "", "Repository slug (defaults to the current repository)")
	flags.StringVar(&createBranchPermissionOpts.kind, "kind", "", fmt.Sprintf("Kind of restriction (%v)", validBranchPermissionKinds))
	flags.StringVar(&createBranchPermissionOpts.pattern, "pattern", "", "Glob pattern of branches this rule applies to")
	flags.StringVar(&createBranchPermissionOpts.branchType, "branch-type", "", "Branching model branch type this rule applies to (production, development, bugfix, release, feature, hotfix)")
	flags.Int32Var(&createBranchPermissionOpts.value, "value", 0, "Value of the restriction (e.g. minimum number of approvals)")
	flags.StringVar(&createBranchPermissionOpts.users, "users", "", "Comma-separated list of user UUIDs exempted from the restriction")
	flags.StringVar(&createBranchPermissionOpts.groups, "groups", "", "Comma-separated list of group slugs exempted from the restriction")
	_ = createBranchPermissionCmd.MarkFlagRequired("kind")
}

func runCreateBranchPermission(_ *cobra.Command, _ []string) error {
	if err := createBranchPermissionOpts.validate(); err != nil {
		return err
	}

	workspace, slug, err := resolveWorkspaceAndSlug(createBranchPermissionOpts.workspace, createBranchPermissionOpts.slug)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/branch-restrictions",
		workspace,
		slug,
	)

	reqBody, err := marshalBranchPermission(createBranchPermissionOpts.toBranchPermission())
	if err != nil {
		return err
	}

	resp, body, err := callBranchPermissionsAPI(http.MethodPost, apiURL, reqBody)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("branch restrictions API returned status %d: %s", resp.StatusCode, string(body))
	}

	var p branchPermission
	if err := json.Unmarshal(body, &p); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}
	fmt.Printf("Branch permission rule created successfully with id %d\n", p.ID)

	return nil
}
