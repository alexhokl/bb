package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/git"
	"github.com/spf13/cobra"
)

type checkoutOptions struct {
	idOptions
	isShowDifftool           bool
	isShowStat               bool
	isNoMergeFromDestination bool
}

var checkoutOpts checkoutOptions

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout the latest code of the branch of the specified pull request",
	RunE: runCheckout,
}

func init() {
	rootCmd.AddCommand(checkoutCmd)

	flags := checkoutCmd.Flags()
	flags.Int32VarP(&checkoutOpts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&checkoutOpts.isShowDifftool, "difftool", "d", false, "Show difftool after checkout")
	flags.BoolVarP(&checkoutOpts.isShowStat, "stat", "s", false, "Show diff stats after checkout")
	flags.BoolVar(&checkoutOpts.isNoMergeFromDestination, "no-merge", false, "Do not merge from destination branch during checkout")
}

func runCheckout(_ *cobra.Command, _ []string) error {
	if err := validateIDOptions(checkoutOpts.idOptions); err != nil {
		return err
	}

	statusOutput, errStatus := git.GetStatus()
	if errStatus != nil {
		return errStatus
	}
	if len(statusOutput) > 0 {
		return errors.New("working directory is not prestine. Please stash your work and try again")
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
		checkoutOpts.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}

	_, errFetch := git.Fetch()
	if errFetch != nil {
		return errFetch
	}
	fmt.Println("Downloaded the latest information from BitBucket.")

	currentBranchName, errCurrent := git.GetCurrentBranchName()
	if errCurrent != nil {
		return errCurrent
	}

	if currentBranchName != pr.Destination.Branch.Name {
		_, errCheckoutDestination := git.Checkout(pr.Destination.Branch.Name)
		if errCheckoutDestination != nil {
			return errCheckoutDestination
		}
		fmt.Printf("Checked out branch %s.\n", pr.Destination.Branch.Name)
	}

	_, errPullDestination := git.Pull()
	if errPullDestination != nil {
		return errPullDestination
	}
	fmt.Printf("Pulled the latest code of branch %s.\n", pr.Destination.Branch.Name)

	_, errCheckout := git.Checkout(pr.Source.Branch.Name)
	if errCheckout != nil {
		return errCheckout
	}
	fmt.Printf("Checked out branch %s.\n", pr.Source.Branch.Name)

	_, errPull := git.Pull()
	if errPull != nil {
		return errPull
	}
	fmt.Printf("Pulled the latest code of branch %s.\n", pr.Source.Branch.Name)

	if !checkoutOpts.isNoMergeFromDestination {
		_, errMerge := git.Merge(pr.Destination.Branch.Name)
		if errMerge != nil {
			return errMerge
		}
		fmt.Printf("Merged from branch %s into %s.\n", pr.Destination.Branch.Name, pr.Source.Branch.Name)
	}

	if checkoutOpts.isShowStat {
		stat, errStat := git.DiffStat(pr.Destination.Branch.Name, "")
		if errStat != nil {
			return errStat
		}
		fmt.Println(stat)
	}

	if checkoutOpts.isShowDifftool {
		errDifftool := git.Difftool(pr.Destination.Branch.Name, "")
		if errDifftool != nil {
			return errDifftool
		}
	}

	return nil
}
