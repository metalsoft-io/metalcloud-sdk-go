# \ExtensionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveExtension**](ExtensionAPI.md#ArchiveExtension) | **Post** /api/v2/extensions/{extensionId}/actions/archive | Archive published extension
[**CreateExtension**](ExtensionAPI.md#CreateExtension) | **Post** /api/v2/extensions | Create extension
[**GetExtension**](ExtensionAPI.md#GetExtension) | **Get** /api/v2/extensions/{extensionId} | Get details for an extension
[**GetExtensions**](ExtensionAPI.md#GetExtensions) | **Get** /api/v2/extensions | Get a list of available extensions
[**MakePublicExtension**](ExtensionAPI.md#MakePublicExtension) | **Post** /api/v2/extensions/{extensionId}/actions/make-public | Makes extension public
[**PublishExtension**](ExtensionAPI.md#PublishExtension) | **Post** /api/v2/extensions/{extensionId}/actions/publish | Publish draft extension
[**UpdateExtension**](ExtensionAPI.md#UpdateExtension) | **Patch** /api/v2/extensions/{extensionId} | Update extension



## ArchiveExtension

> ArchiveExtension(ctx, extensionId).IfMatch(ifMatch).Execute()

Archive published extension



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
	extensionId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.ArchiveExtension(context.Background(), extensionId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.ArchiveExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

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


## CreateExtension

> Extension CreateExtension(ctx).CreateExtension(createExtension).Execute()

Create extension



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
	createExtension := *openapiclient.NewCreateExtension("My App", "My App Description", "application", *openapiclient.NewExtensionDefinition("Kind_example", "SchemaVersion_example", "Name_example", "Label_example", "ExtensionType_example", "Vendor_example", "ExtensionVersion_example", "Icon_example", *openapiclient.NewExtensionDependency("ControllerVersion_example"), []openapiclient.ExtensionDefinitionInputsInner{openapiclient.ExtensionDefinition_inputs_inner{ExtensionInputBoolean: openapiclient.NewExtensionInputBoolean("Label_example", "Name_example", openapiclient.ExtensionInputType("ExtensionInputString"), *openapiclient.NewExtensionInputOptionBoolean())}}, []openapiclient.ExtensionOutput{*openapiclient.NewExtensionOutput("Label_example", "Name_example", "OutputType_example")}, []openapiclient.ExtensionAsset{*openapiclient.NewExtensionAsset("Label_example", "Name_example", "AssetType_example", "Url_example")})) // CreateExtension | The extension details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.CreateExtension(context.Background()).CreateExtension(createExtension).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.CreateExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExtension`: Extension
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.CreateExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExtension** | [**CreateExtension**](CreateExtension.md) | The extension details | 

### Return type

[**Extension**](Extension.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtension

> Extension GetExtension(ctx, extensionId).Execute()

Get details for an extension



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
	extensionId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetExtension(context.Background(), extensionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.GetExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtension`: Extension
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.GetExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Extension**](Extension.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensions

> ExtensionPaginatedList GetExtensions(ctx).Page(page).Limit(limit).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).FilterIsPublic(filterIsPublic).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of available extensions



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
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** -1    **Max Value:** -1   If provided value is greater than max value, max value will be applied.  (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe&filter.label=$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or (optional)
	filterIsPublic := []string{"Inner_example"} // []string | Filter by isPublic query param.  **Format:** filter.isPublic={$not}:OPERATION:VALUE    **Example:** filter.isPublic=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:ASC  **Available Fields** - id  - name  - status  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetExtensions(context.Background()).Page(page).Limit(limit).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).FilterIsPublic(filterIsPublic).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.GetExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensions`: ExtensionPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.GetExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** -1    **Max Value:** -1   If provided value is greater than max value, max value will be applied.  | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe&amp;filter.label&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or | 
 **filterIsPublic** | **[]string** | Filter by isPublic query param.  **Format:** filter.isPublic&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.isPublic&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:ASC  **Available Fields** - id  - name  - status  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  | 

### Return type

[**ExtensionPaginatedList**](ExtensionPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakePublicExtension

> MakePublicExtension(ctx, extensionId).IfMatch(ifMatch).Execute()

Makes extension public



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
	extensionId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.MakePublicExtension(context.Background(), extensionId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.MakePublicExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakePublicExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

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


## PublishExtension

> PublishExtension(ctx, extensionId).IfMatch(ifMatch).Execute()

Publish draft extension



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
	extensionId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.PublishExtension(context.Background(), extensionId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.PublishExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

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


## UpdateExtension

> Extension UpdateExtension(ctx, extensionId).UpdateExtension(updateExtension).IfMatch(ifMatch).Execute()

Update extension



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
	extensionId := float32(8.14) // float32 | 
	updateExtension := *openapiclient.NewUpdateExtension("My App", "My App Description", *openapiclient.NewExtensionDefinition("Kind_example", "SchemaVersion_example", "Name_example", "Label_example", "ExtensionType_example", "Vendor_example", "ExtensionVersion_example", "Icon_example", *openapiclient.NewExtensionDependency("ControllerVersion_example"), []openapiclient.ExtensionDefinitionInputsInner{openapiclient.ExtensionDefinition_inputs_inner{ExtensionInputBoolean: openapiclient.NewExtensionInputBoolean("Label_example", "Name_example", openapiclient.ExtensionInputType("ExtensionInputString"), *openapiclient.NewExtensionInputOptionBoolean())}}, []openapiclient.ExtensionOutput{*openapiclient.NewExtensionOutput("Label_example", "Name_example", "OutputType_example")}, []openapiclient.ExtensionAsset{*openapiclient.NewExtensionAsset("Label_example", "Name_example", "AssetType_example", "Url_example")})) // UpdateExtension | The extension details
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.UpdateExtension(context.Background(), extensionId).UpdateExtension(updateExtension).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.UpdateExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtension`: Extension
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.UpdateExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExtension** | [**UpdateExtension**](UpdateExtension.md) | The extension details | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Extension**](Extension.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

