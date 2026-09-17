# \ContainerInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyContainerTypeOnContainerInstance**](ContainerInstanceAPI.md#ApplyContainerTypeOnContainerInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/actions/apply-type/{containerTypeId} | Applies a Container Type to a Container Instance
[**CreateContainerInstance**](ContainerInstanceAPI.md#CreateContainerInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instances | Creates a Container Instance
[**DeleteContainerInstance**](ContainerInstanceAPI.md#DeleteContainerInstance) | **Delete** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId} | Deletes a Container Instance
[**GetContainerInstanceConfigInfo**](ContainerInstanceAPI.md#GetContainerInstanceConfigInfo) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/config | Get configuration information about the specified Container Instance
[**GetContainerInstanceCredentials**](ContainerInstanceAPI.md#GetContainerInstanceCredentials) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/credentials | Get Container Instance credentials
[**GetContainerInstanceOSInstallationData**](ContainerInstanceAPI.md#GetContainerInstanceOSInstallationData) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/os-installation-data | Get Container instance OS installation data
[**GetContainerInstancePowerStatus**](ContainerInstanceAPI.md#GetContainerInstancePowerStatus) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/power-status | Retrieves the power status of the Container Instance
[**GetContainerInstanceVariables**](ContainerInstanceAPI.md#GetContainerInstanceVariables) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/variables | Get Container instance variables
[**GetInfrastructureContainerInstance**](ContainerInstanceAPI.md#GetInfrastructureContainerInstance) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId} | Get Container Instance information
[**GetInfrastructureContainerInstances**](ContainerInstanceAPI.md#GetInfrastructureContainerInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/container-instances | Get all Container Instances on the infrastructure
[**PatchContainerInstanceMeta**](ContainerInstanceAPI.md#PatchContainerInstanceMeta) | **Patch** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/meta | Updates the meta of a Container Instance
[**RebootContainerInstance**](ContainerInstanceAPI.md#RebootContainerInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/reboot | Reboots the Container Instance
[**ShutdownContainerInstance**](ContainerInstanceAPI.md#ShutdownContainerInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/shutdown | Shuts down the Container Instance
[**StartContainerInstance**](ContainerInstanceAPI.md#StartContainerInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/start | Starts the Container Instance
[**UpdateContainerInstanceConfig**](ContainerInstanceAPI.md#UpdateContainerInstanceConfig) | **Patch** /api/v2/infrastructures/{infrastructureId}/container-instances/{containerInstanceId}/config | Updates Container Instance config information



## ApplyContainerTypeOnContainerInstance

> ContainerInstance ApplyContainerTypeOnContainerInstance(ctx, infrastructureId, containerInstanceId, containerTypeId).IfMatch(ifMatch).Execute()

Applies a Container Type to a Container Instance



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 
	containerTypeId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.ApplyContainerTypeOnContainerInstance(context.Background(), infrastructureId, containerInstanceId, containerTypeId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.ApplyContainerTypeOnContainerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyContainerTypeOnContainerInstance`: ContainerInstance
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.ApplyContainerTypeOnContainerInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 
**containerTypeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyContainerTypeOnContainerInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **string** | Entity tag | 

### Return type

[**ContainerInstance**](ContainerInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerInstance

> ContainerInstance CreateContainerInstance(ctx, infrastructureId).CreateContainerInstance(createContainerInstance).Execute()

Creates a Container Instance



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
	infrastructureId := int64(789) // int64 | 
	createContainerInstance := *openapiclient.NewCreateContainerInstance(int64(123), int64(123)) // CreateContainerInstance | The Container Instance create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.CreateContainerInstance(context.Background(), infrastructureId).CreateContainerInstance(createContainerInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.CreateContainerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerInstance`: ContainerInstance
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.CreateContainerInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContainerInstance** | [**CreateContainerInstance**](CreateContainerInstance.md) | The Container Instance create object | 

### Return type

[**ContainerInstance**](ContainerInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerInstance

> DeleteContainerInstance(ctx, infrastructureId, containerInstanceId).IfMatch(ifMatch).Execute()

Deletes a Container Instance



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerInstanceAPI.DeleteContainerInstance(context.Background(), infrastructureId, containerInstanceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.DeleteContainerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerInstanceRequest struct via the builder pattern


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


## GetContainerInstanceConfigInfo

> ContainerInstanceConfiguration GetContainerInstanceConfigInfo(ctx, infrastructureId, containerInstanceId).Execute()

Get configuration information about the specified Container Instance

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.GetContainerInstanceConfigInfo(context.Background(), infrastructureId, containerInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.GetContainerInstanceConfigInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceConfigInfo`: ContainerInstanceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.GetContainerInstanceConfigInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceConfigInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerInstanceConfiguration**](ContainerInstanceConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstanceCredentials

> ServerInstanceCredentials GetContainerInstanceCredentials(ctx, infrastructureId, containerInstanceId).Execute()

Get Container Instance credentials



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.GetContainerInstanceCredentials(context.Background(), infrastructureId, containerInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.GetContainerInstanceCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceCredentials`: ServerInstanceCredentials
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.GetContainerInstanceCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerInstanceCredentials**](ServerInstanceCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstanceOSInstallationData

> ContainerInstanceContextOSInstallationData GetContainerInstanceOSInstallationData(ctx, infrastructureId, containerInstanceId).Usage(usage).RemoveEmpty(removeEmpty).Execute()

Get Container instance OS installation data

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 
	usage := openapiclient.VariableUsageType("HTTPRequest") // VariableUsageType | Filter by variable usage (optional)
	removeEmpty := int32(56) // int32 | Remove empty fields from the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.GetContainerInstanceOSInstallationData(context.Background(), infrastructureId, containerInstanceId).Usage(usage).RemoveEmpty(removeEmpty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.GetContainerInstanceOSInstallationData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceOSInstallationData`: ContainerInstanceContextOSInstallationData
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.GetContainerInstanceOSInstallationData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceOSInstallationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **usage** | [**VariableUsageType**](VariableUsageType.md) | Filter by variable usage | 
 **removeEmpty** | **int32** | Remove empty fields from the response | 

### Return type

[**ContainerInstanceContextOSInstallationData**](ContainerInstanceContextOSInstallationData.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerInstancePowerStatus

> string GetContainerInstancePowerStatus(ctx, infrastructureId, containerInstanceId).Execute()

Retrieves the power status of the Container Instance



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.GetContainerInstancePowerStatus(context.Background(), infrastructureId, containerInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.GetContainerInstancePowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstancePowerStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.GetContainerInstancePowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstancePowerStatusRequest struct via the builder pattern


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


## GetContainerInstanceVariables

> ContainerInstanceContextVariables GetContainerInstanceVariables(ctx, infrastructureId, containerInstanceId).Usage(usage).Execute()

Get Container instance variables

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 
	usage := openapiclient.VariableUsageType("HTTPRequest") // VariableUsageType | Filter by variable usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.GetContainerInstanceVariables(context.Background(), infrastructureId, containerInstanceId).Usage(usage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.GetContainerInstanceVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerInstanceVariables`: ContainerInstanceContextVariables
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.GetContainerInstanceVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInstanceVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **usage** | [**VariableUsageType**](VariableUsageType.md) | Filter by variable usage | 

### Return type

[**ContainerInstanceContextVariables**](ContainerInstanceContextVariables.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureContainerInstance

> ContainerInstance GetInfrastructureContainerInstance(ctx, infrastructureId, containerInstanceId).Execute()

Get Container Instance information



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.GetInfrastructureContainerInstance(context.Background(), infrastructureId, containerInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.GetInfrastructureContainerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureContainerInstance`: ContainerInstance
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.GetInfrastructureContainerInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureContainerInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerInstance**](ContainerInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureContainerInstances

> ContainerInstancePaginatedList GetInfrastructureContainerInstances(ctx, infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterVmId(filterVmId).FilterConfigVmId(filterConfigVmId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Container Instances on the infrastructure



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
	infrastructureId := int64(789) // int64 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomain := []string{"Inner_example"} // []string | Filter by subdomain query param.  **Format:** filter.subdomain={$not}:OPERATION:VALUE    **Example:** filter.subdomain=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubdomainPermanent := []string{"Inner_example"} // []string | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent={$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterVmId := []string{"Inner_example"} // []string | Filter by vmId query param.  **Format:** filter.vmId={$not}:OPERATION:VALUE    **Example:** filter.vmId=$eq:John Doe&filter.vmId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigVmId := []string{"Inner_example"} // []string | Filter by config.vmId query param.  **Format:** filter.config.vmId={$not}:OPERATION:VALUE    **Example:** filter.config.vmId=$eq:John Doe&filter.config.vmId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=serviceStatus:DESC   **Default Value:** id:DESC  **Available Fields** - id  - serviceStatus  - config.deployStatus  - config.deployType  - containerId  - config.containerId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,serviceStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - serviceStatus  - config.deployStatus  - config.deployType  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.GetInfrastructureContainerInstances(context.Background(), infrastructureId).Page(page).Limit(limit).FilterLabel(filterLabel).FilterSubdomain(filterSubdomain).FilterSubdomainPermanent(filterSubdomainPermanent).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterVmId(filterVmId).FilterConfigVmId(filterConfigVmId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.GetInfrastructureContainerInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfrastructureContainerInstances`: ContainerInstancePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.GetInfrastructureContainerInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureContainerInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomain** | **[]string** | Filter by subdomain query param.  **Format:** filter.subdomain&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomain&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubdomainPermanent** | **[]string** | Filter by subdomainPermanent query param.  **Format:** filter.subdomainPermanent&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subdomainPermanent&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterVmId** | **[]string** | Filter by vmId query param.  **Format:** filter.vmId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vmId&#x3D;$eq:John Doe&amp;filter.vmId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigVmId** | **[]string** | Filter by config.vmId query param.  **Format:** filter.config.vmId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.vmId&#x3D;$eq:John Doe&amp;filter.config.vmId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;serviceStatus:DESC   **Default Value:** id:DESC  **Available Fields** - id  - serviceStatus  - config.deployStatus  - config.deployType  - containerId  - config.containerId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,subdomain,subdomainPermanent,serviceStatus   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - subdomain  - subdomainPermanent  - serviceStatus  - config.deployStatus  - config.deployType  | 

### Return type

[**ContainerInstancePaginatedList**](ContainerInstancePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchContainerInstanceMeta

> ContainerInstance PatchContainerInstanceMeta(ctx, infrastructureId, containerInstanceId).UpdateContainerInstanceMeta(updateContainerInstanceMeta).Execute()

Updates the meta of a Container Instance

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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 
	updateContainerInstanceMeta := *openapiclient.NewUpdateContainerInstanceMeta() // UpdateContainerInstanceMeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.PatchContainerInstanceMeta(context.Background(), infrastructureId, containerInstanceId).UpdateContainerInstanceMeta(updateContainerInstanceMeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.PatchContainerInstanceMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContainerInstanceMeta`: ContainerInstance
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.PatchContainerInstanceMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContainerInstanceMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateContainerInstanceMeta** | [**UpdateContainerInstanceMeta**](UpdateContainerInstanceMeta.md) |  | 

### Return type

[**ContainerInstance**](ContainerInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootContainerInstance

> RebootContainerInstance(ctx, infrastructureId, containerInstanceId).Execute()

Reboots the Container Instance



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerInstanceAPI.RebootContainerInstance(context.Background(), infrastructureId, containerInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.RebootContainerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootContainerInstanceRequest struct via the builder pattern


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


## ShutdownContainerInstance

> ShutdownContainerInstance(ctx, infrastructureId, containerInstanceId).Execute()

Shuts down the Container Instance



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerInstanceAPI.ShutdownContainerInstance(context.Background(), infrastructureId, containerInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.ShutdownContainerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownContainerInstanceRequest struct via the builder pattern


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


## StartContainerInstance

> StartContainerInstance(ctx, infrastructureId, containerInstanceId).Execute()

Starts the Container Instance



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerInstanceAPI.StartContainerInstance(context.Background(), infrastructureId, containerInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.StartContainerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartContainerInstanceRequest struct via the builder pattern


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


## UpdateContainerInstanceConfig

> ContainerInstanceGroupConfiguration UpdateContainerInstanceConfig(ctx, infrastructureId, containerInstanceId).UpdateContainerInstance(updateContainerInstance).IfMatch(ifMatch).Execute()

Updates Container Instance config information



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
	infrastructureId := int64(789) // int64 | 
	containerInstanceId := int64(789) // int64 | 
	updateContainerInstance := *openapiclient.NewUpdateContainerInstance() // UpdateContainerInstance | The Container Instance update object
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerInstanceAPI.UpdateContainerInstanceConfig(context.Background(), infrastructureId, containerInstanceId).UpdateContainerInstance(updateContainerInstance).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerInstanceAPI.UpdateContainerInstanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContainerInstanceConfig`: ContainerInstanceGroupConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ContainerInstanceAPI.UpdateContainerInstanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **int64** |  | 
**containerInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerInstanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateContainerInstance** | [**UpdateContainerInstance**](UpdateContainerInstance.md) | The Container Instance update object | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ContainerInstanceGroupConfiguration**](ContainerInstanceGroupConfiguration.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

