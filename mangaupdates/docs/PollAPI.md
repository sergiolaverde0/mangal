# \PollAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPoll**](PollAPI.md#AddPoll) | **Post** /poll | add a new poll
[**ArchivePoll**](PollAPI.md#ArchivePoll) | **Delete** /poll | archive the active poll
[**RetrieveOldPolls**](PollAPI.md#RetrieveOldPolls) | **Get** /poll/old | get old polls
[**RetrievePoll**](PollAPI.md#RetrievePoll) | **Get** /poll | get the active poll
[**RetrieveVoteStatus**](PollAPI.md#RetrieveVoteStatus) | **Get** /poll/vote/status | get information about whether the user has voted
[**VotePollAnswer**](PollAPI.md#VotePollAnswer) | **Post** /poll/vote/{answer_id} | vote in a poll answer
[**VotePollNoAnswer**](PollAPI.md#VotePollNoAnswer) | **Post** /poll/vote | vote in a poll



## AddPoll

> ApiResponseV1 AddPoll(ctx).PollModelUpdateV1(pollModelUpdateV1).Execute()

add a new poll

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
    pollModelUpdateV1 := *openapiclient.NewPollModelUpdateV1() // PollModelUpdateV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PollAPI.AddPoll(context.Background()).PollModelUpdateV1(pollModelUpdateV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PollAPI.AddPoll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPoll`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PollAPI.AddPoll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pollModelUpdateV1** | [**PollModelUpdateV1**](PollModelUpdateV1.md) |  | 

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


## ArchivePoll

> ApiResponseV1 ArchivePoll(ctx).Execute()

archive the active poll

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
    resp, r, err := apiClient.PollAPI.ArchivePoll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PollAPI.ArchivePoll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivePoll`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PollAPI.ArchivePoll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiArchivePollRequest struct via the builder pattern


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


## RetrieveOldPolls

> ApiResponseV1 RetrieveOldPolls(ctx).Execute()

get old polls

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
    resp, r, err := apiClient.PollAPI.RetrieveOldPolls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PollAPI.RetrieveOldPolls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveOldPolls`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PollAPI.RetrieveOldPolls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOldPollsRequest struct via the builder pattern


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


## RetrievePoll

> PollModelV1 RetrievePoll(ctx).Execute()

get the active poll

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
    resp, r, err := apiClient.PollAPI.RetrievePoll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PollAPI.RetrievePoll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePoll`: PollModelV1
    fmt.Fprintf(os.Stdout, "Response from `PollAPI.RetrievePoll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePollRequest struct via the builder pattern


### Return type

[**PollModelV1**](PollModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVoteStatus

> PollVoteStatusModelV1 RetrieveVoteStatus(ctx).Execute()

get information about whether the user has voted

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
    resp, r, err := apiClient.PollAPI.RetrieveVoteStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PollAPI.RetrieveVoteStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveVoteStatus`: PollVoteStatusModelV1
    fmt.Fprintf(os.Stdout, "Response from `PollAPI.RetrieveVoteStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVoteStatusRequest struct via the builder pattern


### Return type

[**PollVoteStatusModelV1**](PollVoteStatusModelV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VotePollAnswer

> ApiResponseV1 VotePollAnswer(ctx, answerId).Execute()

vote in a poll answer

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
    answerId := int64(56) // int64 | id of answer to vote for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PollAPI.VotePollAnswer(context.Background(), answerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PollAPI.VotePollAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VotePollAnswer`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PollAPI.VotePollAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**answerId** | **int64** | id of answer to vote for | 

### Other Parameters

Other parameters are passed through a pointer to a apiVotePollAnswerRequest struct via the builder pattern


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


## VotePollNoAnswer

> ApiResponseV1 VotePollNoAnswer(ctx).Execute()

vote in a poll

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
    resp, r, err := apiClient.PollAPI.VotePollNoAnswer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PollAPI.VotePollNoAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VotePollNoAnswer`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PollAPI.VotePollNoAnswer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVotePollNoAnswerRequest struct via the builder pattern


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

