package command

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/git"
	"github.com/alexhokl/go-bb-pr/models"
	"github.com/spf13/cobra"
)

type createOptions struct {
	destinationBranchName string
	sourceBranchName      string
	title                 string
	message               string
}

// NewCreateCommand returns definition of command merge
func NewCreateCommand(cli *ManagerCli) *cobra.Command {
	opts := createOptions{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create the specified pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runCreate(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.destinationBranchName, "destination", "d", "", "Destination branch")
	flags.StringVarP(&opts.sourceBranchName, "source", "s", "", "Source branch; current branch if not specified")
	flags.StringVarP(&opts.title, "title", "t", "", "Title; branch name if not specified")
	flags.StringVarP(&opts.message, "message", "m", "", "Long description; auto-generated if not specified")

	return cmd
}

func runCreate(cli *ManagerCli, opts createOptions) error {
	if opts.destinationBranchName == "" {
		return errors.New("Destination branch is not specified")
	}

	_, errFetch := git.Fetch()
	if errFetch != nil {
		return errFetch
	}
	fmt.Println("Downloaded the latest information from BitBucket.")

	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	if opts.sourceBranchName == "" {
		currentBranchName, errCurrentBranch := git.GetCurrentBranchName()
		if errCurrentBranch != nil {
			return errCurrentBranch
		}
		opts.sourceBranchName = currentBranchName
	}

	if opts.sourceBranchName == opts.destinationBranchName {
		return errors.New("Source branch cannot be same as destination branch")
	}

	if opts.title == "" {
		opts.title = opts.sourceBranchName
	}

	if opts.message == "" {
		msg, errMessage := git.GetBranchCommitComments(opts.sourceBranchName, opts.destinationBranchName)
		if errMessage != nil {
			return errMessage
		}
		opts.message = msg
	}

	prRequest := &models.PullRequestCreateRequest{
		Destination: models.CommitBranch{Branch: models.Branch{Name: opts.destinationBranchName}},
		Source:      models.CommitBranch{Branch: models.Branch{Name: opts.sourceBranchName}},
		Title:       opts.title,
		Description: opts.message,
	}

	pr, errCreate := client.CreateRequest(cred, repo, prRequest)
	if errCreate != nil {
		return errCreate
	}
	fmt.Printf("Pull request created.\n%s", pr.ToShortDescription(true))

	return nil
}
