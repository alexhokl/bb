package command

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/alexhokl/bb/api"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type listWorkspaceMembersOptions struct {
	isQuiet bool
}

// NewListWorkspaceDefaultReviewersCommand returns definition of command list
func NewListWorkspaceMembersCommand(cli *ManagerCli) *cobra.Command {
	opts := listWorkspaceMembersOptions{}

	cmd := &cobra.Command{
		Use:   "list-workspace-members",
		Short: "List default reviewers of a workspace",
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
			return runListWorkspaceMembers(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.isQuiet, "quiet", "q", false, "UUIDs only")

	return cmd
}

func runListWorkspaceMembers(cli *ManagerCli, opts listWorkspaceMembersOptions) error {
	cred := cli.UserCredential()
	auth := context.WithValue(context.Background(), api.ContextAccessToken, cred.AccessToken)
	config := api.NewConfiguration()
	apiClient := api.NewAPIClient(config)
	repo := cli.Repo()

	list, _, err := apiClient.WorkspacesApi.WorkspacesWorkspaceMembersGet(auth, repo.Org)
	if err != nil {
		return err
	}

	var members []api.WorkspaceMembership
	// for len(list.Values) != 0 {
	members = append(members, list.Values...)
	// }

	if opts.isQuiet {
		for _, m := range members {
			fmt.Println(strings.ReplaceAll(strings.ReplaceAll(m.User.Uuid, "{", ""), "}", ""))
		}
		return nil
	}

	var data [][]string
	for _, m := range members {
		arr := []string{
			strings.ReplaceAll(strings.ReplaceAll(m.User.Uuid, "{", ""), "}", ""),
			m.User.DisplayName,
		}
		data = append(data, arr)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Display Name"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()

	return nil
}
