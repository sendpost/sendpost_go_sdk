# \AccountalertApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertRouterCount**](AccountalertApi.md#AlertRouterCount) | **Get** /account/alert/count | 
[**AlertRouterCreateAlert**](AccountalertApi.md#AlertRouterCreateAlert) | **Post** /account/alert/ | 
[**AlertRouterGetAll**](AccountalertApi.md#AlertRouterGetAll) | **Get** /account/alert/ | 
[**AlertRouterUpdate**](AccountalertApi.md#AlertRouterUpdate) | **Put** /account/alert/{alertid} | 


# **AlertRouterCount**
> ModelsCountStat AlertRouterCount(ctx, xAccountApiKey, optional)


Count Total Alerts for account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccountalertApiAlertRouterCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountalertApiAlertRouterCountOpts struct

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

# **AlertRouterCreateAlert**
> ModelsAlert AlertRouterCreateAlert(ctx, xAccountApiKey, body)


create an alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsAlertRequest**](ModelsAlertRequest.md)| The List to br sent | 

### Return type

[**ModelsAlert**](models.Alert.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertRouterGetAll**
> []ModelsDetailedAlert AlertRouterGetAll(ctx, xAccountApiKey, optional)


Get All Alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccountalertApiAlertRouterGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountalertApiAlertRouterGetAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 
 **search** | **optional.String**| search term | 

### Return type

[**[]ModelsDetailedAlert**](models.DetailedAlert.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertRouterUpdate**
> Alert AlertRouterUpdate(ctx, xAccountApiKey, alertid, body)


Update an Alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **alertid** | **int64**| The alert you want to update | 
  **body** | [**Alert**](Alert.md)| The alert  Settings | 

### Return type

[**Alert**](.alert.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

