# Suppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Reason** | Pointer to **int64** |  | [optional] 
**SmtpError** | Pointer to **string** |  | [optional] 

## Methods

### NewSuppression

`func NewSuppression() *Suppression`

NewSuppression instantiates a new Suppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionWithDefaults

`func NewSuppressionWithDefaults() *Suppression`

NewSuppressionWithDefaults instantiates a new Suppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Suppression) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Suppression) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Suppression) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Suppression) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEmail

`func (o *Suppression) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Suppression) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Suppression) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Suppression) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *Suppression) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Suppression) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Suppression) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Suppression) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *Suppression) GetReason() int64`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Suppression) GetReasonOk() (*int64, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Suppression) SetReason(v int64)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Suppression) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSmtpError

`func (o *Suppression) GetSmtpError() string`

GetSmtpError returns the SmtpError field if non-nil, zero value otherwise.

### GetSmtpErrorOk

`func (o *Suppression) GetSmtpErrorOk() (*string, bool)`

GetSmtpErrorOk returns a tuple with the SmtpError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpError

`func (o *Suppression) SetSmtpError(v string)`

SetSmtpError sets SmtpError field to given value.

### HasSmtpError

`func (o *Suppression) HasSmtpError() bool`

HasSmtpError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


