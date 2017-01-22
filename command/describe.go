package command

import (
	"fmt"
	"strconv"

	"github.com/alexhokl/go-bb-pr/models"
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

	activities, errActivities := client.ActivityRequest(cred, repo, pullRequestNumber)
	if errActivities != nil {
		return errActivities
	}

	printFunc := getPrint(pr, cred)
	printFunc(pr.ToString())

	fmt.Println("")

	for _, reviewer := range pr.Participants {
		if reviewer.Approved {
			fmt.Printf("Approved by %s\n", reviewer.User.DisplayName)
		}
	}

	fmt.Println("")

	for _, event := range activities {
		if event.Comment != (models.Comment{}) {
			fmt.Println(event.Comment.ToString())
		}
		if event.Update != (models.Update{}) {
			fmt.Println(event.Update.ToString())
		}
	}

	return nil
}
