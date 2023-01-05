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

// An event, associated with a resource or subject type.
type HookEvent struct {
	// The event identifier.
	Event string `json:"event,omitempty"`
	// The category this event belongs to.
	Category string `json:"category,omitempty"`
	// Summary of the webhook event type.
	Label string `json:"label,omitempty"`
	// More detailed description of the webhook event type.
	Description string `json:"description,omitempty"`
}
