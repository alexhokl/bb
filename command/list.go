package command

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewListCommand returns definition of command list
func NewListCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List pull requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(cli)
		},
	}
	return cmd
}

func runList(cli *ManagerCli) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	prList, err := client.ListRequests(cred, repo)
	if err != nil {
		return err
	}

	for _, pr := range prList.Items {
		prInfo, _ := client.GetRequest(cred, repo, pr.ID)
		isApproved := prInfo.IsApproved(cred.Username)
		if isApproved {
			color.Cyan(pr.ToString())
		} else if pr.Author.Username == cred.Username {
			color.Blue(pr.ToString())
		} else {
			color.Red(pr.ToString())
		}
	}
	return nil
}
