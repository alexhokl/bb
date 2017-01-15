package command

import (
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
	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	pr, err := client.GetRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}

	cmdName := "git"
	fetchArgs := []string{"fetch"}
	_, errFetch := exec.Command(cmdName, fetchArgs...).Output()
	if errFetch != nil {
		return errFetch
	}

	checkoutArgs := []string{"checkout", pr.Source.Branch.Name}
	_, errCheckout := exec.Command(cmdName, checkoutArgs...).Output()
	if errCheckout != nil {
		return errCheckout
	}

	pullArgs := []string{"pull"}
	_, errPull := exec.Command(cmdName, pullArgs...).Output()
	if errPull != nil {
		return errPull
	}

	return nil
}
