# \FirmwareBinaryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirmwareBinary**](FirmwareBinaryAPI.md#CreateFirmwareBinary) | **Post** /api/v2/firmware/binary | Create a new firmware binary
[**DeleteFirmwareBinary**](FirmwareBinaryAPI.md#DeleteFirmwareBinary) | **Delete** /api/v2/firmware/binary/{firmwareBinaryId} | Delete Firmware Binary
[**GetFirmwareBinaries**](FirmwareBinaryAPI.md#GetFirmwareBinaries) | **Get** /api/v2/firmware/binary | Get Firmware Binaries
[**GetFirmwareBinary**](FirmwareBinaryAPI.md#GetFirmwareBinary) | **Get** /api/v2/firmware/binary/{firmwareBinaryId} | Get Firmware Binary



## CreateFirmwareBinary

> FirmwareBinary CreateFirmwareBinary(ctx).CreateFirmwareBinary(createFirmwareBinary).Execute()

Create a new firmware binary



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
	createFirmwareBinary := *openapiclient.NewCreateFirmwareBinary(float32(46), "https://downloads.dell.com/xxxxx", "Test", true, openapiclient.FirmwareBinaryUpdateSeverity("critical"), []map[string]interface{}{map[string]interface{}(123)}, []map[string]interface{}{map[string]interface{}(123)}) // CreateFirmwareBinary | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBinaryAPI.CreateFirmwareBinary(context.Background()).CreateFirmwareBinary(createFirmwareBinary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBinaryAPI.CreateFirmwareBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirmwareBinary`: FirmwareBinary
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBinaryAPI.CreateFirmwareBinary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirmwareBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirmwareBinary** | [**CreateFirmwareBinary**](CreateFirmwareBinary.md) |  | 

### Return type

[**FirmwareBinary**](FirmwareBinary.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirmwareBinary

> DeleteFirmwareBinary(ctx, firmwareBinaryId).Execute()

Delete Firmware Binary



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
	firmwareBinaryId := float32(8.14) // float32 | The firmware binary id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirmwareBinaryAPI.DeleteFirmwareBinary(context.Background(), firmwareBinaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBinaryAPI.DeleteFirmwareBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBinaryId** | **float32** | The firmware binary id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirmwareBinaryRequest struct via the builder pattern


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


## GetFirmwareBinaries

> FirmwareBinaryPaginatedList GetFirmwareBinaries(ctx).Page(page).Limit(limit).FilterId(filterId).FilterCatalogId(filterCatalogId).FilterExternalId(filterExternalId).FilterPackageId(filterPackageId).FilterCreatedTimestamp(filterCreatedTimestamp).FilterUpdateSeverity(filterUpdateSeverity).FilterRebootRequired(filterRebootRequired).FilterVendorReleaseTimestamp(filterVendorReleaseTimestamp).FilterPackageVersion(filterPackageVersion).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get Firmware Binaries



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
	filterCatalogId := []string{"Inner_example"} // []string | Filter by catalogId query param.  **Format:** filter.catalogId={$not}:OPERATION:VALUE    **Example:** filter.catalogId=$btw:John Doe&filter.catalogId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterExternalId := []string{"Inner_example"} // []string | Filter by externalId query param.  **Format:** filter.externalId={$not}:OPERATION:VALUE    **Example:** filter.externalId=$btw:John Doe&filter.externalId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPackageId := []string{"Inner_example"} // []string | Filter by packageId query param.  **Format:** filter.packageId={$not}:OPERATION:VALUE    **Example:** filter.packageId=$btw:John Doe&filter.packageId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp={$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp=$btw:John Doe&filter.createdTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdateSeverity := []string{"Inner_example"} // []string | Filter by updateSeverity query param.  **Format:** filter.updateSeverity={$not}:OPERATION:VALUE    **Example:** filter.updateSeverity=$btw:John Doe&filter.updateSeverity=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterRebootRequired := []string{"Inner_example"} // []string | Filter by rebootRequired query param.  **Format:** filter.rebootRequired={$not}:OPERATION:VALUE    **Example:** filter.rebootRequired=$btw:John Doe&filter.rebootRequired=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterVendorReleaseTimestamp := []string{"Inner_example"} // []string | Filter by vendorReleaseTimestamp query param.  **Format:** filter.vendorReleaseTimestamp={$not}:OPERATION:VALUE    **Example:** filter.vendorReleaseTimestamp=$btw:John Doe&filter.vendorReleaseTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPackageVersion := []string{"Inner_example"} // []string | Filter by packageVersion query param.  **Format:** filter.packageVersion={$not}:OPERATION:VALUE    **Example:** filter.packageVersion=$btw:John Doe&filter.packageVersion=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - updateSeverity  - catalogId  - rebootRequired  - vendorReleaseTimestamp  - packageVersion  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,catalogId,externalId,packageId,packageVersion   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBinaryAPI.GetFirmwareBinaries(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterCatalogId(filterCatalogId).FilterExternalId(filterExternalId).FilterPackageId(filterPackageId).FilterCreatedTimestamp(filterCreatedTimestamp).FilterUpdateSeverity(filterUpdateSeverity).FilterRebootRequired(filterRebootRequired).FilterVendorReleaseTimestamp(filterVendorReleaseTimestamp).FilterPackageVersion(filterPackageVersion).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBinaryAPI.GetFirmwareBinaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareBinaries`: FirmwareBinaryPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBinaryAPI.GetFirmwareBinaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareBinariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCatalogId** | **[]string** | Filter by catalogId query param.  **Format:** filter.catalogId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.catalogId&#x3D;$btw:John Doe&amp;filter.catalogId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterExternalId** | **[]string** | Filter by externalId query param.  **Format:** filter.externalId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.externalId&#x3D;$btw:John Doe&amp;filter.externalId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPackageId** | **[]string** | Filter by packageId query param.  **Format:** filter.packageId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.packageId&#x3D;$btw:John Doe&amp;filter.packageId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp&#x3D;$btw:John Doe&amp;filter.createdTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdateSeverity** | **[]string** | Filter by updateSeverity query param.  **Format:** filter.updateSeverity&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updateSeverity&#x3D;$btw:John Doe&amp;filter.updateSeverity&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterRebootRequired** | **[]string** | Filter by rebootRequired query param.  **Format:** filter.rebootRequired&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.rebootRequired&#x3D;$btw:John Doe&amp;filter.rebootRequired&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVendorReleaseTimestamp** | **[]string** | Filter by vendorReleaseTimestamp query param.  **Format:** filter.vendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vendorReleaseTimestamp&#x3D;$btw:John Doe&amp;filter.vendorReleaseTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPackageVersion** | **[]string** | Filter by packageVersion query param.  **Format:** filter.packageVersion&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.packageVersion&#x3D;$btw:John Doe&amp;filter.packageVersion&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - updateSeverity  - catalogId  - rebootRequired  - vendorReleaseTimestamp  - packageVersion  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 
 **select_** | **string** | List of fields to select.  **Example:** id,catalogId,externalId,packageId,packageVersion   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**FirmwareBinaryPaginatedList**](FirmwareBinaryPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirmwareBinary

> FirmwareBinary GetFirmwareBinary(ctx, firmwareBinaryId).Execute()

Get Firmware Binary



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
	firmwareBinaryId := float32(8.14) // float32 | The firmware binary id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBinaryAPI.GetFirmwareBinary(context.Background(), firmwareBinaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBinaryAPI.GetFirmwareBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareBinary`: FirmwareBinary
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBinaryAPI.GetFirmwareBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBinaryId** | **float32** | The firmware binary id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirmwareBinary**](FirmwareBinary.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

