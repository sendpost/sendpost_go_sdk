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

type ModelsIpStat struct {
	Clicked int64 `json:"clicked,omitempty"`
	Day int64 `json:"day,omitempty"`
	Delivered int64 `json:"delivered,omitempty"`
	Dropped int64 `json:"dropped,omitempty"`
	HardBounced int64 `json:"hardBounced,omitempty"`
	Month int64 `json:"month,omitempty"`
	Opened int64 `json:"opened,omitempty"`
	Processed int64 `json:"processed,omitempty"`
	SoftBounced int64 `json:"softBounced,omitempty"`
	Spam int64 `json:"spam,omitempty"`
	Unsubscribed int64 `json:"unsubscribed,omitempty"`
	Year int64 `json:"year,omitempty"`
}
