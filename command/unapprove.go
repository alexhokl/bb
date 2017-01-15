package command

import (
	"strconv"

	"github.com/spf13/cobra"
)

// NewUnapproveCommand returns definition of command unapprove
func NewUnapproveCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unapprove",
		Short: "Unapprove a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUnapprove(cli, args)
		},
	}
	return cmd
}

func runUnapprove(cli *ManagerCli, args []string) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()
	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	return client.UnapproveRequest(cred, repo, pullRequestNumber)
}
