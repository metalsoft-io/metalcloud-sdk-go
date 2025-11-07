# \VMInstanceGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyVMTypeOnVMInstanceGroup**](VMInstanceGroupAPI.md#ApplyVMTypeOnVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance Group
[**CreateVMInstanceGroup**](VMInstanceGroupAPI.md#CreateVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Creates a VM Instance Group
[**CreateVMInstanceGroupNetworkConfigurationConnection**](VMInstanceGroupAPI.md#CreateVMInstanceGroupNetworkConfigurationConnection) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config/networking/connections | Create a network connection for a VM instance group
[**DeleteVMInstanceGroup**](VMInstanceGroupAPI.md#DeleteVMInstanceGroup) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Deletes a VM Instance Group
[**DeleteVMInstanceGroupNetworkConfigurationConnection**](VMInstanceGroupAPI.md#DeleteVMInstanceGroupNetworkConfigurationConnection) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config/networking/connections/{connectionId} | Delete a network connection for a VM instance group
[**GetInfrastructureVMInstanceGroup**](VMInstanceGroupAPI.md#GetInfrastructureVMInstanceGroup) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Get VM Instance Group information
[**GetInfrastructureVMInstanceGroups**](VMInstanceGroupAPI.md#GetInfrastructureVMInstanceGroups) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Get all VM Instance Groups on the infrastructure
[**GetVMInstanceGroupConfigInfo**](VMInstanceGroupAPI.md#GetVMInstanceGroupConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config | Get configuration information about the specified VM Instance Group
[**GetVMInstanceGroupInterfaceInfo**](VMInstanceGroupAPI.md#GetVMInstanceGroupInterfaceInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/interfaces/{vmInstanceGroupInterfaceId} | Get VM Instance Group Interface information
[**GetVMInstanceGroupInterfaces**](VMInstanceGroupAPI.md#GetVMInstanceGroupInterfaces) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/interfaces | Get all VM Instance Group Interfaces on the VM Instance Group
[**GetVMInstanceGroupNetworkConfigurationConnectionById**](VMInstanceGroupAPI.md#GetVMInstanceGroupNetworkConfigurationConnectionById) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config/networking/connections/{connectionId} | Get VM instance group network configuration connection by id
[**GetVMInstanceGroupNetworkConfigurationConnections**](VMInstanceGroupAPI.md#GetVMInstanceGroupNetworkConfigurationConnections) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config/networking/connections | Get VM instance group&#39;s network connections
[**GetVMInstanceGroupVMInstances**](VMInstanceGroupAPI.md#GetVMInstanceGroupVMInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/vm-instances | Get the VM Instances of VM Instance Group
[**GetVmInstanceGroupNetworkConfiguration**](VMInstanceGroupAPI.md#GetVmInstanceGroupNetworkConfiguration) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config/networking | Get vm instance group network configuration
[**PatchVMInstanceGroupMeta**](VMInstanceGroupAPI.md#PatchVMInstanceGroupMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/meta | Updates the meta of a VM Instance Group
[**UpdateVMInstanceGroupConfig**](VMInstanceGroupAPI.md#UpdateVMInstanceGroupConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config | Updates VM Instance Group information
[**UpdateVMInstanceGroupNetworkConfigurationConnection**](VMInstanceGroupAPI.md#UpdateVMInstanceGroupNetworkConfigurationConnection) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/config/networking/connections/{connectionId} | Update a network connection for a VM instance group



## ApplyVMTypeOnVMInstanceGroup

> VMInstanceGroup ApplyVMTypeOnVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId, vmTypeId).IfMatch(ifMatch).Execute()

Applies a VM Type to a VM Instance Group



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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 
	vmTypeId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.ApplyVMTypeOnVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId, vmTypeId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.ApplyVMTypeOnVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyVMTypeOnVMInstanceGroup`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.ApplyVMTypeOnVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 
**vmTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyVMTypeOnVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **string** | Entity tag | 

### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVMInstanceGroup

> VMInstanceGroup CreateVMInstanceGroup(ctx, infrastructureId).CreateVMInstanceGroup(createVMInstanceGroup).Execute()

Creates a VM Instance Group



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
	infrastructureId := float32(8.14) // float32 | 
	createVMInstanceGroup := *openapiclient.NewCreateVMInstanceGroup(float32(123), float32(123), float32(123), float32(123)) // CreateVMInstanceGroup | The VM Instance Group create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.CreateVMInstanceGroup(context.Background(), infrastructureId).CreateVMInstanceGroup(createVMInstanceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.CreateVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVMInstanceGroup`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.CreateVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVMInstanceGroup** | [**CreateVMInstanceGroup**](CreateVMInstanceGroup.md) | The VM Instance Group create object | 

### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVMInstanceGroupNetworkConfigurationConnection

> VMInstanceGroupNetworkConnection CreateVMInstanceGroupNetworkConfigurationConnection(ctx, infrastructureId, vmInstanceGroupId).CreateVMInstanceGroupNetworkConnection(createVMInstanceGroupNetworkConnection).Execute()

Create a network connection for a VM instance group

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
	vmInstanceGroupId := int32(56) // int32 | 
	createVMInstanceGroupNetworkConnection := *openapiclient.NewCreateVMInstanceGroupNetworkConnection("1", true, openapiclient.NetworkEndpointGroupAllowedAccessMode("l2")) // CreateVMInstanceGroupNetworkConnection | The network connection object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.CreateVMInstanceGroupNetworkConfigurationConnection(context.Background(), infrastructureId, vmInstanceGroupId).CreateVMInstanceGroupNetworkConnection(createVMInstanceGroupNetworkConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.CreateVMInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVMInstanceGroupNetworkConfigurationConnection`: VMInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.CreateVMInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 
**vmInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVMInstanceGroupNetworkConnection** | [**CreateVMInstanceGroupNetworkConnection**](CreateVMInstanceGroupNetworkConnection.md) | The network connection object to create | 

### Return type

[**VMInstanceGroupNetworkConnection**](VMInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVMInstanceGroup

> DeleteVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId).IfMatch(ifMatch).Execute()

Deletes a VM Instance Group



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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceGroupAPI.DeleteVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.DeleteVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMInstanceGroupRequest struct via the builder pattern


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


## DeleteVMInstanceGroupNetworkConfigurationConnection

> DeleteVMInstanceGroupNetworkConfigurationConnection(ctx, infrastructureId, vmInstanceGroupId, connectionId).Execute()

Delete a network connection for a VM instance group

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
	vmInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceGroupAPI.DeleteVMInstanceGroupNetworkConfigurationConnection(context.Background(), infrastructureId, vmInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.DeleteVMInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 
**vmInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


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


## GetInfrastructureVMInstanceGroup

> VMInstanceGroup GetInfrastructureVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId).Execute()

Get VM Instance Group information



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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetInfrastructureVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetInfrastructureVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureVMInstanceGroup`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetInfrastructureVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureVMInstanceGroups

> VMInstanceGroupPaginatedList GetInfrastructureVMInstanceGroups(ctx, infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all VM Instance Groups on the infrastructure



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
	infrastructureId := float32(8.14) // float32 | 
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
	resp, r, err := apiClient.VMInstanceGroupAPI.GetInfrastructureVMInstanceGroups(context.Background(), infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetInfrastructureVMInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureVMInstanceGroups`: VMInstanceGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetInfrastructureVMInstanceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureVMInstanceGroupsRequest struct via the builder pattern


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

[**VMInstanceGroupPaginatedList**](VMInstanceGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroupConfigInfo

> VMInstanceGroupConfiguration GetVMInstanceGroupConfigInfo(ctx, infrastructureId, vmInstanceGroupId).Execute()

Get configuration information about the specified VM Instance Group

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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupConfigInfo(context.Background(), infrastructureId, vmInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroupConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroupConfigInfo`: VMInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroupConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMInstanceGroupConfiguration**](VMInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroupInterfaceInfo

> VMInstanceGroupInterface GetVMInstanceGroupInterfaceInfo(ctx, infrastructureId, vmInstanceGroupId, vmInstanceGroupInterfaceId).Execute()

Get VM Instance Group Interface information



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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 
	vmInstanceGroupInterfaceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupInterfaceInfo(context.Background(), infrastructureId, vmInstanceGroupId, vmInstanceGroupInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroupInterfaceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroupInterfaceInfo`: VMInstanceGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroupInterfaceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 
**vmInstanceGroupInterfaceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupInterfaceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VMInstanceGroupInterface**](VMInstanceGroupInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroupInterfaces

> VMInstanceGroupPaginatedList GetVMInstanceGroupInterfaces(ctx, infrastructureId, vmInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterNetworkId(filterNetworkId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all VM Instance Group Interfaces on the VM Instance Group



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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 
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
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupInterfaces(context.Background(), infrastructureId, vmInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterNetworkId(filterNetworkId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroupInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroupInterfaces`: VMInstanceGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroupInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupInterfacesRequest struct via the builder pattern


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

[**VMInstanceGroupPaginatedList**](VMInstanceGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroupNetworkConfigurationConnectionById

> VMInstanceGroupNetworkConnection GetVMInstanceGroupNetworkConfigurationConnectionById(ctx, infrastructureId, vmInstanceGroupId, connectionId).Execute()

Get VM instance group network configuration connection by id



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
	vmInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupNetworkConfigurationConnectionById(context.Background(), infrastructureId, vmInstanceGroupId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroupNetworkConfigurationConnectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroupNetworkConfigurationConnectionById`: VMInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroupNetworkConfigurationConnectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 
**vmInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupNetworkConfigurationConnectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VMInstanceGroupNetworkConnection**](VMInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroupNetworkConfigurationConnections

> VMInstanceGroupNetworkConnectionsList GetVMInstanceGroupNetworkConfigurationConnections(ctx, infrastructureId, vmInstanceGroupId).Execute()

Get VM instance group's network connections



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
	vmInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupNetworkConfigurationConnections(context.Background(), infrastructureId, vmInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroupNetworkConfigurationConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroupNetworkConfigurationConnections`: VMInstanceGroupNetworkConnectionsList
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroupNetworkConfigurationConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 
**vmInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupNetworkConfigurationConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMInstanceGroupNetworkConnectionsList**](VMInstanceGroupNetworkConnectionsList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroupVMInstances

> VMInstancePaginatedList GetVMInstanceGroupVMInstances(ctx, infrastructureId, vmInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get the VM Instances of VM Instance Group



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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 
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
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupVMInstances(context.Background(), infrastructureId, vmInstanceGroupId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroupVMInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroupVMInstances`: VMInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroupVMInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupVMInstancesRequest struct via the builder pattern


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

[**VMInstancePaginatedList**](VMInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmInstanceGroupNetworkConfiguration

> NetworkEndpointGroup GetVmInstanceGroupNetworkConfiguration(ctx, infrastructureId, vmInstanceGroupId).Execute()

Get vm instance group network configuration



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
	vmInstanceGroupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVmInstanceGroupNetworkConfiguration(context.Background(), infrastructureId, vmInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVmInstanceGroupNetworkConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVmInstanceGroupNetworkConfiguration`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVmInstanceGroupNetworkConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 
**vmInstanceGroupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmInstanceGroupNetworkConfigurationRequest struct via the builder pattern


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


## PatchVMInstanceGroupMeta

> VMInstanceGroup PatchVMInstanceGroupMeta(ctx, infrastructureId, vmInstanceGroupId).UpdateVMInstanceGroupMeta(updateVMInstanceGroupMeta).Execute()

Updates the meta of a VM Instance Group

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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 
	updateVMInstanceGroupMeta := *openapiclient.NewUpdateVMInstanceGroupMeta() // UpdateVMInstanceGroupMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.PatchVMInstanceGroupMeta(context.Background(), infrastructureId, vmInstanceGroupId).UpdateVMInstanceGroupMeta(updateVMInstanceGroupMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.PatchVMInstanceGroupMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchVMInstanceGroupMeta`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.PatchVMInstanceGroupMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchVMInstanceGroupMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMInstanceGroupMeta** | [**UpdateVMInstanceGroupMeta**](UpdateVMInstanceGroupMeta.md) |  | 

### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMInstanceGroupConfig

> VMInstanceGroupConfiguration UpdateVMInstanceGroupConfig(ctx, infrastructureId, vmInstanceGroupId).UpdateVMInstanceGroup(updateVMInstanceGroup).IfMatch(ifMatch).Execute()

Updates VM Instance Group information



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
	infrastructureId := float32(8.14) // float32 | 
	vmInstanceGroupId := float32(8.14) // float32 | 
	updateVMInstanceGroup := *openapiclient.NewUpdateVMInstanceGroup() // UpdateVMInstanceGroup | The VM Instance Group update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.UpdateVMInstanceGroupConfig(context.Background(), infrastructureId, vmInstanceGroupId).UpdateVMInstanceGroup(updateVMInstanceGroup).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.UpdateVMInstanceGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMInstanceGroupConfig`: VMInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.UpdateVMInstanceGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMInstanceGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMInstanceGroup** | [**UpdateVMInstanceGroup**](UpdateVMInstanceGroup.md) | The VM Instance Group update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**VMInstanceGroupConfiguration**](VMInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMInstanceGroupNetworkConfigurationConnection

> VMInstanceGroupNetworkConnection UpdateVMInstanceGroupNetworkConfigurationConnection(ctx, infrastructureId, vmInstanceGroupId, connectionId).UpdateVMInstanceGroupNetworkConnection(updateVMInstanceGroupNetworkConnection).Execute()

Update a network connection for a VM instance group

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
	vmInstanceGroupId := int32(56) // int32 | 
	connectionId := int32(56) // int32 | 
	updateVMInstanceGroupNetworkConnection := *openapiclient.NewUpdateVMInstanceGroupNetworkConnection() // UpdateVMInstanceGroupNetworkConnection | The network connection object to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.UpdateVMInstanceGroupNetworkConfigurationConnection(context.Background(), infrastructureId, vmInstanceGroupId, connectionId).UpdateVMInstanceGroupNetworkConnection(updateVMInstanceGroupNetworkConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.UpdateVMInstanceGroupNetworkConfigurationConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMInstanceGroupNetworkConfigurationConnection`: VMInstanceGroupNetworkConnection
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.UpdateVMInstanceGroupNetworkConfigurationConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 
**vmInstanceGroupId** | **int32** |  | 
**connectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMInstanceGroupNetworkConfigurationConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateVMInstanceGroupNetworkConnection** | [**UpdateVMInstanceGroupNetworkConnection**](UpdateVMInstanceGroupNetworkConnection.md) | The network connection object to update | 

### Return type

[**VMInstanceGroupNetworkConnection**](VMInstanceGroupNetworkConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

