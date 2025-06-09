# \UnmanagedServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportUnmanagedServer**](UnmanagedServersAPI.md#ImportUnmanagedServer) | **Post** /api/v2/servers/unmanaged/import | Import Unmanaged Server



## ImportUnmanagedServer

> Server ImportUnmanagedServer(ctx).ServerUnmanagedImport(serverUnmanagedImport).Execute()

Import Unmanaged Server



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
	serverUnmanagedImport := *openapiclient.NewServerUnmanagedImport(float32(123), float32(123), []openapiclient.ServerUnmanagedImportInternalInterface{*openapiclient.NewServerUnmanagedImportInternalInterface("ServerInterfaceMacAddress_example", "IdentifierString_example", "NetworkEquipmentInterfaceIdentifierString_example")}) // ServerUnmanagedImport | The unmanaged server data to import

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnmanagedServersAPI.ImportUnmanagedServer(context.Background()).ServerUnmanagedImport(serverUnmanagedImport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnmanagedServersAPI.ImportUnmanagedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportUnmanagedServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `UnmanagedServersAPI.ImportUnmanagedServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportUnmanagedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverUnmanagedImport** | [**ServerUnmanagedImport**](ServerUnmanagedImport.md) | The unmanaged server data to import | 

### Return type

[**Server**](Server.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

