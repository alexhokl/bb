package command

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
	"github.com/spf13/cobra"
)

// NewDescribeCommand returns definition of command describe
func NewDescribeCommand(cli *ManagerCli) *cobra.Command {
	opts := idOption{}

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runDescribe(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")

	return cmd
}

func runDescribe(cli *ManagerCli, opts idOption) error {
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

	activities, errActivities := client.ActivityRequest(cred, repo, opts.id)
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
