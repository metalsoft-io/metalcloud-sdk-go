# \EndpointInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpointInstance**](EndpointInstanceAPI.md#CreateEndpointInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/endpoint-instances | Add Endpoint Instance to an infrastructure
[**DeleteEndpointInstance**](EndpointInstanceAPI.md#DeleteEndpointInstance) | **Delete** /api/v2/endpoint-instances/{endpointInstanceId} | Delete Endpoint Instance
[**GetEndpointInstance**](EndpointInstanceAPI.md#GetEndpointInstance) | **Get** /api/v2/endpoint-instances/{endpointInstanceId} | Get Endpoint Instance details
[**GetEndpointInstanceConfig**](EndpointInstanceAPI.md#GetEndpointInstanceConfig) | **Get** /api/v2/endpoint-instances/{endpointInstanceId}/config | Get Endpoint Instance config details
[**GetEndpointInstances**](EndpointInstanceAPI.md#GetEndpointInstances) | **Get** /api/v2/endpoint-instances | List Endpoint Instances
[**GetInfrastructureEndpointInstances**](EndpointInstanceAPI.md#GetInfrastructureEndpointInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/endpoint-instances | List Endpoint Instances for an infrastructure
[**UpdateEndpointInstanceConfig**](EndpointInstanceAPI.md#UpdateEndpointInstanceConfig) | **Patch** /api/v2/endpoint-instances/{endpointInstanceId}/config | Update Endpoint Instance configuration
[**UpdateEndpointInstanceMeta**](EndpointInstanceAPI.md#UpdateEndpointInstanceMeta) | **Patch** /api/v2/endpoint-instances/{endpointInstanceId}/meta | Update an Endpoint Instance meta information



## CreateEndpointInstance

> EndpointInstance CreateEndpointInstance(ctx, infrastructureId).EndpointInstanceCreate(endpointInstanceCreate).Execute()

Add Endpoint Instance to an infrastructure



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
	infrastructureId := int32(56) // int32 | 
	endpointInstanceCreate := *openapiclient.NewEndpointInstanceCreate(int32(123)) // EndpointInstanceCreate | The Endpoint Instance to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.CreateEndpointInstance(context.Background(), infrastructureId).EndpointInstanceCreate(endpointInstanceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.CreateEndpointInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointInstance`: EndpointInstance
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.CreateEndpointInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointInstanceCreate** | [**EndpointInstanceCreate**](EndpointInstanceCreate.md) | The Endpoint Instance to create | 

### Return type

[**EndpointInstance**](EndpointInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpointInstance

> DeleteEndpointInstance(ctx, endpointInstanceId).IfMatch(ifMatch).Execute()

Delete Endpoint Instance



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
	endpointInstanceId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceAPI.DeleteEndpointInstance(context.Background(), endpointInstanceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.DeleteEndpointInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointInstanceRequest struct via the builder pattern


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


## GetEndpointInstance

> EndpointInstance GetEndpointInstance(ctx, endpointInstanceId).Execute()

Get Endpoint Instance details



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
	endpointInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetEndpointInstance(context.Background(), endpointInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetEndpointInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstance`: EndpointInstance
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetEndpointInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointInstance**](EndpointInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstanceConfig

> EndpointInstanceConfiguration GetEndpointInstanceConfig(ctx, endpointInstanceId).Execute()

Get Endpoint Instance config details



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
	endpointInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetEndpointInstanceConfig(context.Background(), endpointInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetEndpointInstanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstanceConfig`: EndpointInstanceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetEndpointInstanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointInstanceConfiguration**](EndpointInstanceConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointInstances

> EndpointInstancePaginatedList GetEndpointInstances(ctx).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Endpoint Instances



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
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterGroupId := []string{"Inner_example"} // []string | Filter by groupId query param.           <p>              <b>Format: </b> filter.groupId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.groupId=$not:$like:John Doe&filter.groupId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterEndpointId := []string{"Inner_example"} // []string | Filter by endpointId query param.           <p>              <b>Format: </b> filter.endpointId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.endpointId=$not:$like:John Doe&filter.endpointId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigEndpointId := []string{"Inner_example"} // []string | Filter by config.endpointId query param.           <p>              <b>Format: </b> filter.config.endpointId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.endpointId=$not:$like:John Doe&filter.config.endpointId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetEndpointInstances(context.Background()).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetEndpointInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInstances`: EndpointInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetEndpointInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterGroupId** | **[]string** | Filter by groupId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.groupId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.groupId&#x3D;$not:$like:John Doe&amp;filter.groupId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterEndpointId** | **[]string** | Filter by endpointId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.endpointId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.endpointId&#x3D;$not:$like:John Doe&amp;filter.endpointId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigEndpointId** | **[]string** | Filter by config.endpointId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.endpointId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.endpointId&#x3D;$not:$like:John Doe&amp;filter.config.endpointId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**EndpointInstancePaginatedList**](EndpointInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureEndpointInstances

> EndpointInstancePaginatedList GetInfrastructureEndpointInstances(ctx, infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List Endpoint Instances for an infrastructure



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
	infrastructureId := int32(56) // int32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterGroupId := []string{"Inner_example"} // []string | Filter by groupId query param.           <p>              <b>Format: </b> filter.groupId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.groupId=$not:$like:John Doe&filter.groupId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterEndpointId := []string{"Inner_example"} // []string | Filter by endpointId query param.           <p>              <b>Format: </b> filter.endpointId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.endpointId=$not:$like:John Doe&filter.endpointId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigEndpointId := []string{"Inner_example"} // []string | Filter by config.endpointId query param.           <p>              <b>Format: </b> filter.config.endpointId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.endpointId=$not:$like:John Doe&filter.config.endpointId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.GetInfrastructureEndpointInstances(context.Background(), infrastructureId).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterGroupId(filterGroupId).FilterEndpointId(filterEndpointId).FilterServiceStatus(filterServiceStatus).FilterConfigEndpointId(filterConfigEndpointId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.GetInfrastructureEndpointInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureEndpointInstances`: EndpointInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.GetInfrastructureEndpointInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureEndpointInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterGroupId** | **[]string** | Filter by groupId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.groupId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.groupId&#x3D;$not:$like:John Doe&amp;filter.groupId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterEndpointId** | **[]string** | Filter by endpointId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.endpointId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.endpointId&#x3D;$not:$like:John Doe&amp;filter.endpointId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigEndpointId** | **[]string** | Filter by config.endpointId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.endpointId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.endpointId&#x3D;$not:$like:John Doe&amp;filter.config.endpointId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**EndpointInstancePaginatedList**](EndpointInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceConfig

> EndpointInstanceConfiguration UpdateEndpointInstanceConfig(ctx, endpointInstanceId).EndpointInstanceUpdate(endpointInstanceUpdate).IfMatch(ifMatch).Execute()

Update Endpoint Instance configuration



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
	endpointInstanceId := int32(56) // int32 | 
	endpointInstanceUpdate := *openapiclient.NewEndpointInstanceUpdate() // EndpointInstanceUpdate | The Endpoint Instance configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointInstanceAPI.UpdateEndpointInstanceConfig(context.Background(), endpointInstanceId).EndpointInstanceUpdate(endpointInstanceUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.UpdateEndpointInstanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointInstanceConfig`: EndpointInstanceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EndpointInstanceAPI.UpdateEndpointInstanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointInstanceUpdate** | [**EndpointInstanceUpdate**](EndpointInstanceUpdate.md) | The Endpoint Instance configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**EndpointInstanceConfiguration**](EndpointInstanceConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointInstanceMeta

> UpdateEndpointInstanceMeta(ctx, endpointInstanceId).GenericMeta(genericMeta).Execute()

Update an Endpoint Instance meta information



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
	endpointInstanceId := int32(56) // int32 | 
	genericMeta := *openapiclient.NewGenericMeta() // GenericMeta | The Endpoint Instance meta information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointInstanceAPI.UpdateEndpointInstanceMeta(context.Background(), endpointInstanceId).GenericMeta(genericMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointInstanceAPI.UpdateEndpointInstanceMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointInstanceMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genericMeta** | [**GenericMeta**](GenericMeta.md) | The Endpoint Instance meta information | 

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

