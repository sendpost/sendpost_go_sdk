# \AuthApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthRouterCreate**](AuthApi.md#AuthRouterCreate) | **Post** /auth/create | 
[**AuthRouterGetAuthInfo**](AuthApi.md#AuthRouterGetAuthInfo) | **Post** /auth/info | 
[**AuthRouterUpdateAuthInfo**](AuthApi.md#AuthRouterUpdateAuthInfo) | **Put** /auth/info | 


# **AuthRouterCreate**
> ModelsAccount AuthRouterCreate(ctx, body, xToken)


Create Account, sub-account and member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsEAccount**](ModelsEAccount.md)| The Account content | 
  **xToken** | **string**| Firebase dynamic token | 

### Return type

[**ModelsAccount**](models.Account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthRouterGetAuthInfo**
> ModelsAuthInfo AuthRouterGetAuthInfo(ctx, body, xToken)


Get Auth Info Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsEAccount**](ModelsEAccount.md)| The Account content | 
  **xToken** | **string**| Firebase dynamic token | 

### Return type

[**ModelsAuthInfo**](models.AuthInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthRouterUpdateAuthInfo**
> ModelsAuthInfo AuthRouterUpdateAuthInfo(ctx, body, xToken)


Updates member and accounts related for email verification and on boarding answers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsEAccountMember**](ModelsEAccountMember.md)| The Account Member content | 
  **xToken** | **string**| Firebase dynamic token | 

### Return type

[**ModelsAuthInfo**](models.AuthInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

