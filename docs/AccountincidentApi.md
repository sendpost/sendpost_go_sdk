# \AccountincidentApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IncidentRouterAdd**](AccountincidentApi.md#IncidentRouterAdd) | **Post** /account/incident/{incidentId}/comment | 
[**IncidentRouterCount**](AccountincidentApi.md#IncidentRouterCount) | **Get** /account/incident/count | 
[**IncidentRouterCreate**](AccountincidentApi.md#IncidentRouterCreate) | **Post** /account/incident/ | 
[**IncidentRouterGetAll**](AccountincidentApi.md#IncidentRouterGetAll) | **Get** /account/incident/ | 
[**IncidentRouterGetAllComments**](AccountincidentApi.md#IncidentRouterGetAllComments) | **Get** /account/incident/{incidentId}/comment | 
[**IncidentRouterGetIncident**](AccountincidentApi.md#IncidentRouterGetIncident) | **Get** /account/incident/{incidentId} | 
[**IncidentRouterUpdate**](AccountincidentApi.md#IncidentRouterUpdate) | **Put** /account/incident/{incidentId} | 


# **IncidentRouterAdd**
> ModelsComment IncidentRouterAdd(ctx, xAccountApiKey, incidentId, body)


Add comment to Incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **incidentId** | **int64**| the incident id | 
  **body** | [**ModelsEComment**](ModelsEComment.md)| The Comment content | 

### Return type

[**ModelsComment**](models.Comment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncidentRouterCount**
> ModelsCountStat IncidentRouterCount(ctx, xAccountApiKey, optional)


Count Total Incidents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccountincidentApiIncidentRouterCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountincidentApiIncidentRouterCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.Int64**| status | 
 **tag** | **optional.Int64**| status | 
 **search** | **optional.String**| search term | 

### Return type

[**ModelsCountStat**](models.CountStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncidentRouterCreate**
> ModelsIncident IncidentRouterCreate(ctx, xAccountApiKey, body)


Create Incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsEIncident**](ModelsEIncident.md)| The Incident content | 

### Return type

[**ModelsIncident**](models.Incident.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncidentRouterGetAll**
> []ModelsIncident IncidentRouterGetAll(ctx, xAccountApiKey, optional)


Get All Incidents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccountincidentApiIncidentRouterGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountincidentApiIncidentRouterGetAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 
 **search** | **optional.String**| search term | 
 **status** | **optional.Int64**| status | 
 **tag** | **optional.Int64**| status | 

### Return type

[**[]ModelsIncident**](models.Incident.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncidentRouterGetAllComments**
> []ModelsComment IncidentRouterGetAllComments(ctx, xAccountApiKey, incidentId)


Get All Comments Associated with Incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **incidentId** | **int64**| the IncidentId you want to get comments for | 

### Return type

[**[]ModelsComment**](models.Comment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncidentRouterGetIncident**
> ModelsIncident IncidentRouterGetIncident(ctx, xAccountApiKey, incidentId)


Find Incident by incidentId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **incidentId** | **int64**| the IncidentId you want to get | 

### Return type

[**ModelsIncident**](models.Incident.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncidentRouterUpdate**
> ModelsIncident IncidentRouterUpdate(ctx, xAccountApiKey, incidentId, body)


Update Incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **incidentId** | **int64**| the Incident Id you want to update | 
  **body** | [**ModelsEIncident**](ModelsEIncident.md)| The Incident content | 

### Return type

[**ModelsIncident**](models.Incident.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

