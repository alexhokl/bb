package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

type createRepositoryOptions struct {
	workspace   string
	name        string
	description string
	isPrivate   bool
	scm         string
}

var createRepositoryOpts createRepositoryOptions

func (opts *createRepositoryOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.name == "" {
		return fmt.Errorf("name is required")
	}
	if opts.scm != "git" && opts.scm != "hg" {
		return fmt.Errorf("scm must be either 'git' or 'hg'")
	}
	return nil
}

var createRepositoryCmd = &cobra.Command{
	Use:     "repository",
	Aliases: []string{"repo"},
	Short:   "Create a repository",
	RunE:    runCreateRepository,
}

func init() {
	createCmd.AddCommand(createRepositoryCmd)

	flags := createRepositoryCmd.Flags()
	flags.StringVarP(&createRepositoryOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&createRepositoryOpts.name, "name", "n", "", "Name of the repository")
	flags.StringVarP(&createRepositoryOpts.description, "description", "d", "", "Description of the repository")
	flags.BoolVar(&createRepositoryOpts.isPrivate, "private", true, "Whether the repository is private")
	flags.StringVar(&createRepositoryOpts.scm, "scm", "git", "Source control management type (git or hg)")

	_ = createRepositoryCmd.MarkFlagRequired("workspace")
	_ = createRepositoryCmd.MarkFlagRequired("name")
}

func runCreateRepository(_ *cobra.Command, _ []string) error {
	if err := createRepositoryOpts.validate(); err != nil {
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
		Type_:       "repository",
		Name:        createRepositoryOpts.name,
		Description: createRepositoryOpts.description,
		IsPrivate:   createRepositoryOpts.isPrivate,
		Scm:         createRepositoryOpts.scm,
	}

	opts := &swagger.RepositoriesApiRepositoriesWorkspaceRepoSlugPostOpts{
		Body: optional.NewInterface(repo),
	}

	created, _, err := client.RepositoriesApi.RepositoriesWorkspaceRepoSlugPost(
		auth,
		createRepositoryOpts.name,
		createRepositoryOpts.workspace,
		opts,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Repository created successfully: %s\n", created.FullName)
	return nil
}
