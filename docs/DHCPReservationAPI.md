# \DHCPReservationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDhcpReservation**](DHCPReservationAPI.md#CreateDhcpReservation) | **Post** /api/v2/sites/{siteId}/dhcp/{ipVersion}/reservations | Create a DHCP reservation
[**DeleteDhcpReservation**](DHCPReservationAPI.md#DeleteDhcpReservation) | **Delete** /api/v2/sites/{siteId}/dhcp/{ipVersion}/reservations/{id} | Delete a DHCP reservation
[**GetDhcpReservation**](DHCPReservationAPI.md#GetDhcpReservation) | **Get** /api/v2/sites/{siteId}/dhcp/{ipVersion}/reservations/{id} | Get a DHCP reservation
[**GetDhcpReservations**](DHCPReservationAPI.md#GetDhcpReservations) | **Get** /api/v2/sites/{siteId}/dhcp/{ipVersion}/reservations | List DHCP reservations for a site
[**UpdateDhcpReservation**](DHCPReservationAPI.md#UpdateDhcpReservation) | **Put** /api/v2/sites/{siteId}/dhcp/{ipVersion}/reservations/{id} | Update a DHCP reservation



## CreateDhcpReservation

> DhcpReservation CreateDhcpReservation(ctx, siteId, ipVersion).DhcpReservationCreate(dhcpReservationCreate).Execute()

Create a DHCP reservation

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
	ipVersion := "ipVersion_example" // string | 
	dhcpReservationCreate := *openapiclient.NewDhcpReservationCreate(openapiclient.DhcpReservation_allocation{DhcpAutoAllocation: openapiclient.NewDhcpAutoAllocation("auto", []float32{float32(123)})}) // DhcpReservationCreate | The DHCP reservation to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.CreateDhcpReservation(context.Background(), siteId, ipVersion).DhcpReservationCreate(dhcpReservationCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.CreateDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDhcpReservation`: DhcpReservation
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.CreateDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 
**ipVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpReservationCreate** | [**DhcpReservationCreate**](DhcpReservationCreate.md) | The DHCP reservation to create | 

### Return type

[**DhcpReservation**](DhcpReservation.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDhcpReservation

> DeleteDhcpReservation(ctx, siteId, ipVersion, id).IfMatch(ifMatch).Execute()

Delete a DHCP reservation

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
	ipVersion := "ipVersion_example" // string | 
	id := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DHCPReservationAPI.DeleteDhcpReservation(context.Background(), siteId, ipVersion, id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.DeleteDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 
**ipVersion** | **string** |  | 
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDhcpReservationRequest struct via the builder pattern


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


## GetDhcpReservation

> DhcpReservation GetDhcpReservation(ctx, siteId, ipVersion, id).Execute()

Get a DHCP reservation

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
	ipVersion := "ipVersion_example" // string | 
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.GetDhcpReservation(context.Background(), siteId, ipVersion, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.GetDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpReservation`: DhcpReservation
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.GetDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 
**ipVersion** | **string** |  | 
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DhcpReservation**](DhcpReservation.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDhcpReservations

> DhcpReservationPaginatedList GetDhcpReservations(ctx, siteId, ipVersion).Page(page).Limit(limit).FilterId(filterId).FilterMacAddress(filterMacAddress).FilterCircuitId(filterCircuitId).FilterDeviceType(filterDeviceType).FilterVendor(filterVendor).FilterGiaddr(filterGiaddr).FilterPriority(filterPriority).FilterPlannedDeviceId(filterPlannedDeviceId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List DHCP reservations for a site

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
	ipVersion := "ipVersion_example" // string | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterMacAddress := []string{"Inner_example"} // []string | Filter by macAddress query param.  **Format:** filter.macAddress={$not}:OPERATION:VALUE    **Example:** filter.macAddress=$eq:John Doe&filter.macAddress=$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $null  - $and  - $or (optional)
	filterCircuitId := []string{"Inner_example"} // []string | Filter by circuitId query param.  **Format:** filter.circuitId={$not}:OPERATION:VALUE    **Example:** filter.circuitId=$eq:John Doe&filter.circuitId=$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $null  - $and  - $or (optional)
	filterDeviceType := []string{"Inner_example"} // []string | Filter by deviceType query param.  **Format:** filter.deviceType={$not}:OPERATION:VALUE    **Example:** filter.deviceType=$eq:John Doe&filter.deviceType=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	filterVendor := []string{"Inner_example"} // []string | Filter by vendor query param.  **Format:** filter.vendor={$not}:OPERATION:VALUE    **Example:** filter.vendor=$eq:John Doe&filter.vendor=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	filterGiaddr := []string{"Inner_example"} // []string | Filter by giaddr query param.  **Format:** filter.giaddr={$not}:OPERATION:VALUE    **Example:** filter.giaddr=$eq:John Doe&filter.giaddr=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	filterPriority := []string{"Inner_example"} // []string | Filter by priority query param.  **Format:** filter.priority={$not}:OPERATION:VALUE    **Example:** filter.priority=$eq:John Doe&filter.priority=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	filterPlannedDeviceId := []string{"Inner_example"} // []string | Filter by plannedDeviceId query param.  **Format:** filter.plannedDeviceId={$not}:OPERATION:VALUE    **Example:** filter.plannedDeviceId=$eq:John Doe&filter.plannedDeviceId=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=priority:DESC   **Default Value:** id:ASC  **Available Fields** - id  - priority  - macAddress  - circuitId  - deviceType  - vendor  - giaddr  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** macAddress,circuitId,giaddr   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - macAddress  - circuitId  - giaddr  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.GetDhcpReservations(context.Background(), siteId, ipVersion).Page(page).Limit(limit).FilterId(filterId).FilterMacAddress(filterMacAddress).FilterCircuitId(filterCircuitId).FilterDeviceType(filterDeviceType).FilterVendor(filterVendor).FilterGiaddr(filterGiaddr).FilterPriority(filterPriority).FilterPlannedDeviceId(filterPlannedDeviceId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.GetDhcpReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpReservations`: DhcpReservationPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.GetDhcpReservations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 
**ipVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterMacAddress** | **[]string** | Filter by macAddress query param.  **Format:** filter.macAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.macAddress&#x3D;$eq:John Doe&amp;filter.macAddress&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $null  - $and  - $or | 
 **filterCircuitId** | **[]string** | Filter by circuitId query param.  **Format:** filter.circuitId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.circuitId&#x3D;$eq:John Doe&amp;filter.circuitId&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $null  - $and  - $or | 
 **filterDeviceType** | **[]string** | Filter by deviceType query param.  **Format:** filter.deviceType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.deviceType&#x3D;$eq:John Doe&amp;filter.deviceType&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **filterVendor** | **[]string** | Filter by vendor query param.  **Format:** filter.vendor&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vendor&#x3D;$eq:John Doe&amp;filter.vendor&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **filterGiaddr** | **[]string** | Filter by giaddr query param.  **Format:** filter.giaddr&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.giaddr&#x3D;$eq:John Doe&amp;filter.giaddr&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **filterPriority** | **[]string** | Filter by priority query param.  **Format:** filter.priority&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.priority&#x3D;$eq:John Doe&amp;filter.priority&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **filterPlannedDeviceId** | **[]string** | Filter by plannedDeviceId query param.  **Format:** filter.plannedDeviceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.plannedDeviceId&#x3D;$eq:John Doe&amp;filter.plannedDeviceId&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;priority:DESC   **Default Value:** id:ASC  **Available Fields** - id  - priority  - macAddress  - circuitId  - deviceType  - vendor  - giaddr  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** macAddress,circuitId,giaddr   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - macAddress  - circuitId  - giaddr  | 

### Return type

[**DhcpReservationPaginatedList**](DhcpReservationPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDhcpReservation

> DhcpReservation UpdateDhcpReservation(ctx, siteId, ipVersion, id).DhcpReservationUpdate(dhcpReservationUpdate).IfMatch(ifMatch).Execute()

Update a DHCP reservation

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
	ipVersion := "ipVersion_example" // string | 
	id := int64(789) // int64 | 
	dhcpReservationUpdate := *openapiclient.NewDhcpReservationUpdate() // DhcpReservationUpdate | The DHCP reservation update
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.UpdateDhcpReservation(context.Background(), siteId, ipVersion, id).DhcpReservationUpdate(dhcpReservationUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.UpdateDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDhcpReservation`: DhcpReservation
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.UpdateDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 
**ipVersion** | **string** |  | 
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dhcpReservationUpdate** | [**DhcpReservationUpdate**](DhcpReservationUpdate.md) | The DHCP reservation update | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**DhcpReservation**](DhcpReservation.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

