# \DriveGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDriveGroup**](DriveGroupAPI.md#CreateDriveGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/drive-groups | Create a new Drive Group
[**DeleteDriveGroup**](DriveGroupAPI.md#DeleteDriveGroup) | **Delete** /api/v2/infrastructures/{infrastructureId}/drive-groups/{driveGroupId} | Deletes a Drive Group
[**GetDriveGroupConfigInfo**](DriveGroupAPI.md#GetDriveGroupConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/drive-groups/{driveGroupId}/config | Get configuration information about the specified Drive Group
[**GetDriveGroupDrives**](DriveGroupAPI.md#GetDriveGroupDrives) | **Get** /api/v2/{infrastructureId}/drive-groups/{driveGroupId}/drives | Get the Drives of the Drive Group
[**GetInfrastructureDriveGroup**](DriveGroupAPI.md#GetInfrastructureDriveGroup) | **Get** /api/v2/infrastructures/{infrastructureId}/drive-groups/{driveGroupId} | Get DriveGroup information
[**GetInfrastructureDriveGroups**](DriveGroupAPI.md#GetInfrastructureDriveGroups) | **Get** /api/v2/infrastructures/{infrastructureId}/drive-groups | Get all Drive Groups on the infrastructure
[**PatchDriveGroupConfig**](DriveGroupAPI.md#PatchDriveGroupConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/drive-groups/{driveGroupId}/config | Updates the config of a Drive Group
[**PatchDriveGroupMeta**](DriveGroupAPI.md#PatchDriveGroupMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/drive-groups/{driveGroupId}/meta | Updates the meta of a Drive Group



## CreateDriveGroup

> DriveGroup CreateDriveGroup(ctx, infrastructureId).CreateGroupDrive(createGroupDrive).Execute()

Create a new Drive Group

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
	createGroupDrive := *openapiclient.NewCreateGroupDrive(float32(1), float32(123)) // CreateGroupDrive | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveGroupAPI.CreateDriveGroup(context.Background(), infrastructureId).CreateGroupDrive(createGroupDrive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.CreateDriveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDriveGroup`: DriveGroup
	fmt.Fprintf(os.Stdout, "Response from `DriveGroupAPI.CreateDriveGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupDrive** | [**CreateGroupDrive**](CreateGroupDrive.md) |  | 

### Return type

[**DriveGroup**](DriveGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDriveGroup

> DeleteDriveGroup(ctx, infrastructureId, driveGroupId).IfMatch(ifMatch).Execute()

Deletes a Drive Group

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
	driveGroupId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DriveGroupAPI.DeleteDriveGroup(context.Background(), infrastructureId, driveGroupId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.DeleteDriveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriveGroupRequest struct via the builder pattern


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


## GetDriveGroupConfigInfo

> DriveGroupConfiguration GetDriveGroupConfigInfo(ctx, infrastructureId, driveGroupId).Execute()

Get configuration information about the specified Drive Group

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
	driveGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveGroupAPI.GetDriveGroupConfigInfo(context.Background(), infrastructureId, driveGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.GetDriveGroupConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveGroupConfigInfo`: DriveGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `DriveGroupAPI.GetDriveGroupConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveGroupConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DriveGroupConfiguration**](DriveGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDriveGroupDrives

> DrivePaginatedList GetDriveGroupDrives(ctx, infrastructureId, driveGroupId).Execute()

Get the Drives of the Drive Group



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
	driveGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveGroupAPI.GetDriveGroupDrives(context.Background(), infrastructureId, driveGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.GetDriveGroupDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveGroupDrives`: DrivePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DriveGroupAPI.GetDriveGroupDrives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveGroupDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetInfrastructureDriveGroup

> DriveGroup GetInfrastructureDriveGroup(ctx, infrastructureId, driveGroupId).Execute()

Get DriveGroup information



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
	driveGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveGroupAPI.GetInfrastructureDriveGroup(context.Background(), infrastructureId, driveGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.GetInfrastructureDriveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureDriveGroup`: DriveGroup
	fmt.Fprintf(os.Stdout, "Response from `DriveGroupAPI.GetInfrastructureDriveGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureDriveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DriveGroup**](DriveGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureDriveGroups

> DriveGroupPaginatedList GetInfrastructureDriveGroups(ctx, infrastructureId).Page(page).Limit(limit).FilterId(filterId).FilterServerInstanceGroupId(filterServerInstanceGroupId).FilterTemplateId(filterTemplateId).FilterServiceStatus(filterServiceStatus).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterExtensionInstanceId(filterExtensionInstanceId).FilterTags(filterTags).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Drive Groups on the infrastructure



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServerInstanceGroupId := []string{"Inner_example"} // []string | Filter by serverInstanceGroupId query param.           <p>              <b>Format: </b> filter.serverInstanceGroupId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverInstanceGroupId=$not:$like:John Doe&filter.serverInstanceGroupId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterTemplateId := []string{"Inner_example"} // []string | Filter by templateId query param.           <p>              <b>Format: </b> filter.templateId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.templateId=$not:$like:John Doe&filter.templateId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.           <p>              <b>Format: </b> filter.subdomain={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomain=$not:$like:John Doe&filter.subdomain=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.           <p>              <b>Format: </b> filter.subdomainPermanent={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.subdomainPermanent=$not:$like:John Doe&filter.subdomainPermanent=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterExtensionInstanceId := []string{"Inner_example"} // []string | Filter by extensionInstanceId query param.           <p>              <b>Format: </b> filter.extensionInstanceId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.extensionInstanceId=$not:$like:John Doe&filter.extensionInstanceId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.           <p>              <b>Format: </b> filter.tags={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.tags=$not:$like:John Doe&filter.tags=like:John           </p>           <h4>Available Operations</h4><ul><li>$contains</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,infrastructureId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>infrastructureId</li> <li>serviceStatus</li> <li>templateId</li> <li>serverInstanceGroupId</li> <li>extensionInstanceId</li> <li>config.deployStatus</li> <li>config.deployType</li> <li>config.templateId</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveGroupAPI.GetInfrastructureDriveGroups(context.Background(), infrastructureId).Page(page).Limit(limit).FilterId(filterId).FilterServerInstanceGroupId(filterServerInstanceGroupId).FilterTemplateId(filterTemplateId).FilterServiceStatus(filterServiceStatus).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterExtensionInstanceId(filterExtensionInstanceId).FilterTags(filterTags).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.GetInfrastructureDriveGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureDriveGroups`: DriveGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DriveGroupAPI.GetInfrastructureDriveGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureDriveGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerInstanceGroupId** | **[]string** | Filter by serverInstanceGroupId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverInstanceGroupId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverInstanceGroupId&#x3D;$not:$like:John Doe&amp;filter.serverInstanceGroupId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterTemplateId** | **[]string** | Filter by templateId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.templateId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.templateId&#x3D;$not:$like:John Doe&amp;filter.templateId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterExtensionInstanceId** | **[]string** | Filter by extensionInstanceId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.extensionInstanceId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.extensionInstanceId&#x3D;$not:$like:John Doe&amp;filter.extensionInstanceId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterTags** | **[]string** | Filter by tags query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.tags&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.tags&#x3D;$not:$like:John Doe&amp;filter.tags&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,infrastructureId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;templateId&lt;/li&gt; &lt;li&gt;serverInstanceGroupId&lt;/li&gt; &lt;li&gt;extensionInstanceId&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt; &lt;li&gt;config.templateId&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**DriveGroupPaginatedList**](DriveGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDriveGroupConfig

> DriveGroup PatchDriveGroupConfig(ctx, infrastructureId, driveGroupId).UpdateGroupDrive(updateGroupDrive).IfMatch(ifMatch).Execute()

Updates the config of a Drive Group

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
	driveGroupId := float32(8.14) // float32 | 
	updateGroupDrive := *openapiclient.NewUpdateGroupDrive() // UpdateGroupDrive | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveGroupAPI.PatchDriveGroupConfig(context.Background(), infrastructureId, driveGroupId).UpdateGroupDrive(updateGroupDrive).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.PatchDriveGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDriveGroupConfig`: DriveGroup
	fmt.Fprintf(os.Stdout, "Response from `DriveGroupAPI.PatchDriveGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDriveGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupDrive** | [**UpdateGroupDrive**](UpdateGroupDrive.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**DriveGroup**](DriveGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDriveGroupMeta

> DriveGroup PatchDriveGroupMeta(ctx, infrastructureId, driveGroupId).UpdateDriveGroupMeta(updateDriveGroupMeta).Execute()

Updates the meta of a Drive Group

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
	driveGroupId := float32(8.14) // float32 | 
	updateDriveGroupMeta := *openapiclient.NewUpdateDriveGroupMeta() // UpdateDriveGroupMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveGroupAPI.PatchDriveGroupMeta(context.Background(), infrastructureId, driveGroupId).UpdateDriveGroupMeta(updateDriveGroupMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveGroupAPI.PatchDriveGroupMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDriveGroupMeta`: DriveGroup
	fmt.Fprintf(os.Stdout, "Response from `DriveGroupAPI.PatchDriveGroupMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**driveGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDriveGroupMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDriveGroupMeta** | [**UpdateDriveGroupMeta**](UpdateDriveGroupMeta.md) |  | 

### Return type

[**DriveGroup**](DriveGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

