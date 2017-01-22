package command

import (
	"fmt"
	"strconv"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/spf13/cobra"
)

type approveOptions struct {
	isKeepBranch bool
}

// NewApproveCommand returns definition of command approve
func NewApproveCommand(cli *ManagerCli) *cobra.Command {
	opts := approveOptions{}

	cmd := &cobra.Command{
		Use:   "approve [PR ID]",
		Short: "Approve a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runApprove(cli, args, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.isKeepBranch, "keep-branch", "k", false, "Keep local branch after approval")

	return cmd
}

func runApprove(cli *ManagerCli, args []string, opts approveOptions) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()
	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	pr, errGet := client.GetRequest(cred, repo, pullRequestNumber)
	if errGet != nil {
		return errGet
	}

	err := client.ApproveRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}
	fmt.Printf("Approved pull request [%d].\n", pullRequestNumber)

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
		fmt.Printf("Checked out to branch %s.", pr.Destination.Branch.Name)
	}

	_, errDelete := git.DeleteBranch(pr.Source.Branch.Name)
	if errDelete != nil {
		return errDelete
	}
	fmt.Printf("Deleted branch %s.", pr.Source.Branch.Name)

	return nil
}
