package command

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/spf13/cobra"
)

type checkoutOptions struct {
	idOption
	isShowDifftool           bool
	isShowStat               bool
	isNoMergeFromDestination bool
}

// NewCheckoutCommand returns definition of command checkout
func NewCheckoutCommand(cli *ManagerCli) *cobra.Command {
	opts := checkoutOptions{}

	cmd := &cobra.Command{
		Use:   "checkout",
		Short: "Checkout the latest code of the branch of the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			errRepo := cli.SetRepository()
			if errRepo != nil {
				return errRepo
			}
			errCred := cli.SetCredentials()
			if errCred != nil {
				return errCred
			}
			return runCheckout(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&opts.isShowDifftool, "difftool", "d", false, "Show difftool after checkout")
	flags.BoolVarP(&opts.isShowStat, "stat", "s", false, "Show diff stats after checkout")
	flags.BoolVar(&opts.isNoMergeFromDestination, "no-merge", false, "Do not merge from destination branch during checkout")

	return cmd
}

func runCheckout(cli *ManagerCli, opts checkoutOptions) error {
	if opts.id <= 0 {
		return errors.New("Invalid pull request ID")
	}

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	statusOutput, errStatus := git.GetStatus()
	if errStatus != nil {
		return errStatus
	}
	if len(statusOutput) > 0 {
		return errors.New("Working directory is not prestine. Please stash your work and try again")
	}

	pr, err := client.GetRequest(cred, repo, opts.id)
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

	if !opts.isNoMergeFromDestination {
		_, errMerge := git.Merge(pr.Destination.Branch.Name)
		if errMerge != nil {
			return errMerge
		}
		fmt.Printf("Merged from branch %s into %s.\n", pr.Destination.Branch.Name, pr.Source.Branch.Name)
	}

	if opts.isShowStat {
		stat, errStat := git.DiffStat(pr.Destination.Branch.Name, "")
		if errStat != nil {
			return errStat
		}
		fmt.Println(stat)
	}

	if opts.isShowDifftool {
		errDifftool := git.Difftool(pr.Destination.Branch.Name, "")
		if errDifftool != nil {
			return errDifftool
		}
	}

	return nil
}
