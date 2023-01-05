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

type User struct {
	Type_ string `json:"type"`
	Links *UserLinks `json:"links,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	// The user's Atlassian account ID.
	AccountId string `json:"account_id,omitempty"`
	// The status of the account. Currently the only possible value is \"active\", but more values may be added in the future.
	AccountStatus string `json:"account_status,omitempty"`
	Has2faEnabled bool `json:"has_2fa_enabled,omitempty"`
	// Account name defined by the owner. Should be used instead of the \"username\" field. Note that \"nickname\" cannot be used in place of \"username\" in URLs and queries, as \"nickname\" is not guaranteed to be unique.
	Nickname string `json:"nickname,omitempty"`
	IsStaff bool `json:"is_staff,omitempty"`
	Website string `json:"website,omitempty"`
}
