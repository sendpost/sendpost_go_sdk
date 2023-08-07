# RDSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suppressions** | Pointer to [**[]SuppressionEmail**](SuppressionEmail.md) |  | [optional] 

## Methods

### NewRDSuppression

`func NewRDSuppression() *RDSuppression`

NewRDSuppression instantiates a new RDSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSuppressionWithDefaults

`func NewRDSuppressionWithDefaults() *RDSuppression`

NewRDSuppressionWithDefaults instantiates a new RDSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppressions

`func (o *RDSuppression) GetSuppressions() []SuppressionEmail`

GetSuppressions returns the Suppressions field if non-nil, zero value otherwise.

### GetSuppressionsOk

`func (o *RDSuppression) GetSuppressionsOk() (*[]SuppressionEmail, bool)`

GetSuppressionsOk returns a tuple with the Suppressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressions

`func (o *RDSuppression) SetSuppressions(v []SuppressionEmail)`

SetSuppressions sets Suppressions field to given value.

### HasSuppressions

`func (o *RDSuppression) HasSuppressions() bool`

HasSuppressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


