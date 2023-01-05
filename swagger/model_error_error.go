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

type ErrorError struct {
	Message string `json:"message"`
	Detail string `json:"detail,omitempty"`
	// Optional structured data that is endpoint-specific.
	Data interface{} `json:"data,omitempty"`
}
