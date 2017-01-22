package command

import (
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listOptions struct {
	isQuiet bool
}

type print func(msg string, args ...interface{})

// NewListCommand returns definition of command list
func NewListCommand(cli *ManagerCli) *cobra.Command {
	opts := listOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List pull requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.isQuiet, "quiet", "q", false, "List IDs only")

	return cmd
}

func runList(cli *ManagerCli, opts listOptions) error {
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
		printFunc := getPrint(prInfo, cred)
		if opts.isQuiet {
			printFunc("%d", prInfo.ID)
		} else {
			printFunc(prInfo.ToShortDescription())
		}
	}
	return nil
}

func getPrint(pr *models.PullRequestDetail, cred *models.UserCredential) print {
	if pr.IsApproved(cred.Username) {
		return color.Cyan
	}
	if pr.Author.Username == cred.Username {
		return color.Blue
	}
	return color.Red
}
