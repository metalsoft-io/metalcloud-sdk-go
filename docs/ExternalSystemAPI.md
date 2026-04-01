# \ExternalSystemAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalSystem**](ExternalSystemAPI.md#CreateExternalSystem) | **Post** /api/v2/external-systems | Create a new external system
[**DeleteExternalSystem**](ExternalSystemAPI.md#DeleteExternalSystem) | **Delete** /api/v2/external-systems/{externalSystemId} | Delete external system
[**GetExternalSystemById**](ExternalSystemAPI.md#GetExternalSystemById) | **Get** /api/v2/external-systems/{externalSystemId} | Get external system details
[**GetExternalSystems**](ExternalSystemAPI.md#GetExternalSystems) | **Get** /api/v2/external-systems | List external systems
[**UpdateExternalSystem**](ExternalSystemAPI.md#UpdateExternalSystem) | **Patch** /api/v2/external-systems/{externalSystemId} | Update external system



## CreateExternalSystem

> ExternalSystem CreateExternalSystem(ctx).CreateExternalSystem(createExternalSystem).Execute()

Create a new external system



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
	createExternalSystem := *openapiclient.NewCreateExternalSystem("external-system-1", "My External System") // CreateExternalSystem | The external system to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalSystemAPI.CreateExternalSystem(context.Background()).CreateExternalSystem(createExternalSystem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalSystemAPI.CreateExternalSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalSystem`: ExternalSystem
	fmt.Fprintf(os.Stdout, "Response from `ExternalSystemAPI.CreateExternalSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExternalSystem** | [**CreateExternalSystem**](CreateExternalSystem.md) | The external system to create | 

### Return type

[**ExternalSystem**](ExternalSystem.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalSystem

> DeleteExternalSystem(ctx, externalSystemId).Execute()

Delete external system



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
	externalSystemId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalSystemAPI.DeleteExternalSystem(context.Background(), externalSystemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalSystemAPI.DeleteExternalSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalSystemId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalSystemRequest struct via the builder pattern


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


## GetExternalSystemById

> ExternalSystem GetExternalSystemById(ctx, externalSystemId).Execute()

Get external system details



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
	externalSystemId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalSystemAPI.GetExternalSystemById(context.Background(), externalSystemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalSystemAPI.GetExternalSystemById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalSystemById`: ExternalSystem
	fmt.Fprintf(os.Stdout, "Response from `ExternalSystemAPI.GetExternalSystemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalSystemId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalSystemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalSystem**](ExternalSystem.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalSystems

> ExternalSystemPaginatedList GetExternalSystems(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List external systems



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterCreatedAt := []string{"Inner_example"} // []string | Filter by createdAt query param.  **Format:** filter.createdAt={$not}:OPERATION:VALUE    **Example:** filter.createdAt=$btw:John Doe&filter.createdAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterUpdatedAt := []string{"Inner_example"} // []string | Filter by updatedAt query param.  **Format:** filter.updatedAt={$not}:OPERATION:VALUE    **Example:** filter.updatedAt=$btw:John Doe&filter.updatedAt=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  - createdAt  - updatedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalSystemAPI.GetExternalSystems(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterCreatedAt(filterCreatedAt).FilterUpdatedAt(filterUpdatedAt).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalSystemAPI.GetExternalSystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalSystems`: ExternalSystemPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ExternalSystemAPI.GetExternalSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterCreatedAt** | **[]string** | Filter by createdAt query param.  **Format:** filter.createdAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdAt&#x3D;$btw:John Doe&amp;filter.createdAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterUpdatedAt** | **[]string** | Filter by updatedAt query param.  **Format:** filter.updatedAt&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.updatedAt&#x3D;$btw:John Doe&amp;filter.updatedAt&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** id:DESC  **Available Fields** - id  - label  - name  - createdAt  - updatedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - name  | 

### Return type

[**ExternalSystemPaginatedList**](ExternalSystemPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalSystem

> ExternalSystem UpdateExternalSystem(ctx, externalSystemId).UpdateExternalSystem(updateExternalSystem).IfMatch(ifMatch).Execute()

Update external system



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
	externalSystemId := int32(56) // int32 | 
	updateExternalSystem := *openapiclient.NewUpdateExternalSystem() // UpdateExternalSystem | The external system configuration changes
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalSystemAPI.UpdateExternalSystem(context.Background(), externalSystemId).UpdateExternalSystem(updateExternalSystem).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalSystemAPI.UpdateExternalSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalSystem`: ExternalSystem
	fmt.Fprintf(os.Stdout, "Response from `ExternalSystemAPI.UpdateExternalSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalSystemId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExternalSystem** | [**UpdateExternalSystem**](UpdateExternalSystem.md) | The external system configuration changes | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**ExternalSystem**](ExternalSystem.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

