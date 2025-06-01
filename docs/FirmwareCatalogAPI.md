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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.           <p>              <b>Format: </b> filter.description={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.description=$not:$like:John Doe&filter.description=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterVendor := []string{"Inner_example"} // []string | Filter by vendor query param.           <p>              <b>Format: </b> filter.vendor={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.vendor=$not:$like:John Doe&filter.vendor=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterVendorId := []string{"Inner_example"} // []string | Filter by vendorId query param.           <p>              <b>Format: </b> filter.vendorId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.vendorId=$not:$like:John Doe&filter.vendorId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterVendorUrl := []string{"Inner_example"} // []string | Filter by vendorUrl query param.           <p>              <b>Format: </b> filter.vendorUrl={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.vendorUrl=$not:$like:John Doe&filter.vendorUrl=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterVendorReleaseTimestamp := []string{"Inner_example"} // []string | Filter by vendorReleaseTimestamp query param.           <p>              <b>Format: </b> filter.vendorReleaseTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.vendorReleaseTimestamp=$not:$like:John Doe&filter.vendorReleaseTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterUpdateType := []string{"Inner_example"} // []string | Filter by updateType query param.           <p>              <b>Format: </b> filter.updateType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.updateType=$not:$like:John Doe&filter.updateType=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.           <p>              <b>Format: </b> filter.createdTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.createdTimestamp=$not:$like:John Doe&filter.createdTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>name</li> <li>vendor</li> <li>vendorReleaseTimestamp</li> <li>updateType</li> <li>createdTimestamp</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> name,description           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>name</li> <li>description</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> id,name,description,vendor,vendorId           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterDescription** | **[]string** | Filter by description query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.description&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.description&#x3D;$not:$like:John Doe&amp;filter.description&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterVendor** | **[]string** | Filter by vendor query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.vendor&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.vendor&#x3D;$not:$like:John Doe&amp;filter.vendor&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterVendorId** | **[]string** | Filter by vendorId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.vendorId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.vendorId&#x3D;$not:$like:John Doe&amp;filter.vendorId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterVendorUrl** | **[]string** | Filter by vendorUrl query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.vendorUrl&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.vendorUrl&#x3D;$not:$like:John Doe&amp;filter.vendorUrl&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterVendorReleaseTimestamp** | **[]string** | Filter by vendorReleaseTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.vendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.vendorReleaseTimestamp&#x3D;$not:$like:John Doe&amp;filter.vendorReleaseTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterUpdateType** | **[]string** | Filter by updateType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.updateType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.updateType&#x3D;$not:$like:John Doe&amp;filter.updateType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.createdTimestamp&#x3D;$not:$like:John Doe&amp;filter.createdTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;vendor&lt;/li&gt; &lt;li&gt;vendorReleaseTimestamp&lt;/li&gt; &lt;li&gt;updateType&lt;/li&gt; &lt;li&gt;createdTimestamp&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name,description           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt; &lt;li&gt;description&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,name,description,vendor,vendorId           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

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

