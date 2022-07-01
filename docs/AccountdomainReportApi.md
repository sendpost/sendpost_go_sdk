# \AccountdomainReportApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainReportRouterReputation**](AccountdomainReportApi.md#DomainReportRouterReputation) | **Get** /account/domainReport/reputation | 


# **DomainReportRouterReputation**
> ModelsDomainCheckResult DomainReportRouterReputation(ctx, xAccountApiKey, optional)


get reputation of domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccountdomainReportApiDomainReportRouterReputationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountdomainReportApiDomainReportRouterReputationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ip** | **optional.String**| ip | 
 **host** | **optional.String**| host | 

### Return type

[**ModelsDomainCheckResult**](models.DomainCheckResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

