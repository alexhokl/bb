package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/iohelper"
	"github.com/spf13/cobra"
)

type commentOptions struct {
	idOptions
	message          string
	markdownFilename string
}

var commentOpts commentOptions

func (opts *commentOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("invalid pull request ID")
	}
	if opts.message == "" && opts.markdownFilename == "" {
		return fmt.Errorf("message or file must be specified")
	}
	if opts.message != "" && opts.markdownFilename != "" {
		return fmt.Errorf("message and file cannot be specified at the same time")
	}
	return nil
}

var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add comment to the specified pull request",
	RunE:  runComment,
}

func init() {
	rootCmd.AddCommand(commentCmd)

	flags := commentCmd.Flags()
	flags.Int32VarP(&commentOpts.id, "id", "i", 0, "Pull request ID")
	flags.StringVarP(&commentOpts.message, "message", "m", "", "comment message")
	flags.StringVarP(&commentOpts.markdownFilename, "file", "f", "", "markdown file")

	commentCmd.MarkFlagRequired("id")
}

func runComment(_ *cobra.Command, _ []string) error {
	if err := commentOpts.validate(); err != nil {
		return err
	}

	message := commentOpts.message
	if commentOpts.markdownFilename != "" {
		var errFile error
		message, errFile = iohelper.ReadStringFromFile(commentOpts.markdownFilename)
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

	opts := swagger.PullrequestComment{
		Content: &swagger.CommentContent{
			Raw: message,
		},
	}
	_, _, err = client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost(
		auth,
		commentOpts.id,
		repo.Name,
		repo.Org,
		opts,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Added comment to pull request [%d].\n", commentOpts.id)
	return nil
}
