# \AccountmailReportApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MailReportRouterMailReport**](AccountmailReportApi.md#MailReportRouterMailReport) | **Post** /account/mailReport/ | 
[**MailReportRouterMailReportGetSingleReport**](AccountmailReportApi.md#MailReportRouterMailReportGetSingleReport) | **Get** /account/mailReport/{reportId} | 
[**MailReportRouterMailReportProviders**](AccountmailReportApi.md#MailReportRouterMailReportProviders) | **Get** /account/mailReport/provider | 


# **MailReportRouterMailReport**
> ModelsMailReportResult MailReportRouterMailReport(ctx, xAccountApiKey, body)


get reputation of domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsCreateMailReport**](ModelsCreateMailReport.md)| The IP Email Provider Settings | 

### Return type

[**ModelsMailReportResult**](models.MailReportResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailReportRouterMailReportGetSingleReport**
> ApiGlockappsMailReport MailReportRouterMailReportGetSingleReport(ctx, xAccountApiKey, reportId)


get Providers available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **reportId** | **int64**| the report id you want to get | 

### Return type

[**ApiGlockappsMailReport**](api.GlockappsMailReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailReportRouterMailReportProviders**
> ModelsProviderResult MailReportRouterMailReportProviders(ctx, xAccountApiKey)


get Providers available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**ModelsProviderResult**](models.ProviderResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

