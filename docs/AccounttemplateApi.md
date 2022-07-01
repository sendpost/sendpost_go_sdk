# \AccounttemplateApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountTemplateRouterCopy**](AccounttemplateApi.md#AccountTemplateRouterCopy) | **Post** /account/template/{templateid}/copy | 
[**AccountTemplateRouterCount**](AccounttemplateApi.md#AccountTemplateRouterCount) | **Get** /account/template/count | 
[**AccountTemplateRouterCreate**](AccounttemplateApi.md#AccountTemplateRouterCreate) | **Post** /account/template/ | 
[**AccountTemplateRouterDelete**](AccounttemplateApi.md#AccountTemplateRouterDelete) | **Delete** /account/template/{templateid} | 
[**AccountTemplateRouterGet**](AccounttemplateApi.md#AccountTemplateRouterGet) | **Get** /account/template/{templateid} | 
[**AccountTemplateRouterGetAll**](AccounttemplateApi.md#AccountTemplateRouterGetAll) | **Get** /account/template/ | 
[**AccountTemplateRouterUpdate**](AccounttemplateApi.md#AccountTemplateRouterUpdate) | **Put** /account/template/{templateid} | 


# **AccountTemplateRouterCopy**
> ModelsAccountTemplate AccountTemplateRouterCopy(ctx, xAccountApiKey, templateid)


Copy AccountTemplate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **templateid** | **int64**| The id of the template you want to copy | 

### Return type

[**ModelsAccountTemplate**](models.AccountTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountTemplateRouterCount**
> ModelsCountStat AccountTemplateRouterCount(ctx, xAccountApiKey, optional)


Count total templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccounttemplateApiAccountTemplateRouterCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccounttemplateApiAccountTemplateRouterCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| search term | 

### Return type

[**ModelsCountStat**](models.CountStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountTemplateRouterCreate**
> ModelsAccountTemplate AccountTemplateRouterCreate(ctx, xAccountApiKey, body)


Create a new template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsAccountTemplate**](ModelsAccountTemplate.md)| The AccountTemplate content | 

### Return type

[**ModelsAccountTemplate**](models.AccountTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountTemplateRouterDelete**
> ModelsDeleteResponse AccountTemplateRouterDelete(ctx, xAccountApiKey, templateid)


Delete AccountTemplate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **templateid** | **int64**| The id of the template you want to delete | 

### Return type

[**ModelsDeleteResponse**](models.DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountTemplateRouterGet**
> ModelsAccountTemplate AccountTemplateRouterGet(ctx, xAccountApiKey, templateid)


Get single template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **templateid** | **int64**| ID of the template required | 

### Return type

[**ModelsAccountTemplate**](models.AccountTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountTemplateRouterGetAll**
> []ModelsAccountTemplate AccountTemplateRouterGetAll(ctx, xAccountApiKey, optional)


Get all templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccounttemplateApiAccountTemplateRouterGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccounttemplateApiAccountTemplateRouterGetAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 
 **search** | **optional.String**| search term | 

### Return type

[**[]ModelsAccountTemplate**](models.AccountTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountTemplateRouterUpdate**
> ModelsAccountTemplate AccountTemplateRouterUpdate(ctx, xAccountApiKey, templateid, body)


update template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **templateid** | **int64**| The id of the template you want to update | 
  **body** | [**ModelsAccountTemplate**](ModelsAccountTemplate.md)| The template content | 

### Return type

[**ModelsAccountTemplate**](models.AccountTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

