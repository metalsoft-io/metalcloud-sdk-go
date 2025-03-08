# \FirmwareBaselineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirmwareBaseline**](FirmwareBaselineAPI.md#CreateFirmwareBaseline) | **Post** /api/v2/firmware/baseline | Create a new firmware baseline
[**DeleteFirmwareBaseline**](FirmwareBaselineAPI.md#DeleteFirmwareBaseline) | **Delete** /api/v2/firmware/baseline/{firmwareBaselineId} | Delete Firmware Baseline
[**GetFirmwareBaseline**](FirmwareBaselineAPI.md#GetFirmwareBaseline) | **Get** /api/v2/firmware/baseline/{firmwareBaselineId} | Get Firmware Baseline
[**GetFirmwareBaselines**](FirmwareBaselineAPI.md#GetFirmwareBaselines) | **Get** /api/v2/firmware/baseline | Get Firmware Baselines
[**UpdateFirmwareBaseline**](FirmwareBaselineAPI.md#UpdateFirmwareBaseline) | **Put** /api/v2/firmware/baseline/{firmwareBaselineId} | Update Firmware Baseline



## CreateFirmwareBaseline

> CreateFirmwareBaseline CreateFirmwareBaseline(ctx).CreateFirmwareBaseline(createFirmwareBaseline).Execute()

Create a new firmware baseline



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
	createFirmwareBaseline := *openapiclient.NewCreateFirmwareBaseline("Data center baseline name", openapiclient.BaselineLevelType("datacenter"), []string{"LevelFilter_example"}) // CreateFirmwareBaseline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.CreateFirmwareBaseline(context.Background()).CreateFirmwareBaseline(createFirmwareBaseline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.CreateFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirmwareBaseline`: CreateFirmwareBaseline
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.CreateFirmwareBaseline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirmwareBaselineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirmwareBaseline** | [**CreateFirmwareBaseline**](CreateFirmwareBaseline.md) |  | 

### Return type

[**CreateFirmwareBaseline**](CreateFirmwareBaseline.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirmwareBaseline

> DeleteFirmwareBaseline(ctx, firmwareBaselineId).Execute()

Delete Firmware Baseline



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
	firmwareBaselineId := float32(8.14) // float32 | The firmware baseline id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirmwareBaselineAPI.DeleteFirmwareBaseline(context.Background(), firmwareBaselineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.DeleteFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBaselineId** | **float32** | The firmware baseline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirmwareBaselineRequest struct via the builder pattern


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


## GetFirmwareBaseline

> FirmwareBaseline GetFirmwareBaseline(ctx, firmwareBaselineId).Execute()

Get Firmware Baseline



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
	firmwareBaselineId := float32(8.14) // float32 | The firmware baseline id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.GetFirmwareBaseline(context.Background(), firmwareBaselineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.GetFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareBaseline`: FirmwareBaseline
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.GetFirmwareBaseline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBaselineId** | **float32** | The firmware baseline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareBaselineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirmwareBaseline**](FirmwareBaseline.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirmwareBaselines

> FirmwareBaselinePaginatedList GetFirmwareBaselines(ctx).Page(page).Limit(limit).FilterLevel(filterLevel).FilterLevelFilter(filterLevelFilter).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get Firmware Baselines



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
	filterLevel := []string{"Inner_example"} // []string | Filter by level query param.           <p>              <b>Format: </b> filter.level={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.level=$not:$like:John Doe&filter.level=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterLevelFilter := []string{"Inner_example"} // []string | Filter by levelFilter query param.           <p>              <b>Format: </b> filter.levelFilter={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.levelFilter=$not:$like:John Doe&filter.levelFilter=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>name</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> name,levelFilter           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>name</li> <li>levelFilter</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> id,name,description,catalog,level           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.GetFirmwareBaselines(context.Background()).Page(page).Limit(limit).FilterLevel(filterLevel).FilterLevelFilter(filterLevelFilter).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.GetFirmwareBaselines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareBaselines`: FirmwareBaselinePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.GetFirmwareBaselines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareBaselinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLevel** | **[]string** | Filter by level query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.level&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.level&#x3D;$not:$like:John Doe&amp;filter.level&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterLevelFilter** | **[]string** | Filter by levelFilter query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.levelFilter&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.levelFilter&#x3D;$not:$like:John Doe&amp;filter.levelFilter&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name,levelFilter           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt; &lt;li&gt;levelFilter&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,name,description,catalog,level           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

### Return type

[**FirmwareBaselinePaginatedList**](FirmwareBaselinePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirmwareBaseline

> FirmwareBaseline UpdateFirmwareBaseline(ctx, firmwareBaselineId).UpdateFirmwareBaseline(updateFirmwareBaseline).Execute()

Update Firmware Baseline



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
	firmwareBaselineId := float32(8.14) // float32 | The firmware baseline id
	updateFirmwareBaseline := *openapiclient.NewUpdateFirmwareBaseline("Data center baseline name", openapiclient.BaselineLevelType("datacenter"), []string{"LevelFilter_example"}) // UpdateFirmwareBaseline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.UpdateFirmwareBaseline(context.Background(), firmwareBaselineId).UpdateFirmwareBaseline(updateFirmwareBaseline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.UpdateFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirmwareBaseline`: FirmwareBaseline
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.UpdateFirmwareBaseline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBaselineId** | **float32** | The firmware baseline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirmwareBaselineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirmwareBaseline** | [**UpdateFirmwareBaseline**](UpdateFirmwareBaseline.md) |  | 

### Return type

[**FirmwareBaseline**](FirmwareBaseline.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

