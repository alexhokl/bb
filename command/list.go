package command

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type listOptions struct {
	isQuiet bool
}

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
		isApproved := prInfo.IsApproved(cred.Username)
		if isApproved {
			if opts.isQuiet {
				color.Cyan(fmt.Sprintf("%d", prInfo.ID))
			} else {
				color.Cyan(prInfo.ToShortDescription())
			}
		} else if pr.Author.Username == cred.Username {
			if opts.isQuiet {
				color.Blue(fmt.Sprintf("%d", prInfo.ID))
			} else {
				color.Blue(prInfo.ToShortDescription())
			}
		} else {
			if opts.isQuiet {
				color.Red(fmt.Sprintf("%d", prInfo.ID))
			} else {
				color.Red(prInfo.ToShortDescription())
			}
		}
	}
	return nil
}
