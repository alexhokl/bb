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

type PullrequestLinks struct {
	Self *Link `json:"self,omitempty"`
	Html *Link `json:"html,omitempty"`
	Commits *Link `json:"commits,omitempty"`
	Approve *Link `json:"approve,omitempty"`
	Diff *Link `json:"diff,omitempty"`
	Diffstat *Link `json:"diffstat,omitempty"`
	Comments *Link `json:"comments,omitempty"`
	Activity *Link `json:"activity,omitempty"`
	Merge *Link `json:"merge,omitempty"`
	Decline *Link `json:"decline,omitempty"`
}
