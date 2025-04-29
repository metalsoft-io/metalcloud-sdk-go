# \ExtensionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveExtension**](ExtensionAPI.md#ArchiveExtension) | **Post** /api/v2/extensions/{extensionId}/actions/archive | Archive published extension
[**CreateExtension**](ExtensionAPI.md#CreateExtension) | **Post** /api/v2/extensions | Create extension
[**GetExtension**](ExtensionAPI.md#GetExtension) | **Get** /api/v2/extensions/{extensionId} | Get details for an extension
[**GetExtensions**](ExtensionAPI.md#GetExtensions) | **Get** /api/v2/extensions | Get a list of available extensions
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
	createExtension := *openapiclient.NewCreateExtension("My App", "My App Description", "application", *openapiclient.NewExtensionDefinition("Kind_example", "SchemaVersion_example", "Name_example", "Label_example", "ExtensionType_example", "Vendor_example", "ExtensionVersion_example", "Icon_example", *openapiclient.NewExtensionDependency("ControllerVersion_example"), []openapiclient.ExtensionInput{*openapiclient.NewExtensionInput("Label_example", "Name_example", "InputType_example", openapiclient.ExtensionInput_options{ExtensionInputBoolean: openapiclient.NewExtensionInputBoolean()})}, []openapiclient.ExtensionOutput{*openapiclient.NewExtensionOutput("Label_example", "Name_example", "OutputType_example")}, []openapiclient.ExtensionAsset{*openapiclient.NewExtensionAsset("Label_example", "Name_example", "AssetType_example", "Url_example")})) // CreateExtension | The extension details

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

> ExtensionPaginatedList GetExtensions(ctx).Page(page).Limit(limit).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> -1           </p>       <p>              <b>Max Value: </b> -1           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.           <p>              <b>Format: </b> filter.status={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.status=$not:$like:John Doe&filter.status=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$null</li> <li>$ilike</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>name</li> <li>status</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetExtensions(context.Background()).Page(page).Limit(limit).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; -1           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; -1           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterStatus** | **[]string** | Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;status&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

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
	updateExtension := *openapiclient.NewUpdateExtension("My App", "My App Description", *openapiclient.NewExtensionDefinition("Kind_example", "SchemaVersion_example", "Name_example", "Label_example", "ExtensionType_example", "Vendor_example", "ExtensionVersion_example", "Icon_example", *openapiclient.NewExtensionDependency("ControllerVersion_example"), []openapiclient.ExtensionInput{*openapiclient.NewExtensionInput("Label_example", "Name_example", "InputType_example", openapiclient.ExtensionInput_options{ExtensionInputBoolean: openapiclient.NewExtensionInputBoolean()})}, []openapiclient.ExtensionOutput{*openapiclient.NewExtensionOutput("Label_example", "Name_example", "OutputType_example")}, []openapiclient.ExtensionAsset{*openapiclient.NewExtensionAsset("Label_example", "Name_example", "AssetType_example", "Url_example")})) // UpdateExtension | The extension details
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

