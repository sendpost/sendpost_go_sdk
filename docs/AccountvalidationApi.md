# \AccountvalidationApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateRouterValidateEmailBulk**](AccountvalidationApi.md#ValidateRouterValidateEmailBulk) | **Post** /account/validation/bulk | 
[**ValidationRouterCount**](AccountvalidationApi.md#ValidationRouterCount) | **Get** /account/validation/count | 
[**ValidationRouterDeleteValidation**](AccountvalidationApi.md#ValidationRouterDeleteValidation) | **Delete** /account/validation/ | 
[**ValidationRouterGetAll**](AccountvalidationApi.md#ValidationRouterGetAll) | **Get** /account/validation/ | 
[**ValidationRouterValidateEmailList**](AccountvalidationApi.md#ValidationRouterValidateEmailList) | **Post** /account/validation/ | 


# **ValidateRouterValidateEmailBulk**
> ModelsBulkResponse ValidateRouterValidateEmailBulk(ctx, fileinput, xAccountApiKey)


Validate Emails In File Asynchronously

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileinput** | ***os.File**| CSV whose emails need to be validated | 
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**ModelsBulkResponse**](models.BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidationRouterCount**
> ModelsCountStat ValidationRouterCount(ctx, xAccountApiKey)


Count Total Validations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**ModelsCountStat**](models.CountStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidationRouterDeleteValidation**
> ModelsValidation ValidationRouterDeleteValidation(ctx, xAccountApiKey, body)


Delete a specific validation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsEValidation**](ModelsEValidation.md)| List of emails to be deleted from validation | 

### Return type

[**ModelsValidation**](models.Validation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidationRouterGetAll**
> []ModelsValidation ValidationRouterGetAll(ctx, xAccountApiKey, optional)


Get all validation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccountvalidationApiValidationRouterGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountvalidationApiValidationRouterGetAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 
 **search** | **optional.String**| search | 

### Return type

[**[]ModelsValidation**](models.Validation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidationRouterValidateEmailList**
> ModelsCleanedList ValidationRouterValidateEmailList(ctx, xAccountApiKey, body)


Validate Email List Synchronously

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsEmailList**](ModelsEmailList.md)| The email list to be sent for being validated | 

### Return type

[**ModelsCleanedList**](models.CleanedList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

