# \AccountlabelApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LabelRouterCount**](AccountlabelApi.md#LabelRouterCount) | **Get** /account/label/count | 
[**LabelRouterCreate**](AccountlabelApi.md#LabelRouterCreate) | **Post** /account/label/ | 
[**LabelRouterDelete**](AccountlabelApi.md#LabelRouterDelete) | **Delete** /account/label/{labelId} | 
[**LabelRouterGet**](AccountlabelApi.md#LabelRouterGet) | **Get** /account/label/{labelId} | 
[**LabelRouterGetAll**](AccountlabelApi.md#LabelRouterGetAll) | **Get** /account/label/ | 
[**LabelRouterUpdate**](AccountlabelApi.md#LabelRouterUpdate) | **Put** /account/label/{labelId} | 


# **LabelRouterCount**
> ModelsCountStat LabelRouterCount(ctx, xAccountApiKey, type_)


Count Total Labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **type_** | **string**| get all labels type | 

### Return type

[**ModelsCountStat**](models.CountStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelRouterCreate**
> ModelsLabel LabelRouterCreate(ctx, xAccountApiKey, body)


Create Label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsLabel**](ModelsLabel.md)| The Label content | 

### Return type

[**ModelsLabel**](models.Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelRouterDelete**
> ModelsDeleteResponse LabelRouterDelete(ctx, xAccountApiKey, labelId)


Delete Label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **labelId** | **int64**| The LabelId you want to delete | 

### Return type

[**ModelsDeleteResponse**](models.DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelRouterGet**
> ModelsLabel LabelRouterGet(ctx, xAccountApiKey, labelId)


Find Label by LabelId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **labelId** | **int64**| the LabelId you want to get | 

### Return type

[**ModelsLabel**](models.Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelRouterGetAll**
> []ModelsLabel LabelRouterGetAll(ctx, xAccountApiKey, type_, optional)


Get All Labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **type_** | **string**| get all labels type | 
 **optional** | ***AccountlabelApiLabelRouterGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountlabelApiLabelRouterGetAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 
 **search** | **optional.String**| search term | 

### Return type

[**[]ModelsLabel**](models.Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelRouterUpdate**
> ModelsLabel LabelRouterUpdate(ctx, xAccountApiKey, labelId, body)


Update Label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **labelId** | **int64**| The LabelId you want to update | 
  **body** | [**ModelsLabel**](ModelsLabel.md)| The body | 

### Return type

[**ModelsLabel**](models.Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

