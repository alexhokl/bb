package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type createProjectOptions struct {
	workspace   string
	name        string
	key         string
	description string
	isPrivate   bool
}

var createProjectOpts createProjectOptions

func (opts *createProjectOptions) validate() error {
	if opts.workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	if opts.name == "" {
		return fmt.Errorf("name is required")
	}
	if opts.key == "" {
		return fmt.Errorf("key is required")
	}
	return nil
}

var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create a project in a workspace",
	RunE:  runCreateProject,
}

func init() {
	createCmd.AddCommand(createProjectCmd)

	flags := createProjectCmd.Flags()
	flags.StringVarP(&createProjectOpts.workspace, "workspace", "w", "", "Workspace slug")
	flags.StringVarP(&createProjectOpts.name, "name", "n", "", "Name of the project")
	flags.StringVarP(&createProjectOpts.key, "key", "k", "", "Project key (e.g. PROJ)")
	flags.StringVarP(&createProjectOpts.description, "description", "d", "", "Description of the project")
	flags.BoolVar(&createProjectOpts.isPrivate, "private", true, "Whether the project is private")

	_ = createProjectCmd.MarkFlagRequired("workspace")
	_ = createProjectCmd.MarkFlagRequired("name")
	_ = createProjectCmd.MarkFlagRequired("key")
}

// createProjectRequest is a local struct for serialising the POST body.
// The swagger Project model always emits "type" which Bitbucket rejects with 400.
type createProjectRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
	IsPrivate   bool   `json:"is_private"`
}

func runCreateProject(_ *cobra.Command, _ []string) error {
	if err := createProjectOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	reqBody := createProjectRequest{
		Name:        createProjectOpts.name,
		Key:         createProjectOpts.key,
		Description: createProjectOpts.description,
		IsPrivate:   createProjectOpts.isPrivate,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	apiURL := fmt.Sprintf("https://api.bitbucket.org/2.0/workspaces/%s/projects", createProjectOpts.workspace)
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", savedToken.AccessToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call projects API: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("projects API returned status %d: %s", resp.StatusCode, string(respBody))
	}

	var created struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	}
	if err := json.Unmarshal(respBody, &created); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	fmt.Printf("Project created successfully: %s (%s)\n", created.Name, created.Key)
	return nil
}
