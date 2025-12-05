# \NetworkEndpointGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogicalNetworksToNetworkEndpointGroup**](NetworkEndpointGroupAPI.md#AddLogicalNetworksToNetworkEndpointGroup) | **Post** /api/v2/network-endpoint-groups/{networkEndpointGroupId}/logical-networks | Add a list of logical networks to a network endpoint group
[**CreateNetworkEndpointGroup**](NetworkEndpointGroupAPI.md#CreateNetworkEndpointGroup) | **Post** /api/v2/network-endpoint-groups | Create a new network endpoint group
[**DeleteNetworkEndpointGroup**](NetworkEndpointGroupAPI.md#DeleteNetworkEndpointGroup) | **Delete** /api/v2/network-endpoint-groups/{networkEndpointGroupId} | Delete a network endpoint group
[**GetNetworkEndpointGroupById**](NetworkEndpointGroupAPI.md#GetNetworkEndpointGroupById) | **Get** /api/v2/network-endpoint-groups/{networkEndpointGroupId} | Get a network endpoint group by ID
[**GetNetworkEndpointGroupLogicalNetwork**](NetworkEndpointGroupAPI.md#GetNetworkEndpointGroupLogicalNetwork) | **Get** /api/v2/network-endpoint-groups/{networkEndpointGroupId}/logical-networks/{logicalNetworkId} | Get a logical network by its ID
[**GetNetworkEndpointGroupLogicalNetworks**](NetworkEndpointGroupAPI.md#GetNetworkEndpointGroupLogicalNetworks) | **Get** /api/v2/network-endpoint-groups/{networkEndpointGroupId}/logical-networks | Get a network endpoint group with its logical networks
[**GetNetworkEndpointGroups**](NetworkEndpointGroupAPI.md#GetNetworkEndpointGroups) | **Get** /api/v2/network-endpoint-groups | List all network endpoint groups
[**RemoveLogicalNetworkFromNetworkEndpointGroup**](NetworkEndpointGroupAPI.md#RemoveLogicalNetworkFromNetworkEndpointGroup) | **Delete** /api/v2/network-endpoint-groups/{networkEndpointGroupId}/logical-networks/{logicalNetworkId} | Remove a logical network from a network endpoint group
[**UpdateNetworkEndpointGroup**](NetworkEndpointGroupAPI.md#UpdateNetworkEndpointGroup) | **Patch** /api/v2/network-endpoint-groups/{networkEndpointGroupId} | Update a network endpoint group
[**UpdateNetworkEndpointGroupLogicalNetwork**](NetworkEndpointGroupAPI.md#UpdateNetworkEndpointGroupLogicalNetwork) | **Patch** /api/v2/network-endpoint-groups/{networkEndpointGroupId}/logical-networks/{logicalNetworkId} | Update a logical network in a network endpoint group



## AddLogicalNetworksToNetworkEndpointGroup

> AddLogicalNetworksToNetworkEndpointGroup(ctx, networkEndpointGroupId).CreateNetworkEndpointGroupLogicalNetwork(createNetworkEndpointGroupLogicalNetwork).Execute()

Add a list of logical networks to a network endpoint group

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
	networkEndpointGroupId := int32(56) // int32 | The ID of the network endpoint group
	createNetworkEndpointGroupLogicalNetwork := *openapiclient.NewCreateNetworkEndpointGroupLogicalNetwork("1", true, openapiclient.NetworkEndpointGroupAllowedAccessMode("l2")) // CreateNetworkEndpointGroupLogicalNetwork | The logical network id and settings to add to the network endpoint group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkEndpointGroupAPI.AddLogicalNetworksToNetworkEndpointGroup(context.Background(), networkEndpointGroupId).CreateNetworkEndpointGroupLogicalNetwork(createNetworkEndpointGroupLogicalNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.AddLogicalNetworksToNetworkEndpointGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **int32** | The ID of the network endpoint group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLogicalNetworksToNetworkEndpointGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkEndpointGroupLogicalNetwork** | [**CreateNetworkEndpointGroupLogicalNetwork**](CreateNetworkEndpointGroupLogicalNetwork.md) | The logical network id and settings to add to the network endpoint group | 

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


## CreateNetworkEndpointGroup

> NetworkEndpointGroup CreateNetworkEndpointGroup(ctx).CreateNetworkEndpointGroup(createNetworkEndpointGroup).Execute()

Create a new network endpoint group

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
	createNetworkEndpointGroup := *openapiclient.NewCreateNetworkEndpointGroup("Default Network Endpoint Group") // CreateNetworkEndpointGroup | The network endpoint group create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.CreateNetworkEndpointGroup(context.Background()).CreateNetworkEndpointGroup(createNetworkEndpointGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.CreateNetworkEndpointGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkEndpointGroup`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.CreateNetworkEndpointGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkEndpointGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkEndpointGroup** | [**CreateNetworkEndpointGroup**](CreateNetworkEndpointGroup.md) | The network endpoint group create object | 

### Return type

[**NetworkEndpointGroup**](NetworkEndpointGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkEndpointGroup

> DeleteNetworkEndpointGroup(ctx, networkEndpointGroupId).Execute()

Delete a network endpoint group

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
	networkEndpointGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkEndpointGroupAPI.DeleteNetworkEndpointGroup(context.Background(), networkEndpointGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.DeleteNetworkEndpointGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkEndpointGroupRequest struct via the builder pattern


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


## GetNetworkEndpointGroupById

> NetworkEndpointGroup GetNetworkEndpointGroupById(ctx, networkEndpointGroupId).Execute()

Get a network endpoint group by ID

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
	networkEndpointGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.GetNetworkEndpointGroupById(context.Background(), networkEndpointGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.GetNetworkEndpointGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEndpointGroupById`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.GetNetworkEndpointGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEndpointGroupByIdRequest struct via the builder pattern


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


## GetNetworkEndpointGroupLogicalNetwork

> NetworkEndpointGroupLogicalNetwork GetNetworkEndpointGroupLogicalNetwork(ctx, networkEndpointGroupId, logicalNetworkId).Execute()

Get a logical network by its ID

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
	networkEndpointGroupId := int32(56) // int32 | The ID of the network endpoint group
	logicalNetworkId := int32(56) // int32 | The ID of the logical network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.GetNetworkEndpointGroupLogicalNetwork(context.Background(), networkEndpointGroupId, logicalNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.GetNetworkEndpointGroupLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEndpointGroupLogicalNetwork`: NetworkEndpointGroupLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.GetNetworkEndpointGroupLogicalNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **int32** | The ID of the network endpoint group | 
**logicalNetworkId** | **int32** | The ID of the logical network | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEndpointGroupLogicalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkEndpointGroupLogicalNetwork**](NetworkEndpointGroupLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEndpointGroupLogicalNetworks

> NetworkEndpointGroupLogicalNetworksList GetNetworkEndpointGroupLogicalNetworks(ctx, networkEndpointGroupId).Execute()

Get a network endpoint group with its logical networks

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
	networkEndpointGroupId := int32(56) // int32 | The id of the network endpoint group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.GetNetworkEndpointGroupLogicalNetworks(context.Background(), networkEndpointGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.GetNetworkEndpointGroupLogicalNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEndpointGroupLogicalNetworks`: NetworkEndpointGroupLogicalNetworksList
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.GetNetworkEndpointGroupLogicalNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **int32** | The id of the network endpoint group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEndpointGroupLogicalNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkEndpointGroupLogicalNetworksList**](NetworkEndpointGroupLogicalNetworksList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEndpointGroups

> NetworkEndpointGroupPaginatedList GetNetworkEndpointGroups(ctx).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List all network endpoint groups

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
	resp, r, err := apiClient.NetworkEndpointGroupAPI.GetNetworkEndpointGroups(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.GetNetworkEndpointGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEndpointGroups`: NetworkEndpointGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.GetNetworkEndpointGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEndpointGroupsRequest struct via the builder pattern


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

[**NetworkEndpointGroupPaginatedList**](NetworkEndpointGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLogicalNetworkFromNetworkEndpointGroup

> NetworkEndpointGroupLogicalNetwork RemoveLogicalNetworkFromNetworkEndpointGroup(ctx, networkEndpointGroupId, logicalNetworkId).Execute()

Remove a logical network from a network endpoint group

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
	networkEndpointGroupId := int32(56) // int32 | The ID of the network endpoint group
	logicalNetworkId := int32(56) // int32 | The ID of the logical network to remove from the network endpoint group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.RemoveLogicalNetworkFromNetworkEndpointGroup(context.Background(), networkEndpointGroupId, logicalNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.RemoveLogicalNetworkFromNetworkEndpointGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveLogicalNetworkFromNetworkEndpointGroup`: NetworkEndpointGroupLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.RemoveLogicalNetworkFromNetworkEndpointGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **int32** | The ID of the network endpoint group | 
**logicalNetworkId** | **int32** | The ID of the logical network to remove from the network endpoint group | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLogicalNetworkFromNetworkEndpointGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkEndpointGroupLogicalNetwork**](NetworkEndpointGroupLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkEndpointGroup

> NetworkEndpointGroup UpdateNetworkEndpointGroup(ctx, networkEndpointGroupId).UpdateNetworkEndpointGroup(updateNetworkEndpointGroup).IfMatch(ifMatch).Execute()

Update a network endpoint group

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
	networkEndpointGroupId := int32(56) // int32 | The ID of the network endpoint group to update
	updateNetworkEndpointGroup := *openapiclient.NewUpdateNetworkEndpointGroup() // UpdateNetworkEndpointGroup | Network endpoint group updates
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.UpdateNetworkEndpointGroup(context.Background(), networkEndpointGroupId).UpdateNetworkEndpointGroup(updateNetworkEndpointGroup).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.UpdateNetworkEndpointGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEndpointGroup`: NetworkEndpointGroup
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.UpdateNetworkEndpointGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **int32** | The ID of the network endpoint group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkEndpointGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkEndpointGroup** | [**UpdateNetworkEndpointGroup**](UpdateNetworkEndpointGroup.md) | Network endpoint group updates | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkEndpointGroup**](NetworkEndpointGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkEndpointGroupLogicalNetwork

> NetworkEndpointGroupLogicalNetwork UpdateNetworkEndpointGroupLogicalNetwork(ctx, networkEndpointGroupId, logicalNetworkId).UpdateNetworkEndpointGroupLogicalNetwork(updateNetworkEndpointGroupLogicalNetwork).IfMatch(ifMatch).Execute()

Update a logical network in a network endpoint group

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
	networkEndpointGroupId := int32(56) // int32 | The ID of the network endpoint group
	logicalNetworkId := int32(56) // int32 | The ID of the logical network
	updateNetworkEndpointGroupLogicalNetwork := *openapiclient.NewUpdateNetworkEndpointGroupLogicalNetwork() // UpdateNetworkEndpointGroupLogicalNetwork | The logical network id and settings to add to the network endpoint group
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.UpdateNetworkEndpointGroupLogicalNetwork(context.Background(), networkEndpointGroupId, logicalNetworkId).UpdateNetworkEndpointGroupLogicalNetwork(updateNetworkEndpointGroupLogicalNetwork).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.UpdateNetworkEndpointGroupLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEndpointGroupLogicalNetwork`: NetworkEndpointGroupLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.UpdateNetworkEndpointGroupLogicalNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkEndpointGroupId** | **int32** | The ID of the network endpoint group | 
**logicalNetworkId** | **int32** | The ID of the logical network | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkEndpointGroupLogicalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkEndpointGroupLogicalNetwork** | [**UpdateNetworkEndpointGroupLogicalNetwork**](UpdateNetworkEndpointGroupLogicalNetwork.md) | The logical network id and settings to add to the network endpoint group | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkEndpointGroupLogicalNetwork**](NetworkEndpointGroupLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

