# \AccountmemberApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberRouterDelete**](AccountmemberApi.md#MemberRouterDelete) | **Delete** /account/member/{memberId} | 
[**MemberRouterGet**](AccountmemberApi.md#MemberRouterGet) | **Get** /account/member/{memberId} | 
[**MemberRouterGetAll**](AccountmemberApi.md#MemberRouterGetAll) | **Get** /account/member/ | 


# **MemberRouterDelete**
> ModelsDeleteResponse MemberRouterDelete(ctx, xAccountApiKey, memberId)


Delete Member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **memberId** | **int64**| The MemberId you want to delete | 

### Return type

[**ModelsDeleteResponse**](models.DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberRouterGet**
> ModelsMember MemberRouterGet(ctx, xAccountApiKey, memberId)


Find Member by MemberId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **memberId** | **int64**| the MemberId you want to get | 

### Return type

[**ModelsMember**](models.Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberRouterGetAll**
> []ModelsMember MemberRouterGetAll(ctx, xAccountApiKey)


Get All Members

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**[]ModelsMember**](models.Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

