package command

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// NewUnapproveCommand returns definition of command unapprove
func NewUnapproveCommand(cli *ManagerCli) *cobra.Command {
	opts := idOption{}

	cmd := &cobra.Command{
		Use:   "unapprove",
		Short: "Un-approve the specified pull request",
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
			return runUnapprove(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")

	return cmd
}

func runUnapprove(cli *ManagerCli, opts idOption) error {
	if opts.id <= 0 {
		return errors.New("invalid pull request ID")
	}

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	err := client.UnapproveRequest(cred, repo, opts.id)
	if err != nil {
		return err
	}

	fmt.Printf("Un-approved pull request [%d].\n", opts.id)
	return nil
}
