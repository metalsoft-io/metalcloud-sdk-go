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

> FirmwareCatalog CreateFirmwareCatalogs(ctx).CreateFirmwareCatalogDto(createFirmwareCatalogDto).Execute()

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
	createFirmwareCatalogDto := *openapiclient.NewCreateFirmwareCatalogDto("Dell PowerEdge R740", openapiclient.FirmwareVendorType("dell"), openapiclient.CatalogUpdateType("online")) // CreateFirmwareCatalogDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareCatalogAPI.CreateFirmwareCatalogs(context.Background()).CreateFirmwareCatalogDto(createFirmwareCatalogDto).Execute()
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
 **createFirmwareCatalogDto** | [**CreateFirmwareCatalogDto**](CreateFirmwareCatalogDto.md) |  | 

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

> FirmwareCatalogPaginatedList GetFirmwareCatalogs(ctx).Page(page).Limit(limit).FilterServerFirmwareCatalogId(filterServerFirmwareCatalogId).FilterServerFirmwareCatalogName(filterServerFirmwareCatalogName).FilterServerFirmwareCatalogDescription(filterServerFirmwareCatalogDescription).FilterServerFirmwareCatalogVendor(filterServerFirmwareCatalogVendor).FilterServerFirmwareCatalogVendorId(filterServerFirmwareCatalogVendorId).FilterServerFirmwareCatalogVendorUrl(filterServerFirmwareCatalogVendorUrl).FilterServerFirmwareCatalogVendorReleaseTimestamp(filterServerFirmwareCatalogVendorReleaseTimestamp).FilterServerFirmwareCatalogUpdateType(filterServerFirmwareCatalogUpdateType).FilterServerFirmwareCatalogCreatedTimestamp(filterServerFirmwareCatalogCreatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

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
	filterServerFirmwareCatalogId := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogId query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogId=$not:$like:John Doe&filter.serverFirmwareCatalogId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogName := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogName query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogName=$not:$like:John Doe&filter.serverFirmwareCatalogName=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogDescription := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogDescription query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogDescription={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogDescription=$not:$like:John Doe&filter.serverFirmwareCatalogDescription=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogVendor := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogVendor query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogVendor={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogVendor=$not:$like:John Doe&filter.serverFirmwareCatalogVendor=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogVendorId := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogVendorId query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogVendorId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogVendorId=$not:$like:John Doe&filter.serverFirmwareCatalogVendorId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogVendorUrl := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogVendorUrl query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogVendorUrl={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogVendorUrl=$not:$like:John Doe&filter.serverFirmwareCatalogVendorUrl=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogVendorReleaseTimestamp := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogVendorReleaseTimestamp query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogVendorReleaseTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogVendorReleaseTimestamp=$not:$like:John Doe&filter.serverFirmwareCatalogVendorReleaseTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogUpdateType := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogUpdateType query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogUpdateType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogUpdateType=$not:$like:John Doe&filter.serverFirmwareCatalogUpdateType=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareCatalogCreatedTimestamp := []string{"Inner_example"} // []string | Filter by serverFirmwareCatalogCreatedTimestamp query param.           <p>              <b>Format: </b> filter.serverFirmwareCatalogCreatedTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareCatalogCreatedTimestamp=$not:$like:John Doe&filter.serverFirmwareCatalogCreatedTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> serverFirmwareCatalogId:DESC           </p>       <h4>Available Fields</h4><ul><li>serverFirmwareCatalogId</li> <li>serverFirmwareCatalogName</li> <li>serverFirmwareCatalogVendor</li> <li>serverFirmwareCatalogVendorReleaseTimestamp</li> <li>serverFirmwareCatalogUpdateType</li> <li>serverFirmwareCatalogCreatedTimestamp</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> serverFirmwareCatalogName,serverFirmwareCatalogDescription           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>serverFirmwareCatalogName</li> <li>serverFirmwareCatalogDescription</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> serverFirmwareCatalogId,serverFirmwareCatalogName,serverFirmwareCatalogDescription,serverFirmwareCatalogVendor,serverFirmwareCatalogVendorId           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareCatalogAPI.GetFirmwareCatalogs(context.Background()).Page(page).Limit(limit).FilterServerFirmwareCatalogId(filterServerFirmwareCatalogId).FilterServerFirmwareCatalogName(filterServerFirmwareCatalogName).FilterServerFirmwareCatalogDescription(filterServerFirmwareCatalogDescription).FilterServerFirmwareCatalogVendor(filterServerFirmwareCatalogVendor).FilterServerFirmwareCatalogVendorId(filterServerFirmwareCatalogVendorId).FilterServerFirmwareCatalogVendorUrl(filterServerFirmwareCatalogVendorUrl).FilterServerFirmwareCatalogVendorReleaseTimestamp(filterServerFirmwareCatalogVendorReleaseTimestamp).FilterServerFirmwareCatalogUpdateType(filterServerFirmwareCatalogUpdateType).FilterServerFirmwareCatalogCreatedTimestamp(filterServerFirmwareCatalogCreatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
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
 **filterServerFirmwareCatalogId** | **[]string** | Filter by serverFirmwareCatalogId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogName** | **[]string** | Filter by serverFirmwareCatalogName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogName&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogDescription** | **[]string** | Filter by serverFirmwareCatalogDescription query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogDescription&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogDescription&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogDescription&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogVendor** | **[]string** | Filter by serverFirmwareCatalogVendor query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendor&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendor&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendor&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogVendorId** | **[]string** | Filter by serverFirmwareCatalogVendorId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendorId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendorId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendorId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogVendorUrl** | **[]string** | Filter by serverFirmwareCatalogVendorUrl query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendorUrl&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendorUrl&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendorUrl&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogVendorReleaseTimestamp** | **[]string** | Filter by serverFirmwareCatalogVendorReleaseTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendorReleaseTimestamp&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendorReleaseTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogUpdateType** | **[]string** | Filter by serverFirmwareCatalogUpdateType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogUpdateType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogUpdateType&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogUpdateType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareCatalogCreatedTimestamp** | **[]string** | Filter by serverFirmwareCatalogCreatedTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogCreatedTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogCreatedTimestamp&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogCreatedTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; serverFirmwareCatalogId:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;serverFirmwareCatalogId&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogName&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogVendor&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogVendorReleaseTimestamp&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogUpdateType&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogCreatedTimestamp&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; serverFirmwareCatalogName,serverFirmwareCatalogDescription           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;serverFirmwareCatalogName&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogDescription&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; serverFirmwareCatalogId,serverFirmwareCatalogName,serverFirmwareCatalogDescription,serverFirmwareCatalogVendor,serverFirmwareCatalogVendorId           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

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

> FirmwareCatalog UpdateFirmwareCatalog(ctx, firmwareCatalogId).UpdateFirmwareCatalogDto(updateFirmwareCatalogDto).Execute()

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
	updateFirmwareCatalogDto := *openapiclient.NewUpdateFirmwareCatalogDto("Dell PowerEdge R740", openapiclient.FirmwareVendorType("dell"), openapiclient.CatalogUpdateType("online")) // UpdateFirmwareCatalogDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareCatalogAPI.UpdateFirmwareCatalog(context.Background(), firmwareCatalogId).UpdateFirmwareCatalogDto(updateFirmwareCatalogDto).Execute()
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

 **updateFirmwareCatalogDto** | [**UpdateFirmwareCatalogDto**](UpdateFirmwareCatalogDto.md) |  | 

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

