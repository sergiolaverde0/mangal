# \GenreAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGenre**](GenreAPI.md#AddGenre) | **Post** /genres | add a genre
[**DeleteGenre**](GenreAPI.md#DeleteGenre) | **Delete** /genres/{id} | delete a genre
[**RetrieveGenreById**](GenreAPI.md#RetrieveGenreById) | **Get** /genres/{id} | get genres
[**RetrieveGenres**](GenreAPI.md#RetrieveGenres) | **Get** /genres | get genres
[**UpdateGenre**](GenreAPI.md#UpdateGenre) | **Patch** /genres/{id} | update a genre



## AddGenre

> ApiResponseV1 AddGenre(ctx).GenreModelUpdateV1(genreModelUpdateV1).Execute()

add a genre

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
    genreModelUpdateV1 := *openapiclient.NewGenreModelUpdateV1() // GenreModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenreAPI.AddGenre(context.Background()).GenreModelUpdateV1(genreModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenreAPI.AddGenre``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGenre`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GenreAPI.AddGenre`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGenreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genreModelUpdateV1** | [**GenreModelUpdateV1**](GenreModelUpdateV1.md) |  | 

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


## DeleteGenre

> ApiResponseV1 DeleteGenre(ctx, id).Execute()

delete a genre

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
    id := int64(56) // int64 | id of genre

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenreAPI.DeleteGenre(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenreAPI.DeleteGenre``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGenre`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GenreAPI.DeleteGenre`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of genre | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGenreRequest struct via the builder pattern


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


## RetrieveGenreById

> GenreModelStatsV1 RetrieveGenreById(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get genres

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
    id := int64(56) // int64 | Genre id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenreAPI.RetrieveGenreById(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenreAPI.RetrieveGenreById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveGenreById`: GenreModelStatsV1
    fmt.Fprintf(os.Stdout, "Response from `GenreAPI.RetrieveGenreById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Genre id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveGenreByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**GenreModelStatsV1**](GenreModelStatsV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveGenres

> []GenreModelStatsV1 RetrieveGenres(ctx).Execute()

get genres

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
    resp, r, err := apiClient.GenreAPI.RetrieveGenres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenreAPI.RetrieveGenres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveGenres`: []GenreModelStatsV1
    fmt.Fprintf(os.Stdout, "Response from `GenreAPI.RetrieveGenres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveGenresRequest struct via the builder pattern


### Return type

[**[]GenreModelStatsV1**](GenreModelStatsV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGenre

> ApiResponseV1 UpdateGenre(ctx, id).GenreModelUpdateV1(genreModelUpdateV1).Execute()

update a genre

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
    id := int64(56) // int64 | id of genre
    genreModelUpdateV1 := *openapiclient.NewGenreModelUpdateV1() // GenreModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenreAPI.UpdateGenre(context.Background(), id).GenreModelUpdateV1(genreModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenreAPI.UpdateGenre``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGenre`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `GenreAPI.UpdateGenre`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of genre | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGenreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genreModelUpdateV1** | [**GenreModelUpdateV1**](GenreModelUpdateV1.md) |  | 

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

