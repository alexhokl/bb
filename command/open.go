package command

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	clihelper "github.com/alexhokl/helper/cli"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the web page of the specified pull request in a browser",
	RunE:  runOpen,
}

func init() {
	rootCmd.AddCommand(openCmd)

	flags := openCmd.Flags()
	flags.Int32VarP(&idOpts.id, "id", "i", 0, "Pull request ID")
}

func runOpen(_ *cobra.Command, _ []string) error {
	if err := idOpts.validate(); err != nil {
		return err
	}

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

	pr, _, err := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet(
		auth,
		idOpts.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}

	cmdName, cmdArgs := clihelper.GetOpenCommand(pr.Links.Html.Href)
	_, errOpen := exec.Command(cmdName, cmdArgs...).Output()
	if errOpen != nil {
		return errOpen
	}

	fmt.Printf("Opened %s in a browser.\n", pr.Links.Html.Href)
	return nil
}
