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
[**GetUser**](UsersAPI.md#GetUser) | **Get** /api/v2/users/{userId} | Get user
[**GetUserChildDelegates**](UsersAPI.md#GetUserChildDelegates) | **Get** /api/v2/users/{userId}/child-delegates | Get user child delegates by ID
[**GetUserConfiguration**](UsersAPI.md#GetUserConfiguration) | **Get** /api/v2/users/{userId}/config | Get user configuration by ID
[**GetUserLimits**](UsersAPI.md#GetUserLimits) | **Get** /api/v2/users/{userId}/limits | Get user limits
[**GetUserParentDelegates**](UsersAPI.md#GetUserParentDelegates) | **Get** /api/v2/users/{userId}/parent-delegates | Get user parent delegates by ID
[**GetUserPermissions**](UsersAPI.md#GetUserPermissions) | **Get** /api/v2/users/{userId}/permissions | Get user resource permissions by ID
[**GetUserSshKey**](UsersAPI.md#GetUserSshKey) | **Get** /api/v2/users/{userId}/ssh-keys/{keyId} | Get specific SSH key
[**GetUserSshKeys**](UsersAPI.md#GetUserSshKeys) | **Get** /api/v2/users/{userId}/ssh-keys | Get user SSH keys
[**GetUserSuspendReasons**](UsersAPI.md#GetUserSuspendReasons) | **Get** /api/v2/users/{userId}/suspend-reasons | Get user suspend reasons by ID
[**GetUsers**](UsersAPI.md#GetUsers) | **Get** /api/v2/users | Get users
[**RemoveUserDelegate**](UsersAPI.md#RemoveUserDelegate) | **Post** /api/v2/users/{userId}/actions/remove-delegate/{delegateId} | Remove a delegate from a user
[**SuspendUser**](UsersAPI.md#SuspendUser) | **Post** /api/v2/users/{userId}/actions/suspend | Suspend a user
[**UnarchiveUser**](UsersAPI.md#UnarchiveUser) | **Post** /api/v2/users/{userId}/actions/unarchive | Unarchive user
[**UnsuspendUser**](UsersAPI.md#UnsuspendUser) | **Post** /api/v2/users/{userId}/actions/unsuspend | Unsuspend a user
[**UpdateUserConfig**](UsersAPI.md#UpdateUserConfig) | **Patch** /api/v2/users/{userId}/config | Update user configuration
[**UpdateUserLimits**](UsersAPI.md#UpdateUserLimits) | **Patch** /api/v2/users/{userId}/actions/change-limits | Update user limits
[**UpdateUserMeta**](UsersAPI.md#UpdateUserMeta) | **Patch** /api/v2/users/{userId}/meta | Update user metadata
[**UpdateUserPermissions**](UsersAPI.md#UpdateUserPermissions) | **Post** /api/v2/users/{userId}/actions/change-resource-permissions | Update user resource permissions



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
	userId := float32(8.14) // float32 | ID of the user
	delegateId := float32(8.14) // float32 | ID of the delegate to add

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
**userId** | **float32** | ID of the user | 
**delegateId** | **float32** | ID of the delegate to add | 

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
	userId := float32(8.14) // float32 | 
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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 
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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 
	changeUserAccount := *openapiclient.NewChangeUserAccount(float32(123)) // ChangeUserAccount | The new account id
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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 
	keyId := float32(8.14) // float32 | 

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
**userId** | **float32** |  | 
**keyId** | **float32** |  | 

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


## GetUser

> User GetUser(ctx, userId).Recursion(recursion).Execute()

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
	userId := float32(8.14) // float32 | 
	recursion := float32(8.14) // float32 | The recursion level of the displayed details. Default is 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), userId).Recursion(recursion).Execute()
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
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursion** | **float32** | The recursion level of the displayed details. Default is 0. | 

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
	userId := float32(8.14) // float32 | 

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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 

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
**userId** | **float32** |  | 

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


## GetUserLimits

> UserLimits GetUserLimits(ctx, userId).Execute()

Get user limits



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
	resp, r, err := apiClient.UsersAPI.GetUserLimits(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserLimits`: UserLimits
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserLimits**](UserLimits.md)

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
	userId := float32(8.14) // float32 | 

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
**userId** | **float32** |  | 

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


## GetUserPermissions

> UserPermissions GetUserPermissions(ctx, userId).Execute()

Get user resource permissions by ID

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
	resp, r, err := apiClient.UsersAPI.GetUserPermissions(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPermissions`: UserPermissions
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserPermissions**](UserPermissions.md)

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
	userId := float32(8.14) // float32 | 
	keyId := float32(8.14) // float32 | 

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
**userId** | **float32** |  | 
**keyId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 

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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 

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
**userId** | **float32** |  | 

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

> UserPaginatedList GetUsers(ctx).Page(page).Limit(limit).FilterId(filterId).FilterDisplayName(filterDisplayName).FilterEmail(filterEmail).FilterAccountId(filterAccountId).FilterArchived(filterArchived).FilterInfrastructureIdDefault(filterInfrastructureIdDefault).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterDisplayName := []string{"Inner_example"} // []string | Filter by displayName query param.           <p>              <b>Format: </b> filter.displayName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.displayName=$not:$like:John Doe&filter.displayName=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterEmail := []string{"Inner_example"} // []string | Filter by email query param.           <p>              <b>Format: </b> filter.email={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.email=$not:$like:John Doe&filter.email=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAccountId := []string{"Inner_example"} // []string | Filter by accountId query param.           <p>              <b>Format: </b> filter.accountId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.accountId=$not:$like:John Doe&filter.accountId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterArchived := []string{"Inner_example"} // []string | Filter by archived query param.           <p>              <b>Format: </b> filter.archived={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.archived=$not:$like:John Doe&filter.archived=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterInfrastructureIdDefault := []string{"Inner_example"} // []string | Filter by infrastructureIdDefault query param.           <p>              <b>Format: </b> filter.infrastructureIdDefault={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureIdDefault=$not:$like:John Doe&filter.infrastructureIdDefault=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> displayName:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>displayName</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> displayName,email           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>displayName</li> <li>email</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsers(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterDisplayName(filterDisplayName).FilterEmail(filterEmail).FilterAccountId(filterAccountId).FilterArchived(filterArchived).FilterInfrastructureIdDefault(filterInfrastructureIdDefault).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterDisplayName** | **[]string** | Filter by displayName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.displayName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.displayName&#x3D;$not:$like:John Doe&amp;filter.displayName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterEmail** | **[]string** | Filter by email query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.email&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.email&#x3D;$not:$like:John Doe&amp;filter.email&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterAccountId** | **[]string** | Filter by accountId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.accountId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.accountId&#x3D;$not:$like:John Doe&amp;filter.accountId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterArchived** | **[]string** | Filter by archived query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.archived&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.archived&#x3D;$not:$like:John Doe&amp;filter.archived&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterInfrastructureIdDefault** | **[]string** | Filter by infrastructureIdDefault query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureIdDefault&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureIdDefault&#x3D;$not:$like:John Doe&amp;filter.infrastructureIdDefault&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; displayName:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;displayName&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; displayName,email           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;displayName&lt;/li&gt; &lt;li&gt;email&lt;/li&gt;&lt;/ul&gt;          | 

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
	userId := float32(8.14) // float32 | ID of the user
	delegateId := float32(8.14) // float32 | ID of the delegate to remove

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
**userId** | **float32** | ID of the user | 
**delegateId** | **float32** | ID of the delegate to remove | 

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
	userId := float32(8.14) // float32 | 
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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 
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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 
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
**userId** | **float32** |  | 

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
	userId := float32(8.14) // float32 | 
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
**userId** | **float32** |  | 

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


## UpdateUserLimits

> UserLimits UpdateUserLimits(ctx, userId).UserLimits(userLimits).IfMatch(ifMatch).Execute()

Update user limits



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
	userLimits := *openapiclient.NewUserLimits() // UserLimits | The new limits
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserLimits(context.Background(), userId).UserLimits(userLimits).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserLimits`: UserLimits
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userLimits** | [**UserLimits**](UserLimits.md) | The new limits | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**UserLimits**](UserLimits.md)

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
	userId := float32(8.14) // float32 | 
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
**userId** | **float32** |  | 

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


## UpdateUserPermissions

> UserPermissions UpdateUserPermissions(ctx, userId).UpdateUserPermissions(updateUserPermissions).IfMatch(ifMatch).Execute()

Update user resource permissions

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
	updateUserPermissions := *openapiclient.NewUpdateUserPermissions() // UpdateUserPermissions | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserPermissions(context.Background(), userId).UpdateUserPermissions(updateUserPermissions).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserPermissions`: UserPermissions
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserPermissions** | [**UpdateUserPermissions**](UpdateUserPermissions.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**UserPermissions**](UserPermissions.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

