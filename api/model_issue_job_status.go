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

// The status of an import or export job
type IssueJobStatus struct {
	Type_ string `json:"type,omitempty"`
	// The status of the import/export job
	Status string `json:"status,omitempty"`
	// The phase of the import/export job
	Phase string `json:"phase,omitempty"`
	// The total number of issues being imported/exported
	Total int32 `json:"total,omitempty"`
	// The total number of issues already imported/exported
	Count int32 `json:"count,omitempty"`
	// The percentage of issues already imported/exported
	Pct float64 `json:"pct,omitempty"`
}
