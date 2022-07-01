# \TrackApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackRouterTrackEmailOpen**](TrackApi.md#TrackRouterTrackEmailOpen) | **Get** /track/open/{accountId}/{subAccountId}/{ipId}/{emailType}/{messageId}/1.png | 
[**TrackRouterTrackLinkClick**](TrackApi.md#TrackRouterTrackLinkClick) | **Get** /track/click/{accountId}/{subAccountId}/{ipId}/{emailType}/{messageId} | 
[**TrackRouterTrackUnsubscribe**](TrackApi.md#TrackRouterTrackUnsubscribe) | **Get** /track/unsubscribe/{accountId}/{subAccountId}/{ipId}/{emailType}/{messageId} | 


# **TrackRouterTrackEmailOpen**
> TrackRouterTrackEmailOpen(ctx, ipId, accountId, subAccountId, messageId, emailType)


Track Email Open

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipId** | **int64**| The id of ip from which this email message was sent | 
  **accountId** | **int64**| The id of account from which email is being sent | 
  **subAccountId** | **int64**| The id of sub-account from which email is being sent | 
  **messageId** | **string**| The UUID of message which was sent | 
  **emailType** | **string**| The type of email such as gmail, yahoo etc. which was sent. This is inferred from to email address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackRouterTrackLinkClick**
> TrackRouterTrackLinkClick(ctx, ipId, accountId, subAccountId, messageId, emailType, redirecturl)


Track Link Click

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipId** | **int64**| The id of ip from which this email message was sent | 
  **accountId** | **int64**| The id of account from which email is being sent | 
  **subAccountId** | **int64**| The id of sub-account from which email is being sent | 
  **messageId** | **string**| The UUID of message which was sent | 
  **emailType** | **string**| The type of email such as gmail, yahoo etc. which was sent. This is inferred from to email address | 
  **redirecturl** | **string**| The encoded redirect URL | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackRouterTrackUnsubscribe**
> TrackRouterTrackUnsubscribe(ctx, ipId, accountId, subAccountId, messageId, emailType)


track link click

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipId** | **int64**| The id of ip from which this email message was sent | 
  **accountId** | **int64**| The id of account from which email is being sent | 
  **subAccountId** | **int64**| The id of sub-account from which email is being sent | 
  **messageId** | **string**| The UUID of message which was sent | 
  **emailType** | **string**| The type of email such as gmail, yahoo etc. which was sent. This is inferred from to email address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

