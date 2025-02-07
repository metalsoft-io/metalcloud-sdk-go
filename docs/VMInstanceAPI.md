# \VMInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyVMTypeOnVMInstance**](VMInstanceAPI.md#ApplyVMTypeOnVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance
[**CreateVMInstance**](VMInstanceAPI.md#CreateVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances | Creates a VM Instance
[**DeleteVMInstance**](VMInstanceAPI.md#DeleteVMInstance) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Deletes a VM Instance
[**GetVMInstance**](VMInstanceAPI.md#GetVMInstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Get VM Instance information
[**GetVMInstancePowerStatus**](VMInstanceAPI.md#GetVMInstancePowerStatus) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/power-status | Retrieves the power status of the VM Instance
[**RebootVMInstance**](VMInstanceAPI.md#RebootVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/reboot | Reboots the VM Instance
[**ShutdownVMInstance**](VMInstanceAPI.md#ShutdownVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/shutdown | Shuts down the VM Instance
[**StartVMInstance**](VMInstanceAPI.md#StartVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/start | Starts the VM Instance
[**UpdateVMInstance**](VMInstanceAPI.md#UpdateVMInstance) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Updates VM Instance information



## ApplyVMTypeOnVMInstance

> VMInstance ApplyVMTypeOnVMInstance(ctx, infrastructureId, vmInstanceId, vmTypeId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.ApplyVMTypeOnVMInstance(context.Background(), infrastructureId, vmInstanceId, vmTypeId).Execute()
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

> DeleteVMInstance(ctx, infrastructureId, vmInstanceId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceAPI.DeleteVMInstance(context.Background(), infrastructureId, vmInstanceId).Execute()
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


## GetVMInstance

> VMInstance GetVMInstance(ctx, infrastructureId, vmInstanceId).Execute()

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
	resp, r, err := apiClient.VMInstanceAPI.GetVMInstance(context.Background(), infrastructureId, vmInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.GetVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstance`: VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.GetVMInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceRequest struct via the builder pattern


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


## UpdateVMInstance

> VMInstance UpdateVMInstance(ctx, infrastructureId, vmInstanceId).UpdateVMInstance(updateVMInstance).Execute()

Updates VM Instance information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceAPI.UpdateVMInstance(context.Background(), infrastructureId, vmInstanceId).UpdateVMInstance(updateVMInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceAPI.UpdateVMInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMInstance`: VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceAPI.UpdateVMInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMInstance** | [**UpdateVMInstance**](UpdateVMInstance.md) | The VM Instance update object | 

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

