# \AccountdomainStatApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainStatRouterGetAllAggregateDomainStatsByGroup**](AccountdomainStatApi.md#DomainStatRouterGetAllAggregateDomainStatsByGroup) | **Get** /account/domainStat/{domainId}/aggregate/provider | 
[**DomainStatRouterGetStatsForASingleDomainStats**](AccountdomainStatApi.md#DomainStatRouterGetStatsForASingleDomainStats) | **Get** /account/domainStat/{domainId}/aggregate | 


# **DomainStatRouterGetAllAggregateDomainStatsByGroup**
> ModelsStat DomainStatRouterGetAllAggregateDomainStatsByGroup(ctx, xAccountApiKey, domainId, provider, optional)


Get All Aggregate Stats by Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **domainId** | **int64**| the domainName you want to get | 
  **provider** | **string**| the group whose stats you want | 
 **optional** | ***AccountdomainStatApiDomainStatRouterGetAllAggregateDomainStatsByGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountdomainStatApiDomainStatRouterGetAllAggregateDomainStatsByGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 

### Return type

[**ModelsStat**](models.Stat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainStatRouterGetStatsForASingleDomainStats**
> ModelsStat DomainStatRouterGetStatsForASingleDomainStats(ctx, xAccountApiKey, domainId, optional)


Get All Aggregate Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **domainId** | **string**| the domainID you want to get | 
 **optional** | ***AccountdomainStatApiDomainStatRouterGetStatsForASingleDomainStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountdomainStatApiDomainStatRouterGetStatsForASingleDomainStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 

### Return type

[**ModelsStat**](models.Stat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

