package command

import (
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

// NewOpenCommand returns definition of command checkout
func NewOpenCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open the web page of the specified pull request in a browser",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runOpen(cli, args)
		},
	}
	return cmd
}

func runOpen(cli *ManagerCli, args []string) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()
	pullRequestNumber, errParse := strconv.Atoi(args[0])
	if errParse != nil {
		return errParse
	}

	pr, err := client.GetRequest(cred, repo, pullRequestNumber)
	if err != nil {
		return err
	}

	cmdName := "open"
	cmdArgs := []string{pr.Links.Html.Href}
	if runtime.GOOS == "windows" {
		cmdName = "cmd"
		cmdArgs = []string{"/C", "start", pr.Links.Html.Href}
	}
	_, errOpen := exec.Command(cmdName, cmdArgs...).Output()
	if errOpen != nil {
		return errOpen
	}

	return nil
}
