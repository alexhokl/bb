package command

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/spf13/cobra"
)

type mergeOptions struct {
	idOption
	isKeepBranch bool
}

// NewMergeCommand returns definition of command merge
func NewMergeCommand(cli *ManagerCli) *cobra.Command {
	opts := mergeOptions{}

	cmd := &cobra.Command{
		Use:   "merge",
		Short: "Merge the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runMerge(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&opts.isKeepBranch, "keep-branch", "k", false, "Keep local branch after merging")

	return cmd
}

func runMerge(cli *ManagerCli, opts mergeOptions) error {
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

	err := client.MergeRequest(cred, repo, opts.id)
	if err != nil {
		return err
	}
	fmt.Printf("Merged pull request [%d].\n", opts.id)

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
