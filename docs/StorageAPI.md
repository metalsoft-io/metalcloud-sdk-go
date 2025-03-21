# \StorageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorage**](StorageAPI.md#CreateStorage) | **Post** /api/v2/storages | Creates a Storage
[**CreateStorageNetworkDeviceConfiguration**](StorageAPI.md#CreateStorageNetworkDeviceConfiguration) | **Post** /api/v2/storages/{storageId}/network-device-configurations | Creates a Storage Network Device Configuration
[**DeleteStorage**](StorageAPI.md#DeleteStorage) | **Delete** /api/v2/storages/{storageId} | Deletes a Storage
[**DeleteStorageNetworkDeviceConfiguration**](StorageAPI.md#DeleteStorageNetworkDeviceConfiguration) | **Delete** /api/v2/storages/{storageId}/network-device-configurations/{storageNetworkDeviceConfigurationId} | Deletes a Storage Network Device Configuration
[**ExecuteStorageAction**](StorageAPI.md#ExecuteStorageAction) | **Post** /api/v2/storages/{storageId}/actions/update-cache | Execute a Storage action
[**GetStorage**](StorageAPI.md#GetStorage) | **Get** /api/v2/storages/{storageId} | Retrieves a Storage
[**GetStorageBuckets**](StorageAPI.md#GetStorageBuckets) | **Get** /api/v2/storages/{storageId}/buckets | Get all Buckets linked to the specified storage
[**GetStorageDrives**](StorageAPI.md#GetStorageDrives) | **Get** /api/v2/storages/{storageId}/drives | Get all Drives linked to the specified storage
[**GetStorageFileShares**](StorageAPI.md#GetStorageFileShares) | **Get** /api/v2/storages/{storageId}/file-shares | Get all File Shares linked to the specified storage
[**GetStorageIscsiBootServers**](StorageAPI.md#GetStorageIscsiBootServers) | **Get** /api/v2/storages/{storageId}/iscsi-boot-servers | Returns information about servers which are setup to boot from iSCSI block devices
[**GetStorageNetworkDeviceConfigurations**](StorageAPI.md#GetStorageNetworkDeviceConfigurations) | **Get** /api/v2/storages/{storageId}/network-device-configurations | Retrieves Storage Network Device Configurations
[**GetStorageSharedDrives**](StorageAPI.md#GetStorageSharedDrives) | **Get** /api/v2/storages/{storageId}/shared-drives | Get all Shared Drives linked to the specified storage
[**GetStorageStatistics**](StorageAPI.md#GetStorageStatistics) | **Get** /api/v2/storages/statistics | Get Storages statistics
[**GetStorages**](StorageAPI.md#GetStorages) | **Get** /api/v2/storages | Get a list of Storages
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
	createStorage := *openapiclient.NewCreateStorage(float32(123), "StorageDriver_example", "StorageTechnology_example", "StorageType_example", "Name_example", "ManagementHost_example", "Username_example", "PasswordEncrypted_example", float32(123), "SubnetType_example", "Password_example") // CreateStorage | The Storage create object

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
	createStorageNetworkDeviceConfiguration := *openapiclient.NewCreateStorageNetworkDeviceConfiguration(float32(123), "StoragePhysicalInterfaceIdentifier_example", "NetworkEquipmentInterfaceIdentifier_example", []float32{float32(123)}) // CreateStorageNetworkDeviceConfiguration | The new Storage Network Device Configuration object

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

> DeleteStorage(ctx, storageId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageAPI.DeleteStorage(context.Background(), storageId).Execute()
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


## ExecuteStorageAction

> ExecuteStorageAction(ctx, storageId).Execute()

Execute a Storage action



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
	r, err := apiClient.StorageAPI.ExecuteStorageAction(context.Background(), storageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.ExecuteStorageAction``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExecuteStorageActionRequest struct via the builder pattern


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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.           <p>              <b>Format: </b> filter.subdomain={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomain=$not:$like:John Doe&filter.subdomain=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.           <p>              <b>Format: </b> filter.subdomainPermanent={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomainPermanent=$not:$like:John Doe&filter.subdomainPermanent=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.           <p>              <b>Format: </b> filter.storagePoolId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storagePoolId=$not:$like:John Doe&filter.storagePoolId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>infrastructureId</li> <li>storagePoolId</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,storagePoolId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>storagePoolId</li> <li>serviceStatus</li> <li>infrastructureId</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>          (optional)

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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storagePoolId&#x3D;$not:$like:John Doe&amp;filter.storagePoolId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,storagePoolId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;          | 

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


## GetStorageDrives

> DrivePaginatedList GetStorageDrives(ctx, storageId).Page(page).Limit(limit).FilterId(filterId).FilterInstanceId(filterInstanceId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterStoragePoolId(filterStoragePoolId).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterInstanceId := []string{"Inner_example"} // []string | Filter by instanceId query param.           <p>              <b>Format: </b> filter.instanceId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.instanceId=$not:$like:John Doe&filter.instanceId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.           <p>              <b>Format: </b> filter.subdomain={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomain=$not:$like:John Doe&filter.subdomain=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.           <p>              <b>Format: </b> filter.subdomainPermanent={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomainPermanent=$not:$like:John Doe&filter.subdomainPermanent=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.           <p>              <b>Format: </b> filter.storagePoolId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storagePoolId=$not:$like:John Doe&filter.storagePoolId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterProvisioningProtocol := []string{"Inner_example"} // []string | Filter by provisioningProtocol query param.           <p>              <b>Format: </b> filter.provisioningProtocol={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.provisioningProtocol=$not:$like:John Doe&filter.provisioningProtocol=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>storagePoolId</li> <li>infrastructureId</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,infrastructureId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>infrastructureId</li> <li>storagePoolId</li> <li>serviceStatus</li> <li>wwn</li> <li>networkVlanId</li> <li>provisioningProtocol</li> <li>templateId</li> <li>config.deployStatus</li> <li>config.deployType</li> <li>config.templateId</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageDrives(context.Background(), storageId).Page(page).Limit(limit).FilterId(filterId).FilterInstanceId(filterInstanceId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterStoragePoolId(filterStoragePoolId).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageDrives`: DrivePaginatedList
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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterInstanceId** | **[]string** | Filter by instanceId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.instanceId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.instanceId&#x3D;$not:$like:John Doe&amp;filter.instanceId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storagePoolId&#x3D;$not:$like:John Doe&amp;filter.storagePoolId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterProvisioningProtocol** | **[]string** | Filter by provisioningProtocol query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.provisioningProtocol&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.provisioningProtocol&#x3D;$not:$like:John Doe&amp;filter.provisioningProtocol&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,infrastructureId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;wwn&lt;/li&gt; &lt;li&gt;networkVlanId&lt;/li&gt; &lt;li&gt;provisioningProtocol&lt;/li&gt; &lt;li&gt;templateId&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt; &lt;li&gt;config.templateId&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**DrivePaginatedList**](DrivePaginatedList.md)

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.           <p>              <b>Format: </b> filter.subdomain={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomain=$not:$like:John Doe&filter.subdomain=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.           <p>              <b>Format: </b> filter.subdomainPermanent={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomainPermanent=$not:$like:John Doe&filter.subdomainPermanent=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.           <p>              <b>Format: </b> filter.storagePoolId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storagePoolId=$not:$like:John Doe&filter.storagePoolId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>infrastructureId</li> <li>storagePoolId</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,storagePoolId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>storagePoolId</li> <li>serviceStatus</li> <li>infrastructureId</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>          (optional)

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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storagePoolId&#x3D;$not:$like:John Doe&amp;filter.storagePoolId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,storagePoolId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;          | 

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


## GetStorageIscsiBootServers

> []IscsiBootServerInfo GetStorageIscsiBootServers(ctx, storageId).Execute()

Returns information about servers which are setup to boot from iSCSI block devices



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
	resp, r, err := apiClient.StorageAPI.GetStorageIscsiBootServers(context.Background(), storageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageIscsiBootServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageIscsiBootServers`: []IscsiBootServerInfo
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageIscsiBootServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageIscsiBootServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IscsiBootServerInfo**](IscsiBootServerInfo.md)

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


## GetStorageSharedDrives

> SharedDrivePaginatedList GetStorageSharedDrives(ctx, storageId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterInfrastructureId(filterInfrastructureId).FilterStoragePoolId(filterStoragePoolId).FilterServiceStatus(filterServiceStatus).FilterWwn(filterWwn).FilterIoLimitPolicy(filterIoLimitPolicy).FilterNetworkVlanId(filterNetworkVlanId).FilterAllocationAffinity(filterAllocationAffinity).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Shared Drives linked to the specified storage



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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.           <p>              <b>Format: </b> filter.subdomain={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomain=$not:$like:John Doe&filter.subdomain=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.           <p>              <b>Format: </b> filter.subdomainPermanent={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomainPermanent=$not:$like:John Doe&filter.subdomainPermanent=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.           <p>              <b>Format: </b> filter.storagePoolId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storagePoolId=$not:$like:John Doe&filter.storagePoolId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterWwn := []string{"Inner_example"} // []string | Filter by wwn query param.           <p>              <b>Format: </b> filter.wwn={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.wwn=$not:$like:John Doe&filter.wwn=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterIoLimitPolicy := []string{"Inner_example"} // []string | Filter by ioLimitPolicy query param.           <p>              <b>Format: </b> filter.ioLimitPolicy={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.ioLimitPolicy=$not:$like:John Doe&filter.ioLimitPolicy=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterNetworkVlanId := []string{"Inner_example"} // []string | Filter by networkVlanId query param.           <p>              <b>Format: </b> filter.networkVlanId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.networkVlanId=$not:$like:John Doe&filter.networkVlanId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAllocationAffinity := []string{"Inner_example"} // []string | Filter by allocationAffinity query param.           <p>              <b>Format: </b> filter.allocationAffinity={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.allocationAffinity=$not:$like:John Doe&filter.allocationAffinity=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterProvisioningProtocol := []string{"Inner_example"} // []string | Filter by provisioningProtocol query param.           <p>              <b>Format: </b> filter.provisioningProtocol={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.provisioningProtocol=$not:$like:John Doe&filter.provisioningProtocol=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>storagePoolId</li> <li>infrastructureId</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,infrastructureId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>infrastructureId</li> <li>storagePoolId</li> <li>serviceStatus</li> <li>wwn</li> <li>ioLimitPolicy</li> <li>networkVlanId</li> <li>allocationAffinity</li> <li>provisioningProtocol</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageSharedDrives(context.Background(), storageId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterInfrastructureId(filterInfrastructureId).FilterStoragePoolId(filterStoragePoolId).FilterServiceStatus(filterServiceStatus).FilterWwn(filterWwn).FilterIoLimitPolicy(filterIoLimitPolicy).FilterNetworkVlanId(filterNetworkVlanId).FilterAllocationAffinity(filterAllocationAffinity).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageSharedDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageSharedDrives`: SharedDrivePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageSharedDrives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageSharedDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storagePoolId&#x3D;$not:$like:John Doe&amp;filter.storagePoolId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterWwn** | **[]string** | Filter by wwn query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.wwn&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.wwn&#x3D;$not:$like:John Doe&amp;filter.wwn&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterIoLimitPolicy** | **[]string** | Filter by ioLimitPolicy query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.ioLimitPolicy&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.ioLimitPolicy&#x3D;$not:$like:John Doe&amp;filter.ioLimitPolicy&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterNetworkVlanId** | **[]string** | Filter by networkVlanId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.networkVlanId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.networkVlanId&#x3D;$not:$like:John Doe&amp;filter.networkVlanId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterAllocationAffinity** | **[]string** | Filter by allocationAffinity query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.allocationAffinity&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.allocationAffinity&#x3D;$not:$like:John Doe&amp;filter.allocationAffinity&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterProvisioningProtocol** | **[]string** | Filter by provisioningProtocol query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.provisioningProtocol&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.provisioningProtocol&#x3D;$not:$like:John Doe&amp;filter.provisioningProtocol&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,infrastructureId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;wwn&lt;/li&gt; &lt;li&gt;ioLimitPolicy&lt;/li&gt; &lt;li&gt;networkVlanId&lt;/li&gt; &lt;li&gt;allocationAffinity&lt;/li&gt; &lt;li&gt;provisioningProtocol&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;          | 

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


## GetStorageStatistics

> map[string]interface{} GetStorageStatistics(ctx).IncludeMaintenance(includeMaintenance).IncludeExperimental(includeExperimental).MinimumSpace(minimumSpace).Execute()

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
	includeMaintenance := "includeMaintenance_example" // string | Include storages in maintenance. (optional)
	includeExperimental := "includeExperimental_example" // string | Include experimental storages. (optional)
	minimumSpace := float32(8.14) // float32 | Minimum space. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageStatistics(context.Background()).IncludeMaintenance(includeMaintenance).IncludeExperimental(includeExperimental).MinimumSpace(minimumSpace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageStatistics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeMaintenance** | **string** | Include storages in maintenance. | 
 **includeExperimental** | **string** | Include experimental storages. | 
 **minimumSpace** | **float32** | Minimum space. | 

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


## GetStorages

> StoragePaginatedList GetStorages(ctx).Page(page).Limit(limit).FilterStorageId(filterStorageId).FilterUserId(filterUserId).FilterSiteId(filterSiteId).FilterStorageTechnology(filterStorageTechnology).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterStorageId := []string{"Inner_example"} // []string | Filter by storageId query param.           <p>              <b>Format: </b> filter.storageId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storageId=$not:$like:John Doe&filter.storageId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterUserId := []string{"Inner_example"} // []string | Filter by userId query param.           <p>              <b>Format: </b> filter.userId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.userId=$not:$like:John Doe&filter.userId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterStorageTechnology := []string{"Inner_example"} // []string | Filter by storageTechnology query param.           <p>              <b>Format: </b> filter.storageTechnology={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storageTechnology=$not:$like:John Doe&filter.storageTechnology=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> storageId:DESC           </p>       <h4>Available Fields</h4><ul><li>storageId</li> <li>userId</li> <li>siteId</li> <li>storageDriver</li> <li>storageTechnology</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> storageId,userId,siteId,storageDriver,storageTechnology           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>storageId</li> <li>userId</li> <li>siteId</li> <li>storageDriver</li> <li>storageTechnology</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorages(context.Background()).Page(page).Limit(limit).FilterStorageId(filterStorageId).FilterUserId(filterUserId).FilterSiteId(filterSiteId).FilterStorageTechnology(filterStorageTechnology).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterStorageId** | **[]string** | Filter by storageId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storageId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storageId&#x3D;$not:$like:John Doe&amp;filter.storageId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterUserId** | **[]string** | Filter by userId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.userId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.userId&#x3D;$not:$like:John Doe&amp;filter.userId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterStorageTechnology** | **[]string** | Filter by storageTechnology query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storageTechnology&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storageTechnology&#x3D;$not:$like:John Doe&amp;filter.storageTechnology&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; storageId:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;storageId&lt;/li&gt; &lt;li&gt;userId&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;storageDriver&lt;/li&gt; &lt;li&gt;storageTechnology&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; storageId,userId,siteId,storageDriver,storageTechnology           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;storageId&lt;/li&gt; &lt;li&gt;userId&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;storageDriver&lt;/li&gt; &lt;li&gt;storageTechnology&lt;/li&gt;&lt;/ul&gt;          | 

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


## UpdateStorage

> Storage UpdateStorage(ctx, storageId).UpdateStorageInternal(updateStorageInternal).Execute()

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
	updateStorageInternal := *openapiclient.NewUpdateStorageInternal() // UpdateStorageInternal | The Storage update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.UpdateStorage(context.Background(), storageId).UpdateStorageInternal(updateStorageInternal).Execute()
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

 **updateStorageInternal** | [**UpdateStorageInternal**](UpdateStorageInternal.md) | The Storage update object | 

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

