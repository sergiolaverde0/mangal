# \GroupsAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroup**](GroupsAPI.md#AddGroup) | **Post** /groups | add an group
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{id} | delete a group
[**RejectGroup**](GroupsAPI.md#RejectGroup) | **Post** /groups/{id}/reject | reject and delete a group
[**RetrieveGroup**](GroupsAPI.md#RetrieveGroup) | **Get** /groups/{id} | get a specific group
[**RetrieveGroupSeries**](GroupsAPI.md#RetrieveGroupSeries) | **Get** /groups/{id}/series | get the list of series and release frequency for a specific group
[**SearchGroupsPost**](GroupsAPI.md#SearchGroupsPost) | **Post** /groups/search | search groups
[**UpdateGroup**](GroupsAPI.md#UpdateGroup) | **Patch** /groups/{id} | update an group



## AddGroup

> ApiResponseV1 AddGroup(ctx).GroupsModelUpdateV1(groupsModelUpdateV1).Execute()

add an group

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
    groupsModelUpdateV1 := *openapiclient.NewGroupsModelUpdateV1() // GroupsModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsAPI.AddGroup(context.Background()).GroupsModelUpdateV1(groupsModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGroup`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupsModelUpdateV1** | [**GroupsModelUpdateV1**](GroupsModelUpdateV1.md) |  | 

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


## DeleteGroup

> ApiResponseV1 DeleteGroup(ctx, id).Execute()

delete a group

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
    id := int64(56) // int64 | id of group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroup`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## RejectGroup

> ApiResponseV1 RejectGroup(ctx, id).Execute()

reject and delete a group

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
    id := int64(56) // int64 | id of group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsAPI.RejectGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RejectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectGroup`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.RejectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectGroupRequest struct via the builder pattern


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


## RetrieveGroup

> GroupsModelV1 RetrieveGroup(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific group

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
    id := int64(789) // int64 | Group id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsAPI.RetrieveGroup(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RetrieveGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveGroup`: GroupsModelV1
    fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.RetrieveGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**GroupsModelV1**](GroupsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveGroupSeries

> GroupsSeriesListResponseV1 RetrieveGroupSeries(ctx, id).Execute()

get the list of series and release frequency for a specific group

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
    id := int64(789) // int64 | Group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsAPI.RetrieveGroupSeries(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RetrieveGroupSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveGroupSeries`: GroupsSeriesListResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.RetrieveGroupSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveGroupSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupsSeriesListResponseV1**](GroupsSeriesListResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGroupsPost

> GroupsSearchResponseV1 SearchGroupsPost(ctx).GroupsSearchRequestV1(groupsSearchRequestV1).Execute()

search groups

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
    groupsSearchRequestV1 := *openapiclient.NewGroupsSearchRequestV1() // GroupsSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsAPI.SearchGroupsPost(context.Background()).GroupsSearchRequestV1(groupsSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SearchGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchGroupsPost`: GroupsSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SearchGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupsSearchRequestV1** | [**GroupsSearchRequestV1**](GroupsSearchRequestV1.md) |  | 

### Return type

[**GroupsSearchResponseV1**](GroupsSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> ApiResponseV1 UpdateGroup(ctx, id).GroupsModelUpdateV1(groupsModelUpdateV1).Execute()

update an group

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
    id := int64(56) // int64 | id of group
    groupsModelUpdateV1 := *openapiclient.NewGroupsModelUpdateV1() // GroupsModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsAPI.UpdateGroup(context.Background(), id).GroupsModelUpdateV1(groupsModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsModelUpdateV1** | [**GroupsModelUpdateV1**](GroupsModelUpdateV1.md) |  | 

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

