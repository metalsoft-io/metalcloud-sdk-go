# \VMAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVM**](VMAPI.md#GetVM) | **Get** /api/v2/vms/{vmId} | Retrieves the VM information
[**GetVMPowerStatus**](VMAPI.md#GetVMPowerStatus) | **Get** /api/v2/vms/{vmId}/power-status | Retrieves the power status of the VM
[**GetVMRemoteConsoleInfo**](VMAPI.md#GetVMRemoteConsoleInfo) | **Get** /api/v2/vms/{vmId}/remote-console-info | Get Remote Console information
[**GetVMs**](VMAPI.md#GetVMs) | **Get** /api/v2/vms | Get a list of VMs
[**RebootVM**](VMAPI.md#RebootVM) | **Post** /api/v2/vms/{vmId}/reboot | Reboots the VM
[**ShutdownVM**](VMAPI.md#ShutdownVM) | **Post** /api/v2/vms/{vmId}/shutdown | Shuts down the VM
[**StartVM**](VMAPI.md#StartVM) | **Post** /api/v2/vms/{vmId}/start | Starts the VM
[**UpdateVM**](VMAPI.md#UpdateVM) | **Patch** /api/v2/vms/{vmId} | Updates VM information



## GetVM

> VM GetVM(ctx, vmId).Execute()

Retrieves the VM information



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
	vmId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMAPI.GetVM(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.GetVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVM`: VM
	fmt.Fprintf(os.Stdout, "Response from `VMAPI.GetVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VM**](VM.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMPowerStatus

> string GetVMPowerStatus(ctx, vmId).Execute()

Retrieves the power status of the VM



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
	vmId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMAPI.GetVMPowerStatus(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.GetVMPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMPowerStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `VMAPI.GetVMPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMPowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMRemoteConsoleInfo

> RemoteConsoleInfo GetVMRemoteConsoleInfo(ctx, vmId).Execute()

Get Remote Console information



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
	vmId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMAPI.GetVMRemoteConsoleInfo(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.GetVMRemoteConsoleInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMRemoteConsoleInfo`: RemoteConsoleInfo
	fmt.Fprintf(os.Stdout, "Response from `VMAPI.GetVMRemoteConsoleInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMRemoteConsoleInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteConsoleInfo**](RemoteConsoleInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMs

> VMPaginatedList GetVMs(ctx).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterName(filterName).FilterAddress(filterAddress).FilterHost(filterHost).FilterHosts(filterHosts).FilterTypeId(filterTypeId).FilterPoolId(filterPoolId).FilterAdministrationState(filterAdministrationState).FilterNumaNodes(filterNumaNodes).FilterInfrastructureId(filterInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of VMs



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
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAddress := []string{"Inner_example"} // []string | Filter by address query param.  **Format:** filter.address={$not}:OPERATION:VALUE    **Example:** filter.address=$btw:John Doe&filter.address=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHost := []string{"Inner_example"} // []string | Filter by host query param.  **Format:** filter.host={$not}:OPERATION:VALUE    **Example:** filter.host=$btw:John Doe&filter.host=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHosts := []string{"Inner_example"} // []string | Filter by hosts query param.  **Format:** filter.hosts={$not}:OPERATION:VALUE    **Example:** filter.hosts=$btw:John Doe&filter.hosts=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterTypeId := []string{"Inner_example"} // []string | Filter by typeId query param.  **Format:** filter.typeId={$not}:OPERATION:VALUE    **Example:** filter.typeId=$btw:John Doe&filter.typeId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPoolId := []string{"Inner_example"} // []string | Filter by poolId query param.  **Format:** filter.poolId={$not}:OPERATION:VALUE    **Example:** filter.poolId=$btw:John Doe&filter.poolId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAdministrationState := []string{"Inner_example"} // []string | Filter by administrationState query param.  **Format:** filter.administrationState={$not}:OPERATION:VALUE    **Example:** filter.administrationState=$btw:John Doe&filter.administrationState=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNumaNodes := []string{"Inner_example"} // []string | Filter by numaNodes query param.  **Format:** filter.numaNodes={$not}:OPERATION:VALUE    **Example:** filter.numaNodes=$btw:John Doe&filter.numaNodes=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$btw:John Doe&filter.infrastructureId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - host  - administrationState  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,host,hosts,administrationState   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - host  - hosts  - administrationState  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMAPI.GetVMs(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterName(filterName).FilterAddress(filterAddress).FilterHost(filterHost).FilterHosts(filterHosts).FilterTypeId(filterTypeId).FilterPoolId(filterPoolId).FilterAdministrationState(filterAdministrationState).FilterNumaNodes(filterNumaNodes).FilterInfrastructureId(filterInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.GetVMs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMs`: VMPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMAPI.GetVMs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAddress** | **[]string** | Filter by address query param.  **Format:** filter.address&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.address&#x3D;$btw:John Doe&amp;filter.address&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHost** | **[]string** | Filter by host query param.  **Format:** filter.host&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.host&#x3D;$btw:John Doe&amp;filter.host&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHosts** | **[]string** | Filter by hosts query param.  **Format:** filter.hosts&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.hosts&#x3D;$btw:John Doe&amp;filter.hosts&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterTypeId** | **[]string** | Filter by typeId query param.  **Format:** filter.typeId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.typeId&#x3D;$btw:John Doe&amp;filter.typeId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPoolId** | **[]string** | Filter by poolId query param.  **Format:** filter.poolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.poolId&#x3D;$btw:John Doe&amp;filter.poolId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAdministrationState** | **[]string** | Filter by administrationState query param.  **Format:** filter.administrationState&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.administrationState&#x3D;$btw:John Doe&amp;filter.administrationState&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNumaNodes** | **[]string** | Filter by numaNodes query param.  **Format:** filter.numaNodes&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.numaNodes&#x3D;$btw:John Doe&amp;filter.numaNodes&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$btw:John Doe&amp;filter.infrastructureId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - host  - administrationState  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,host,hosts,administrationState   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - host  - hosts  - administrationState  | 

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


## RebootVM

> RebootVM(ctx, vmId).Execute()

Reboots the VM



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
	vmId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMAPI.RebootVM(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.RebootVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootVMRequest struct via the builder pattern


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


## ShutdownVM

> ShutdownVM(ctx, vmId).Execute()

Shuts down the VM



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
	vmId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMAPI.ShutdownVM(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.ShutdownVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownVMRequest struct via the builder pattern


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


## StartVM

> StartVM(ctx, vmId).Execute()

Starts the VM



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
	vmId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMAPI.StartVM(context.Background(), vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.StartVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVMRequest struct via the builder pattern


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


## UpdateVM

> VM UpdateVM(ctx, vmId).UpdateVM(updateVM).Execute()

Updates VM information



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
	vmId := int64(789) // int64 | 
	updateVM := *openapiclient.NewUpdateVM() // UpdateVM | The VM update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMAPI.UpdateVM(context.Background(), vmId).UpdateVM(updateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMAPI.UpdateVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVM`: VM
	fmt.Fprintf(os.Stdout, "Response from `VMAPI.UpdateVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVM** | [**UpdateVM**](UpdateVM.md) | The VM update object | 

### Return type

[**VM**](VM.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

