# \SubnetAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnet**](SubnetAPI.md#CreateSubnet) | **Post** /api/v2/subnets | Create a subnet.
[**DeleteSubnet**](SubnetAPI.md#DeleteSubnet) | **Delete** /api/v2/subnets/{subnetId} | Delete Subnet
[**GetSubnet**](SubnetAPI.md#GetSubnet) | **Get** /api/v2/subnets/{subnetId} | Retrieves the Subnet information
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


## GetSubnets

> SubnetPaginatedList GetSubnets(ctx).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterParentSubnetId(filterParentSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 1000           </p>       <p>              <b>Max Value: </b> 5000           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li> <li>$gt</li></ul> (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterIpVersion := []string{"Inner_example"} // []string | Filter by ipVersion query param.           <p>              <b>Format: </b> filter.ipVersion={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.ipVersion=$not:$like:John Doe&filter.ipVersion=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterParentSubnetId := []string{"Inner_example"} // []string | Filter by parentSubnetId query param.           <p>              <b>Format: </b> filter.parentSubnetId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.parentSubnetId=$not:$like:John Doe&filter.parentSubnetId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$null</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:ASC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,name           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>name</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.GetSubnets(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterLabel(filterLabel).FilterName(filterName).FilterIpVersion(filterIpVersion).FilterParentSubnetId(filterParentSubnetId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1000           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 5000           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt;&lt;/ul&gt; | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterIpVersion** | **[]string** | Filter by ipVersion query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.ipVersion&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.ipVersion&#x3D;$not:$like:John Doe&amp;filter.ipVersion&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterParentSubnetId** | **[]string** | Filter by parentSubnetId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.parentSubnetId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.parentSubnetId&#x3D;$not:$like:John Doe&amp;filter.parentSubnetId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;          | 

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

