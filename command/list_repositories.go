package command

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type listRepositoriesOptions struct {
	workspace  string
	projectKey string
}

var listRepositoriesOpts listRepositoriesOptions

func (opts *listRepositoriesOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	return nil
}

var listRepositoriesCmd = &cobra.Command{
	Use:     "repositories",
	Aliases: []string{"repo", "repos"},
	Short:   "List repositories in a workspace",
	RunE:    runListRepositories,
}

func init() {
	listCmd.AddCommand(listRepositoriesCmd)
	flags := listRepositoriesCmd.Flags()
	flags.StringVarP(&listRepositoriesOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&listRepositoriesOpts.projectKey, "project-key", "p", "", "Filter by project key")
	_ = listRepositoriesCmd.MarkFlagRequired("workspace")
}

type repositoryPage struct {
	Values []swagger.Repository `json:"values"`
	Next   string               `json:"next"`
}

func runListRepositories(_ *cobra.Command, _ []string) error {
	if err := listRepositoriesOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	nextURL := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s", listRepositoriesOpts.workspace)
	if listRepositoriesOpts.projectKey != "" {
		params := url.Values{}
		params.Set("q", fmt.Sprintf("project.key=\"%s\"", listRepositoriesOpts.projectKey))
		nextURL += "?" + params.Encode()
	}

	for nextURL != "" {
		req, err := http.NewRequest(http.MethodGet, nextURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to call repositories API: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("repositories API returned status %d: %s", resp.StatusCode, string(body))
		}

		var page repositoryPage
		if err := json.Unmarshal(body, &page); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}

		for _, r := range page.Values {
			fmt.Printf("%s\n", r.FullName)
		}

		nextURL = page.Next
	}

	return nil
}
