# \AccountinvitationApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvitationRouterCreate**](AccountinvitationApi.md#InvitationRouterCreate) | **Post** /account/invitation/ | 
[**InvitationRouterDelete**](AccountinvitationApi.md#InvitationRouterDelete) | **Delete** /account/invitation/{invitationId} | 
[**InvitationRouterGetAll**](AccountinvitationApi.md#InvitationRouterGetAll) | **Get** /account/invitation/ | 


# **InvitationRouterCreate**
> ModelsInvitation InvitationRouterCreate(ctx, xAccountApiKey, body)


Create Invitation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsEInvitation**](ModelsEInvitation.md)| The Invitation content | 

### Return type

[**ModelsInvitation**](models.Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvitationRouterDelete**
> ModelsDeleteResponse InvitationRouterDelete(ctx, xAccountApiKey, invitationId)


Delete Invitation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **invitationId** | **int64**| The InvitationId you want to delete | 

### Return type

[**ModelsDeleteResponse**](models.DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvitationRouterGetAll**
> []ModelsInvitation InvitationRouterGetAll(ctx, xAccountApiKey)


Get All Invitations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**[]ModelsInvitation**](models.Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

