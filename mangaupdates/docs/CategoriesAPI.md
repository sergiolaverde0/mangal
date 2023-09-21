# \CategoriesAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCombineSeriesCategories**](CategoriesAPI.md#BulkCombineSeriesCategories) | **Post** /categories/bulk/combine | combine two categories across the database
[**BulkDeleteSeriesCategories**](CategoriesAPI.md#BulkDeleteSeriesCategories) | **Post** /categories/bulk/delete | delete a category across the database
[**FindCategoryByExact**](CategoriesAPI.md#FindCategoryByExact) | **Post** /categories/findByExact | find a category by name
[**FindCategoryByPrefix**](CategoriesAPI.md#FindCategoryByPrefix) | **Post** /categories/findByPrefix | find a category by prefix
[**SearchCategoriesPost**](CategoriesAPI.md#SearchCategoriesPost) | **Post** /categories/search | search categories



## BulkCombineSeriesCategories

> ApiResponseV1 BulkCombineSeriesCategories(ctx).SeriesCategoryUpdateModelV1(seriesCategoryUpdateModelV1).Execute()

combine two categories across the database

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
    seriesCategoryUpdateModelV1 := *openapiclient.NewSeriesCategoryUpdateModelV1() // SeriesCategoryUpdateModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.BulkCombineSeriesCategories(context.Background()).SeriesCategoryUpdateModelV1(seriesCategoryUpdateModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.BulkCombineSeriesCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCombineSeriesCategories`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.BulkCombineSeriesCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCombineSeriesCategoriesRequest struct via the builder pattern


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


## BulkDeleteSeriesCategories

> ApiResponseV1 BulkDeleteSeriesCategories(ctx).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()

delete a category across the database

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
    categoriesModelUpdateV1 := *openapiclient.NewCategoriesModelUpdateV1() // CategoriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.BulkDeleteSeriesCategories(context.Background()).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.BulkDeleteSeriesCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteSeriesCategories`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.BulkDeleteSeriesCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteSeriesCategoriesRequest struct via the builder pattern


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


## FindCategoryByExact

> CategoriesModelV1 FindCategoryByExact(ctx).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()

find a category by name

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
    categoriesModelUpdateV1 := *openapiclient.NewCategoriesModelUpdateV1() // CategoriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.FindCategoryByExact(context.Background()).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.FindCategoryByExact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCategoryByExact`: CategoriesModelV1
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.FindCategoryByExact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindCategoryByExactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoriesModelUpdateV1** | [**CategoriesModelUpdateV1**](CategoriesModelUpdateV1.md) |  | 

### Return type

[**CategoriesModelV1**](CategoriesModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCategoryByPrefix

> []CategoriesModelV1 FindCategoryByPrefix(ctx).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()

find a category by prefix

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
    categoriesModelUpdateV1 := *openapiclient.NewCategoriesModelUpdateV1() // CategoriesModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.FindCategoryByPrefix(context.Background()).CategoriesModelUpdateV1(categoriesModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.FindCategoryByPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCategoryByPrefix`: []CategoriesModelV1
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.FindCategoryByPrefix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindCategoryByPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoriesModelUpdateV1** | [**CategoriesModelUpdateV1**](CategoriesModelUpdateV1.md) |  | 

### Return type

[**[]CategoriesModelV1**](CategoriesModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCategoriesPost

> CategoriesSearchResponseV1 SearchCategoriesPost(ctx).CategoriesSearchRequestV1(categoriesSearchRequestV1).Execute()

search categories

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
    categoriesSearchRequestV1 := *openapiclient.NewCategoriesSearchRequestV1() // CategoriesSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.SearchCategoriesPost(context.Background()).CategoriesSearchRequestV1(categoriesSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.SearchCategoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCategoriesPost`: CategoriesSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.SearchCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoriesSearchRequestV1** | [**CategoriesSearchRequestV1**](CategoriesSearchRequestV1.md) |  | 

### Return type

[**CategoriesSearchResponseV1**](CategoriesSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

