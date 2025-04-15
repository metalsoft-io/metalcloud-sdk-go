# \DriveAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrive**](DriveAPI.md#CreateDrive) | **Post** /api/v2/infrastructures/{infrastructureId}/drives | Create a new Drive
[**DeleteDrive**](DriveAPI.md#DeleteDrive) | **Delete** /api/v2/infrastructures/{infrastructureId}/drives/{driveId} | Deletes a Drive
[**GetDrive**](DriveAPI.md#GetDrive) | **Get** /api/v2/drives/{driveId} | Get Drive information
[**GetDriveConfigInfo**](DriveAPI.md#GetDriveConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/config | Get configuration information about the specified Drive
[**GetDriveHosts**](DriveAPI.md#GetDriveHosts) | **Get** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/hosts | Get the Hosts of Drive
[**GetInfrastructureDrive**](DriveAPI.md#GetInfrastructureDrive) | **Get** /api/v2/infrastructures/{infrastructureId}/drives/{driveId} | Get Drive information
[**GetInfrastructureDrives**](DriveAPI.md#GetInfrastructureDrives) | **Get** /api/v2/infrastructures/{infrastructureId}/drives | Get all Drives on the infrastructure
[**PatchDriveConfig**](DriveAPI.md#PatchDriveConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/config | Updates the config of a Drive
[**PatchDriveMeta**](DriveAPI.md#PatchDriveMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/drives/{driveId}/meta | Updates the meta of a Drive
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
	createSharedDrive := *openapiclient.NewCreateSharedDrive(float32(123)) // CreateSharedDrive | 

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.           <p>              <b>Format: </b> filter.subdomain={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomain=$not:$like:John Doe&filter.subdomain=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.           <p>              <b>Format: </b> filter.subdomainPermanent={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomainPermanent=$not:$like:John Doe&filter.subdomainPermanent=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.           <p>              <b>Format: </b> filter.storagePoolId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storagePoolId=$not:$like:John Doe&filter.storagePoolId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterWwn := []string{"Inner_example"} // []string | Filter by wwn query param.           <p>              <b>Format: </b> filter.wwn={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.wwn=$not:$like:John Doe&filter.wwn=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterQoS := []string{"Inner_example"} // []string | Filter by qoS query param.           <p>              <b>Format: </b> filter.qoS={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.qoS=$not:$like:John Doe&filter.qoS=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterLogicalNetworkId := []string{"Inner_example"} // []string | Filter by logicalNetworkId query param.           <p>              <b>Format: </b> filter.logicalNetworkId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.logicalNetworkId=$not:$like:John Doe&filter.logicalNetworkId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterAllocationAffinity := []string{"Inner_example"} // []string | Filter by allocationAffinity query param.           <p>              <b>Format: </b> filter.allocationAffinity={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.allocationAffinity=$not:$like:John Doe&filter.allocationAffinity=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterProvisioningProtocol := []string{"Inner_example"} // []string | Filter by provisioningProtocol query param.           <p>              <b>Format: </b> filter.provisioningProtocol={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.provisioningProtocol=$not:$like:John Doe&filter.provisioningProtocol=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigLogicalNetworkId := []string{"Inner_example"} // []string | Filter by config.logicalNetworkId query param.           <p>              <b>Format: </b> filter.config.logicalNetworkId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.logicalNetworkId=$not:$like:John Doe&filter.config.logicalNetworkId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>storagePoolId</li> <li>infrastructureId</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,infrastructureId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>infrastructureId</li> <li>storagePoolId</li> <li>serviceStatus</li> <li>wwn</li> <li>qoS</li> <li>logicalNetworkId</li> <li>allocationAffinity</li> <li>provisioningProtocol</li> <li>config.deployStatus</li> <li>config.deployType</li> <li>config.logicalNetworkId</li></ul>          (optional)

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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storagePoolId&#x3D;$not:$like:John Doe&amp;filter.storagePoolId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterWwn** | **[]string** | Filter by wwn query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.wwn&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.wwn&#x3D;$not:$like:John Doe&amp;filter.wwn&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterQoS** | **[]string** | Filter by qoS query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.qoS&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.qoS&#x3D;$not:$like:John Doe&amp;filter.qoS&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterLogicalNetworkId** | **[]string** | Filter by logicalNetworkId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.logicalNetworkId&#x3D;$not:$like:John Doe&amp;filter.logicalNetworkId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterAllocationAffinity** | **[]string** | Filter by allocationAffinity query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.allocationAffinity&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.allocationAffinity&#x3D;$not:$like:John Doe&amp;filter.allocationAffinity&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterProvisioningProtocol** | **[]string** | Filter by provisioningProtocol query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.provisioningProtocol&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.provisioningProtocol&#x3D;$not:$like:John Doe&amp;filter.provisioningProtocol&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigLogicalNetworkId** | **[]string** | Filter by config.logicalNetworkId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.logicalNetworkId&#x3D;$not:$like:John Doe&amp;filter.config.logicalNetworkId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,infrastructureId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;storagePoolId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;wwn&lt;/li&gt; &lt;li&gt;qoS&lt;/li&gt; &lt;li&gt;logicalNetworkId&lt;/li&gt; &lt;li&gt;allocationAffinity&lt;/li&gt; &lt;li&gt;provisioningProtocol&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt; &lt;li&gt;config.logicalNetworkId&lt;/li&gt;&lt;/ul&gt;          | 

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
	updateSharedDriveMeta := *openapiclient.NewUpdateSharedDriveMeta() // UpdateSharedDriveMeta | 

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

