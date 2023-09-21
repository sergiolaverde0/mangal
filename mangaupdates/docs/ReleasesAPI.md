# \ReleasesAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRelease**](ReleasesAPI.md#AddRelease) | **Post** /releases | add an release
[**DeleteRelease**](ReleasesAPI.md#DeleteRelease) | **Delete** /releases/{id} | delete a release
[**ListReleasesByDay**](ReleasesAPI.md#ListReleasesByDay) | **Get** /releases/days | show releases by day
[**ModerateReleasesPost**](ReleasesAPI.md#ModerateReleasesPost) | **Post** /releases/moderate | search releases to moderate
[**ReleaseRssFeed**](ReleasesAPI.md#ReleaseRssFeed) | **Get** /releases/rss | releases rss feed
[**RetrieveRelease**](ReleasesAPI.md#RetrieveRelease) | **Get** /releases/{id} | get a specific release
[**SearchReleasesPost**](ReleasesAPI.md#SearchReleasesPost) | **Post** /releases/search | search releases
[**UpdateRelease**](ReleasesAPI.md#UpdateRelease) | **Patch** /releases/{id} | update an release



## AddRelease

> ApiResponseV1 AddRelease(ctx).ReleaseModelUpdateV1(releaseModelUpdateV1).Execute()

add an release

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
    releaseModelUpdateV1 := []openapiclient.ReleaseModelUpdateV1{*openapiclient.NewReleaseModelUpdateV1()} // []ReleaseModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.AddRelease(context.Background()).ReleaseModelUpdateV1(releaseModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.AddRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRelease`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.AddRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseModelUpdateV1** | [**[]ReleaseModelUpdateV1**](ReleaseModelUpdateV1.md) |  | 

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


## DeleteRelease

> ApiResponseV1 DeleteRelease(ctx, id).Execute()

delete a release

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
    id := int64(56) // int64 | id of release

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.DeleteRelease(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.DeleteRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRelease`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.DeleteRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of release | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseRequest struct via the builder pattern


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


## ListReleasesByDay

> ReleaseSearchResponseV1 ListReleasesByDay(ctx).Page(page).IncludeMetadata(includeMetadata).Execute()

show releases by day

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
    page := int64(789) // int64 | Start page (optional)
    includeMetadata := true // bool | Include series metadata (if available) (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.ListReleasesByDay(context.Background()).Page(page).IncludeMetadata(includeMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.ListReleasesByDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReleasesByDay`: ReleaseSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.ListReleasesByDay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReleasesByDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Start page | 
 **includeMetadata** | **bool** | Include series metadata (if available) | [default to false]

### Return type

[**ReleaseSearchResponseV1**](ReleaseSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModerateReleasesPost

> ReleaseModerateResponseV1 ModerateReleasesPost(ctx).ReleaseModerateRequestV1(releaseModerateRequestV1).Execute()

search releases to moderate

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
    releaseModerateRequestV1 := *openapiclient.NewReleaseModerateRequestV1() // ReleaseModerateRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.ModerateReleasesPost(context.Background()).ReleaseModerateRequestV1(releaseModerateRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.ModerateReleasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModerateReleasesPost`: ReleaseModerateResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.ModerateReleasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModerateReleasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseModerateRequestV1** | [**ReleaseModerateRequestV1**](ReleaseModerateRequestV1.md) |  | 

### Return type

[**ReleaseModerateResponseV1**](ReleaseModerateResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseRssFeed

> string ReleaseRssFeed(ctx).Execute()

releases rss feed

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
    resp, r, err := apiClient.ReleasesAPI.ReleaseRssFeed(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.ReleaseRssFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseRssFeed`: string
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.ReleaseRssFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseRssFeedRequest struct via the builder pattern


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRelease

> ReleaseModelV1 RetrieveRelease(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific release

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
    id := int64(56) // int64 | Release id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.RetrieveRelease(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.RetrieveRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRelease`: ReleaseModelV1
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.RetrieveRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Release id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ReleaseModelV1**](ReleaseModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReleasesPost

> ReleaseSearchResponseV1 SearchReleasesPost(ctx).ReleaseSearchRequestV1(releaseSearchRequestV1).Execute()

search releases

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
    releaseSearchRequestV1 := *openapiclient.NewReleaseSearchRequestV1() // ReleaseSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.SearchReleasesPost(context.Background()).ReleaseSearchRequestV1(releaseSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.SearchReleasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchReleasesPost`: ReleaseSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.SearchReleasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchReleasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseSearchRequestV1** | [**ReleaseSearchRequestV1**](ReleaseSearchRequestV1.md) |  | 

### Return type

[**ReleaseSearchResponseV1**](ReleaseSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRelease

> ApiResponseV1 UpdateRelease(ctx, id).ReleaseModelUpdateV1(releaseModelUpdateV1).Execute()

update an release

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
    id := int64(56) // int64 | id of release
    releaseModelUpdateV1 := *openapiclient.NewReleaseModelUpdateV1() // ReleaseModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.UpdateRelease(context.Background(), id).ReleaseModelUpdateV1(releaseModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.UpdateRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRelease`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.UpdateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of release | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseModelUpdateV1** | [**ReleaseModelUpdateV1**](ReleaseModelUpdateV1.md) |  | 

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

