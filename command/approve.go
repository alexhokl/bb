package command

import (
	"strconv"

	"github.com/spf13/cobra"
)

// NewApproveCommand returns definition of command approve
func NewApproveCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "Approve a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runApprove(cli, args)
		},
	}
	return cmd
}

func runApprove(cli *ManagerCli, args []string) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()
	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	return client.ApproveRequest(cred, repo, pullRequestNumber)
}
