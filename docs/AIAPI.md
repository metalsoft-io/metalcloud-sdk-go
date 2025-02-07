# \AIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateEliResponse**](AIAPI.md#GenerateEliResponse) | **Post** /api/v2/ai/generate | Request from AI a response for the given input



## GenerateEliResponse

> AIGenerateResponse GenerateEliResponse(ctx).AIGenerateRequest(aIGenerateRequest).Execute()

Request from AI a response for the given input



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
	aIGenerateRequest := *openapiclient.NewAIGenerateRequest("Datacenter_example", "Prompt_example") // AIGenerateRequest | The AI generate request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAPI.GenerateEliResponse(context.Background()).AIGenerateRequest(aIGenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.GenerateEliResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateEliResponse`: AIGenerateResponse
	fmt.Fprintf(os.Stdout, "Response from `AIAPI.GenerateEliResponse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateEliResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aIGenerateRequest** | [**AIGenerateRequest**](AIGenerateRequest.md) | The AI generate request | 

### Return type

[**AIGenerateResponse**](AIGenerateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

