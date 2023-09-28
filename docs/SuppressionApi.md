# \SuppressionAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Count**](SuppressionAPI.md#Count) | **Get** /subaccount/suppression/count | 
[**CreateSuppressions**](SuppressionAPI.md#CreateSuppressions) | **Post** /subaccount/suppression/ | 
[**DeleteSuppression**](SuppressionAPI.md#DeleteSuppression) | **Delete** /subaccount/suppression/ | 
[**GetSuppressions**](SuppressionAPI.md#GetSuppressions) | **Get** /subaccount/suppression/ | 



## Count

> CountStat Count(ctx).XSubAccountApiKey(xSubAccountApiKey).From(from).To(to).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sendpost/sendpost_go_sdk"
)

func main() {
    xSubAccountApiKey := "xSubAccountApiKey_example" // string | Sub-Account API Key
    from := "from_example" // string | from date (optional)
    to := "to_example" // string | to date (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuppressionAPI.Count(context.Background()).XSubAccountApiKey(xSubAccountApiKey).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuppressionAPI.Count``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Count`: CountStat
    fmt.Fprintf(os.Stdout, "Response from `SuppressionAPI.Count`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSubAccountApiKey** | **string** | Sub-Account API Key | 
 **from** | **string** | from date | 
 **to** | **string** | to date | 

### Return type

[**CountStat**](CountStat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSuppressions

> []Suppression CreateSuppressions(ctx).XSubAccountApiKey(xSubAccountApiKey).RSuppression(rSuppression).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sendpost/sendpost_go_sdk"
)

func main() {
    xSubAccountApiKey := "xSubAccountApiKey_example" // string | Sub-Account API Key
    rSuppression := *openapiclient.NewRSuppression() // RSuppression | Suppression content (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuppressionAPI.CreateSuppressions(context.Background()).XSubAccountApiKey(xSubAccountApiKey).RSuppression(rSuppression).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuppressionAPI.CreateSuppressions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSuppressions`: []Suppression
    fmt.Fprintf(os.Stdout, "Response from `SuppressionAPI.CreateSuppressions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSubAccountApiKey** | **string** | Sub-Account API Key | 
 **rSuppression** | [**RSuppression**](RSuppression.md) | Suppression content | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSuppression

> []DeleteResponse DeleteSuppression(ctx).XSubAccountApiKey(xSubAccountApiKey).RDSuppression(rDSuppression).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sendpost/sendpost_go_sdk"
)

func main() {
    xSubAccountApiKey := "xSubAccountApiKey_example" // string | Sub-Account API Key
    rDSuppression := *openapiclient.NewRDSuppression() // RDSuppression | Suppression content (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuppressionAPI.DeleteSuppression(context.Background()).XSubAccountApiKey(xSubAccountApiKey).RDSuppression(rDSuppression).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuppressionAPI.DeleteSuppression``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSuppression`: []DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `SuppressionAPI.DeleteSuppression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSubAccountApiKey** | **string** | Sub-Account API Key | 
 **rDSuppression** | [**RDSuppression**](RDSuppression.md) | Suppression content | 

### Return type

[**[]DeleteResponse**](DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressions

> []Suppression GetSuppressions(ctx).XSubAccountApiKey(xSubAccountApiKey).Offset(offset).Limit(limit).Search(search).From(from).To(to).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sendpost/sendpost_go_sdk"
)

func main() {
    xSubAccountApiKey := "xSubAccountApiKey_example" // string | Sub-Account API Key
    offset := int64(789) // int64 | offset (optional)
    limit := int64(789) // int64 | limit (optional)
    search := "search_example" // string | search (optional)
    from := "from_example" // string | from date (optional)
    to := "to_example" // string | to date (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuppressionAPI.GetSuppressions(context.Background()).XSubAccountApiKey(xSubAccountApiKey).Offset(offset).Limit(limit).Search(search).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuppressionAPI.GetSuppressions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuppressions`: []Suppression
    fmt.Fprintf(os.Stdout, "Response from `SuppressionAPI.GetSuppressions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSubAccountApiKey** | **string** | Sub-Account API Key | 
 **offset** | **int64** | offset | 
 **limit** | **int64** | limit | 
 **search** | **string** | search | 
 **from** | **string** | from date | 
 **to** | **string** | to date | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

