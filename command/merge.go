package command

import (
	"fmt"
	"strconv"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/spf13/cobra"
)

type mergeOptions struct {
	isKeepBranch bool
}

// NewMergeCommand returns definition of command merge
func NewMergeCommand(cli *ManagerCli) *cobra.Command {
	opts := mergeOptions{}

	cmd := &cobra.Command{
		Use:   "merge [PR ID]",
		Short: "Merge a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMerge(cli, args, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.isKeepBranch, "keep-branch", "k", false, "Keep local branch after merging")

	return cmd
}

func runMerge(cli *ManagerCli, args []string, opts mergeOptions) error {
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

	err := client.MergeRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}
	fmt.Printf("Merged pull request [%d].\n", pullRequestNumber)

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
