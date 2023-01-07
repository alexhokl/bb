package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

var unapproveCmd = &cobra.Command{
	Use:   "unapprove",
	Short: "Un-approve the specified pull request",
	RunE:  runUnapprove,
}

func init() {
	rootCmd.AddCommand(unapproveCmd)

	flags := unapproveCmd.Flags()
	flags.Int32VarP(&idOpts.id, "id", "i", 0, "Pull request ID")
}

func runUnapprove(_ *cobra.Command, _ []string) error {
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

	_, err = client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete(
		auth,
		idOpts.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Un-approved pull request [%d].\n", idOpts.id)
	return nil
}
