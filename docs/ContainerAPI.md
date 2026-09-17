# \ContainerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainer**](ContainerAPI.md#GetContainer) | **Get** /api/v2/containers/{containerId} | Retrieves the Container information
[**GetContainerPowerStatus**](ContainerAPI.md#GetContainerPowerStatus) | **Get** /api/v2/containers/{containerId}/power-status | Retrieves the power status of the Container
[**GetContainerRemoteConsoleInfo**](ContainerAPI.md#GetContainerRemoteConsoleInfo) | **Get** /api/v2/containers/{containerId}/remote-console-info | Get Remote Console information
[**GetContainers**](ContainerAPI.md#GetContainers) | **Get** /api/v2/containers | Get a list of Containers
[**RebootContainer**](ContainerAPI.md#RebootContainer) | **Post** /api/v2/containers/{containerId}/reboot | Reboots the Container
[**ShutdownContainer**](ContainerAPI.md#ShutdownContainer) | **Post** /api/v2/containers/{containerId}/shutdown | Shuts down the Container
[**StartContainer**](ContainerAPI.md#StartContainer) | **Post** /api/v2/containers/{containerId}/start | Starts the Container
[**UpdateContainer**](ContainerAPI.md#UpdateContainer) | **Patch** /api/v2/containers/{containerId} | Updates Container information



## GetContainer

> Container GetContainer(ctx, containerId).Execute()

Retrieves the Container information



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
	containerId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPI.GetContainer(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.GetContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainer`: Container
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPI.GetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Container**](Container.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerPowerStatus

> string GetContainerPowerStatus(ctx, containerId).Execute()

Retrieves the power status of the Container



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
	containerId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPI.GetContainerPowerStatus(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.GetContainerPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerPowerStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPI.GetContainerPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerPowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerRemoteConsoleInfo

> RemoteConsoleInfo GetContainerRemoteConsoleInfo(ctx, containerId).Execute()

Get Remote Console information



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
	containerId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPI.GetContainerRemoteConsoleInfo(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.GetContainerRemoteConsoleInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerRemoteConsoleInfo`: RemoteConsoleInfo
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPI.GetContainerRemoteConsoleInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRemoteConsoleInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteConsoleInfo**](RemoteConsoleInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainers

> ContainerPaginatedList GetContainers(ctx).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterName(filterName).FilterAddress(filterAddress).FilterHost(filterHost).FilterHosts(filterHosts).FilterTypeId(filterTypeId).FilterPoolId(filterPoolId).FilterAdministrationState(filterAdministrationState).FilterNumaNodes(filterNumaNodes).FilterInfrastructureId(filterInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Containers



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
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAddress := []string{"Inner_example"} // []string | Filter by address query param.  **Format:** filter.address={$not}:OPERATION:VALUE    **Example:** filter.address=$btw:John Doe&filter.address=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHost := []string{"Inner_example"} // []string | Filter by host query param.  **Format:** filter.host={$not}:OPERATION:VALUE    **Example:** filter.host=$btw:John Doe&filter.host=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterHosts := []string{"Inner_example"} // []string | Filter by hosts query param.  **Format:** filter.hosts={$not}:OPERATION:VALUE    **Example:** filter.hosts=$btw:John Doe&filter.hosts=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterTypeId := []string{"Inner_example"} // []string | Filter by typeId query param.  **Format:** filter.typeId={$not}:OPERATION:VALUE    **Example:** filter.typeId=$btw:John Doe&filter.typeId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterPoolId := []string{"Inner_example"} // []string | Filter by poolId query param.  **Format:** filter.poolId={$not}:OPERATION:VALUE    **Example:** filter.poolId=$btw:John Doe&filter.poolId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAdministrationState := []string{"Inner_example"} // []string | Filter by administrationState query param.  **Format:** filter.administrationState={$not}:OPERATION:VALUE    **Example:** filter.administrationState=$btw:John Doe&filter.administrationState=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNumaNodes := []string{"Inner_example"} // []string | Filter by numaNodes query param.  **Format:** filter.numaNodes={$not}:OPERATION:VALUE    **Example:** filter.numaNodes=$btw:John Doe&filter.numaNodes=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$btw:John Doe&filter.infrastructureId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - host  - administrationState  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,host,hosts,administrationState   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - host  - hosts  - administrationState  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPI.GetContainers(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterName(filterName).FilterAddress(filterAddress).FilterHost(filterHost).FilterHosts(filterHosts).FilterTypeId(filterTypeId).FilterPoolId(filterPoolId).FilterAdministrationState(filterAdministrationState).FilterNumaNodes(filterNumaNodes).FilterInfrastructureId(filterInfrastructureId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.GetContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainers`: ContainerPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPI.GetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAddress** | **[]string** | Filter by address query param.  **Format:** filter.address&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.address&#x3D;$btw:John Doe&amp;filter.address&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHost** | **[]string** | Filter by host query param.  **Format:** filter.host&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.host&#x3D;$btw:John Doe&amp;filter.host&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterHosts** | **[]string** | Filter by hosts query param.  **Format:** filter.hosts&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.hosts&#x3D;$btw:John Doe&amp;filter.hosts&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterTypeId** | **[]string** | Filter by typeId query param.  **Format:** filter.typeId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.typeId&#x3D;$btw:John Doe&amp;filter.typeId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterPoolId** | **[]string** | Filter by poolId query param.  **Format:** filter.poolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.poolId&#x3D;$btw:John Doe&amp;filter.poolId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAdministrationState** | **[]string** | Filter by administrationState query param.  **Format:** filter.administrationState&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.administrationState&#x3D;$btw:John Doe&amp;filter.administrationState&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNumaNodes** | **[]string** | Filter by numaNodes query param.  **Format:** filter.numaNodes&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.numaNodes&#x3D;$btw:John Doe&amp;filter.numaNodes&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$btw:John Doe&amp;filter.infrastructureId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  - host  - administrationState  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,host,hosts,administrationState   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - host  - hosts  - administrationState  | 

### Return type

[**ContainerPaginatedList**](ContainerPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootContainer

> RebootContainer(ctx, containerId).Execute()

Reboots the Container



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
	containerId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerAPI.RebootContainer(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.RebootContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootContainerRequest struct via the builder pattern


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


## ShutdownContainer

> ShutdownContainer(ctx, containerId).Execute()

Shuts down the Container



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
	containerId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerAPI.ShutdownContainer(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.ShutdownContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownContainerRequest struct via the builder pattern


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


## StartContainer

> StartContainer(ctx, containerId).Execute()

Starts the Container



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
	containerId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerAPI.StartContainer(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.StartContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartContainerRequest struct via the builder pattern


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


## UpdateContainer

> Container UpdateContainer(ctx, containerId).UpdateContainer(updateContainer).Execute()

Updates Container information



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
	containerId := int64(789) // int64 | 
	updateContainer := *openapiclient.NewUpdateContainer() // UpdateContainer | The Container update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerAPI.UpdateContainer(context.Background(), containerId).UpdateContainer(updateContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerAPI.UpdateContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContainer`: Container
	fmt.Fprintf(os.Stdout, "Response from `ContainerAPI.UpdateContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContainer** | [**UpdateContainer**](UpdateContainer.md) | The Container update object | 

### Return type

[**Container**](Container.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

