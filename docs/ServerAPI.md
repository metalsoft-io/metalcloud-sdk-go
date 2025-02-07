# \ServerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecommissionServer**](ServerAPI.md#DecommissionServer) | **Post** /api/v2/servers/{serverId}/actions/decommission | Decommissions a Server
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
[**ReRegisterServer**](ServerAPI.md#ReRegisterServer) | **Post** /api/v2/servers/{serverId}/actions/re-register | Re-register a server
[**RegisterServer**](ServerAPI.md#RegisterServer) | **Post** /api/v2/servers | Initialize server registration
[**ResetServerToFactoryDefaults**](ServerAPI.md#ResetServerToFactoryDefaults) | **Post** /api/v2/servers/{serverId}/actions/factory-reset | Resets a server to factory defaults
[**SetServerPowerState**](ServerAPI.md#SetServerPowerState) | **Post** /api/v2/servers/{serverId}/actions/set-power | Sets the power state of a server
[**UpdateServer**](ServerAPI.md#UpdateServer) | **Patch** /api/v2/servers/{serverId} | Updates the server information
[**UpdateServerEnableSnmp**](ServerAPI.md#UpdateServerEnableSnmp) | **Post** /api/v2/servers/{serverId}/actions/enable-snmp | Enables SNMP on a Server
[**UpdateServerIpmiCredentials**](ServerAPI.md#UpdateServerIpmiCredentials) | **Post** /api/v2/servers/{serverId}/actions/update-ipmi-credentials | Update Server ipmi credentials



## DecommissionServer

> DecommissionServer(ctx, serverId).IfMatch(ifMatch).Execute()

Decommissions a Server



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
	r, err := apiClient.ServerAPI.DecommissionServer(context.Background(), serverId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.DecommissionServer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDecommissionServerRequest struct via the builder pattern


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

> ServerPaginatedList GetServers(ctx).Page(page).Limit(limit).FilterServerTypeId(filterServerTypeId).FilterServerId(filterServerId).FilterSiteId(filterSiteId).FilterVendor(filterVendor).FilterSerialNumber(filterSerialNumber).FilterManagementAddress(filterManagementAddress).FilterModel(filterModel).FilterAdministrationState(filterAdministrationState).FilterServerDiskWipe(filterServerDiskWipe).FilterPowerStatus(filterPowerStatus).FilterServerDhcpStatus(filterServerDhcpStatus).FilterServerClass(filterServerClass).FilterServerStatus(filterServerStatus).FilterRequiresManualCleaning(filterRequiresManualCleaning).FilterInstanceInfrastructureInfrastructureId(filterInstanceInfrastructureInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterServerTypeId := []string{"Inner_example"} // []string | Filter by serverTypeId query param.           <p>              <b>Format: </b> filter.serverTypeId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverTypeId=$not:$like:John Doe&filter.serverTypeId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.           <p>              <b>Format: </b> filter.serverId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverId=$not:$like:John Doe&filter.serverId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterVendor := []string{"Inner_example"} // []string | Filter by vendor query param.           <p>              <b>Format: </b> filter.vendor={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.vendor=$not:$like:John Doe&filter.vendor=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSerialNumber := []string{"Inner_example"} // []string | Filter by serialNumber query param.           <p>              <b>Format: </b> filter.serialNumber={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serialNumber=$not:$like:John Doe&filter.serialNumber=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterManagementAddress := []string{"Inner_example"} // []string | Filter by managementAddress query param.           <p>              <b>Format: </b> filter.managementAddress={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.managementAddress=$not:$like:John Doe&filter.managementAddress=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterModel := []string{"Inner_example"} // []string | Filter by model query param.           <p>              <b>Format: </b> filter.model={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.model=$not:$like:John Doe&filter.model=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAdministrationState := []string{"Inner_example"} // []string | Filter by administrationState query param.           <p>              <b>Format: </b> filter.administrationState={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.administrationState=$not:$like:John Doe&filter.administrationState=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServerDiskWipe := []string{"Inner_example"} // []string | Filter by serverDiskWipe query param.           <p>              <b>Format: </b> filter.serverDiskWipe={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverDiskWipe=$not:$like:John Doe&filter.serverDiskWipe=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterPowerStatus := []string{"Inner_example"} // []string | Filter by powerStatus query param.           <p>              <b>Format: </b> filter.powerStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.powerStatus=$not:$like:John Doe&filter.powerStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServerDhcpStatus := []string{"Inner_example"} // []string | Filter by serverDhcpStatus query param.           <p>              <b>Format: </b> filter.serverDhcpStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverDhcpStatus=$not:$like:John Doe&filter.serverDhcpStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServerClass := []string{"Inner_example"} // []string | Filter by serverClass query param.           <p>              <b>Format: </b> filter.serverClass={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverClass=$not:$like:John Doe&filter.serverClass=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServerStatus := []string{"Inner_example"} // []string | Filter by serverStatus query param.           <p>              <b>Format: </b> filter.serverStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverStatus=$not:$like:John Doe&filter.serverStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterRequiresManualCleaning := []string{"Inner_example"} // []string | Filter by requiresManualCleaning query param.           <p>              <b>Format: </b> filter.requiresManualCleaning={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.requiresManualCleaning=$not:$like:John Doe&filter.requiresManualCleaning=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterInstanceInfrastructureInfrastructureId := []string{"Inner_example"} // []string | Filter by instance.infrastructure.infrastructureId query param.           <p>              <b>Format: </b> filter.instance.infrastructure.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.instance.infrastructure.infrastructureId=$not:$like:John Doe&filter.instance.infrastructure.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> serverId:DESC           </p>       <h4>Available Fields</h4><ul><li>serverId</li> <li>serverTypeId</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> vendor,model           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>vendor</li> <li>model</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServers(context.Background()).Page(page).Limit(limit).FilterServerTypeId(filterServerTypeId).FilterServerId(filterServerId).FilterSiteId(filterSiteId).FilterVendor(filterVendor).FilterSerialNumber(filterSerialNumber).FilterManagementAddress(filterManagementAddress).FilterModel(filterModel).FilterAdministrationState(filterAdministrationState).FilterServerDiskWipe(filterServerDiskWipe).FilterPowerStatus(filterPowerStatus).FilterServerDhcpStatus(filterServerDhcpStatus).FilterServerClass(filterServerClass).FilterServerStatus(filterServerStatus).FilterRequiresManualCleaning(filterRequiresManualCleaning).FilterInstanceInfrastructureInfrastructureId(filterInstanceInfrastructureInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterServerTypeId** | **[]string** | Filter by serverTypeId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverTypeId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverTypeId&#x3D;$not:$like:John Doe&amp;filter.serverTypeId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerId** | **[]string** | Filter by serverId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverId&#x3D;$not:$like:John Doe&amp;filter.serverId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterVendor** | **[]string** | Filter by vendor query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.vendor&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.vendor&#x3D;$not:$like:John Doe&amp;filter.vendor&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSerialNumber** | **[]string** | Filter by serialNumber query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serialNumber&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serialNumber&#x3D;$not:$like:John Doe&amp;filter.serialNumber&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterManagementAddress** | **[]string** | Filter by managementAddress query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.managementAddress&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.managementAddress&#x3D;$not:$like:John Doe&amp;filter.managementAddress&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterModel** | **[]string** | Filter by model query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.model&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.model&#x3D;$not:$like:John Doe&amp;filter.model&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterAdministrationState** | **[]string** | Filter by administrationState query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.administrationState&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.administrationState&#x3D;$not:$like:John Doe&amp;filter.administrationState&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerDiskWipe** | **[]string** | Filter by serverDiskWipe query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverDiskWipe&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverDiskWipe&#x3D;$not:$like:John Doe&amp;filter.serverDiskWipe&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterPowerStatus** | **[]string** | Filter by powerStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.powerStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.powerStatus&#x3D;$not:$like:John Doe&amp;filter.powerStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerDhcpStatus** | **[]string** | Filter by serverDhcpStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverDhcpStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverDhcpStatus&#x3D;$not:$like:John Doe&amp;filter.serverDhcpStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerClass** | **[]string** | Filter by serverClass query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverClass&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverClass&#x3D;$not:$like:John Doe&amp;filter.serverClass&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerStatus** | **[]string** | Filter by serverStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverStatus&#x3D;$not:$like:John Doe&amp;filter.serverStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterRequiresManualCleaning** | **[]string** | Filter by requiresManualCleaning query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.requiresManualCleaning&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.requiresManualCleaning&#x3D;$not:$like:John Doe&amp;filter.requiresManualCleaning&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterInstanceInfrastructureInfrastructureId** | **[]string** | Filter by instance.infrastructure.infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.instance.infrastructure.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.instance.infrastructure.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.instance.infrastructure.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; serverId:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;serverId&lt;/li&gt; &lt;li&gt;serverTypeId&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; vendor,model           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;vendor&lt;/li&gt; &lt;li&gt;model&lt;/li&gt;&lt;/ul&gt;          | 

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
	updateServerIpmiCredentials := *openapiclient.NewUpdateServerIpmiCredentials("Host_example", "Username_example", "Password_example", false) // UpdateServerIpmiCredentials | The Server Ipmi credentials object
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

