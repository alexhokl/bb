package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

type listPullRequestOptions struct {
	isQuiet               bool
	isInDetail            bool
	isIncldeCreationTime  bool
	destinationBranchName string
	sourceBranchName      string
}

var listPullRequestOpts listPullRequestOptions

var listPullRequestCmd = &cobra.Command{
	Use:   "pr",
	Short: "List pull requests",
	RunE: runListPullRequest,
}

func init() {
	listCmd.AddCommand(listPullRequestCmd)

	flags := listCmd.Flags()
	flags.BoolVarP(&listPullRequestOpts.isQuiet, "quiet", "q", false, "List IDs only")
	flags.BoolVar(&listPullRequestOpts.isInDetail, "detail", false, "List in detail")
	flags.BoolVar(
		&listPullRequestOpts.isIncldeCreationTime, "created-time", false, "Include created time")
	flags.StringVarP(&listPullRequestOpts.destinationBranchName, "destination", "d", "", "Destination branch")
	flags.StringVarP(&listPullRequestOpts.sourceBranchName, "source", "s", "", "Source branch")
}

func runListPullRequest(_ *cobra.Command, _ []string) error {
	if err := validateListOptions(listPullRequestOpts); err != nil {
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

	opts := &swagger.PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsGetOpts{
		// possible values OPEN, MERGED, DECLINED, SUPERSEDED
		State: optional.NewString("OPEN"),
	}
	page, _, err := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsGet(
		auth,
		repo.Name,
		repo.Org,
		opts,
	)
	if err != nil {
		return err
	}

	var prList []swagger.Pullrequest


	prList = append(prList, page.Values...)

	// TODO: handle pagination
	// for page.Next != "" {
	// }

	if len(prList) == 0 {
		fmt.Println("There are no open pull requests.")
		return nil
	}

	for _, pr := range prList {
		if listPullRequestOpts.sourceBranchName != "" && pr.Source.Branch.Name != listPullRequestOpts.sourceBranchName {
			continue
		}
		if listPullRequestOpts.destinationBranchName != "" && pr.Destination.Branch.Name != listPullRequestOpts.destinationBranchName {
			continue
		}
		if listPullRequestOpts.isQuiet {
			models.PrintID(&pr)
		} else if listPullRequestOpts.isInDetail {
			models.PrintShortDescription(&pr, listPullRequestOpts.isIncldeCreationTime)
		} else {
			models.PrintOneLiner(&pr)
		}
	}
	return nil
}

func validateListOptions(opts listPullRequestOptions) error {
	if opts.sourceBranchName != "" && opts.destinationBranchName != "" && opts.sourceBranchName == opts.destinationBranchName {
		return errors.New("source branch cannot be same as destination branch")
	}
	return nil
}
