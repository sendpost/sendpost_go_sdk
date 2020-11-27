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

type ModelsIncident struct {
	Author *ModelsMember `json:"author,omitempty"`
	Created int64 `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	RelatedIP *ModelsIp `json:"relatedIP,omitempty"`
	RelatedSubAccount *ModelsSubAccount `json:"relatedSubAccount,omitempty"`
	Status *ModelsIncidentStatus `json:"status,omitempty"`
	Summary string `json:"summary,omitempty"`
	Tags []ModelsTag `json:"tags,omitempty"`
	Updated int64 `json:"updated,omitempty"`
}