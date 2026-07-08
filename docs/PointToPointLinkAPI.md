# \PointToPointLinkAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePointToPointLink**](PointToPointLinkAPI.md#CreatePointToPointLink) | **Post** /api/v2/point-to-point-links | Create a PointToPointLink between two specific switch interfaces.
[**CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy**](PointToPointLinkAPI.md#CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy) | **Post** /api/v2/point-to-point-links/{id}/config/ipv4/subnet-allocation-strategies | Add an IPv4 subnet allocation strategy to a PointToPointLink candidate.
[**CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy**](PointToPointLinkAPI.md#CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy) | **Post** /api/v2/point-to-point-links/{id}/config/ipv6/subnet-allocation-strategies | Add an IPv6 subnet allocation strategy to a PointToPointLink candidate.
[**CreatePointToPointLinkConfigStaticRoute**](PointToPointLinkAPI.md#CreatePointToPointLinkConfigStaticRoute) | **Post** /api/v2/point-to-point-links/{id}/config/{family}/static-routes | Add a static route to a PointToPointLink candidate (per address family).
[**DeletePointToPointLink**](PointToPointLinkAPI.md#DeletePointToPointLink) | **Delete** /api/v2/point-to-point-links/{id} | Mark a PointToPointLink for deletion (sets candidate deployType&#x3D;DELETE).
[**DeletePointToPointLinkConfigIpv4SubnetAllocationStrategy**](PointToPointLinkAPI.md#DeletePointToPointLinkConfigIpv4SubnetAllocationStrategy) | **Delete** /api/v2/point-to-point-links/{id}/config/ipv4/subnet-allocation-strategies/{strategyId} | Delete an IPv4 subnet allocation strategy from the candidate.
[**DeletePointToPointLinkConfigIpv6SubnetAllocationStrategy**](PointToPointLinkAPI.md#DeletePointToPointLinkConfigIpv6SubnetAllocationStrategy) | **Delete** /api/v2/point-to-point-links/{id}/config/ipv6/subnet-allocation-strategies/{strategyId} | Delete an IPv6 subnet allocation strategy from the candidate.
[**DeletePointToPointLinkConfigStaticRoute**](PointToPointLinkAPI.md#DeletePointToPointLinkConfigStaticRoute) | **Delete** /api/v2/point-to-point-links/{id}/config/{family}/static-routes/{routeId} | Delete a static route from the candidate.
[**GetPointToPointLink**](PointToPointLinkAPI.md#GetPointToPointLink) | **Get** /api/v2/point-to-point-links/{id} | Get a PointToPointLink by id.
[**GetPointToPointLinkConfig**](PointToPointLinkAPI.md#GetPointToPointLinkConfig) | **Get** /api/v2/point-to-point-links/{id}/config | Get the candidate (proposed) configuration of a PointToPointLink.
[**GetPointToPointLinkConfigIpv4SubnetAllocationStrategies**](PointToPointLinkAPI.md#GetPointToPointLinkConfigIpv4SubnetAllocationStrategies) | **Get** /api/v2/point-to-point-links/{id}/config/ipv4/subnet-allocation-strategies | List all IPv4 subnet allocation strategies on a PointToPointLink candidate.
[**GetPointToPointLinkConfigIpv4SubnetAllocationStrategy**](PointToPointLinkAPI.md#GetPointToPointLinkConfigIpv4SubnetAllocationStrategy) | **Get** /api/v2/point-to-point-links/{id}/config/ipv4/subnet-allocation-strategies/{strategyId} | Get one IPv4 subnet allocation strategy.
[**GetPointToPointLinkConfigIpv6SubnetAllocationStrategies**](PointToPointLinkAPI.md#GetPointToPointLinkConfigIpv6SubnetAllocationStrategies) | **Get** /api/v2/point-to-point-links/{id}/config/ipv6/subnet-allocation-strategies | List all IPv6 subnet allocation strategies on a PointToPointLink candidate.
[**GetPointToPointLinkConfigIpv6SubnetAllocationStrategy**](PointToPointLinkAPI.md#GetPointToPointLinkConfigIpv6SubnetAllocationStrategy) | **Get** /api/v2/point-to-point-links/{id}/config/ipv6/subnet-allocation-strategies/{strategyId} | Get one IPv6 subnet allocation strategy.
[**GetPointToPointLinkConfigStaticRoute**](PointToPointLinkAPI.md#GetPointToPointLinkConfigStaticRoute) | **Get** /api/v2/point-to-point-links/{id}/config/{family}/static-routes/{routeId} | Get one static route.
[**GetPointToPointLinkConfigStaticRoutes**](PointToPointLinkAPI.md#GetPointToPointLinkConfigStaticRoutes) | **Get** /api/v2/point-to-point-links/{id}/config/{family}/static-routes | List static routes on a PointToPointLink candidate (per address family).
[**GetPointToPointLinks**](PointToPointLinkAPI.md#GetPointToPointLinks) | **Get** /api/v2/point-to-point-links | List PointToPointLinks, optionally filtered by Route Domain or interface id.
[**UpdatePointToPointLink**](PointToPointLinkAPI.md#UpdatePointToPointLink) | **Patch** /api/v2/point-to-point-links/{id} | Update a PointToPointLink (label, name, annotations, description).
[**UpdatePointToPointLinkConfig**](PointToPointLinkAPI.md#UpdatePointToPointLinkConfig) | **Patch** /api/v2/point-to-point-links/{id}/config | Update the candidate configuration of a PointToPointLink.



## CreatePointToPointLink

> PointToPointLink CreatePointToPointLink(ctx).CreatePointToPointLink(createPointToPointLink).Execute()

Create a PointToPointLink between two specific switch interfaces.

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
	createPointToPointLink := *openapiclient.NewCreatePointToPointLink() // CreatePointToPointLink | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.CreatePointToPointLink(context.Background()).CreatePointToPointLink(createPointToPointLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.CreatePointToPointLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePointToPointLink`: PointToPointLink
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.CreatePointToPointLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePointToPointLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPointToPointLink** | [**CreatePointToPointLink**](CreatePointToPointLink.md) |  | 

### Return type

[**PointToPointLink**](PointToPointLink.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy

> CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy(ctx, id).IfMatch(ifMatch).CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest(createPointToPointLinkConfigIpv4SubnetAllocationStrategyRequest).Execute()

Add an IPv4 subnet allocation strategy to a PointToPointLink candidate.

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
	ifMatch := "ifMatch_example" // string | Candidate revision (entity tag)
	createPointToPointLinkConfigIpv4SubnetAllocationStrategyRequest := openapiclient.createPointToPointLinkConfigIpv4SubnetAllocationStrategy_request{CreateAutoIpv4PointToPointAllocationStrategy: openapiclient.NewCreateAutoIpv4PointToPointAllocationStrategy(openapiclient.PointToPointAllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global")), []int64{int64(123)})} // CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest(createPointToPointLinkConfigIpv4SubnetAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy`: CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Candidate revision (entity tag) | 
 **createPointToPointLinkConfigIpv4SubnetAllocationStrategyRequest** | [**CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest**](CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest.md) |  | 

### Return type

[**CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response**](CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy

> CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy(ctx, id).IfMatch(ifMatch).CreatePointToPointLinkConfigIpv6SubnetAllocationStrategyRequest(createPointToPointLinkConfigIpv6SubnetAllocationStrategyRequest).Execute()

Add an IPv6 subnet allocation strategy to a PointToPointLink candidate.

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
	ifMatch := "ifMatch_example" // string | Candidate revision (entity tag)
	createPointToPointLinkConfigIpv6SubnetAllocationStrategyRequest := openapiclient.createPointToPointLinkConfigIpv6SubnetAllocationStrategy_request{CreateAutoIpv6PointToPointAllocationStrategy: openapiclient.NewCreateAutoIpv6PointToPointAllocationStrategy(openapiclient.PointToPointAllocationStrategyKind("manual"), *openapiclient.NewCreateResourceScope(openapiclient.ResourceScopeKind("global")), []int64{int64(123)})} // CreatePointToPointLinkConfigIpv6SubnetAllocationStrategyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy(context.Background(), id).IfMatch(ifMatch).CreatePointToPointLinkConfigIpv6SubnetAllocationStrategyRequest(createPointToPointLinkConfigIpv6SubnetAllocationStrategyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy`: CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePointToPointLinkConfigIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Candidate revision (entity tag) | 
 **createPointToPointLinkConfigIpv6SubnetAllocationStrategyRequest** | [**CreatePointToPointLinkConfigIpv6SubnetAllocationStrategyRequest**](CreatePointToPointLinkConfigIpv6SubnetAllocationStrategyRequest.md) |  | 

### Return type

[**CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response**](CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePointToPointLinkConfigStaticRoute

> PointToPointStaticRoute CreatePointToPointLinkConfigStaticRoute(ctx, id, family).IfMatch(ifMatch).CreatePointToPointStaticRoute(createPointToPointStaticRoute).Execute()

Add a static route to a PointToPointLink candidate (per address family).

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
	family := "family_example" // string | 
	ifMatch := "ifMatch_example" // string | Candidate revision (entity tag)
	createPointToPointStaticRoute := *openapiclient.NewCreatePointToPointStaticRoute("10.0.0.0/24") // CreatePointToPointStaticRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.CreatePointToPointLinkConfigStaticRoute(context.Background(), id, family).IfMatch(ifMatch).CreatePointToPointStaticRoute(createPointToPointStaticRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.CreatePointToPointLinkConfigStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePointToPointLinkConfigStaticRoute`: PointToPointStaticRoute
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.CreatePointToPointLinkConfigStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**family** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePointToPointLinkConfigStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Candidate revision (entity tag) | 
 **createPointToPointStaticRoute** | [**CreatePointToPointStaticRoute**](CreatePointToPointStaticRoute.md) |  | 

### Return type

[**PointToPointStaticRoute**](PointToPointStaticRoute.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePointToPointLink

> DeletePointToPointLink(ctx, id).IfMatch(ifMatch).Execute()

Mark a PointToPointLink for deletion (sets candidate deployType=DELETE).

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
	ifMatch := "ifMatch_example" // string | Entity tag (revision)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PointToPointLinkAPI.DeletePointToPointLink(context.Background(), id).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.DeletePointToPointLink``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePointToPointLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag (revision) | 

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


## DeletePointToPointLinkConfigIpv4SubnetAllocationStrategy

> DeletePointToPointLinkConfigIpv4SubnetAllocationStrategy(ctx, id, strategyId).IfMatch(ifMatch).Execute()

Delete an IPv4 subnet allocation strategy from the candidate.

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
	strategyId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Candidate revision (entity tag)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PointToPointLinkAPI.DeletePointToPointLinkConfigIpv4SubnetAllocationStrategy(context.Background(), id, strategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.DeletePointToPointLinkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**strategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Candidate revision (entity tag) | 

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


## DeletePointToPointLinkConfigIpv6SubnetAllocationStrategy

> DeletePointToPointLinkConfigIpv6SubnetAllocationStrategy(ctx, id, strategyId).IfMatch(ifMatch).Execute()

Delete an IPv6 subnet allocation strategy from the candidate.

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
	strategyId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Candidate revision (entity tag)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PointToPointLinkAPI.DeletePointToPointLinkConfigIpv6SubnetAllocationStrategy(context.Background(), id, strategyId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.DeletePointToPointLinkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**strategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePointToPointLinkConfigIpv6SubnetAllocationStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Candidate revision (entity tag) | 

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


## DeletePointToPointLinkConfigStaticRoute

> DeletePointToPointLinkConfigStaticRoute(ctx, id, family, routeId).IfMatch(ifMatch).Execute()

Delete a static route from the candidate.

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
	family := "family_example" // string | 
	routeId := float32(8.14) // float32 | 
	ifMatch := "ifMatch_example" // string | Candidate revision (entity tag)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PointToPointLinkAPI.DeletePointToPointLinkConfigStaticRoute(context.Background(), id, family, routeId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.DeletePointToPointLinkConfigStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**family** | **string** |  | 
**routeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePointToPointLinkConfigStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **string** | Candidate revision (entity tag) | 

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


## GetPointToPointLink

> PointToPointLink GetPointToPointLink(ctx, id).Execute()

Get a PointToPointLink by id.

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
	resp, r, err := apiClient.PointToPointLinkAPI.GetPointToPointLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointToPointLink`: PointToPointLink
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.GetPointToPointLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointToPointLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PointToPointLink**](PointToPointLink.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPointToPointLinkConfig

> PointToPointLinkConfig GetPointToPointLinkConfig(ctx, id).Execute()

Get the candidate (proposed) configuration of a PointToPointLink.

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
	resp, r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinkConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointToPointLinkConfig`: PointToPointLinkConfig
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.GetPointToPointLinkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointToPointLinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PointToPointLinkConfig**](PointToPointLinkConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPointToPointLinkConfigIpv4SubnetAllocationStrategies

> GetPointToPointLinkConfigIpv4SubnetAllocationStrategies(ctx, id).Execute()

List all IPv4 subnet allocation strategies on a PointToPointLink candidate.

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
	r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinkConfigIpv4SubnetAllocationStrategies(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinkConfigIpv4SubnetAllocationStrategies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetPointToPointLinkConfigIpv4SubnetAllocationStrategiesRequest struct via the builder pattern


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


## GetPointToPointLinkConfigIpv4SubnetAllocationStrategy

> GetPointToPointLinkConfigIpv4SubnetAllocationStrategy(ctx, id, strategyId).Execute()

Get one IPv4 subnet allocation strategy.

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
	strategyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinkConfigIpv4SubnetAllocationStrategy(context.Background(), id, strategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinkConfigIpv4SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**strategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointToPointLinkConfigIpv4SubnetAllocationStrategyRequest struct via the builder pattern


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


## GetPointToPointLinkConfigIpv6SubnetAllocationStrategies

> GetPointToPointLinkConfigIpv6SubnetAllocationStrategies(ctx, id).Execute()

List all IPv6 subnet allocation strategies on a PointToPointLink candidate.

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
	r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinkConfigIpv6SubnetAllocationStrategies(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinkConfigIpv6SubnetAllocationStrategies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetPointToPointLinkConfigIpv6SubnetAllocationStrategiesRequest struct via the builder pattern


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


## GetPointToPointLinkConfigIpv6SubnetAllocationStrategy

> GetPointToPointLinkConfigIpv6SubnetAllocationStrategy(ctx, id, strategyId).Execute()

Get one IPv6 subnet allocation strategy.

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
	strategyId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinkConfigIpv6SubnetAllocationStrategy(context.Background(), id, strategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinkConfigIpv6SubnetAllocationStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**strategyId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointToPointLinkConfigIpv6SubnetAllocationStrategyRequest struct via the builder pattern


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


## GetPointToPointLinkConfigStaticRoute

> PointToPointStaticRoute GetPointToPointLinkConfigStaticRoute(ctx, id, family, routeId).Execute()

Get one static route.

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
	family := "family_example" // string | 
	routeId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinkConfigStaticRoute(context.Background(), id, family, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinkConfigStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointToPointLinkConfigStaticRoute`: PointToPointStaticRoute
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.GetPointToPointLinkConfigStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**family** | **string** |  | 
**routeId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointToPointLinkConfigStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PointToPointStaticRoute**](PointToPointStaticRoute.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPointToPointLinkConfigStaticRoutes

> []PointToPointStaticRoute GetPointToPointLinkConfigStaticRoutes(ctx, id, family).Execute()

List static routes on a PointToPointLink candidate (per address family).

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
	family := "family_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinkConfigStaticRoutes(context.Background(), id, family).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinkConfigStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointToPointLinkConfigStaticRoutes`: []PointToPointStaticRoute
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.GetPointToPointLinkConfigStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 
**family** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointToPointLinkConfigStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PointToPointStaticRoute**](PointToPointStaticRoute.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPointToPointLinks

> []PointToPointLink GetPointToPointLinks(ctx).RouteDomainId(routeDomainId).InterfaceId(interfaceId).Execute()

List PointToPointLinks, optionally filtered by Route Domain or interface id.

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
	routeDomainId := "routeDomainId_example" // string | Filter by Route Domain id (\"null\" matches links without a Route Domain). (optional)
	interfaceId := "interfaceId_example" // string | Filter by interface id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.GetPointToPointLinks(context.Background()).RouteDomainId(routeDomainId).InterfaceId(interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.GetPointToPointLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointToPointLinks`: []PointToPointLink
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.GetPointToPointLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPointToPointLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeDomainId** | **string** | Filter by Route Domain id (\&quot;null\&quot; matches links without a Route Domain). | 
 **interfaceId** | **string** | Filter by interface id. | 

### Return type

[**[]PointToPointLink**](PointToPointLink.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePointToPointLink

> PointToPointLink UpdatePointToPointLink(ctx, id).IfMatch(ifMatch).UpdatePointToPointLink(updatePointToPointLink).Execute()

Update a PointToPointLink (label, name, annotations, description).

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
	ifMatch := "ifMatch_example" // string | Entity tag (revision)
	updatePointToPointLink := *openapiclient.NewUpdatePointToPointLink() // UpdatePointToPointLink | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.UpdatePointToPointLink(context.Background(), id).IfMatch(ifMatch).UpdatePointToPointLink(updatePointToPointLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.UpdatePointToPointLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePointToPointLink`: PointToPointLink
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.UpdatePointToPointLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePointToPointLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag (revision) | 
 **updatePointToPointLink** | [**UpdatePointToPointLink**](UpdatePointToPointLink.md) |  | 

### Return type

[**PointToPointLink**](PointToPointLink.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePointToPointLinkConfig

> PointToPointLinkConfig UpdatePointToPointLinkConfig(ctx, id).IfMatch(ifMatch).UpdatePointToPointLinkConfig(updatePointToPointLinkConfig).Execute()

Update the candidate configuration of a PointToPointLink.

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
	ifMatch := "ifMatch_example" // string | Entity tag (revision)
	updatePointToPointLinkConfig := *openapiclient.NewUpdatePointToPointLinkConfig() // UpdatePointToPointLinkConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointToPointLinkAPI.UpdatePointToPointLinkConfig(context.Background(), id).IfMatch(ifMatch).UpdatePointToPointLinkConfig(updatePointToPointLinkConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointToPointLinkAPI.UpdatePointToPointLinkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePointToPointLinkConfig`: PointToPointLinkConfig
	fmt.Fprintf(os.Stdout, "Response from `PointToPointLinkAPI.UpdatePointToPointLinkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePointToPointLinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag (revision) | 
 **updatePointToPointLinkConfig** | [**UpdatePointToPointLinkConfig**](UpdatePointToPointLinkConfig.md) |  | 

### Return type

[**PointToPointLinkConfig**](PointToPointLinkConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

