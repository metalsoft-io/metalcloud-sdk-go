# \VMInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyVMTypeOnVMInstance**](VMInstanceAPI.md#ApplyVMTypeOnVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance
[**CreateVMInstance**](VMInstanceAPI.md#CreateVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances | Creates a VM Instance
[**DeleteVMInstance**](VMInstanceAPI.md#DeleteVMInstance) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Deletes a VM Instance
[**GetInfrastructureVMInstance**](VMInstanceAPI.md#GetInfrastructureVMInstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Get VM Instance information
[**GetInfrastructureVMInstances**](VMInstanceAPI.md#GetInfrastructureVMInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances | Get all VM Instances on the infrastructure
[**GetVMInstanceConfigInfo**](VMInstanceAPI.md#GetVMInstanceConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/config | Get configuration information about the specified VM Instance
[**GetVMInstancePowerStatus**](VMInstanceAPI.md#GetVMInstancePowerStatus) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/power-status | Retrieves the power status of the VM Instance
[**PatchVMInstanceMeta**](VMInstanceAPI.md#PatchVMInstanceMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/meta | Updates the meta of a VM Instance
[**RebootVMInstance**](VMInstanceAPI.md#RebootVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/reboot | Reboots the VM Instance
[**ShutdownVMInstance**](VMInstanceAPI.md#ShutdownVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/shutdown | Shuts down the VM Instance
[**StartVMInstance**](VMInstanceAPI.md#StartVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/start | Starts the VM Instance
[**UpdateVMInstanceConfig**](VMInstanceAPI.md#UpdateVMInstanceConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/config | Updates VM Instance config information



## ApplyVMTypeOnVMInstance

> VMInstance ApplyVMTypeOnVMInstance(ctx, infrastructureId, vmInstanceId, vmTypeId).IfMatch(ifMatch).Execute()

Applies a VM Type to a VM Instance



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
	vmInstanceId := float32(8.14) // float32 | 
	vmTypeId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.ApplyVMTypeOnVMInstance(context.Background(), infrastructureId, vmInstanceId, vmTypeId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.ApplyVMTypeOnVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyVMTypeOnVMInstance`: VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.ApplyVMTypeOnVMInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 
**vmTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyVMTypeOnVMInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **string** | Entity tag | 

### Return type

[**VMInstance**](VMInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVMInstance

> VMInstance CreateVMInstance(ctx, infrastructureId).CreateVMInstance(createVMInstance).Execute()

Creates a VM Instance



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
	createVMInstance := *openapiclient.NewCreateVMInstance(float32(123), float32(123)) // CreateVMInstance | The VM Instance create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.CreateVMInstance(context.Background(), infrastructureId).CreateVMInstance(createVMInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.CreateVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVMInstance`: VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.CreateVMInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVMInstance** | [**CreateVMInstance**](CreateVMInstance.md) | The VM Instance create object | 

### Return type

[**VMInstance**](VMInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVMInstance

> DeleteVMInstance(ctx, infrastructureId, vmInstanceId).IfMatch(ifMatch).Execute()

Deletes a VM Instance



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
	vmInstanceId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceAPI.DeleteVMInstance(context.Background(), infrastructureId, vmInstanceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.DeleteVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMInstanceRequest struct via the builder pattern


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


## GetInfrastructureVMInstance

> VMInstance GetInfrastructureVMInstance(ctx, infrastructureId, vmInstanceId).Execute()

Get VM Instance information



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
	vmInstanceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.GetInfrastructureVMInstance(context.Background(), infrastructureId, vmInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.GetInfrastructureVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureVMInstance`: VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.GetInfrastructureVMInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureVMInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMInstance**](VMInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureVMInstances

> VMInstancePaginatedList GetInfrastructureVMInstances(ctx, infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all VM Instances on the infrastructure



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
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,subdomain,subdomainPermanent,serviceStatus           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>subdomain</li> <li>subdomainPermanent</li> <li>serviceStatus</li> <li>config.deployStatus</li> <li>config.deployType</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.GetInfrastructureVMInstances(context.Background(), infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.GetInfrastructureVMInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureVMInstances`: VMInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.GetInfrastructureVMInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureVMInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomain&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomain&#x3D;$not:$like:John Doe&amp;filter.subdomain&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.subdomainPermanent&#x3D;$not:$like:John Doe&amp;filter.subdomainPermanent&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,subdomain,subdomainPermanent,serviceStatus           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;subdomain&lt;/li&gt; &lt;li&gt;subdomainPermanent&lt;/li&gt; &lt;li&gt;serviceStatus&lt;/li&gt; &lt;li&gt;config.deployStatus&lt;/li&gt; &lt;li&gt;config.deployType&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**VMInstancePaginatedList**](VMInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceConfigInfo

> VMInstanceConfiguration GetVMInstanceConfigInfo(ctx, infrastructureId, vmInstanceId).Execute()

Get configuration information about the specified VM Instance

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
	vmInstanceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.GetVMInstanceConfigInfo(context.Background(), infrastructureId, vmInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.GetVMInstanceConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceConfigInfo`: VMInstanceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.GetVMInstanceConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMInstanceConfiguration**](VMInstanceConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstancePowerStatus

> string GetVMInstancePowerStatus(ctx, infrastructureId, vmInstanceId).Execute()

Retrieves the power status of the VM Instance



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
	vmInstanceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.GetVMInstancePowerStatus(context.Background(), infrastructureId, vmInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.GetVMInstancePowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstancePowerStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.GetVMInstancePowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstancePowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchVMInstanceMeta

> VMInstance PatchVMInstanceMeta(ctx, infrastructureId, vmInstanceId).UpdateVMInstanceMeta(updateVMInstanceMeta).Execute()

Updates the meta of a VM Instance

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
	vmInstanceId := float32(8.14) // float32 | 
	updateVMInstanceMeta := *openapiclient.NewUpdateVMInstanceMeta() // UpdateVMInstanceMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.PatchVMInstanceMeta(context.Background(), infrastructureId, vmInstanceId).UpdateVMInstanceMeta(updateVMInstanceMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.PatchVMInstanceMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchVMInstanceMeta`: VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.PatchVMInstanceMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchVMInstanceMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMInstanceMeta** | [**UpdateVMInstanceMeta**](UpdateVMInstanceMeta.md) |  | 

### Return type

[**VMInstance**](VMInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVMInstance

> RebootVMInstance(ctx, infrastructureId, vmInstanceId).Execute()

Reboots the VM Instance



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
	vmInstanceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceAPI.RebootVMInstance(context.Background(), infrastructureId, vmInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.RebootVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootVMInstanceRequest struct via the builder pattern


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


## ShutdownVMInstance

> ShutdownVMInstance(ctx, infrastructureId, vmInstanceId).Execute()

Shuts down the VM Instance



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
	vmInstanceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceAPI.ShutdownVMInstance(context.Background(), infrastructureId, vmInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.ShutdownVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownVMInstanceRequest struct via the builder pattern


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


## StartVMInstance

> StartVMInstance(ctx, infrastructureId, vmInstanceId).Execute()

Starts the VM Instance



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
	vmInstanceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceAPI.StartVMInstance(context.Background(), infrastructureId, vmInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.StartVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVMInstanceRequest struct via the builder pattern


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


## UpdateVMInstanceConfig

> VMInstance UpdateVMInstanceConfig(ctx, infrastructureId, vmInstanceId).UpdateVMInstance(updateVMInstance).IfMatch(ifMatch).Execute()

Updates VM Instance config information



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
	vmInstanceId := float32(8.14) // float32 | 
	updateVMInstance := *openapiclient.NewUpdateVMInstance() // UpdateVMInstance | The VM Instance update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.UpdateVMInstanceConfig(context.Background(), infrastructureId, vmInstanceId).UpdateVMInstance(updateVMInstance).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.UpdateVMInstanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMInstanceConfig`: VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.UpdateVMInstanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMInstanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMInstance** | [**UpdateVMInstance**](UpdateVMInstance.md) | The VM Instance update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**VMInstance**](VMInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

