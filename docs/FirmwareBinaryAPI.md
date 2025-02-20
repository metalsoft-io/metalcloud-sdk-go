# \FirmwareBinaryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirmwareBinary**](FirmwareBinaryAPI.md#CreateFirmwareBinary) | **Post** /api/v2/firmware/binary | Create a new firmware binary
[**DeleteFirmwareBinary**](FirmwareBinaryAPI.md#DeleteFirmwareBinary) | **Delete** /api/v2/firmware/binary/{firmwareBinaryId} | Delete Firmware Binary
[**GetFirmwareBinaries**](FirmwareBinaryAPI.md#GetFirmwareBinaries) | **Get** /api/v2/firmware/binary | Get Firmware Binaries
[**GetFirmwareBinary**](FirmwareBinaryAPI.md#GetFirmwareBinary) | **Get** /api/v2/firmware/binary/{firmwareBinaryId} | Get Firmware Binary
[**UpdateFirmwareBinary**](FirmwareBinaryAPI.md#UpdateFirmwareBinary) | **Put** /api/v2/firmware/binary/{firmwareBinaryId} | Update Firmware Binary



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
	createFirmwareBinary := *openapiclient.NewCreateFirmwareBinary(float32(46), "https://downloads.dell.com/xxxxx", "Test", true, openapiclient.FirmwareBinaryUpdateSeverity("critical"), "[{"id": "123121", "model": "R240"}, {"id": "123122", "model": "R740"}]", "[{"id": "PowerEdge"}, {"id": "ThinkSystem"}, {"id": "ProLiant"}]") // CreateFirmwareBinary | 

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

> FirmwareBinaryPaginatedList GetFirmwareBinaries(ctx).Page(page).Limit(limit).FilterServerFirmwareBinaryId(filterServerFirmwareBinaryId).FilterServerFirmwareBinaryCatalogId(filterServerFirmwareBinaryCatalogId).FilterServerFirmwareBinaryExternalId(filterServerFirmwareBinaryExternalId).FilterServerFirmwareBinaryPackageId(filterServerFirmwareBinaryPackageId).FilterServerFirmwareBinaryCreatedTimestamp(filterServerFirmwareBinaryCreatedTimestamp).FilterServerFirmwareBinaryUpdateSeverity(filterServerFirmwareBinaryUpdateSeverity).FilterServerFirmwareBinaryRebootRequired(filterServerFirmwareBinaryRebootRequired).FilterServerFirmwareBinaryVendorReleaseTimestamp(filterServerFirmwareBinaryVendorReleaseTimestamp).FilterServerFirmwareBinaryPackageVersion(filterServerFirmwareBinaryPackageVersion).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

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
	filterServerFirmwareBinaryId := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryId query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryId=$not:$like:John Doe&filter.serverFirmwareBinaryId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryCatalogId := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryCatalogId query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryCatalogId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryCatalogId=$not:$like:John Doe&filter.serverFirmwareBinaryCatalogId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryExternalId := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryExternalId query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryExternalId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryExternalId=$not:$like:John Doe&filter.serverFirmwareBinaryExternalId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryPackageId := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryPackageId query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryPackageId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryPackageId=$not:$like:John Doe&filter.serverFirmwareBinaryPackageId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryCreatedTimestamp := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryCreatedTimestamp query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryCreatedTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryCreatedTimestamp=$not:$like:John Doe&filter.serverFirmwareBinaryCreatedTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryUpdateSeverity := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryUpdateSeverity query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryUpdateSeverity={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryUpdateSeverity=$not:$like:John Doe&filter.serverFirmwareBinaryUpdateSeverity=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryRebootRequired := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryRebootRequired query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryRebootRequired={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryRebootRequired=$not:$like:John Doe&filter.serverFirmwareBinaryRebootRequired=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryVendorReleaseTimestamp := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryVendorReleaseTimestamp query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryVendorReleaseTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryVendorReleaseTimestamp=$not:$like:John Doe&filter.serverFirmwareBinaryVendorReleaseTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerFirmwareBinaryPackageVersion := []string{"Inner_example"} // []string | Filter by serverFirmwareBinaryPackageVersion query param.           <p>              <b>Format: </b> filter.serverFirmwareBinaryPackageVersion={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverFirmwareBinaryPackageVersion=$not:$like:John Doe&filter.serverFirmwareBinaryPackageVersion=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> serverFirmwareBinaryId:DESC           </p>       <h4>Available Fields</h4><ul><li>serverFirmwareBinaryId</li> <li>serverFirmwareBinaryName</li> <li>serverFirmwareBinaryUpdateSeverity</li> <li>serverFirmwareBinaryCatalogId</li> <li>serverFirmwareBinaryRebootRequired</li> <li>serverFirmwareBinaryVendorReleaseTimestamp</li> <li>serverFirmwareBinaryPackageVersion</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> serverFirmwareBinaryName           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>serverFirmwareBinaryName</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> serverFirmwareBinaryId,serverFirmwareBinaryCatalogId,serverFirmwareBinaryExternalId,serverFirmwareBinaryPackageId,serverFirmwareBinaryPackageVersion           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBinaryAPI.GetFirmwareBinaries(context.Background()).Page(page).Limit(limit).FilterServerFirmwareBinaryId(filterServerFirmwareBinaryId).FilterServerFirmwareBinaryCatalogId(filterServerFirmwareBinaryCatalogId).FilterServerFirmwareBinaryExternalId(filterServerFirmwareBinaryExternalId).FilterServerFirmwareBinaryPackageId(filterServerFirmwareBinaryPackageId).FilterServerFirmwareBinaryCreatedTimestamp(filterServerFirmwareBinaryCreatedTimestamp).FilterServerFirmwareBinaryUpdateSeverity(filterServerFirmwareBinaryUpdateSeverity).FilterServerFirmwareBinaryRebootRequired(filterServerFirmwareBinaryRebootRequired).FilterServerFirmwareBinaryVendorReleaseTimestamp(filterServerFirmwareBinaryVendorReleaseTimestamp).FilterServerFirmwareBinaryPackageVersion(filterServerFirmwareBinaryPackageVersion).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
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
 **filterServerFirmwareBinaryId** | **[]string** | Filter by serverFirmwareBinaryId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryCatalogId** | **[]string** | Filter by serverFirmwareBinaryCatalogId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryCatalogId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryCatalogId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryCatalogId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryExternalId** | **[]string** | Filter by serverFirmwareBinaryExternalId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryExternalId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryExternalId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryExternalId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryPackageId** | **[]string** | Filter by serverFirmwareBinaryPackageId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryPackageId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryPackageId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryPackageId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryCreatedTimestamp** | **[]string** | Filter by serverFirmwareBinaryCreatedTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryCreatedTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryCreatedTimestamp&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryCreatedTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryUpdateSeverity** | **[]string** | Filter by serverFirmwareBinaryUpdateSeverity query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryUpdateSeverity&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryUpdateSeverity&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryUpdateSeverity&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryRebootRequired** | **[]string** | Filter by serverFirmwareBinaryRebootRequired query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryRebootRequired&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryRebootRequired&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryRebootRequired&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryVendorReleaseTimestamp** | **[]string** | Filter by serverFirmwareBinaryVendorReleaseTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryVendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryVendorReleaseTimestamp&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryVendorReleaseTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerFirmwareBinaryPackageVersion** | **[]string** | Filter by serverFirmwareBinaryPackageVersion query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareBinaryPackageVersion&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareBinaryPackageVersion&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareBinaryPackageVersion&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; serverFirmwareBinaryId:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;serverFirmwareBinaryId&lt;/li&gt; &lt;li&gt;serverFirmwareBinaryName&lt;/li&gt; &lt;li&gt;serverFirmwareBinaryUpdateSeverity&lt;/li&gt; &lt;li&gt;serverFirmwareBinaryCatalogId&lt;/li&gt; &lt;li&gt;serverFirmwareBinaryRebootRequired&lt;/li&gt; &lt;li&gt;serverFirmwareBinaryVendorReleaseTimestamp&lt;/li&gt; &lt;li&gt;serverFirmwareBinaryPackageVersion&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; serverFirmwareBinaryName           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;serverFirmwareBinaryName&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; serverFirmwareBinaryId,serverFirmwareBinaryCatalogId,serverFirmwareBinaryExternalId,serverFirmwareBinaryPackageId,serverFirmwareBinaryPackageVersion           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

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


## UpdateFirmwareBinary

> FirmwareBinary UpdateFirmwareBinary(ctx, firmwareBinaryId).UpdateFirmwareBinary(updateFirmwareBinary).Execute()

Update Firmware Binary



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
	updateFirmwareBinary := *openapiclient.NewUpdateFirmwareBinary(float32(46), "https://downloads.dell.com/xxxxx", "Test", true, openapiclient.FirmwareBinaryUpdateSeverity("critical"), "[{"id": "123121", "model": "R240"}, {"id": "123122", "model": "R740"}]", "[{"id": "PowerEdge"}, {"id": "ThinkSystem"}, {"id": "ProLiant"}]") // UpdateFirmwareBinary | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBinaryAPI.UpdateFirmwareBinary(context.Background(), firmwareBinaryId).UpdateFirmwareBinary(updateFirmwareBinary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBinaryAPI.UpdateFirmwareBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirmwareBinary`: FirmwareBinary
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBinaryAPI.UpdateFirmwareBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBinaryId** | **float32** | The firmware binary id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirmwareBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirmwareBinary** | [**UpdateFirmwareBinary**](UpdateFirmwareBinary.md) |  | 

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

