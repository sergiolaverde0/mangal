# \ListsAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomList**](ListsAPI.md#AddCustomList) | **Post** /lists | add a custom user list
[**AddListSeries**](ListsAPI.md#AddListSeries) | **Post** /lists/series | add a series to a list
[**AddListSeriesBulk**](ListsAPI.md#AddListSeriesBulk) | **Post** /lists/{id}/series/bulk | add a list of series to a list
[**DeleteCustomList**](ListsAPI.md#DeleteCustomList) | **Delete** /lists/{id} | remove a custom list
[**DeleteListSeries**](ListsAPI.md#DeleteListSeries) | **Post** /lists/series/delete | remove a series from a list
[**RetrieveListById**](ListsAPI.md#RetrieveListById) | **Get** /lists/{id} | retrieve list metadata and options
[**RetrieveListSeries**](ListsAPI.md#RetrieveListSeries) | **Get** /lists/series/{series_id} | retrieve list series item
[**RetrieveLists**](ListsAPI.md#RetrieveLists) | **Get** /lists | retrieve list of user lists
[**RetrievePublicListStats**](ListsAPI.md#RetrievePublicListStats) | **Get** /lists/public/{user_id}/stats | retrieve stats for user public lists
[**RetrievePublicLists**](ListsAPI.md#RetrievePublicLists) | **Get** /lists/public/{user_id} | retrieve list of user lists
[**RetrieveSimilarUsersBySeries**](ListsAPI.md#RetrieveSimilarUsersBySeries) | **Get** /lists/similar/{list_name}/{series_id} | retrieve users who have similar interests based on series
[**SearchListsPost**](ListsAPI.md#SearchListsPost) | **Post** /lists/{id}/search | search lists
[**SearchPublicListsPost**](ListsAPI.md#SearchPublicListsPost) | **Post** /lists/public/{user_id}/search/{id} | search lists
[**UpdateList**](ListsAPI.md#UpdateList) | **Patch** /lists/{id} | update a user list
[**UpdateListSeries**](ListsAPI.md#UpdateListSeries) | **Post** /lists/series/update | update a series list item



## AddCustomList

> ApiResponseV1 AddCustomList(ctx).ListsModelUpdateV1(listsModelUpdateV1).Execute()

add a custom user list

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
    listsModelUpdateV1 := *openapiclient.NewListsModelUpdateV1() // ListsModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.AddCustomList(context.Background()).ListsModelUpdateV1(listsModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.AddCustomList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomList`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.AddCustomList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listsModelUpdateV1** | [**ListsModelUpdateV1**](ListsModelUpdateV1.md) |  | 

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


## AddListSeries

> AddListSeries(ctx).ListsSeriesModelUpdateV1(listsSeriesModelUpdateV1).Execute()

add a series to a list

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
    listsSeriesModelUpdateV1 := []openapiclient.ListsSeriesModelUpdateV1{*openapiclient.NewListsSeriesModelUpdateV1(*openapiclient.NewListsSeriesModelUpdateV1Series(int64(123)))} // []ListsSeriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ListsAPI.AddListSeries(context.Background()).ListsSeriesModelUpdateV1(listsSeriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.AddListSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddListSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listsSeriesModelUpdateV1** | [**[]ListsSeriesModelUpdateV1**](ListsSeriesModelUpdateV1.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddListSeriesBulk

> ApiResponseV1 AddListSeriesBulk(ctx, id).ListsBulkAddModelV1(listsBulkAddModelV1).Execute()

add a list of series to a list

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
    id := int64(56) // int64 | id of list
    listsBulkAddModelV1 := []openapiclient.ListsBulkAddModelV1{*openapiclient.NewListsBulkAddModelV1()} // []ListsBulkAddModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.AddListSeriesBulk(context.Background(), id).ListsBulkAddModelV1(listsBulkAddModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.AddListSeriesBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddListSeriesBulk`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.AddListSeriesBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddListSeriesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listsBulkAddModelV1** | [**[]ListsBulkAddModelV1**](ListsBulkAddModelV1.md) |  | 

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


## DeleteCustomList

> ApiResponseV1 DeleteCustomList(ctx, id).Execute()

remove a custom list

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
    id := int64(56) // int64 | id of list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.DeleteCustomList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteCustomList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomList`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.DeleteCustomList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomListRequest struct via the builder pattern


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


## DeleteListSeries

> ApiResponseV1 DeleteListSeries(ctx).RequestBody(requestBody).Execute()

remove a series from a list

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
    requestBody := []int64{int64(123)} // []int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.DeleteListSeries(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteListSeries`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.DeleteListSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int64** |  | 

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


## RetrieveListById

> ListsModelV1 RetrieveListById(ctx, id).UnrenderedFields(unrenderedFields).Execute()

retrieve list metadata and options

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
    id := int64(56) // int64 | List id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.RetrieveListById(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.RetrieveListById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListById`: ListsModelV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.RetrieveListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | List id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ListsModelV1**](ListsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListSeries

> ListsSeriesModelV1 RetrieveListSeries(ctx, seriesId).UnrenderedFields(unrenderedFields).Execute()

retrieve list series item

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
    seriesId := int64(56) // int64 | Series id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.RetrieveListSeries(context.Background(), seriesId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.RetrieveListSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListSeries`: ListsSeriesModelV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.RetrieveListSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesId** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ListsSeriesModelV1**](ListsSeriesModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveLists

> []ListsModelV1 RetrieveLists(ctx).Execute()

retrieve list of user lists

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
    resp, r, err := apiClient.ListsAPI.RetrieveLists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.RetrieveLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveLists`: []ListsModelV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.RetrieveLists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListsRequest struct via the builder pattern


### Return type

[**[]ListsModelV1**](ListsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePublicListStats

> ListsPublicStatsModelV1 RetrievePublicListStats(ctx, userId).Execute()

retrieve stats for user public lists

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
    userId := int64(56) // int64 | User id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.RetrievePublicListStats(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.RetrievePublicListStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePublicListStats`: ListsPublicStatsModelV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.RetrievePublicListStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePublicListStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListsPublicStatsModelV1**](ListsPublicStatsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePublicLists

> []ListsModelV1 RetrievePublicLists(ctx, userId).Execute()

retrieve list of user lists

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
    userId := int64(56) // int64 | User id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.RetrievePublicLists(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.RetrievePublicLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePublicLists`: []ListsModelV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.RetrievePublicLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePublicListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListsModelV1**](ListsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSimilarUsersBySeries

> ListsSimilarUsersResponseV1 RetrieveSimilarUsersBySeries(ctx, listName, seriesId).Execute()

retrieve users who have similar interests based on series

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
    listName := "listName_example" // string | name of list
    seriesId := int64(56) // int64 | Series id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.RetrieveSimilarUsersBySeries(context.Background(), listName, seriesId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.RetrieveSimilarUsersBySeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSimilarUsersBySeries`: ListsSimilarUsersResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.RetrieveSimilarUsersBySeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listName** | **string** | name of list | 
**seriesId** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSimilarUsersBySeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListsSimilarUsersResponseV1**](ListsSimilarUsersResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListsPost

> ListsSearchResponseV1 SearchListsPost(ctx, id).ListsSearchRequestV1(listsSearchRequestV1).Execute()

search lists

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
    id := int64(56) // int64 | list id to search
    listsSearchRequestV1 := *openapiclient.NewListsSearchRequestV1() // ListsSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.SearchListsPost(context.Background(), id).ListsSearchRequestV1(listsSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.SearchListsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchListsPost`: ListsSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.SearchListsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | list id to search | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listsSearchRequestV1** | [**ListsSearchRequestV1**](ListsSearchRequestV1.md) |  | 

### Return type

[**ListsSearchResponseV1**](ListsSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPublicListsPost

> ListsPublicSearchResponseV1 SearchPublicListsPost(ctx, userId, id).ListsSearchRequestV1(listsSearchRequestV1).Execute()

search lists

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
    userId := int64(56) // int64 | User id
    id := int64(56) // int64 | list id to search
    listsSearchRequestV1 := *openapiclient.NewListsSearchRequestV1() // ListsSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.SearchPublicListsPost(context.Background(), userId, id).ListsSearchRequestV1(listsSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.SearchPublicListsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPublicListsPost`: ListsPublicSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.SearchPublicListsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 
**id** | **int64** | list id to search | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchPublicListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **listsSearchRequestV1** | [**ListsSearchRequestV1**](ListsSearchRequestV1.md) |  | 

### Return type

[**ListsPublicSearchResponseV1**](ListsPublicSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateList

> ApiResponseV1 UpdateList(ctx, id).ListsModelUpdateV1(listsModelUpdateV1).Execute()

update a user list

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
    id := int64(56) // int64 | id of list
    listsModelUpdateV1 := *openapiclient.NewListsModelUpdateV1() // ListsModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.UpdateList(context.Background(), id).ListsModelUpdateV1(listsModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.UpdateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateList`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.UpdateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listsModelUpdateV1** | [**ListsModelUpdateV1**](ListsModelUpdateV1.md) |  | 

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


## UpdateListSeries

> ApiResponseV1 UpdateListSeries(ctx).ListsSeriesModelUpdateV1(listsSeriesModelUpdateV1).Execute()

update a series list item

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
    listsSeriesModelUpdateV1 := []openapiclient.ListsSeriesModelUpdateV1{*openapiclient.NewListsSeriesModelUpdateV1(*openapiclient.NewListsSeriesModelUpdateV1Series(int64(123)))} // []ListsSeriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListsAPI.UpdateListSeries(context.Background()).ListsSeriesModelUpdateV1(listsSeriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.UpdateListSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListSeries`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ListsAPI.UpdateListSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listsSeriesModelUpdateV1** | [**[]ListsSeriesModelUpdateV1**](ListsSeriesModelUpdateV1.md) |  | 

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

