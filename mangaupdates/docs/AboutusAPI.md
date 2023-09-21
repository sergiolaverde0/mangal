# \AboutusAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAboutusCategory**](AboutusAPI.md#AddAboutusCategory) | **Post** /aboutus/category | add a category
[**AddAboutusCategoryUser**](AboutusAPI.md#AddAboutusCategoryUser) | **Post** /aboutus/category/{category_id}/users | add a user to a category
[**DeleteAboutusCategory**](AboutusAPI.md#DeleteAboutusCategory) | **Delete** /aboutus/category/{category_id} | remove a category
[**DeleteAboutusCategoryUser**](AboutusAPI.md#DeleteAboutusCategoryUser) | **Delete** /aboutus/category/{category_id}/users/{entry_id} | remove a user from a category
[**ReorderAboutus**](AboutusAPI.md#ReorderAboutus) | **Post** /aboutus/reorder | reorder aboutus
[**RetrieveAboutusCategoriesAndUsers**](AboutusAPI.md#RetrieveAboutusCategoriesAndUsers) | **Get** /aboutus/users | returns categories and users
[**RetrieveAboutusCategory**](AboutusAPI.md#RetrieveAboutusCategory) | **Get** /aboutus/category/{category_id} | returns a single category
[**RetrieveAboutusDescription**](AboutusAPI.md#RetrieveAboutusDescription) | **Get** /aboutus | returns description of site
[**UpdateAboutusCategory**](AboutusAPI.md#UpdateAboutusCategory) | **Patch** /aboutus/category/{category_id} | update a category
[**UpdateAboutusDescription**](AboutusAPI.md#UpdateAboutusDescription) | **Post** /aboutus | update description of site



## AddAboutusCategory

> ApiResponseV1 AddAboutusCategory(ctx).AboutusCategoryModelUpdateV1(aboutusCategoryModelUpdateV1).Execute()

add a category

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
    aboutusCategoryModelUpdateV1 := *openapiclient.NewAboutusCategoryModelUpdateV1() // AboutusCategoryModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.AddAboutusCategory(context.Background()).AboutusCategoryModelUpdateV1(aboutusCategoryModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.AddAboutusCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAboutusCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.AddAboutusCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAboutusCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aboutusCategoryModelUpdateV1** | [**AboutusCategoryModelUpdateV1**](AboutusCategoryModelUpdateV1.md) |  | 

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


## AddAboutusCategoryUser

> ApiResponseV1 AddAboutusCategoryUser(ctx, categoryId).AboutusUserModelUpdateV1(aboutusUserModelUpdateV1).Execute()

add a user to a category

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
    categoryId := int64(56) // int64 | Aboutus Category id
    aboutusUserModelUpdateV1 := *openapiclient.NewAboutusUserModelUpdateV1() // AboutusUserModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.AddAboutusCategoryUser(context.Background(), categoryId).AboutusUserModelUpdateV1(aboutusUserModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.AddAboutusCategoryUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAboutusCategoryUser`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.AddAboutusCategoryUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Aboutus Category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAboutusCategoryUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aboutusUserModelUpdateV1** | [**AboutusUserModelUpdateV1**](AboutusUserModelUpdateV1.md) |  | 

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


## DeleteAboutusCategory

> ApiResponseV1 DeleteAboutusCategory(ctx, categoryId).Execute()

remove a category

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
    categoryId := int64(56) // int64 | Aboutus Category id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.DeleteAboutusCategory(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.DeleteAboutusCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAboutusCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.DeleteAboutusCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Aboutus Category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAboutusCategoryRequest struct via the builder pattern


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


## DeleteAboutusCategoryUser

> ApiResponseV1 DeleteAboutusCategoryUser(ctx, categoryId, entryId).Execute()

remove a user from a category

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
    categoryId := int64(56) // int64 | Aboutus Category id
    entryId := int64(56) // int64 | Aboutus Category User Entry id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.DeleteAboutusCategoryUser(context.Background(), categoryId, entryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.DeleteAboutusCategoryUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAboutusCategoryUser`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.DeleteAboutusCategoryUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Aboutus Category id | 
**entryId** | **int64** | Aboutus Category User Entry id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAboutusCategoryUserRequest struct via the builder pattern


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


## ReorderAboutus

> ApiResponseV1 ReorderAboutus(ctx).AboutusCategoryReorderModelV1(aboutusCategoryReorderModelV1).Execute()

reorder aboutus

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
    aboutusCategoryReorderModelV1 := []openapiclient.AboutusCategoryReorderModelV1{*openapiclient.NewAboutusCategoryReorderModelV1(int64(123))} // []AboutusCategoryReorderModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.ReorderAboutus(context.Background()).AboutusCategoryReorderModelV1(aboutusCategoryReorderModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.ReorderAboutus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReorderAboutus`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.ReorderAboutus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderAboutusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aboutusCategoryReorderModelV1** | [**[]AboutusCategoryReorderModelV1**](AboutusCategoryReorderModelV1.md) |  | 

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


## RetrieveAboutusCategoriesAndUsers

> []AboutusCategoryModelV1 RetrieveAboutusCategoriesAndUsers(ctx).Execute()

returns categories and users

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
    resp, r, err := apiClient.AboutusAPI.RetrieveAboutusCategoriesAndUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.RetrieveAboutusCategoriesAndUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAboutusCategoriesAndUsers`: []AboutusCategoryModelV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.RetrieveAboutusCategoriesAndUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAboutusCategoriesAndUsersRequest struct via the builder pattern


### Return type

[**[]AboutusCategoryModelV1**](AboutusCategoryModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAboutusCategory

> AboutusCategoryModelV1 RetrieveAboutusCategory(ctx, categoryId).UnrenderedFields(unrenderedFields).Execute()

returns a single category

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
    categoryId := int64(56) // int64 | Aboutus Category id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.RetrieveAboutusCategory(context.Background(), categoryId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.RetrieveAboutusCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAboutusCategory`: AboutusCategoryModelV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.RetrieveAboutusCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Aboutus Category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAboutusCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**AboutusCategoryModelV1**](AboutusCategoryModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAboutusDescription

> AboutusDescriptionModelV1 RetrieveAboutusDescription(ctx).UnrenderedFields(unrenderedFields).Execute()

returns description of site

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
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.RetrieveAboutusDescription(context.Background()).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.RetrieveAboutusDescription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAboutusDescription`: AboutusDescriptionModelV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.RetrieveAboutusDescription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAboutusDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**AboutusDescriptionModelV1**](AboutusDescriptionModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAboutusCategory

> ApiResponseV1 UpdateAboutusCategory(ctx, categoryId).AboutusCategoryModelUpdateV1(aboutusCategoryModelUpdateV1).Execute()

update a category

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
    categoryId := int64(56) // int64 | id of category
    aboutusCategoryModelUpdateV1 := *openapiclient.NewAboutusCategoryModelUpdateV1() // AboutusCategoryModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.UpdateAboutusCategory(context.Background(), categoryId).AboutusCategoryModelUpdateV1(aboutusCategoryModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.UpdateAboutusCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAboutusCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.UpdateAboutusCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | id of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAboutusCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aboutusCategoryModelUpdateV1** | [**AboutusCategoryModelUpdateV1**](AboutusCategoryModelUpdateV1.md) |  | 

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


## UpdateAboutusDescription

> ApiResponseV1 UpdateAboutusDescription(ctx).AboutusDescriptionModelV1(aboutusDescriptionModelV1).Execute()

update description of site

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
    aboutusDescriptionModelV1 := *openapiclient.NewAboutusDescriptionModelV1() // AboutusDescriptionModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AboutusAPI.UpdateAboutusDescription(context.Background()).AboutusDescriptionModelV1(aboutusDescriptionModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutusAPI.UpdateAboutusDescription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAboutusDescription`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AboutusAPI.UpdateAboutusDescription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAboutusDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aboutusDescriptionModelV1** | [**AboutusDescriptionModelV1**](AboutusDescriptionModelV1.md) |  | 

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

