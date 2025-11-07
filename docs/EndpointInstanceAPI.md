# \EndpointInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpointInstance**](EndpointInstanceAPI.md#CreateEndpointInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/endpoint-instances | Add Endpoint Instance to an infrastructure
[**DeleteEndpointInstance**](EndpointInstanceAPI.md#DeleteEndpointInstance) | **Delete** /api/v2/endpoint-instances/{endpointInstanceId} | Delete Endpoint Instance
[**GetEndpointInstance**](EndpointInstanceAPI.md#GetEndpointInstance) | **Get** /api/v2/endpoint-instances/{endpointInstanceId} | Get Endpoint Instance details
[**GetEndpointInstanceConfig**](EndpointInstanceAPI.md#GetEndpointInstanceConfig) | **Get** /api/v2/endpoint-instances/{endpointInstanceId}/config | Get Endpoint Instance config details
[**GetEndpointInstances**](EndpointInstanceAPI.md#GetEndpointInstances) | **Get** /api/v2/endpoint-instances | List Endpoint Instances
[**GetInfrastructureEndpointInstances**](EndpointInstanceAPI.md#GetInfrastructureEndpointInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/endpoint-instances | List Endpoint Instances for an infrastructure
[**UpdateEndpointInstanceConfig**](EndpointInstanceAPI.md#UpdateEndpointInstanceConfig) | **Patch** /api/v2/endpoint-instances/{endpointInstanceId}/config | Update Endpoint Instance configuration
[**UpdateEndpointInstanceMeta**](EndpointInstanceAPI.md#UpdateEndpointInstanceMeta) | **Patch** /api/v2/endpoint-instances/{endpointInstanceId}/meta | Update an Endpoint Instance meta information



## CreateEndpointInstance

> EndpointInstance CreateEndpointInstance(ctx, infrastructureId).EndpointInstanceCreate(endpointInstanceCreate).Execute()

Add Endpoint Instance to an infrastructure



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
	infrastructureId := int32(56) // int32 | 
	endpointInstanceCreate := *openapiclient.NewEndpointInstanceCreate(int32(123)) // EndpointInstanceCreate | The Endpoint Instance to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.CreateEndpointInstance(context.Background(), infrastructureId).EndpointInstanceCreate(endpointInstanceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.CreateEndpointInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointInstance`: EndpointInstance
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.CreateEndpointInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointInstanceCreate** | [**EndpointInstanceCreate**](EndpointInstanceCreate.md) | The Endpoint Instance to create | 

### Return type

[**EndpointInstance**](EndpointInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpointInstance

> DeleteEndpointInstance(ctx, endpointInstanceId).IfMatch(ifMatch).Execute()

Delete Endpoint Instance



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
	endpointInstanceId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceAPI.DeleteEndpointInstance(context.Background(), endpointInstanceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.DeleteEndpointInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointInstanceRequest struct via the builder pattern


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


## GetEndpointInstance

> EndpointInstance GetEndpointInstance(ctx, endpointInstanceId).Execute()

Get Endpoint Instance details



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
	endpointInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetEndpointInstance(context.Background(), endpointInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetEndpointInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstance`: EndpointInstance
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetEndpointInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointInstance**](EndpointInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceConfig

> EndpointInstanceConfiguration GetEndpointInstanceConfig(ctx, endpointInstanceId).Execute()

Get Endpoint Instance config details



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
	endpointInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetEndpointInstanceConfig(context.Background(), endpointInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetEndpointInstanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceConfig`: EndpointInstanceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetEndpointInstanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointInstanceConfiguration**](EndpointInstanceConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstances

> EndpointInstancePaginatedList GetEndpointInstances(ctx).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Endpoint Instances



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
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterGroupId := []string{"Inner_example"} // []string | Filter by groupId query param.  **Format:** filter.groupId={$not}:OPERATION:VALUE    **Example:** filter.groupId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterEndpointId := []string{"Inner_example"} // []string | Filter by endpointId query param.  **Format:** filter.endpointId={$not}:OPERATION:VALUE    **Example:** filter.endpointId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigEndpointId := []string{"Inner_example"} // []string | Filter by config.endpointId query param.  **Format:** filter.config.endpointId={$not}:OPERATION:VALUE    **Example:** filter.config.endpointId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetEndpointInstances(context.Background()).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetEndpointInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstances`: EndpointInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetEndpointInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterGroupId** | **[]string** | Filter by groupId query param.  **Format:** filter.groupId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.groupId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterEndpointId** | **[]string** | Filter by endpointId query param.  **Format:** filter.endpointId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.endpointId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigEndpointId** | **[]string** | Filter by config.endpointId query param.  **Format:** filter.config.endpointId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.endpointId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  | 

### Return type

[**EndpointInstancePaginatedList**](EndpointInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureEndpointInstances

> EndpointInstancePaginatedList GetInfrastructureEndpointInstances(ctx, infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Endpoint Instances for an infrastructure



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
	infrastructureId := int32(56) // int32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterGroupId := []string{"Inner_example"} // []string | Filter by groupId query param.  **Format:** filter.groupId={$not}:OPERATION:VALUE    **Example:** filter.groupId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterEndpointId := []string{"Inner_example"} // []string | Filter by endpointId query param.  **Format:** filter.endpointId={$not}:OPERATION:VALUE    **Example:** filter.endpointId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigEndpointId := []string{"Inner_example"} // []string | Filter by config.endpointId query param.  **Format:** filter.config.endpointId={$not}:OPERATION:VALUE    **Example:** filter.config.endpointId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetInfrastructureEndpointInstances(context.Background(), infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetInfrastructureEndpointInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureEndpointInstances`: EndpointInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetInfrastructureEndpointInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureEndpointInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterGroupId** | **[]string** | Filter by groupId query param.  **Format:** filter.groupId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.groupId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterEndpointId** | **[]string** | Filter by endpointId query param.  **Format:** filter.endpointId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.endpointId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigEndpointId** | **[]string** | Filter by config.endpointId query param.  **Format:** filter.config.endpointId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.endpointId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  | 

### Return type

[**EndpointInstancePaginatedList**](EndpointInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceConfig

> EndpointInstanceConfiguration UpdateEndpointInstanceConfig(ctx, endpointInstanceId).EndpointInstanceUpdate(endpointInstanceUpdate).IfMatch(ifMatch).Execute()

Update Endpoint Instance configuration



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
	endpointInstanceId := int32(56) // int32 | 
	endpointInstanceUpdate := *openapiclient.NewEndpointInstanceUpdate() // EndpointInstanceUpdate | The Endpoint Instance configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.UpdateEndpointInstanceConfig(context.Background(), endpointInstanceId).EndpointInstanceUpdate(endpointInstanceUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.UpdateEndpointInstanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointInstanceConfig`: EndpointInstanceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.UpdateEndpointInstanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointInstanceUpdate** | [**EndpointInstanceUpdate**](EndpointInstanceUpdate.md) | The Endpoint Instance configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**EndpointInstanceConfiguration**](EndpointInstanceConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceMeta

> UpdateEndpointInstanceMeta(ctx, endpointInstanceId).GenericMeta(genericMeta).Execute()

Update an Endpoint Instance meta information



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
	endpointInstanceId := int32(56) // int32 | 
	genericMeta := *openapiclient.NewGenericMeta() // GenericMeta | The Endpoint Instance meta information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceAPI.UpdateEndpointInstanceMeta(context.Background(), endpointInstanceId).GenericMeta(genericMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.UpdateEndpointInstanceMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genericMeta** | [**GenericMeta**](GenericMeta.md) | The Endpoint Instance meta information | 

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

