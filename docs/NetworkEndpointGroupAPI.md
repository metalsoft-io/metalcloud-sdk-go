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

> CreateNetworkEndpointGroupLogicalNetwork AddLogicalNetworksToNetworkEndpointGroup(ctx, networkEndpointGroupId).CreateNetworkEndpointGroupLogicalNetwork(createNetworkEndpointGroupLogicalNetwork).Execute()

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
	createNetworkEndpointGroupLogicalNetwork := *openapiclient.NewCreateNetworkEndpointGroupLogicalNetwork("1", true, "l2") // CreateNetworkEndpointGroupLogicalNetwork | The logical network id and settings to add to the network endpoint group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.AddLogicalNetworksToNetworkEndpointGroup(context.Background(), networkEndpointGroupId).CreateNetworkEndpointGroupLogicalNetwork(createNetworkEndpointGroupLogicalNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.AddLogicalNetworksToNetworkEndpointGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLogicalNetworksToNetworkEndpointGroup`: CreateNetworkEndpointGroupLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworkEndpointGroupAPI.AddLogicalNetworksToNetworkEndpointGroup`: %v\n", resp)
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

[**CreateNetworkEndpointGroupLogicalNetwork**](CreateNetworkEndpointGroupLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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

> NetworkEndpointGroupLogicalNetworkDto GetNetworkEndpointGroupLogicalNetwork(ctx, networkEndpointGroupId, logicalNetworkId).Execute()

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
	// response from `GetNetworkEndpointGroupLogicalNetwork`: NetworkEndpointGroupLogicalNetworkDto
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

[**NetworkEndpointGroupLogicalNetworkDto**](NetworkEndpointGroupLogicalNetworkDto.md)

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

> NetworkEndpointGroupPaginatedList GetNetworkEndpointGroups(ctx).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.           <p>              <b>Format: </b> filter.description={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.description=$not:$like:John Doe&filter.description=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterFabricConfigurationFabricType := []string{"Inner_example"} // []string | Filter by fabricConfiguration.fabricType query param.           <p>              <b>Format: </b> filter.fabricConfiguration.fabricType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.fabricConfiguration.fabricType=$not:$like:John Doe&filter.fabricConfiguration.fabricType=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>siteId</li> <li>name</li> <li>createdTimestamp</li> <li>updatedTimestamp</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> name,description,siteId,fabricConfiguration.fabricType           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>name</li> <li>description</li> <li>siteId</li> <li>fabricConfiguration.fabricType</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.GetNetworkEndpointGroups(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterDescription** | **[]string** | Filter by description query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.description&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.description&#x3D;$not:$like:John Doe&amp;filter.description&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterFabricConfigurationFabricType** | **[]string** | Filter by fabricConfiguration.fabricType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.fabricConfiguration.fabricType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.fabricConfiguration.fabricType&#x3D;$not:$like:John Doe&amp;filter.fabricConfiguration.fabricType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;createdTimestamp&lt;/li&gt; &lt;li&gt;updatedTimestamp&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name,description,siteId,fabricConfiguration.fabricType           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt; &lt;li&gt;description&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;fabricConfiguration.fabricType&lt;/li&gt;&lt;/ul&gt;          | 

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

> NetworkEndpointGroupLogicalNetworkDto RemoveLogicalNetworkFromNetworkEndpointGroup(ctx, networkEndpointGroupId, logicalNetworkId).Execute()

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
	// response from `RemoveLogicalNetworkFromNetworkEndpointGroup`: NetworkEndpointGroupLogicalNetworkDto
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

[**NetworkEndpointGroupLogicalNetworkDto**](NetworkEndpointGroupLogicalNetworkDto.md)

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

> NetworkEndpointGroupLogicalNetworkDto UpdateNetworkEndpointGroupLogicalNetwork(ctx, networkEndpointGroupId, logicalNetworkId).UpdateNetworkEndpointGroupLogicalNetworkDto(updateNetworkEndpointGroupLogicalNetworkDto).IfMatch(ifMatch).Execute()

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
	updateNetworkEndpointGroupLogicalNetworkDto := *openapiclient.NewUpdateNetworkEndpointGroupLogicalNetworkDto() // UpdateNetworkEndpointGroupLogicalNetworkDto | The logical network id and settings to add to the network endpoint group
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkEndpointGroupAPI.UpdateNetworkEndpointGroupLogicalNetwork(context.Background(), networkEndpointGroupId, logicalNetworkId).UpdateNetworkEndpointGroupLogicalNetworkDto(updateNetworkEndpointGroupLogicalNetworkDto).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkEndpointGroupAPI.UpdateNetworkEndpointGroupLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEndpointGroupLogicalNetwork`: NetworkEndpointGroupLogicalNetworkDto
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


 **updateNetworkEndpointGroupLogicalNetworkDto** | [**UpdateNetworkEndpointGroupLogicalNetworkDto**](UpdateNetworkEndpointGroupLogicalNetworkDto.md) | The logical network id and settings to add to the network endpoint group | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkEndpointGroupLogicalNetworkDto**](NetworkEndpointGroupLogicalNetworkDto.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

