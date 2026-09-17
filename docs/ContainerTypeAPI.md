# \ContainerTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerType**](ContainerTypeAPI.md#CreateContainerType) | **Post** /api/v2/container-types | Creates a Container Type
[**DeleteContainerType**](ContainerTypeAPI.md#DeleteContainerType) | **Delete** /api/v2/container-types/{containerTypeId} | Deletes a Container Type
[**GetContainerType**](ContainerTypeAPI.md#GetContainerType) | **Get** /api/v2/container-types/{containerTypeId} | Get Container Type information
[**GetContainerTypes**](ContainerTypeAPI.md#GetContainerTypes) | **Get** /api/v2/container-types | Get all Container Types
[**GetContainersByContainerType**](ContainerTypeAPI.md#GetContainersByContainerType) | **Get** /api/v2/container-types/{containerTypeId}/containers | Returns all Containers linked to the Container Type
[**UpdateContainerType**](ContainerTypeAPI.md#UpdateContainerType) | **Patch** /api/v2/container-types/{containerTypeId} | Updates Container Type information



## CreateContainerType

> ContainerType CreateContainerType(ctx).CreateContainerType(createContainerType).Execute()

Creates a Container Type



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
	createContainerType := *openapiclient.NewCreateContainerType("Name_example", float32(2), float32(2)) // CreateContainerType | The Container Type create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerTypeAPI.CreateContainerType(context.Background()).CreateContainerType(createContainerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerTypeAPI.CreateContainerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerType`: ContainerType
	fmt.Fprintf(os.Stdout, "Response from `ContainerTypeAPI.CreateContainerType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContainerType** | [**CreateContainerType**](CreateContainerType.md) | The Container Type create object | 

### Return type

[**ContainerType**](ContainerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerType

> DeleteContainerType(ctx, containerTypeId).Execute()

Deletes a Container Type



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
	containerTypeId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerTypeAPI.DeleteContainerType(context.Background(), containerTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerTypeAPI.DeleteContainerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerTypeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerTypeRequest struct via the builder pattern


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


## GetContainerType

> ContainerType GetContainerType(ctx, containerTypeId).Execute()

Get Container Type information



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
	containerTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerTypeAPI.GetContainerType(context.Background(), containerTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerTypeAPI.GetContainerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerType`: ContainerType
	fmt.Fprintf(os.Stdout, "Response from `ContainerTypeAPI.GetContainerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerType**](ContainerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerTypes

> ContainerTypePaginatedList GetContainerTypes(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDisplayName(filterDisplayName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Container Types



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
	filterDisplayName := []string{"Inner_example"} // []string | Filter by displayName query param.  **Format:** filter.displayName={$not}:OPERATION:VALUE    **Example:** filter.displayName=$btw:John Doe&filter.displayName=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  - displayName  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label,name,displayName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  - name  - displayName  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerTypeAPI.GetContainerTypes(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDisplayName(filterDisplayName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerTypeAPI.GetContainerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerTypes`: ContainerTypePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ContainerTypeAPI.GetContainerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$btw:John Doe&amp;filter.label&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDisplayName** | **[]string** | Filter by displayName query param.  **Format:** filter.displayName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.displayName&#x3D;$btw:John Doe&amp;filter.displayName&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  - displayName  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** label,name,displayName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  - name  - displayName  | 

### Return type

[**ContainerTypePaginatedList**](ContainerTypePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainersByContainerType

> ContainerPaginatedList GetContainersByContainerType(ctx, containerTypeId).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterName(filterName).FilterAddress(filterAddress).FilterHost(filterHost).FilterHosts(filterHosts).FilterTypeId(filterTypeId).FilterPoolId(filterPoolId).FilterAdministrationState(filterAdministrationState).FilterNumaNodes(filterNumaNodes).FilterInfrastructureId(filterInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Returns all Containers linked to the Container Type



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
	containerTypeId := int64(789) // int64 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAddress := []string{"Inner_example"} // []string | Filter by address query param.  **Format:** filter.address={$not}:OPERATION:VALUE    **Example:** filter.address=$btw:John Doe&filter.address=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHost := []string{"Inner_example"} // []string | Filter by host query param.  **Format:** filter.host={$not}:OPERATION:VALUE    **Example:** filter.host=$btw:John Doe&filter.host=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHosts := []string{"Inner_example"} // []string | Filter by hosts query param.  **Format:** filter.hosts={$not}:OPERATION:VALUE    **Example:** filter.hosts=$btw:John Doe&filter.hosts=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterTypeId := []string{"Inner_example"} // []string | Filter by typeId query param.  **Format:** filter.typeId={$not}:OPERATION:VALUE    **Example:** filter.typeId=$btw:John Doe&filter.typeId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPoolId := []string{"Inner_example"} // []string | Filter by poolId query param.  **Format:** filter.poolId={$not}:OPERATION:VALUE    **Example:** filter.poolId=$btw:John Doe&filter.poolId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAdministrationState := []string{"Inner_example"} // []string | Filter by administrationState query param.  **Format:** filter.administrationState={$not}:OPERATION:VALUE    **Example:** filter.administrationState=$btw:John Doe&filter.administrationState=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNumaNodes := []string{"Inner_example"} // []string | Filter by numaNodes query param.  **Format:** filter.numaNodes={$not}:OPERATION:VALUE    **Example:** filter.numaNodes=$btw:John Doe&filter.numaNodes=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$btw:John Doe&filter.infrastructureId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - host  - administrationState  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,host,hosts,administrationState   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - host  - hosts  - administrationState  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerTypeAPI.GetContainersByContainerType(context.Background(), containerTypeId).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterName(filterName).FilterAddress(filterAddress).FilterHost(filterHost).FilterHosts(filterHosts).FilterTypeId(filterTypeId).FilterPoolId(filterPoolId).FilterAdministrationState(filterAdministrationState).FilterNumaNodes(filterNumaNodes).FilterInfrastructureId(filterInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerTypeAPI.GetContainersByContainerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainersByContainerType`: ContainerPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ContainerTypeAPI.GetContainersByContainerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerTypeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersByContainerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAddress** | **[]string** | Filter by address query param.  **Format:** filter.address&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.address&#x3D;$btw:John Doe&amp;filter.address&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHost** | **[]string** | Filter by host query param.  **Format:** filter.host&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.host&#x3D;$btw:John Doe&amp;filter.host&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHosts** | **[]string** | Filter by hosts query param.  **Format:** filter.hosts&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.hosts&#x3D;$btw:John Doe&amp;filter.hosts&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterTypeId** | **[]string** | Filter by typeId query param.  **Format:** filter.typeId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.typeId&#x3D;$btw:John Doe&amp;filter.typeId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPoolId** | **[]string** | Filter by poolId query param.  **Format:** filter.poolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.poolId&#x3D;$btw:John Doe&amp;filter.poolId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAdministrationState** | **[]string** | Filter by administrationState query param.  **Format:** filter.administrationState&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.administrationState&#x3D;$btw:John Doe&amp;filter.administrationState&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNumaNodes** | **[]string** | Filter by numaNodes query param.  **Format:** filter.numaNodes&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.numaNodes&#x3D;$btw:John Doe&amp;filter.numaNodes&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$btw:John Doe&amp;filter.infrastructureId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - host  - administrationState  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,host,hosts,administrationState   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - host  - hosts  - administrationState  | 

### Return type

[**ContainerPaginatedList**](ContainerPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerType

> ContainerType UpdateContainerType(ctx, containerTypeId).UpdateContainerType(updateContainerType).Execute()

Updates Container Type information



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
	containerTypeId := int64(789) // int64 | 
	updateContainerType := *openapiclient.NewUpdateContainerType() // UpdateContainerType | The Container Type update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerTypeAPI.UpdateContainerType(context.Background(), containerTypeId).UpdateContainerType(updateContainerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerTypeAPI.UpdateContainerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContainerType`: ContainerType
	fmt.Fprintf(os.Stdout, "Response from `ContainerTypeAPI.UpdateContainerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerTypeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContainerType** | [**UpdateContainerType**](UpdateContainerType.md) | The Container Type update object | 

### Return type

[**ContainerType**](ContainerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

