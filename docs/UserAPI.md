# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableUser2FA**](UserAPI.md#DisableUser2FA) | **Post** /api/v2/user/actions/2fa-disable | Disable 2FA
[**EnableUser2FA**](UserAPI.md#EnableUser2FA) | **Post** /api/v2/user/actions/2fa-enable | Enable 2FA
[**GenerateUser2FASecret**](UserAPI.md#GenerateUser2FASecret) | **Post** /api/v2/user/actions/2fa-generate | Generate 2FA secret
[**GetUserApiKey**](UserAPI.md#GetUserApiKey) | **Get** /api/v2/user/api-key | Get user API key
[**InitiateEmailChange**](UserAPI.md#InitiateEmailChange) | **Post** /api/v2/user/actions/initiate-email-change | Initiates the user email change
[**InitiatePasswordReset**](UserAPI.md#InitiatePasswordReset) | **Post** /api/v2/user/actions/initiate-password-reset | Initiate reset user password
[**RegenerateUserApiKey**](UserAPI.md#RegenerateUserApiKey) | **Post** /api/v2/user/actions/regenerate-api-key | Regenerate user API key
[**RegenerateUserJwtSalt**](UserAPI.md#RegenerateUserJwtSalt) | **Post** /api/v2/user/actions/regenerate-jwt-salt | Regenerate user JWT salt. Also logs out all user sessions.
[**UpdateUserPassword**](UserAPI.md#UpdateUserPassword) | **Post** /api/v2/user/actions/change-password | Change user password
[**UserControllerHandleEmailVerify**](UserAPI.md#UserControllerHandleEmailVerify) | **Get** /api/v2/user/actions/verify-email | Handle user email verify action
[**UserControllerHandleUserResetPassword**](UserAPI.md#UserControllerHandleUserResetPassword) | **Get** /api/v2/user/actions/reset-password | Handle user reset password action



## DisableUser2FA

> DisableUser2FA(ctx).IfMatch(ifMatch).Execute()

Disable 2FA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DisableUser2FA(context.Background()).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DisableUser2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableUser2FARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifMatch** | **string** | Entity tag | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableUser2FA

> EnableUser2FA(ctx).TwoFactorAuthenticationToken(twoFactorAuthenticationToken).IfMatch(ifMatch).Execute()

Enable 2FA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	twoFactorAuthenticationToken := *openapiclient.NewTwoFactorAuthenticationToken("Token_example") // TwoFactorAuthenticationToken | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.EnableUser2FA(context.Background()).TwoFactorAuthenticationToken(twoFactorAuthenticationToken).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.EnableUser2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableUser2FARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthenticationToken** | [**TwoFactorAuthenticationToken**](TwoFactorAuthenticationToken.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUser2FASecret

> TwoFactorAuthenticationSecret GenerateUser2FASecret(ctx).Execute()

Generate 2FA secret

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GenerateUser2FASecret(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GenerateUser2FASecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUser2FASecret`: TwoFactorAuthenticationSecret
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GenerateUser2FASecret`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUser2FASecretRequest struct via the builder pattern


### Return type

[**TwoFactorAuthenticationSecret**](TwoFactorAuthenticationSecret.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApiKey

> UserApiKey GetUserApiKey(ctx, userId).Execute()

Get user API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserApiKey(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserApiKey`: UserApiKey
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateEmailChange

> User InitiateEmailChange(ctx).ChangeUserEmail(changeUserEmail).Execute()

Initiates the user email change

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	changeUserEmail := *openapiclient.NewChangeUserEmail("Email_example") // ChangeUserEmail | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.InitiateEmailChange(context.Background()).ChangeUserEmail(changeUserEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.InitiateEmailChange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateEmailChange`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.InitiateEmailChange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateEmailChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeUserEmail** | [**ChangeUserEmail**](ChangeUserEmail.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiatePasswordReset

> InitiatePasswordReset(ctx).PasswordReset(passwordReset).Execute()

Initiate reset user password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	passwordReset := *openapiclient.NewPasswordReset("Email_example") // PasswordReset | The password reset request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.InitiatePasswordReset(context.Background()).PasswordReset(passwordReset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.InitiatePasswordReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiatePasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordReset** | [**PasswordReset**](PasswordReset.md) | The password reset request | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateUserApiKey

> User RegenerateUserApiKey(ctx).IfMatch(ifMatch).Execute()

Regenerate user API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RegenerateUserApiKey(context.Background()).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RegenerateUserApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateUserApiKey`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RegenerateUserApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifMatch** | **string** | Entity tag | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateUserJwtSalt

> User RegenerateUserJwtSalt(ctx).IfMatch(ifMatch).Execute()

Regenerate user JWT salt. Also logs out all user sessions.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RegenerateUserJwtSalt(context.Background()).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RegenerateUserJwtSalt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateUserJwtSalt`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RegenerateUserJwtSalt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateUserJwtSaltRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifMatch** | **string** | Entity tag | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> User UpdateUserPassword(ctx).UserUpdatePassword(userUpdatePassword).IfMatch(ifMatch).Execute()

Change user password

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	userUpdatePassword := *openapiclient.NewUserUpdatePassword("NewPassword_example") // UserUpdatePassword | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateUserPassword(context.Background()).UserUpdatePassword(userUpdatePassword).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserPassword`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUserPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUpdatePassword** | [**UserUpdatePassword**](UserUpdatePassword.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserControllerHandleEmailVerify

> UserControllerHandleEmailVerify(ctx).Token(token).Execute()

Handle user email verify action

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	token := "token_example" // string | The encrypted request token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UserControllerHandleEmailVerify(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserControllerHandleEmailVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserControllerHandleEmailVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The encrypted request token. | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserControllerHandleUserResetPassword

> UserControllerHandleUserResetPassword(ctx).Token(token).Execute()

Handle user reset password action

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	token := "token_example" // string | The encrypted request token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UserControllerHandleUserResetPassword(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserControllerHandleUserResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserControllerHandleUserResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The encrypted request token. | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

