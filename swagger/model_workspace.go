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

import (
	"time"
)

type Workspace struct {
	Type_ string `json:"type"`
	Links *WorkspaceLinks `json:"links,omitempty"`
	// The workspace's immutable id.
	Uuid string `json:"uuid,omitempty"`
	// The name of the workspace.
	Name string `json:"name,omitempty"`
	// The short label that identifies this workspace.
	Slug string `json:"slug,omitempty"`
	// Indicates whether the workspace is publicly accessible, or whether it is private to the members and consequently only visible to members.
	IsPrivate bool `json:"is_private,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
}
