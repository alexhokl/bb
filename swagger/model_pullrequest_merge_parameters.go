/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The metadata that describes a pull request merge.
type PullrequestMergeParameters struct {
	Type_ string `json:"type"`
	// The commit message that will be used on the resulting commit.
	Message string `json:"message,omitempty"`
	// Whether the source branch should be deleted. If this is not provided, we fallback to the value used when the pull request was created, which defaults to False
	CloseSourceBranch bool `json:"close_source_branch,omitempty"`
	// The merge strategy that will be used to merge the pull request.
	MergeStrategy string `json:"merge_strategy,omitempty"`
}
