# \ServerRegistrationProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerRegistrationProfile**](ServerRegistrationProfileAPI.md#CreateServerRegistrationProfile) | **Post** /api/v2/servers/registration-profiles | Creates a Server Registration Profile
[**DeleteServerRegistrationProfile**](ServerRegistrationProfileAPI.md#DeleteServerRegistrationProfile) | **Delete** /api/v2/servers/registration-profiles/{serverRegistrationProfileId} | Deletes a Server Registration Profile
[**GetServerRegistrationProfileInfo**](ServerRegistrationProfileAPI.md#GetServerRegistrationProfileInfo) | **Get** /api/v2/servers/registration-profiles/{serverRegistrationProfileId} | Get Server Registration Profile information
[**GetServerRegistrationProfileInfoForServer**](ServerRegistrationProfileAPI.md#GetServerRegistrationProfileInfoForServer) | **Get** /api/v2/servers/registration-profiles/search/for-server/{serverId} | Get Server Registration Profile information for specific server
[**GetServerRegistrationProfiles**](ServerRegistrationProfileAPI.md#GetServerRegistrationProfiles) | **Get** /api/v2/servers/registration-profiles | Get a list of Server Registration Profiles
[**SearchServerRegistrationProfileInfo**](ServerRegistrationProfileAPI.md#SearchServerRegistrationProfileInfo) | **Get** /api/v2/servers/registration-profiles/search | Search for a Server Registration Profile by data to match
[**UpdateServerRegistrationProfile**](ServerRegistrationProfileAPI.md#UpdateServerRegistrationProfile) | **Patch** /api/v2/servers/registration-profiles/{serverRegistrationProfileId} | Updates a Server Registration Profile



## CreateServerRegistrationProfile

> ServerRegistrationProfile CreateServerRegistrationProfile(ctx).ServerRegistrationProfileCreate(serverRegistrationProfileCreate).Execute()

Creates a Server Registration Profile



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
	serverRegistrationProfileCreate := *openapiclient.NewServerRegistrationProfileCreate("Name_example", *openapiclient.NewServerRegistrationProfileSettings()) // ServerRegistrationProfileCreate | The Server Registration Profile create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerRegistrationProfileAPI.CreateServerRegistrationProfile(context.Background()).ServerRegistrationProfileCreate(serverRegistrationProfileCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerRegistrationProfileAPI.CreateServerRegistrationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerRegistrationProfile`: ServerRegistrationProfile
	fmt.Fprintf(os.Stdout, "Response from `ServerRegistrationProfileAPI.CreateServerRegistrationProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerRegistrationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverRegistrationProfileCreate** | [**ServerRegistrationProfileCreate**](ServerRegistrationProfileCreate.md) | The Server Registration Profile create object | 

### Return type

[**ServerRegistrationProfile**](ServerRegistrationProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerRegistrationProfile

> DeleteServerRegistrationProfile(ctx, serverRegistrationProfileId).IfMatch(ifMatch).Execute()

Deletes a Server Registration Profile



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
	serverRegistrationProfileId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerRegistrationProfileAPI.DeleteServerRegistrationProfile(context.Background(), serverRegistrationProfileId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerRegistrationProfileAPI.DeleteServerRegistrationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverRegistrationProfileId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerRegistrationProfileRequest struct via the builder pattern


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


## GetServerRegistrationProfileInfo

> ServerRegistrationProfile GetServerRegistrationProfileInfo(ctx, serverRegistrationProfileId).Execute()

Get Server Registration Profile information



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
	serverRegistrationProfileId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerRegistrationProfileAPI.GetServerRegistrationProfileInfo(context.Background(), serverRegistrationProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerRegistrationProfileAPI.GetServerRegistrationProfileInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerRegistrationProfileInfo`: ServerRegistrationProfile
	fmt.Fprintf(os.Stdout, "Response from `ServerRegistrationProfileAPI.GetServerRegistrationProfileInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverRegistrationProfileId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRegistrationProfileInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerRegistrationProfile**](ServerRegistrationProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerRegistrationProfileInfoForServer

> ServerRegistrationProfile GetServerRegistrationProfileInfoForServer(ctx, serverId).Execute()

Get Server Registration Profile information for specific server



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
	serverId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerRegistrationProfileAPI.GetServerRegistrationProfileInfoForServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerRegistrationProfileAPI.GetServerRegistrationProfileInfoForServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerRegistrationProfileInfoForServer`: ServerRegistrationProfile
	fmt.Fprintf(os.Stdout, "Response from `ServerRegistrationProfileAPI.GetServerRegistrationProfileInfoForServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRegistrationProfileInfoForServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerRegistrationProfile**](ServerRegistrationProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerRegistrationProfiles

> ServerRegistrationProfilePaginatedList GetServerRegistrationProfiles(ctx).Page(page).Limit(limit).FilterName(filterName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Server Registration Profiles



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
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>name</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,name           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>name</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerRegistrationProfileAPI.GetServerRegistrationProfiles(context.Background()).Page(page).Limit(limit).FilterName(filterName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerRegistrationProfileAPI.GetServerRegistrationProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerRegistrationProfiles`: ServerRegistrationProfilePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerRegistrationProfileAPI.GetServerRegistrationProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRegistrationProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**ServerRegistrationProfilePaginatedList**](ServerRegistrationProfilePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchServerRegistrationProfileInfo

> ServerRegistrationProfile SearchServerRegistrationProfileInfo(ctx).SiteId(siteId).Execute()

Search for a Server Registration Profile by data to match



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
	siteId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerRegistrationProfileAPI.SearchServerRegistrationProfileInfo(context.Background()).SiteId(siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerRegistrationProfileAPI.SearchServerRegistrationProfileInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchServerRegistrationProfileInfo`: ServerRegistrationProfile
	fmt.Fprintf(os.Stdout, "Response from `ServerRegistrationProfileAPI.SearchServerRegistrationProfileInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchServerRegistrationProfileInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteId** | **float32** |  | 

### Return type

[**ServerRegistrationProfile**](ServerRegistrationProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerRegistrationProfile

> ServerRegistrationProfile UpdateServerRegistrationProfile(ctx, serverRegistrationProfileId).ServerRegistrationProfileUpdate(serverRegistrationProfileUpdate).IfMatch(ifMatch).Execute()

Updates a Server Registration Profile



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
	serverRegistrationProfileId := float32(8.14) // float32 | 
	serverRegistrationProfileUpdate := *openapiclient.NewServerRegistrationProfileUpdate() // ServerRegistrationProfileUpdate | The Server Registration Profile update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerRegistrationProfileAPI.UpdateServerRegistrationProfile(context.Background(), serverRegistrationProfileId).ServerRegistrationProfileUpdate(serverRegistrationProfileUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerRegistrationProfileAPI.UpdateServerRegistrationProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerRegistrationProfile`: ServerRegistrationProfile
	fmt.Fprintf(os.Stdout, "Response from `ServerRegistrationProfileAPI.UpdateServerRegistrationProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverRegistrationProfileId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerRegistrationProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverRegistrationProfileUpdate** | [**ServerRegistrationProfileUpdate**](ServerRegistrationProfileUpdate.md) | The Server Registration Profile update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ServerRegistrationProfile**](ServerRegistrationProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

