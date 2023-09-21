# \FaqAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFaqCategory**](FaqAPI.md#AddFaqCategory) | **Post** /faq | add a faq category
[**AddFaqQuestion**](FaqAPI.md#AddFaqQuestion) | **Post** /faq/{category_id}/questions | add a faq question
[**DeleteFaqCategory**](FaqAPI.md#DeleteFaqCategory) | **Delete** /faq/{category_id} | delete a faq category
[**DeleteFaqQuestion**](FaqAPI.md#DeleteFaqQuestion) | **Delete** /faq/{category_id}/questions/{question_id} | delete an faq
[**ReorderFaq**](FaqAPI.md#ReorderFaq) | **Post** /faq/reorder | reorder faq
[**RetrieveAllFaqCategoriesAndQuestions**](FaqAPI.md#RetrieveAllFaqCategoriesAndQuestions) | **Get** /faq | retrieve all categories and questions
[**RetrieveAllFaqCategoryQuestions**](FaqAPI.md#RetrieveAllFaqCategoryQuestions) | **Get** /faq/{category_id}/questions | retrieve all quesitons for a category
[**RetrieveFaqCategory**](FaqAPI.md#RetrieveFaqCategory) | **Get** /faq/{category_id} | get a specific category
[**RetrieveFaqQuestion**](FaqAPI.md#RetrieveFaqQuestion) | **Get** /faq/{category_id}/questions/{question_id} | get a specific question for a category
[**UpdateFaqCategory**](FaqAPI.md#UpdateFaqCategory) | **Patch** /faq/{category_id} | update a faq category
[**UpdateFaqQuestion**](FaqAPI.md#UpdateFaqQuestion) | **Patch** /faq/{category_id}/questions/{question_id} | update a faq question



## AddFaqCategory

> ApiResponseV1 AddFaqCategory(ctx).FaqCategoryModelUpdateV1(faqCategoryModelUpdateV1).Execute()

add a faq category

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
    faqCategoryModelUpdateV1 := *openapiclient.NewFaqCategoryModelUpdateV1() // FaqCategoryModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.AddFaqCategory(context.Background()).FaqCategoryModelUpdateV1(faqCategoryModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.AddFaqCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFaqCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.AddFaqCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFaqCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **faqCategoryModelUpdateV1** | [**FaqCategoryModelUpdateV1**](FaqCategoryModelUpdateV1.md) |  | 

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


## AddFaqQuestion

> ApiResponseV1 AddFaqQuestion(ctx, categoryId).FaqQuestionModelUpdateV1(faqQuestionModelUpdateV1).Execute()

add a faq question

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
    categoryId := int64(56) // int64 | id of category to add question to
    faqQuestionModelUpdateV1 := *openapiclient.NewFaqQuestionModelUpdateV1() // FaqQuestionModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.AddFaqQuestion(context.Background(), categoryId).FaqQuestionModelUpdateV1(faqQuestionModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.AddFaqQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFaqQuestion`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.AddFaqQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | id of category to add question to | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFaqQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **faqQuestionModelUpdateV1** | [**FaqQuestionModelUpdateV1**](FaqQuestionModelUpdateV1.md) |  | 

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


## DeleteFaqCategory

> ApiResponseV1 DeleteFaqCategory(ctx, categoryId).Execute()

delete a faq category

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
    categoryId := int64(56) // int64 | id of faq category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.DeleteFaqCategory(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.DeleteFaqCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFaqCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.DeleteFaqCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | id of faq category | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFaqCategoryRequest struct via the builder pattern


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


## DeleteFaqQuestion

> ApiResponseV1 DeleteFaqQuestion(ctx, categoryId, questionId).Execute()

delete an faq

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
    categoryId := int64(56) // int64 | Faq category id
    questionId := int64(56) // int64 | Faq question id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.DeleteFaqQuestion(context.Background(), categoryId, questionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.DeleteFaqQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFaqQuestion`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.DeleteFaqQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Faq category id | 
**questionId** | **int64** | Faq question id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFaqQuestionRequest struct via the builder pattern


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


## ReorderFaq

> ApiResponseV1 ReorderFaq(ctx).FaqCategoryReorderModelV1(faqCategoryReorderModelV1).Execute()

reorder faq

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
    faqCategoryReorderModelV1 := []openapiclient.FaqCategoryReorderModelV1{*openapiclient.NewFaqCategoryReorderModelV1(int64(123))} // []FaqCategoryReorderModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.ReorderFaq(context.Background()).FaqCategoryReorderModelV1(faqCategoryReorderModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.ReorderFaq``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReorderFaq`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.ReorderFaq`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderFaqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **faqCategoryReorderModelV1** | [**[]FaqCategoryReorderModelV1**](FaqCategoryReorderModelV1.md) |  | 

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


## RetrieveAllFaqCategoriesAndQuestions

> []FaqCategoryQuestionsModelV1 RetrieveAllFaqCategoriesAndQuestions(ctx).Execute()

retrieve all categories and questions

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
    resp, r, err := apiClient.FaqAPI.RetrieveAllFaqCategoriesAndQuestions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.RetrieveAllFaqCategoriesAndQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAllFaqCategoriesAndQuestions`: []FaqCategoryQuestionsModelV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.RetrieveAllFaqCategoriesAndQuestions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllFaqCategoriesAndQuestionsRequest struct via the builder pattern


### Return type

[**[]FaqCategoryQuestionsModelV1**](FaqCategoryQuestionsModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllFaqCategoryQuestions

> []FaqQuestionModelV1 RetrieveAllFaqCategoryQuestions(ctx, categoryId).Execute()

retrieve all quesitons for a category

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
    categoryId := int64(56) // int64 | Faq category id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.RetrieveAllFaqCategoryQuestions(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.RetrieveAllFaqCategoryQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAllFaqCategoryQuestions`: []FaqQuestionModelV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.RetrieveAllFaqCategoryQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Faq category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllFaqCategoryQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FaqQuestionModelV1**](FaqQuestionModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFaqCategory

> FaqCategoryModelV1 RetrieveFaqCategory(ctx, categoryId).UnrenderedFields(unrenderedFields).Execute()

get a specific category

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
    categoryId := int64(56) // int64 | Faq category id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.RetrieveFaqCategory(context.Background(), categoryId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.RetrieveFaqCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveFaqCategory`: FaqCategoryModelV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.RetrieveFaqCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Faq category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFaqCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**FaqCategoryModelV1**](FaqCategoryModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFaqQuestion

> FaqQuestionModelV1 RetrieveFaqQuestion(ctx, categoryId, questionId).UnrenderedFields(unrenderedFields).Execute()

get a specific question for a category

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
    categoryId := int64(56) // int64 | Faq category id
    questionId := int64(56) // int64 | Faq question id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.RetrieveFaqQuestion(context.Background(), categoryId, questionId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.RetrieveFaqQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveFaqQuestion`: FaqQuestionModelV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.RetrieveFaqQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Faq category id | 
**questionId** | **int64** | Faq question id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFaqQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**FaqQuestionModelV1**](FaqQuestionModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFaqCategory

> ApiResponseV1 UpdateFaqCategory(ctx, categoryId).FaqCategoryModelUpdateV1(faqCategoryModelUpdateV1).Execute()

update a faq category

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
    categoryId := int64(56) // int64 | id of faq category
    faqCategoryModelUpdateV1 := *openapiclient.NewFaqCategoryModelUpdateV1() // FaqCategoryModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.UpdateFaqCategory(context.Background(), categoryId).FaqCategoryModelUpdateV1(faqCategoryModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.UpdateFaqCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFaqCategory`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.UpdateFaqCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | id of faq category | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFaqCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **faqCategoryModelUpdateV1** | [**FaqCategoryModelUpdateV1**](FaqCategoryModelUpdateV1.md) |  | 

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


## UpdateFaqQuestion

> ApiResponseV1 UpdateFaqQuestion(ctx, categoryId, questionId).FaqQuestionModelUpdateV1(faqQuestionModelUpdateV1).Execute()

update a faq question

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
    categoryId := int64(56) // int64 | Faq category id
    questionId := int64(56) // int64 | Faq question id
    faqQuestionModelUpdateV1 := *openapiclient.NewFaqQuestionModelUpdateV1() // FaqQuestionModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaqAPI.UpdateFaqQuestion(context.Background(), categoryId, questionId).FaqQuestionModelUpdateV1(faqQuestionModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaqAPI.UpdateFaqQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFaqQuestion`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `FaqAPI.UpdateFaqQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Faq category id | 
**questionId** | **int64** | Faq question id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFaqQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **faqQuestionModelUpdateV1** | [**FaqQuestionModelUpdateV1**](FaqQuestionModelUpdateV1.md) |  | 

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

