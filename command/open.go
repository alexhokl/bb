package command

import (
	"errors"
	"fmt"
	"os/exec"

	clihelper "github.com/alexhokl/helper/cli"
	"github.com/spf13/cobra"
)

// NewOpenCommand returns definition of command checkout
func NewOpenCommand(cli *ManagerCli) *cobra.Command {
	opts := idOption{}

	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open the web page of the specified pull request in a browser",
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
			return runOpen(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")

	return cmd
}

func runOpen(cli *ManagerCli, opts idOption) error {
	if opts.id <= 0 {
		return errors.New("Invalid pull request ID")
	}

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	pr, err := client.GetRequest(cred, repo, opts.id)
	if err != nil {
		return err
	}

	cmdName, cmdArgs := clihelper.GetOpenCommand(pr.Links.Html.Href)
	_, errOpen := exec.Command(cmdName, cmdArgs...).Output()
	if errOpen != nil {
		return errOpen
	}

	fmt.Printf("Opened %s in a browser.\n", pr.Links.Html.Href)
	return nil
}
