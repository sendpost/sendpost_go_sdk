# \ClusterApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterRouterDeleteItemFromCacheOfEveryNodeInCluster**](ClusterApi.md#ClusterRouterDeleteItemFromCacheOfEveryNodeInCluster) | **Delete** /cluster/cache | 
[**ClusterRouterDeleteItemFromCacheOfSpecificNodeInCluster**](ClusterApi.md#ClusterRouterDeleteItemFromCacheOfSpecificNodeInCluster) | **Delete** /cluster/cache/node | 


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

# **ClusterRouterDeleteItemFromCacheOfSpecificNodeInCluster**
> ClusterRouterDeleteItemFromCacheOfSpecificNodeInCluster(ctx, xSystemApiKey, optional)


Delete item from cache of specific node in cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSystemApiKey** | **string**| System API Key | 
 **optional** | ***ClusterApiClusterRouterDeleteItemFromCacheOfSpecificNodeInClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiClusterRouterDeleteItemFromCacheOfSpecificNodeInClusterOpts struct

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

