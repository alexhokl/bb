package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/git"
	"github.com/spf13/cobra"
)

type approveOptions struct {
	idOptions
	isKeepBranch bool
}

var approveOpts approveOptions

func (opts *approveOptions) validate() error {
	return opts.idOptions.validate()
}

var approveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve the specified pull request",
	RunE:  runApprove,
}

func init() {
	rootCmd.AddCommand(approveCmd)

	flags := approveCmd.Flags()
	flags.Int32VarP(&approveOpts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&approveOpts.isKeepBranch, "keep-branch", "k", false, "Keep local branch after approval")

	approveCmd.MarkFlagRequired("id")
}

func runApprove(_ *cobra.Command, _ []string) error {
	if err := approveOpts.validate(); err != nil {
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

	pr, _, err := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet(
		auth,
		idOpts.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}


	_, _, err = client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost(
		auth,
		approveOpts.idOptions.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Approved pull request [%d].\n", idOpts.id)

	if approveOpts.isKeepBranch {
		return nil
	}

	currentBranchName, errCurrentBranch := git.GetCurrentBranchName()
	if errCurrentBranch != nil {
		return errCurrentBranch
	}

	if currentBranchName == pr.Source.Branch.Name {
		_, errCheckout := git.Checkout(pr.Destination.Branch.Name)
		if errCheckout != nil {
			return errCheckout
		}
		fmt.Printf("Checked out to branch %s.\n", pr.Destination.Branch.Name)
	}

	isBranchExists, errBranchCheck := git.IsBranchExists(pr.Source.Branch.Name)
	if errBranchCheck != nil {
		return errCurrentBranch
	}

	if isBranchExists {
		_, errDelete := git.DeleteBranch(pr.Source.Branch.Name)
		if errDelete != nil {
			return errDelete
		}
		fmt.Printf("Deleted branch %s.\n", pr.Source.Branch.Name)
	}

	return nil
}
