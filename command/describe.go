package command

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewDescribeCommand returns definition of command describe
func NewDescribeCommand(cli *ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDescribe(cli, args)
		},
	}
	return cmd
}

func runDescribe(cli *ManagerCli, args []string) error {
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

	isApproved := pr.IsApproved(cred.Username)
	if isApproved {
		color.Cyan(pr.ToString())
	} else if pr.Author.Username == cred.Username {
		color.Blue(pr.ToString())
	} else {
		color.Red(pr.ToString())
	}

	for _, reviewer := range pr.Participants {
		if reviewer.Approved {
			fmt.Printf("Approved by %s\n", reviewer.User.DisplayName)
		}
	}

	return nil
}
