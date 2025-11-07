# \FirmwareCatalogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirmwareCatalogs**](FirmwareCatalogAPI.md#CreateFirmwareCatalogs) | **Post** /api/v2/firmware/catalog | Create Firmware Catalog
[**DeleteFirmwareCatalog**](FirmwareCatalogAPI.md#DeleteFirmwareCatalog) | **Delete** /api/v2/firmware/catalog/{firmwareCatalogId} | Delete Firmware Catalog
[**GetFirmwareCatalog**](FirmwareCatalogAPI.md#GetFirmwareCatalog) | **Get** /api/v2/firmware/catalog/{firmwareCatalogId} | Get Firmware Catalog
[**GetFirmwareCatalogs**](FirmwareCatalogAPI.md#GetFirmwareCatalogs) | **Get** /api/v2/firmware/catalog | Get Firmware Catalogs
[**UpdateFirmwareCatalog**](FirmwareCatalogAPI.md#UpdateFirmwareCatalog) | **Put** /api/v2/firmware/catalog/{firmwareCatalogId} | Update Firmware Catalog



## CreateFirmwareCatalogs

> FirmwareCatalog CreateFirmwareCatalogs(ctx).CreateFirmwareCatalog(createFirmwareCatalog).Execute()

Create Firmware Catalog



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
	createFirmwareCatalog := *openapiclient.NewCreateFirmwareCatalog("Dell PowerEdge R740", openapiclient.ServerFirmwareCatalogVendor("dell"), openapiclient.CatalogUpdateType("online")) // CreateFirmwareCatalog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareCatalogAPI.CreateFirmwareCatalogs(context.Background()).CreateFirmwareCatalog(createFirmwareCatalog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareCatalogAPI.CreateFirmwareCatalogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirmwareCatalogs`: FirmwareCatalog
	fmt.Fprintf(os.Stdout, "Response from `FirmwareCatalogAPI.CreateFirmwareCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirmwareCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirmwareCatalog** | [**CreateFirmwareCatalog**](CreateFirmwareCatalog.md) |  | 

### Return type

[**FirmwareCatalog**](FirmwareCatalog.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirmwareCatalog

> DeleteFirmwareCatalog(ctx, firmwareCatalogId).Execute()

Delete Firmware Catalog



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
	firmwareCatalogId := float32(8.14) // float32 | The firmware catalog id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirmwareCatalogAPI.DeleteFirmwareCatalog(context.Background(), firmwareCatalogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareCatalogAPI.DeleteFirmwareCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareCatalogId** | **float32** | The firmware catalog id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirmwareCatalogRequest struct via the builder pattern


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


## GetFirmwareCatalog

> FirmwareCatalog GetFirmwareCatalog(ctx, firmwareCatalogId).Execute()

Get Firmware Catalog



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
	firmwareCatalogId := float32(8.14) // float32 | The firmware catalog id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareCatalogAPI.GetFirmwareCatalog(context.Background(), firmwareCatalogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareCatalogAPI.GetFirmwareCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareCatalog`: FirmwareCatalog
	fmt.Fprintf(os.Stdout, "Response from `FirmwareCatalogAPI.GetFirmwareCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareCatalogId** | **float32** | The firmware catalog id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirmwareCatalog**](FirmwareCatalog.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirmwareCatalogs

> FirmwareCatalogPaginatedList GetFirmwareCatalogs(ctx).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterVendor(filterVendor).FilterVendorId(filterVendorId).FilterVendorUrl(filterVendorUrl).FilterVendorReleaseTimestamp(filterVendorReleaseTimestamp).FilterUpdateType(filterUpdateType).FilterCreatedTimestamp(filterCreatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get Firmware Catalogs



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
	filterVendor := []string{"Inner_example"} // []string | Filter by vendor query param.  **Format:** filter.vendor={$not}:OPERATION:VALUE    **Example:** filter.vendor=$btw:John Doe&filter.vendor=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterVendorId := []string{"Inner_example"} // []string | Filter by vendorId query param.  **Format:** filter.vendorId={$not}:OPERATION:VALUE    **Example:** filter.vendorId=$btw:John Doe&filter.vendorId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterVendorUrl := []string{"Inner_example"} // []string | Filter by vendorUrl query param.  **Format:** filter.vendorUrl={$not}:OPERATION:VALUE    **Example:** filter.vendorUrl=$btw:John Doe&filter.vendorUrl=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterVendorReleaseTimestamp := []string{"Inner_example"} // []string | Filter by vendorReleaseTimestamp query param.  **Format:** filter.vendorReleaseTimestamp={$not}:OPERATION:VALUE    **Example:** filter.vendorReleaseTimestamp=$btw:John Doe&filter.vendorReleaseTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdateType := []string{"Inner_example"} // []string | Filter by updateType query param.  **Format:** filter.updateType={$not}:OPERATION:VALUE    **Example:** filter.updateType=$btw:John Doe&filter.updateType=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp={$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp=$btw:John Doe&filter.createdTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - vendor  - vendorReleaseTimestamp  - updateType  - createdTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - description  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,name,description,vendor,vendorId   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareCatalogAPI.GetFirmwareCatalogs(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDescription(filterDescription).FilterVendor(filterVendor).FilterVendorId(filterVendorId).FilterVendorUrl(filterVendorUrl).FilterVendorReleaseTimestamp(filterVendorReleaseTimestamp).FilterUpdateType(filterUpdateType).FilterCreatedTimestamp(filterCreatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareCatalogAPI.GetFirmwareCatalogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareCatalogs`: FirmwareCatalogPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `FirmwareCatalogAPI.GetFirmwareCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDescription** | **[]string** | Filter by description query param.  **Format:** filter.description&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.description&#x3D;$btw:John Doe&amp;filter.description&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVendor** | **[]string** | Filter by vendor query param.  **Format:** filter.vendor&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vendor&#x3D;$btw:John Doe&amp;filter.vendor&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVendorId** | **[]string** | Filter by vendorId query param.  **Format:** filter.vendorId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vendorId&#x3D;$btw:John Doe&amp;filter.vendorId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVendorUrl** | **[]string** | Filter by vendorUrl query param.  **Format:** filter.vendorUrl&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vendorUrl&#x3D;$btw:John Doe&amp;filter.vendorUrl&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVendorReleaseTimestamp** | **[]string** | Filter by vendorReleaseTimestamp query param.  **Format:** filter.vendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vendorReleaseTimestamp&#x3D;$btw:John Doe&amp;filter.vendorReleaseTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdateType** | **[]string** | Filter by updateType query param.  **Format:** filter.updateType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updateType&#x3D;$btw:John Doe&amp;filter.updateType&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp&#x3D;$btw:John Doe&amp;filter.createdTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - vendor  - vendorReleaseTimestamp  - updateType  - createdTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - description  | 
 **select_** | **string** | List of fields to select.  **Example:** id,name,description,vendor,vendorId   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**FirmwareCatalogPaginatedList**](FirmwareCatalogPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirmwareCatalog

> FirmwareCatalog UpdateFirmwareCatalog(ctx, firmwareCatalogId).UpdateFirmwareCatalog(updateFirmwareCatalog).Execute()

Update Firmware Catalog



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
	firmwareCatalogId := float32(8.14) // float32 | The firmware catalog id
	updateFirmwareCatalog := *openapiclient.NewUpdateFirmwareCatalog("Dell PowerEdge R740", openapiclient.ServerFirmwareCatalogVendor("dell"), openapiclient.CatalogUpdateType("online")) // UpdateFirmwareCatalog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareCatalogAPI.UpdateFirmwareCatalog(context.Background(), firmwareCatalogId).UpdateFirmwareCatalog(updateFirmwareCatalog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareCatalogAPI.UpdateFirmwareCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirmwareCatalog`: FirmwareCatalog
	fmt.Fprintf(os.Stdout, "Response from `FirmwareCatalogAPI.UpdateFirmwareCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareCatalogId** | **float32** | The firmware catalog id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirmwareCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirmwareCatalog** | [**UpdateFirmwareCatalog**](UpdateFirmwareCatalog.md) |  | 

### Return type

[**FirmwareCatalog**](FirmwareCatalog.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

