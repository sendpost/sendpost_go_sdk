# \AccountvalidateApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateRouterValidateEmailBulk**](AccountvalidateApi.md#ValidateRouterValidateEmailBulk) | **Post** /account/validate/bulk | 
[**ValidateRouterValidateEmailList**](AccountvalidateApi.md#ValidateRouterValidateEmailList) | **Post** /account/validate/ | 


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

# **ValidateRouterValidateEmailList**
> ModelsCleanedList ValidateRouterValidateEmailList(ctx, xAccountApiKey, body)


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

