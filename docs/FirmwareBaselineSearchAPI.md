# \FirmwareBaselineSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchFirmwareBaselines**](FirmwareBaselineSearchAPI.md#SearchFirmwareBaselines) | **Post** /api/v2/firmware/search/baseline | Search Firmware Baselines



## SearchFirmwareBaselines

> map[string]interface{} SearchFirmwareBaselines(ctx).SearchFirmwareBinary(searchFirmwareBinary).Execute()

Search Firmware Baselines



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
	searchFirmwareBinary := *openapiclient.NewSearchFirmwareBinary(openapiclient.FirmwareVendorType("dell"), *openapiclient.NewBaselineFilter()) // SearchFirmwareBinary | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareBaselineSearchAPI.SearchFirmwareBaselines(context.Background()).SearchFirmwareBinary(searchFirmwareBinary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareBaselineSearchAPI.SearchFirmwareBaselines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFirmwareBaselines`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FirmwareBaselineSearchAPI.SearchFirmwareBaselines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFirmwareBaselinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchFirmwareBinary** | [**SearchFirmwareBinary**](SearchFirmwareBinary.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

