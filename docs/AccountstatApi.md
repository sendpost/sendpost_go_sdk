# \AccountstatApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountStatRouterGetAllAccountStats**](AccountstatApi.md#AccountStatRouterGetAllAccountStats) | **Get** /account/stat/ | 
[**AccountStatRouterGetAllAccountStatsByGroup**](AccountstatApi.md#AccountStatRouterGetAllAccountStatsByGroup) | **Get** /account/stat/group | 
[**AccountStatRouterGetAllAggregateAccountStats**](AccountstatApi.md#AccountStatRouterGetAllAggregateAccountStats) | **Get** /account/stat/aggregate | 
[**AccountStatRouterGetAllAggregateAccountStatsByGroup**](AccountstatApi.md#AccountStatRouterGetAllAggregateAccountStatsByGroup) | **Get** /account/stat/aggregate/group | 
[**AccountStatRouterGetAllAggregateIPStats**](AccountstatApi.md#AccountStatRouterGetAllAggregateIPStats) | **Get** /account/stat/aggregate/ip | 
[**AccountStatRouterGetAllAggregateSubAccountStats**](AccountstatApi.md#AccountStatRouterGetAllAggregateSubAccountStats) | **Get** /account/stat/aggregate/subaccount | 


# **AccountStatRouterGetAllAccountStats**
> []ModelsRStat AccountStatRouterGetAllAccountStats(ctx, xAccountApiKey, optional)


Get All Account Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Sub-Account API Key | 
 **optional** | ***AccountstatApiAccountStatRouterGetAllAccountStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountstatApiAccountStatRouterGetAllAccountStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 

### Return type

[**[]ModelsRStat**](models.RStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountStatRouterGetAllAccountStatsByGroup**
> []ModelsRStat AccountStatRouterGetAllAccountStatsByGroup(ctx, xAccountApiKey, group, optional)


Get All Account Stats by Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Sub-Account API Key | 
  **group** | **string**| the group whose stats you want | 
 **optional** | ***AccountstatApiAccountStatRouterGetAllAccountStatsByGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountstatApiAccountStatRouterGetAllAccountStatsByGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| from date | 
 **to** | **optional.String**| to date | 

### Return type

[**[]ModelsRStat**](models.RStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountStatRouterGetAllAggregateAccountStats**
> ModelsStat AccountStatRouterGetAllAggregateAccountStats(ctx, xAccountApiKey, optional)


Get All Aggregate Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
 **optional** | ***AccountstatApiAccountStatRouterGetAllAggregateAccountStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountstatApiAccountStatRouterGetAllAggregateAccountStatsOpts struct

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

# **AccountStatRouterGetAllAggregateAccountStatsByGroup**
> ModelsStat AccountStatRouterGetAllAggregateAccountStatsByGroup(ctx, xAccountApiKey, group, optional)


Get All Aggregate Stats by Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Sub-Account API Key | 
  **group** | **string**| the group whose stats you want | 
 **optional** | ***AccountstatApiAccountStatRouterGetAllAggregateAccountStatsByGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountstatApiAccountStatRouterGetAllAggregateAccountStatsByGroupOpts struct

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

# **AccountStatRouterGetAllAggregateIPStats**
> []ModelsAgipStat AccountStatRouterGetAllAggregateIPStats(ctx, xAccountApiKey, from, to, optional)


Get All Aggregate IP Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **from** | **string**| from date | 
  **to** | **string**| to date | 
 **optional** | ***AccountstatApiAccountStatRouterGetAllAggregateIPStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountstatApiAccountStatRouterGetAllAggregateIPStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filterBy** | **optional.String**| filterBy | 
 **filterValue** | **optional.Int64**| filterValue | 
 **orderBy** | **optional.String**| orderBy | 
 **sortOrder** | **optional.String**| sortOrder | 
 **search** | **optional.String**| search term | 
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsAgipStat**](models.AGIPStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountStatRouterGetAllAggregateSubAccountStats**
> []ModelsAgSubAccountStat AccountStatRouterGetAllAggregateSubAccountStats(ctx, xAccountApiKey, from, to, optional)


Get All Aggregate SubAccount Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **from** | **string**| from date | 
  **to** | **string**| to date | 
 **optional** | ***AccountstatApiAccountStatRouterGetAllAggregateSubAccountStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountstatApiAccountStatRouterGetAllAggregateSubAccountStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filterBy** | **optional.String**| filterBy | 
 **filterValue** | **optional.Int64**| filterValue | 
 **orderBy** | **optional.String**| orderBy | 
 **sortOrder** | **optional.String**| sortOrder | 
 **search** | **optional.String**| search term | 
 **offset** | **optional.Int64**| offset | 
 **limit** | **optional.Int64**| limit | 

### Return type

[**[]ModelsAgSubAccountStat**](models.AGSubAccountStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

