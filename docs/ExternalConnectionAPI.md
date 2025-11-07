# \ExternalConnectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalConnection**](ExternalConnectionAPI.md#CreateExternalConnection) | **Post** /api/v2/external-connections | Create a new external connection
[**CreateExternalConnectionInterface**](ExternalConnectionAPI.md#CreateExternalConnectionInterface) | **Post** /api/v2/external-connections/{externalConnectionId}/interfaces | Create a new external connection interface
[**CreateExternalConnectionLogicalNetwork**](ExternalConnectionAPI.md#CreateExternalConnectionLogicalNetwork) | **Post** /api/v2/external-connections/{externalConnectionId}/logical-networks | Create a new external connection logical network
[**DeleteExternalConnection**](ExternalConnectionAPI.md#DeleteExternalConnection) | **Delete** /api/v2/external-connections/{externalConnectionId} | Delete external connection
[**DeleteExternalConnectionInterface**](ExternalConnectionAPI.md#DeleteExternalConnectionInterface) | **Delete** /api/v2/external-connections/{externalConnectionId}/interfaces/{externalConnectionInterfaceId} | Delete external connection interface
[**DeleteExternalConnectionLogicalNetwork**](ExternalConnectionAPI.md#DeleteExternalConnectionLogicalNetwork) | **Delete** /api/v2/external-connections/{externalConnectionId}/logical-networks/{id} | Start deletion of an external connection logical network. Only those in pending_activation state will be deleted.
[**GetExternalConnectionById**](ExternalConnectionAPI.md#GetExternalConnectionById) | **Get** /api/v2/external-connections/{externalConnectionId} | Get external connection details
[**GetExternalConnectionInterfaceById**](ExternalConnectionAPI.md#GetExternalConnectionInterfaceById) | **Get** /api/v2/external-connections/{externalConnectionId}/interfaces/{externalConnectionInterfaceId} | Get external connection interface details
[**GetExternalConnectionInterfaces**](ExternalConnectionAPI.md#GetExternalConnectionInterfaces) | **Get** /api/v2/external-connections/{externalConnectionId}/interfaces | List external connection interfaces
[**GetExternalConnectionLogicalNetworkById**](ExternalConnectionAPI.md#GetExternalConnectionLogicalNetworkById) | **Get** /api/v2/external-connections/{externalConnectionId}/logical-networks/{id} | Get an external connection logical network by ID
[**GetExternalConnectionLogicalNetworks**](ExternalConnectionAPI.md#GetExternalConnectionLogicalNetworks) | **Get** /api/v2/external-connections/{externalConnectionId}/logical-networks | Get all logical networks of an external connection
[**GetExternalConnections**](ExternalConnectionAPI.md#GetExternalConnections) | **Get** /api/v2/external-connections | List external connections
[**GetNetworkDeviceInterfacesAndExternalConnections**](ExternalConnectionAPI.md#GetNetworkDeviceInterfacesAndExternalConnections) | **Get** /api/v2/external-connections/network-devices/{networkDeviceId}/interfaces | List network device interfaces associated with external connections
[**UpdateExternalConnection**](ExternalConnectionAPI.md#UpdateExternalConnection) | **Patch** /api/v2/external-connections/{externalConnectionId} | Update external connection
[**UpdateExternalConnectionInterface**](ExternalConnectionAPI.md#UpdateExternalConnectionInterface) | **Patch** /api/v2/external-connections/{externalConnectionId}/interfaces/{externalConnectionInterfaceId} | Update external connection interface



## CreateExternalConnection

> ExternalConnection CreateExternalConnection(ctx).CreateExternalConnection(createExternalConnection).Execute()

Create a new external connection



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
	createExternalConnection := *openapiclient.NewCreateExternalConnection("external-connection-1", "Server 1", int32(123)) // CreateExternalConnection | The external connection to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.CreateExternalConnection(context.Background()).CreateExternalConnection(createExternalConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.CreateExternalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalConnection`: ExternalConnection
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.CreateExternalConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExternalConnection** | [**CreateExternalConnection**](CreateExternalConnection.md) | The external connection to create | 

### Return type

[**ExternalConnection**](ExternalConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalConnectionInterface

> ExternalConnectionInterface CreateExternalConnectionInterface(ctx, externalConnectionId).CreateExternalConnectionInterface(createExternalConnectionInterface).Execute()

Create a new external connection interface



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
	externalConnectionId := int32(56) // int32 | 
	createExternalConnectionInterface := *openapiclient.NewCreateExternalConnectionInterface(float32(1)) // CreateExternalConnectionInterface | The external connection interface to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.CreateExternalConnectionInterface(context.Background(), externalConnectionId).CreateExternalConnectionInterface(createExternalConnectionInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.CreateExternalConnectionInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalConnectionInterface`: ExternalConnectionInterface
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.CreateExternalConnectionInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalConnectionInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createExternalConnectionInterface** | [**CreateExternalConnectionInterface**](CreateExternalConnectionInterface.md) | The external connection interface to create | 

### Return type

[**ExternalConnectionInterface**](ExternalConnectionInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalConnectionLogicalNetwork

> ExternalConnectionLogicalNetwork CreateExternalConnectionLogicalNetwork(ctx, externalConnectionId).CreateExternalConnectionLogicalNetwork(createExternalConnectionLogicalNetwork).Execute()

Create a new external connection logical network

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
	externalConnectionId := int32(56) // int32 | The id of the external connection
	createExternalConnectionLogicalNetwork := *openapiclient.NewCreateExternalConnectionLogicalNetwork(int32(123)) // CreateExternalConnectionLogicalNetwork | The external connection logical network to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.CreateExternalConnectionLogicalNetwork(context.Background(), externalConnectionId).CreateExternalConnectionLogicalNetwork(createExternalConnectionLogicalNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.CreateExternalConnectionLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalConnectionLogicalNetwork`: ExternalConnectionLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.CreateExternalConnectionLogicalNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** | The id of the external connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalConnectionLogicalNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createExternalConnectionLogicalNetwork** | [**CreateExternalConnectionLogicalNetwork**](CreateExternalConnectionLogicalNetwork.md) | The external connection logical network to create | 

### Return type

[**ExternalConnectionLogicalNetwork**](ExternalConnectionLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalConnection

> DeleteExternalConnection(ctx, externalConnectionId).Execute()

Delete external connection



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
	externalConnectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalConnectionAPI.DeleteExternalConnection(context.Background(), externalConnectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.DeleteExternalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalConnectionRequest struct via the builder pattern


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


## DeleteExternalConnectionInterface

> DeleteExternalConnectionInterface(ctx, externalConnectionId, externalConnectionInterfaceId).IfMatch(ifMatch).Execute()

Delete external connection interface



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
	externalConnectionId := int32(56) // int32 | 
	externalConnectionInterfaceId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalConnectionAPI.DeleteExternalConnectionInterface(context.Background(), externalConnectionId, externalConnectionInterfaceId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.DeleteExternalConnectionInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 
**externalConnectionInterfaceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalConnectionInterfaceRequest struct via the builder pattern


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


## DeleteExternalConnectionLogicalNetwork

> DeleteExternalConnectionLogicalNetwork(ctx, externalConnectionId, id).Execute()

Start deletion of an external connection logical network. Only those in pending_activation state will be deleted.

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
	externalConnectionId := int32(56) // int32 | The ID of the external connection
	id := int32(56) // int32 | The ID of the external connection logical network to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalConnectionAPI.DeleteExternalConnectionLogicalNetwork(context.Background(), externalConnectionId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.DeleteExternalConnectionLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** | The ID of the external connection | 
**id** | **int32** | The ID of the external connection logical network to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalConnectionLogicalNetworkRequest struct via the builder pattern


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


## GetExternalConnectionById

> ExternalConnection GetExternalConnectionById(ctx, externalConnectionId).Execute()

Get external connection details



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
	externalConnectionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnectionById(context.Background(), externalConnectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.GetExternalConnectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionById`: ExternalConnection
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.GetExternalConnectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalConnection**](ExternalConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionInterfaceById

> ExternalConnectionInterface GetExternalConnectionInterfaceById(ctx, externalConnectionId, externalConnectionInterfaceId).Execute()

Get external connection interface details



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
	externalConnectionId := int32(56) // int32 | 
	externalConnectionInterfaceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnectionInterfaceById(context.Background(), externalConnectionId, externalConnectionInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.GetExternalConnectionInterfaceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionInterfaceById`: ExternalConnectionInterface
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.GetExternalConnectionInterfaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 
**externalConnectionInterfaceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionInterfaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExternalConnectionInterface**](ExternalConnectionInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionInterfaces

> ExternalConnectionInterfacePaginatedList GetExternalConnectionInterfaces(ctx, externalConnectionId).Page(page).Limit(limit).FilterId(filterId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()

List external connection interfaces



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
	externalConnectionId := int32(56) // int32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.  **Format:** filter.createdAt={$not}:OPERATION:VALUE    **Example:** filter.createdAt=$btw:John Doe&filter.createdAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.  **Format:** filter.updatedAt={$not}:OPERATION:VALUE    **Example:** filter.updatedAt=$btw:John Doe&filter.updatedAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnectionInterfaces(context.Background(), externalConnectionId).Page(page).Limit(limit).FilterId(filterId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.GetExternalConnectionInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionInterfaces`: ExternalConnectionInterfacePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.GetExternalConnectionInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.  **Format:** filter.createdAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdAt&#x3D;$btw:John Doe&amp;filter.createdAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.  **Format:** filter.updatedAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updatedAt&#x3D;$btw:John Doe&amp;filter.updatedAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  | 

### Return type

[**ExternalConnectionInterfacePaginatedList**](ExternalConnectionInterfacePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionLogicalNetworkById

> ExternalConnectionLogicalNetwork GetExternalConnectionLogicalNetworkById(ctx, externalConnectionId, id).Execute()

Get an external connection logical network by ID



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
	externalConnectionId := int32(56) // int32 | The id of the external connection
	id := int32(56) // int32 | The ID of the external connection logical network to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnectionLogicalNetworkById(context.Background(), externalConnectionId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.GetExternalConnectionLogicalNetworkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionLogicalNetworkById`: ExternalConnectionLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.GetExternalConnectionLogicalNetworkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** | The id of the external connection | 
**id** | **int32** | The ID of the external connection logical network to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionLogicalNetworkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExternalConnectionLogicalNetwork**](ExternalConnectionLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalConnectionLogicalNetworks

> ExternalConnectionLogicalNetworkPaginatedList GetExternalConnectionLogicalNetworks(ctx, externalConnectionId).Page(page).Limit(limit).FilterId(filterId).FilterExternalConnectionId(filterExternalConnectionId).FilterLogicalNetworkId(filterLogicalNetworkId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()

Get all logical networks of an external connection



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
	externalConnectionId := int32(56) // int32 | The id of the external connection
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterExternalConnectionId := []string{"Inner_example"} // []string | Filter by externalConnectionId query param.  **Format:** filter.externalConnectionId={$not}:OPERATION:VALUE    **Example:** filter.externalConnectionId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLogicalNetworkId := []string{"Inner_example"} // []string | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId={$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.  **Format:** filter.createdAt={$not}:OPERATION:VALUE    **Example:** filter.createdAt=$btw:John Doe&filter.createdAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.  **Format:** filter.updatedAt={$not}:OPERATION:VALUE    **Example:** filter.updatedAt=$btw:John Doe&filter.updatedAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  - externalConnectionId  - logicalNetworkId  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnectionLogicalNetworks(context.Background(), externalConnectionId).Page(page).Limit(limit).FilterId(filterId).FilterExternalConnectionId(filterExternalConnectionId).FilterLogicalNetworkId(filterLogicalNetworkId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.GetExternalConnectionLogicalNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnectionLogicalNetworks`: ExternalConnectionLogicalNetworkPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.GetExternalConnectionLogicalNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** | The id of the external connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionLogicalNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterExternalConnectionId** | **[]string** | Filter by externalConnectionId query param.  **Format:** filter.externalConnectionId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.externalConnectionId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLogicalNetworkId** | **[]string** | Filter by logicalNetworkId query param.  **Format:** filter.logicalNetworkId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.logicalNetworkId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
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


## GetExternalConnections

> ExternalConnectionPaginatedList GetExternalConnections(ctx).Page(page).Limit(limit).FilterId(filterId).FilterFabricId(filterFabricId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List external connections



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
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterFabricId := []string{"Inner_example"} // []string | Filter by fabricId query param.  **Format:** filter.fabricId={$not}:OPERATION:VALUE    **Example:** filter.fabricId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.  **Format:** filter.createdAt={$not}:OPERATION:VALUE    **Example:** filter.createdAt=$btw:John Doe&filter.createdAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.  **Format:** filter.updatedAt={$not}:OPERATION:VALUE    **Example:** filter.updatedAt=$btw:John Doe&filter.updatedAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdAt:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdAt  - updatedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnections(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterFabricId(filterFabricId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.GetExternalConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalConnections`: ExternalConnectionPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.GetExternalConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
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


## GetNetworkDeviceInterfacesAndExternalConnections

> NetworkDeviceExternalConnectionInterfaces GetNetworkDeviceInterfacesAndExternalConnections(ctx, networkDeviceId).Execute()

List network device interfaces associated with external connections

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
	networkDeviceId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetNetworkDeviceInterfacesAndExternalConnections(context.Background(), networkDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.GetNetworkDeviceInterfacesAndExternalConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDeviceInterfacesAndExternalConnections`: NetworkDeviceExternalConnectionInterfaces
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.GetNetworkDeviceInterfacesAndExternalConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDeviceId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDeviceInterfacesAndExternalConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkDeviceExternalConnectionInterfaces**](NetworkDeviceExternalConnectionInterfaces.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalConnection

> ExternalConnection UpdateExternalConnection(ctx, externalConnectionId).UpdateExternalConnection(updateExternalConnection).IfMatch(ifMatch).Execute()

Update external connection



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
	externalConnectionId := int32(56) // int32 | 
	updateExternalConnection := *openapiclient.NewUpdateExternalConnection() // UpdateExternalConnection | The external connection configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.UpdateExternalConnection(context.Background(), externalConnectionId).UpdateExternalConnection(updateExternalConnection).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.UpdateExternalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalConnection`: ExternalConnection
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.UpdateExternalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalConnection** | [**UpdateExternalConnection**](UpdateExternalConnection.md) | The external connection configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ExternalConnection**](ExternalConnection.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalConnectionInterface

> ExternalConnectionInterface UpdateExternalConnectionInterface(ctx, externalConnectionId, externalConnectionInterfaceId).UpdateExternalConnectionInterface(updateExternalConnectionInterface).IfMatch(ifMatch).Execute()

Update external connection interface



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
	externalConnectionId := int32(56) // int32 | 
	externalConnectionInterfaceId := int32(56) // int32 | 
	updateExternalConnectionInterface := *openapiclient.NewUpdateExternalConnectionInterface(float32(1)) // UpdateExternalConnectionInterface | The external connection interface configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.UpdateExternalConnectionInterface(context.Background(), externalConnectionId, externalConnectionInterfaceId).UpdateExternalConnectionInterface(updateExternalConnectionInterface).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.UpdateExternalConnectionInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalConnectionInterface`: ExternalConnectionInterface
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.UpdateExternalConnectionInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **int32** |  | 
**externalConnectionInterfaceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalConnectionInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateExternalConnectionInterface** | [**UpdateExternalConnectionInterface**](UpdateExternalConnectionInterface.md) | The external connection interface configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ExternalConnectionInterface**](ExternalConnectionInterface.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

