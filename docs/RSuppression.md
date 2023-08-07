# RSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardBounce** | Pointer to [**[]SuppressionEmail**](SuppressionEmail.md) |  | [optional] 
**Manual** | Pointer to [**[]SuppressionEmail**](SuppressionEmail.md) |  | [optional] 
**SpamComplaint** | Pointer to [**[]SuppressionEmail**](SuppressionEmail.md) |  | [optional] 
**Unsubscribe** | Pointer to [**[]SuppressionEmail**](SuppressionEmail.md) |  | [optional] 

## Methods

### NewRSuppression

`func NewRSuppression() *RSuppression`

NewRSuppression instantiates a new RSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRSuppressionWithDefaults

`func NewRSuppressionWithDefaults() *RSuppression`

NewRSuppressionWithDefaults instantiates a new RSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardBounce

`func (o *RSuppression) GetHardBounce() []SuppressionEmail`

GetHardBounce returns the HardBounce field if non-nil, zero value otherwise.

### GetHardBounceOk

`func (o *RSuppression) GetHardBounceOk() (*[]SuppressionEmail, bool)`

GetHardBounceOk returns a tuple with the HardBounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounce

`func (o *RSuppression) SetHardBounce(v []SuppressionEmail)`

SetHardBounce sets HardBounce field to given value.

### HasHardBounce

`func (o *RSuppression) HasHardBounce() bool`

HasHardBounce returns a boolean if a field has been set.

### GetManual

`func (o *RSuppression) GetManual() []SuppressionEmail`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *RSuppression) GetManualOk() (*[]SuppressionEmail, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *RSuppression) SetManual(v []SuppressionEmail)`

SetManual sets Manual field to given value.

### HasManual

`func (o *RSuppression) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetSpamComplaint

`func (o *RSuppression) GetSpamComplaint() []SuppressionEmail`

GetSpamComplaint returns the SpamComplaint field if non-nil, zero value otherwise.

### GetSpamComplaintOk

`func (o *RSuppression) GetSpamComplaintOk() (*[]SuppressionEmail, bool)`

GetSpamComplaintOk returns a tuple with the SpamComplaint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamComplaint

`func (o *RSuppression) SetSpamComplaint(v []SuppressionEmail)`

SetSpamComplaint sets SpamComplaint field to given value.

### HasSpamComplaint

`func (o *RSuppression) HasSpamComplaint() bool`

HasSpamComplaint returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *RSuppression) GetUnsubscribe() []SuppressionEmail`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *RSuppression) GetUnsubscribeOk() (*[]SuppressionEmail, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *RSuppression) SetUnsubscribe(v []SuppressionEmail)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *RSuppression) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


