# \SmtpApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SMTPRouterReceiveWebhooksRaisedFromSMTPServers**](SmtpApi.md#SMTPRouterReceiveWebhooksRaisedFromSMTPServers) | **Post** /smtp/webhook | 


# **SMTPRouterReceiveWebhooksRaisedFromSMTPServers**
> SMTPRouterReceiveWebhooksRaisedFromSMTPServers(ctx, body)


Receive webhooks raised from SMTP servers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsWMessage**](ModelsWMessage.md)| The Webhook content | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

