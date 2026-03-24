package command

import (
	"fmt"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

type listUsersOptions struct {
	workspace string
}

var listUsersOpts listUsersOptions

func (opts *listUsersOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	return nil
}

var listUsersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"user"},
	Short:   "List users of a workspace",
	RunE:    runListUsers,
}

func init() {
	listCmd.AddCommand(listUsersCmd)
	flags := listUsersCmd.Flags()
	flags.StringVarP(&listUsersOpts.workspace, "workspace", "w", "", "Workspace slug or UUID")
	_ = listUsersCmd.MarkFlagRequired("workspace")
}

func runListUsers(_ *cobra.Command, _ []string) error {
	if err := listUsersOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	page, _, err := client.WorkspacesApi.WorkspacesWorkspaceMembersGet(auth, listUsersOpts.workspace)
	if err != nil {
		return err
	}

	for i := range page.Values {
		models.PrintWorkspaceMembership(&page.Values[i])
	}
	return nil
}
