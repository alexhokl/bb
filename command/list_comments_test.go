package command

import (
	"reflect"
	"testing"
	"time"

	"github.com/alexhokl/bb/swagger"
)

func TestListCommentsOptions_Validate(t *testing.T) {
	tests := []struct {
		name    string
		opts    listCommentsOptions
		wantErr bool
	}{
		{
			name:    "valid text format",
			opts:    listCommentsOptions{idOptions: idOptions{id: 1}, format: "text"},
			wantErr: false,
		},
		{
			name:    "valid json format",
			opts:    listCommentsOptions{idOptions: idOptions{id: 1}, format: "json"},
			wantErr: false,
		},
		{
			name:    "invalid format",
			opts:    listCommentsOptions{idOptions: idOptions{id: 1}, format: "xml"},
			wantErr: true,
		},
		{
			name:    "invalid zero ID",
			opts:    listCommentsOptions{idOptions: idOptions{id: 0}, format: "text"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("listCommentsOptions.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFormatCommentLocation(t *testing.T) {
	tests := []struct {
		name   string
		inline *swagger.CommentInline
		want   string
	}{
		{name: "general comment", inline: nil, want: ""},
		{name: "new line only", inline: &swagger.CommentInline{Path: "a.go", To: 10}, want: "a.go:10"},
		{name: "old line only", inline: &swagger.CommentInline{Path: "a.go", From: 7}, want: "a.go:7 (old)"},
		{name: "both lines", inline: &swagger.CommentInline{Path: "a.go", From: 7, To: 10}, want: "a.go:10"},
		{name: "file only", inline: &swagger.CommentInline{Path: "a.go"}, want: "a.go"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatCommentLocation(tt.inline); got != tt.want {
				t.Errorf("formatCommentLocation() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildCommentThreads(t *testing.T) {
	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	comment := func(id int32, minute int, parentID int32) swagger.PullrequestComment {
		c := swagger.PullrequestComment{Id: id, CreatedOn: base.Add(time.Duration(minute) * time.Minute)}
		if parentID != 0 {
			c.Parent = &swagger.Comment{Id: parentID}
		}
		return c
	}

	tests := []struct {
		name     string
		comments []swagger.PullrequestComment
		want     map[int32][]int32
		order    []int32
	}{
		{
			name:     "empty",
			comments: nil,
			want:     map[int32][]int32{},
			order:    nil,
		},
		{
			name:     "orders top-level comments by creation time",
			comments: []swagger.PullrequestComment{comment(2, 5, 0), comment(1, 1, 0)},
			want:     map[int32][]int32{1: nil, 2: nil},
			order:    []int32{1, 2},
		},
		{
			name: "nests replies and replies to replies under the top-level comment",
			comments: []swagger.PullrequestComment{
				comment(1, 1, 0),
				comment(3, 3, 2),
				comment(2, 2, 1),
				comment(4, 4, 0),
			},
			want:  map[int32][]int32{1: {2, 3}, 4: nil},
			order: []int32{1, 4},
		},
		{
			name:     "reply with missing parent becomes top-level",
			comments: []swagger.PullrequestComment{comment(1, 1, 0), comment(5, 2, 99)},
			want:     map[int32][]int32{1: nil, 5: nil},
			order:    []int32{1, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threads := buildCommentThreads(tt.comments)
			var order []int32
			got := map[int32][]int32{}
			for _, thread := range threads {
				order = append(order, thread.comment.Id)
				var replyIDs []int32
				for _, r := range thread.replies {
					replyIDs = append(replyIDs, r.Id)
				}
				got[thread.comment.Id] = replyIDs
			}
			if !reflect.DeepEqual(order, tt.order) {
				t.Errorf("thread order = %v, want %v", order, tt.order)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threads = %v, want %v", got, tt.want)
			}
		})
	}
}
