# \EndpointInstanceGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpointInstanceGroup**](EndpointInstanceGroupAPI.md#CreateEndpointInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/endpoint-instance-groups | Add a Endpoint Instance Group to an infrastructure. By default it will not have any instance.
[**CreateEndpointInstanceGroupLogicalNetworkACL**](EndpointInstanceGroupAPI.md#CreateEndpointInstanceGroupLogicalNetworkACL) | **Post** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId}/security/rules | Create a security rule for a logical network
[**CreateEndpointInstanceGroupNetworkConfigurationConnection**](EndpointInstanceGroupAPI.md#CreateEndpointInstanceGroupNetworkConfigurationConnection) | **Post** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections | Create a network connection for a endpoint instance group
[**DeleteEndpointInstanceGroup**](EndpointInstanceGroupAPI.md#DeleteEndpointInstanceGroup) | **Delete** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId} | Delete Endpoint Instance Group. Will not take effect if there are instances in this group.
[**DeleteEndpointInstanceGroupLogicalNetworkACL**](EndpointInstanceGroupAPI.md#DeleteEndpointInstanceGroupLogicalNetworkACL) | **Delete** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId}/security/rules/{ruleId} | Delete a security rule for a logical network
[**DeleteEndpointInstanceGroupNetworkConfigurationConnection**](EndpointInstanceGroupAPI.md#DeleteEndpointInstanceGroupNetworkConfigurationConnection) | **Delete** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId} | Delete a network connection for a endpoint instance group
[**GetEndpointInstanceGroup**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroup) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId} | Get Endpoint Instance Groups details
[**GetEndpointInstanceGroupConfig**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroupConfig) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config | Get Endpoint Instance Group config details
[**GetEndpointInstanceGroupEndpointInstances**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroupEndpointInstances) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/endpoint-instances | List Endpoint Instances for a Endpoint Instance Group
[**GetEndpointInstanceGroupLogicalNetworkACL**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroupLogicalNetworkACL) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId}/security/rules | Get the security rules for a logical network
[**GetEndpointInstanceGroupLogicalNetworkACLById**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroupLogicalNetworkACLById) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId}/security/rules/{ruleId} | Get a security rule for a logical network by id
[**GetEndpointInstanceGroupNetworkConfiguration**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroupNetworkConfiguration) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking | Get endpoint instance group network configuration
[**GetEndpointInstanceGroupNetworkConfigurationConnectionById**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroupNetworkConfigurationConnectionById) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId} | Get endpoint instance group network configuration connection by id
[**GetEndpointInstanceGroupNetworkConfigurationConnections**](EndpointInstanceGroupAPI.md#GetEndpointInstanceGroupNetworkConfigurationConnections) | **Get** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections | Get endpoint instance group&#39;s network connections
[**GetInfrastructureEndpointInstanceGroups**](EndpointInstanceGroupAPI.md#GetInfrastructureEndpointInstanceGroups) | **Get** /api/v2/infrastructures/{infrastructureId}/endpoint-instance-groups | List Endpoint Instance Groups for an infrastructure
[**UpdateEndpointInstanceGroupConfig**](EndpointInstanceGroupAPI.md#UpdateEndpointInstanceGroupConfig) | **Patch** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config | Updates Endpoint Instance Group configuration
[**UpdateEndpointInstanceGroupLogicalNetworkACL**](EndpointInstanceGroupAPI.md#UpdateEndpointInstanceGroupLogicalNetworkACL) | **Patch** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId}/security/rules/{ruleId} | Update a security rule for a logical network
[**UpdateEndpointInstanceGroupMeta**](EndpointInstanceGroupAPI.md#UpdateEndpointInstanceGroupMeta) | **Patch** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/meta | Update an Endpoint Instance Group meta information
[**UpdateEndpointInstanceGroupNetworkConfiguration**](EndpointInstanceGroupAPI.md#UpdateEndpointInstanceGroupNetworkConfiguration) | **Put** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking | Create or return the network configuration of the specified endpoint instance group
[**UpdateEndpointInstanceGroupNetworkConfigurationConnection**](EndpointInstanceGroupAPI.md#UpdateEndpointInstanceGroupNetworkConfigurationConnection) | **Patch** /api/v2/endpoint-instance-groups/{endpointInstanceGroupId}/config/networking/connections/{connectionId} | Update a network connection for a endpoint instance group



## CreateEndpointInstanceGroup

> EndpointInstanceGroup CreateEndpointInstanceGroup(ctx, infrastructureId).EndpointInstanceGroupCreate(endpointInstanceGroupCreate).Execute()

Add a Endpoint Instance Group to an infrastructure. By default it will not have any instance.

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
	endpointInstanceGroupCreate := *openapiclient.NewEndpointInstanceGroupCreate() // EndpointInstanceGroupCreate | The Endpoint Instance Group to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.CreateEndpointInstanceGroup(context.Background(), infrastructureId).EndpointInstanceGroupCreate(endpointInstanceGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.CreateEndpointInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointInstanceGroup`: EndpointInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.CreateEndpointInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointInstanceGroupCreate** | [**EndpointInstanceGroupCreate**](EndpointInstanceGroupCreate.md) | The Endpoint Instance Group to create | 

### Return type

[**EndpointInstanceGroup**](EndpointInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpointInstanceGroupLogicalNetworkACL

> LogicalNetworkACL CreateEndpointInstanceGroupLogicalNetworkACL(ctx, endpointInstanceGroupId, connectionId).CreateLogicalNetworkACL(createLogicalNetworkACL).Execute()

Create a security rule for a logical network

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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	createLogicalNetworkACL := *openapiclient.NewCreateLogicalNetworkACL(openapiclient.ACLType("ipv4"), openapiclient.ACLDirection("in"), int32(1), openapiclient.ACLForwardingAction("allow"), openapiclient.ACLEnforcementPoint("svi")) // CreateLogicalNetworkACL | The security rule to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.CreateEndpointInstanceGroupLogicalNetworkACL(context.Background(), endpointInstanceGroupId, connectionId).CreateLogicalNetworkACL(createLogicalNetworkACL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.CreateEndpointInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointInstanceGroupLogicalNetworkACL`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.CreateEndpointInstanceGroupLogicalNetworkACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLogicalNetworkACL** | [**CreateLogicalNetworkACL**](CreateLogicalNetworkACL.md) | The security rule to create | 

### Return type

[**LogicalNetworkACL**](LogicalNetworkACL.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpointInstanceGroupNetworkConfigurationConnection

> EndpointInstanceGroupNetworkConnection CreateEndpointInstanceGroupNetworkConfigurationConnection(ctx, endpointInstanceGroupId).CreateEndpointInstanceGroupNetworkConnection(createEndpointInstanceGroupNetworkConnection).Execute()

Create a network connection for a endpoint instance group

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
	endpointInstanceGroupId := int32(56) // int32 | 
	createEndpointInstanceGroupNetworkConnection := *openapiclient.NewCreateEndpointInstanceGroupNetworkConnection("1", true, openapiclient.NetworkEndpointGroupAllowedAccessMode("l2")) // CreateEndpointInstanceGroupNetworkConnection | The network connection object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.CreateEndpointInstanceGroupNetworkConfigurationConnection(context.Background(), endpointInstanceGroupId).CreateEndpointInstanceGroupNetworkConnection(createEndpointInstanceGroupNetworkConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.CreateEndpointInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointInstanceGroupNetworkConfigurationConnection`: EndpointInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.CreateEndpointInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEndpointInstanceGroupNetworkConnection** | [**CreateEndpointInstanceGroupNetworkConnection**](CreateEndpointInstanceGroupNetworkConnection.md) | The network connection object to create | 

### Return type

[**EndpointInstanceGroupNetworkConnection**](EndpointInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpointInstanceGroup

> DeleteEndpointInstanceGroup(ctx, endpointInstanceGroupId).IfMatch(ifMatch).Execute()

Delete Endpoint Instance Group. Will not take effect if there are instances in this group.



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
	endpointInstanceGroupId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceGroupAPI.DeleteEndpointInstanceGroup(context.Background(), endpointInstanceGroupId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.DeleteEndpointInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointInstanceGroupRequest struct via the builder pattern


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


## DeleteEndpointInstanceGroupLogicalNetworkACL

> DeleteEndpointInstanceGroupLogicalNetworkACL(ctx, endpointInstanceGroupId, connectionId, ruleId).Execute()

Delete a security rule for a logical network

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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	ruleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceGroupAPI.DeleteEndpointInstanceGroupLogicalNetworkACL(context.Background(), endpointInstanceGroupId, connectionId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.DeleteEndpointInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


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


## DeleteEndpointInstanceGroupNetworkConfigurationConnection

> DeleteEndpointInstanceGroupNetworkConfigurationConnection(ctx, endpointInstanceGroupId, connectionId).Execute()

Delete a network connection for a endpoint instance group

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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceGroupAPI.DeleteEndpointInstanceGroupNetworkConfigurationConnection(context.Background(), endpointInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.DeleteEndpointInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


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


## GetEndpointInstanceGroup

> EndpointInstanceGroup GetEndpointInstanceGroup(ctx, endpointInstanceGroupId).Execute()

Get Endpoint Instance Groups details



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
	endpointInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroup(context.Background(), endpointInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroup`: EndpointInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointInstanceGroup**](EndpointInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceGroupConfig

> EndpointInstanceGroupConfiguration GetEndpointInstanceGroupConfig(ctx, endpointInstanceGroupId).Execute()

Get Endpoint Instance Group config details



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
	endpointInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroupConfig(context.Background(), endpointInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroupConfig`: EndpointInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointInstanceGroupConfiguration**](EndpointInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceGroupEndpointInstances

> EndpointInstancePaginatedList GetEndpointInstanceGroupEndpointInstances(ctx, endpointInstanceGroupId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Endpoint Instances for a Endpoint Instance Group



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
	endpointInstanceGroupId := int32(56) // int32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterGroupId := []string{"Inner_example"} // []string | Filter by groupId query param.           <p>              <b>Format: </b> filter.groupId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.groupId=$not:$like:John Doe&filter.groupId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterEndpointId := []string{"Inner_example"} // []string | Filter by endpointId query param.           <p>              <b>Format: </b> filter.endpointId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.endpointId=$not:$like:John Doe&filter.endpointId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigEndpointId := []string{"Inner_example"} // []string | Filter by config.endpointId query param.           <p>              <b>Format: </b> filter.config.endpointId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.endpointId=$not:$like:John Doe&filter.config.endpointId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroupEndpointInstances(context.Background(), endpointInstanceGroupId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroupEndpointInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroupEndpointInstances`: EndpointInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroupEndpointInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupEndpointInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterGroupId** | **[]string** | Filter by groupId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.groupId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.groupId&#x3D;$not:$like:John Doe&amp;filter.groupId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterEndpointId** | **[]string** | Filter by endpointId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.endpointId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.endpointId&#x3D;$not:$like:John Doe&amp;filter.endpointId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigEndpointId** | **[]string** | Filter by config.endpointId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.endpointId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.endpointId&#x3D;$not:$like:John Doe&amp;filter.config.endpointId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

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


## GetEndpointInstanceGroupLogicalNetworkACL

> LogicalNetworkACL GetEndpointInstanceGroupLogicalNetworkACL(ctx, endpointInstanceGroupId, connectionId).Execute()

Get the security rules for a logical network

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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroupLogicalNetworkACL(context.Background(), endpointInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroupLogicalNetworkACL`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroupLogicalNetworkACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LogicalNetworkACL**](LogicalNetworkACL.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceGroupLogicalNetworkACLById

> LogicalNetworkACL GetEndpointInstanceGroupLogicalNetworkACLById(ctx, endpointInstanceGroupId, connectionId, ruleId).Execute()

Get a security rule for a logical network by id

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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	ruleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroupLogicalNetworkACLById(context.Background(), endpointInstanceGroupId, connectionId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroupLogicalNetworkACLById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroupLogicalNetworkACLById`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroupLogicalNetworkACLById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupLogicalNetworkACLByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LogicalNetworkACL**](LogicalNetworkACL.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceGroupNetworkConfiguration

> NetworkEndpointGroup GetEndpointInstanceGroupNetworkConfiguration(ctx, endpointInstanceGroupId).Execute()

Get endpoint instance group network configuration



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
	endpointInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfiguration(context.Background(), endpointInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroupNetworkConfiguration`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupNetworkConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkEndpointGroup**](NetworkEndpointGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceGroupNetworkConfigurationConnectionById

> EndpointInstanceGroupNetworkConnection GetEndpointInstanceGroupNetworkConfigurationConnectionById(ctx, endpointInstanceGroupId, connectionId).Execute()

Get endpoint instance group network configuration connection by id



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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfigurationConnectionById(context.Background(), endpointInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfigurationConnectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroupNetworkConfigurationConnectionById`: EndpointInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfigurationConnectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupNetworkConfigurationConnectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EndpointInstanceGroupNetworkConnection**](EndpointInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceGroupNetworkConfigurationConnections

> EndpointInstanceGroupNetworkConnectionsList GetEndpointInstanceGroupNetworkConfigurationConnections(ctx, endpointInstanceGroupId).Execute()

Get endpoint instance group's network connections

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
	endpointInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfigurationConnections(context.Background(), endpointInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfigurationConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceGroupNetworkConfigurationConnections`: EndpointInstanceGroupNetworkConnectionsList
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetEndpointInstanceGroupNetworkConfigurationConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceGroupNetworkConfigurationConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointInstanceGroupNetworkConnectionsList**](EndpointInstanceGroupNetworkConnectionsList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureEndpointInstanceGroups

> EndpointInstanceGroupPaginatedList GetInfrastructureEndpointInstanceGroups(ctx, infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterExtensionInstanceId(filterExtensionInstanceId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Endpoint Instance Groups for an infrastructure



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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterExtensionInstanceId := []string{"Inner_example"} // []string | Filter by extensionInstanceId query param.           <p>              <b>Format: </b> filter.extensionInstanceId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.extensionInstanceId=$not:$like:John Doe&filter.extensionInstanceId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.GetInfrastructureEndpointInstanceGroups(context.Background(), infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterExtensionInstanceId(filterExtensionInstanceId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.GetInfrastructureEndpointInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureEndpointInstanceGroups`: EndpointInstanceGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.GetInfrastructureEndpointInstanceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureEndpointInstanceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterExtensionInstanceId** | **[]string** | Filter by extensionInstanceId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.extensionInstanceId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.extensionInstanceId&#x3D;$not:$like:John Doe&amp;filter.extensionInstanceId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**EndpointInstanceGroupPaginatedList**](EndpointInstanceGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceGroupConfig

> EndpointInstanceGroupConfiguration UpdateEndpointInstanceGroupConfig(ctx, endpointInstanceGroupId).EndpointInstanceGroupUpdate(endpointInstanceGroupUpdate).IfMatch(ifMatch).Execute()

Updates Endpoint Instance Group configuration



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
	endpointInstanceGroupId := int32(56) // int32 | 
	endpointInstanceGroupUpdate := *openapiclient.NewEndpointInstanceGroupUpdate() // EndpointInstanceGroupUpdate | The Endpoint Instance Group configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupConfig(context.Background(), endpointInstanceGroupId).EndpointInstanceGroupUpdate(endpointInstanceGroupUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointInstanceGroupConfig`: EndpointInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointInstanceGroupUpdate** | [**EndpointInstanceGroupUpdate**](EndpointInstanceGroupUpdate.md) | The Endpoint Instance Group configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**EndpointInstanceGroupConfiguration**](EndpointInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceGroupLogicalNetworkACL

> LogicalNetworkACL UpdateEndpointInstanceGroupLogicalNetworkACL(ctx, endpointInstanceGroupId, connectionId, ruleId).UpdateLogicalNetworkACL(updateLogicalNetworkACL).Execute()

Update a security rule for a logical network

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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	ruleId := int32(56) // int32 | 
	updateLogicalNetworkACL := *openapiclient.NewUpdateLogicalNetworkACL() // UpdateLogicalNetworkACL | The security rule to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupLogicalNetworkACL(context.Background(), endpointInstanceGroupId, connectionId, ruleId).UpdateLogicalNetworkACL(updateLogicalNetworkACL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointInstanceGroupLogicalNetworkACL`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupLogicalNetworkACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateLogicalNetworkACL** | [**UpdateLogicalNetworkACL**](UpdateLogicalNetworkACL.md) | The security rule to update | 

### Return type

[**LogicalNetworkACL**](LogicalNetworkACL.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceGroupMeta

> UpdateEndpointInstanceGroupMeta(ctx, endpointInstanceGroupId).GenericMeta(genericMeta).Execute()

Update an Endpoint Instance Group meta information



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
	endpointInstanceGroupId := int32(56) // int32 | 
	genericMeta := *openapiclient.NewGenericMeta() // GenericMeta | The Endpoint Instance Group meta information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupMeta(context.Background(), endpointInstanceGroupId).GenericMeta(genericMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceGroupMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genericMeta** | [**GenericMeta**](GenericMeta.md) | The Endpoint Instance Group meta information | 

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


## UpdateEndpointInstanceGroupNetworkConfiguration

> NetworkEndpointGroup UpdateEndpointInstanceGroupNetworkConfiguration(ctx, endpointInstanceGroupId).Execute()

Create or return the network configuration of the specified endpoint instance group

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
	endpointInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupNetworkConfiguration(context.Background(), endpointInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupNetworkConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointInstanceGroupNetworkConfiguration`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupNetworkConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceGroupNetworkConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkEndpointGroup**](NetworkEndpointGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceGroupNetworkConfigurationConnection

> EndpointInstanceGroupNetworkConnection UpdateEndpointInstanceGroupNetworkConfigurationConnection(ctx, endpointInstanceGroupId, connectionId).UpdateNetworkEndpointGroupLogicalNetwork(updateNetworkEndpointGroupLogicalNetwork).Execute()

Update a network connection for a endpoint instance group

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
	endpointInstanceGroupId := int32(56) // int32 | 
	connectionId := float32(8.14) // float32 | 
	updateNetworkEndpointGroupLogicalNetwork := *openapiclient.NewUpdateNetworkEndpointGroupLogicalNetwork() // UpdateNetworkEndpointGroupLogicalNetwork | The network connection object to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupNetworkConfigurationConnection(context.Background(), endpointInstanceGroupId, connectionId).UpdateNetworkEndpointGroupLogicalNetwork(updateNetworkEndpointGroupLogicalNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointInstanceGroupNetworkConfigurationConnection`: EndpointInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceGroupAPI.UpdateEndpointInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceGroupId** | **int32** |  | 
**connectionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkEndpointGroupLogicalNetwork** | [**UpdateNetworkEndpointGroupLogicalNetwork**](UpdateNetworkEndpointGroupLogicalNetwork.md) | The network connection object to update | 

### Return type

[**EndpointInstanceGroupNetworkConnection**](EndpointInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

