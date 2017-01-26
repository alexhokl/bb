package command

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/spf13/cobra"
)

type approveOptions struct {
	idOption
	isKeepBranch bool
}

// NewApproveCommand returns definition of command approve
func NewApproveCommand(cli *ManagerCli) *cobra.Command {
	opts := approveOptions{}

	cmd := &cobra.Command{
		Use:   "approve",
		Short: "Approve the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runApprove(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&opts.isKeepBranch, "keep-branch", "k", false, "Keep local branch after approval")

	return cmd
}

func runApprove(cli *ManagerCli, opts approveOptions) error {
	if opts.id <= 0 {
		return errors.New("Invalid pull request ID")
	}

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	pr, errGet := client.GetRequest(cred, repo, opts.id)
	if errGet != nil {
		return errGet
	}

	err := client.ApproveRequest(cred, repo, opts.id)
	if err != nil {
		return err
	}
	fmt.Printf("Approved pull request [%d].\n", opts.id)

	if opts.isKeepBranch {
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
