# \FirmwareBaselineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirmwareBaseline**](FirmwareBaselineAPI.md#CreateFirmwareBaseline) | **Post** /api/v2/firmware/baseline | Create a new firmware baseline
[**DeleteFirmwareBaseline**](FirmwareBaselineAPI.md#DeleteFirmwareBaseline) | **Delete** /api/v2/firmware/baseline/{firmwareBaselineId} | Delete Firmware Baseline
[**GetFirmwareBaseline**](FirmwareBaselineAPI.md#GetFirmwareBaseline) | **Get** /api/v2/firmware/baseline/{firmwareBaselineId} | Get Firmware Baseline
[**GetFirmwareBaselines**](FirmwareBaselineAPI.md#GetFirmwareBaselines) | **Get** /api/v2/firmware/baseline | Get Firmware Baselines
[**UpdateFirmwareBaseline**](FirmwareBaselineAPI.md#UpdateFirmwareBaseline) | **Put** /api/v2/firmware/baseline/{firmwareBaselineId} | Update Firmware Baseline



## CreateFirmwareBaseline

> CreateFirmwareBaseline CreateFirmwareBaseline(ctx).CreateFirmwareBaseline(createFirmwareBaseline).Execute()

Create a new firmware baseline



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
	createFirmwareBaseline := *openapiclient.NewCreateFirmwareBaseline("Data center baseline name") // CreateFirmwareBaseline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.CreateFirmwareBaseline(context.Background()).CreateFirmwareBaseline(createFirmwareBaseline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.CreateFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirmwareBaseline`: CreateFirmwareBaseline
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.CreateFirmwareBaseline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirmwareBaselineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirmwareBaseline** | [**CreateFirmwareBaseline**](CreateFirmwareBaseline.md) |  | 

### Return type

[**CreateFirmwareBaseline**](CreateFirmwareBaseline.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirmwareBaseline

> DeleteFirmwareBaseline(ctx, firmwareBaselineId).Execute()

Delete Firmware Baseline



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
	firmwareBaselineId := float32(8.14) // float32 | The firmware baseline id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirmwareBaselineAPI.DeleteFirmwareBaseline(context.Background(), firmwareBaselineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.DeleteFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBaselineId** | **float32** | The firmware baseline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirmwareBaselineRequest struct via the builder pattern


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


## GetFirmwareBaseline

> FirmwareBaseline GetFirmwareBaseline(ctx, firmwareBaselineId).Execute()

Get Firmware Baseline



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
	firmwareBaselineId := float32(8.14) // float32 | The firmware baseline id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.GetFirmwareBaseline(context.Background(), firmwareBaselineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.GetFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareBaseline`: FirmwareBaseline
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.GetFirmwareBaseline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBaselineId** | **float32** | The firmware baseline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareBaselineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirmwareBaseline**](FirmwareBaseline.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirmwareBaselines

> FirmwareBaselinePaginatedList GetFirmwareBaselines(ctx).Page(page).Limit(limit).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get Firmware Baselines



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
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,name,description,catalog,createdTimestamp   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.GetFirmwareBaselines(context.Background()).Page(page).Limit(limit).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.GetFirmwareBaselines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareBaselines`: FirmwareBaselinePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.GetFirmwareBaselines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareBaselinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** id:DESC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  | 
 **select_** | **string** | List of fields to select.  **Example:** id,name,description,catalog,createdTimestamp   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**FirmwareBaselinePaginatedList**](FirmwareBaselinePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirmwareBaseline

> FirmwareBaseline UpdateFirmwareBaseline(ctx, firmwareBaselineId).UpdateFirmwareBaseline(updateFirmwareBaseline).Execute()

Update Firmware Baseline



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
	firmwareBaselineId := float32(8.14) // float32 | The firmware baseline id
	updateFirmwareBaseline := *openapiclient.NewUpdateFirmwareBaseline("Data center baseline name") // UpdateFirmwareBaseline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineAPI.UpdateFirmwareBaseline(context.Background(), firmwareBaselineId).UpdateFirmwareBaseline(updateFirmwareBaseline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineAPI.UpdateFirmwareBaseline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirmwareBaseline`: FirmwareBaseline
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineAPI.UpdateFirmwareBaseline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firmwareBaselineId** | **float32** | The firmware baseline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirmwareBaselineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirmwareBaseline** | [**UpdateFirmwareBaseline**](UpdateFirmwareBaseline.md) |  | 

### Return type

[**FirmwareBaseline**](FirmwareBaseline.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

