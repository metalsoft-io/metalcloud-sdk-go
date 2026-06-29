# \ConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationAPI.md#GetConfiguration) | **Get** /api/v2/config | Get configuration
[**ReplaceConfiguration**](ConfigurationAPI.md#ReplaceConfiguration) | **Put** /api/v2/config/{filter} | Replace configuration
[**UpdateConfiguration**](ConfigurationAPI.md#UpdateConfiguration) | **Patch** /api/v2/config/{filter} | Partially update configuration



## GetConfiguration

> Configurations GetConfiguration(ctx).Filter(filter).Execute()

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
	// response from `GetConfiguration`: Configurations
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

[**Configurations**](Configurations.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceConfiguration

> ReplaceConfigurationRequest ReplaceConfiguration(ctx, filter).ReplaceConfigurationRequest(replaceConfigurationRequest).Execute()

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
	replaceConfigurationRequest := openapiclient.replaceConfiguration_request{AuthConfiguration: openapiclient.NewAuthConfiguration()} // ReplaceConfigurationRequest | Configuration object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ReplaceConfiguration(context.Background(), filter).ReplaceConfigurationRequest(replaceConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ReplaceConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceConfiguration`: ReplaceConfigurationRequest
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ReplaceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Service to update configuration for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceConfigurationRequest** | [**ReplaceConfigurationRequest**](ReplaceConfigurationRequest.md) | Configuration object | 

### Return type

[**ReplaceConfigurationRequest**](ReplaceConfigurationRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ReplaceConfigurationRequest UpdateConfiguration(ctx, filter).ReplaceConfigurationRequest(replaceConfigurationRequest).Execute()

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
	replaceConfigurationRequest := openapiclient.replaceConfiguration_request{AuthConfiguration: openapiclient.NewAuthConfiguration()} // ReplaceConfigurationRequest | Partial configuration object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.UpdateConfiguration(context.Background(), filter).ReplaceConfigurationRequest(replaceConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.UpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfiguration`: ReplaceConfigurationRequest
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Service to update configuration for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceConfigurationRequest** | [**ReplaceConfigurationRequest**](ReplaceConfigurationRequest.md) | Partial configuration object | 

### Return type

[**ReplaceConfigurationRequest**](ReplaceConfigurationRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

