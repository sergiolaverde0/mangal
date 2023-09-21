# \AccountAPI

All URIs are relative to *https://api.mangaupdates.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Captcha**](AccountAPI.md#Captcha) | **Get** /account/captcha | retrieve the public captcha key
[**ConfirmAndChangePassword**](AccountAPI.md#ConfirmAndChangePassword) | **Post** /account/forgotpass/confirm/{auth_hash} | update a password change using an auth hash
[**ConfirmDeleteAccount**](AccountAPI.md#ConfirmDeleteAccount) | **Post** /account/delete/confirm/{auth_hash} | confirm deletion of your account
[**ConfirmRegistration**](AccountAPI.md#ConfirmRegistration) | **Post** /account/register/confirm/{auth_hash} | confirm a new member registration
[**DeleteAccount**](AccountAPI.md#DeleteAccount) | **Post** /account/delete/{captcha_response} | delete your account
[**ForgotPassword**](AccountAPI.md#ForgotPassword) | **Post** /account/forgotpass/{captcha_response} | send a forgotten password email
[**Login**](AccountAPI.md#Login) | **Put** /account/login | create a session token
[**Logout**](AccountAPI.md#Logout) | **Post** /account/logout | remove a session token
[**Profile**](AccountAPI.md#Profile) | **Get** /account/profile | get the profile for the current user
[**RegisterMember**](AccountAPI.md#RegisterMember) | **Post** /account/register/{captcha_response} | register a new member
[**ResendAuthEmail**](AccountAPI.md#ResendAuthEmail) | **Post** /account/resendauth/{id} | send an auth email to a user
[**SendForgotEmail**](AccountAPI.md#SendForgotEmail) | **Post** /account/sendforgot/{id} | send a forgotten password email to a user



## Captcha

> ApiResponseV1 Captcha(ctx).Execute()

retrieve the public captcha key

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
    resp, r, err := apiClient.AccountAPI.Captcha(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Captcha``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Captcha`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Captcha`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCaptchaRequest struct via the builder pattern


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


## ConfirmAndChangePassword

> ApiResponseV1 ConfirmAndChangePassword(ctx, authHash).UserModelUpdatePasswordV1(userModelUpdatePasswordV1).Execute()

update a password change using an auth hash

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
    authHash := "authHash_example" // string | auth hash from email confirmation
    userModelUpdatePasswordV1 := *openapiclient.NewUserModelUpdatePasswordV1() // UserModelUpdatePasswordV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ConfirmAndChangePassword(context.Background(), authHash).UserModelUpdatePasswordV1(userModelUpdatePasswordV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ConfirmAndChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmAndChangePassword`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ConfirmAndChangePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authHash** | **string** | auth hash from email confirmation | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmAndChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userModelUpdatePasswordV1** | [**UserModelUpdatePasswordV1**](UserModelUpdatePasswordV1.md) |  | 

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


## ConfirmDeleteAccount

> ApiResponseV1 ConfirmDeleteAccount(ctx, authHash).Execute()

confirm deletion of your account

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
    authHash := "authHash_example" // string | auth hash from email confirmation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ConfirmDeleteAccount(context.Background(), authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ConfirmDeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmDeleteAccount`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ConfirmDeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authHash** | **string** | auth hash from email confirmation | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmDeleteAccountRequest struct via the builder pattern


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


## ConfirmRegistration

> ApiResponseV1 ConfirmRegistration(ctx, authHash).Execute()

confirm a new member registration

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
    authHash := "authHash_example" // string | auth hash from email confirmation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ConfirmRegistration(context.Background(), authHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ConfirmRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmRegistration`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ConfirmRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authHash** | **string** | auth hash from email confirmation | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmRegistrationRequest struct via the builder pattern


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


## DeleteAccount

> ApiResponseV1 DeleteAccount(ctx, captchaResponse).Execute()

delete your account

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
    captchaResponse := "captchaResponse_example" // string | response of captcha

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.DeleteAccount(context.Background(), captchaResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccount`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captchaResponse** | **string** | response of captcha | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## ForgotPassword

> ApiResponseV1 ForgotPassword(ctx, captchaResponse).AccountForgotPassModelV1(accountForgotPassModelV1).Execute()

send a forgotten password email

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
    captchaResponse := "captchaResponse_example" // string | response of captcha
    accountForgotPassModelV1 := *openapiclient.NewAccountForgotPassModelV1() // AccountForgotPassModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ForgotPassword(context.Background(), captchaResponse).AccountForgotPassModelV1(accountForgotPassModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ForgotPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForgotPassword`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ForgotPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captchaResponse** | **string** | response of captcha | 

### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountForgotPassModelV1** | [**AccountForgotPassModelV1**](AccountForgotPassModelV1.md) |  | 

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


## Login

> ApiResponseV1 Login(ctx).AccountLoginModelV1(accountLoginModelV1).Execute()

create a session token

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
    accountLoginModelV1 := *openapiclient.NewAccountLoginModelV1() // AccountLoginModelV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.Login(context.Background()).AccountLoginModelV1(accountLoginModelV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountLoginModelV1** | [**AccountLoginModelV1**](AccountLoginModelV1.md) |  | 

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


## Logout

> ApiResponseV1 Logout(ctx).Execute()

remove a session token

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
    resp, r, err := apiClient.AccountAPI.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logout`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Logout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## Profile

> UserModelV1 Profile(ctx).Execute()

get the profile for the current user

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
    resp, r, err := apiClient.AccountAPI.Profile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Profile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Profile`: UserModelV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Profile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProfileRequest struct via the builder pattern


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


## RegisterMember

> ApiResponseV1 RegisterMember(ctx, captchaResponse).UserModelRegisterV1(userModelRegisterV1).Execute()

register a new member

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
    captchaResponse := "captchaResponse_example" // string | response of captcha
    userModelRegisterV1 := *openapiclient.NewUserModelRegisterV1() // UserModelRegisterV1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.RegisterMember(context.Background(), captchaResponse).UserModelRegisterV1(userModelRegisterV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RegisterMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterMember`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RegisterMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captchaResponse** | **string** | response of captcha | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userModelRegisterV1** | [**UserModelRegisterV1**](UserModelRegisterV1.md) |  | 

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


## ResendAuthEmail

> ApiResponseV1 ResendAuthEmail(ctx, id).Execute()

send an auth email to a user

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
    resp, r, err := apiClient.AccountAPI.ResendAuthEmail(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ResendAuthEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendAuthEmail`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ResendAuthEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendAuthEmailRequest struct via the builder pattern


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


## SendForgotEmail

> ApiResponseV1 SendForgotEmail(ctx, id).Execute()

send a forgotten password email to a user

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
    resp, r, err := apiClient.AccountAPI.SendForgotEmail(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.SendForgotEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendForgotEmail`: ApiResponseV1
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.SendForgotEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Member id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendForgotEmailRequest struct via the builder pattern


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

