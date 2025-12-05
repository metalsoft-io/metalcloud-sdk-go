# \VMPoolAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVMPool**](VMPoolAPI.md#CreateVMPool) | **Post** /api/v2/vm-pools | Creates a VM Pool
[**DeleteVMPool**](VMPoolAPI.md#DeleteVMPool) | **Delete** /api/v2/vm-pools/{vmPoolId} | Deletes a VM Pool
[**GetVMPool**](VMPoolAPI.md#GetVMPool) | **Get** /api/v2/vm-pools/{vmPoolId} | Get VM Pool information
[**GetVMPoolClusterHost**](VMPoolAPI.md#GetVMPoolClusterHost) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId} | Retrieves a VM Cluster Host
[**GetVMPoolClusterHostInterface**](VMPoolAPI.md#GetVMPoolClusterHostInterface) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Retrieves a VM Cluster Host Interface
[**GetVMPoolClusterHostInterfaces**](VMPoolAPI.md#GetVMPoolClusterHostInterfaces) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces | Retrieves a list of VM Cluster Host Interfaces
[**GetVMPoolClusterHostStatistics**](VMPoolAPI.md#GetVMPoolClusterHostStatistics) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/statistics | Retrieves VM Cluster Host Statistics
[**GetVMPoolClusterHostVMs**](VMPoolAPI.md#GetVMPoolClusterHostVMs) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/vms | Retrieves a list of VM Cluster Host VMs
[**GetVMPoolClusterHosts**](VMPoolAPI.md#GetVMPoolClusterHosts) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts | Get list of VM Cluster Hosts linked to the VM Pool
[**GetVMPoolCredentials**](VMPoolAPI.md#GetVMPoolCredentials) | **Get** /api/v2/vm-pools/{vmPoolId}/credentials | Get VM Pool credentials
[**GetVMPoolVMs**](VMPoolAPI.md#GetVMPoolVMs) | **Get** /api/v2/vm-pools/{vmPoolId}/vms | Returns all VMs linked to the VM Pool
[**GetVMPools**](VMPoolAPI.md#GetVMPools) | **Get** /api/v2/vm-pools | Get all VM Pools
[**GetVmPoolStatistics**](VMPoolAPI.md#GetVmPoolStatistics) | **Get** /api/v2/vm-pools/{vmPoolId}/statistics | Get VM Pool statistics
[**ImportVMPoolVMs**](VMPoolAPI.md#ImportVMPoolVMs) | **Post** /api/v2/vm-pools/{vmPoolId}/actions/import-vms | Import VMs into VM Pool
[**RefreshVMPoolInformation**](VMPoolAPI.md#RefreshVMPoolInformation) | **Post** /api/v2/vm-pools/{vmPoolId}/actions/refresh-information | Refresh VM Pool information
[**SyncVMPool**](VMPoolAPI.md#SyncVMPool) | **Post** /api/v2/vm-pools/{vmPoolId}/actions/sync | Sync VM Pool
[**UpdateVMPool**](VMPoolAPI.md#UpdateVMPool) | **Patch** /api/v2/vm-pools/{vmPoolId} | Updates VM Pool information
[**UpdateVMPoolClusterHostInterface**](VMPoolAPI.md#UpdateVMPoolClusterHostInterface) | **Patch** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Updates a VM Cluster Host Interface



## CreateVMPool

> VMPool CreateVMPool(ctx).CreateVMPool(createVMPool).Execute()

Creates a VM Pool



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
	createVMPool := *openapiclient.NewCreateVMPool(float32(123), "ManagementHost_example", float32(123), "Name_example", "Type_example", float32(123)) // CreateVMPool | The VM Pool create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.CreateVMPool(context.Background()).CreateVMPool(createVMPool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.CreateVMPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVMPool`: VMPool
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.CreateVMPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVMPool** | [**CreateVMPool**](CreateVMPool.md) | The VM Pool create object | 

### Return type

[**VMPool**](VMPool.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVMPool

> DeleteVMPool(ctx, vmPoolId).Execute()

Deletes a VM Pool



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
	vmPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMPoolAPI.DeleteVMPool(context.Background(), vmPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.DeleteVMPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMPoolRequest struct via the builder pattern


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


## GetVMPool

> VMPool GetVMPool(ctx, vmPoolId).Execute()

Get VM Pool information



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
	vmPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPool(context.Background(), vmPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPool`: VMPool
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VMPool**](VMPool.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolClusterHost

> VMPoolHosts GetVMPoolClusterHost(ctx, vmPoolId, vmPoolClusterHostId).Execute()

Retrieves a VM Cluster Host



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
	vmPoolId := float32(8.14) // float32 | 
	vmPoolClusterHostId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolClusterHost(context.Background(), vmPoolId, vmPoolClusterHostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolClusterHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolClusterHost`: VMPoolHosts
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolClusterHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 
**vmPoolClusterHostId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolClusterHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMPoolHosts**](VMPoolHosts.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolClusterHostInterface

> VMPoolHostInterfaces GetVMPoolClusterHostInterface(ctx, vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId).Execute()

Retrieves a VM Cluster Host Interface



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
	vmPoolId := float32(8.14) // float32 | 
	vmPoolClusterHostId := float32(8.14) // float32 | 
	vmPoolClusterHostInterfaceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolClusterHostInterface(context.Background(), vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolClusterHostInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolClusterHostInterface`: VMPoolHostInterfaces
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolClusterHostInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 
**vmPoolClusterHostId** | **float32** |  | 
**vmPoolClusterHostInterfaceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolClusterHostInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VMPoolHostInterfaces**](VMPoolHostInterfaces.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolClusterHostInterfaces

> []VMPoolHostInterfaces GetVMPoolClusterHostInterfaces(ctx, vmPoolId, vmPoolClusterHostId).Execute()

Retrieves a list of VM Cluster Host Interfaces



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
	vmPoolId := float32(8.14) // float32 | 
	vmPoolClusterHostId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolClusterHostInterfaces(context.Background(), vmPoolId, vmPoolClusterHostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolClusterHostInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolClusterHostInterfaces`: []VMPoolHostInterfaces
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolClusterHostInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 
**vmPoolClusterHostId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolClusterHostInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]VMPoolHostInterfaces**](VMPoolHostInterfaces.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolClusterHostStatistics

> VMPoolStatistics GetVMPoolClusterHostStatistics(ctx, vmPoolId, vmPoolClusterHostId).Execute()

Retrieves VM Cluster Host Statistics



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
	vmPoolId := float32(8.14) // float32 | 
	vmPoolClusterHostId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolClusterHostStatistics(context.Background(), vmPoolId, vmPoolClusterHostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolClusterHostStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolClusterHostStatistics`: VMPoolStatistics
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolClusterHostStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 
**vmPoolClusterHostId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolClusterHostStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMPoolStatistics**](VMPoolStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolClusterHostVMs

> VMPaginatedList GetVMPoolClusterHostVMs(ctx, vmPoolId, vmPoolClusterHostId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Retrieves a list of VM Cluster Host VMs



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
	vmPoolId := float32(8.14) // float32 | 
	vmPoolClusterHostId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAddress := []string{"Inner_example"} // []string | Filter by address query param.  **Format:** filter.address={$not}:OPERATION:VALUE    **Example:** filter.address=$btw:John Doe&filter.address=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolClusterHostVMs(context.Background(), vmPoolId, vmPoolClusterHostId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolClusterHostVMs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolClusterHostVMs`: VMPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolClusterHostVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 
**vmPoolClusterHostId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolClusterHostVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAddress** | **[]string** | Filter by address query param.  **Format:** filter.address&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.address&#x3D;$btw:John Doe&amp;filter.address&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 

### Return type

[**VMPaginatedList**](VMPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolClusterHosts

> VMPoolHostsPaginatedList GetVMPoolClusterHosts(ctx, vmPoolId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get list of VM Cluster Hosts linked to the VM Pool



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
	vmPoolId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAddress := []string{"Inner_example"} // []string | Filter by address query param.  **Format:** filter.address={$not}:OPERATION:VALUE    **Example:** filter.address=$btw:John Doe&filter.address=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolClusterHosts(context.Background(), vmPoolId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolClusterHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolClusterHosts`: VMPoolHostsPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolClusterHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolClusterHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAddress** | **[]string** | Filter by address query param.  **Format:** filter.address&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.address&#x3D;$btw:John Doe&amp;filter.address&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 

### Return type

[**VMPoolHostsPaginatedList**](VMPoolHostsPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolCredentials

> VMPoolCredentials GetVMPoolCredentials(ctx, vmPoolId).Execute()

Get VM Pool credentials



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
	vmPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolCredentials(context.Background(), vmPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolCredentials`: VMPoolCredentials
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VMPoolCredentials**](VMPoolCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPoolVMs

> VMPaginatedList GetVMPoolVMs(ctx, vmPoolId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Returns all VMs linked to the VM Pool



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
	vmPoolId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAddress := []string{"Inner_example"} // []string | Filter by address query param.  **Format:** filter.address={$not}:OPERATION:VALUE    **Example:** filter.address=$btw:John Doe&filter.address=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPoolVMs(context.Background(), vmPoolId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPoolVMs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPoolVMs`: VMPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPoolVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAddress** | **[]string** | Filter by address query param.  **Format:** filter.address&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.address&#x3D;$btw:John Doe&amp;filter.address&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 

### Return type

[**VMPaginatedList**](VMPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPools

> VMPoolPaginatedList GetVMPools(ctx).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterSiteId(filterSiteId).FilterManagementHost(filterManagementHost).FilterStatus(filterStatus).FilterType(filterType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all VM Pools



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementHost := []string{"Inner_example"} // []string | Filter by managementHost query param.  **Format:** filter.managementHost={$not}:OPERATION:VALUE    **Example:** filter.managementHost=$btw:John Doe&filter.managementHost=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterType := []string{"Inner_example"} // []string | Filter by type query param.  **Format:** filter.type={$not}:OPERATION:VALUE    **Example:** filter.type=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - type  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVMPools(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterSiteId(filterSiteId).FilterManagementHost(filterManagementHost).FilterStatus(filterStatus).FilterType(filterType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVMPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPools`: VMPoolPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVMPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementHost** | **[]string** | Filter by managementHost query param.  **Format:** filter.managementHost&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementHost&#x3D;$btw:John Doe&amp;filter.managementHost&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterType** | **[]string** | Filter by type query param.  **Format:** filter.type&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.type&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - type  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 

### Return type

[**VMPoolPaginatedList**](VMPoolPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmPoolStatistics

> VMPoolStatistics GetVmPoolStatistics(ctx, vmPoolId).Execute()

Get VM Pool statistics



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
	vmPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.GetVmPoolStatistics(context.Background(), vmPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.GetVmPoolStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVmPoolStatistics`: VMPoolStatistics
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.GetVmPoolStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmPoolStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VMPoolStatistics**](VMPoolStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportVMPoolVMs

> ImportVMPoolVMs(ctx, vmPoolId).VMPoolImportVMs(vMPoolImportVMs).Execute()

Import VMs into VM Pool



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
	vmPoolId := float32(8.14) // float32 | 
	vMPoolImportVMs := *openapiclient.NewVMPoolImportVMs([]string{"VmNames_example"}) // VMPoolImportVMs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMPoolAPI.ImportVMPoolVMs(context.Background(), vmPoolId).VMPoolImportVMs(vMPoolImportVMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.ImportVMPoolVMs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportVMPoolVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vMPoolImportVMs** | [**VMPoolImportVMs**](VMPoolImportVMs.md) |  | 

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


## RefreshVMPoolInformation

> VMPool RefreshVMPoolInformation(ctx, vmPoolId).Execute()

Refresh VM Pool information



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
	vmPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.RefreshVMPoolInformation(context.Background(), vmPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.RefreshVMPoolInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshVMPoolInformation`: VMPool
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.RefreshVMPoolInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshVMPoolInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VMPool**](VMPool.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncVMPool

> JobInfo SyncVMPool(ctx, vmPoolId).Execute()

Sync VM Pool



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
	vmPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.SyncVMPool(context.Background(), vmPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.SyncVMPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncVMPool`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.SyncVMPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncVMPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMPool

> VMPool UpdateVMPool(ctx, vmPoolId).UpdateVMPool(updateVMPool).Execute()

Updates VM Pool information



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
	vmPoolId := float32(8.14) // float32 | 
	updateVMPool := *openapiclient.NewUpdateVMPool() // UpdateVMPool | The VM Pool update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.UpdateVMPool(context.Background(), vmPoolId).UpdateVMPool(updateVMPool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.UpdateVMPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMPool`: VMPool
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.UpdateVMPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVMPool** | [**UpdateVMPool**](UpdateVMPool.md) | The VM Pool update object | 

### Return type

[**VMPool**](VMPool.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMPoolClusterHostInterface

> VMPoolHostInterfaces UpdateVMPoolClusterHostInterface(ctx, vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId).UpdateVMPoolHostInterface(updateVMPoolHostInterface).Execute()

Updates a VM Cluster Host Interface



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
	vmPoolId := float32(8.14) // float32 | 
	vmPoolClusterHostId := float32(8.14) // float32 | 
	vmPoolClusterHostInterfaceId := float32(8.14) // float32 | 
	updateVMPoolHostInterface := *openapiclient.NewUpdateVMPoolHostInterface() // UpdateVMPoolHostInterface | The VM Pool Cluster Host Interface update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPoolAPI.UpdateVMPoolClusterHostInterface(context.Background(), vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId).UpdateVMPoolHostInterface(updateVMPoolHostInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPoolAPI.UpdateVMPoolClusterHostInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMPoolClusterHostInterface`: VMPoolHostInterfaces
	fmt.Fprintf(os.Stdout, "Response from `VMPoolAPI.UpdateVMPoolClusterHostInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmPoolId** | **float32** |  | 
**vmPoolClusterHostId** | **float32** |  | 
**vmPoolClusterHostInterfaceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMPoolClusterHostInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateVMPoolHostInterface** | [**UpdateVMPoolHostInterface**](UpdateVMPoolHostInterface.md) | The VM Pool Cluster Host Interface update object | 

### Return type

[**VMPoolHostInterfaces**](VMPoolHostInterfaces.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

