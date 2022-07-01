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

type ModelsBlackListEngine struct {
	Confidence string `json:"confidence,omitempty"`
	Detected bool `json:"detected,omitempty"`
	Elapsed string `json:"elapsed,omitempty"`
	Engine string `json:"engine,omitempty"`
	Reference string `json:"reference,omitempty"`
}
