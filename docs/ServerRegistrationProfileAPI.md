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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$btw:John Doe&filter.name=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  (optional)

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
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$btw:John Doe&amp;filter.name&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  | 

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

