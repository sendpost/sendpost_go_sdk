# \AccountpaymentApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentRouterCreateCustomerPortal**](AccountpaymentApi.md#PaymentRouterCreateCustomerPortal) | **Post** /account/payment/portal | 
[**PaymentRouterCreatePaymentSubscription**](AccountpaymentApi.md#PaymentRouterCreatePaymentSubscription) | **Post** /account/payment/subscription | 
[**PaymentRouterHandlePaymentWebhook**](AccountpaymentApi.md#PaymentRouterHandlePaymentWebhook) | **Post** /account/payment/webhook | 


# **PaymentRouterCreateCustomerPortal**
> ModelsBillingPortalSession PaymentRouterCreateCustomerPortal(ctx, xAccountApiKey)


Create Customer Portal for account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 

### Return type

[**ModelsBillingPortalSession**](models.BillingPortalSession.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentRouterCreatePaymentSubscription**
> ModelsPaymentStatus PaymentRouterCreatePaymentSubscription(ctx, xAccountApiKey, body)


Create Payment Subscription for Stripe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccountApiKey** | **string**| Account API Key | 
  **body** | [**ModelsPaymentOptions**](ModelsPaymentOptions.md)| PaymentOptions content | 

### Return type

[**ModelsPaymentStatus**](models.PaymentStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentRouterHandlePaymentWebhook**
> PaymentRouterHandlePaymentWebhook(ctx, )


Handle Payment Related Webhooks

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

