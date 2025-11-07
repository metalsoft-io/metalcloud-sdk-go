# \EventAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventAPI.md#GetEvent) | **Get** /api/v2/events/{eventId} | Get Event information
[**GetEvents**](EventAPI.md#GetEvents) | **Get** /api/v2/events | Get all Events



## GetEvent

> Event GetEvent(ctx, eventId).Execute()

Get Event information



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
	eventId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.GetEvent(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> EventPaginatedList GetEvents(ctx).Page(page).Limit(limit).FilterId(filterId).FilterUserIdAuthenticated(filterUserIdAuthenticated).FilterType(filterType).FilterSeverity(filterSeverity).FilterVisibility(filterVisibility).FilterInfrastructureId(filterInfrastructureId).FilterUserId(filterUserId).FilterStoragePoolId(filterStoragePoolId).FilterServerId(filterServerId).FilterJobId(filterJobId).FilterSiteId(filterSiteId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Events



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterUserIdAuthenticated := []string{"Inner_example"} // []string | Filter by userIdAuthenticated query param.  **Format:** filter.userIdAuthenticated={$not}:OPERATION:VALUE    **Example:** filter.userIdAuthenticated=$eq:John Doe&filter.userIdAuthenticated=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterType := []string{"Inner_example"} // []string | Filter by type query param.  **Format:** filter.type={$not}:OPERATION:VALUE    **Example:** filter.type=$eq:John Doe&filter.type=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterSeverity := []string{"Inner_example"} // []string | Filter by severity query param.  **Format:** filter.severity={$not}:OPERATION:VALUE    **Example:** filter.severity=$eq:John Doe&filter.severity=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterVisibility := []string{"Inner_example"} // []string | Filter by visibility query param.  **Format:** filter.visibility={$not}:OPERATION:VALUE    **Example:** filter.visibility=$eq:John Doe&filter.visibility=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$eq:John Doe&filter.infrastructureId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterUserId := []string{"Inner_example"} // []string | Filter by userId query param.  **Format:** filter.userId={$not}:OPERATION:VALUE    **Example:** filter.userId=$eq:John Doe&filter.userId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.  **Format:** filter.storagePoolId={$not}:OPERATION:VALUE    **Example:** filter.storagePoolId=$eq:John Doe&filter.storagePoolId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$eq:John Doe&filter.serverId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterJobId := []string{"Inner_example"} // []string | Filter by jobId query param.  **Format:** filter.jobId={$not}:OPERATION:VALUE    **Example:** filter.jobId=$eq:John Doe&filter.jobId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$eq:John Doe&filter.siteId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:DESC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,type,severity,visibility,infrastructureId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - type  - severity  - visibility  - infrastructureId  - userId  - serverId  - jobId  - siteId  - title  - message  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.GetEvents(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterUserIdAuthenticated(filterUserIdAuthenticated).FilterType(filterType).FilterSeverity(filterSeverity).FilterVisibility(filterVisibility).FilterInfrastructureId(filterInfrastructureId).FilterUserId(filterUserId).FilterStoragePoolId(filterStoragePoolId).FilterServerId(filterServerId).FilterJobId(filterJobId).FilterSiteId(filterSiteId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: EventPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterUserIdAuthenticated** | **[]string** | Filter by userIdAuthenticated query param.  **Format:** filter.userIdAuthenticated&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.userIdAuthenticated&#x3D;$eq:John Doe&amp;filter.userIdAuthenticated&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterType** | **[]string** | Filter by type query param.  **Format:** filter.type&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.type&#x3D;$eq:John Doe&amp;filter.type&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterSeverity** | **[]string** | Filter by severity query param.  **Format:** filter.severity&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.severity&#x3D;$eq:John Doe&amp;filter.severity&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterVisibility** | **[]string** | Filter by visibility query param.  **Format:** filter.visibility&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.visibility&#x3D;$eq:John Doe&amp;filter.visibility&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$eq:John Doe&amp;filter.infrastructureId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterUserId** | **[]string** | Filter by userId query param.  **Format:** filter.userId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.userId&#x3D;$eq:John Doe&amp;filter.userId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.  **Format:** filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.storagePoolId&#x3D;$eq:John Doe&amp;filter.storagePoolId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$eq:John Doe&amp;filter.serverId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterJobId** | **[]string** | Filter by jobId query param.  **Format:** filter.jobId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobId&#x3D;$eq:John Doe&amp;filter.jobId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$eq:John Doe&amp;filter.siteId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:DESC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,type,severity,visibility,infrastructureId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - type  - severity  - visibility  - infrastructureId  - userId  - serverId  - jobId  - siteId  - title  - message  | 

### Return type

[**EventPaginatedList**](EventPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

