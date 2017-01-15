package command

import (
	"strconv"

	"github.com/spf13/cobra"
)

// NewMergeCommand returns definition of command merge
func NewMergeCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "merge",
		Short: "Merge a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMerge(cli, args)
		},
	}
	return cmd
}

func runMerge(cli *ManagerCli, args []string) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()
	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	return client.MergeRequest(cred, repo, pullRequestNumber)
}
