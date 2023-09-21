# \ConvoAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbandonConvo**](ConvoAPI.md#AbandonConvo) | **Post** /convo/{id}/abandon | abandon a convo
[**AbandonConvoBulk**](ConvoAPI.md#AbandonConvoBulk) | **Post** /convo/bulk/abandon | abandon convos in bulk
[**AddConvo**](ConvoAPI.md#AddConvo) | **Post** /convo | add an convo
[**AddConvoMessage**](ConvoAPI.md#AddConvoMessage) | **Post** /convo/{id}/messages | add a message to a convo
[**ConvoInbox**](ConvoAPI.md#ConvoInbox) | **Get** /convo/inbox | display unread messages
[**ConvoInboxCount**](ConvoAPI.md#ConvoInboxCount) | **Get** /convo/inbox/count | retrieve number of unread messages
[**ConvoReceived**](ConvoAPI.md#ConvoReceived) | **Post** /convo/received | display received (and read) messages
[**ConvoSent**](ConvoAPI.md#ConvoSent) | **Post** /convo/sent | display sent messages
[**DeleteConvo**](ConvoAPI.md#DeleteConvo) | **Delete** /convo/{id} | delete a convo
[**DeleteConvoBulk**](ConvoAPI.md#DeleteConvoBulk) | **Post** /convo/bulk/delete | delete convos in bulk
[**IgnoreUser**](ConvoAPI.md#IgnoreUser) | **Post** /convo/ignore/{user_id} | ignore a user
[**InviteUserToConvo**](ConvoAPI.md#InviteUserToConvo) | **Post** /convo/{id}/invite | invite a user to a convo
[**IsUserIgnored**](ConvoAPI.md#IsUserIgnored) | **Get** /convo/ignore/{user_id} | return whether the user is ignored
[**JoinConvo**](ConvoAPI.md#JoinConvo) | **Post** /convo/{id}/join | join a convo
[**KickUserFromConvo**](ConvoAPI.md#KickUserFromConvo) | **Post** /convo/{id}/kick/{user_id} | kick a user from a convo
[**ListConvoMessages**](ConvoAPI.md#ListConvoMessages) | **Post** /convo/{id}/messages/list | list convo messages
[**RetrieveConvo**](ConvoAPI.md#RetrieveConvo) | **Get** /convo/{id} | get a specific convo
[**RetrieveConvoMessage**](ConvoAPI.md#RetrieveConvoMessage) | **Get** /convo/{id}/messages/{message_id} | get a specific convo message
[**RetrieveConvoMessageLocation**](ConvoAPI.md#RetrieveConvoMessageLocation) | **Get** /convo/{id}/messages/{message_id}/location | get a specific convo message location
[**RetrieveConvoParticipants**](ConvoAPI.md#RetrieveConvoParticipants) | **Get** /convo/{id}/participants | get list of convo participants
[**SearchConvoMessagesPost**](ConvoAPI.md#SearchConvoMessagesPost) | **Post** /convo/{id}/messages/search | search convo
[**SearchConvoPost**](ConvoAPI.md#SearchConvoPost) | **Post** /convo/search | search convo
[**UnIgnoreUser**](ConvoAPI.md#UnIgnoreUser) | **Delete** /convo/ignore/{user_id} | remove ignore for a user
[**UpdateConvo**](ConvoAPI.md#UpdateConvo) | **Patch** /convo/{id} | update a convo
[**UpdateConvoMessage**](ConvoAPI.md#UpdateConvoMessage) | **Patch** /convo/{id}/messages/{message_id} | update a convo message



## AbandonConvo

> ApiResponseV1 AbandonConvo(ctx, id).Execute()

abandon a convo

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
    id := int64(56) // int64 | Convo id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.AbandonConvo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.AbandonConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbandonConvo`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.AbandonConvo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbandonConvoRequest struct via the builder pattern


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


## AbandonConvoBulk

> ApiResponseV1 AbandonConvoBulk(ctx).ConvoBulkModelV1(convoBulkModelV1).Execute()

abandon convos in bulk

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
    convoBulkModelV1 := *openapiclient.NewConvoBulkModelV1() // ConvoBulkModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.AbandonConvoBulk(context.Background()).ConvoBulkModelV1(convoBulkModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.AbandonConvoBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbandonConvoBulk`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.AbandonConvoBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAbandonConvoBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convoBulkModelV1** | [**ConvoBulkModelV1**](ConvoBulkModelV1.md) |  | 

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


## AddConvo

> ApiResponseV1 AddConvo(ctx).ConvoModelAddV1(convoModelAddV1).Execute()

add an convo

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
    convoModelAddV1 := *openapiclient.NewConvoModelAddV1() // ConvoModelAddV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.AddConvo(context.Background()).ConvoModelAddV1(convoModelAddV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.AddConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConvo`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.AddConvo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConvoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convoModelAddV1** | [**ConvoModelAddV1**](ConvoModelAddV1.md) |  | 

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


## AddConvoMessage

> ApiResponseV1 AddConvoMessage(ctx, id).ConvoMessageModelUpdateV1(convoMessageModelUpdateV1).Execute()

add a message to a convo

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
    id := int64(56) // int64 | Convo id
    convoMessageModelUpdateV1 := *openapiclient.NewConvoMessageModelUpdateV1() // ConvoMessageModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.AddConvoMessage(context.Background(), id).ConvoMessageModelUpdateV1(convoMessageModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.AddConvoMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConvoMessage`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.AddConvoMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddConvoMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **convoMessageModelUpdateV1** | [**ConvoMessageModelUpdateV1**](ConvoMessageModelUpdateV1.md) |  | 

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


## ConvoInbox

> ConvoSearchResponseV1 ConvoInbox(ctx).Execute()

display unread messages

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
    resp, r, err := apiClient.ConvoAPI.ConvoInbox(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.ConvoInbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvoInbox`: ConvoSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.ConvoInbox`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConvoInboxRequest struct via the builder pattern


### Return type

[**ConvoSearchResponseV1**](ConvoSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvoInboxCount

> ConvoSearchResponseV1 ConvoInboxCount(ctx).Execute()

retrieve number of unread messages

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
    resp, r, err := apiClient.ConvoAPI.ConvoInboxCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.ConvoInboxCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvoInboxCount`: ConvoSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.ConvoInboxCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConvoInboxCountRequest struct via the builder pattern


### Return type

[**ConvoSearchResponseV1**](ConvoSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvoReceived

> ConvoSearchResponseV1 ConvoReceived(ctx).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()

display received (and read) messages

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
    perPageSearchRequestV1 := *openapiclient.NewPerPageSearchRequestV1() // PerPageSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.ConvoReceived(context.Background()).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.ConvoReceived``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvoReceived`: ConvoSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.ConvoReceived`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvoReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPageSearchRequestV1** | [**PerPageSearchRequestV1**](PerPageSearchRequestV1.md) |  | 

### Return type

[**ConvoSearchResponseV1**](ConvoSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvoSent

> ConvoSearchResponseV1 ConvoSent(ctx).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()

display sent messages

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
    perPageSearchRequestV1 := *openapiclient.NewPerPageSearchRequestV1() // PerPageSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.ConvoSent(context.Background()).PerPageSearchRequestV1(perPageSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.ConvoSent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConvoSent`: ConvoSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.ConvoSent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvoSentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPageSearchRequestV1** | [**PerPageSearchRequestV1**](PerPageSearchRequestV1.md) |  | 

### Return type

[**ConvoSearchResponseV1**](ConvoSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConvo

> ApiResponseV1 DeleteConvo(ctx, id).Execute()

delete a convo

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
    id := int64(56) // int64 | Convo id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.DeleteConvo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.DeleteConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConvo`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.DeleteConvo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConvoRequest struct via the builder pattern


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


## DeleteConvoBulk

> ApiResponseV1 DeleteConvoBulk(ctx).ConvoBulkModelV1(convoBulkModelV1).Execute()

delete convos in bulk

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
    convoBulkModelV1 := *openapiclient.NewConvoBulkModelV1() // ConvoBulkModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.DeleteConvoBulk(context.Background()).ConvoBulkModelV1(convoBulkModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.DeleteConvoBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConvoBulk`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.DeleteConvoBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConvoBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convoBulkModelV1** | [**ConvoBulkModelV1**](ConvoBulkModelV1.md) |  | 

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


## IgnoreUser

> ApiResponseV1 IgnoreUser(ctx, userId).Execute()

ignore a user

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
    resp, r, err := apiClient.ConvoAPI.IgnoreUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.IgnoreUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IgnoreUser`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.IgnoreUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIgnoreUserRequest struct via the builder pattern


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


## InviteUserToConvo

> ApiResponseV1 InviteUserToConvo(ctx, id).ConvoParticipantModelAddV1(convoParticipantModelAddV1).Execute()

invite a user to a convo

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
    id := int64(56) // int64 | Convo id
    convoParticipantModelAddV1 := []openapiclient.ConvoParticipantModelAddV1{*openapiclient.NewConvoParticipantModelAddV1()} // []ConvoParticipantModelAddV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.InviteUserToConvo(context.Background(), id).ConvoParticipantModelAddV1(convoParticipantModelAddV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.InviteUserToConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteUserToConvo`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.InviteUserToConvo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserToConvoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **convoParticipantModelAddV1** | [**[]ConvoParticipantModelAddV1**](ConvoParticipantModelAddV1.md) |  | 

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


## IsUserIgnored

> ConvoUserIgnoreModelV1 IsUserIgnored(ctx, userId).Execute()

return whether the user is ignored

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
    resp, r, err := apiClient.ConvoAPI.IsUserIgnored(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.IsUserIgnored``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsUserIgnored`: ConvoUserIgnoreModelV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.IsUserIgnored`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsUserIgnoredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConvoUserIgnoreModelV1**](ConvoUserIgnoreModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinConvo

> ApiResponseV1 JoinConvo(ctx, id).Execute()

join a convo

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
    id := int64(56) // int64 | Convo id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.JoinConvo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.JoinConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinConvo`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.JoinConvo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinConvoRequest struct via the builder pattern


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


## KickUserFromConvo

> ApiResponseV1 KickUserFromConvo(ctx, id, userId).Execute()

kick a user from a convo

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
    id := int64(56) // int64 | Convo id
    userId := int64(56) // int64 | User id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.KickUserFromConvo(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.KickUserFromConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KickUserFromConvo`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.KickUserFromConvo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiKickUserFromConvoRequest struct via the builder pattern


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


## ListConvoMessages

> ConvoMessageSearchResponseV1 ListConvoMessages(ctx, id).ConvoMessageListRequestV1(convoMessageListRequestV1).Execute()

list convo messages

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
    id := int64(56) // int64 | Convo id
    convoMessageListRequestV1 := *openapiclient.NewConvoMessageListRequestV1() // ConvoMessageListRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.ListConvoMessages(context.Background(), id).ConvoMessageListRequestV1(convoMessageListRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.ListConvoMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConvoMessages`: ConvoMessageSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.ListConvoMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConvoMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **convoMessageListRequestV1** | [**ConvoMessageListRequestV1**](ConvoMessageListRequestV1.md) |  | 

### Return type

[**ConvoMessageSearchResponseV1**](ConvoMessageSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveConvo

> ConvoModelV1 RetrieveConvo(ctx, id).UnrenderedFields(unrenderedFields).Execute()

get a specific convo

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
    id := int64(56) // int64 | Convo id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.RetrieveConvo(context.Background(), id).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.RetrieveConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveConvo`: ConvoModelV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.RetrieveConvo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConvoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ConvoModelV1**](ConvoModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveConvoMessage

> ConvoMessageModelV1 RetrieveConvoMessage(ctx, id, messageId).UnrenderedFields(unrenderedFields).Execute()

get a specific convo message

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
    id := int64(789) // int64 | Convo id
    messageId := int64(789) // int64 | Convo message id
    unrenderedFields := true // bool | Output fields in unrendered form for editing (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.RetrieveConvoMessage(context.Background(), id, messageId).UnrenderedFields(unrenderedFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.RetrieveConvoMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveConvoMessage`: ConvoMessageModelV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.RetrieveConvoMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 
**messageId** | **int64** | Convo message id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConvoMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unrenderedFields** | **bool** | Output fields in unrendered form for editing | 

### Return type

[**ConvoMessageModelV1**](ConvoMessageModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveConvoMessageLocation

> ApiResponseV1 RetrieveConvoMessageLocation(ctx, id, messageId).Execute()

get a specific convo message location

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
    id := int64(789) // int64 | Convo id
    messageId := int64(789) // int64 | Convo message id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.RetrieveConvoMessageLocation(context.Background(), id, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.RetrieveConvoMessageLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveConvoMessageLocation`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.RetrieveConvoMessageLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 
**messageId** | **int64** | Convo message id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConvoMessageLocationRequest struct via the builder pattern


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


## RetrieveConvoParticipants

> []ConvoParticipantModelV1 RetrieveConvoParticipants(ctx, id).Execute()

get list of convo participants

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
    id := int64(56) // int64 | Convo id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.RetrieveConvoParticipants(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.RetrieveConvoParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveConvoParticipants`: []ConvoParticipantModelV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.RetrieveConvoParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveConvoParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConvoParticipantModelV1**](ConvoParticipantModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchConvoMessagesPost

> ConvoMessageSearchResponseV1 SearchConvoMessagesPost(ctx, id).ConvoMessageSearchRequestV1(convoMessageSearchRequestV1).Execute()

search convo

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
    id := int64(56) // int64 | Convo id
    convoMessageSearchRequestV1 := *openapiclient.NewConvoMessageSearchRequestV1() // ConvoMessageSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.SearchConvoMessagesPost(context.Background(), id).ConvoMessageSearchRequestV1(convoMessageSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.SearchConvoMessagesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchConvoMessagesPost`: ConvoMessageSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.SearchConvoMessagesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchConvoMessagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **convoMessageSearchRequestV1** | [**ConvoMessageSearchRequestV1**](ConvoMessageSearchRequestV1.md) |  | 

### Return type

[**ConvoMessageSearchResponseV1**](ConvoMessageSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchConvoPost

> ConvoSearchResponseV1 SearchConvoPost(ctx).ConvoSearchRequestV1(convoSearchRequestV1).Execute()

search convo

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
    convoSearchRequestV1 := *openapiclient.NewConvoSearchRequestV1() // ConvoSearchRequestV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.SearchConvoPost(context.Background()).ConvoSearchRequestV1(convoSearchRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.SearchConvoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchConvoPost`: ConvoSearchResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.SearchConvoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchConvoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convoSearchRequestV1** | [**ConvoSearchRequestV1**](ConvoSearchRequestV1.md) |  | 

### Return type

[**ConvoSearchResponseV1**](ConvoSearchResponseV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnIgnoreUser

> ApiResponseV1 UnIgnoreUser(ctx, userId).Execute()

remove ignore for a user

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
    resp, r, err := apiClient.ConvoAPI.UnIgnoreUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.UnIgnoreUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnIgnoreUser`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.UnIgnoreUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnIgnoreUserRequest struct via the builder pattern


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


## UpdateConvo

> ApiResponseV1 UpdateConvo(ctx, id).ConvoModelUpdateV1(convoModelUpdateV1).Execute()

update a convo

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
    id := int64(56) // int64 | Convo id
    convoModelUpdateV1 := *openapiclient.NewConvoModelUpdateV1() // ConvoModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.UpdateConvo(context.Background(), id).ConvoModelUpdateV1(convoModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.UpdateConvo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConvo`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.UpdateConvo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConvoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **convoModelUpdateV1** | [**ConvoModelUpdateV1**](ConvoModelUpdateV1.md) |  | 

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


## UpdateConvoMessage

> ApiResponseV1 UpdateConvoMessage(ctx, id, messageId).ConvoMessageModelUpdateV1(convoMessageModelUpdateV1).Execute()

update a convo message

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
    id := int64(56) // int64 | Convo id
    messageId := int64(56) // int64 | Convo message id
    convoMessageModelUpdateV1 := *openapiclient.NewConvoMessageModelUpdateV1() // ConvoMessageModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvoAPI.UpdateConvoMessage(context.Background(), id, messageId).ConvoMessageModelUpdateV1(convoMessageModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvoAPI.UpdateConvoMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConvoMessage`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ConvoAPI.UpdateConvoMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Convo id | 
**messageId** | **int64** | Convo message id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConvoMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **convoMessageModelUpdateV1** | [**ConvoMessageModelUpdateV1**](ConvoMessageModelUpdateV1.md) |  | 

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

