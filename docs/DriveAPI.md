# \DriveAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrive**](DriveAPI.md#CreateDrive) | **Post** /api/v2/infrastructures/{infrastructureId}/drives | Create a new Drive
[**CreateDriveSnapshot**](DriveAPI.md#CreateDriveSnapshot) | **Post** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/snapshots | Create a snapshot of the specified Drive
[**DeleteDrive**](DriveAPI.md#DeleteDrive) | **Delete** /api/v2/infrastructures/{infrastructureId}/drives/{driveId} | Deletes a Drive
[**DeleteDriveSnapshot**](DriveAPI.md#DeleteDriveSnapshot) | **Delete** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/snapshots | Delete a snapshot of the specified Drive
[**GetDrive**](DriveAPI.md#GetDrive) | **Get** /api/v2/drives/{driveId} | Get Drive information
[**GetDriveConfigInfo**](DriveAPI.md#GetDriveConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/config | Get configuration information about the specified Drive
[**GetDriveHosts**](DriveAPI.md#GetDriveHosts) | **Get** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/hosts | Get the Hosts of Drive
[**GetDriveSnapshots**](DriveAPI.md#GetDriveSnapshots) | **Get** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/snapshots | Get snapshots of the specified Drive
[**GetInfrastructureDrive**](DriveAPI.md#GetInfrastructureDrive) | **Get** /api/v2/infrastructures/{infrastructureId}/drives/{driveId} | Get Drive information
[**GetInfrastructureDrives**](DriveAPI.md#GetInfrastructureDrives) | **Get** /api/v2/infrastructures/{infrastructureId}/drives | Get all Drives on the infrastructure
[**PatchDriveConfig**](DriveAPI.md#PatchDriveConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/config | Updates the config of a Drive
[**PatchDriveMeta**](DriveAPI.md#PatchDriveMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/meta | Updates the meta of a Drive
[**RestoreDriveToSnapshot**](DriveAPI.md#RestoreDriveToSnapshot) | **Post** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/snapshots/restore | Restore a Drive to a specified snapshot
[**UpdateDriveServerInstanceGroupHostsBulk**](DriveAPI.md#UpdateDriveServerInstanceGroupHostsBulk) | **Post** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/actions/modify-server-instance-group-hosts-bulk | Updates Server Instance Group Hosts on the Drive



## CreateDrive

> SharedDrive CreateDrive(ctx, infrastructureId).CreateSharedDrive(createSharedDrive).Execute()

Create a new Drive

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
	createSharedDrive := *openapiclient.NewCreateSharedDrive(float32(123), float32(123), float32(123)) // CreateSharedDrive | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.CreateDrive(context.Background(), infrastructureId).CreateSharedDrive(createSharedDrive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.CreateDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDrive`: SharedDrive
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.CreateDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSharedDrive** | [**CreateSharedDrive**](CreateSharedDrive.md) |  | 

### Return type

[**SharedDrive**](SharedDrive.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDriveSnapshot

> SharedDriveSnapshot CreateDriveSnapshot(ctx, infrastructureId, driveId).Execute()

Create a snapshot of the specified Drive

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
	driveId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.CreateDriveSnapshot(context.Background(), infrastructureId, driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.CreateDriveSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDriveSnapshot`: SharedDriveSnapshot
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.CreateDriveSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriveSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharedDriveSnapshot**](SharedDriveSnapshot.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDrive

> DeleteDrive(ctx, infrastructureId, driveId).IfMatch(ifMatch).Execute()

Deletes a Drive

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
	driveId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DriveAPI.DeleteDrive(context.Background(), infrastructureId, driveId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.DeleteDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriveRequest struct via the builder pattern


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


## DeleteDriveSnapshot

> DeleteDriveSnapshot(ctx, infrastructureId, driveId).DeleteSharedDriveSnapshot(deleteSharedDriveSnapshot).Execute()

Delete a snapshot of the specified Drive

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
	driveId := float32(8.14) // float32 | 
	deleteSharedDriveSnapshot := *openapiclient.NewDeleteSharedDriveSnapshot("Name_example") // DeleteSharedDriveSnapshot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DriveAPI.DeleteDriveSnapshot(context.Background(), infrastructureId, driveId).DeleteSharedDriveSnapshot(deleteSharedDriveSnapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.DeleteDriveSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriveSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteSharedDriveSnapshot** | [**DeleteSharedDriveSnapshot**](DeleteSharedDriveSnapshot.md) |  | 

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


## GetDrive

> SharedDrive GetDrive(ctx, driveId).Execute()

Get Drive information



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
	driveId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.GetDrive(context.Background(), driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.GetDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDrive`: SharedDrive
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.GetDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SharedDrive**](SharedDrive.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDriveConfigInfo

> SharedDriveConfiguration GetDriveConfigInfo(ctx, infrastructureId, driveId).Execute()

Get configuration information about the specified Drive

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
	driveId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.GetDriveConfigInfo(context.Background(), infrastructureId, driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.GetDriveConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveConfigInfo`: SharedDriveConfiguration
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.GetDriveConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharedDriveConfiguration**](SharedDriveConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDriveHosts

> SharedDriveHosts GetDriveHosts(ctx, infrastructureId, driveId).Execute()

Get the Hosts of Drive



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
	driveId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.GetDriveHosts(context.Background(), infrastructureId, driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.GetDriveHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveHosts`: SharedDriveHosts
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.GetDriveHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharedDriveHosts**](SharedDriveHosts.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDriveSnapshots

> []SharedDriveSnapshot GetDriveSnapshots(ctx, infrastructureId, driveId).Execute()

Get snapshots of the specified Drive

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
	driveId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.GetDriveSnapshots(context.Background(), infrastructureId, driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.GetDriveSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveSnapshots`: []SharedDriveSnapshot
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.GetDriveSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SharedDriveSnapshot**](SharedDriveSnapshot.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureDrive

> SharedDrive GetInfrastructureDrive(ctx, infrastructureId, driveId).Execute()

Get Drive information



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
	driveId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.GetInfrastructureDrive(context.Background(), infrastructureId, driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.GetInfrastructureDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureDrive`: SharedDrive
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.GetInfrastructureDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharedDrive**](SharedDrive.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureDrives

> SharedDrivePaginatedList GetInfrastructureDrives(ctx, infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterInfrastructureId(filterInfrastructureId).FilterStoragePoolId(filterStoragePoolId).FilterServiceStatus(filterServiceStatus).FilterWwn(filterWwn).FilterQoS(filterQoS).FilterLogicalNetworkId(filterLogicalNetworkId).FilterAllocationAffinity(filterAllocationAffinity).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterConfigLogicalNetworkId(filterConfigLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Drives on the infrastructure



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
	resp, r, err := apiClient.DriveAPI.GetInfrastructureDrives(context.Background(), infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterInfrastructureId(filterInfrastructureId).FilterStoragePoolId(filterStoragePoolId).FilterServiceStatus(filterServiceStatus).FilterWwn(filterWwn).FilterQoS(filterQoS).FilterLogicalNetworkId(filterLogicalNetworkId).FilterAllocationAffinity(filterAllocationAffinity).FilterProvisioningProtocol(filterProvisioningProtocol).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterConfigLogicalNetworkId(filterConfigLogicalNetworkId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.GetInfrastructureDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureDrives`: SharedDrivePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.GetInfrastructureDrives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureDrivesRequest struct via the builder pattern


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


## PatchDriveConfig

> SharedDrive PatchDriveConfig(ctx, infrastructureId, driveId).UpdateSharedDrive(updateSharedDrive).IfMatch(ifMatch).Execute()

Updates the config of a Drive

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
	driveId := float32(8.14) // float32 | 
	updateSharedDrive := *openapiclient.NewUpdateSharedDrive() // UpdateSharedDrive | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.PatchDriveConfig(context.Background(), infrastructureId, driveId).UpdateSharedDrive(updateSharedDrive).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.PatchDriveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDriveConfig`: SharedDrive
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.PatchDriveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDriveConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSharedDrive** | [**UpdateSharedDrive**](UpdateSharedDrive.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**SharedDrive**](SharedDrive.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDriveMeta

> SharedDrive PatchDriveMeta(ctx, infrastructureId, driveId).UpdateSharedDriveMeta(updateSharedDriveMeta).Execute()

Updates the meta of a Drive

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
	driveId := float32(8.14) // float32 | 
	updateSharedDriveMeta := *openapiclient.NewUpdateSharedDriveMeta("Name_example") // UpdateSharedDriveMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.PatchDriveMeta(context.Background(), infrastructureId, driveId).UpdateSharedDriveMeta(updateSharedDriveMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.PatchDriveMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDriveMeta`: SharedDrive
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.PatchDriveMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDriveMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSharedDriveMeta** | [**UpdateSharedDriveMeta**](UpdateSharedDriveMeta.md) |  | 

### Return type

[**SharedDrive**](SharedDrive.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDriveToSnapshot

> RestoreDriveToSnapshot(ctx, infrastructureId, driveId).RestoreSharedDriveSnapshot(restoreSharedDriveSnapshot).Execute()

Restore a Drive to a specified snapshot

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
	driveId := float32(8.14) // float32 | 
	restoreSharedDriveSnapshot := *openapiclient.NewRestoreSharedDriveSnapshot("Name_example") // RestoreSharedDriveSnapshot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DriveAPI.RestoreDriveToSnapshot(context.Background(), infrastructureId, driveId).RestoreSharedDriveSnapshot(restoreSharedDriveSnapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.RestoreDriveToSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDriveToSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restoreSharedDriveSnapshot** | [**RestoreSharedDriveSnapshot**](RestoreSharedDriveSnapshot.md) |  | 

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


## UpdateDriveServerInstanceGroupHostsBulk

> SharedDriveHosts UpdateDriveServerInstanceGroupHostsBulk(ctx, infrastructureId, driveId).SharedDriveHostsModifyBulk(sharedDriveHostsModifyBulk).Execute()

Updates Server Instance Group Hosts on the Drive



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
	driveId := float32(8.14) // float32 | 
	sharedDriveHostsModifyBulk := *openapiclient.NewSharedDriveHostsModifyBulk([]openapiclient.SharedDriveHostBulkOperation{*openapiclient.NewSharedDriveHostBulkOperation(float32(123), "OperationType_example")}) // SharedDriveHostsModifyBulk | The Drive Server Instance Group Hosts update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveAPI.UpdateDriveServerInstanceGroupHostsBulk(context.Background(), infrastructureId, driveId).SharedDriveHostsModifyBulk(sharedDriveHostsModifyBulk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveAPI.UpdateDriveServerInstanceGroupHostsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDriveServerInstanceGroupHostsBulk`: SharedDriveHosts
	fmt.Fprintf(os.Stdout, "Response from `DriveAPI.UpdateDriveServerInstanceGroupHostsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriveServerInstanceGroupHostsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharedDriveHostsModifyBulk** | [**SharedDriveHostsModifyBulk**](SharedDriveHostsModifyBulk.md) | The Drive Server Instance Group Hosts update object | 

### Return type

[**SharedDriveHosts**](SharedDriveHosts.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

