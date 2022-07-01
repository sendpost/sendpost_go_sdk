# \AccountadminApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountAdminRouterAssumeAccountGetToken**](AccountadminApi.md#AccountAdminRouterAssumeAccountGetToken) | **Get** /account/admin/assume | 


# **AccountAdminRouterAssumeAccountGetToken**
> ModelsRAssumeAccount AccountAdminRouterAssumeAccountGetToken(ctx, xAccountApiKey, email, optional)


Responds back with a custom token for frontend to start login using firebase sdk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **email** | **string**| email | 
 **optional** | ***AccountadminApiAccountAdminRouterAssumeAccountGetTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountadminApiAccountAdminRouterAssumeAccountGetTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uid** | **optional.String**| firebase uid if you have | 

### Return type

[**ModelsRAssumeAccount**](models.RAssumeAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

