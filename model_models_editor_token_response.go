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

type ModelsEditorTokenResponse struct {
	Expires string `json:".expires,omitempty"`
	Issued string `json:".issued,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	AsclientId string `json:"as:client_id,omitempty"`
	Asregion string `json:"as:region,omitempty"`
	ExpiresIn int32 `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	UserName string `json:"userName,omitempty"`
}
