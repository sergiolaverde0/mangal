# \ForumAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddForumAdmin**](ForumAPI.md#AddForumAdmin) | **Put** /forums/{forum_id}/admins/{user_id} | add a forum admin
[**AddPollVote**](ForumAPI.md#AddPollVote) | **Post** /forums/{forum_id}/topics/{topic_id}/poll/vote/{choice_id} | add a vote to a forum poll
[**AddPost**](ForumAPI.md#AddPost) | **Post** /forums/{forum_id}/topics/{topic_id} | add a forum post
[**AddTemporaryPollImage**](ForumAPI.md#AddTemporaryPollImage) | **Post** /forums/temp_poll_images | add a temporary poll image
[**AddTopic**](ForumAPI.md#AddTopic) | **Post** /forums/{forum_id} | add a forum topic
[**DeletePost**](ForumAPI.md#DeletePost) | **Delete** /forums/{forum_id}/topics/{topic_id}/posts/{post_id} | delete a post
[**DeletePostReport**](ForumAPI.md#DeletePostReport) | **Delete** /forums/{forum_id}/topics/{topic_id}/posts/{post_id}/report | delete a post report
[**DeleteTopic**](ForumAPI.md#DeleteTopic) | **Delete** /forums/{forum_id}/topics/{topic_id} | delete a topic
[**GetCurrentWarnForUser**](ForumAPI.md#GetCurrentWarnForUser) | **Get** /forums/warn/{user_id} | gets the current warn status for user
[**ListCategories**](ForumAPI.md#ListCategories) | **Get** /forums | show forum categories and forums
[**ListGlobalTopics**](ForumAPI.md#ListGlobalTopics) | **Get** /forums/global | list global topics
[**ListPopularForums**](ForumAPI.md#ListPopularForums) | **Get** /forums/popular | show popular forums
[**ListPosts**](ForumAPI.md#ListPosts) | **Post** /forums/{forum_id}/topics/{topic_id}/list | list posts in topic
[**ListPostsByMe**](ForumAPI.md#ListPostsByMe) | **Get** /forums/{forum_id}/topics/{topic_id}/my_posts | list posts in topic that I made
[**ListReportedPosts**](ForumAPI.md#ListReportedPosts) | **Get** /forums/report | show reported posts
[**ListTopics**](ForumAPI.md#ListTopics) | **Post** /forums/{forum_id}/list | list topics
[**ListWarnHistoryForUser**](ForumAPI.md#ListWarnHistoryForUser) | **Get** /forums/warn/{user_id}/history | show warn history for a user
[**LookupPost**](ForumAPI.md#LookupPost) | **Get** /forums/lookup/post/{post_id} | lookup a post to find the forum and topic id
[**LookupSeries**](ForumAPI.md#LookupSeries) | **Get** /forums/lookup/series/{series_id} | lookup a series to find the forum id
[**LookupTopic**](ForumAPI.md#LookupTopic) | **Get** /forums/lookup/topic/{topic_id} | lookup a topic to find the forum id
[**RemoveForumAdmin**](ForumAPI.md#RemoveForumAdmin) | **Delete** /forums/{forum_id}/admins/{user_id} | remove a forum admin
[**ReportPost**](ForumAPI.md#ReportPost) | **Post** /forums/{forum_id}/topics/{topic_id}/posts/{post_id}/report | report a forum post
[**RetrieveForum**](ForumAPI.md#RetrieveForum) | **Get** /forums/{forum_id} | retrieve a forum
[**RetrievePost**](ForumAPI.md#RetrievePost) | **Get** /forums/{forum_id}/topics/{topic_id}/posts/{post_id} | retrieve a forum post
[**RetrievePostLocation**](ForumAPI.md#RetrievePostLocation) | **Get** /forums/{forum_id}/topics/{topic_id}/posts/{post_id}/location | retrieve a forum post location within topic
[**RetrieveTemporaryPollImages**](ForumAPI.md#RetrieveTemporaryPollImages) | **Get** /forums/temp_poll_images | retrieve temporary poll images
[**RetrieveTopic**](ForumAPI.md#RetrieveTopic) | **Get** /forums/{forum_id}/topics/{topic_id} | retrieve a forum topic
[**RetrieveVote**](ForumAPI.md#RetrieveVote) | **Get** /forums/{forum_id}/topics/{topic_id}/poll/vote | retrieve my vote from the poll
[**SearchForumPost**](ForumAPI.md#SearchForumPost) | **Post** /forums/search | search forum
[**SearchSpecificForumPost**](ForumAPI.md#SearchSpecificForumPost) | **Post** /forums/{forum_id}/search | search specific forum
[**SearchSpecificTopicPost**](ForumAPI.md#SearchSpecificTopicPost) | **Post** /forums/{forum_id}/topics/{topic_id}/search | search specific topic
[**ShowLogPost**](ForumAPI.md#ShowLogPost) | **Post** /forums/log | show forum admin log
[**UpdatePost**](ForumAPI.md#UpdatePost) | **Patch** /forums/{forum_id}/topics/{topic_id}/posts/{post_id} | update a forum post
[**UpdateTopic**](ForumAPI.md#UpdateTopic) | **Patch** /forums/{forum_id}/topics/{topic_id} | update a forum topic
[**UpdateTopicPoll**](ForumAPI.md#UpdateTopicPoll) | **Patch** /forums/{forum_id}/topics/{topic_id}/poll | update a forum topic poll (if present)
[**UpdateUserWarnLevel**](ForumAPI.md#UpdateUserWarnLevel) | **Put** /forums/warn/{user_id} | update a user warn level



## AddForumAdmin

> ApiResponseV1 AddForumAdmin(ctx, forumId, userId).Execute()

add a forum admin

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
    forumId := int64(56) // int64 | Forum id
    userId := int64(56) // int64 | User id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.AddForumAdmin(context.Background(), forumId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.AddForumAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddForumAdmin`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.AddForumAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddForumAdminRequest struct via the builder pattern


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


## AddPollVote

> ApiResponseV1 AddPollVote(ctx, forumId, topicId, choiceId).Execute()

add a vote to a forum poll

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    choiceId := int64(56) // int64 | Choice id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.AddPollVote(context.Background(), forumId, topicId, choiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.AddPollVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPollVote`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.AddPollVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 
**choiceId** | **int64** | Choice id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPollVoteRequest struct via the builder pattern


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


## AddPost

> ApiResponseV1 AddPost(ctx, forumId, topicId).ForumPostModelUpdateV1(forumPostModelUpdateV1).Execute()

add a forum post

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    forumPostModelUpdateV1 := *openapiclient.NewForumPostModelUpdateV1() // ForumPostModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.AddPost(context.Background(), forumId, topicId).ForumPostModelUpdateV1(forumPostModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.AddPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPost`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.AddPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forumPostModelUpdateV1** | [**ForumPostModelUpdateV1**](ForumPostModelUpdateV1.md) |  | 

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


## AddTemporaryPollImage

> ApiResponseV1 AddTemporaryPollImage(ctx).Image(image).Caption(caption).Execute()

add a temporary poll image

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
    image := os.NewFile(1234, "some_file") // *os.File | Image to update (optional)
    caption := "caption_example" // string | Image caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.AddTemporaryPollImage(context.Background()).Image(image).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.AddTemporaryPollImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTemporaryPollImage`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.AddTemporaryPollImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTemporaryPollImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | ***os.File** | Image to update | 
 **caption** | **string** | Image caption | 

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


## AddTopic

> ApiResponseV1 AddTopic(ctx, forumId).ForumTopicModelAddV1(forumTopicModelAddV1).Execute()

add a forum topic

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
    forumId := int64(56) // int64 | Forum id
    forumTopicModelAddV1 := *openapiclient.NewForumTopicModelAddV1() // ForumTopicModelAddV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.AddTopic(context.Background(), forumId).ForumTopicModelAddV1(forumTopicModelAddV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.AddTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTopic`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.AddTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forumTopicModelAddV1** | [**ForumTopicModelAddV1**](ForumTopicModelAddV1.md) |  | 

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


## DeletePost

> ApiResponseV1 DeletePost(ctx, forumId, topicId, postId).Execute()

delete a post

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    postId := int64(56) // int64 | Post id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.DeletePost(context.Background(), forumId, topicId, postId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.DeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePost`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.DeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 
**postId** | **int64** | Post id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostRequest struct via the builder pattern


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


## DeletePostReport

> ApiResponseV1 DeletePostReport(ctx, forumId, topicId, postId).Execute()

delete a post report

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    postId := int64(56) // int64 | Post id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.DeletePostReport(context.Background(), forumId, topicId, postId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.DeletePostReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePostReport`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.DeletePostReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 
**postId** | **int64** | Post id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostReportRequest struct via the builder pattern


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


## DeleteTopic

> ApiResponseV1 DeleteTopic(ctx, forumId, topicId).Execute()

delete a topic

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.DeleteTopic(context.Background(), forumId, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.DeleteTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTopic`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.DeleteTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopicRequest struct via the builder pattern


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


## GetCurrentWarnForUser

> ForumWarnModelV1 GetCurrentWarnForUser(ctx, userId).Execute()

gets the current warn status for user

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
    resp, r, err := apiClient.ForumAPI.GetCurrentWarnForUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.GetCurrentWarnForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentWarnForUser`: ForumWarnModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.GetCurrentWarnForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentWarnForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumWarnModelV1**](ForumWarnModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategories

> []ForumCategoryModelListV1 ListCategories(ctx).Execute()

show forum categories and forums

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
    resp, r, err := apiClient.ForumAPI.ListCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategories`: []ForumCategoryModelListV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


### Return type

[**[]ForumCategoryModelListV1**](ForumCategoryModelListV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalTopics

> ForumTopicListResponseV1 ListGlobalTopics(ctx).Execute()

list global topics

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
    resp, r, err := apiClient.ForumAPI.ListGlobalTopics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListGlobalTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalTopics`: ForumTopicListResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListGlobalTopics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalTopicsRequest struct via the builder pattern


### Return type

[**ForumTopicListResponseV1**](ForumTopicListResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPopularForums

> []ForumForumModelListV1 ListPopularForums(ctx).Execute()

show popular forums

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
    resp, r, err := apiClient.ForumAPI.ListPopularForums(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListPopularForums``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPopularForums`: []ForumForumModelListV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListPopularForums`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPopularForumsRequest struct via the builder pattern


### Return type

[**[]ForumForumModelListV1**](ForumForumModelListV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPosts

> ForumPostListResponseV1 ListPosts(ctx, forumId, topicId).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()

list posts in topic

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    perPageSearchRequestV1 := *openapiclient.NewPerPageSearchRequestV1() // PerPageSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.ListPosts(context.Background(), forumId, topicId).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPosts`: ForumPostListResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPageSearchRequestV1** | [**PerPageSearchRequestV1**](PerPageSearchRequestV1.md) |  | 

### Return type

[**ForumPostListResponseV1**](ForumPostListResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostsByMe

> ForumPostByUserResponseV1 ListPostsByMe(ctx, forumId, topicId).Execute()

list posts in topic that I made

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.ListPostsByMe(context.Background(), forumId, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListPostsByMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPostsByMe`: ForumPostByUserResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListPostsByMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPostsByMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ForumPostByUserResponseV1**](ForumPostByUserResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReportedPosts

> []ForumPostReportModelV1 ListReportedPosts(ctx).Execute()

show reported posts

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
    resp, r, err := apiClient.ForumAPI.ListReportedPosts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListReportedPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReportedPosts`: []ForumPostReportModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListReportedPosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListReportedPostsRequest struct via the builder pattern


### Return type

[**[]ForumPostReportModelV1**](ForumPostReportModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTopics

> ForumTopicListResponseV1 ListTopics(ctx, forumId).ForumTopicListRequestV1(forumTopicListRequestV1).WithFirstPost(withFirstPost).Execute()

list topics

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
    forumId := int64(56) // int64 | Forum id
    forumTopicListRequestV1 := *openapiclient.NewForumTopicListRequestV1() // ForumTopicListRequestV1 | 
    withFirstPost := true // bool | Also return the first post of each topic (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.ListTopics(context.Background(), forumId).ForumTopicListRequestV1(forumTopicListRequestV1).WithFirstPost(withFirstPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTopics`: ForumTopicListResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forumTopicListRequestV1** | [**ForumTopicListRequestV1**](ForumTopicListRequestV1.md) |  | 
 **withFirstPost** | **bool** | Also return the first post of each topic | 

### Return type

[**ForumTopicListResponseV1**](ForumTopicListResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWarnHistoryForUser

> []ForumWarnModelV1 ListWarnHistoryForUser(ctx, userId).Execute()

show warn history for a user

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
    resp, r, err := apiClient.ForumAPI.ListWarnHistoryForUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ListWarnHistoryForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWarnHistoryForUser`: []ForumWarnModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ListWarnHistoryForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWarnHistoryForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ForumWarnModelV1**](ForumWarnModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupPost

> ForumLookupResponseV1 LookupPost(ctx, postId).Execute()

lookup a post to find the forum and topic id

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
    postId := int64(56) // int64 | Post id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.LookupPost(context.Background(), postId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.LookupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupPost`: ForumLookupResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.LookupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **int64** | Post id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumLookupResponseV1**](ForumLookupResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupSeries

> ForumLookupResponseV1 LookupSeries(ctx, seriesId).Execute()

lookup a series to find the forum id

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.LookupSeries(context.Background(), seriesId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.LookupSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupSeries`: ForumLookupResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.LookupSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesId** | **int64** | Series id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumLookupResponseV1**](ForumLookupResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupTopic

> ForumLookupResponseV1 LookupTopic(ctx, topicId).Execute()

lookup a topic to find the forum id

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
    topicId := int64(56) // int64 | Topic id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.LookupTopic(context.Background(), topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.LookupTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupTopic`: ForumLookupResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.LookupTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumLookupResponseV1**](ForumLookupResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveForumAdmin

> ApiResponseV1 RemoveForumAdmin(ctx, forumId, userId).Execute()

remove a forum admin

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
    forumId := int64(56) // int64 | Forum id
    userId := int64(56) // int64 | User id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.RemoveForumAdmin(context.Background(), forumId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.RemoveForumAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveForumAdmin`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.RemoveForumAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveForumAdminRequest struct via the builder pattern


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


## ReportPost

> ApiResponseV1 ReportPost(ctx, forumId, topicId, postId).ForumPostReportModelUpdateV1(forumPostReportModelUpdateV1).Execute()

report a forum post

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    postId := int64(56) // int64 | Post id
    forumPostReportModelUpdateV1 := *openapiclient.NewForumPostReportModelUpdateV1() // ForumPostReportModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.ReportPost(context.Background(), forumId, topicId, postId).ForumPostReportModelUpdateV1(forumPostReportModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ReportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportPost`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ReportPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 
**postId** | **int64** | Post id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forumPostReportModelUpdateV1** | [**ForumPostReportModelUpdateV1**](ForumPostReportModelUpdateV1.md) |  | 

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


## RetrieveForum

> ForumForumModelV1 RetrieveForum(ctx, forumId).UnrenderedFields(unrenderedFields).Execute()

retrieve a forum

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
    forumId := int64(56) // int64 | Forum id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.RetrieveForum(context.Background(), forumId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.RetrieveForum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveForum`: ForumForumModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.RetrieveForum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveForumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ForumForumModelV1**](ForumForumModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePost

> ForumPostModelV1 RetrievePost(ctx, forumId, topicId, postId).UnrenderedFields(unrenderedFields).Execute()

retrieve a forum post

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    postId := int64(56) // int64 | Post id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.RetrievePost(context.Background(), forumId, topicId, postId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.RetrievePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePost`: ForumPostModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.RetrievePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 
**postId** | **int64** | Post id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ForumPostModelV1**](ForumPostModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePostLocation

> ApiResponseV1 RetrievePostLocation(ctx, forumId, topicId, postId).Execute()

retrieve a forum post location within topic

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    postId := int64(56) // int64 | Post id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.RetrievePostLocation(context.Background(), forumId, topicId, postId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.RetrievePostLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePostLocation`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.RetrievePostLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 
**postId** | **int64** | Post id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePostLocationRequest struct via the builder pattern


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


## RetrieveTemporaryPollImages

> []ForumPollTempImageModelV1 RetrieveTemporaryPollImages(ctx).Execute()

retrieve temporary poll images

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
    resp, r, err := apiClient.ForumAPI.RetrieveTemporaryPollImages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.RetrieveTemporaryPollImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveTemporaryPollImages`: []ForumPollTempImageModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.RetrieveTemporaryPollImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTemporaryPollImagesRequest struct via the builder pattern


### Return type

[**[]ForumPollTempImageModelV1**](ForumPollTempImageModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTopic

> ForumTopicModelV1 RetrieveTopic(ctx, forumId, topicId).UnrenderedFields(unrenderedFields).Execute()

retrieve a forum topic

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.RetrieveTopic(context.Background(), forumId, topicId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.RetrieveTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveTopic`: ForumTopicModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.RetrieveTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ForumTopicModelV1**](ForumTopicModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVote

> ForumPollVoteModelV1 RetrieveVote(ctx, forumId, topicId).Execute()

retrieve my vote from the poll

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.RetrieveVote(context.Background(), forumId, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.RetrieveVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveVote`: ForumPollVoteModelV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.RetrieveVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ForumPollVoteModelV1**](ForumPollVoteModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchForumPost

> ForumSearchResponseV1 SearchForumPost(ctx).ForumSearchRequestV1(forumSearchRequestV1).Execute()

search forum

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
    forumSearchRequestV1 := *openapiclient.NewForumSearchRequestV1() // ForumSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.SearchForumPost(context.Background()).ForumSearchRequestV1(forumSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.SearchForumPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchForumPost`: ForumSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.SearchForumPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forumSearchRequestV1** | [**ForumSearchRequestV1**](ForumSearchRequestV1.md) |  | 

### Return type

[**ForumSearchResponseV1**](ForumSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSpecificForumPost

> ForumSearchResponseV1 SearchSpecificForumPost(ctx, forumId).ForumSearchRequestV1(forumSearchRequestV1).Execute()

search specific forum

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
    forumId := int64(56) // int64 | Forum id
    forumSearchRequestV1 := *openapiclient.NewForumSearchRequestV1() // ForumSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.SearchSpecificForumPost(context.Background(), forumId).ForumSearchRequestV1(forumSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.SearchSpecificForumPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSpecificForumPost`: ForumSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.SearchSpecificForumPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpecificForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forumSearchRequestV1** | [**ForumSearchRequestV1**](ForumSearchRequestV1.md) |  | 

### Return type

[**ForumSearchResponseV1**](ForumSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSpecificTopicPost

> ForumSearchResponseV1 SearchSpecificTopicPost(ctx, forumId, topicId).ForumSearchRequestV1(forumSearchRequestV1).Execute()

search specific topic

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    forumSearchRequestV1 := *openapiclient.NewForumSearchRequestV1() // ForumSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.SearchSpecificTopicPost(context.Background(), forumId, topicId).ForumSearchRequestV1(forumSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.SearchSpecificTopicPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSpecificTopicPost`: ForumSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.SearchSpecificTopicPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpecificTopicPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forumSearchRequestV1** | [**ForumSearchRequestV1**](ForumSearchRequestV1.md) |  | 

### Return type

[**ForumSearchResponseV1**](ForumSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLogPost

> ForumAdminHistorySearchResponseV1 ShowLogPost(ctx).ForumAdminHistorySearchRequestV1(forumAdminHistorySearchRequestV1).Execute()

show forum admin log

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
    forumAdminHistorySearchRequestV1 := *openapiclient.NewForumAdminHistorySearchRequestV1() // ForumAdminHistorySearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.ShowLogPost(context.Background()).ForumAdminHistorySearchRequestV1(forumAdminHistorySearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.ShowLogPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowLogPost`: ForumAdminHistorySearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.ShowLogPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowLogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forumAdminHistorySearchRequestV1** | [**ForumAdminHistorySearchRequestV1**](ForumAdminHistorySearchRequestV1.md) |  | 

### Return type

[**ForumAdminHistorySearchResponseV1**](ForumAdminHistorySearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePost

> ApiResponseV1 UpdatePost(ctx, forumId, topicId, postId).ForumPostModelUpdateV1(forumPostModelUpdateV1).Execute()

update a forum post

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    postId := int64(56) // int64 | Post id
    forumPostModelUpdateV1 := *openapiclient.NewForumPostModelUpdateV1() // ForumPostModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.UpdatePost(context.Background(), forumId, topicId, postId).ForumPostModelUpdateV1(forumPostModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.UpdatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePost`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.UpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 
**postId** | **int64** | Post id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forumPostModelUpdateV1** | [**ForumPostModelUpdateV1**](ForumPostModelUpdateV1.md) |  | 

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


## UpdateTopic

> ApiResponseV1 UpdateTopic(ctx, forumId, topicId).ForumTopicModelUpdateV1(forumTopicModelUpdateV1).Execute()

update a forum topic

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    forumTopicModelUpdateV1 := *openapiclient.NewForumTopicModelUpdateV1() // ForumTopicModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.UpdateTopic(context.Background(), forumId, topicId).ForumTopicModelUpdateV1(forumTopicModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.UpdateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopic`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.UpdateTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forumTopicModelUpdateV1** | [**ForumTopicModelUpdateV1**](ForumTopicModelUpdateV1.md) |  | 

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


## UpdateTopicPoll

> ApiResponseV1 UpdateTopicPoll(ctx, forumId, topicId).ForumPollModelUpdateV1(forumPollModelUpdateV1).Execute()

update a forum topic poll (if present)

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
    forumId := int64(56) // int64 | Forum id
    topicId := int64(56) // int64 | Topic id
    forumPollModelUpdateV1 := *openapiclient.NewForumPollModelUpdateV1() // ForumPollModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.UpdateTopicPoll(context.Background(), forumId, topicId).ForumPollModelUpdateV1(forumPollModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.UpdateTopicPoll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopicPoll`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.UpdateTopicPoll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forumId** | **int64** | Forum id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopicPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forumPollModelUpdateV1** | [**ForumPollModelUpdateV1**](ForumPollModelUpdateV1.md) |  | 

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


## UpdateUserWarnLevel

> ApiResponseV1 UpdateUserWarnLevel(ctx, userId).ForumWarnModelUpdateV1(forumWarnModelUpdateV1).Execute()

update a user warn level

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
    forumWarnModelUpdateV1 := *openapiclient.NewForumWarnModelUpdateV1(int64(123), "Reason_example") // ForumWarnModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ForumAPI.UpdateUserWarnLevel(context.Background(), userId).ForumWarnModelUpdateV1(forumWarnModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ForumAPI.UpdateUserWarnLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserWarnLevel`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ForumAPI.UpdateUserWarnLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserWarnLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forumWarnModelUpdateV1** | [**ForumWarnModelUpdateV1**](ForumWarnModelUpdateV1.md) |  | 

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

