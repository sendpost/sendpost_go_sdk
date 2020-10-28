# \AccounteventApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventRouterCountAllEventsFromAAccountForAGivenTimeRange**](AccounteventApi.md#EventRouterCountAllEventsFromAAccountForAGivenTimeRange) | **Get** /account/event/count | 
[**EventRouterCountAllEventsFromANodeOfASubAccountForAGivenTimeRange**](AccounteventApi.md#EventRouterCountAllEventsFromANodeOfASubAccountForAGivenTimeRange) | **Get** /account/event/node/count | 
[**EventRouterGet**](AccounteventApi.md#EventRouterGet) | **Get** /account/event/{eventId} | 
[**EventRouterGetAllEventTimestampKeysOfASubAccountFromASpecificNodeForAGivenTimeRange**](AccounteventApi.md#EventRouterGetAllEventTimestampKeysOfASubAccountFromASpecificNodeForAGivenTimeRange) | **Get** /account/event/node/timestampkeys | 
[**EventRouterGetAllEventsOfAAccountFromASpecificNode**](AccounteventApi.md#EventRouterGetAllEventsOfAAccountFromASpecificNode) | **Post** /account/event/node | 
[**EventRouterGetEventInNode**](AccounteventApi.md#EventRouterGetEventInNode) | **Get** /account/event/node/{eventId} | 


# **EventRouterCountAllEventsFromAAccountForAGivenTimeRange**
> ModelsCountStat EventRouterCountAllEventsFromAAccountForAGivenTimeRange(ctx, xAccountApiKey, optional)


Count all events from a account for a given time-range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccounteventApiEventRouterCountAllEventsFromAAccountForAGivenTimeRangeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccounteventApiEventRouterCountAllEventsFromAAccountForAGivenTimeRangeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| search term | 
 **type_** | **optional.String**| search type | 
 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 
 **source** | **optional.String**| data source from which to get timestamp keys subaccount or ip | 
 **sourceId** | **optional.String**| source id from which to get timestamp keys subaccount or ip | 

### Return type

[**ModelsCountStat**](models.CountStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventRouterCountAllEventsFromANodeOfASubAccountForAGivenTimeRange**
> ModelsCountStat EventRouterCountAllEventsFromANodeOfASubAccountForAGivenTimeRange(ctx, xAccountApiKey, optional)


Count all events from a node of a sub-account for a given time-range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccounteventApiEventRouterCountAllEventsFromANodeOfASubAccountForAGivenTimeRangeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccounteventApiEventRouterCountAllEventsFromANodeOfASubAccountForAGivenTimeRangeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| search term | 
 **type_** | **optional.String**| search type | 
 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 
 **source** | **optional.String**| data source from which to get timestamp keys subaccount or ip | 
 **sourceId** | **optional.String**| source id from which to get timestamp keys subaccount or ip | 

### Return type

[**ModelsCountStat**](models.CountStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventRouterGet**
> ModelsQEvent EventRouterGet(ctx, xAccountApiKey, eventId)


Find Event By Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **eventId** | **string**| the eventId that you want to retrieve | 

### Return type

[**ModelsQEvent**](models.QEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventRouterGetAllEventTimestampKeysOfASubAccountFromASpecificNodeForAGivenTimeRange**
> []ModelsQEvent EventRouterGetAllEventTimestampKeysOfASubAccountFromASpecificNodeForAGivenTimeRange(ctx, xAccountApiKey, optional)


Find all events of a sub-account from a specific node for a give time-range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccounteventApiEventRouterGetAllEventTimestampKeysOfASubAccountFromASpecificNodeForAGivenTimeRangeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccounteventApiEventRouterGetAllEventTimestampKeysOfASubAccountFromASpecificNodeForAGivenTimeRangeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| search term | 
 **type_** | **optional.String**| search type | 
 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 
 **source** | **optional.String**| data source from which to get timestamp keys subaccount or ip | 
 **sourceId** | **optional.String**| source id from which to get timestamp keys subaccount or ip | 

### Return type

[**[]ModelsQEvent**](models.QEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventRouterGetAllEventsOfAAccountFromASpecificNode**
> []ModelsQEvent EventRouterGetAllEventsOfAAccountFromASpecificNode(ctx, xAccountApiKey)


Find all events of a account from a specific node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**[]ModelsQEvent**](models.QEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventRouterGetEventInNode**
> ModelsQEvent EventRouterGetEventInNode(ctx, xAccountApiKey, eventId)


Find Event From Node by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **eventId** | **string**| the eventId that you want to retrieve | 

### Return type

[**ModelsQEvent**](models.QEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

