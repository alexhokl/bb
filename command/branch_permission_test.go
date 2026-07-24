package command

import (
	"testing"
)

func TestListBranchPermissionsOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    listBranchPermissionsOptions
		wantErr bool
	}{
		{
			name: "valid minimal",
			opts: listBranchPermissionsOptions{
				workspace: "ws",
				slug:      "repo",
			},
			wantErr: false,
		},
		{
			name: "valid with kind filter",
			opts: listBranchPermissionsOptions{
				workspace: "ws",
				slug:      "repo",
				kind:      "push",
			},
			wantErr: false,
		},
		{
			name:    "valid without workspace or slug (falls back to git)",
			opts:    listBranchPermissionsOptions{},
			wantErr: false,
		},
		{
			name: "invalid kind",
			opts: listBranchPermissionsOptions{
				workspace: "ws",
				slug:      "repo",
				kind:      "not-a-kind",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("listBranchPermissionsOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBranchPermissionOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    getBranchPermissionOptions
		wantErr bool
	}{
		{
			name: "valid",
			opts: getBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				id:        1,
			},
			wantErr: false,
		},
		{
			name: "valid without workspace or slug (falls back to git)",
			opts: getBranchPermissionOptions{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid zero id",
			opts: getBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("getBranchPermissionOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateBranchPermissionOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    createBranchPermissionOptions
		wantErr bool
	}{
		{
			name: "valid with pattern",
			opts: createBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				kind:      "push",
				pattern:   "main",
			},
			wantErr: false,
		},
		{
			name: "valid with branch type",
			opts: createBranchPermissionOptions{
				workspace:  "ws",
				slug:       "repo",
				kind:       "delete",
				branchType: "production",
			},
			wantErr: false,
		},
		{
			name: "valid without workspace or slug (falls back to git)",
			opts: createBranchPermissionOptions{
				kind:    "push",
				pattern: "main",
			},
			wantErr: false,
		},
		{
			name: "invalid missing kind",
			opts: createBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				pattern:   "main",
			},
			wantErr: true,
		},
		{
			name: "invalid kind",
			opts: createBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				kind:      "not-a-kind",
				pattern:   "main",
			},
			wantErr: true,
		},
		{
			name: "invalid missing pattern and branch type",
			opts: createBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				kind:      "push",
			},
			wantErr: true,
		},
		{
			name: "invalid both pattern and branch type",
			opts: createBranchPermissionOptions{
				workspace:  "ws",
				slug:       "repo",
				kind:       "push",
				pattern:    "main",
				branchType: "production",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("createBranchPermissionOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateBranchPermissionOptions_ToBranchPermission(t *testing.T) {
	opts := createBranchPermissionOptions{
		kind:    "push",
		pattern: "main",
		value:   2,
		users:   "user-1, user-2",
		groups:  "group-1",
	}
	p := opts.toBranchPermission()

	if p.Kind != "push" {
		t.Errorf("Kind = %v, want push", p.Kind)
	}
	if p.BranchMatchKind != "glob" {
		t.Errorf("BranchMatchKind = %v, want glob", p.BranchMatchKind)
	}
	if p.Pattern != "main" {
		t.Errorf("Pattern = %v, want main", p.Pattern)
	}
	if p.Value == nil || *p.Value != 2 {
		t.Errorf("Value = %v, want 2", p.Value)
	}
	if len(p.Users) != 2 || p.Users[0].Uuid != "user-1" || p.Users[1].Uuid != "user-2" {
		t.Errorf("Users = %v, want [user-1 user-2]", p.Users)
	}
	if len(p.Groups) != 1 || p.Groups[0].Uuid != "group-1" {
		t.Errorf("Groups = %v, want [group-1]", p.Groups)
	}
}

func TestCreateBranchPermissionOptions_ToBranchPermission_BranchType(t *testing.T) {
	opts := createBranchPermissionOptions{
		kind:       "delete",
		branchType: "production",
	}
	p := opts.toBranchPermission()

	if p.BranchMatchKind != "branching_model" {
		t.Errorf("BranchMatchKind = %v, want branching_model", p.BranchMatchKind)
	}
	if p.BranchType != "production" {
		t.Errorf("BranchType = %v, want production", p.BranchType)
	}
	if p.Pattern != "" {
		t.Errorf("Pattern = %v, want empty", p.Pattern)
	}
}

func TestUpdateBranchPermissionOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    updateBranchPermissionOptions
		wantErr bool
	}{
		{
			name: "valid minimal",
			opts: updateBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				id:        1,
			},
			wantErr: false,
		},
		{
			name: "valid with kind and pattern",
			opts: updateBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				id:        1,
				kind:      "push",
				pattern:   "main",
			},
			wantErr: false,
		},
		{
			name: "valid without workspace or slug (falls back to git)",
			opts: updateBranchPermissionOptions{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid zero id",
			opts: updateBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
			},
			wantErr: true,
		},
		{
			name: "invalid kind",
			opts: updateBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				id:        1,
				kind:      "not-a-kind",
			},
			wantErr: true,
		},
		{
			name: "invalid both pattern and branch type",
			opts: updateBranchPermissionOptions{
				workspace:  "ws",
				slug:       "repo",
				id:         1,
				pattern:    "main",
				branchType: "production",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("updateBranchPermissionOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteBranchPermissionOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    deleteBranchPermissionOptions
		wantErr bool
	}{
		{
			name: "valid",
			opts: deleteBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
				id:        1,
			},
			wantErr: false,
		},
		{
			name: "valid without workspace or slug (falls back to git)",
			opts: deleteBranchPermissionOptions{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid zero id",
			opts: deleteBranchPermissionOptions{
				workspace: "ws",
				slug:      "repo",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteBranchPermissionOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidBranchPermissionKind(t *testing.T) {
	if !isValidBranchPermissionKind("push") {
		t.Error("expected push to be a valid kind")
	}
	if isValidBranchPermissionKind("not-a-kind") {
		t.Error("expected not-a-kind to be invalid")
	}
}

func TestParseAccountList(t *testing.T) {
	got := parseAccountList("a, b ,c")
	if len(got) != 3 {
		t.Fatalf("expected 3 accounts, got %d", len(got))
	}
	if got[0].Uuid != "a" || got[1].Uuid != "b" || got[2].Uuid != "c" {
		t.Errorf("unexpected parsed accounts: %+v", got)
	}

	if got := parseAccountList(""); got != nil {
		t.Errorf("expected nil for empty string, got %v", got)
	}
}

func TestResolveWorkspaceAndSlug(t *testing.T) {
	// When both are explicitly provided, they should be returned as-is
	// without needing to consult git.
	workspace, slug, err := resolveWorkspaceAndSlug("ws", "repo")
	if err != nil {
		t.Fatalf("resolveWorkspaceAndSlug() unexpected error = %v", err)
	}
	if workspace != "ws" || slug != "repo" {
		t.Errorf("resolveWorkspaceAndSlug() = (%v, %v), want (ws, repo)", workspace, slug)
	}
}

func TestSummarizeAccounts(t *testing.T) {
	tests := []struct {
		name     string
		accounts []branchPermissionAccount
		want     string
	}{
		{
			name:     "empty",
			accounts: nil,
			want:     "-",
		},
		{
			name:     "one",
			accounts: []branchPermissionAccount{{Uuid: "a"}},
			want:     "1 user",
		},
		{
			name:     "many",
			accounts: []branchPermissionAccount{{Uuid: "a"}, {Uuid: "b"}, {Uuid: "c"}},
			want:     "3 users",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := summarizeAccounts(tt.accounts, "user", "users")
			if got != tt.want {
				t.Errorf("summarizeAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBranchPermissionTableRow(t *testing.T) {
	value := int32(2)

	tests := []struct {
		name string
		p    branchPermission
		want []string
	}{
		{
			name: "glob with value and accounts",
			p: branchPermission{
				ID:              34,
				Kind:            "push",
				BranchMatchKind: "glob",
				Pattern:         "release/*",
				Value:           &value,
				Users:           []branchPermissionAccount{{Uuid: "u1"}},
				Groups:          []branchPermissionAccount{{Uuid: "g1"}, {Uuid: "g2"}},
			},
			want: []string{"34", "push", "glob", "release/*", "2", "1 user", "2 groups"},
		},
		{
			name: "branching model without value or accounts",
			p: branchPermission{
				ID:              1,
				Kind:            "delete",
				BranchMatchKind: "branching_model",
				BranchType:      "production",
			},
			want: []string{"1", "delete", "branching_model", "production", "", "-", "-"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := branchPermissionTableRow(&tt.p)
			if len(got) != len(tt.want) {
				t.Fatalf("branchPermissionTableRow() = %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("branchPermissionTableRow()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
