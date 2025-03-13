# \TemplateAssetAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateAsset**](TemplateAssetAPI.md#CreateTemplateAsset) | **Post** /api/v2/template-assets | Create template asset
[**DeleteTemplateAsset**](TemplateAssetAPI.md#DeleteTemplateAsset) | **Delete** /api/v2/template-assets/{templateAssetId} | Delete template asset
[**GetTemplateAsset**](TemplateAssetAPI.md#GetTemplateAsset) | **Get** /api/v2/template-assets/{templateAssetId} | Get details for an template asset
[**GetTemplateAssets**](TemplateAssetAPI.md#GetTemplateAssets) | **Get** /api/v2/template-assets | Get a list of available template assets
[**UpdateTemplateAsset**](TemplateAssetAPI.md#UpdateTemplateAsset) | **Put** /api/v2/template-assets/{templateAssetId} | Update template asset



## CreateTemplateAsset

> TemplateAsset CreateTemplateAsset(ctx).TemplateAssetCreate(templateAssetCreate).Execute()

Create template asset



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
	templateAssetCreate := *openapiclient.NewTemplateAssetCreate(int32(4), "build_source_image", *openapiclient.NewTemplateAssetFile("user-data", "text/plain", false, "/boot/grub/grub.cfg")) // TemplateAssetCreate | The template asset details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAssetAPI.CreateTemplateAsset(context.Background()).TemplateAssetCreate(templateAssetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssetAPI.CreateTemplateAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateAsset`: TemplateAsset
	fmt.Fprintf(os.Stdout, "Response from `TemplateAssetAPI.CreateTemplateAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateAssetCreate** | [**TemplateAssetCreate**](TemplateAssetCreate.md) | The template asset details | 

### Return type

[**TemplateAsset**](TemplateAsset.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplateAsset

> DeleteTemplateAsset(ctx, templateAssetId).Execute()

Delete template asset



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
	templateAssetId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplateAssetAPI.DeleteTemplateAsset(context.Background(), templateAssetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssetAPI.DeleteTemplateAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateAssetId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateAssetRequest struct via the builder pattern


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


## GetTemplateAsset

> TemplateAsset GetTemplateAsset(ctx, templateAssetId).Execute()

Get details for an template asset



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
	templateAssetId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAssetAPI.GetTemplateAsset(context.Background(), templateAssetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssetAPI.GetTemplateAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateAsset`: TemplateAsset
	fmt.Fprintf(os.Stdout, "Response from `TemplateAssetAPI.GetTemplateAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateAssetId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateAsset**](TemplateAsset.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateAssets

> TemplateAssetPaginatedList GetTemplateAssets(ctx).Page(page).Limit(limit).FilterTemplateId(filterTemplateId).FilterUsage(filterUsage).FilterFileMimeType(filterFileMimeType).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get a list of available template assets



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
	filterTemplateId := []string{"Inner_example"} // []string | Filter by templateId query param.           <p>              <b>Format: </b> filter.templateId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.templateId=$not:$like:John Doe&filter.templateId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterUsage := []string{"Inner_example"} // []string | Filter by usage query param.           <p>              <b>Format: </b> filter.usage={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.usage=$not:$like:John Doe&filter.usage=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterFileMimeType := []string{"Inner_example"} // []string | Filter by file.mimeType query param.           <p>              <b>Format: </b> filter.file.mimeType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.file.mimeType=$not:$like:John Doe&filter.file.mimeType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.           <p>              <b>Format: </b> filter.tags={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.tags=$not:$like:John Doe&filter.tags=like:John           </p>           <h4>Available Operations</h4><ul><li>$ilike</li> <li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> templateId:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>templateId</li> <li>usage</li> <li>file.mimeType</li> <li>createdAt</li> <li>modifiedAt</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,templateId,usage,file.name,file.mimeType           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>templateId</li> <li>usage</li> <li>file.name</li> <li>file.mimeType</li> <li>createdBy</li> <li>modifiedBy</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> id,templateId,usage,file.name,file.mimeType           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAssetAPI.GetTemplateAssets(context.Background()).Page(page).Limit(limit).FilterTemplateId(filterTemplateId).FilterUsage(filterUsage).FilterFileMimeType(filterFileMimeType).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssetAPI.GetTemplateAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateAssets`: TemplateAssetPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `TemplateAssetAPI.GetTemplateAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterTemplateId** | **[]string** | Filter by templateId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.templateId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.templateId&#x3D;$not:$like:John Doe&amp;filter.templateId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterUsage** | **[]string** | Filter by usage query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.usage&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.usage&#x3D;$not:$like:John Doe&amp;filter.usage&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterFileMimeType** | **[]string** | Filter by file.mimeType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.file.mimeType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.file.mimeType&#x3D;$not:$like:John Doe&amp;filter.file.mimeType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterTags** | **[]string** | Filter by tags query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.tags&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.tags&#x3D;$not:$like:John Doe&amp;filter.tags&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; templateId:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;templateId&lt;/li&gt; &lt;li&gt;usage&lt;/li&gt; &lt;li&gt;file.mimeType&lt;/li&gt; &lt;li&gt;createdAt&lt;/li&gt; &lt;li&gt;modifiedAt&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,templateId,usage,file.name,file.mimeType           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;templateId&lt;/li&gt; &lt;li&gt;usage&lt;/li&gt; &lt;li&gt;file.name&lt;/li&gt; &lt;li&gt;file.mimeType&lt;/li&gt; &lt;li&gt;createdBy&lt;/li&gt; &lt;li&gt;modifiedBy&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,templateId,usage,file.name,file.mimeType           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

### Return type

[**TemplateAssetPaginatedList**](TemplateAssetPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplateAsset

> TemplateAsset UpdateTemplateAsset(ctx, templateAssetId).TemplateAssetCreate(templateAssetCreate).Execute()

Update template asset



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
	templateAssetId := float32(8.14) // float32 | 
	templateAssetCreate := *openapiclient.NewTemplateAssetCreate(int32(4), "build_source_image", *openapiclient.NewTemplateAssetFile("user-data", "text/plain", false, "/boot/grub/grub.cfg")) // TemplateAssetCreate | The template asset details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAssetAPI.UpdateTemplateAsset(context.Background(), templateAssetId).TemplateAssetCreate(templateAssetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssetAPI.UpdateTemplateAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplateAsset`: TemplateAsset
	fmt.Fprintf(os.Stdout, "Response from `TemplateAssetAPI.UpdateTemplateAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateAssetId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateAssetCreate** | [**TemplateAssetCreate**](TemplateAssetCreate.md) | The template asset details | 

### Return type

[**TemplateAsset**](TemplateAsset.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

