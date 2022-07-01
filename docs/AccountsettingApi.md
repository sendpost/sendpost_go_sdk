# \AccountsettingApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountSettingRouterUpdate**](AccountsettingApi.md#AccountSettingRouterUpdate) | **Put** /account/setting/ | 


# **AccountSettingRouterUpdate**
> ModelsAccount AccountSettingRouterUpdate(ctx, xAccountApiKey, body)


update account settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsEAccountSetting**](ModelsEAccountSetting.md)| The account settings to be updated | 

### Return type

[**ModelsAccount**](models.Account.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

