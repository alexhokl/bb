package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

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

func getUUIDFromEnvironmentVariableKey(repo *models.Repository, client *swagger.APIClient, auth context.Context, envUUID string, key string) (string, error) {
	page, _, err := client.PipelinesApi.GetDeploymentVariables(
		auth,
		repo.Org,
		repo.Name,
		envUUID,
	)
	if err != nil {
		return "", err
	}

	var variables []swagger.DeploymentVariable
	variables = append(variables, page.Values...)

	if len(variables) == 0 {
		return "", fmt.Errorf("no environment variable %s found", key)
	}

	for _, v := range variables {
		if v.Key == key {
			return v.Uuid, nil
		}
	}
	return "", fmt.Errorf("no environment variable %s found", key)
}

// readStdinIfPiped returns the trimmed content of stdin when data is being
// piped into the process. If stdin is a TTY (no piped input), it returns an
// empty string and nil error so callers can fall back to flag-based input.
func readStdinIfPiped() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to stat stdin: %w", err)
	}
	if (info.Mode() & os.ModeCharDevice) != 0 {
		return "", nil
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, os.Stdin); err != nil {
		return "", fmt.Errorf("failed to read stdin: %w", err)
	}
	return strings.TrimSpace(buf.String()), nil
}

// getCurrentUserUUID returns the UUID of the currently authenticated user.
func getCurrentUserUUID(client *swagger.APIClient, auth context.Context) (string, error) {
	account, _, err := client.UsersApi.UserGet(auth)
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}
	if account.Uuid == "" {
		return "", fmt.Errorf("current user has no UUID")
	}
	return account.Uuid, nil
}
