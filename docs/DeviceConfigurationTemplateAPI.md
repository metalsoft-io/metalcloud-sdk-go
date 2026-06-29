# \DeviceConfigurationTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAssignDeviceConfigurationTemplateProfiles**](DeviceConfigurationTemplateAPI.md#BulkAssignDeviceConfigurationTemplateProfiles) | **Post** /api/v2/device-configuration-templates/profile/bulk | Bulk-assign a Device Configuration Template across a fabric (all its devices) or an explicit device list.
[**CreateDeviceConfigurationTemplate**](DeviceConfigurationTemplateAPI.md#CreateDeviceConfigurationTemplate) | **Post** /api/v2/device-configuration-templates/config | Create a Device Configuration Template
[**CreateDeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateAPI.md#CreateDeviceConfigurationTemplateProfile) | **Post** /api/v2/device-configuration-templates/profile | Create a Device Configuration Template Profile (template ↔ device ↔ fabric binding)
[**DeleteDeviceConfigurationTemplate**](DeviceConfigurationTemplateAPI.md#DeleteDeviceConfigurationTemplate) | **Delete** /api/v2/device-configuration-templates/config/{id} | Delete a Device Configuration Template
[**DeleteDeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateAPI.md#DeleteDeviceConfigurationTemplateProfile) | **Delete** /api/v2/device-configuration-templates/profile/{id} | Delete a Device Configuration Template Profile
[**FindApplicableDeviceConfigurationTemplateProfiles**](DeviceConfigurationTemplateAPI.md#FindApplicableDeviceConfigurationTemplateProfiles) | **Post** /api/v2/device-configuration-templates/profile/actions/find-applicable | List Device Configuration Template Profiles applicable to a target device or fabric
[**GetDeviceConfigurationTemplate**](DeviceConfigurationTemplateAPI.md#GetDeviceConfigurationTemplate) | **Get** /api/v2/device-configuration-templates/config/{id} | Get a Device Configuration Template by ID
[**GetDeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateAPI.md#GetDeviceConfigurationTemplateProfile) | **Get** /api/v2/device-configuration-templates/profile/{id} | Get a Device Configuration Template Profile by ID
[**GetDeviceConfigurationTemplateProfiles**](DeviceConfigurationTemplateAPI.md#GetDeviceConfigurationTemplateProfiles) | **Get** /api/v2/device-configuration-templates/profile | List Device Configuration Template Profiles
[**GetDeviceConfigurationTemplates**](DeviceConfigurationTemplateAPI.md#GetDeviceConfigurationTemplates) | **Get** /api/v2/device-configuration-templates/config | List Device Configuration Templates
[**RenderApplicableDeviceConfigurationTemplateProfiles**](DeviceConfigurationTemplateAPI.md#RenderApplicableDeviceConfigurationTemplateProfiles) | **Post** /api/v2/device-configuration-templates/profile/actions/render-applicable | Render every Device Configuration Template Profile applicable to a target device or fabric
[**RenderDeviceConfigurationTemplate**](DeviceConfigurationTemplateAPI.md#RenderDeviceConfigurationTemplate) | **Post** /api/v2/device-configuration-templates/config/actions/render | Render an unsaved Device Configuration Template body against a variables payload (stateless)
[**RenderDeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateAPI.md#RenderDeviceConfigurationTemplateProfile) | **Post** /api/v2/device-configuration-templates/profile/{id}/actions/render | Render a Device Configuration Template Profile against a target NetworkDevice
[**RenderSavedDeviceConfigurationTemplate**](DeviceConfigurationTemplateAPI.md#RenderSavedDeviceConfigurationTemplate) | **Post** /api/v2/device-configuration-templates/config/{id}/actions/render | Render a saved Device Configuration Template by ID, merging customVariablesJson with request variables.
[**UpdateDeviceConfigurationTemplate**](DeviceConfigurationTemplateAPI.md#UpdateDeviceConfigurationTemplate) | **Patch** /api/v2/device-configuration-templates/config/{id} | Update a Device Configuration Template
[**UpdateDeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateAPI.md#UpdateDeviceConfigurationTemplateProfile) | **Patch** /api/v2/device-configuration-templates/profile/{id} | Update a Device Configuration Template Profile



## BulkAssignDeviceConfigurationTemplateProfiles

> BulkAssignDeviceConfigurationTemplateProfileResult BulkAssignDeviceConfigurationTemplateProfiles(ctx).BulkAssignDeviceConfigurationTemplateProfile(bulkAssignDeviceConfigurationTemplateProfile).Execute()

Bulk-assign a Device Configuration Template across a fabric (all its devices) or an explicit device list.

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
	bulkAssignDeviceConfigurationTemplateProfile := *openapiclient.NewBulkAssignDeviceConfigurationTemplateProfile(int64(1)) // BulkAssignDeviceConfigurationTemplateProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.BulkAssignDeviceConfigurationTemplateProfiles(context.Background()).BulkAssignDeviceConfigurationTemplateProfile(bulkAssignDeviceConfigurationTemplateProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.BulkAssignDeviceConfigurationTemplateProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAssignDeviceConfigurationTemplateProfiles`: BulkAssignDeviceConfigurationTemplateProfileResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.BulkAssignDeviceConfigurationTemplateProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAssignDeviceConfigurationTemplateProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkAssignDeviceConfigurationTemplateProfile** | [**BulkAssignDeviceConfigurationTemplateProfile**](BulkAssignDeviceConfigurationTemplateProfile.md) |  | 

### Return type

[**BulkAssignDeviceConfigurationTemplateProfileResult**](BulkAssignDeviceConfigurationTemplateProfileResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceConfigurationTemplate

> DeviceConfigurationTemplate CreateDeviceConfigurationTemplate(ctx).CreateDeviceConfigurationTemplate(createDeviceConfigurationTemplate).Execute()

Create a Device Configuration Template

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
	createDeviceConfigurationTemplate := *openapiclient.NewCreateDeviceConfigurationTemplate("nvidia-spectrum-x-base", openapiclient.SwitchDriver("cisco_aci51"), openapiclient.NetworkTemplateExecutionType("cli")) // CreateDeviceConfigurationTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.CreateDeviceConfigurationTemplate(context.Background()).CreateDeviceConfigurationTemplate(createDeviceConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.CreateDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceConfigurationTemplate`: DeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.CreateDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeviceConfigurationTemplate** | [**CreateDeviceConfigurationTemplate**](CreateDeviceConfigurationTemplate.md) |  | 

### Return type

[**DeviceConfigurationTemplate**](DeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceConfigurationTemplateProfile

> DeviceConfigurationTemplateProfile CreateDeviceConfigurationTemplateProfile(ctx).CreateDeviceConfigurationTemplateProfile(createDeviceConfigurationTemplateProfile).Execute()

Create a Device Configuration Template Profile (template ↔ device ↔ fabric binding)

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
	createDeviceConfigurationTemplateProfile := *openapiclient.NewCreateDeviceConfigurationTemplateProfile(int64(1)) // CreateDeviceConfigurationTemplateProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.CreateDeviceConfigurationTemplateProfile(context.Background()).CreateDeviceConfigurationTemplateProfile(createDeviceConfigurationTemplateProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.CreateDeviceConfigurationTemplateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceConfigurationTemplateProfile`: DeviceConfigurationTemplateProfile
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.CreateDeviceConfigurationTemplateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceConfigurationTemplateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeviceConfigurationTemplateProfile** | [**CreateDeviceConfigurationTemplateProfile**](CreateDeviceConfigurationTemplateProfile.md) |  | 

### Return type

[**DeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceConfigurationTemplate

> DeleteDeviceConfigurationTemplate(ctx, id).IfMatch(ifMatch).Execute()

Delete a Device Configuration Template

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
	id := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceConfigurationTemplateAPI.DeleteDeviceConfigurationTemplate(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.DeleteDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceConfigurationTemplateRequest struct via the builder pattern


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


## DeleteDeviceConfigurationTemplateProfile

> DeleteDeviceConfigurationTemplateProfile(ctx, id).IfMatch(ifMatch).Execute()

Delete a Device Configuration Template Profile

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
	id := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceConfigurationTemplateAPI.DeleteDeviceConfigurationTemplateProfile(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.DeleteDeviceConfigurationTemplateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceConfigurationTemplateProfileRequest struct via the builder pattern


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


## FindApplicableDeviceConfigurationTemplateProfiles

> FindApplicableDeviceConfigurationTemplateProfilesResult FindApplicableDeviceConfigurationTemplateProfiles(ctx).FindApplicableDeviceConfigurationTemplateProfiles(findApplicableDeviceConfigurationTemplateProfiles).Execute()

List Device Configuration Template Profiles applicable to a target device or fabric



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
	findApplicableDeviceConfigurationTemplateProfiles := *openapiclient.NewFindApplicableDeviceConfigurationTemplateProfiles() // FindApplicableDeviceConfigurationTemplateProfiles | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.FindApplicableDeviceConfigurationTemplateProfiles(context.Background()).FindApplicableDeviceConfigurationTemplateProfiles(findApplicableDeviceConfigurationTemplateProfiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.FindApplicableDeviceConfigurationTemplateProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindApplicableDeviceConfigurationTemplateProfiles`: FindApplicableDeviceConfigurationTemplateProfilesResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.FindApplicableDeviceConfigurationTemplateProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindApplicableDeviceConfigurationTemplateProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findApplicableDeviceConfigurationTemplateProfiles** | [**FindApplicableDeviceConfigurationTemplateProfiles**](FindApplicableDeviceConfigurationTemplateProfiles.md) |  | 

### Return type

[**FindApplicableDeviceConfigurationTemplateProfilesResult**](FindApplicableDeviceConfigurationTemplateProfilesResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceConfigurationTemplate

> DeviceConfigurationTemplate GetDeviceConfigurationTemplate(ctx, id).Execute()

Get a Device Configuration Template by ID

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
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceConfigurationTemplate`: DeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceConfigurationTemplate**](DeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceConfigurationTemplateProfile

> DeviceConfigurationTemplateProfile GetDeviceConfigurationTemplateProfile(ctx, id).Execute()

Get a Device Configuration Template Profile by ID

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
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplateProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceConfigurationTemplateProfile`: DeviceConfigurationTemplateProfile
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigurationTemplateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceConfigurationTemplateProfiles

> DeviceConfigurationTemplateProfilePaginatedList GetDeviceConfigurationTemplateProfiles(ctx).Tag(tag).Page(page).Limit(limit).FilterId(filterId).FilterDeviceConfigurationTemplateId(filterDeviceConfigurationTemplateId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterNetworkFabricId(filterNetworkFabricId).FilterLifecycleStage(filterLifecycleStage).FilterIsEnabled(filterIsEnabled).FilterPriority(filterPriority).FilterApplyMode(filterApplyMode).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Device Configuration Template Profiles

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
	tag := "canary" // string | Filter by exact tag membership. Composes with `filter.*` and `sortBy=` like any other query param. (optional)
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe&filter.id=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterDeviceConfigurationTemplateId := []string{"Inner_example"} // []string | Filter by deviceConfigurationTemplateId query param.  **Format:** filter.deviceConfigurationTemplateId={$not}:OPERATION:VALUE    **Example:** filter.deviceConfigurationTemplateId=$eq:John Doe&filter.deviceConfigurationTemplateId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkDeviceId := []string{"Inner_example"} // []string | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId=$eq:John Doe&filter.networkDeviceId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterNetworkFabricId := []string{"Inner_example"} // []string | Filter by networkFabricId query param.  **Format:** filter.networkFabricId={$not}:OPERATION:VALUE    **Example:** filter.networkFabricId=$eq:John Doe&filter.networkFabricId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterLifecycleStage := []string{"Inner_example"} // []string | Filter by lifecycleStage query param.  **Format:** filter.lifecycleStage={$not}:OPERATION:VALUE    **Example:** filter.lifecycleStage=$eq:John Doe&filter.lifecycleStage=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterIsEnabled := []string{"Inner_example"} // []string | Filter by isEnabled query param.  **Format:** filter.isEnabled={$not}:OPERATION:VALUE    **Example:** filter.isEnabled=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterPriority := []string{"Inner_example"} // []string | Filter by priority query param.  **Format:** filter.priority={$not}:OPERATION:VALUE    **Example:** filter.priority=$eq:John Doe&filter.priority=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterApplyMode := []string{"Inner_example"} // []string | Filter by applyMode query param.  **Format:** filter.applyMode={$not}:OPERATION:VALUE    **Example:** filter.applyMode=$eq:John Doe&filter.applyMode=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=deviceConfigurationTemplateId:DESC   **Default Value:** priority:ASC  **Available Fields** - id  - deviceConfigurationTemplateId  - networkDeviceId  - networkFabricId  - lifecycleStage  - isEnabled  - priority  - applyMode  - createdTimestamp  - updatedTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplateProfiles(context.Background()).Tag(tag).Page(page).Limit(limit).FilterId(filterId).FilterDeviceConfigurationTemplateId(filterDeviceConfigurationTemplateId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterNetworkFabricId(filterNetworkFabricId).FilterLifecycleStage(filterLifecycleStage).FilterIsEnabled(filterIsEnabled).FilterPriority(filterPriority).FilterApplyMode(filterApplyMode).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplateProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceConfigurationTemplateProfiles`: DeviceConfigurationTemplateProfilePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplateProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigurationTemplateProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** | Filter by exact tag membership. Composes with &#x60;filter.*&#x60; and &#x60;sortBy&#x3D;&#x60; like any other query param. | 
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterDeviceConfigurationTemplateId** | **[]string** | Filter by deviceConfigurationTemplateId query param.  **Format:** filter.deviceConfigurationTemplateId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.deviceConfigurationTemplateId&#x3D;$eq:John Doe&amp;filter.deviceConfigurationTemplateId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkDeviceId** | **[]string** | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId&#x3D;$eq:John Doe&amp;filter.networkDeviceId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterNetworkFabricId** | **[]string** | Filter by networkFabricId query param.  **Format:** filter.networkFabricId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkFabricId&#x3D;$eq:John Doe&amp;filter.networkFabricId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterLifecycleStage** | **[]string** | Filter by lifecycleStage query param.  **Format:** filter.lifecycleStage&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.lifecycleStage&#x3D;$eq:John Doe&amp;filter.lifecycleStage&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterIsEnabled** | **[]string** | Filter by isEnabled query param.  **Format:** filter.isEnabled&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.isEnabled&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterPriority** | **[]string** | Filter by priority query param.  **Format:** filter.priority&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.priority&#x3D;$eq:John Doe&amp;filter.priority&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterApplyMode** | **[]string** | Filter by applyMode query param.  **Format:** filter.applyMode&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.applyMode&#x3D;$eq:John Doe&amp;filter.applyMode&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;deviceConfigurationTemplateId:DESC   **Default Value:** priority:ASC  **Available Fields** - id  - deviceConfigurationTemplateId  - networkDeviceId  - networkFabricId  - lifecycleStage  - isEnabled  - priority  - applyMode  - createdTimestamp  - updatedTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:**    **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields**   | 

### Return type

[**DeviceConfigurationTemplateProfilePaginatedList**](DeviceConfigurationTemplateProfilePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceConfigurationTemplates

> DeviceConfigurationTemplatePaginatedList GetDeviceConfigurationTemplates(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDeviceDriver(filterDeviceDriver).FilterExecutionType(filterExecutionType).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Device Configuration Templates

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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe&filter.label=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe&filter.name=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterDeviceDriver := []string{"Inner_example"} // []string | Filter by deviceDriver query param.  **Format:** filter.deviceDriver={$not}:OPERATION:VALUE    **Example:** filter.deviceDriver=$eq:John Doe&filter.deviceDriver=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterExecutionType := []string{"Inner_example"} // []string | Filter by executionType query param.  **Format:** filter.executionType={$not}:OPERATION:VALUE    **Example:** filter.executionType=$eq:John Doe&filter.executionType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.  **Format:** filter.tags={$not}:OPERATION:VALUE    **Example:** filter.tags=$eq:John Doe&filter.tags=$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  - deviceDriver  - executionType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,name,deviceDriver,executionType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  - deviceDriver  - executionType  - tags  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplates(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDeviceDriver(filterDeviceDriver).FilterExecutionType(filterExecutionType).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceConfigurationTemplates`: DeviceConfigurationTemplatePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.GetDeviceConfigurationTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigurationTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe&amp;filter.label&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe&amp;filter.name&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterDeviceDriver** | **[]string** | Filter by deviceDriver query param.  **Format:** filter.deviceDriver&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.deviceDriver&#x3D;$eq:John Doe&amp;filter.deviceDriver&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterExecutionType** | **[]string** | Filter by executionType query param.  **Format:** filter.executionType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.executionType&#x3D;$eq:John Doe&amp;filter.executionType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterTags** | **[]string** | Filter by tags query param.  **Format:** filter.tags&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.tags&#x3D;$eq:John Doe&amp;filter.tags&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  - deviceDriver  - executionType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,name,deviceDriver,executionType   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  - deviceDriver  - executionType  - tags  | 

### Return type

[**DeviceConfigurationTemplatePaginatedList**](DeviceConfigurationTemplatePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderApplicableDeviceConfigurationTemplateProfiles

> RenderedApplicableDeviceConfigurationTemplateProfilesResult RenderApplicableDeviceConfigurationTemplateProfiles(ctx).RenderApplicableDeviceConfigurationTemplateProfiles(renderApplicableDeviceConfigurationTemplateProfiles).Execute()

Render every Device Configuration Template Profile applicable to a target device or fabric



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
	renderApplicableDeviceConfigurationTemplateProfiles := *openapiclient.NewRenderApplicableDeviceConfigurationTemplateProfiles() // RenderApplicableDeviceConfigurationTemplateProfiles | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.RenderApplicableDeviceConfigurationTemplateProfiles(context.Background()).RenderApplicableDeviceConfigurationTemplateProfiles(renderApplicableDeviceConfigurationTemplateProfiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.RenderApplicableDeviceConfigurationTemplateProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderApplicableDeviceConfigurationTemplateProfiles`: RenderedApplicableDeviceConfigurationTemplateProfilesResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.RenderApplicableDeviceConfigurationTemplateProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderApplicableDeviceConfigurationTemplateProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renderApplicableDeviceConfigurationTemplateProfiles** | [**RenderApplicableDeviceConfigurationTemplateProfiles**](RenderApplicableDeviceConfigurationTemplateProfiles.md) |  | 

### Return type

[**RenderedApplicableDeviceConfigurationTemplateProfilesResult**](RenderedApplicableDeviceConfigurationTemplateProfilesResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderDeviceConfigurationTemplate

> RenderedDeviceConfigurationTemplate RenderDeviceConfigurationTemplate(ctx).RenderDeviceConfigurationTemplate(renderDeviceConfigurationTemplate).Execute()

Render an unsaved Device Configuration Template body against a variables payload (stateless)

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
	renderDeviceConfigurationTemplate := *openapiclient.NewRenderDeviceConfigurationTemplate("cm91dGVyIGJncCB7eyBhc24gfX0K") // RenderDeviceConfigurationTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.RenderDeviceConfigurationTemplate(context.Background()).RenderDeviceConfigurationTemplate(renderDeviceConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.RenderDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderDeviceConfigurationTemplate`: RenderedDeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.RenderDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renderDeviceConfigurationTemplate** | [**RenderDeviceConfigurationTemplate**](RenderDeviceConfigurationTemplate.md) |  | 

### Return type

[**RenderedDeviceConfigurationTemplate**](RenderedDeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderDeviceConfigurationTemplateProfile

> RenderedDeviceConfigurationTemplate RenderDeviceConfigurationTemplateProfile(ctx, id).RenderDeviceConfigurationTemplateProfile(renderDeviceConfigurationTemplateProfile).Execute()

Render a Device Configuration Template Profile against a target NetworkDevice



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
	id := int64(789) // int64 | 
	renderDeviceConfigurationTemplateProfile := *openapiclient.NewRenderDeviceConfigurationTemplateProfile(int64(42)) // RenderDeviceConfigurationTemplateProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.RenderDeviceConfigurationTemplateProfile(context.Background(), id).RenderDeviceConfigurationTemplateProfile(renderDeviceConfigurationTemplateProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.RenderDeviceConfigurationTemplateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderDeviceConfigurationTemplateProfile`: RenderedDeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.RenderDeviceConfigurationTemplateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenderDeviceConfigurationTemplateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renderDeviceConfigurationTemplateProfile** | [**RenderDeviceConfigurationTemplateProfile**](RenderDeviceConfigurationTemplateProfile.md) |  | 

### Return type

[**RenderedDeviceConfigurationTemplate**](RenderedDeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderSavedDeviceConfigurationTemplate

> RenderedDeviceConfigurationTemplate RenderSavedDeviceConfigurationTemplate(ctx, id).RenderSavedDeviceConfigurationTemplate(renderSavedDeviceConfigurationTemplate).Execute()

Render a saved Device Configuration Template by ID, merging customVariablesJson with request variables.

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
	id := int64(789) // int64 | 
	renderSavedDeviceConfigurationTemplate := *openapiclient.NewRenderSavedDeviceConfigurationTemplate() // RenderSavedDeviceConfigurationTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.RenderSavedDeviceConfigurationTemplate(context.Background(), id).RenderSavedDeviceConfigurationTemplate(renderSavedDeviceConfigurationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.RenderSavedDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderSavedDeviceConfigurationTemplate`: RenderedDeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.RenderSavedDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenderSavedDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renderSavedDeviceConfigurationTemplate** | [**RenderSavedDeviceConfigurationTemplate**](RenderSavedDeviceConfigurationTemplate.md) |  | 

### Return type

[**RenderedDeviceConfigurationTemplate**](RenderedDeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceConfigurationTemplate

> DeviceConfigurationTemplate UpdateDeviceConfigurationTemplate(ctx, id).UpdateDeviceConfigurationTemplate(updateDeviceConfigurationTemplate).IfMatch(ifMatch).Execute()

Update a Device Configuration Template

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
	id := int64(789) // int64 | 
	updateDeviceConfigurationTemplate := *openapiclient.NewUpdateDeviceConfigurationTemplate() // UpdateDeviceConfigurationTemplate | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.UpdateDeviceConfigurationTemplate(context.Background(), id).UpdateDeviceConfigurationTemplate(updateDeviceConfigurationTemplate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.UpdateDeviceConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceConfigurationTemplate`: DeviceConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.UpdateDeviceConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceConfigurationTemplate** | [**UpdateDeviceConfigurationTemplate**](UpdateDeviceConfigurationTemplate.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**DeviceConfigurationTemplate**](DeviceConfigurationTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceConfigurationTemplateProfile

> DeviceConfigurationTemplateProfile UpdateDeviceConfigurationTemplateProfile(ctx, id).UpdateDeviceConfigurationTemplateProfile(updateDeviceConfigurationTemplateProfile).IfMatch(ifMatch).Execute()

Update a Device Configuration Template Profile

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
	id := int64(789) // int64 | 
	updateDeviceConfigurationTemplateProfile := *openapiclient.NewUpdateDeviceConfigurationTemplateProfile() // UpdateDeviceConfigurationTemplateProfile | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceConfigurationTemplateAPI.UpdateDeviceConfigurationTemplateProfile(context.Background(), id).UpdateDeviceConfigurationTemplateProfile(updateDeviceConfigurationTemplateProfile).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceConfigurationTemplateAPI.UpdateDeviceConfigurationTemplateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceConfigurationTemplateProfile`: DeviceConfigurationTemplateProfile
	fmt.Fprintf(os.Stdout, "Response from `DeviceConfigurationTemplateAPI.UpdateDeviceConfigurationTemplateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceConfigurationTemplateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceConfigurationTemplateProfile** | [**UpdateDeviceConfigurationTemplateProfile**](UpdateDeviceConfigurationTemplateProfile.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**DeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

