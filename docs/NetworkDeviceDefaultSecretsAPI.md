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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterMacAddressOrSerialNumber := []string{"Inner_example"} // []string | Filter by macAddressOrSerialNumber query param.  **Format:** filter.macAddressOrSerialNumber={$not}:OPERATION:VALUE    **Example:** filter.macAddressOrSerialNumber=$eq:John Doe&filter.macAddressOrSerialNumber=$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or (optional)
	filterSecretName := []string{"Inner_example"} // []string | Filter by secretName query param.  **Format:** filter.secretName={$not}:OPERATION:VALUE    **Example:** filter.secretName=$eq:John Doe&filter.secretName=$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,siteId,macAddressOrSerialNumber,secretName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - siteId  - macAddressOrSerialNumber  - secretName  (optional)

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
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterMacAddressOrSerialNumber** | **[]string** | Filter by macAddressOrSerialNumber query param.  **Format:** filter.macAddressOrSerialNumber&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.macAddressOrSerialNumber&#x3D;$eq:John Doe&amp;filter.macAddressOrSerialNumber&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or | 
 **filterSecretName** | **[]string** | Filter by secretName query param.  **Format:** filter.secretName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.secretName&#x3D;$eq:John Doe&amp;filter.secretName&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,siteId,macAddressOrSerialNumber,secretName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - siteId  - macAddressOrSerialNumber  - secretName  | 

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

