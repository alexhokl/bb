package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/git"
	"github.com/alexhokl/helper/jsonhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get

type pullRequestActivityListResponse struct {
	Items []models.PullRequestActivity `json:"values"`
	Next  string                       `json:"next"`
}

type describeOptions struct {
	idOptions
	isShowDifftool bool
}

var describeOpts describeOptions

func (opts describeOptions) Validate() error {
	return opts.idOptions.validate()
}

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe the specified pull request",
	RunE:  runDescribe,
}

func init() {
	rootCmd.AddCommand(describeCmd)

	flags := describeCmd.Flags()
	flags.Int32VarP(&describeOpts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&describeOpts.isShowDifftool, "difftool", "d", false, "Open difftool after description")
}

func runDescribe(_ *cobra.Command, _ []string) error {
	if err := describeOpts.Validate(); err != nil {
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
		describeOpts.id,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}

	respBody, _, errActivities := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet(
		auth,
		describeOpts.id,
		repo.Name,
		repo.Org,
	)
	if errActivities != nil {
		return errActivities
	}
	activities, errParse := parseActivityResponse(respBody)
	if errParse != nil {
		return errParse
	}

	models.PrintShortDescription(&pr, true)
	models.PrintDescription(&pr)

	fmt.Println("---")

	for _, event := range activities {
		if event.Comment != (models.Comment{}) {
			fmt.Println(event.Comment.ToString())
		}
		if event.Update != (models.Update{}) {
			fmt.Println(event.Update.ToString())
		}
	}

	stat, errStat := git.DiffStat(pr.Destination.Branch.Name, pr.Source.Branch.Name)
	if errStat != nil {
		return errStat
	}
	fmt.Println("Diff:")
	fmt.Println(stat)

	if describeOpts.isShowDifftool {
		errDifftool := git.Difftool(pr.Destination.Branch.Name, pr.Source.Branch.Name)
		if errDifftool != nil {
			return errDifftool
		}
	}

	return nil
}

func parseActivityResponse(body []byte) ([]models.PullRequestActivity, error) {
	var list []models.PullRequestActivity

	// TODO: handle pagination
	var listResponse pullRequestActivityListResponse
	errParse := jsonhelper.ParseJSONFromBytes(&listResponse, body)
	if errParse != nil {
		return nil, fmt.Errorf("failed to parse response body to JSON [%v]", errParse)
	}
	if list == nil {
		list = listResponse.Items
	} else {
		list = append(list, listResponse.Items...)
	}

	// getting around a bug of the API
	updatedList := list[:len(list)-1]

	return updatedList, nil
}
