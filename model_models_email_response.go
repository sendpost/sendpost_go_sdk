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

type ModelsEmailResponse struct {
	ErrorCode *ModelsEmailErrorCode `json:"errorCode,omitempty"`
	Message string `json:"message,omitempty"`
	MessageId string `json:"messageId,omitempty"`
	SubmittedAt int64 `json:"submittedAt,omitempty"`
	To string `json:"to,omitempty"`
}
