package command

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"

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

	cmdName := "open"
	cmdArgs := []string{pr.Links.Html.Href}
	if runtime.GOOS == "windows" {
		cmdName = "cmd"
		cmdArgs = []string{"/C", "start", pr.Links.Html.Href}
	}
	if runtime.GOOS == "linux" {
		cmdName = "xdg-open"
	}
	_, errOpen := exec.Command(cmdName, cmdArgs...).Output()
	if errOpen != nil {
		return errOpen
	}

	fmt.Printf("Opened %s in a browser.\n", pr.Links.Html.Href)
	return nil
}
