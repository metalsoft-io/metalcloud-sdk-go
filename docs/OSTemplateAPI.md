# \OSTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSTemplate**](OSTemplateAPI.md#CreateOSTemplate) | **Post** /api/v2/os-templates | Create OS template
[**DeleteOSTemplate**](OSTemplateAPI.md#DeleteOSTemplate) | **Delete** /api/v2/os-templates/{osTemplateId} | Delete OS template
[**GetOSTemplate**](OSTemplateAPI.md#GetOSTemplate) | **Get** /api/v2/os-templates/{osTemplateId} | Get details for an OS template
[**GetOSTemplates**](OSTemplateAPI.md#GetOSTemplates) | **Get** /api/v2/os-templates | Get a list of available OS templates
[**UpdateOSTemplate**](OSTemplateAPI.md#UpdateOSTemplate) | **Put** /api/v2/os-templates/{osTemplateId} | Update OS template



## CreateOSTemplate

> OSTemplate CreateOSTemplate(ctx).OSTemplateCreate(oSTemplateCreate).Execute()

Create OS template



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
	oSTemplateCreate := *openapiclient.NewOSTemplateCreate("Ubuntu 24.04 LTS", *openapiclient.NewOSTemplateDevice("server", "uefi", "x86_64"), *openapiclient.NewOSTemplateInstall("oob", "local_drive", "wait_for_power_off"), *openapiclient.NewOSTemplateImageBuild(true), *openapiclient.NewOSTemplateOs("Ubuntu | CentOS | Suse | Windows | ESXi | ...", "22.04 | 8 | 15 | 2019 | ...", *openapiclient.NewOSTemplateOsCredential("root", "plain"))) // OSTemplateCreate | The OS template details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.CreateOSTemplate(context.Background()).OSTemplateCreate(oSTemplateCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.CreateOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSTemplate`: OSTemplate
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.CreateOSTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSTemplateCreate** | [**OSTemplateCreate**](OSTemplateCreate.md) | The OS template details | 

### Return type

[**OSTemplate**](OSTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSTemplate

> DeleteOSTemplate(ctx, osTemplateId).Execute()

Delete OS template



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
	osTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSTemplateAPI.DeleteOSTemplate(context.Background(), osTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.DeleteOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSTemplateRequest struct via the builder pattern


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


## GetOSTemplate

> OSTemplate GetOSTemplate(ctx, osTemplateId).Execute()

Get details for an OS template



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
	osTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.GetOSTemplate(context.Background(), osTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.GetOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSTemplate`: OSTemplate
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.GetOSTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSTemplate**](OSTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSTemplates

> OSTemplateList GetOSTemplates(ctx).Page(page).Limit(limit).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of available OS templates



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
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> -1           </p>       <p>              <b>Max Value: </b> -1           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.           <p>              <b>Format: </b> filter.status={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.status=$not:$like:John Doe&filter.status=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$null</li> <li>$ilike</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>device.type</li> <li>os.name</li> <li>visibility</li> <li>status</li> <li>createdBy</li> <li>createdAt</li> <li>modifiedAt</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,device.type           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>device.type</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.GetOSTemplates(context.Background()).Page(page).Limit(limit).FilterStatus(filterStatus).FilterName(filterName).FilterLabel(filterLabel).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.GetOSTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSTemplates`: OSTemplateList
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.GetOSTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOSTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; -1           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; -1           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterStatus** | **[]string** | Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;device.type&lt;/li&gt; &lt;li&gt;os.name&lt;/li&gt; &lt;li&gt;visibility&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;createdBy&lt;/li&gt; &lt;li&gt;createdAt&lt;/li&gt; &lt;li&gt;modifiedAt&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,device.type           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;device.type&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**OSTemplateList**](OSTemplateList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSTemplate

> OSTemplate UpdateOSTemplate(ctx, osTemplateId).OSTemplateUpdate(oSTemplateUpdate).IfMatch(ifMatch).Execute()

Update OS template



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
	osTemplateId := float32(8.14) // float32 | 
	oSTemplateUpdate := *openapiclient.NewOSTemplateUpdate("Ubuntu 24.04 LTS", *openapiclient.NewOSTemplateDevice("server", "uefi", "x86_64"), *openapiclient.NewOSTemplateInstall("oob", "local_drive", "wait_for_power_off"), *openapiclient.NewOSTemplateImageBuild(true), *openapiclient.NewOSTemplateOs("Ubuntu | CentOS | Suse | Windows | ESXi | ...", "22.04 | 8 | 15 | 2019 | ...", *openapiclient.NewOSTemplateOsCredential("root", "plain"))) // OSTemplateUpdate | The OS template details
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.UpdateOSTemplate(context.Background(), osTemplateId).OSTemplateUpdate(oSTemplateUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.UpdateOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSTemplate`: OSTemplate
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.UpdateOSTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oSTemplateUpdate** | [**OSTemplateUpdate**](OSTemplateUpdate.md) | The OS template details | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**OSTemplate**](OSTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

