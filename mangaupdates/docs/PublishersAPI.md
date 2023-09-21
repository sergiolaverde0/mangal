# \PublishersAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPublisher**](PublishersAPI.md#AddPublisher) | **Post** /publishers | add an publisher
[**DeletePublisher**](PublishersAPI.md#DeletePublisher) | **Delete** /publishers/{id} | delete a publisher
[**RetrievePublicationSeries**](PublishersAPI.md#RetrievePublicationSeries) | **Get** /publishers/publication | get the list of series for a specific publication
[**RetrievePublisher**](PublishersAPI.md#RetrievePublisher) | **Get** /publishers/{id} | get a specific publisher
[**RetrievePublisherSeries**](PublishersAPI.md#RetrievePublisherSeries) | **Get** /publishers/{id}/series | get the list of series for a specific publisher
[**SearchPublishersPost**](PublishersAPI.md#SearchPublishersPost) | **Post** /publishers/search | search publishers
[**UpdatePublisher**](PublishersAPI.md#UpdatePublisher) | **Patch** /publishers/{id} | update a publisher



## AddPublisher

> ApiResponseV1 AddPublisher(ctx).PublishersModelUpdateV1(publishersModelUpdateV1).Execute()

add an publisher

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
    publishersModelUpdateV1 := *openapiclient.NewPublishersModelUpdateV1() // PublishersModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublishersAPI.AddPublisher(context.Background()).PublishersModelUpdateV1(publishersModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.AddPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPublisher`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.AddPublisher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishersModelUpdateV1** | [**PublishersModelUpdateV1**](PublishersModelUpdateV1.md) |  | 

### Return type

[**ApiResponseV1**](ApiResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublisher

> ApiResponseV1 DeletePublisher(ctx, id).Execute()

delete a publisher

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
    id := int64(56) // int64 | id of publisher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublishersAPI.DeletePublisher(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.DeletePublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePublisher`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.DeletePublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of publisher | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponseV1**](ApiResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePublicationSeries

> PublishersPublicationResponseV1 RetrievePublicationSeries(ctx).Pubname(pubname).Execute()

get the list of series for a specific publication

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
    pubname := "pubname_example" // string | Publication name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublishersAPI.RetrievePublicationSeries(context.Background()).Pubname(pubname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.RetrievePublicationSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePublicationSeries`: PublishersPublicationResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.RetrievePublicationSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePublicationSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubname** | **string** | Publication name | 

### Return type

[**PublishersPublicationResponseV1**](PublishersPublicationResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePublisher

> PublishersModelV1 RetrievePublisher(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific publisher

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
    id := int64(56) // int64 | Publisher id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublishersAPI.RetrievePublisher(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.RetrievePublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePublisher`: PublishersModelV1
    fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.RetrievePublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Publisher id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**PublishersModelV1**](PublishersModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePublisherSeries

> PublishersSeriesListResponseV1 RetrievePublisherSeries(ctx, id).Execute()

get the list of series for a specific publisher

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
    id := int64(789) // int64 | Publisher id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublishersAPI.RetrievePublisherSeries(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.RetrievePublisherSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePublisherSeries`: PublishersSeriesListResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.RetrievePublisherSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Publisher id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePublisherSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublishersSeriesListResponseV1**](PublishersSeriesListResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPublishersPost

> PublishersSearchResponseV1 SearchPublishersPost(ctx).PublishersSearchRequestV1(publishersSearchRequestV1).Execute()

search publishers

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
    publishersSearchRequestV1 := *openapiclient.NewPublishersSearchRequestV1() // PublishersSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublishersAPI.SearchPublishersPost(context.Background()).PublishersSearchRequestV1(publishersSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.SearchPublishersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPublishersPost`: PublishersSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.SearchPublishersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPublishersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishersSearchRequestV1** | [**PublishersSearchRequestV1**](PublishersSearchRequestV1.md) |  | 

### Return type

[**PublishersSearchResponseV1**](PublishersSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublisher

> ApiResponseV1 UpdatePublisher(ctx, id).PublishersModelUpdateV1(publishersModelUpdateV1).Execute()

update a publisher

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
    id := int64(56) // int64 | id of publisher
    publishersModelUpdateV1 := *openapiclient.NewPublishersModelUpdateV1() // PublishersModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublishersAPI.UpdatePublisher(context.Background(), id).PublishersModelUpdateV1(publishersModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.UpdatePublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePublisher`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.UpdatePublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of publisher | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publishersModelUpdateV1** | [**PublishersModelUpdateV1**](PublishersModelUpdateV1.md) |  | 

### Return type

[**ApiResponseV1**](ApiResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

