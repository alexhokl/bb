package command

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/spf13/cobra"
)

// NewCheckoutCommand returns definition of command checkout
func NewCheckoutCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checkout [PR ID]",
		Short: "Checkout the latest code of the branch of a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCheckout(cli, args)
		},
	}
	return cmd
}

func runCheckout(cli *ManagerCli, args []string) error {
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

	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	pr, err := client.GetRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}

	_, errFetch := git.Fetch()
	if errFetch != nil {
		return errFetch
	}
	fmt.Println("Downloaded the latest information from BitBucket.")

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

	return nil
}
