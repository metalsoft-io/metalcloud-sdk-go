# \LogicalNetworkInterconnectAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogicalNetworkToLogicalNetworkInterconnect**](LogicalNetworkInterconnectAPI.md#AddLogicalNetworkToLogicalNetworkInterconnect) | **Post** /api/v2/logical-network-interconnects/{id}/links | Add a logical network to a logical network interconnect
[**CreateLogicalNetworkInterconnect**](LogicalNetworkInterconnectAPI.md#CreateLogicalNetworkInterconnect) | **Post** /api/v2/logical-network-interconnects | Create a new logical network interconnect
[**DeleteLogicalNetworkInterconnect**](LogicalNetworkInterconnectAPI.md#DeleteLogicalNetworkInterconnect) | **Delete** /api/v2/logical-network-interconnects/{id} | Delete a logical network interconnect by ID
[**GetLogicalNetworkInterconnectById**](LogicalNetworkInterconnectAPI.md#GetLogicalNetworkInterconnectById) | **Get** /api/v2/logical-network-interconnects/{id} | Get a logical network interconnect by ID
[**GetLogicalNetworkInterconnectLinkById**](LogicalNetworkInterconnectAPI.md#GetLogicalNetworkInterconnectLinkById) | **Get** /api/v2/logical-network-interconnects/{id}/links/{linkId} | Get a logical network interconnect link association by link id
[**GetLogicalNetworkInterconnectLinks**](LogicalNetworkInterconnectAPI.md#GetLogicalNetworkInterconnectLinks) | **Get** /api/v2/logical-network-interconnects/{id}/links | Get all logical network associated to a logical network interconnect
[**GetLogicalNetworkInterconnects**](LogicalNetworkInterconnectAPI.md#GetLogicalNetworkInterconnects) | **Get** /api/v2/logical-network-interconnects | Get all logical network interconnects
[**RemoveLogicalNetworkFromLogicalNetworkInterconnect**](LogicalNetworkInterconnectAPI.md#RemoveLogicalNetworkFromLogicalNetworkInterconnect) | **Delete** /api/v2/logical-network-interconnects/{id}/links/{linkId} | Remove a logical network from a logical network interconnect
[**UpdateLogicalNetworkInterconnect**](LogicalNetworkInterconnectAPI.md#UpdateLogicalNetworkInterconnect) | **Patch** /api/v2/logical-network-interconnects/{id} | Update a logical network interconnect



## AddLogicalNetworkToLogicalNetworkInterconnect

> LogicalNetworkInterconnectLogicalNetwork AddLogicalNetworkToLogicalNetworkInterconnect(ctx, id).AddLogicalNetworkToInterconnect(addLogicalNetworkToInterconnect).Execute()

Add a logical network to a logical network interconnect



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
	id := int32(56) // int32 | The id of the logical network interconnect
	addLogicalNetworkToInterconnect := *openapiclient.NewAddLogicalNetworkToInterconnect(int32(1)) // AddLogicalNetworkToInterconnect | The logical network to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkInterconnectAPI.AddLogicalNetworkToLogicalNetworkInterconnect(context.Background(), id).AddLogicalNetworkToInterconnect(addLogicalNetworkToInterconnect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.AddLogicalNetworkToLogicalNetworkInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLogicalNetworkToLogicalNetworkInterconnect`: LogicalNetworkInterconnectLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkInterconnectAPI.AddLogicalNetworkToLogicalNetworkInterconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the logical network interconnect | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLogicalNetworkToLogicalNetworkInterconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addLogicalNetworkToInterconnect** | [**AddLogicalNetworkToInterconnect**](AddLogicalNetworkToInterconnect.md) | The logical network to add | 

### Return type

[**LogicalNetworkInterconnectLogicalNetwork**](LogicalNetworkInterconnectLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkInterconnect

> LogicalNetworkInterconnect CreateLogicalNetworkInterconnect(ctx).CreateLogicalNetworkInterconnect(createLogicalNetworkInterconnect).Execute()

Create a new logical network interconnect

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
	createLogicalNetworkInterconnect := *openapiclient.NewCreateLogicalNetworkInterconnect("logical-network-interconnect-01", "Primary Logical Network Interconnect", int32(1)) // CreateLogicalNetworkInterconnect | The logical network interconnect to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkInterconnectAPI.CreateLogicalNetworkInterconnect(context.Background()).CreateLogicalNetworkInterconnect(createLogicalNetworkInterconnect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.CreateLogicalNetworkInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkInterconnect`: LogicalNetworkInterconnect
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkInterconnectAPI.CreateLogicalNetworkInterconnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkInterconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogicalNetworkInterconnect** | [**CreateLogicalNetworkInterconnect**](CreateLogicalNetworkInterconnect.md) | The logical network interconnect to create | 

### Return type

[**LogicalNetworkInterconnect**](LogicalNetworkInterconnect.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkInterconnect

> DeleteLogicalNetworkInterconnect(ctx, id).IfMatch(ifMatch).Execute()

Delete a logical network interconnect by ID

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
	id := int32(56) // int32 | The ID of the logical network interconnect to delete
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkInterconnectAPI.DeleteLogicalNetworkInterconnect(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.DeleteLogicalNetworkInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the logical network interconnect to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkInterconnectRequest struct via the builder pattern


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


## GetLogicalNetworkInterconnectById

> LogicalNetworkInterconnect GetLogicalNetworkInterconnectById(ctx, id).Execute()

Get a logical network interconnect by ID



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
	id := int32(56) // int32 | The ID of the logical network interconnect to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkInterconnectById`: LogicalNetworkInterconnect
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the logical network interconnect to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkInterconnectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogicalNetworkInterconnect**](LogicalNetworkInterconnect.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkInterconnectLinkById

> LogicalNetworkInterconnectLogicalNetwork GetLogicalNetworkInterconnectLinkById(ctx, id, linkId).Execute()

Get a logical network interconnect link association by link id



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
	id := int32(56) // int32 | The id of the logical network interconnect
	linkId := int32(56) // int32 | The id of the link

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectLinkById(context.Background(), id, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkInterconnectLinkById`: LogicalNetworkInterconnectLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectLinkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the logical network interconnect | 
**linkId** | **int32** | The id of the link | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkInterconnectLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LogicalNetworkInterconnectLogicalNetwork**](LogicalNetworkInterconnectLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkInterconnectLinks

> LogicalNetworkInterconnectLogicalNetworkPaginatedList GetLogicalNetworkInterconnectLinks(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterLogicalNetworkId(filterLogicalNetworkId).FilterStatus(filterStatus).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all logical network associated to a logical network interconnect



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
	id := int32(56) // int32 | The ID of the logical network interconnect
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterLogicalNetworkId := []string{"Inner_example"} // []string | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId=$btw:John Doe&filter.logicalNetworkId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.  **Format:** filter.createdAt={$not}:OPERATION:VALUE    **Example:** filter.createdAt=$btw:John Doe&filter.createdAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.  **Format:** filter.updatedAt={$not}:OPERATION:VALUE    **Example:** filter.updatedAt=$btw:John Doe&filter.updatedAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=logicalNetworkId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - logicalNetworkId  - createdAt  - updatedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectLinks(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterLogicalNetworkId(filterLogicalNetworkId).FilterStatus(filterStatus).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkInterconnectLinks`: LogicalNetworkInterconnectLogicalNetworkPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnectLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the logical network interconnect | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkInterconnectLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterLogicalNetworkId** | **[]string** | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId&#x3D;$btw:John Doe&amp;filter.logicalNetworkId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.  **Format:** filter.createdAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdAt&#x3D;$btw:John Doe&amp;filter.createdAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.  **Format:** filter.updatedAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updatedAt&#x3D;$btw:John Doe&amp;filter.updatedAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;logicalNetworkId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - logicalNetworkId  - createdAt  - updatedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   | 

### Return type

[**LogicalNetworkInterconnectLogicalNetworkPaginatedList**](LogicalNetworkInterconnectLogicalNetworkPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkInterconnects

> LogicalNetworkInterconnectPaginatedList GetLogicalNetworkInterconnects(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterKind(filterKind).FilterStatus(filterStatus).FilterFabricInterconnectId(filterFabricInterconnectId).FilterTransportId(filterTransportId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all logical network interconnects



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
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFabricInterconnectId := []string{"Inner_example"} // []string | Filter by fabricInterconnectId query param.  **Format:** filter.fabricInterconnectId={$not}:OPERATION:VALUE    **Example:** filter.fabricInterconnectId=$btw:John Doe&filter.fabricInterconnectId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterTransportId := []string{"Inner_example"} // []string | Filter by transportId query param.  **Format:** filter.transportId={$not}:OPERATION:VALUE    **Example:** filter.transportId=$btw:John Doe&filter.transportId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - label  - createdAt  - updatedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnects(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterKind(filterKind).FilterStatus(filterStatus).FilterFabricInterconnectId(filterFabricInterconnectId).FilterTransportId(filterTransportId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkInterconnects`: LogicalNetworkInterconnectPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkInterconnectAPI.GetLogicalNetworkInterconnects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkInterconnectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$btw:John Doe&amp;filter.label&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFabricInterconnectId** | **[]string** | Filter by fabricInterconnectId query param.  **Format:** filter.fabricInterconnectId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricInterconnectId&#x3D;$btw:John Doe&amp;filter.fabricInterconnectId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterTransportId** | **[]string** | Filter by transportId query param.  **Format:** filter.transportId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.transportId&#x3D;$btw:John Doe&amp;filter.transportId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - label  - createdAt  - updatedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  | 

### Return type

[**LogicalNetworkInterconnectPaginatedList**](LogicalNetworkInterconnectPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLogicalNetworkFromLogicalNetworkInterconnect

> RemoveLogicalNetworkFromLogicalNetworkInterconnect(ctx, id, linkId).Execute()

Remove a logical network from a logical network interconnect



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
	id := int32(56) // int32 | The id of the logical network interconnect
	linkId := int32(56) // int32 | The id of link association to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkInterconnectAPI.RemoveLogicalNetworkFromLogicalNetworkInterconnect(context.Background(), id, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.RemoveLogicalNetworkFromLogicalNetworkInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the logical network interconnect | 
**linkId** | **int32** | The id of link association to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLogicalNetworkFromLogicalNetworkInterconnectRequest struct via the builder pattern


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


## UpdateLogicalNetworkInterconnect

> LogicalNetworkInterconnect UpdateLogicalNetworkInterconnect(ctx, id).UpdateLogicalNetworkInterconnectDto(updateLogicalNetworkInterconnectDto).IfMatch(ifMatch).Execute()

Update a logical network interconnect

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
	id := int32(56) // int32 | The ID of the logical network interconnect to update
	updateLogicalNetworkInterconnectDto := *openapiclient.NewUpdateLogicalNetworkInterconnectDto() // UpdateLogicalNetworkInterconnectDto | The logical network interconnect update entity
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkInterconnectAPI.UpdateLogicalNetworkInterconnect(context.Background(), id).UpdateLogicalNetworkInterconnectDto(updateLogicalNetworkInterconnectDto).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkInterconnectAPI.UpdateLogicalNetworkInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalNetworkInterconnect`: LogicalNetworkInterconnect
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkInterconnectAPI.UpdateLogicalNetworkInterconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the logical network interconnect to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogicalNetworkInterconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLogicalNetworkInterconnectDto** | [**UpdateLogicalNetworkInterconnectDto**](UpdateLogicalNetworkInterconnectDto.md) | The logical network interconnect update entity | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**LogicalNetworkInterconnect**](LogicalNetworkInterconnect.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

