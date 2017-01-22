package command

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// NewDeclineCommand returns definition of command decline
func NewDeclineCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decline",
		Short: "Decline a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDecline(cli, args)
		},
	}
	return cmd
}

func runDecline(cli *ManagerCli, args []string) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()
	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	err := client.DeclineRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}

	fmt.Printf("Declined pull request [%d].\n", pullRequestNumber)
	return nil
}
