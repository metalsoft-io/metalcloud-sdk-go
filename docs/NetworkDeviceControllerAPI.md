# \NetworkDeviceControllerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDeviceController**](NetworkDeviceControllerAPI.md#CreateNetworkDeviceController) | **Post** /api/v2/network-device-controllers | Create Network Device Controller
[**DeleteNetworkDeviceController**](NetworkDeviceControllerAPI.md#DeleteNetworkDeviceController) | **Delete** /api/v2/network-device-controllers/{networkDeviceControllerId} | Delete Network Device Controller
[**GetNetworkDeviceController**](NetworkDeviceControllerAPI.md#GetNetworkDeviceController) | **Get** /api/v2/network-device-controllers/{networkDeviceControllerId} | Get Network Device Controller
[**GetNetworkDeviceControllerCredentials**](NetworkDeviceControllerAPI.md#GetNetworkDeviceControllerCredentials) | **Get** /api/v2/network-device-controllers/{networkDeviceControllerId}/credentials | Get Network Device Controller credentials
[**GetNetworkDeviceControllers**](NetworkDeviceControllerAPI.md#GetNetworkDeviceControllers) | **Get** /api/v2/network-device-controllers | Get paginated Network Device Controllers
[**UpdateNetworkDeviceController**](NetworkDeviceControllerAPI.md#UpdateNetworkDeviceController) | **Patch** /api/v2/network-device-controllers/{networkDeviceControllerId} | Update Network Device Controller



## CreateNetworkDeviceController

> NetworkDeviceController CreateNetworkDeviceController(ctx).CreateNetworkDeviceController(createNetworkDeviceController).Execute()

Create Network Device Controller

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
	createNetworkDeviceController := *openapiclient.NewCreateNetworkDeviceController("DC-North", openapiclient.SwitchControllerDriver("cisco_aci51"), "192.168.1.10", int32(22), "admin", "ManagementPassword_example") // CreateNetworkDeviceController | The Network Device Controller create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceControllerAPI.CreateNetworkDeviceController(context.Background()).CreateNetworkDeviceController(createNetworkDeviceController).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceControllerAPI.CreateNetworkDeviceController``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDeviceController`: NetworkDeviceController
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceControllerAPI.CreateNetworkDeviceController`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDeviceController** | [**CreateNetworkDeviceController**](CreateNetworkDeviceController.md) | The Network Device Controller create object | 

### Return type

[**NetworkDeviceController**](NetworkDeviceController.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDeviceController

> DeleteNetworkDeviceController(ctx, networkDeviceControllerId).IfMatch(ifMatch).Execute()

Delete Network Device Controller

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
	networkDeviceControllerId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceControllerAPI.DeleteNetworkDeviceController(context.Background(), networkDeviceControllerId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceControllerAPI.DeleteNetworkDeviceController``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceControllerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceControllerRequest struct via the builder pattern


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


## GetNetworkDeviceController

> NetworkDeviceController GetNetworkDeviceController(ctx, networkDeviceControllerId).Execute()

Get Network Device Controller

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
	networkDeviceControllerId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceControllerAPI.GetNetworkDeviceController(context.Background(), networkDeviceControllerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceControllerAPI.GetNetworkDeviceController``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceController`: NetworkDeviceController
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceControllerAPI.GetNetworkDeviceController`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceControllerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceController**](NetworkDeviceController.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceControllerCredentials

> NetworkDeviceControllerCredentials GetNetworkDeviceControllerCredentials(ctx, networkDeviceControllerId).Execute()

Get Network Device Controller credentials



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
	networkDeviceControllerId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceControllerAPI.GetNetworkDeviceControllerCredentials(context.Background(), networkDeviceControllerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceControllerAPI.GetNetworkDeviceControllerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceControllerCredentials`: NetworkDeviceControllerCredentials
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceControllerAPI.GetNetworkDeviceControllerCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceControllerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceControllerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceControllerCredentials**](NetworkDeviceControllerCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceControllers

> NetworkDeviceControllerPaginatedList GetNetworkDeviceControllers(ctx).Page(page).Limit(limit).FilterId(filterId).FilterDatacenterName(filterDatacenterName).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get paginated Network Device Controllers

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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDatacenterName := []string{"Inner_example"} // []string | Filter by datacenterName query param.  **Format:** filter.datacenterName={$not}:OPERATION:VALUE    **Example:** filter.datacenterName=$btw:John Doe&filter.datacenterName=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementAddress := []string{"Inner_example"} // []string | Filter by managementAddress query param.  **Format:** filter.managementAddress={$not}:OPERATION:VALUE    **Example:** filter.managementAddress=$btw:John Doe&filter.managementAddress=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterManagementPort := []string{"Inner_example"} // []string | Filter by managementPort query param.  **Format:** filter.managementPort={$not}:OPERATION:VALUE    **Example:** filter.managementPort=$btw:John Doe&filter.managementPort=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterIdentifierString := []string{"Inner_example"} // []string | Filter by identifierString query param.  **Format:** filter.identifierString={$not}:OPERATION:VALUE    **Example:** filter.identifierString=$btw:John Doe&filter.identifierString=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** managementAddress,identifierString   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - managementAddress  - identifierString  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceControllerAPI.GetNetworkDeviceControllers(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterDatacenterName(filterDatacenterName).FilterManagementAddress(filterManagementAddress).FilterManagementPort(filterManagementPort).FilterIdentifierString(filterIdentifierString).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceControllerAPI.GetNetworkDeviceControllers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceControllers`: NetworkDeviceControllerPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceControllerAPI.GetNetworkDeviceControllers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceControllersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDatacenterName** | **[]string** | Filter by datacenterName query param.  **Format:** filter.datacenterName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.datacenterName&#x3D;$btw:John Doe&amp;filter.datacenterName&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementAddress** | **[]string** | Filter by managementAddress query param.  **Format:** filter.managementAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementAddress&#x3D;$btw:John Doe&amp;filter.managementAddress&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterManagementPort** | **[]string** | Filter by managementPort query param.  **Format:** filter.managementPort&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.managementPort&#x3D;$btw:John Doe&amp;filter.managementPort&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterIdentifierString** | **[]string** | Filter by identifierString query param.  **Format:** filter.identifierString&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.identifierString&#x3D;$btw:John Doe&amp;filter.identifierString&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** managementAddress,identifierString   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - managementAddress  - identifierString  | 

### Return type

[**NetworkDeviceControllerPaginatedList**](NetworkDeviceControllerPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDeviceController

> NetworkDeviceController UpdateNetworkDeviceController(ctx, networkDeviceControllerId).UpdateNetworkDeviceController(updateNetworkDeviceController).IfMatch(ifMatch).Execute()

Update Network Device Controller

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
	networkDeviceControllerId := int32(56) // int32 | 
	updateNetworkDeviceController := *openapiclient.NewUpdateNetworkDeviceController() // UpdateNetworkDeviceController | The Network Device update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceControllerAPI.UpdateNetworkDeviceController(context.Background(), networkDeviceControllerId).UpdateNetworkDeviceController(updateNetworkDeviceController).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceControllerAPI.UpdateNetworkDeviceController``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceController`: NetworkDeviceController
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceControllerAPI.UpdateNetworkDeviceController`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceControllerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDeviceController** | [**UpdateNetworkDeviceController**](UpdateNetworkDeviceController.md) | The Network Device update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**NetworkDeviceController**](NetworkDeviceController.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

