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

type ModelsSmtpStat struct {
	Count int64 `json:"count,omitempty"`
	SmtpCode int64 `json:"smtpCode,omitempty"`
	SmtpDescription string `json:"smtpDescription,omitempty"`
}
