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

type ModelsOnboardingChecklist struct {
	Id int64 `json:"id,omitempty"`
	IsDomainAdded bool `json:"isDomainAdded,omitempty"`
	IsDomainVerified bool `json:"isDomainVerified,omitempty"`
	IsFirstEmailSent bool `json:"isFirstEmailSent,omitempty"`
}
