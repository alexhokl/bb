package models

import (
	"strings"
	"testing"
	"time"
)

func TestComment_ToString(t *testing.T) {
	fixedTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	tests := []struct {
		name            string
		comment         Comment
		wantContains    []string
		wantNotContains []string
	}{
		{
			name: "basic comment",
			comment: Comment{
				Content:   RawContent{Raw: "This is a test comment"},
				UpdatedOn: fixedTime,
				User:      User{DisplayName: "John Doe"},
			},
			wantContains: []string{
				"Comment by",
				"John Doe",
				"This is a test comment",
			},
		},
		{
			name: "comment with special characters",
			comment: Comment{
				Content:   RawContent{Raw: "Fix: bug #123"},
				UpdatedOn: fixedTime,
				User:      User{DisplayName: "Jane Smith"},
			},
			wantContains: []string{
				"Jane Smith",
				"Fix: bug #123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.comment.ToString()
			for _, want := range tt.wantContains {
				if !strings.Contains(result, want) {
					t.Errorf("Comment.ToString() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestUpdate_ToString(t *testing.T) {
	fixedTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	tests := []struct {
		name         string
		update       Update
		wantContains []string
	}{
		{
			name: "basic update",
			update: Update{
				Date:   fixedTime,
				Author: User{DisplayName: "John Doe"},
				Source: Commit{
					Commit: CommitInfo{Hash: "abc123def"},
				},
			},
			wantContains: []string{
				"Commit made by",
				"John Doe",
				"abc123def",
			},
		},
		{
			name: "update with full hash",
			update: Update{
				Date:   fixedTime,
				Author: User{DisplayName: "Jane Smith"},
				Source: Commit{
					Commit: CommitInfo{Hash: "1234567890abcdef1234567890abcdef12345678"},
				},
			},
			wantContains: []string{
				"Jane Smith",
				"1234567890abcdef1234567890abcdef12345678",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.update.ToString()
			for _, want := range tt.wantContains {
				if !strings.Contains(result, want) {
					t.Errorf("Update.ToString() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestRepository_Fields(t *testing.T) {
	repo := Repository{
		Org:  "myorg",
		Name: "myrepo",
	}

	if repo.Org != "myorg" {
		t.Errorf("Repository.Org = %v, want %v", repo.Org, "myorg")
	}
	if repo.Name != "myrepo" {
		t.Errorf("Repository.Name = %v, want %v", repo.Name, "myrepo")
	}
}

func TestUser_Fields(t *testing.T) {
	user := User{
		Username:    "johndoe",
		DisplayName: "John Doe",
	}

	if user.Username != "johndoe" {
		t.Errorf("User.Username = %v, want %v", user.Username, "johndoe")
	}
	if user.DisplayName != "John Doe" {
		t.Errorf("User.DisplayName = %v, want %v", user.DisplayName, "John Doe")
	}
}

func TestBranch_Fields(t *testing.T) {
	branch := Branch{Name: "feature-branch"}

	if branch.Name != "feature-branch" {
		t.Errorf("Branch.Name = %v, want %v", branch.Name, "feature-branch")
	}
}

func TestRawContent_Fields(t *testing.T) {
	content := RawContent{Raw: "Some raw content"}

	if content.Raw != "Some raw content" {
		t.Errorf("RawContent.Raw = %v, want %v", content.Raw, "Some raw content")
	}
}

func TestCommitInfo_Fields(t *testing.T) {
	tests := []struct {
		name       string
		commitInfo CommitInfo
		wantHash   string
		wantRaw    string
	}{
		{
			name: "short hash",
			commitInfo: CommitInfo{
				Hash:    "abc123",
				Summary: RawContent{Raw: "Fix bug"},
			},
			wantHash: "abc123",
			wantRaw:  "Fix bug",
		},
		{
			name: "full hash",
			commitInfo: CommitInfo{
				Hash:    "1234567890abcdef1234567890abcdef12345678",
				Summary: RawContent{Raw: "Add feature"},
			},
			wantHash: "1234567890abcdef1234567890abcdef12345678",
			wantRaw:  "Add feature",
		},
		{
			name: "empty summary",
			commitInfo: CommitInfo{
				Hash:    "def456",
				Summary: RawContent{},
			},
			wantHash: "def456",
			wantRaw:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.commitInfo.Hash != tt.wantHash {
				t.Errorf("CommitInfo.Hash = %v, want %v", tt.commitInfo.Hash, tt.wantHash)
			}
			if tt.commitInfo.Summary.Raw != tt.wantRaw {
				t.Errorf("CommitInfo.Summary.Raw = %v, want %v", tt.commitInfo.Summary.Raw, tt.wantRaw)
			}
		})
	}
}

func TestCommitBranch_Fields(t *testing.T) {
	cb := CommitBranch{
		Branch: Branch{Name: "main"},
	}

	if cb.Branch.Name != "main" {
		t.Errorf("CommitBranch.Branch.Name = %v, want %v", cb.Branch.Name, "main")
	}
}

func TestCommit_Fields(t *testing.T) {
	commit := Commit{
		CommitBranch: CommitBranch{
			Branch: Branch{Name: "feature"},
		},
		Commit: CommitInfo{
			Hash:    "abc123",
			Summary: RawContent{Raw: "Test commit"},
		},
	}

	if commit.Branch.Name != "feature" {
		t.Errorf("Commit.Branch.Name = %v, want %v", commit.Branch.Name, "feature")
	}
	if commit.Commit.Hash != "abc123" {
		t.Errorf("Commit.Commit.Hash = %v, want %v", commit.Commit.Hash, "abc123")
	}
	if commit.Commit.Summary.Raw != "Test commit" {
		t.Errorf("Commit.Commit.Summary.Raw = %v, want %v", commit.Commit.Summary.Raw, "Test commit")
	}
}

func TestPullRequestActivity_Fields(t *testing.T) {
	fixedTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	activity := PullRequestActivity{
		Update: Update{
			Date:   fixedTime,
			Author: User{DisplayName: "John"},
		},
		Comment: Comment{
			Content: RawContent{Raw: "LGTM"},
			User:    User{DisplayName: "Jane"},
		},
	}

	if activity.Update.Author.DisplayName != "John" {
		t.Errorf("PullRequestActivity.Update.Author.DisplayName = %v, want %v", activity.Update.Author.DisplayName, "John")
	}
	if activity.Comment.User.DisplayName != "Jane" {
		t.Errorf("PullRequestActivity.Comment.User.DisplayName = %v, want %v", activity.Comment.User.DisplayName, "Jane")
	}
	if activity.Comment.Content.Raw != "LGTM" {
		t.Errorf("PullRequestActivity.Comment.Content.Raw = %v, want %v", activity.Comment.Content.Raw, "LGTM")
	}
}

func TestComment_Fields(t *testing.T) {
	createdTime := time.Date(2024, 1, 10, 8, 0, 0, 0, time.UTC)
	updatedTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	comment := Comment{
		Content:   RawContent{Raw: "This needs review"},
		CreatedOn: createdTime,
		UpdatedOn: updatedTime,
		User: User{
			Username:    "reviewer",
			DisplayName: "Code Reviewer",
		},
	}

	if comment.Content.Raw != "This needs review" {
		t.Errorf("Comment.Content.Raw = %v, want %v", comment.Content.Raw, "This needs review")
	}
	if !comment.CreatedOn.Equal(createdTime) {
		t.Errorf("Comment.CreatedOn = %v, want %v", comment.CreatedOn, createdTime)
	}
	if !comment.UpdatedOn.Equal(updatedTime) {
		t.Errorf("Comment.UpdatedOn = %v, want %v", comment.UpdatedOn, updatedTime)
	}
	if comment.User.Username != "reviewer" {
		t.Errorf("Comment.User.Username = %v, want %v", comment.User.Username, "reviewer")
	}
}

func TestUpdate_Fields(t *testing.T) {
	updateTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	update := Update{
		Date: updateTime,
		Source: Commit{
			CommitBranch: CommitBranch{
				Branch: Branch{Name: "feature-x"},
			},
			Commit: CommitInfo{
				Hash:    "deadbeef",
				Summary: RawContent{Raw: "Update feature"},
			},
		},
		Author: User{
			Username:    "dev",
			DisplayName: "Developer",
		},
	}

	if !update.Date.Equal(updateTime) {
		t.Errorf("Update.Date = %v, want %v", update.Date, updateTime)
	}
	if update.Source.Branch.Name != "feature-x" {
		t.Errorf("Update.Source.Branch.Name = %v, want %v", update.Source.Branch.Name, "feature-x")
	}
	if update.Source.Commit.Hash != "deadbeef" {
		t.Errorf("Update.Source.Commit.Hash = %v, want %v", update.Source.Commit.Hash, "deadbeef")
	}
	if update.Author.DisplayName != "Developer" {
		t.Errorf("Update.Author.DisplayName = %v, want %v", update.Author.DisplayName, "Developer")
	}
}

func TestComment_ToString_EmptyFields(t *testing.T) {
	comment := Comment{}
	result := comment.ToString()

	if !strings.Contains(result, "Comment by") {
		t.Errorf("Comment.ToString() should contain 'Comment by', got %v", result)
	}
}

func TestUpdate_ToString_EmptyFields(t *testing.T) {
	update := Update{}
	result := update.ToString()

	if !strings.Contains(result, "Commit made by") {
		t.Errorf("Update.ToString() should contain 'Commit made by', got %v", result)
	}
}
