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

> DnsZoneDto CreateDNSZone(ctx).CreateDnsZoneDto(createDnsZoneDto).Execute()

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
	createDnsZoneDto := *openapiclient.NewCreateDnsZoneDto("example.com", true, []string{"NameServers_example"}) // CreateDnsZoneDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.CreateDNSZone(context.Background()).CreateDnsZoneDto(createDnsZoneDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.CreateDNSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDNSZone`: DnsZoneDto
	fmt.Fprintf(os.Stdout, "Response from `DNSZoneAPI.CreateDNSZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDNSZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDnsZoneDto** | [**CreateDnsZoneDto**](CreateDnsZoneDto.md) |  | 

### Return type

[**DnsZoneDto**](DnsZoneDto.md)

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

> DnsZoneDto GetDNSZoneById(ctx, dnsZoneId).Execute()

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
	// response from `GetDNSZoneById`: DnsZoneDto
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

[**DnsZoneDto**](DnsZoneDto.md)

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.           <p>              <b>Format: </b> filter.label={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.label=$not:$like:John Doe&filter.label=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$null</li> <li>$ilike</li></ul> (optional)
	filterZoneName := []string{"Inner_example"} // []string | Filter by zoneName query param.           <p>              <b>Format: </b> filter.zoneName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.zoneName=$not:$like:John Doe&filter.zoneName=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterZoneType := []string{"Inner_example"} // []string | Filter by zoneType query param.           <p>              <b>Format: </b> filter.zoneType={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.zoneType=$not:$like:John Doe&filter.zoneType=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.           <p>              <b>Format: </b> filter.status={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.status=$not:$like:John Doe&filter.status=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterIsDefault := []string{"Inner_example"} // []string | Filter by isDefault query param.           <p>              <b>Format: </b> filter.isDefault={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.isDefault=$not:$like:John Doe&filter.isDefault=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.           <p>              <b>Format: </b> filter.tags={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.tags=$not:$like:John Doe&filter.tags=like:John           </p>           <h4>Available Operations</h4><ul><li>$ilike</li> <li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> label:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>zoneName</li> <li>status</li> <li>isDefault</li> <li>createdBy</li> <li>updatedBy</li> <li>createdAt</li> <li>updatedAt</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,label,zoneName,status,createdBy           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>label</li> <li>zoneName</li> <li>status</li> <li>createdBy</li> <li>updatedBy</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> id,label,description,zoneName,zoneType           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

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
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterLabel** | **[]string** | Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt; | 
 **filterZoneName** | **[]string** | Filter by zoneName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.zoneName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.zoneName&#x3D;$not:$like:John Doe&amp;filter.zoneName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterZoneType** | **[]string** | Filter by zoneType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.zoneType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.zoneType&#x3D;$not:$like:John Doe&amp;filter.zoneType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterStatus** | **[]string** | Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterIsDefault** | **[]string** | Filter by isDefault query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.isDefault&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.isDefault&#x3D;$not:$like:John Doe&amp;filter.isDefault&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterTags** | **[]string** | Filter by tags query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.tags&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.tags&#x3D;$not:$like:John Doe&amp;filter.tags&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; label:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;zoneName&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;isDefault&lt;/li&gt; &lt;li&gt;createdBy&lt;/li&gt; &lt;li&gt;updatedBy&lt;/li&gt; &lt;li&gt;createdAt&lt;/li&gt; &lt;li&gt;updatedAt&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,zoneName,status,createdBy           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;zoneName&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;createdBy&lt;/li&gt; &lt;li&gt;updatedBy&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,label,description,zoneName,zoneType           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

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
	page := float32(8.14) // float32 | Page number to retrieve.If you provide invalid value the default page number will applied         <p>              <b>Example: </b> 1           </p>         <p>              <b>Default Value: </b> 1           </p>          (optional)
	limit := float32(8.14) // float32 | Number of records per page.       <p>              <b>Example: </b> 20           </p>       <p>              <b>Default Value: </b> 20           </p>       <p>              <b>Max Value: </b> 100           </p>        If provided value is greater than max value, max value will be applied.        (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterName := []string{"Inner_example"} // []string | Filter by name query param.           <p>              <b>Format: </b> filter.name={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.name=$not:$like:John Doe&filter.name=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$ilike</li></ul> (optional)
	filterZoneName := []string{"Inner_example"} // []string | Filter by zoneName query param.           <p>              <b>Format: </b> filter.zoneName={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.zoneName=$not:$like:John Doe&filter.zoneName=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.           <p>              <b>Format: </b> filter.status={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.status=$not:$like:John Doe&filter.status=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterTags := []string{"Inner_example"} // []string | Filter by tags query param.           <p>              <b>Format: </b> filter.tags={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.tags=$not:$like:John Doe&filter.tags=like:John           </p>           <h4>Available Operations</h4><ul><li>$ilike</li> <li>$eq</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> name:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li> <li>status</li> <li>siteId</li> <li>infrastructureId</li> <li>name</li> <li>zoneName</li> <li>type</li> <li>ttl</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,siteId,infrastructureId,name,zoneName           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>siteId</li> <li>infrastructureId</li> <li>name</li> <li>zoneName</li></ul>          (optional)
	select_ := "select__example" // string | List of fields to select.       <p>              <b>Example: </b> id,status,siteId,infrastructureId,name           </p>       <p>              <b>Default Value: </b> By default all fields returns. If you want to select only some fields, provide them in query param           </p>        (optional)

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

 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterName** | **[]string** | Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt; | 
 **filterZoneName** | **[]string** | Filter by zoneName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.zoneName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.zoneName&#x3D;$not:$like:John Doe&amp;filter.zoneName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterStatus** | **[]string** | Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterTags** | **[]string** | Filter by tags query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.tags&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.tags&#x3D;$not:$like:John Doe&amp;filter.tags&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; name:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;zoneName&lt;/li&gt; &lt;li&gt;type&lt;/li&gt; &lt;li&gt;ttl&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,siteId,infrastructureId,name,zoneName           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;zoneName&lt;/li&gt;&lt;/ul&gt;          | 
 **select_** | **string** | List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,status,siteId,infrastructureId,name           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;        | 

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

> DnsZoneDto UpdateDNSZone(ctx, dnsZoneId).UpdateDnsZoneDto(updateDnsZoneDto).Execute()

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
	updateDnsZoneDto := *openapiclient.NewUpdateDnsZoneDto() // UpdateDnsZoneDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSZoneAPI.UpdateDNSZone(context.Background(), dnsZoneId).UpdateDnsZoneDto(updateDnsZoneDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSZoneAPI.UpdateDNSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDNSZone`: DnsZoneDto
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

 **updateDnsZoneDto** | [**UpdateDnsZoneDto**](UpdateDnsZoneDto.md) |  | 

### Return type

[**DnsZoneDto**](DnsZoneDto.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

