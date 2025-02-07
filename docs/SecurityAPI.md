# \SecurityAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProviders**](SecurityAPI.md#GetProviders) | **Get** /api/v2/authentication/providers | Get available authentication providers
[**UpdateProvider**](SecurityAPI.md#UpdateProvider) | **Patch** /api/v2/authentication/providers/{name} | Updates authentication provider



## GetProviders

> []AuthenticationProvider GetProviders(ctx).Execute()

Get available authentication providers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviders`: []AuthenticationProvider
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersRequest struct via the builder pattern


### Return type

[**[]AuthenticationProvider**](AuthenticationProvider.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvider

> AuthenticationProvider UpdateProvider(ctx, name).AuthenticationProviderUpdate(authenticationProviderUpdate).Execute()

Updates authentication provider



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
	name := "name_example" // string | The provider name
	authenticationProviderUpdate := *openapiclient.NewAuthenticationProviderUpdate(false, []string{"Domains_example"}) // AuthenticationProviderUpdate | The provider updates

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateProvider(context.Background(), name).AuthenticationProviderUpdate(authenticationProviderUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProvider`: AuthenticationProvider
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The provider name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationProviderUpdate** | [**AuthenticationProviderUpdate**](AuthenticationProviderUpdate.md) | The provider updates | 

### Return type

[**AuthenticationProvider**](AuthenticationProvider.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

