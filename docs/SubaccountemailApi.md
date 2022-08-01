# \SubaccountemailApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailRouterSendEmail**](SubaccountemailApi.md#EmailRouterSendEmail) | **Post** /subaccount/email/ | 
[**EmailRouterSendEmailWithTemplate**](SubaccountemailApi.md#EmailRouterSendEmailWithTemplate) | **Post** /subaccount/email/template | 


# **EmailRouterSendEmail**
> []ModelsEmailResponse EmailRouterSendEmail(ctx, xSubAccountApiKey, body, optional)


Send Email To Contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSubAccountApiKey** | **string**| Sub-Account API Key | 
  **body** | [**ModelsEmailMessage**](ModelsEmailMessage.md)| The Email Message | 
 **optional** | ***SubaccountemailApiEmailRouterSendEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubaccountemailApiEmailRouterSendEmailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSendPostMockEmail** | **optional.Bool**| Mock email header | 
 **xSendPostMockTimeShift** | **optional.String**| Mock email time shift | 

### Return type

[**[]ModelsEmailResponse**](models.EmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailRouterSendEmailWithTemplate**
> []ModelsEmailResponse EmailRouterSendEmailWithTemplate(ctx, xSubAccountApiKey, body)


Send Email To Contacts With Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSubAccountApiKey** | **string**| Sub-Account API Key | 
  **body** | [**ModelsEmailMessage**](ModelsEmailMessage.md)| The Email Message | 

### Return type

[**[]ModelsEmailResponse**](models.EmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

