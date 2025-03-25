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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterCatalogId := []string{"Inner_example"} // []string | Filter by catalogId query param.           <p>              <b>Format: </b> filter.catalogId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.catalogId=$not:$like:John Doe&filter.catalogId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterExternalId := []string{"Inner_example"} // []string | Filter by externalId query param.           <p>              <b>Format: </b> filter.externalId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.externalId=$not:$like:John Doe&filter.externalId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterPackageId := []string{"Inner_example"} // []string | Filter by packageId query param.           <p>              <b>Format: </b> filter.packageId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.packageId=$not:$like:John Doe&filter.packageId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.           <p>              <b>Format: </b> filter.createdTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.createdTimestamp=$not:$like:John Doe&filter.createdTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterUpdateSeverity := []string{"Inner_example"} // []string | Filter by updateSeverity query param.           <p>              <b>Format: </b> filter.updateSeverity={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.updateSeverity=$not:$like:John Doe&filter.updateSeverity=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterRebootRequired := []string{"Inner_example"} // []string | Filter by rebootRequired query param.           <p>              <b>Format: </b> filter.rebootRequired={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.rebootRequired=$not:$like:John Doe&filter.rebootRequired=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterVendorReleaseTimestamp := []string{"Inner_example"} // []string | Filter by vendorReleaseTimestamp query param.           <p>              <b>Format: </b> filter.vendorReleaseTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.vendorReleaseTimestamp=$not:$like:John Doe&filter.vendorReleaseTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterPackageVersion := []string{"Inner_example"} // []string | Filter by packageVersion query param.           <p>              <b>Format: </b> filter.packageVersion={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.packageVersion=$not:$like:John Doe&filter.packageVersion=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>name</li> <li>updateSeverity</li> <li>catalogId</li> <li>rebootRequired</li> <li>vendorReleaseTimestamp</li> <li>packageVersion</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> name           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>name</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> id,catalogId,externalId,packageId,packageVersion           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterCatalogId** | **[]string** | Filter by catalogId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.catalogId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.catalogId&#x3D;$not:$like:John Doe&amp;filter.catalogId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterExternalId** | **[]string** | Filter by externalId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.externalId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.externalId&#x3D;$not:$like:John Doe&amp;filter.externalId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterPackageId** | **[]string** | Filter by packageId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.packageId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.packageId&#x3D;$not:$like:John Doe&amp;filter.packageId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.createdTimestamp&#x3D;$not:$like:John Doe&amp;filter.createdTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterUpdateSeverity** | **[]string** | Filter by updateSeverity query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.updateSeverity&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.updateSeverity&#x3D;$not:$like:John Doe&amp;filter.updateSeverity&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterRebootRequired** | **[]string** | Filter by rebootRequired query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.rebootRequired&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.rebootRequired&#x3D;$not:$like:John Doe&amp;filter.rebootRequired&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterVendorReleaseTimestamp** | **[]string** | Filter by vendorReleaseTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.vendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.vendorReleaseTimestamp&#x3D;$not:$like:John Doe&amp;filter.vendorReleaseTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterPackageVersion** | **[]string** | Filter by packageVersion query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.packageVersion&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.packageVersion&#x3D;$not:$like:John Doe&amp;filter.packageVersion&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;updateSeverity&lt;/li&gt; &lt;li&gt;catalogId&lt;/li&gt; &lt;li&gt;rebootRequired&lt;/li&gt; &lt;li&gt;vendorReleaseTimestamp&lt;/li&gt; &lt;li&gt;packageVersion&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,catalogId,externalId,packageId,packageVersion           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

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

