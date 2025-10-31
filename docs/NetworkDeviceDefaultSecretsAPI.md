# \NetworkDeviceDefaultSecretsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDeviceDefaultSecrets**](NetworkDeviceDefaultSecretsAPI.md#CreateNetworkDeviceDefaultSecrets) | **Post** /api/v2/network-devices/default-secrets | Creates a NetworkDevice Default Secrets
[**DeleteNetworkDeviceDefaultSecrets**](NetworkDeviceDefaultSecretsAPI.md#DeleteNetworkDeviceDefaultSecrets) | **Delete** /api/v2/network-devices/default-secrets/{networkDeviceDefaultSecretsId} | Deletes a NetworkDevice Default Secrets
[**GetNetworkDeviceDefaultSecretsCredentials**](NetworkDeviceDefaultSecretsAPI.md#GetNetworkDeviceDefaultSecretsCredentials) | **Get** /api/v2/network-devices/default-secrets/{networkDeviceDefaultSecretsId}/credentials | Get NetworkDevice Default Secrets unencrypted
[**GetNetworkDeviceDefaultSecretsInfo**](NetworkDeviceDefaultSecretsAPI.md#GetNetworkDeviceDefaultSecretsInfo) | **Get** /api/v2/network-devices/default-secrets/{networkDeviceDefaultSecretsId} | Get NetworkDevice Default Secrets information
[**GetNetworkDevicesDefaultSecrets**](NetworkDeviceDefaultSecretsAPI.md#GetNetworkDevicesDefaultSecrets) | **Get** /api/v2/network-devices/default-secrets | Get a list of NetworkDevice Default Secrets
[**UpdateNetworkDeviceDefaultSecrets**](NetworkDeviceDefaultSecretsAPI.md#UpdateNetworkDeviceDefaultSecrets) | **Patch** /api/v2/network-devices/default-secrets/{networkDeviceDefaultSecretsId} | Updates a NetworkDevice Default Secrets



## CreateNetworkDeviceDefaultSecrets

> NetworkDeviceDefaultSecrets CreateNetworkDeviceDefaultSecrets(ctx).CreateNetworkDeviceDefaultSecrets(createNetworkDeviceDefaultSecrets).Execute()

Creates a NetworkDevice Default Secrets



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
	createNetworkDeviceDefaultSecrets := *openapiclient.NewCreateNetworkDeviceDefaultSecrets(float32(123), "MacAddressOrSerialNumber_example", "SecretName_example", "SecretValue_example") // CreateNetworkDeviceDefaultSecrets | The NetworkDevice Default Secrets create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceDefaultSecretsAPI.CreateNetworkDeviceDefaultSecrets(context.Background()).CreateNetworkDeviceDefaultSecrets(createNetworkDeviceDefaultSecrets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceDefaultSecretsAPI.CreateNetworkDeviceDefaultSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDeviceDefaultSecrets`: NetworkDeviceDefaultSecrets
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceDefaultSecretsAPI.CreateNetworkDeviceDefaultSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDeviceDefaultSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDeviceDefaultSecrets** | [**CreateNetworkDeviceDefaultSecrets**](CreateNetworkDeviceDefaultSecrets.md) | The NetworkDevice Default Secrets create object | 

### Return type

[**NetworkDeviceDefaultSecrets**](NetworkDeviceDefaultSecrets.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDeviceDefaultSecrets

> DeleteNetworkDeviceDefaultSecrets(ctx, networkDeviceDefaultSecretsId).Execute()

Deletes a NetworkDevice Default Secrets



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
	networkDeviceDefaultSecretsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkDeviceDefaultSecretsAPI.DeleteNetworkDeviceDefaultSecrets(context.Background(), networkDeviceDefaultSecretsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceDefaultSecretsAPI.DeleteNetworkDeviceDefaultSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceDefaultSecretsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDeviceDefaultSecretsRequest struct via the builder pattern


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


## GetNetworkDeviceDefaultSecretsCredentials

> NetworkDeviceDefaultSecretsCredentials GetNetworkDeviceDefaultSecretsCredentials(ctx, networkDeviceDefaultSecretsId).Execute()

Get NetworkDevice Default Secrets unencrypted



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
	networkDeviceDefaultSecretsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceDefaultSecretsAPI.GetNetworkDeviceDefaultSecretsCredentials(context.Background(), networkDeviceDefaultSecretsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceDefaultSecretsAPI.GetNetworkDeviceDefaultSecretsCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceDefaultSecretsCredentials`: NetworkDeviceDefaultSecretsCredentials
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceDefaultSecretsAPI.GetNetworkDeviceDefaultSecretsCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceDefaultSecretsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceDefaultSecretsCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceDefaultSecretsCredentials**](NetworkDeviceDefaultSecretsCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDeviceDefaultSecretsInfo

> NetworkDeviceDefaultSecrets GetNetworkDeviceDefaultSecretsInfo(ctx, networkDeviceDefaultSecretsId).Execute()

Get NetworkDevice Default Secrets information



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
	networkDeviceDefaultSecretsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceDefaultSecretsAPI.GetNetworkDeviceDefaultSecretsInfo(context.Background(), networkDeviceDefaultSecretsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceDefaultSecretsAPI.GetNetworkDeviceDefaultSecretsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceDefaultSecretsInfo`: NetworkDeviceDefaultSecrets
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceDefaultSecretsAPI.GetNetworkDeviceDefaultSecretsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceDefaultSecretsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceDefaultSecretsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceDefaultSecrets**](NetworkDeviceDefaultSecrets.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevicesDefaultSecrets

> NetworkDeviceDefaultSecretsPaginatedList GetNetworkDevicesDefaultSecrets(ctx).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterMacAddressOrSerialNumber(filterMacAddressOrSerialNumber).FilterSecretName(filterSecretName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of NetworkDevice Default Secrets



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
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterMacAddressOrSerialNumber := []string{"Inner_example"} // []string | Filter by macAddressOrSerialNumber query param.           <p>              <b>Format: </b> filter.macAddressOrSerialNumber={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.macAddressOrSerialNumber=$not:$like:John Doe&filter.macAddressOrSerialNumber=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$ilike</li></ul> (optional)
	filterSecretName := []string{"Inner_example"} // []string | Filter by secretName query param.           <p>              <b>Format: </b> filter.secretName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.secretName=$not:$like:John Doe&filter.secretName=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$ilike</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>siteId</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,siteId,macAddressOrSerialNumber,secretName           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>siteId</li> <li>macAddressOrSerialNumber</li> <li>secretName</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceDefaultSecretsAPI.GetNetworkDevicesDefaultSecrets(context.Background()).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterMacAddressOrSerialNumber(filterMacAddressOrSerialNumber).FilterSecretName(filterSecretName).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceDefaultSecretsAPI.GetNetworkDevicesDefaultSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevicesDefaultSecrets`: NetworkDeviceDefaultSecretsPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceDefaultSecretsAPI.GetNetworkDevicesDefaultSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicesDefaultSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterMacAddressOrSerialNumber** | **[]string** | Filter by macAddressOrSerialNumber query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.macAddressOrSerialNumber&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.macAddressOrSerialNumber&#x3D;$not:$like:John Doe&amp;filter.macAddressOrSerialNumber&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt; | 
 **filterSecretName** | **[]string** | Filter by secretName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.secretName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.secretName&#x3D;$not:$like:John Doe&amp;filter.secretName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,siteId,macAddressOrSerialNumber,secretName           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;macAddressOrSerialNumber&lt;/li&gt; &lt;li&gt;secretName&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**NetworkDeviceDefaultSecretsPaginatedList**](NetworkDeviceDefaultSecretsPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDeviceDefaultSecrets

> NetworkDeviceDefaultSecrets UpdateNetworkDeviceDefaultSecrets(ctx, networkDeviceDefaultSecretsId).UpdateNetworkDeviceDefaultSecrets(updateNetworkDeviceDefaultSecrets).Execute()

Updates a NetworkDevice Default Secrets



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
	networkDeviceDefaultSecretsId := float32(8.14) // float32 | 
	updateNetworkDeviceDefaultSecrets := *openapiclient.NewUpdateNetworkDeviceDefaultSecrets() // UpdateNetworkDeviceDefaultSecrets | The NetworkDevice Default Secrets update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkDeviceDefaultSecretsAPI.UpdateNetworkDeviceDefaultSecrets(context.Background(), networkDeviceDefaultSecretsId).UpdateNetworkDeviceDefaultSecrets(updateNetworkDeviceDefaultSecrets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDeviceDefaultSecretsAPI.UpdateNetworkDeviceDefaultSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDeviceDefaultSecrets`: NetworkDeviceDefaultSecrets
	fmt.Fprintf(os.Stdout, "Response from `NetworkDeviceDefaultSecretsAPI.UpdateNetworkDeviceDefaultSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceDefaultSecretsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDeviceDefaultSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDeviceDefaultSecrets** | [**UpdateNetworkDeviceDefaultSecrets**](UpdateNetworkDeviceDefaultSecrets.md) | The NetworkDevice Default Secrets update object | 

### Return type

[**NetworkDeviceDefaultSecrets**](NetworkDeviceDefaultSecrets.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

