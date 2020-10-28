# \AccountmessageApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MessageRouterGet**](AccountmessageApi.md#MessageRouterGet) | **Get** /account/message/{messageId} | 
[**MessageRouterGetAllEventsForAMessageId**](AccountmessageApi.md#MessageRouterGetAllEventsForAMessageId) | **Get** /account/message/{messageId}/events | 
[**MessageRouterGetAllEventsForAMessageIdFromANode**](AccountmessageApi.md#MessageRouterGetAllEventsForAMessageIdFromANode) | **Get** /account/message/node/{messageId}/events | 
[**MessageRouterGetMessageFromNode**](AccountmessageApi.md#MessageRouterGetMessageFromNode) | **Get** /account/message/node/{messageId} | 


# **MessageRouterGet**
> ModelsQEmailMessage MessageRouterGet(ctx, xAccountApiKey, messageId)


Find Message By Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Sub-Account API Key | 
  **messageId** | **string**| the messageId that you want to retrieve | 

### Return type

[**ModelsQEmailMessage**](models.QEmailMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MessageRouterGetAllEventsForAMessageId**
> []ModelsQEvent MessageRouterGetAllEventsForAMessageId(ctx, xAccountApiKey, messageId)


Find all events associated with a message id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **messageId** | **string**| the messageId that you want to retrieve | 

### Return type

[**[]ModelsQEvent**](models.QEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MessageRouterGetAllEventsForAMessageIdFromANode**
> []ModelsQEvent MessageRouterGetAllEventsForAMessageIdFromANode(ctx, xAccountApiKey, messageId)


Find all message events associated with a message id from a specific node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **messageId** | **string**| the messageId that you want to retrieve | 

### Return type

[**[]ModelsQEvent**](models.QEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MessageRouterGetMessageFromNode**
> ModelsQEmailMessage MessageRouterGetMessageFromNode(ctx, xAccountApiKey, messageId)


Find Message from node by specific Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **messageId** | **string**| the messageId that you want to retrieve | 

### Return type

[**ModelsQEmailMessage**](models.QEmailMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

