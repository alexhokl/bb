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

type ProjectBranchingModelDevelopment struct {
	// Name of the target branch. If inherited by a repository, it will default to the main branch if the specified branch does not exist.
	Name string `json:"name"`
	// Indicates if the setting points at an explicit branch (`false`) or tracks the main branch (`true`).
	UseMainbranch bool `json:"use_mainbranch"`
}