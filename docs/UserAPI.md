# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveUser**](UserAPI.md#ArchiveUser) | **Post** /api/v2/users/{userId}/actions/archive | Archive user
[**ChangeUserAccount**](UserAPI.md#ChangeUserAccount) | **Post** /api/v2/users/{userId}/actions/change-account | Change account for user
[**CreateUser**](UserAPI.md#CreateUser) | **Post** /api/v2/users | Creates a user
[**GetUser**](UserAPI.md#GetUser) | **Get** /api/v2/users/{userId} | Get user
[**GetUserLimits**](UserAPI.md#GetUserLimits) | **Get** /api/v2/users/{userId}/limits | Get user limits
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /api/v2/users | Get users
[**UnarchiveUser**](UserAPI.md#UnarchiveUser) | **Post** /api/v2/users/{userId}/actions/unarchive | Unarchive user
[**UpdateUser**](UserAPI.md#UpdateUser) | **Patch** /api/v2/users/{userId} | Update user
[**UpdateUserLimits**](UserAPI.md#UpdateUserLimits) | **Patch** /api/v2/users/{userId}/limits | Update user limits



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
	resp, r, err := apiClient.UserAPI.ArchiveUser(context.Background(), userId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ArchiveUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ArchiveUser`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.ChangeUserAccount(context.Background(), userId).ChangeUserAccount(changeUserAccount).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ChangeUserAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeUserAccount`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ChangeUserAccount`: %v\n", resp)
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


## CreateUser

> User CreateUser(ctx).CreateUser(createUser).Execute()

Creates a user



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
	createUser := *openapiclient.NewCreateUser("DisplayName_example", "Email_example") // CreateUser | The user to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUser(context.Background()).CreateUser(createUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


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
	resp, r, err := apiClient.UserAPI.GetUser(context.Background(), userId).Recursion(recursion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.GetUserLimits(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserLimits`: UserLimits
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserLimits`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterDisplayName(filterDisplayName).FilterEmail(filterEmail).FilterAccountId(filterAccountId).FilterArchived(filterArchived).FilterInfrastructureIdDefault(filterInfrastructureIdDefault).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: UserPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.UnarchiveUser(context.Background(), userId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnarchiveUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnarchiveUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UnarchiveUser`: %v\n", resp)
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


## UpdateUser

> User UpdateUser(ctx, userId).UpdateUser(updateUser).IfMatch(ifMatch).Execute()

Update user



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
	resp, r, err := apiClient.UserAPI.UpdateUser(context.Background(), userId).UpdateUser(updateUser).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUser** | [**UpdateUser**](UpdateUser.md) | The user updates | 
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
	resp, r, err := apiClient.UserAPI.UpdateUserLimits(context.Background(), userId).UserLimits(userLimits).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserLimits`: UserLimits
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUserLimits`: %v\n", resp)
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

