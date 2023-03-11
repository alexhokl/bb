package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-environments-environment-uuid-delete

var deleteEnvironmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"env"},
	Short:   "delete an environment",
	RunE:    runDeleteEnvironment,
}

type deleteEnvironmentOptions struct {
	uuid string
}

var deleteEnvironmentOpts deleteEnvironmentOptions

func init() {
	deleteCmd.AddCommand(deleteEnvironmentCmd)

	flags := deleteEnvironmentCmd.Flags()
	flags.StringVar(&deleteEnvironmentOpts.uuid, "id", "", "UUID of environment with curly braces")

	deleteEnvironmentCmd.MarkFlagRequired("id")
}

func runDeleteEnvironment(_ *cobra.Command, _ []string) error {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	repo, err := getRepositoryInfoFromCurrentPath()
	if err != nil {
		return err
	}

	_, err = client.DeploymentsApi.DeleteEnvironmentForRepository(
		auth,
		repo.Org,
		repo.Name,
		deleteEnvironmentOpts.uuid,
	)
	if err != nil {
		return err
	}

	fmt.Println("Environment deleted successfully")
	return nil
}
