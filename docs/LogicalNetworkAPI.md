# \LogicalNetworkAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyProfilesToLogicalNetworkConfig**](LogicalNetworkAPI.md#ApplyProfilesToLogicalNetworkConfig) | **Post** /api/v2/logical-networks/{id}/config/actions/apply-profiles | Apply profiles to Logical Network config
[**CreateLogicalNetwork**](LogicalNetworkAPI.md#CreateLogicalNetwork) | **Post** /api/v2/logical-networks | Create a Logical Network.
[**CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies | Create Ipv4 Subnet allocation strategy.
[**CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies | Create Ipv6 Subnet allocation strategy.
[**CreateLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigVlanAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies | Create Vlan allocation strategy.
[**CreateLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigVniAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies | Create Vni allocation strategy.
[**CreateLogicalNetworkFromProfile**](LogicalNetworkAPI.md#CreateLogicalNetworkFromProfile) | **Post** /api/v2/logical-networks/actions/create-from-profile | Create a Logical Network from a profile.
[**DeleteLogicalNetwork**](LogicalNetworkAPI.md#DeleteLogicalNetwork) | **Delete** /api/v2/logical-networks/{id} | Delete a Logical Network.
[**DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv4 Subnet allocation strategy.
[**DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv6 Subnet allocation strategy.
[**DeleteLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigVlanAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies/{allocationStrategyId} | Delete Vlan allocation strategy.
[**DeleteLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigVniAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies/{allocationStrategyId} | Delete Vni allocation strategy.
[**GetLogicalNetwork**](LogicalNetworkAPI.md#GetLogicalNetwork) | **Get** /api/v2/logical-networks/{id} | Get a Logical Network.
[**GetLogicalNetworkConfig**](LogicalNetworkAPI.md#GetLogicalNetworkConfig) | **Get** /api/v2/logical-networks/{id}/config | Get the config for a Logical Network.
[**GetLogicalNetworkConfigIpv4SubnetAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv4SubnetAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies | Get all Ipv4 Subnet allocation strategies.
[**GetLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv4 Subnet allocation strategy.
[**GetLogicalNetworkConfigIpv6SubnetAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv6SubnetAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies | Get all Ipv6 Subnet allocation strategies.
[**GetLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv6 Subnet allocation strategy.
[**GetLogicalNetworkConfigVlanAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVlanAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies | Get all Vlan allocation strategies.
[**GetLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVlanAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies/{allocationStrategyId} | Get a Vlan allocation strategy.
[**GetLogicalNetworkConfigVniAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVniAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies | Get all Vni allocation strategies.
[**GetLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVniAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies/{allocationStrategyId} | Get a Vni allocation strategy.
[**GetLogicalNetworks**](LogicalNetworkAPI.md#GetLogicalNetworks) | **Get** /api/v2/logical-networks | Get all Logical Networks
[**ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv4 Subnet allocation strategy
[**ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv6 Subnet allocation strategy
[**ReplaceLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigVlanAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies/{allocationStrategyId} | Replace Vlan allocation strategy
[**ReplaceLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigVniAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies/{allocationStrategyId} | Replace Vni allocation strategy
[**UpdateLogicalNetwork**](LogicalNetworkAPI.md#UpdateLogicalNetwork) | **Patch** /api/v2/logical-networks/{id} | Update Logical Network
[**UpdateLogicalNetworkConfig**](LogicalNetworkAPI.md#UpdateLogicalNetworkConfig) | **Patch** /api/v2/logical-networks/{id}/config | Update Logical Network config



## ApplyProfilesToLogicalNetworkConfig

> GetLogicalNetworkConfig200Response ApplyProfilesToLogicalNetworkConfig(ctx, id).IfMatch(ifMatch).ApplyProfilesToLogicalNetworkConfig(applyProfilesToLogicalNetworkConfig).Execute()

Apply profiles to Logical Network config

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
	ifMatch := "ifMatch_example" // string | Entity tag
	applyProfilesToLogicalNetworkConfig := *openapiclient.NewApplyProfilesToLogicalNetworkConfig() // ApplyProfilesToLogicalNetworkConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ApplyProfilesToLogicalNetworkConfig(context.Background(), id).IfMatch(ifMatch).ApplyProfilesToLogicalNetworkConfig(applyProfilesToLogicalNetworkConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ApplyProfilesToLogicalNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyProfilesToLogicalNetworkConfig`: GetLogicalNetworkConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.ApplyProfilesToLogicalNetworkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyProfilesToLogicalNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **applyProfilesToLogicalNetworkConfig** | [**ApplyProfilesToLogicalNetworkConfig**](ApplyProfilesToLogicalNetworkConfig.md) |  | 

### Return type

[**GetLogicalNetworkConfig200Response**](GetLogicalNetworkConfig200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetwork

> CreateLogicalNetwork201Response CreateLogicalNetwork(ctx).CreateLogicalNetworkRequest(createLogicalNetworkRequest).Execute()

Create a Logical Network.

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
	createLogicalNetworkRequest := openapiclient.createLogicalNetwork_request{CreateVlanLogicalNetwork: openapiclient.NewCreateVlanLogicalNetwork(openapiclient.LogicalNetworkKind("vlan"), int32(123), *openapiclient.NewCreateVlanLogicalNetworkVlanProperties([]openapiclient.CreateVlanAllocationStrategy{openapiclient.CreateVlanAllocationStrategy{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))}}), *openapiclient.NewCreateVlanLogicalNetworkIpv4Properties([]openapiclient.CreateIpv4SubnetAllocationStrategy{openapiclient.CreateIpv4SubnetAllocationStrategy{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))}}))} // CreateLogicalNetworkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetwork(context.Background()).CreateLogicalNetworkRequest(createLogicalNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetwork`: CreateLogicalNetwork201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogicalNetworkRequest** | [**CreateLogicalNetworkRequest**](CreateLogicalNetworkRequest.md) |  | 

### Return type

[**CreateLogicalNetwork201Response**](CreateLogicalNetwork201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv4SubnetAllocationStrategy_request{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkConfigIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv6SubnetAllocationStrategy_request{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkConfigIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkConfigVlanAllocationStrategy

> CreateLogicalNetworkProfileVlanAllocationStrategy201Response CreateLogicalNetworkConfigVlanAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateLogicalNetworkProfileVlanAllocationStrategyRequest(createLogicalNetworkProfileVlanAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileVlanAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileVlanAllocationStrategy_request{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateLogicalNetworkProfileVlanAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigVlanAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateLogicalNetworkProfileVlanAllocationStrategyRequest(createLogicalNetworkProfileVlanAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigVlanAllocationStrategy`: CreateLogicalNetworkProfileVlanAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetworkConfigVlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkConfigVlanAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileVlanAllocationStrategyRequest** | [**CreateLogicalNetworkProfileVlanAllocationStrategyRequest**](CreateLogicalNetworkProfileVlanAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileVlanAllocationStrategy201Response**](CreateLogicalNetworkProfileVlanAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkConfigVniAllocationStrategy

> CreateLogicalNetworkProfileVniAllocationStrategy201Response CreateLogicalNetworkConfigVniAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileVniAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileVniAllocationStrategy_request{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateLogicalNetworkProfileVniAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigVniAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigVniAllocationStrategy`: CreateLogicalNetworkProfileVniAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetworkConfigVniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkConfigVniAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileVniAllocationStrategyRequest** | [**CreateLogicalNetworkProfileVniAllocationStrategyRequest**](CreateLogicalNetworkProfileVniAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileVniAllocationStrategy201Response**](CreateLogicalNetworkProfileVniAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkFromProfile

> CreateLogicalNetwork201Response CreateLogicalNetworkFromProfile(ctx).CreateLogicalNetworkFromProfile(createLogicalNetworkFromProfile).Execute()

Create a Logical Network from a profile.

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
	createLogicalNetworkFromProfile := *openapiclient.NewCreateLogicalNetworkFromProfile(int32(123)) // CreateLogicalNetworkFromProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkFromProfile(context.Background()).CreateLogicalNetworkFromProfile(createLogicalNetworkFromProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkFromProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkFromProfile`: CreateLogicalNetwork201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetworkFromProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkFromProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogicalNetworkFromProfile** | [**CreateLogicalNetworkFromProfile**](CreateLogicalNetworkFromProfile.md) |  | 

### Return type

[**CreateLogicalNetwork201Response**](CreateLogicalNetwork201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogicalNetwork

> DeleteLogicalNetwork(ctx, id).IfMatch(ifMatch).Execute()

Delete a Logical Network.

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
	ifMatch := "ifMatch_example" // string | Entity tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkAPI.DeleteLogicalNetwork(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DeleteLogicalNetwork``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkRequest struct via the builder pattern


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


## DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy

> DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkAPI.DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkConfigIpv4SubnetAllocationStrategyRequest struct via the builder pattern


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


## DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy

> DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkAPI.DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkConfigIpv6SubnetAllocationStrategyRequest struct via the builder pattern


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


## DeleteLogicalNetworkConfigVlanAllocationStrategy

> DeleteLogicalNetworkConfigVlanAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkAPI.DeleteLogicalNetworkConfigVlanAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DeleteLogicalNetworkConfigVlanAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkConfigVlanAllocationStrategyRequest struct via the builder pattern


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


## DeleteLogicalNetworkConfigVniAllocationStrategy

> DeleteLogicalNetworkConfigVniAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkAPI.DeleteLogicalNetworkConfigVniAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DeleteLogicalNetworkConfigVniAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkConfigVniAllocationStrategyRequest struct via the builder pattern


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


## GetLogicalNetwork

> CreateLogicalNetwork201Response GetLogicalNetwork(ctx, id).Execute()

Get a Logical Network.

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetwork`: CreateLogicalNetwork201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateLogicalNetwork201Response**](CreateLogicalNetwork201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfig

> GetLogicalNetworkConfig200Response GetLogicalNetworkConfig(ctx, id).Execute()

Get the config for a Logical Network.

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfig`: GetLogicalNetworkConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLogicalNetworkConfig200Response**](GetLogicalNetworkConfig200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigIpv4SubnetAllocationStrategies

> PaginatedIpv4SubnetAllocationStrategy GetLogicalNetworkConfigIpv4SubnetAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigIpv4SubnetAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigIpv4SubnetAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigIpv4SubnetAllocationStrategies`: PaginatedIpv4SubnetAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigIpv4SubnetAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigIpv4SubnetAllocationStrategiesRequest struct via the builder pattern


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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigIpv4SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response GetLogicalNetworkConfigIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigIpv4SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigIpv4SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigIpv6SubnetAllocationStrategies

> PaginatedIpv6SubnetAllocationStrategy GetLogicalNetworkConfigIpv6SubnetAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigIpv6SubnetAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigIpv6SubnetAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigIpv6SubnetAllocationStrategies`: PaginatedIpv6SubnetAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigIpv6SubnetAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigIpv6SubnetAllocationStrategiesRequest struct via the builder pattern


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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigIpv6SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response GetLogicalNetworkConfigIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigIpv6SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigIpv6SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigVlanAllocationStrategies

> PaginatedVlanAllocationStrategy GetLogicalNetworkConfigVlanAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigVlanAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigVlanAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigVlanAllocationStrategies`: PaginatedVlanAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigVlanAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigVlanAllocationStrategiesRequest struct via the builder pattern


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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigVlanAllocationStrategy

> CreateLogicalNetworkProfileVlanAllocationStrategy201Response GetLogicalNetworkConfigVlanAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigVlanAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigVlanAllocationStrategy`: CreateLogicalNetworkProfileVlanAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigVlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigVlanAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateLogicalNetworkProfileVlanAllocationStrategy201Response**](CreateLogicalNetworkProfileVlanAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigVniAllocationStrategies

> PaginatedVniAllocationStrategy GetLogicalNetworkConfigVniAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigVniAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigVniAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigVniAllocationStrategies`: PaginatedVniAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigVniAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigVniAllocationStrategiesRequest struct via the builder pattern


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

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigVniAllocationStrategy

> CreateLogicalNetworkProfileVniAllocationStrategy201Response GetLogicalNetworkConfigVniAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigVniAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigVniAllocationStrategy`: CreateLogicalNetworkProfileVniAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigVniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigVniAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateLogicalNetworkProfileVniAllocationStrategy201Response**](CreateLogicalNetworkProfileVniAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworks

> PaginatedLogicalNetwork GetLogicalNetworks(ctx).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Logical Networks

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
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> -1           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.           <p>              <b>Format: </b> filter.kind={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.kind=$not:$like:John Doe&filter.kind=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterFabricId := []string{"Inner_example"} // []string | Filter by fabricId query param.           <p>              <b>Format: </b> filter.fabricId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.fabricId=$not:$like:John Doe&filter.fabricId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$null</li></ul> (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.           <p>              <b>Format: </b> filter.serviceStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serviceStatus=$not:$like:John Doe&filter.serviceStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.           <p>              <b>Format: </b> filter.config.deployStatus={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployStatus=$not:$like:John Doe&filter.config.deployStatus=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.           <p>              <b>Format: </b> filter.config.deployType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.config.deployType=$not:$like:John Doe&filter.config.deployType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,name           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>name</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworks(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworks`: PaginatedLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; -1           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterKind** | **[]string** | Filter by kind query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.kind&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.kind&#x3D;$not:$like:John Doe&amp;filter.kind&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterFabricId** | **[]string** | Filter by fabricId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.fabricId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.fabricId&#x3D;$not:$like:John Doe&amp;filter.fabricId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt;&lt;/ul&gt; | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serviceStatus&#x3D;$not:$like:John Doe&amp;filter.serviceStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployStatus&#x3D;$not:$like:John Doe&amp;filter.config.deployStatus&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.config.deployType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.config.deployType&#x3D;$not:$like:John Doe&amp;filter.config.deployType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**PaginatedLogicalNetwork**](PaginatedLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv4SubnetAllocationStrategy_request{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv4SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy

> CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileIpv6SubnetAllocationStrategy_request{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest(createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy`: CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest** | [**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response**](CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkConfigVlanAllocationStrategy

> CreateLogicalNetworkProfileVlanAllocationStrategy201Response ReplaceLogicalNetworkConfigVlanAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileVlanAllocationStrategyRequest(createLogicalNetworkProfileVlanAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileVlanAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileVlanAllocationStrategy_request{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateLogicalNetworkProfileVlanAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigVlanAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileVlanAllocationStrategyRequest(createLogicalNetworkProfileVlanAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigVlanAllocationStrategy`: CreateLogicalNetworkProfileVlanAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.ReplaceLogicalNetworkConfigVlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkConfigVlanAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileVlanAllocationStrategyRequest** | [**CreateLogicalNetworkProfileVlanAllocationStrategyRequest**](CreateLogicalNetworkProfileVlanAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileVlanAllocationStrategy201Response**](CreateLogicalNetworkProfileVlanAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkConfigVniAllocationStrategy

> CreateLogicalNetworkProfileVniAllocationStrategy201Response ReplaceLogicalNetworkConfigVniAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createLogicalNetworkProfileVniAllocationStrategyRequest := openapiclient.createLogicalNetworkProfileVniAllocationStrategy_request{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateLogicalNetworkProfileVniAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigVniAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateLogicalNetworkProfileVniAllocationStrategyRequest(createLogicalNetworkProfileVniAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigVniAllocationStrategy`: CreateLogicalNetworkProfileVniAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.ReplaceLogicalNetworkConfigVniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkConfigVniAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 
 **createLogicalNetworkProfileVniAllocationStrategyRequest** | [**CreateLogicalNetworkProfileVniAllocationStrategyRequest**](CreateLogicalNetworkProfileVniAllocationStrategyRequest.md) |  | 

### Return type

[**CreateLogicalNetworkProfileVniAllocationStrategy201Response**](CreateLogicalNetworkProfileVniAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalNetwork

> CreateLogicalNetwork201Response UpdateLogicalNetwork(ctx, id).IfMatch(ifMatch).UpdateLogicalNetwork(updateLogicalNetwork).Execute()

Update Logical Network

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
	ifMatch := "ifMatch_example" // string | Entity tag
	updateLogicalNetwork := *openapiclient.NewUpdateLogicalNetwork() // UpdateLogicalNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.UpdateLogicalNetwork(context.Background(), id).IfMatch(ifMatch).UpdateLogicalNetwork(updateLogicalNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.UpdateLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalNetwork`: CreateLogicalNetwork201Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.UpdateLogicalNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogicalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **updateLogicalNetwork** | [**UpdateLogicalNetwork**](UpdateLogicalNetwork.md) |  | 

### Return type

[**CreateLogicalNetwork201Response**](CreateLogicalNetwork201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalNetworkConfig

> GetLogicalNetworkConfig200Response UpdateLogicalNetworkConfig(ctx, id).IfMatch(ifMatch).UpdateLogicalNetworkConfig(updateLogicalNetworkConfig).Execute()

Update Logical Network config

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
	ifMatch := "ifMatch_example" // string | Entity tag
	updateLogicalNetworkConfig := *openapiclient.NewUpdateLogicalNetworkConfig(map[string]interface{}(123), map[string]interface{}(123), map[string]interface{}(123), map[string]interface{}(123)) // UpdateLogicalNetworkConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.UpdateLogicalNetworkConfig(context.Background(), id).IfMatch(ifMatch).UpdateLogicalNetworkConfig(updateLogicalNetworkConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.UpdateLogicalNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalNetworkConfig`: GetLogicalNetworkConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.UpdateLogicalNetworkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogicalNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **updateLogicalNetworkConfig** | [**UpdateLogicalNetworkConfig**](UpdateLogicalNetworkConfig.md) |  | 

### Return type

[**GetLogicalNetworkConfig200Response**](GetLogicalNetworkConfig200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

