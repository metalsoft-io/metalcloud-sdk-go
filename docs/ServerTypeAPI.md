# \ServerTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerType**](ServerTypeAPI.md#CreateServerType) | **Post** /api/v2/server-types | Creates a Server Type
[**DeleteServerType**](ServerTypeAPI.md#DeleteServerType) | **Delete** /api/v2/server-types/{serverTypeId} | Deletes a Server Type
[**GetServerTypeInfo**](ServerTypeAPI.md#GetServerTypeInfo) | **Get** /api/v2/server-types/{serverTypeId} | Get Server Type information
[**GetServerTypes**](ServerTypeAPI.md#GetServerTypes) | **Get** /api/v2/server-types | Get a list of Server Types
[**GetServerTypesStatisticsBatch**](ServerTypeAPI.md#GetServerTypesStatisticsBatch) | **Post** /api/v2/server-types/statistics | Get Server Type statistics batch
[**RemoveUnusedServerTypes**](ServerTypeAPI.md#RemoveUnusedServerTypes) | **Post** /api/v2/server-types/actions/clean-unused | Deletes unused server types
[**UpdateServerType**](ServerTypeAPI.md#UpdateServerType) | **Patch** /api/v2/server-types/{serverTypeId} | Updates a Server Type



## CreateServerType

> ServerType CreateServerType(ctx).CreateServerType(createServerType).Execute()

Creates a Server Type



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
	createServerType := *openapiclient.NewCreateServerType(float32(123), float32(123), float32(123), float32(123), "Name_example", "Label_example", []float32{float32(123)}, []string{"ProcessorNames_example"}, float32(123), "ServerClass_example") // CreateServerType | The Server Type create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.CreateServerType(context.Background()).CreateServerType(createServerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.CreateServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerType`: ServerType
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.CreateServerType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerType** | [**CreateServerType**](CreateServerType.md) | The Server Type create object | 

### Return type

[**ServerType**](ServerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerType

> DeleteServerType(ctx, serverTypeId).Execute()

Deletes a Server Type



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
	serverTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerTypeAPI.DeleteServerType(context.Background(), serverTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.DeleteServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerTypeRequest struct via the builder pattern


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


## GetServerTypeInfo

> ServerType GetServerTypeInfo(ctx, serverTypeId).Execute()

Get Server Type information



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
	serverTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.GetServerTypeInfo(context.Background(), serverTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.GetServerTypeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerTypeInfo`: ServerType
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.GetServerTypeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerTypeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerType**](ServerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerTypes

> ServerTypePaginatedList GetServerTypes(ctx).Page(page).Limit(limit).FilterName(filterName).FilterLabel(filterLabel).FilterRamGbytes(filterRamGbytes).FilterProcessorCount(filterProcessorCount).FilterProcessorCoreCount(filterProcessorCoreCount).FilterDescription(filterDescription).FilterNetworkTotalCapacityMbps(filterNetworkTotalCapacityMbps).FilterIsExperimental(filterIsExperimental).FilterForUnmanagedServersOnly(filterForUnmanagedServersOnly).FilterForGenericEndpointsOnly(filterForGenericEndpointsOnly).FilterSupportsOobProvisioning(filterSupportsOobProvisioning).FilterAllowedVendorSkuIds(filterAllowedVendorSkuIds).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Server Types



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
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterRamGbytes := []string{"Inner_example"} // []string | Filter by ramGbytes query param.           <p>              <b>Format: </b> filter.ramGbytes={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.ramGbytes=$not:$like:John Doe&filter.ramGbytes=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterProcessorCount := []string{"Inner_example"} // []string | Filter by processorCount query param.           <p>              <b>Format: </b> filter.processorCount={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.processorCount=$not:$like:John Doe&filter.processorCount=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterProcessorCoreCount := []string{"Inner_example"} // []string | Filter by processorCoreCount query param.           <p>              <b>Format: </b> filter.processorCoreCount={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.processorCoreCount=$not:$like:John Doe&filter.processorCoreCount=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.           <p>              <b>Format: </b> filter.description={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.description=$not:$like:John Doe&filter.description=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterNetworkTotalCapacityMbps := []string{"Inner_example"} // []string | Filter by networkTotalCapacityMbps query param.           <p>              <b>Format: </b> filter.networkTotalCapacityMbps={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.networkTotalCapacityMbps=$not:$like:John Doe&filter.networkTotalCapacityMbps=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterIsExperimental := []string{"Inner_example"} // []string | Filter by isExperimental query param.           <p>              <b>Format: </b> filter.isExperimental={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.isExperimental=$not:$like:John Doe&filter.isExperimental=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterForUnmanagedServersOnly := []string{"Inner_example"} // []string | Filter by forUnmanagedServersOnly query param.           <p>              <b>Format: </b> filter.forUnmanagedServersOnly={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.forUnmanagedServersOnly=$not:$like:John Doe&filter.forUnmanagedServersOnly=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterForGenericEndpointsOnly := []string{"Inner_example"} // []string | Filter by forGenericEndpointsOnly query param.           <p>              <b>Format: </b> filter.forGenericEndpointsOnly={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.forGenericEndpointsOnly=$not:$like:John Doe&filter.forGenericEndpointsOnly=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSupportsOobProvisioning := []string{"Inner_example"} // []string | Filter by supportsOobProvisioning query param.           <p>              <b>Format: </b> filter.supportsOobProvisioning={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.supportsOobProvisioning=$not:$like:John Doe&filter.supportsOobProvisioning=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAllowedVendorSkuIds := []string{"Inner_example"} // []string | Filter by allowedVendorSkuIds query param.           <p>              <b>Format: </b> filter.allowedVendorSkuIds={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.allowedVendorSkuIds=$not:$like:John Doe&filter.allowedVendorSkuIds=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>name</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> name,label,description           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>name</li> <li>label</li> <li>description</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.GetServerTypes(context.Background()).Page(page).Limit(limit).FilterName(filterName).FilterLabel(filterLabel).FilterRamGbytes(filterRamGbytes).FilterProcessorCount(filterProcessorCount).FilterProcessorCoreCount(filterProcessorCoreCount).FilterDescription(filterDescription).FilterNetworkTotalCapacityMbps(filterNetworkTotalCapacityMbps).FilterIsExperimental(filterIsExperimental).FilterForUnmanagedServersOnly(filterForUnmanagedServersOnly).FilterForGenericEndpointsOnly(filterForGenericEndpointsOnly).FilterSupportsOobProvisioning(filterSupportsOobProvisioning).FilterAllowedVendorSkuIds(filterAllowedVendorSkuIds).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.GetServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerTypes`: ServerTypePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.GetServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterRamGbytes** | **[]string** | Filter by ramGbytes query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.ramGbytes&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.ramGbytes&#x3D;$not:$like:John Doe&amp;filter.ramGbytes&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterProcessorCount** | **[]string** | Filter by processorCount query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.processorCount&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.processorCount&#x3D;$not:$like:John Doe&amp;filter.processorCount&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterProcessorCoreCount** | **[]string** | Filter by processorCoreCount query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.processorCoreCount&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.processorCoreCount&#x3D;$not:$like:John Doe&amp;filter.processorCoreCount&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterDescription** | **[]string** | Filter by description query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.description&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.description&#x3D;$not:$like:John Doe&amp;filter.description&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterNetworkTotalCapacityMbps** | **[]string** | Filter by networkTotalCapacityMbps query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.networkTotalCapacityMbps&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.networkTotalCapacityMbps&#x3D;$not:$like:John Doe&amp;filter.networkTotalCapacityMbps&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterIsExperimental** | **[]string** | Filter by isExperimental query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.isExperimental&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.isExperimental&#x3D;$not:$like:John Doe&amp;filter.isExperimental&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterForUnmanagedServersOnly** | **[]string** | Filter by forUnmanagedServersOnly query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.forUnmanagedServersOnly&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.forUnmanagedServersOnly&#x3D;$not:$like:John Doe&amp;filter.forUnmanagedServersOnly&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterForGenericEndpointsOnly** | **[]string** | Filter by forGenericEndpointsOnly query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.forGenericEndpointsOnly&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.forGenericEndpointsOnly&#x3D;$not:$like:John Doe&amp;filter.forGenericEndpointsOnly&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSupportsOobProvisioning** | **[]string** | Filter by supportsOobProvisioning query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.supportsOobProvisioning&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.supportsOobProvisioning&#x3D;$not:$like:John Doe&amp;filter.supportsOobProvisioning&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterAllowedVendorSkuIds** | **[]string** | Filter by allowedVendorSkuIds query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.allowedVendorSkuIds&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.allowedVendorSkuIds&#x3D;$not:$like:John Doe&amp;filter.allowedVendorSkuIds&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name,label,description           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;description&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**ServerTypePaginatedList**](ServerTypePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerTypesStatisticsBatch

> ServerTypeStatisticsBatch GetServerTypesStatisticsBatch(ctx).ServerTypeStatisticsBatchOptions(serverTypeStatisticsBatchOptions).Execute()

Get Server Type statistics batch



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
	serverTypeStatisticsBatchOptions := *openapiclient.NewServerTypeStatisticsBatchOptions(float32(123)) // ServerTypeStatisticsBatchOptions | The server type statistics batch options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.GetServerTypesStatisticsBatch(context.Background()).ServerTypeStatisticsBatchOptions(serverTypeStatisticsBatchOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.GetServerTypesStatisticsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerTypesStatisticsBatch`: ServerTypeStatisticsBatch
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.GetServerTypesStatisticsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerTypesStatisticsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverTypeStatisticsBatchOptions** | [**ServerTypeStatisticsBatchOptions**](ServerTypeStatisticsBatchOptions.md) | The server type statistics batch options | 

### Return type

[**ServerTypeStatisticsBatch**](ServerTypeStatisticsBatch.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUnusedServerTypes

> RemoveUnusedServerTypes(ctx).Execute()

Deletes unused server types



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
	r, err := apiClient.ServerTypeAPI.RemoveUnusedServerTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.RemoveUnusedServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUnusedServerTypesRequest struct via the builder pattern


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


## UpdateServerType

> ServerType UpdateServerType(ctx, serverTypeId).UpdateServerType(updateServerType).Execute()

Updates a Server Type



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
	serverTypeId := float32(8.14) // float32 | 
	updateServerType := *openapiclient.NewUpdateServerType("Label_example") // UpdateServerType | The Server Type update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTypeAPI.UpdateServerType(context.Background(), serverTypeId).UpdateServerType(updateServerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTypeAPI.UpdateServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerType`: ServerType
	fmt.Fprintf(os.Stdout, "Response from `ServerTypeAPI.UpdateServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerType** | [**UpdateServerType**](UpdateServerType.md) | The Server Type update object | 

### Return type

[**ServerType**](ServerType.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

