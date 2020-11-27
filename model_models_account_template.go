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

type ModelsAccountTemplate struct {
	Created int64 `json:"created,omitempty"`
	Html string `json:"html,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Template string `json:"template,omitempty"`
	Text string `json:"text,omitempty"`
	Updated int64 `json:"updated,omitempty"`
}
