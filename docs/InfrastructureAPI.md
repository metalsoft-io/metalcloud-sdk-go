# \InfrastructureAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInfrastructureUser**](InfrastructureAPI.md#AddInfrastructureUser) | **Post** /api/v2/infrastructures/{infrastructureId}/users | Adds a user to the specified infrastructure
[**CreateInfrastructure**](InfrastructureAPI.md#CreateInfrastructure) | **Post** /api/v2/infrastructures | Creates a new infrastructure
[**DeleteInfrastructure**](InfrastructureAPI.md#DeleteInfrastructure) | **Delete** /api/v2/infrastructures/{infrastructureId} | Deletes the specified infrastructure
[**DeployInfrastructure**](InfrastructureAPI.md#DeployInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/deploy | Deploys the specified infrastructure
[**GetInfrastructure**](InfrastructureAPI.md#GetInfrastructure) | **Get** /api/v2/infrastructures/{infrastructureId} | Retrieves the specified infrastructure
[**GetInfrastructureConfigInfo**](InfrastructureAPI.md#GetInfrastructureConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/config | Get configuration information about the specified Infrastructure
[**GetInfrastructureResourceUtilizationDetailed**](InfrastructureAPI.md#GetInfrastructureResourceUtilizationDetailed) | **Post** /api/v2/infrastructures/actions/get/resource-utilization-detailed | Gets detailed resource utilization for infrastructures
[**GetInfrastructureStatistics**](InfrastructureAPI.md#GetInfrastructureStatistics) | **Get** /api/v2/infrastructures/{infrastructureId}/statistics | Retrieves statistics for the specified infrastructure
[**GetInfrastructureUserLimits**](InfrastructureAPI.md#GetInfrastructureUserLimits) | **Get** /api/v2/infrastructures/{infrastructureId}/user-limits | Retrieves the specified infrastructure user limits
[**GetInfrastructureUsers**](InfrastructureAPI.md#GetInfrastructureUsers) | **Get** /api/v2/infrastructures/{infrastructureId}/users | Retrieves the specified infrastructure users
[**GetInfrastructures**](InfrastructureAPI.md#GetInfrastructures) | **Get** /api/v2/infrastructures | Get all infrastructures
[**RemoveInfrastructureUser**](InfrastructureAPI.md#RemoveInfrastructureUser) | **Delete** /api/v2/infrastructures/{infrastructureId}/users/{userId} | Removes a user from the specified infrastructure
[**RevertInfrastructure**](InfrastructureAPI.md#RevertInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/revert | Reverts the specified infrastructure
[**UpdateInfrastructureConfiguration**](InfrastructureAPI.md#UpdateInfrastructureConfiguration) | **Patch** /api/v2/infrastructures/{infrastructureId}/config | Updates the specified infrastructure configuration
[**UpdateInfrastructureMetadata**](InfrastructureAPI.md#UpdateInfrastructureMetadata) | **Patch** /api/v2/infrastructures/{infrastructureId}/metadata | Updates the specified infrastructure metadata



## AddInfrastructureUser

> AddInfrastructureUser(ctx, infrastructureId).AddUserToInfrastructure(addUserToInfrastructure).Execute()

Adds a user to the specified infrastructure



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
	infrastructureId := float32(8.14) // float32 | 
	addUserToInfrastructure := *openapiclient.NewAddUserToInfrastructure("UserEmail_example", false) // AddUserToInfrastructure | Additional information for the user to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfrastructureAPI.AddInfrastructureUser(context.Background(), infrastructureId).AddUserToInfrastructure(addUserToInfrastructure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.AddInfrastructureUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInfrastructureUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUserToInfrastructure** | [**AddUserToInfrastructure**](AddUserToInfrastructure.md) | Additional information for the user to add | 

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


## CreateInfrastructure

> Infrastructure CreateInfrastructure(ctx).InfrastructureCreate(infrastructureCreate).Execute()

Creates a new infrastructure



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
	infrastructureCreate := *openapiclient.NewInfrastructureCreate("Label_example", float32(123)) // InfrastructureCreate | The infrastructure to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.CreateInfrastructure(context.Background()).InfrastructureCreate(infrastructureCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.CreateInfrastructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInfrastructure`: Infrastructure
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.CreateInfrastructure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInfrastructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **infrastructureCreate** | [**InfrastructureCreate**](InfrastructureCreate.md) | The infrastructure to create | 

### Return type

[**Infrastructure**](Infrastructure.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInfrastructure

> DeleteInfrastructure(ctx, infrastructureId).IfMatch(ifMatch).Execute()

Deletes the specified infrastructure



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
	infrastructureId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfrastructureAPI.DeleteInfrastructure(context.Background(), infrastructureId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.DeleteInfrastructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInfrastructureRequest struct via the builder pattern


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


## DeployInfrastructure

> Infrastructure DeployInfrastructure(ctx, infrastructureId).InfrastructureDeployOptions(infrastructureDeployOptions).Execute()

Deploys the specified infrastructure



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
	infrastructureId := float32(8.14) // float32 | 
	infrastructureDeployOptions := *openapiclient.NewInfrastructureDeployOptions(false, *openapiclient.NewInfrastructureDeployShutdownOptions(false, float32(123), false, false)) // InfrastructureDeployOptions | The infrastructure deploy options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.DeployInfrastructure(context.Background(), infrastructureId).InfrastructureDeployOptions(infrastructureDeployOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.DeployInfrastructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployInfrastructure`: Infrastructure
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.DeployInfrastructure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployInfrastructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **infrastructureDeployOptions** | [**InfrastructureDeployOptions**](InfrastructureDeployOptions.md) | The infrastructure deploy options | 

### Return type

[**Infrastructure**](Infrastructure.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructure

> Infrastructure GetInfrastructure(ctx, infrastructureId).Execute()

Retrieves the specified infrastructure



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
	infrastructureId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructure(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructure`: Infrastructure
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Infrastructure**](Infrastructure.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureConfigInfo

> InfrastructureConfiguration GetInfrastructureConfigInfo(ctx, infrastructureId).Execute()

Get configuration information about the specified Infrastructure

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
	infrastructureId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructureConfigInfo(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureConfigInfo`: InfrastructureConfiguration
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructureConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InfrastructureConfiguration**](InfrastructureConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureResourceUtilizationDetailed

> InfrastructureResourceUtilizationResponse GetInfrastructureResourceUtilizationDetailed(ctx, infrastructureId).Execute()

Gets detailed resource utilization for infrastructures



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
	infrastructureId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructureResourceUtilizationDetailed(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureResourceUtilizationDetailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureResourceUtilizationDetailed`: InfrastructureResourceUtilizationResponse
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructureResourceUtilizationDetailed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureResourceUtilizationDetailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InfrastructureResourceUtilizationResponse**](InfrastructureResourceUtilizationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureStatistics

> GetInfrastructureStatistics(ctx, infrastructureId).Execute()

Retrieves statistics for the specified infrastructure



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
	infrastructureId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfrastructureAPI.GetInfrastructureStatistics(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureStatisticsRequest struct via the builder pattern


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


## GetInfrastructureUserLimits

> UserLimits GetInfrastructureUserLimits(ctx, infrastructureId).Execute()

Retrieves the specified infrastructure user limits



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
	infrastructureId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructureUserLimits(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureUserLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureUserLimits`: UserLimits
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructureUserLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureUserLimitsRequest struct via the builder pattern


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


## GetInfrastructureUsers

> UserPaginatedList GetInfrastructureUsers(ctx, infrastructureId).Execute()

Retrieves the specified infrastructure users



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
	infrastructureId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructureUsers(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureUsers`: UserPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructureUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetInfrastructures

> InfrastructurePaginatedList GetInfrastructures(ctx).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSiteId(filterSiteId).FilterUserIdOwner(filterUserIdOwner).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterDeployId(filterDeployId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).AccountId(accountId).Execute()

Get all infrastructures



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterUserIdOwner := []string{"Inner_example"} // []string | Filter by userIdOwner query param.           <p>              <b>Format: </b> filter.userIdOwner={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.userIdOwner=$not:$like:John Doe&filter.userIdOwner=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.           <p>              <b>Format: </b> filter.subdomain={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomain=$not:$like:John Doe&filter.subdomain=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.           <p>              <b>Format: </b> filter.subdomainPermanent={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomainPermanent=$not:$like:John Doe&filter.subdomainPermanent=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterDeployId := []string{"Inner_example"} // []string | Filter by deployId query param.           <p>              <b>Format: </b> filter.deployId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.deployId=$not:$like:John Doe&filter.deployId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>siteId</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>serviceStatus</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,siteId,subdomain,subdomainPermanent           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>siteId</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>serviceStatus</li></ul>          (optional)
	accountId := float32(8.14) // float32 | The account ID to filter user infrastructures by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructures(context.Background()).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSiteId(filterSiteId).FilterUserIdOwner(filterUserIdOwner).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterDeployId(filterDeployId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructures`: InfrastructurePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterUserIdOwner** | **[]string** | Filter by userIdOwner query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.userIdOwner&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.userIdOwner&#x3D;$not:$like:John Doe&amp;filter.userIdOwner&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterDeployId** | **[]string** | Filter by deployId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.deployId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.deployId&#x3D;$not:$like:John Doe&amp;filter.deployId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,siteId,subdomain,subdomainPermanent           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt;&lt;/ul&gt;          | 
 **accountId** | **float32** | The account ID to filter user infrastructures by | 

### Return type

[**InfrastructurePaginatedList**](InfrastructurePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInfrastructureUser

> RemoveInfrastructureUser(ctx, infrastructureId, userId).Execute()

Removes a user from the specified infrastructure



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
	infrastructureId := float32(8.14) // float32 | 
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfrastructureAPI.RemoveInfrastructureUser(context.Background(), infrastructureId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.RemoveInfrastructureUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInfrastructureUserRequest struct via the builder pattern


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


## RevertInfrastructure

> RevertInfrastructure(ctx, infrastructureId).IfMatch(ifMatch).Execute()

Reverts the specified infrastructure



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
	infrastructureId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfrastructureAPI.RevertInfrastructure(context.Background(), infrastructureId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.RevertInfrastructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertInfrastructureRequest struct via the builder pattern


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


## UpdateInfrastructureConfiguration

> Infrastructure UpdateInfrastructureConfiguration(ctx, infrastructureId).UpdateInfrastructure(updateInfrastructure).IfMatch(ifMatch).Execute()

Updates the specified infrastructure configuration



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
	infrastructureId := float32(8.14) // float32 | 
	updateInfrastructure := *openapiclient.NewUpdateInfrastructure() // UpdateInfrastructure | The infrastructure configuration to update
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.UpdateInfrastructureConfiguration(context.Background(), infrastructureId).UpdateInfrastructure(updateInfrastructure).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.UpdateInfrastructureConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInfrastructureConfiguration`: Infrastructure
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.UpdateInfrastructureConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInfrastructureConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInfrastructure** | [**UpdateInfrastructure**](UpdateInfrastructure.md) | The infrastructure configuration to update | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Infrastructure**](Infrastructure.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInfrastructureMetadata

> Infrastructure UpdateInfrastructureMetadata(ctx, infrastructureId).UpdateInfrastructureMeta(updateInfrastructureMeta).Execute()

Updates the specified infrastructure metadata



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
	infrastructureId := float32(8.14) // float32 | 
	updateInfrastructureMeta := *openapiclient.NewUpdateInfrastructureMeta() // UpdateInfrastructureMeta | The infrastructure metadata to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.UpdateInfrastructureMetadata(context.Background(), infrastructureId).UpdateInfrastructureMeta(updateInfrastructureMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.UpdateInfrastructureMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInfrastructureMetadata`: Infrastructure
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.UpdateInfrastructureMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInfrastructureMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInfrastructureMeta** | [**UpdateInfrastructureMeta**](UpdateInfrastructureMeta.md) | The infrastructure metadata to update | 

### Return type

[**Infrastructure**](Infrastructure.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

