package command

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/alexhokl/go-bb-pr/models"
	"github.com/spf13/cobra"
)

type describeOptions struct {
	idOption
	isShowDifftool bool
}

// NewDescribeCommand returns definition of command describe
func NewDescribeCommand(cli *ManagerCli) *cobra.Command {
	opts := describeOptions{}

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe the specified pull request",
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
			return runDescribe(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&opts.isShowDifftool, "difftool", "d", false, "Open difftool after description")

	return cmd
}

func runDescribe(cli *ManagerCli, opts describeOptions) error {
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

	pr.PrintShortDescription(true)
	pr.PrintDescription()

	fmt.Println("---")

	for _, event := range activities {
		if event.Comment != (models.Comment{}) {
			fmt.Println(event.Comment.ToString())
		}
		if event.Update != (models.Update{}) {
			fmt.Println(event.Update.ToString())
		}
	}

	stat, errStat := git.DiffStat(pr.Destination.Branch.Name, pr.Source.Branch.Name)
	if errStat != nil {
		return errStat
	}
	fmt.Println("Diff:")
	fmt.Println(stat)

	if opts.isShowDifftool {
		errDifftool := git.Difftool(pr.Destination.Branch.Name, pr.Source.Branch.Name)
		if errDifftool != nil {
			return errDifftool
		}
	}

	return nil
}
