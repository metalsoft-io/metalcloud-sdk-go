# \DNSRecordSetAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDNSRecordSets**](DNSRecordSetAPI.md#ListDNSRecordSets) | **Get** /api/v2/dns-recordsets | List all RecordSets



## ListDNSRecordSets

> DNSRecordSetPaginatedList ListDNSRecordSets(ctx).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterName(filterName).FilterZoneName(filterZoneName).FilterStatus(filterStatus).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

List all RecordSets

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
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.  **Format:** filter.name={$not}:OPERATION:VALUE    **Example:** filter.name=$eq:John Doe&filter.name=$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or (optional)
	filterZoneName := []string{"Inner_example"} // []string | Filter by zoneName query param.  **Format:** filter.zoneName={$not}:OPERATION:VALUE    **Example:** filter.zoneName=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.  **Format:** filter.tags={$not}:OPERATION:VALUE    **Example:** filter.tags=$eq:John Doe&filter.tags=$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=status:DESC   **Default Value:** name:DESC  **Available Fields** - id  - status  - siteId  - infrastructureId  - name  - zoneName  - type  - ttl  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,siteId,infrastructureId,name,zoneName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - siteId  - infrastructureId  - name  - zoneName  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,status,siteId,infrastructureId,name   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSRecordSetAPI.ListDNSRecordSets(context.Background()).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterName(filterName).FilterZoneName(filterZoneName).FilterStatus(filterStatus).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordSetAPI.ListDNSRecordSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDNSRecordSets`: DNSRecordSetPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DNSRecordSetAPI.ListDNSRecordSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNSRecordSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterName** | **[]string** | Filter by name query param.  **Format:** filter.name&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.name&#x3D;$eq:John Doe&amp;filter.name&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $ilike  - $and  - $or | 
 **filterZoneName** | **[]string** | Filter by zoneName query param.  **Format:** filter.zoneName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.zoneName&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterTags** | **[]string** | Filter by tags query param.  **Format:** filter.tags&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.tags&#x3D;$eq:John Doe&amp;filter.tags&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;status:DESC   **Default Value:** name:DESC  **Available Fields** - id  - status  - siteId  - infrastructureId  - name  - zoneName  - type  - ttl  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,siteId,infrastructureId,name,zoneName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - siteId  - infrastructureId  - name  - zoneName  | 
 **select_** | **string** | List of fields to select.  **Example:** id,status,siteId,infrastructureId,name   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**DNSRecordSetPaginatedList**](DNSRecordSetPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

