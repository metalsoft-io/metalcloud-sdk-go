# \DNSZoneAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDNSZone**](DNSZoneAPI.md#CreateDNSZone) | **Post** /api/v2/dns-zones | Create a new DNS Zone
[**DeleteDNSZone**](DNSZoneAPI.md#DeleteDNSZone) | **Delete** /api/v2/dns-zones/{dnsZoneId} | Delete a DNS Zone by ID
[**GetDNSRecordSetById**](DNSZoneAPI.md#GetDNSRecordSetById) | **Get** /api/v2/dns-zones/{id}/recordsets/{recordSetId} | Get a DNS RecordSet by ID
[**GetDNSZoneById**](DNSZoneAPI.md#GetDNSZoneById) | **Get** /api/v2/dns-zones/{dnsZoneId} | Get a DNS Zone by ID
[**GetDNSZoneNameservers**](DNSZoneAPI.md#GetDNSZoneNameservers) | **Get** /api/v2/dns-zones/{dnsZoneId}/nameservers | Get the nameservers of a DNS zone
[**GetDNSZones**](DNSZoneAPI.md#GetDNSZones) | **Get** /api/v2/dns-zones | Get all DNS Zones
[**ListDNSRecordSetsByZoneId**](DNSZoneAPI.md#ListDNSRecordSetsByZoneId) | **Get** /api/v2/dns-zones/{id}/recordsets | List RecordSets in a DNS zone by ID
[**UpdateDNSZone**](DNSZoneAPI.md#UpdateDNSZone) | **Patch** /api/v2/dns-zones/{dnsZoneId} | Update a DNS Zone by ID



## CreateDNSZone

> DnsZone CreateDNSZone(ctx).CreateDnsZone(createDnsZone).Execute()

Create a new DNS Zone



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
	createDnsZone := *openapiclient.NewCreateDnsZone("example.com", true, []string{"NameServers_example"}) // CreateDnsZone | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.CreateDNSZone(context.Background()).CreateDnsZone(createDnsZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.CreateDNSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDNSZone`: DnsZone
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.CreateDNSZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDNSZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDnsZone** | [**CreateDnsZone**](CreateDnsZone.md) |  | 

### Return type

[**DnsZone**](DnsZone.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDNSZone

> DeleteDNSZone(ctx, dnsZoneId).Execute()

Delete a DNS Zone by ID



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
	dnsZoneId := float32(8.14) // float32 | The ID of the DNS Zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DNSZoneAPI.DeleteDNSZone(context.Background(), dnsZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.DeleteDNSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsZoneId** | **float32** | The ID of the DNS Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDNSZoneRequest struct via the builder pattern


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


## GetDNSRecordSetById

> DnsRecordSet GetDNSRecordSetById(ctx, id, recordSetId).Execute()

Get a DNS RecordSet by ID



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
	id := float32(8.14) // float32 | The ID of the DNS Zone
	recordSetId := float32(8.14) // float32 | The ID of the DNS RecordSet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.GetDNSRecordSetById(context.Background(), id, recordSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.GetDNSRecordSetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSRecordSetById`: DnsRecordSet
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.GetDNSRecordSetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the DNS Zone | 
**recordSetId** | **float32** | The ID of the DNS RecordSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSRecordSetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DnsRecordSet**](DnsRecordSet.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDNSZoneById

> DnsZone GetDNSZoneById(ctx, dnsZoneId).Execute()

Get a DNS Zone by ID



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
	dnsZoneId := float32(8.14) // float32 | The ID of the DNS Zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.GetDNSZoneById(context.Background(), dnsZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.GetDNSZoneById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSZoneById`: DnsZone
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.GetDNSZoneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsZoneId** | **float32** | The ID of the DNS Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSZoneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnsZone**](DnsZone.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDNSZoneNameservers

> []string GetDNSZoneNameservers(ctx, dnsZoneId).Execute()

Get the nameservers of a DNS zone



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
	dnsZoneId := float32(8.14) // float32 | The ID of the DNS Zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.GetDNSZoneNameservers(context.Background(), dnsZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.GetDNSZoneNameservers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSZoneNameservers`: []string
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.GetDNSZoneNameservers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsZoneId** | **float32** | The ID of the DNS Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSZoneNameserversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDNSZones

> DNSZonePaginatedList GetDNSZones(ctx).Page(page).Limit(limit).FilterLabel(filterLabel).FilterZoneName(filterZoneName).FilterZoneType(filterZoneType).FilterStatus(filterStatus).FilterIsDefault(filterIsDefault).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

Get all DNS Zones



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$eq:John Doe&filter.label=$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or (optional)
	filterZoneName := []string{"Inner_example"} // []string | Filter by zoneName query param.  **Format:** filter.zoneName={$not}:OPERATION:VALUE    **Example:** filter.zoneName=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterZoneType := []string{"Inner_example"} // []string | Filter by zoneType query param.  **Format:** filter.zoneType={$not}:OPERATION:VALUE    **Example:** filter.zoneType=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterIsDefault := []string{"Inner_example"} // []string | Filter by isDefault query param.  **Format:** filter.isDefault={$not}:OPERATION:VALUE    **Example:** filter.isDefault=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.  **Format:** filter.tags={$not}:OPERATION:VALUE    **Example:** filter.tags=$eq:John Doe&filter.tags=$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=label:DESC   **Default Value:** label:DESC  **Available Fields** - id  - label  - zoneName  - status  - isDefault  - createdBy  - updatedBy  - createdAt  - updatedAt  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,label,zoneName,status,createdBy   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - zoneName  - status  - createdBy  - updatedBy  (optional)
	select_ := "select__example" // string | List of fields to select.  **Example:** id,label,description,zoneName,zoneType   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.GetDNSZones(context.Background()).Page(page).Limit(limit).FilterLabel(filterLabel).FilterZoneName(filterZoneName).FilterZoneType(filterZoneType).FilterStatus(filterStatus).FilterIsDefault(filterIsDefault).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.GetDNSZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSZones`: DNSZonePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.GetDNSZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$eq:John Doe&amp;filter.label&#x3D;$ilike:John Doe  **Available Operations** - $eq  - $null  - $ilike  - $and  - $or | 
 **filterZoneName** | **[]string** | Filter by zoneName query param.  **Format:** filter.zoneName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.zoneName&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterZoneType** | **[]string** | Filter by zoneType query param.  **Format:** filter.zoneType&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.zoneType&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterIsDefault** | **[]string** | Filter by isDefault query param.  **Format:** filter.isDefault&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.isDefault&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterTags** | **[]string** | Filter by tags query param.  **Format:** filter.tags&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.tags&#x3D;$eq:John Doe&amp;filter.tags&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;label:DESC   **Default Value:** label:DESC  **Available Fields** - id  - label  - zoneName  - status  - isDefault  - createdBy  - updatedBy  - createdAt  - updatedAt  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,label,zoneName,status,createdBy   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - label  - zoneName  - status  - createdBy  - updatedBy  | 
 **select_** | **string** | List of fields to select.  **Example:** id,label,description,zoneName,zoneType   **Default Value:** By default all fields returns. If you want to select only some fields, provide them in query param   | 

### Return type

[**DNSZonePaginatedList**](DNSZonePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDNSRecordSetsByZoneId

> DNSRecordSetPaginatedList ListDNSRecordSetsByZoneId(ctx, id).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterName(filterName).FilterZoneName(filterZoneName).FilterStatus(filterStatus).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()

List RecordSets in a DNS zone by ID



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
	id := float32(8.14) // float32 | The ID of the DNS Zone
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
	resp, r, err := apiClient.DNSZoneAPI.ListDNSRecordSetsByZoneId(context.Background(), id).Page(page).Limit(limit).FilterInfrastructureId(filterInfrastructureId).FilterName(filterName).FilterZoneName(filterZoneName).FilterStatus(filterStatus).FilterTags(filterTags).SortBy(sortBy).Search(search).SearchBy(searchBy).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.ListDNSRecordSetsByZoneId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDNSRecordSetsByZoneId`: DNSRecordSetPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.ListDNSRecordSetsByZoneId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the DNS Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDNSRecordSetsByZoneIdRequest struct via the builder pattern


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


## UpdateDNSZone

> DnsZone UpdateDNSZone(ctx, dnsZoneId).UpdateDnsZone(updateDnsZone).Execute()

Update a DNS Zone by ID



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
	dnsZoneId := float32(8.14) // float32 | The ID of the DNS Zone
	updateDnsZone := *openapiclient.NewUpdateDnsZone() // UpdateDnsZone | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.UpdateDNSZone(context.Background(), dnsZoneId).UpdateDnsZone(updateDnsZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.UpdateDNSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDNSZone`: DnsZone
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.UpdateDNSZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsZoneId** | **float32** | The ID of the DNS Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDNSZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDnsZone** | [**UpdateDnsZone**](UpdateDnsZone.md) |  | 

### Return type

[**DnsZone**](DnsZone.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

