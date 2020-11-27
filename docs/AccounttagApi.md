# \AccounttagApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagRouterCreate**](AccounttagApi.md#TagRouterCreate) | **Post** /account/tag/ | 
[**TagRouterDelete**](AccounttagApi.md#TagRouterDelete) | **Delete** /account/tag/{tagid} | 
[**TagRouterGetAll**](AccounttagApi.md#TagRouterGetAll) | **Get** /account/tag/ | 


# **TagRouterCreate**
> ModelsTag TagRouterCreate(ctx, xAccountApiKey, body)


Create Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsTag**](ModelsTag.md)| The Tag content | 

### Return type

[**ModelsTag**](models.Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagRouterDelete**
> ModelsDeleteResponse TagRouterDelete(ctx, xAccountApiKey, tagid)


Delete Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **tagid** | **int64**| The AccountTagId you want to delete | 

### Return type

[**ModelsDeleteResponse**](models.DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagRouterGetAll**
> []ModelsTag TagRouterGetAll(ctx, xAccountApiKey)


Get All Tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**[]ModelsTag**](models.Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

