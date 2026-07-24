package command

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/alexhokl/helper/authhelper"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-branch-restrictions/

// branchPermissionAccount is a minimal representation of a user or group
// referenced by a branch restriction rule.
type branchPermissionAccount struct {
	Type_    string `json:"type,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Slug     string `json:"slug,omitempty"`
}

// branchPermission represents a branch restriction rule as returned/accepted
// by the Bitbucket API. This is defined locally (rather than reusing the
// generated swagger.Branchrestriction model) because the generated model is
// missing several fields (id, kind, pattern, branch_match_kind, branch_type,
// value) that are required to manage branch restrictions.
type branchPermission struct {
	Type_           string                    `json:"type,omitempty"`
	ID              int32                     `json:"id,omitempty"`
	Kind            string                    `json:"kind,omitempty"`
	BranchMatchKind string                    `json:"branch_match_kind,omitempty"`
	BranchType      string                    `json:"branch_type,omitempty"`
	Pattern         string                    `json:"pattern,omitempty"`
	Value           *int32                    `json:"value,omitempty"`
	Users           []branchPermissionAccount `json:"users,omitempty"`
	Groups          []branchPermissionAccount `json:"groups,omitempty"`
}

type paginatedBranchPermissions struct {
	Size     int32              `json:"size,omitempty"`
	Page     int32              `json:"page,omitempty"`
	Pagelen  int32              `json:"pagelen,omitempty"`
	Next     string             `json:"next,omitempty"`
	Previous string             `json:"previous,omitempty"`
	Values   []branchPermission `json:"values"`
}

var validBranchPermissionKinds = []string{
	"push",
	"force",
	"delete",
	"restrict_merges",
	"require_tasks_to_be_completed",
	"require_approvals_to_merge",
	"require_default_reviewer_approvals_to_merge",
	"require_no_changes_requested",
	"require_passing_builds_to_merge",
	"require_commits_behind",
	"reset_pullrequest_approvals_on_change",
	"smart_reset_pullrequest_approvals",
	"reset_pullrequest_changes_requested_on_change",
	"require_all_dependencies_merged",
	"enforce_merge_checks",
	"allow_auto_merge_when_builds_pass",
}

func isValidBranchPermissionKind(kind string) bool {
	for _, v := range validBranchPermissionKinds {
		if v == kind {
			return true
		}
	}
	return false
}

func parseAccountList(csv string) []branchPermissionAccount {
	if csv == "" {
		return nil
	}
	parts := strings.Split(csv, ",")
	accounts := make([]branchPermissionAccount, 0, len(parts))
	for _, p := range parts {
		v := strings.TrimSpace(p)
		if v == "" {
			continue
		}
		accounts = append(accounts, branchPermissionAccount{Uuid: v})
	}
	return accounts
}

func parseGroupList(csv string) []branchPermissionAccount {
	if csv == "" {
		return nil
	}
	parts := strings.Split(csv, ",")
	accounts := make([]branchPermissionAccount, 0, len(parts))
	for _, p := range parts {
		v := strings.TrimSpace(p)
		if v == "" {
			continue
		}
		accounts = append(accounts, branchPermissionAccount{Slug: v})
	}
	return accounts
}

func printBranchPermission(p *branchPermission) {
	fmt.Printf("id\t%d\n", p.ID)
	fmt.Printf("kind\t%s\n", p.Kind)
	fmt.Printf("branch_match_kind\t%s\n", p.BranchMatchKind)
	if p.BranchType != "" {
		fmt.Printf("branch_type\t%s\n", p.BranchType)
	}
	if p.Pattern != "" {
		fmt.Printf("pattern\t%s\n", p.Pattern)
	}
	if p.Value != nil {
		fmt.Printf("value\t%d\n", *p.Value)
	}
	for _, u := range p.Users {
		fmt.Printf("user\t%s\n", u.Uuid)
	}
	for _, g := range p.Groups {
		fmt.Printf("group\t%s\n", g.Slug)
	}
}

// summarizeAccounts returns a short human-readable summary of the number of
// accounts (users or groups) referenced by a branch restriction rule, e.g.
// "-", "1 user", "3 users".
func summarizeAccounts(accounts []branchPermissionAccount, singular, plural string) string {
	n := len(accounts)
	if n == 0 {
		return "-"
	}
	if n == 1 {
		return fmt.Sprintf("1 %s", singular)
	}
	return fmt.Sprintf("%d %s", n, plural)
}

// branchPermissionTableRow converts a branch restriction rule into the
// columns of its table row: ID, KIND, MATCH, PATTERN/TYPE, VALUE, USERS,
// GROUPS.
func branchPermissionTableRow(p *branchPermission) []string {
	patternOrType := p.Pattern
	if p.BranchMatchKind == "branching_model" {
		patternOrType = p.BranchType
	}

	value := ""
	if p.Value != nil {
		value = fmt.Sprintf("%d", *p.Value)
	}

	return []string{
		fmt.Sprintf("%d", p.ID),
		p.Kind,
		p.BranchMatchKind,
		patternOrType,
		value,
		summarizeAccounts(p.Users, "user", "users"),
		summarizeAccounts(p.Groups, "group", "groups"),
	}
}

// printBranchPermissionsTable prints a list of branch restriction rules as
// an aligned table.
func printBranchPermissionsTable(permissions []branchPermission) {
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tKIND\tMATCH\tPATTERN/TYPE\tVALUE\tUSERS\tGROUPS")
	for i := range permissions {
		row := branchPermissionTableRow(&permissions[i])
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}

func callBranchPermissionsAPI(method, url string, body []byte) (*http.Response, []byte, error) {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return nil, nil, err
	}

	var reqBody io.Reader
	if body != nil {
		reqBody = strings.NewReader(string(body))
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to call branch restrictions API: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return resp, respBody, nil
}

func marshalBranchPermission(p *branchPermission) ([]byte, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	return body, nil
}

// resolveWorkspaceAndSlug fills in workspace and slug from the git remote of
// the current directory when either of them is not explicitly provided.
func resolveWorkspaceAndSlug(workspace, slug string) (string, string, error) {
	if workspace != "" && slug != "" {
		return workspace, slug, nil
	}

	repo, err := getRepositoryInfoFromCurrentPath()
	if err != nil {
		return "", "", err
	}

	if workspace == "" {
		workspace = repo.Org
	}
	if slug == "" {
		slug = repo.Name
	}

	return workspace, slug, nil
}
