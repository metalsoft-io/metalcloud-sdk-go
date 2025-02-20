# \NetworkDeviceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworkDeviceDefaults**](NetworkDeviceAPI.md#AddNetworkDeviceDefaults) | **Post** /api/v2/network-devices/defaults | Add network device defaults
[**ChangeNetworkDeviceStatus**](NetworkDeviceAPI.md#ChangeNetworkDeviceStatus) | **Patch** /api/v2/network-devices/{networkDeviceId}/actions/change-status | Change status of a network device
[**DiscoverNetworkDevice**](NetworkDeviceAPI.md#DiscoverNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/discover | Discover network device interfaces, hardware and software configuration
[**EnableNetworkDeviceSyslog**](NetworkDeviceAPI.md#EnableNetworkDeviceSyslog) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/syslog-subscribe | Enables remote syslog for a network device
[**GetNetworkDevice**](NetworkDeviceAPI.md#GetNetworkDevice) | **Get** /api/v2/network-devices/{networkDeviceId} | Get Network Device
[**GetNetworkDeviceDefaults**](NetworkDeviceAPI.md#GetNetworkDeviceDefaults) | **Get** /api/v2/network-devices/defaults/{datacenterName} | Get network device defaults for a datacenter
[**GetNetworkDevicePorts**](NetworkDeviceAPI.md#GetNetworkDevicePorts) | **Get** /api/v2/network-devices/{networkDeviceId}/ports | Get all ports for network device
[**GetNetworkDevices**](NetworkDeviceAPI.md#GetNetworkDevices) | **Get** /api/v2/network-devices | Get paginated Network Devices
[**NetworkDeviceControllerAddTagsToNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerAddTagsToNetworkDevice) | **Post** /api/v2/network-devices/network-devices/{networkDeviceId}/tags | Add tags to Network Device
[**NetworkDeviceControllerCreateNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerCreateNetworkDevice) | **Post** /api/v2/network-devices | Create Network Device
[**NetworkDeviceControllerDecommissionNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerDecommissionNetworkDevice) | **Delete** /api/v2/network-devices/{networkDeviceId}/decommission | Decommission network device
[**NetworkDeviceControllerDeleteNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerDeleteNetworkDevice) | **Delete** /api/v2/network-devices/{networkDeviceId} | Delete Network Device
[**NetworkDeviceControllerGetIscsiBootServers**](NetworkDeviceAPI.md#NetworkDeviceControllerGetIscsiBootServers) | **Get** /api/v2/network-devices/{networkDeviceId}/iscsi-boot-servers | Returns information about servers which are setup to boot from iSCSI block devices. This is useful in the event of a switch device reboot
[**NetworkDeviceControllerGetNetworkDeviceStatistics**](NetworkDeviceAPI.md#NetworkDeviceControllerGetNetworkDeviceStatistics) | **Get** /api/v2/network-devices/statistics | Get Network Device Statistics
[**NetworkDeviceControllerGetTagsForNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerGetTagsForNetworkDevice) | **Get** /api/v2/network-devices/network-devices/{networkDeviceId}/tags | Get tags for Network Device
[**NetworkDeviceControllerRemoveTagsFromNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerRemoveTagsFromNetworkDevice) | **Patch** /api/v2/network-devices/network-devices/{networkDeviceId}/tags | Clear tags from Network Device
[**NetworkDeviceControllerReplaceNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerReplaceNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/replace | Replace network device
[**NetworkDeviceControllerUpdateNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerUpdateNetworkDevice) | **Patch** /api/v2/network-devices/{networkDeviceId} | Update Network Device
[**NetworkDeviceControllerUpdateTagsForNetworkDevice**](NetworkDeviceAPI.md#NetworkDeviceControllerUpdateTagsForNetworkDevice) | **Put** /api/v2/network-devices/network-devices/{networkDeviceId}/tags | Update tags for Network Device
[**ReProvisionNetworkEquipment**](NetworkDeviceAPI.md#ReProvisionNetworkEquipment) | **Post** /api/v2/network-devices/re-provision | Re-provision network equipment
[**RemoveNetworkDeviceDefaults**](NetworkDeviceAPI.md#RemoveNetworkDeviceDefaults) | **Delete** /api/v2/network-devices/defaults | Remove network device defaults
[**ResetNetworkDevice**](NetworkDeviceAPI.md#ResetNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/reset | Resets a network device to default state
[**SetNetworkDevicePortStatus**](NetworkDeviceAPI.md#SetNetworkDevicePortStatus) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-port-status | Set port status



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


## ChangeNetworkDeviceStatus

> ChangeNetworkDeviceStatus(ctx, networkDeviceId).NetworkDeviceStatus(networkDeviceStatus).Execute()

Change status of a network device



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
	networkDeviceStatus := *openapiclient.NewNetworkDeviceStatus("active") // NetworkDeviceStatus | Network device status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.ChangeNetworkDeviceStatus(context.Background(), networkDeviceId).NetworkDeviceStatus(networkDeviceStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ChangeNetworkDeviceStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChangeNetworkDeviceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDeviceStatus** | [**NetworkDeviceStatus**](NetworkDeviceStatus.md) | Network device status | 

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


## DiscoverNetworkDevice

> DiscoverNetworkDevice(ctx, networkDeviceId).Body(body).Execute()

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DiscoverNetworkDevice(context.Background(), networkDeviceId).Body(body).Execute()
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

 **body** | **map[string]interface{}** |  | 

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


## GetNetworkDeviceDefaults

> []map[string]interface{} GetNetworkDeviceDefaults(ctx, datacenterName).Execute()

Get network device defaults for a datacenter

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
	datacenterName := "datacenterName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceDefaults(context.Background(), datacenterName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceDefaults`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicePorts

> GetNetworkDevicePorts(ctx, networkDeviceId).Execute()

Get all ports for network device

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
	r, err := apiClient.NetworkDeviceAPI.GetNetworkDevicePorts(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevicePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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


## GetNetworkDevices

> NetworkDevicePaginatedList GetNetworkDevices(ctx).Page(page).Limit(limit).FilterSwitchId(filterSwitchId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterSwitchId := []string{"Inner_example"} // []string | Filter by switchId query param.           <p>              <b>Format: </b> filter.switchId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.switchId=$not:$like:John Doe&filter.switchId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.           <p>              <b>Format: </b> filter.status={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.status=$not:$like:John Doe&filter.status=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterDatacenterName := []string{"Inner_example"} // []string | Filter by datacenterName query param.           <p>              <b>Format: </b> filter.datacenterName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.datacenterName=$not:$like:John Doe&filter.datacenterName=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterChassisIdentifier := []string{"Inner_example"} // []string | Filter by chassisIdentifier query param.           <p>              <b>Format: </b> filter.chassisIdentifier={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.chassisIdentifier=$not:$like:John Doe&filter.chassisIdentifier=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterManagementAddress := []string{"Inner_example"} // []string | Filter by managementAddress query param.           <p>              <b>Format: </b> filter.managementAddress={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.managementAddress=$not:$like:John Doe&filter.managementAddress=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterManagementPort := []string{"Inner_example"} // []string | Filter by managementPort query param.           <p>              <b>Format: </b> filter.managementPort={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.managementPort=$not:$like:John Doe&filter.managementPort=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterProvisionerType := []string{"Inner_example"} // []string | Filter by provisionerType query param.           <p>              <b>Format: </b> filter.provisionerType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.provisionerType=$not:$like:John Doe&filter.provisionerType=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterPosition := []string{"Inner_example"} // []string | Filter by position query param.           <p>              <b>Format: </b> filter.position={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.position=$not:$like:John Doe&filter.position=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterIdentifierString := []string{"Inner_example"} // []string | Filter by identifierString query param.           <p>              <b>Format: </b> filter.identifierString={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.identifierString=$not:$like:John Doe&filter.identifierString=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> switchId:ASC           </p>       <h4>Available Fields</h4><ul><li>switchId</li> <li>status</li> <li>datacenterName</li> <li>status</li> <li>position</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> switchId,status,datacenterName,managementAddress,provisionerType           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>switchId</li> <li>status</li> <li>datacenterName</li> <li>managementAddress</li> <li>provisionerType</li> <li>position</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevices(context.Background()).Page(page).Limit(limit).FilterSwitchId(filterSwitchId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterSwitchId** | **[]string** | Filter by switchId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.switchId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.switchId&#x3D;$not:$like:John Doe&amp;filter.switchId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterStatus** | **[]string** | Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterDatacenterName** | **[]string** | Filter by datacenterName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.datacenterName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.datacenterName&#x3D;$not:$like:John Doe&amp;filter.datacenterName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterChassisIdentifier** | **[]string** | Filter by chassisIdentifier query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.chassisIdentifier&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.chassisIdentifier&#x3D;$not:$like:John Doe&amp;filter.chassisIdentifier&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterManagementAddress** | **[]string** | Filter by managementAddress query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.managementAddress&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.managementAddress&#x3D;$not:$like:John Doe&amp;filter.managementAddress&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterManagementPort** | **[]string** | Filter by managementPort query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.managementPort&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.managementPort&#x3D;$not:$like:John Doe&amp;filter.managementPort&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterProvisionerType** | **[]string** | Filter by provisionerType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.provisionerType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.provisionerType&#x3D;$not:$like:John Doe&amp;filter.provisionerType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterPosition** | **[]string** | Filter by position query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.position&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.position&#x3D;$not:$like:John Doe&amp;filter.position&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterIdentifierString** | **[]string** | Filter by identifierString query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.identifierString&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.identifierString&#x3D;$not:$like:John Doe&amp;filter.identifierString&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; switchId:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;switchId&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;datacenterName&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;position&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; switchId,status,datacenterName,managementAddress,provisionerType           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;switchId&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;datacenterName&lt;/li&gt; &lt;li&gt;managementAddress&lt;/li&gt; &lt;li&gt;provisionerType&lt;/li&gt; &lt;li&gt;position&lt;/li&gt;&lt;/ul&gt;          | 

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


## NetworkDeviceControllerAddTagsToNetworkDevice

> map[string]interface{} NetworkDeviceControllerAddTagsToNetworkDevice(ctx, networkDeviceId).NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest(networkDeviceControllerUpdateTagsForNetworkDeviceRequest).Execute()

Add tags to Network Device

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
	networkDeviceControllerUpdateTagsForNetworkDeviceRequest := *openapiclient.NewNetworkDeviceControllerUpdateTagsForNetworkDeviceRequest() // NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest | The tags to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerAddTagsToNetworkDevice(context.Background(), networkDeviceId).NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest(networkDeviceControllerUpdateTagsForNetworkDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerAddTagsToNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkDeviceControllerAddTagsToNetworkDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.NetworkDeviceControllerAddTagsToNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerAddTagsToNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDeviceControllerUpdateTagsForNetworkDeviceRequest** | [**NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest**](NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest.md) | The tags to add | 

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


## NetworkDeviceControllerCreateNetworkDevice

> NetworkDevice NetworkDeviceControllerCreateNetworkDevice(ctx).CreateNetworkDevice(createNetworkDevice).Execute()

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
	createNetworkDevice := *openapiclient.NewCreateNetworkDevice("DatacenterName_example", "ProvisionerType_example", "Driver_example", "Position_example") // CreateNetworkDevice | The Network Device create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerCreateNetworkDevice(context.Background()).CreateNetworkDevice(createNetworkDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerCreateNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkDeviceControllerCreateNetworkDevice`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.NetworkDeviceControllerCreateNetworkDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerCreateNetworkDeviceRequest struct via the builder pattern


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


## NetworkDeviceControllerDecommissionNetworkDevice

> NetworkDeviceControllerDecommissionNetworkDevice(ctx, networkDeviceId).Execute()

Decommission network device

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
	r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerDecommissionNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerDecommissionNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerDecommissionNetworkDeviceRequest struct via the builder pattern


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


## NetworkDeviceControllerDeleteNetworkDevice

> NetworkDeviceControllerDeleteNetworkDevice(ctx, networkDeviceId).Execute()

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
	r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerDeleteNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerDeleteNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerDeleteNetworkDeviceRequest struct via the builder pattern


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


## NetworkDeviceControllerGetIscsiBootServers

> NetworkDeviceControllerGetIscsiBootServers(ctx, networkDeviceId).Execute()

Returns information about servers which are setup to boot from iSCSI block devices. This is useful in the event of a switch device reboot

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
	r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerGetIscsiBootServers(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerGetIscsiBootServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerGetIscsiBootServersRequest struct via the builder pattern


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


## NetworkDeviceControllerGetNetworkDeviceStatistics

> NetworkDeviceControllerGetNetworkDeviceStatistics(ctx).Execute()

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
	r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerGetNetworkDeviceStatistics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerGetNetworkDeviceStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerGetNetworkDeviceStatisticsRequest struct via the builder pattern


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


## NetworkDeviceControllerGetTagsForNetworkDevice

> []string NetworkDeviceControllerGetTagsForNetworkDevice(ctx, networkDeviceId).Execute()

Get tags for Network Device

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
	resp, r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerGetTagsForNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerGetTagsForNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkDeviceControllerGetTagsForNetworkDevice`: []string
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.NetworkDeviceControllerGetTagsForNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerGetTagsForNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDeviceControllerRemoveTagsFromNetworkDevice

> NetworkDeviceControllerRemoveTagsFromNetworkDevice(ctx, networkDeviceId).Body(body).Execute()

Clear tags from Network Device

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
	body := "body_example" // string | The tags to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerRemoveTagsFromNetworkDevice(context.Background(), networkDeviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerRemoveTagsFromNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerRemoveTagsFromNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The tags to remove | 

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


## NetworkDeviceControllerReplaceNetworkDevice

> map[string]interface{} NetworkDeviceControllerReplaceNetworkDevice(ctx, networkDeviceId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerReplaceNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerReplaceNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkDeviceControllerReplaceNetworkDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.NetworkDeviceControllerReplaceNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerReplaceNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDeviceControllerUpdateNetworkDevice

> NetworkDevice NetworkDeviceControllerUpdateNetworkDevice(ctx, networkDeviceId).UpdateNetworkDevice(updateNetworkDevice).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerUpdateNetworkDevice(context.Background(), networkDeviceId).UpdateNetworkDevice(updateNetworkDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerUpdateNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkDeviceControllerUpdateNetworkDevice`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.NetworkDeviceControllerUpdateNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerUpdateNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDevice** | [**UpdateNetworkDevice**](UpdateNetworkDevice.md) | The Network Device update object | 

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


## NetworkDeviceControllerUpdateTagsForNetworkDevice

> map[string]interface{} NetworkDeviceControllerUpdateTagsForNetworkDevice(ctx, networkDeviceId).NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest(networkDeviceControllerUpdateTagsForNetworkDeviceRequest).Execute()

Update tags for Network Device

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
	networkDeviceControllerUpdateTagsForNetworkDeviceRequest := *openapiclient.NewNetworkDeviceControllerUpdateTagsForNetworkDeviceRequest() // NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest | The tags to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.NetworkDeviceControllerUpdateTagsForNetworkDevice(context.Background(), networkDeviceId).NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest(networkDeviceControllerUpdateTagsForNetworkDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceControllerUpdateTagsForNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkDeviceControllerUpdateTagsForNetworkDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.NetworkDeviceControllerUpdateTagsForNetworkDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDeviceControllerUpdateTagsForNetworkDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDeviceControllerUpdateTagsForNetworkDeviceRequest** | [**NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest**](NetworkDeviceControllerUpdateTagsForNetworkDeviceRequest.md) | The tags to update | 

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


## ReProvisionNetworkEquipment

> JobInfo ReProvisionNetworkEquipment(ctx).Body(body).Execute()

Re-provision network equipment



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
	body := map[string]interface{}{ ... } // map[string]interface{} | The network equipment re-provision options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ReProvisionNetworkEquipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ReProvisionNetworkEquipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReProvisionNetworkEquipment`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ReProvisionNetworkEquipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReProvisionNetworkEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The network equipment re-provision options | 

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

> RemoveNetworkDeviceDefaults(ctx).RequestBody(requestBody).Execute()

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
	requestBody := []float32{float32(123)} // []float32 | Network device defaults IDs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.RemoveNetworkDeviceDefaults(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.RemoveNetworkDeviceDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkDeviceDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]float32** | Network device defaults IDs | 

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

