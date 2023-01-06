package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-environments-post

var createEnvironmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"env"},
	Short:   "Create an environment",
	RunE:    runCreateEnvironment,
}

type createEnvironmentOptions struct {
	name           string
	enviromentType string
}

var createEnvironmentOpts createEnvironmentOptions

func init() {
	createCmd.AddCommand(createEnvironmentCmd)

	flags := createEnvironmentCmd.Flags()
	flags.StringVar(&createEnvironmentOpts.name, "name", "", "name of environment")
	flags.StringVar(&createEnvironmentOpts.enviromentType, "environment-type", "", "type of environment (e.g. staging, production)")

	createEnvironmentCmd.MarkFlagRequired("name")
	createEnvironmentCmd.MarkFlagRequired("environment-type")
}

func runCreateEnvironment(_ *cobra.Command, _ []string) error {
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

	environment := swagger.DeploymentEnvironment{
		Name: createEnvironmentOpts.name,
		EnvironmentType: swagger.EnvironmentType{
			Name: createEnvironmentOpts.enviromentType,
		},
	}

	_, _, err = client.DeploymentsApi.CreateEnvironment(
		auth,
		repo.Org,
		repo.Name,
		environment,
	)
	if err != nil {
		return err
	}

	fmt.Println("Environment created successfully")
	return nil
}
