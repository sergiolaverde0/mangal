# \MiscAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOnlineUsers**](MiscAPI.md#ListOnlineUsers) | **Get** /misc/online | list online users
[**RetrieveSlowTransactionStatus**](MiscAPI.md#RetrieveSlowTransactionStatus) | **Get** /misc/slow-transaction-status/{transaction_id} | get the status of a bulk transaction
[**SiteStats**](MiscAPI.md#SiteStats) | **Get** /misc/stats | show various site stats
[**Time**](MiscAPI.md#Time) | **Get** /misc/time | get the current time



## ListOnlineUsers

> MiscOnlineUsersModelV1 ListOnlineUsers(ctx).Execute()

list online users

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/belphemur/mangal/mangaupdates"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscAPI.ListOnlineUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.ListOnlineUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOnlineUsers`: MiscOnlineUsersModelV1
    fmt.Fprintf(os.Stdout, "Response from `MiscAPI.ListOnlineUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOnlineUsersRequest struct via the builder pattern


### Return type

[**MiscOnlineUsersModelV1**](MiscOnlineUsersModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSlowTransactionStatus

> MiscSlowTransactionStatusResponseV1 RetrieveSlowTransactionStatus(ctx, transactionId).Execute()

get the status of a bulk transaction

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/belphemur/mangal/mangaupdates"
)

func main() {
    transactionId := "transactionId_example" // string | the transaction id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscAPI.RetrieveSlowTransactionStatus(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.RetrieveSlowTransactionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSlowTransactionStatus`: MiscSlowTransactionStatusResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MiscAPI.RetrieveSlowTransactionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | the transaction id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSlowTransactionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MiscSlowTransactionStatusResponseV1**](MiscSlowTransactionStatusResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteStats

> MiscStatsModelV1 SiteStats(ctx).Execute()

show various site stats

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/belphemur/mangal/mangaupdates"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscAPI.SiteStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.SiteStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteStats`: MiscStatsModelV1
    fmt.Fprintf(os.Stdout, "Response from `MiscAPI.SiteStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSiteStatsRequest struct via the builder pattern


### Return type

[**MiscStatsModelV1**](MiscStatsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Time

> TimeV1 Time(ctx).Execute()

get the current time

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/belphemur/mangal/mangaupdates"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscAPI.Time(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscAPI.Time``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Time`: TimeV1
    fmt.Fprintf(os.Stdout, "Response from `MiscAPI.Time`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTimeRequest struct via the builder pattern


### Return type

[**TimeV1**](TimeV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

