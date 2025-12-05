# \EndpointAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateEndpoints**](EndpointAPI.md#BulkCreateEndpoints) | **Post** /api/v2/endpoints/actions/bulk-create | Bulk create endpoints
[**CreateEndpoint**](EndpointAPI.md#CreateEndpoint) | **Post** /api/v2/endpoints | Create a new endpoint
[**CreateEndpointInterface**](EndpointAPI.md#CreateEndpointInterface) | **Post** /api/v2/endpoints/{endpointId}/interfaces | Create a new endpoint interface
[**DeleteEndpoint**](EndpointAPI.md#DeleteEndpoint) | **Delete** /api/v2/endpoints/{endpointId} | Delete endpoint
[**DeleteEndpointInterface**](EndpointAPI.md#DeleteEndpointInterface) | **Delete** /api/v2/endpoints/{endpointId}/interfaces/{endpointInterfaceId} | Delete endpoint interface
[**GetEndpointById**](EndpointAPI.md#GetEndpointById) | **Get** /api/v2/endpoints/{endpointId} | Get endpoint details
[**GetEndpointInterfaceById**](EndpointAPI.md#GetEndpointInterfaceById) | **Get** /api/v2/endpoints/{endpointId}/interfaces/{endpointInterfaceId} | Get endpoint interface details
[**GetEndpointInterfaces**](EndpointAPI.md#GetEndpointInterfaces) | **Get** /api/v2/endpoints/{endpointId}/interfaces | List endpoint interfaces
[**GetEndpoints**](EndpointAPI.md#GetEndpoints) | **Get** /api/v2/endpoints | List endpoints
[**GetNetworkDeviceInterfacesAndEndpoints**](EndpointAPI.md#GetNetworkDeviceInterfacesAndEndpoints) | **Get** /api/v2/endpoints/network-devices/{networkDeviceId}/interfaces | List network device interfaces associated with endpoints
[**UpdateEndpoint**](EndpointAPI.md#UpdateEndpoint) | **Patch** /api/v2/endpoints/{endpointId} | Update endpoint
[**UpdateEndpointInterface**](EndpointAPI.md#UpdateEndpointInterface) | **Patch** /api/v2/endpoints/{endpointId}/interfaces/{endpointInterfaceId} | Update endpoint interface



## BulkCreateEndpoints

> BulkCreateEndpoints(ctx).BulkCreateEndpoints(bulkCreateEndpoints).Execute()

Bulk create endpoints



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
	bulkCreateEndpoints := *openapiclient.NewBulkCreateEndpoints([]openapiclient.CreateEndpoint{*openapiclient.NewCreateEndpoint(int32(123), "Server 1", "lan-finance-1")}) // BulkCreateEndpoints | An object containing an array of endpoints to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointAPI.BulkCreateEndpoints(context.Background()).BulkCreateEndpoints(bulkCreateEndpoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.BulkCreateEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkCreateEndpoints** | [**BulkCreateEndpoints**](BulkCreateEndpoints.md) | An object containing an array of endpoints to create | 

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


## CreateEndpoint

> Endpoint CreateEndpoint(ctx).CreateEndpoint(createEndpoint).Execute()

Create a new endpoint



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
	createEndpoint := *openapiclient.NewCreateEndpoint(int32(123), "Server 1", "lan-finance-1") // CreateEndpoint | The endpoint to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.CreateEndpoint(context.Background()).CreateEndpoint(createEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.CreateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpoint`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.CreateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpoint** | [**CreateEndpoint**](CreateEndpoint.md) | The endpoint to create | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpointInterface

> EndpointInterface CreateEndpointInterface(ctx, endpointId).CreateEndpointInterface(createEndpointInterface).Execute()

Create a new endpoint interface



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
	endpointId := int32(56) // int32 | 
	createEndpointInterface := *openapiclient.NewCreateEndpointInterface(float32(1)) // CreateEndpointInterface | The endpoint interface to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.CreateEndpointInterface(context.Background(), endpointId).CreateEndpointInterface(createEndpointInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.CreateEndpointInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointInterface`: EndpointInterface
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.CreateEndpointInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEndpointInterface** | [**CreateEndpointInterface**](CreateEndpointInterface.md) | The endpoint interface to create | 

### Return type

[**EndpointInterface**](EndpointInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpoint

> DeleteEndpoint(ctx, endpointId).IfMatch(ifMatch).Execute()

Delete endpoint



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
	endpointId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointAPI.DeleteEndpoint(context.Background(), endpointId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.DeleteEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointRequest struct via the builder pattern


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


## DeleteEndpointInterface

> DeleteEndpointInterface(ctx, endpointId, endpointInterfaceId).IfMatch(ifMatch).Execute()

Delete endpoint interface



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
	endpointId := int32(56) // int32 | 
	endpointInterfaceId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointAPI.DeleteEndpointInterface(context.Background(), endpointId, endpointInterfaceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.DeleteEndpointInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 
**endpointInterfaceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointInterfaceRequest struct via the builder pattern


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


## GetEndpointById

> Endpoint GetEndpointById(ctx, endpointId).Execute()

Get endpoint details



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
	endpointId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.GetEndpointById(context.Background(), endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetEndpointById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointById`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetEndpointById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInterfaceById

> EndpointInterface GetEndpointInterfaceById(ctx, endpointId, endpointInterfaceId).Execute()

Get endpoint interface details



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
	endpointId := int32(56) // int32 | 
	endpointInterfaceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.GetEndpointInterfaceById(context.Background(), endpointId, endpointInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetEndpointInterfaceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInterfaceById`: EndpointInterface
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetEndpointInterfaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 
**endpointInterfaceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInterfaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EndpointInterface**](EndpointInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInterfaces

> EndpointInterfacePaginatedList GetEndpointInterfaces(ctx, endpointId).Page(page).Limit(limit).FilterId(filterId).FilterMacAddress(filterMacAddress).FilterCreatedTimestamp(filterCreatedTimestamp).FilterUpdatedTimestamp(filterUpdatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List endpoint interfaces



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
	endpointId := int32(56) // int32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterMacAddress := []string{"Inner_example"} // []string | Filter by macAddress query param.  **Format:** filter.macAddress={$not}:OPERATION:VALUE    **Example:** filter.macAddress=$btw:John Doe&filter.macAddress=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp={$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp=$btw:John Doe&filter.createdTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedTimestamp := []string{"Inner_example"} // []string | Filter by updatedTimestamp query param.  **Format:** filter.updatedTimestamp={$not}:OPERATION:VALUE    **Example:** filter.updatedTimestamp=$btw:John Doe&filter.updatedTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdTimestamp:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdTimestamp  - updatedTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** macAddress   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - macAddress  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.GetEndpointInterfaces(context.Background(), endpointId).Page(page).Limit(limit).FilterId(filterId).FilterMacAddress(filterMacAddress).FilterCreatedTimestamp(filterCreatedTimestamp).FilterUpdatedTimestamp(filterUpdatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetEndpointInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInterfaces`: EndpointInterfacePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetEndpointInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterMacAddress** | **[]string** | Filter by macAddress query param.  **Format:** filter.macAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.macAddress&#x3D;$btw:John Doe&amp;filter.macAddress&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp&#x3D;$btw:John Doe&amp;filter.createdTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdatedTimestamp** | **[]string** | Filter by updatedTimestamp query param.  **Format:** filter.updatedTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updatedTimestamp&#x3D;$btw:John Doe&amp;filter.updatedTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdTimestamp:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdTimestamp  - updatedTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** macAddress   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - macAddress  | 

### Return type

[**EndpointInterfacePaginatedList**](EndpointInterfacePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoints

> EndpointPaginatedList GetEndpoints(ctx).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterExternalId(filterExternalId).FilterCreatedTimestamp(filterCreatedTimestamp).FilterUpdatedTimestamp(filterUpdatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List endpoints



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterExternalId := []string{"Inner_example"} // []string | Filter by externalId query param.  **Format:** filter.externalId={$not}:OPERATION:VALUE    **Example:** filter.externalId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp={$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp=$btw:John Doe&filter.createdTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedTimestamp := []string{"Inner_example"} // []string | Filter by updatedTimestamp query param.  **Format:** filter.updatedTimestamp={$not}:OPERATION:VALUE    **Example:** filter.updatedTimestamp=$btw:John Doe&filter.updatedTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdTimestamp:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdTimestamp  - updatedTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.GetEndpoints(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterExternalId(filterExternalId).FilterCreatedTimestamp(filterCreatedTimestamp).FilterUpdatedTimestamp(filterUpdatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpoints`: EndpointPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterExternalId** | **[]string** | Filter by externalId query param.  **Format:** filter.externalId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.externalId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp&#x3D;$btw:John Doe&amp;filter.createdTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdatedTimestamp** | **[]string** | Filter by updatedTimestamp query param.  **Format:** filter.updatedTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updatedTimestamp&#x3D;$btw:John Doe&amp;filter.updatedTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdTimestamp:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdTimestamp  - updatedTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  | 

### Return type

[**EndpointPaginatedList**](EndpointPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceInterfacesAndEndpoints

> NetworkDeviceEndpointInterfaces GetNetworkDeviceInterfacesAndEndpoints(ctx, networkDeviceId).Execute()

List network device interfaces associated with endpoints

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
	networkDeviceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.GetNetworkDeviceInterfacesAndEndpoints(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetNetworkDeviceInterfacesAndEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceInterfacesAndEndpoints`: NetworkDeviceEndpointInterfaces
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetNetworkDeviceInterfacesAndEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceInterfacesAndEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceEndpointInterfaces**](NetworkDeviceEndpointInterfaces.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoint

> Endpoint UpdateEndpoint(ctx, endpointId).UpdateEndpoint(updateEndpoint).IfMatch(ifMatch).Execute()

Update endpoint



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
	endpointId := int32(56) // int32 | 
	updateEndpoint := *openapiclient.NewUpdateEndpoint() // UpdateEndpoint | The endpoint configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.UpdateEndpoint(context.Background(), endpointId).UpdateEndpoint(updateEndpoint).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.UpdateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpoint`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.UpdateEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEndpoint** | [**UpdateEndpoint**](UpdateEndpoint.md) | The endpoint configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInterface

> EndpointInterface UpdateEndpointInterface(ctx, endpointId, endpointInterfaceId).UpdateEndpointInterface(updateEndpointInterface).IfMatch(ifMatch).Execute()

Update endpoint interface



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
	endpointId := int32(56) // int32 | 
	endpointInterfaceId := int32(56) // int32 | 
	updateEndpointInterface := *openapiclient.NewUpdateEndpointInterface() // UpdateEndpointInterface | The endpoint interface configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.UpdateEndpointInterface(context.Background(), endpointId, endpointInterfaceId).UpdateEndpointInterface(updateEndpointInterface).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.UpdateEndpointInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointInterface`: EndpointInterface
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.UpdateEndpointInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** |  | 
**endpointInterfaceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEndpointInterface** | [**UpdateEndpointInterface**](UpdateEndpointInterface.md) | The endpoint interface configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**EndpointInterface**](EndpointInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

