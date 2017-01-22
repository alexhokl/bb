package command

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// NewCheckoutCommand returns definition of command checkout
func NewCheckoutCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checkout",
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

	cmdName := "git"
	statusArgs := []string{"status", "-s"}
	statusOutput, errStatus := exec.Command(cmdName, statusArgs...).Output()
	if errStatus != nil {
		return errStatus
	}
	if len(statusOutput) > 0 {
		return errors.New("Working directory is not prestine. Please stash your work and try again.")
	}

	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	pr, err := client.GetRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}

	fetchArgs := []string{"fetch"}
	_, errFetch := exec.Command(cmdName, fetchArgs...).Output()
	if errFetch != nil {
		return errFetch
	}
	fmt.Println("Downloaded the latest information from BitBucket.")

	checkoutArgs := []string{"checkout", pr.Source.Branch.Name}
	_, errCheckout := exec.Command(cmdName, checkoutArgs...).Output()
	if errCheckout != nil {
		return errCheckout
	}
	fmt.Printf("Checked out branch %s.\n", pr.Source.Branch.Name)

	pullArgs := []string{"pull"}
	_, errPull := exec.Command(cmdName, pullArgs...).Output()
	if errPull != nil {
		return errPull
	}
	fmt.Printf("Pulled the latest code of branch %s.\n", pr.Source.Branch.Name)

	return nil
}
