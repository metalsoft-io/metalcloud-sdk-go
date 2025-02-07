# \ServerFirmwareAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
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
	generateFirmwareUpgradeAudit := *openapiclient.NewGenerateFirmwareUpgradeAudit([]float32{float32(123)}) // GenerateFirmwareUpgradeAudit | The Firmware Upgrade Audit options

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterExternalId := []string{"Inner_example"} // []string | Filter by externalId query param.           <p>              <b>Format: </b> filter.externalId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.externalId=$not:$like:John Doe&filter.externalId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.           <p>              <b>Format: </b> filter.serverId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverId=$not:$like:John Doe&filter.serverId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterFirmwareTargetVersion := []string{"Inner_example"} // []string | Filter by firmwareTargetVersion query param.           <p>              <b>Format: </b> filter.firmwareTargetVersion={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.firmwareTargetVersion=$not:$like:John Doe&filter.firmwareTargetVersion=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterFirmwareVersion := []string{"Inner_example"} // []string | Filter by firmwareVersion query param.           <p>              <b>Format: </b> filter.firmwareVersion={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.firmwareVersion=$not:$like:John Doe&filter.firmwareVersion=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterFirmwareUpdateTimestamp := []string{"Inner_example"} // []string | Filter by firmwareUpdateTimestamp query param.           <p>              <b>Format: </b> filter.firmwareUpdateTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.firmwareUpdateTimestamp=$not:$like:John Doe&filter.firmwareUpdateTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterFirmwareScheduledTimestamp := []string{"Inner_example"} // []string | Filter by firmwareScheduledTimestamp query param.           <p>              <b>Format: </b> filter.firmwareScheduledTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.firmwareScheduledTimestamp=$not:$like:John Doe&filter.firmwareScheduledTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterFirmwareStatus := []string{"Inner_example"} // []string | Filter by firmwareStatus query param.           <p>              <b>Format: </b> filter.firmwareStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.firmwareStatus=$not:$like:John Doe&filter.firmwareStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>externalId</li> <li>serverId</li> <li>name</li> <li>firmwareVersion</li> <li>firmwareScheduledTimestamp</li> <li>firmwareStatus</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,externalId,serverId,name,firmwareVersion           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>externalId</li> <li>serverId</li> <li>name</li> <li>firmwareVersion</li> <li>firmwareScheduledTimestamp</li> <li>firmwareStatus</li></ul>          (optional)

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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterExternalId** | **[]string** | Filter by externalId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.externalId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.externalId&#x3D;$not:$like:John Doe&amp;filter.externalId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerId** | **[]string** | Filter by serverId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverId&#x3D;$not:$like:John Doe&amp;filter.serverId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterFirmwareTargetVersion** | **[]string** | Filter by firmwareTargetVersion query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.firmwareTargetVersion&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.firmwareTargetVersion&#x3D;$not:$like:John Doe&amp;filter.firmwareTargetVersion&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterFirmwareVersion** | **[]string** | Filter by firmwareVersion query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.firmwareVersion&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.firmwareVersion&#x3D;$not:$like:John Doe&amp;filter.firmwareVersion&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterFirmwareUpdateTimestamp** | **[]string** | Filter by firmwareUpdateTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.firmwareUpdateTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.firmwareUpdateTimestamp&#x3D;$not:$like:John Doe&amp;filter.firmwareUpdateTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterFirmwareScheduledTimestamp** | **[]string** | Filter by firmwareScheduledTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.firmwareScheduledTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.firmwareScheduledTimestamp&#x3D;$not:$like:John Doe&amp;filter.firmwareScheduledTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterFirmwareStatus** | **[]string** | Filter by firmwareStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.firmwareStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.firmwareStatus&#x3D;$not:$like:John Doe&amp;filter.firmwareStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;externalId&lt;/li&gt; &lt;li&gt;serverId&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;firmwareVersion&lt;/li&gt; &lt;li&gt;firmwareScheduledTimestamp&lt;/li&gt; &lt;li&gt;firmwareStatus&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,externalId,serverId,name,firmwareVersion           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;externalId&lt;/li&gt; &lt;li&gt;serverId&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;firmwareVersion&lt;/li&gt; &lt;li&gt;firmwareScheduledTimestamp&lt;/li&gt; &lt;li&gt;firmwareStatus&lt;/li&gt;&lt;/ul&gt;          | 

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

