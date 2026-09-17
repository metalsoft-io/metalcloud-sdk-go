# \ContainerInstanceGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyContainerTypeOnContainerInstanceGroup**](ContainerInstanceGroupAPI.md#ApplyContainerTypeOnContainerInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/actions/apply-type/{containerTypeId} | Applies a Container Type to a Container Instance Group
[**CreateContainerInstanceGroup**](ContainerInstanceGroupAPI.md#CreateContainerInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instance-groups | Creates a Container Instance Group
[**CreateContainerInstanceGroupNetworkConfigurationConnection**](ContainerInstanceGroupAPI.md#CreateContainerInstanceGroupNetworkConfigurationConnection) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config/networking/connections | Create a network connection for a Container instance group
[**DeleteContainerInstanceGroup**](ContainerInstanceGroupAPI.md#DeleteContainerInstanceGroup) | **Delete** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId} | Deletes a Container Instance Group
[**DeleteContainerInstanceGroupNetworkConfigurationConnection**](ContainerInstanceGroupAPI.md#DeleteContainerInstanceGroupNetworkConfigurationConnection) | **Delete** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config/networking/connections/{connectionId} | Delete a network connection for a Container instance group
[**GetContainerInstanceGroupConfigInfo**](ContainerInstanceGroupAPI.md#GetContainerInstanceGroupConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config | Get configuration information about the specified Container Instance Group
[**GetContainerInstanceGroupContainerInstances**](ContainerInstanceGroupAPI.md#GetContainerInstanceGroupContainerInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/container-instances | Get the Container Instances of Container Instance Group
[**GetContainerInstanceGroupInterfaceInfo**](ContainerInstanceGroupAPI.md#GetContainerInstanceGroupInterfaceInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/interfaces/{containerInstanceGroupInterfaceId} | Get Container Instance Group Interface information
[**GetContainerInstanceGroupInterfaces**](ContainerInstanceGroupAPI.md#GetContainerInstanceGroupInterfaces) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/interfaces | Get all Container Instance Group Interfaces on the Container Instance Group
[**GetContainerInstanceGroupNetworkConfiguration**](ContainerInstanceGroupAPI.md#GetContainerInstanceGroupNetworkConfiguration) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config/networking | Get container instance group network configuration
[**GetContainerInstanceGroupNetworkConfigurationConnectionById**](ContainerInstanceGroupAPI.md#GetContainerInstanceGroupNetworkConfigurationConnectionById) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config/networking/connections/{connectionId} | Get Container instance group network configuration connection by id
[**GetContainerInstanceGroupNetworkConfigurationConnections**](ContainerInstanceGroupAPI.md#GetContainerInstanceGroupNetworkConfigurationConnections) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config/networking/connections | Get Container instance group&#39;s network connections
[**GetInfrastructureContainerInstanceGroup**](ContainerInstanceGroupAPI.md#GetInfrastructureContainerInstanceGroup) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId} | Get Container Instance Group information
[**GetInfrastructureContainerInstanceGroups**](ContainerInstanceGroupAPI.md#GetInfrastructureContainerInstanceGroups) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instance-groups | Get all Container Instance Groups on the infrastructure
[**PatchContainerInstanceGroupMeta**](ContainerInstanceGroupAPI.md#PatchContainerInstanceGroupMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/meta | Updates the meta of a Container Instance Group
[**UpdateContainerInstanceGroupConfig**](ContainerInstanceGroupAPI.md#UpdateContainerInstanceGroupConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config | Updates Container Instance Group information
[**UpdateContainerInstanceGroupNetworkConfigurationConnection**](ContainerInstanceGroupAPI.md#UpdateContainerInstanceGroupNetworkConfigurationConnection) | **Patch** /api/v2/infrastructures/{infrastructureId}/container-instance-groups/{containerInstanceGroupId}/config/networking/connections/{connectionId} | Update a network connection for a Container instance group



## ApplyContainerTypeOnContainerInstanceGroup

> ContainerInstanceGroup ApplyContainerTypeOnContainerInstanceGroup(ctx, infrastructureId, containerInstanceGroupId, containerTypeId).IfMatch(ifMatch).Execute()

Applies a Container Type to a Container Instance Group



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	containerTypeId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.ApplyContainerTypeOnContainerInstanceGroup(context.Background(), infrastructureId, containerInstanceGroupId, containerTypeId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.ApplyContainerTypeOnContainerInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyContainerTypeOnContainerInstanceGroup`: ContainerInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.ApplyContainerTypeOnContainerInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 
**containerTypeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyContainerTypeOnContainerInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **string** | Entity tag | 

### Return type

[**ContainerInstanceGroup**](ContainerInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerInstanceGroup

> ContainerInstanceGroup CreateContainerInstanceGroup(ctx, infrastructureId).CreateContainerInstanceGroup(createContainerInstanceGroup).Execute()

Creates a Container Instance Group



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
	infrastructureId := int64(789) // int64 | 
	createContainerInstanceGroup := *openapiclient.NewCreateContainerInstanceGroup(float32(123), int64(123), int64(123), int64(123)) // CreateContainerInstanceGroup | The Container Instance Group create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.CreateContainerInstanceGroup(context.Background(), infrastructureId).CreateContainerInstanceGroup(createContainerInstanceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.CreateContainerInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerInstanceGroup`: ContainerInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.CreateContainerInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContainerInstanceGroup** | [**CreateContainerInstanceGroup**](CreateContainerInstanceGroup.md) | The Container Instance Group create object | 

### Return type

[**ContainerInstanceGroup**](ContainerInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerInstanceGroupNetworkConfigurationConnection

> ContainerInstanceGroupNetworkConnection CreateContainerInstanceGroupNetworkConfigurationConnection(ctx, infrastructureId, containerInstanceGroupId).CreateContainerInstanceGroupNetworkConnection(createContainerInstanceGroupNetworkConnection).Execute()

Create a network connection for a Container instance group

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	createContainerInstanceGroupNetworkConnection := *openapiclient.NewCreateContainerInstanceGroupNetworkConnection("1", true, openapiclient.NetworkEndpointGroupAllowedAccessMode("l2")) // CreateContainerInstanceGroupNetworkConnection | The network connection object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.CreateContainerInstanceGroupNetworkConfigurationConnection(context.Background(), infrastructureId, containerInstanceGroupId).CreateContainerInstanceGroupNetworkConnection(createContainerInstanceGroupNetworkConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.CreateContainerInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerInstanceGroupNetworkConfigurationConnection`: ContainerInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.CreateContainerInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createContainerInstanceGroupNetworkConnection** | [**CreateContainerInstanceGroupNetworkConnection**](CreateContainerInstanceGroupNetworkConnection.md) | The network connection object to create | 

### Return type

[**ContainerInstanceGroupNetworkConnection**](ContainerInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerInstanceGroup

> DeleteContainerInstanceGroup(ctx, infrastructureId, containerInstanceGroupId).IfMatch(ifMatch).Execute()

Deletes a Container Instance Group



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerInstanceGroupAPI.DeleteContainerInstanceGroup(context.Background(), infrastructureId, containerInstanceGroupId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.DeleteContainerInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerInstanceGroupRequest struct via the builder pattern


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


## DeleteContainerInstanceGroupNetworkConfigurationConnection

> DeleteContainerInstanceGroupNetworkConfigurationConnection(ctx, infrastructureId, containerInstanceGroupId, connectionId).Execute()

Delete a network connection for a Container instance group

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	connectionId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerInstanceGroupAPI.DeleteContainerInstanceGroupNetworkConfigurationConnection(context.Background(), infrastructureId, containerInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.DeleteContainerInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 
**connectionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


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


## GetContainerInstanceGroupConfigInfo

> ContainerInstanceGroupConfiguration GetContainerInstanceGroupConfigInfo(ctx, infrastructureId, containerInstanceGroupId).Execute()

Get configuration information about the specified Container Instance Group

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetContainerInstanceGroupConfigInfo(context.Background(), infrastructureId, containerInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetContainerInstanceGroupConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceGroupConfigInfo`: ContainerInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetContainerInstanceGroupConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceGroupConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerInstanceGroupConfiguration**](ContainerInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstanceGroupContainerInstances

> ContainerInstancePaginatedList GetContainerInstanceGroupContainerInstances(ctx, infrastructureId, containerInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterVmId(filterVmId).FilterConfigVmId(filterConfigVmId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get the Container Instances of Container Instance Group



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterVmId := []string{"Inner_example"} // []string | Filter by vmId query param.  **Format:** filter.vmId={$not}:OPERATION:VALUE    **Example:** filter.vmId=$eq:John Doe&filter.vmId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigVmId := []string{"Inner_example"} // []string | Filter by config.vmId query param.  **Format:** filter.config.vmId={$not}:OPERATION:VALUE    **Example:** filter.config.vmId=$eq:John Doe&filter.config.vmId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=serviceStatus:DESC   **Default Value:** id:DESC  **Available Fields** - id  - serviceStatus  - config.deployStatus  - config.deployType  - containerId  - config.containerId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,serviceStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - serviceStatus  - config.deployStatus  - config.deployType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetContainerInstanceGroupContainerInstances(context.Background(), infrastructureId, containerInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterVmId(filterVmId).FilterConfigVmId(filterConfigVmId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetContainerInstanceGroupContainerInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceGroupContainerInstances`: ContainerInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetContainerInstanceGroupContainerInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceGroupContainerInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.  **Format:** filter.subdomain&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomain&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterVmId** | **[]string** | Filter by vmId query param.  **Format:** filter.vmId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vmId&#x3D;$eq:John Doe&amp;filter.vmId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigVmId** | **[]string** | Filter by config.vmId query param.  **Format:** filter.config.vmId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.vmId&#x3D;$eq:John Doe&amp;filter.config.vmId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;serviceStatus:DESC   **Default Value:** id:DESC  **Available Fields** - id  - serviceStatus  - config.deployStatus  - config.deployType  - containerId  - config.containerId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,serviceStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - serviceStatus  - config.deployStatus  - config.deployType  | 

### Return type

[**ContainerInstancePaginatedList**](ContainerInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstanceGroupInterfaceInfo

> ContainerInstanceGroupInterface GetContainerInstanceGroupInterfaceInfo(ctx, infrastructureId, containerInstanceGroupId, containerInstanceGroupInterfaceId).Execute()

Get Container Instance Group Interface information



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	containerInstanceGroupInterfaceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetContainerInstanceGroupInterfaceInfo(context.Background(), infrastructureId, containerInstanceGroupId, containerInstanceGroupInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetContainerInstanceGroupInterfaceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceGroupInterfaceInfo`: ContainerInstanceGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetContainerInstanceGroupInterfaceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 
**containerInstanceGroupInterfaceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceGroupInterfaceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ContainerInstanceGroupInterface**](ContainerInstanceGroupInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstanceGroupInterfaces

> ContainerInstanceGroupInterfacePaginatedList GetContainerInstanceGroupInterfaces(ctx, infrastructureId, containerInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterNetworkId(filterNetworkId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Container Instance Group Interfaces on the Container Instance Group



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterNetworkId := []string{"Inner_example"} // []string | Filter by networkId query param.  **Format:** filter.networkId={$not}:OPERATION:VALUE    **Example:** filter.networkId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=networkId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkId  - serviceStatus  - config.deployStatus  - config.deployType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,networkId,serviceStatus,config.deployStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - networkId  - serviceStatus  - config.deployStatus  - config.deployType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetContainerInstanceGroupInterfaces(context.Background(), infrastructureId, containerInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterNetworkId(filterNetworkId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetContainerInstanceGroupInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceGroupInterfaces`: ContainerInstanceGroupInterfacePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetContainerInstanceGroupInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceGroupInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterNetworkId** | **[]string** | Filter by networkId query param.  **Format:** filter.networkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;networkId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkId  - serviceStatus  - config.deployStatus  - config.deployType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,networkId,serviceStatus,config.deployStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - networkId  - serviceStatus  - config.deployStatus  - config.deployType  | 

### Return type

[**ContainerInstanceGroupInterfacePaginatedList**](ContainerInstanceGroupInterfacePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstanceGroupNetworkConfiguration

> NetworkEndpointGroup GetContainerInstanceGroupNetworkConfiguration(ctx, infrastructureId, containerInstanceGroupId).Execute()

Get container instance group network configuration



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfiguration(context.Background(), infrastructureId, containerInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceGroupNetworkConfiguration`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceGroupNetworkConfigurationRequest struct via the builder pattern


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


## GetContainerInstanceGroupNetworkConfigurationConnectionById

> ContainerInstanceGroupNetworkConnection GetContainerInstanceGroupNetworkConfigurationConnectionById(ctx, infrastructureId, containerInstanceGroupId, connectionId).Execute()

Get Container instance group network configuration connection by id



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	connectionId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfigurationConnectionById(context.Background(), infrastructureId, containerInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfigurationConnectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceGroupNetworkConfigurationConnectionById`: ContainerInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfigurationConnectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 
**connectionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceGroupNetworkConfigurationConnectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ContainerInstanceGroupNetworkConnection**](ContainerInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstanceGroupNetworkConfigurationConnections

> ContainerInstanceGroupNetworkConnectionsList GetContainerInstanceGroupNetworkConfigurationConnections(ctx, infrastructureId, containerInstanceGroupId).Execute()

Get Container instance group's network connections



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfigurationConnections(context.Background(), infrastructureId, containerInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfigurationConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceGroupNetworkConfigurationConnections`: ContainerInstanceGroupNetworkConnectionsList
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetContainerInstanceGroupNetworkConfigurationConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceGroupNetworkConfigurationConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerInstanceGroupNetworkConnectionsList**](ContainerInstanceGroupNetworkConnectionsList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureContainerInstanceGroup

> ContainerInstanceGroup GetInfrastructureContainerInstanceGroup(ctx, infrastructureId, containerInstanceGroupId).Execute()

Get Container Instance Group information



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetInfrastructureContainerInstanceGroup(context.Background(), infrastructureId, containerInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetInfrastructureContainerInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureContainerInstanceGroup`: ContainerInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetInfrastructureContainerInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureContainerInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerInstanceGroup**](ContainerInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureContainerInstanceGroups

> ContainerInstanceGroupPaginatedList GetInfrastructureContainerInstanceGroups(ctx, infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Container Instance Groups on the infrastructure



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
	infrastructureId := int64(789) // int64 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=serviceStatus:DESC   **Default Value:** id:DESC  **Available Fields** - id  - serviceStatus  - config.deployStatus  - config.deployType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,serviceStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - serviceStatus  - config.deployStatus  - config.deployType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.GetInfrastructureContainerInstanceGroups(context.Background(), infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.GetInfrastructureContainerInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureContainerInstanceGroups`: ContainerInstanceGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.GetInfrastructureContainerInstanceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureContainerInstanceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.  **Format:** filter.subdomain&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomain&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;serviceStatus:DESC   **Default Value:** id:DESC  **Available Fields** - id  - serviceStatus  - config.deployStatus  - config.deployType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,serviceStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - serviceStatus  - config.deployStatus  - config.deployType  | 

### Return type

[**ContainerInstanceGroupPaginatedList**](ContainerInstanceGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchContainerInstanceGroupMeta

> ContainerInstanceGroup PatchContainerInstanceGroupMeta(ctx, infrastructureId, containerInstanceGroupId).UpdateContainerInstanceGroupMeta(updateContainerInstanceGroupMeta).Execute()

Updates the meta of a Container Instance Group

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	updateContainerInstanceGroupMeta := *openapiclient.NewUpdateContainerInstanceGroupMeta() // UpdateContainerInstanceGroupMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.PatchContainerInstanceGroupMeta(context.Background(), infrastructureId, containerInstanceGroupId).UpdateContainerInstanceGroupMeta(updateContainerInstanceGroupMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.PatchContainerInstanceGroupMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContainerInstanceGroupMeta`: ContainerInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.PatchContainerInstanceGroupMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContainerInstanceGroupMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateContainerInstanceGroupMeta** | [**UpdateContainerInstanceGroupMeta**](UpdateContainerInstanceGroupMeta.md) |  | 

### Return type

[**ContainerInstanceGroup**](ContainerInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerInstanceGroupConfig

> ContainerInstanceGroupConfiguration UpdateContainerInstanceGroupConfig(ctx, infrastructureId, containerInstanceGroupId).UpdateContainerInstanceGroup(updateContainerInstanceGroup).IfMatch(ifMatch).Execute()

Updates Container Instance Group information



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	updateContainerInstanceGroup := *openapiclient.NewUpdateContainerInstanceGroup() // UpdateContainerInstanceGroup | The Container Instance Group update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.UpdateContainerInstanceGroupConfig(context.Background(), infrastructureId, containerInstanceGroupId).UpdateContainerInstanceGroup(updateContainerInstanceGroup).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.UpdateContainerInstanceGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContainerInstanceGroupConfig`: ContainerInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.UpdateContainerInstanceGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerInstanceGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateContainerInstanceGroup** | [**UpdateContainerInstanceGroup**](UpdateContainerInstanceGroup.md) | The Container Instance Group update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ContainerInstanceGroupConfiguration**](ContainerInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerInstanceGroupNetworkConfigurationConnection

> ContainerInstanceGroupNetworkConnection UpdateContainerInstanceGroupNetworkConfigurationConnection(ctx, infrastructureId, containerInstanceGroupId, connectionId).UpdateContainerInstanceGroupNetworkConnection(updateContainerInstanceGroupNetworkConnection).Execute()

Update a network connection for a Container instance group

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceGroupId := int64(789) // int64 | 
	connectionId := int64(789) // int64 | 
	updateContainerInstanceGroupNetworkConnection := *openapiclient.NewUpdateContainerInstanceGroupNetworkConnection() // UpdateContainerInstanceGroupNetworkConnection | The network connection object to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceGroupAPI.UpdateContainerInstanceGroupNetworkConfigurationConnection(context.Background(), infrastructureId, containerInstanceGroupId, connectionId).UpdateContainerInstanceGroupNetworkConnection(updateContainerInstanceGroupNetworkConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceGroupAPI.UpdateContainerInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContainerInstanceGroupNetworkConfigurationConnection`: ContainerInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceGroupAPI.UpdateContainerInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceGroupId** | **int64** |  | 
**connectionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateContainerInstanceGroupNetworkConnection** | [**UpdateContainerInstanceGroupNetworkConnection**](UpdateContainerInstanceGroupNetworkConnection.md) | The network connection object to update | 

### Return type

[**ContainerInstanceGroupNetworkConnection**](ContainerInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

