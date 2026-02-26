# \LogicalNetworkAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyProfilesToLogicalNetworkConfig**](LogicalNetworkAPI.md#ApplyProfilesToLogicalNetworkConfig) | **Post** /api/v2/logical-networks/{id}/config/actions/apply-profiles | Apply profiles to Logical Network config
[**CreateLogicalNetwork**](LogicalNetworkAPI.md#CreateLogicalNetwork) | **Post** /api/v2/logical-networks | Create a Logical Network.
[**CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies | Create Ipv4 Subnet allocation strategy.
[**CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies | Create Ipv6 Subnet allocation strategy.
[**CreateLogicalNetworkConfigPkeyAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigPkeyAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/pkey/pkey-allocation-strategies | Create Pkey allocation strategy.
[**CreateLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigVlanAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies | Create Vlan allocation strategy.
[**CreateLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigVniAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies | Create Vni allocation strategy.
[**CreateLogicalNetworkConfigZoneAllocationStrategy**](LogicalNetworkAPI.md#CreateLogicalNetworkConfigZoneAllocationStrategy) | **Post** /api/v2/logical-networks/{id}/config/zone/zone-allocation-strategies | Create Zone allocation strategy.
[**CreateLogicalNetworkFromProfile**](LogicalNetworkAPI.md#CreateLogicalNetworkFromProfile) | **Post** /api/v2/logical-networks/actions/create-from-profile | Create a Logical Network from a profile.
[**DeleteLogicalNetwork**](LogicalNetworkAPI.md#DeleteLogicalNetwork) | **Delete** /api/v2/logical-networks/{id} | Delete a Logical Network.
[**DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv4 Subnet allocation strategy.
[**DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Delete Ipv6 Subnet allocation strategy.
[**DeleteLogicalNetworkConfigPkeyAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigPkeyAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/pkey/pkey-allocation-strategies/{allocationStrategyId} | Delete Pkey allocation strategy.
[**DeleteLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigVlanAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies/{allocationStrategyId} | Delete Vlan allocation strategy.
[**DeleteLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigVniAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies/{allocationStrategyId} | Delete Vni allocation strategy.
[**DeleteLogicalNetworkConfigZoneAllocationStrategy**](LogicalNetworkAPI.md#DeleteLogicalNetworkConfigZoneAllocationStrategy) | **Delete** /api/v2/logical-networks/{id}/config/zone/zone-allocation-strategies/{allocationStrategyId} | Delete Zone allocation strategy.
[**DetachExternalConnectionLogicalNetwork**](LogicalNetworkAPI.md#DetachExternalConnectionLogicalNetwork) | **Delete** /api/v2/logical-networks/{id}/external-connections/{externalConnectionId} | Detaches an external connection from a logical network
[**GetLogicalNetwork**](LogicalNetworkAPI.md#GetLogicalNetwork) | **Get** /api/v2/logical-networks/{id} | Get a Logical Network.
[**GetLogicalNetworkAttachedExternalConnectionLogicalNetworks**](LogicalNetworkAPI.md#GetLogicalNetworkAttachedExternalConnectionLogicalNetworks) | **Get** /api/v2/logical-networks/{id}/external-connection-logical-networks | Get all external connection logical networks
[**GetLogicalNetworkAttachedExternalConnections**](LogicalNetworkAPI.md#GetLogicalNetworkAttachedExternalConnections) | **Get** /api/v2/logical-networks/{id}/external-connections | List external connections attached to a logical network
[**GetLogicalNetworkConfig**](LogicalNetworkAPI.md#GetLogicalNetworkConfig) | **Get** /api/v2/logical-networks/{id}/config | Get the config for a Logical Network.
[**GetLogicalNetworkConfigIpv4SubnetAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv4SubnetAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies | Get all Ipv4 Subnet allocation strategies.
[**GetLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv4 Subnet allocation strategy.
[**GetLogicalNetworkConfigIpv6SubnetAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv6SubnetAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies | Get all Ipv6 Subnet allocation strategies.
[**GetLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Get a Ipv6 Subnet allocation strategy.
[**GetLogicalNetworkConfigPkeyAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigPkeyAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/pkey/pkey-allocation-strategies | Get all Pkey allocation strategies.
[**GetLogicalNetworkConfigPkeyAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigPkeyAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/pkey/pkey-allocation-strategies/{allocationStrategyId} | Get a Pkey allocation strategy.
[**GetLogicalNetworkConfigVlanAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVlanAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies | Get all Vlan allocation strategies.
[**GetLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVlanAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies/{allocationStrategyId} | Get a Vlan allocation strategy.
[**GetLogicalNetworkConfigVniAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVniAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies | Get all Vni allocation strategies.
[**GetLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigVniAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies/{allocationStrategyId} | Get a Vni allocation strategy.
[**GetLogicalNetworkConfigZoneAllocationStrategies**](LogicalNetworkAPI.md#GetLogicalNetworkConfigZoneAllocationStrategies) | **Get** /api/v2/logical-networks/{id}/config/zone/zone-allocation-strategies | Get all Zone allocation strategies.
[**GetLogicalNetworkConfigZoneAllocationStrategy**](LogicalNetworkAPI.md#GetLogicalNetworkConfigZoneAllocationStrategy) | **Get** /api/v2/logical-networks/{id}/config/zone/zone-allocation-strategies/{allocationStrategyId} | Get a Zone allocation strategy.
[**GetLogicalNetworks**](LogicalNetworkAPI.md#GetLogicalNetworks) | **Get** /api/v2/logical-networks | Get all Logical Networks
[**ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/ipv4/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv4 Subnet allocation strategy
[**ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/ipv6/subnet-allocation-strategies/{allocationStrategyId} | Replace Ipv6 Subnet allocation strategy
[**ReplaceLogicalNetworkConfigPkeyAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigPkeyAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/pkey/pkey-allocation-strategies/{allocationStrategyId} | Replace Pkey allocation strategy
[**ReplaceLogicalNetworkConfigVlanAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigVlanAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/vlan/vlan-allocation-strategies/{allocationStrategyId} | Replace Vlan allocation strategy
[**ReplaceLogicalNetworkConfigVniAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigVniAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/vxlan/vni-allocation-strategies/{allocationStrategyId} | Replace Vni allocation strategy
[**ReplaceLogicalNetworkConfigZoneAllocationStrategy**](LogicalNetworkAPI.md#ReplaceLogicalNetworkConfigZoneAllocationStrategy) | **Put** /api/v2/logical-networks/{id}/config/zone/zone-allocation-strategies/{allocationStrategyId} | Replace Zone allocation strategy
[**UpdateLogicalNetwork**](LogicalNetworkAPI.md#UpdateLogicalNetwork) | **Patch** /api/v2/logical-networks/{id} | Update Logical Network
[**UpdateLogicalNetworkConfig**](LogicalNetworkAPI.md#UpdateLogicalNetworkConfig) | **Patch** /api/v2/logical-networks/{id}/config | Update Logical Network config



## ApplyProfilesToLogicalNetworkConfig

> LogicalNetworkConfig ApplyProfilesToLogicalNetworkConfig(ctx, id).IfMatch(ifMatch).ApplyProfilesToLogicalNetworkConfig(applyProfilesToLogicalNetworkConfig).Execute()

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
	// response from `ApplyProfilesToLogicalNetworkConfig`: LogicalNetworkConfig
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

[**LogicalNetworkConfig**](LogicalNetworkConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetwork

> LogicalNetwork CreateLogicalNetwork(ctx).CreateLogicalNetwork(createLogicalNetwork).Execute()

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
	createLogicalNetwork := *openapiclient.NewCreateLogicalNetwork(openapiclient.LogicalNetworkKind("vlan"), int32(123)) // CreateLogicalNetwork | The logical network to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetwork(context.Background()).CreateLogicalNetwork(createLogicalNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetwork`: LogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogicalNetwork** | [**CreateLogicalNetwork**](CreateLogicalNetwork.md) | The logical network to create | 

### Return type

[**LogicalNetwork**](LogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy

> Ipv4SubnetAllocationStrategy CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).Execute()

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
	createIpv4SubnetAllocationStrategy := openapiclient.CreateIpv4SubnetAllocationStrategy{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv4SubnetAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigIpv4SubnetAllocationStrategy`: Ipv4SubnetAllocationStrategy
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
 **createIpv4SubnetAllocationStrategy** | [**CreateIpv4SubnetAllocationStrategy**](CreateIpv4SubnetAllocationStrategy.md) |  | 

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


## CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy

> Ipv6SubnetAllocationStrategy CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).Execute()

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
	createIpv6SubnetAllocationStrategy := openapiclient.CreateIpv6SubnetAllocationStrategy{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv6SubnetAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigIpv6SubnetAllocationStrategy`: Ipv6SubnetAllocationStrategy
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
 **createIpv6SubnetAllocationStrategy** | [**CreateIpv6SubnetAllocationStrategy**](CreateIpv6SubnetAllocationStrategy.md) |  | 

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


## CreateLogicalNetworkConfigPkeyAllocationStrategy

> PkeyAllocationStrategy CreateLogicalNetworkConfigPkeyAllocationStrategy(ctx, id).IfMatch(ifMatch).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createPkeyAllocationStrategy := openapiclient.CreatePkeyAllocationStrategy{CreateAutoPkeyAllocationStrategy: openapiclient.NewCreateAutoPkeyAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreatePkeyAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigPkeyAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigPkeyAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigPkeyAllocationStrategy`: PkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetworkConfigPkeyAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkConfigPkeyAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **createPkeyAllocationStrategy** | [**CreatePkeyAllocationStrategy**](CreatePkeyAllocationStrategy.md) |  | 

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


## CreateLogicalNetworkConfigVlanAllocationStrategy

> VlanAllocationStrategy CreateLogicalNetworkConfigVlanAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()

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
	createVlanAllocationStrategy := openapiclient.CreateVlanAllocationStrategy{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVlanAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigVlanAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigVlanAllocationStrategy`: VlanAllocationStrategy
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
 **createVlanAllocationStrategy** | [**CreateVlanAllocationStrategy**](CreateVlanAllocationStrategy.md) |  | 

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


## CreateLogicalNetworkConfigVniAllocationStrategy

> VniAllocationStrategy CreateLogicalNetworkConfigVniAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()

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
	createVniAllocationStrategy := openapiclient.CreateVniAllocationStrategy{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVniAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigVniAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigVniAllocationStrategy`: VniAllocationStrategy
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
 **createVniAllocationStrategy** | [**CreateVniAllocationStrategy**](CreateVniAllocationStrategy.md) |  | 

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


## CreateLogicalNetworkConfigZoneAllocationStrategy

> ZoneAllocationStrategy CreateLogicalNetworkConfigZoneAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateZoneAllocationStrategy(createZoneAllocationStrategy).Execute()

Create Zone allocation strategy.

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
	createZoneAllocationStrategy := openapiclient.CreateZoneAllocationStrategy{CreateAutoZoneAllocationStrategy: openapiclient.NewCreateAutoZoneAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateZoneAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.CreateLogicalNetworkConfigZoneAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateZoneAllocationStrategy(createZoneAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.CreateLogicalNetworkConfigZoneAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalNetworkConfigZoneAllocationStrategy`: ZoneAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.CreateLogicalNetworkConfigZoneAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalNetworkConfigZoneAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **createZoneAllocationStrategy** | [**CreateZoneAllocationStrategy**](CreateZoneAllocationStrategy.md) |  | 

### Return type

[**ZoneAllocationStrategy**](ZoneAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalNetworkFromProfile

> LogicalNetwork CreateLogicalNetworkFromProfile(ctx).CreateLogicalNetworkFromProfile(createLogicalNetworkFromProfile).Execute()

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
	// response from `CreateLogicalNetworkFromProfile`: LogicalNetwork
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

[**LogicalNetwork**](LogicalNetwork.md)

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


## DeleteLogicalNetworkConfigPkeyAllocationStrategy

> DeleteLogicalNetworkConfigPkeyAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkAPI.DeleteLogicalNetworkConfigPkeyAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DeleteLogicalNetworkConfigPkeyAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkConfigPkeyAllocationStrategyRequest struct via the builder pattern


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


## DeleteLogicalNetworkConfigZoneAllocationStrategy

> DeleteLogicalNetworkConfigZoneAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete Zone allocation strategy.

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
	r, err := apiClient.LogicalNetworkAPI.DeleteLogicalNetworkConfigZoneAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DeleteLogicalNetworkConfigZoneAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogicalNetworkConfigZoneAllocationStrategyRequest struct via the builder pattern


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


## DetachExternalConnectionLogicalNetwork

> DetachExternalConnectionLogicalNetwork(ctx, id, externalConnectionId).Execute()

Detaches an external connection from a logical network



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
	externalConnectionId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalNetworkAPI.DetachExternalConnectionLogicalNetwork(context.Background(), id, externalConnectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.DetachExternalConnectionLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**externalConnectionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachExternalConnectionLogicalNetworkRequest struct via the builder pattern


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


## GetLogicalNetwork

> LogicalNetwork GetLogicalNetwork(ctx, id).Execute()

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
	// response from `GetLogicalNetwork`: LogicalNetwork
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

[**LogicalNetwork**](LogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkAttachedExternalConnectionLogicalNetworks

> ExternalConnectionLogicalNetworkPaginatedList GetLogicalNetworkAttachedExternalConnectionLogicalNetworks(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterExternalConnectionId(filterExternalConnectionId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()

Get all external connection logical networks



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
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterExternalConnectionId := []string{"Inner_example"} // []string | Filter by externalConnectionId query param.  **Format:** filter.externalConnectionId={$not}:OPERATION:VALUE    **Example:** filter.externalConnectionId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.  **Format:** filter.createdAt={$not}:OPERATION:VALUE    **Example:** filter.createdAt=$btw:John Doe&filter.createdAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.  **Format:** filter.updatedAt={$not}:OPERATION:VALUE    **Example:** filter.updatedAt=$btw:John Doe&filter.updatedAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  - externalConnectionId  - logicalNetworkId  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkAttachedExternalConnectionLogicalNetworks(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterExternalConnectionId(filterExternalConnectionId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkAttachedExternalConnectionLogicalNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkAttachedExternalConnectionLogicalNetworks`: ExternalConnectionLogicalNetworkPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkAttachedExternalConnectionLogicalNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkAttachedExternalConnectionLogicalNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterExternalConnectionId** | **[]string** | Filter by externalConnectionId query param.  **Format:** filter.externalConnectionId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.externalConnectionId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.  **Format:** filter.createdAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdAt&#x3D;$btw:John Doe&amp;filter.createdAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.  **Format:** filter.updatedAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updatedAt&#x3D;$btw:John Doe&amp;filter.updatedAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  - externalConnectionId  - logicalNetworkId  | 

### Return type

[**ExternalConnectionLogicalNetworkPaginatedList**](ExternalConnectionLogicalNetworkPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkAttachedExternalConnections

> ExternalConnectionPaginatedList GetLogicalNetworkAttachedExternalConnections(ctx, id).Page(page).Limit(limit).FilterFabricId(filterFabricId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List external connections attached to a logical network



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
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterFabricId := []string{"Inner_example"} // []string | Filter by fabricId query param.  **Format:** filter.fabricId={$not}:OPERATION:VALUE    **Example:** filter.fabricId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.  **Format:** filter.createdAt={$not}:OPERATION:VALUE    **Example:** filter.createdAt=$btw:John Doe&filter.createdAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.  **Format:** filter.updatedAt={$not}:OPERATION:VALUE    **Example:** filter.updatedAt=$btw:John Doe&filter.updatedAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkAttachedExternalConnections(context.Background(), id).Page(page).Limit(limit).FilterFabricId(filterFabricId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkAttachedExternalConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkAttachedExternalConnections`: ExternalConnectionPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkAttachedExternalConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkAttachedExternalConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterFabricId** | **[]string** | Filter by fabricId query param.  **Format:** filter.fabricId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.  **Format:** filter.createdAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdAt&#x3D;$btw:John Doe&amp;filter.createdAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.  **Format:** filter.updatedAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updatedAt&#x3D;$btw:John Doe&amp;filter.updatedAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  | 

### Return type

[**ExternalConnectionPaginatedList**](ExternalConnectionPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfig

> LogicalNetworkConfig GetLogicalNetworkConfig(ctx, id).Execute()

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
	// response from `GetLogicalNetworkConfig`: LogicalNetworkConfig
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

[**LogicalNetworkConfig**](LogicalNetworkConfig.md)

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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


## GetLogicalNetworkConfigIpv4SubnetAllocationStrategy

> Ipv4SubnetAllocationStrategy GetLogicalNetworkConfigIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkConfigIpv4SubnetAllocationStrategy`: Ipv4SubnetAllocationStrategy
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

[**Ipv4SubnetAllocationStrategy**](Ipv4SubnetAllocationStrategy.md)

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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


## GetLogicalNetworkConfigIpv6SubnetAllocationStrategy

> Ipv6SubnetAllocationStrategy GetLogicalNetworkConfigIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkConfigIpv6SubnetAllocationStrategy`: Ipv6SubnetAllocationStrategy
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

[**Ipv6SubnetAllocationStrategy**](Ipv6SubnetAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigPkeyAllocationStrategies

> PaginatedPkeyAllocationStrategy GetLogicalNetworkConfigPkeyAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigPkeyAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigPkeyAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigPkeyAllocationStrategies`: PaginatedPkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigPkeyAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigPkeyAllocationStrategiesRequest struct via the builder pattern


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


## GetLogicalNetworkConfigPkeyAllocationStrategy

> PkeyAllocationStrategy GetLogicalNetworkConfigPkeyAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigPkeyAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigPkeyAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigPkeyAllocationStrategy`: PkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigPkeyAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigPkeyAllocationStrategyRequest struct via the builder pattern


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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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


## GetLogicalNetworkConfigVlanAllocationStrategy

> VlanAllocationStrategy GetLogicalNetworkConfigVlanAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkConfigVlanAllocationStrategy`: VlanAllocationStrategy
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

[**VlanAllocationStrategy**](VlanAllocationStrategy.md)

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$btw:John Doe&filter.kind=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)

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


## GetLogicalNetworkConfigVniAllocationStrategy

> VniAllocationStrategy GetLogicalNetworkConfigVniAllocationStrategy(ctx, id, allocationStrategyId).Execute()

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
	// response from `GetLogicalNetworkConfigVniAllocationStrategy`: VniAllocationStrategy
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

[**VniAllocationStrategy**](VniAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigZoneAllocationStrategies

> PaginatedZoneAllocationStrategy GetLogicalNetworkConfigZoneAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all Zone allocation strategies.

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigZoneAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigZoneAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigZoneAllocationStrategies`: PaginatedZoneAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigZoneAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigZoneAllocationStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 

### Return type

[**PaginatedZoneAllocationStrategy**](PaginatedZoneAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworkConfigZoneAllocationStrategy

> ZoneAllocationStrategy GetLogicalNetworkConfigZoneAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a Zone allocation strategy.

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
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworkConfigZoneAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworkConfigZoneAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworkConfigZoneAllocationStrategy`: ZoneAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworkConfigZoneAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworkConfigZoneAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ZoneAllocationStrategy**](ZoneAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalNetworks

> PaginatedLogicalNetworkList GetLogicalNetworks(ctx).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterRouteDomainId(filterRouteDomainId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** -1   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe&filter.id=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterKind := []string{"Inner_example"} // []string | Filter by kind query param.  **Format:** filter.kind={$not}:OPERATION:VALUE    **Example:** filter.kind=$eq:John Doe&filter.kind=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterFabricId := []string{"Inner_example"} // []string | Filter by fabricId query param.  **Format:** filter.fabricId={$not}:OPERATION:VALUE    **Example:** filter.fabricId=$eq:John Doe&filter.fabricId=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$eq:John Doe&filter.infrastructureId=$in:John Doe  **Available Operations** - $eq  - $in  - $not  - $null  - $and  - $or (optional)
	filterServiceStatus := []string{"Inner_example"} // []string | Filter by serviceStatus query param.  **Format:** filter.serviceStatus={$not}:OPERATION:VALUE    **Example:** filter.serviceStatus=$eq:John Doe&filter.serviceStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployStatus := []string{"Inner_example"} // []string | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus={$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus=$eq:John Doe&filter.config.deployStatus=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterConfigDeployType := []string{"Inner_example"} // []string | Filter by config.deployType query param.  **Format:** filter.config.deployType={$not}:OPERATION:VALUE    **Example:** filter.config.deployType=$eq:John Doe&filter.config.deployType=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterRouteDomainId := []string{"Inner_example"} // []string | Filter by routeDomainId query param.  **Format:** filter.routeDomainId={$not}:OPERATION:VALUE    **Example:** filter.routeDomainId=$eq:John Doe&filter.routeDomainId=$in:John Doe  **Available Operations** - $eq  - $in  - $not  - $null  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.GetLogicalNetworks(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).FilterRouteDomainId(filterRouteDomainId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.GetLogicalNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalNetworks`: PaginatedLogicalNetworkList
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.GetLogicalNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** -1   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$eq:John Doe&amp;filter.kind&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterFabricId** | **[]string** | Filter by fabricId query param.  **Format:** filter.fabricId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.fabricId&#x3D;$eq:John Doe&amp;filter.fabricId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$eq:John Doe&amp;filter.infrastructureId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $not  - $null  - $and  - $or | 
 **filterServiceStatus** | **[]string** | Filter by serviceStatus query param.  **Format:** filter.serviceStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serviceStatus&#x3D;$eq:John Doe&amp;filter.serviceStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployStatus** | **[]string** | Filter by config.deployStatus query param.  **Format:** filter.config.deployStatus&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployStatus&#x3D;$eq:John Doe&amp;filter.config.deployStatus&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterConfigDeployType** | **[]string** | Filter by config.deployType query param.  **Format:** filter.config.deployType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.config.deployType&#x3D;$eq:John Doe&amp;filter.config.deployType&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterRouteDomainId** | **[]string** | Filter by routeDomainId query param.  **Format:** filter.routeDomainId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.routeDomainId&#x3D;$eq:John Doe&amp;filter.routeDomainId&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $not  - $null  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  | 

### Return type

[**PaginatedLogicalNetworkList**](PaginatedLogicalNetworkList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy

> Ipv4SubnetAllocationStrategy ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).Execute()

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
	createIpv4SubnetAllocationStrategy := openapiclient.CreateIpv4SubnetAllocationStrategy{CreateAutoIpv4SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv4SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv4SubnetAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateIpv4SubnetAllocationStrategy(createIpv4SubnetAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigIpv4SubnetAllocationStrategy`: Ipv4SubnetAllocationStrategy
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
 **createIpv4SubnetAllocationStrategy** | [**CreateIpv4SubnetAllocationStrategy**](CreateIpv4SubnetAllocationStrategy.md) |  | 

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


## ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy

> Ipv6SubnetAllocationStrategy ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).Execute()

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
	createIpv6SubnetAllocationStrategy := openapiclient.CreateIpv6SubnetAllocationStrategy{CreateAutoIpv6SubnetAllocationStrategy: openapiclient.NewCreateAutoIpv6SubnetAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), []int32{int32(123)}, int32(123))} // CreateIpv6SubnetAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateIpv6SubnetAllocationStrategy(createIpv6SubnetAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigIpv6SubnetAllocationStrategy`: Ipv6SubnetAllocationStrategy
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
 **createIpv6SubnetAllocationStrategy** | [**CreateIpv6SubnetAllocationStrategy**](CreateIpv6SubnetAllocationStrategy.md) |  | 

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


## ReplaceLogicalNetworkConfigPkeyAllocationStrategy

> PkeyAllocationStrategy ReplaceLogicalNetworkConfigPkeyAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).Execute()

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
	ifMatch := "ifMatch_example" // string | Entity tag
	createPkeyAllocationStrategy := openapiclient.CreatePkeyAllocationStrategy{CreateAutoPkeyAllocationStrategy: openapiclient.NewCreateAutoPkeyAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreatePkeyAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigPkeyAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreatePkeyAllocationStrategy(createPkeyAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigPkeyAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigPkeyAllocationStrategy`: PkeyAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.ReplaceLogicalNetworkConfigPkeyAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkConfigPkeyAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 
 **createPkeyAllocationStrategy** | [**CreatePkeyAllocationStrategy**](CreatePkeyAllocationStrategy.md) |  | 

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


## ReplaceLogicalNetworkConfigVlanAllocationStrategy

> VlanAllocationStrategy ReplaceLogicalNetworkConfigVlanAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()

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
	createVlanAllocationStrategy := openapiclient.CreateVlanAllocationStrategy{CreateAutoVlanAllocationStrategy: openapiclient.NewCreateAutoVlanAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVlanAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigVlanAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigVlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigVlanAllocationStrategy`: VlanAllocationStrategy
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
 **createVlanAllocationStrategy** | [**CreateVlanAllocationStrategy**](CreateVlanAllocationStrategy.md) |  | 

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


## ReplaceLogicalNetworkConfigVniAllocationStrategy

> VniAllocationStrategy ReplaceLogicalNetworkConfigVniAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()

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
	createVniAllocationStrategy := openapiclient.CreateVniAllocationStrategy{CreateAutoVniAllocationStrategy: openapiclient.NewCreateAutoVniAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateVniAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigVniAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigVniAllocationStrategy`: VniAllocationStrategy
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
 **createVniAllocationStrategy** | [**CreateVniAllocationStrategy**](CreateVniAllocationStrategy.md) |  | 

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


## ReplaceLogicalNetworkConfigZoneAllocationStrategy

> ZoneAllocationStrategy ReplaceLogicalNetworkConfigZoneAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateZoneAllocationStrategy(createZoneAllocationStrategy).Execute()

Replace Zone allocation strategy

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
	createZoneAllocationStrategy := openapiclient.CreateZoneAllocationStrategy{CreateAutoZoneAllocationStrategy: openapiclient.NewCreateAutoZoneAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)))} // CreateZoneAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.ReplaceLogicalNetworkConfigZoneAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateZoneAllocationStrategy(createZoneAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.ReplaceLogicalNetworkConfigZoneAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceLogicalNetworkConfigZoneAllocationStrategy`: ZoneAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `LogicalNetworkAPI.ReplaceLogicalNetworkConfigZoneAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogicalNetworkConfigZoneAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 
 **createZoneAllocationStrategy** | [**CreateZoneAllocationStrategy**](CreateZoneAllocationStrategy.md) |  | 

### Return type

[**ZoneAllocationStrategy**](ZoneAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalNetwork

> LogicalNetwork UpdateLogicalNetwork(ctx, id).IfMatch(ifMatch).UpdateLogicalNetwork(updateLogicalNetwork).Execute()

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
	// response from `UpdateLogicalNetwork`: LogicalNetwork
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

[**LogicalNetwork**](LogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalNetworkConfig

> LogicalNetworkConfig UpdateLogicalNetworkConfig(ctx, id).IfMatch(ifMatch).UpdateLogicalNetworkConfigGlobalSettings(updateLogicalNetworkConfigGlobalSettings).Execute()

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
	updateLogicalNetworkConfigGlobalSettings := *openapiclient.NewUpdateLogicalNetworkConfigGlobalSettings() // UpdateLogicalNetworkConfigGlobalSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalNetworkAPI.UpdateLogicalNetworkConfig(context.Background(), id).IfMatch(ifMatch).UpdateLogicalNetworkConfigGlobalSettings(updateLogicalNetworkConfigGlobalSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalNetworkAPI.UpdateLogicalNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalNetworkConfig`: LogicalNetworkConfig
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
 **updateLogicalNetworkConfigGlobalSettings** | [**UpdateLogicalNetworkConfigGlobalSettings**](UpdateLogicalNetworkConfigGlobalSettings.md) |  | 

### Return type

[**LogicalNetworkConfig**](LogicalNetworkConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

