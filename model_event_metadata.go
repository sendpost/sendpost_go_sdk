/*
SendPost API

Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.

API version: 1.0.0
Contact: hello@sendpost.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"encoding/json"
)

// checks if the EventMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventMetadata{}

// EventMetadata struct for EventMetadata
type EventMetadata struct {
	ClickedURL *string `json:"clickedURL,omitempty"`
	Device *Device `json:"device,omitempty"`
	Geo *City `json:"geo,omitempty"`
	Os *Os `json:"os,omitempty"`
	SmtpCode *int64 `json:"smtpCode,omitempty"`
	SmtpDescription *string `json:"smtpDescription,omitempty"`
	UserAgent *UserAgent `json:"userAgent,omitempty"`
}

// NewEventMetadata instantiates a new EventMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMetadata() *EventMetadata {
	this := EventMetadata{}
	return &this
}

// NewEventMetadataWithDefaults instantiates a new EventMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMetadataWithDefaults() *EventMetadata {
	this := EventMetadata{}
	return &this
}

// GetClickedURL returns the ClickedURL field value if set, zero value otherwise.
func (o *EventMetadata) GetClickedURL() string {
	if o == nil || IsNil(o.ClickedURL) {
		var ret string
		return ret
	}
	return *o.ClickedURL
}

// GetClickedURLOk returns a tuple with the ClickedURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetClickedURLOk() (*string, bool) {
	if o == nil || IsNil(o.ClickedURL) {
		return nil, false
	}
	return o.ClickedURL, true
}

// HasClickedURL returns a boolean if a field has been set.
func (o *EventMetadata) HasClickedURL() bool {
	if o != nil && !IsNil(o.ClickedURL) {
		return true
	}

	return false
}

// SetClickedURL gets a reference to the given string and assigns it to the ClickedURL field.
func (o *EventMetadata) SetClickedURL(v string) {
	o.ClickedURL = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *EventMetadata) GetDevice() Device {
	if o == nil || IsNil(o.Device) {
		var ret Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetDeviceOk() (*Device, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *EventMetadata) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Device and assigns it to the Device field.
func (o *EventMetadata) SetDevice(v Device) {
	o.Device = &v
}

// GetGeo returns the Geo field value if set, zero value otherwise.
func (o *EventMetadata) GetGeo() City {
	if o == nil || IsNil(o.Geo) {
		var ret City
		return ret
	}
	return *o.Geo
}

// GetGeoOk returns a tuple with the Geo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetGeoOk() (*City, bool) {
	if o == nil || IsNil(o.Geo) {
		return nil, false
	}
	return o.Geo, true
}

// HasGeo returns a boolean if a field has been set.
func (o *EventMetadata) HasGeo() bool {
	if o != nil && !IsNil(o.Geo) {
		return true
	}

	return false
}

// SetGeo gets a reference to the given City and assigns it to the Geo field.
func (o *EventMetadata) SetGeo(v City) {
	o.Geo = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *EventMetadata) GetOs() Os {
	if o == nil || IsNil(o.Os) {
		var ret Os
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetOsOk() (*Os, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *EventMetadata) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given Os and assigns it to the Os field.
func (o *EventMetadata) SetOs(v Os) {
	o.Os = &v
}

// GetSmtpCode returns the SmtpCode field value if set, zero value otherwise.
func (o *EventMetadata) GetSmtpCode() int64 {
	if o == nil || IsNil(o.SmtpCode) {
		var ret int64
		return ret
	}
	return *o.SmtpCode
}

// GetSmtpCodeOk returns a tuple with the SmtpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetSmtpCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.SmtpCode) {
		return nil, false
	}
	return o.SmtpCode, true
}

// HasSmtpCode returns a boolean if a field has been set.
func (o *EventMetadata) HasSmtpCode() bool {
	if o != nil && !IsNil(o.SmtpCode) {
		return true
	}

	return false
}

// SetSmtpCode gets a reference to the given int64 and assigns it to the SmtpCode field.
func (o *EventMetadata) SetSmtpCode(v int64) {
	o.SmtpCode = &v
}

// GetSmtpDescription returns the SmtpDescription field value if set, zero value otherwise.
func (o *EventMetadata) GetSmtpDescription() string {
	if o == nil || IsNil(o.SmtpDescription) {
		var ret string
		return ret
	}
	return *o.SmtpDescription
}

// GetSmtpDescriptionOk returns a tuple with the SmtpDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetSmtpDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpDescription) {
		return nil, false
	}
	return o.SmtpDescription, true
}

// HasSmtpDescription returns a boolean if a field has been set.
func (o *EventMetadata) HasSmtpDescription() bool {
	if o != nil && !IsNil(o.SmtpDescription) {
		return true
	}

	return false
}

// SetSmtpDescription gets a reference to the given string and assigns it to the SmtpDescription field.
func (o *EventMetadata) SetSmtpDescription(v string) {
	o.SmtpDescription = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *EventMetadata) GetUserAgent() UserAgent {
	if o == nil || IsNil(o.UserAgent) {
		var ret UserAgent
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetUserAgentOk() (*UserAgent, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *EventMetadata) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given UserAgent and assigns it to the UserAgent field.
func (o *EventMetadata) SetUserAgent(v UserAgent) {
	o.UserAgent = &v
}

func (o EventMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClickedURL) {
		toSerialize["clickedURL"] = o.ClickedURL
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Geo) {
		toSerialize["geo"] = o.Geo
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.SmtpCode) {
		toSerialize["smtpCode"] = o.SmtpCode
	}
	if !IsNil(o.SmtpDescription) {
		toSerialize["smtpDescription"] = o.SmtpDescription
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	return toSerialize, nil
}

type NullableEventMetadata struct {
	value *EventMetadata
	isSet bool
}

func (v NullableEventMetadata) Get() *EventMetadata {
	return v.value
}

func (v *NullableEventMetadata) Set(val *EventMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMetadata(val *EventMetadata) *NullableEventMetadata {
	return &NullableEventMetadata{value: val, isSet: true}
}

func (v NullableEventMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

