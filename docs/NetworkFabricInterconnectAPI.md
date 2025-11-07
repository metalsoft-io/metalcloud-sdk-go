# \NetworkFabricInterconnectAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterconnectLink**](NetworkFabricInterconnectAPI.md#CreateInterconnectLink) | **Post** /api/v2/network-fabric-interconnects/{id}/links | Create a new link within an interconnect
[**CreateNetworkFabricInterconnect**](NetworkFabricInterconnectAPI.md#CreateNetworkFabricInterconnect) | **Post** /api/v2/network-fabric-interconnects | Create a new network fabric interconnect
[**DeleteInterconnectLink**](NetworkFabricInterconnectAPI.md#DeleteInterconnectLink) | **Delete** /api/v2/network-fabric-interconnects/{id}/links/{linkId} | Delete a specific fabric interconnect link
[**DeleteNetworkFabricInterconnect**](NetworkFabricInterconnectAPI.md#DeleteNetworkFabricInterconnect) | **Delete** /api/v2/network-fabric-interconnects/{id} | Delete a network fabric interconnect by ID
[**GetFabricInterconnectAvailableFabrics**](NetworkFabricInterconnectAPI.md#GetFabricInterconnectAvailableFabrics) | **Get** /api/v2/network-fabric-interconnects/{id}/fabrics-available | Get the available fabrics for a fabric interconnect
[**GetFabricInterconnectFabrics**](NetworkFabricInterconnectAPI.md#GetFabricInterconnectFabrics) | **Get** /api/v2/network-fabric-interconnects/{id}/fabrics | Get the fabrics for a fabric interconnect
[**GetInterconnectLink**](NetworkFabricInterconnectAPI.md#GetInterconnectLink) | **Get** /api/v2/network-fabric-interconnects/{id}/links/{linkId} | Get a specific fabric interconnect link
[**GetInterconnectLinks**](NetworkFabricInterconnectAPI.md#GetInterconnectLinks) | **Get** /api/v2/network-fabric-interconnects/{id}/links | Get all fabric interconnect links
[**GetNetworkFabricInterconnectById**](NetworkFabricInterconnectAPI.md#GetNetworkFabricInterconnectById) | **Get** /api/v2/network-fabric-interconnects/{id} | Get a network fabric interconnect by ID
[**GetNetworkFabricInterconnects**](NetworkFabricInterconnectAPI.md#GetNetworkFabricInterconnects) | **Get** /api/v2/network-fabric-interconnects | Get all network fabric interconnects
[**UpdateNetworkFabricInterconnect**](NetworkFabricInterconnectAPI.md#UpdateNetworkFabricInterconnect) | **Patch** /api/v2/network-fabric-interconnects/{id} | Update a network fabric interconnect



## CreateInterconnectLink

> NetworkFabricInterconnectLink CreateInterconnectLink(ctx, id).CreateNetworkFabricInterconnectLink(createNetworkFabricInterconnectLink).Execute()

Create a new link within an interconnect



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
	id := int32(56) // int32 | The ID of the network fabric interconnect
	createNetworkFabricInterconnectLink := *openapiclient.NewCreateNetworkFabricInterconnectLink(int32(2), int32(1), int32(4), int32(1)) // CreateNetworkFabricInterconnectLink | The network fabric link to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.CreateInterconnectLink(context.Background(), id).CreateNetworkFabricInterconnectLink(createNetworkFabricInterconnectLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.CreateInterconnectLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInterconnectLink`: NetworkFabricInterconnectLink
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.CreateInterconnectLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInterconnectLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFabricInterconnectLink** | [**CreateNetworkFabricInterconnectLink**](CreateNetworkFabricInterconnectLink.md) | The network fabric link to create | 

### Return type

[**NetworkFabricInterconnectLink**](NetworkFabricInterconnectLink.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFabricInterconnect

> NetworkFabricInterconnect CreateNetworkFabricInterconnect(ctx).CreateNetworkFabricInterconnect(createNetworkFabricInterconnect).Execute()

Create a new network fabric interconnect

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
	createNetworkFabricInterconnect := *openapiclient.NewCreateNetworkFabricInterconnect(openapiclient.NetworkFabricInterconnectType("dci-evpn"), "dci-evpn-interconnect-01") // CreateNetworkFabricInterconnect | The network fabric interconnect to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.CreateNetworkFabricInterconnect(context.Background()).CreateNetworkFabricInterconnect(createNetworkFabricInterconnect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.CreateNetworkFabricInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFabricInterconnect`: NetworkFabricInterconnect
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.CreateNetworkFabricInterconnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFabricInterconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkFabricInterconnect** | [**CreateNetworkFabricInterconnect**](CreateNetworkFabricInterconnect.md) | The network fabric interconnect to create | 

### Return type

[**NetworkFabricInterconnect**](NetworkFabricInterconnect.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterconnectLink

> DeleteInterconnectLink(ctx, id, linkId).IfMatch(ifMatch).Execute()

Delete a specific fabric interconnect link



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
	id := int32(56) // int32 | The ID of the network fabric interconnect
	linkId := int32(56) // int32 | The ID of the network fabric link
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkFabricInterconnectAPI.DeleteInterconnectLink(context.Background(), id, linkId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.DeleteInterconnectLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect | 
**linkId** | **int32** | The ID of the network fabric link | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterconnectLinkRequest struct via the builder pattern


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


## DeleteNetworkFabricInterconnect

> DeleteNetworkFabricInterconnect(ctx, id).Execute()

Delete a network fabric interconnect by ID

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
	id := int32(56) // int32 | The ID of the network fabric interconnect to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkFabricInterconnectAPI.DeleteNetworkFabricInterconnect(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.DeleteNetworkFabricInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFabricInterconnectRequest struct via the builder pattern


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


## GetFabricInterconnectAvailableFabrics

> NetworkFabricPaginatedList GetFabricInterconnectAvailableFabrics(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get the available fabrics for a fabric interconnect



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
	id := int32(56) // int32 | The ID of the network fabric interconnect
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
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.GetFabricInterconnectAvailableFabrics(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).FilterSiteId(filterSiteId).FilterFabricConfigurationFabricType(filterFabricConfigurationFabricType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.GetFabricInterconnectAvailableFabrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFabricInterconnectAvailableFabrics`: NetworkFabricPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.GetFabricInterconnectAvailableFabrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFabricInterconnectAvailableFabricsRequest struct via the builder pattern


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


## GetFabricInterconnectFabrics

> NetworkFabricList GetFabricInterconnectFabrics(ctx, id).Execute()

Get the fabrics for a fabric interconnect



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
	id := int32(56) // int32 | The ID of the network fabric interconnect

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.GetFabricInterconnectFabrics(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.GetFabricInterconnectFabrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFabricInterconnectFabrics`: NetworkFabricList
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.GetFabricInterconnectFabrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFabricInterconnectFabricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkFabricList**](NetworkFabricList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnectLink

> NetworkFabricInterconnectLink GetInterconnectLink(ctx, id, linkId).Execute()

Get a specific fabric interconnect link



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
	id := int32(56) // int32 | The ID of the network fabric interconnect
	linkId := int32(56) // int32 | The ID of the network fabric link

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.GetInterconnectLink(context.Background(), id, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.GetInterconnectLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterconnectLink`: NetworkFabricInterconnectLink
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.GetInterconnectLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect | 
**linkId** | **int32** | The ID of the network fabric link | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkFabricInterconnectLink**](NetworkFabricInterconnectLink.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnectLinks

> NetworkFabricInterconnectLinksPaginatedList GetInterconnectLinks(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterInterconnectId(filterInterconnectId).FilterFabricAId(filterFabricAId).FilterFabricANetworkEquipmentId(filterFabricANetworkEquipmentId).FilterFabricBId(filterFabricBId).FilterFabricBNetworkEquipmentId(filterFabricBNetworkEquipmentId).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get all fabric interconnect links



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
	id := int32(56) // int32 | The ID of the network fabric interconnect
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInterconnectId := []string{"Inner_example"} // []string | Filter by interconnectId query param.  **Format:** filter.interconnectId={$not}:OPERATION:VALUE    **Example:** filter.interconnectId=$btw:John Doe&filter.interconnectId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFabricAId := []string{"Inner_example"} // []string | Filter by fabricAId query param.  **Format:** filter.fabricAId={$not}:OPERATION:VALUE    **Example:** filter.fabricAId=$btw:John Doe&filter.fabricAId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFabricANetworkEquipmentId := []string{"Inner_example"} // []string | Filter by fabricANetworkEquipmentId query param.  **Format:** filter.fabricANetworkEquipmentId={$not}:OPERATION:VALUE    **Example:** filter.fabricANetworkEquipmentId=$btw:John Doe&filter.fabricANetworkEquipmentId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFabricBId := []string{"Inner_example"} // []string | Filter by fabricBId query param.  **Format:** filter.fabricBId={$not}:OPERATION:VALUE    **Example:** filter.fabricBId=$btw:John Doe&filter.fabricBId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFabricBNetworkEquipmentId := []string{"Inner_example"} // []string | Filter by fabricBNetworkEquipmentId query param.  **Format:** filter.fabricBNetworkEquipmentId={$not}:OPERATION:VALUE    **Example:** filter.fabricBNetworkEquipmentId=$btw:John Doe&filter.fabricBNetworkEquipmentId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=interconnectId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - interconnectId  - status  - fabricAId  - fabricANetworkEquipmentId  - fabricBId  - fabricBNetworkEquipmentId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,revision,status,interconnectId,fabricAId   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.GetInterconnectLinks(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterInterconnectId(filterInterconnectId).FilterFabricAId(filterFabricAId).FilterFabricANetworkEquipmentId(filterFabricANetworkEquipmentId).FilterFabricBId(filterFabricBId).FilterFabricBNetworkEquipmentId(filterFabricBNetworkEquipmentId).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.GetInterconnectLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterconnectLinks`: NetworkFabricInterconnectLinksPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.GetInterconnectLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInterconnectId** | **[]string** | Filter by interconnectId query param.  **Format:** filter.interconnectId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.interconnectId&#x3D;$btw:John Doe&amp;filter.interconnectId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFabricAId** | **[]string** | Filter by fabricAId query param.  **Format:** filter.fabricAId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricAId&#x3D;$btw:John Doe&amp;filter.fabricAId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFabricANetworkEquipmentId** | **[]string** | Filter by fabricANetworkEquipmentId query param.  **Format:** filter.fabricANetworkEquipmentId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricANetworkEquipmentId&#x3D;$btw:John Doe&amp;filter.fabricANetworkEquipmentId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFabricBId** | **[]string** | Filter by fabricBId query param.  **Format:** filter.fabricBId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricBId&#x3D;$btw:John Doe&amp;filter.fabricBId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFabricBNetworkEquipmentId** | **[]string** | Filter by fabricBNetworkEquipmentId query param.  **Format:** filter.fabricBNetworkEquipmentId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricBNetworkEquipmentId&#x3D;$btw:John Doe&amp;filter.fabricBNetworkEquipmentId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;interconnectId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - interconnectId  - status  - fabricAId  - fabricANetworkEquipmentId  - fabricBId  - fabricBNetworkEquipmentId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   | 
 **select_** | **string** | List of fields to select.  **Example:** id,revision,status,interconnectId,fabricAId   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**NetworkFabricInterconnectLinksPaginatedList**](NetworkFabricInterconnectLinksPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFabricInterconnectById

> NetworkFabricInterconnect GetNetworkFabricInterconnectById(ctx, id).Execute()

Get a network fabric interconnect by ID



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
	id := int32(56) // int32 | The ID of the network fabric interconnect to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.GetNetworkFabricInterconnectById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.GetNetworkFabricInterconnectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFabricInterconnectById`: NetworkFabricInterconnect
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.GetNetworkFabricInterconnectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFabricInterconnectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkFabricInterconnect**](NetworkFabricInterconnect.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFabricInterconnects

> NetworkFabricInterconnectPaginatedList GetNetworkFabricInterconnects(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all network fabric interconnects



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$btw:John Doe&filter.label=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.  **Format:** filter.description={$not}:OPERATION:VALUE    **Example:** filter.description=$btw:John Doe&filter.description=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - createdTimestamp  - updatedTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.GetNetworkFabricInterconnects(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDescription(filterDescription).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.GetNetworkFabricInterconnects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFabricInterconnects`: NetworkFabricInterconnectPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.GetNetworkFabricInterconnects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFabricInterconnectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$btw:John Doe&amp;filter.label&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDescription** | **[]string** | Filter by description query param.  **Format:** filter.description&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.description&#x3D;$btw:John Doe&amp;filter.description&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - createdTimestamp  - updatedTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  | 

### Return type

[**NetworkFabricInterconnectPaginatedList**](NetworkFabricInterconnectPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFabricInterconnect

> NetworkFabricInterconnect UpdateNetworkFabricInterconnect(ctx, id).UpdateNetworkFabricInterconnect(updateNetworkFabricInterconnect).IfMatch(ifMatch).Execute()

Update a network fabric interconnect

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
	id := int32(56) // int32 | The ID of the network fabric interconnect to update
	updateNetworkFabricInterconnect := *openapiclient.NewUpdateNetworkFabricInterconnect() // UpdateNetworkFabricInterconnect | The network fabric interconnect update entity
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkFabricInterconnectAPI.UpdateNetworkFabricInterconnect(context.Background(), id).UpdateNetworkFabricInterconnect(updateNetworkFabricInterconnect).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkFabricInterconnectAPI.UpdateNetworkFabricInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFabricInterconnect`: NetworkFabricInterconnect
	fmt.Fprintf(os.Stdout, "Response from `NetworkFabricInterconnectAPI.UpdateNetworkFabricInterconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the network fabric interconnect to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFabricInterconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFabricInterconnect** | [**UpdateNetworkFabricInterconnect**](UpdateNetworkFabricInterconnect.md) | The network fabric interconnect update entity | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkFabricInterconnect**](NetworkFabricInterconnect.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

