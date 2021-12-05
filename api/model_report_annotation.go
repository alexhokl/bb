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
import (
	"time"
)

type ReportAnnotation struct {
	Type_ string `json:"type"`
	// ID of the annotation provided by the annotation creator. It can be used to identify the annotation as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the annotation creator for updating or deleting this specific annotation. Needs to be unique.
	ExternalId string `json:"external_id,omitempty"`
	// The UUID that can be used to identify the annotation.
	Uuid string `json:"uuid,omitempty"`
	// The type of the report.
	AnnotationType string `json:"annotation_type,omitempty"`
	// The path of the file on which this annotation should be placed. This is the path of the file relative to the git repository. If no path is provided, then it will appear in the overview modal on all pull requests where the tip of the branch is the given commit, regardless of which files were modified.
	Path string `json:"path,omitempty"`
	// The line number that the annotation should belong to. If no line number is provided, then it will default to 0 and in a pull request it will appear at the top of the file specified by the path field.
	Line int32 `json:"line,omitempty"`
	// The message to display to users.
	Summary string `json:"summary,omitempty"`
	// The details to show to users when clicking on the annotation.
	Details string `json:"details,omitempty"`
	// The state of the report. May be set to PENDING and later updated.
	Result string `json:"result,omitempty"`
	// The severity of the annotation.
	Severity string `json:"severity,omitempty"`
	// A URL linking to the annotation in an external tool.
	Link string `json:"link,omitempty"`
	// The timestamp when the report was created.
	CreatedOn time.Time `json:"created_on,omitempty"`
	// The timestamp when the report was updated.
	UpdatedOn time.Time `json:"updated_on,omitempty"`
}