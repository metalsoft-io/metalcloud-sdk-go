# \LogicalNetworkProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfile) | **Post** /api/v2/logical-network-profiles | Create a Logical Network Profile.
[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies | Create Ipv4 Subnet allocation strategy.
[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies | Create Ipv6 Subnet allocation strategy.
[**CreateLogicalNetworkProfilePkeyAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfilePkeyAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/pkey/pkey-allocation-strategies | Create Pkey allocation strategy.
[**CreateLogicalNetworkProfileVlanAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfileVlanAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies | Create Vlan allocation strategy.
[**CreateLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#CreateLogicalNetworkProfileVniAllocationStrategy) | **Post** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies | Create Vni allocation strategy.
[**DeleteLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfile) | **Delete** /api/v2/logical-network-profiles/{id} | Delete a Logical Network Profile.
[**DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv4 Subnet allocation strategy.
[**DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv6 Subnet allocation strategy.
[**DeleteLogicalNetworkProfilePkeyAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfilePkeyAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/pkey/pkey-allocation-strategies/{allocationStrategyId} | Delete Pkey allocation strategy.
[**DeleteLogicalNetworkProfileVlanAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileVlanAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies/{allocationStrategyId} | Delete Vlan allocation strategy.
[**DeleteLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#DeleteLogicalNetworkProfileVniAllocationStrategy) | **Delete** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies/{allocationStrategyId} | Delete Vni allocation strategy.
[**GetLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfile) | **Get** /api/v2/logical-network-profiles/{id} | Get a Logical Network Profile.
[**GetLogicalNetworkProfileIpv4SubnetAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv4SubnetAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies | Get all Ipv4 Subnet allocation strategies.
[**GetLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv4 Subnet allocation strategy.
[**GetLogicalNetworkProfileIpv6SubnetAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv6SubnetAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies | Get all Ipv6 Subnet allocation strategies.
[**GetLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv6 Subnet allocation strategy.
[**GetLogicalNetworkProfilePkeyAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfilePkeyAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/pkey/pkey-allocation-strategies | Get all Pkey allocation strategies.
[**GetLogicalNetworkProfilePkeyAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfilePkeyAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/pkey/pkey-allocation-strategies/{allocationStrategyId} | Get a Pkey allocation strategy.
[**GetLogicalNetworkProfileVlanAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVlanAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies | Get all Vlan allocation strategies.
[**GetLogicalNetworkProfileVlanAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVlanAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies/{allocationStrategyId} | Get a Vlan allocation strategy.
[**GetLogicalNetworkProfileVniAllocationStrategies**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVniAllocationStrategies) | **Get** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies | Get all Vni allocation strategies.
[**GetLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfileVniAllocationStrategy) | **Get** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies/{allocationStrategyId} | Get a Vni allocation strategy.
[**GetLogicalNetworkProfiles**](LogicalNetworkProfileAPI.md#GetLogicalNetworkProfiles) | **Get** /api/v2/logical-network-profiles | Get all Logical Network Profiles.
[**ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv4 Subnet allocation strategy
[**ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv6 Subnet allocation strategy
[**ReplaceLogicalNetworkProfilePkeyAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfilePkeyAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/pkey/pkey-allocation-strategies/{allocationStrategyId} | Replace Pkey allocation strategy
[**ReplaceLogicalNetworkProfileVlanAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileVlanAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/vlan/vlan-allocation-strategies/{allocationStrategyId} | Replace Vlan allocation strategy
[**ReplaceLogicalNetworkProfileVniAllocationStrategy**](LogicalNetworkProfileAPI.md#ReplaceLogicalNetworkProfileVniAllocationStrategy) | **Put** /api/v2/logical-network-profiles/{id}/vxlan/vni-allocation-strategies/{allocationStrategyId} | Replace Vni allocation strategy
[**UpdateLogicalNetworkProfile**](LogicalNetworkProfileAPI.md#UpdateLogicalNetworkProfile) | **Patch** /api/v2/logical-network-profiles/{id} | Update Logical Network Profile



## CreateLogicalNetworkProfile

> LogicalNetworkProfile CreateLogicalNetworkProfile(ctx).CreateLogicalNetworkProfile(createLogicalNetworkProfile).Execute()

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
	createLogicalNetworkProfile := *openapiclient.NewCreateLogicalNetworkProfile(openapiclient.LogicalNetworkKind("vlan"), int32(123)) // CreateLogicalNetworkProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfile(context.Background()).CreateLogicalNetworkProfile(createLogicalNetworkProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfile`: LogicalNetworkProfile
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.CreateLogicalNetworkProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogicalNetworkProfile** | [**CreateLogicalNetworkProfile**](CreateLogicalNetworkProfile.md) |  | 

### Return type

[**LogicalNetworkProfile**](LogicalNetworkProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy

> Ipv4SubnetAllocationStrategy CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy(ctx, id).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).IfMatch(ifMatch).Execute()

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
	createIpv4SubnetAllocationStrategy := openapiclient.CreateIpv4SubnetAllocationStrategy{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv4SubnetAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy(context.Background(), id).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy`: Ipv4SubnetAllocationStrategy
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

 **createIpv4SubnetAllocationStrategy** | [**CreateIpv4SubnetAllocationStrategy**](CreateIpv4SubnetAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Ipv4SubnetAllocationStrategy**](Ipv4SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy

> Ipv6SubnetAllocationStrategy CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy(ctx, id).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).IfMatch(ifMatch).Execute()

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
	createIpv6SubnetAllocationStrategy := openapiclient.CreateIpv6SubnetAllocationStrategy{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv6SubnetAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy(context.Background(), id).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy`: Ipv6SubnetAllocationStrategy
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

 **createIpv6SubnetAllocationStrategy** | [**CreateIpv6SubnetAllocationStrategy**](CreateIpv6SubnetAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Ipv6SubnetAllocationStrategy**](Ipv6SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfilePkeyAllocationStrategy

> PkeyAllocationStrategy CreateLogicalNetworkProfilePkeyAllocationStrategy(ctx, id).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).IfMatch(ifMatch).Execute()

Create Pkey allocation strategy.

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
	createPkeyAllocationStrategy := openapiclient.CreatePkeyAllocationStrategy{CreateAutoPkeyAllocationStrategy: openapiclient.NewCreateAutoPkeyAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreatePkeyAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfilePkeyAllocationStrategy(context.Background(), id).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfilePkeyAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfilePkeyAllocationStrategy`: PkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.CreateLogicalNetworkProfilePkeyAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfilePkeyAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPkeyAllocationStrategy** | [**CreatePkeyAllocationStrategy**](CreatePkeyAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**PkeyAllocationStrategy**](PkeyAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfileVlanAllocationStrategy

> VlanAllocationStrategy CreateLogicalNetworkProfileVlanAllocationStrategy(ctx, id).CreateVlanAllocationStrategy(createVlanAllocationStrategy).IfMatch(ifMatch).Execute()

Create Vlan allocation strategy.

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
	createVlanAllocationStrategy := openapiclient.CreateVlanAllocationStrategy{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVlanAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVlanAllocationStrategy(context.Background(), id).CreateVlanAllocationStrategy(createVlanAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfileVlanAllocationStrategy`: VlanAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkProfileVlanAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVlanAllocationStrategy** | [**CreateVlanAllocationStrategy**](CreateVlanAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**VlanAllocationStrategy**](VlanAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkProfileVniAllocationStrategy

> VniAllocationStrategy CreateLogicalNetworkProfileVniAllocationStrategy(ctx, id).CreateVniAllocationStrategy(createVniAllocationStrategy).IfMatch(ifMatch).Execute()

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
	createVniAllocationStrategy := openapiclient.CreateVniAllocationStrategy{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVniAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVniAllocationStrategy(context.Background(), id).CreateVniAllocationStrategy(createVniAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.CreateLogicalNetworkProfileVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkProfileVniAllocationStrategy`: VniAllocationStrategy
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

 **createVniAllocationStrategy** | [**CreateVniAllocationStrategy**](CreateVniAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**VniAllocationStrategy**](VniAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetworkProfilePkeyAllocationStrategy

> DeleteLogicalNetworkProfilePkeyAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete Pkey allocation strategy.

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
	r, err := apiClient.LogicalNetworkProfileAPI.DeleteLogicalNetworkProfilePkeyAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.DeleteLogicalNetworkProfilePkeyAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkProfilePkeyAllocationStrategyRequest struct via the builder pattern


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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfile

> LogicalNetworkProfile GetLogicalNetworkProfile(ctx, id).Execute()

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
	// response from `GetLogicalNetworkProfile`: LogicalNetworkProfile
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

[**LogicalNetworkProfile**](LogicalNetworkProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 

### Return type

[**PaginatedIpv4SubnetAllocationStrategy**](PaginatedIpv4SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileIpv4SubnetAllocationStrategy

> Ipv4SubnetAllocationStrategy GetLogicalNetworkProfileIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkProfileIpv4SubnetAllocationStrategy`: Ipv4SubnetAllocationStrategy
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

[**Ipv4SubnetAllocationStrategy**](Ipv4SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 

### Return type

[**PaginatedIpv6SubnetAllocationStrategy**](PaginatedIpv6SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileIpv6SubnetAllocationStrategy

> Ipv6SubnetAllocationStrategy GetLogicalNetworkProfileIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkProfileIpv6SubnetAllocationStrategy`: Ipv6SubnetAllocationStrategy
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

[**Ipv6SubnetAllocationStrategy**](Ipv6SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfilePkeyAllocationStrategies

> PaginatedPkeyAllocationStrategy GetLogicalNetworkProfilePkeyAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all Pkey allocation strategies.

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfilePkeyAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfilePkeyAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfilePkeyAllocationStrategies`: PaginatedPkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfilePkeyAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfilePkeyAllocationStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 

### Return type

[**PaginatedPkeyAllocationStrategy**](PaginatedPkeyAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfilePkeyAllocationStrategy

> PkeyAllocationStrategy GetLogicalNetworkProfilePkeyAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a Pkey allocation strategy.

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
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfilePkeyAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfilePkeyAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfilePkeyAllocationStrategy`: PkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfilePkeyAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfilePkeyAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PkeyAllocationStrategy**](PkeyAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 

### Return type

[**PaginatedVlanAllocationStrategy**](PaginatedVlanAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileVlanAllocationStrategy

> VlanAllocationStrategy GetLogicalNetworkProfileVlanAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkProfileVlanAllocationStrategy`: VlanAllocationStrategy
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

[**VlanAllocationStrategy**](VlanAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 

### Return type

[**PaginatedVniAllocationStrategy**](PaginatedVniAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfileVniAllocationStrategy

> VniAllocationStrategy GetLogicalNetworkProfileVniAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkProfileVniAllocationStrategy`: VniAllocationStrategy
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

[**VniAllocationStrategy**](VniAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkProfiles

> PaginatedLogicalNetworkProfileList GetLogicalNetworkProfiles(ctx).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** -1   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe&filter.id=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$eq:John Doe&filter.kind=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterFabricId := []string{"Inner_example"} // []string | Filter by fabricId query param.  **Format:** filter.fabricId={$not}:OPERATION:VALUE    **Example:** filter.fabricId=$eq:John Doe&filter.fabricId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.GetLogicalNetworkProfiles(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.GetLogicalNetworkProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkProfiles`: PaginatedLogicalNetworkProfileList
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.GetLogicalNetworkProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** -1   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$eq:John Doe&amp;filter.kind&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterFabricId** | **[]string** | Filter by fabricId query param.  **Format:** filter.fabricId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricId&#x3D;$eq:John Doe&amp;filter.fabricId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  | 

### Return type

[**PaginatedLogicalNetworkProfileList**](PaginatedLogicalNetworkProfileList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy

> Ipv4SubnetAllocationStrategy ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).IfMatch(ifMatch).Execute()

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
	createIpv4SubnetAllocationStrategy := openapiclient.CreateIpv4SubnetAllocationStrategy{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv4SubnetAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileIpv4SubnetAllocationStrategy`: Ipv4SubnetAllocationStrategy
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


 **createIpv4SubnetAllocationStrategy** | [**CreateIpv4SubnetAllocationStrategy**](CreateIpv4SubnetAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Ipv4SubnetAllocationStrategy**](Ipv4SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy

> Ipv6SubnetAllocationStrategy ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).IfMatch(ifMatch).Execute()

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
	createIpv6SubnetAllocationStrategy := openapiclient.CreateIpv6SubnetAllocationStrategy{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv6SubnetAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileIpv6SubnetAllocationStrategy`: Ipv6SubnetAllocationStrategy
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


 **createIpv6SubnetAllocationStrategy** | [**CreateIpv6SubnetAllocationStrategy**](CreateIpv6SubnetAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**Ipv6SubnetAllocationStrategy**](Ipv6SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfilePkeyAllocationStrategy

> PkeyAllocationStrategy ReplaceLogicalNetworkProfilePkeyAllocationStrategy(ctx, id, allocationStrategyId).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).IfMatch(ifMatch).Execute()

Replace Pkey allocation strategy

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
	createPkeyAllocationStrategy := openapiclient.CreatePkeyAllocationStrategy{CreateAutoPkeyAllocationStrategy: openapiclient.NewCreateAutoPkeyAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreatePkeyAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfilePkeyAllocationStrategy(context.Background(), id, allocationStrategyId).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfilePkeyAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfilePkeyAllocationStrategy`: PkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfilePkeyAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkProfilePkeyAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createPkeyAllocationStrategy** | [**CreatePkeyAllocationStrategy**](CreatePkeyAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**PkeyAllocationStrategy**](PkeyAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileVlanAllocationStrategy

> VlanAllocationStrategy ReplaceLogicalNetworkProfileVlanAllocationStrategy(ctx, id, allocationStrategyId).CreateVlanAllocationStrategy(createVlanAllocationStrategy).IfMatch(ifMatch).Execute()

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
	createVlanAllocationStrategy := openapiclient.CreateVlanAllocationStrategy{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVlanAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVlanAllocationStrategy(context.Background(), id, allocationStrategyId).CreateVlanAllocationStrategy(createVlanAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileVlanAllocationStrategy`: VlanAllocationStrategy
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


 **createVlanAllocationStrategy** | [**CreateVlanAllocationStrategy**](CreateVlanAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**VlanAllocationStrategy**](VlanAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkProfileVniAllocationStrategy

> VniAllocationStrategy ReplaceLogicalNetworkProfileVniAllocationStrategy(ctx, id, allocationStrategyId).CreateVniAllocationStrategy(createVniAllocationStrategy).IfMatch(ifMatch).Execute()

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
	createVniAllocationStrategy := openapiclient.CreateVniAllocationStrategy{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVniAllocationStrategy | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVniAllocationStrategy(context.Background(), id, allocationStrategyId).CreateVniAllocationStrategy(createVniAllocationStrategy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.ReplaceLogicalNetworkProfileVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkProfileVniAllocationStrategy`: VniAllocationStrategy
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


 **createVniAllocationStrategy** | [**CreateVniAllocationStrategy**](CreateVniAllocationStrategy.md) |  | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**VniAllocationStrategy**](VniAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalNetworkProfile

> LogicalNetworkProfile UpdateLogicalNetworkProfile(ctx, id).UpdateLogicalNetworkProfile(updateLogicalNetworkProfile).IfMatch(ifMatch).Execute()

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
	updateLogicalNetworkProfile := *openapiclient.NewUpdateLogicalNetworkProfile() // UpdateLogicalNetworkProfile | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkProfileAPI.UpdateLogicalNetworkProfile(context.Background(), id).UpdateLogicalNetworkProfile(updateLogicalNetworkProfile).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkProfileAPI.UpdateLogicalNetworkProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalNetworkProfile`: LogicalNetworkProfile
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

[**LogicalNetworkProfile**](LogicalNetworkProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

