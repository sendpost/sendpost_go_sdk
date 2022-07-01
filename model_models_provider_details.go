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

type ModelsProviderDetails struct {
	ConsumerStats *ModelsConsumerStats `json:"consumerStats,omitempty"`
	DeferredConsumerStats *ModelsConsumerStats `json:"deferredConsumerStats,omitempty"`
	ProviderSettings *ModelsIpProviderSettings `json:"providerSettings,omitempty"`
	Type_ string `json:"type,omitempty"`
}
