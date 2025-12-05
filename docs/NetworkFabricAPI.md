# \NetworkFabricAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptNetworkFabricDeploy**](NetworkFabricAPI.md#AcceptNetworkFabricDeploy) | **Post** /api/v2/network-fabrics/{networkFabricId}/actions/accept-deploy | Accepts the deployment of the specified network fabric
[**ActivateNetworkFabric**](NetworkFabricAPI.md#ActivateNetworkFabric) | **Post** /api/v2/network-fabrics/{networkFabricId}/actions/activate | Activate a network fabric
[**AddNetworkDevicesToFabric**](NetworkFabricAPI.md#AddNetworkDevicesToFabric) | **Post** /api/v2/network-fabrics/{networkFabricId}/network-devices | Add a list of network devices to a fabric
[**CreateNetworkFabric**](NetworkFabricAPI.md#CreateNetworkFabric) | **Post** /api/v2/network-fabrics | Create a new network fabric
[**CreateNetworkFabricLink**](NetworkFabricAPI.md#CreateNetworkFabricLink) | **Post** /api/v2/network-fabrics/{networkFabricId}/links | Create a new network fabric link
[**DeleteNetworkFabric**](NetworkFabricAPI.md#DeleteNetworkFabric) | **Delete** /api/v2/network-fabrics/{networkFabricId} | Delete a network fabric
[**DeleteNetworkFabricLink**](NetworkFabricAPI.md#DeleteNetworkFabricLink) | **Delete** /api/v2/network-fabrics/{networkFabricId}/links/{networkFabricLinkId} | Remove a network fabric link
[**DeployNetworkFabric**](NetworkFabricAPI.md#DeployNetworkFabric) | **Post** /api/v2/network-fabrics/{networkFabricId}/actions/deploy | Deploys the specified network fabric
[**GetFabricNetworkDevices**](NetworkFabricAPI.md#GetFabricNetworkDevices) | **Get** /api/v2/network-fabrics/{networkFabricId}/network-devices | Get paginated Network Devices
[**GetNetworkFabricById**](NetworkFabricAPI.md#GetNetworkFabricById) | **Get** /api/v2/network-fabrics/{networkFabricId} | Get a network fabric by ID
[**GetNetworkFabricLinks**](NetworkFabricAPI.md#GetNetworkFabricLinks) | **Get** /api/v2/network-fabrics/{networkFabricId}/links | Get paginated Network Fabric Links
[**GetNetworkFabrics**](NetworkFabricAPI.md#GetNetworkFabrics) | **Get** /api/v2/network-fabrics | List all network fabrics
[**RejectNetworkFabricDeploy**](NetworkFabricAPI.md#RejectNetworkFabricDeploy) | **Post** /api/v2/network-fabrics/{networkFabricId}/actions/reject-deploy | Rejects the deployment of the specified network fabric
[**RemoveNetworkDeviceFromFabric**](NetworkFabricAPI.md#RemoveNetworkDeviceFromFabric) | **Delete** /api/v2/network-fabrics/{networkFabricId}/network-devices/{networkDeviceId} | Remove a network device from a fabric
[**UpdateNetworkFabric**](NetworkFabricAPI.md#UpdateNetworkFabric) | **Patch** /api/v2/network-fabrics/{networkFabricId} | Update a network fabric



## AcceptNetworkFabricDeploy

> AcceptNetworkFabricDeploy(ctx, networkFabricId).Execute()

Accepts the deployment of the specified network fabric



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
	networkFabricId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkFabricAPI.AcceptNetworkFabricDeploy(context.Background(), networkFabricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.AcceptNetworkFabricDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptNetworkFabricDeployRequest struct via the builder pattern


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


## ActivateNetworkFabric

> NetworkFabric ActivateNetworkFabric(ctx, networkFabricId).IfMatch(ifMatch).Execute()

Activate a network fabric

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
	networkFabricId := int32(56) // int32 | The ID of the network fabric to activate
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.ActivateNetworkFabric(context.Background(), networkFabricId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.ActivateNetworkFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateNetworkFabric`: NetworkFabric
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.ActivateNetworkFabric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the network fabric to activate | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateNetworkFabricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkFabric**](NetworkFabric.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddNetworkDevicesToFabric

> NetworkFabric AddNetworkDevicesToFabric(ctx, networkFabricId).NetworkDevicesToFabric(networkDevicesToFabric).Execute()

Add a list of network devices to a fabric

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
	networkFabricId := int32(56) // int32 | The ID of the fabric
	networkDevicesToFabric := *openapiclient.NewNetworkDevicesToFabric([]float32{float32(123)}) // NetworkDevicesToFabric | The network device list containing the IDs of the network devices to add to the fabric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.AddNetworkDevicesToFabric(context.Background(), networkFabricId).NetworkDevicesToFabric(networkDevicesToFabric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.AddNetworkDevicesToFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNetworkDevicesToFabric`: NetworkFabric
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.AddNetworkDevicesToFabric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the fabric | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkDevicesToFabricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDevicesToFabric** | [**NetworkDevicesToFabric**](NetworkDevicesToFabric.md) | The network device list containing the IDs of the network devices to add to the fabric | 

### Return type

[**NetworkFabric**](NetworkFabric.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFabric

> NetworkFabric CreateNetworkFabric(ctx).CreateNetworkFabric(createNetworkFabric).Execute()

Create a new network fabric

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
	createNetworkFabric := *openapiclient.NewCreateNetworkFabric("My Network Fabric", openapiclient.NetworkFabric_fabricConfiguration{EthernetFabric: openapiclient.NewEthernetFabric(openapiclient.FabricType("ethernet"))}) // CreateNetworkFabric | The network fabric create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.CreateNetworkFabric(context.Background()).CreateNetworkFabric(createNetworkFabric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.CreateNetworkFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFabric`: NetworkFabric
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.CreateNetworkFabric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFabricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkFabric** | [**CreateNetworkFabric**](CreateNetworkFabric.md) | The network fabric create object | 

### Return type

[**NetworkFabric**](NetworkFabric.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFabricLink

> NetworkFabricLink CreateNetworkFabricLink(ctx, networkFabricId).CreateNetworkFabricLink(createNetworkFabricLink).Execute()

Create a new network fabric link

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
	networkFabricId := int32(56) // int32 | The ID of the fabric
	createNetworkFabricLink := *openapiclient.NewCreateNetworkFabricLink(float32(1), float32(2), "point-to-point", float32(0), "inherited", "disabled") // CreateNetworkFabricLink | The network fabric link object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.CreateNetworkFabricLink(context.Background(), networkFabricId).CreateNetworkFabricLink(createNetworkFabricLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.CreateNetworkFabricLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFabricLink`: NetworkFabricLink
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.CreateNetworkFabricLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the fabric | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFabricLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFabricLink** | [**CreateNetworkFabricLink**](CreateNetworkFabricLink.md) | The network fabric link object | 

### Return type

[**NetworkFabricLink**](NetworkFabricLink.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkFabric

> DeleteNetworkFabric(ctx, networkFabricId).Execute()

Delete a network fabric

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
	networkFabricId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkFabricAPI.DeleteNetworkFabric(context.Background(), networkFabricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.DeleteNetworkFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFabricRequest struct via the builder pattern


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


## DeleteNetworkFabricLink

> DeleteNetworkFabricLink(ctx, networkFabricId, networkFabricLinkId).Execute()

Remove a network fabric link

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
	networkFabricId := int32(56) // int32 | The ID of the fabric to remove the network device from
	networkFabricLinkId := int32(56) // int32 | The ID of the network fabric link

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkFabricAPI.DeleteNetworkFabricLink(context.Background(), networkFabricId, networkFabricLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.DeleteNetworkFabricLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the fabric to remove the network device from | 
**networkFabricLinkId** | **int32** | The ID of the network fabric link | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFabricLinkRequest struct via the builder pattern


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


## DeployNetworkFabric

> JobInfo DeployNetworkFabric(ctx, networkFabricId).NetworkFabricDeployOptions(networkFabricDeployOptions).Execute()

Deploys the specified network fabric



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
	networkFabricId := float32(8.14) // float32 | 
	networkFabricDeployOptions := *openapiclient.NewNetworkFabricDeployOptions(true) // NetworkFabricDeployOptions | Network fabric deploy options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.DeployNetworkFabric(context.Background(), networkFabricId).NetworkFabricDeployOptions(networkFabricDeployOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.DeployNetworkFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployNetworkFabric`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.DeployNetworkFabric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployNetworkFabricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkFabricDeployOptions** | [**NetworkFabricDeployOptions**](NetworkFabricDeployOptions.md) | Network fabric deploy options | 

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


## GetFabricNetworkDevices

> NetworkDevicePaginatedList GetFabricNetworkDevices(ctx, networkFabricId).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	networkFabricId := int32(56) // int32 | The ID of the fabric
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
	resp, r, err := apiClient.NetworkFabricAPI.GetFabricNetworkDevices(context.Background(), networkFabricId).Page(page).Limit(limit).FilterId(filterId).FilterStatus(filterStatus).FilterDatacenterName(filterDatacenterName).FilterSiteId(filterSiteId).FilterChassisIdentifier(filterChassisIdentifier).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterProvisionerType(filterProvisionerType).FilterPosition(filterPosition).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.GetFabricNetworkDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFabricNetworkDevices`: NetworkDevicePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.GetFabricNetworkDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the fabric | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFabricNetworkDevicesRequest struct via the builder pattern


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


## GetNetworkFabricById

> NetworkFabric GetNetworkFabricById(ctx, networkFabricId).Execute()

Get a network fabric by ID

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
	networkFabricId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.GetNetworkFabricById(context.Background(), networkFabricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.GetNetworkFabricById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFabricById`: NetworkFabric
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.GetNetworkFabricById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFabricByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkFabric**](NetworkFabric.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFabricLinks

> NetworkFabricLinkPaginatedList GetNetworkFabricLinks(ctx, networkFabricId).Page(page).Limit(limit).FilterId(filterId).FilterNetworkFabricId(filterNetworkFabricId).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get paginated Network Fabric Links

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
	networkFabricId := int32(56) // int32 | The ID of the fabric
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNetworkFabricId := []string{"Inner_example"} // []string | Filter by networkFabricId query param.  **Format:** filter.networkFabricId={$not}:OPERATION:VALUE    **Example:** filter.networkFabricId=$btw:John Doe&filter.networkFabricId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=networkFabricId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkFabricId  - status  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.GetNetworkFabricLinks(context.Background(), networkFabricId).Page(page).Limit(limit).FilterId(filterId).FilterNetworkFabricId(filterNetworkFabricId).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.GetNetworkFabricLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFabricLinks`: NetworkFabricLinkPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.GetNetworkFabricLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the fabric | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFabricLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNetworkFabricId** | **[]string** | Filter by networkFabricId query param.  **Format:** filter.networkFabricId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkFabricId&#x3D;$btw:John Doe&amp;filter.networkFabricId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;networkFabricId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkFabricId  - status  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   | 

### Return type

[**NetworkFabricLinkPaginatedList**](NetworkFabricLinkPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFabrics

> NetworkFabricPaginatedList GetNetworkFabrics(ctx).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List all network fabrics

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
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.  **Format:** filter.description={$not}:OPERATION:VALUE    **Example:** filter.description=$btw:John Doe&filter.description=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFabricConfigurationFabricType := []string{"Inner_example"} // []string | Filter by fabricConfiguration.fabricType query param.  **Format:** filter.fabricConfiguration.fabricType={$not}:OPERATION:VALUE    **Example:** filter.fabricConfiguration.fabricType=$btw:John Doe&filter.fabricConfiguration.fabricType=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  - name  - createdTimestamp  - updatedTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,description,status,siteId,fabricConfiguration.fabricType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - description  - status  - siteId  - fabricConfiguration.fabricType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.GetNetworkFabrics(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.GetNetworkFabrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFabrics`: NetworkFabricPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.GetNetworkFabrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFabricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDescription** | **[]string** | Filter by description query param.  **Format:** filter.description&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.description&#x3D;$btw:John Doe&amp;filter.description&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFabricConfigurationFabricType** | **[]string** | Filter by fabricConfiguration.fabricType query param.  **Format:** filter.fabricConfiguration.fabricType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricConfiguration.fabricType&#x3D;$btw:John Doe&amp;filter.fabricConfiguration.fabricType&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  - name  - createdTimestamp  - updatedTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,description,status,siteId,fabricConfiguration.fabricType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - description  - status  - siteId  - fabricConfiguration.fabricType  | 

### Return type

[**NetworkFabricPaginatedList**](NetworkFabricPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectNetworkFabricDeploy

> RejectNetworkFabricDeploy(ctx, networkFabricId).Execute()

Rejects the deployment of the specified network fabric



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
	networkFabricId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkFabricAPI.RejectNetworkFabricDeploy(context.Background(), networkFabricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.RejectNetworkFabricDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectNetworkFabricDeployRequest struct via the builder pattern


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


## RemoveNetworkDeviceFromFabric

> NetworkFabric RemoveNetworkDeviceFromFabric(ctx, networkFabricId, networkDeviceId).Execute()

Remove a network device from a fabric

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
	networkFabricId := int32(56) // int32 | The ID of the fabric to remove the network device from
	networkDeviceId := int32(56) // int32 | The ID of the network device to remove from the fabric

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.RemoveNetworkDeviceFromFabric(context.Background(), networkFabricId, networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.RemoveNetworkDeviceFromFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNetworkDeviceFromFabric`: NetworkFabric
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.RemoveNetworkDeviceFromFabric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the fabric to remove the network device from | 
**networkDeviceId** | **int32** | The ID of the network device to remove from the fabric | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkDeviceFromFabricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkFabric**](NetworkFabric.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFabric

> NetworkFabric UpdateNetworkFabric(ctx, networkFabricId).UpdateNetworkFabric(updateNetworkFabric).IfMatch(ifMatch).Execute()

Update a network fabric

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
	networkFabricId := int32(56) // int32 | The ID of the network fabric to update
	updateNetworkFabric := *openapiclient.NewUpdateNetworkFabric(openapiclient.NetworkFabric_fabricConfiguration{EthernetFabric: openapiclient.NewEthernetFabric(openapiclient.FabricType("ethernet"))}) // UpdateNetworkFabric | Network fabric updates
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricAPI.UpdateNetworkFabric(context.Background(), networkFabricId).UpdateNetworkFabric(updateNetworkFabric).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricAPI.UpdateNetworkFabric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFabric`: NetworkFabric
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricAPI.UpdateNetworkFabric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkFabricId** | **int32** | The ID of the network fabric to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFabricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFabric** | [**UpdateNetworkFabric**](UpdateNetworkFabric.md) | Network fabric updates | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkFabric**](NetworkFabric.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

