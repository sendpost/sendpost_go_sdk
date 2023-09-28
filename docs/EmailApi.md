# \EmailAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendEmail**](EmailAPI.md#SendEmail) | **Post** /subaccount/email/ | 
[**SendEmailWithTemplate**](EmailAPI.md#SendEmailWithTemplate) | **Post** /subaccount/email/template | 



## SendEmail

> []EmailResponse SendEmail(ctx).XSubAccountApiKey(xSubAccountApiKey).EmailMessage(emailMessage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sendpost/sendpost_go_sdk/v1"
)

func main() {
    xSubAccountApiKey := "xSubAccountApiKey_example" // string | Sub-Account API Key
    emailMessage := *openapiclient.NewEmailMessage() // EmailMessage | Email message (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailAPI.SendEmail(context.Background()).XSubAccountApiKey(xSubAccountApiKey).EmailMessage(emailMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.SendEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendEmail`: []EmailResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailAPI.SendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSubAccountApiKey** | **string** | Sub-Account API Key | 
 **emailMessage** | [**EmailMessage**](EmailMessage.md) | Email message | 

### Return type

[**[]EmailResponse**](EmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailWithTemplate

> []EmailResponse SendEmailWithTemplate(ctx).XSubAccountApiKey(xSubAccountApiKey).EmailMessage(emailMessage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sendpost/sendpost_go_sdk/v1"
)

func main() {
    xSubAccountApiKey := "xSubAccountApiKey_example" // string | Sub-Account API Key
    emailMessage := *openapiclient.NewEmailMessage() // EmailMessage | Email message (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailAPI.SendEmailWithTemplate(context.Background()).XSubAccountApiKey(xSubAccountApiKey).EmailMessage(emailMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.SendEmailWithTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendEmailWithTemplate`: []EmailResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailAPI.SendEmailWithTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailWithTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSubAccountApiKey** | **string** | Sub-Account API Key | 
 **emailMessage** | [**EmailMessage**](EmailMessage.md) | Email message | 

### Return type

[**[]EmailResponse**](EmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

