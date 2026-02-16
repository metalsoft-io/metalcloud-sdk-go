# \NetworkDeviceBGPConfigurationTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDeviceBGPConfigurationTemplate**](NetworkDeviceBGPConfigurationTemplateAPI.md#CreateNetworkDeviceBGPConfigurationTemplate) | **Post** /api/v2/network-device-bgp-configuration-templates | Creates a Network Device BGP Configuration Template
[**DeleteNetworkDeviceBGPConfigurationTemplate**](NetworkDeviceBGPConfigurationTemplateAPI.md#DeleteNetworkDeviceBGPConfigurationTemplate) | **Delete** /api/v2/network-device-bgp-configuration-templates/{networkDeviceBGPConfigurationTemplateId} | Deletes a Network Device BGP Configuration Template
[**GetNetworkDeviceBGPConfigurationTemplate**](NetworkDeviceBGPConfigurationTemplateAPI.md#GetNetworkDeviceBGPConfigurationTemplate) | **Get** /api/v2/network-device-bgp-configuration-templates/{networkDeviceBGPConfigurationTemplateId} | Get Network Device BGP Configuration Template information
[**GetNetworkDeviceBGPConfigurationTemplates**](NetworkDeviceBGPConfigurationTemplateAPI.md#GetNetworkDeviceBGPConfigurationTemplates) | **Get** /api/v2/network-device-bgp-configuration-templates | Get all Network Device BGP Configuration Templates
[**UpdateNetworkDeviceBGPConfigurationTemplate**](NetworkDeviceBGPConfigurationTemplateAPI.md#UpdateNetworkDeviceBGPConfigurationTemplate) | **Patch** /api/v2/network-device-bgp-configuration-templates/{networkDeviceBGPConfigurationTemplateId} | Updates Network Device BGP Configuration Template information



## CreateNetworkDeviceBGPConfigurationTemplate

> NetworkDeviceBGPConfigurationTemplate CreateNetworkDeviceBGPConfigurationTemplate(ctx).CreateNetworkDeviceBGPConfigurationTemplate(createNetworkDeviceBGPConfigurationTemplate).Execute()

Creates a Network Device BGP Configuration Template



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
	createNetworkDeviceBGPConfigurationTemplate := *openapiclient.NewCreateNetworkDeviceBGPConfigurationTemplate("Action_example", "NetworkType_example", "NetworkDeviceDriver_example", "NetworkDevicePosition_example", "RemoteNetworkDevicePosition_example", "BgpNumbering_example", "BgpLinkConfiguration_example", "ExecutionType_example", "LibraryLabel_example", "Configuration_example") // CreateNetworkDeviceBGPConfigurationTemplate | The Network Device BGP Configuration Template create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPConfigurationTemplateAPI.CreateNetworkDeviceBGPConfigurationTemplate(context.Background()).CreateNetworkDeviceBGPConfigurationTemplate(createNetworkDeviceBGPConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPConfigurationTemplateAPI.CreateNetworkDeviceBGPConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDeviceBGPConfigurationTemplate`: NetworkDeviceBGPConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPConfigurationTemplateAPI.CreateNetworkDeviceBGPConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceBGPConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDeviceBGPConfigurationTemplate** | [**CreateNetworkDeviceBGPConfigurationTemplate**](CreateNetworkDeviceBGPConfigurationTemplate.md) | The Network Device BGP Configuration Template create object | 

### Return type

[**NetworkDeviceBGPConfigurationTemplate**](NetworkDeviceBGPConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDeviceBGPConfigurationTemplate

> DeleteNetworkDeviceBGPConfigurationTemplate(ctx, networkDeviceBGPConfigurationTemplateId).Execute()

Deletes a Network Device BGP Configuration Template



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
	networkDeviceBGPConfigurationTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceBGPConfigurationTemplateAPI.DeleteNetworkDeviceBGPConfigurationTemplate(context.Background(), networkDeviceBGPConfigurationTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPConfigurationTemplateAPI.DeleteNetworkDeviceBGPConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceBGPConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceBGPConfigurationTemplateRequest struct via the builder pattern


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


## GetNetworkDeviceBGPConfigurationTemplate

> NetworkDeviceBGPConfigurationTemplate GetNetworkDeviceBGPConfigurationTemplate(ctx, networkDeviceBGPConfigurationTemplateId).Execute()

Get Network Device BGP Configuration Template information



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
	networkDeviceBGPConfigurationTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPConfigurationTemplateAPI.GetNetworkDeviceBGPConfigurationTemplate(context.Background(), networkDeviceBGPConfigurationTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPConfigurationTemplateAPI.GetNetworkDeviceBGPConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceBGPConfigurationTemplate`: NetworkDeviceBGPConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPConfigurationTemplateAPI.GetNetworkDeviceBGPConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceBGPConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceBGPConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceBGPConfigurationTemplate**](NetworkDeviceBGPConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceBGPConfigurationTemplates

> NetworkDeviceBGPConfigurationTemplatePaginatedList GetNetworkDeviceBGPConfigurationTemplates(ctx).Page(page).Limit(limit).FilterId(filterId).FilterAction(filterAction).FilterNetworkType(filterNetworkType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterNetworkDevicePosition(filterNetworkDevicePosition).FilterRemoteNetworkDevicePosition(filterRemoteNetworkDevicePosition).FilterBgpNumbering(filterBgpNumbering).FilterBgpLinkConfiguration(filterBgpLinkConfiguration).FilterExecutionType(filterExecutionType).FilterLibraryLabel(filterLibraryLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Network Device BGP Configuration Templates



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
	filterAction := []string{"Inner_example"} // []string | Filter by action query param.  **Format:** filter.action={$not}:OPERATION:VALUE    **Example:** filter.action=$eq:John Doe&filter.action=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkType := []string{"Inner_example"} // []string | Filter by networkType query param.  **Format:** filter.networkType={$not}:OPERATION:VALUE    **Example:** filter.networkType=$eq:John Doe&filter.networkType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkDeviceDriver := []string{"Inner_example"} // []string | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver=$eq:John Doe&filter.networkDeviceDriver=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkDevicePosition := []string{"Inner_example"} // []string | Filter by networkDevicePosition query param.  **Format:** filter.networkDevicePosition={$not}:OPERATION:VALUE    **Example:** filter.networkDevicePosition=$eq:John Doe&filter.networkDevicePosition=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterRemoteNetworkDevicePosition := []string{"Inner_example"} // []string | Filter by remoteNetworkDevicePosition query param.  **Format:** filter.remoteNetworkDevicePosition={$not}:OPERATION:VALUE    **Example:** filter.remoteNetworkDevicePosition=$eq:John Doe&filter.remoteNetworkDevicePosition=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterBgpNumbering := []string{"Inner_example"} // []string | Filter by bgpNumbering query param.  **Format:** filter.bgpNumbering={$not}:OPERATION:VALUE    **Example:** filter.bgpNumbering=$eq:John Doe&filter.bgpNumbering=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterBgpLinkConfiguration := []string{"Inner_example"} // []string | Filter by bgpLinkConfiguration query param.  **Format:** filter.bgpLinkConfiguration={$not}:OPERATION:VALUE    **Example:** filter.bgpLinkConfiguration=$eq:John Doe&filter.bgpLinkConfiguration=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterExecutionType := []string{"Inner_example"} // []string | Filter by executionType query param.  **Format:** filter.executionType={$not}:OPERATION:VALUE    **Example:** filter.executionType=$eq:John Doe&filter.executionType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterLibraryLabel := []string{"Inner_example"} // []string | Filter by libraryLabel query param.  **Format:** filter.libraryLabel={$not}:OPERATION:VALUE    **Example:** filter.libraryLabel=$eq:John Doe&filter.libraryLabel=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=action:DESC   **Default Value:** id:DESC  **Available Fields** - id  - action  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,action,networkType,networkDeviceDriver,networkDevicePosition   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - action  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPConfigurationTemplateAPI.GetNetworkDeviceBGPConfigurationTemplates(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterAction(filterAction).FilterNetworkType(filterNetworkType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterNetworkDevicePosition(filterNetworkDevicePosition).FilterRemoteNetworkDevicePosition(filterRemoteNetworkDevicePosition).FilterBgpNumbering(filterBgpNumbering).FilterBgpLinkConfiguration(filterBgpLinkConfiguration).FilterExecutionType(filterExecutionType).FilterLibraryLabel(filterLibraryLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPConfigurationTemplateAPI.GetNetworkDeviceBGPConfigurationTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceBGPConfigurationTemplates`: NetworkDeviceBGPConfigurationTemplatePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPConfigurationTemplateAPI.GetNetworkDeviceBGPConfigurationTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceBGPConfigurationTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterAction** | **[]string** | Filter by action query param.  **Format:** filter.action&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.action&#x3D;$eq:John Doe&amp;filter.action&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkType** | **[]string** | Filter by networkType query param.  **Format:** filter.networkType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkType&#x3D;$eq:John Doe&amp;filter.networkType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkDeviceDriver** | **[]string** | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver&#x3D;$eq:John Doe&amp;filter.networkDeviceDriver&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkDevicePosition** | **[]string** | Filter by networkDevicePosition query param.  **Format:** filter.networkDevicePosition&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDevicePosition&#x3D;$eq:John Doe&amp;filter.networkDevicePosition&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterRemoteNetworkDevicePosition** | **[]string** | Filter by remoteNetworkDevicePosition query param.  **Format:** filter.remoteNetworkDevicePosition&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.remoteNetworkDevicePosition&#x3D;$eq:John Doe&amp;filter.remoteNetworkDevicePosition&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterBgpNumbering** | **[]string** | Filter by bgpNumbering query param.  **Format:** filter.bgpNumbering&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.bgpNumbering&#x3D;$eq:John Doe&amp;filter.bgpNumbering&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterBgpLinkConfiguration** | **[]string** | Filter by bgpLinkConfiguration query param.  **Format:** filter.bgpLinkConfiguration&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.bgpLinkConfiguration&#x3D;$eq:John Doe&amp;filter.bgpLinkConfiguration&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterExecutionType** | **[]string** | Filter by executionType query param.  **Format:** filter.executionType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.executionType&#x3D;$eq:John Doe&amp;filter.executionType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterLibraryLabel** | **[]string** | Filter by libraryLabel query param.  **Format:** filter.libraryLabel&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.libraryLabel&#x3D;$eq:John Doe&amp;filter.libraryLabel&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;action:DESC   **Default Value:** id:DESC  **Available Fields** - id  - action  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,action,networkType,networkDeviceDriver,networkDevicePosition   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - action  - networkType  - networkDeviceDriver  - networkDevicePosition  - remoteNetworkDevicePosition  - bgpNumbering  - bgpLinkConfiguration  - executionType  - libraryLabel  | 

### Return type

[**NetworkDeviceBGPConfigurationTemplatePaginatedList**](NetworkDeviceBGPConfigurationTemplatePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDeviceBGPConfigurationTemplate

> NetworkDeviceBGPConfigurationTemplate UpdateNetworkDeviceBGPConfigurationTemplate(ctx, networkDeviceBGPConfigurationTemplateId).UpdateNetworkDeviceBGPConfigurationTemplate(updateNetworkDeviceBGPConfigurationTemplate).Execute()

Updates Network Device BGP Configuration Template information



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
	networkDeviceBGPConfigurationTemplateId := float32(8.14) // float32 | 
	updateNetworkDeviceBGPConfigurationTemplate := *openapiclient.NewUpdateNetworkDeviceBGPConfigurationTemplate() // UpdateNetworkDeviceBGPConfigurationTemplate | The Network Device BGP Configuration Template update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPConfigurationTemplateAPI.UpdateNetworkDeviceBGPConfigurationTemplate(context.Background(), networkDeviceBGPConfigurationTemplateId).UpdateNetworkDeviceBGPConfigurationTemplate(updateNetworkDeviceBGPConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPConfigurationTemplateAPI.UpdateNetworkDeviceBGPConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceBGPConfigurationTemplate`: NetworkDeviceBGPConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPConfigurationTemplateAPI.UpdateNetworkDeviceBGPConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceBGPConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceBGPConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDeviceBGPConfigurationTemplate** | [**UpdateNetworkDeviceBGPConfigurationTemplate**](UpdateNetworkDeviceBGPConfigurationTemplate.md) | The Network Device BGP Configuration Template update object | 

### Return type

[**NetworkDeviceBGPConfigurationTemplate**](NetworkDeviceBGPConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

