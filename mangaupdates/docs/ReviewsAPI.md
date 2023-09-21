# \ReviewsAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReview**](ReviewsAPI.md#AddReview) | **Post** /reviews | add a review
[**AddReviewComment**](ReviewsAPI.md#AddReviewComment) | **Post** /reviews/{id}/comments | add a review comment
[**DeleteReview**](ReviewsAPI.md#DeleteReview) | **Delete** /reviews/{id} | delete a review
[**DeleteReviewComment**](ReviewsAPI.md#DeleteReviewComment) | **Delete** /reviews/{id}/comments/{comment_id} | delete a review comment
[**RetrieveReview**](ReviewsAPI.md#RetrieveReview) | **Get** /reviews/{id} | get a specific review
[**RetrieveReviewComment**](ReviewsAPI.md#RetrieveReviewComment) | **Get** /reviews/{id}/comments/{comment_id} | get a specific review comment
[**ReviewCommentsModerationPost**](ReviewsAPI.md#ReviewCommentsModerationPost) | **Post** /reviews/comments/moderation | moderate review comments
[**SearchReviewCommentsPost**](ReviewsAPI.md#SearchReviewCommentsPost) | **Post** /reviews/{id}/comments/search | search review comments
[**SearchReviewsPost**](ReviewsAPI.md#SearchReviewsPost) | **Post** /reviews/search | search reviews
[**UpdateReview**](ReviewsAPI.md#UpdateReview) | **Patch** /reviews/{id} | update a review
[**UpdateReviewComment**](ReviewsAPI.md#UpdateReviewComment) | **Patch** /reviews/{id}/comments/{comment_id} | update a review comment



## AddReview

> ApiResponseV1 AddReview(ctx).ReviewModelUpdateV1(reviewModelUpdateV1).Execute()

add a review

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
    reviewModelUpdateV1 := *openapiclient.NewReviewModelUpdateV1() // ReviewModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.AddReview(context.Background()).ReviewModelUpdateV1(reviewModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.AddReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddReview`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.AddReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewModelUpdateV1** | [**ReviewModelUpdateV1**](ReviewModelUpdateV1.md) |  | 

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


## AddReviewComment

> ApiResponseV1 AddReviewComment(ctx, id).ReviewCommentModelUpdateV1(reviewCommentModelUpdateV1).Execute()

add a review comment

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
    id := int64(789) // int64 | Review id
    reviewCommentModelUpdateV1 := *openapiclient.NewReviewCommentModelUpdateV1() // ReviewCommentModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.AddReviewComment(context.Background(), id).ReviewCommentModelUpdateV1(reviewCommentModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.AddReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddReviewComment`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.AddReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Review id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewCommentModelUpdateV1** | [**ReviewCommentModelUpdateV1**](ReviewCommentModelUpdateV1.md) |  | 

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


## DeleteReview

> ApiResponseV1 DeleteReview(ctx, id).Execute()

delete a review

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
    id := int64(56) // int64 | id of review

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.DeleteReview(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.DeleteReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReview`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.DeleteReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of review | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReviewRequest struct via the builder pattern


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


## DeleteReviewComment

> ApiResponseV1 DeleteReviewComment(ctx, id, commentId).Execute()

delete a review comment

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
    id := int64(56) // int64 | id of review
    commentId := int64(56) // int64 | id of review comment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.DeleteReviewComment(context.Background(), id, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.DeleteReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReviewComment`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.DeleteReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of review | 
**commentId** | **int64** | id of review comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReviewCommentRequest struct via the builder pattern


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


## RetrieveReview

> ReviewModelV1 RetrieveReview(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific review

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
    id := int64(789) // int64 | Review id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.RetrieveReview(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.RetrieveReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveReview`: ReviewModelV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.RetrieveReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Review id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ReviewModelV1**](ReviewModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveReviewComment

> ReviewCommentModelV1 RetrieveReviewComment(ctx, id, commentId).UnrenderedFields(unrenderedFields).Execute()

get a specific review comment

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
    id := int64(789) // int64 | Review id
    commentId := int64(789) // int64 | Review comment id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.RetrieveReviewComment(context.Background(), id, commentId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.RetrieveReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveReviewComment`: ReviewCommentModelV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.RetrieveReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Review id | 
**commentId** | **int64** | Review comment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ReviewCommentModelV1**](ReviewCommentModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewCommentsModerationPost

> ReviewCommentSearchResponseV1 ReviewCommentsModerationPost(ctx).ReviewCommentSearchRequestV1(reviewCommentSearchRequestV1).Execute()

moderate review comments

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
    reviewCommentSearchRequestV1 := *openapiclient.NewReviewCommentSearchRequestV1() // ReviewCommentSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.ReviewCommentsModerationPost(context.Background()).ReviewCommentSearchRequestV1(reviewCommentSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.ReviewCommentsModerationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReviewCommentsModerationPost`: ReviewCommentSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.ReviewCommentsModerationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReviewCommentsModerationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewCommentSearchRequestV1** | [**ReviewCommentSearchRequestV1**](ReviewCommentSearchRequestV1.md) |  | 

### Return type

[**ReviewCommentSearchResponseV1**](ReviewCommentSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReviewCommentsPost

> ReviewCommentSearchResponseV1 SearchReviewCommentsPost(ctx, id).ReviewCommentSearchRequestV1(reviewCommentSearchRequestV1).Execute()

search review comments

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
    id := int64(789) // int64 | Review id
    reviewCommentSearchRequestV1 := *openapiclient.NewReviewCommentSearchRequestV1() // ReviewCommentSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.SearchReviewCommentsPost(context.Background(), id).ReviewCommentSearchRequestV1(reviewCommentSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.SearchReviewCommentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchReviewCommentsPost`: ReviewCommentSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.SearchReviewCommentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Review id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchReviewCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewCommentSearchRequestV1** | [**ReviewCommentSearchRequestV1**](ReviewCommentSearchRequestV1.md) |  | 

### Return type

[**ReviewCommentSearchResponseV1**](ReviewCommentSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReviewsPost

> ReviewSearchResponseV1 SearchReviewsPost(ctx).ReviewSearchRequestV1(reviewSearchRequestV1).Execute()

search reviews

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
    reviewSearchRequestV1 := *openapiclient.NewReviewSearchRequestV1() // ReviewSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.SearchReviewsPost(context.Background()).ReviewSearchRequestV1(reviewSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.SearchReviewsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchReviewsPost`: ReviewSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.SearchReviewsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchReviewsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewSearchRequestV1** | [**ReviewSearchRequestV1**](ReviewSearchRequestV1.md) |  | 

### Return type

[**ReviewSearchResponseV1**](ReviewSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReview

> ApiResponseV1 UpdateReview(ctx, id).ReviewModelUpdateV1(reviewModelUpdateV1).Execute()

update a review

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
    id := int64(56) // int64 | id of review
    reviewModelUpdateV1 := *openapiclient.NewReviewModelUpdateV1() // ReviewModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.UpdateReview(context.Background(), id).ReviewModelUpdateV1(reviewModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.UpdateReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReview`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.UpdateReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of review | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewModelUpdateV1** | [**ReviewModelUpdateV1**](ReviewModelUpdateV1.md) |  | 

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


## UpdateReviewComment

> ApiResponseV1 UpdateReviewComment(ctx, id, commentId).ReviewCommentModelUpdateV1(reviewCommentModelUpdateV1).Execute()

update a review comment

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
    id := int64(56) // int64 | id of review
    commentId := int64(56) // int64 | id of review comment
    reviewCommentModelUpdateV1 := *openapiclient.NewReviewCommentModelUpdateV1() // ReviewCommentModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReviewsAPI.UpdateReviewComment(context.Background(), id, commentId).ReviewCommentModelUpdateV1(reviewCommentModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.UpdateReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReviewComment`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.UpdateReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of review | 
**commentId** | **int64** | id of review comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reviewCommentModelUpdateV1** | [**ReviewCommentModelUpdateV1**](ReviewCommentModelUpdateV1.md) |  | 

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

