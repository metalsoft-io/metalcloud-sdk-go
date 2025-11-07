# \ServerDefaultCredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerDefaultCredentials**](ServerDefaultCredentialsAPI.md#CreateServerDefaultCredentials) | **Post** /api/v2/servers/default-credentials | Creates a Server Default Credentials
[**DeleteServerDefaultCredentials**](ServerDefaultCredentialsAPI.md#DeleteServerDefaultCredentials) | **Delete** /api/v2/servers/default-credentials/{serverDefaultCredentialsId} | Deletes a Server Default Credentials
[**GetServerDefaultCredentialsCredentials**](ServerDefaultCredentialsAPI.md#GetServerDefaultCredentialsCredentials) | **Get** /api/v2/servers/default-credentials/{serverDefaultCredentialsId}/credentials | Get Server Default Credentials unencrypted
[**GetServerDefaultCredentialsInfo**](ServerDefaultCredentialsAPI.md#GetServerDefaultCredentialsInfo) | **Get** /api/v2/servers/default-credentials/{serverDefaultCredentialsId} | Get Server Default Credentials information
[**GetServersDefaultCredentials**](ServerDefaultCredentialsAPI.md#GetServersDefaultCredentials) | **Get** /api/v2/servers/default-credentials | Get a list of Server Default Credentials
[**UpdateServerDefaultCredentials**](ServerDefaultCredentialsAPI.md#UpdateServerDefaultCredentials) | **Patch** /api/v2/servers/default-credentials/{serverDefaultCredentialsId} | Updates a Server Default Credentials



## CreateServerDefaultCredentials

> ServerDefaultCredentials CreateServerDefaultCredentials(ctx).CreateServerDefaultCredentials(createServerDefaultCredentials).Execute()

Creates a Server Default Credentials



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
	createServerDefaultCredentials := *openapiclient.NewCreateServerDefaultCredentials(float32(123), "ServerSerialNumber_example", "ServerMacAddress_example", "DefaultUsername_example", "DefaultPassword_example") // CreateServerDefaultCredentials | The Server Default Credentials create object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.CreateServerDefaultCredentials(context.Background()).CreateServerDefaultCredentials(createServerDefaultCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.CreateServerDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerDefaultCredentials`: ServerDefaultCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.CreateServerDefaultCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerDefaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServerDefaultCredentials** | [**CreateServerDefaultCredentials**](CreateServerDefaultCredentials.md) | The Server Default Credentials create object | 

### Return type

[**ServerDefaultCredentials**](ServerDefaultCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerDefaultCredentials

> DeleteServerDefaultCredentials(ctx, serverDefaultCredentialsId).Execute()

Deletes a Server Default Credentials



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerDefaultCredentialsAPI.DeleteServerDefaultCredentials(context.Background(), serverDefaultCredentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.DeleteServerDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerDefaultCredentialsRequest struct via the builder pattern


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


## GetServerDefaultCredentialsCredentials

> ServerDefaultCredentialsCredentials GetServerDefaultCredentialsCredentials(ctx, serverDefaultCredentialsId).Execute()

Get Server Default Credentials unencrypted



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.GetServerDefaultCredentialsCredentials(context.Background(), serverDefaultCredentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDefaultCredentialsCredentials`: ServerDefaultCredentialsCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDefaultCredentialsCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerDefaultCredentialsCredentials**](ServerDefaultCredentialsCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDefaultCredentialsInfo

> ServerDefaultCredentials GetServerDefaultCredentialsInfo(ctx, serverDefaultCredentialsId).Execute()

Get Server Default Credentials information



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.GetServerDefaultCredentialsInfo(context.Background(), serverDefaultCredentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerDefaultCredentialsInfo`: ServerDefaultCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.GetServerDefaultCredentialsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDefaultCredentialsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerDefaultCredentials**](ServerDefaultCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServersDefaultCredentials

> ServerDefaultCredentialsPaginatedList GetServersDefaultCredentials(ctx).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterServerSerialNumber(filterServerSerialNumber).FilterServerMacAddress(filterServerMacAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Server Default Credentials



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
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerSerialNumber := []string{"Inner_example"} // []string | Filter by serverSerialNumber query param.  **Format:** filter.serverSerialNumber={$not}:OPERATION:VALUE    **Example:** filter.serverSerialNumber=$btw:John Doe&filter.serverSerialNumber=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerMacAddress := []string{"Inner_example"} // []string | Filter by serverMacAddress query param.  **Format:** filter.serverMacAddress={$not}:OPERATION:VALUE    **Example:** filter.serverMacAddress=$btw:John Doe&filter.serverMacAddress=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,siteId,serverSerialNumber,serverMacAddress   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - siteId  - serverSerialNumber  - serverMacAddress  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.GetServersDefaultCredentials(context.Background()).Page(page).Limit(limit).FilterSiteId(filterSiteId).FilterServerSerialNumber(filterServerSerialNumber).FilterServerMacAddress(filterServerMacAddress).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.GetServersDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServersDefaultCredentials`: ServerDefaultCredentialsPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.GetServersDefaultCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServersDefaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerSerialNumber** | **[]string** | Filter by serverSerialNumber query param.  **Format:** filter.serverSerialNumber&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverSerialNumber&#x3D;$btw:John Doe&amp;filter.serverSerialNumber&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerMacAddress** | **[]string** | Filter by serverMacAddress query param.  **Format:** filter.serverMacAddress&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverMacAddress&#x3D;$btw:John Doe&amp;filter.serverMacAddress&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;siteId:DESC   **Default Value:** id:DESC  **Available Fields** - id  - siteId  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,siteId,serverSerialNumber,serverMacAddress   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - siteId  - serverSerialNumber  - serverMacAddress  | 

### Return type

[**ServerDefaultCredentialsPaginatedList**](ServerDefaultCredentialsPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDefaultCredentials

> ServerDefaultCredentials UpdateServerDefaultCredentials(ctx, serverDefaultCredentialsId).UpdateServerDefaultCredentials(updateServerDefaultCredentials).Execute()

Updates a Server Default Credentials



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
	serverDefaultCredentialsId := float32(8.14) // float32 | 
	updateServerDefaultCredentials := *openapiclient.NewUpdateServerDefaultCredentials() // UpdateServerDefaultCredentials | The Server Default Credentials update object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerDefaultCredentialsAPI.UpdateServerDefaultCredentials(context.Background(), serverDefaultCredentialsId).UpdateServerDefaultCredentials(updateServerDefaultCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerDefaultCredentialsAPI.UpdateServerDefaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerDefaultCredentials`: ServerDefaultCredentials
	fmt.Fprintf(os.Stdout, "Response from `ServerDefaultCredentialsAPI.UpdateServerDefaultCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverDefaultCredentialsId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDefaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerDefaultCredentials** | [**UpdateServerDefaultCredentials**](UpdateServerDefaultCredentials.md) | The Server Default Credentials update object | 

### Return type

[**ServerDefaultCredentials**](ServerDefaultCredentials.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

