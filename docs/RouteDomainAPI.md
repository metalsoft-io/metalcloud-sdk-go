# \RouteDomainAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteDomain**](RouteDomainAPI.md#CreateRouteDomain) | **Post** /api/v2/route-domains | Create a Route Domain.
[**CreateRouteDomainConfigL3VlanAllocationStrategy**](RouteDomainAPI.md#CreateRouteDomainConfigL3VlanAllocationStrategy) | **Post** /api/v2/route-domains/{id}/config/l3-vlan-allocation-strategies | Create L3 Vlan allocation strategy.
[**CreateRouteDomainConfigVniAllocationStrategy**](RouteDomainAPI.md#CreateRouteDomainConfigVniAllocationStrategy) | **Post** /api/v2/route-domains/{id}/config/l3-vni-allocation-strategies | Create Vni allocation strategy.
[**CreateRouteDomainConfigVrfAllocationStrategy**](RouteDomainAPI.md#CreateRouteDomainConfigVrfAllocationStrategy) | **Post** /api/v2/route-domains/{id}/config/vrf-allocation-strategies | Create Vrf allocation strategy.
[**DeleteRouteDomain**](RouteDomainAPI.md#DeleteRouteDomain) | **Delete** /api/v2/route-domains/{id} | Delete a Route Domain.
[**DeleteRouteDomainConfigL3VlanAllocationStrategy**](RouteDomainAPI.md#DeleteRouteDomainConfigL3VlanAllocationStrategy) | **Delete** /api/v2/route-domains/{id}/config/l3-vlan-allocation-strategies/{allocationStrategyId} | Delete L3 Vlan allocation strategy.
[**DeleteRouteDomainConfigL3VniAllocationStrategy**](RouteDomainAPI.md#DeleteRouteDomainConfigL3VniAllocationStrategy) | **Delete** /api/v2/route-domains/{id}/config/l3-vni-allocation-strategies/{allocationStrategyId} | Delete L3 Vni allocation strategy.
[**DeleteRouteDomainConfigVrfAllocationStrategy**](RouteDomainAPI.md#DeleteRouteDomainConfigVrfAllocationStrategy) | **Delete** /api/v2/route-domains/{id}/config/vrf-allocation-strategies/{allocationStrategyId} | Delete Vrf allocation strategy.
[**GetRouteDomain**](RouteDomainAPI.md#GetRouteDomain) | **Get** /api/v2/route-domains/{id} | Get a Route Domain.
[**GetRouteDomainConfig**](RouteDomainAPI.md#GetRouteDomainConfig) | **Get** /api/v2/route-domains/{id}/config | Get the config for a Route Domain.
[**GetRouteDomainConfigL3VlanAllocationStrategies**](RouteDomainAPI.md#GetRouteDomainConfigL3VlanAllocationStrategies) | **Get** /api/v2/route-domains/{id}/config/l3-vlan-allocation-strategies | Get all L3 Vlan allocation strategies.
[**GetRouteDomainConfigL3VlanAllocationStrategy**](RouteDomainAPI.md#GetRouteDomainConfigL3VlanAllocationStrategy) | **Get** /api/v2/route-domains/{id}/config/l3-vlan-allocation-strategies/{allocationStrategyId} | Get a L3 Vlan allocation strategy.
[**GetRouteDomainConfigL3VniAllocationStrategies**](RouteDomainAPI.md#GetRouteDomainConfigL3VniAllocationStrategies) | **Get** /api/v2/route-domains/{id}/config/l3-vni-allocation-strategies | Get all L3 Vni allocation strategies.
[**GetRouteDomainConfigL3VniAllocationStrategy**](RouteDomainAPI.md#GetRouteDomainConfigL3VniAllocationStrategy) | **Get** /api/v2/route-domains/{id}/config/l3-vni-allocation-strategies/{allocationStrategyId} | Get a L3 Vni allocation strategy.
[**GetRouteDomainConfigVrfAllocationStrategies**](RouteDomainAPI.md#GetRouteDomainConfigVrfAllocationStrategies) | **Get** /api/v2/route-domains/{id}/config/vrf-allocation-strategies | Get all Vrf allocation strategies.
[**GetRouteDomainConfigVrfAllocationStrategy**](RouteDomainAPI.md#GetRouteDomainConfigVrfAllocationStrategy) | **Get** /api/v2/route-domains/{id}/config/vrf-allocation-strategies/{allocationStrategyId} | Get a Vrf allocation strategy.
[**GetRouteDomains**](RouteDomainAPI.md#GetRouteDomains) | **Get** /api/v2/route-domains | Get all Route Domains
[**ReplaceRouteDomainConfigL3VlanAllocationStrategy**](RouteDomainAPI.md#ReplaceRouteDomainConfigL3VlanAllocationStrategy) | **Put** /api/v2/route-domains/{id}/config/l3-vlan-allocation-strategies/{allocationStrategyId} | Replace L3 Vlan allocation strategy
[**ReplaceRouteDomainConfigL3VniAllocationStrategy**](RouteDomainAPI.md#ReplaceRouteDomainConfigL3VniAllocationStrategy) | **Put** /api/v2/route-domains/{id}/config/l3-vni-allocation-strategies/{allocationStrategyId} | Replace L3 Vni allocation strategy
[**ReplaceRouteDomainConfigVrfAllocationStrategy**](RouteDomainAPI.md#ReplaceRouteDomainConfigVrfAllocationStrategy) | **Put** /api/v2/route-domains/{id}/config/vrf-allocation-strategies/{allocationStrategyId} | Replace Vrf allocation strategy
[**UpdateRouteDomain**](RouteDomainAPI.md#UpdateRouteDomain) | **Patch** /api/v2/route-domains/{id} | Update Route Domain
[**UpdateRouteDomainConfig**](RouteDomainAPI.md#UpdateRouteDomainConfig) | **Patch** /api/v2/route-domains/{id}/config | Update Route Domain config



## CreateRouteDomain

> RouteDomain CreateRouteDomain(ctx).CreateRouteDomain(createRouteDomain).Execute()

Create a Route Domain.

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
	createRouteDomain := *openapiclient.NewCreateRouteDomain(openapiclient.RouteDomainKind("vrf_lite"), []openapiclient.CreateVrfAllocationStrategy{openapiclient.CreateVrfAllocationStrategy{CreateManualVrfAllocationStrategy: openapiclient.NewCreateManualVrfAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), "Name_example")}}) // CreateRouteDomain | The route domain to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteDomainAPI.CreateRouteDomain(context.Background()).CreateRouteDomain(createRouteDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.CreateRouteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteDomain`: RouteDomain
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.CreateRouteDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteDomain** | [**CreateRouteDomain**](CreateRouteDomain.md) | The route domain to create | 

### Return type

[**RouteDomain**](RouteDomain.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRouteDomainConfigL3VlanAllocationStrategy

> VlanAllocationStrategy CreateRouteDomainConfigL3VlanAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()

Create L3 Vlan allocation strategy.

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
	resp, r, err := apiClient.RouteDomainAPI.CreateRouteDomainConfigL3VlanAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.CreateRouteDomainConfigL3VlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteDomainConfigL3VlanAllocationStrategy`: VlanAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.CreateRouteDomainConfigL3VlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteDomainConfigL3VlanAllocationStrategyRequest struct via the builder pattern


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


## CreateRouteDomainConfigVniAllocationStrategy

> VniAllocationStrategy CreateRouteDomainConfigVniAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()

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
	resp, r, err := apiClient.RouteDomainAPI.CreateRouteDomainConfigVniAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.CreateRouteDomainConfigVniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteDomainConfigVniAllocationStrategy`: VniAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.CreateRouteDomainConfigVniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteDomainConfigVniAllocationStrategyRequest struct via the builder pattern


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


## CreateRouteDomainConfigVrfAllocationStrategy

> ManualVrfAllocationStrategy CreateRouteDomainConfigVrfAllocationStrategy(ctx, id).IfMatch(ifMatch).CreateManualVrfAllocationStrategy(createManualVrfAllocationStrategy).Execute()

Create Vrf allocation strategy.

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
	createManualVrfAllocationStrategy := *openapiclient.NewCreateManualVrfAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), "Name_example") // CreateManualVrfAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteDomainAPI.CreateRouteDomainConfigVrfAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreateManualVrfAllocationStrategy(createManualVrfAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.CreateRouteDomainConfigVrfAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteDomainConfigVrfAllocationStrategy`: ManualVrfAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.CreateRouteDomainConfigVrfAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteDomainConfigVrfAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **createManualVrfAllocationStrategy** | [**CreateManualVrfAllocationStrategy**](CreateManualVrfAllocationStrategy.md) |  | 

### Return type

[**ManualVrfAllocationStrategy**](ManualVrfAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteDomain

> DeleteRouteDomain(ctx, id).IfMatch(ifMatch).Execute()

Delete a Route Domain.

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
	r, err := apiClient.RouteDomainAPI.DeleteRouteDomain(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.DeleteRouteDomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRouteDomainRequest struct via the builder pattern


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


## DeleteRouteDomainConfigL3VlanAllocationStrategy

> DeleteRouteDomainConfigL3VlanAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete L3 Vlan allocation strategy.

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
	r, err := apiClient.RouteDomainAPI.DeleteRouteDomainConfigL3VlanAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.DeleteRouteDomainConfigL3VlanAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRouteDomainConfigL3VlanAllocationStrategyRequest struct via the builder pattern


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


## DeleteRouteDomainConfigL3VniAllocationStrategy

> DeleteRouteDomainConfigL3VniAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete L3 Vni allocation strategy.

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
	r, err := apiClient.RouteDomainAPI.DeleteRouteDomainConfigL3VniAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.DeleteRouteDomainConfigL3VniAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRouteDomainConfigL3VniAllocationStrategyRequest struct via the builder pattern


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


## DeleteRouteDomainConfigVrfAllocationStrategy

> DeleteRouteDomainConfigVrfAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).Execute()

Delete Vrf allocation strategy.

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
	r, err := apiClient.RouteDomainAPI.DeleteRouteDomainConfigVrfAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.DeleteRouteDomainConfigVrfAllocationStrategy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRouteDomainConfigVrfAllocationStrategyRequest struct via the builder pattern


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


## GetRouteDomain

> RouteDomain GetRouteDomain(ctx, id).Execute()

Get a Route Domain.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomain`: RouteDomain
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteDomain**](RouteDomain.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteDomainConfig

> RouteDomainConfig GetRouteDomainConfig(ctx, id).Execute()

Get the config for a Route Domain.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomainConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomainConfig`: RouteDomainConfig
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteDomainConfig**](RouteDomainConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteDomainConfigL3VlanAllocationStrategies

> PaginatedVlanAllocationStrategy GetRouteDomainConfigL3VlanAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all L3 Vlan allocation strategies.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomainConfigL3VlanAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomainConfigL3VlanAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomainConfigL3VlanAllocationStrategies`: PaginatedVlanAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomainConfigL3VlanAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainConfigL3VlanAllocationStrategiesRequest struct via the builder pattern


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


## GetRouteDomainConfigL3VlanAllocationStrategy

> VlanAllocationStrategy GetRouteDomainConfigL3VlanAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a L3 Vlan allocation strategy.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomainConfigL3VlanAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomainConfigL3VlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomainConfigL3VlanAllocationStrategy`: VlanAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomainConfigL3VlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainConfigL3VlanAllocationStrategyRequest struct via the builder pattern


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


## GetRouteDomainConfigL3VniAllocationStrategies

> PaginatedVniAllocationStrategy GetRouteDomainConfigL3VniAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all L3 Vni allocation strategies.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomainConfigL3VniAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomainConfigL3VniAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomainConfigL3VniAllocationStrategies`: PaginatedVniAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomainConfigL3VniAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainConfigL3VniAllocationStrategiesRequest struct via the builder pattern


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


## GetRouteDomainConfigL3VniAllocationStrategy

> VniAllocationStrategy GetRouteDomainConfigL3VniAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a L3 Vni allocation strategy.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomainConfigL3VniAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomainConfigL3VniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomainConfigL3VniAllocationStrategy`: VniAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomainConfigL3VniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainConfigL3VniAllocationStrategyRequest struct via the builder pattern


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


## GetRouteDomainConfigVrfAllocationStrategies

> PaginatedVrfAllocationStrategy GetRouteDomainConfigVrfAllocationStrategies(ctx, id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()

Get all Vrf allocation strategies.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomainConfigVrfAllocationStrategies(context.Background(), id).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomainConfigVrfAllocationStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomainConfigVrfAllocationStrategies`: PaginatedVrfAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomainConfigVrfAllocationStrategies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainConfigVrfAllocationStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 100    **Max Value:** 1000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterKind** | **[]string** | Filter by kind query param.  **Format:** filter.kind&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.kind&#x3D;$btw:John Doe&amp;filter.kind&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 

### Return type

[**PaginatedVrfAllocationStrategy**](PaginatedVrfAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteDomainConfigVrfAllocationStrategy

> ManualVrfAllocationStrategy GetRouteDomainConfigVrfAllocationStrategy(ctx, id, allocationStrategyId).Execute()

Get a Vrf allocation strategy.

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
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomainConfigVrfAllocationStrategy(context.Background(), id, allocationStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomainConfigVrfAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomainConfigVrfAllocationStrategy`: ManualVrfAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomainConfigVrfAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainConfigVrfAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManualVrfAllocationStrategy**](ManualVrfAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteDomains

> PaginatedRouteDomainList GetRouteDomains(ctx).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Route Domains

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
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteDomainAPI.GetRouteDomains(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterKind(filterKind).FilterLabel(filterLabel).FilterName(filterName).FilterFabricId(filterFabricId).FilterInfrastructureId(filterInfrastructureId).FilterServiceStatus(filterServiceStatus).FilterConfigDeployStatus(filterConfigDeployStatus).FilterConfigDeployType(filterConfigDeployType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.GetRouteDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteDomains`: PaginatedRouteDomainList
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.GetRouteDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteDomainsRequest struct via the builder pattern


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
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  | 

### Return type

[**PaginatedRouteDomainList**](PaginatedRouteDomainList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRouteDomainConfigL3VlanAllocationStrategy

> VlanAllocationStrategy ReplaceRouteDomainConfigL3VlanAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()

Replace L3 Vlan allocation strategy

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
	resp, r, err := apiClient.RouteDomainAPI.ReplaceRouteDomainConfigL3VlanAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateVlanAllocationStrategy(createVlanAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.ReplaceRouteDomainConfigL3VlanAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRouteDomainConfigL3VlanAllocationStrategy`: VlanAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.ReplaceRouteDomainConfigL3VlanAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRouteDomainConfigL3VlanAllocationStrategyRequest struct via the builder pattern


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


## ReplaceRouteDomainConfigL3VniAllocationStrategy

> VniAllocationStrategy ReplaceRouteDomainConfigL3VniAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()

Replace L3 Vni allocation strategy

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
	resp, r, err := apiClient.RouteDomainAPI.ReplaceRouteDomainConfigL3VniAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateVniAllocationStrategy(createVniAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.ReplaceRouteDomainConfigL3VniAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRouteDomainConfigL3VniAllocationStrategy`: VniAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.ReplaceRouteDomainConfigL3VniAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRouteDomainConfigL3VniAllocationStrategyRequest struct via the builder pattern


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


## ReplaceRouteDomainConfigVrfAllocationStrategy

> ManualVrfAllocationStrategy ReplaceRouteDomainConfigVrfAllocationStrategy(ctx, id, allocationStrategyId).IfMatch(ifMatch).CreateManualVrfAllocationStrategy(createManualVrfAllocationStrategy).Execute()

Replace Vrf allocation strategy

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
	createManualVrfAllocationStrategy := *openapiclient.NewCreateManualVrfAllocationStrategy(openapiclient.AllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global"), float32(123)), "Name_example") // CreateManualVrfAllocationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteDomainAPI.ReplaceRouteDomainConfigVrfAllocationStrategy(context.Background(), id, allocationStrategyId).IfMatch(ifMatch).CreateManualVrfAllocationStrategy(createManualVrfAllocationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.ReplaceRouteDomainConfigVrfAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRouteDomainConfigVrfAllocationStrategy`: ManualVrfAllocationStrategy
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.ReplaceRouteDomainConfigVrfAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**allocationStrategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRouteDomainConfigVrfAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Entity tag | 
 **createManualVrfAllocationStrategy** | [**CreateManualVrfAllocationStrategy**](CreateManualVrfAllocationStrategy.md) |  | 

### Return type

[**ManualVrfAllocationStrategy**](ManualVrfAllocationStrategy.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteDomain

> RouteDomain UpdateRouteDomain(ctx, id).IfMatch(ifMatch).UpdateRouteDomain(updateRouteDomain).Execute()

Update Route Domain

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
	updateRouteDomain := *openapiclient.NewUpdateRouteDomain() // UpdateRouteDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteDomainAPI.UpdateRouteDomain(context.Background(), id).IfMatch(ifMatch).UpdateRouteDomain(updateRouteDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.UpdateRouteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRouteDomain`: RouteDomain
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.UpdateRouteDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **updateRouteDomain** | [**UpdateRouteDomain**](UpdateRouteDomain.md) |  | 

### Return type

[**RouteDomain**](RouteDomain.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteDomainConfig

> RouteDomainConfig UpdateRouteDomainConfig(ctx, id).IfMatch(ifMatch).Body(body).Execute()

Update Route Domain config

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteDomainAPI.UpdateRouteDomainConfig(context.Background(), id).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteDomainAPI.UpdateRouteDomainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRouteDomainConfig`: RouteDomainConfig
	fmt.Fprintf(os.Stdout, "Response from `RouteDomainAPI.UpdateRouteDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**RouteDomainConfig**](RouteDomainConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

