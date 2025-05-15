# \ServerInstanceGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerInstanceGroup**](ServerInstanceGroupAPI.md#CreateServerInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/server-instance-groups | Add a Server Instance Group to an infrastructure. By default it will not have any instance.
[**CreateServerInstanceGroupLogicalNetworkACL**](ServerInstanceGroupAPI.md#CreateServerInstanceGroupLogicalNetworkACL) | **Post** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId}/security/rules | Create a security rule for a logical network
[**CreateServerInstanceGroupNetworkConfigurationConnection**](ServerInstanceGroupAPI.md#CreateServerInstanceGroupNetworkConfigurationConnection) | **Post** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections | Create a network connection for a server instance group
[**DeleteServerInstanceGroup**](ServerInstanceGroupAPI.md#DeleteServerInstanceGroup) | **Delete** /api/v2/server-instance-groups/{serverInstanceGroupId} | Delete Server Instance Group. Will not take effect if there are instances in this group.
[**DeleteServerInstanceGroupLogicalNetworkACL**](ServerInstanceGroupAPI.md#DeleteServerInstanceGroupLogicalNetworkACL) | **Delete** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId}/security/rules/{ruleId} | Delete a security rule for a logical network
[**DeleteServerInstanceGroupNetworkConfigurationConnection**](ServerInstanceGroupAPI.md#DeleteServerInstanceGroupNetworkConfigurationConnection) | **Delete** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId} | Delete a network connection for a server instance group
[**GetInfrastructureServerInstanceGroups**](ServerInstanceGroupAPI.md#GetInfrastructureServerInstanceGroups) | **Get** /api/v2/infrastructures/{infrastructureId}/server-instance-groups | List Server Instance Groups for an infrastructure
[**GetServerInstanceGroup**](ServerInstanceGroupAPI.md#GetServerInstanceGroup) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId} | Get Server Instance Groups details
[**GetServerInstanceGroupConfig**](ServerInstanceGroupAPI.md#GetServerInstanceGroupConfig) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/config | Get Server Instance Group config details
[**GetServerInstanceGroupDriveGroups**](ServerInstanceGroupAPI.md#GetServerInstanceGroupDriveGroups) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/drive-groups | Get Server Instance Group Drive Groups
[**GetServerInstanceGroupInterface**](ServerInstanceGroupAPI.md#GetServerInstanceGroupInterface) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/interfaces/{interfaceId} | Get Server Instance Group Interface details
[**GetServerInstanceGroupInterfaces**](ServerInstanceGroupAPI.md#GetServerInstanceGroupInterfaces) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/interfaces | Get Server Instance Group Interfaces
[**GetServerInstanceGroupLogicalNetworkACL**](ServerInstanceGroupAPI.md#GetServerInstanceGroupLogicalNetworkACL) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId}/security/rules | Get the security rules for a logical network
[**GetServerInstanceGroupLogicalNetworkACLById**](ServerInstanceGroupAPI.md#GetServerInstanceGroupLogicalNetworkACLById) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId}/security/rules/{ruleId} | Get a security rule for a logical network by id
[**GetServerInstanceGroupNetworkConfiguration**](ServerInstanceGroupAPI.md#GetServerInstanceGroupNetworkConfiguration) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking | Get server instance group network configuration
[**GetServerInstanceGroupNetworkConfigurationConnectionById**](ServerInstanceGroupAPI.md#GetServerInstanceGroupNetworkConfigurationConnectionById) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId} | Get server instance group network configuration connection by id
[**GetServerInstanceGroupNetworkConfigurationConnections**](ServerInstanceGroupAPI.md#GetServerInstanceGroupNetworkConfigurationConnections) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections | Get server instance group&#39;s network connections
[**GetServerInstanceGroupServerInstances**](ServerInstanceGroupAPI.md#GetServerInstanceGroupServerInstances) | **Get** /api/v2/server-instance-groups/{serverInstanceGroupId}/server-instances | List Server Instances for a Server Instance Group
[**UpdateServerInstanceGroupConfig**](ServerInstanceGroupAPI.md#UpdateServerInstanceGroupConfig) | **Patch** /api/v2/server-instance-groups/{serverInstanceGroupId}/config | Updates Server Instance Group configuration
[**UpdateServerInstanceGroupLogicalNetworkACL**](ServerInstanceGroupAPI.md#UpdateServerInstanceGroupLogicalNetworkACL) | **Patch** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId}/security/rules/{ruleId} | Update a security rule for a logical network
[**UpdateServerInstanceGroupMeta**](ServerInstanceGroupAPI.md#UpdateServerInstanceGroupMeta) | **Patch** /api/v2/server-instance-groups/{serverInstanceGroupId}/meta | Update an Server Instance Group meta information
[**UpdateServerInstanceGroupNetworkConfiguration**](ServerInstanceGroupAPI.md#UpdateServerInstanceGroupNetworkConfiguration) | **Put** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking | Create or return the network configuration of the specified server instance group
[**UpdateServerInstanceGroupNetworkConfigurationConnection**](ServerInstanceGroupAPI.md#UpdateServerInstanceGroupNetworkConfigurationConnection) | **Patch** /api/v2/server-instance-groups/{serverInstanceGroupId}/config/networking/connections/{connectionId} | Update a network connection for a server instance group



## CreateServerInstanceGroup

> ServerInstanceGroup CreateServerInstanceGroup(ctx, infrastructureId).ServerInstanceGroupCreate(serverInstanceGroupCreate).Execute()

Add a Server Instance Group to an infrastructure. By default it will not have any instance.

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
	serverInstanceGroupCreate := *openapiclient.NewServerInstanceGroupCreate(int32(5)) // ServerInstanceGroupCreate | The Server Instance Group to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.CreateServerInstanceGroup(context.Background(), infrastructureId).ServerInstanceGroupCreate(serverInstanceGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.CreateServerInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerInstanceGroup`: ServerInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.CreateServerInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverInstanceGroupCreate** | [**ServerInstanceGroupCreate**](ServerInstanceGroupCreate.md) | The Server Instance Group to create | 

### Return type

[**ServerInstanceGroup**](ServerInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerInstanceGroupLogicalNetworkACL

> LogicalNetworkACL CreateServerInstanceGroupLogicalNetworkACL(ctx, serverInstanceGroupId, connectionId).CreateLogicalNetworkACL(createLogicalNetworkACL).Execute()

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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	createLogicalNetworkACL := *openapiclient.NewCreateLogicalNetworkACL(openapiclient.ACLType("ipv4"), openapiclient.ACLDirection("in"), int32(1), openapiclient.ACLForwardingAction("allow"), openapiclient.ACLEnforcementPoint("svi")) // CreateLogicalNetworkACL | The security rule to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.CreateServerInstanceGroupLogicalNetworkACL(context.Background(), serverInstanceGroupId, connectionId).CreateLogicalNetworkACL(createLogicalNetworkACL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.CreateServerInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerInstanceGroupLogicalNetworkACL`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.CreateServerInstanceGroupLogicalNetworkACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


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


## CreateServerInstanceGroupNetworkConfigurationConnection

> ServerInstanceGroupNetworkConnection CreateServerInstanceGroupNetworkConfigurationConnection(ctx, serverInstanceGroupId).CreateServerInstanceGroupNetworkConnection(createServerInstanceGroupNetworkConnection).Execute()

Create a network connection for a server instance group

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
	serverInstanceGroupId := int32(56) // int32 | 
	createServerInstanceGroupNetworkConnection := *openapiclient.NewCreateServerInstanceGroupNetworkConnection("1", true, openapiclient.NetworkEndpointGroupAllowedAccessMode("l2")) // CreateServerInstanceGroupNetworkConnection | The network connection object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.CreateServerInstanceGroupNetworkConfigurationConnection(context.Background(), serverInstanceGroupId).CreateServerInstanceGroupNetworkConnection(createServerInstanceGroupNetworkConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.CreateServerInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerInstanceGroupNetworkConfigurationConnection`: ServerInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.CreateServerInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServerInstanceGroupNetworkConnection** | [**CreateServerInstanceGroupNetworkConnection**](CreateServerInstanceGroupNetworkConnection.md) | The network connection object to create | 

### Return type

[**ServerInstanceGroupNetworkConnection**](ServerInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerInstanceGroup

> DeleteServerInstanceGroup(ctx, serverInstanceGroupId).IfMatch(ifMatch).Execute()

Delete Server Instance Group. Will not take effect if there are instances in this group.



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
	serverInstanceGroupId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerInstanceGroupAPI.DeleteServerInstanceGroup(context.Background(), serverInstanceGroupId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.DeleteServerInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerInstanceGroupRequest struct via the builder pattern


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


## DeleteServerInstanceGroupLogicalNetworkACL

> DeleteServerInstanceGroupLogicalNetworkACL(ctx, serverInstanceGroupId, connectionId, ruleId).Execute()

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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	ruleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerInstanceGroupAPI.DeleteServerInstanceGroupLogicalNetworkACL(context.Background(), serverInstanceGroupId, connectionId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.DeleteServerInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


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


## DeleteServerInstanceGroupNetworkConfigurationConnection

> DeleteServerInstanceGroupNetworkConfigurationConnection(ctx, serverInstanceGroupId, connectionId).Execute()

Delete a network connection for a server instance group

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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerInstanceGroupAPI.DeleteServerInstanceGroupNetworkConfigurationConnection(context.Background(), serverInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.DeleteServerInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


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


## GetInfrastructureServerInstanceGroups

> ServerInstanceGroupPaginatedList GetInfrastructureServerInstanceGroups(ctx, infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterExtensionInstanceId(filterExtensionInstanceId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Server Instance Groups for an infrastructure



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
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetInfrastructureServerInstanceGroups(context.Background(), infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterExtensionInstanceId(filterExtensionInstanceId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetInfrastructureServerInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureServerInstanceGroups`: ServerInstanceGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetInfrastructureServerInstanceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureServerInstanceGroupsRequest struct via the builder pattern


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

[**ServerInstanceGroupPaginatedList**](ServerInstanceGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroup

> ServerInstanceGroup GetServerInstanceGroup(ctx, serverInstanceGroupId).Execute()

Get Server Instance Groups details



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
	serverInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroup(context.Background(), serverInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroup`: ServerInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerInstanceGroup**](ServerInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroupConfig

> ServerInstanceGroupConfiguration GetServerInstanceGroupConfig(ctx, serverInstanceGroupId).Execute()

Get Server Instance Group config details



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
	serverInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupConfig(context.Background(), serverInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupConfig`: ServerInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerInstanceGroupConfiguration**](ServerInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroupDriveGroups

> DriveGroupList GetServerInstanceGroupDriveGroups(ctx, serverInstanceGroupId).Execute()

Get Server Instance Group Drive Groups



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
	serverInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupDriveGroups(context.Background(), serverInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupDriveGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupDriveGroups`: DriveGroupList
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupDriveGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupDriveGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriveGroupList**](DriveGroupList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroupInterface

> ServerInstanceGroupInterface GetServerInstanceGroupInterface(ctx, serverInstanceGroupId, interfaceId).Execute()

Get Server Instance Group Interface details



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
	serverInstanceGroupId := int32(56) // int32 | 
	interfaceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupInterface(context.Background(), serverInstanceGroupId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupInterface`: ServerInstanceGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**interfaceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerInstanceGroupInterface**](ServerInstanceGroupInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroupInterfaces

> ServerInstanceGroupInterfacePaginatedList GetServerInstanceGroupInterfaces(ctx, serverInstanceGroupId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get Server Instance Group Interfaces



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
	serverInstanceGroupId := int32(56) // int32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>infrastructureId</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,infrastructureId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>infrastructureId</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupInterfaces(context.Background(), serverInstanceGroupId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupInterfaces`: ServerInstanceGroupInterfacePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,infrastructureId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**ServerInstanceGroupInterfacePaginatedList**](ServerInstanceGroupInterfacePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroupLogicalNetworkACL

> LogicalNetworkACL GetServerInstanceGroupLogicalNetworkACL(ctx, serverInstanceGroupId, connectionId).Execute()

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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupLogicalNetworkACL(context.Background(), serverInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupLogicalNetworkACL`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupLogicalNetworkACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


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


## GetServerInstanceGroupLogicalNetworkACLById

> LogicalNetworkACL GetServerInstanceGroupLogicalNetworkACLById(ctx, serverInstanceGroupId, connectionId, ruleId).Execute()

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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	ruleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupLogicalNetworkACLById(context.Background(), serverInstanceGroupId, connectionId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupLogicalNetworkACLById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupLogicalNetworkACLById`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupLogicalNetworkACLById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupLogicalNetworkACLByIdRequest struct via the builder pattern


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


## GetServerInstanceGroupNetworkConfiguration

> NetworkEndpointGroup GetServerInstanceGroupNetworkConfiguration(ctx, serverInstanceGroupId).Execute()

Get server instance group network configuration



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
	serverInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfiguration(context.Background(), serverInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupNetworkConfiguration`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupNetworkConfigurationRequest struct via the builder pattern


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


## GetServerInstanceGroupNetworkConfigurationConnectionById

> ServerInstanceGroupNetworkConnection GetServerInstanceGroupNetworkConfigurationConnectionById(ctx, serverInstanceGroupId, connectionId).Execute()

Get server instance group network configuration connection by id



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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationConnectionById(context.Background(), serverInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationConnectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupNetworkConfigurationConnectionById`: ServerInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationConnectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupNetworkConfigurationConnectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerInstanceGroupNetworkConnection**](ServerInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroupNetworkConfigurationConnections

> ServerInstanceGroupNetworkConnectionsList GetServerInstanceGroupNetworkConfigurationConnections(ctx, serverInstanceGroupId).Execute()

Get server instance group's network connections

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
	serverInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationConnections(context.Background(), serverInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupNetworkConfigurationConnections`: ServerInstanceGroupNetworkConnectionsList
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupNetworkConfigurationConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerInstanceGroupNetworkConnectionsList**](ServerInstanceGroupNetworkConnectionsList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceGroupServerInstances

> ServerInstancePaginatedList GetServerInstanceGroupServerInstances(ctx, serverInstanceGroupId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterServerId(filterServerId).FilterServiceStatus(filterServiceStatus).FilterConfigServerId(filterConfigServerId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Server Instances for a Server Instance Group



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
	serverInstanceGroupId := int32(56) // int32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterGroupId := []string{"Inner_example"} // []string | Filter by groupId query param.           <p>              <b>Format: </b> filter.groupId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.groupId=$not:$like:John Doe&filter.groupId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.           <p>              <b>Format: </b> filter.serverId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverId=$not:$like:John Doe&filter.serverId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigServerId := []string{"Inner_example"} // []string | Filter by config.serverId query param.           <p>              <b>Format: </b> filter.config.serverId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.serverId=$not:$like:John Doe&filter.config.serverId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupServerInstances(context.Background(), serverInstanceGroupId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterServerId(filterServerId).FilterServiceStatus(filterServiceStatus).FilterConfigServerId(filterConfigServerId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.GetServerInstanceGroupServerInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceGroupServerInstances`: ServerInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.GetServerInstanceGroupServerInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceGroupServerInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterGroupId** | **[]string** | Filter by groupId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.groupId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.groupId&#x3D;$not:$like:John Doe&amp;filter.groupId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerId** | **[]string** | Filter by serverId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverId&#x3D;$not:$like:John Doe&amp;filter.serverId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigServerId** | **[]string** | Filter by config.serverId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.serverId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.serverId&#x3D;$not:$like:John Doe&amp;filter.config.serverId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**ServerInstancePaginatedList**](ServerInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerInstanceGroupConfig

> ServerInstanceGroupConfiguration UpdateServerInstanceGroupConfig(ctx, serverInstanceGroupId).ServerInstanceGroupUpdate(serverInstanceGroupUpdate).IfMatch(ifMatch).Execute()

Updates Server Instance Group configuration



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
	serverInstanceGroupId := int32(56) // int32 | 
	serverInstanceGroupUpdate := *openapiclient.NewServerInstanceGroupUpdate() // ServerInstanceGroupUpdate | The Server Instance Group configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.UpdateServerInstanceGroupConfig(context.Background(), serverInstanceGroupId).ServerInstanceGroupUpdate(serverInstanceGroupUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.UpdateServerInstanceGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerInstanceGroupConfig`: ServerInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.UpdateServerInstanceGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverInstanceGroupUpdate** | [**ServerInstanceGroupUpdate**](ServerInstanceGroupUpdate.md) | The Server Instance Group configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ServerInstanceGroupConfiguration**](ServerInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerInstanceGroupLogicalNetworkACL

> LogicalNetworkACL UpdateServerInstanceGroupLogicalNetworkACL(ctx, serverInstanceGroupId, connectionId, ruleId).UpdateLogicalNetworkACL(updateLogicalNetworkACL).Execute()

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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	ruleId := int32(56) // int32 | 
	updateLogicalNetworkACL := *openapiclient.NewUpdateLogicalNetworkACL() // UpdateLogicalNetworkACL | The security rule to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.UpdateServerInstanceGroupLogicalNetworkACL(context.Background(), serverInstanceGroupId, connectionId, ruleId).UpdateLogicalNetworkACL(updateLogicalNetworkACL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.UpdateServerInstanceGroupLogicalNetworkACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerInstanceGroupLogicalNetworkACL`: LogicalNetworkACL
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.UpdateServerInstanceGroupLogicalNetworkACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceGroupLogicalNetworkACLRequest struct via the builder pattern


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


## UpdateServerInstanceGroupMeta

> UpdateServerInstanceGroupMeta(ctx, serverInstanceGroupId).GenericMeta(genericMeta).Execute()

Update an Server Instance Group meta information



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
	serverInstanceGroupId := int32(56) // int32 | 
	genericMeta := *openapiclient.NewGenericMeta() // GenericMeta | The Server Instance Group meta information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerInstanceGroupAPI.UpdateServerInstanceGroupMeta(context.Background(), serverInstanceGroupId).GenericMeta(genericMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.UpdateServerInstanceGroupMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceGroupMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genericMeta** | [**GenericMeta**](GenericMeta.md) | The Server Instance Group meta information | 

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


## UpdateServerInstanceGroupNetworkConfiguration

> NetworkEndpointGroup UpdateServerInstanceGroupNetworkConfiguration(ctx, serverInstanceGroupId).Execute()

Create or return the network configuration of the specified server instance group

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
	serverInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.UpdateServerInstanceGroupNetworkConfiguration(context.Background(), serverInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.UpdateServerInstanceGroupNetworkConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerInstanceGroupNetworkConfiguration`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.UpdateServerInstanceGroupNetworkConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceGroupNetworkConfigurationRequest struct via the builder pattern


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


## UpdateServerInstanceGroupNetworkConfigurationConnection

> ServerInstanceGroupNetworkConnection UpdateServerInstanceGroupNetworkConfigurationConnection(ctx, serverInstanceGroupId, connectionId).UpdateNetworkEndpointGroupLogicalNetwork(updateNetworkEndpointGroupLogicalNetwork).Execute()

Update a network connection for a server instance group

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
	serverInstanceGroupId := int32(56) // int32 | 
	connectionId := float32(8.14) // float32 | 
	updateNetworkEndpointGroupLogicalNetwork := *openapiclient.NewUpdateNetworkEndpointGroupLogicalNetwork() // UpdateNetworkEndpointGroupLogicalNetwork | The network connection object to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceGroupAPI.UpdateServerInstanceGroupNetworkConfigurationConnection(context.Background(), serverInstanceGroupId, connectionId).UpdateNetworkEndpointGroupLogicalNetwork(updateNetworkEndpointGroupLogicalNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceGroupAPI.UpdateServerInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerInstanceGroupNetworkConfigurationConnection`: ServerInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceGroupAPI.UpdateServerInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceGroupId** | **int32** |  | 
**connectionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkEndpointGroupLogicalNetwork** | [**UpdateNetworkEndpointGroupLogicalNetwork**](UpdateNetworkEndpointGroupLogicalNetwork.md) | The network connection object to update | 

### Return type

[**ServerInstanceGroupNetworkConnection**](ServerInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

