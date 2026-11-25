package command

import (
	"testing"
)

func TestUpdateRepositoryOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    updateRepositoryOptions
		wantErr bool
	}{
		{
			name:    "project key only",
			opts:    updateRepositoryOptions{workspace: "w", slug: "s", projectKey: "P"},
			wantErr: false,
		},
		{
			name:    "main branch only",
			opts:    updateRepositoryOptions{workspace: "w", slug: "s", mainBranch: "main"},
			wantErr: false,
		},
		{
			name:    "project key and main branch",
			opts:    updateRepositoryOptions{workspace: "w", slug: "s", projectKey: "P", mainBranch: "main"},
			wantErr: false,
		},
		{
			name:    "nothing to update",
			opts:    updateRepositoryOptions{workspace: "w", slug: "s"},
			wantErr: true,
		},
		{
			name:    "missing workspace",
			opts:    updateRepositoryOptions{slug: "s", mainBranch: "main"},
			wantErr: true,
		},
		{
			name:    "missing slug",
			opts:    updateRepositoryOptions{workspace: "w", mainBranch: "main"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("updateRepositoryOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildRepositoryUpdate(t *testing.T) {
	t.Run("main branch only", func(t *testing.T) {
		repo := buildRepositoryUpdate(updateRepositoryOptions{mainBranch: "production"})
		if repo.Project != nil {
			t.Errorf("Project = %v, want nil", repo.Project)
		}
		if repo.Mainbranch == nil || repo.Mainbranch.Name != "production" {
			t.Errorf("Mainbranch = %v, want production", repo.Mainbranch)
		}
	})

	t.Run("project key only", func(t *testing.T) {
		repo := buildRepositoryUpdate(updateRepositoryOptions{projectKey: "P"})
		if repo.Mainbranch != nil {
			t.Errorf("Mainbranch = %v, want nil", repo.Mainbranch)
		}
		if repo.Project == nil || repo.Project.Key != "P" {
			t.Errorf("Project = %v, want key P", repo.Project)
		}
	})
}
