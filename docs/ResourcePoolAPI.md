# \ResourcePoolAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourcePoolUser**](ResourcePoolAPI.md#AddResourcePoolUser) | **Post** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Add a user to a Resource Pool
[**AddServerToResourcePool**](ResourcePoolAPI.md#AddServerToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Add a server to a Resource Pool
[**AddSubnetPoolToResourcePool**](ResourcePoolAPI.md#AddSubnetPoolToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Add a subnet pool to a resource pool
[**CreateResourcePool**](ResourcePoolAPI.md#CreateResourcePool) | **Post** /api/v2/resource-pools | Creates a Resource Pool
[**DeleteResourcePool**](ResourcePoolAPI.md#DeleteResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId} | Deletes a Resource Pool
[**GetResourcePool**](ResourcePoolAPI.md#GetResourcePool) | **Get** /api/v2/resource-pools/{resourcePoolId} | Get Resource Pool information
[**GetResourcePoolServers**](ResourcePoolAPI.md#GetResourcePoolServers) | **Get** /api/v2/resource-pools/{resourcePoolId}/servers | Get all servers that are part of a Resource Pool
[**GetResourcePoolSubnetPools**](ResourcePoolAPI.md#GetResourcePoolSubnetPools) | **Get** /api/v2/resource-pools/{resourcePoolId}/subnet-pools | Get all subnet pools that are part of a resource pool
[**GetResourcePoolUsers**](ResourcePoolAPI.md#GetResourcePoolUsers) | **Get** /api/v2/resource-pools/{resourcePoolId}/users | Get all users that have access to a Resource Pool
[**GetResourcePools**](ResourcePoolAPI.md#GetResourcePools) | **Get** /api/v2/resource-pools | Get all Resource Pools
[**GetUserResourcePools**](ResourcePoolAPI.md#GetUserResourcePools) | **Get** /api/v2/resource-pools/user/{userId} | Get all Resource Pools that a user has access to
[**RemoveResourcePoolUser**](ResourcePoolAPI.md#RemoveResourcePoolUser) | **Delete** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Remove a user from a Resource Pool
[**RemoveServerFromResourcePool**](ResourcePoolAPI.md#RemoveServerFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Remove a server from a Resource Pool
[**RemoveSubnetPoolFromResourcePool**](ResourcePoolAPI.md#RemoveSubnetPoolFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Remove a subnet from a resource pool
[**UpdateResourcePool**](ResourcePoolAPI.md#UpdateResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId} | Updates Resource Pool information



## AddResourcePoolUser

> AddResourcePoolUser(ctx, resourcePoolId, userId).Execute()

Add a user to a Resource Pool



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
	resourcePoolId := float32(8.14) // float32 | 
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcePoolAPI.AddResourcePoolUser(context.Background(), resourcePoolId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.AddResourcePoolUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourcePoolUserRequest struct via the builder pattern


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


## AddServerToResourcePool

> AddServerToResourcePool(ctx, resourcePoolId, serverId).Execute()

Add a server to a Resource Pool



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
	resourcePoolId := float32(8.14) // float32 | 
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcePoolAPI.AddServerToResourcePool(context.Background(), resourcePoolId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.AddServerToResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServerToResourcePoolRequest struct via the builder pattern


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


## AddSubnetPoolToResourcePool

> AddSubnetPoolToResourcePool(ctx, resourcePoolId, subnetPoolId).Execute()

Add a subnet pool to a resource pool



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
	resourcePoolId := float32(8.14) // float32 | 
	subnetPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcePoolAPI.AddSubnetPoolToResourcePool(context.Background(), resourcePoolId, subnetPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.AddSubnetPoolToResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 
**subnetPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSubnetPoolToResourcePoolRequest struct via the builder pattern


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


## CreateResourcePool

> ResourcePool CreateResourcePool(ctx).CreateResourcePool(createResourcePool).Execute()

Creates a Resource Pool



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
	createResourcePool := *openapiclient.NewCreateResourcePool("ResourcePoolLabel_example", "ResourcePoolDescription_example") // CreateResourcePool | The Resource Pool create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.CreateResourcePool(context.Background()).CreateResourcePool(createResourcePool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.CreateResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResourcePool`: ResourcePool
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.CreateResourcePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createResourcePool** | [**CreateResourcePool**](CreateResourcePool.md) | The Resource Pool create object | 

### Return type

[**ResourcePool**](ResourcePool.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourcePool

> DeleteResourcePool(ctx, resourcePoolId).Execute()

Deletes a Resource Pool



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
	resourcePoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcePoolAPI.DeleteResourcePool(context.Background(), resourcePoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.DeleteResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourcePoolRequest struct via the builder pattern


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


## GetResourcePool

> ResourcePool GetResourcePool(ctx, resourcePoolId).Execute()

Get Resource Pool information



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
	resourcePoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.GetResourcePool(context.Background(), resourcePoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.GetResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourcePool`: ResourcePool
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.GetResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourcePool**](ResourcePool.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcePoolServers

> []float32 GetResourcePoolServers(ctx, resourcePoolId).Execute()

Get all servers that are part of a Resource Pool



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
	resourcePoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.GetResourcePoolServers(context.Background(), resourcePoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.GetResourcePoolServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourcePoolServers`: []float32
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.GetResourcePoolServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePoolServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]float32**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcePoolSubnetPools

> []float32 GetResourcePoolSubnetPools(ctx, resourcePoolId).Execute()

Get all subnet pools that are part of a resource pool



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
	resourcePoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.GetResourcePoolSubnetPools(context.Background(), resourcePoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.GetResourcePoolSubnetPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourcePoolSubnetPools`: []float32
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.GetResourcePoolSubnetPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePoolSubnetPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]float32**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcePoolUsers

> []map[string]interface{} GetResourcePoolUsers(ctx, resourcePoolId).Execute()

Get all users that have access to a Resource Pool



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
	resourcePoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.GetResourcePoolUsers(context.Background(), resourcePoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.GetResourcePoolUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourcePoolUsers`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.GetResourcePoolUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePoolUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcePools

> ResourcePoolPaginatedList GetResourcePools(ctx).Page(page).Limit(limit).FilterResourcePoolId(filterResourcePoolId).FilterResourcePoolLabel(filterResourcePoolLabel).FilterResourcePoolDescription(filterResourcePoolDescription).FilterResourcePoolCreatedTimestamp(filterResourcePoolCreatedTimestamp).FilterResourcePoolUpdatedTimestamp(filterResourcePoolUpdatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Resource Pools



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
	filterResourcePoolId := []string{"Inner_example"} // []string | Filter by resourcePoolId query param.           <p>              <b>Format: </b> filter.resourcePoolId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.resourcePoolId=$not:$like:John Doe&filter.resourcePoolId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterResourcePoolLabel := []string{"Inner_example"} // []string | Filter by resourcePoolLabel query param.           <p>              <b>Format: </b> filter.resourcePoolLabel={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.resourcePoolLabel=$not:$like:John Doe&filter.resourcePoolLabel=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterResourcePoolDescription := []string{"Inner_example"} // []string | Filter by resourcePoolDescription query param.           <p>              <b>Format: </b> filter.resourcePoolDescription={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.resourcePoolDescription=$not:$like:John Doe&filter.resourcePoolDescription=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterResourcePoolCreatedTimestamp := []string{"Inner_example"} // []string | Filter by resourcePoolCreatedTimestamp query param.           <p>              <b>Format: </b> filter.resourcePoolCreatedTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.resourcePoolCreatedTimestamp=$not:$like:John Doe&filter.resourcePoolCreatedTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterResourcePoolUpdatedTimestamp := []string{"Inner_example"} // []string | Filter by resourcePoolUpdatedTimestamp query param.           <p>              <b>Format: </b> filter.resourcePoolUpdatedTimestamp={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.resourcePoolUpdatedTimestamp=$not:$like:John Doe&filter.resourcePoolUpdatedTimestamp=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> resourcePoolId:DESC           </p>       <h4>Available Fields</h4><ul><li>resourcePoolId</li> <li>resourcePoolLabel</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> resourcePoolLabel,resourcePoolDescription           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>resourcePoolLabel</li> <li>resourcePoolDescription</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.GetResourcePools(context.Background()).Page(page).Limit(limit).FilterResourcePoolId(filterResourcePoolId).FilterResourcePoolLabel(filterResourcePoolLabel).FilterResourcePoolDescription(filterResourcePoolDescription).FilterResourcePoolCreatedTimestamp(filterResourcePoolCreatedTimestamp).FilterResourcePoolUpdatedTimestamp(filterResourcePoolUpdatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.GetResourcePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourcePools`: ResourcePoolPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.GetResourcePools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterResourcePoolId** | **[]string** | Filter by resourcePoolId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.resourcePoolId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.resourcePoolId&#x3D;$not:$like:John Doe&amp;filter.resourcePoolId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterResourcePoolLabel** | **[]string** | Filter by resourcePoolLabel query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.resourcePoolLabel&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.resourcePoolLabel&#x3D;$not:$like:John Doe&amp;filter.resourcePoolLabel&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterResourcePoolDescription** | **[]string** | Filter by resourcePoolDescription query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.resourcePoolDescription&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.resourcePoolDescription&#x3D;$not:$like:John Doe&amp;filter.resourcePoolDescription&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterResourcePoolCreatedTimestamp** | **[]string** | Filter by resourcePoolCreatedTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.resourcePoolCreatedTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.resourcePoolCreatedTimestamp&#x3D;$not:$like:John Doe&amp;filter.resourcePoolCreatedTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterResourcePoolUpdatedTimestamp** | **[]string** | Filter by resourcePoolUpdatedTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.resourcePoolUpdatedTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.resourcePoolUpdatedTimestamp&#x3D;$not:$like:John Doe&amp;filter.resourcePoolUpdatedTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; resourcePoolId:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;resourcePoolId&lt;/li&gt; &lt;li&gt;resourcePoolLabel&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; resourcePoolLabel,resourcePoolDescription           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;resourcePoolLabel&lt;/li&gt; &lt;li&gt;resourcePoolDescription&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**ResourcePoolPaginatedList**](ResourcePoolPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserResourcePools

> []map[string]interface{} GetUserResourcePools(ctx, userId).Execute()

Get all Resource Pools that a user has access to



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
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.GetUserResourcePools(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.GetUserResourcePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserResourcePools`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.GetUserResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveResourcePoolUser

> RemoveResourcePoolUser(ctx, resourcePoolId, userId).Execute()

Remove a user from a Resource Pool



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
	resourcePoolId := float32(8.14) // float32 | 
	userId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcePoolAPI.RemoveResourcePoolUser(context.Background(), resourcePoolId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.RemoveResourcePoolUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveResourcePoolUserRequest struct via the builder pattern


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


## RemoveServerFromResourcePool

> RemoveServerFromResourcePool(ctx, resourcePoolId, serverId).Execute()

Remove a server from a Resource Pool



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
	resourcePoolId := float32(8.14) // float32 | 
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcePoolAPI.RemoveServerFromResourcePool(context.Background(), resourcePoolId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.RemoveServerFromResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServerFromResourcePoolRequest struct via the builder pattern


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


## RemoveSubnetPoolFromResourcePool

> RemoveSubnetPoolFromResourcePool(ctx, resourcePoolId, subnetPoolId).Execute()

Remove a subnet from a resource pool



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
	resourcePoolId := float32(8.14) // float32 | 
	subnetPoolId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcePoolAPI.RemoveSubnetPoolFromResourcePool(context.Background(), resourcePoolId, subnetPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.RemoveSubnetPoolFromResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 
**subnetPoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubnetPoolFromResourcePoolRequest struct via the builder pattern


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


## UpdateResourcePool

> ResourcePool UpdateResourcePool(ctx, resourcePoolId).UpdateResourcePool(updateResourcePool).Execute()

Updates Resource Pool information



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
	resourcePoolId := float32(8.14) // float32 | 
	updateResourcePool := *openapiclient.NewUpdateResourcePool() // UpdateResourcePool | The Resource Pool update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePoolAPI.UpdateResourcePool(context.Background(), resourcePoolId).UpdateResourcePool(updateResourcePool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolAPI.UpdateResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourcePool`: ResourcePool
	fmt.Fprintf(os.Stdout, "Response from `ResourcePoolAPI.UpdateResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePoolId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateResourcePool** | [**UpdateResourcePool**](UpdateResourcePool.md) | The Resource Pool update object | 

### Return type

[**ResourcePool**](ResourcePool.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

