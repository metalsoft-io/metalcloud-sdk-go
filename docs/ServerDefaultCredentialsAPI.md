# \ServerDefaultCredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerDefaultCredentials**](ServerDefaultCredentialsAPI.md#CreateServerDefaultCredentials) | **Post** /api/v2/servers/default-credentials | Creates a Server Default Credentials
[**DeleteServerDefaultCredentials**](ServerDefaultCredentialsAPI.md#DeleteServerDefaultCredentials) | **Delete** /api/v2/servers/default-credentials/{serverDefaultCredentialsId} | Deletes a Server Default Credentials
[**GetServerDefaultCredentialsCredentials**](ServerDefaultCredentialsAPI.md#GetServerDefaultCredentialsCredentials) | **Get** /api/v2/servers/default-credentials/{serverDefaultCredentialsId}/credentials | Get Server Default Credentials unencrypted
[**GetServerDefaultCredentialsInfo**](ServerDefaultCredentialsAPI.md#GetServerDefaultCredentialsInfo) | **Get** /api/v2/servers/default-credentials/{serverDefaultCredentialsId} | Get Server Default Credentials information
[**GetServersDefaultCredentials**](ServerDefaultCredentialsAPI.md#GetServersDefaultCredentials) | **Get** /api/v2/servers/default-credentials | Get a list of Server Default Credentials
[**UpdateServerDefaultCredentials**](ServerDefaultCredentialsAPI.md#UpdateServerDefaultCredentials) | **Patch** /api/v2/servers/default-credentials/{serverDefaultCredentialsId} | Updates a Server Default Credentials



## CreateServerDefaultCredentials

> ServerDefaultCredentials CreateServerDefaultCredentials(ctx).CreateServerDefaultCredentials(createServerDefaultCredentials).Execute()

Creates a Server Default Credentials



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
	createServerDefaultCredentials := *openapiclient.NewCreateServerDefaultCredentials(float32(123), "ServerSerialNumber_example", "ServerMacAddress_example", "DefaultUsername_example", "DefaultPassword_example") // CreateServerDefaultCredentials | The Server Default Credentials create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.CreateServerDefaultCredentials(context.Background()).CreateServerDefaultCredentials(createServerDefaultCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.CreateServerDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerDefaultCredentials`: ServerDefaultCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.CreateServerDefaultCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerDefaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerDefaultCredentials** | [**CreateServerDefaultCredentials**](CreateServerDefaultCredentials.md) | The Server Default Credentials create object | 

### Return type

[**ServerDefaultCredentials**](ServerDefaultCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerDefaultCredentials

> DeleteServerDefaultCredentials(ctx, serverDefaultCredentialsId).Execute()

Deletes a Server Default Credentials



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerDefaultCredentialsAPI.DeleteServerDefaultCredentials(context.Background(), serverDefaultCredentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.DeleteServerDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerDefaultCredentialsRequest struct via the builder pattern


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


## GetServerDefaultCredentialsCredentials

> ServerDefaultCredentialsCredentials GetServerDefaultCredentialsCredentials(ctx, serverDefaultCredentialsId).Execute()

Get Server Default Credentials unencrypted



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.GetServerDefaultCredentialsCredentials(context.Background(), serverDefaultCredentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDefaultCredentialsCredentials`: ServerDefaultCredentialsCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDefaultCredentialsCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerDefaultCredentialsCredentials**](ServerDefaultCredentialsCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDefaultCredentialsInfo

> ServerDefaultCredentials GetServerDefaultCredentialsInfo(ctx, serverDefaultCredentialsId).Execute()

Get Server Default Credentials information



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.GetServerDefaultCredentialsInfo(context.Background(), serverDefaultCredentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDefaultCredentialsInfo`: ServerDefaultCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDefaultCredentialsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerDefaultCredentials**](ServerDefaultCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServersDefaultCredentials

> ServerDefaultCredentialsPaginatedList GetServersDefaultCredentials(ctx).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterServerSerialNumber(filterServerSerialNumber).FilterServerMacAddress(filterServerMacAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Server Default Credentials



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
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerSerialNumber := []string{"Inner_example"} // []string | Filter by serverSerialNumber query param.           <p>              <b>Format: </b> filter.serverSerialNumber={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverSerialNumber=$not:$like:John Doe&filter.serverSerialNumber=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterServerMacAddress := []string{"Inner_example"} // []string | Filter by serverMacAddress query param.           <p>              <b>Format: </b> filter.serverMacAddress={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverMacAddress=$not:$like:John Doe&filter.serverMacAddress=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>siteId</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,siteId,serverSerialNumber,serverMacAddress           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>siteId</li> <li>serverSerialNumber</li> <li>serverMacAddress</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.GetServersDefaultCredentials(context.Background()).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterServerSerialNumber(filterServerSerialNumber).FilterServerMacAddress(filterServerMacAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.GetServersDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServersDefaultCredentials`: ServerDefaultCredentialsPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.GetServersDefaultCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServersDefaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerSerialNumber** | **[]string** | Filter by serverSerialNumber query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverSerialNumber&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverSerialNumber&#x3D;$not:$like:John Doe&amp;filter.serverSerialNumber&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerMacAddress** | **[]string** | Filter by serverMacAddress query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverMacAddress&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverMacAddress&#x3D;$not:$like:John Doe&amp;filter.serverMacAddress&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,siteId,serverSerialNumber,serverMacAddress           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;serverSerialNumber&lt;/li&gt; &lt;li&gt;serverMacAddress&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**ServerDefaultCredentialsPaginatedList**](ServerDefaultCredentialsPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDefaultCredentials

> ServerDefaultCredentials UpdateServerDefaultCredentials(ctx, serverDefaultCredentialsId).UpdateServerDefaultCredentials(updateServerDefaultCredentials).Execute()

Updates a Server Default Credentials



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 
	updateServerDefaultCredentials := *openapiclient.NewUpdateServerDefaultCredentials() // UpdateServerDefaultCredentials | The Server Default Credentials update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.UpdateServerDefaultCredentials(context.Background(), serverDefaultCredentialsId).UpdateServerDefaultCredentials(updateServerDefaultCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.UpdateServerDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerDefaultCredentials`: ServerDefaultCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.UpdateServerDefaultCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDefaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerDefaultCredentials** | [**UpdateServerDefaultCredentials**](UpdateServerDefaultCredentials.md) | The Server Default Credentials update object | 

### Return type

[**ServerDefaultCredentials**](ServerDefaultCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

