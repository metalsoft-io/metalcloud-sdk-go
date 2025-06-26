# \BGPTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpTemplate**](BGPTemplateAPI.md#CreateBgpTemplate) | **Post** /api/v2/bgp-templates | Creates a BGP Template
[**DeleteBgpTemplate**](BGPTemplateAPI.md#DeleteBgpTemplate) | **Delete** /api/v2/bgp-templates/{bgpTemplateId} | Deletes a BGP Template
[**GetBgpTemplate**](BGPTemplateAPI.md#GetBgpTemplate) | **Get** /api/v2/bgp-templates/{bgpTemplateId} | Get BGP Template information
[**GetBgpTemplates**](BGPTemplateAPI.md#GetBgpTemplates) | **Get** /api/v2/bgp-templates | Get all BGP Templates
[**UpdateBgpTemplate**](BGPTemplateAPI.md#UpdateBgpTemplate) | **Patch** /api/v2/bgp-templates/{bgpTemplateId} | Updates BGP Template information



## CreateBgpTemplate

> BgpTemplate CreateBgpTemplate(ctx).CreateBgpTemplate(createBgpTemplate).Execute()

Creates a BGP Template



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
	createBgpTemplate := *openapiclient.NewCreateBgpTemplate("NetworkType_example", "NetworkDeviceDriver_example", "NetworkDevicePosition_example", "RemoteNetworkDevicePosition_example", float32(123), "BgpNumbering_example", "BgpLinkConfiguration_example", "ExecutionType_example", "Configuration_example") // CreateBgpTemplate | The BGP Template create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPTemplateAPI.CreateBgpTemplate(context.Background()).CreateBgpTemplate(createBgpTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTemplateAPI.CreateBgpTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBgpTemplate`: BgpTemplate
	fmt.Fprintf(os.Stdout, "Response from `BGPTemplateAPI.CreateBgpTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBgpTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBgpTemplate** | [**CreateBgpTemplate**](CreateBgpTemplate.md) | The BGP Template create object | 

### Return type

[**BgpTemplate**](BgpTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBgpTemplate

> DeleteBgpTemplate(ctx, bgpTemplateId).Execute()

Deletes a BGP Template



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
	bgpTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BGPTemplateAPI.DeleteBgpTemplate(context.Background(), bgpTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTemplateAPI.DeleteBgpTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBgpTemplateRequest struct via the builder pattern


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


## GetBgpTemplate

> BgpTemplate GetBgpTemplate(ctx, bgpTemplateId).Execute()

Get BGP Template information



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
	bgpTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPTemplateAPI.GetBgpTemplate(context.Background(), bgpTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTemplateAPI.GetBgpTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBgpTemplate`: BgpTemplate
	fmt.Fprintf(os.Stdout, "Response from `BGPTemplateAPI.GetBgpTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpTemplate**](BgpTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBgpTemplates

> BgpTemplatePaginatedList GetBgpTemplates(ctx).Page(page).Limit(limit).FilterId(filterId).FilterNetworkType(filterNetworkType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterNetworkDevicePosition(filterNetworkDevicePosition).FilterRemoteNetworkDevicePosition(filterRemoteNetworkDevicePosition).FilterMlagPair(filterMlagPair).FilterBgpNumbering(filterBgpNumbering).FilterBgpLinkConfiguration(filterBgpLinkConfiguration).FilterExecutionType(filterExecutionType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all BGP Templates



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterNetworkType := []string{"Inner_example"} // []string | Filter by networkType query param.           <p>              <b>Format: </b> filter.networkType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.networkType=$not:$like:John Doe&filter.networkType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterNetworkDeviceDriver := []string{"Inner_example"} // []string | Filter by networkDeviceDriver query param.           <p>              <b>Format: </b> filter.networkDeviceDriver={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.networkDeviceDriver=$not:$like:John Doe&filter.networkDeviceDriver=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterNetworkDevicePosition := []string{"Inner_example"} // []string | Filter by networkDevicePosition query param.           <p>              <b>Format: </b> filter.networkDevicePosition={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.networkDevicePosition=$not:$like:John Doe&filter.networkDevicePosition=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterRemoteNetworkDevicePosition := []string{"Inner_example"} // []string | Filter by remoteNetworkDevicePosition query param.           <p>              <b>Format: </b> filter.remoteNetworkDevicePosition={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.remoteNetworkDevicePosition=$not:$like:John Doe&filter.remoteNetworkDevicePosition=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterMlagPair := []string{"Inner_example"} // []string | Filter by mlagPair query param.           <p>              <b>Format: </b> filter.mlagPair={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.mlagPair=$not:$like:John Doe&filter.mlagPair=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterBgpNumbering := []string{"Inner_example"} // []string | Filter by bgpNumbering query param.           <p>              <b>Format: </b> filter.bgpNumbering={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.bgpNumbering=$not:$like:John Doe&filter.bgpNumbering=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterBgpLinkConfiguration := []string{"Inner_example"} // []string | Filter by bgpLinkConfiguration query param.           <p>              <b>Format: </b> filter.bgpLinkConfiguration={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.bgpLinkConfiguration=$not:$like:John Doe&filter.bgpLinkConfiguration=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterExecutionType := []string{"Inner_example"} // []string | Filter by executionType query param.           <p>              <b>Format: </b> filter.executionType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.executionType=$not:$like:John Doe&filter.executionType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>networkType</li> <li>networkDeviceDriver</li> <li>networkDevicePosition</li> <li>remoteNetworkDevicePosition</li> <li>mlagPair</li> <li>bgpNumbering</li> <li>bgpLinkConfiguration</li> <li>executionType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,networkType,networkDeviceDriver,networkDevicePosition,remoteNetworkDevicePosition           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>networkType</li> <li>networkDeviceDriver</li> <li>networkDevicePosition</li> <li>remoteNetworkDevicePosition</li> <li>mlagPair</li> <li>bgpNumbering</li> <li>bgpLinkConfiguration</li> <li>executionType</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPTemplateAPI.GetBgpTemplates(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterNetworkType(filterNetworkType).FilterNetworkDeviceDriver(filterNetworkDeviceDriver).FilterNetworkDevicePosition(filterNetworkDevicePosition).FilterRemoteNetworkDevicePosition(filterRemoteNetworkDevicePosition).FilterMlagPair(filterMlagPair).FilterBgpNumbering(filterBgpNumbering).FilterBgpLinkConfiguration(filterBgpLinkConfiguration).FilterExecutionType(filterExecutionType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTemplateAPI.GetBgpTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBgpTemplates`: BgpTemplatePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `BGPTemplateAPI.GetBgpTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterNetworkType** | **[]string** | Filter by networkType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.networkType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.networkType&#x3D;$not:$like:John Doe&amp;filter.networkType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterNetworkDeviceDriver** | **[]string** | Filter by networkDeviceDriver query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.networkDeviceDriver&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.networkDeviceDriver&#x3D;$not:$like:John Doe&amp;filter.networkDeviceDriver&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterNetworkDevicePosition** | **[]string** | Filter by networkDevicePosition query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.networkDevicePosition&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.networkDevicePosition&#x3D;$not:$like:John Doe&amp;filter.networkDevicePosition&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterRemoteNetworkDevicePosition** | **[]string** | Filter by remoteNetworkDevicePosition query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.remoteNetworkDevicePosition&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.remoteNetworkDevicePosition&#x3D;$not:$like:John Doe&amp;filter.remoteNetworkDevicePosition&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterMlagPair** | **[]string** | Filter by mlagPair query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.mlagPair&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.mlagPair&#x3D;$not:$like:John Doe&amp;filter.mlagPair&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterBgpNumbering** | **[]string** | Filter by bgpNumbering query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.bgpNumbering&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.bgpNumbering&#x3D;$not:$like:John Doe&amp;filter.bgpNumbering&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterBgpLinkConfiguration** | **[]string** | Filter by bgpLinkConfiguration query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.bgpLinkConfiguration&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.bgpLinkConfiguration&#x3D;$not:$like:John Doe&amp;filter.bgpLinkConfiguration&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterExecutionType** | **[]string** | Filter by executionType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.executionType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.executionType&#x3D;$not:$like:John Doe&amp;filter.executionType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;networkType&lt;/li&gt; &lt;li&gt;networkDeviceDriver&lt;/li&gt; &lt;li&gt;networkDevicePosition&lt;/li&gt; &lt;li&gt;remoteNetworkDevicePosition&lt;/li&gt; &lt;li&gt;mlagPair&lt;/li&gt; &lt;li&gt;bgpNumbering&lt;/li&gt; &lt;li&gt;bgpLinkConfiguration&lt;/li&gt; &lt;li&gt;executionType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,networkType,networkDeviceDriver,networkDevicePosition,remoteNetworkDevicePosition           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;networkType&lt;/li&gt; &lt;li&gt;networkDeviceDriver&lt;/li&gt; &lt;li&gt;networkDevicePosition&lt;/li&gt; &lt;li&gt;remoteNetworkDevicePosition&lt;/li&gt; &lt;li&gt;mlagPair&lt;/li&gt; &lt;li&gt;bgpNumbering&lt;/li&gt; &lt;li&gt;bgpLinkConfiguration&lt;/li&gt; &lt;li&gt;executionType&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**BgpTemplatePaginatedList**](BgpTemplatePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBgpTemplate

> BgpTemplate UpdateBgpTemplate(ctx, bgpTemplateId).UpdateBgpTemplate(updateBgpTemplate).Execute()

Updates BGP Template information



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
	bgpTemplateId := float32(8.14) // float32 | 
	updateBgpTemplate := *openapiclient.NewUpdateBgpTemplate() // UpdateBgpTemplate | The BGP Template update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPTemplateAPI.UpdateBgpTemplate(context.Background(), bgpTemplateId).UpdateBgpTemplate(updateBgpTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTemplateAPI.UpdateBgpTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBgpTemplate`: BgpTemplate
	fmt.Fprintf(os.Stdout, "Response from `BGPTemplateAPI.UpdateBgpTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBgpTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBgpTemplate** | [**UpdateBgpTemplate**](UpdateBgpTemplate.md) | The BGP Template update object | 

### Return type

[**BgpTemplate**](BgpTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

