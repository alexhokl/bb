package command

import (
	"testing"
)

func TestUpdatePullRequestOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    updatePullRequestOptions
		wantErr bool
	}{
		{
			name:    "title only",
			opts:    updatePullRequestOptions{idOptions: idOptions{id: 1}, title: "title"},
			wantErr: false,
		},
		{
			name:    "message only",
			opts:    updatePullRequestOptions{idOptions: idOptions{id: 1}, message: "message"},
			wantErr: false,
		},
		{
			name:    "append file",
			opts:    updatePullRequestOptions{idOptions: idOptions{id: 1}, markdownFilename: "a.md", isAppend: true},
			wantErr: false,
		},
		{
			name:    "invalid ID",
			opts:    updatePullRequestOptions{title: "title"},
			wantErr: true,
		},
		{
			name:    "nothing to update",
			opts:    updatePullRequestOptions{idOptions: idOptions{id: 1}},
			wantErr: true,
		},
		{
			name:    "message and file",
			opts:    updatePullRequestOptions{idOptions: idOptions{id: 1}, message: "message", markdownFilename: "a.md"},
			wantErr: true,
		},
		{
			name:    "append without description",
			opts:    updatePullRequestOptions{idOptions: idOptions{id: 1}, title: "title", isAppend: true},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("updatePullRequestOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildDescription(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		message  string
		isAppend bool
		want     string
	}{
		{
			name:    "replace",
			current: "old",
			message: "new",
			want:    "new",
		},
		{
			name:     "append",
			current:  "old",
			message:  "new",
			isAppend: true,
			want:     "old\n\nnew",
		},
		{
			name:     "append trims trailing new lines",
			current:  "old\n\n\n",
			message:  "new",
			isAppend: true,
			want:     "old\n\nnew",
		},
		{
			name:     "append to empty description",
			current:  "",
			message:  "new",
			isAppend: true,
			want:     "new",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildDescription(tt.current, tt.message, tt.isAppend)
			if got != tt.want {
				t.Errorf("buildDescription() = %q, want %q", got, tt.want)
			}
		})
	}
}
