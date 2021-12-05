/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

type BranchingModelDevelopment struct {
	Branch *Branch `json:"branch,omitempty"`
	// Name of the target branch. Will be listed here even when the target branch does not exist. Will be `null` if targeting the main branch and the repository is empty.
	Name string `json:"name"`
	// Indicates if the setting points at an explicit branch (`false`) or tracks the main branch (`true`).
	UseMainbranch bool `json:"use_mainbranch"`
	// Indicates if the indicated branch exists on the repository (`false`)or not (`true`). This is useful for determining a fallback to the mainbranch when a repository is inheriting its project's branching model.
	BranchDoesNotExist bool `json:"branch_does_not_exist,omitempty"`
}