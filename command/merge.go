package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/git"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

// reference: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-merge-post

type MergeStrategy string

const (
	MergeCommit MergeStrategy = "merge_commit"
	NoCommit    MergeStrategy = "no_commit"
	Squash      MergeStrategy = "squash"
)

type mergeOptions struct {
	idOptions
	isKeepBranch        bool
	isCloseSourceBranch bool
	mergeStrategy       string
}

var mergeOpts mergeOptions

func (opts mergeOptions) Validate() error {
	return opts.idOptions.validate()
}

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge the specified pull request",
	RunE:  runMerge,
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	flags := mergeCmd.Flags()
	flags.Int32VarP(&mergeOpts.id, "id", "i", 0, "Pull request ID")
	flags.BoolVarP(&mergeOpts.isKeepBranch, "keep-branch", "k", false, "Keep local branch after merging")
	flags.BoolVar(&mergeOpts.isCloseSourceBranch, "close-source-branch", true, "Close source branch after merging")
	flags.StringVar(&mergeOpts.mergeStrategy, "strategy", "merge_commit", "Merge strategy (merge_commit, no_commit, squash)")

	mergeCmd.MarkFlagRequired("id")
}

func runMerge(_ *cobra.Command, _ []string) error {
	if err := mergeOpts.Validate(); err != nil {
		return err
	}
	strategy, err := getMergeStrategy(mergeOpts.mergeStrategy)
	if err != nil {
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

	mergeParams := swagger.PullrequestMergeParameters{
		CloseSourceBranch: mergeOpts.isCloseSourceBranch,
		MergeStrategy:     string(strategy),
		Message:           "",
		Type_:             "merge",
	}
	opts := &swagger.PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePostOpts{
		Async: optional.NewBool(false),
		Body:  optional.NewInterface(mergeParams),
	}
	_, _, err = client.PullrequestsApi.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost(
		auth,
		describeOpts.id,
		repo.Name,
		repo.Org,
		opts,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Merged pull request [%d].\n", mergeOpts.id)

	if mergeOpts.isKeepBranch {
		return nil
	}

	currentBranchName, errCurrentBranch := git.GetCurrentBranchName()
	if errCurrentBranch != nil {
		return errCurrentBranch
	}

	if currentBranchName == pr.Source.Branch.Name {
		_, errCheckout := git.Checkout(pr.Destination.Branch.Name)
		if errCheckout != nil {
			return errCheckout
		}
		fmt.Printf("Checked out to branch %s.\n", pr.Destination.Branch.Name)
	}

	isBranchExists, errBranchCheck := git.IsBranchExists(pr.Source.Branch.Name)
	if errBranchCheck != nil {
		return errCurrentBranch
	}

	if isBranchExists {
		_, errDelete := git.DeleteBranch(pr.Source.Branch.Name)
		if errDelete != nil {
			return errDelete
		}
		fmt.Printf("Deleted branch %s.\n", pr.Source.Branch.Name)
	}

	return nil
}

func getMergeStrategy(input string) (MergeStrategy, error) {
	switch input {
	case "merge_commit":
		return MergeCommit, nil
	case "no_commit":
		return NoCommit, nil
	case "squash":
		return Squash, nil
	default:
		return "", fmt.Errorf("invalid merge strategy %s", input)
	}
}
