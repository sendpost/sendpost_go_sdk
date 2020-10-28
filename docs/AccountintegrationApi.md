# \AccountintegrationApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountIntegrationRouterCount**](AccountintegrationApi.md#AccountIntegrationRouterCount) | **Get** /account/integration/count | 
[**AccountIntegrationRouterCreate**](AccountintegrationApi.md#AccountIntegrationRouterCreate) | **Post** /account/integration/{itype} | 
[**AccountIntegrationRouterDelete**](AccountintegrationApi.md#AccountIntegrationRouterDelete) | **Delete** /account/integration/{itype} | 
[**AccountIntegrationRouterDisableGlockappsIPMonitoring**](AccountintegrationApi.md#AccountIntegrationRouterDisableGlockappsIPMonitoring) | **Delete** /account/integration/glockapps/monitor/{ipid} | 
[**AccountIntegrationRouterEnableGlockappsIPMonitoring**](AccountintegrationApi.md#AccountIntegrationRouterEnableGlockappsIPMonitoring) | **Post** /account/integration/glockapps/monitor/{ipid} | 
[**AccountIntegrationRouterGetAll**](AccountintegrationApi.md#AccountIntegrationRouterGetAll) | **Get** /account/integration/ | 
[**AccountIntegrationRouterGetMonitoredIPStats**](AccountintegrationApi.md#AccountIntegrationRouterGetMonitoredIPStats) | **Get** /account/integration/glockapps/monitor/stat/{ipid} | 
[**AccountIntegrationRouterUpdate**](AccountintegrationApi.md#AccountIntegrationRouterUpdate) | **Put** /account/integration/{itype} | 


# **AccountIntegrationRouterCount**
> ModelsCountStat AccountIntegrationRouterCount(ctx, xAccountApiKey)


Count Total AccountIntegrations

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

# **AccountIntegrationRouterCreate**
> ModelsIntegration AccountIntegrationRouterCreate(ctx, xAccountApiKey, itype, body)


Create Integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **itype** | **string**| The integration type you want to create | 
  **body** | [**ModelsEIntegration**](ModelsEIntegration.md)| The Integration content | 

### Return type

[**ModelsIntegration**](models.Integration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountIntegrationRouterDelete**
> ModelsDeleteResponse AccountIntegrationRouterDelete(ctx, xAccountApiKey, itype)


Delete Integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **itype** | **string**| The integration type you want to update | 

### Return type

[**ModelsDeleteResponse**](models.DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountIntegrationRouterDisableGlockappsIPMonitoring**
> ModelsDeleteResponse AccountIntegrationRouterDisableGlockappsIPMonitoring(ctx, xAccountApiKey, ipid)


Disable IP Monitoring for a single IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **ipid** | **int64**| the IPId you want to disable monitoring for | 

### Return type

[**ModelsDeleteResponse**](models.DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountIntegrationRouterEnableGlockappsIPMonitoring**
> ModelsResponse AccountIntegrationRouterEnableGlockappsIPMonitoring(ctx, xAccountApiKey, ipid)


Enable IP Monitoring for a single IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **ipid** | **int64**| the IPId you want to enable monitoring for | 

### Return type

[**ModelsResponse**](models.Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountIntegrationRouterGetAll**
> []ModelsIntegration AccountIntegrationRouterGetAll(ctx, xAccountApiKey)


Get All Integrations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**[]ModelsIntegration**](models.Integration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountIntegrationRouterGetMonitoredIPStats**
> []ModelsRGlockappsMonitorStat AccountIntegrationRouterGetMonitoredIPStats(ctx, xAccountApiKey, ipid, optional)


Get Monitored IP Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **ipid** | **int64**| the IPId for which you want monitored stats | 
 **optional** | ***AccountintegrationApiAccountIntegrationRouterGetMonitoredIPStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountintegrationApiAccountIntegrationRouterGetMonitoredIPStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 

### Return type

[**[]ModelsRGlockappsMonitorStat**](models.RGlockappsMonitorStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountIntegrationRouterUpdate**
> ModelsIntegration AccountIntegrationRouterUpdate(ctx, xAccountApiKey, itype, body)


Update Integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **itype** | **string**| The integration type you want to update | 
  **body** | [**ModelsEIntegration**](ModelsEIntegration.md)| The Integration content | 

### Return type

[**ModelsIntegration**](models.Integration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

