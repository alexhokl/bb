package command

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/collection"
	"github.com/alexhokl/helper/jsonhelper"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get

type commitListResponse struct {
	Items []models.CommitInfo `json:"values"`
	Next  string              `json:"next"`
}

type listJiraIDsOptions struct {
	isCommaSeparated bool
	pullRequestID    int32
	idPrefixes       []string
}

var listJiraOpts listJiraIDsOptions

func (opts listJiraIDsOptions) Validate() error {
	if len(opts.idPrefixes) == 0 {
		return fmt.Errorf("JIRA ID prefixes are not specified")
	}
	return nil
}

var listJiraCmd = &cobra.Command{
	Use:   "jira",
	Short: "List JIRA IDs of a pull request which has the specified JIRA ID prefix(es)",
	RunE:  runListJiraIDs,
}

func init() {
	listCmd.AddCommand(listJiraCmd)

	flags := listJiraCmd.Flags()
	flags.Int32VarP(&listJiraOpts.pullRequestID, "id", "i", 0, "Pull request ID")
	flags.BoolVar(&listJiraOpts.isCommaSeparated, "comma", false, "comma separated list")
	flags.StringArrayVarP(&listJiraOpts.idPrefixes, "prefixes", "p", []string{}, "Comma separated list of JIRA ID prefixes")
}

func runListJiraIDs(_ *cobra.Command, _ []string) error {
	if err := listJiraOpts.Validate(); err != nil {
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

	responseBodyBytes, _, err := client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet(
		auth,
		listJiraOpts.pullRequestID,
		repo.Name,
		repo.Org,
	)
	if err != nil {
		return err
	}
	list, errParse := parseCommitListResponse(responseBodyBytes)
	if errParse != nil {
		return errParse
	}
	if len(list) == 0 {
		return fmt.Errorf("no commits found")
	}

	var builder strings.Builder
	builder.WriteString(`(((`)
	for index, p := range listJiraOpts.idPrefixes {
		if index == 0 {
			builder.WriteString(p)
		} else {
			builder.WriteString(fmt.Sprintf("|%s", p))
		}
	}
	builder.WriteString(`)-\d+))`)

	regex := regexp.MustCompile(builder.String())

	var ids []string
	for _, c := range list {
		matches := regex.FindAllString(c.Summary.Raw, -1)
		ids = append(ids, matches...)
	}
	distinctIDs := collection.GetDistinct(ids)
	sort.Strings(distinctIDs)

	if listJiraOpts.isCommaSeparated {
		fmt.Print(collection.GetDelimitedString(distinctIDs, ", "))
		fmt.Println()
	} else {
		for _, i := range distinctIDs {
			fmt.Println(i)
		}
	}

	return nil
}

func parseCommitListResponse(body []byte) ([]models.CommitInfo, error) {
	var list []models.CommitInfo

	// TODO: handle pagination
	var listResponse commitListResponse
	errParse := jsonhelper.ParseJSONFromBytes(&listResponse, body)
	if errParse != nil {
		return nil, fmt.Errorf("failed to parse response body to JSON [%v]", errParse)
	}
	if list == nil {
		list = listResponse.Items
	} else {
		list = append(list, listResponse.Items...)
	}

	return list, nil
}
