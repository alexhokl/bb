package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-deployments-config-environments-environment-uuid-variables-variable-uuid-delete

var deleteEnvironmentVariableCmd = &cobra.Command{
	Use:     "environment-variable",
	Aliases: []string{"env-var"},
	Short:   "Delete a environment variable",
	RunE:    runDeleteEnvironmentVariable,
}

type deleteEnvironmentVariableOptions struct {
	name string
	key  string
}

var deleteEnvironmentVariableOpts deleteEnvironmentVariableOptions

func init() {
	deleteCmd.AddCommand(deleteEnvironmentVariableCmd)

	flags := deleteEnvironmentVariableCmd.Flags()
	flags.StringVar(&deleteEnvironmentVariableOpts.name, "env", "", "name of environment")
	flags.StringVar(&deleteEnvironmentVariableOpts.key, "key", "", "key of environment variable")

	deleteEnvironmentVariableCmd.MarkFlagRequired("env")
	deleteEnvironmentVariableCmd.MarkFlagRequired("key")
}

func runDeleteEnvironmentVariable(_ *cobra.Command, _ []string) error {
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

	envUUID, err := getUUIDFromEnvironmentName(repo, client, auth, deleteEnvironmentVariableOpts.name)
	if err != nil {
		return err
	}

	variableUUID, err := getUUIDFromEnvironmentVariableKey(repo, client, auth, envUUID, deleteEnvironmentVariableOpts.key)
	if err != nil {
		return err
	}

	_, err = client.PipelinesApi.DeleteDeploymentVariable(
		auth,
		repo.Org,
		repo.Name,
		envUUID,
		variableUUID,
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Environment variable [%s] from environment [%s] successfully\n",
		deleteEnvironmentVariableOpts.key,
		deleteEnvironmentVariableOpts.name,
	)
	return nil
}
