package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/iohelper"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

type updatePullRequestOptions struct {
	idOptions
	title            string
	message          string
	markdownFilename string
	isAppend         bool
}

var updatePullRequestOpts updatePullRequestOptions

func (opts *updatePullRequestOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("invalid pull request ID")
	}
	if opts.message != "" && opts.markdownFilename != "" {
		return fmt.Errorf("message and file cannot be specified at the same time")
	}
	isDescriptionSpecified := opts.message != "" || opts.markdownFilename != ""
	if opts.title == "" && !isDescriptionSpecified {
		return fmt.Errorf("title, message or file must be specified")
	}
	if opts.isAppend && !isDescriptionSpecified {
		return fmt.Errorf("append requires message or file")
	}
	return nil
}

var updatePullRequestCmd = &cobra.Command{
	Use:     "pull-request",
	Aliases: []string{"pull-requests", "pr"},
	Short:   "Update the title or description of the specified pull request",
	RunE:    runUpdatePullRequest,
}

func init() {
	updateCmd.AddCommand(updatePullRequestCmd)

	flags := updatePullRequestCmd.Flags()
	flags.Int32VarP(&updatePullRequestOpts.id, "id", "i", 0, "Pull request ID")
	flags.StringVarP(&updatePullRequestOpts.title, "title", "t", "", "New title; unchanged if not specified")
	flags.StringVarP(&updatePullRequestOpts.message, "message", "m", "", "New description")
	flags.StringVarP(&updatePullRequestOpts.markdownFilename, "file", "f", "", "markdown file containing the new description")
	flags.BoolVarP(&updatePullRequestOpts.isAppend, "append", "a", false, "Append to the current description instead of replacing it")

	_ = updatePullRequestCmd.MarkFlagRequired("id")
}

func runUpdatePullRequest(_ *cobra.Command, _ []string) error {
	if err := updatePullRequestOpts.validate(); err != nil {
		return err
	}

	message := updatePullRequestOpts.message
	if updatePullRequestOpts.markdownFilename != "" {
		var errFile error
		message, errFile = iohelper.ReadStringFromFile(updatePullRequestOpts.markdownFilename)
		if errFile != nil {
			return errFile
		}
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	repo, err := getRepositoryInfoFromCurrentPath()
	if err != nil {
		return err
	}

	current, _, err := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet(
		auth,
		updatePullRequestOpts.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}

	title := current.Title
	if updatePullRequestOpts.title != "" {
		title = updatePullRequestOpts.title
	}
	currentDescription := ""
	if current.Summary != nil {
		currentDescription = current.Summary.Raw
	}
	description := currentDescription
	if updatePullRequestOpts.message != "" || updatePullRequestOpts.markdownFilename != "" {
		description = buildDescription(currentDescription, message, updatePullRequestOpts.isAppend)
	}

	pullRequest := swagger.Pullrequest{
		Title: title,
		Summary: &swagger.CommentContent{
			Raw: description,
		},
	}
	opts := &swagger.PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPutOpts{
		Body: optional.NewInterface(pullRequest),
	}
	updated, _, err := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut(
		auth,
		updatePullRequestOpts.id,
		repo.Name,
		repo.Org,
		opts,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Updated pull request [%d].\n", updatePullRequestOpts.id)
	models.PrintShortDescription(&updated, true)
	return nil
}

// buildDescription returns the description to be saved, either replacing the
// current one or appending to it with a blank line in between
func buildDescription(current string, message string, isAppend bool) string {
	current = strings.TrimRight(current, "\n")
	if !isAppend || current == "" {
		return message
	}
	return current + "\n\n" + message
}
