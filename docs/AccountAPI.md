# \AccountAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountControllerGetUserConfiguration**](AccountAPI.md#AccountControllerGetUserConfiguration) | **Get** /api/v2/accounts/{accountId}/config | Get account configuration by ID
[**ArchiveAccount**](AccountAPI.md#ArchiveAccount) | **Post** /api/v2/accounts/{accountId}/actions/archive | Archive account
[**CreateAccount**](AccountAPI.md#CreateAccount) | **Post** /api/v2/accounts | Create account
[**GetAccount**](AccountAPI.md#GetAccount) | **Get** /api/v2/accounts/{accountId} | Get account by id
[**GetAccountLimits**](AccountAPI.md#GetAccountLimits) | **Get** /api/v2/accounts/{accountId}/limits | Get account limits
[**GetAccountUsers**](AccountAPI.md#GetAccountUsers) | **Get** /api/v2/accounts/{accountId}/users | Get users for account
[**GetAccounts**](AccountAPI.md#GetAccounts) | **Get** /api/v2/accounts | Get all accounts
[**UnarchiveAccount**](AccountAPI.md#UnarchiveAccount) | **Post** /api/v2/accounts/{accountId}/actions/unarchive | Unarchive account
[**UpdateAccountConfig**](AccountAPI.md#UpdateAccountConfig) | **Patch** /api/v2/accounts/{accountId}/config | Update account configuration
[**UpdateAccountLimits**](AccountAPI.md#UpdateAccountLimits) | **Patch** /api/v2/accounts/{accountId}/actions/change-limits | Update account limits



## AccountControllerGetUserConfiguration

> AccountConfig AccountControllerGetUserConfiguration(ctx, accountId).Execute()

Get account configuration by ID

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
	accountId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountControllerGetUserConfiguration(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountControllerGetUserConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountControllerGetUserConfiguration`: AccountConfig
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountControllerGetUserConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountControllerGetUserConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountConfig**](AccountConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveAccount

> Account ArchiveAccount(ctx, accountId).IfMatch(ifMatch).Execute()

Archive account



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
	accountId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.ArchiveAccount(context.Background(), accountId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ArchiveAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ArchiveAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

[**Account**](Account.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> Account CreateAccount(ctx).CreateAccount(createAccount).Execute()

Create account



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
	createAccount := *openapiclient.NewCreateAccount("Name_example") // CreateAccount | The account to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CreateAccount(context.Background()).CreateAccount(createAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccount** | [**CreateAccount**](CreateAccount.md) | The account to create | 

### Return type

[**Account**](Account.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId).Recursion(recursion).Execute()

Get account by id



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
	accountId := float32(8.14) // float32 | 
	recursion := float32(8.14) // float32 | The recursion level of the displayed details. Default is 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccount(context.Background(), accountId).Recursion(recursion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursion** | **float32** | The recursion level of the displayed details. Default is 0. | 

### Return type

[**Account**](Account.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountLimits

> AccountLimits GetAccountLimits(ctx, accountId).Execute()

Get account limits



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
	accountId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountLimits(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountLimits`: AccountLimits
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountLimits**](AccountLimits.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountUsers

> UserPaginatedList GetAccountUsers(ctx, accountId).Page(page).Limit(limit).FilterId(filterId).FilterDisplayName(filterDisplayName).FilterEmail(filterEmail).FilterAccountId(filterAccountId).FilterArchived(filterArchived).FilterInfrastructureIdDefault(filterInfrastructureIdDefault).FilterAccessLevel(filterAccessLevel).FilterIsBillable(filterIsBillable).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get users for account



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
	accountId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterDisplayName := []string{"Inner_example"} // []string | Filter by displayName query param.           <p>              <b>Format: </b> filter.displayName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.displayName=$not:$like:John Doe&filter.displayName=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterEmail := []string{"Inner_example"} // []string | Filter by email query param.           <p>              <b>Format: </b> filter.email={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.email=$not:$like:John Doe&filter.email=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAccountId := []string{"Inner_example"} // []string | Filter by accountId query param.           <p>              <b>Format: </b> filter.accountId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.accountId=$not:$like:John Doe&filter.accountId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterArchived := []string{"Inner_example"} // []string | Filter by archived query param.           <p>              <b>Format: </b> filter.archived={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.archived=$not:$like:John Doe&filter.archived=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterInfrastructureIdDefault := []string{"Inner_example"} // []string | Filter by infrastructureIdDefault query param.           <p>              <b>Format: </b> filter.infrastructureIdDefault={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureIdDefault=$not:$like:John Doe&filter.infrastructureIdDefault=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAccessLevel := []string{"Inner_example"} // []string | Filter by accessLevel query param.           <p>              <b>Format: </b> filter.accessLevel={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.accessLevel=$not:$like:John Doe&filter.accessLevel=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterIsBillable := []string{"Inner_example"} // []string | Filter by isBillable query param.           <p>              <b>Format: </b> filter.isBillable={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.isBillable=$not:$like:John Doe&filter.isBillable=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> displayName:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>displayName</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> displayName,email           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>displayName</li> <li>email</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountUsers(context.Background(), accountId).Page(page).Limit(limit).FilterId(filterId).FilterDisplayName(filterDisplayName).FilterEmail(filterEmail).FilterAccountId(filterAccountId).FilterArchived(filterArchived).FilterInfrastructureIdDefault(filterInfrastructureIdDefault).FilterAccessLevel(filterAccessLevel).FilterIsBillable(filterIsBillable).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountUsers`: UserPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountUsersRequest struct via the builder pattern


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
 **filterAccessLevel** | **[]string** | Filter by accessLevel query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.accessLevel&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.accessLevel&#x3D;$not:$like:John Doe&amp;filter.accessLevel&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterIsBillable** | **[]string** | Filter by isBillable query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.isBillable&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.isBillable&#x3D;$not:$like:John Doe&amp;filter.isBillable&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
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


## GetAccounts

> AccountPaginatedList GetAccounts(ctx).Page(page).Limit(limit).FilterName(filterName).FilterParentAccountId(filterParentAccountId).FilterPrimaryContactId(filterPrimaryContactId).FilterSecondaryContactId(filterSecondaryContactId).FilterArchived(filterArchived).FilterCode(filterCode).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all accounts



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
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterParentAccountId := []string{"Inner_example"} // []string | Filter by parentAccountId query param.           <p>              <b>Format: </b> filter.parentAccountId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.parentAccountId=$not:$like:John Doe&filter.parentAccountId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterPrimaryContactId := []string{"Inner_example"} // []string | Filter by primaryContactId query param.           <p>              <b>Format: </b> filter.primaryContactId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.primaryContactId=$not:$like:John Doe&filter.primaryContactId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSecondaryContactId := []string{"Inner_example"} // []string | Filter by secondaryContactId query param.           <p>              <b>Format: </b> filter.secondaryContactId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.secondaryContactId=$not:$like:John Doe&filter.secondaryContactId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterArchived := []string{"Inner_example"} // []string | Filter by archived query param.           <p>              <b>Format: </b> filter.archived={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.archived=$not:$like:John Doe&filter.archived=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterCode := []string{"Inner_example"} // []string | Filter by code query param.           <p>              <b>Format: </b> filter.code={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.code=$not:$like:John Doe&filter.code=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> name:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>name</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> name,code           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>name</li> <li>code</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccounts(context.Background()).Page(page).Limit(limit).FilterName(filterName).FilterParentAccountId(filterParentAccountId).FilterPrimaryContactId(filterPrimaryContactId).FilterSecondaryContactId(filterSecondaryContactId).FilterArchived(filterArchived).FilterCode(filterCode).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: AccountPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterParentAccountId** | **[]string** | Filter by parentAccountId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.parentAccountId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.parentAccountId&#x3D;$not:$like:John Doe&amp;filter.parentAccountId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterPrimaryContactId** | **[]string** | Filter by primaryContactId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.primaryContactId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.primaryContactId&#x3D;$not:$like:John Doe&amp;filter.primaryContactId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSecondaryContactId** | **[]string** | Filter by secondaryContactId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.secondaryContactId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.secondaryContactId&#x3D;$not:$like:John Doe&amp;filter.secondaryContactId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterArchived** | **[]string** | Filter by archived query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.archived&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.archived&#x3D;$not:$like:John Doe&amp;filter.archived&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterCode** | **[]string** | Filter by code query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.code&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.code&#x3D;$not:$like:John Doe&amp;filter.code&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; name:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name,code           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt; &lt;li&gt;code&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**AccountPaginatedList**](AccountPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveAccount

> Account UnarchiveAccount(ctx, accountId).IfMatch(ifMatch).Execute()

Unarchive account



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
	accountId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UnarchiveAccount(context.Background(), accountId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UnarchiveAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnarchiveAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UnarchiveAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

[**Account**](Account.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountConfig

> AccountConfig UpdateAccountConfig(ctx, accountId).UpdateAccount(updateAccount).IfMatch(ifMatch).Execute()

Update account configuration



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
	accountId := float32(8.14) // float32 | 
	updateAccount := *openapiclient.NewUpdateAccount("Name_example") // UpdateAccount | The account updates
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountConfig(context.Background(), accountId).UpdateAccount(updateAccount).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountConfig`: AccountConfig
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccount** | [**UpdateAccount**](UpdateAccount.md) | The account updates | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**AccountConfig**](AccountConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountLimits

> AccountLimits UpdateAccountLimits(ctx, accountId).AccountLimits(accountLimits).IfMatch(ifMatch).Execute()

Update account limits



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
	accountId := float32(8.14) // float32 | 
	accountLimits := *openapiclient.NewAccountLimits() // AccountLimits | The account limits updates
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UpdateAccountLimits(context.Background(), accountId).AccountLimits(accountLimits).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccountLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountLimits`: AccountLimits
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccountLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountLimits** | [**AccountLimits**](AccountLimits.md) | The account limits updates | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**AccountLimits**](AccountLimits.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

