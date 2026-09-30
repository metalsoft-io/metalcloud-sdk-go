# \NetworkDeviceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcknowledgeNetworkDeviceDrift**](NetworkDeviceAPI.md#AcknowledgeNetworkDeviceDrift) | **Post** /api/v2/network-devices/{networkDeviceId}/drift-history/{driftId}/actions/acknowledge | Acknowledge a network device drift
[**AddNetworkDevicePortIp**](NetworkDeviceAPI.md#AddNetworkDevicePortIp) | **Post** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/{family}/addresses | Stage adding an IP address of an address family to a port
[**ArchiveNetworkDevice**](NetworkDeviceAPI.md#ArchiveNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/archive | Archives a network device
[**CreateNetworkDevice**](NetworkDeviceAPI.md#CreateNetworkDevice) | **Post** /api/v2/network-devices | Create Network Device
[**CreateNetworkDeviceBreakout**](NetworkDeviceAPI.md#CreateNetworkDeviceBreakout) | **Post** /api/v2/network-devices/{networkDeviceId}/breakouts | Create a port breakout (Cumulus only)
[**CreateNetworkDevicePort**](NetworkDeviceAPI.md#CreateNetworkDevicePort) | **Post** /api/v2/network-devices/{networkDeviceId}/ports | Stage creation of a logical port on a network device (not supported yet — returns 501)
[**DeleteNetworkDevice**](NetworkDeviceAPI.md#DeleteNetworkDevice) | **Delete** /api/v2/network-devices/{networkDeviceId} | Delete Network Device
[**DeleteNetworkDeviceBreakout**](NetworkDeviceAPI.md#DeleteNetworkDeviceBreakout) | **Delete** /api/v2/network-devices/{networkDeviceId}/breakouts/{breakoutId} | Remove a port breakout (un-breakout)
[**DeleteNetworkDevicePort**](NetworkDeviceAPI.md#DeleteNetworkDevicePort) | **Delete** /api/v2/network-devices/{networkDeviceId}/ports/{portId} | Stage removal of a port
[**DeleteNetworkDeviceSecret**](NetworkDeviceAPI.md#DeleteNetworkDeviceSecret) | **Delete** /api/v2/network-devices/{networkDeviceId}/secrets/{name} | Delete one Network Device secret
[**DeleteNetworkDeviceSecrets**](NetworkDeviceAPI.md#DeleteNetworkDeviceSecrets) | **Delete** /api/v2/network-devices/{networkDeviceId}/secrets | Delete all secrets of a Network Device
[**DisableNetworkDeviceSnmpMonitoring**](NetworkDeviceAPI.md#DisableNetworkDeviceSnmpMonitoring) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/snmp-monitoring-unsubscribe | Disables SNMP monitoring for a network device
[**DisableNetworkDeviceSnmpMonitoringBatch**](NetworkDeviceAPI.md#DisableNetworkDeviceSnmpMonitoringBatch) | **Post** /api/v2/network-devices/actions/snmp-monitoring-unsubscribe/batch | Disables SNMP monitoring for a batch of network devices
[**DisableNetworkDeviceSnmpService**](NetworkDeviceAPI.md#DisableNetworkDeviceSnmpService) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/snmp-service-disable | Disables SNMP service on a network device
[**DisableNetworkDeviceSyslog**](NetworkDeviceAPI.md#DisableNetworkDeviceSyslog) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/syslog-unsubscribe | Disables remote syslog for a network device
[**DiscoverNetworkDevice**](NetworkDeviceAPI.md#DiscoverNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/discover | Discover network device interfaces, hardware and software configuration
[**EnableNetworkDeviceSnmpMonitoring**](NetworkDeviceAPI.md#EnableNetworkDeviceSnmpMonitoring) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/snmp-monitoring-subscribe | Enables SNMP monitoring for a network device
[**EnableNetworkDeviceSnmpMonitoringBatch**](NetworkDeviceAPI.md#EnableNetworkDeviceSnmpMonitoringBatch) | **Post** /api/v2/network-devices/actions/snmp-monitoring-subscribe/batch | Enables SNMP monitoring for a batch of network devices
[**EnableNetworkDeviceSnmpService**](NetworkDeviceAPI.md#EnableNetworkDeviceSnmpService) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/snmp-service-enable | Enables SNMP service on a network device
[**EnableNetworkDeviceSyslog**](NetworkDeviceAPI.md#EnableNetworkDeviceSyslog) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/syslog-subscribe | Enables remote syslog for a network device
[**GetNetworkDevice**](NetworkDeviceAPI.md#GetNetworkDevice) | **Get** /api/v2/network-devices/{networkDeviceId} | Get Network Device
[**GetNetworkDeviceBreakout**](NetworkDeviceAPI.md#GetNetworkDeviceBreakout) | **Get** /api/v2/network-devices/{networkDeviceId}/breakouts/{breakoutId} | Get a port breakout by id
[**GetNetworkDeviceBreakoutConfig**](NetworkDeviceAPI.md#GetNetworkDeviceBreakoutConfig) | **Get** /api/v2/network-devices/{networkDeviceId}/breakouts/{breakoutId}/config | Get the staged (desired) breakout groups for a port breakout
[**GetNetworkDeviceCredentials**](NetworkDeviceAPI.md#GetNetworkDeviceCredentials) | **Get** /api/v2/network-devices/{networkDeviceId}/credentials | Get Network Device credentials
[**GetNetworkDeviceDriftHistory**](NetworkDeviceAPI.md#GetNetworkDeviceDriftHistory) | **Get** /api/v2/network-devices/{networkDeviceId}/drift-history | Get drift history for a specific network device
[**GetNetworkDeviceDriftHistoryById**](NetworkDeviceAPI.md#GetNetworkDeviceDriftHistoryById) | **Get** /api/v2/network-devices/{networkDeviceId}/drift-history/{driftId} | Get a specific drift history for a specific network device
[**GetNetworkDeviceDriverCapabilities**](NetworkDeviceAPI.md#GetNetworkDeviceDriverCapabilities) | **Get** /api/v2/network-devices/drivers/capabilities | Get Network Device driver capabilities
[**GetNetworkDeviceHealthSummary**](NetworkDeviceAPI.md#GetNetworkDeviceHealthSummary) | **Get** /api/v2/network-devices/{networkDeviceId}/health-summary | Get Network Device health summary
[**GetNetworkDevicePort**](NetworkDeviceAPI.md#GetNetworkDevicePort) | **Get** /api/v2/network-devices/{networkDeviceId}/ports/{portId} | Get one port on a network device
[**GetNetworkDevicePortConfig**](NetworkDeviceAPI.md#GetNetworkDevicePortConfig) | **Get** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/config | Get the desired-state config for a port
[**GetNetworkDevicePortIp**](NetworkDeviceAPI.md#GetNetworkDevicePortIp) | **Get** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/{family}/addresses/{ipId} | Get one IP-address aspect row on a port
[**GetNetworkDevicePortVirtualFunction**](NetworkDeviceAPI.md#GetNetworkDevicePortVirtualFunction) | **Get** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/virtual-functions/{virtualFunctionId} | Get a specific virtual function for a specific port on a network device
[**GetNetworkDevicePortVirtualFunctions**](NetworkDeviceAPI.md#GetNetworkDevicePortVirtualFunctions) | **Get** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/virtual-functions | Get virtual functions for a specific port on a network device
[**GetNetworkDevicePorts**](NetworkDeviceAPI.md#GetNetworkDevicePorts) | **Get** /api/v2/network-devices/{networkDeviceId}/ports | Get paginated ports for a network device from the database
[**GetNetworkDeviceSNMPMonitoringAgentInfoBatch**](NetworkDeviceAPI.md#GetNetworkDeviceSNMPMonitoringAgentInfoBatch) | **Get** /api/v2/network-devices/snmp-monitoring/agent-info/batch | Get Network Device SNMP Monitoring Agent Info Batch
[**GetNetworkDeviceSecretCredentials**](NetworkDeviceAPI.md#GetNetworkDeviceSecretCredentials) | **Get** /api/v2/network-devices/{networkDeviceId}/secrets/{name}/credentials | Get one Network Device secret unencrypted
[**GetNetworkDeviceSecrets**](NetworkDeviceAPI.md#GetNetworkDeviceSecrets) | **Get** /api/v2/network-devices/{networkDeviceId}/secrets | List the names of the secrets set on a Network Device
[**GetNetworkDeviceSnapshots**](NetworkDeviceAPI.md#GetNetworkDeviceSnapshots) | **Get** /api/v2/network-devices/{networkDeviceId}/snapshots | Get Network Device snapshots
[**GetNetworkDeviceStatistics**](NetworkDeviceAPI.md#GetNetworkDeviceStatistics) | **Get** /api/v2/network-devices/statistics | Get Network Device Statistics
[**GetNetworkDeviceVendor**](NetworkDeviceAPI.md#GetNetworkDeviceVendor) | **Get** /api/v2/network-devices/vendors/{vendorId} | Get Network Device Vendor
[**GetNetworkDeviceVendors**](NetworkDeviceAPI.md#GetNetworkDeviceVendors) | **Get** /api/v2/network-devices/vendors | Get paginated Network Device Vendors
[**GetNetworkDeviceVirtualFunction**](NetworkDeviceAPI.md#GetNetworkDeviceVirtualFunction) | **Get** /api/v2/network-devices/{networkDeviceId}/virtual-functions/{virtualFunctionId} | Get a specific virtual function for a specific network device
[**GetNetworkDeviceVirtualFunctions**](NetworkDeviceAPI.md#GetNetworkDeviceVirtualFunctions) | **Get** /api/v2/network-devices/{networkDeviceId}/virtual-functions | Get virtual functions for a specific network device
[**GetNetworkDevices**](NetworkDeviceAPI.md#GetNetworkDevices) | **Get** /api/v2/network-devices | Get paginated Network Devices
[**GetPorts**](NetworkDeviceAPI.md#GetPorts) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/ports | Port statistics for network device directly from the device
[**ListNetworkDeviceBreakouts**](NetworkDeviceAPI.md#ListNetworkDeviceBreakouts) | **Get** /api/v2/network-devices/{networkDeviceId}/breakouts | List all port breakouts on a network device
[**ListNetworkDevicePortIps**](NetworkDeviceAPI.md#ListNetworkDevicePortIps) | **Get** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/{family}/addresses | List IP addresses of an address family on a port
[**MarkNetworkDeviceInstallationReady**](NetworkDeviceAPI.md#MarkNetworkDeviceInstallationReady) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/mark-installation-ready | Marks a planned network device as installable
[**NetworkDeviceSyncTargetSnapshotWithLatestSnapshot**](NetworkDeviceAPI.md#NetworkDeviceSyncTargetSnapshotWithLatestSnapshot) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/sync-target-snapshot-with-latest-snapshot | Updates the target snapshot id to be the latest snapshot of the network device
[**ReProvisionNetworkDevice**](NetworkDeviceAPI.md#ReProvisionNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/re-provision | Re-provision network device
[**RemoveNetworkDevicePortIp**](NetworkDeviceAPI.md#RemoveNetworkDevicePortIp) | **Delete** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/{family}/addresses/{ipId} | Stage removal of an IP address from a port
[**ReplaceNetworkDevice**](NetworkDeviceAPI.md#ReplaceNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/replace | Replace network device
[**ReplaceNetworkDevicePortIps**](NetworkDeviceAPI.md#ReplaceNetworkDevicePortIps) | **Put** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/{family}/addresses | Replace the full IP-address set of an address family on a port (diff against current)
[**ResetNetworkDevice**](NetworkDeviceAPI.md#ResetNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/reset | Resets a network device to default state
[**ReturnNetworkDeviceToPlanned**](NetworkDeviceAPI.md#ReturnNetworkDeviceToPlanned) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/return-to-planned | Returns a network device to planned
[**RevertNetworkDeviceDefectiveState**](NetworkDeviceAPI.md#RevertNetworkDeviceDefectiveState) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/revert-defective-state | Revert network device defective state
[**RunExtensionOnNetworkDevice**](NetworkDeviceAPI.md#RunExtensionOnNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/run-extension | Runs an extension of type action on the network device
[**SetNetworkDeviceAsDefective**](NetworkDeviceAPI.md#SetNetworkDeviceAsDefective) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-as-defective | Set network device as defective
[**SetNetworkDeviceHealthMonitoringFilter**](NetworkDeviceAPI.md#SetNetworkDeviceHealthMonitoringFilter) | **Get** /api/v2/network-devices/health-monitoring-filter | Set the health monitoring WebSocket filter for the current user
[**SetNetworkDevicePortStatus**](NetworkDeviceAPI.md#SetNetworkDevicePortStatus) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-port-status | Set port status
[**SetNetworkDeviceSecret**](NetworkDeviceAPI.md#SetNetworkDeviceSecret) | **Patch** /api/v2/network-devices/{networkDeviceId}/secrets | Create or replace one Network Device secret
[**StartNetworkDeviceRegistration**](NetworkDeviceAPI.md#StartNetworkDeviceRegistration) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/start-registration | Starts registration for a network device
[**UpdateNetworkDevice**](NetworkDeviceAPI.md#UpdateNetworkDevice) | **Patch** /api/v2/network-devices/{networkDeviceId} | Update Network Device
[**UpdateNetworkDeviceBreakoutConfig**](NetworkDeviceAPI.md#UpdateNetworkDeviceBreakoutConfig) | **Patch** /api/v2/network-devices/{networkDeviceId}/breakouts/{breakoutId}/config | Edit the staged breakout groups (reconciles child interfaces)
[**UpdateNetworkDevicePortConfig**](NetworkDeviceAPI.md#UpdateNetworkDevicePortConfig) | **Patch** /api/v2/network-devices/{networkDeviceId}/ports/{portId}/config | Stage an update to a port&#39;s desired-state config
[**UpdateNetworkDeviceVendor**](NetworkDeviceAPI.md#UpdateNetworkDeviceVendor) | **Patch** /api/v2/network-devices/vendors/{vendorId} | Update Network Device Vendor



## AcknowledgeNetworkDeviceDrift

> NetworkDeviceDriftHistory AcknowledgeNetworkDeviceDrift(ctx, networkDeviceId, driftId).Execute()

Acknowledge a network device drift

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
	networkDeviceId := int64(789) // int64 | Network device ID
	driftId := int64(789) // int64 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.AcknowledgeNetworkDeviceDrift(context.Background(), networkDeviceId, driftId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.AcknowledgeNetworkDeviceDrift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcknowledgeNetworkDeviceDrift`: NetworkDeviceDriftHistory
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.AcknowledgeNetworkDeviceDrift`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 
**driftId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeNetworkDeviceDriftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkDeviceDriftHistory**](NetworkDeviceDriftHistory.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddNetworkDevicePortIp

> NetworkEquipmentInterfaceIp AddNetworkDevicePortIp(ctx, networkDeviceId, portId, family).AddNetworkEquipmentInterfaceIp(addNetworkEquipmentInterfaceIp).IfMatch(ifMatch).Execute()

Stage adding an IP address of an address family to a port

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 
	family := "family_example" // string | 
	addNetworkEquipmentInterfaceIp := *openapiclient.NewAddNetworkEquipmentInterfaceIp("Address_example", int32(123)) // AddNetworkEquipmentInterfaceIp | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.AddNetworkDevicePortIp(context.Background(), networkDeviceId, portId, family).AddNetworkEquipmentInterfaceIp(addNetworkEquipmentInterfaceIp).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.AddNetworkDevicePortIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNetworkDevicePortIp`: NetworkEquipmentInterfaceIp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.AddNetworkDevicePortIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 
**family** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkDevicePortIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addNetworkEquipmentInterfaceIp** | [**AddNetworkEquipmentInterfaceIp**](AddNetworkEquipmentInterfaceIp.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkEquipmentInterfaceIp**](NetworkEquipmentInterfaceIp.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveNetworkDevice

> ArchiveNetworkDevice(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

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
	networkDeviceId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.ArchiveNetworkDevice(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ArchiveNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveNetworkDeviceRequest struct via the builder pattern


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


## CreateNetworkDeviceBreakout

> NetworkEquipmentBreakout CreateNetworkDeviceBreakout(ctx, networkDeviceId).CreateNetworkEquipmentBreakout(createNetworkEquipmentBreakout).Execute()

Create a port breakout (Cumulus only)

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
	networkDeviceId := int64(789) // int64 | 
	createNetworkEquipmentBreakout := *openapiclient.NewCreateNetworkEquipmentBreakout("swp1", []openapiclient.BreakoutGroup{*openapiclient.NewBreakoutGroup(int32(4))}) // CreateNetworkEquipmentBreakout | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.CreateNetworkDeviceBreakout(context.Background(), networkDeviceId).CreateNetworkEquipmentBreakout(createNetworkEquipmentBreakout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.CreateNetworkDeviceBreakout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDeviceBreakout`: NetworkEquipmentBreakout
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.CreateNetworkDeviceBreakout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceBreakoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkEquipmentBreakout** | [**CreateNetworkEquipmentBreakout**](CreateNetworkEquipmentBreakout.md) |  | 

### Return type

[**NetworkEquipmentBreakout**](NetworkEquipmentBreakout.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkDevicePort

> NetworkEquipmentInterface CreateNetworkDevicePort(ctx, networkDeviceId).CreateNetworkEquipmentInterface(createNetworkEquipmentInterface).Execute()

Stage creation of a logical port on a network device (not supported yet — returns 501)

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
	networkDeviceId := int64(789) // int64 | 
	createNetworkEquipmentInterface := *openapiclient.NewCreateNetworkEquipmentInterface("Kind_example", "Name_example") // CreateNetworkEquipmentInterface | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.CreateNetworkDevicePort(context.Background(), networkDeviceId).CreateNetworkEquipmentInterface(createNetworkEquipmentInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.CreateNetworkDevicePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDevicePort`: NetworkEquipmentInterface
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.CreateNetworkDevicePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDevicePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkEquipmentInterface** | [**CreateNetworkEquipmentInterface**](CreateNetworkEquipmentInterface.md) |  | 

### Return type

[**NetworkEquipmentInterface**](NetworkEquipmentInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDevice

> DeleteNetworkDevice(ctx, networkDeviceId).Execute()

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
	networkDeviceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DeleteNetworkDevice(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DeleteNetworkDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceRequest struct via the builder pattern


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


## DeleteNetworkDeviceBreakout

> DeleteNetworkDeviceBreakout(ctx, networkDeviceId, breakoutId).Execute()

Remove a port breakout (un-breakout)

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
	networkDeviceId := int64(789) // int64 | 
	breakoutId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DeleteNetworkDeviceBreakout(context.Background(), networkDeviceId, breakoutId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DeleteNetworkDeviceBreakout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**breakoutId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceBreakoutRequest struct via the builder pattern


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


## DeleteNetworkDevicePort

> DeleteNetworkDevicePort(ctx, networkDeviceId, portId).IfMatch(ifMatch).Execute()

Stage removal of a port

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DeleteNetworkDevicePort(context.Background(), networkDeviceId, portId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DeleteNetworkDevicePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDevicePortRequest struct via the builder pattern


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


## DeleteNetworkDeviceSecret

> DeleteNetworkDeviceSecret(ctx, networkDeviceId, name).Execute()

Delete one Network Device secret

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
	networkDeviceId := int64(789) // int64 | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DeleteNetworkDeviceSecret(context.Background(), networkDeviceId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DeleteNetworkDeviceSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceSecretRequest struct via the builder pattern


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


## DeleteNetworkDeviceSecrets

> DeleteNetworkDeviceSecrets(ctx, networkDeviceId).Execute()

Delete all secrets of a Network Device

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
	networkDeviceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DeleteNetworkDeviceSecrets(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DeleteNetworkDeviceSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceSecretsRequest struct via the builder pattern


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


## DisableNetworkDeviceSnmpMonitoring

> DisableNetworkDeviceSnmpMonitoring(ctx, networkDeviceId).Execute()

Disables SNMP monitoring for a network device



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
	networkDeviceId := int64(789) // int64 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DisableNetworkDeviceSnmpMonitoring(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DisableNetworkDeviceSnmpMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableNetworkDeviceSnmpMonitoringRequest struct via the builder pattern


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


## DisableNetworkDeviceSnmpMonitoringBatch

> DisableNetworkDeviceSnmpMonitoringBatch(ctx).NetworkDeviceSNMPMonitoringChangeStatus(networkDeviceSNMPMonitoringChangeStatus).Execute()

Disables SNMP monitoring for a batch of network devices



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
	networkDeviceSNMPMonitoringChangeStatus := *openapiclient.NewNetworkDeviceSNMPMonitoringChangeStatus() // NetworkDeviceSNMPMonitoringChangeStatus | The Network Device SNMP Monitoring Agent Info Batch object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.DisableNetworkDeviceSnmpMonitoringBatch(context.Background()).NetworkDeviceSNMPMonitoringChangeStatus(networkDeviceSNMPMonitoringChangeStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DisableNetworkDeviceSnmpMonitoringBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableNetworkDeviceSnmpMonitoringBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkDeviceSNMPMonitoringChangeStatus** | [**NetworkDeviceSNMPMonitoringChangeStatus**](NetworkDeviceSNMPMonitoringChangeStatus.md) | The Network Device SNMP Monitoring Agent Info Batch object | 

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


## DisableNetworkDeviceSnmpService

> JobInfo DisableNetworkDeviceSnmpService(ctx, networkDeviceId).Execute()

Disables SNMP service on a network device



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
	networkDeviceId := int64(789) // int64 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.DisableNetworkDeviceSnmpService(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DisableNetworkDeviceSnmpService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableNetworkDeviceSnmpService`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.DisableNetworkDeviceSnmpService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableNetworkDeviceSnmpServiceRequest struct via the builder pattern


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


## DisableNetworkDeviceSyslog

> JobInfo DisableNetworkDeviceSyslog(ctx, networkDeviceId).Execute()

Disables remote syslog for a network device



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
	networkDeviceId := int64(789) // int64 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.DisableNetworkDeviceSyslog(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.DisableNetworkDeviceSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableNetworkDeviceSyslog`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.DisableNetworkDeviceSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableNetworkDeviceSyslogRequest struct via the builder pattern


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
	networkDeviceId := int64(789) // int64 | Network device identifier
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
**networkDeviceId** | **int64** | Network device identifier | 

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


## EnableNetworkDeviceSnmpMonitoring

> EnableNetworkDeviceSnmpMonitoring(ctx, networkDeviceId).Execute()

Enables SNMP monitoring for a network device



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
	networkDeviceId := int64(789) // int64 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.EnableNetworkDeviceSnmpMonitoring(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.EnableNetworkDeviceSnmpMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableNetworkDeviceSnmpMonitoringRequest struct via the builder pattern


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


## EnableNetworkDeviceSnmpMonitoringBatch

> EnableNetworkDeviceSnmpMonitoringBatch(ctx).NetworkDeviceSNMPMonitoringChangeStatus(networkDeviceSNMPMonitoringChangeStatus).Execute()

Enables SNMP monitoring for a batch of network devices



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
	networkDeviceSNMPMonitoringChangeStatus := *openapiclient.NewNetworkDeviceSNMPMonitoringChangeStatus() // NetworkDeviceSNMPMonitoringChangeStatus | The Network Device SNMP Monitoring Agent Info Batch object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.EnableNetworkDeviceSnmpMonitoringBatch(context.Background()).NetworkDeviceSNMPMonitoringChangeStatus(networkDeviceSNMPMonitoringChangeStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.EnableNetworkDeviceSnmpMonitoringBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableNetworkDeviceSnmpMonitoringBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkDeviceSNMPMonitoringChangeStatus** | [**NetworkDeviceSNMPMonitoringChangeStatus**](NetworkDeviceSNMPMonitoringChangeStatus.md) | The Network Device SNMP Monitoring Agent Info Batch object | 

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


## EnableNetworkDeviceSnmpService

> JobInfo EnableNetworkDeviceSnmpService(ctx, networkDeviceId).NetworkDeviceSNMPConfig(networkDeviceSNMPConfig).Execute()

Enables SNMP service on a network device



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
	networkDeviceId := int64(789) // int64 | Network device ID
	networkDeviceSNMPConfig := *openapiclient.NewNetworkDeviceSNMPConfig() // NetworkDeviceSNMPConfig | The SNMP service configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.EnableNetworkDeviceSnmpService(context.Background(), networkDeviceId).NetworkDeviceSNMPConfig(networkDeviceSNMPConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.EnableNetworkDeviceSnmpService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableNetworkDeviceSnmpService`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.EnableNetworkDeviceSnmpService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableNetworkDeviceSnmpServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDeviceSNMPConfig** | [**NetworkDeviceSNMPConfig**](NetworkDeviceSNMPConfig.md) | The SNMP service configuration | 

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


## EnableNetworkDeviceSyslog

> JobInfo EnableNetworkDeviceSyslog(ctx, networkDeviceId).Execute()

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
	networkDeviceId := int64(789) // int64 | Network device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.EnableNetworkDeviceSyslog(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.EnableNetworkDeviceSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableNetworkDeviceSyslog`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.EnableNetworkDeviceSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableNetworkDeviceSyslogRequest struct via the builder pattern


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
	networkDeviceId := int64(789) // int64 | 

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
**networkDeviceId** | **int64** |  | 

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


## GetNetworkDeviceBreakout

> NetworkEquipmentBreakout GetNetworkDeviceBreakout(ctx, networkDeviceId, breakoutId).Execute()

Get a port breakout by id

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
	networkDeviceId := int64(789) // int64 | 
	breakoutId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceBreakout(context.Background(), networkDeviceId, breakoutId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceBreakout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceBreakout`: NetworkEquipmentBreakout
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceBreakout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**breakoutId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceBreakoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkEquipmentBreakout**](NetworkEquipmentBreakout.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceBreakoutConfig

> NetworkEquipmentBreakoutConfigDto GetNetworkDeviceBreakoutConfig(ctx, networkDeviceId, breakoutId).Execute()

Get the staged (desired) breakout groups for a port breakout

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
	networkDeviceId := int64(789) // int64 | 
	breakoutId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceBreakoutConfig(context.Background(), networkDeviceId, breakoutId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceBreakoutConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceBreakoutConfig`: NetworkEquipmentBreakoutConfigDto
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceBreakoutConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**breakoutId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceBreakoutConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkEquipmentBreakoutConfigDto**](NetworkEquipmentBreakoutConfigDto.md)

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
	networkDeviceId := int64(789) // int64 | 

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
**networkDeviceId** | **int64** |  | 

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


## GetNetworkDeviceDriftHistory

> NetworkDeviceDriftHistoryPaginatedList GetNetworkDeviceDriftHistory(ctx, networkDeviceId).Page(page).Limit(limit).FilterId(filterId).FilterSnapshotId(filterSnapshotId).FilterAcknowledgedBy(filterAcknowledgedBy).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get drift history for a specific network device

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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSnapshotId := []string{"Inner_example"} // []string | Filter by snapshotId query param.  **Format:** filter.snapshotId={$not}:OPERATION:VALUE    **Example:** filter.snapshotId=$btw:John Doe&filter.snapshotId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAcknowledgedBy := []string{"Inner_example"} // []string | Filter by acknowledgedBy query param.  **Format:** filter.acknowledgedBy={$not}:OPERATION:VALUE    **Example:** filter.acknowledgedBy=$btw:John Doe&filter.acknowledgedBy=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=snapshotId:DESC   **Default Value:** id:ASC  **Available Fields** - id  - snapshotId  - acknowledgedBy  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,snapshotId,acknowledgedBy   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - snapshotId  - acknowledgedBy  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceDriftHistory(context.Background(), networkDeviceId).Page(page).Limit(limit).FilterId(filterId).FilterSnapshotId(filterSnapshotId).FilterAcknowledgedBy(filterAcknowledgedBy).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceDriftHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceDriftHistory`: NetworkDeviceDriftHistoryPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceDriftHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceDriftHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSnapshotId** | **[]string** | Filter by snapshotId query param.  **Format:** filter.snapshotId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.snapshotId&#x3D;$btw:John Doe&amp;filter.snapshotId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAcknowledgedBy** | **[]string** | Filter by acknowledgedBy query param.  **Format:** filter.acknowledgedBy&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.acknowledgedBy&#x3D;$btw:John Doe&amp;filter.acknowledgedBy&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;snapshotId:DESC   **Default Value:** id:ASC  **Available Fields** - id  - snapshotId  - acknowledgedBy  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,snapshotId,acknowledgedBy   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - snapshotId  - acknowledgedBy  | 

### Return type

[**NetworkDeviceDriftHistoryPaginatedList**](NetworkDeviceDriftHistoryPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceDriftHistoryById

> NetworkDeviceDriftHistory GetNetworkDeviceDriftHistoryById(ctx, networkDeviceId, driftId).Execute()

Get a specific drift history for a specific network device

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
	driftId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceDriftHistoryById(context.Background(), networkDeviceId, driftId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceDriftHistoryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceDriftHistoryById`: NetworkDeviceDriftHistory
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceDriftHistoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 
**driftId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceDriftHistoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkDeviceDriftHistory**](NetworkDeviceDriftHistory.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceDriverCapabilities

> []NetworkDeviceDriverCapabilities GetNetworkDeviceDriverCapabilities(ctx).Search(search).Execute()

Get Network Device driver capabilities

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
	search := "search_example" // string | Filter by driver name (substring) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceDriverCapabilities(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceDriverCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceDriverCapabilities`: []NetworkDeviceDriverCapabilities
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceDriverCapabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceDriverCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Filter by driver name (substring) | 

### Return type

[**[]NetworkDeviceDriverCapabilities**](NetworkDeviceDriverCapabilities.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceHealthSummary

> NetworkDeviceHealthSummaryDto GetNetworkDeviceHealthSummary(ctx, networkDeviceId).Execute()

Get Network Device health summary



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
	networkDeviceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceHealthSummary(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceHealthSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceHealthSummary`: NetworkDeviceHealthSummaryDto
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceHealthSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceHealthSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceHealthSummaryDto**](NetworkDeviceHealthSummaryDto.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicePort

> NetworkEquipmentInterface GetNetworkDevicePort(ctx, networkDeviceId, portId).Execute()

Get one port on a network device

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevicePort(context.Background(), networkDeviceId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevicePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevicePort`: NetworkEquipmentInterface
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevicePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkEquipmentInterface**](NetworkEquipmentInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicePortConfig

> NetworkEquipmentInterfaceConfig GetNetworkDevicePortConfig(ctx, networkDeviceId, portId).Execute()

Get the desired-state config for a port

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevicePortConfig(context.Background(), networkDeviceId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevicePortConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevicePortConfig`: NetworkEquipmentInterfaceConfig
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevicePortConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicePortConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkEquipmentInterfaceConfig**](NetworkEquipmentInterfaceConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicePortIp

> NetworkEquipmentInterfaceIp GetNetworkDevicePortIp(ctx, networkDeviceId, portId, family, ipId).Execute()

Get one IP-address aspect row on a port

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 
	family := "family_example" // string | 
	ipId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevicePortIp(context.Background(), networkDeviceId, portId, family, ipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevicePortIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevicePortIp`: NetworkEquipmentInterfaceIp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevicePortIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 
**family** | **string** |  | 
**ipId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicePortIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**NetworkEquipmentInterfaceIp**](NetworkEquipmentInterfaceIp.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicePortVirtualFunction

> NetworkDeviceInterfaceVirtualFunction GetNetworkDevicePortVirtualFunction(ctx, networkDeviceId, portId, virtualFunctionId).Execute()

Get a specific virtual function for a specific port on a network device

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
	portId := float32(8.14) // float32 | 
	virtualFunctionId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevicePortVirtualFunction(context.Background(), networkDeviceId, portId, virtualFunctionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevicePortVirtualFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevicePortVirtualFunction`: NetworkDeviceInterfaceVirtualFunction
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevicePortVirtualFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 
**portId** | **float32** |  | 
**virtualFunctionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicePortVirtualFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NetworkDeviceInterfaceVirtualFunction**](NetworkDeviceInterfaceVirtualFunction.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicePortVirtualFunctions

> NetworkDeviceInterfaceVirtualFunctionsPaginatedList GetNetworkDevicePortVirtualFunctions(ctx, networkDeviceId, portId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIndex(filterIndex).FilterLogicalNetworkId(filterLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get virtual functions for a specific port on a network device

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
	portId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterIndex := []string{"Inner_example"} // []string | Filter by index query param.  **Format:** filter.index={$not}:OPERATION:VALUE    **Example:** filter.index=$btw:John Doe&filter.index=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterLogicalNetworkId := []string{"Inner_example"} // []string | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId=$btw:John Doe&filter.logicalNetworkId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:ASC  **Available Fields** - id  - name  - index  - logicalNetworkId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,index,logicalNetworkId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - index  - logicalNetworkId  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevicePortVirtualFunctions(context.Background(), networkDeviceId, portId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIndex(filterIndex).FilterLogicalNetworkId(filterLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDevicePortVirtualFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevicePortVirtualFunctions`: NetworkDeviceInterfaceVirtualFunctionsPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDevicePortVirtualFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 
**portId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicePortVirtualFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterIndex** | **[]string** | Filter by index query param.  **Format:** filter.index&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.index&#x3D;$btw:John Doe&amp;filter.index&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterLogicalNetworkId** | **[]string** | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId&#x3D;$btw:John Doe&amp;filter.logicalNetworkId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:ASC  **Available Fields** - id  - name  - index  - logicalNetworkId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,index,logicalNetworkId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - index  - logicalNetworkId  | 

### Return type

[**NetworkDeviceInterfaceVirtualFunctionsPaginatedList**](NetworkDeviceInterfaceVirtualFunctionsPaginatedList.md)

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


## GetNetworkDeviceSNMPMonitoringAgentInfoBatch

> NetworkDeviceSNMPMonitoringAgentInfoBatch GetNetworkDeviceSNMPMonitoringAgentInfoBatch(ctx).NetworkDeviceIds(networkDeviceIds).SiteIds(siteIds).FabricIds(fabricIds).Execute()

Get Network Device SNMP Monitoring Agent Info Batch

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
	networkDeviceIds := []int64{int64(123)} // []int64 |  (optional)
	siteIds := []int64{int64(123)} // []int64 |  (optional)
	fabricIds := []int64{int64(123)} // []int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceSNMPMonitoringAgentInfoBatch(context.Background()).NetworkDeviceIds(networkDeviceIds).SiteIds(siteIds).FabricIds(fabricIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceSNMPMonitoringAgentInfoBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceSNMPMonitoringAgentInfoBatch`: NetworkDeviceSNMPMonitoringAgentInfoBatch
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceSNMPMonitoringAgentInfoBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceSNMPMonitoringAgentInfoBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkDeviceIds** | **[]int64** |  | 
 **siteIds** | **[]int64** |  | 
 **fabricIds** | **[]int64** |  | 

### Return type

[**NetworkDeviceSNMPMonitoringAgentInfoBatch**](NetworkDeviceSNMPMonitoringAgentInfoBatch.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceSecretCredentials

> NetworkDeviceSecretCredentials GetNetworkDeviceSecretCredentials(ctx, networkDeviceId, name).Execute()

Get one Network Device secret unencrypted



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
	networkDeviceId := int64(789) // int64 | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceSecretCredentials(context.Background(), networkDeviceId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceSecretCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceSecretCredentials`: NetworkDeviceSecretCredentials
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceSecretCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceSecretCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkDeviceSecretCredentials**](NetworkDeviceSecretCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceSecrets

> NetworkDeviceSecretNames GetNetworkDeviceSecrets(ctx, networkDeviceId).Execute()

List the names of the secrets set on a Network Device



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
	networkDeviceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceSecrets(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceSecrets`: NetworkDeviceSecretNames
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceSecretNames**](NetworkDeviceSecretNames.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceSnapshots

> NetworkDeviceSnapshotPaginatedList GetNetworkDeviceSnapshots(ctx, networkDeviceId).Page(page).Limit(limit).Kind(kind).Search(search).Execute()

Get Network Device snapshots



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
	networkDeviceId := int64(789) // int64 | 
	page := float32(8.14) // float32 | Page number (1-based) (optional) (default to 1)
	limit := float32(8.14) // float32 | Number of items per page (max 200) (optional) (default to 20)
	kind := "kind_example" // string | Filter by commit kind (optional)
	search := "search_example" // string | Search by partial oid or message (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceSnapshots(context.Background(), networkDeviceId).Page(page).Limit(limit).Kind(kind).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceSnapshots`: NetworkDeviceSnapshotPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number (1-based) | [default to 1]
 **limit** | **float32** | Number of items per page (max 200) | [default to 20]
 **kind** | **string** | Filter by commit kind | 
 **search** | **string** | Search by partial oid or message | 

### Return type

[**NetworkDeviceSnapshotPaginatedList**](NetworkDeviceSnapshotPaginatedList.md)

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


## GetNetworkDeviceVendor

> NetworkDeviceVendors GetNetworkDeviceVendor(ctx, vendorId).Execute()

Get Network Device Vendor

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
	vendorId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceVendor(context.Background(), vendorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceVendor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceVendor`: NetworkDeviceVendors
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceVendor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceVendorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceVendors**](NetworkDeviceVendors.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceVendors

> NetworkDeviceVendorsPaginatedList GetNetworkDeviceVendors(ctx).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get paginated Network Device Vendors

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
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=kind:DESC   **Default Value:** id:ASC  **Available Fields** - id  - kind  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,kind   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - kind  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceVendors(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceVendors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceVendors`: NetworkDeviceVendorsPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;kind:DESC   **Default Value:** id:ASC  **Available Fields** - id  - kind  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,kind   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - kind  | 

### Return type

[**NetworkDeviceVendorsPaginatedList**](NetworkDeviceVendorsPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceVirtualFunction

> NetworkDeviceInterfaceVirtualFunction GetNetworkDeviceVirtualFunction(ctx, networkDeviceId, virtualFunctionId).Execute()

Get a specific virtual function for a specific network device

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
	virtualFunctionId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceVirtualFunction(context.Background(), networkDeviceId, virtualFunctionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceVirtualFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceVirtualFunction`: NetworkDeviceInterfaceVirtualFunction
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceVirtualFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 
**virtualFunctionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceVirtualFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkDeviceInterfaceVirtualFunction**](NetworkDeviceInterfaceVirtualFunction.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceVirtualFunctions

> NetworkDeviceInterfaceVirtualFunctionsPaginatedList GetNetworkDeviceVirtualFunctions(ctx, networkDeviceId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIndex(filterIndex).FilterLogicalNetworkId(filterLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get virtual functions for a specific network device

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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterIndex := []string{"Inner_example"} // []string | Filter by index query param.  **Format:** filter.index={$not}:OPERATION:VALUE    **Example:** filter.index=$btw:John Doe&filter.index=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterLogicalNetworkId := []string{"Inner_example"} // []string | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId=$btw:John Doe&filter.logicalNetworkId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:ASC  **Available Fields** - id  - name  - index  - logicalNetworkId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,index,logicalNetworkId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - index  - logicalNetworkId  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDeviceVirtualFunctions(context.Background(), networkDeviceId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIndex(filterIndex).FilterLogicalNetworkId(filterLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.GetNetworkDeviceVirtualFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceVirtualFunctions`: NetworkDeviceInterfaceVirtualFunctionsPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.GetNetworkDeviceVirtualFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceVirtualFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterIndex** | **[]string** | Filter by index query param.  **Format:** filter.index&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.index&#x3D;$btw:John Doe&amp;filter.index&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterLogicalNetworkId** | **[]string** | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId&#x3D;$btw:John Doe&amp;filter.logicalNetworkId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:ASC  **Available Fields** - id  - name  - index  - logicalNetworkId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,index,logicalNetworkId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - index  - logicalNetworkId  | 

### Return type

[**NetworkDeviceInterfaceVirtualFunctionsPaginatedList**](NetworkDeviceInterfaceVirtualFunctionsPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevices

> NetworkDevicePaginatedList GetNetworkDevices(ctx).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterDriver(filterDriver).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).FilterServerId(filterServerId).FilterHealthStatus(filterHealthStatus).FilterDriftDetectionSyncStatus(filterDriftDetectionSyncStatus).FilterBackupEnabled(filterBackupEnabled).FilterDriftDetectionEnabled(filterDriftDetectionEnabled).FilterNetworkFabricId(filterNetworkFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	filterDriver := []string{"Inner_example"} // []string | Filter by driver query param.  **Format:** filter.driver={$not}:OPERATION:VALUE    **Example:** filter.driver=$btw:John Doe&filter.driver=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterChassisIdentifier := []string{"Inner_example"} // []string | Filter by chassisIdentifier query param.  **Format:** filter.chassisIdentifier={$not}:OPERATION:VALUE    **Example:** filter.chassisIdentifier=$btw:John Doe&filter.chassisIdentifier=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementAddress := []string{"Inner_example"} // []string | Filter by managementAddress query param.  **Format:** filter.managementAddress={$not}:OPERATION:VALUE    **Example:** filter.managementAddress=$btw:John Doe&filter.managementAddress=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementPort := []string{"Inner_example"} // []string | Filter by managementPort query param.  **Format:** filter.managementPort={$not}:OPERATION:VALUE    **Example:** filter.managementPort=$btw:John Doe&filter.managementPort=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterProvisionerType := []string{"Inner_example"} // []string | Filter by provisionerType query param.  **Format:** filter.provisionerType={$not}:OPERATION:VALUE    **Example:** filter.provisionerType=$btw:John Doe&filter.provisionerType=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPosition := []string{"Inner_example"} // []string | Filter by position query param.  **Format:** filter.position={$not}:OPERATION:VALUE    **Example:** filter.position=$btw:John Doe&filter.position=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterIdentifierString := []string{"Inner_example"} // []string | Filter by identifierString query param.  **Format:** filter.identifierString={$not}:OPERATION:VALUE    **Example:** filter.identifierString=$btw:John Doe&filter.identifierString=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$btw:John Doe&filter.serverId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHealthStatus := []string{"Inner_example"} // []string | Filter by healthStatus query param.  **Format:** filter.healthStatus={$not}:OPERATION:VALUE    **Example:** filter.healthStatus=$btw:John Doe&filter.healthStatus=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDriftDetectionSyncStatus := []string{"Inner_example"} // []string | Filter by driftDetectionSyncStatus query param.  **Format:** filter.driftDetectionSyncStatus={$not}:OPERATION:VALUE    **Example:** filter.driftDetectionSyncStatus=$btw:John Doe&filter.driftDetectionSyncStatus=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterBackupEnabled := []string{"Inner_example"} // []string | Filter by backupEnabled query param.  **Format:** filter.backupEnabled={$not}:OPERATION:VALUE    **Example:** filter.backupEnabled=$btw:John Doe&filter.backupEnabled=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDriftDetectionEnabled := []string{"Inner_example"} // []string | Filter by driftDetectionEnabled query param.  **Format:** filter.driftDetectionEnabled={$not}:OPERATION:VALUE    **Example:** filter.driftDetectionEnabled=$btw:John Doe&filter.driftDetectionEnabled=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNetworkFabricId := []string{"Inner_example"} // []string | Filter by networkFabricId query param.  **Format:** filter.networkFabricId={$not}:OPERATION:VALUE    **Example:** filter.networkFabricId=$eq:John Doe&filter.networkFabricId=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=identifierString:DESC   **Default Value:** id:ASC  **Available Fields** - id  - identifierString  - status  - siteId  - position  - driver  - managementAddress  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,identifierString,status,position,driver   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - identifierString  - status  - position  - driver  - managementAddress  - description  - driver  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.GetNetworkDevices(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterDriver(filterDriver).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).FilterServerId(filterServerId).FilterHealthStatus(filterHealthStatus).FilterDriftDetectionSyncStatus(filterDriftDetectionSyncStatus).FilterBackupEnabled(filterBackupEnabled).FilterDriftDetectionEnabled(filterDriftDetectionEnabled).FilterNetworkFabricId(filterNetworkFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **filterDriver** | **[]string** | Filter by driver query param.  **Format:** filter.driver&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.driver&#x3D;$btw:John Doe&amp;filter.driver&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterChassisIdentifier** | **[]string** | Filter by chassisIdentifier query param.  **Format:** filter.chassisIdentifier&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.chassisIdentifier&#x3D;$btw:John Doe&amp;filter.chassisIdentifier&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementAddress** | **[]string** | Filter by managementAddress query param.  **Format:** filter.managementAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementAddress&#x3D;$btw:John Doe&amp;filter.managementAddress&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementPort** | **[]string** | Filter by managementPort query param.  **Format:** filter.managementPort&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementPort&#x3D;$btw:John Doe&amp;filter.managementPort&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterProvisionerType** | **[]string** | Filter by provisionerType query param.  **Format:** filter.provisionerType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.provisionerType&#x3D;$btw:John Doe&amp;filter.provisionerType&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPosition** | **[]string** | Filter by position query param.  **Format:** filter.position&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.position&#x3D;$btw:John Doe&amp;filter.position&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterIdentifierString** | **[]string** | Filter by identifierString query param.  **Format:** filter.identifierString&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.identifierString&#x3D;$btw:John Doe&amp;filter.identifierString&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$btw:John Doe&amp;filter.serverId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHealthStatus** | **[]string** | Filter by healthStatus query param.  **Format:** filter.healthStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.healthStatus&#x3D;$btw:John Doe&amp;filter.healthStatus&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDriftDetectionSyncStatus** | **[]string** | Filter by driftDetectionSyncStatus query param.  **Format:** filter.driftDetectionSyncStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.driftDetectionSyncStatus&#x3D;$btw:John Doe&amp;filter.driftDetectionSyncStatus&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterBackupEnabled** | **[]string** | Filter by backupEnabled query param.  **Format:** filter.backupEnabled&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.backupEnabled&#x3D;$btw:John Doe&amp;filter.backupEnabled&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDriftDetectionEnabled** | **[]string** | Filter by driftDetectionEnabled query param.  **Format:** filter.driftDetectionEnabled&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.driftDetectionEnabled&#x3D;$btw:John Doe&amp;filter.driftDetectionEnabled&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNetworkFabricId** | **[]string** | Filter by networkFabricId query param.  **Format:** filter.networkFabricId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkFabricId&#x3D;$eq:John Doe&amp;filter.networkFabricId&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;identifierString:DESC   **Default Value:** id:ASC  **Available Fields** - id  - identifierString  - status  - siteId  - position  - driver  - managementAddress  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,identifierString,status,position,driver   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - identifierString  - status  - position  - driver  - managementAddress  - description  - driver  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  | 

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


## ListNetworkDeviceBreakouts

> []NetworkEquipmentBreakout ListNetworkDeviceBreakouts(ctx, networkDeviceId).Execute()

List all port breakouts on a network device

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
	networkDeviceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ListNetworkDeviceBreakouts(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ListNetworkDeviceBreakouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkDeviceBreakouts`: []NetworkEquipmentBreakout
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ListNetworkDeviceBreakouts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkDeviceBreakoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NetworkEquipmentBreakout**](NetworkEquipmentBreakout.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkDevicePortIps

> []NetworkEquipmentInterfaceIp ListNetworkDevicePortIps(ctx, networkDeviceId, portId, family).Execute()

List IP addresses of an address family on a port

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 
	family := "family_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ListNetworkDevicePortIps(context.Background(), networkDeviceId, portId, family).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ListNetworkDevicePortIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkDevicePortIps`: []NetworkEquipmentInterfaceIp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ListNetworkDevicePortIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 
**family** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkDevicePortIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]NetworkEquipmentInterfaceIp**](NetworkEquipmentInterfaceIp.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNetworkDeviceInstallationReady

> NetworkDevice MarkNetworkDeviceInstallationReady(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

Marks a planned network device as installable



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
	networkDeviceId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.MarkNetworkDeviceInstallationReady(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.MarkNetworkDeviceInstallationReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNetworkDeviceInstallationReady`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.MarkNetworkDeviceInstallationReady`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNetworkDeviceInstallationReadyRequest struct via the builder pattern


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


## NetworkDeviceSyncTargetSnapshotWithLatestSnapshot

> NetworkDeviceSyncTargetSnapshotWithLatestSnapshot(ctx, networkDeviceId).Execute()

Updates the target snapshot id to be the latest snapshot of the network device



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
	r, err := apiClient.NetworkDeviceAPI.NetworkDeviceSyncTargetSnapshotWithLatestSnapshot(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.NetworkDeviceSyncTargetSnapshotWithLatestSnapshot``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNetworkDeviceSyncTargetSnapshotWithLatestSnapshotRequest struct via the builder pattern


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
	networkDeviceId := int64(789) // int64 | 
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
**networkDeviceId** | **int64** |  | 

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


## RemoveNetworkDevicePortIp

> RemoveNetworkDevicePortIp(ctx, networkDeviceId, portId, family, ipId).IfMatch(ifMatch).Execute()

Stage removal of an IP address from a port

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 
	family := "family_example" // string | 
	ipId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.RemoveNetworkDevicePortIp(context.Background(), networkDeviceId, portId, family, ipId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.RemoveNetworkDevicePortIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 
**family** | **string** |  | 
**ipId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkDevicePortIpRequest struct via the builder pattern


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
	networkDeviceId := int64(789) // int64 | Network device ID
	switchReplace := *openapiclient.NewSwitchReplace(int64(123)) // SwitchReplace | Network device replacement details

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
**networkDeviceId** | **int64** | Network device ID | 

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


## ReplaceNetworkDevicePortIps

> []NetworkEquipmentInterfaceIp ReplaceNetworkDevicePortIps(ctx, networkDeviceId, portId, family).ReplaceNetworkEquipmentInterfaceIps(replaceNetworkEquipmentInterfaceIps).IfMatch(ifMatch).Execute()

Replace the full IP-address set of an address family on a port (diff against current)

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 
	family := "family_example" // string | 
	replaceNetworkEquipmentInterfaceIps := *openapiclient.NewReplaceNetworkEquipmentInterfaceIps([]openapiclient.AddNetworkEquipmentInterfaceIp{*openapiclient.NewAddNetworkEquipmentInterfaceIp("Address_example", int32(123))}) // ReplaceNetworkEquipmentInterfaceIps | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ReplaceNetworkDevicePortIps(context.Background(), networkDeviceId, portId, family).ReplaceNetworkEquipmentInterfaceIps(replaceNetworkEquipmentInterfaceIps).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ReplaceNetworkDevicePortIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceNetworkDevicePortIps`: []NetworkEquipmentInterfaceIp
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ReplaceNetworkDevicePortIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 
**family** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNetworkDevicePortIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **replaceNetworkEquipmentInterfaceIps** | [**ReplaceNetworkEquipmentInterfaceIps**](ReplaceNetworkEquipmentInterfaceIps.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**[]NetworkEquipmentInterfaceIp**](NetworkEquipmentInterfaceIp.md)

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
	networkDeviceId := int64(789) // int64 | Network device ID

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
**networkDeviceId** | **int64** | Network device ID | 

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


## ReturnNetworkDeviceToPlanned

> NetworkDevice ReturnNetworkDeviceToPlanned(ctx, networkDeviceId).ReturnNetworkDeviceToPlanned(returnNetworkDeviceToPlanned).IfMatch(ifMatch).Execute()

Returns a network device to planned



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
	networkDeviceId := int64(789) // int64 | 
	returnNetworkDeviceToPlanned := *openapiclient.NewReturnNetworkDeviceToPlanned() // ReturnNetworkDeviceToPlanned | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.ReturnNetworkDeviceToPlanned(context.Background(), networkDeviceId).ReturnNetworkDeviceToPlanned(returnNetworkDeviceToPlanned).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.ReturnNetworkDeviceToPlanned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnNetworkDeviceToPlanned`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.ReturnNetworkDeviceToPlanned`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnNetworkDeviceToPlannedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnNetworkDeviceToPlanned** | [**ReturnNetworkDeviceToPlanned**](ReturnNetworkDeviceToPlanned.md) |  | 
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


## RevertNetworkDeviceDefectiveState

> NetworkDevice RevertNetworkDeviceDefectiveState(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

Revert network device defective state



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
	networkDeviceId := int64(789) // int64 | Network device ID
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.RevertNetworkDeviceDefectiveState(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.RevertNetworkDeviceDefectiveState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertNetworkDeviceDefectiveState`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.RevertNetworkDeviceDefectiveState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertNetworkDeviceDefectiveStateRequest struct via the builder pattern


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
	networkDeviceId := int64(789) // int64 | 
	runExtensionOnPhysicalDevice := *openapiclient.NewRunExtensionOnPhysicalDevice(int64(10), map[string]interface{}(123)) // RunExtensionOnPhysicalDevice | The extension information
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
**networkDeviceId** | **int64** |  | 

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


## SetNetworkDeviceAsDefective

> NetworkDevice SetNetworkDeviceAsDefective(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

Set network device as defective



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
	networkDeviceId := int64(789) // int64 | Network device ID
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.SetNetworkDeviceAsDefective(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.SetNetworkDeviceAsDefective``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNetworkDeviceAsDefective`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.SetNetworkDeviceAsDefective`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** | Network device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNetworkDeviceAsDefectiveRequest struct via the builder pattern


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


## SetNetworkDeviceHealthMonitoringFilter

> SetNetworkDeviceHealthMonitoringFilter(ctx).SocketId(socketId).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterDriver(filterDriver).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).FilterServerId(filterServerId).FilterHealthStatus(filterHealthStatus).FilterDriftDetectionSyncStatus(filterDriftDetectionSyncStatus).FilterBackupEnabled(filterBackupEnabled).FilterDriftDetectionEnabled(filterDriftDetectionEnabled).FilterNetworkFabricId(filterNetworkFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Set the health monitoring WebSocket filter for the current user

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
	socketId := "socketId_example" // string | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDatacenterName := []string{"Inner_example"} // []string | Filter by datacenterName query param.  **Format:** filter.datacenterName={$not}:OPERATION:VALUE    **Example:** filter.datacenterName=$btw:John Doe&filter.datacenterName=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDriver := []string{"Inner_example"} // []string | Filter by driver query param.  **Format:** filter.driver={$not}:OPERATION:VALUE    **Example:** filter.driver=$btw:John Doe&filter.driver=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterChassisIdentifier := []string{"Inner_example"} // []string | Filter by chassisIdentifier query param.  **Format:** filter.chassisIdentifier={$not}:OPERATION:VALUE    **Example:** filter.chassisIdentifier=$btw:John Doe&filter.chassisIdentifier=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementAddress := []string{"Inner_example"} // []string | Filter by managementAddress query param.  **Format:** filter.managementAddress={$not}:OPERATION:VALUE    **Example:** filter.managementAddress=$btw:John Doe&filter.managementAddress=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementPort := []string{"Inner_example"} // []string | Filter by managementPort query param.  **Format:** filter.managementPort={$not}:OPERATION:VALUE    **Example:** filter.managementPort=$btw:John Doe&filter.managementPort=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterProvisionerType := []string{"Inner_example"} // []string | Filter by provisionerType query param.  **Format:** filter.provisionerType={$not}:OPERATION:VALUE    **Example:** filter.provisionerType=$btw:John Doe&filter.provisionerType=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPosition := []string{"Inner_example"} // []string | Filter by position query param.  **Format:** filter.position={$not}:OPERATION:VALUE    **Example:** filter.position=$btw:John Doe&filter.position=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterIdentifierString := []string{"Inner_example"} // []string | Filter by identifierString query param.  **Format:** filter.identifierString={$not}:OPERATION:VALUE    **Example:** filter.identifierString=$btw:John Doe&filter.identifierString=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$btw:John Doe&filter.serverId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHealthStatus := []string{"Inner_example"} // []string | Filter by healthStatus query param.  **Format:** filter.healthStatus={$not}:OPERATION:VALUE    **Example:** filter.healthStatus=$btw:John Doe&filter.healthStatus=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDriftDetectionSyncStatus := []string{"Inner_example"} // []string | Filter by driftDetectionSyncStatus query param.  **Format:** filter.driftDetectionSyncStatus={$not}:OPERATION:VALUE    **Example:** filter.driftDetectionSyncStatus=$btw:John Doe&filter.driftDetectionSyncStatus=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterBackupEnabled := []string{"Inner_example"} // []string | Filter by backupEnabled query param.  **Format:** filter.backupEnabled={$not}:OPERATION:VALUE    **Example:** filter.backupEnabled=$btw:John Doe&filter.backupEnabled=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDriftDetectionEnabled := []string{"Inner_example"} // []string | Filter by driftDetectionEnabled query param.  **Format:** filter.driftDetectionEnabled={$not}:OPERATION:VALUE    **Example:** filter.driftDetectionEnabled=$btw:John Doe&filter.driftDetectionEnabled=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNetworkFabricId := []string{"Inner_example"} // []string | Filter by networkFabricId query param.  **Format:** filter.networkFabricId={$not}:OPERATION:VALUE    **Example:** filter.networkFabricId=$eq:John Doe&filter.networkFabricId=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=identifierString:DESC   **Default Value:** id:ASC  **Available Fields** - id  - identifierString  - status  - siteId  - position  - driver  - managementAddress  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,identifierString,status,position,driver   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - identifierString  - status  - position  - driver  - managementAddress  - description  - driver  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceAPI.SetNetworkDeviceHealthMonitoringFilter(context.Background()).SocketId(socketId).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterDriver(filterDriver).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).FilterServerId(filterServerId).FilterHealthStatus(filterHealthStatus).FilterDriftDetectionSyncStatus(filterDriftDetectionSyncStatus).FilterBackupEnabled(filterBackupEnabled).FilterDriftDetectionEnabled(filterDriftDetectionEnabled).FilterNetworkFabricId(filterNetworkFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.SetNetworkDeviceHealthMonitoringFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNetworkDeviceHealthMonitoringFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socketId** | **string** |  | 
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDatacenterName** | **[]string** | Filter by datacenterName query param.  **Format:** filter.datacenterName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.datacenterName&#x3D;$btw:John Doe&amp;filter.datacenterName&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDriver** | **[]string** | Filter by driver query param.  **Format:** filter.driver&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.driver&#x3D;$btw:John Doe&amp;filter.driver&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterChassisIdentifier** | **[]string** | Filter by chassisIdentifier query param.  **Format:** filter.chassisIdentifier&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.chassisIdentifier&#x3D;$btw:John Doe&amp;filter.chassisIdentifier&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementAddress** | **[]string** | Filter by managementAddress query param.  **Format:** filter.managementAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementAddress&#x3D;$btw:John Doe&amp;filter.managementAddress&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementPort** | **[]string** | Filter by managementPort query param.  **Format:** filter.managementPort&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementPort&#x3D;$btw:John Doe&amp;filter.managementPort&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterProvisionerType** | **[]string** | Filter by provisionerType query param.  **Format:** filter.provisionerType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.provisionerType&#x3D;$btw:John Doe&amp;filter.provisionerType&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPosition** | **[]string** | Filter by position query param.  **Format:** filter.position&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.position&#x3D;$btw:John Doe&amp;filter.position&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterIdentifierString** | **[]string** | Filter by identifierString query param.  **Format:** filter.identifierString&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.identifierString&#x3D;$btw:John Doe&amp;filter.identifierString&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$btw:John Doe&amp;filter.serverId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHealthStatus** | **[]string** | Filter by healthStatus query param.  **Format:** filter.healthStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.healthStatus&#x3D;$btw:John Doe&amp;filter.healthStatus&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDriftDetectionSyncStatus** | **[]string** | Filter by driftDetectionSyncStatus query param.  **Format:** filter.driftDetectionSyncStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.driftDetectionSyncStatus&#x3D;$btw:John Doe&amp;filter.driftDetectionSyncStatus&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterBackupEnabled** | **[]string** | Filter by backupEnabled query param.  **Format:** filter.backupEnabled&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.backupEnabled&#x3D;$btw:John Doe&amp;filter.backupEnabled&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDriftDetectionEnabled** | **[]string** | Filter by driftDetectionEnabled query param.  **Format:** filter.driftDetectionEnabled&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.driftDetectionEnabled&#x3D;$btw:John Doe&amp;filter.driftDetectionEnabled&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNetworkFabricId** | **[]string** | Filter by networkFabricId query param.  **Format:** filter.networkFabricId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkFabricId&#x3D;$eq:John Doe&amp;filter.networkFabricId&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;identifierString:DESC   **Default Value:** id:ASC  **Available Fields** - id  - identifierString  - status  - siteId  - position  - driver  - managementAddress  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,identifierString,status,position,driver   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - identifierString  - status  - position  - driver  - managementAddress  - description  - driver  - healthStatus  - driftDetectionSyncStatus  - backupEnabled  - driftDetectionEnabled  | 

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
	networkDeviceId := int64(789) // int64 | Network device ID
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
**networkDeviceId** | **int64** | Network device ID | 

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


## SetNetworkDeviceSecret

> NetworkDeviceSecretNames SetNetworkDeviceSecret(ctx, networkDeviceId).SetNetworkDeviceSecret(setNetworkDeviceSecret).Execute()

Create or replace one Network Device secret



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
	networkDeviceId := int64(789) // int64 | 
	setNetworkDeviceSecret := *openapiclient.NewSetNetworkDeviceSecret("peer_group_backbone_password", "Secret_example") // SetNetworkDeviceSecret | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.SetNetworkDeviceSecret(context.Background(), networkDeviceId).SetNetworkDeviceSecret(setNetworkDeviceSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.SetNetworkDeviceSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNetworkDeviceSecret`: NetworkDeviceSecretNames
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.SetNetworkDeviceSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNetworkDeviceSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNetworkDeviceSecret** | [**SetNetworkDeviceSecret**](SetNetworkDeviceSecret.md) |  | 

### Return type

[**NetworkDeviceSecretNames**](NetworkDeviceSecretNames.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartNetworkDeviceRegistration

> NetworkDevice StartNetworkDeviceRegistration(ctx, networkDeviceId).IfMatch(ifMatch).Execute()

Starts registration for a network device



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
	networkDeviceId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.StartNetworkDeviceRegistration(context.Background(), networkDeviceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.StartNetworkDeviceRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartNetworkDeviceRegistration`: NetworkDevice
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.StartNetworkDeviceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartNetworkDeviceRegistrationRequest struct via the builder pattern


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
	networkDeviceId := int64(789) // int64 | 
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
**networkDeviceId** | **int64** |  | 

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


## UpdateNetworkDeviceBreakoutConfig

> NetworkEquipmentBreakoutConfigDto UpdateNetworkDeviceBreakoutConfig(ctx, networkDeviceId, breakoutId).UpdateNetworkEquipmentBreakoutConfigDto(updateNetworkEquipmentBreakoutConfigDto).IfMatch(ifMatch).Execute()

Edit the staged breakout groups (reconciles child interfaces)

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
	networkDeviceId := int64(789) // int64 | 
	breakoutId := int64(789) // int64 | 
	updateNetworkEquipmentBreakoutConfigDto := *openapiclient.NewUpdateNetworkEquipmentBreakoutConfigDto([]openapiclient.BreakoutGroup{*openapiclient.NewBreakoutGroup(int32(4))}) // UpdateNetworkEquipmentBreakoutConfigDto | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.UpdateNetworkDeviceBreakoutConfig(context.Background(), networkDeviceId, breakoutId).UpdateNetworkEquipmentBreakoutConfigDto(updateNetworkEquipmentBreakoutConfigDto).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.UpdateNetworkDeviceBreakoutConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceBreakoutConfig`: NetworkEquipmentBreakoutConfigDto
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.UpdateNetworkDeviceBreakoutConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**breakoutId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceBreakoutConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkEquipmentBreakoutConfigDto** | [**UpdateNetworkEquipmentBreakoutConfigDto**](UpdateNetworkEquipmentBreakoutConfigDto.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkEquipmentBreakoutConfigDto**](NetworkEquipmentBreakoutConfigDto.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDevicePortConfig

> NetworkEquipmentInterfaceConfig UpdateNetworkDevicePortConfig(ctx, networkDeviceId, portId).UpdateNetworkEquipmentInterfaceConfig(updateNetworkEquipmentInterfaceConfig).IfMatch(ifMatch).Execute()

Stage an update to a port's desired-state config

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
	networkDeviceId := int64(789) // int64 | 
	portId := int64(789) // int64 | 
	updateNetworkEquipmentInterfaceConfig := *openapiclient.NewUpdateNetworkEquipmentInterfaceConfig() // UpdateNetworkEquipmentInterfaceConfig | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.UpdateNetworkDevicePortConfig(context.Background(), networkDeviceId, portId).UpdateNetworkEquipmentInterfaceConfig(updateNetworkEquipmentInterfaceConfig).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.UpdateNetworkDevicePortConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDevicePortConfig`: NetworkEquipmentInterfaceConfig
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.UpdateNetworkDevicePortConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **int64** |  | 
**portId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDevicePortConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkEquipmentInterfaceConfig** | [**UpdateNetworkEquipmentInterfaceConfig**](UpdateNetworkEquipmentInterfaceConfig.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkEquipmentInterfaceConfig**](NetworkEquipmentInterfaceConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDeviceVendor

> NetworkDeviceVendors UpdateNetworkDeviceVendor(ctx, vendorId).UpdateNetworkDeviceVendorsDto(updateNetworkDeviceVendorsDto).IfMatch(ifMatch).Execute()

Update Network Device Vendor

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
	vendorId := int64(789) // int64 | 
	updateNetworkDeviceVendorsDto := *openapiclient.NewUpdateNetworkDeviceVendorsDto() // UpdateNetworkDeviceVendorsDto | The Network Device Vendor update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceAPI.UpdateNetworkDeviceVendor(context.Background(), vendorId).UpdateNetworkDeviceVendorsDto(updateNetworkDeviceVendorsDto).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceAPI.UpdateNetworkDeviceVendor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceVendor`: NetworkDeviceVendors
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceAPI.UpdateNetworkDeviceVendor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceVendorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDeviceVendorsDto** | [**UpdateNetworkDeviceVendorsDto**](UpdateNetworkDeviceVendorsDto.md) | The Network Device Vendor update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkDeviceVendors**](NetworkDeviceVendors.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

