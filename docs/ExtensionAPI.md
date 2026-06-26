# \ExtensionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateExtension**](ExtensionAPI.md#ActivateExtension) | **Post** /api/v2/extensions/{extensionId}/actions/activate | Activates draft or suspended extension
[**ArchiveExtension**](ExtensionAPI.md#ArchiveExtension) | **Post** /api/v2/extensions/{extensionId}/actions/archive | Archive published extension
[**CreateExtension**](ExtensionAPI.md#CreateExtension) | **Post** /api/v2/extensions | Create extension
[**DeleteExtension**](ExtensionAPI.md#DeleteExtension) | **Post** /api/v2/extensions/{extensionId}/actions/delete | Delete extension
[**DeleteExtensionSiteConfig**](ExtensionAPI.md#DeleteExtensionSiteConfig) | **Delete** /api/v2/extensions/{extensionId}/siteConfig/{siteId} | 
[**GetExtension**](ExtensionAPI.md#GetExtension) | **Get** /api/v2/extensions/{extensionId} | Get details for an extension
[**GetExtensionSiteConfig**](ExtensionAPI.md#GetExtensionSiteConfig) | **Get** /api/v2/extensions/{extensionId}/siteConfig/{siteId} | 
[**GetExtensionSiteConfigs**](ExtensionAPI.md#GetExtensionSiteConfigs) | **Get** /api/v2/extensions/{extensionId}/siteConfig | 
[**GetExtensionSiteCredentials**](ExtensionAPI.md#GetExtensionSiteCredentials) | **Get** /api/v2/extensions/{extensionId}/siteConfig/{siteId}/credentials | Get Extension Site Config credentials
[**GetExtensions**](ExtensionAPI.md#GetExtensions) | **Get** /api/v2/extensions | Get a list of available extensions
[**GetSiteExtensionConfigs**](ExtensionAPI.md#GetSiteExtensionConfigs) | **Get** /api/v2/extensions/siteConfig/{siteId} | 
[**PublishExtension**](ExtensionAPI.md#PublishExtension) | **Post** /api/v2/extensions/{extensionId}/actions/publish | Activates draft or suspended extension
[**SetExtensionSiteConfig**](ExtensionAPI.md#SetExtensionSiteConfig) | **Post** /api/v2/extensions/{extensionId}/siteConfig/{siteId} | 
[**SuspendExtension**](ExtensionAPI.md#SuspendExtension) | **Post** /api/v2/extensions/{extensionId}/actions/suspend | Suspend extension
[**UpdateExtension**](ExtensionAPI.md#UpdateExtension) | **Patch** /api/v2/extensions/{extensionId} | Update extension



## ActivateExtension

> ActivateExtension(ctx, extensionId).IfMatch(ifMatch).Execute()

Activates draft or suspended extension



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
	extensionId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.ActivateExtension(context.Background(), extensionId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.ActivateExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateExtensionRequest struct via the builder pattern


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
	extensionId := int64(789) // int64 | 
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
**extensionId** | **int64** |  | 

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
	createExtension := *openapiclient.NewCreateExtension("My App", "My App Description", "application", *openapiclient.NewExtensionDefinition("Kind_example", "SchemaVersion_example", "Name_example", "Label_example", "ExtensionType_example", "Vendor_example", "ExtensionVersion_example", "Icon_example", *openapiclient.NewExtensionDependency("ControllerVersion_example"), []openapiclient.ExtensionDefinitionInputsDataItem{openapiclient.ExtensionDefinition_inputsDataItem{ExtensionInputBoolean: openapiclient.NewExtensionInputBoolean("Label_example", "Name_example", openapiclient.ExtensionInputType("ExtensionInputString"), map[string]interface{}(123))}}, []openapiclient.ExtensionOutput{*openapiclient.NewExtensionOutput("Label_example", "Name_example", "OutputType_example")}, []openapiclient.ExtensionAsset{*openapiclient.NewExtensionAsset("Label_example", "Name_example", openapiclient.ExtensionAssetType("AnsibleBundle"))})) // CreateExtension | The extension details

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


## DeleteExtension

> DeleteExtension(ctx, extensionId).IfMatch(ifMatch).Execute()

Delete extension



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
	extensionId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.DeleteExtension(context.Background(), extensionId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.DeleteExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionRequest struct via the builder pattern


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


## DeleteExtensionSiteConfig

> DeleteExtensionSiteConfig(ctx, extensionId, siteId).Execute()





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
	extensionId := int64(789) // int64 | 
	siteId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.DeleteExtensionSiteConfig(context.Background(), extensionId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.DeleteExtensionSiteConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionSiteConfigRequest struct via the builder pattern


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
	extensionId := int64(789) // int64 | 

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
**extensionId** | **int64** |  | 

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


## GetExtensionSiteConfig

> []ExtensionConfigValue GetExtensionSiteConfig(ctx, extensionId, siteId).Execute()



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
	extensionId := int64(789) // int64 | 
	siteId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetExtensionSiteConfig(context.Background(), extensionId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.GetExtensionSiteConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionSiteConfig`: []ExtensionConfigValue
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.GetExtensionSiteConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionSiteConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ExtensionConfigValue**](ExtensionConfigValue.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionSiteConfigs

> []ExtensionSiteConfig GetExtensionSiteConfigs(ctx, extensionId).Execute()



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
	extensionId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetExtensionSiteConfigs(context.Background(), extensionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.GetExtensionSiteConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionSiteConfigs`: []ExtensionSiteConfig
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.GetExtensionSiteConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionSiteConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ExtensionSiteConfig**](ExtensionSiteConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionSiteCredentials

> map[string]interface{} GetExtensionSiteCredentials(ctx, extensionId, siteId).Execute()

Get Extension Site Config credentials



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
	extensionId := int64(789) // int64 | 
	siteId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetExtensionSiteCredentials(context.Background(), extensionId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.GetExtensionSiteCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionSiteCredentials`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.GetExtensionSiteCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionSiteCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensions

> ExtensionPaginatedList GetExtensions(ctx).Page(page).Limit(limit).FilterKind(filterKind).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$eq:John Doe&filter.status=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe&filter.label=$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:ASC  **Available Fields** - id  - name  - status  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetExtensions(context.Background()).Page(page).Limit(limit).FilterKind(filterKind).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$eq:John Doe&amp;filter.status&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe&amp;filter.label&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or | 
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


## GetSiteExtensionConfigs

> []SiteExtensionConfig GetSiteExtensionConfigs(ctx, siteId).Execute()



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
	siteId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionAPI.GetSiteExtensionConfigs(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.GetSiteExtensionConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteExtensionConfigs`: []SiteExtensionConfig
	fmt.Fprintf(os.Stdout, "Response from `ExtensionAPI.GetSiteExtensionConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteExtensionConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SiteExtensionConfig**](SiteExtensionConfig.md)

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

Activates draft or suspended extension



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
	extensionId := int64(789) // int64 | 
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
**extensionId** | **int64** |  | 

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


## SetExtensionSiteConfig

> SetExtensionSiteConfig(ctx, extensionId, siteId).ExtensionConfigValue(extensionConfigValue).Execute()





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
	extensionId := int64(789) // int64 | 
	siteId := int64(789) // int64 | 
	extensionConfigValue := []openapiclient.ExtensionConfigValue{*openapiclient.NewExtensionConfigValue("Label_example", openapiclient.ExtensionConfigValue_value{Bool: new(bool)})} // []ExtensionConfigValue | The extension configuration details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.SetExtensionSiteConfig(context.Background(), extensionId, siteId).ExtensionConfigValue(extensionConfigValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.SetExtensionSiteConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetExtensionSiteConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extensionConfigValue** | [**[]ExtensionConfigValue**](ExtensionConfigValue.md) | The extension configuration details | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendExtension

> SuspendExtension(ctx, extensionId).IfMatch(ifMatch).Execute()

Suspend extension



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
	extensionId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionAPI.SuspendExtension(context.Background(), extensionId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionAPI.SuspendExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendExtensionRequest struct via the builder pattern


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
	extensionId := int64(789) // int64 | 
	updateExtension := *openapiclient.NewUpdateExtension(*openapiclient.NewExtensionDefinition("Kind_example", "SchemaVersion_example", "Name_example", "Label_example", "ExtensionType_example", "Vendor_example", "ExtensionVersion_example", "Icon_example", *openapiclient.NewExtensionDependency("ControllerVersion_example"), []openapiclient.ExtensionDefinitionInputsDataItem{openapiclient.ExtensionDefinition_inputsDataItem{ExtensionInputBoolean: openapiclient.NewExtensionInputBoolean("Label_example", "Name_example", openapiclient.ExtensionInputType("ExtensionInputString"), map[string]interface{}(123))}}, []openapiclient.ExtensionOutput{*openapiclient.NewExtensionOutput("Label_example", "Name_example", "OutputType_example")}, []openapiclient.ExtensionAsset{*openapiclient.NewExtensionAsset("Label_example", "Name_example", openapiclient.ExtensionAssetType("AnsibleBundle"))})) // UpdateExtension | The extension details
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
**extensionId** | **int64** |  | 

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

