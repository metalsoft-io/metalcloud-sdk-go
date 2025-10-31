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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li></ul>          (optional)

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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

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

