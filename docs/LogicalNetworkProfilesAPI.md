# \LogicalNetworkProfilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogicalNetworkProfile**](LogicalNetworkProfilesAPI.md#CreateLogicalNetworkProfile) | **Post** /api/v2/logical-network-profiles | Create a new logical network profile
[**DeleteLogicalNetworkProfile**](LogicalNetworkProfilesAPI.md#DeleteLogicalNetworkProfile) | **Delete** /api/v2/logical-network-profiles/{logicalNetworkProfileId} | Delete a logical network profile
[**GetAllLogicalNetworkProfiles**](LogicalNetworkProfilesAPI.md#GetAllLogicalNetworkProfiles) | **Get** /api/v2/logical-network-profiles | Get all logical network profiles
[**GetLogicalNetworkProfileById**](LogicalNetworkProfilesAPI.md#GetLogicalNetworkProfileById) | **Get** /api/v2/logical-network-profiles/{logicalNetworkProfileId} | Get a logical network profile by ID
[**UpdateLogicalNetworkProfile**](LogicalNetworkProfilesAPI.md#UpdateLogicalNetworkProfile) | **Put** /api/v2/logical-network-profiles/{logicalNetworkProfileId} | Update a logical network profile



## CreateLogicalNetworkProfile

> map[string]interface{} CreateLogicalNetworkProfile(ctx).Body(body).Execute()

Create a new logical network profile



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
	body := map[string]interface{}{ ... } // map[string]interface{} | The Logical Network Profile create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfilesAPI.CreateLogicalNetworkProfile(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfilesAPI.CreateLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfilesAPI.CreateLogicalNetworkProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The Logical Network Profile create object | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkProfile

> DeleteLogicalNetworkProfile(ctx, logicalNetworkProfileId).Execute()

Delete a logical network profile



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
	logicalNetworkProfileId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkProfilesAPI.DeleteLogicalNetworkProfile(context.Background(), logicalNetworkProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfilesAPI.DeleteLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalNetworkProfileId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkProfileRequest struct via the builder pattern


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


## GetAllLogicalNetworkProfiles

> map[string]interface{} GetAllLogicalNetworkProfiles(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDescription(filterDescription).FilterLogicalNetworkType(filterLogicalNetworkType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all logical network profiles



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
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.           <p>              <b>Format: </b> filter.description={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.description=$not:$like:John Doe&filter.description=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterLogicalNetworkType := []string{"Inner_example"} // []string | Filter by logicalNetworkType query param.           <p>              <b>Format: </b> filter.logicalNetworkType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.logicalNetworkType=$not:$like:John Doe&filter.logicalNetworkType=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>name</li> <li>logicalNetworkType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label,name,description           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li> <li>name</li> <li>description</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfilesAPI.GetAllLogicalNetworkProfiles(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterDescription(filterDescription).FilterLogicalNetworkType(filterLogicalNetworkType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfilesAPI.GetAllLogicalNetworkProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLogicalNetworkProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfilesAPI.GetAllLogicalNetworkProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLogicalNetworkProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterDescription** | **[]string** | Filter by description query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.description&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.description&#x3D;$not:$like:John Doe&amp;filter.description&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterLogicalNetworkType** | **[]string** | Filter by logicalNetworkType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.logicalNetworkType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.logicalNetworkType&#x3D;$not:$like:John Doe&amp;filter.logicalNetworkType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;logicalNetworkType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label,name,description           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;description&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileById

> map[string]interface{} GetLogicalNetworkProfileById(ctx, logicalNetworkProfileId).Execute()

Get a logical network profile by ID



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
	logicalNetworkProfileId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfilesAPI.GetLogicalNetworkProfileById(context.Background(), logicalNetworkProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfilesAPI.GetLogicalNetworkProfileById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileById`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfilesAPI.GetLogicalNetworkProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalNetworkProfileId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalNetworkProfile

> map[string]interface{} UpdateLogicalNetworkProfile(ctx, logicalNetworkProfileId).Body(body).Execute()

Update a logical network profile



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
	logicalNetworkProfileId := float32(8.14) // float32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | The Logical Network Profile update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfilesAPI.UpdateLogicalNetworkProfile(context.Background(), logicalNetworkProfileId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfilesAPI.UpdateLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalNetworkProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfilesAPI.UpdateLogicalNetworkProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalNetworkProfileId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogicalNetworkProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | The Logical Network Profile update object | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

