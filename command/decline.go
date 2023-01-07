package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type declineOptions struct {
	idOptions
	message string
}

var declineOpts declineOptions

func (opts *declineOptions) validate() error {
	return opts.idOptions.validate()
}

var declineCmd = &cobra.Command{
	Use:   "decline",
	Short: "Decline the specified pull request",
	RunE:  runDecline,
}

func init() {
	rootCmd.AddCommand(declineCmd)

	flags := declineCmd.Flags()
	flags.Int32VarP(&declineOpts.id, "id", "i", 0, "Pull request ID")
	flags.StringVarP(&declineOpts.message, "message", "m", "", "comment message")

	declineCmd.MarkFlagRequired("id")
}

func runDecline(_ *cobra.Command, _ []string) error {
	if err := declineOpts.validate(); err != nil {
		return err
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

	_, _, err = client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost(
		auth,
		declineOpts.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Declined pull request [%d].\n", declineOpts.id)

	if declineOpts.message != "" {
		opts := swagger.PullrequestComment{
			Content: &swagger.CommentContent{
				Raw: declineOpts.message,
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
		fmt.Printf("Added comment to pull request [%d].\n", declineOpts.id)
	}

	return nil
}
