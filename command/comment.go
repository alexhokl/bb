package command

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

type commentOptions struct {
	idOption
	message string
}

// NewCommentCommand returns definition of command comment
func NewCommentCommand(cli *ManagerCli) *cobra.Command {
	opts := commentOptions{}

	cmd := &cobra.Command{
		Use:   "comment",
		Short: "Add comment to the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runComment(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")
	flags.StringVarP(&opts.message, "message", "m", "", "comment message")

	return cmd
}

func runComment(cli *ManagerCli, opts commentOptions) error {
	if opts.id <= 0 {
		return errors.New("Invalid pull request ID")
	}
	if opts.message == "" {
		return errors.New("Message cannot be empty")
	}

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	err := client.AddComment(cred, repo, opts.id, opts.message)
	if err != nil {
		return err
	}

	fmt.Printf("Added comment to pull request [%d].\n", opts.id)
	return nil
}
