# \VMInstanceGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyVMTypeOnVMInstanceGroup**](VMInstanceGroupAPI.md#ApplyVMTypeOnVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance Group
[**CreateVMInstanceGroup**](VMInstanceGroupAPI.md#CreateVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Creates a VM Instance Group
[**CreateVMInterfaceOnVMInstanceGroup**](VMInstanceGroupAPI.md#CreateVMInterfaceOnVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/interfaces | Creates a new Virtual Interface for the VM Instance Group
[**DeleteVMInstanceGroup**](VMInstanceGroupAPI.md#DeleteVMInstanceGroup) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Deletes a VM Instance Group
[**GetVMInstanceGroup**](VMInstanceGroupAPI.md#GetVMInstanceGroup) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Get VM Instance Group information
[**GetVMInstanceGroupVMInstances**](VMInstanceGroupAPI.md#GetVMInstanceGroupVMInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/vm-instances | Get the VM Instances of VM Instance Group
[**GetVMInstanceGroups**](VMInstanceGroupAPI.md#GetVMInstanceGroups) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Get all VM Instance Groups
[**UpdateNetworkProfileOnVMInstanceGroupNetwork**](VMInstanceGroupAPI.md#UpdateNetworkProfileOnVMInstanceGroupNetwork) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/network/{networkId} | Applies the given Network Profile to the specified VM Instance Group Network
[**UpdateVMInstanceGroup**](VMInstanceGroupAPI.md#UpdateVMInstanceGroup) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Updates VM Instance Group information



## ApplyVMTypeOnVMInstanceGroup

> VMInstanceGroup ApplyVMTypeOnVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId, vmTypeId).Execute()

Applies a VM Type to a VM Instance Group



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
	vmInstanceGroupId := float32(8.14) // float32 | 
	vmTypeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.ApplyVMTypeOnVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId, vmTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.ApplyVMTypeOnVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyVMTypeOnVMInstanceGroup`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.ApplyVMTypeOnVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 
**vmTypeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyVMTypeOnVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVMInstanceGroup

> VMInstanceGroup CreateVMInstanceGroup(ctx, infrastructureId).CreateVMInstanceGroup(createVMInstanceGroup).Execute()

Creates a VM Instance Group



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
	createVMInstanceGroup := *openapiclient.NewCreateVMInstanceGroup(float32(1), float32(123), float32(123)) // CreateVMInstanceGroup | The VM Instance Group create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.CreateVMInstanceGroup(context.Background(), infrastructureId).CreateVMInstanceGroup(createVMInstanceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.CreateVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVMInstanceGroup`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.CreateVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVMInstanceGroup** | [**CreateVMInstanceGroup**](CreateVMInstanceGroup.md) | The VM Instance Group create object | 

### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVMInterfaceOnVMInstanceGroup

> VMInstanceGroupInterface CreateVMInterfaceOnVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId).CreateVMInstanceGroupInterface(createVMInstanceGroupInterface).Execute()

Creates a new Virtual Interface for the VM Instance Group



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
	vmInstanceGroupId := float32(8.14) // float32 | 
	createVMInstanceGroupInterface := *openapiclient.NewCreateVMInstanceGroupInterface() // CreateVMInstanceGroupInterface | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.CreateVMInterfaceOnVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).CreateVMInstanceGroupInterface(createVMInstanceGroupInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.CreateVMInterfaceOnVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVMInterfaceOnVMInstanceGroup`: VMInstanceGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.CreateVMInterfaceOnVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMInterfaceOnVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVMInstanceGroupInterface** | [**CreateVMInstanceGroupInterface**](CreateVMInstanceGroupInterface.md) |  | 

### Return type

[**VMInstanceGroupInterface**](VMInstanceGroupInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVMInstanceGroup

> DeleteVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId).Execute()

Deletes a VM Instance Group



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
	vmInstanceGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VMInstanceGroupAPI.DeleteVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.DeleteVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMInstanceGroupRequest struct via the builder pattern


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


## GetVMInstanceGroup

> VMInstanceGroup GetVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId).Execute()

Get VM Instance Group information



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
	vmInstanceGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroup`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroupVMInstances

> []VMInstance GetVMInstanceGroupVMInstances(ctx, infrastructureId, vmInstanceGroupId).Execute()

Get the VM Instances of VM Instance Group



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
	vmInstanceGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupVMInstances(context.Background(), infrastructureId, vmInstanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroupVMInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroupVMInstances`: []VMInstance
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroupVMInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupVMInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]VMInstance**](VMInstance.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMInstanceGroups

> []VMInstanceGroup GetVMInstanceGroups(ctx, infrastructureId).Execute()

Get all VM Instance Groups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroups(context.Background(), infrastructureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.GetVMInstanceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMInstanceGroups`: []VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.GetVMInstanceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMInstanceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkProfileOnVMInstanceGroupNetwork

> VMInstanceGroup UpdateNetworkProfileOnVMInstanceGroupNetwork(ctx, infrastructureId, vmInstanceGroupId, networkId).UpdateVMInstanceGroupNetwork(updateVMInstanceGroupNetwork).Execute()

Applies the given Network Profile to the specified VM Instance Group Network



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
	vmInstanceGroupId := float32(8.14) // float32 | 
	networkId := float32(8.14) // float32 | 
	updateVMInstanceGroupNetwork := *openapiclient.NewUpdateVMInstanceGroupNetwork() // UpdateVMInstanceGroupNetwork | The VM Instance Group Network update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.UpdateNetworkProfileOnVMInstanceGroupNetwork(context.Background(), infrastructureId, vmInstanceGroupId, networkId).UpdateVMInstanceGroupNetwork(updateVMInstanceGroupNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.UpdateNetworkProfileOnVMInstanceGroupNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkProfileOnVMInstanceGroupNetwork`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.UpdateNetworkProfileOnVMInstanceGroupNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 
**networkId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkProfileOnVMInstanceGroupNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateVMInstanceGroupNetwork** | [**UpdateVMInstanceGroupNetwork**](UpdateVMInstanceGroupNetwork.md) | The VM Instance Group Network update object | 

### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMInstanceGroup

> VMInstanceGroup UpdateVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId).UpdateVMInstanceGroup(updateVMInstanceGroup).Execute()

Updates VM Instance Group information



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
	vmInstanceGroupId := float32(8.14) // float32 | 
	updateVMInstanceGroup := *openapiclient.NewUpdateVMInstanceGroup() // UpdateVMInstanceGroup | The VM Instance Group update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMInstanceGroupAPI.UpdateVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).UpdateVMInstanceGroup(updateVMInstanceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMInstanceGroupAPI.UpdateVMInstanceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVMInstanceGroup`: VMInstanceGroup
	fmt.Fprintf(os.Stdout, "Response from `VMInstanceGroupAPI.UpdateVMInstanceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**infrastructureId** | **float32** |  | 
**vmInstanceGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMInstanceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMInstanceGroup** | [**UpdateVMInstanceGroup**](UpdateVMInstanceGroup.md) | The VM Instance Group update object | 

### Return type

[**VMInstanceGroup**](VMInstanceGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

