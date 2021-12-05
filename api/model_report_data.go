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

// A key-value element that will be displayed along with the report.
type ReportData struct {
	// The type of data contained in the value field. If not provided, then the value will be detected as a boolean, number or string.
	Type_ string `json:"type,omitempty"`
	// A string describing what this data field represents.
	Title string `json:"title,omitempty"`
	// The value of the data element.
	Value *interface{} `json:"value,omitempty"`
}
