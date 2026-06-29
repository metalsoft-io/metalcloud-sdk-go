# \SiteAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceAuthProvider**](SiteAPI.md#CreateDeviceAuthProvider) | **Post** /api/v2/device-auth-providers | Create a new device auth provider
[**CreateSite**](SiteAPI.md#CreateSite) | **Post** /api/v2/sites | Creates a Site
[**DecommissionSite**](SiteAPI.md#DecommissionSite) | **Post** /api/v2/sites/{siteId}/actions/decommission | Decommissions a Site
[**DeleteDeviceAuthProvider**](SiteAPI.md#DeleteDeviceAuthProvider) | **Delete** /api/v2/device-auth-providers/{id} | Delete a device auth provider by ID
[**GetAgents**](SiteAPI.md#GetAgents) | **Get** /api/v2/sites/{siteId}/controllers | Get a list of agents for a site
[**GetDeviceAuthProviderById**](SiteAPI.md#GetDeviceAuthProviderById) | **Get** /api/v2/device-auth-providers/{id} | Get a device auth provider by ID
[**GetDeviceAuthProviderCredentials**](SiteAPI.md#GetDeviceAuthProviderCredentials) | **Get** /api/v2/device-auth-providers/{id}/credentials | Get credentials for a device auth provider
[**GetRegistryUrls**](SiteAPI.md#GetRegistryUrls) | **Get** /api/v2/sites/controllers/actions/get/registry-urls | Get registry URLs for site controllers
[**GetSite**](SiteAPI.md#GetSite) | **Get** /api/v2/sites/{siteId} | Get Site information
[**GetSiteConfig**](SiteAPI.md#GetSiteConfig) | **Get** /api/v2/sites/{siteId}/config | Get Site Config information
[**GetSiteControllerOneLiner**](SiteAPI.md#GetSiteControllerOneLiner) | **Post** /api/v2/sites/{siteId}/controllers/actions/get/one-liner | Get a one liner to configure a site controller
[**GetSites**](SiteAPI.md#GetSites) | **Get** /api/v2/sites | Get a list of Sites
[**GetSitesStatistics**](SiteAPI.md#GetSitesStatistics) | **Get** /api/v2/sites/statistics | Get Sites statistics
[**ListDeviceAuthProviders**](SiteAPI.md#ListDeviceAuthProviders) | **Get** /api/v2/device-auth-providers | List device auth providers
[**UpdateDeviceAuthProvider**](SiteAPI.md#UpdateDeviceAuthProvider) | **Patch** /api/v2/device-auth-providers/{id} | Update a device auth provider by ID
[**UpdateDeviceAuthProviderSharedSecret**](SiteAPI.md#UpdateDeviceAuthProviderSharedSecret) | **Patch** /api/v2/device-auth-providers/{id}/shared-secret | Update the shared secret of a device auth provider
[**UpdateSite**](SiteAPI.md#UpdateSite) | **Patch** /api/v2/sites/{siteId} | Updates a Site
[**UpdateSiteConfig**](SiteAPI.md#UpdateSiteConfig) | **Patch** /api/v2/sites/{siteId}/config | Updates a Site Config



## CreateDeviceAuthProvider

> DeviceAuthProvider CreateDeviceAuthProvider(ctx).CreateDeviceAuthProvider(createDeviceAuthProvider).Execute()

Create a new device auth provider

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
	createDeviceAuthProvider := *openapiclient.NewCreateDeviceAuthProvider("tacacs-primary", "Primary TACACS+ Server", int64(1), "tacacs", "192.168.1.10", int32(49), "my-shared-secret", "admin") // CreateDeviceAuthProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.CreateDeviceAuthProvider(context.Background()).CreateDeviceAuthProvider(createDeviceAuthProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.CreateDeviceAuthProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceAuthProvider`: DeviceAuthProvider
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.CreateDeviceAuthProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceAuthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeviceAuthProvider** | [**CreateDeviceAuthProvider**](CreateDeviceAuthProvider.md) |  | 

### Return type

[**DeviceAuthProvider**](DeviceAuthProvider.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSite

> Site CreateSite(ctx).SiteCreate(siteCreate).Execute()

Creates a Site



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
	siteCreate := *openapiclient.NewSiteCreate("reading-uk", "Reading, UK") // SiteCreate | The Site create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.CreateSite(context.Background()).SiteCreate(siteCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.CreateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.CreateSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteCreate** | [**SiteCreate**](SiteCreate.md) | The Site create object | 

### Return type

[**Site**](Site.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecommissionSite

> DecommissionSite(ctx, siteId).IfMatch(ifMatch).Execute()

Decommissions a Site



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
	siteId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteAPI.DecommissionSite(context.Background(), siteId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.DecommissionSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecommissionSiteRequest struct via the builder pattern


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


## DeleteDeviceAuthProvider

> DeleteDeviceAuthProvider(ctx, id).IfMatch(ifMatch).Execute()

Delete a device auth provider by ID

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
	id := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteAPI.DeleteDeviceAuthProvider(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.DeleteDeviceAuthProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceAuthProviderRequest struct via the builder pattern


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


## GetAgents

> []AgentInfo GetAgents(ctx, siteId).Execute()

Get a list of agents for a site



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
	siteId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetAgents(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgents`: []AgentInfo
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AgentInfo**](AgentInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAuthProviderById

> DeviceAuthProvider GetDeviceAuthProviderById(ctx, id).Execute()

Get a device auth provider by ID

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
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetDeviceAuthProviderById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetDeviceAuthProviderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAuthProviderById`: DeviceAuthProvider
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetDeviceAuthProviderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAuthProviderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceAuthProvider**](DeviceAuthProvider.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAuthProviderCredentials

> DeviceAuthProviderCredentials GetDeviceAuthProviderCredentials(ctx, id).Execute()

Get credentials for a device auth provider



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
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetDeviceAuthProviderCredentials(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetDeviceAuthProviderCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAuthProviderCredentials`: DeviceAuthProviderCredentials
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetDeviceAuthProviderCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAuthProviderCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceAuthProviderCredentials**](DeviceAuthProviderCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryUrls

> []string GetRegistryUrls(ctx).Execute()

Get registry URLs for site controllers



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
	resp, r, err := apiClient.SiteAPI.GetRegistryUrls(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetRegistryUrls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryUrls`: []string
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetRegistryUrls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryUrlsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSite

> Site GetSite(ctx, siteId).Execute()

Get Site information



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
	siteId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetSite(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Site**](Site.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteConfig

> SiteConfig GetSiteConfig(ctx, siteId).Execute()

Get Site Config information



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
	siteId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetSiteConfig(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteConfig`: SiteConfig
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteConfig**](SiteConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteControllerOneLiner

> string GetSiteControllerOneLiner(ctx, siteId).GenerateSiteControllerOneliner(generateSiteControllerOneliner).Execute()

Get a one liner to configure a site controller



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
	siteId := int64(789) // int64 | 
	generateSiteControllerOneliner := *openapiclient.NewGenerateSiteControllerOneliner(false, false, false, "Registry_example", "GitHubTag_example", false, "api.example.com", "v6.4.0", "your-secure-tunnel-token", false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false) // GenerateSiteControllerOneliner | Data needed for the controller one-liner generation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetSiteControllerOneLiner(context.Background(), siteId).GenerateSiteControllerOneliner(generateSiteControllerOneliner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteControllerOneLiner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteControllerOneLiner`: string
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteControllerOneLiner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteControllerOneLinerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateSiteControllerOneliner** | [**GenerateSiteControllerOneliner**](GenerateSiteControllerOneliner.md) | Data needed for the controller one-liner generation | 

### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSites

> SitePaginatedList GetSites(ctx).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterSlug(filterSlug).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Sites



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSlug := []string{"Inner_example"} // []string | Filter by slug query param.  **Format:** filter.slug={$not}:OPERATION:VALUE    **Example:** filter.slug=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:DESC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,slug   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - slug  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetSites(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterSlug(filterSlug).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSites`: SitePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSlug** | **[]string** | Filter by slug query param.  **Format:** filter.slug&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.slug&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:DESC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,slug   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - slug  | 

### Return type

[**SitePaginatedList**](SitePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSitesStatistics

> SiteStatistics GetSitesStatistics(ctx).Execute()

Get Sites statistics



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
	resp, r, err := apiClient.SiteAPI.GetSitesStatistics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSitesStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSitesStatistics`: SiteStatistics
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSitesStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesStatisticsRequest struct via the builder pattern


### Return type

[**SiteStatistics**](SiteStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceAuthProviders

> DeviceAuthProviderPaginatedList ListDeviceAuthProviders(ctx).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterLabel(filterLabel).FilterName(filterName).FilterKind(filterKind).FilterStatus(filterStatus).FilterIpAddress(filterIpAddress).FilterPort(filterPort).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

List device auth providers



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
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe&filter.label=$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe&filter.name=$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$eq:John Doe&filter.status=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterIpAddress := []string{"Inner_example"} // []string | Filter by ipAddress query param.  **Format:** filter.ipAddress={$not}:OPERATION:VALUE    **Example:** filter.ipAddress=$eq:John Doe&filter.ipAddress=$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or (optional)
	filterPort := []string{"Inner_example"} // []string | Filter by port query param.  **Format:** filter.port={$not}:OPERATION:VALUE    **Example:** filter.port=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  - label  - name  - kind  - status  - ipAddress  - port  - createdAt  - updatedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,name,kind,status   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  - kind  - status  - ipAddress  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,siteId,label,name,annotations   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.ListDeviceAuthProviders(context.Background()).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterLabel(filterLabel).FilterName(filterName).FilterKind(filterKind).FilterStatus(filterStatus).FilterIpAddress(filterIpAddress).FilterPort(filterPort).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.ListDeviceAuthProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeviceAuthProviders`: DeviceAuthProviderPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.ListDeviceAuthProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceAuthProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe&amp;filter.label&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe&amp;filter.name&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$eq:John Doe&amp;filter.status&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterIpAddress** | **[]string** | Filter by ipAddress query param.  **Format:** filter.ipAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.ipAddress&#x3D;$eq:John Doe&amp;filter.ipAddress&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or | 
 **filterPort** | **[]string** | Filter by port query param.  **Format:** filter.port&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.port&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  - label  - name  - kind  - status  - ipAddress  - port  - createdAt  - updatedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,name,kind,status   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  - kind  - status  - ipAddress  | 
 **select_** | **string** | List of fields to select.  **Example:** id,siteId,label,name,annotations   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**DeviceAuthProviderPaginatedList**](DeviceAuthProviderPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceAuthProvider

> DeviceAuthProvider UpdateDeviceAuthProvider(ctx, id).UpdateDeviceAuthProvider(updateDeviceAuthProvider).IfMatch(ifMatch).Execute()

Update a device auth provider by ID



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
	id := int64(789) // int64 | 
	updateDeviceAuthProvider := *openapiclient.NewUpdateDeviceAuthProvider() // UpdateDeviceAuthProvider | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.UpdateDeviceAuthProvider(context.Background(), id).UpdateDeviceAuthProvider(updateDeviceAuthProvider).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.UpdateDeviceAuthProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceAuthProvider`: DeviceAuthProvider
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.UpdateDeviceAuthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceAuthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceAuthProvider** | [**UpdateDeviceAuthProvider**](UpdateDeviceAuthProvider.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**DeviceAuthProvider**](DeviceAuthProvider.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceAuthProviderSharedSecret

> UpdateDeviceAuthProviderSharedSecret(ctx, id).UpdateDeviceAuthProviderSharedSecret(updateDeviceAuthProviderSharedSecret).IfMatch(ifMatch).Execute()

Update the shared secret of a device auth provider



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
	id := int64(789) // int64 | 
	updateDeviceAuthProviderSharedSecret := *openapiclient.NewUpdateDeviceAuthProviderSharedSecret("new-shared-secret") // UpdateDeviceAuthProviderSharedSecret | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteAPI.UpdateDeviceAuthProviderSharedSecret(context.Background(), id).UpdateDeviceAuthProviderSharedSecret(updateDeviceAuthProviderSharedSecret).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.UpdateDeviceAuthProviderSharedSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceAuthProviderSharedSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceAuthProviderSharedSecret** | [**UpdateDeviceAuthProviderSharedSecret**](UpdateDeviceAuthProviderSharedSecret.md) |  | 
 **ifMatch** | **string** | Entity tag | 

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


## UpdateSite

> Site UpdateSite(ctx, siteId).SiteUpdate(siteUpdate).IfMatch(ifMatch).Execute()

Updates a Site



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
	siteId := int64(789) // int64 | 
	siteUpdate := *openapiclient.NewSiteUpdate() // SiteUpdate | The Site update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.UpdateSite(context.Background(), siteId).SiteUpdate(siteUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.UpdateSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.UpdateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteUpdate** | [**SiteUpdate**](SiteUpdate.md) | The Site update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Site**](Site.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteConfig

> SiteConfig UpdateSiteConfig(ctx, siteId).SiteConfigUpdate(siteConfigUpdate).IfMatch(ifMatch).Execute()

Updates a Site Config



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
	siteId := int64(789) // int64 | 
	siteConfigUpdate := *openapiclient.NewSiteConfigUpdate() // SiteConfigUpdate | The Site Config update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.UpdateSiteConfig(context.Background(), siteId).SiteConfigUpdate(siteConfigUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.UpdateSiteConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteConfig`: SiteConfig
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.UpdateSiteConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteConfigUpdate** | [**SiteConfigUpdate**](SiteConfigUpdate.md) | The Site Config update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**SiteConfig**](SiteConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

