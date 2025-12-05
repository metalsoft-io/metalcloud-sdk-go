# \FileShareAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFileShareSnapshot**](FileShareAPI.md#CreateFileShareSnapshot) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/snapshots | Create a snapshot of the specified File Share
[**CreateInfrastructureFileShare**](FileShareAPI.md#CreateInfrastructureFileShare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares | Creates a File Share
[**DeleteFileShare**](FileShareAPI.md#DeleteFileShare) | **Delete** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Deletes a File Share
[**DeleteFileShareSnapshot**](FileShareAPI.md#DeleteFileShareSnapshot) | **Delete** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/snapshots | Delete a snapshot of the specified File Share
[**GetFileShare**](FileShareAPI.md#GetFileShare) | **Get** /api/v2/file-shares/{fileShareId} | Get File Share information
[**GetFileShareConfigInfo**](FileShareAPI.md#GetFileShareConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/config | Get configuration information about the specified File Share
[**GetFileShareHosts**](FileShareAPI.md#GetFileShareHosts) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/hosts | Get the Hosts of File Share
[**GetFileShareSnapshots**](FileShareAPI.md#GetFileShareSnapshots) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/snapshots | Get snapshots of the specified File Share
[**GetInfrastructureFileShare**](FileShareAPI.md#GetInfrastructureFileShare) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Get File Share information
[**GetInfrastructureFileShares**](FileShareAPI.md#GetInfrastructureFileShares) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares | Get all File Shares
[**PatchFileShareMeta**](FileShareAPI.md#PatchFileShareMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/meta | Updates the meta of a File Share
[**RestoreFileShareToSnapshot**](FileShareAPI.md#RestoreFileShareToSnapshot) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/snapshots/restore | Restore a File Share to a specified snapshot
[**UpdateFileShareConfig**](FileShareAPI.md#UpdateFileShareConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/config | Updates File Share config information
[**UpdateFileShareInstanceArrayHostsBulk**](FileShareAPI.md#UpdateFileShareInstanceArrayHostsBulk) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/modify-instance-array-hosts-bulk | Updates Instance Array Hosts on the File Share



## CreateFileShareSnapshot

> FileShareSnapshot CreateFileShareSnapshot(ctx, infrastructureId, fileShareId).Execute()

Create a snapshot of the specified File Share

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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.CreateFileShareSnapshot(context.Background(), infrastructureId, fileShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.CreateFileShareSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileShareSnapshot`: FileShareSnapshot
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.CreateFileShareSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileShareSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileShareSnapshot**](FileShareSnapshot.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInfrastructureFileShare

> FileShare CreateInfrastructureFileShare(ctx, infrastructureId).CreateFileShare(createFileShare).Execute()

Creates a File Share



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
	infrastructureId := float32(8.14) // float32 | 
	createFileShare := *openapiclient.NewCreateFileShare(float32(123), float32(123)) // CreateFileShare | The File Share create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.CreateInfrastructureFileShare(context.Background(), infrastructureId).CreateFileShare(createFileShare).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.CreateInfrastructureFileShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInfrastructureFileShare`: FileShare
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.CreateInfrastructureFileShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInfrastructureFileShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileShare** | [**CreateFileShare**](CreateFileShare.md) | The File Share create object | 

### Return type

[**FileShare**](FileShare.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileShare

> DeleteFileShare(ctx, infrastructureId, fileShareId).IfMatch(ifMatch).Execute()

Deletes a File Share



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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileShareAPI.DeleteFileShare(context.Background(), infrastructureId, fileShareId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.DeleteFileShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileShareRequest struct via the builder pattern


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


## DeleteFileShareSnapshot

> DeleteFileShareSnapshot(ctx, infrastructureId, fileShareId).DeleteFileShareSnapshot(deleteFileShareSnapshot).Execute()

Delete a snapshot of the specified File Share

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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 
	deleteFileShareSnapshot := *openapiclient.NewDeleteFileShareSnapshot("Name_example") // DeleteFileShareSnapshot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileShareAPI.DeleteFileShareSnapshot(context.Background(), infrastructureId, fileShareId).DeleteFileShareSnapshot(deleteFileShareSnapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.DeleteFileShareSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileShareSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteFileShareSnapshot** | [**DeleteFileShareSnapshot**](DeleteFileShareSnapshot.md) |  | 

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


## GetFileShare

> FileShare GetFileShare(ctx, fileShareId).Execute()

Get File Share information



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
	fileShareId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.GetFileShare(context.Background(), fileShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.GetFileShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileShare`: FileShare
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.GetFileShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileShare**](FileShare.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileShareConfigInfo

> FileShareConfiguration GetFileShareConfigInfo(ctx, infrastructureId, fileShareId).Execute()

Get configuration information about the specified File Share

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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.GetFileShareConfigInfo(context.Background(), infrastructureId, fileShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.GetFileShareConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileShareConfigInfo`: FileShareConfiguration
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.GetFileShareConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileShareConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileShareConfiguration**](FileShareConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileShareHosts

> FileShareHosts GetFileShareHosts(ctx, infrastructureId, fileShareId).Execute()

Get the Hosts of File Share



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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.GetFileShareHosts(context.Background(), infrastructureId, fileShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.GetFileShareHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileShareHosts`: FileShareHosts
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.GetFileShareHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileShareHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileShareHosts**](FileShareHosts.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileShareSnapshots

> []FileShareSnapshot GetFileShareSnapshots(ctx, infrastructureId, fileShareId).Execute()

Get snapshots of the specified File Share

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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.GetFileShareSnapshots(context.Background(), infrastructureId, fileShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.GetFileShareSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileShareSnapshots`: []FileShareSnapshot
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.GetFileShareSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileShareSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]FileShareSnapshot**](FileShareSnapshot.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureFileShare

> FileShare GetInfrastructureFileShare(ctx, infrastructureId, fileShareId).Execute()

Get File Share information



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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.GetInfrastructureFileShare(context.Background(), infrastructureId, fileShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.GetInfrastructureFileShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureFileShare`: FileShare
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.GetInfrastructureFileShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureFileShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileShare**](FileShare.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureFileShares

> FileSharePaginatedList GetInfrastructureFileShares(ctx, infrastructureId).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterStoragePoolId(filterStoragePoolId).FilterLogicalNetworkId(filterLogicalNetworkId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterConfigLogicalNetworkId(filterConfigLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all File Shares



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
	infrastructureId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe&filter.serviceStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.  **Format:** filter.storagePoolId={$not}:OPERATION:VALUE    **Example:** filter.storagePoolId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLogicalNetworkId := []string{"Inner_example"} // []string | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigLogicalNetworkId := []string{"Inner_example"} // []string | Filter by config.logicalNetworkId query param.  **Format:** filter.config.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.config.logicalNetworkId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=infrastructureId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - infrastructureId  - storagePoolId  - serviceStatus  - config.deployStatus  - config.deployType  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,storagePoolId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - storagePoolId  - serviceStatus  - infrastructureId  - logicalNetworkId  - config.deployStatus  - config.deployType  - config.logicalNetworkId  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.GetInfrastructureFileShares(context.Background(), infrastructureId).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterStoragePoolId(filterStoragePoolId).FilterLogicalNetworkId(filterLogicalNetworkId).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterConfigLogicalNetworkId(filterConfigLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.GetInfrastructureFileShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureFileShares`: FileSharePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.GetInfrastructureFileShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureFileSharesRequest struct via the builder pattern


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
 **filterLogicalNetworkId** | **[]string** | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigLogicalNetworkId** | **[]string** | Filter by config.logicalNetworkId query param.  **Format:** filter.config.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.logicalNetworkId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;infrastructureId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - infrastructureId  - storagePoolId  - serviceStatus  - config.deployStatus  - config.deployType  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,storagePoolId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - storagePoolId  - serviceStatus  - infrastructureId  - logicalNetworkId  - config.deployStatus  - config.deployType  - config.logicalNetworkId  | 

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


## PatchFileShareMeta

> FileShare PatchFileShareMeta(ctx, infrastructureId, fileShareId).UpdateFileShareMeta(updateFileShareMeta).Execute()

Updates the meta of a File Share

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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 
	updateFileShareMeta := *openapiclient.NewUpdateFileShareMeta("Name_example") // UpdateFileShareMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.PatchFileShareMeta(context.Background(), infrastructureId, fileShareId).UpdateFileShareMeta(updateFileShareMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.PatchFileShareMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFileShareMeta`: FileShare
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.PatchFileShareMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFileShareMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFileShareMeta** | [**UpdateFileShareMeta**](UpdateFileShareMeta.md) |  | 

### Return type

[**FileShare**](FileShare.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFileShareToSnapshot

> RestoreFileShareToSnapshot(ctx, infrastructureId, fileShareId).RestoreFileShareSnapshot(restoreFileShareSnapshot).Execute()

Restore a File Share to a specified snapshot

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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 
	restoreFileShareSnapshot := *openapiclient.NewRestoreFileShareSnapshot("Name_example") // RestoreFileShareSnapshot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileShareAPI.RestoreFileShareToSnapshot(context.Background(), infrastructureId, fileShareId).RestoreFileShareSnapshot(restoreFileShareSnapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.RestoreFileShareToSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFileShareToSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restoreFileShareSnapshot** | [**RestoreFileShareSnapshot**](RestoreFileShareSnapshot.md) |  | 

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


## UpdateFileShareConfig

> FileShareConfiguration UpdateFileShareConfig(ctx, infrastructureId, fileShareId).UpdateFileShare(updateFileShare).IfMatch(ifMatch).Execute()

Updates File Share config information



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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 
	updateFileShare := *openapiclient.NewUpdateFileShare() // UpdateFileShare | The File Share update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.UpdateFileShareConfig(context.Background(), infrastructureId, fileShareId).UpdateFileShare(updateFileShare).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.UpdateFileShareConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileShareConfig`: FileShareConfiguration
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.UpdateFileShareConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileShareConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFileShare** | [**UpdateFileShare**](UpdateFileShare.md) | The File Share update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**FileShareConfiguration**](FileShareConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileShareInstanceArrayHostsBulk

> FileShareHosts UpdateFileShareInstanceArrayHostsBulk(ctx, infrastructureId, fileShareId).FileShareHostsModifyBulk(fileShareHostsModifyBulk).Execute()

Updates Instance Array Hosts on the File Share



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
	infrastructureId := float32(8.14) // float32 | 
	fileShareId := float32(8.14) // float32 | 
	fileShareHostsModifyBulk := *openapiclient.NewFileShareHostsModifyBulk([]openapiclient.FileShareHostBulkOperation{*openapiclient.NewFileShareHostBulkOperation(float32(123), "OperationType_example")}) // FileShareHostsModifyBulk | The File Share Instance Array Hosts update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileShareAPI.UpdateFileShareInstanceArrayHostsBulk(context.Background(), infrastructureId, fileShareId).FileShareHostsModifyBulk(fileShareHostsModifyBulk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileShareAPI.UpdateFileShareInstanceArrayHostsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileShareInstanceArrayHostsBulk`: FileShareHosts
	fmt.Fprintf(os.Stdout, "Response from `FileShareAPI.UpdateFileShareInstanceArrayHostsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**fileShareId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileShareInstanceArrayHostsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileShareHostsModifyBulk** | [**FileShareHostsModifyBulk**](FileShareHostsModifyBulk.md) | The File Share Instance Array Hosts update object | 

### Return type

[**FileShareHosts**](FileShareHosts.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

