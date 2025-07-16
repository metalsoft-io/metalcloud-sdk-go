# \ExternalConnectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalConnection**](ExternalConnectionAPI.md#CreateExternalConnection) | **Post** /api/v2/external-connections | Create a new external connection
[**CreateExternalConnectionInterface**](ExternalConnectionAPI.md#CreateExternalConnectionInterface) | **Post** /api/v2/external-connections/{externalConnectionId}/interfaces | Create a new external connection interface
[**CreateExternalConnectionLogicalNetwork**](ExternalConnectionAPI.md#CreateExternalConnectionLogicalNetwork) | **Post** /api/v2/external-connections/{externalConnectionId}/logical-networks | Create a new external connection logical network
[**DeleteExternalConnection**](ExternalConnectionAPI.md#DeleteExternalConnection) | **Delete** /api/v2/external-connections/{externalConnectionId} | Delete external connection
[**DeleteExternalConnectionInterface**](ExternalConnectionAPI.md#DeleteExternalConnectionInterface) | **Delete** /api/v2/external-connections/{externalConnectionId}/interfaces/{externalConnectionInterfaceId} | Delete external connection interface
[**DeleteExternalConnectionLogicalNetwork**](ExternalConnectionAPI.md#DeleteExternalConnectionLogicalNetwork) | **Delete** /api/v2/external-connections/{externalConnectionId}/logical-networks/{id} | Start deletion of an external connection logical network
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
	createExternalConnection := *openapiclient.NewCreateExternalConnection("external-connection-1", "Server 1", float32(123)) // CreateExternalConnection | The external connection to create

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
	createExternalConnectionLogicalNetwork := *openapiclient.NewCreateExternalConnectionLogicalNetwork(float32(123)) // CreateExternalConnectionLogicalNetwork | The external connection logical network to create

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

> ExternalConnectionLogicalNetwork DeleteExternalConnectionLogicalNetwork(ctx, externalConnectionId, id).Execute()

Start deletion of an external connection logical network

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
	resp, r, err := apiClient.ExternalConnectionAPI.DeleteExternalConnectionLogicalNetwork(context.Background(), externalConnectionId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalConnectionAPI.DeleteExternalConnectionLogicalNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalConnectionLogicalNetwork`: ExternalConnectionLogicalNetwork
	fmt.Fprintf(os.Stdout, "Response from `ExternalConnectionAPI.DeleteExternalConnectionLogicalNetwork`: %v\n", resp)
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

[**ExternalConnectionLogicalNetwork**](ExternalConnectionLogicalNetwork.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.           <p>              <b>Format: </b> filter.createdAt={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.createdAt=$not:$like:John Doe&filter.createdAt=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.           <p>              <b>Format: </b> filter.updatedAt={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.updatedAt=$not:$like:John Doe&filter.updatedAt=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>createdAt</li> <li>updatedAt</li></ul>        (optional)

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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.createdAt&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.createdAt&#x3D;$not:$like:John Doe&amp;filter.createdAt&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.updatedAt&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.updatedAt&#x3D;$not:$like:John Doe&amp;filter.updatedAt&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;createdAt&lt;/li&gt; &lt;li&gt;updatedAt&lt;/li&gt;&lt;/ul&gt;        | 

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

> ExternalConnectionLogicalNetworkPaginatedList GetExternalConnectionLogicalNetworks(ctx, externalConnectionId).Page(page).Limit(limit).FilterId(filterId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.           <p>              <b>Format: </b> filter.createdAt={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.createdAt=$not:$like:John Doe&filter.createdAt=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.           <p>              <b>Format: </b> filter.updatedAt={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.updatedAt=$not:$like:John Doe&filter.updatedAt=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>createdAt</li> <li>updatedAt</li></ul>        (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnectionLogicalNetworks(context.Background(), externalConnectionId).Page(page).Limit(limit).FilterId(filterId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Execute()
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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.createdAt&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.createdAt&#x3D;$not:$like:John Doe&amp;filter.createdAt&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.updatedAt&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.updatedAt&#x3D;$not:$like:John Doe&amp;filter.updatedAt&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;createdAt&lt;/li&gt; &lt;li&gt;updatedAt&lt;/li&gt;&lt;/ul&gt;        | 

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

> ExternalConnectionPaginatedList GetExternalConnections(ctx).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterExternalId(filterExternalId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterExternalId := []string{"Inner_example"} // []string | Filter by externalId query param.           <p>              <b>Format: </b> filter.externalId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.externalId=$not:$like:John Doe&filter.externalId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.           <p>              <b>Format: </b> filter.createdAt={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.createdAt=$not:$like:John Doe&filter.createdAt=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.           <p>              <b>Format: </b> filter.updatedAt={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.updatedAt=$not:$like:John Doe&filter.updatedAt=like:John           </p>           <h4>Available Operations</h4><ul><li>$and</li> <li>$or</li> <li>$not</li> <li>$eq</li> <li>$gt</li> <li>$gte</li> <li>$in</li> <li>$null</li> <li>$lt</li> <li>$lte</li> <li>$btw</li> <li>$ilike</li> <li>$sw</li> <li>$contains</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>createdAt</li> <li>updatedAt</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> name,label           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>name</li> <li>label</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalConnectionAPI.GetExternalConnections(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterSiteId(filterSiteId).FilterExternalId(filterExternalId).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterExternalId** | **[]string** | Filter by externalId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.externalId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.externalId&#x3D;$not:$like:John Doe&amp;filter.externalId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.createdAt&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.createdAt&#x3D;$not:$like:John Doe&amp;filter.createdAt&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.updatedAt&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.updatedAt&#x3D;$not:$like:John Doe&amp;filter.updatedAt&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;createdAt&lt;/li&gt; &lt;li&gt;updatedAt&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name,label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt; &lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;          | 

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

