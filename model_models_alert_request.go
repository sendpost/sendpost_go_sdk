/*
 * SendPost API
 *
 * SendPost API to send transactional emails reliably
 *
 * API version: 1.0.0
 * Contact: hello@sendx.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsAlertRequest struct {
	Active bool `json:"active,omitempty"`
	Emails *interface{} `json:"emails,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	EntityValue int64 `json:"entityValue,omitempty"`
	Label *ModelsAlertLabel `json:"label,omitempty"`
	Member int64 `json:"member,omitempty"`
	Name string `json:"name,omitempty"`
	NotificationType *ModelsNotificationType `json:"notificationType,omitempty"`
	Parameter int64 `json:"parameter,omitempty"`
	Provider string `json:"provider,omitempty"`
	SlackUrl string `json:"slackUrl,omitempty"`
	Threshold string `json:"threshold,omitempty"`
}
