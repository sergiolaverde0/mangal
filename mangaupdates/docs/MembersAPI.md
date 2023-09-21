# \MembersAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMember**](MembersAPI.md#AddMember) | **Post** /members | add a member
[**AddMemberAvatar**](MembersAPI.md#AddMemberAvatar) | **Post** /members/{id}/avatar | add a new member avatar
[**AddMemberChangeRequest**](MembersAPI.md#AddMemberChangeRequest) | **Post** /members/{id}/requests | add a change request
[**AddMemberGenreFilter**](MembersAPI.md#AddMemberGenreFilter) | **Post** /members/{id}/genre/{genre_id}/filter | filter a genre for a user
[**AddMemberGenreHighlight**](MembersAPI.md#AddMemberGenreHighlight) | **Post** /members/{id}/genre/{genre_id}/highlight | highlight a genre for a user
[**AddOrUpdateUserGroup**](MembersAPI.md#AddOrUpdateUserGroup) | **Put** /membergroups/{id} | add or update a user group
[**AddUserGroupFilter**](MembersAPI.md#AddUserGroupFilter) | **Post** /members/{id}/group/{group_id}/filter | filter a group for a user
[**AddUserTopicSubscription**](MembersAPI.md#AddUserTopicSubscription) | **Post** /members/{id}/topics/{topic_id} | add a topic subscription for a user
[**ApproveMemberUpgrade**](MembersAPI.md#ApproveMemberUpgrade) | **Post** /members/{id}/upgrade/approve | upgrade a member
[**DeleteMember**](MembersAPI.md#DeleteMember) | **Delete** /members/{id} | delete a member
[**DeleteMemberAvatar**](MembersAPI.md#DeleteMemberAvatar) | **Delete** /members/{id}/avatar/{avatar_id} | delete a member avatar
[**DeleteMemberChangeRequest**](MembersAPI.md#DeleteMemberChangeRequest) | **Delete** /members/{id}/requests/{request_id} | add a change request
[**DeleteUserGroup**](MembersAPI.md#DeleteUserGroup) | **Delete** /membergroups/{id} | delete a user group
[**RejectMemberUpgrade**](MembersAPI.md#RejectMemberUpgrade) | **Post** /members/{id}/upgrade/reject | reject a member upgrade
[**RemoveMemberGenreFilter**](MembersAPI.md#RemoveMemberGenreFilter) | **Delete** /members/{id}/genre/{genre_id}/filter | remove a filter for a genre for a user
[**RemoveMemberGenreHighlight**](MembersAPI.md#RemoveMemberGenreHighlight) | **Delete** /members/{id}/genre/{genre_id}/highlight | remove a highlight for a genre for a user
[**RemoveUserGroupFilter**](MembersAPI.md#RemoveUserGroupFilter) | **Delete** /members/{id}/group/{group_id}/filter | remove a filter for a group for a user
[**RemoveUserTopicSubscription**](MembersAPI.md#RemoveUserTopicSubscription) | **Delete** /members/{id}/topics/{topic_id} | remove a topic subscription for a user
[**ResetGenreSettings**](MembersAPI.md#ResetGenreSettings) | **Post** /members/{id}/genre/reset | reset genre highlights and filters for a user
[**RetrieveMember**](MembersAPI.md#RetrieveMember) | **Get** /members/{id} | get a specific members
[**RetrieveMemberAvatars**](MembersAPI.md#RetrieveMemberAvatars) | **Get** /members/{id}/avatars | get avatars for a specific user
[**RetrieveMemberChangeRequest**](MembersAPI.md#RetrieveMemberChangeRequest) | **Get** /members/{id}/requests/{request_id} | get change requests for a specific user
[**RetrieveMemberGenreFilters**](MembersAPI.md#RetrieveMemberGenreFilters) | **Get** /members/{id}/genre/filters | get genre filters for a specific user
[**RetrieveMemberGenreHighlights**](MembersAPI.md#RetrieveMemberGenreHighlights) | **Get** /members/{id}/genre/highlights | get highlights for a specific user
[**RetrieveMemberGroupFilters**](MembersAPI.md#RetrieveMemberGroupFilters) | **Get** /members/{id}/group/filters | get group filters for a specific user
[**RetrieveMemberTopicSubscription**](MembersAPI.md#RetrieveMemberTopicSubscription) | **Get** /members/{id}/topics/{topic_id} | get a subscription to a specific topic for a user
[**RetrieveMemberTopicSubscriptions**](MembersAPI.md#RetrieveMemberTopicSubscriptions) | **Get** /members/{id}/topics | get topic subscriptions for a specific user
[**RetrieveUserGroupById**](MembersAPI.md#RetrieveUserGroupById) | **Get** /membergroups/{id} | get user group
[**RetrieveUserGroups**](MembersAPI.md#RetrieveUserGroups) | **Get** /membergroups | get user groups
[**SearchMemberChangeRequests**](MembersAPI.md#SearchMemberChangeRequests) | **Get** /members/{id}/requests | search change requests for a specific user
[**SearchMembersPost**](MembersAPI.md#SearchMembersPost) | **Post** /members/search | search members
[**UpdateMember**](MembersAPI.md#UpdateMember) | **Patch** /members/{id} | update a member
[**UpdateMemberChangeRequest**](MembersAPI.md#UpdateMemberChangeRequest) | **Patch** /members/{id}/requests/{request_id} | update a change request



## AddMember

> ApiResponseV1 AddMember(ctx).UserModelUpdateV1(userModelUpdateV1).Execute()

add a member

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
    userModelUpdateV1 := *openapiclient.NewUserModelUpdateV1() // UserModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddMember(context.Background()).UserModelUpdateV1(userModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMember`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userModelUpdateV1** | [**UserModelUpdateV1**](UserModelUpdateV1.md) |  | 

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


## AddMemberAvatar

> ApiResponseV1 AddMemberAvatar(ctx, id).Image(image).Title(title).Execute()

add a new member avatar

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
    id := int64(789) // int64 | Member id
    image := os.NewFile(1234, "some_file") // *os.File | Image to update (optional)
    title := "title_example" // string | Title of the new avatar (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddMemberAvatar(context.Background(), id).Image(image).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddMemberAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMemberAvatar`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddMemberAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | ***os.File** | Image to update | 
 **title** | **string** | Title of the new avatar | 

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


## AddMemberChangeRequest

> ApiResponseV1 AddMemberChangeRequest(ctx, id).UserChangeRequestModelUpdateV1(userChangeRequestModelUpdateV1).Execute()

add a change request

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
    id := int64(789) // int64 | Member id
    userChangeRequestModelUpdateV1 := *openapiclient.NewUserChangeRequestModelUpdateV1() // UserChangeRequestModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddMemberChangeRequest(context.Background(), id).UserChangeRequestModelUpdateV1(userChangeRequestModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddMemberChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMemberChangeRequest`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddMemberChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userChangeRequestModelUpdateV1** | [**UserChangeRequestModelUpdateV1**](UserChangeRequestModelUpdateV1.md) |  | 

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


## AddMemberGenreFilter

> ApiResponseV1 AddMemberGenreFilter(ctx, id, genreId).Execute()

filter a genre for a user

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
    id := int64(789) // int64 | id of member
    genreId := int64(56) // int64 | genre id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddMemberGenreFilter(context.Background(), id, genreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddMemberGenreFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMemberGenreFilter`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddMemberGenreFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**genreId** | **int64** | genre id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberGenreFilterRequest struct via the builder pattern


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


## AddMemberGenreHighlight

> ApiResponseV1 AddMemberGenreHighlight(ctx, id, genreId).UserGenreHighlightModelUpdateV1(userGenreHighlightModelUpdateV1).Execute()

highlight a genre for a user

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
    id := int64(789) // int64 | id of member
    genreId := int64(56) // int64 | genre id
    userGenreHighlightModelUpdateV1 := *openapiclient.NewUserGenreHighlightModelUpdateV1() // UserGenreHighlightModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddMemberGenreHighlight(context.Background(), id, genreId).UserGenreHighlightModelUpdateV1(userGenreHighlightModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddMemberGenreHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMemberGenreHighlight`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddMemberGenreHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**genreId** | **int64** | genre id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberGenreHighlightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userGenreHighlightModelUpdateV1** | [**UserGenreHighlightModelUpdateV1**](UserGenreHighlightModelUpdateV1.md) |  | 

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


## AddOrUpdateUserGroup

> ApiResponseV1 AddOrUpdateUserGroup(ctx, id).UserGroupModelUpdateV1(userGroupModelUpdateV1).Execute()

add or update a user group

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
    id := "id_example" // string | user group id
    userGroupModelUpdateV1 := *openapiclient.NewUserGroupModelUpdateV1("Name_example", "Description_example") // UserGroupModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddOrUpdateUserGroup(context.Background(), id).UserGroupModelUpdateV1(userGroupModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddOrUpdateUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrUpdateUserGroup`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddOrUpdateUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | user group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrUpdateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userGroupModelUpdateV1** | [**UserGroupModelUpdateV1**](UserGroupModelUpdateV1.md) |  | 

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


## AddUserGroupFilter

> ApiResponseV1 AddUserGroupFilter(ctx, id, groupId).Execute()

filter a group for a user

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
    id := int64(789) // int64 | id of member
    groupId := int64(56) // int64 | group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddUserGroupFilter(context.Background(), id, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddUserGroupFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserGroupFilter`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddUserGroupFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**groupId** | **int64** | group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupFilterRequest struct via the builder pattern


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


## AddUserTopicSubscription

> ApiResponseV1 AddUserTopicSubscription(ctx, id, topicId).Execute()

add a topic subscription for a user

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
    id := int64(789) // int64 | id of member
    topicId := int64(56) // int64 | topic id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AddUserTopicSubscription(context.Background(), id, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddUserTopicSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserTopicSubscription`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AddUserTopicSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**topicId** | **int64** | topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserTopicSubscriptionRequest struct via the builder pattern


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


## ApproveMemberUpgrade

> ApiResponseV1 ApproveMemberUpgrade(ctx, id).Execute()

upgrade a member

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
    id := int64(56) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.ApproveMemberUpgrade(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ApproveMemberUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveMemberUpgrade`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ApproveMemberUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveMemberUpgradeRequest struct via the builder pattern


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


## DeleteMember

> ApiResponseV1 DeleteMember(ctx, id).Execute()

delete a member

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
    id := int64(56) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.DeleteMember(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMember`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.DeleteMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


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


## DeleteMemberAvatar

> ApiResponseV1 DeleteMemberAvatar(ctx, id, avatarId).Execute()

delete a member avatar

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
    id := int64(789) // int64 | Member id
    avatarId := int64(789) // int64 | Avatar id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.DeleteMemberAvatar(context.Background(), id, avatarId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteMemberAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMemberAvatar`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.DeleteMemberAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 
**avatarId** | **int64** | Avatar id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberAvatarRequest struct via the builder pattern


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


## DeleteMemberChangeRequest

> ApiResponseV1 DeleteMemberChangeRequest(ctx, id, requestId).Execute()

add a change request

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
    id := int64(789) // int64 | Member id
    requestId := int64(789) // int64 | Change request id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.DeleteMemberChangeRequest(context.Background(), id, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteMemberChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMemberChangeRequest`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.DeleteMemberChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 
**requestId** | **int64** | Change request id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberChangeRequestRequest struct via the builder pattern


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


## DeleteUserGroup

> ApiResponseV1 DeleteUserGroup(ctx, id).Execute()

delete a user group

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
    id := "id_example" // string | id of user group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.DeleteUserGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserGroup`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.DeleteUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of user group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupRequest struct via the builder pattern


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


## RejectMemberUpgrade

> ApiResponseV1 RejectMemberUpgrade(ctx, id).Execute()

reject a member upgrade

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
    id := int64(56) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RejectMemberUpgrade(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RejectMemberUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectMemberUpgrade`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RejectMemberUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectMemberUpgradeRequest struct via the builder pattern


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


## RemoveMemberGenreFilter

> ApiResponseV1 RemoveMemberGenreFilter(ctx, id, genreId).Execute()

remove a filter for a genre for a user

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
    id := int64(789) // int64 | id of member
    genreId := int64(56) // int64 | genre id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RemoveMemberGenreFilter(context.Background(), id, genreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RemoveMemberGenreFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMemberGenreFilter`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RemoveMemberGenreFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**genreId** | **int64** | genre id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMemberGenreFilterRequest struct via the builder pattern


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


## RemoveMemberGenreHighlight

> ApiResponseV1 RemoveMemberGenreHighlight(ctx, id, genreId).Execute()

remove a highlight for a genre for a user

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
    id := int64(789) // int64 | id of member
    genreId := int64(56) // int64 | genre id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RemoveMemberGenreHighlight(context.Background(), id, genreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RemoveMemberGenreHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMemberGenreHighlight`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RemoveMemberGenreHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**genreId** | **int64** | genre id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMemberGenreHighlightRequest struct via the builder pattern


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


## RemoveUserGroupFilter

> ApiResponseV1 RemoveUserGroupFilter(ctx, id, groupId).Execute()

remove a filter for a group for a user

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
    id := int64(789) // int64 | id of member
    groupId := int64(56) // int64 | group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RemoveUserGroupFilter(context.Background(), id, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RemoveUserGroupFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserGroupFilter`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RemoveUserGroupFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**groupId** | **int64** | group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserGroupFilterRequest struct via the builder pattern


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


## RemoveUserTopicSubscription

> ApiResponseV1 RemoveUserTopicSubscription(ctx, id, topicId).Execute()

remove a topic subscription for a user

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
    id := int64(789) // int64 | id of member
    topicId := int64(56) // int64 | topic id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RemoveUserTopicSubscription(context.Background(), id, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RemoveUserTopicSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserTopicSubscription`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RemoveUserTopicSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 
**topicId** | **int64** | topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserTopicSubscriptionRequest struct via the builder pattern


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


## ResetGenreSettings

> ApiResponseV1 ResetGenreSettings(ctx, id).Execute()

reset genre highlights and filters for a user

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
    id := int64(789) // int64 | id of member

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.ResetGenreSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ResetGenreSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetGenreSettings`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ResetGenreSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of member | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetGenreSettingsRequest struct via the builder pattern


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


## RetrieveMember

> UserModelV1 RetrieveMember(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific members

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
    id := int64(789) // int64 | Member id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMember(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMember`: UserModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**UserModelV1**](UserModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberAvatars

> []AvatarModelV1 RetrieveMemberAvatars(ctx, id).Execute()

get avatars for a specific user

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
    id := int64(789) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMemberAvatars(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMemberAvatars``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMemberAvatars`: []AvatarModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMemberAvatars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AvatarModelV1**](AvatarModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberChangeRequest

> UserChangeRequestModelV1 RetrieveMemberChangeRequest(ctx, id, requestId).Execute()

get change requests for a specific user

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
    id := int64(789) // int64 | Member id
    requestId := int64(789) // int64 | Change request id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMemberChangeRequest(context.Background(), id, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMemberChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMemberChangeRequest`: UserChangeRequestModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMemberChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 
**requestId** | **int64** | Change request id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserChangeRequestModelV1**](UserChangeRequestModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberGenreFilters

> []UserGenreFilterModelV1 RetrieveMemberGenreFilters(ctx, id).Execute()

get genre filters for a specific user

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
    id := int64(789) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMemberGenreFilters(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMemberGenreFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMemberGenreFilters`: []UserGenreFilterModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMemberGenreFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberGenreFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserGenreFilterModelV1**](UserGenreFilterModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberGenreHighlights

> []UserGenreHighlightModelV1 RetrieveMemberGenreHighlights(ctx, id).Execute()

get highlights for a specific user

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
    id := int64(789) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMemberGenreHighlights(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMemberGenreHighlights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMemberGenreHighlights`: []UserGenreHighlightModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMemberGenreHighlights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberGenreHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserGenreHighlightModelV1**](UserGenreHighlightModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberGroupFilters

> []UserGroupFilterModelV1 RetrieveMemberGroupFilters(ctx, id).Execute()

get group filters for a specific user

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
    id := int64(789) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMemberGroupFilters(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMemberGroupFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMemberGroupFilters`: []UserGroupFilterModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMemberGroupFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberGroupFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserGroupFilterModelV1**](UserGroupFilterModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberTopicSubscription

> UserSubscribedTopicModelV1 RetrieveMemberTopicSubscription(ctx, id, topicId).Execute()

get a subscription to a specific topic for a user

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
    id := int64(789) // int64 | Member id
    topicId := int64(789) // int64 | Topic id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMemberTopicSubscription(context.Background(), id, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMemberTopicSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMemberTopicSubscription`: UserSubscribedTopicModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMemberTopicSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 
**topicId** | **int64** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberTopicSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserSubscribedTopicModelV1**](UserSubscribedTopicModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberTopicSubscriptions

> []UserSubscribedTopicModelV1 RetrieveMemberTopicSubscriptions(ctx, id).Execute()

get topic subscriptions for a specific user

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
    id := int64(789) // int64 | Member id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveMemberTopicSubscriptions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveMemberTopicSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMemberTopicSubscriptions`: []UserSubscribedTopicModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveMemberTopicSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberTopicSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserSubscribedTopicModelV1**](UserSubscribedTopicModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUserGroupById

> UserGroupModelV1 RetrieveUserGroupById(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get user group

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
    id := "id_example" // string | user group id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.RetrieveUserGroupById(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveUserGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUserGroupById`: UserGroupModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveUserGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | user group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**UserGroupModelV1**](UserGroupModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUserGroups

> []UserGroupModelV1 RetrieveUserGroups(ctx).Execute()

get user groups

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
    resp, r, err := apiClient.MembersAPI.RetrieveUserGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.RetrieveUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUserGroups`: []UserGroupModelV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.RetrieveUserGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserGroupsRequest struct via the builder pattern


### Return type

[**[]UserGroupModelV1**](UserGroupModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMemberChangeRequests

> UserChangeRequestSearchResponseV1 SearchMemberChangeRequests(ctx, id).Page(page).Perpage(perpage).Orderby(orderby).Asc(asc).Execute()

search change requests for a specific user

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
    id := int64(789) // int64 | Member id
    page := int64(56) // int64 | Start page (optional)
    perpage := int64(56) // int64 | Items per page (optional)
    orderby := "orderby_example" // string | order by field (optional) (default to "time")
    asc := "asc_example" // string | Direction of results (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.SearchMemberChangeRequests(context.Background(), id).Page(page).Perpage(perpage).Orderby(orderby).Asc(asc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.SearchMemberChangeRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMemberChangeRequests`: UserChangeRequestSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.SearchMemberChangeRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchMemberChangeRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Start page | 
 **perpage** | **int64** | Items per page | 
 **orderby** | **string** | order by field | [default to &quot;time&quot;]
 **asc** | **string** | Direction of results | [default to &quot;asc&quot;]

### Return type

[**UserChangeRequestSearchResponseV1**](UserChangeRequestSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMembersPost

> UserSearchResponseV1 SearchMembersPost(ctx).UserSearchRequestV1(userSearchRequestV1).Execute()

search members

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
    userSearchRequestV1 := *openapiclient.NewUserSearchRequestV1() // UserSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.SearchMembersPost(context.Background()).UserSearchRequestV1(userSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.SearchMembersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMembersPost`: UserSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.SearchMembersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMembersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userSearchRequestV1** | [**UserSearchRequestV1**](UserSearchRequestV1.md) |  | 

### Return type

[**UserSearchResponseV1**](UserSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ApiResponseV1 UpdateMember(ctx, id).UserModelUpdateV1(userModelUpdateV1).Execute()

update a member

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
    id := int64(56) // int64 | Member id
    userModelUpdateV1 := *openapiclient.NewUserModelUpdateV1() // UserModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.UpdateMember(context.Background(), id).UserModelUpdateV1(userModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userModelUpdateV1** | [**UserModelUpdateV1**](UserModelUpdateV1.md) |  | 

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


## UpdateMemberChangeRequest

> ApiResponseV1 UpdateMemberChangeRequest(ctx, id, requestId).UserChangeRequestModelUpdateV1(userChangeRequestModelUpdateV1).Execute()

update a change request

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
    id := int64(789) // int64 | Member id
    requestId := int64(789) // int64 | Change request id
    userChangeRequestModelUpdateV1 := *openapiclient.NewUserChangeRequestModelUpdateV1() // UserChangeRequestModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.UpdateMemberChangeRequest(context.Background(), id, requestId).UserChangeRequestModelUpdateV1(userChangeRequestModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.UpdateMemberChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemberChangeRequest`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.UpdateMemberChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 
**requestId** | **int64** | Change request id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userChangeRequestModelUpdateV1** | [**UserChangeRequestModelUpdateV1**](UserChangeRequestModelUpdateV1.md) |  | 

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

