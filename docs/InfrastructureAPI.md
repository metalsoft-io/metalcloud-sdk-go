# \InfrastructureAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInfrastructureUser**](InfrastructureAPI.md#AddInfrastructureUser) | **Post** /api/v2/infrastructures/{infrastructureId}/users | Adds a user to the specified infrastructure
[**CancelDeployInfrastructure**](InfrastructureAPI.md#CancelDeployInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/cancel-deploy | Cancels the specified infrastructure deployment
[**CreateInfrastructure**](InfrastructureAPI.md#CreateInfrastructure) | **Post** /api/v2/infrastructures | Creates a new infrastructure
[**DeleteInfrastructure**](InfrastructureAPI.md#DeleteInfrastructure) | **Delete** /api/v2/infrastructures/{infrastructureId} | Deletes the specified infrastructure
[**DeployInfrastructure**](InfrastructureAPI.md#DeployInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/deploy | Deploys the specified infrastructure
[**GetAllInfrastructureStatistics**](InfrastructureAPI.md#GetAllInfrastructureStatistics) | **Get** /api/v2/infrastructures/statistics | Retrieves statistics for all infrastructures
[**GetInfrastructure**](InfrastructureAPI.md#GetInfrastructure) | **Get** /api/v2/infrastructures/{infrastructureId} | Retrieves the specified infrastructure
[**GetInfrastructureConfigInfo**](InfrastructureAPI.md#GetInfrastructureConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/config | Get configuration information about the specified Infrastructure
[**GetInfrastructureResourceUtilizationDetailed**](InfrastructureAPI.md#GetInfrastructureResourceUtilizationDetailed) | **Post** /api/v2/infrastructures/actions/get/resource-utilization-detailed | Gets detailed resource utilization for infrastructures
[**GetInfrastructureResourceUtilizationSummary**](InfrastructureAPI.md#GetInfrastructureResourceUtilizationSummary) | **Post** /api/v2/infrastructures/actions/get/resource-utilization-summarized | Gets summary resource utilization for infrastructures
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


## CancelDeployInfrastructure

> Infrastructure CancelDeployInfrastructure(ctx, infrastructureId).Execute()

Cancels the specified infrastructure deployment



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
	resp, r, err := apiClient.InfrastructureAPI.CancelDeployInfrastructure(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.CancelDeployInfrastructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelDeployInfrastructure`: Infrastructure
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.CancelDeployInfrastructure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDeployInfrastructureRequest struct via the builder pattern


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
	infrastructureCreate := *openapiclient.NewInfrastructureCreate(float32(123)) // InfrastructureCreate | The infrastructure to create

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


## GetAllInfrastructureStatistics

> InfrastructuresStatistics GetAllInfrastructureStatistics(ctx).Execute()

Retrieves statistics for all infrastructures



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
	resp, r, err := apiClient.InfrastructureAPI.GetAllInfrastructureStatistics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetAllInfrastructureStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllInfrastructureStatistics`: InfrastructuresStatistics
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetAllInfrastructureStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInfrastructureStatisticsRequest struct via the builder pattern


### Return type

[**InfrastructuresStatistics**](InfrastructuresStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
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

> InfrastructureResourceUtilizationResponse GetInfrastructureResourceUtilizationDetailed(ctx).GetResourceUtilizationDetailed(getResourceUtilizationDetailed).Execute()

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
	getResourceUtilizationDetailed := *openapiclient.NewGetResourceUtilizationDetailed(float32(123), "StartTimestamp_example", "EndTimestamp_example") // GetResourceUtilizationDetailed | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructureResourceUtilizationDetailed(context.Background()).GetResourceUtilizationDetailed(getResourceUtilizationDetailed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureResourceUtilizationDetailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureResourceUtilizationDetailed`: InfrastructureResourceUtilizationResponse
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructureResourceUtilizationDetailed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureResourceUtilizationDetailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getResourceUtilizationDetailed** | [**GetResourceUtilizationDetailed**](GetResourceUtilizationDetailed.md) |  | 

### Return type

[**InfrastructureResourceUtilizationResponse**](InfrastructureResourceUtilizationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureResourceUtilizationSummary

> InfrastructureResourceUtilizationSummaryResponse GetInfrastructureResourceUtilizationSummary(ctx).GetResourceUtilizationSummarized(getResourceUtilizationSummarized).Execute()

Gets summary resource utilization for infrastructures



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
	getResourceUtilizationSummarized := *openapiclient.NewGetResourceUtilizationSummarized(float32(123), "StartTimestamp_example", "EndTimestamp_example") // GetResourceUtilizationSummarized | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructureResourceUtilizationSummary(context.Background()).GetResourceUtilizationSummarized(getResourceUtilizationSummarized).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureResourceUtilizationSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureResourceUtilizationSummary`: InfrastructureResourceUtilizationSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructureResourceUtilizationSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureResourceUtilizationSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getResourceUtilizationSummarized** | [**GetResourceUtilizationSummarized**](GetResourceUtilizationSummarized.md) |  | 

### Return type

[**InfrastructureResourceUtilizationSummaryResponse**](InfrastructureResourceUtilizationSummaryResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureStatistics

> InfrastructureStatistics GetInfrastructureStatistics(ctx, infrastructureId).Execute()

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
	resp, r, err := apiClient.InfrastructureAPI.GetInfrastructureStatistics(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureAPI.GetInfrastructureStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureStatistics`: InfrastructureStatistics
	fmt.Fprintf(os.Stdout, "Response from `InfrastructureAPI.GetInfrastructureStatistics`: %v\n", resp)
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

[**InfrastructureStatistics**](InfrastructureStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterUserIdOwner := []string{"Inner_example"} // []string | Filter by userIdOwner query param.  **Format:** filter.userIdOwner={$not}:OPERATION:VALUE    **Example:** filter.userIdOwner=$eq:John Doe&filter.userIdOwner=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterDeployId := []string{"Inner_example"} // []string | Filter by deployId query param.  **Format:** filter.deployId={$not}:OPERATION:VALUE    **Example:** filter.deployId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  - subdomain  - subdomainPermanent  - serviceStatus  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,siteId,subdomain,subdomainPermanent   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - siteId  - subdomain  - subdomainPermanent  - serviceStatus  (optional)
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
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterUserIdOwner** | **[]string** | Filter by userIdOwner query param.  **Format:** filter.userIdOwner&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.userIdOwner&#x3D;$eq:John Doe&amp;filter.userIdOwner&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.  **Format:** filter.subdomain&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomain&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterDeployId** | **[]string** | Filter by deployId query param.  **Format:** filter.deployId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.deployId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  - subdomain  - subdomainPermanent  - serviceStatus  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,siteId,subdomain,subdomainPermanent   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - siteId  - subdomain  - subdomainPermanent  - serviceStatus  | 
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
	updateInfrastructureMeta := *openapiclient.NewUpdateInfrastructureMeta("Name_example") // UpdateInfrastructureMeta | The infrastructure metadata to update

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

