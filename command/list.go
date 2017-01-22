package command

import (
	"fmt"

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

	if len(prList) == 0 {
		fmt.Println("There are no open pull requests.")
		return nil
	}

	for _, pr := range prList {
		prInfo, _ := client.GetRequest(cred, repo, pr.ID)
		isApproved := prInfo.IsApproved(cred.Username)
		if isApproved {
			color.Cyan(prInfo.ToShortDescription())
		} else if pr.Author.Username == cred.Username {
			color.Blue(prInfo.ToShortDescription())
		} else {
			color.Red(prInfo.ToShortDescription())
		}
	}
	return nil
}
