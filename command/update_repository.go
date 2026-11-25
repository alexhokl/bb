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
	mainBranch string
}

var updateRepositoryOpts updateRepositoryOptions

func (opts *updateRepositoryOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.slug == "" {
		return fmt.Errorf("slug is required")
	}
	if opts.projectKey == "" && opts.mainBranch == "" {
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
	flags.StringVarP(&updateRepositoryOpts.mainBranch, "main-branch", "b", "", "Name of the branch to set as the main branch")

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

	opts := &swagger.RepositoriesApiRepositoriesWorkspaceRepoSlugPutOpts{
		Body: optional.NewInterface(buildRepositoryUpdate(updateRepositoryOpts)),
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
	if updated.Mainbranch != nil {
		fmt.Printf("Main branch: %s\n", updated.Mainbranch.Name)
	}
	return nil
}

// buildRepositoryUpdate returns the request body of a repository update,
// containing only the fields given in opts
func buildRepositoryUpdate(opts updateRepositoryOptions) swagger.Repository {
	repo := swagger.Repository{Type_: "repository"}
	if opts.projectKey != "" {
		repo.Project = &swagger.Project{Key: opts.projectKey}
	}
	if opts.mainBranch != "" {
		repo.Mainbranch = &swagger.Branch{Type_: "branch", Name: opts.mainBranch}
	}
	return repo
}
