# \AccountrecipientApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecipientRouterGetAllMessagesForARecipient**](AccountrecipientApi.md#RecipientRouterGetAllMessagesForARecipient) | **Get** /account/recipient/{recipient}/messages | 
[**RecipientRouterGetAllMessagesForARecipientFromANode**](AccountrecipientApi.md#RecipientRouterGetAllMessagesForARecipientFromANode) | **Get** /account/recipient/node/{recipient}/messages | 


# **RecipientRouterGetAllMessagesForARecipient**
> []ModelsQEmailMessage RecipientRouterGetAllMessagesForARecipient(ctx, xAccountApiKey, recipient)


Find all messages sent to a specific recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **recipient** | **string**| email of the recipient | 

### Return type

[**[]ModelsQEmailMessage**](models.QEmailMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecipientRouterGetAllMessagesForARecipientFromANode**
> []ModelsQEmailMessage RecipientRouterGetAllMessagesForARecipientFromANode(ctx, xAccountApiKey, recipient)


Find all message sent to a recipient from a specific node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **recipient** | **string**| email of the recipient | 

### Return type

[**[]ModelsQEmailMessage**](models.QEmailMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

