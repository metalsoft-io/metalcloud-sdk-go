# \ServerFirmwareAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchScheduleServerFirmwareUpgrade**](ServerFirmwareAPI.md#BatchScheduleServerFirmwareUpgrade) | **Post** /api/v2/servers/{serverId}/firmware/actions/batch-schedule-upgrade | Schedules a firmware upgrade for the specified Server
[**FetchAndUpdateServerFirmwareAvailableVersions**](ServerFirmwareAPI.md#FetchAndUpdateServerFirmwareAvailableVersions) | **Post** /api/v2/servers/{serverId}/firmware/actions/fetch-and-update-available-versions | Fetches and updates the available firmware versions for the specified Server
[**GenerateServersFirmwareUpgradeAudit**](ServerFirmwareAPI.md#GenerateServersFirmwareUpgradeAudit) | **Post** /api/v2/servers/firmware/actions/generate-upgrade-audit | Generates a Firmware Upgrade Audit for the specified Servers
[**GetServerComponentInfo**](ServerFirmwareAPI.md#GetServerComponentInfo) | **Get** /api/v2/servers/{serverId}/firmware/components/{componentId} | Get Server component information
[**GetServerComponents**](ServerFirmwareAPI.md#GetServerComponents) | **Get** /api/v2/servers/{serverId}/firmware/components | Get a list of Server Components
[**GetServerFirmwareInventory**](ServerFirmwareAPI.md#GetServerFirmwareInventory) | **Post** /api/v2/servers/{serverId}/firmware/inventory | Retrieves server firmware inventory from redfish
[**ScheduleServerFirmwareUpgrade**](ServerFirmwareAPI.md#ScheduleServerFirmwareUpgrade) | **Post** /api/v2/servers/{serverId}/firmware/actions/schedule-upgrade | Schedules a firmware upgrade for the specified Server
[**UpdateServerComponent**](ServerFirmwareAPI.md#UpdateServerComponent) | **Patch** /api/v2/servers/{serverId}/firmware/components/{componentId} | Updates a Server Component
[**UpdateServerFirmwareInfo**](ServerFirmwareAPI.md#UpdateServerFirmwareInfo) | **Post** /api/v2/servers/{serverId}/firmware/actions/update-info | Updates the firmware information of the Server Components
[**UpgradeFirmwareOfServer**](ServerFirmwareAPI.md#UpgradeFirmwareOfServer) | **Post** /api/v2/servers/{serverId}/firmware/actions/upgrade | Upgrades the firmware of all updatable components on Server
[**UpgradeFirmwareOfServerComponent**](ServerFirmwareAPI.md#UpgradeFirmwareOfServerComponent) | **Post** /api/v2/servers/{serverId}/firmware/components/{serverComponentId}/actions/upgrade | Upgrades the firmware of the specified component on Server
[**UpgradeFirmwareOfServersBatch**](ServerFirmwareAPI.md#UpgradeFirmwareOfServersBatch) | **Post** /api/v2/servers/firmware/actions/batch-upgrade | Upgrades the firmware of all updatable components on the specified Servers



## BatchScheduleServerFirmwareUpgrade

> BatchScheduleServerFirmwareUpgrade(ctx, serverId).BatchScheduleServerFirmwareUpgrade(batchScheduleServerFirmwareUpgrade).Execute()

Schedules a firmware upgrade for the specified Server



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
	serverId := float32(8.14) // float32 | 
	batchScheduleServerFirmwareUpgrade := *openapiclient.NewBatchScheduleServerFirmwareUpgrade([]float32{float32(123)}) // BatchScheduleServerFirmwareUpgrade | The Schedule Firmware Upgrade object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerFirmwareAPI.BatchScheduleServerFirmwareUpgrade(context.Background(), serverId).BatchScheduleServerFirmwareUpgrade(batchScheduleServerFirmwareUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.BatchScheduleServerFirmwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchScheduleServerFirmwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchScheduleServerFirmwareUpgrade** | [**BatchScheduleServerFirmwareUpgrade**](BatchScheduleServerFirmwareUpgrade.md) | The Schedule Firmware Upgrade object | 

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


## FetchAndUpdateServerFirmwareAvailableVersions

> FetchAndUpdateServerFirmwareAvailableVersions(ctx, serverId).Execute()

Fetches and updates the available firmware versions for the specified Server



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerFirmwareAPI.FetchAndUpdateServerFirmwareAvailableVersions(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.FetchAndUpdateServerFirmwareAvailableVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAndUpdateServerFirmwareAvailableVersionsRequest struct via the builder pattern


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


## GenerateServersFirmwareUpgradeAudit

> map[string]interface{} GenerateServersFirmwareUpgradeAudit(ctx).GenerateFirmwareUpgradeAudit(generateFirmwareUpgradeAudit).Execute()

Generates a Firmware Upgrade Audit for the specified Servers



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
	generateFirmwareUpgradeAudit := *openapiclient.NewGenerateFirmwareUpgradeAudit() // GenerateFirmwareUpgradeAudit | The Firmware Upgrade Audit options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.GenerateServersFirmwareUpgradeAudit(context.Background()).GenerateFirmwareUpgradeAudit(generateFirmwareUpgradeAudit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.GenerateServersFirmwareUpgradeAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateServersFirmwareUpgradeAudit`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.GenerateServersFirmwareUpgradeAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateServersFirmwareUpgradeAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateFirmwareUpgradeAudit** | [**GenerateFirmwareUpgradeAudit**](GenerateFirmwareUpgradeAudit.md) | The Firmware Upgrade Audit options | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerComponentInfo

> ServerComponent GetServerComponentInfo(ctx, serverId, componentId).Execute()

Get Server component information



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
	serverId := float32(8.14) // float32 | 
	componentId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.GetServerComponentInfo(context.Background(), serverId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.GetServerComponentInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerComponentInfo`: ServerComponent
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.GetServerComponentInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 
**componentId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerComponentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerComponent**](ServerComponent.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerComponents

> ServerComponentPaginatedList GetServerComponents(ctx, serverId).Page(page).Limit(limit).FilterId(filterId).FilterExternalId(filterExternalId).FilterServerId(filterServerId).FilterName(filterName).FilterFirmwareTargetVersion(filterFirmwareTargetVersion).FilterFirmwareVersion(filterFirmwareVersion).FilterFirmwareUpdateTimestamp(filterFirmwareUpdateTimestamp).FilterFirmwareScheduledTimestamp(filterFirmwareScheduledTimestamp).FilterFirmwareStatus(filterFirmwareStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Server Components



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
	serverId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterExternalId := []string{"Inner_example"} // []string | Filter by externalId query param.  **Format:** filter.externalId={$not}:OPERATION:VALUE    **Example:** filter.externalId=$btw:John Doe&filter.externalId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$btw:John Doe&filter.serverId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFirmwareTargetVersion := []string{"Inner_example"} // []string | Filter by firmwareTargetVersion query param.  **Format:** filter.firmwareTargetVersion={$not}:OPERATION:VALUE    **Example:** filter.firmwareTargetVersion=$btw:John Doe&filter.firmwareTargetVersion=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFirmwareVersion := []string{"Inner_example"} // []string | Filter by firmwareVersion query param.  **Format:** filter.firmwareVersion={$not}:OPERATION:VALUE    **Example:** filter.firmwareVersion=$btw:John Doe&filter.firmwareVersion=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFirmwareUpdateTimestamp := []string{"Inner_example"} // []string | Filter by firmwareUpdateTimestamp query param.  **Format:** filter.firmwareUpdateTimestamp={$not}:OPERATION:VALUE    **Example:** filter.firmwareUpdateTimestamp=$btw:John Doe&filter.firmwareUpdateTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFirmwareScheduledTimestamp := []string{"Inner_example"} // []string | Filter by firmwareScheduledTimestamp query param.  **Format:** filter.firmwareScheduledTimestamp={$not}:OPERATION:VALUE    **Example:** filter.firmwareScheduledTimestamp=$btw:John Doe&filter.firmwareScheduledTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFirmwareStatus := []string{"Inner_example"} // []string | Filter by firmwareStatus query param.  **Format:** filter.firmwareStatus={$not}:OPERATION:VALUE    **Example:** filter.firmwareStatus=$btw:John Doe&filter.firmwareStatus=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=externalId:DESC   **Default Value:** id:ASC  **Available Fields** - id  - externalId  - serverId  - name  - firmwareVersion  - firmwareScheduledTimestamp  - firmwareStatus  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,externalId,serverId,name,firmwareVersion   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - externalId  - serverId  - name  - firmwareVersion  - firmwareScheduledTimestamp  - firmwareStatus  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.GetServerComponents(context.Background(), serverId).Page(page).Limit(limit).FilterId(filterId).FilterExternalId(filterExternalId).FilterServerId(filterServerId).FilterName(filterName).FilterFirmwareTargetVersion(filterFirmwareTargetVersion).FilterFirmwareVersion(filterFirmwareVersion).FilterFirmwareUpdateTimestamp(filterFirmwareUpdateTimestamp).FilterFirmwareScheduledTimestamp(filterFirmwareScheduledTimestamp).FilterFirmwareStatus(filterFirmwareStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.GetServerComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerComponents`: ServerComponentPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.GetServerComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterExternalId** | **[]string** | Filter by externalId query param.  **Format:** filter.externalId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.externalId&#x3D;$btw:John Doe&amp;filter.externalId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$btw:John Doe&amp;filter.serverId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFirmwareTargetVersion** | **[]string** | Filter by firmwareTargetVersion query param.  **Format:** filter.firmwareTargetVersion&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.firmwareTargetVersion&#x3D;$btw:John Doe&amp;filter.firmwareTargetVersion&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFirmwareVersion** | **[]string** | Filter by firmwareVersion query param.  **Format:** filter.firmwareVersion&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.firmwareVersion&#x3D;$btw:John Doe&amp;filter.firmwareVersion&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFirmwareUpdateTimestamp** | **[]string** | Filter by firmwareUpdateTimestamp query param.  **Format:** filter.firmwareUpdateTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.firmwareUpdateTimestamp&#x3D;$btw:John Doe&amp;filter.firmwareUpdateTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFirmwareScheduledTimestamp** | **[]string** | Filter by firmwareScheduledTimestamp query param.  **Format:** filter.firmwareScheduledTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.firmwareScheduledTimestamp&#x3D;$btw:John Doe&amp;filter.firmwareScheduledTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFirmwareStatus** | **[]string** | Filter by firmwareStatus query param.  **Format:** filter.firmwareStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.firmwareStatus&#x3D;$btw:John Doe&amp;filter.firmwareStatus&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;externalId:DESC   **Default Value:** id:ASC  **Available Fields** - id  - externalId  - serverId  - name  - firmwareVersion  - firmwareScheduledTimestamp  - firmwareStatus  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,externalId,serverId,name,firmwareVersion   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - externalId  - serverId  - name  - firmwareVersion  - firmwareScheduledTimestamp  - firmwareStatus  | 

### Return type

[**ServerComponentPaginatedList**](ServerComponentPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerFirmwareInventory

> map[string]interface{} GetServerFirmwareInventory(ctx, serverId).Execute()

Retrieves server firmware inventory from redfish



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.GetServerFirmwareInventory(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.GetServerFirmwareInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerFirmwareInventory`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.GetServerFirmwareInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerFirmwareInventoryRequest struct via the builder pattern


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


## ScheduleServerFirmwareUpgrade

> ScheduleServerFirmwareUpgrade(ctx, serverId).ScheduleFirmwareUpgrade(scheduleFirmwareUpgrade).Execute()

Schedules a firmware upgrade for the specified Server



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
	serverId := float32(8.14) // float32 | 
	scheduleFirmwareUpgrade := *openapiclient.NewScheduleFirmwareUpgrade() // ScheduleFirmwareUpgrade | The Schedule Firmware Upgrade object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerFirmwareAPI.ScheduleServerFirmwareUpgrade(context.Background(), serverId).ScheduleFirmwareUpgrade(scheduleFirmwareUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.ScheduleServerFirmwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleServerFirmwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleFirmwareUpgrade** | [**ScheduleFirmwareUpgrade**](ScheduleFirmwareUpgrade.md) | The Schedule Firmware Upgrade object | 

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


## UpdateServerComponent

> ServerComponent UpdateServerComponent(ctx, serverId, componentId).UpdateServerComponent(updateServerComponent).Execute()

Updates a Server Component



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
	serverId := float32(8.14) // float32 | 
	componentId := float32(8.14) // float32 | 
	updateServerComponent := *openapiclient.NewUpdateServerComponent() // UpdateServerComponent | The Server Component update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.UpdateServerComponent(context.Background(), serverId, componentId).UpdateServerComponent(updateServerComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.UpdateServerComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerComponent`: ServerComponent
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.UpdateServerComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 
**componentId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateServerComponent** | [**UpdateServerComponent**](UpdateServerComponent.md) | The Server Component update object | 

### Return type

[**ServerComponent**](ServerComponent.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerFirmwareInfo

> UpdateServerFirmwareInfo(ctx, serverId).Execute()

Updates the firmware information of the Server Components



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerFirmwareAPI.UpdateServerFirmwareInfo(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.UpdateServerFirmwareInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerFirmwareInfoRequest struct via the builder pattern


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


## UpgradeFirmwareOfServer

> JobInfo UpgradeFirmwareOfServer(ctx, serverId).Execute()

Upgrades the firmware of all updatable components on Server



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.UpgradeFirmwareOfServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.UpgradeFirmwareOfServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeFirmwareOfServer`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.UpgradeFirmwareOfServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeFirmwareOfServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeFirmwareOfServerComponent

> JobInfo UpgradeFirmwareOfServerComponent(ctx, serverId, serverComponentId).FirmwareUpgrade(firmwareUpgrade).Execute()

Upgrades the firmware of the specified component on Server



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
	serverId := float32(8.14) // float32 | 
	serverComponentId := float32(8.14) // float32 | 
	firmwareUpgrade := *openapiclient.NewFirmwareUpgrade() // FirmwareUpgrade | The Firmware Upgrade object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.UpgradeFirmwareOfServerComponent(context.Background(), serverId, serverComponentId).FirmwareUpgrade(firmwareUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.UpgradeFirmwareOfServerComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeFirmwareOfServerComponent`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.UpgradeFirmwareOfServerComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 
**serverComponentId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeFirmwareOfServerComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firmwareUpgrade** | [**FirmwareUpgrade**](FirmwareUpgrade.md) | The Firmware Upgrade object | 

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeFirmwareOfServersBatch

> BatchServerFirmwareUpgradeResponse UpgradeFirmwareOfServersBatch(ctx).BatchServerFirmwareUpgrade(batchServerFirmwareUpgrade).Execute()

Upgrades the firmware of all updatable components on the specified Servers



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
	batchServerFirmwareUpgrade := *openapiclient.NewBatchServerFirmwareUpgrade([]float32{float32(123)}) // BatchServerFirmwareUpgrade | The Firmware Upgrade object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerFirmwareAPI.UpgradeFirmwareOfServersBatch(context.Background()).BatchServerFirmwareUpgrade(batchServerFirmwareUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerFirmwareAPI.UpgradeFirmwareOfServersBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeFirmwareOfServersBatch`: BatchServerFirmwareUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `ServerFirmwareAPI.UpgradeFirmwareOfServersBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeFirmwareOfServersBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchServerFirmwareUpgrade** | [**BatchServerFirmwareUpgrade**](BatchServerFirmwareUpgrade.md) | The Firmware Upgrade object | 

### Return type

[**BatchServerFirmwareUpgradeResponse**](BatchServerFirmwareUpgradeResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

