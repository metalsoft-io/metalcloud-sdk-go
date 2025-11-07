# \ExtensionInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtensionInstance**](ExtensionInstanceAPI.md#CreateExtensionInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/extension-instances | Add extension instance to an infrastructure
[**DeleteExtensionInstance**](ExtensionInstanceAPI.md#DeleteExtensionInstance) | **Delete** /api/v2/extension-instances/{extensionInstanceId} | Delete extension instance
[**GetExtensionInstance**](ExtensionInstanceAPI.md#GetExtensionInstance) | **Get** /api/v2/extension-instances/{extensionInstanceId} | Get extension instance details
[**GetExtensionInstanceCredentials**](ExtensionInstanceAPI.md#GetExtensionInstanceCredentials) | **Get** /api/v2/extension-instances/{extensionInstanceId}/credentials | Get Extension Instance credentials
[**GetExtensionInstances**](ExtensionInstanceAPI.md#GetExtensionInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/extension-instances | Get extension instances list
[**UpdateExtensionInstance**](ExtensionInstanceAPI.md#UpdateExtensionInstance) | **Patch** /api/v2/extension-instances/{extensionInstanceId}/config | Update extension instance configuration



## CreateExtensionInstance

> ExtensionInstance CreateExtensionInstance(ctx, infrastructureId).CreateExtensionInstance(createExtensionInstance).Execute()

Add extension instance to an infrastructure



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
	createExtensionInstance := *openapiclient.NewCreateExtensionInstance() // CreateExtensionInstance | The extension instance to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionInstanceAPI.CreateExtensionInstance(context.Background(), infrastructureId).CreateExtensionInstance(createExtensionInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionInstanceAPI.CreateExtensionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExtensionInstance`: ExtensionInstance
	fmt.Fprintf(os.Stdout, "Response from `ExtensionInstanceAPI.CreateExtensionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExtensionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createExtensionInstance** | [**CreateExtensionInstance**](CreateExtensionInstance.md) | The extension instance to create | 

### Return type

[**ExtensionInstance**](ExtensionInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtensionInstance

> DeleteExtensionInstance(ctx, extensionInstanceId).IfMatch(ifMatch).Execute()

Delete extension instance



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
	extensionInstanceId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionInstanceAPI.DeleteExtensionInstance(context.Background(), extensionInstanceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionInstanceAPI.DeleteExtensionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionInstanceRequest struct via the builder pattern


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


## GetExtensionInstance

> ExtensionInstance GetExtensionInstance(ctx, extensionInstanceId).Execute()

Get extension instance details



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
	extensionInstanceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionInstanceAPI.GetExtensionInstance(context.Background(), extensionInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionInstanceAPI.GetExtensionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionInstance`: ExtensionInstance
	fmt.Fprintf(os.Stdout, "Response from `ExtensionInstanceAPI.GetExtensionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionInstance**](ExtensionInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionInstanceCredentials

> ExtensionInstanceCredentials GetExtensionInstanceCredentials(ctx, extensionInstanceId).Execute()

Get Extension Instance credentials



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
	extensionInstanceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionInstanceAPI.GetExtensionInstanceCredentials(context.Background(), extensionInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionInstanceAPI.GetExtensionInstanceCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionInstanceCredentials`: ExtensionInstanceCredentials
	fmt.Fprintf(os.Stdout, "Response from `ExtensionInstanceAPI.GetExtensionInstanceCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionInstanceCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionInstanceCredentials**](ExtensionInstanceCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionInstances

> ExtensionInstancePaginatedList GetExtensionInstances(ctx, infrastructureId).Page(page).Limit(limit).FilterExtensionId(filterExtensionId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get extension instances list



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
	filterExtensionId := []string{"Inner_example"} // []string | Filter by extensionId query param.  **Format:** filter.extensionId={$not}:OPERATION:VALUE    **Example:** filter.extensionId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - extensionId  - infrastructureId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionInstanceAPI.GetExtensionInstances(context.Background(), infrastructureId).Page(page).Limit(limit).FilterExtensionId(filterExtensionId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionInstanceAPI.GetExtensionInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionInstances`: ExtensionInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ExtensionInstanceAPI.GetExtensionInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterExtensionId** | **[]string** | Filter by extensionId query param.  **Format:** filter.extensionId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.extensionId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - extensionId  - infrastructureId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  | 

### Return type

[**ExtensionInstancePaginatedList**](ExtensionInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtensionInstance

> ExtensionInstance UpdateExtensionInstance(ctx, extensionInstanceId).UpdateExtensionInstance(updateExtensionInstance).IfMatch(ifMatch).Execute()

Update extension instance configuration



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
	extensionInstanceId := float32(8.14) // float32 | 
	updateExtensionInstance := *openapiclient.NewUpdateExtensionInstance() // UpdateExtensionInstance | The extension instance changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionInstanceAPI.UpdateExtensionInstance(context.Background(), extensionInstanceId).UpdateExtensionInstance(updateExtensionInstance).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionInstanceAPI.UpdateExtensionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtensionInstance`: ExtensionInstance
	fmt.Fprintf(os.Stdout, "Response from `ExtensionInstanceAPI.UpdateExtensionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExtensionInstance** | [**UpdateExtensionInstance**](UpdateExtensionInstance.md) | The extension instance changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ExtensionInstance**](ExtensionInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

