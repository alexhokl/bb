package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type listProjectsOptions struct {
	workspace string
}

var listProjectsOpts listProjectsOptions

func (opts *listProjectsOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	return nil
}

var listProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"project"},
	Short:   "List projects in a workspace",
	RunE:    runListProjects,
}

func init() {
	listCmd.AddCommand(listProjectsCmd)
	flags := listProjectsCmd.Flags()
	flags.StringVarP(&listProjectsOpts.workspace, "workspace", "w", "", "Workspace slug")
	_ = listProjectsCmd.MarkFlagRequired("workspace")
}

func runListProjects(_ *cobra.Command, _ []string) error {
	if err := listProjectsOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	page, _, err := client.WorkspacesApi.WorkspacesWorkspaceProjectsGet(auth, listProjectsOpts.workspace)
	if err != nil {
		return err
	}

	for _, p := range page.Values {
		fmt.Printf("%s\t%s\n", p.Key, p.Name)
	}
	return nil
}
