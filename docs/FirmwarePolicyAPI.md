# \FirmwarePolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyFirmwarePoliciesWithInstanceArrays**](FirmwarePolicyAPI.md#ApplyFirmwarePoliciesWithInstanceArrays) | **Post** /api/v2/firmware/policies/actions/apply-with-instance-arrays | Applies all Firmware Policies linked to instance arrays
[**ApplyFirmwarePoliciesWithoutInstanceArrays**](FirmwarePolicyAPI.md#ApplyFirmwarePoliciesWithoutInstanceArrays) | **Post** /api/v2/firmware/policies/actions/apply-without-instance-arrays | Applies all Firmware Policies not linked to instance arrays
[**CreateFirmwarePolicy**](FirmwarePolicyAPI.md#CreateFirmwarePolicy) | **Post** /api/v2/firmware/policies | Creates a Firmware Policy
[**DeleteFirmwarePolicy**](FirmwarePolicyAPI.md#DeleteFirmwarePolicy) | **Delete** /api/v2/firmware/policies/{firmwarePolicyId} | Deletes a Firmware Policy
[**GenerateFirmwarePolicyAudit**](FirmwarePolicyAPI.md#GenerateFirmwarePolicyAudit) | **Post** /api/v2/firmware/policies/{firmwarePolicyId}/actions/generate-audit | Returns the server components ids that match this policy
[**GetFirmwarePolicies**](FirmwarePolicyAPI.md#GetFirmwarePolicies) | **Get** /api/v2/firmware/policies | Get a list of Firmware Policies
[**GetFirmwarePolicyInfo**](FirmwarePolicyAPI.md#GetFirmwarePolicyInfo) | **Get** /api/v2/firmware/policies/{firmwarePolicyId} | Get Firmware Policy information
[**GetGlobalFirmwareConfiguration**](FirmwarePolicyAPI.md#GetGlobalFirmwareConfiguration) | **Get** /api/v2/firmware/configuration | Get Global Firmware Configuration
[**UpdateFirmwarePolicy**](FirmwarePolicyAPI.md#UpdateFirmwarePolicy) | **Patch** /api/v2/firmware/policies/{firmwarePolicyId} | Updates a Firmware Policy
[**UpdateGlobalFirmwareConfiguration**](FirmwarePolicyAPI.md#UpdateGlobalFirmwareConfiguration) | **Patch** /api/v2/firmware/configuration | Updates Global Firmware Policy Configuration



## ApplyFirmwarePoliciesWithInstanceArrays

> ServerFirmwareUpgradePolicyApplyResult ApplyFirmwarePoliciesWithInstanceArrays(ctx).Execute()

Applies all Firmware Policies linked to instance arrays



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
	resp, r, err := apiClient.FirmwarePolicyAPI.ApplyFirmwarePoliciesWithInstanceArrays(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithInstanceArrays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFirmwarePoliciesWithInstanceArrays`: ServerFirmwareUpgradePolicyApplyResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithInstanceArrays`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApplyFirmwarePoliciesWithInstanceArraysRequest struct via the builder pattern


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


## ApplyFirmwarePoliciesWithoutInstanceArrays

> ServerFirmwareUpgradePolicyApplyResult ApplyFirmwarePoliciesWithoutInstanceArrays(ctx).Execute()

Applies all Firmware Policies not linked to instance arrays



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
	resp, r, err := apiClient.FirmwarePolicyAPI.ApplyFirmwarePoliciesWithoutInstanceArrays(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithoutInstanceArrays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFirmwarePoliciesWithoutInstanceArrays`: ServerFirmwareUpgradePolicyApplyResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwarePolicyAPI.ApplyFirmwarePoliciesWithoutInstanceArrays`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApplyFirmwarePoliciesWithoutInstanceArraysRequest struct via the builder pattern


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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.           <p>              <b>Format: </b> filter.status={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.status=$not:$like:John Doe&filter.status=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterAction := []string{"Inner_example"} // []string | Filter by action query param.           <p>              <b>Format: </b> filter.action={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.action=$not:$like:John Doe&filter.action=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerInstanceGroupIds := []string{"Inner_example"} // []string | Filter by serverInstanceGroupIds query param.           <p>              <b>Format: </b> filter.serverInstanceGroupIds={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverInstanceGroupIds=$not:$like:John Doe&filter.serverInstanceGroupIds=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterUserIdOwner := []string{"Inner_example"} // []string | Filter by userIdOwner query param.           <p>              <b>Format: </b> filter.userIdOwner={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.userIdOwner=$not:$like:John Doe&filter.userIdOwner=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>status</li> <li>action</li> <li>userIdOwner</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,status,action,userIdOwner           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>status</li> <li>action</li> <li>userIdOwner</li> <li>serverInstanceGroupIds</li></ul>          (optional)

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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterStatus** | **[]string** | Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterAction** | **[]string** | Filter by action query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.action&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.action&#x3D;$not:$like:John Doe&amp;filter.action&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerInstanceGroupIds** | **[]string** | Filter by serverInstanceGroupIds query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverInstanceGroupIds&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverInstanceGroupIds&#x3D;$not:$like:John Doe&amp;filter.serverInstanceGroupIds&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterUserIdOwner** | **[]string** | Filter by userIdOwner query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.userIdOwner&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.userIdOwner&#x3D;$not:$like:John Doe&amp;filter.userIdOwner&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;action&lt;/li&gt; &lt;li&gt;userIdOwner&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,status,action,userIdOwner           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;action&lt;/li&gt; &lt;li&gt;userIdOwner&lt;/li&gt; &lt;li&gt;serverInstanceGroupIds&lt;/li&gt;&lt;/ul&gt;          | 

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

