# \CustomIsoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BootCustomIsoIntoServer**](CustomIsoAPI.md#BootCustomIsoIntoServer) | **Post** /api/v2/custom-isos/{customIsoId}/actions/boot-into-server/{serverId} | Boot Custom Iso into Server
[**CreateCustomIso**](CustomIsoAPI.md#CreateCustomIso) | **Post** /api/v2/custom-isos | Creates a Custom Iso.
[**DeleteCustomIso**](CustomIsoAPI.md#DeleteCustomIso) | **Delete** /api/v2/custom-isos/{customIsoId} | Delete Custom Iso
[**GetCustomIso**](CustomIsoAPI.md#GetCustomIso) | **Get** /api/v2/custom-isos/{customIsoId} | Get Custom Iso information
[**GetCustomIsos**](CustomIsoAPI.md#GetCustomIsos) | **Get** /api/v2/custom-isos | Get all Custom Isos
[**MakeCustomIsoPublic**](CustomIsoAPI.md#MakeCustomIsoPublic) | **Post** /api/v2/custom-isos/{customIsoId}/actions/make-public | Make Custom Iso public
[**UpdateCustomIso**](CustomIsoAPI.md#UpdateCustomIso) | **Patch** /api/v2/custom-isos/{customIsoId} | Updates Custom Iso information



## BootCustomIsoIntoServer

> JobInfo BootCustomIsoIntoServer(ctx, customIsoId, serverId).Execute()

Boot Custom Iso into Server

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
	customIsoId := float32(8.14) // float32 | 
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomIsoAPI.BootCustomIsoIntoServer(context.Background(), customIsoId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomIsoAPI.BootCustomIsoIntoServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BootCustomIsoIntoServer`: JobInfo
	fmt.Fprintf(os.Stdout, "Response from `CustomIsoAPI.BootCustomIsoIntoServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIsoId** | **float32** |  | 
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBootCustomIsoIntoServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomIso

> CustomIso CreateCustomIso(ctx).CreateCustomIso(createCustomIso).Execute()

Creates a Custom Iso.

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
	createCustomIso := *openapiclient.NewCreateCustomIso("Name_example", "AccessUrl_example") // CreateCustomIso | Custom Iso create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomIsoAPI.CreateCustomIso(context.Background()).CreateCustomIso(createCustomIso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomIsoAPI.CreateCustomIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomIso`: CustomIso
	fmt.Fprintf(os.Stdout, "Response from `CustomIsoAPI.CreateCustomIso`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomIsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomIso** | [**CreateCustomIso**](CreateCustomIso.md) | Custom Iso create object | 

### Return type

[**CustomIso**](CustomIso.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomIso

> DeleteCustomIso(ctx, customIsoId).Execute()

Delete Custom Iso

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
	customIsoId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomIsoAPI.DeleteCustomIso(context.Background(), customIsoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomIsoAPI.DeleteCustomIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIsoId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomIsoRequest struct via the builder pattern


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


## GetCustomIso

> CustomIso GetCustomIso(ctx, customIsoId).Execute()

Get Custom Iso information



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
	customIsoId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomIsoAPI.GetCustomIso(context.Background(), customIsoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomIsoAPI.GetCustomIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomIso`: CustomIso
	fmt.Fprintf(os.Stdout, "Response from `CustomIsoAPI.GetCustomIso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIsoId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomIsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomIso**](CustomIso.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomIsos

> CustomIsoPaginatedList GetCustomIsos(ctx).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDisplayName(filterDisplayName).FilterAccessUrl(filterAccessUrl).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Custom Isos



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterDisplayName := []string{"Inner_example"} // []string | Filter by displayName query param.           <p>              <b>Format: </b> filter.displayName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.displayName=$not:$like:John Doe&filter.displayName=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAccessUrl := []string{"Inner_example"} // []string | Filter by accessUrl query param.           <p>              <b>Format: </b> filter.accessUrl={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.accessUrl=$not:$like:John Doe&filter.accessUrl=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,name,displayName,accessUrl           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>name</li> <li>displayName</li> <li>accessUrl</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomIsoAPI.GetCustomIsos(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterDisplayName(filterDisplayName).FilterAccessUrl(filterAccessUrl).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomIsoAPI.GetCustomIsos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomIsos`: CustomIsoPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `CustomIsoAPI.GetCustomIsos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomIsosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterDisplayName** | **[]string** | Filter by displayName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.displayName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.displayName&#x3D;$not:$like:John Doe&amp;filter.displayName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterAccessUrl** | **[]string** | Filter by accessUrl query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.accessUrl&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.accessUrl&#x3D;$not:$like:John Doe&amp;filter.accessUrl&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,name,displayName,accessUrl           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;displayName&lt;/li&gt; &lt;li&gt;accessUrl&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**CustomIsoPaginatedList**](CustomIsoPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeCustomIsoPublic

> CustomIso MakeCustomIsoPublic(ctx, customIsoId).Execute()

Make Custom Iso public

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
	customIsoId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomIsoAPI.MakeCustomIsoPublic(context.Background(), customIsoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomIsoAPI.MakeCustomIsoPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeCustomIsoPublic`: CustomIso
	fmt.Fprintf(os.Stdout, "Response from `CustomIsoAPI.MakeCustomIsoPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIsoId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeCustomIsoPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomIso**](CustomIso.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomIso

> CustomIso UpdateCustomIso(ctx, customIsoId).UpdateCustomIso(updateCustomIso).Execute()

Updates Custom Iso information

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
	customIsoId := float32(8.14) // float32 | 
	updateCustomIso := *openapiclient.NewUpdateCustomIso() // UpdateCustomIso | The updated Custom Iso object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomIsoAPI.UpdateCustomIso(context.Background(), customIsoId).UpdateCustomIso(updateCustomIso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomIsoAPI.UpdateCustomIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomIso`: CustomIso
	fmt.Fprintf(os.Stdout, "Response from `CustomIsoAPI.UpdateCustomIso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIsoId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomIsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomIso** | [**UpdateCustomIso**](UpdateCustomIso.md) | The updated Custom Iso object | 

### Return type

[**CustomIso**](CustomIso.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

