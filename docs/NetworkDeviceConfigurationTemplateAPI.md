# \NetworkDeviceConfigurationTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplateAPI.md#CreateNetworkDeviceConfigurationTemplate) | **Post** /api/v2/network-device-configuration-templates | Creates a Network Device Configuration Template
[**DeleteNetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplateAPI.md#DeleteNetworkDeviceConfigurationTemplate) | **Delete** /api/v2/network-device-configuration-templates/{networkDeviceConfigurationTemplateId} | Deletes a Network Device Configuration Template
[**GetNetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplateAPI.md#GetNetworkDeviceConfigurationTemplate) | **Get** /api/v2/network-device-configuration-templates/{networkDeviceConfigurationTemplateId} | Get Network Device Configuration Template information
[**GetNetworkDeviceConfigurationTemplates**](NetworkDeviceConfigurationTemplateAPI.md#GetNetworkDeviceConfigurationTemplates) | **Get** /api/v2/network-device-configuration-templates | Get all Network Device Configuration Templates
[**UpdateNetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplateAPI.md#UpdateNetworkDeviceConfigurationTemplate) | **Patch** /api/v2/network-device-configuration-templates/{networkDeviceConfigurationTemplateId} | Updates Network Device Configuration Template information



## CreateNetworkDeviceConfigurationTemplate

> NetworkDeviceConfigurationTemplate CreateNetworkDeviceConfigurationTemplate(ctx).CreateNetworkDeviceConfigurationTemplate(createNetworkDeviceConfigurationTemplate).Execute()

Creates a Network Device Configuration Template



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
	createNetworkDeviceConfigurationTemplate := *openapiclient.NewCreateNetworkDeviceConfigurationTemplate("NetworkType_example", "NetworkDeviceDriver_example", "NetworkDevicePosition_example", "RemoteNetworkDevicePosition_example", float32(123), "BgpNumbering_example", "BgpLinkConfiguration_example", "ExecutionType_example", "LibraryLabel_example", "Configuration_example") // CreateNetworkDeviceConfigurationTemplate | The Network Device Configuration Template create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceConfigurationTemplateAPI.CreateNetworkDeviceConfigurationTemplate(context.Background()).CreateNetworkDeviceConfigurationTemplate(createNetworkDeviceConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceConfigurationTemplateAPI.CreateNetworkDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDeviceConfigurationTemplate`: NetworkDeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceConfigurationTemplateAPI.CreateNetworkDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDeviceConfigurationTemplate** | [**CreateNetworkDeviceConfigurationTemplate**](CreateNetworkDeviceConfigurationTemplate.md) | The Network Device Configuration Template create object | 

### Return type

[**NetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDeviceConfigurationTemplate

> DeleteNetworkDeviceConfigurationTemplate(ctx, networkDeviceConfigurationTemplateId).Execute()

Deletes a Network Device Configuration Template



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
	networkDeviceConfigurationTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceConfigurationTemplateAPI.DeleteNetworkDeviceConfigurationTemplate(context.Background(), networkDeviceConfigurationTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceConfigurationTemplateAPI.DeleteNetworkDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceConfigurationTemplateRequest struct via the builder pattern


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


## GetNetworkDeviceConfigurationTemplate

> NetworkDeviceConfigurationTemplate GetNetworkDeviceConfigurationTemplate(ctx, networkDeviceConfigurationTemplateId).Execute()

Get Network Device Configuration Template information



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
	networkDeviceConfigurationTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceConfigurationTemplateAPI.GetNetworkDeviceConfigurationTemplate(context.Background(), networkDeviceConfigurationTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceConfigurationTemplateAPI.GetNetworkDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceConfigurationTemplate`: NetworkDeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceConfigurationTemplateAPI.GetNetworkDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceConfigurationTemplates

> NetworkDeviceConfigurationTemplatePaginatedList GetNetworkDeviceConfigurationTemplates(ctx).Page(page).Limit(limit).FilterId(filterId).FilterNetworkType(filterNetworkType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterNetworkDevicePosition(filterNetworkDevicePosition).FilterRemoteNetworkDevicePosition(filterRemoteNetworkDevicePosition).FilterMlagPair(filterMlagPair).FilterBgpNumbering(filterBgpNumbering).FilterBgpLinkConfiguration(filterBgpLinkConfiguration).FilterExecutionType(filterExecutionType).FilterLibraryLabel(filterLibraryLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Network Device Configuration Templates



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe&filter.id=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkType := []string{"Inner_example"} // []string | Filter by networkType query param.  **Format:** filter.networkType={$not}:OPERATION:VALUE    **Example:** filter.networkType=$eq:John Doe&filter.networkType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkDeviceDriver := []string{"Inner_example"} // []string | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver=$eq:John Doe&filter.networkDeviceDriver=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkDevicePosition := []string{"Inner_example"} // []string | Filter by networkDevicePosition query param.  **Format:** filter.networkDevicePosition={$not}:OPERATION:VALUE    **Example:** filter.networkDevicePosition=$eq:John Doe&filter.networkDevicePosition=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterRemoteNetworkDevicePosition := []string{"Inner_example"} // []string | Filter by remoteNetworkDevicePosition query param.  **Format:** filter.remoteNetworkDevicePosition={$not}:OPERATION:VALUE    **Example:** filter.remoteNetworkDevicePosition=$eq:John Doe&filter.remoteNetworkDevicePosition=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterMlagPair := []string{"Inner_example"} // []string | Filter by mlagPair query param.  **Format:** filter.mlagPair={$not}:OPERATION:VALUE    **Example:** filter.mlagPair=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterBgpNumbering := []string{"Inner_example"} // []string | Filter by bgpNumbering query param.  **Format:** filter.bgpNumbering={$not}:OPERATION:VALUE    **Example:** filter.bgpNumbering=$eq:John Doe&filter.bgpNumbering=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterBgpLinkConfiguration := []string{"Inner_example"} // []string | Filter by bgpLinkConfiguration query param.  **Format:** filter.bgpLinkConfiguration={$not}:OPERATION:VALUE    **Example:** filter.bgpLinkConfiguration=$eq:John Doe&filter.bgpLinkConfiguration=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterExecutionType := []string{"Inner_example"} // []string | Filter by executionType query param.  **Format:** filter.executionType={$not}:OPERATION:VALUE    **Example:** filter.executionType=$eq:John Doe&filter.executionType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterLibraryLabel := []string{"Inner_example"} // []string | Filter by libraryLabel query param.  **Format:** filter.libraryLabel={$not}:OPERATION:VALUE    **Example:** filter.libraryLabel=$eq:John Doe&filter.libraryLabel=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=networkType:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - mlagPair  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,networkType,networkDeviceDriver,networkDevicePosition,remoteNetworkDevicePosition   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - mlagPair  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceConfigurationTemplateAPI.GetNetworkDeviceConfigurationTemplates(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterNetworkType(filterNetworkType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterNetworkDevicePosition(filterNetworkDevicePosition).FilterRemoteNetworkDevicePosition(filterRemoteNetworkDevicePosition).FilterMlagPair(filterMlagPair).FilterBgpNumbering(filterBgpNumbering).FilterBgpLinkConfiguration(filterBgpLinkConfiguration).FilterExecutionType(filterExecutionType).FilterLibraryLabel(filterLibraryLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceConfigurationTemplateAPI.GetNetworkDeviceConfigurationTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceConfigurationTemplates`: NetworkDeviceConfigurationTemplatePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceConfigurationTemplateAPI.GetNetworkDeviceConfigurationTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceConfigurationTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkType** | **[]string** | Filter by networkType query param.  **Format:** filter.networkType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkType&#x3D;$eq:John Doe&amp;filter.networkType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkDeviceDriver** | **[]string** | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver&#x3D;$eq:John Doe&amp;filter.networkDeviceDriver&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkDevicePosition** | **[]string** | Filter by networkDevicePosition query param.  **Format:** filter.networkDevicePosition&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDevicePosition&#x3D;$eq:John Doe&amp;filter.networkDevicePosition&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterRemoteNetworkDevicePosition** | **[]string** | Filter by remoteNetworkDevicePosition query param.  **Format:** filter.remoteNetworkDevicePosition&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.remoteNetworkDevicePosition&#x3D;$eq:John Doe&amp;filter.remoteNetworkDevicePosition&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterMlagPair** | **[]string** | Filter by mlagPair query param.  **Format:** filter.mlagPair&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.mlagPair&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterBgpNumbering** | **[]string** | Filter by bgpNumbering query param.  **Format:** filter.bgpNumbering&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.bgpNumbering&#x3D;$eq:John Doe&amp;filter.bgpNumbering&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterBgpLinkConfiguration** | **[]string** | Filter by bgpLinkConfiguration query param.  **Format:** filter.bgpLinkConfiguration&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.bgpLinkConfiguration&#x3D;$eq:John Doe&amp;filter.bgpLinkConfiguration&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterExecutionType** | **[]string** | Filter by executionType query param.  **Format:** filter.executionType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.executionType&#x3D;$eq:John Doe&amp;filter.executionType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterLibraryLabel** | **[]string** | Filter by libraryLabel query param.  **Format:** filter.libraryLabel&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.libraryLabel&#x3D;$eq:John Doe&amp;filter.libraryLabel&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;networkType:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - mlagPair  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,networkType,networkDeviceDriver,networkDevicePosition,remoteNetworkDevicePosition   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - mlagPair  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  | 

### Return type

[**NetworkDeviceConfigurationTemplatePaginatedList**](NetworkDeviceConfigurationTemplatePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDeviceConfigurationTemplate

> NetworkDeviceConfigurationTemplate UpdateNetworkDeviceConfigurationTemplate(ctx, networkDeviceConfigurationTemplateId).UpdateNetworkDeviceConfigurationTemplate(updateNetworkDeviceConfigurationTemplate).Execute()

Updates Network Device Configuration Template information



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
	networkDeviceConfigurationTemplateId := float32(8.14) // float32 | 
	updateNetworkDeviceConfigurationTemplate := *openapiclient.NewUpdateNetworkDeviceConfigurationTemplate() // UpdateNetworkDeviceConfigurationTemplate | The Network Device Configuration Template update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceConfigurationTemplateAPI.UpdateNetworkDeviceConfigurationTemplate(context.Background(), networkDeviceConfigurationTemplateId).UpdateNetworkDeviceConfigurationTemplate(updateNetworkDeviceConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceConfigurationTemplateAPI.UpdateNetworkDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceConfigurationTemplate`: NetworkDeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceConfigurationTemplateAPI.UpdateNetworkDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDeviceConfigurationTemplate** | [**UpdateNetworkDeviceConfigurationTemplate**](UpdateNetworkDeviceConfigurationTemplate.md) | The Network Device Configuration Template update object | 

### Return type

[**NetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

