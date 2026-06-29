# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserDelegate**](UsersAPI.md#AddUserDelegate) | **Post** /api/v2/users/{userId}/actions/add-delegate/{delegateId} | Add a delegate to a user
[**AddUserSshKey**](UsersAPI.md#AddUserSshKey) | **Post** /api/v2/users/{userId}/ssh-keys | Add SSH key for user
[**ArchiveUser**](UsersAPI.md#ArchiveUser) | **Post** /api/v2/users/{userId}/actions/archive | Archive user
[**ChangeUserAccount**](UsersAPI.md#ChangeUserAccount) | **Post** /api/v2/users/{userId}/actions/change-account | Change account for user
[**CreateUserAuthorized**](UsersAPI.md#CreateUserAuthorized) | **Post** /api/v2/users | Creates another user
[**DeleteUserSshKey**](UsersAPI.md#DeleteUserSshKey) | **Delete** /api/v2/users/{userId}/ssh-keys/{keyId} | Delete SSH key for user
[**GetQuotaLimitsBreakdown**](UsersAPI.md#GetQuotaLimitsBreakdown) | **Get** /api/v2/users/{userId}/quota-limits-breakdown | Get quota limits breakdown for a user
[**GetUser**](UsersAPI.md#GetUser) | **Get** /api/v2/users/{userId} | Get user
[**GetUserChildDelegates**](UsersAPI.md#GetUserChildDelegates) | **Get** /api/v2/users/{userId}/child-delegates | Get user child delegates by ID
[**GetUserConfiguration**](UsersAPI.md#GetUserConfiguration) | **Get** /api/v2/users/{userId}/config | Get user configuration by ID
[**GetUserParentDelegates**](UsersAPI.md#GetUserParentDelegates) | **Get** /api/v2/users/{userId}/parent-delegates | Get user parent delegates by ID
[**GetUserSshKey**](UsersAPI.md#GetUserSshKey) | **Get** /api/v2/users/{userId}/ssh-keys/{keyId} | Get specific SSH key
[**GetUserSshKeys**](UsersAPI.md#GetUserSshKeys) | **Get** /api/v2/users/{userId}/ssh-keys | Get user SSH keys
[**GetUserSuspendReasons**](UsersAPI.md#GetUserSuspendReasons) | **Get** /api/v2/users/{userId}/suspend-reasons | Get user suspend reasons by ID
[**GetUsers**](UsersAPI.md#GetUsers) | **Get** /api/v2/users | Get users
[**RemoveUserDelegate**](UsersAPI.md#RemoveUserDelegate) | **Post** /api/v2/users/{userId}/actions/remove-delegate/{delegateId} | Remove a delegate from a user
[**ResendEmailVerification**](UsersAPI.md#ResendEmailVerification) | **Post** /api/v2/users/{userId}/actions/resend-email-verification | Resend email verification
[**ResendUserInvitation**](UsersAPI.md#ResendUserInvitation) | **Post** /api/v2/users/{userId}/actions/resend-user-invitation | Resend user invitation
[**SendPasswordResetByAdmin**](UsersAPI.md#SendPasswordResetByAdmin) | **Post** /api/v2/users/{userId}/actions/send-password-reset | Send password reset by admin
[**SetUserPasswordByAdmin**](UsersAPI.md#SetUserPasswordByAdmin) | **Post** /api/v2/users/{userId}/actions/set-password | Set user password by admin
[**SuspendUser**](UsersAPI.md#SuspendUser) | **Post** /api/v2/users/{userId}/actions/suspend | Suspend a user
[**UnarchiveUser**](UsersAPI.md#UnarchiveUser) | **Post** /api/v2/users/{userId}/actions/unarchive | Unarchive user
[**UnsuspendUser**](UsersAPI.md#UnsuspendUser) | **Post** /api/v2/users/{userId}/actions/unsuspend | Unsuspend a user
[**UpdateUserConfig**](UsersAPI.md#UpdateUserConfig) | **Patch** /api/v2/users/{userId}/config | Update user configuration
[**UpdateUserMeta**](UsersAPI.md#UpdateUserMeta) | **Patch** /api/v2/users/{userId}/meta | Update user metadata



## AddUserDelegate

> User AddUserDelegate(ctx, userId, delegateId).Execute()

Add a delegate to a user

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
	userId := int64(789) // int64 | ID of the user
	delegateId := int64(789) // int64 | ID of the delegate to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AddUserDelegate(context.Background(), userId, delegateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddUserDelegate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserDelegate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddUserDelegate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | ID of the user | 
**delegateId** | **int64** | ID of the delegate to add | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserDelegateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AddUserSshKey

> UserSSHKeys AddUserSshKey(ctx, userId).CreateUserSSHKey(createUserSSHKey).Execute()

Add SSH key for user

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
	userId := int64(789) // int64 | 
	createUserSSHKey := *openapiclient.NewCreateUserSSHKey("SshKey_example") // CreateUserSSHKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AddUserSshKey(context.Background(), userId).CreateUserSSHKey(createUserSSHKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddUserSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserSshKey`: UserSSHKeys
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddUserSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserSSHKey** | [**CreateUserSSHKey**](CreateUserSSHKey.md) |  | 

### Return type

[**UserSSHKeys**](UserSSHKeys.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveUser

> User ArchiveUser(ctx, userId).IfMatch(ifMatch).Execute()

Archive user



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
	userId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ArchiveUser(context.Background(), userId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ArchiveUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ArchiveUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveUserRequest struct via the builder pattern


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


## ChangeUserAccount

> User ChangeUserAccount(ctx, userId).ChangeUserAccount(changeUserAccount).IfMatch(ifMatch).Execute()

Change account for user



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
	userId := int64(789) // int64 | 
	changeUserAccount := *openapiclient.NewChangeUserAccount(int64(123)) // ChangeUserAccount | The new account id
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ChangeUserAccount(context.Background(), userId).ChangeUserAccount(changeUserAccount).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ChangeUserAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeUserAccount`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ChangeUserAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeUserAccount** | [**ChangeUserAccount**](ChangeUserAccount.md) | The new account id | 
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


## CreateUserAuthorized

> User CreateUserAuthorized(ctx).CreateUser(createUser).Execute()

Creates another user



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
	createUser := *openapiclient.NewCreateUser("DisplayName_example", "AccessLevel_example", "Email_example") // CreateUser | The user to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateUserAuthorized(context.Background()).CreateUser(createUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserAuthorized``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserAuthorized`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserAuthorized`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserAuthorizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUser** | [**CreateUser**](CreateUser.md) | The user to create | 

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


## DeleteUserSshKey

> DeleteUserSshKey(ctx, userId, keyId).Execute()

Delete SSH key for user

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
	userId := int64(789) // int64 | 
	keyId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DeleteUserSshKey(context.Background(), userId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 
**keyId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetQuotaLimitsBreakdown

> QuotaLimitsBreakdown GetQuotaLimitsBreakdown(ctx, userId).IncludeUsage(includeUsage).Execute()

Get quota limits breakdown for a user



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
	userId := int64(789) // int64 | 
	includeUsage := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetQuotaLimitsBreakdown(context.Background(), userId).IncludeUsage(includeUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetQuotaLimitsBreakdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotaLimitsBreakdown`: QuotaLimitsBreakdown
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetQuotaLimitsBreakdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotaLimitsBreakdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUsage** | **bool** |  | 

### Return type

[**QuotaLimitsBreakdown**](QuotaLimitsBreakdown.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Execute()

Get user



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
	userId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetUserChildDelegates

> UserList GetUserChildDelegates(ctx, userId).Execute()

Get user child delegates by ID

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
	userId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserChildDelegates(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserChildDelegates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserChildDelegates`: UserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserChildDelegates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserChildDelegatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserList**](UserList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserConfiguration

> UserConfiguration GetUserConfiguration(ctx, userId).Execute()

Get user configuration by ID

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
	userId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserConfiguration(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserConfiguration`: UserConfiguration
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserConfiguration**](UserConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserParentDelegates

> UserList GetUserParentDelegates(ctx, userId).Execute()

Get user parent delegates by ID

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
	userId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserParentDelegates(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserParentDelegates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserParentDelegates`: UserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserParentDelegates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserParentDelegatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserList**](UserList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSshKey

> UserSSHKeys GetUserSshKey(ctx, userId, keyId).Execute()

Get specific SSH key

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
	userId := int64(789) // int64 | 
	keyId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserSshKey(context.Background(), userId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSshKey`: UserSSHKeys
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 
**keyId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserSSHKeys**](UserSSHKeys.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSshKeys

> UserSSHKeysList GetUserSshKeys(ctx, userId).Execute()

Get user SSH keys

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
	userId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserSshKeys(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSshKeys`: UserSSHKeysList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSshKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSshKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSSHKeysList**](UserSSHKeysList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSuspendReasons

> UserSuspendReasonList GetUserSuspendReasons(ctx, userId).Execute()

Get user suspend reasons by ID

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
	userId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserSuspendReasons(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSuspendReasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSuspendReasons`: UserSuspendReasonList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSuspendReasons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSuspendReasonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSuspendReasonList**](UserSuspendReasonList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> UserPaginatedList GetUsers(ctx).Page(page).Limit(limit).FilterId(filterId).FilterDisplayName(filterDisplayName).FilterEmail(filterEmail).FilterAccountId(filterAccountId).FilterArchived(filterArchived).FilterInfrastructureIdDefault(filterInfrastructureIdDefault).FilterAccessLevel(filterAccessLevel).FilterIsBillable(filterIsBillable).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get users



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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterDisplayName := []string{"Inner_example"} // []string | Filter by displayName query param.  **Format:** filter.displayName={$not}:OPERATION:VALUE    **Example:** filter.displayName=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterEmail := []string{"Inner_example"} // []string | Filter by email query param.  **Format:** filter.email={$not}:OPERATION:VALUE    **Example:** filter.email=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterAccountId := []string{"Inner_example"} // []string | Filter by accountId query param.  **Format:** filter.accountId={$not}:OPERATION:VALUE    **Example:** filter.accountId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterArchived := []string{"Inner_example"} // []string | Filter by archived query param.  **Format:** filter.archived={$not}:OPERATION:VALUE    **Example:** filter.archived=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterInfrastructureIdDefault := []string{"Inner_example"} // []string | Filter by infrastructureIdDefault query param.  **Format:** filter.infrastructureIdDefault={$not}:OPERATION:VALUE    **Example:** filter.infrastructureIdDefault=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterAccessLevel := []string{"Inner_example"} // []string | Filter by accessLevel query param.  **Format:** filter.accessLevel={$not}:OPERATION:VALUE    **Example:** filter.accessLevel=$eq:John Doe&filter.accessLevel=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterIsBillable := []string{"Inner_example"} // []string | Filter by isBillable query param.  **Format:** filter.isBillable={$not}:OPERATION:VALUE    **Example:** filter.isBillable=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=displayName:DESC   **Default Value:** displayName:DESC  **Available Fields** - id  - displayName  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** displayName,email   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - displayName  - email  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsers(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterDisplayName(filterDisplayName).FilterEmail(filterEmail).FilterAccountId(filterAccountId).FilterArchived(filterArchived).FilterInfrastructureIdDefault(filterInfrastructureIdDefault).FilterAccessLevel(filterAccessLevel).FilterIsBillable(filterIsBillable).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: UserPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterDisplayName** | **[]string** | Filter by displayName query param.  **Format:** filter.displayName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.displayName&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterEmail** | **[]string** | Filter by email query param.  **Format:** filter.email&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.email&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterAccountId** | **[]string** | Filter by accountId query param.  **Format:** filter.accountId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.accountId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterArchived** | **[]string** | Filter by archived query param.  **Format:** filter.archived&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.archived&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterInfrastructureIdDefault** | **[]string** | Filter by infrastructureIdDefault query param.  **Format:** filter.infrastructureIdDefault&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureIdDefault&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterAccessLevel** | **[]string** | Filter by accessLevel query param.  **Format:** filter.accessLevel&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.accessLevel&#x3D;$eq:John Doe&amp;filter.accessLevel&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterIsBillable** | **[]string** | Filter by isBillable query param.  **Format:** filter.isBillable&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.isBillable&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;displayName:DESC   **Default Value:** displayName:DESC  **Available Fields** - id  - displayName  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** displayName,email   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - displayName  - email  | 

### Return type

[**UserPaginatedList**](UserPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserDelegate

> User RemoveUserDelegate(ctx, userId, delegateId).Execute()

Remove a delegate from a user

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
	userId := int64(789) // int64 | ID of the user
	delegateId := int64(789) // int64 | ID of the delegate to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.RemoveUserDelegate(context.Background(), userId, delegateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveUserDelegate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserDelegate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveUserDelegate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | ID of the user | 
**delegateId** | **int64** | ID of the delegate to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserDelegateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ResendEmailVerification

> User ResendEmailVerification(ctx, userId).ResendUserVerificationEmail(resendUserVerificationEmail).Execute()

Resend email verification

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
	userId := int64(789) // int64 | 
	resendUserVerificationEmail := *openapiclient.NewResendUserVerificationEmail() // ResendUserVerificationEmail | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ResendEmailVerification(context.Background(), userId).ResendUserVerificationEmail(resendUserVerificationEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ResendEmailVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendEmailVerification`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ResendEmailVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resendUserVerificationEmail** | [**ResendUserVerificationEmail**](ResendUserVerificationEmail.md) |  | 

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


## ResendUserInvitation

> User ResendUserInvitation(ctx, userId).ResendUserInvitation(resendUserInvitation).Execute()

Resend user invitation

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
	userId := int64(789) // int64 | 
	resendUserInvitation := *openapiclient.NewResendUserInvitation() // ResendUserInvitation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ResendUserInvitation(context.Background(), userId).ResendUserInvitation(resendUserInvitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ResendUserInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendUserInvitation`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ResendUserInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendUserInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resendUserInvitation** | [**ResendUserInvitation**](ResendUserInvitation.md) |  | 

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


## SendPasswordResetByAdmin

> User SendPasswordResetByAdmin(ctx, userId).PasswordResetByAdmin(passwordResetByAdmin).Execute()

Send password reset by admin

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
	userId := int64(789) // int64 | 
	passwordResetByAdmin := *openapiclient.NewPasswordResetByAdmin() // PasswordResetByAdmin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SendPasswordResetByAdmin(context.Background(), userId).PasswordResetByAdmin(passwordResetByAdmin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SendPasswordResetByAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPasswordResetByAdmin`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SendPasswordResetByAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendPasswordResetByAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordResetByAdmin** | [**PasswordResetByAdmin**](PasswordResetByAdmin.md) |  | 

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


## SetUserPasswordByAdmin

> User SetUserPasswordByAdmin(ctx, userId).SetUserPasswordByAdmin(setUserPasswordByAdmin).IfMatch(ifMatch).Execute()

Set user password by admin

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
	userId := int64(789) // int64 | 
	setUserPasswordByAdmin := *openapiclient.NewSetUserPasswordByAdmin("Password_example") // SetUserPasswordByAdmin | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SetUserPasswordByAdmin(context.Background(), userId).SetUserPasswordByAdmin(setUserPasswordByAdmin).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SetUserPasswordByAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserPasswordByAdmin`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SetUserPasswordByAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserPasswordByAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setUserPasswordByAdmin** | [**SetUserPasswordByAdmin**](SetUserPasswordByAdmin.md) |  | 
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


## SuspendUser

> UserSuspendReason SuspendUser(ctx, userId).UserSuspend(userSuspend).IfMatch(ifMatch).Execute()

Suspend a user

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
	userId := int64(789) // int64 | 
	userSuspend := *openapiclient.NewUserSuspend("SuspendReason_example", "SuspendReasonPublicComment_example") // UserSuspend | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SuspendUser(context.Background(), userId).UserSuspend(userSuspend).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SuspendUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuspendUser`: UserSuspendReason
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SuspendUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSuspend** | [**UserSuspend**](UserSuspend.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**UserSuspendReason**](UserSuspendReason.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveUser

> User UnarchiveUser(ctx, userId).IfMatch(ifMatch).Execute()

Unarchive user



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
	userId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UnarchiveUser(context.Background(), userId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnarchiveUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnarchiveUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UnarchiveUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveUserRequest struct via the builder pattern


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


## UnsuspendUser

> UnsuspendUser(ctx, userId).IfMatch(ifMatch).Execute()

Unsuspend a user

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
	userId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UnsuspendUser(context.Background(), userId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnsuspendUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsuspendUserRequest struct via the builder pattern


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


## UpdateUserConfig

> UserConfiguration UpdateUserConfig(ctx, userId).UpdateUser(updateUser).IfMatch(ifMatch).Execute()

Update user configuration



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
	userId := int64(789) // int64 | 
	updateUser := *openapiclient.NewUpdateUser() // UpdateUser | The user updates
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserConfig(context.Background(), userId).UpdateUser(updateUser).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserConfig`: UserConfiguration
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUser** | [**UpdateUser**](UpdateUser.md) | The user updates | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**UserConfiguration**](UserConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserMeta

> UserMeta UpdateUserMeta(ctx, userId).UserMeta(userMeta).Execute()

Update user metadata

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
	userId := int64(789) // int64 | 
	userMeta := *openapiclient.NewUserMeta() // UserMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserMeta(context.Background(), userId).UserMeta(userMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserMeta`: UserMeta
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userMeta** | [**UserMeta**](UserMeta.md) |  | 

### Return type

[**UserMeta**](UserMeta.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

