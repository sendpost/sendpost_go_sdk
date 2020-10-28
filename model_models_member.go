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

type ModelsMember struct {
	Email string `json:"Email,omitempty"`
	Id int64 `json:"Id,omitempty"`
	IsForbidden bool `json:"IsForbidden,omitempty"`
	IsVerified bool `json:"IsVerified,omitempty"`
	Created int64 `json:"created,omitempty"`
}
