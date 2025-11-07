# \ServerTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerType**](ServerTypeAPI.md#CreateServerType) | **Post** /api/v2/server-types | Creates a Server Type
[**DeleteServerType**](ServerTypeAPI.md#DeleteServerType) | **Delete** /api/v2/server-types/{serverTypeId} | Deletes a Server Type
[**GetServerTypeInfo**](ServerTypeAPI.md#GetServerTypeInfo) | **Get** /api/v2/server-types/{serverTypeId} | Get Server Type information
[**GetServerTypes**](ServerTypeAPI.md#GetServerTypes) | **Get** /api/v2/server-types | Get a list of Server Types
[**GetServerTypesStatisticsBatch**](ServerTypeAPI.md#GetServerTypesStatisticsBatch) | **Post** /api/v2/server-types/statistics | Get Server Type statistics batch
[**RemoveUnusedServerTypes**](ServerTypeAPI.md#RemoveUnusedServerTypes) | **Post** /api/v2/server-types/actions/clean-unused | Deletes unused server types
[**UpdateServerType**](ServerTypeAPI.md#UpdateServerType) | **Patch** /api/v2/server-types/{serverTypeId} | Updates a Server Type



## CreateServerType

> ServerType CreateServerType(ctx).CreateServerType(createServerType).Execute()

Creates a Server Type



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
	createServerType := *openapiclient.NewCreateServerType(float32(123), float32(123), float32(123), float32(123), "Name_example", "Label_example", []float32{float32(123)}, []string{"ProcessorNames_example"}, float32(123), "ServerClass_example") // CreateServerType | The Server Type create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.CreateServerType(context.Background()).CreateServerType(createServerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.CreateServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerType`: ServerType
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.CreateServerType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerType** | [**CreateServerType**](CreateServerType.md) | The Server Type create object | 

### Return type

[**ServerType**](ServerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerType

> DeleteServerType(ctx, serverTypeId).Execute()

Deletes a Server Type



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
	serverTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerTypeAPI.DeleteServerType(context.Background(), serverTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.DeleteServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerTypeRequest struct via the builder pattern


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


## GetServerTypeInfo

> ServerType GetServerTypeInfo(ctx, serverTypeId).Execute()

Get Server Type information



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
	serverTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.GetServerTypeInfo(context.Background(), serverTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.GetServerTypeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerTypeInfo`: ServerType
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.GetServerTypeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerTypeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerType**](ServerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerTypes

> ServerTypePaginatedList GetServerTypes(ctx).Page(page).Limit(limit).FilterName(filterName).FilterLabel(filterLabel).FilterRamGbytes(filterRamGbytes).FilterProcessorCount(filterProcessorCount).FilterProcessorCoreCount(filterProcessorCoreCount).FilterDescription(filterDescription).FilterNetworkTotalCapacityMbps(filterNetworkTotalCapacityMbps).FilterIsExperimental(filterIsExperimental).FilterForUnmanagedServersOnly(filterForUnmanagedServersOnly).FilterForGenericEndpointsOnly(filterForGenericEndpointsOnly).FilterSupportsOobProvisioning(filterSupportsOobProvisioning).FilterAllowedVendorSkuIds(filterAllowedVendorSkuIds).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Server Types



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterRamGbytes := []string{"Inner_example"} // []string | Filter by ramGbytes query param.  **Format:** filter.ramGbytes={$not}:OPERATION:VALUE    **Example:** filter.ramGbytes=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterProcessorCount := []string{"Inner_example"} // []string | Filter by processorCount query param.  **Format:** filter.processorCount={$not}:OPERATION:VALUE    **Example:** filter.processorCount=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterProcessorCoreCount := []string{"Inner_example"} // []string | Filter by processorCoreCount query param.  **Format:** filter.processorCoreCount={$not}:OPERATION:VALUE    **Example:** filter.processorCoreCount=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.  **Format:** filter.description={$not}:OPERATION:VALUE    **Example:** filter.description=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterNetworkTotalCapacityMbps := []string{"Inner_example"} // []string | Filter by networkTotalCapacityMbps query param.  **Format:** filter.networkTotalCapacityMbps={$not}:OPERATION:VALUE    **Example:** filter.networkTotalCapacityMbps=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterIsExperimental := []string{"Inner_example"} // []string | Filter by isExperimental query param.  **Format:** filter.isExperimental={$not}:OPERATION:VALUE    **Example:** filter.isExperimental=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterForUnmanagedServersOnly := []string{"Inner_example"} // []string | Filter by forUnmanagedServersOnly query param.  **Format:** filter.forUnmanagedServersOnly={$not}:OPERATION:VALUE    **Example:** filter.forUnmanagedServersOnly=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterForGenericEndpointsOnly := []string{"Inner_example"} // []string | Filter by forGenericEndpointsOnly query param.  **Format:** filter.forGenericEndpointsOnly={$not}:OPERATION:VALUE    **Example:** filter.forGenericEndpointsOnly=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSupportsOobProvisioning := []string{"Inner_example"} // []string | Filter by supportsOobProvisioning query param.  **Format:** filter.supportsOobProvisioning={$not}:OPERATION:VALUE    **Example:** filter.supportsOobProvisioning=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterAllowedVendorSkuIds := []string{"Inner_example"} // []string | Filter by allowedVendorSkuIds query param.  **Format:** filter.allowedVendorSkuIds={$not}:OPERATION:VALUE    **Example:** filter.allowedVendorSkuIds=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.GetServerTypes(context.Background()).Page(page).Limit(limit).FilterName(filterName).FilterLabel(filterLabel).FilterRamGbytes(filterRamGbytes).FilterProcessorCount(filterProcessorCount).FilterProcessorCoreCount(filterProcessorCoreCount).FilterDescription(filterDescription).FilterNetworkTotalCapacityMbps(filterNetworkTotalCapacityMbps).FilterIsExperimental(filterIsExperimental).FilterForUnmanagedServersOnly(filterForUnmanagedServersOnly).FilterForGenericEndpointsOnly(filterForGenericEndpointsOnly).FilterSupportsOobProvisioning(filterSupportsOobProvisioning).FilterAllowedVendorSkuIds(filterAllowedVendorSkuIds).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.GetServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerTypes`: ServerTypePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.GetServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterRamGbytes** | **[]string** | Filter by ramGbytes query param.  **Format:** filter.ramGbytes&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.ramGbytes&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterProcessorCount** | **[]string** | Filter by processorCount query param.  **Format:** filter.processorCount&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.processorCount&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterProcessorCoreCount** | **[]string** | Filter by processorCoreCount query param.  **Format:** filter.processorCoreCount&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.processorCoreCount&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterDescription** | **[]string** | Filter by description query param.  **Format:** filter.description&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.description&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterNetworkTotalCapacityMbps** | **[]string** | Filter by networkTotalCapacityMbps query param.  **Format:** filter.networkTotalCapacityMbps&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkTotalCapacityMbps&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterIsExperimental** | **[]string** | Filter by isExperimental query param.  **Format:** filter.isExperimental&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.isExperimental&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterForUnmanagedServersOnly** | **[]string** | Filter by forUnmanagedServersOnly query param.  **Format:** filter.forUnmanagedServersOnly&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.forUnmanagedServersOnly&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterForGenericEndpointsOnly** | **[]string** | Filter by forGenericEndpointsOnly query param.  **Format:** filter.forGenericEndpointsOnly&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.forGenericEndpointsOnly&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSupportsOobProvisioning** | **[]string** | Filter by supportsOobProvisioning query param.  **Format:** filter.supportsOobProvisioning&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.supportsOobProvisioning&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterAllowedVendorSkuIds** | **[]string** | Filter by allowedVendorSkuIds query param.  **Format:** filter.allowedVendorSkuIds&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.allowedVendorSkuIds&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  | 

### Return type

[**ServerTypePaginatedList**](ServerTypePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerTypesStatisticsBatch

> ServerTypeStatisticsBatch GetServerTypesStatisticsBatch(ctx).ServerTypeStatisticsBatchOptions(serverTypeStatisticsBatchOptions).Execute()

Get Server Type statistics batch



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
	serverTypeStatisticsBatchOptions := *openapiclient.NewServerTypeStatisticsBatchOptions(float32(123)) // ServerTypeStatisticsBatchOptions | The server type statistics batch options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.GetServerTypesStatisticsBatch(context.Background()).ServerTypeStatisticsBatchOptions(serverTypeStatisticsBatchOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.GetServerTypesStatisticsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerTypesStatisticsBatch`: ServerTypeStatisticsBatch
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.GetServerTypesStatisticsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerTypesStatisticsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverTypeStatisticsBatchOptions** | [**ServerTypeStatisticsBatchOptions**](ServerTypeStatisticsBatchOptions.md) | The server type statistics batch options | 

### Return type

[**ServerTypeStatisticsBatch**](ServerTypeStatisticsBatch.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUnusedServerTypes

> RemoveUnusedServerTypes(ctx).Execute()

Deletes unused server types



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
	r, err := apiClient.ServerTypeAPI.RemoveUnusedServerTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.RemoveUnusedServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUnusedServerTypesRequest struct via the builder pattern


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


## UpdateServerType

> ServerType UpdateServerType(ctx, serverTypeId).UpdateServerType(updateServerType).Execute()

Updates a Server Type



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
	serverTypeId := float32(8.14) // float32 | 
	updateServerType := *openapiclient.NewUpdateServerType("Label_example") // UpdateServerType | The Server Type update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.UpdateServerType(context.Background(), serverTypeId).UpdateServerType(updateServerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.UpdateServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerType`: ServerType
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.UpdateServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerType** | [**UpdateServerType**](UpdateServerType.md) | The Server Type update object | 

### Return type

[**ServerType**](ServerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

