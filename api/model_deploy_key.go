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

type DeployKey struct {
	Type_ string `json:"type"`
	// The deploy key value.
	Key string `json:"key,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	// The comment parsed from the deploy key (if present)
	Comment string `json:"comment,omitempty"`
	// The user-defined label for the deploy key
	Label string `json:"label,omitempty"`
	AddedOn time.Time `json:"added_on,omitempty"`
	LastUsed time.Time `json:"last_used,omitempty"`
	Links *BranchingModelSettingsLinks `json:"links,omitempty"`
	Owner *Account `json:"owner,omitempty"`
}
