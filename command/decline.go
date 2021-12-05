package command

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// NewDeclineCommand returns definition of command decline
func NewDeclineCommand(cli *ManagerCli) *cobra.Command {
	opts := commentOptions{}

	cmd := &cobra.Command{
		Use:   "decline",
		Short: "Decline the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			errRepo := cli.SetRepository()
			if errRepo != nil {
				return errRepo
			}
			errCred := cli.SetCredentials()
			if errCred != nil {
				return errCred
			}
			return runDecline(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")
	flags.StringVarP(&opts.message, "message", "m", "", "comment message")

	return cmd
}

func runDecline(cli *ManagerCli, opts commentOptions) error {
	if opts.id <= 0 {
		return errors.New("invalid pull request ID")
	}

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	err := client.DeclineRequest(cred, repo, opts.id)
	if err != nil {
		return err
	}

	fmt.Printf("Declined pull request [%d].\n", opts.id)

	if opts.message != "" {
		err := client.AddComment(cred, repo, opts.id, opts.message)
		if err != nil {
			return err
		}
		fmt.Printf("Added comment to pull request [%d].\n", opts.id)
	}

	return nil
}
