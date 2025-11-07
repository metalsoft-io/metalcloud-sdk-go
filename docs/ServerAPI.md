# \ServerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveServer**](ServerAPI.md#ArchiveServer) | **Post** /api/v2/servers/{serverId}/actions/archive | Archives a Server
[**ConnectServerInterface**](ServerAPI.md#ConnectServerInterface) | **Post** /api/v2/servers/{serverId}/actions/connect-interface | Connects a server interface to a switch
[**DeleteServer**](ServerAPI.md#DeleteServer) | **Delete** /api/v2/servers/{serverId} | Deletes a Server
[**EnableServerSyslog**](ServerAPI.md#EnableServerSyslog) | **Post** /api/v2/servers/{serverId}/actions/syslog-subscribe | Enables remote syslog for a server
[**GetServerCapabilities**](ServerAPI.md#GetServerCapabilities) | **Get** /api/v2/servers/{serverId}/capabilities | Get Server capabilities
[**GetServerCredentials**](ServerAPI.md#GetServerCredentials) | **Get** /api/v2/servers/{serverId}/credentials | Get Server credentials
[**GetServerInfo**](ServerAPI.md#GetServerInfo) | **Get** /api/v2/servers/{serverId} | Get Server information
[**GetServerPowerStatus**](ServerAPI.md#GetServerPowerStatus) | **Post** /api/v2/servers/{serverId}/actions/get-power | Gets the power state of a server
[**GetServerRemoteConsoleInfo**](ServerAPI.md#GetServerRemoteConsoleInfo) | **Get** /api/v2/servers/{serverId}/remote-console-info | Get Remote Console information
[**GetServerVNCInfo**](ServerAPI.md#GetServerVNCInfo) | **Get** /api/v2/servers/{serverId}/vnc-info | Get VNC information
[**GetServers**](ServerAPI.md#GetServers) | **Get** /api/v2/servers | Get a list of Servers
[**GetServersStatistics**](ServerAPI.md#GetServersStatistics) | **Get** /api/v2/servers/statistics | Get Servers statistics
[**IdentifyServer**](ServerAPI.md#IdentifyServer) | **Post** /api/v2/servers/{serverId}/actions/identify-server | identify the server chassis by blinking the LED
[**ReRegisterServer**](ServerAPI.md#ReRegisterServer) | **Post** /api/v2/servers/{serverId}/actions/re-register | Re-register a server
[**RegisterProductionServer**](ServerAPI.md#RegisterProductionServer) | **Post** /api/v2/servers/actions/register-production | Initialize a production (live) server
[**RegisterServer**](ServerAPI.md#RegisterServer) | **Post** /api/v2/servers | Initialize server registration
[**ResetServerToFactoryDefaults**](ServerAPI.md#ResetServerToFactoryDefaults) | **Post** /api/v2/servers/{serverId}/actions/factory-reset | Resets a server to factory defaults
[**SetServerInterfacesDefaultFabric**](ServerAPI.md#SetServerInterfacesDefaultFabric) | **Post** /api/v2/servers/{serverId}/actions/set-interfaces-default-fabric | Sets the default fabric for the specified server interfaces
[**SetServerInterfacesRedundancyGroup**](ServerAPI.md#SetServerInterfacesRedundancyGroup) | **Post** /api/v2/servers/{serverId}/actions/set-interfaces-redundancy-group | Sets the redundancy group index for the specified server interfaces
[**SetServerPowerState**](ServerAPI.md#SetServerPowerState) | **Post** /api/v2/servers/{serverId}/actions/set-power | Sets the power state of a server
[**UpdateServer**](ServerAPI.md#UpdateServer) | **Patch** /api/v2/servers/{serverId} | Updates the server information
[**UpdateServerEnableSnmp**](ServerAPI.md#UpdateServerEnableSnmp) | **Post** /api/v2/servers/{serverId}/actions/enable-snmp | Enables SNMP on a Server
[**UpdateServerIpmiCredentials**](ServerAPI.md#UpdateServerIpmiCredentials) | **Post** /api/v2/servers/{serverId}/actions/update-ipmi-credentials | Update Server ipmi credentials



## ArchiveServer

> ArchiveServer(ctx, serverId).IfMatch(ifMatch).Execute()

Archives a Server



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
	serverId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.ArchiveServer(context.Background(), serverId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.ArchiveServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveServerRequest struct via the builder pattern


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


## ConnectServerInterface

> ConnectServerInterface(ctx, serverId).ServerConnectInterface(serverConnectInterface).Execute()

Connects a server interface to a switch



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
	serverId := float32(8.14) // float32 | 
	serverConnectInterface := *openapiclient.NewServerConnectInterface(float32(1), "Ethernet1/1", "switch1.example.com") // ServerConnectInterface | The server interface connection options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.ConnectServerInterface(context.Background(), serverId).ServerConnectInterface(serverConnectInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.ConnectServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverConnectInterface** | [**ServerConnectInterface**](ServerConnectInterface.md) | The server interface connection options | 

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


## DeleteServer

> DeleteServer(ctx, serverId).IfMatch(ifMatch).Execute()

Deletes a Server



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
	serverId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.DeleteServer(context.Background(), serverId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.DeleteServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerRequest struct via the builder pattern


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


## EnableServerSyslog

> EnableServerSyslog(ctx, serverId).IfMatch(ifMatch).Execute()

Enables remote syslog for a server



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
	serverId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.EnableServerSyslog(context.Background(), serverId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.EnableServerSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableServerSyslogRequest struct via the builder pattern


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


## GetServerCapabilities

> ServerCapabilities GetServerCapabilities(ctx, serverId).Execute()

Get Server capabilities



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerCapabilities(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCapabilities`: ServerCapabilities
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerCapabilities**](ServerCapabilities.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerCredentials

> ServerCredentials GetServerCredentials(ctx, serverId).Execute()

Get Server credentials



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerCredentials(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCredentials`: ServerCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerCredentials**](ServerCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInfo

> Server GetServerInfo(ctx, serverId).Execute()

Get Server information



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerInfo(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInfo`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerPowerStatus

> string GetServerPowerStatus(ctx, serverId).Execute()

Gets the power state of a server



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerPowerStatus(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerPowerStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerPowerStatusRequest struct via the builder pattern


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


## GetServerRemoteConsoleInfo

> RemoteConsoleInfo GetServerRemoteConsoleInfo(ctx, serverId).Execute()

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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerRemoteConsoleInfo(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerRemoteConsoleInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerRemoteConsoleInfo`: RemoteConsoleInfo
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerRemoteConsoleInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRemoteConsoleInfoRequest struct via the builder pattern


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


## GetServerVNCInfo

> ServerVNCInfo GetServerVNCInfo(ctx, serverId).Execute()

Get VNC information



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerVNCInfo(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerVNCInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerVNCInfo`: ServerVNCInfo
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerVNCInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerVNCInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerVNCInfo**](ServerVNCInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServers

> ServerPaginatedList GetServers(ctx).Page(page).Limit(limit).FilterServerTypeId(filterServerTypeId).FilterServerId(filterServerId).FilterSiteId(filterSiteId).FilterVendor(filterVendor).FilterSerialNumber(filterSerialNumber).FilterManagementAddress(filterManagementAddress).FilterModel(filterModel).FilterAdministrationState(filterAdministrationState).FilterServerDiskWipe(filterServerDiskWipe).FilterPowerStatus(filterPowerStatus).FilterServerDhcpStatus(filterServerDhcpStatus).FilterServerClass(filterServerClass).FilterServerStatus(filterServerStatus).FilterRequiresManualCleaning(filterRequiresManualCleaning).FilterInstanceInfrastructureInfrastructureId(filterInstanceInfrastructureInfrastructureId).FilterInstanceInfrastructureUserIdOwner(filterInstanceInfrastructureUserIdOwner).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Servers



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
	filterServerTypeId := []string{"Inner_example"} // []string | Filter by serverTypeId query param.  **Format:** filter.serverTypeId={$not}:OPERATION:VALUE    **Example:** filter.serverTypeId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterVendor := []string{"Inner_example"} // []string | Filter by vendor query param.  **Format:** filter.vendor={$not}:OPERATION:VALUE    **Example:** filter.vendor=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSerialNumber := []string{"Inner_example"} // []string | Filter by serialNumber query param.  **Format:** filter.serialNumber={$not}:OPERATION:VALUE    **Example:** filter.serialNumber=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterManagementAddress := []string{"Inner_example"} // []string | Filter by managementAddress query param.  **Format:** filter.managementAddress={$not}:OPERATION:VALUE    **Example:** filter.managementAddress=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterModel := []string{"Inner_example"} // []string | Filter by model query param.  **Format:** filter.model={$not}:OPERATION:VALUE    **Example:** filter.model=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterAdministrationState := []string{"Inner_example"} // []string | Filter by administrationState query param.  **Format:** filter.administrationState={$not}:OPERATION:VALUE    **Example:** filter.administrationState=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServerDiskWipe := []string{"Inner_example"} // []string | Filter by serverDiskWipe query param.  **Format:** filter.serverDiskWipe={$not}:OPERATION:VALUE    **Example:** filter.serverDiskWipe=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterPowerStatus := []string{"Inner_example"} // []string | Filter by powerStatus query param.  **Format:** filter.powerStatus={$not}:OPERATION:VALUE    **Example:** filter.powerStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServerDhcpStatus := []string{"Inner_example"} // []string | Filter by serverDhcpStatus query param.  **Format:** filter.serverDhcpStatus={$not}:OPERATION:VALUE    **Example:** filter.serverDhcpStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServerClass := []string{"Inner_example"} // []string | Filter by serverClass query param.  **Format:** filter.serverClass={$not}:OPERATION:VALUE    **Example:** filter.serverClass=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServerStatus := []string{"Inner_example"} // []string | Filter by serverStatus query param.  **Format:** filter.serverStatus={$not}:OPERATION:VALUE    **Example:** filter.serverStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterRequiresManualCleaning := []string{"Inner_example"} // []string | Filter by requiresManualCleaning query param.  **Format:** filter.requiresManualCleaning={$not}:OPERATION:VALUE    **Example:** filter.requiresManualCleaning=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterInstanceInfrastructureInfrastructureId := []string{"Inner_example"} // []string | Filter by instance.infrastructure.infrastructureId query param.  **Format:** filter.instance.infrastructure.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.instance.infrastructure.infrastructureId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterInstanceInfrastructureUserIdOwner := []string{"Inner_example"} // []string | Filter by instance.infrastructure.userIdOwner query param.  **Format:** filter.instance.infrastructure.userIdOwner={$not}:OPERATION:VALUE    **Example:** filter.instance.infrastructure.userIdOwner=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=serverId:DESC&sortBy=serverTypeId:DESC   **Default Value:** serverId:DESC  **Available Fields** - serverId  - serverTypeId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** vendor,model   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - vendor  - model  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServers(context.Background()).Page(page).Limit(limit).FilterServerTypeId(filterServerTypeId).FilterServerId(filterServerId).FilterSiteId(filterSiteId).FilterVendor(filterVendor).FilterSerialNumber(filterSerialNumber).FilterManagementAddress(filterManagementAddress).FilterModel(filterModel).FilterAdministrationState(filterAdministrationState).FilterServerDiskWipe(filterServerDiskWipe).FilterPowerStatus(filterPowerStatus).FilterServerDhcpStatus(filterServerDhcpStatus).FilterServerClass(filterServerClass).FilterServerStatus(filterServerStatus).FilterRequiresManualCleaning(filterRequiresManualCleaning).FilterInstanceInfrastructureInfrastructureId(filterInstanceInfrastructureInfrastructureId).FilterInstanceInfrastructureUserIdOwner(filterInstanceInfrastructureUserIdOwner).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServers`: ServerPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterServerTypeId** | **[]string** | Filter by serverTypeId query param.  **Format:** filter.serverTypeId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverTypeId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterVendor** | **[]string** | Filter by vendor query param.  **Format:** filter.vendor&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vendor&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSerialNumber** | **[]string** | Filter by serialNumber query param.  **Format:** filter.serialNumber&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serialNumber&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterManagementAddress** | **[]string** | Filter by managementAddress query param.  **Format:** filter.managementAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementAddress&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterModel** | **[]string** | Filter by model query param.  **Format:** filter.model&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.model&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterAdministrationState** | **[]string** | Filter by administrationState query param.  **Format:** filter.administrationState&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.administrationState&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServerDiskWipe** | **[]string** | Filter by serverDiskWipe query param.  **Format:** filter.serverDiskWipe&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverDiskWipe&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterPowerStatus** | **[]string** | Filter by powerStatus query param.  **Format:** filter.powerStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.powerStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServerDhcpStatus** | **[]string** | Filter by serverDhcpStatus query param.  **Format:** filter.serverDhcpStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverDhcpStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServerClass** | **[]string** | Filter by serverClass query param.  **Format:** filter.serverClass&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverClass&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServerStatus** | **[]string** | Filter by serverStatus query param.  **Format:** filter.serverStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterRequiresManualCleaning** | **[]string** | Filter by requiresManualCleaning query param.  **Format:** filter.requiresManualCleaning&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.requiresManualCleaning&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterInstanceInfrastructureInfrastructureId** | **[]string** | Filter by instance.infrastructure.infrastructureId query param.  **Format:** filter.instance.infrastructure.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.instance.infrastructure.infrastructureId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterInstanceInfrastructureUserIdOwner** | **[]string** | Filter by instance.infrastructure.userIdOwner query param.  **Format:** filter.instance.infrastructure.userIdOwner&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.instance.infrastructure.userIdOwner&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;serverId:DESC&amp;sortBy&#x3D;serverTypeId:DESC   **Default Value:** serverId:DESC  **Available Fields** - serverId  - serverTypeId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** vendor,model   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - vendor  - model  | 

### Return type

[**ServerPaginatedList**](ServerPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServersStatistics

> ServerStatistics GetServersStatistics(ctx).Execute()

Get Servers statistics



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
	resp, r, err := apiClient.ServerAPI.GetServersStatistics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServersStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServersStatistics`: ServerStatistics
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServersStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServersStatisticsRequest struct via the builder pattern


### Return type

[**ServerStatistics**](ServerStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentifyServer

> IdentifyServer(ctx, serverId).Execute()

identify the server chassis by blinking the LED



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.IdentifyServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.IdentifyServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentifyServerRequest struct via the builder pattern


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


## ReRegisterServer

> ReRegisterServerResponse ReRegisterServer(ctx, serverId).IfMatch(ifMatch).Execute()

Re-register a server



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
	serverId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.ReRegisterServer(context.Background(), serverId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.ReRegisterServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReRegisterServer`: ReRegisterServerResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.ReRegisterServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReRegisterServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

[**ReRegisterServerResponse**](ReRegisterServerResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterProductionServer

> RegisterServerResponse RegisterProductionServer(ctx).RegisterProductionServer(registerProductionServer).Execute()

Initialize a production (live) server



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
	registerProductionServer := *openapiclient.NewRegisterProductionServer(float32(123), *openapiclient.NewRegisterProductionServerSettings(float32(123))) // RegisterProductionServer | The production server registration information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.RegisterProductionServer(context.Background()).RegisterProductionServer(registerProductionServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.RegisterProductionServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterProductionServer`: RegisterServerResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.RegisterProductionServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterProductionServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerProductionServer** | [**RegisterProductionServer**](RegisterProductionServer.md) | The production server registration information | 

### Return type

[**RegisterServerResponse**](RegisterServerResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterServer

> RegisterServerResponse RegisterServer(ctx).RegisterServer(registerServer).Execute()

Initialize server registration



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
	registerServer := *openapiclient.NewRegisterServer(float32(123)) // RegisterServer | The server registration information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.RegisterServer(context.Background()).RegisterServer(registerServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.RegisterServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterServer`: RegisterServerResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.RegisterServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerServer** | [**RegisterServer**](RegisterServer.md) | The server registration information | 

### Return type

[**RegisterServerResponse**](RegisterServerResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetServerToFactoryDefaults

> ResetServerToFactoryDefaults(ctx, serverId).IfMatch(ifMatch).Execute()

Resets a server to factory defaults



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
	serverId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.ResetServerToFactoryDefaults(context.Background(), serverId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.ResetServerToFactoryDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetServerToFactoryDefaultsRequest struct via the builder pattern


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


## SetServerInterfacesDefaultFabric

> SetServerInterfacesDefaultFabric(ctx, serverId).ServerInterfacesDefaultFabric(serverInterfacesDefaultFabric).Execute()

Sets the default fabric for the specified server interfaces



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
	serverId := float32(8.14) // float32 | 
	serverInterfacesDefaultFabric := *openapiclient.NewServerInterfacesDefaultFabric([]float32{float32(123)}, NullableFloat32(1)) // ServerInterfacesDefaultFabric | The server interfaces default fabric option

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.SetServerInterfacesDefaultFabric(context.Background(), serverId).ServerInterfacesDefaultFabric(serverInterfacesDefaultFabric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.SetServerInterfacesDefaultFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetServerInterfacesDefaultFabricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverInterfacesDefaultFabric** | [**ServerInterfacesDefaultFabric**](ServerInterfacesDefaultFabric.md) | The server interfaces default fabric option | 

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


## SetServerInterfacesRedundancyGroup

> SetServerInterfacesRedundancyGroup(ctx, serverId).ServerInterfacesRedundancyGroup(serverInterfacesRedundancyGroup).Execute()

Sets the redundancy group index for the specified server interfaces



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
	serverId := float32(8.14) // float32 | 
	serverInterfacesRedundancyGroup := *openapiclient.NewServerInterfacesRedundancyGroup([]float32{float32(123)}, NullableFloat32(1)) // ServerInterfacesRedundancyGroup | The server interfaces redundancy group option

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.SetServerInterfacesRedundancyGroup(context.Background(), serverId).ServerInterfacesRedundancyGroup(serverInterfacesRedundancyGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.SetServerInterfacesRedundancyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetServerInterfacesRedundancyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverInterfacesRedundancyGroup** | [**ServerInterfacesRedundancyGroup**](ServerInterfacesRedundancyGroup.md) | The server interfaces redundancy group option | 

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


## SetServerPowerState

> SetServerPowerState(ctx, serverId).ServerPowerSet(serverPowerSet).IfMatch(ifMatch).Execute()

Sets the power state of a server



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
	serverId := float32(8.14) // float32 | 
	serverPowerSet := *openapiclient.NewServerPowerSet("PowerCommand_example") // ServerPowerSet | The server power options
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.SetServerPowerState(context.Background(), serverId).ServerPowerSet(serverPowerSet).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.SetServerPowerState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetServerPowerStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverPowerSet** | [**ServerPowerSet**](ServerPowerSet.md) | The server power options | 
 **ifMatch** | **string** | Entity tag | 

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


## UpdateServer

> Server UpdateServer(ctx, serverId).UpdateServer(updateServer).IfMatch(ifMatch).Execute()

Updates the server information



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
	serverId := float32(8.14) // float32 | 
	updateServer := *openapiclient.NewUpdateServer() // UpdateServer | The server information update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.UpdateServer(context.Background(), serverId).UpdateServer(updateServer).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.UpdateServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.UpdateServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServer** | [**UpdateServer**](UpdateServer.md) | The server information update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Server**](Server.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerEnableSnmp

> float32 UpdateServerEnableSnmp(ctx, serverId).IfMatch(ifMatch).Execute()

Enables SNMP on a Server



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
	serverId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.UpdateServerEnableSnmp(context.Background(), serverId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.UpdateServerEnableSnmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerEnableSnmp`: float32
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.UpdateServerEnableSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerEnableSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

**float32**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerIpmiCredentials

> ServerCredentials UpdateServerIpmiCredentials(ctx, serverId).UpdateServerIpmiCredentials(updateServerIpmiCredentials).IfMatch(ifMatch).Execute()

Update Server ipmi credentials



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
	serverId := float32(8.14) // float32 | 
	updateServerIpmiCredentials := *openapiclient.NewUpdateServerIpmiCredentials(false) // UpdateServerIpmiCredentials | The Server Ipmi credentials object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.UpdateServerIpmiCredentials(context.Background(), serverId).UpdateServerIpmiCredentials(updateServerIpmiCredentials).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.UpdateServerIpmiCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerIpmiCredentials`: ServerCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.UpdateServerIpmiCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerIpmiCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerIpmiCredentials** | [**UpdateServerIpmiCredentials**](UpdateServerIpmiCredentials.md) | The Server Ipmi credentials object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ServerCredentials**](ServerCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

