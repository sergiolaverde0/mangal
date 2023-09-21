# \AuthorsAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthor**](AuthorsAPI.md#AddAuthor) | **Post** /authors | add an author
[**DeleteAuthor**](AuthorsAPI.md#DeleteAuthor) | **Delete** /authors/{id} | delete an author
[**DeleteImage**](AuthorsAPI.md#DeleteImage) | **Delete** /authors/{id}/image | delete the image of an author
[**LockAuthorField**](AuthorsAPI.md#LockAuthorField) | **Post** /authors/{id}/locks/{item}/lock | lock a field of an author
[**RetrieveAuthor**](AuthorsAPI.md#RetrieveAuthor) | **Get** /authors/{id} | get a specific author
[**RetrieveAuthorLocks**](AuthorsAPI.md#RetrieveAuthorLocks) | **Get** /authors/{id}/locks | get locks for a specific author
[**RetrieveAuthorSeries**](AuthorsAPI.md#RetrieveAuthorSeries) | **Post** /authors/{id}/series | get the list of series for a specific author
[**SearchAuthorsPost**](AuthorsAPI.md#SearchAuthorsPost) | **Post** /authors/search | search authors
[**UnlockAuthorField**](AuthorsAPI.md#UnlockAuthorField) | **Post** /authors/{id}/locks/{item}/unlock | unlock a field of an author
[**UpdateAuthor**](AuthorsAPI.md#UpdateAuthor) | **Patch** /authors/{id} | update an author
[**UpdateImage**](AuthorsAPI.md#UpdateImage) | **Post** /authors/{id}/image | update the image of an author



## AddAuthor

> ApiResponseV1 AddAuthor(ctx).AuthorsModelUpdateV1(authorsModelUpdateV1).Execute()

add an author

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
    authorsModelUpdateV1 := *openapiclient.NewAuthorsModelUpdateV1() // AuthorsModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.AddAuthor(context.Background()).AuthorsModelUpdateV1(authorsModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.AddAuthor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAuthor`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.AddAuthor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorsModelUpdateV1** | [**AuthorsModelUpdateV1**](AuthorsModelUpdateV1.md) |  | 

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


## DeleteAuthor

> ApiResponseV1 DeleteAuthor(ctx, id).Execute()

delete an author

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
    id := int64(56) // int64 | id of author

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.DeleteAuthor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.DeleteAuthor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAuthor`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.DeleteAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of author | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorRequest struct via the builder pattern


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


## DeleteImage

> ApiResponseV1 DeleteImage(ctx, id).Execute()

delete the image of an author

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
    id := int64(789) // int64 | id of author

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.DeleteImage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteImage`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.DeleteImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of author | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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


## LockAuthorField

> ApiResponseV1 LockAuthorField(ctx, id, item).AuthorsLockModelUpdateV1(authorsLockModelUpdateV1).Execute()

lock a field of an author

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
    id := int64(789) // int64 | id of author
    item := "item_example" // string | field name
    authorsLockModelUpdateV1 := *openapiclient.NewAuthorsLockModelUpdateV1() // AuthorsLockModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.LockAuthorField(context.Background(), id, item).AuthorsLockModelUpdateV1(authorsLockModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.LockAuthorField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockAuthorField`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.LockAuthorField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of author | 
**item** | **string** | field name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockAuthorFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorsLockModelUpdateV1** | [**AuthorsLockModelUpdateV1**](AuthorsLockModelUpdateV1.md) |  | 

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


## RetrieveAuthor

> AuthorsModelV1 RetrieveAuthor(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific author

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
    id := int64(56) // int64 | Author id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.RetrieveAuthor(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.RetrieveAuthor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAuthor`: AuthorsModelV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.RetrieveAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Author id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**AuthorsModelV1**](AuthorsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAuthorLocks

> []AuthorsLockModelV1 RetrieveAuthorLocks(ctx, id).Execute()

get locks for a specific author

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
    id := int64(789) // int64 | Author id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.RetrieveAuthorLocks(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.RetrieveAuthorLocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAuthorLocks`: []AuthorsLockModelV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.RetrieveAuthorLocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Author id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAuthorLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AuthorsLockModelV1**](AuthorsLockModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAuthorSeries

> AuthorsSeriesListResponseV1 RetrieveAuthorSeries(ctx, id).AuthorsSeriesListRequestV1(authorsSeriesListRequestV1).Execute()

get the list of series for a specific author

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
    id := int64(789) // int64 | Author id
    authorsSeriesListRequestV1 := *openapiclient.NewAuthorsSeriesListRequestV1() // AuthorsSeriesListRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.RetrieveAuthorSeries(context.Background(), id).AuthorsSeriesListRequestV1(authorsSeriesListRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.RetrieveAuthorSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAuthorSeries`: AuthorsSeriesListResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.RetrieveAuthorSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Author id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAuthorSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorsSeriesListRequestV1** | [**AuthorsSeriesListRequestV1**](AuthorsSeriesListRequestV1.md) |  | 

### Return type

[**AuthorsSeriesListResponseV1**](AuthorsSeriesListResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAuthorsPost

> AuthorsSearchResponseV1 SearchAuthorsPost(ctx).AuthorsSearchRequestV1(authorsSearchRequestV1).Execute()

search authors

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
    authorsSearchRequestV1 := *openapiclient.NewAuthorsSearchRequestV1() // AuthorsSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.SearchAuthorsPost(context.Background()).AuthorsSearchRequestV1(authorsSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.SearchAuthorsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAuthorsPost`: AuthorsSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.SearchAuthorsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAuthorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorsSearchRequestV1** | [**AuthorsSearchRequestV1**](AuthorsSearchRequestV1.md) |  | 

### Return type

[**AuthorsSearchResponseV1**](AuthorsSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockAuthorField

> ApiResponseV1 UnlockAuthorField(ctx, id, item).Execute()

unlock a field of an author

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
    id := int64(789) // int64 | id of author
    item := "item_example" // string | field name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.UnlockAuthorField(context.Background(), id, item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.UnlockAuthorField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockAuthorField`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.UnlockAuthorField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of author | 
**item** | **string** | field name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockAuthorFieldRequest struct via the builder pattern


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


## UpdateAuthor

> ApiResponseV1 UpdateAuthor(ctx, id).AuthorsModelUpdateV1(authorsModelUpdateV1).Execute()

update an author

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
    id := int64(56) // int64 | id of author
    authorsModelUpdateV1 := *openapiclient.NewAuthorsModelUpdateV1() // AuthorsModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.UpdateAuthor(context.Background(), id).AuthorsModelUpdateV1(authorsModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.UpdateAuthor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthor`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.UpdateAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of author | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorsModelUpdateV1** | [**AuthorsModelUpdateV1**](AuthorsModelUpdateV1.md) |  | 

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


## UpdateImage

> ApiResponseV1 UpdateImage(ctx, id).Image(image).Execute()

update the image of an author

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
    id := int64(789) // int64 | id of author
    image := os.NewFile(1234, "some_file") // *os.File | Image to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorsAPI.UpdateImage(context.Background(), id).Image(image).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.UpdateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImage`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of author | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


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

