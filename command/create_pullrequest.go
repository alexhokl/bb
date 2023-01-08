package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/git"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

type createPullRequestOptions struct {
	destinationBranchName string
	sourceBranchName      string
	title                 string
	message               string
}

var createPullRequestOpts createPullRequestOptions

func (opts createPullRequestOptions) Validate() error {
	if opts.destinationBranchName == "" {
		return fmt.Errorf("destination branch name is required")
	}
	return nil
}

var createPullRequestCmd = &cobra.Command{
	Use:   "pull-request",
	Aliases: []string{"pr"},
	Short: "Create the specified pull request",
	RunE:  runCreate,
}

func init() {
	rootCmd.AddCommand(createPullRequestCmd)

	flags := createPullRequestCmd.Flags()
	flags.StringVarP(&createPullRequestOpts.destinationBranchName, "destination", "d", "", "Destination branch")
	flags.StringVarP(&createPullRequestOpts.sourceBranchName, "source", "s", "", "Source branch; current branch if not specified")
	flags.StringVarP(&createPullRequestOpts.title, "title", "t", "", "Title; branch name if not specified")
	flags.StringVarP(&createPullRequestOpts.message, "message", "m", "", "Long description; auto-generated if not specified")
}

func runCreate(_ *cobra.Command, _ []string) error {
	if err := createPullRequestOpts.Validate(); err != nil {
		return err
	}

	_, errFetch := git.Fetch()
	if errFetch != nil {
		return errFetch
	}
	fmt.Println("Downloaded the latest information from BitBucket.")

	if createPullRequestOpts.sourceBranchName == "" {
		currentBranchName, errCurrentBranch := git.GetCurrentBranchName()
		if errCurrentBranch != nil {
			return errCurrentBranch
		}
		createPullRequestOpts.sourceBranchName = currentBranchName
	}

	if createPullRequestOpts.sourceBranchName == createPullRequestOpts.destinationBranchName {
		return errors.New("source branch cannot be same as destination branch")
	}

	if createPullRequestOpts.title == "" {
		createPullRequestOpts.title = createPullRequestOpts.sourceBranchName
	}

	if createPullRequestOpts.message == "" {
		msg, errMessage := git.GetBranchCommitComments(createPullRequestOpts.sourceBranchName, createPullRequestOpts.destinationBranchName)
		if errMessage != nil {
			return errMessage
		}
		createPullRequestOpts.message = msg
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

	pullRequest := swagger.Pullrequest{
		Title: createPullRequestOpts.title,
		Summary: &swagger.CommentContent{
			Raw: createPullRequestOpts.message,
		},
		Destination: &swagger.PullrequestEndpoint{
			Branch: &swagger.PullRequestBranch{
				Name: createPullRequestOpts.destinationBranchName,
			},
		},
		Source: &swagger.PullrequestEndpoint{
			Branch: &swagger.PullRequestBranch{
				Name: createPullRequestOpts.sourceBranchName,
			},
		},
	}

	opts := &swagger.PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPostOpts{
		Body: optional.NewInterface(pullRequest),
	}
	pr, _, err := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPost(
		auth,
		repo.Name,
		repo.Org,
		opts,
	)
	if err != nil {
		return err
	}

	fmt.Println("Pull request created.")
	models.PrintShortDescription(&pr, true)
	models.PrintDescription(&pr)

	return nil
}
