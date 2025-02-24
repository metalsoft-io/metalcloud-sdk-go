# \EventAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventAPI.md#GetEvent) | **Get** /api/v2/events/{customIsoId} | Get Custom Iso information
[**GetEvents**](EventAPI.md#GetEvents) | **Get** /api/v2/events | Get all Events



## GetEvent

> Event GetEvent(ctx, customIsoId).Execute()

Get Custom Iso information



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
	customIsoId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.GetEvent(context.Background(), customIsoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customIsoId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> EventPaginatedList GetEvents(ctx).Page(page).Limit(limit).FilterId(filterId).FilterUserIdAuthenticated(filterUserIdAuthenticated).FilterType(filterType).FilterSeverity(filterSeverity).FilterVisibility(filterVisibility).FilterInfrastructureId(filterInfrastructureId).FilterUserId(filterUserId).FilterStoragePoolId(filterStoragePoolId).FilterServerId(filterServerId).FilterJobId(filterJobId).FilterSiteId(filterSiteId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all Events



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.           <p>              <b>Format: </b> filter.id={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.id=$not:$like:John Doe&filter.id=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li></ul> (optional)
	filterUserIdAuthenticated := []string{"Inner_example"} // []string | Filter by userIdAuthenticated query param.           <p>              <b>Format: </b> filter.userIdAuthenticated={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.userIdAuthenticated=$not:$like:John Doe&filter.userIdAuthenticated=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterType := []string{"Inner_example"} // []string | Filter by type query param.           <p>              <b>Format: </b> filter.type={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.type=$not:$like:John Doe&filter.type=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterSeverity := []string{"Inner_example"} // []string | Filter by severity query param.           <p>              <b>Format: </b> filter.severity={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.severity=$not:$like:John Doe&filter.severity=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterVisibility := []string{"Inner_example"} // []string | Filter by visibility query param.           <p>              <b>Format: </b> filter.visibility={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.visibility=$not:$like:John Doe&filter.visibility=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.           <p>              <b>Format: </b> filter.infrastructureId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.infrastructureId=$not:$like:John Doe&filter.infrastructureId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterUserId := []string{"Inner_example"} // []string | Filter by userId query param.           <p>              <b>Format: </b> filter.userId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.userId=$not:$like:John Doe&filter.userId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterStoragePoolId := []string{"Inner_example"} // []string | Filter by storagePoolId query param.           <p>              <b>Format: </b> filter.storagePoolId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.storagePoolId=$not:$like:John Doe&filter.storagePoolId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.           <p>              <b>Format: </b> filter.serverId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.serverId=$not:$like:John Doe&filter.serverId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterJobId := []string{"Inner_example"} // []string | Filter by jobId query param.           <p>              <b>Format: </b> filter.jobId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.jobId=$not:$like:John Doe&filter.jobId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.           <p>              <b>Format: </b> filter.siteId={$not}:OPERATION:VALUE           </p>           <p>              <b>Example: </b> filter.siteId=$not:$like:John Doe&filter.siteId=like:John           </p>           <h4>Available Operations</h4><ul><li>$eq</li> <li>$in</li></ul> (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by.       <p>To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting</p>       <p>              <b>Format: </b> fieldName:DIRECTION           </p>       <p>              <b>Example: </b> sortBy=id:DESC&sortBy=createdAt:ASC           </p>       <p>              <b>Default Value: </b> id:DESC           </p>       <h4>Available Fields</h4><ul><li>id</li></ul>        (optional)
	search := "search_example" // string | Search term to filter result values         <p>              <b>Example: </b> John           </p>         <p>              <b>Default Value: </b> No default value           </p>          (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values         <p>              <b>Example: </b> id,type,severity,visibility,infrastructureId           </p>         <p>              <b>Default Value: </b> By default all fields mentioned below will be used to search by term           </p>         <h4>Available Fields</h4><ul><li>id</li> <li>type</li> <li>severity</li> <li>visibility</li> <li>infrastructureId</li> <li>userId</li> <li>serverId</li> <li>jobId</li> <li>siteId</li> <li>title</li> <li>message</li></ul>          (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.GetEvents(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterUserIdAuthenticated(filterUserIdAuthenticated).FilterType(filterType).FilterSeverity(filterSeverity).FilterVisibility(filterVisibility).FilterInfrastructureId(filterInfrastructureId).FilterUserId(filterUserId).FilterStoragePoolId(filterStoragePoolId).FilterServerId(filterServerId).FilterJobId(filterJobId).FilterSiteId(filterSiteId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: EventPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;          | 
 **limit** | **float32** | Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.        | 
 **filterId** | **[]string** | Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt; | 
 **filterUserIdAuthenticated** | **[]string** | Filter by userIdAuthenticated query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.userIdAuthenticated&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.userIdAuthenticated&#x3D;$not:$like:John Doe&amp;filter.userIdAuthenticated&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterType** | **[]string** | Filter by type query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.type&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.type&#x3D;$not:$like:John Doe&amp;filter.type&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterSeverity** | **[]string** | Filter by severity query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.severity&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.severity&#x3D;$not:$like:John Doe&amp;filter.severity&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterVisibility** | **[]string** | Filter by visibility query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.visibility&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.visibility&#x3D;$not:$like:John Doe&amp;filter.visibility&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.infrastructureId&#x3D;$not:$like:John Doe&amp;filter.infrastructureId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterUserId** | **[]string** | Filter by userId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.userId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.userId&#x3D;$not:$like:John Doe&amp;filter.userId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterStoragePoolId** | **[]string** | Filter by storagePoolId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.storagePoolId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.storagePoolId&#x3D;$not:$like:John Doe&amp;filter.storagePoolId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterServerId** | **[]string** | Filter by serverId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverId&#x3D;$not:$like:John Doe&amp;filter.serverId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterJobId** | **[]string** | Filter by jobId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.jobId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.jobId&#x3D;$not:$like:John Doe&amp;filter.jobId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **filterSiteId** | **[]string** | Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **[]string** | Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;        | 
 **search** | **string** | Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;          | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,type,severity,visibility,infrastructureId           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;type&lt;/li&gt; &lt;li&gt;severity&lt;/li&gt; &lt;li&gt;visibility&lt;/li&gt; &lt;li&gt;infrastructureId&lt;/li&gt; &lt;li&gt;userId&lt;/li&gt; &lt;li&gt;serverId&lt;/li&gt; &lt;li&gt;jobId&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;title&lt;/li&gt; &lt;li&gt;message&lt;/li&gt;&lt;/ul&gt;          | 

### Return type

[**EventPaginatedList**](EventPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

