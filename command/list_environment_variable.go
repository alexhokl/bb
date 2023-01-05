package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-deployments-config-environments-environment-uuid-variables-get

var listEnvironmentVariableCmd = &cobra.Command{
	Use:   "environment-variable",
	Aliases: []string{"env-var"},
	Short: "List environment variables",
	RunE: runListEnvironmentVariables,
}

type listEnvironmentVariableOptions struct {
	uuid string
}

var listEnvironmentVariableOpts listEnvironmentVariableOptions

func init() {
	listCmd.AddCommand(listEnvironmentVariableCmd)

	flags := listEnvironmentVariableCmd.Flags()
	flags.StringVar(&listEnvironmentVariableOpts.uuid, "id", "", "UUID of environment with curly braces")

	listEnvironmentVariableCmd.MarkFlagRequired("id")
}

func runListEnvironmentVariables(_ *cobra.Command, _ []string) error {
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

	page, _, err := client.PipelinesApi.GetDeploymentVariables(
		auth,
		repo.Org,
		repo.Name,
		listEnvironmentVariableOpts.uuid,
	)
	if err != nil {
		return err
	}

	var environmentVariables []swagger.DeploymentVariable
	environmentVariables = append(environmentVariables, page.Values...)

	if len(environmentVariables) == 0 {
		fmt.Println("There are no environment variables configured")
		return nil
	}

	fmt.Println("Key,Value")
	for _, v := range environmentVariables {
		 fmt.Printf("%s,%s\n", v.Key, v.Value)
	}
	return nil
}
