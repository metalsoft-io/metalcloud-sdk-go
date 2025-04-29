# \LogicalNetworkProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfile) | **Post** /api/v2/logical-network-profiles | Create a Logical Network Profile.
[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies | Create Ipv4 Subnet allocation strategy.
[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies | Create Ipv6 Subnet allocation strategy.
[**CreateLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfileVniAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies | Create Vni allocation strategy.
[**DeleteLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfile) | **Delete** /api/v2/logical-network-profiles/{id} | Delete a Logical Network Profile.
[**DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv4 Subnet allocation strategy.
[**DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv6 Subnet allocation strategy.
[**DeleteLogicalNetworkProfileVlanAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileVlanAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies/{allocationStrategyId} | Delete Vlan allocation strategy.
[**DeleteLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileVniAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies/{allocationStrategyId} | Delete Vni allocation strategy.
[**GetLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfile) | **Get** /api/v2/logical-network-profiles/{id} | Get a Logical Network Profile.
[**GetLogicalNetworkProfileIpv4SubnetAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv4SubnetAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies | Get all Ipv4 Subnet allocation strategies.
[**GetLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv4 Subnet allocation strategy.
[**GetLogicalNetworkProfileIpv6SubnetAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv6SubnetAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies | Get all Ipv6 Subnet allocation strategies.
[**GetLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv6 Subnet allocation strategy.
[**GetLogicalNetworkProfileVlanAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVlanAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies | Get all Vlan allocation strategies.
[**GetLogicalNetworkProfileVlanAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVlanAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies/{allocationStrategyId} | Get a Vlan allocation strategy.
[**GetLogicalNetworkProfileVniAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVniAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies | Get all Vni allocation strategies.
[**GetLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVniAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies/{allocationStrategyId} | Get a Vni allocation strategy.
[**GetLogicalNetworkProfiles**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfiles) | **Get** /api/v2/logical-network-profiles | Get all Logical Network Profiles.
[**ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv4 Subnet allocation strategy
[**ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv6 Subnet allocation strategy
[**ReplaceLogicalNetworkProfileVlanAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileVlanAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies/{allocationStrategyId} | Replace Vlan allocation strategy
[**ReplaceLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileVniAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies/{allocationStrategyId} | Replace Vni allocation strategy
[**UpdateLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#UpdateLogicalNetworkProfile) | **Patch** /api/v2/logical-network-profiles/{id} | Update Logical Network Profile



## CreateLogicalNetworkProfile

> CreateLogicalNetworkProfile201Response CreateLogicalNetworkProfile(ctx).CreateLogicalNetworkProfileRequest(createLogicalNetworkProfileRequest).Execute()

Create a Logical Network Profile.

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
	createLogicalNetworkProfileRequest := openapiclient.createLogicalNetworkProfile_request{CreateVlanLogicalNetworkProfile: openapiclient.NewCreateVlanLogicalNetworkProfile(openapiclient.LogicalNetworkKind("vlan"), int32(123), *openapiclient.NewCreateVlanLogicalNetworkVlanProperties([]openapiclient.CreateVlanAllocationStrategy{openapiclient.CreateVlanAllocationStrategy{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))}}), *openapiclient.NewCreateVlanLogicalNetworkIpv4Properties([]openapiclient.CreateIpv4SubnetAllocationStrategy{openapiclient.CreateIpv4SubnetAllocationStrategy{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))}}))} // CreateLogicalNetworkProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfile(context.Background()).CreateLogicalNetworkProfileRequest(createLogicalNetworkProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfile`: CreateLogicalNetworkProfile201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.CreateLogicalNetworkProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogicalNetworkProfileRequest** | [**CreateLogicalNetworkProfileRequest**](CreateLogicalNetworkProfileRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfile201Response**](CreateLogicalNetworkProfile201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy(ctx, id).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()

Create Ipv4 Subnet allocation strategy.

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
	id := float32(8.14) // float32 | 
	createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv4SubnetAllocationStrategy_request{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy(context.Background(), id).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy(ctx, id).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()

Create Ipv6 Subnet allocation strategy.

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
	id := float32(8.14) // float32 | 
	createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv6SubnetAllocationStrategy_request{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy(context.Background(), id).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfileVniAllocationStrategy

> CreateLogicalNetworkProfileVniAllocationStrategy201Response CreateLogicalNetworkProfileVniAllocationStrategy(ctx, id).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).IfMatch(ifMatch).Execute()

Create Vni allocation strategy.

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
	id := float32(8.14) // float32 | 
	createLogicalNetworkProfileVniAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileVniAllocationStrategy_request{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateLogicalNetworkProfileVniAllocationStrategyRequest | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVniAllocationStrategy(context.Background(), id).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfileVniAllocationStrategy`: CreateLogicalNetworkProfileVniAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfileVniAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLogicalNetworkProfileVniAllocationStrategyRequest** | [**CreateLogicalNetworkProfileVniAllocationStrategyRequest**](CreateLogicalNetworkProfileVniAllocationStrategyRequest.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**CreateLogicalNetworkProfileVniAllocationStrategy201Response**](CreateLogicalNetworkProfileVniAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkProfile

> DeleteLogicalNetworkProfile(ctx, id).IfMatch(ifMatch).Execute()

Delete a Logical Network Profile.

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
	id := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkProfileAPI.DeleteLogicalNetworkProfile(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.DeleteLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy

> DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete Ipv4 Subnet allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy

> DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete Ipv6 Subnet allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkProfileVlanAllocationStrategy

> DeleteLogicalNetworkProfileVlanAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete Vlan allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileVlanAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkProfileVlanAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkProfileVniAllocationStrategy

> DeleteLogicalNetworkProfileVniAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete Vni allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileVniAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.DeleteLogicalNetworkProfileVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkProfileVniAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfile

> CreateLogicalNetworkProfile201Response GetLogicalNetworkProfile(ctx, id).Execute()

Get a Logical Network Profile.

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
	id := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfile`: CreateLogicalNetworkProfile201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateLogicalNetworkProfile201Response**](CreateLogicalNetworkProfile201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileIpv4SubnetAllocationStrategies

> PaginatedIpv4SubnetAllocationStrategy GetLogicalNetworkProfileIpv4SubnetAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all Ipv4 Subnet allocation strategies.

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
	id := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 100           </p>       <p>              <b>Max Value: </b> 1000           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.           <p>              <b>Format: </b> filter.kind={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.kind=$not:$like:John Doe&filter.kind=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv4SubnetAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv4SubnetAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileIpv4SubnetAllocationStrategies`: PaginatedIpv4SubnetAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv4SubnetAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileIpv4SubnetAllocationStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 100           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 1000           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterKind** | **[]string** | Filter by kind query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.kind&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.kind&#x3D;$not:$like:John Doe&amp;filter.kind&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 

### Return type

[**PaginatedIpv4SubnetAllocationStrategy**](PaginatedIpv4SubnetAllocationStrategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileIpv4SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response GetLogicalNetworkProfileIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a Ipv4 Subnet allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileIpv4SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv4SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileIpv6SubnetAllocationStrategies

> PaginatedIpv6SubnetAllocationStrategy GetLogicalNetworkProfileIpv6SubnetAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all Ipv6 Subnet allocation strategies.

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
	id := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 100           </p>       <p>              <b>Max Value: </b> 1000           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.           <p>              <b>Format: </b> filter.kind={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.kind=$not:$like:John Doe&filter.kind=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv6SubnetAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv6SubnetAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileIpv6SubnetAllocationStrategies`: PaginatedIpv6SubnetAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv6SubnetAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileIpv6SubnetAllocationStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 100           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 1000           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterKind** | **[]string** | Filter by kind query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.kind&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.kind&#x3D;$not:$like:John Doe&amp;filter.kind&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 

### Return type

[**PaginatedIpv6SubnetAllocationStrategy**](PaginatedIpv6SubnetAllocationStrategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileIpv6SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response GetLogicalNetworkProfileIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a Ipv6 Subnet allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileIpv6SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileIpv6SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileVlanAllocationStrategies

> PaginatedVlanAllocationStrategy GetLogicalNetworkProfileVlanAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all Vlan allocation strategies.

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
	id := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 100           </p>       <p>              <b>Max Value: </b> 1000           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.           <p>              <b>Format: </b> filter.kind={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.kind=$not:$like:John Doe&filter.kind=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileVlanAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVlanAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileVlanAllocationStrategies`: PaginatedVlanAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVlanAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileVlanAllocationStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 100           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 1000           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterKind** | **[]string** | Filter by kind query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.kind&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.kind&#x3D;$not:$like:John Doe&amp;filter.kind&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 

### Return type

[**PaginatedVlanAllocationStrategy**](PaginatedVlanAllocationStrategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileVlanAllocationStrategy

> GetLogicalNetworkProfileVlanAllocationStrategy200Response GetLogicalNetworkProfileVlanAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a Vlan allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileVlanAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileVlanAllocationStrategy`: GetLogicalNetworkProfileVlanAllocationStrategy200Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileVlanAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLogicalNetworkProfileVlanAllocationStrategy200Response**](GetLogicalNetworkProfileVlanAllocationStrategy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileVniAllocationStrategies

> PaginatedVniAllocationStrategy GetLogicalNetworkProfileVniAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all Vni allocation strategies.

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
	id := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 100           </p>       <p>              <b>Max Value: </b> 1000           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.           <p>              <b>Format: </b> filter.kind={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.kind=$not:$like:John Doe&filter.kind=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileVniAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVniAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileVniAllocationStrategies`: PaginatedVniAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVniAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileVniAllocationStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 100           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 1000           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterKind** | **[]string** | Filter by kind query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.kind&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.kind&#x3D;$not:$like:John Doe&amp;filter.kind&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 

### Return type

[**PaginatedVniAllocationStrategy**](PaginatedVniAllocationStrategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileVniAllocationStrategy

> CreateLogicalNetworkProfileVniAllocationStrategy201Response GetLogicalNetworkProfileVniAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a Vni allocation strategy.

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfileVniAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfileVniAllocationStrategy`: CreateLogicalNetworkProfileVniAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfileVniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfileVniAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateLogicalNetworkProfileVniAllocationStrategy201Response**](CreateLogicalNetworkProfileVniAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfiles

> PaginatedLogicalNetworkProfile GetLogicalNetworkProfiles(ctx).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Logical Network Profiles.

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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.           <p>              <b>Format: </b> filter.kind={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.kind=$not:$like:John Doe&filter.kind=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterFabricId := []string{"Inner_example"} // []string | Filter by fabricId query param.           <p>              <b>Format: </b> filter.fabricId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.fabricId=$not:$like:John Doe&filter.fabricId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,name           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>name</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfiles(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfiles`: PaginatedLogicalNetworkProfile
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterKind** | **[]string** | Filter by kind query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.kind&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.kind&#x3D;$not:$like:John Doe&amp;filter.kind&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterFabricId** | **[]string** | Filter by fabricId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.fabricId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.fabricId&#x3D;$not:$like:John Doe&amp;filter.fabricId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**PaginatedLogicalNetworkProfile**](PaginatedLogicalNetworkProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()

Replace Ipv4 Subnet allocation strategy

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv4SubnetAllocationStrategy_request{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()

Replace Ipv6 Subnet allocation strategy

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv6SubnetAllocationStrategy_request{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileVlanAllocationStrategy

> GetLogicalNetworkProfileVlanAllocationStrategy200Response ReplaceLogicalNetworkProfileVlanAllocationStrategy(ctx, id, allocationStrategyId).ReplaceLogicalNetworkProfileVlanAllocationStrategyRequest(replaceLogicalNetworkProfileVlanAllocationStrategyRequest).IfMatch(ifMatch).Execute()

Replace Vlan allocation strategy

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	replaceLogicalNetworkProfileVlanAllocationStrategyRequest := openapiclient.replaceLogicalNetworkProfileVlanAllocationStrategy_request{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // ReplaceLogicalNetworkProfileVlanAllocationStrategyRequest | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVlanAllocationStrategy(context.Background(), id, allocationStrategyId).ReplaceLogicalNetworkProfileVlanAllocationStrategyRequest(replaceLogicalNetworkProfileVlanAllocationStrategyRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileVlanAllocationStrategy`: GetLogicalNetworkProfileVlanAllocationStrategy200Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkProfileVlanAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceLogicalNetworkProfileVlanAllocationStrategyRequest** | [**ReplaceLogicalNetworkProfileVlanAllocationStrategyRequest**](ReplaceLogicalNetworkProfileVlanAllocationStrategyRequest.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**GetLogicalNetworkProfileVlanAllocationStrategy200Response**](GetLogicalNetworkProfileVlanAllocationStrategy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileVniAllocationStrategy

> CreateLogicalNetworkProfileVniAllocationStrategy201Response ReplaceLogicalNetworkProfileVniAllocationStrategy(ctx, id, allocationStrategyId).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).IfMatch(ifMatch).Execute()

Replace Vni allocation strategy

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
	id := float32(8.14) // float32 | 
	allocationStrategyId := float32(8.14) // float32 | 
	createLogicalNetworkProfileVniAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileVniAllocationStrategy_request{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateLogicalNetworkProfileVniAllocationStrategyRequest | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVniAllocationStrategy(context.Background(), id, allocationStrategyId).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileVniAllocationStrategy`: CreateLogicalNetworkProfileVniAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkProfileVniAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLogicalNetworkProfileVniAllocationStrategyRequest** | [**CreateLogicalNetworkProfileVniAllocationStrategyRequest**](CreateLogicalNetworkProfileVniAllocationStrategyRequest.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**CreateLogicalNetworkProfileVniAllocationStrategy201Response**](CreateLogicalNetworkProfileVniAllocationStrategy201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalNetworkProfile

> CreateLogicalNetworkProfile201Response UpdateLogicalNetworkProfile(ctx, id).UpdateLogicalNetworkProfile(updateLogicalNetworkProfile).IfMatch(ifMatch).Execute()

Update Logical Network Profile

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
	id := float32(8.14) // float32 | 
	updateLogicalNetworkProfile := *openapiclient.NewUpdateLogicalNetworkProfile(*openapiclient.NewUpdateLogicalNetworkProfileVlanProperties(), *openapiclient.NewUpdateLogicalNetworkProfileVxlanProperties(), *openapiclient.NewUpdateLogicalNetworkProfileIpv4Properties(), *openapiclient.NewUpdateLogicalNetworkProfileIpv6Properties()) // UpdateLogicalNetworkProfile | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.UpdateLogicalNetworkProfile(context.Background(), id).UpdateLogicalNetworkProfile(updateLogicalNetworkProfile).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.UpdateLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalNetworkProfile`: CreateLogicalNetworkProfile201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.UpdateLogicalNetworkProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogicalNetworkProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLogicalNetworkProfile** | [**UpdateLogicalNetworkProfile**](UpdateLogicalNetworkProfile.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**CreateLogicalNetworkProfile201Response**](CreateLogicalNetworkProfile201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

