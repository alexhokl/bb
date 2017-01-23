package command

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// NewDeclineCommand returns definition of command decline
func NewDeclineCommand(cli *ManagerCli) *cobra.Command {
	opts := idOption{}

	cmd := &cobra.Command{
		Use:   "decline",
		Short: "Decline the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runDecline(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")

	return cmd
}

func runDecline(cli *ManagerCli, opts idOption) error {
	if opts.id <= 0 {
		return errors.New("Invalid pull request ID")
	}

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	err := client.DeclineRequest(cred, repo, opts.id)
	if err != nil {
		return err
	}

	fmt.Printf("Declined pull request [%d].\n", opts.id)
	return nil
}
