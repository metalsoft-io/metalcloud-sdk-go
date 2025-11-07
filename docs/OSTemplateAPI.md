# \OSTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSTemplate**](OSTemplateAPI.md#CreateOSTemplate) | **Post** /api/v2/os-templates | Create OS template
[**DeleteOSTemplate**](OSTemplateAPI.md#DeleteOSTemplate) | **Delete** /api/v2/os-templates/{osTemplateId} | Delete OS template
[**GetOSTemplate**](OSTemplateAPI.md#GetOSTemplate) | **Get** /api/v2/os-templates/{osTemplateId} | Get details for an OS template
[**GetOSTemplateCredentials**](OSTemplateAPI.md#GetOSTemplateCredentials) | **Get** /api/v2/os-templates/{osTemplateId}/credentials | Get OS template credentials
[**GetOSTemplates**](OSTemplateAPI.md#GetOSTemplates) | **Get** /api/v2/os-templates | Get a list of available OS templates
[**UpdateOSTemplate**](OSTemplateAPI.md#UpdateOSTemplate) | **Put** /api/v2/os-templates/{osTemplateId} | Update OS template



## CreateOSTemplate

> OSTemplate CreateOSTemplate(ctx).OSTemplateCreate(oSTemplateCreate).Execute()

Create OS template



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
	oSTemplateCreate := *openapiclient.NewOSTemplateCreate("Ubuntu 24.04 LTS", *openapiclient.NewOSTemplateDevice("server", "uefi", "x86_64"), *openapiclient.NewOSTemplateInstall("oob", "local_drive", "wait_for_power_off"), *openapiclient.NewOSTemplateImageBuild(true), *openapiclient.NewOSTemplateOs("Ubuntu | CentOS | Suse | Windows | ESXi | ...", "22.04 | 8 | 15 | 2019 | ...", *openapiclient.NewOSTemplateOsCredential("root", "plain"))) // OSTemplateCreate | The OS template details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.CreateOSTemplate(context.Background()).OSTemplateCreate(oSTemplateCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.CreateOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSTemplate`: OSTemplate
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.CreateOSTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSTemplateCreate** | [**OSTemplateCreate**](OSTemplateCreate.md) | The OS template details | 

### Return type

[**OSTemplate**](OSTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSTemplate

> DeleteOSTemplate(ctx, osTemplateId).Execute()

Delete OS template



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
	osTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSTemplateAPI.DeleteOSTemplate(context.Background(), osTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.DeleteOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSTemplateRequest struct via the builder pattern


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


## GetOSTemplate

> OSTemplate GetOSTemplate(ctx, osTemplateId).Execute()

Get details for an OS template



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
	osTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.GetOSTemplate(context.Background(), osTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.GetOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSTemplate`: OSTemplate
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.GetOSTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSTemplate**](OSTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSTemplateCredentials

> OSTemplateOsCredential GetOSTemplateCredentials(ctx, osTemplateId).Execute()

Get OS template credentials



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
	osTemplateId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.GetOSTemplateCredentials(context.Background(), osTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.GetOSTemplateCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSTemplateCredentials`: OSTemplateOsCredential
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.GetOSTemplateCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSTemplateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSTemplateOsCredential**](OSTemplateOsCredential.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSTemplates

> OSTemplatePaginatedList GetOSTemplates(ctx).Page(page).Limit(limit).FilterDeviceType(filterDeviceType).FilterLabel(filterLabel).FilterOsName(filterOsName).FilterInstallOnieStrings(filterInstallOnieStrings).FilterStatus(filterStatus).FilterVisibility(filterVisibility).FilterCreatedBy(filterCreatedBy).FilterModifiedBy(filterModifiedBy).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get a list of available OS templates



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
	filterDeviceType := []string{"Inner_example"} // []string | Filter by device.type query param.  **Format:** filter.device.type={$not}:OPERATION:VALUE    **Example:** filter.device.type=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe&filter.label=$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or (optional)
	filterOsName := []string{"Inner_example"} // []string | Filter by os.name query param.  **Format:** filter.os.name={$not}:OPERATION:VALUE    **Example:** filter.os.name=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterInstallOnieStrings := []string{"Inner_example"} // []string | Filter by install.onieStrings query param.  **Format:** filter.install.onieStrings={$not}:OPERATION:VALUE    **Example:** filter.install.onieStrings=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterVisibility := []string{"Inner_example"} // []string | Filter by visibility query param.  **Format:** filter.visibility={$not}:OPERATION:VALUE    **Example:** filter.visibility=$eq:John Doe&filter.visibility=$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or (optional)
	filterCreatedBy := []string{"Inner_example"} // []string | Filter by createdBy query param.  **Format:** filter.createdBy={$not}:OPERATION:VALUE    **Example:** filter.createdBy=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterModifiedBy := []string{"Inner_example"} // []string | Filter by modifiedBy query param.  **Format:** filter.modifiedBy={$not}:OPERATION:VALUE    **Example:** filter.modifiedBy=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.  **Format:** filter.tags={$not}:OPERATION:VALUE    **Example:** filter.tags=$eq:John Doe&filter.tags=$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=device.type:DESC   **Default Value:** id:DESC  **Available Fields** - id  - device.type  - visibility  - status  - createdBy  - createdAt  - modifiedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,label,device.type,os.name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - label  - device.type  - os.name  - os.version  - install.onieStrings  - visibility  - status  - createdBy  - modifiedBy  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,name,description,label,device.type   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.GetOSTemplates(context.Background()).Page(page).Limit(limit).FilterDeviceType(filterDeviceType).FilterLabel(filterLabel).FilterOsName(filterOsName).FilterInstallOnieStrings(filterInstallOnieStrings).FilterStatus(filterStatus).FilterVisibility(filterVisibility).FilterCreatedBy(filterCreatedBy).FilterModifiedBy(filterModifiedBy).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.GetOSTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSTemplates`: OSTemplatePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.GetOSTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOSTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterDeviceType** | **[]string** | Filter by device.type query param.  **Format:** filter.device.type&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.device.type&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe&amp;filter.label&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or | 
 **filterOsName** | **[]string** | Filter by os.name query param.  **Format:** filter.os.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.os.name&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterInstallOnieStrings** | **[]string** | Filter by install.onieStrings query param.  **Format:** filter.install.onieStrings&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.install.onieStrings&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterVisibility** | **[]string** | Filter by visibility query param.  **Format:** filter.visibility&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.visibility&#x3D;$eq:John Doe&amp;filter.visibility&#x3D;$in:John Doe  **Available Operations** - $eq  - $in  - $and  - $or | 
 **filterCreatedBy** | **[]string** | Filter by createdBy query param.  **Format:** filter.createdBy&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdBy&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterModifiedBy** | **[]string** | Filter by modifiedBy query param.  **Format:** filter.modifiedBy&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.modifiedBy&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterTags** | **[]string** | Filter by tags query param.  **Format:** filter.tags&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.tags&#x3D;$eq:John Doe&amp;filter.tags&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;device.type:DESC   **Default Value:** id:DESC  **Available Fields** - id  - device.type  - visibility  - status  - createdBy  - createdAt  - modifiedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,label,device.type,os.name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - label  - device.type  - os.name  - os.version  - install.onieStrings  - visibility  - status  - createdBy  - modifiedBy  | 
 **select_** | **string** | List of fields to select.  **Example:** id,name,description,label,device.type   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**OSTemplatePaginatedList**](OSTemplatePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSTemplate

> OSTemplate UpdateOSTemplate(ctx, osTemplateId).OSTemplateUpdate(oSTemplateUpdate).IfMatch(ifMatch).Execute()

Update OS template



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
	osTemplateId := float32(8.14) // float32 | 
	oSTemplateUpdate := *openapiclient.NewOSTemplateUpdate("Ubuntu 24.04 LTS", *openapiclient.NewOSTemplateDevice("server", "uefi", "x86_64"), *openapiclient.NewOSTemplateInstall("oob", "local_drive", "wait_for_power_off"), *openapiclient.NewOSTemplateImageBuild(true), *openapiclient.NewOSTemplateOs("Ubuntu | CentOS | Suse | Windows | ESXi | ...", "22.04 | 8 | 15 | 2019 | ...", *openapiclient.NewOSTemplateOsCredential("root", "plain"))) // OSTemplateUpdate | The OS template details
	ifMatch := "ifMatch_example" // string | Entity tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSTemplateAPI.UpdateOSTemplate(context.Background(), osTemplateId).OSTemplateUpdate(oSTemplateUpdate).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSTemplateAPI.UpdateOSTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSTemplate`: OSTemplate
	fmt.Fprintf(os.Stdout, "Response from `OSTemplateAPI.UpdateOSTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osTemplateId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oSTemplateUpdate** | [**OSTemplateUpdate**](OSTemplateUpdate.md) | The OS template details | 
 **ifMatch** | **string** | Entity tag | 

### Return type

[**OSTemplate**](OSTemplate.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

