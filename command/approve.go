package command

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// NewApproveCommand returns definition of command approve
func NewApproveCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [PR ID]",
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

	err := client.ApproveRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}

	fmt.Printf("Approved pull request [%d].\n", pullRequestNumber)
	return nil
}
