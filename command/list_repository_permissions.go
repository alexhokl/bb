package command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type listRepositoryPermissionsOptions struct {
	workspace string
	slug      string
}

var listRepositoryPermissionsOpts listRepositoryPermissionsOptions

func (opts *listRepositoryPermissionsOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.slug == "" {
		return fmt.Errorf("slug is required")
	}
	return nil
}

var listRepositoryPermissionsCmd = &cobra.Command{
	Use:   "repository-permissions",
	Short: "List group and user permissions of a repository",
	RunE:  runListRepositoryPermissions,
}

func init() {
	listCmd.AddCommand(listRepositoryPermissionsCmd)
	flags := listRepositoryPermissionsCmd.Flags()
	flags.StringVarP(&listRepositoryPermissionsOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&listRepositoryPermissionsOpts.slug, "slug", "s", "", "Repository slug")
	_ = listRepositoryPermissionsCmd.MarkFlagRequired("workspace")
	_ = listRepositoryPermissionsCmd.MarkFlagRequired("slug")
}

func runListRepositoryPermissions(_ *cobra.Command, _ []string) error {
	if err := listRepositoryPermissionsOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	workspace := listRepositoryPermissionsOpts.workspace
	slug := listRepositoryPermissionsOpts.slug

	groupPage, _, err := client.RepositoriesApi.RepositoriesWorkspaceRepoSlugPermissionsConfigGroupsGet(auth, slug, workspace)
	if err != nil {
		return fmt.Errorf("failed to list group permissions: %w", err)
	}
	for _, p := range groupPage.Values {
		if p.Group != nil {
			fmt.Printf("group\t%s\t%s\n", p.Permission, p.Group.Slug)
		}
	}
	for nextURL := groupPage.Next; nextURL != ""; {
		page, err := fetchGroupPermissionPage(savedToken.AccessToken, nextURL)
		if err != nil {
			return err
		}
		for _, p := range page.Values {
			if p.Group != nil {
				fmt.Printf("group\t%s\t%s\n", p.Permission, p.Group.Slug)
			}
		}
		nextURL = page.Next
	}

	userPage, _, err := client.RepositoriesApi.RepositoriesWorkspaceRepoSlugPermissionsConfigUsersGet(auth, slug, workspace)
	if err != nil {
		return fmt.Errorf("failed to list user permissions: %w", err)
	}
	for _, p := range userPage.Values {
		if p.User != nil {
			fmt.Printf("user\t%s\t%s\n", p.Permission, p.User.Nickname)
		}
	}
	for nextURL := userPage.Next; nextURL != ""; {
		page, err := fetchUserPermissionPage(savedToken.AccessToken, nextURL)
		if err != nil {
			return err
		}
		for _, p := range page.Values {
			if p.User != nil {
				fmt.Printf("user\t%s\t%s\n", p.Permission, p.User.Nickname)
			}
		}
		nextURL = page.Next
	}

	return nil
}

func fetchGroupPermissionPage(token, url string) (*swagger.PaginatedRepositoryGroupPermissions, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call permissions API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("permissions API returned status %d: %s", resp.StatusCode, string(body))
	}

	var page swagger.PaginatedRepositoryGroupPermissions
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &page, nil
}

func fetchUserPermissionPage(token, url string) (*swagger.PaginatedRepositoryUserPermissions, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call permissions API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("permissions API returned status %d: %s", resp.StatusCode, string(body))
	}

	var page swagger.PaginatedRepositoryUserPermissions
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &page, nil
}
