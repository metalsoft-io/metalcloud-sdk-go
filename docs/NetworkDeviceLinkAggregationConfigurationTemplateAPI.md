# \NetworkDeviceLinkAggregationConfigurationTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDeviceLinkAggregationConfigurationTemplate**](NetworkDeviceLinkAggregationConfigurationTemplateAPI.md#CreateNetworkDeviceLinkAggregationConfigurationTemplate) | **Post** /api/v2/network-device-link-aggregation-configuration-templates | Creates a Network Device Link Aggregation Configuration Template
[**DeleteNetworkDeviceLinkAggregationConfigurationTemplate**](NetworkDeviceLinkAggregationConfigurationTemplateAPI.md#DeleteNetworkDeviceLinkAggregationConfigurationTemplate) | **Delete** /api/v2/network-device-link-aggregation-configuration-templates/{networkDeviceLinkAggregationConfigurationTemplateId} | Deletes a Network Device Link Aggregation Configuration Template
[**GetNetworkDeviceLinkAggregationConfigurationTemplate**](NetworkDeviceLinkAggregationConfigurationTemplateAPI.md#GetNetworkDeviceLinkAggregationConfigurationTemplate) | **Get** /api/v2/network-device-link-aggregation-configuration-templates/{networkDeviceLinkAggregationConfigurationTemplateId} | Get Network Device Link Aggregation Configuration Template information
[**GetNetworkDeviceLinkAggregationConfigurationTemplates**](NetworkDeviceLinkAggregationConfigurationTemplateAPI.md#GetNetworkDeviceLinkAggregationConfigurationTemplates) | **Get** /api/v2/network-device-link-aggregation-configuration-templates | Get all Network Device Link Aggregation Configuration Templates
[**UpdateNetworkDeviceLinkAggregationConfigurationTemplate**](NetworkDeviceLinkAggregationConfigurationTemplateAPI.md#UpdateNetworkDeviceLinkAggregationConfigurationTemplate) | **Patch** /api/v2/network-device-link-aggregation-configuration-templates/{networkDeviceLinkAggregationConfigurationTemplateId} | Updates Network Device Link Aggregation Configuration Template information



## CreateNetworkDeviceLinkAggregationConfigurationTemplate

> NetworkDeviceLinkAggregationConfigurationTemplate CreateNetworkDeviceLinkAggregationConfigurationTemplate(ctx).CreateNetworkDeviceLinkAggregationConfigurationTemplate(createNetworkDeviceLinkAggregationConfigurationTemplate).Execute()

Creates a Network Device Link Aggregation Configuration Template



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
	createNetworkDeviceLinkAggregationConfigurationTemplate := *openapiclient.NewCreateNetworkDeviceLinkAggregationConfigurationTemplate("Action_example", "AggregationType_example", "NetworkDeviceDriver_example", "ExecutionType_example", "LibraryLabel_example", "Configuration_example") // CreateNetworkDeviceLinkAggregationConfigurationTemplate | The Network Device Link Aggregation Configuration Template create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceLinkAggregationConfigurationTemplateAPI.CreateNetworkDeviceLinkAggregationConfigurationTemplate(context.Background()).CreateNetworkDeviceLinkAggregationConfigurationTemplate(createNetworkDeviceLinkAggregationConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceLinkAggregationConfigurationTemplateAPI.CreateNetworkDeviceLinkAggregationConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDeviceLinkAggregationConfigurationTemplate`: NetworkDeviceLinkAggregationConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceLinkAggregationConfigurationTemplateAPI.CreateNetworkDeviceLinkAggregationConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceLinkAggregationConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDeviceLinkAggregationConfigurationTemplate** | [**CreateNetworkDeviceLinkAggregationConfigurationTemplate**](CreateNetworkDeviceLinkAggregationConfigurationTemplate.md) | The Network Device Link Aggregation Configuration Template create object | 

### Return type

[**NetworkDeviceLinkAggregationConfigurationTemplate**](NetworkDeviceLinkAggregationConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDeviceLinkAggregationConfigurationTemplate

> DeleteNetworkDeviceLinkAggregationConfigurationTemplate(ctx, networkDeviceLinkAggregationConfigurationTemplateId).Execute()

Deletes a Network Device Link Aggregation Configuration Template



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
	networkDeviceLinkAggregationConfigurationTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceLinkAggregationConfigurationTemplateAPI.DeleteNetworkDeviceLinkAggregationConfigurationTemplate(context.Background(), networkDeviceLinkAggregationConfigurationTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceLinkAggregationConfigurationTemplateAPI.DeleteNetworkDeviceLinkAggregationConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceLinkAggregationConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceLinkAggregationConfigurationTemplateRequest struct via the builder pattern


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


## GetNetworkDeviceLinkAggregationConfigurationTemplate

> NetworkDeviceLinkAggregationConfigurationTemplate GetNetworkDeviceLinkAggregationConfigurationTemplate(ctx, networkDeviceLinkAggregationConfigurationTemplateId).Execute()

Get Network Device Link Aggregation Configuration Template information



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
	networkDeviceLinkAggregationConfigurationTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceLinkAggregationConfigurationTemplateAPI.GetNetworkDeviceLinkAggregationConfigurationTemplate(context.Background(), networkDeviceLinkAggregationConfigurationTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceLinkAggregationConfigurationTemplateAPI.GetNetworkDeviceLinkAggregationConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceLinkAggregationConfigurationTemplate`: NetworkDeviceLinkAggregationConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceLinkAggregationConfigurationTemplateAPI.GetNetworkDeviceLinkAggregationConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceLinkAggregationConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceLinkAggregationConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceLinkAggregationConfigurationTemplate**](NetworkDeviceLinkAggregationConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceLinkAggregationConfigurationTemplates

> NetworkDeviceLinkAggregationConfigurationTemplatePaginatedList GetNetworkDeviceLinkAggregationConfigurationTemplates(ctx).Page(page).Limit(limit).FilterId(filterId).FilterAction(filterAction).FilterAggregationType(filterAggregationType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterExecutionType(filterExecutionType).FilterLibraryLabel(filterLibraryLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Network Device Link Aggregation Configuration Templates



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
	filterAggregationType := []string{"Inner_example"} // []string | Filter by aggregationType query param.  **Format:** filter.aggregationType={$not}:OPERATION:VALUE    **Example:** filter.aggregationType=$eq:John Doe&filter.aggregationType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkDeviceDriver := []string{"Inner_example"} // []string | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver=$eq:John Doe&filter.networkDeviceDriver=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterExecutionType := []string{"Inner_example"} // []string | Filter by executionType query param.  **Format:** filter.executionType={$not}:OPERATION:VALUE    **Example:** filter.executionType=$eq:John Doe&filter.executionType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterLibraryLabel := []string{"Inner_example"} // []string | Filter by libraryLabel query param.  **Format:** filter.libraryLabel={$not}:OPERATION:VALUE    **Example:** filter.libraryLabel=$eq:John Doe&filter.libraryLabel=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=action:DESC   **Default Value:** id:DESC  **Available Fields** - id  - action  - aggregationType  - networkDeviceDriver  - executionType  - libraryLabel  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,action,aggregationType,networkDeviceDriver,executionType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - action  - aggregationType  - networkDeviceDriver  - executionType  - libraryLabel  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceLinkAggregationConfigurationTemplateAPI.GetNetworkDeviceLinkAggregationConfigurationTemplates(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterAction(filterAction).FilterAggregationType(filterAggregationType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterExecutionType(filterExecutionType).FilterLibraryLabel(filterLibraryLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceLinkAggregationConfigurationTemplateAPI.GetNetworkDeviceLinkAggregationConfigurationTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceLinkAggregationConfigurationTemplates`: NetworkDeviceLinkAggregationConfigurationTemplatePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceLinkAggregationConfigurationTemplateAPI.GetNetworkDeviceLinkAggregationConfigurationTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceLinkAggregationConfigurationTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterAction** | **[]string** | Filter by action query param.  **Format:** filter.action&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.action&#x3D;$eq:John Doe&amp;filter.action&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterAggregationType** | **[]string** | Filter by aggregationType query param.  **Format:** filter.aggregationType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.aggregationType&#x3D;$eq:John Doe&amp;filter.aggregationType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkDeviceDriver** | **[]string** | Filter by networkDeviceDriver query param.  **Format:** filter.networkDeviceDriver&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceDriver&#x3D;$eq:John Doe&amp;filter.networkDeviceDriver&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterExecutionType** | **[]string** | Filter by executionType query param.  **Format:** filter.executionType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.executionType&#x3D;$eq:John Doe&amp;filter.executionType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterLibraryLabel** | **[]string** | Filter by libraryLabel query param.  **Format:** filter.libraryLabel&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.libraryLabel&#x3D;$eq:John Doe&amp;filter.libraryLabel&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;action:DESC   **Default Value:** id:DESC  **Available Fields** - id  - action  - aggregationType  - networkDeviceDriver  - executionType  - libraryLabel  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,action,aggregationType,networkDeviceDriver,executionType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - action  - aggregationType  - networkDeviceDriver  - executionType  - libraryLabel  | 

### Return type

[**NetworkDeviceLinkAggregationConfigurationTemplatePaginatedList**](NetworkDeviceLinkAggregationConfigurationTemplatePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDeviceLinkAggregationConfigurationTemplate

> NetworkDeviceLinkAggregationConfigurationTemplate UpdateNetworkDeviceLinkAggregationConfigurationTemplate(ctx, networkDeviceLinkAggregationConfigurationTemplateId).UpdateNetworkDeviceLinkAggregationConfigurationTemplate(updateNetworkDeviceLinkAggregationConfigurationTemplate).Execute()

Updates Network Device Link Aggregation Configuration Template information



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
	networkDeviceLinkAggregationConfigurationTemplateId := float32(8.14) // float32 | 
	updateNetworkDeviceLinkAggregationConfigurationTemplate := *openapiclient.NewUpdateNetworkDeviceLinkAggregationConfigurationTemplate() // UpdateNetworkDeviceLinkAggregationConfigurationTemplate | The Network Device Link Aggregation Configuration Template update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceLinkAggregationConfigurationTemplateAPI.UpdateNetworkDeviceLinkAggregationConfigurationTemplate(context.Background(), networkDeviceLinkAggregationConfigurationTemplateId).UpdateNetworkDeviceLinkAggregationConfigurationTemplate(updateNetworkDeviceLinkAggregationConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceLinkAggregationConfigurationTemplateAPI.UpdateNetworkDeviceLinkAggregationConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceLinkAggregationConfigurationTemplate`: NetworkDeviceLinkAggregationConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceLinkAggregationConfigurationTemplateAPI.UpdateNetworkDeviceLinkAggregationConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceLinkAggregationConfigurationTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceLinkAggregationConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDeviceLinkAggregationConfigurationTemplate** | [**UpdateNetworkDeviceLinkAggregationConfigurationTemplate**](UpdateNetworkDeviceLinkAggregationConfigurationTemplate.md) | The Network Device Link Aggregation Configuration Template update object | 

### Return type

[**NetworkDeviceLinkAggregationConfigurationTemplate**](NetworkDeviceLinkAggregationConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

