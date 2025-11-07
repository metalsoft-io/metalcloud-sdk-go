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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterParentAccountId := []string{"Inner_example"} // []string | Filter by parentAccountId query param.  **Format:** filter.parentAccountId={$not}:OPERATION:VALUE    **Example:** filter.parentAccountId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterPrimaryContactId := []string{"Inner_example"} // []string | Filter by primaryContactId query param.  **Format:** filter.primaryContactId={$not}:OPERATION:VALUE    **Example:** filter.primaryContactId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSecondaryContactId := []string{"Inner_example"} // []string | Filter by secondaryContactId query param.  **Format:** filter.secondaryContactId={$not}:OPERATION:VALUE    **Example:** filter.secondaryContactId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterArchived := []string{"Inner_example"} // []string | Filter by archived query param.  **Format:** filter.archived={$not}:OPERATION:VALUE    **Example:** filter.archived=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterCode := []string{"Inner_example"} // []string | Filter by code query param.  **Format:** filter.code={$not}:OPERATION:VALUE    **Example:** filter.code=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** name:DESC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,code   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - code  (optional)

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
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterParentAccountId** | **[]string** | Filter by parentAccountId query param.  **Format:** filter.parentAccountId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.parentAccountId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterPrimaryContactId** | **[]string** | Filter by primaryContactId query param.  **Format:** filter.primaryContactId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.primaryContactId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSecondaryContactId** | **[]string** | Filter by secondaryContactId query param.  **Format:** filter.secondaryContactId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.secondaryContactId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterArchived** | **[]string** | Filter by archived query param.  **Format:** filter.archived&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.archived&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterCode** | **[]string** | Filter by code query param.  **Format:** filter.code&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.code&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** name:DESC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,code   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - code  | 

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

