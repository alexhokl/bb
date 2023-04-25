package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/iohelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-deployments-config-environments-environment-uuid-variables-post

var createEnvironmentVariableCmd = &cobra.Command{
	Use:   "environment-variable",
	Aliases: []string{"env-var"},
	Short: "Create a environment variable",
	RunE: runCreateEnvironmentVariables,
}

type createEnvironmentVariableOptions struct {
	name string
	key string
	value string
	valueFilePath string
	isSecret bool
}

var createEnvironmentVariableOpts createEnvironmentVariableOptions

func init() {
	createCmd.AddCommand(createEnvironmentVariableCmd)

	flags := createEnvironmentVariableCmd.Flags()
	flags.StringVar(&createEnvironmentVariableOpts.name, "env", "", "name of environment")
	flags.StringVar(&createEnvironmentVariableOpts.key, "key", "", "key of environment variable")
	flags.StringVar(&createEnvironmentVariableOpts.value, "value", "", "value of environment variable")
	flags.StringVar(&createEnvironmentVariableOpts.valueFilePath, "file", "", "path to file containing value of environment variable")
	flags.BoolVar(&createEnvironmentVariableOpts.isSecret, "secret", false, "value of environment variable")

	createEnvironmentVariableCmd.MarkFlagRequired("env")
	createEnvironmentVariableCmd.MarkFlagRequired("key")
}

func runCreateEnvironmentVariables(_ *cobra.Command, _ []string) error {
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

	envUUID, err := getUUIDFromEnvironmentName(repo, client, auth, createEnvironmentVariableOpts.name)
	if err != nil {
		return err
	}

	value := createEnvironmentVariableOpts.value
	if value == "" {
		value, err = iohelper.ReadStringFromFile(createEnvironmentVariableOpts.valueFilePath)
		if err != nil {
			return err
		}
	}

	environmentVariable := swagger.DeploymentVariable{
		Key: createEnvironmentVariableOpts.key,
		Value: value,
		Secured: createEnvironmentVariableOpts.isSecret,
	}

	_, _, err = client.PipelinesApi.CreateDeploymentVariable(
		auth,
		repo.Org,
		repo.Name,
		envUUID,
		environmentVariable,
	)
	if err != nil {
		return err
	}

	fmt.Println("Environment variable created successfully")
	return nil
}

func validateCreateEnvironmentVariableOptions(opts createEnvironmentVariableOptions) error {
	if opts.value == "" && opts.valueFilePath == "" {
		return fmt.Errorf("either --value or --file must be specified")
	}
	if opts.value != "" && opts.valueFilePath != "" {
		return fmt.Errorf("cannot provide both --value and --file")
	}
	return nil
}
