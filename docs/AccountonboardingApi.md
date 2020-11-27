# \AccountonboardingApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnboardingRouterGetOnboardingChecklist**](AccountonboardingApi.md#OnboardingRouterGetOnboardingChecklist) | **Get** /account/onboarding/checklist | 


# **OnboardingRouterGetOnboardingChecklist**
> ModelsOnboardingChecklist OnboardingRouterGetOnboardingChecklist(ctx, xAccountApiKey)


Gets Onboarding Checklist data for account if not present creates a default entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**ModelsOnboardingChecklist**](models.OnboardingChecklist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

