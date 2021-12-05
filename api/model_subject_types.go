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

// The mapping of resource/subject types pointing to their individual event types.
type SubjectTypes struct {
	Repository *SubjectTypesRepository `json:"repository,omitempty"`
	User *SubjectTypesRepository `json:"user,omitempty"`
	Team *SubjectTypesRepository `json:"team,omitempty"`
}
