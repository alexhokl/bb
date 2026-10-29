package command

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/datetime"
	"github.com/spf13/cobra"

	"github.com/alexhokl/bb/swagger"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-get

const (
	commentFormatText = "text"
	commentFormatJSON = "json"
)

var validCommentFormats = []string{commentFormatText, commentFormatJSON}

type listCommentsOptions struct {
	idOptions
	format string
}

var listCommentsOpts listCommentsOptions

func (opts *listCommentsOptions) validate() error {
	if err := opts.idOptions.validate(); err != nil {
		return err
	}
	if opts.format != commentFormatText && opts.format != commentFormatJSON {
		return fmt.Errorf("format must be one of: %s", strings.Join(validCommentFormats, ", "))
	}
	return nil
}

// commentThread is a top-level comment together with all of its replies
type commentThread struct {
	comment swagger.PullrequestComment
	replies []swagger.PullrequestComment
}

var listCommentsCmd = &cobra.Command{
	Use:     "comments",
	Aliases: []string{"comment"},
	Short:   "List comments of the specified pull request",
	RunE:    runListComments,
}

func init() {
	listCmd.AddCommand(listCommentsCmd)
	flags := listCommentsCmd.Flags()
	flags.Int32VarP(&listCommentsOpts.id, "pr-id", "i", 0, "Pull request ID")
	flags.StringVarP(&listCommentsOpts.format, "format", "f", commentFormatText, "Output format (text or json)")

	_ = listCommentsCmd.MarkFlagRequired("pr-id")
}

func runListComments(_ *cobra.Command, _ []string) error {
	if err := listCommentsOpts.validate(); err != nil {
		return err
	}

	repo, err := getRepositoryInfoFromCurrentPath()
	if err != nil {
		return err
	}

	comments, err := getPullRequestComments(repo.Org, repo.Name, listCommentsOpts.id)
	if err != nil {
		return err
	}

	if listCommentsOpts.format == commentFormatJSON {
		if comments == nil {
			comments = []swagger.PullrequestComment{}
		}
		output, err := json.MarshalIndent(comments, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal comments: %w", err)
		}
		fmt.Println(string(output))
		return nil
	}

	if len(comments) == 0 {
		fmt.Printf("There are no comments on pull request [%d].\n", listCommentsOpts.id)
		return nil
	}

	for i, thread := range buildCommentThreads(comments) {
		if i > 0 {
			fmt.Println()
		}
		printComment(&thread.comment, "")
		for j := range thread.replies {
			printComment(&thread.replies[j], "    ")
		}
	}

	return nil
}

func getPullRequestComments(workspace, slug string, id int32) ([]swagger.PullrequestComment, error) {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return nil, err
	}

	nextURL := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/pullrequests/%d/comments?pagelen=100",
		workspace,
		slug,
		id,
	)

	var comments []swagger.PullrequestComment
	for nextURL != "" {
		req, err := http.NewRequest(http.MethodGet, nextURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to call pull request comments API: %w", err)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("pull request comments API returned status %d: %s", resp.StatusCode, string(body))
		}

		var page swagger.PaginatedPullrequestComments
		if err := json.Unmarshal(body, &page); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}
		comments = append(comments, page.Values...)
		nextURL = page.Next
	}

	return comments, nil
}

// buildCommentThreads groups comments into threads ordered by creation time.
// Replies (including replies to replies) are placed under their top-level
// comment; a reply whose parent is not in the list becomes a top-level comment.
func buildCommentThreads(comments []swagger.PullrequestComment) []commentThread {
	sorted := make([]swagger.PullrequestComment, len(comments))
	copy(sorted, comments)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].CreatedOn.Before(sorted[j].CreatedOn)
	})

	byID := make(map[int32]*swagger.PullrequestComment, len(sorted))
	for i := range sorted {
		byID[sorted[i].Id] = &sorted[i]
	}

	rootID := func(c *swagger.PullrequestComment) int32 {
		visited := map[int32]bool{c.Id: true}
		for c.Parent != nil {
			parent, ok := byID[c.Parent.Id]
			if !ok || visited[parent.Id] {
				break
			}
			visited[parent.Id] = true
			c = parent
		}
		return c.Id
	}

	var threads []commentThread
	threadIndex := make(map[int32]int)
	var replies []*swagger.PullrequestComment
	for i := range sorted {
		c := &sorted[i]
		if rootID(c) == c.Id {
			threadIndex[c.Id] = len(threads)
			threads = append(threads, commentThread{comment: *c})
		} else {
			replies = append(replies, c)
		}
	}
	for _, c := range replies {
		idx := threadIndex[rootID(c)]
		threads[idx].replies = append(threads[idx].replies, *c)
	}

	return threads
}

// formatCommentLocation returns the file and line an inline comment is
// anchored to, or an empty string for a general comment
func formatCommentLocation(inline *swagger.CommentInline) string {
	if inline == nil || inline.Path == "" {
		return ""
	}
	if inline.To > 0 {
		return fmt.Sprintf("%s:%d", inline.Path, inline.To)
	}
	if inline.From > 0 {
		return fmt.Sprintf("%s:%d (old)", inline.Path, inline.From)
	}
	return inline.Path
}

func printComment(c *swagger.PullrequestComment, indent string) {
	author := "unknown"
	if c.User != nil && c.User.DisplayName != "" {
		author = c.User.DisplayName
	}

	header := fmt.Sprintf("%s#%d %s (%s)", indent, c.Id, author, datetime.GetLocalDateTimeString(&c.CreatedOn))
	if location := formatCommentLocation(c.Inline); location != "" {
		header += fmt.Sprintf(" [%s]", location)
	}
	fmt.Println(header)

	content := ""
	if c.Content != nil {
		content = c.Content.Raw
	}
	if c.Deleted {
		content = "(deleted)"
	}
	for _, line := range strings.Split(strings.TrimRight(content, "\n"), "\n") {
		fmt.Printf("%s    %s\n", indent, line)
	}
}
