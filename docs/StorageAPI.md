# \StorageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorage**](StorageAPI.md#CreateStorage) | **Post** /api/v2/storages | Creates a Storage
[**CreateStorageNetworkDeviceConfiguration**](StorageAPI.md#CreateStorageNetworkDeviceConfiguration) | **Post** /api/v2/storages/{storageId}/network-device-configurations | Creates a Storage Network Device Configuration
[**DeleteStorage**](StorageAPI.md#DeleteStorage) | **Delete** /api/v2/storages/{storageId} | Deletes a Storage
[**DeleteStorageNetworkDeviceConfiguration**](StorageAPI.md#DeleteStorageNetworkDeviceConfiguration) | **Delete** /api/v2/storages/{storageId}/network-device-configurations/{storageNetworkDeviceConfigurationId} | Deletes a Storage Network Device Configuration
[**GetStorage**](StorageAPI.md#GetStorage) | **Get** /api/v2/storages/{storageId} | Retrieves a Storage
[**GetStorageBuckets**](StorageAPI.md#GetStorageBuckets) | **Get** /api/v2/storages/{storageId}/buckets | Get all Buckets linked to the specified storage
[**GetStorageCredentials**](StorageAPI.md#GetStorageCredentials) | **Get** /api/v2/storages/{storageId}/credentials | Get Storage credentials
[**GetStorageDrives**](StorageAPI.md#GetStorageDrives) | **Get** /api/v2/storages/{storageId}/drives | Get all Drives linked to the specified storage
[**GetStorageFileShares**](StorageAPI.md#GetStorageFileShares) | **Get** /api/v2/storages/{storageId}/file-shares | Get all File Shares linked to the specified storage
[**GetStorageNetworkDeviceConfigurations**](StorageAPI.md#GetStorageNetworkDeviceConfigurations) | **Get** /api/v2/storages/{storageId}/network-device-configurations | Retrieves Storage Network Device Configurations
[**GetStorageStatistics**](StorageAPI.md#GetStorageStatistics) | **Get** /api/v2/storages/{storageId}/statistics | Get Storages statistics
[**GetStorages**](StorageAPI.md#GetStorages) | **Get** /api/v2/storages | Get a list of Storages
[**GetStoragesStatistics**](StorageAPI.md#GetStoragesStatistics) | **Get** /api/v2/storages/statistics | Get statistics for all Storages
[**UpdateStorage**](StorageAPI.md#UpdateStorage) | **Patch** /api/v2/storages/{storageId} | Updates a Storage
[**UpdateStorageNetworkDeviceConfiguration**](StorageAPI.md#UpdateStorageNetworkDeviceConfiguration) | **Patch** /api/v2/storages/{storageId}/network-device-configurations/{storageNetworkDeviceConfigurationId} | Updates a Storage Network Device Configuration



## CreateStorage

> RegisterStorageResponse CreateStorage(ctx).CreateStorage(createStorage).Execute()

Creates a Storage



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
	createStorage := *openapiclient.NewCreateStorage(float32(123), "Driver_example", []string{"Technologies_example"}, "Name_example", "ManagementHost_example", "SubnetType_example") // CreateStorage | The Storage create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.CreateStorage(context.Background()).CreateStorage(createStorage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.CreateStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStorage`: RegisterStorageResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.CreateStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorage** | [**CreateStorage**](CreateStorage.md) | The Storage create object | 

### Return type

[**RegisterStorageResponse**](RegisterStorageResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorageNetworkDeviceConfiguration

> CreateStorageNetworkDeviceConfiguration(ctx, storageId).CreateStorageNetworkDeviceConfiguration(createStorageNetworkDeviceConfiguration).Execute()

Creates a Storage Network Device Configuration



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
	storageId := float32(8.14) // float32 | 
	createStorageNetworkDeviceConfiguration := *openapiclient.NewCreateStorageNetworkDeviceConfiguration(float32(123), "StoragePhysicalInterfaceIdentifier_example", "NetworkDeviceInterfaceIdentifier_example") // CreateStorageNetworkDeviceConfiguration | The new Storage Network Device Configuration object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageAPI.CreateStorageNetworkDeviceConfiguration(context.Background(), storageId).CreateStorageNetworkDeviceConfiguration(createStorageNetworkDeviceConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.CreateStorageNetworkDeviceConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageNetworkDeviceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStorageNetworkDeviceConfiguration** | [**CreateStorageNetworkDeviceConfiguration**](CreateStorageNetworkDeviceConfiguration.md) | The new Storage Network Device Configuration object | 

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


## DeleteStorage

> DeleteStorage(ctx, storageId).IfMatch(ifMatch).Execute()

Deletes a Storage



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
	storageId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageAPI.DeleteStorage(context.Background(), storageId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.DeleteStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageRequest struct via the builder pattern


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


## DeleteStorageNetworkDeviceConfiguration

> DeleteStorageNetworkDeviceConfiguration(ctx, storageId, storageNetworkDeviceConfigurationId).Execute()

Deletes a Storage Network Device Configuration



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
	storageId := float32(8.14) // float32 | 
	storageNetworkDeviceConfigurationId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageAPI.DeleteStorageNetworkDeviceConfiguration(context.Background(), storageId, storageNetworkDeviceConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.DeleteStorageNetworkDeviceConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 
**storageNetworkDeviceConfigurationId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageNetworkDeviceConfigurationRequest struct via the builder pattern


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


## GetStorage

> Storage GetStorage(ctx, storageId).Execute()

Retrieves a Storage



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
	storageId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorage(context.Background(), storageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorage`: Storage
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Storage**](Storage.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageBuckets

> BucketPaginatedList GetStorageBuckets(ctx, storageId).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterStoragePoolId(filterStoragePoolId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Buckets linked to the specified storage



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
	storageId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe&filter.serviceStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.  **Format:** filter.storagePoolId={$not}:OPERATION:VALUE    **Example:** filter.storagePoolId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=infrastructureId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - infrastructureId  - storagePoolId  - serviceStatus  - config.deployStatus  - config.deployType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,storagePoolId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - storagePoolId  - serviceStatus  - infrastructureId  - config.deployStatus  - config.deployType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageBuckets(context.Background(), storageId).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterStoragePoolId(filterStoragePoolId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageBuckets`: BucketPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.  **Format:** filter.subdomain&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomain&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe&amp;filter.serviceStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.  **Format:** filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.storagePoolId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;infrastructureId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - infrastructureId  - storagePoolId  - serviceStatus  - config.deployStatus  - config.deployType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,storagePoolId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - storagePoolId  - serviceStatus  - infrastructureId  - config.deployStatus  - config.deployType  | 

### Return type

[**BucketPaginatedList**](BucketPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageCredentials

> StorageCredentials GetStorageCredentials(ctx, storageId).Execute()

Get Storage credentials



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
	storageId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageCredentials(context.Background(), storageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageCredentials`: StorageCredentials
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageCredentials**](StorageCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageDrives

> SharedDrivePaginatedList GetStorageDrives(ctx, storageId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterInfrastructureId(filterInfrastructureId).FilterStoragePoolId(filterStoragePoolId).FilterServiceStatus(filterServiceStatus).FilterWwn(filterWwn).FilterQoS(filterQoS).FilterLogicalNetworkId(filterLogicalNetworkId).FilterAllocationAffinity(filterAllocationAffinity).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterConfigLogicalNetworkId(filterConfigLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Drives linked to the specified storage



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
	storageId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.  **Format:** filter.storagePoolId={$not}:OPERATION:VALUE    **Example:** filter.storagePoolId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterWwn := []string{"Inner_example"} // []string | Filter by wwn query param.  **Format:** filter.wwn={$not}:OPERATION:VALUE    **Example:** filter.wwn=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterQoS := []string{"Inner_example"} // []string | Filter by qoS query param.  **Format:** filter.qoS={$not}:OPERATION:VALUE    **Example:** filter.qoS=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLogicalNetworkId := []string{"Inner_example"} // []string | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterAllocationAffinity := []string{"Inner_example"} // []string | Filter by allocationAffinity query param.  **Format:** filter.allocationAffinity={$not}:OPERATION:VALUE    **Example:** filter.allocationAffinity=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterProvisioningProtocol := []string{"Inner_example"} // []string | Filter by provisioningProtocol query param.  **Format:** filter.provisioningProtocol={$not}:OPERATION:VALUE    **Example:** filter.provisioningProtocol=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigLogicalNetworkId := []string{"Inner_example"} // []string | Filter by config.logicalNetworkId query param.  **Format:** filter.config.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.config.logicalNetworkId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=storagePoolId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - storagePoolId  - infrastructureId  - serviceStatus  - config.deployStatus  - config.deployType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,infrastructureId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - infrastructureId  - storagePoolId  - serviceStatus  - wwn  - qoS  - logicalNetworkId  - allocationAffinity  - provisioningProtocol  - config.deployStatus  - config.deployType  - config.logicalNetworkId  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageDrives(context.Background(), storageId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterInfrastructureId(filterInfrastructureId).FilterStoragePoolId(filterStoragePoolId).FilterServiceStatus(filterServiceStatus).FilterWwn(filterWwn).FilterQoS(filterQoS).FilterLogicalNetworkId(filterLogicalNetworkId).FilterAllocationAffinity(filterAllocationAffinity).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterConfigLogicalNetworkId(filterConfigLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageDrives`: SharedDrivePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageDrives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.  **Format:** filter.subdomain&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomain&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.  **Format:** filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.storagePoolId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterWwn** | **[]string** | Filter by wwn query param.  **Format:** filter.wwn&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.wwn&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterQoS** | **[]string** | Filter by qoS query param.  **Format:** filter.qoS&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.qoS&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLogicalNetworkId** | **[]string** | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterAllocationAffinity** | **[]string** | Filter by allocationAffinity query param.  **Format:** filter.allocationAffinity&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.allocationAffinity&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterProvisioningProtocol** | **[]string** | Filter by provisioningProtocol query param.  **Format:** filter.provisioningProtocol&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.provisioningProtocol&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigLogicalNetworkId** | **[]string** | Filter by config.logicalNetworkId query param.  **Format:** filter.config.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.logicalNetworkId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;storagePoolId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - storagePoolId  - infrastructureId  - serviceStatus  - config.deployStatus  - config.deployType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,infrastructureId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - infrastructureId  - storagePoolId  - serviceStatus  - wwn  - qoS  - logicalNetworkId  - allocationAffinity  - provisioningProtocol  - config.deployStatus  - config.deployType  - config.logicalNetworkId  | 

### Return type

[**SharedDrivePaginatedList**](SharedDrivePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageFileShares

> FileSharePaginatedList GetStorageFileShares(ctx, storageId).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterStoragePoolId(filterStoragePoolId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all File Shares linked to the specified storage



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
	storageId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe&filter.serviceStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.  **Format:** filter.storagePoolId={$not}:OPERATION:VALUE    **Example:** filter.storagePoolId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=infrastructureId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - infrastructureId  - storagePoolId  - serviceStatus  - config.deployStatus  - config.deployType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,storagePoolId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - storagePoolId  - serviceStatus  - infrastructureId  - config.deployStatus  - config.deployType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageFileShares(context.Background(), storageId).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterStoragePoolId(filterStoragePoolId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageFileShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageFileShares`: FileSharePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageFileShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageFileSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.  **Format:** filter.subdomain&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomain&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe&amp;filter.serviceStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.  **Format:** filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.storagePoolId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;infrastructureId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - infrastructureId  - storagePoolId  - serviceStatus  - config.deployStatus  - config.deployType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,storagePoolId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - storagePoolId  - serviceStatus  - infrastructureId  - config.deployStatus  - config.deployType  | 

### Return type

[**FileSharePaginatedList**](FileSharePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageNetworkDeviceConfigurations

> []StorageNetworkDeviceConfiguration GetStorageNetworkDeviceConfigurations(ctx, storageId).Execute()

Retrieves Storage Network Device Configurations



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
	storageId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageNetworkDeviceConfigurations(context.Background(), storageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageNetworkDeviceConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageNetworkDeviceConfigurations`: []StorageNetworkDeviceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageNetworkDeviceConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageNetworkDeviceConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StorageNetworkDeviceConfiguration**](StorageNetworkDeviceConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageStatistics

> StorageStatistics GetStorageStatistics(ctx, storageId).Execute()

Get Storages statistics



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
	storageId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageStatistics(context.Background(), storageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageStatistics`: StorageStatistics
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageStatistics**](StorageStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorages

> StoragePaginatedList GetStorages(ctx).Page(page).Limit(limit).FilterId(filterId).FilterUserId(filterUserId).FilterSiteId(filterSiteId).FilterTechnologies(filterTechnologies).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Storages



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUserId := []string{"Inner_example"} // []string | Filter by userId query param.  **Format:** filter.userId={$not}:OPERATION:VALUE    **Example:** filter.userId=$btw:John Doe&filter.userId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterTechnologies := []string{"Inner_example"} // []string | Filter by technologies query param.  **Format:** filter.technologies={$not}:OPERATION:VALUE    **Example:** filter.technologies=$in:John Doe  **Available Operations** - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=userId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - userId  - siteId  - driver  - technologies  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,userId,siteId,driver,technologies   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - userId  - siteId  - driver  - technologies  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorages(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterUserId(filterUserId).FilterSiteId(filterSiteId).FilterTechnologies(filterTechnologies).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorages`: StoragePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUserId** | **[]string** | Filter by userId query param.  **Format:** filter.userId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.userId&#x3D;$btw:John Doe&amp;filter.userId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterTechnologies** | **[]string** | Filter by technologies query param.  **Format:** filter.technologies&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.technologies&#x3D;$in:John Doe  **Available Operations** - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;userId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - userId  - siteId  - driver  - technologies  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,userId,siteId,driver,technologies   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - userId  - siteId  - driver  - technologies  | 

### Return type

[**StoragePaginatedList**](StoragePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoragesStatistics

> StoragesStatistics GetStoragesStatistics(ctx).IncludeMaintenance(includeMaintenance).IncludeExperimental(includeExperimental).MinimumSpace(minimumSpace).Execute()

Get statistics for all Storages



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
	includeMaintenance := "includeMaintenance_example" // string | Include storages in maintenance. (optional)
	includeExperimental := "includeExperimental_example" // string | Include experimental storages. (optional)
	minimumSpace := float32(8.14) // float32 | Minimum space. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStoragesStatistics(context.Background()).IncludeMaintenance(includeMaintenance).IncludeExperimental(includeExperimental).MinimumSpace(minimumSpace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStoragesStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStoragesStatistics`: StoragesStatistics
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStoragesStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoragesStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeMaintenance** | **string** | Include storages in maintenance. | 
 **includeExperimental** | **string** | Include experimental storages. | 
 **minimumSpace** | **float32** | Minimum space. | 

### Return type

[**StoragesStatistics**](StoragesStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStorage

> Storage UpdateStorage(ctx, storageId).UpdateStorage(updateStorage).IfMatch(ifMatch).Execute()

Updates a Storage



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
	storageId := float32(8.14) // float32 | 
	updateStorage := *openapiclient.NewUpdateStorage() // UpdateStorage | The Storage update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.UpdateStorage(context.Background(), storageId).UpdateStorage(updateStorage).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.UpdateStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStorage`: Storage
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.UpdateStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStorage** | [**UpdateStorage**](UpdateStorage.md) | The Storage update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Storage**](Storage.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStorageNetworkDeviceConfiguration

> UpdateStorageNetworkDeviceConfiguration(ctx, storageId, storageNetworkDeviceConfigurationId).UpdateStorageNetworkDeviceConfiguration(updateStorageNetworkDeviceConfiguration).Execute()

Updates a Storage Network Device Configuration



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
	storageId := float32(8.14) // float32 | 
	storageNetworkDeviceConfigurationId := float32(8.14) // float32 | 
	updateStorageNetworkDeviceConfiguration := *openapiclient.NewUpdateStorageNetworkDeviceConfiguration() // UpdateStorageNetworkDeviceConfiguration | The updated Storage Network Device Configuration object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageAPI.UpdateStorageNetworkDeviceConfiguration(context.Background(), storageId, storageNetworkDeviceConfigurationId).UpdateStorageNetworkDeviceConfiguration(updateStorageNetworkDeviceConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.UpdateStorageNetworkDeviceConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 
**storageNetworkDeviceConfigurationId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageNetworkDeviceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateStorageNetworkDeviceConfiguration** | [**UpdateStorageNetworkDeviceConfiguration**](UpdateStorageNetworkDeviceConfiguration.md) | The updated Storage Network Device Configuration object | 

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

