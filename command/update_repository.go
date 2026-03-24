package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

type updateRepositoryOptions struct {
	workspace  string
	slug       string
	projectKey string
}

var updateRepositoryOpts updateRepositoryOptions

func (opts *updateRepositoryOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.slug == "" {
		return fmt.Errorf("slug is required")
	}
	if opts.projectKey == "" {
		return fmt.Errorf("at least one field to update must be specified")
	}
	return nil
}

var updateRepositoryCmd = &cobra.Command{
	Use:     "repository",
	Aliases: []string{"repo"},
	Short:   "Update a repository",
	RunE:    runUpdateRepository,
}

func init() {
	updateCmd.AddCommand(updateRepositoryCmd)

	flags := updateRepositoryCmd.Flags()
	flags.StringVarP(&updateRepositoryOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&updateRepositoryOpts.slug, "slug", "s", "", "Repository slug")
	flags.StringVarP(&updateRepositoryOpts.projectKey, "project-key", "p", "", "Key of the project to associate the repository with")

	_ = updateRepositoryCmd.MarkFlagRequired("workspace")
	_ = updateRepositoryCmd.MarkFlagRequired("slug")
}

func runUpdateRepository(_ *cobra.Command, _ []string) error {
	if err := updateRepositoryOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	repo := swagger.Repository{
		Type_:   "repository",
		Project: &swagger.Project{Key: updateRepositoryOpts.projectKey},
	}

	opts := &swagger.RepositoriesApiRepositoriesWorkspaceRepoSlugPutOpts{
		Body: optional.NewInterface(repo),
	}

	updated, _, err := client.RepositoriesApi.RepositoriesWorkspaceRepoSlugPut(
		auth,
		updateRepositoryOpts.slug,
		updateRepositoryOpts.workspace,
		opts,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Repository updated successfully: %s\n", updated.FullName)
	if updated.Project != nil {
		fmt.Printf("Project: %s (%s)\n", updated.Project.Name, updated.Project.Key)
	}
	return nil
}
