# \NetworkDeviceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworkDeviceDefaults**](NetworkDeviceAPI.md#AddNetworkDeviceDefaults) | **Post** /api/v2/network-devices/defaults | Add network device defaults
[**ArchiveNetworkDevice**](NetworkDeviceAPI.md#ArchiveNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/archive | Archives a network device
[**CreateNetworkDevice**](NetworkDeviceAPI.md#CreateNetworkDevice) | **Post** /api/v2/network-devices | Create Network Device
[**DeleteNetworkDevice**](NetworkDeviceAPI.md#DeleteNetworkDevice) | **Delete** /api/v2/network-devices/{networkDeviceId} | Delete Network Device
[**DiscoverNetworkDevice**](NetworkDeviceAPI.md#DiscoverNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/discover | Discover network device interfaces, hardware and software configuration
[**EnableNetworkDeviceSyslog**](NetworkDeviceAPI.md#EnableNetworkDeviceSyslog) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/syslog-subscribe | Enables remote syslog for a network device
[**GetNetworkDevice**](NetworkDeviceAPI.md#GetNetworkDevice) | **Get** /api/v2/network-devices/{networkDeviceId} | Get Network Device
[**GetNetworkDeviceCredentials**](NetworkDeviceAPI.md#GetNetworkDeviceCredentials) | **Get** /api/v2/network-devices/{networkDeviceId}/credentials | Get Network Device credentials
[**GetNetworkDeviceDefaults**](NetworkDeviceAPI.md#GetNetworkDeviceDefaults) | **Get** /api/v2/network-devices/defaults/{siteId} | Get network device defaults for a site
[**GetNetworkDevicePorts**](NetworkDeviceAPI.md#GetNetworkDevicePorts) | **Get** /api/v2/network-devices/{networkDeviceId}/ports | Get paginated ports for a network device from the database
[**GetNetworkDeviceStatistics**](NetworkDeviceAPI.md#GetNetworkDeviceStatistics) | **Get** /api/v2/network-devices/statistics | Get Network Device Statistics
[**GetNetworkDevices**](NetworkDeviceAPI.md#GetNetworkDevices) | **Get** /api/v2/network-devices | Get paginated Network Devices
[**GetPorts**](NetworkDeviceAPI.md#GetPorts) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/ports | Port statistics for network device directly from the device
[**ReProvisionNetworkDevice**](NetworkDeviceAPI.md#ReProvisionNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/re-provision | Re-provision network device
[**RemoveNetworkDeviceDefaults**](NetworkDeviceAPI.md#RemoveNetworkDeviceDefaults) | **Delete** /api/v2/network-devices/defaults/{siteId}/{id} | Remove network device defaults
[**ReplaceNetworkDevice**](NetworkDeviceAPI.md#ReplaceNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/replace | Replace network device
[**ResetNetworkDevice**](NetworkDeviceAPI.md#ResetNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/reset | Resets a network device to default state
[**RevertNetworkDeviceFailedState**](NetworkDeviceAPI.md#RevertNetworkDeviceFailedState) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/revert-failed-state | Revert network device failed state
[**RunExtensionOnNetworkDevice**](NetworkDeviceAPI.md#RunExtensionOnNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/run-extension | Runs an extension of type action on the network device
[**SetNetworkDeviceAsFailed**](NetworkDeviceAPI.md#SetNetworkDeviceAsFailed) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-as-failed | Set network device as failed
[**SetNetworkDevicePortStatus**](NetworkDeviceAPI.md#SetNetworkDevicePortStatus) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-port-status | Set port status
[**UpdateNetworkDevice**](NetworkDeviceAPI.md#UpdateNetworkDevice) | **Patch** /api/v2/network-devices/{networkDeviceId} | Update Network Device



## AddNetworkDeviceDefaults

> AddNetworkDeviceDefaults(ctx).CreateNetworkDeviceDefaults(createNetworkDeviceDefaults).Execute()

Add network device defaults

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
	createNetworkDeviceDefaults := *openapiclient.NewCreateNetworkDeviceDefaults("DatacenterName_example", "00:1A:2B:3C:4D:5E") // CreateNetworkDeviceDefaults | Network device defaults

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.AddNetworkDeviceDefaults(context.Background()).CreateNetworkDeviceDefaults(createNetworkDeviceDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.AddNetworkDeviceDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkDeviceDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDeviceDefaults** | [**CreateNetworkDeviceDefaults**](CreateNetworkDeviceDefaults.md) | Network device defaults | 

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


## ArchiveNetworkDevice

> JobInfo ArchiveNetworkDevice(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

Archives a network device

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
	networkDeviceId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ArchiveNetworkDevice(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ArchiveNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveNetworkDevice`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ArchiveNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

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


## CreateNetworkDevice

> NetworkDevice CreateNetworkDevice(ctx).CreateNetworkDevice(createNetworkDevice).Execute()

Create Network Device

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
	createNetworkDevice := *openapiclient.NewCreateNetworkDevice(openapiclient.NetworkDeviceDriver("cisco_aci51"), "leaf", "admin", "ManagementPassword_example") // CreateNetworkDevice | The Network Device create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.CreateNetworkDevice(context.Background()).CreateNetworkDevice(createNetworkDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.CreateNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDevice`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.CreateNetworkDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDevice** | [**CreateNetworkDevice**](CreateNetworkDevice.md) | The Network Device create object | 

### Return type

[**NetworkDevice**](NetworkDevice.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDevice

> JobInfo DeleteNetworkDevice(ctx, networkDeviceId).Execute()

Delete Network Device

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
	networkDeviceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.DeleteNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DeleteNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkDevice`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.DeleteNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceRequest struct via the builder pattern


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


## DiscoverNetworkDevice

> DiscoverNetworkDevice(ctx, networkDeviceId).DiscoveryQuery(discoveryQuery).Execute()

Discover network device interfaces, hardware and software configuration



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
	networkDeviceId := float32(8.14) // float32 | Network device identifier
	discoveryQuery := *openapiclient.NewDiscoveryQuery([]string{"Discover_example"}, false) // DiscoveryQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DiscoverNetworkDevice(context.Background(), networkDeviceId).DiscoveryQuery(discoveryQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DiscoverNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **discoveryQuery** | [**DiscoveryQuery**](DiscoveryQuery.md) |  | 

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


## EnableNetworkDeviceSyslog

> EnableNetworkDeviceSyslog(ctx, networkDeviceId).Execute()

Enables remote syslog for a network device



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
	networkDeviceId := float32(8.14) // float32 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.EnableNetworkDeviceSyslog(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.EnableNetworkDeviceSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableNetworkDeviceSyslogRequest struct via the builder pattern


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


## GetNetworkDevice

> NetworkDevice GetNetworkDevice(ctx, networkDeviceId).Execute()

Get Network Device

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
	networkDeviceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevice`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDevice**](NetworkDevice.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceCredentials

> NetworkDeviceCredentials GetNetworkDeviceCredentials(ctx, networkDeviceId).Execute()

Get Network Device credentials



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
	networkDeviceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceCredentials(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceCredentials`: NetworkDeviceCredentials
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceCredentials**](NetworkDeviceCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceDefaults

> []NetworkDeviceDefaults GetNetworkDeviceDefaults(ctx, siteId).Execute()

Get network device defaults for a site

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
	siteId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceDefaults(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceDefaults`: []NetworkDeviceDefaults
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NetworkDeviceDefaults**](NetworkDeviceDefaults.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicePorts

> NetworkDeviceInterfacesPaginatedList GetNetworkDevicePorts(ctx, networkDeviceId).Page(page).Limit(limit).FilterInterfaceId(filterInterfaceId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterInterfaceName(filterInterfaceName).FilterServerId(filterServerId).FilterServerInterfaceId(filterServerInterfaceId).FilterDirtyBit(filterDirtyBit).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get paginated ports for a network device from the database

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
	networkDeviceId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterInterfaceId := []string{"Inner_example"} // []string | Filter by interfaceId query param.  **Format:** filter.interfaceId={$not}:OPERATION:VALUE    **Example:** filter.interfaceId=$btw:John Doe&filter.interfaceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNetworkDeviceId := []string{"Inner_example"} // []string | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId=$btw:John Doe&filter.networkDeviceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInterfaceName := []string{"Inner_example"} // []string | Filter by interfaceName query param.  **Format:** filter.interfaceName={$not}:OPERATION:VALUE    **Example:** filter.interfaceName=$btw:John Doe&filter.interfaceName=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$btw:John Doe&filter.serverId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerInterfaceId := []string{"Inner_example"} // []string | Filter by serverInterfaceId query param.  **Format:** filter.serverInterfaceId={$not}:OPERATION:VALUE    **Example:** filter.serverInterfaceId=$btw:John Doe&filter.serverInterfaceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDirtyBit := []string{"Inner_example"} // []string | Filter by dirtyBit query param.  **Format:** filter.dirtyBit={$not}:OPERATION:VALUE    **Example:** filter.dirtyBit=$btw:John Doe&filter.dirtyBit=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=interfaceId:DESC&sortBy=networkDeviceId:DESC   **Default Value:** interfaceId:ASC  **Available Fields** - interfaceId  - networkDeviceId  - serverId  - cachedUpdatedTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** interfaceName,macAddress,lldpInformation   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - interfaceName  - macAddress  - lldpInformation  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevicePorts(context.Background(), networkDeviceId).Page(page).Limit(limit).FilterInterfaceId(filterInterfaceId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterInterfaceName(filterInterfaceName).FilterServerId(filterServerId).FilterServerInterfaceId(filterServerInterfaceId).FilterDirtyBit(filterDirtyBit).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevicePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevicePorts`: NetworkDeviceInterfacesPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevicePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterInterfaceId** | **[]string** | Filter by interfaceId query param.  **Format:** filter.interfaceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.interfaceId&#x3D;$btw:John Doe&amp;filter.interfaceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNetworkDeviceId** | **[]string** | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId&#x3D;$btw:John Doe&amp;filter.networkDeviceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInterfaceName** | **[]string** | Filter by interfaceName query param.  **Format:** filter.interfaceName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.interfaceName&#x3D;$btw:John Doe&amp;filter.interfaceName&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$btw:John Doe&amp;filter.serverId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerInterfaceId** | **[]string** | Filter by serverInterfaceId query param.  **Format:** filter.serverInterfaceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverInterfaceId&#x3D;$btw:John Doe&amp;filter.serverInterfaceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDirtyBit** | **[]string** | Filter by dirtyBit query param.  **Format:** filter.dirtyBit&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.dirtyBit&#x3D;$btw:John Doe&amp;filter.dirtyBit&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;interfaceId:DESC&amp;sortBy&#x3D;networkDeviceId:DESC   **Default Value:** interfaceId:ASC  **Available Fields** - interfaceId  - networkDeviceId  - serverId  - cachedUpdatedTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** interfaceName,macAddress,lldpInformation   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - interfaceName  - macAddress  - lldpInformation  | 

### Return type

[**NetworkDeviceInterfacesPaginatedList**](NetworkDeviceInterfacesPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceStatistics

> NetworkDeviceStatistics GetNetworkDeviceStatistics(ctx).Execute()

Get Network Device Statistics

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
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceStatistics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceStatistics`: NetworkDeviceStatistics
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceStatisticsRequest struct via the builder pattern


### Return type

[**NetworkDeviceStatistics**](NetworkDeviceStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevices

> NetworkDevicePaginatedList GetNetworkDevices(ctx).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get paginated Network Devices

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
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDatacenterName := []string{"Inner_example"} // []string | Filter by datacenterName query param.  **Format:** filter.datacenterName={$not}:OPERATION:VALUE    **Example:** filter.datacenterName=$btw:John Doe&filter.datacenterName=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterChassisIdentifier := []string{"Inner_example"} // []string | Filter by chassisIdentifier query param.  **Format:** filter.chassisIdentifier={$not}:OPERATION:VALUE    **Example:** filter.chassisIdentifier=$btw:John Doe&filter.chassisIdentifier=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementAddress := []string{"Inner_example"} // []string | Filter by managementAddress query param.  **Format:** filter.managementAddress={$not}:OPERATION:VALUE    **Example:** filter.managementAddress=$btw:John Doe&filter.managementAddress=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementPort := []string{"Inner_example"} // []string | Filter by managementPort query param.  **Format:** filter.managementPort={$not}:OPERATION:VALUE    **Example:** filter.managementPort=$btw:John Doe&filter.managementPort=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterProvisionerType := []string{"Inner_example"} // []string | Filter by provisionerType query param.  **Format:** filter.provisionerType={$not}:OPERATION:VALUE    **Example:** filter.provisionerType=$btw:John Doe&filter.provisionerType=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPosition := []string{"Inner_example"} // []string | Filter by position query param.  **Format:** filter.position={$not}:OPERATION:VALUE    **Example:** filter.position=$btw:John Doe&filter.position=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterIdentifierString := []string{"Inner_example"} // []string | Filter by identifierString query param.  **Format:** filter.identifierString={$not}:OPERATION:VALUE    **Example:** filter.identifierString=$btw:John Doe&filter.identifierString=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=status:DESC   **Default Value:** id:ASC  **Available Fields** - id  - status  - siteId  - status  - position  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,status,siteId,managementAddress,position   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - status  - siteId  - managementAddress  - position  - identifierString  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevices(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevices`: NetworkDevicePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDatacenterName** | **[]string** | Filter by datacenterName query param.  **Format:** filter.datacenterName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.datacenterName&#x3D;$btw:John Doe&amp;filter.datacenterName&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterChassisIdentifier** | **[]string** | Filter by chassisIdentifier query param.  **Format:** filter.chassisIdentifier&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.chassisIdentifier&#x3D;$btw:John Doe&amp;filter.chassisIdentifier&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementAddress** | **[]string** | Filter by managementAddress query param.  **Format:** filter.managementAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementAddress&#x3D;$btw:John Doe&amp;filter.managementAddress&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementPort** | **[]string** | Filter by managementPort query param.  **Format:** filter.managementPort&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementPort&#x3D;$btw:John Doe&amp;filter.managementPort&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterProvisionerType** | **[]string** | Filter by provisionerType query param.  **Format:** filter.provisionerType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.provisionerType&#x3D;$btw:John Doe&amp;filter.provisionerType&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPosition** | **[]string** | Filter by position query param.  **Format:** filter.position&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.position&#x3D;$btw:John Doe&amp;filter.position&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterIdentifierString** | **[]string** | Filter by identifierString query param.  **Format:** filter.identifierString&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.identifierString&#x3D;$btw:John Doe&amp;filter.identifierString&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;status:DESC   **Default Value:** id:ASC  **Available Fields** - id  - status  - siteId  - status  - position  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,status,siteId,managementAddress,position   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - status  - siteId  - managementAddress  - position  - identifierString  | 

### Return type

[**NetworkDevicePaginatedList**](NetworkDevicePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPorts

> NetworkDevicePorts GetPorts(ctx, networkDeviceId).Execute()

Port statistics for network device directly from the device

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
	networkDeviceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetPorts(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPorts`: NetworkDevicePorts
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDevicePorts**](NetworkDevicePorts.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReProvisionNetworkDevice

> JobInfo ReProvisionNetworkDevice(ctx, networkDeviceId).NetworkEquipmentReprovision(networkEquipmentReprovision).Execute()

Re-provision network device



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
	networkDeviceId := float32(8.14) // float32 | 
	networkEquipmentReprovision := *openapiclient.NewNetworkEquipmentReprovision("ReprovisionType_example") // NetworkEquipmentReprovision | The network device re-provision options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ReProvisionNetworkDevice(context.Background(), networkDeviceId).NetworkEquipmentReprovision(networkEquipmentReprovision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ReProvisionNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReProvisionNetworkDevice`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ReProvisionNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReProvisionNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkEquipmentReprovision** | [**NetworkEquipmentReprovision**](NetworkEquipmentReprovision.md) | The network device re-provision options | 

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNetworkDeviceDefaults

> RemoveNetworkDeviceDefaults(ctx, siteId, id).Execute()

Remove network device defaults

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
	siteId := float32(8.14) // float32 | 
	id := float32(8.14) // float32 | The ID of the network device default to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.RemoveNetworkDeviceDefaults(context.Background(), siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.RemoveNetworkDeviceDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **float32** |  | 
**id** | **float32** | The ID of the network device default to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkDeviceDefaultsRequest struct via the builder pattern


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


## ReplaceNetworkDevice

> map[string]interface{} ReplaceNetworkDevice(ctx, networkDeviceId).SwitchReplace(switchReplace).Execute()

Replace network device

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
	networkDeviceId := float32(8.14) // float32 | Network device ID
	switchReplace := *openapiclient.NewSwitchReplace(float32(123)) // SwitchReplace | Network device replacement details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ReplaceNetworkDevice(context.Background(), networkDeviceId).SwitchReplace(switchReplace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ReplaceNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceNetworkDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ReplaceNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **switchReplace** | [**SwitchReplace**](SwitchReplace.md) | Network device replacement details | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetNetworkDevice

> ResetNetworkDevice(ctx, networkDeviceId).Execute()

Resets a network device to default state



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
	networkDeviceId := float32(8.14) // float32 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.ResetNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ResetNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetNetworkDeviceRequest struct via the builder pattern


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


## RevertNetworkDeviceFailedState

> NetworkDevice RevertNetworkDeviceFailedState(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

Revert network device failed state



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
	networkDeviceId := float32(8.14) // float32 | Network device ID
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.RevertNetworkDeviceFailedState(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.RevertNetworkDeviceFailedState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertNetworkDeviceFailedState`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.RevertNetworkDeviceFailedState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertNetworkDeviceFailedStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkDevice**](NetworkDevice.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunExtensionOnNetworkDevice

> JobInfo RunExtensionOnNetworkDevice(ctx, networkDeviceId).RunExtensionOnPhysicalDevice(runExtensionOnPhysicalDevice).IfMatch(ifMatch).Execute()

Runs an extension of type action on the network device

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
	networkDeviceId := float32(8.14) // float32 | 
	runExtensionOnPhysicalDevice := *openapiclient.NewRunExtensionOnPhysicalDevice(float32(10), map[string]interface{}(123)) // RunExtensionOnPhysicalDevice | The extension information
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.RunExtensionOnNetworkDevice(context.Background(), networkDeviceId).RunExtensionOnPhysicalDevice(runExtensionOnPhysicalDevice).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.RunExtensionOnNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunExtensionOnNetworkDevice`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.RunExtensionOnNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunExtensionOnNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runExtensionOnPhysicalDevice** | [**RunExtensionOnPhysicalDevice**](RunExtensionOnPhysicalDevice.md) | The extension information | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNetworkDeviceAsFailed

> NetworkDevice SetNetworkDeviceAsFailed(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

Set network device as failed



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
	networkDeviceId := float32(8.14) // float32 | Network device ID
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.SetNetworkDeviceAsFailed(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.SetNetworkDeviceAsFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNetworkDeviceAsFailed`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.SetNetworkDeviceAsFailed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNetworkDeviceAsFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkDevice**](NetworkDevice.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNetworkDevicePortStatus

> SetNetworkDevicePortStatus(ctx, networkDeviceId).NetworkDevicePortStatus(networkDevicePortStatus).Execute()

Set port status

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
	networkDeviceId := float32(8.14) // float32 | Network device ID
	networkDevicePortStatus := *openapiclient.NewNetworkDevicePortStatus([]string{"Ports_example"}, true) // NetworkDevicePortStatus | Port status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.SetNetworkDevicePortStatus(context.Background(), networkDeviceId).NetworkDevicePortStatus(networkDevicePortStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.SetNetworkDevicePortStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNetworkDevicePortStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDevicePortStatus** | [**NetworkDevicePortStatus**](NetworkDevicePortStatus.md) | Port status | 

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


## UpdateNetworkDevice

> NetworkDevice UpdateNetworkDevice(ctx, networkDeviceId).UpdateNetworkDevice(updateNetworkDevice).IfMatch(ifMatch).Execute()

Update Network Device

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
	networkDeviceId := float32(8.14) // float32 | 
	updateNetworkDevice := *openapiclient.NewUpdateNetworkDevice() // UpdateNetworkDevice | The Network Device update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.UpdateNetworkDevice(context.Background(), networkDeviceId).UpdateNetworkDevice(updateNetworkDevice).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.UpdateNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDevice`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.UpdateNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDevice** | [**UpdateNetworkDevice**](UpdateNetworkDevice.md) | The Network Device update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkDevice**](NetworkDevice.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

