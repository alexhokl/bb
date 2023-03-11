package command

import (
	"fmt"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"golang.org/x/net/context"
)

type idOptions struct {
	id int32
}

var idOpts idOptions

func (opts idOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("invalid pull request ID")
	}

	return nil
}

func getUUIDFromEnvironmentName(repo *models.Repository, client *swagger.APIClient, auth context.Context, name string) (string, error) {
	page, _, err := client.DeploymentsApi.GetEnvironmentsForRepository(
		auth,
		repo.Org,
		repo.Name,
	)
	if err != nil {
		return "", err
	}

	var environments []swagger.DeploymentEnvironment
	environments = append(environments, page.Values...)

	if len(environments) == 0 {
		return "", fmt.Errorf("no environment %s found", name)
	}

	for _, e := range environments {
		if e.Name == name {
			return e.Uuid, nil
		}
	}
	return "", fmt.Errorf("no environment %s found", name)
}
