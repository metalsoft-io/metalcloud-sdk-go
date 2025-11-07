# \VMTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVMType**](VMTypeAPI.md#CreateVMType) | **Post** /api/v2/vm-types | Creates a VM Type
[**DeleteVMType**](VMTypeAPI.md#DeleteVMType) | **Delete** /api/v2/vm-types/{vmTypeId} | Deletes a VM Type
[**GetVMType**](VMTypeAPI.md#GetVMType) | **Get** /api/v2/vm-types/{vmTypeId} | Get VM Type information
[**GetVMTypes**](VMTypeAPI.md#GetVMTypes) | **Get** /api/v2/vm-types | Get all VM Types
[**GetVMsByVMType**](VMTypeAPI.md#GetVMsByVMType) | **Get** /api/v2/vm-types/{vmTypeId}/vms | Returns all VMs linked to the VM Type
[**UpdateVMType**](VMTypeAPI.md#UpdateVMType) | **Patch** /api/v2/vm-types/{vmTypeId} | Updates VM Type information



## CreateVMType

> VMType CreateVMType(ctx).CreateVMType(createVMType).Execute()

Creates a VM Type



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
	createVMType := *openapiclient.NewCreateVMType("Name_example", float32(2), float32(2)) // CreateVMType | The VM Type create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMTypeAPI.CreateVMType(context.Background()).CreateVMType(createVMType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMTypeAPI.CreateVMType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVMType`: VMType
	fmt.Fprintf(os.Stdout, "Response from `VMTypeAPI.CreateVMType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVMType** | [**CreateVMType**](CreateVMType.md) | The VM Type create object | 

### Return type

[**VMType**](VMType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVMType

> DeleteVMType(ctx, vmTypeId).Execute()

Deletes a VM Type



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
	vmTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMTypeAPI.DeleteVMType(context.Background(), vmTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMTypeAPI.DeleteVMType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMTypeRequest struct via the builder pattern


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


## GetVMType

> VMType GetVMType(ctx, vmTypeId).Execute()

Get VM Type information



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
	vmTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMTypeAPI.GetVMType(context.Background(), vmTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMTypeAPI.GetVMType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMType`: VMType
	fmt.Fprintf(os.Stdout, "Response from `VMTypeAPI.GetVMType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VMType**](VMType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMTypes

> VMTypePaginatedList GetVMTypes(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDisplayName(filterDisplayName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all VM Types



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
	resp, r, err := apiClient.VMTypeAPI.GetVMTypes(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDisplayName(filterDisplayName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMTypeAPI.GetVMTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMTypes`: VMTypePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMTypeAPI.GetVMTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVMTypesRequest struct via the builder pattern


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

[**VMTypePaginatedList**](VMTypePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMsByVMType

> VMPaginatedList GetVMsByVMType(ctx, vmTypeId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Returns all VMs linked to the VM Type



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
	vmTypeId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAddress := []string{"Inner_example"} // []string | Filter by address query param.  **Format:** filter.address={$not}:OPERATION:VALUE    **Example:** filter.address=$btw:John Doe&filter.address=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMTypeAPI.GetVMsByVMType(context.Background(), vmTypeId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterAddress(filterAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMTypeAPI.GetVMsByVMType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMsByVMType`: VMPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMTypeAPI.GetVMsByVMType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMsByVMTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAddress** | **[]string** | Filter by address query param.  **Format:** filter.address&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.address&#x3D;$btw:John Doe&amp;filter.address&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 

### Return type

[**VMPaginatedList**](VMPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMType

> VMType UpdateVMType(ctx, vmTypeId).UpdateVMType(updateVMType).Execute()

Updates VM Type information



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
	vmTypeId := float32(8.14) // float32 | 
	updateVMType := *openapiclient.NewUpdateVMType() // UpdateVMType | The VM Type update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMTypeAPI.UpdateVMType(context.Background(), vmTypeId).UpdateVMType(updateVMType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMTypeAPI.UpdateVMType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMType`: VMType
	fmt.Fprintf(os.Stdout, "Response from `VMTypeAPI.UpdateVMType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVMType** | [**UpdateVMType**](UpdateVMType.md) | The VM Type update object | 

### Return type

[**VMType**](VMType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

