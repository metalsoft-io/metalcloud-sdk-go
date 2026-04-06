# \NetworkDeviceBGPInterconnectConfigurationTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplateAPI.md#CreateNetworkDeviceBGPInterconnectConfigurationTemplate) | **Post** /api/v2/network-device-bgp-interconnect-configuration-templates | Creates a Network Device BGP Interconnect Configuration Template
[**DeleteNetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplateAPI.md#DeleteNetworkDeviceBGPInterconnectConfigurationTemplate) | **Delete** /api/v2/network-device-bgp-interconnect-configuration-templates/{id} | Deletes a Network Device BGP Interconnect Configuration Template
[**GetNetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplateAPI.md#GetNetworkDeviceBGPInterconnectConfigurationTemplate) | **Get** /api/v2/network-device-bgp-interconnect-configuration-templates/{id} | Get Network Device BGP Interconnect Configuration Template information
[**GetNetworkDeviceBGPInterconnectConfigurationTemplates**](NetworkDeviceBGPInterconnectConfigurationTemplateAPI.md#GetNetworkDeviceBGPInterconnectConfigurationTemplates) | **Get** /api/v2/network-device-bgp-interconnect-configuration-templates | Get all Network Device BGP Interconnect Configuration Templates
[**UpdateNetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplateAPI.md#UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) | **Patch** /api/v2/network-device-bgp-interconnect-configuration-templates/{id} | Updates Network Device BGP Interconnect Configuration Template information



## CreateNetworkDeviceBGPInterconnectConfigurationTemplate

> NetworkDeviceBGPInterconnectConfigurationTemplate CreateNetworkDeviceBGPInterconnectConfigurationTemplate(ctx).CreateNetworkDeviceBGPInterconnectConfigurationTemplate(createNetworkDeviceBGPInterconnectConfigurationTemplate).Execute()

Creates a Network Device BGP Interconnect Configuration Template



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
	createNetworkDeviceBGPInterconnectConfigurationTemplate := *openapiclient.NewCreateNetworkDeviceBGPInterconnectConfigurationTemplate("Label_example", "Name_example", "sonic_enterprise", "cli") // CreateNetworkDeviceBGPInterconnectConfigurationTemplate | The Network Device BGP Interconnect Configuration Template create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPInterconnectConfigurationTemplateAPI.CreateNetworkDeviceBGPInterconnectConfigurationTemplate(context.Background()).CreateNetworkDeviceBGPInterconnectConfigurationTemplate(createNetworkDeviceBGPInterconnectConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.CreateNetworkDeviceBGPInterconnectConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDeviceBGPInterconnectConfigurationTemplate`: NetworkDeviceBGPInterconnectConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.CreateNetworkDeviceBGPInterconnectConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceBGPInterconnectConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDeviceBGPInterconnectConfigurationTemplate** | [**CreateNetworkDeviceBGPInterconnectConfigurationTemplate**](CreateNetworkDeviceBGPInterconnectConfigurationTemplate.md) | The Network Device BGP Interconnect Configuration Template create object | 

### Return type

[**NetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDeviceBGPInterconnectConfigurationTemplate

> DeleteNetworkDeviceBGPInterconnectConfigurationTemplate(ctx, id).IfMatch(ifMatch).Execute()

Deletes a Network Device BGP Interconnect Configuration Template



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
	id := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceBGPInterconnectConfigurationTemplateAPI.DeleteNetworkDeviceBGPInterconnectConfigurationTemplate(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.DeleteNetworkDeviceBGPInterconnectConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceBGPInterconnectConfigurationTemplateRequest struct via the builder pattern


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


## GetNetworkDeviceBGPInterconnectConfigurationTemplate

> NetworkDeviceBGPInterconnectConfigurationTemplate GetNetworkDeviceBGPInterconnectConfigurationTemplate(ctx, id).Execute()

Get Network Device BGP Interconnect Configuration Template information



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
	id := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPInterconnectConfigurationTemplateAPI.GetNetworkDeviceBGPInterconnectConfigurationTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.GetNetworkDeviceBGPInterconnectConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceBGPInterconnectConfigurationTemplate`: NetworkDeviceBGPInterconnectConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.GetNetworkDeviceBGPInterconnectConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceBGPInterconnectConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceBGPInterconnectConfigurationTemplates

> NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList GetNetworkDeviceBGPInterconnectConfigurationTemplates(ctx).Page(page).Limit(limit).FilterId(filterId).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterExecutionType(filterExecutionType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Network Device BGP Interconnect Configuration Templates



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
	filterNetworkDeviceDriver := []string{"Inner_example"} // []string | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver=$eq:John Doe&filter.networkDeviceDriver=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterExecutionType := []string{"Inner_example"} // []string | Filter by executionType query param.  **Format:** filter.executionType={$not}:OPERATION:VALUE    **Example:** filter.executionType=$eq:John Doe&filter.executionType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=networkDeviceDriver:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkDeviceDriver  - executionType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,networkDeviceDriver,executionType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - networkDeviceDriver  - executionType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPInterconnectConfigurationTemplateAPI.GetNetworkDeviceBGPInterconnectConfigurationTemplates(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterExecutionType(filterExecutionType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.GetNetworkDeviceBGPInterconnectConfigurationTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceBGPInterconnectConfigurationTemplates`: NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.GetNetworkDeviceBGPInterconnectConfigurationTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceBGPInterconnectConfigurationTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkDeviceDriver** | **[]string** | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver&#x3D;$eq:John Doe&amp;filter.networkDeviceDriver&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterExecutionType** | **[]string** | Filter by executionType query param.  **Format:** filter.executionType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.executionType&#x3D;$eq:John Doe&amp;filter.executionType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;networkDeviceDriver:DESC   **Default Value:** id:DESC  **Available Fields** - id  - networkDeviceDriver  - executionType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,networkDeviceDriver,executionType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - networkDeviceDriver  - executionType  | 

### Return type

[**NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList**](NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDeviceBGPInterconnectConfigurationTemplate

> NetworkDeviceBGPInterconnectConfigurationTemplate UpdateNetworkDeviceBGPInterconnectConfigurationTemplate(ctx, id).UpdateNetworkDeviceBGPInterconnectConfigurationTemplate(updateNetworkDeviceBGPInterconnectConfigurationTemplate).IfMatch(ifMatch).Execute()

Updates Network Device BGP Interconnect Configuration Template information



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
	id := float32(8.14) // float32 | 
	updateNetworkDeviceBGPInterconnectConfigurationTemplate := *openapiclient.NewUpdateNetworkDeviceBGPInterconnectConfigurationTemplate() // UpdateNetworkDeviceBGPInterconnectConfigurationTemplate | The Network Device BGP Interconnect Configuration Template update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceBGPInterconnectConfigurationTemplateAPI.UpdateNetworkDeviceBGPInterconnectConfigurationTemplate(context.Background(), id).UpdateNetworkDeviceBGPInterconnectConfigurationTemplate(updateNetworkDeviceBGPInterconnectConfigurationTemplate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.UpdateNetworkDeviceBGPInterconnectConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceBGPInterconnectConfigurationTemplate`: NetworkDeviceBGPInterconnectConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceBGPInterconnectConfigurationTemplateAPI.UpdateNetworkDeviceBGPInterconnectConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceBGPInterconnectConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDeviceBGPInterconnectConfigurationTemplate** | [**UpdateNetworkDeviceBGPInterconnectConfigurationTemplate**](UpdateNetworkDeviceBGPInterconnectConfigurationTemplate.md) | The Network Device BGP Interconnect Configuration Template update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

