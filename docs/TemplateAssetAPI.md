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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterTemplateId := []string{"Inner_example"} // []string | Filter by templateId query param.  **Format:** filter.templateId={$not}:OPERATION:VALUE    **Example:** filter.templateId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterUsage := []string{"Inner_example"} // []string | Filter by usage query param.  **Format:** filter.usage={$not}:OPERATION:VALUE    **Example:** filter.usage=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterFileMimeType := []string{"Inner_example"} // []string | Filter by file.mimeType query param.  **Format:** filter.file.mimeType={$not}:OPERATION:VALUE    **Example:** filter.file.mimeType=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.  **Format:** filter.tags={$not}:OPERATION:VALUE    **Example:** filter.tags=$eq:John Doe&filter.tags=$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=templateId:DESC   **Default Value:** templateId:DESC  **Available Fields** - id  - templateId  - usage  - file.mimeType  - createdAt  - modifiedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,templateId,usage,file.name,file.mimeType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - templateId  - usage  - file.name  - file.mimeType  - createdBy  - modifiedBy  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,templateId,usage,file.name,file.mimeType   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

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
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterTemplateId** | **[]string** | Filter by templateId query param.  **Format:** filter.templateId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.templateId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterUsage** | **[]string** | Filter by usage query param.  **Format:** filter.usage&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.usage&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterFileMimeType** | **[]string** | Filter by file.mimeType query param.  **Format:** filter.file.mimeType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.file.mimeType&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterTags** | **[]string** | Filter by tags query param.  **Format:** filter.tags&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.tags&#x3D;$eq:John Doe&amp;filter.tags&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;templateId:DESC   **Default Value:** templateId:DESC  **Available Fields** - id  - templateId  - usage  - file.mimeType  - createdAt  - modifiedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,templateId,usage,file.name,file.mimeType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - templateId  - usage  - file.name  - file.mimeType  - createdBy  - modifiedBy  | 
 **select_** | **string** | List of fields to select.  **Example:** id,templateId,usage,file.name,file.mimeType   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

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

