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

type ModelsComment struct {
	Author *ModelsMember `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
	Created int64 `json:"created,omitempty"`
	Id int64 `json:"id,omitempty"`
	Incident *ModelsIncident `json:"incident,omitempty"`
}
