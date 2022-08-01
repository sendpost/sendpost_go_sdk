/*
 * SendPost API
 *
 * Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.
 *
 * API version: 1.0.0
 * Contact: hello@sendpost.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsSubAccount struct {
	ApiKey string `json:"apiKey,omitempty"`
	Created int64 `json:"created,omitempty"`
	Id int64 `json:"id,omitempty"`
	Labels []ModelsLabel `json:"labels,omitempty"`
	Name string `json:"name,omitempty"`
	SmtpAuths []ModelsSmtpAuth `json:"smtpAuths,omitempty"`
	Type_ *ModelsSubAccountType `json:"type,omitempty"`
}
