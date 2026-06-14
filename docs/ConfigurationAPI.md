# \ConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationAPI.md#GetConfiguration) | **Get** /api/v2/config | Get configuration
[**PatchConfiguration**](ConfigurationAPI.md#PatchConfiguration) | **Patch** /api/v2/config/{filter} | Partially update configuration
[**PutConfiguration**](ConfigurationAPI.md#PutConfiguration) | **Put** /api/v2/config/{filter} | Replace configuration



## GetConfiguration

> ConfigurationsDto GetConfiguration(ctx).Filter(filter).Execute()

Get configuration



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
	filter := "filter_example" // string | Service to retrieve configuration for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetConfiguration(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: ConfigurationsDto
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Service to retrieve configuration for. | 

### Return type

[**ConfigurationsDto**](ConfigurationsDto.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConfiguration

> PutConfigurationRequest PatchConfiguration(ctx, filter).PutConfigurationRequest(putConfigurationRequest).Execute()

Partially update configuration



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
	filter := "filter_example" // string | Service to update configuration for.
	putConfigurationRequest := openapiclient.putConfiguration_request{AuthConfigurationDto: openapiclient.NewAuthConfigurationDto()} // PutConfigurationRequest | Partial configuration object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.PatchConfiguration(context.Background(), filter).PutConfigurationRequest(putConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PatchConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConfiguration`: PutConfigurationRequest
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PatchConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Service to update configuration for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putConfigurationRequest** | [**PutConfigurationRequest**](PutConfigurationRequest.md) | Partial configuration object | 

### Return type

[**PutConfigurationRequest**](PutConfigurationRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConfiguration

> PutConfigurationRequest PutConfiguration(ctx, filter).PutConfigurationRequest(putConfigurationRequest).Execute()

Replace configuration



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
	filter := "filter_example" // string | Service to update configuration for.
	putConfigurationRequest := openapiclient.putConfiguration_request{AuthConfigurationDto: openapiclient.NewAuthConfigurationDto()} // PutConfigurationRequest | Configuration object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.PutConfiguration(context.Background(), filter).PutConfigurationRequest(putConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PutConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConfiguration`: PutConfigurationRequest
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PutConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Service to update configuration for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putConfigurationRequest** | [**PutConfigurationRequest**](PutConfigurationRequest.md) | Configuration object | 

### Return type

[**PutConfigurationRequest**](PutConfigurationRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

