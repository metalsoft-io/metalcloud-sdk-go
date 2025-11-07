# \ServerCleanupPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerCleanupPolicy**](ServerCleanupPolicyAPI.md#CreateServerCleanupPolicy) | **Post** /api/v2/servers/cleanup-policies | Creates a Server Cleanup Policy
[**DeleteServerCleanupPolicy**](ServerCleanupPolicyAPI.md#DeleteServerCleanupPolicy) | **Delete** /api/v2/servers/cleanup-policies/{serverCleanupPolicyId} | Deletes a Server Cleanup Policy
[**GetServerCleanupPolicies**](ServerCleanupPolicyAPI.md#GetServerCleanupPolicies) | **Get** /api/v2/servers/cleanup-policies | Get a list of Server Cleanup Policies
[**GetServerCleanupPolicyInfo**](ServerCleanupPolicyAPI.md#GetServerCleanupPolicyInfo) | **Get** /api/v2/servers/cleanup-policies/{serverCleanupPolicyId} | Get Server Cleanup Policy information
[**UpdateServerCleanupPolicy**](ServerCleanupPolicyAPI.md#UpdateServerCleanupPolicy) | **Patch** /api/v2/servers/cleanup-policies/{serverCleanupPolicyId} | Updates a Server Cleanup Policy



## CreateServerCleanupPolicy

> ServerCleanupPolicy CreateServerCleanupPolicy(ctx).CreateServerCleanupPolicy(createServerCleanupPolicy).Execute()

Creates a Server Cleanup Policy



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
	createServerCleanupPolicy := *openapiclient.NewCreateServerCleanupPolicy("Label_example", float32(1), float32(1), float32(1), float32(1), "RAID0", "RAID1", "RAID5", "RAID10", []string{"SkipRaidActions_example"}) // CreateServerCleanupPolicy | The Server Cleanup Policy create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerCleanupPolicyAPI.CreateServerCleanupPolicy(context.Background()).CreateServerCleanupPolicy(createServerCleanupPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerCleanupPolicyAPI.CreateServerCleanupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerCleanupPolicy`: ServerCleanupPolicy
	fmt.Fprintf(os.Stdout, "Response from `ServerCleanupPolicyAPI.CreateServerCleanupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerCleanupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerCleanupPolicy** | [**CreateServerCleanupPolicy**](CreateServerCleanupPolicy.md) | The Server Cleanup Policy create object | 

### Return type

[**ServerCleanupPolicy**](ServerCleanupPolicy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerCleanupPolicy

> DeleteServerCleanupPolicy(ctx, serverCleanupPolicyId).Execute()

Deletes a Server Cleanup Policy



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
	serverCleanupPolicyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerCleanupPolicyAPI.DeleteServerCleanupPolicy(context.Background(), serverCleanupPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerCleanupPolicyAPI.DeleteServerCleanupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverCleanupPolicyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerCleanupPolicyRequest struct via the builder pattern


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


## GetServerCleanupPolicies

> ServerCleanupPolicyPaginatedList GetServerCleanupPolicies(ctx).Page(page).Limit(limit).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Server Cleanup Policies



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$btw:John Doe&filter.label=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerCleanupPolicyAPI.GetServerCleanupPolicies(context.Background()).Page(page).Limit(limit).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerCleanupPolicyAPI.GetServerCleanupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCleanupPolicies`: ServerCleanupPolicyPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerCleanupPolicyAPI.GetServerCleanupPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCleanupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$btw:John Doe&amp;filter.label&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  | 

### Return type

[**ServerCleanupPolicyPaginatedList**](ServerCleanupPolicyPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerCleanupPolicyInfo

> ServerCleanupPolicy GetServerCleanupPolicyInfo(ctx, serverCleanupPolicyId).Execute()

Get Server Cleanup Policy information



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
	serverCleanupPolicyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerCleanupPolicyAPI.GetServerCleanupPolicyInfo(context.Background(), serverCleanupPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerCleanupPolicyAPI.GetServerCleanupPolicyInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerCleanupPolicyInfo`: ServerCleanupPolicy
	fmt.Fprintf(os.Stdout, "Response from `ServerCleanupPolicyAPI.GetServerCleanupPolicyInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverCleanupPolicyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCleanupPolicyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerCleanupPolicy**](ServerCleanupPolicy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerCleanupPolicy

> ServerCleanupPolicy UpdateServerCleanupPolicy(ctx, serverCleanupPolicyId).UpdateServerCleanupPolicy(updateServerCleanupPolicy).Execute()

Updates a Server Cleanup Policy



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
	serverCleanupPolicyId := float32(8.14) // float32 | 
	updateServerCleanupPolicy := *openapiclient.NewUpdateServerCleanupPolicy() // UpdateServerCleanupPolicy | The Server Cleanup Policy update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerCleanupPolicyAPI.UpdateServerCleanupPolicy(context.Background(), serverCleanupPolicyId).UpdateServerCleanupPolicy(updateServerCleanupPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerCleanupPolicyAPI.UpdateServerCleanupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerCleanupPolicy`: ServerCleanupPolicy
	fmt.Fprintf(os.Stdout, "Response from `ServerCleanupPolicyAPI.UpdateServerCleanupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverCleanupPolicyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerCleanupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerCleanupPolicy** | [**UpdateServerCleanupPolicy**](UpdateServerCleanupPolicy.md) | The Server Cleanup Policy update object | 

### Return type

[**ServerCleanupPolicy**](ServerCleanupPolicy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

