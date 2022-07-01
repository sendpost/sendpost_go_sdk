# \ClusterApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterRouterAddItemsToSuppressionFilterOfEveryNodeInCluster**](ClusterApi.md#ClusterRouterAddItemsToSuppressionFilterOfEveryNodeInCluster) | **Post** /cluster/suppression/filter | 
[**ClusterRouterDeleteItemFromCacheOfEveryNodeInCluster**](ClusterApi.md#ClusterRouterDeleteItemFromCacheOfEveryNodeInCluster) | **Delete** /cluster/cache | 
[**ClusterRouterDeleteItemsFromSuppressionFilterOfEveryNodeInCluster**](ClusterApi.md#ClusterRouterDeleteItemsFromSuppressionFilterOfEveryNodeInCluster) | **Delete** /cluster/suppression/filter | 
[**ClusterRouterGetItemFromCacheOfEveryNodeInCluster**](ClusterApi.md#ClusterRouterGetItemFromCacheOfEveryNodeInCluster) | **Get** /cluster/cache | 
[**ClusterRouterGetItemFromCacheOfSpecificNodeInCluster**](ClusterApi.md#ClusterRouterGetItemFromCacheOfSpecificNodeInCluster) | **Get** /cluster/cache/node | 
[**ClusterRouterGetItemFromCacheOfSpecificNodeInCluster_0**](ClusterApi.md#ClusterRouterGetItemFromCacheOfSpecificNodeInCluster_0) | **Delete** /cluster/cache/node | 


# **ClusterRouterAddItemsToSuppressionFilterOfEveryNodeInCluster**
> ClusterRouterAddItemsToSuppressionFilterOfEveryNodeInCluster(ctx, body)


Add items to suppression filter of every node in cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsSuppression**](ModelsSuppression.md)| Add suppressions to suppression filter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterRouterDeleteItemFromCacheOfEveryNodeInCluster**
> ClusterRouterDeleteItemFromCacheOfEveryNodeInCluster(ctx, xSystemApiKey, optional)


Delete item from cache of every node in cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSystemApiKey** | **string**| System API Key | 
 **optional** | ***ClusterApiClusterRouterDeleteItemFromCacheOfEveryNodeInClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiClusterRouterDeleteItemFromCacheOfEveryNodeInClusterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| cache name | 
 **key** | **optional.String**| cache item key to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterRouterDeleteItemsFromSuppressionFilterOfEveryNodeInCluster**
> ClusterRouterDeleteItemsFromSuppressionFilterOfEveryNodeInCluster(ctx, body)


Delete items from suppression filter of every node in cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsSuppression**](ModelsSuppression.md)| Delete suppressions from suppression filter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterRouterGetItemFromCacheOfEveryNodeInCluster**
> ModelsAllClusterCache ClusterRouterGetItemFromCacheOfEveryNodeInCluster(ctx, xSystemApiKey, optional)


Get item from cache of every node in cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSystemApiKey** | **string**| System API Key | 
 **optional** | ***ClusterApiClusterRouterGetItemFromCacheOfEveryNodeInClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiClusterRouterGetItemFromCacheOfEveryNodeInClusterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| cache name | 
 **key** | **optional.String**| cache item key to delete | 

### Return type

[**ModelsAllClusterCache**](models.AllClusterCache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterRouterGetItemFromCacheOfSpecificNodeInCluster**
> ModelsClusterCache ClusterRouterGetItemFromCacheOfSpecificNodeInCluster(ctx, xSystemApiKey, optional)


Delete item from cache of specific node in cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSystemApiKey** | **string**| System API Key | 
 **optional** | ***ClusterApiClusterRouterGetItemFromCacheOfSpecificNodeInClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiClusterRouterGetItemFromCacheOfSpecificNodeInClusterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| cache name | 
 **key** | **optional.String**| cache item key to delete | 

### Return type

[**ModelsClusterCache**](models.ClusterCache.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterRouterGetItemFromCacheOfSpecificNodeInCluster_0**
> ClusterRouterGetItemFromCacheOfSpecificNodeInCluster_0(ctx, xSystemApiKey, optional)


Get item from cache of specific node in cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSystemApiKey** | **string**| System API Key | 
 **optional** | ***ClusterApiClusterRouterGetItemFromCacheOfSpecificNodeInCluster_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiClusterRouterGetItemFromCacheOfSpecificNodeInCluster_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| cache name | 
 **key** | **optional.String**| cache item key to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

