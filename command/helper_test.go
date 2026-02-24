package command

import (
	"testing"
)

func TestIdOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    idOptions
		wantErr bool
	}{
		{
			name:    "valid positive ID",
			opts:    idOptions{id: 1},
			wantErr: false,
		},
		{
			name:    "valid large ID",
			opts:    idOptions{id: 12345},
			wantErr: false,
		},
		{
			name:    "invalid zero ID",
			opts:    idOptions{id: 0},
			wantErr: true,
		},
		{
			name:    "invalid negative ID",
			opts:    idOptions{id: -1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("idOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreatePullRequestOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    createPullRequestOptions
		wantErr bool
	}{
		{
			name: "valid with destination branch",
			opts: createPullRequestOptions{
				destinationBranchName: "main",
			},
			wantErr: false,
		},
		{
			name: "valid with all fields",
			opts: createPullRequestOptions{
				destinationBranchName: "main",
				sourceBranchName:      "feature",
				title:                 "Test PR",
				message:               "Test message",
			},
			wantErr: false,
		},
		{
			name:    "invalid empty destination branch",
			opts:    createPullRequestOptions{},
			wantErr: true,
		},
		{
			name: "invalid empty destination with source",
			opts: createPullRequestOptions{
				sourceBranchName: "feature",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("createPullRequestOptions.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateListOptions(t *testing.T) {
	tests := []struct {
		name    string
		opts    listPullRequestOptions
		wantErr bool
	}{
		{
			name:    "valid empty options",
			opts:    listPullRequestOptions{},
			wantErr: false,
		},
		{
			name: "valid with different branches",
			opts: listPullRequestOptions{
				sourceBranchName:      "feature",
				destinationBranchName: "main",
			},
			wantErr: false,
		},
		{
			name: "valid with only source branch",
			opts: listPullRequestOptions{
				sourceBranchName: "feature",
			},
			wantErr: false,
		},
		{
			name: "valid with only destination branch",
			opts: listPullRequestOptions{
				destinationBranchName: "main",
			},
			wantErr: false,
		},
		{
			name: "invalid same source and destination",
			opts: listPullRequestOptions{
				sourceBranchName:      "main",
				destinationBranchName: "main",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateListOptions(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateListOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApproveOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    approveOptions
		wantErr bool
	}{
		{
			name: "valid with positive ID",
			opts: approveOptions{
				idOptions:    idOptions{id: 1},
				isKeepBranch: false,
			},
			wantErr: false,
		},
		{
			name: "valid with keep branch",
			opts: approveOptions{
				idOptions:    idOptions{id: 100},
				isKeepBranch: true,
			},
			wantErr: false,
		},
		{
			name: "invalid zero ID",
			opts: approveOptions{
				idOptions: idOptions{id: 0},
			},
			wantErr: true,
		},
		{
			name: "invalid negative ID",
			opts: approveOptions{
				idOptions: idOptions{id: -5},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("approveOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckoutOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    checkoutOptions
		wantErr bool
	}{
		{
			name: "valid with positive ID",
			opts: checkoutOptions{
				idOptions: idOptions{id: 1},
			},
			wantErr: false,
		},
		{
			name: "valid with all flags",
			opts: checkoutOptions{
				idOptions:                idOptions{id: 42},
				isShowDifftool:           true,
				isShowStat:               true,
				isNoMergeFromDestination: true,
			},
			wantErr: false,
		},
		{
			name: "invalid zero ID",
			opts: checkoutOptions{
				idOptions: idOptions{id: 0},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("checkoutOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    commentOptions
		wantErr bool
	}{
		{
			name: "valid with message",
			opts: commentOptions{
				idOptions: idOptions{id: 1},
				message:   "This is a comment",
			},
			wantErr: false,
		},
		{
			name: "valid with file",
			opts: commentOptions{
				idOptions:        idOptions{id: 1},
				markdownFilename: "comment.md",
			},
			wantErr: false,
		},
		{
			name: "invalid zero ID",
			opts: commentOptions{
				idOptions: idOptions{id: 0},
				message:   "comment",
			},
			wantErr: true,
		},
		{
			name: "invalid negative ID",
			opts: commentOptions{
				idOptions: idOptions{id: -1},
				message:   "comment",
			},
			wantErr: true,
		},
		{
			name: "invalid no message or file",
			opts: commentOptions{
				idOptions: idOptions{id: 1},
			},
			wantErr: true,
		},
		{
			name: "invalid both message and file",
			opts: commentOptions{
				idOptions:        idOptions{id: 1},
				message:          "comment",
				markdownFilename: "file.md",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("commentOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeclineOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    declineOptions
		wantErr bool
	}{
		{
			name: "valid with positive ID",
			opts: declineOptions{
				idOptions: idOptions{id: 1},
			},
			wantErr: false,
		},
		{
			name: "valid with message",
			opts: declineOptions{
				idOptions: idOptions{id: 1},
				message:   "Declining due to issues",
			},
			wantErr: false,
		},
		{
			name: "invalid zero ID",
			opts: declineOptions{
				idOptions: idOptions{id: 0},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("declineOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMergeOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    mergeOptions
		wantErr bool
	}{
		{
			name: "valid with positive ID",
			opts: mergeOptions{
				idOptions: idOptions{id: 1},
			},
			wantErr: false,
		},
		{
			name: "valid with all options",
			opts: mergeOptions{
				idOptions:           idOptions{id: 100},
				isKeepBranch:        true,
				isCloseSourceBranch: false,
				mergeStrategy:       "squash",
			},
			wantErr: false,
		},
		{
			name: "invalid zero ID",
			opts: mergeOptions{
				idOptions: idOptions{id: 0},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("mergeOptions.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMergeStrategy(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    MergeStrategy
		wantErr bool
	}{
		{
			name:    "merge commit",
			input:   "merge_commit",
			want:    MergeCommit,
			wantErr: false,
		},
		{
			name:    "no commit",
			input:   "no_commit",
			want:    NoCommit,
			wantErr: false,
		},
		{
			name:    "squash",
			input:   "squash",
			want:    Squash,
			wantErr: false,
		},
		{
			name:    "invalid strategy",
			input:   "invalid",
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "case sensitive - uppercase",
			input:   "SQUASH",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMergeStrategy(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMergeStrategy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMergeStrategy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCreateEnvironmentVariableOptions(t *testing.T) {
	tests := []struct {
		name    string
		opts    createEnvironmentVariableOptions
		wantErr bool
	}{
		{
			name: "valid with value",
			opts: createEnvironmentVariableOptions{
				name:  "production",
				key:   "API_KEY",
				value: "secret123",
			},
			wantErr: false,
		},
		{
			name: "valid with file",
			opts: createEnvironmentVariableOptions{
				name:          "production",
				key:           "API_KEY",
				valueFilePath: "/path/to/secret",
			},
			wantErr: false,
		},
		{
			name: "invalid no value or file",
			opts: createEnvironmentVariableOptions{
				name: "production",
				key:  "API_KEY",
			},
			wantErr: true,
		},
		{
			name: "invalid both value and file",
			opts: createEnvironmentVariableOptions{
				name:          "production",
				key:           "API_KEY",
				value:         "secret123",
				valueFilePath: "/path/to/secret",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateEnvironmentVariableOptions(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCreateEnvironmentVariableOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMergeStrategy_Constants(t *testing.T) {
	if MergeCommit != "merge_commit" {
		t.Errorf("MergeCommit = %v, want %v", MergeCommit, "merge_commit")
	}
	if NoCommit != "no_commit" {
		t.Errorf("NoCommit = %v, want %v", NoCommit, "no_commit")
	}
	if Squash != "squash" {
		t.Errorf("Squash = %v, want %v", Squash, "squash")
	}
}
