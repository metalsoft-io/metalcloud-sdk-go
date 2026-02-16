# \SubnetAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnet**](SubnetAPI.md#CreateSubnet) | **Post** /api/v2/subnets | Create a subnet.
[**DeleteSubnet**](SubnetAPI.md#DeleteSubnet) | **Delete** /api/v2/subnets/{subnetId} | Delete Subnet
[**GetSubnet**](SubnetAPI.md#GetSubnet) | **Get** /api/v2/subnets/{subnetId} | Retrieves the Subnet information
[**GetSubnetIpRanges**](SubnetAPI.md#GetSubnetIpRanges) | **Get** /api/v2/subnets/{subnetId}/ip-ranges | List all Subnet IP Ranges
[**GetSubnetIps**](SubnetAPI.md#GetSubnetIps) | **Get** /api/v2/subnets/{subnetId}/ips | List all Subnet IPs
[**GetSubnets**](SubnetAPI.md#GetSubnets) | **Get** /api/v2/subnets | List all Subnets
[**UpdateSubnet**](SubnetAPI.md#UpdateSubnet) | **Patch** /api/v2/subnets/{subnetId} | Updates Subnet



## CreateSubnet

> Subnet CreateSubnet(ctx).CreateSubnet(createSubnet).Execute()

Create a subnet.

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
	createSubnet := *openapiclient.NewCreateSubnet("NetworkAddress_example", int32(123), false) // CreateSubnet | The Subnet to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.CreateSubnet(context.Background()).CreateSubnet(createSubnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.CreateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubnet`: Subnet
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubnet** | [**CreateSubnet**](CreateSubnet.md) | The Subnet to create | 

### Return type

[**Subnet**](Subnet.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubnet

> DeleteSubnet(ctx, subnetId).IfMatch(ifMatch).Execute()

Delete Subnet



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
	subnetId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubnetAPI.DeleteSubnet(context.Background(), subnetId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.DeleteSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


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


## GetSubnet

> Subnet GetSubnet(ctx, subnetId).Execute()

Retrieves the Subnet information



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
	subnetId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.GetSubnet(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.GetSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnet`: Subnet
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.GetSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subnet**](Subnet.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubnetIpRanges

> PaginatedIpRangeList GetSubnetIpRanges(ctx, subnetId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterSubnetId(filterSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List all Subnet IP Ranges



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
	subnetId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 1000    **Max Value:** 5000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe&filter.id=$gt:John Doe  **Available Operations** - $eq  - $in  - $gt  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterIpVersion := []string{"Inner_example"} // []string | Filter by ipVersion query param.  **Format:** filter.ipVersion={$not}:OPERATION:VALUE    **Example:** filter.ipVersion=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubnetId := []string{"Inner_example"} // []string | Filter by subnetId query param.  **Format:** filter.subnetId={$not}:OPERATION:VALUE    **Example:** filter.subnetId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.GetSubnetIpRanges(context.Background(), subnetId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterSubnetId(filterSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.GetSubnetIpRanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnetIpRanges`: PaginatedIpRangeList
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.GetSubnetIpRanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetIpRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 1000    **Max Value:** 5000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$gt:John Doe  **Available Operations** - $eq  - $in  - $gt  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterIpVersion** | **[]string** | Filter by ipVersion query param.  **Format:** filter.ipVersion&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.ipVersion&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubnetId** | **[]string** | Filter by subnetId query param.  **Format:** filter.subnetId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subnetId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 

### Return type

[**PaginatedIpRangeList**](PaginatedIpRangeList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubnetIps

> PaginatedIpList GetSubnetIps(ctx, subnetId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterSubnetId(filterSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List all Subnet IPs



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
	subnetId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 1000    **Max Value:** 5000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe&filter.id=$gt:John Doe  **Available Operations** - $eq  - $in  - $gt  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterIpVersion := []string{"Inner_example"} // []string | Filter by ipVersion query param.  **Format:** filter.ipVersion={$not}:OPERATION:VALUE    **Example:** filter.ipVersion=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterSubnetId := []string{"Inner_example"} // []string | Filter by subnetId query param.  **Format:** filter.subnetId={$not}:OPERATION:VALUE    **Example:** filter.subnetId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.GetSubnetIps(context.Background(), subnetId).Page(page).Limit(limit).FilterId(filterId).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterSubnetId(filterSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.GetSubnetIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnetIps`: PaginatedIpList
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.GetSubnetIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 1000    **Max Value:** 5000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$gt:John Doe  **Available Operations** - $eq  - $in  - $gt  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterIpVersion** | **[]string** | Filter by ipVersion query param.  **Format:** filter.ipVersion&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.ipVersion&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterSubnetId** | **[]string** | Filter by subnetId query param.  **Format:** filter.subnetId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.subnetId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 

### Return type

[**PaginatedIpList**](PaginatedIpList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubnets

> SubnetPaginatedList GetSubnets(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterIsPool(filterIsPool).FilterParentSubnetId(filterParentSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

List all Subnets



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
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 1000    **Max Value:** 5000   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe&filter.id=$gt:John Doe  **Available Operations** - $eq  - $in  - $gt  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterIpVersion := []string{"Inner_example"} // []string | Filter by ipVersion query param.  **Format:** filter.ipVersion={$not}:OPERATION:VALUE    **Example:** filter.ipVersion=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterIsPool := []string{"Inner_example"} // []string | Filter by isPool query param.  **Format:** filter.isPool={$not}:OPERATION:VALUE    **Example:** filter.isPool=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterParentSubnetId := []string{"Inner_example"} // []string | Filter by parentSubnetId query param.  **Format:** filter.parentSubnetId={$not}:OPERATION:VALUE    **Example:** filter.parentSubnetId=$eq:John Doe&filter.parentSubnetId=$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC   **Default Value:** id:ASC  **Available Fields** - id  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  - name  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.GetSubnets(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterIsPool(filterIsPool).FilterParentSubnetId(filterParentSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.GetSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnets`: SubnetPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.GetSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 1000    **Max Value:** 5000   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe&amp;filter.id&#x3D;$gt:John Doe  **Available Operations** - $eq  - $in  - $gt  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterIpVersion** | **[]string** | Filter by ipVersion query param.  **Format:** filter.ipVersion&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.ipVersion&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterIsPool** | **[]string** | Filter by isPool query param.  **Format:** filter.isPool&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.isPool&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterParentSubnetId** | **[]string** | Filter by parentSubnetId query param.  **Format:** filter.parentSubnetId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.parentSubnetId&#x3D;$eq:John Doe&amp;filter.parentSubnetId&#x3D;$null:John Doe  **Available Operations** - $eq  - $null  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC   **Default Value:** id:ASC  **Available Fields** - id  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** label,name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  - name  | 

### Return type

[**SubnetPaginatedList**](SubnetPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubnet

> Subnet UpdateSubnet(ctx, subnetId).IfMatch(ifMatch).UpdateSubnet(updateSubnet).Execute()

Updates Subnet



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
	subnetId := int32(56) // int32 | 
	ifMatch := "ifMatch_example" // string | Entity tag
	updateSubnet := *openapiclient.NewUpdateSubnet() // UpdateSubnet | The Subnet changes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.UpdateSubnet(context.Background(), subnetId).IfMatch(ifMatch).UpdateSubnet(updateSubnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.UpdateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubnet`: Subnet
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.UpdateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Entity tag | 
 **updateSubnet** | [**UpdateSubnet**](UpdateSubnet.md) | The Subnet changes | 

### Return type

[**Subnet**](Subnet.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

