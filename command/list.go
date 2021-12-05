package command

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

type listOptions struct {
	isQuiet               bool
	isInDetail            bool
	isIncldeCreationTime  bool
	destinationBranchName string
	sourceBranchName      string
}

// NewListCommand returns definition of command list
func NewListCommand(cli *ManagerCli) *cobra.Command {
	opts := listOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List pull requests",
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
			return runList(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.isQuiet, "quiet", "q", false, "List IDs only")
	flags.BoolVar(&opts.isInDetail, "detail", false, "List in detail")
	flags.BoolVar(
		&opts.isIncldeCreationTime, "created-time", false, "Include created time")
	flags.StringVarP(&opts.destinationBranchName, "destination", "d", "", "Destination branch")
	flags.StringVarP(&opts.sourceBranchName, "source", "s", "", "Source branch")

	return cmd
}

func runList(cli *ManagerCli, opts listOptions) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	if opts.sourceBranchName != "" && opts.destinationBranchName != "" && opts.sourceBranchName == opts.destinationBranchName {
		return errors.New("source branch cannot be same as destination branch")
	}

	prList, err := client.ListRequests(cred, repo)
	if err != nil {
		return err
	}

	if len(prList) == 0 {
		fmt.Println("There are no open pull requests.")
		return nil
	}

	for _, pr := range prList {
		if opts.sourceBranchName != "" && pr.Source.Branch.Name != opts.sourceBranchName {
			continue
		}
		if opts.destinationBranchName != "" && pr.Destination.Branch.Name != opts.destinationBranchName {
			continue
		}
		prInfo, _ := client.GetRequest(cred, repo, pr.ID)
		if opts.isQuiet {
			prInfo.PrintID()
		} else if opts.isInDetail {
			prInfo.PrintShortDescription(opts.isIncldeCreationTime)
		} else {
			prInfo.PrintOneLiner()
		}
	}
	return nil
}
