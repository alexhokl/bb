package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-environments-get

var listEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Aliases: []string{"env"},
	Short: "List environments",
	RunE: runListEnvironment,
}

func init() {
	listCmd.AddCommand(listEnvironmentCmd)
}

func runListEnvironment(_ *cobra.Command, _ []string) error {
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

	page, _, err := client.DeploymentsApi.GetEnvironmentsForRepository(
		auth,
		repo.Org,
		repo.Name,
	)
	if err != nil {
		return err
	}

	var environments []swagger.DeploymentEnvironment


	environments = append(environments, page.Values...)

	if len(environments) == 0 {
		fmt.Println("There are no environments.")
		return nil
	}

	fmt.Println("UUID,Name")
	for _, e := range environments {
		 fmt.Printf("%s,%s\n", e.Uuid, e.Name)
	}
	return nil
}
