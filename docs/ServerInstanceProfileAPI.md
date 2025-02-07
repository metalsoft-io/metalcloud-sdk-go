# \ServerInstanceProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerInstanceProfile**](ServerInstanceProfileAPI.md#CreateServerInstanceProfile) | **Post** /api/v2/server-instance-profiles | Create Server Instance Profile
[**DeleteServerInstanceProfile**](ServerInstanceProfileAPI.md#DeleteServerInstanceProfile) | **Delete** /api/v2/server-instance-profiles/{serverInstanceProfileId} | Delete Server Instance Profile
[**GetServerInstanceProfile**](ServerInstanceProfileAPI.md#GetServerInstanceProfile) | **Get** /api/v2/server-instance-profiles/{serverInstanceProfileId} | Get Server Instance Profile details
[**GetServerInstanceProfiles**](ServerInstanceProfileAPI.md#GetServerInstanceProfiles) | **Get** /api/v2/server-instance-profiles | List all Server Instance Profiles
[**UpdateServerInstanceProfile**](ServerInstanceProfileAPI.md#UpdateServerInstanceProfile) | **Patch** /api/v2/server-instance-profiles/{serverInstanceProfileId} | Update Server Instance Profile



## CreateServerInstanceProfile

> CreateServerInstanceProfile(ctx).ServerInstanceProfileCreate(serverInstanceProfileCreate).Execute()

Create Server Instance Profile



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
	serverInstanceProfileCreate := *openapiclient.NewServerInstanceProfileCreate() // ServerInstanceProfileCreate | The Server Instance Profile to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerInstanceProfileAPI.CreateServerInstanceProfile(context.Background()).ServerInstanceProfileCreate(serverInstanceProfileCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceProfileAPI.CreateServerInstanceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInstanceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverInstanceProfileCreate** | [**ServerInstanceProfileCreate**](ServerInstanceProfileCreate.md) | The Server Instance Profile to create | 

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


## DeleteServerInstanceProfile

> DeleteServerInstanceProfile(ctx, serverInstanceProfileId).IfMatch(ifMatch).Execute()

Delete Server Instance Profile



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
	serverInstanceProfileId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerInstanceProfileAPI.DeleteServerInstanceProfile(context.Background(), serverInstanceProfileId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceProfileAPI.DeleteServerInstanceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerInstanceProfileRequest struct via the builder pattern


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


## GetServerInstanceProfile

> ServerInstanceProfile GetServerInstanceProfile(ctx, serverInstanceProfileId).Execute()

Get Server Instance Profile details



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
	serverInstanceProfileId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceProfileAPI.GetServerInstanceProfile(context.Background(), serverInstanceProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceProfileAPI.GetServerInstanceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceProfile`: ServerInstanceProfile
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceProfileAPI.GetServerInstanceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerInstanceProfile**](ServerInstanceProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInstanceProfiles

> ServerInstanceProfilePaginatedList GetServerInstanceProfiles(ctx).Page(page).Limit(limit).FilterLabel(filterLabel).FilterOwnerId(filterOwnerId).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

List all Server Instance Profiles



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterOwnerId := []string{"Inner_example"} // []string | Filter by ownerId query param.           <p>              <b>Format: </b> filter.ownerId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.ownerId=$not:$like:John Doe&filter.ownerId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> label:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>ownerId</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> label,ownerId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>label</li> <li>ownerId</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> id,label,ownerId           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceProfileAPI.GetServerInstanceProfiles(context.Background()).Page(page).Limit(limit).FilterLabel(filterLabel).FilterOwnerId(filterOwnerId).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceProfileAPI.GetServerInstanceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInstanceProfiles`: ServerInstanceProfilePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceProfileAPI.GetServerInstanceProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterOwnerId** | **[]string** | Filter by ownerId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.ownerId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.ownerId&#x3D;$not:$like:John Doe&amp;filter.ownerId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; label:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;ownerId&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label,ownerId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt; &lt;li&gt;ownerId&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,ownerId           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

### Return type

[**ServerInstanceProfilePaginatedList**](ServerInstanceProfilePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerInstanceProfile

> ServerInstanceProfile UpdateServerInstanceProfile(ctx, serverInstanceProfileId).ServerInstanceProfileUpdate(serverInstanceProfileUpdate).IfMatch(ifMatch).Execute()

Update Server Instance Profile



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
	serverInstanceProfileId := int32(56) // int32 | 
	serverInstanceProfileUpdate := *openapiclient.NewServerInstanceProfileUpdate() // ServerInstanceProfileUpdate | The Server Instance Profile changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerInstanceProfileAPI.UpdateServerInstanceProfile(context.Background(), serverInstanceProfileId).ServerInstanceProfileUpdate(serverInstanceProfileUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceProfileAPI.UpdateServerInstanceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerInstanceProfile`: ServerInstanceProfile
	fmt.Fprintf(os.Stdout, "Response from `ServerInstanceProfileAPI.UpdateServerInstanceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverInstanceProfileUpdate** | [**ServerInstanceProfileUpdate**](ServerInstanceProfileUpdate.md) | The Server Instance Profile changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ServerInstanceProfile**](ServerInstanceProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

