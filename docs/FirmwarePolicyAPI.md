# \FirmwarePolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyFirmwarePoliciesWithServerInstanceGroups**](FirmwarePolicyAPI.md#ApplyFirmwarePoliciesWithServerInstanceGroups) | **Post** /api/v2/firmware/policies/actions/apply-with-server-instance-groups | Applies all Firmware Policies linked to server instance groups
[**ApplyFirmwarePoliciesWithoutServerInstanceGroups**](FirmwarePolicyAPI.md#ApplyFirmwarePoliciesWithoutServerInstanceGroups) | **Post** /api/v2/firmware/policies/actions/apply-without-server-instance-groups | Applies all Firmware Policies not linked to server instance groups
[**CreateFirmwarePolicy**](FirmwarePolicyAPI.md#CreateFirmwarePolicy) | **Post** /api/v2/firmware/policies | Creates a Firmware Policy
[**DeleteFirmwarePolicy**](FirmwarePolicyAPI.md#DeleteFirmwarePolicy) | **Delete** /api/v2/firmware/policies/{firmwarePolicyId} | Deletes a Firmware Policy
[**GenerateFirmwarePolicyAudit**](FirmwarePolicyAPI.md#GenerateFirmwarePolicyAudit) | **Post** /api/v2/firmware/policies/{firmwarePolicyId}/actions/generate-audit | Returns the server components ids that match this policy
[**GetFirmwarePolicies**](FirmwarePolicyAPI.md#GetFirmwarePolicies) | **Get** /api/v2/firmware/policies | Get a list of Firmware Policies
[**GetFirmwarePolicyInfo**](FirmwarePolicyAPI.md#GetFirmwarePolicyInfo) | **Get** /api/v2/firmware/policies/{firmwarePolicyId} | Get Firmware Policy information
[**GetGlobalFirmwareConfiguration**](FirmwarePolicyAPI.md#GetGlobalFirmwareConfiguration) | **Get** /api/v2/firmware/configuration | Get Global Firmware Configuration
[**UpdateFirmwarePolicy**](FirmwarePolicyAPI.md#UpdateFirmwarePolicy) | **Patch** /api/v2/firmware/policies/{firmwarePolicyId} | Updates a Firmware Policy
[**UpdateGlobalFirmwareConfiguration**](FirmwarePolicyAPI.md#UpdateGlobalFirmwareConfiguration) | **Patch** /api/v2/firmware/configuration | Updates Global Firmware Policy Configuration



## ApplyFirmwarePoliciesWithServerInstanceGroups

> ServerFirmwareUpgradePolicyApplyResult ApplyFirmwarePoliciesWithServerInstanceGroups(ctx).Execute()

Applies all Firmware Policies linked to server instance groups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.ApplyFirmwarePoliciesWithServerInstanceGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithServerInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFirmwarePoliciesWithServerInstanceGroups`: ServerFirmwareUpgradePolicyApplyResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithServerInstanceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApplyFirmwarePoliciesWithServerInstanceGroupsRequest struct via the builder pattern


### Return type

[**ServerFirmwareUpgradePolicyApplyResult**](ServerFirmwareUpgradePolicyApplyResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyFirmwarePoliciesWithoutServerInstanceGroups

> ServerFirmwareUpgradePolicyApplyResult ApplyFirmwarePoliciesWithoutServerInstanceGroups(ctx).Execute()

Applies all Firmware Policies not linked to server instance groups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.ApplyFirmwarePoliciesWithoutServerInstanceGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithoutServerInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFirmwarePoliciesWithoutServerInstanceGroups`: ServerFirmwareUpgradePolicyApplyResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithoutServerInstanceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApplyFirmwarePoliciesWithoutServerInstanceGroupsRequest struct via the builder pattern


### Return type

[**ServerFirmwareUpgradePolicyApplyResult**](ServerFirmwareUpgradePolicyApplyResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirmwarePolicy

> ServerFirmwareUpgradePolicy CreateFirmwarePolicy(ctx).CreateServerFirmwareUpgradePolicy(createServerFirmwareUpgradePolicy).Execute()

Creates a Firmware Policy



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
	createServerFirmwareUpgradePolicy := *openapiclient.NewCreateServerFirmwareUpgradePolicy("Firmware upgrade policy 1", "accept") // CreateServerFirmwareUpgradePolicy | The Firmware Policy create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.CreateFirmwarePolicy(context.Background()).CreateServerFirmwareUpgradePolicy(createServerFirmwareUpgradePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.CreateFirmwarePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirmwarePolicy`: ServerFirmwareUpgradePolicy
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.CreateFirmwarePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirmwarePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerFirmwareUpgradePolicy** | [**CreateServerFirmwareUpgradePolicy**](CreateServerFirmwareUpgradePolicy.md) | The Firmware Policy create object | 

### Return type

[**ServerFirmwareUpgradePolicy**](ServerFirmwareUpgradePolicy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirmwarePolicy

> DeleteFirmwarePolicy(ctx, firmwarePolicyId).Execute()

Deletes a Firmware Policy



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
	firmwarePolicyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirmwarePolicyAPI.DeleteFirmwarePolicy(context.Background(), firmwarePolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.DeleteFirmwarePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwarePolicyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirmwarePolicyRequest struct via the builder pattern


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


## GenerateFirmwarePolicyAudit

> ServerFirmwareUpgradePolicyAudit GenerateFirmwarePolicyAudit(ctx, firmwarePolicyId).Execute()

Returns the server components ids that match this policy



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
	firmwarePolicyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.GenerateFirmwarePolicyAudit(context.Background(), firmwarePolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.GenerateFirmwarePolicyAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateFirmwarePolicyAudit`: ServerFirmwareUpgradePolicyAudit
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.GenerateFirmwarePolicyAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwarePolicyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateFirmwarePolicyAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerFirmwareUpgradePolicyAudit**](ServerFirmwareUpgradePolicyAudit.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirmwarePolicies

> ServerFirmwareUpgradePolicyPaginatedList GetFirmwarePolicies(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterStatus(filterStatus).FilterAction(filterAction).FilterServerInstanceGroupIds(filterServerInstanceGroupIds).FilterUserIdOwner(filterUserIdOwner).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Firmware Policies



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$btw:John Doe&filter.label=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterAction := []string{"Inner_example"} // []string | Filter by action query param.  **Format:** filter.action={$not}:OPERATION:VALUE    **Example:** filter.action=$btw:John Doe&filter.action=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerInstanceGroupIds := []string{"Inner_example"} // []string | Filter by serverInstanceGroupIds query param.  **Format:** filter.serverInstanceGroupIds={$not}:OPERATION:VALUE    **Example:** filter.serverInstanceGroupIds=$btw:John Doe&filter.serverInstanceGroupIds=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUserIdOwner := []string{"Inner_example"} // []string | Filter by userIdOwner query param.  **Format:** filter.userIdOwner={$not}:OPERATION:VALUE    **Example:** filter.userIdOwner=$btw:John Doe&filter.userIdOwner=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - status  - action  - userIdOwner  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,status,action,userIdOwner   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - status  - action  - userIdOwner  - serverInstanceGroupIds  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.GetFirmwarePolicies(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterStatus(filterStatus).FilterAction(filterAction).FilterServerInstanceGroupIds(filterServerInstanceGroupIds).FilterUserIdOwner(filterUserIdOwner).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.GetFirmwarePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwarePolicies`: ServerFirmwareUpgradePolicyPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.GetFirmwarePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwarePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$btw:John Doe&amp;filter.label&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterAction** | **[]string** | Filter by action query param.  **Format:** filter.action&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.action&#x3D;$btw:John Doe&amp;filter.action&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerInstanceGroupIds** | **[]string** | Filter by serverInstanceGroupIds query param.  **Format:** filter.serverInstanceGroupIds&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverInstanceGroupIds&#x3D;$btw:John Doe&amp;filter.serverInstanceGroupIds&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUserIdOwner** | **[]string** | Filter by userIdOwner query param.  **Format:** filter.userIdOwner&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.userIdOwner&#x3D;$btw:John Doe&amp;filter.userIdOwner&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - status  - action  - userIdOwner  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,status,action,userIdOwner   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - status  - action  - userIdOwner  - serverInstanceGroupIds  | 

### Return type

[**ServerFirmwareUpgradePolicyPaginatedList**](ServerFirmwareUpgradePolicyPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirmwarePolicyInfo

> ServerFirmwareUpgradePolicy GetFirmwarePolicyInfo(ctx, firmwarePolicyId).Execute()

Get Firmware Policy information



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
	firmwarePolicyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.GetFirmwarePolicyInfo(context.Background(), firmwarePolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.GetFirmwarePolicyInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwarePolicyInfo`: ServerFirmwareUpgradePolicy
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.GetFirmwarePolicyInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwarePolicyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwarePolicyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerFirmwareUpgradePolicy**](ServerFirmwareUpgradePolicy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalFirmwareConfiguration

> GlobalFirmwareUpgradeConfiguration GetGlobalFirmwareConfiguration(ctx).Execute()

Get Global Firmware Configuration



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.GetGlobalFirmwareConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.GetGlobalFirmwareConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalFirmwareConfiguration`: GlobalFirmwareUpgradeConfiguration
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.GetGlobalFirmwareConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalFirmwareConfigurationRequest struct via the builder pattern


### Return type

[**GlobalFirmwareUpgradeConfiguration**](GlobalFirmwareUpgradeConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirmwarePolicy

> ServerFirmwareUpgradePolicy UpdateFirmwarePolicy(ctx, firmwarePolicyId).UpdateServerFirmwareUpgradePolicy(updateServerFirmwareUpgradePolicy).Execute()

Updates a Firmware Policy



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
	firmwarePolicyId := float32(8.14) // float32 | 
	updateServerFirmwareUpgradePolicy := *openapiclient.NewUpdateServerFirmwareUpgradePolicy() // UpdateServerFirmwareUpgradePolicy | The Firmware Policy update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.UpdateFirmwarePolicy(context.Background(), firmwarePolicyId).UpdateServerFirmwareUpgradePolicy(updateServerFirmwareUpgradePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.UpdateFirmwarePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirmwarePolicy`: ServerFirmwareUpgradePolicy
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.UpdateFirmwarePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwarePolicyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirmwarePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerFirmwareUpgradePolicy** | [**UpdateServerFirmwareUpgradePolicy**](UpdateServerFirmwareUpgradePolicy.md) | The Firmware Policy update object | 

### Return type

[**ServerFirmwareUpgradePolicy**](ServerFirmwareUpgradePolicy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalFirmwareConfiguration

> GlobalFirmwareUpgradeConfiguration UpdateGlobalFirmwareConfiguration(ctx).UpdateGlobalFirmwareUpgradeConfiguration(updateGlobalFirmwareUpgradeConfiguration).Execute()

Updates Global Firmware Policy Configuration



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
	updateGlobalFirmwareUpgradeConfiguration := *openapiclient.NewUpdateGlobalFirmwareUpgradeConfiguration() // UpdateGlobalFirmwareUpgradeConfiguration | The Global Firmware Policy update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwarePolicyAPI.UpdateGlobalFirmwareConfiguration(context.Background()).UpdateGlobalFirmwareUpgradeConfiguration(updateGlobalFirmwareUpgradeConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.UpdateGlobalFirmwareConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGlobalFirmwareConfiguration`: GlobalFirmwareUpgradeConfiguration
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.UpdateGlobalFirmwareConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalFirmwareConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGlobalFirmwareUpgradeConfiguration** | [**UpdateGlobalFirmwareUpgradeConfiguration**](UpdateGlobalFirmwareUpgradeConfiguration.md) | The Global Firmware Policy update object | 

### Return type

[**GlobalFirmwareUpgradeConfiguration**](GlobalFirmwareUpgradeConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

