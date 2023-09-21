# \SeriesAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSeries**](SeriesAPI.md#AddSeries) | **Post** /series | add an series
[**AddSeriesCategoryVote**](SeriesAPI.md#AddSeriesCategoryVote) | **Post** /series/{id}/categories/vote | add a vote for a category on a series
[**AddSeriesComment**](SeriesAPI.md#AddSeriesComment) | **Post** /series/{id}/comments | add a series comment
[**AddSeriesCommentUsefulFlag**](SeriesAPI.md#AddSeriesCommentUsefulFlag) | **Put** /series/{id}/comments/{comment_id}/useful | set usefulness of a series comment
[**CombineSeriesCategories**](SeriesAPI.md#CombineSeriesCategories) | **Post** /series/{id}/categories/combine | combine two series categories
[**DeleteSeries**](SeriesAPI.md#DeleteSeries) | **Delete** /series/{id} | delete a series
[**DeleteSeriesCategory**](SeriesAPI.md#DeleteSeriesCategory) | **Post** /series/{id}/categories/delete | deletes a series category
[**DeleteSeriesComment**](SeriesAPI.md#DeleteSeriesComment) | **Delete** /series/{id}/comments/{comment_id} | delete a series comment
[**DeleteSeriesImage**](SeriesAPI.md#DeleteSeriesImage) | **Delete** /series/{id}/image | delete the image of an series
[**DeleteUserSeriesRating**](SeriesAPI.md#DeleteUserSeriesRating) | **Delete** /series/{id}/rating | delete a series rating for a user
[**LockSeriesField**](SeriesAPI.md#LockSeriesField) | **Post** /series/{id}/locks/{item}/lock | lock a field of an series
[**RemoveSeriesCategoryVote**](SeriesAPI.md#RemoveSeriesCategoryVote) | **Post** /series/{id}/categories/vote/delete | remove series category vote for user
[**RemoveSeriesCommentUsefulFlag**](SeriesAPI.md#RemoveSeriesCommentUsefulFlag) | **Delete** /series/{id}/comments/{comment_id}/useful | remove usefulness of a series comment
[**RenameSeriesCategory**](SeriesAPI.md#RenameSeriesCategory) | **Post** /series/{id}/categories/rename | renames a series category
[**ReportSeriesComment**](SeriesAPI.md#ReportSeriesComment) | **Post** /series/{id}/comments/{comment_id}/report | report a series comment
[**RetrieveMySeriesComment**](SeriesAPI.md#RetrieveMySeriesComment) | **Get** /series/{id}/comments/my_comment | get my series comment
[**RetrieveSeries**](SeriesAPI.md#RetrieveSeries) | **Get** /series/{id} | get a specific series
[**RetrieveSeriesCategoryVotes**](SeriesAPI.md#RetrieveSeriesCategoryVotes) | **Get** /series/{id}/categories/votes | get category votes for the current user
[**RetrieveSeriesComment**](SeriesAPI.md#RetrieveSeriesComment) | **Get** /series/{id}/comments/{comment_id} | get a specific series comment
[**RetrieveSeriesCommentLocation**](SeriesAPI.md#RetrieveSeriesCommentLocation) | **Get** /series/{id}/comments/{comment_id}/location | get a specific series comment location
[**RetrieveSeriesGroups**](SeriesAPI.md#RetrieveSeriesGroups) | **Get** /series/{id}/groups | get the list of groups scanlating a specific series
[**RetrieveSeriesLocks**](SeriesAPI.md#RetrieveSeriesLocks) | **Get** /series/{id}/locks | get a specific series lock
[**RetrieveSeriesRankLocation**](SeriesAPI.md#RetrieveSeriesRankLocation) | **Get** /series/{id}/rank/{type} | get a specific series rank location
[**RetrieveSeriesRatingRainbow**](SeriesAPI.md#RetrieveSeriesRatingRainbow) | **Get** /series/{id}/ratingrainbow | get a the series rating rainbow
[**RetrieveUserSeriesRating**](SeriesAPI.md#RetrieveUserSeriesRating) | **Get** /series/{id}/rating | get a specific series rating for a user
[**SearchSeriesCommentsPost**](SeriesAPI.md#SearchSeriesCommentsPost) | **Post** /series/{id}/comments/search | search series comments
[**SearchSeriesHistoryPost**](SeriesAPI.md#SearchSeriesHistoryPost) | **Post** /series/{id}/history | search series history
[**SearchSeriesPost**](SeriesAPI.md#SearchSeriesPost) | **Post** /series/search | search series
[**SeriesCommentsModerationPost**](SeriesAPI.md#SeriesCommentsModerationPost) | **Post** /series/comments/moderation | moderate series comments
[**SeriesReleaseRssFeed**](SeriesAPI.md#SeriesReleaseRssFeed) | **Get** /series/{id}/rss | releases rss feed for a specific series
[**UnlockSeriesField**](SeriesAPI.md#UnlockSeriesField) | **Post** /series/{id}/locks/{item}/unlock | unlock a field of an series
[**UpdateSeries**](SeriesAPI.md#UpdateSeries) | **Patch** /series/{id} | update an series
[**UpdateSeriesComment**](SeriesAPI.md#UpdateSeriesComment) | **Patch** /series/{id}/comments/{comment_id} | update a series comment
[**UpdateSeriesImage**](SeriesAPI.md#UpdateSeriesImage) | **Post** /series/{id}/image | update the image of an series
[**UpdateUserSeriesRating**](SeriesAPI.md#UpdateUserSeriesRating) | **Put** /series/{id}/rating | update the user rating for a series



## AddSeries

> ApiResponseV1 AddSeries(ctx).SeriesModelUpdateV1(seriesModelUpdateV1).Execute()

add an series

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
    seriesModelUpdateV1 := *openapiclient.NewSeriesModelUpdateV1() // SeriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.AddSeries(context.Background()).SeriesModelUpdateV1(seriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.AddSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSeries`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.AddSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesModelUpdateV1** | [**SeriesModelUpdateV1**](SeriesModelUpdateV1.md) |  | 

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


## AddSeriesCategoryVote

> ApiResponseV1 AddSeriesCategoryVote(ctx, id).SeriesCategoryVoteModelV1(seriesCategoryVoteModelV1).Execute()

add a vote for a category on a series

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
    id := int64(56) // int64 | id of series
    seriesCategoryVoteModelV1 := *openapiclient.NewSeriesCategoryVoteModelV1() // SeriesCategoryVoteModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.AddSeriesCategoryVote(context.Background(), id).SeriesCategoryVoteModelV1(seriesCategoryVoteModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.AddSeriesCategoryVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSeriesCategoryVote`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.AddSeriesCategoryVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSeriesCategoryVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesCategoryVoteModelV1** | [**SeriesCategoryVoteModelV1**](SeriesCategoryVoteModelV1.md) |  | 

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


## AddSeriesComment

> ApiResponseV1 AddSeriesComment(ctx, id).SeriesCommentModelUpdateV1(seriesCommentModelUpdateV1).Execute()

add a series comment

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
    id := int64(56) // int64 | id of series
    seriesCommentModelUpdateV1 := *openapiclient.NewSeriesCommentModelUpdateV1() // SeriesCommentModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.AddSeriesComment(context.Background(), id).SeriesCommentModelUpdateV1(seriesCommentModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.AddSeriesComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSeriesComment`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.AddSeriesComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSeriesCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesCommentModelUpdateV1** | [**SeriesCommentModelUpdateV1**](SeriesCommentModelUpdateV1.md) |  | 

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


## AddSeriesCommentUsefulFlag

> ApiResponseV1 AddSeriesCommentUsefulFlag(ctx, id, commentId).SeriesCommentUsefulModelV1(seriesCommentUsefulModelV1).Execute()

set usefulness of a series comment

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
    id := int64(56) // int64 | id of series
    commentId := int64(56) // int64 | id of series comment
    seriesCommentUsefulModelV1 := *openapiclient.NewSeriesCommentUsefulModelV1() // SeriesCommentUsefulModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.AddSeriesCommentUsefulFlag(context.Background(), id, commentId).SeriesCommentUsefulModelV1(seriesCommentUsefulModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.AddSeriesCommentUsefulFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSeriesCommentUsefulFlag`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.AddSeriesCommentUsefulFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 
**commentId** | **int64** | id of series comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSeriesCommentUsefulFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seriesCommentUsefulModelV1** | [**SeriesCommentUsefulModelV1**](SeriesCommentUsefulModelV1.md) |  | 

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


## CombineSeriesCategories

> ApiResponseV1 CombineSeriesCategories(ctx, id).SeriesCategoryUpdateModelV1(seriesCategoryUpdateModelV1).Execute()

combine two series categories

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
    id := int64(56) // int64 | id of series
    seriesCategoryUpdateModelV1 := *openapiclient.NewSeriesCategoryUpdateModelV1() // SeriesCategoryUpdateModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.CombineSeriesCategories(context.Background(), id).SeriesCategoryUpdateModelV1(seriesCategoryUpdateModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.CombineSeriesCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CombineSeriesCategories`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.CombineSeriesCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiCombineSeriesCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesCategoryUpdateModelV1** | [**SeriesCategoryUpdateModelV1**](SeriesCategoryUpdateModelV1.md) |  | 

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


## DeleteSeries

> ApiResponseV1 DeleteSeries(ctx, id).Execute()

delete a series

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
    id := int64(56) // int64 | id of series

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.DeleteSeries(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.DeleteSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSeries`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.DeleteSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSeriesRequest struct via the builder pattern


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


## DeleteSeriesCategory

> ApiResponseV1 DeleteSeriesCategory(ctx, id).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()

deletes a series category

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
    id := int64(56) // int64 | id of series
    categoriesModelUpdateV1 := *openapiclient.NewCategoriesModelUpdateV1() // CategoriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.DeleteSeriesCategory(context.Background(), id).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.DeleteSeriesCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSeriesCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.DeleteSeriesCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSeriesCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoriesModelUpdateV1** | [**CategoriesModelUpdateV1**](CategoriesModelUpdateV1.md) |  | 

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


## DeleteSeriesComment

> ApiResponseV1 DeleteSeriesComment(ctx, id, commentId).Execute()

delete a series comment

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
    id := int64(56) // int64 | id of series
    commentId := int64(56) // int64 | id of series comment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.DeleteSeriesComment(context.Background(), id, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.DeleteSeriesComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSeriesComment`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.DeleteSeriesComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 
**commentId** | **int64** | id of series comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSeriesCommentRequest struct via the builder pattern


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


## DeleteSeriesImage

> ApiResponseV1 DeleteSeriesImage(ctx, id).Execute()

delete the image of an series

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
    id := int64(789) // int64 | id of series

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.DeleteSeriesImage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.DeleteSeriesImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSeriesImage`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.DeleteSeriesImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSeriesImageRequest struct via the builder pattern


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


## DeleteUserSeriesRating

> ApiResponseV1 DeleteUserSeriesRating(ctx, id).Execute()

delete a series rating for a user

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
    id := int64(56) // int64 | id of series

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.DeleteUserSeriesRating(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.DeleteUserSeriesRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserSeriesRating`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.DeleteUserSeriesRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSeriesRatingRequest struct via the builder pattern


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


## LockSeriesField

> ApiResponseV1 LockSeriesField(ctx, id, item).SeriesLockModelUpdateV1(seriesLockModelUpdateV1).Execute()

lock a field of an series

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
    id := int64(789) // int64 | id of series
    item := "item_example" // string | field name
    seriesLockModelUpdateV1 := *openapiclient.NewSeriesLockModelUpdateV1() // SeriesLockModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.LockSeriesField(context.Background(), id, item).SeriesLockModelUpdateV1(seriesLockModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.LockSeriesField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockSeriesField`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.LockSeriesField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 
**item** | **string** | field name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockSeriesFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seriesLockModelUpdateV1** | [**SeriesLockModelUpdateV1**](SeriesLockModelUpdateV1.md) |  | 

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


## RemoveSeriesCategoryVote

> ApiResponseV1 RemoveSeriesCategoryVote(ctx, id).SeriesCategoryVoteDeleteModelV1(seriesCategoryVoteDeleteModelV1).Execute()

remove series category vote for user

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
    id := int64(56) // int64 | id of series
    seriesCategoryVoteDeleteModelV1 := *openapiclient.NewSeriesCategoryVoteDeleteModelV1() // SeriesCategoryVoteDeleteModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RemoveSeriesCategoryVote(context.Background(), id).SeriesCategoryVoteDeleteModelV1(seriesCategoryVoteDeleteModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RemoveSeriesCategoryVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSeriesCategoryVote`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RemoveSeriesCategoryVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSeriesCategoryVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesCategoryVoteDeleteModelV1** | [**SeriesCategoryVoteDeleteModelV1**](SeriesCategoryVoteDeleteModelV1.md) |  | 

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


## RemoveSeriesCommentUsefulFlag

> ApiResponseV1 RemoveSeriesCommentUsefulFlag(ctx, id, commentId).Execute()

remove usefulness of a series comment

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
    id := int64(56) // int64 | id of series
    commentId := int64(56) // int64 | id of series comment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RemoveSeriesCommentUsefulFlag(context.Background(), id, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RemoveSeriesCommentUsefulFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSeriesCommentUsefulFlag`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RemoveSeriesCommentUsefulFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 
**commentId** | **int64** | id of series comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSeriesCommentUsefulFlagRequest struct via the builder pattern


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


## RenameSeriesCategory

> ApiResponseV1 RenameSeriesCategory(ctx, id).SeriesCategoryUpdateModelV1(seriesCategoryUpdateModelV1).Execute()

renames a series category

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
    id := int64(56) // int64 | id of series
    seriesCategoryUpdateModelV1 := *openapiclient.NewSeriesCategoryUpdateModelV1() // SeriesCategoryUpdateModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RenameSeriesCategory(context.Background(), id).SeriesCategoryUpdateModelV1(seriesCategoryUpdateModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RenameSeriesCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameSeriesCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RenameSeriesCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameSeriesCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesCategoryUpdateModelV1** | [**SeriesCategoryUpdateModelV1**](SeriesCategoryUpdateModelV1.md) |  | 

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


## ReportSeriesComment

> ApiResponseV1 ReportSeriesComment(ctx, id, commentId).SeriesCommentReportModelV1(seriesCommentReportModelV1).Execute()

report a series comment

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
    id := int64(56) // int64 | id of series
    commentId := int64(56) // int64 | id of series comment
    seriesCommentReportModelV1 := *openapiclient.NewSeriesCommentReportModelV1() // SeriesCommentReportModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.ReportSeriesComment(context.Background(), id, commentId).SeriesCommentReportModelV1(seriesCommentReportModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.ReportSeriesComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportSeriesComment`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.ReportSeriesComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 
**commentId** | **int64** | id of series comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportSeriesCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seriesCommentReportModelV1** | [**SeriesCommentReportModelV1**](SeriesCommentReportModelV1.md) |  | 

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


## RetrieveMySeriesComment

> SeriesCommentModelV1 RetrieveMySeriesComment(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get my series comment

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
    id := int64(789) // int64 | Series id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveMySeriesComment(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveMySeriesComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMySeriesComment`: SeriesCommentModelV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveMySeriesComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMySeriesCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**SeriesCommentModelV1**](SeriesCommentModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSeries

> SeriesModelV1 RetrieveSeries(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific series

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
    id := int64(56) // int64 | Series id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeries(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeries`: SeriesModelV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**SeriesModelV1**](SeriesModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSeriesCategoryVotes

> []SeriesCategoryVoteModelV1 RetrieveSeriesCategoryVotes(ctx, id).Execute()

get category votes for the current user

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
    id := int64(789) // int64 | Series id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeriesCategoryVotes(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeriesCategoryVotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeriesCategoryVotes`: []SeriesCategoryVoteModelV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeriesCategoryVotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesCategoryVotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SeriesCategoryVoteModelV1**](SeriesCategoryVoteModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSeriesComment

> SeriesCommentModelV1 RetrieveSeriesComment(ctx, id, commentId).UnrenderedFields(unrenderedFields).Execute()

get a specific series comment

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
    id := int64(789) // int64 | Series id
    commentId := int64(789) // int64 | Series comment id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeriesComment(context.Background(), id, commentId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeriesComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeriesComment`: SeriesCommentModelV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeriesComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 
**commentId** | **int64** | Series comment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**SeriesCommentModelV1**](SeriesCommentModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSeriesCommentLocation

> ApiResponseV1 RetrieveSeriesCommentLocation(ctx, id, commentId).Execute()

get a specific series comment location

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
    id := int64(789) // int64 | Series id
    commentId := int64(789) // int64 | Series comment id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeriesCommentLocation(context.Background(), id, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeriesCommentLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeriesCommentLocation`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeriesCommentLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 
**commentId** | **int64** | Series comment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesCommentLocationRequest struct via the builder pattern


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


## RetrieveSeriesGroups

> SeriesGroupListResponseV1 RetrieveSeriesGroups(ctx, id).Execute()

get the list of groups scanlating a specific series

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
    id := int64(789) // int64 | Series id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeriesGroups(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeriesGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeriesGroups`: SeriesGroupListResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeriesGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SeriesGroupListResponseV1**](SeriesGroupListResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSeriesLocks

> []SeriesLockModelV1 RetrieveSeriesLocks(ctx, id).Execute()

get a specific series lock

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
    id := int64(789) // int64 | Series id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeriesLocks(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeriesLocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeriesLocks`: []SeriesLockModelV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeriesLocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SeriesLockModelV1**](SeriesLockModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSeriesRankLocation

> ApiResponseV1 RetrieveSeriesRankLocation(ctx, id, type_).Execute()

get a specific series rank location

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
    id := int64(789) // int64 | Series id
    type_ := "type__example" // string | Stat type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeriesRankLocation(context.Background(), id, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeriesRankLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeriesRankLocation`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeriesRankLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 
**type_** | **string** | Stat type | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesRankLocationRequest struct via the builder pattern


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


## RetrieveSeriesRatingRainbow

> SeriesRatingRainbowModelV1 RetrieveSeriesRatingRainbow(ctx, id).Execute()

get a the series rating rainbow

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
    id := int64(789) // int64 | Series id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveSeriesRatingRainbow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveSeriesRatingRainbow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSeriesRatingRainbow`: SeriesRatingRainbowModelV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveSeriesRatingRainbow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSeriesRatingRainbowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SeriesRatingRainbowModelV1**](SeriesRatingRainbowModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUserSeriesRating

> SeriesRatingModelV1 RetrieveUserSeriesRating(ctx, id).Execute()

get a specific series rating for a user

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
    id := int64(789) // int64 | Series id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.RetrieveUserSeriesRating(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.RetrieveUserSeriesRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUserSeriesRating`: SeriesRatingModelV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.RetrieveUserSeriesRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserSeriesRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SeriesRatingModelV1**](SeriesRatingModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSeriesCommentsPost

> SeriesCommentSearchResponseV1 SearchSeriesCommentsPost(ctx, id).SeriesCommentSearchRequestV1(seriesCommentSearchRequestV1).Execute()

search series comments

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
    id := int64(789) // int64 | Series id
    seriesCommentSearchRequestV1 := *openapiclient.NewSeriesCommentSearchRequestV1() // SeriesCommentSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.SearchSeriesCommentsPost(context.Background(), id).SeriesCommentSearchRequestV1(seriesCommentSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.SearchSeriesCommentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSeriesCommentsPost`: SeriesCommentSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.SearchSeriesCommentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSeriesCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesCommentSearchRequestV1** | [**SeriesCommentSearchRequestV1**](SeriesCommentSearchRequestV1.md) |  | 

### Return type

[**SeriesCommentSearchResponseV1**](SeriesCommentSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSeriesHistoryPost

> SeriesHistorySearchResponseV1 SearchSeriesHistoryPost(ctx, id).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()

search series history

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
    id := int64(789) // int64 | Series id
    perPageSearchRequestV1 := *openapiclient.NewPerPageSearchRequestV1() // PerPageSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.SearchSeriesHistoryPost(context.Background(), id).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.SearchSeriesHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSeriesHistoryPost`: SeriesHistorySearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.SearchSeriesHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSeriesHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPageSearchRequestV1** | [**PerPageSearchRequestV1**](PerPageSearchRequestV1.md) |  | 

### Return type

[**SeriesHistorySearchResponseV1**](SeriesHistorySearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSeriesPost

> SeriesSearchResponseV1 SearchSeriesPost(ctx).SeriesSearchRequestV1(seriesSearchRequestV1).Execute()

search series

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
    seriesSearchRequestV1 := *openapiclient.NewSeriesSearchRequestV1() // SeriesSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.SearchSeriesPost(context.Background()).SeriesSearchRequestV1(seriesSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.SearchSeriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSeriesPost`: SeriesSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.SearchSeriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSeriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesSearchRequestV1** | [**SeriesSearchRequestV1**](SeriesSearchRequestV1.md) |  | 

### Return type

[**SeriesSearchResponseV1**](SeriesSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesCommentsModerationPost

> SeriesCommentModerationResponseV1 SeriesCommentsModerationPost(ctx).SeriesCommentSearchRequestV1(seriesCommentSearchRequestV1).Execute()

moderate series comments

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
    seriesCommentSearchRequestV1 := *openapiclient.NewSeriesCommentSearchRequestV1() // SeriesCommentSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.SeriesCommentsModerationPost(context.Background()).SeriesCommentSearchRequestV1(seriesCommentSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.SeriesCommentsModerationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SeriesCommentsModerationPost`: SeriesCommentModerationResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.SeriesCommentsModerationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeriesCommentsModerationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesCommentSearchRequestV1** | [**SeriesCommentSearchRequestV1**](SeriesCommentSearchRequestV1.md) |  | 

### Return type

[**SeriesCommentModerationResponseV1**](SeriesCommentModerationResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesReleaseRssFeed

> string SeriesReleaseRssFeed(ctx, id).Execute()

releases rss feed for a specific series

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
    id := int64(789) // int64 | id of series

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.SeriesReleaseRssFeed(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.SeriesReleaseRssFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SeriesReleaseRssFeed`: string
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.SeriesReleaseRssFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiSeriesReleaseRssFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UnlockSeriesField

> ApiResponseV1 UnlockSeriesField(ctx, id, item).Execute()

unlock a field of an series

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
    id := int64(789) // int64 | id of series
    item := "item_example" // string | field name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.UnlockSeriesField(context.Background(), id, item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.UnlockSeriesField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockSeriesField`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.UnlockSeriesField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 
**item** | **string** | field name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockSeriesFieldRequest struct via the builder pattern


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


## UpdateSeries

> ApiResponseV1 UpdateSeries(ctx, id).SeriesModelUpdateV1(seriesModelUpdateV1).Execute()

update an series

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
    id := int64(56) // int64 | id of series
    seriesModelUpdateV1 := *openapiclient.NewSeriesModelUpdateV1() // SeriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.UpdateSeries(context.Background(), id).SeriesModelUpdateV1(seriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.UpdateSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSeries`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.UpdateSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesModelUpdateV1** | [**SeriesModelUpdateV1**](SeriesModelUpdateV1.md) |  | 

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


## UpdateSeriesComment

> ApiResponseV1 UpdateSeriesComment(ctx, id, commentId).SeriesCommentModelUpdateV1(seriesCommentModelUpdateV1).Execute()

update a series comment

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
    id := int64(56) // int64 | id of series
    commentId := int64(56) // int64 | id of series comment
    seriesCommentModelUpdateV1 := *openapiclient.NewSeriesCommentModelUpdateV1() // SeriesCommentModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.UpdateSeriesComment(context.Background(), id, commentId).SeriesCommentModelUpdateV1(seriesCommentModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.UpdateSeriesComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSeriesComment`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.UpdateSeriesComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 
**commentId** | **int64** | id of series comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSeriesCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seriesCommentModelUpdateV1** | [**SeriesCommentModelUpdateV1**](SeriesCommentModelUpdateV1.md) |  | 

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


## UpdateSeriesImage

> ApiResponseV1 UpdateSeriesImage(ctx, id).Image(image).Execute()

update the image of an series

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
    id := int64(789) // int64 | id of series
    image := os.NewFile(1234, "some_file") // *os.File | Image to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.UpdateSeriesImage(context.Background(), id).Image(image).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.UpdateSeriesImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSeriesImage`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.UpdateSeriesImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSeriesImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | ***os.File** | Image to update | 

### Return type

[**ApiResponseV1**](ApiResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSeriesRating

> ApiResponseV1 UpdateUserSeriesRating(ctx, id).SeriesRatingModelV1(seriesRatingModelV1).Execute()

update the user rating for a series

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
    id := int64(789) // int64 | id of series
    seriesRatingModelV1 := *openapiclient.NewSeriesRatingModelV1(float32(123)) // SeriesRatingModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesAPI.UpdateUserSeriesRating(context.Background(), id).SeriesRatingModelV1(seriesRatingModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.UpdateUserSeriesRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserSeriesRating`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.UpdateUserSeriesRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of series | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSeriesRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesRatingModelV1** | [**SeriesRatingModelV1**](SeriesRatingModelV1.md) |  | 

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

