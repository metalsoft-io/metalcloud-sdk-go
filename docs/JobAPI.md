# \JobAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCronJob**](JobAPI.md#CreateCronJob) | **Post** /api/v2/cron-jobs | Creates a new cron job
[**DeleteCronJob**](JobAPI.md#DeleteCronJob) | **Delete** /api/v2/cron-jobs/{cronJobId} | Delete cron job
[**GetCronJob**](JobAPI.md#GetCronJob) | **Get** /api/v2/cron-jobs/{cronJobId} | Get cron job information
[**GetCronJobs**](JobAPI.md#GetCronJobs) | **Get** /api/v2/cron-jobs | Get a list of cron jobs
[**GetCronJobsSupportedFunctions**](JobAPI.md#GetCronJobsSupportedFunctions) | **Get** /api/v2/cron-jobs/supported-functions | Get a list of supported functions for cron jobs
[**GetJob**](JobAPI.md#GetJob) | **Get** /api/v2/jobs/{jobId} | Get Job information
[**GetJobExceptions**](JobAPI.md#GetJobExceptions) | **Get** /api/v2/jobs/{jobId}/exceptions | Get a list of Job Exceptions
[**GetJobFromArchive**](JobAPI.md#GetJobFromArchive) | **Get** /api/v2/jobs/archive/{jobId} | Get Job from archive information
[**GetJobGroup**](JobAPI.md#GetJobGroup) | **Get** /api/v2/job-groups/{jobGroupId} | Get Job Group information
[**GetJobGroupStatistics**](JobAPI.md#GetJobGroupStatistics) | **Get** /api/v2/job-groups/{jobGroupId}/statistics | Get Job Group statistics
[**GetJobGroups**](JobAPI.md#GetJobGroups) | **Get** /api/v2/job-groups | Get a list of Job Groups
[**GetJobs**](JobAPI.md#GetJobs) | **Get** /api/v2/jobs | Get a list of Jobs
[**GetJobsFromArchive**](JobAPI.md#GetJobsFromArchive) | **Get** /api/v2/jobs/archive | Get a list of Jobs from archive
[**GetJobsStatistics**](JobAPI.md#GetJobsStatistics) | **Get** /api/v2/jobs/statistics | Get Jobs statistics
[**IssueCommandForJob**](JobAPI.md#IssueCommandForJob) | **Post** /api/v2/jobs/{jobId}/actions/issue-command | Issues a command for a job that changes the operational state of the job
[**RetryJob**](JobAPI.md#RetryJob) | **Post** /api/v2/jobs/{jobId}/actions/retry | Retries a job
[**SkipJob**](JobAPI.md#SkipJob) | **Post** /api/v2/jobs/{jobId}/actions/skip | Skips a job
[**UpdateCronJob**](JobAPI.md#UpdateCronJob) | **Patch** /api/v2/cron-jobs/{cronJobId} | Updates an existing cron job



## CreateCronJob

> CreateCronJob(ctx).CreateCronJob(createCronJob).Execute()

Creates a new cron job



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
	createCronJob := *openapiclient.NewCreateCronJob("Label_example", "FunctionName_example", []map[string]interface{}{map[string]interface{}(123)}, "Schedule_example", float32(123), float32(123), float32(123)) // CreateCronJob | The cron job details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.CreateCronJob(context.Background()).CreateCronJob(createCronJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.CreateCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCronJob** | [**CreateCronJob**](CreateCronJob.md) | The cron job details | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCronJob

> DeleteCronJob(ctx, cronJobId).Execute()

Delete cron job



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
	cronJobId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.DeleteCronJob(context.Background(), cronJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.DeleteCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cronJobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCronJobRequest struct via the builder pattern


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


## GetCronJob

> CronJob GetCronJob(ctx, cronJobId).Execute()

Get cron job information



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
	cronJobId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetCronJob(context.Background(), cronJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCronJob`: CronJob
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetCronJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cronJobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CronJob**](CronJob.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCronJobs

> CronJobPaginatedList GetCronJobs(ctx).Page(page).Limit(limit).FilterLabel(filterLabel).FilterDescription(filterDescription).FilterFunctionName(filterFunctionName).FilterDisabled(filterDisabled).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of cron jobs



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
	filterLabel := []string{"Inner_example"} // []string | Filter by label query param.  **Format:** filter.label={$not}:OPERATION:VALUE    **Example:** filter.label=$ilike:John Doe  **Available Operations** - $ilike  - $and  - $or (optional)
	filterDescription := []string{"Inner_example"} // []string | Filter by description query param.  **Format:** filter.description={$not}:OPERATION:VALUE    **Example:** filter.description=$ilike:John Doe  **Available Operations** - $ilike  - $and  - $or (optional)
	filterFunctionName := []string{"Inner_example"} // []string | Filter by functionName query param.  **Format:** filter.functionName={$not}:OPERATION:VALUE    **Example:** filter.functionName=$ilike:John Doe  **Available Operations** - $ilike  - $and  - $or (optional)
	filterDisabled := []string{"Inner_example"} // []string | Filter by disabled query param.  **Format:** filter.disabled={$not}:OPERATION:VALUE    **Example:** filter.disabled=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=label:DESC&sortBy=functionName:DESC   **Default Value:** id:ASC  **Available Fields** - label  - functionName  - disabled  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** label,description,functionName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  - description  - functionName  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetCronJobs(context.Background()).Page(page).Limit(limit).FilterLabel(filterLabel).FilterDescription(filterDescription).FilterFunctionName(filterFunctionName).FilterDisabled(filterDisabled).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetCronJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCronJobs`: CronJobPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetCronJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCronJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterLabel** | **[]string** | Filter by label query param.  **Format:** filter.label&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.label&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $and  - $or | 
 **filterDescription** | **[]string** | Filter by description query param.  **Format:** filter.description&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.description&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $and  - $or | 
 **filterFunctionName** | **[]string** | Filter by functionName query param.  **Format:** filter.functionName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.functionName&#x3D;$ilike:John Doe  **Available Operations** - $ilike  - $and  - $or | 
 **filterDisabled** | **[]string** | Filter by disabled query param.  **Format:** filter.disabled&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.disabled&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;label:DESC&amp;sortBy&#x3D;functionName:DESC   **Default Value:** id:ASC  **Available Fields** - label  - functionName  - disabled  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** label,description,functionName   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - label  - description  - functionName  | 

### Return type

[**CronJobPaginatedList**](CronJobPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCronJobsSupportedFunctions

> []CronJobFunction GetCronJobsSupportedFunctions(ctx).Execute()

Get a list of supported functions for cron jobs



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
	resp, r, err := apiClient.JobAPI.GetCronJobsSupportedFunctions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetCronJobsSupportedFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCronJobsSupportedFunctions`: []CronJobFunction
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetCronJobsSupportedFunctions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronJobsSupportedFunctionsRequest struct via the builder pattern


### Return type

[**[]CronJobFunction**](CronJobFunction.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> Job GetJob(ctx, jobId).Execute()

Get Job information



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
	jobId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobExceptions

> JobExceptionPaginatedList GetJobExceptions(ctx, jobId).Page(page).Limit(limit).FilterExceptionId(filterExceptionId).FilterArchiveId(filterArchiveId).FilterCreatedTimestamp(filterCreatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Job Exceptions



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
	jobId := float32(8.14) // float32 | 
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterExceptionId := []string{"Inner_example"} // []string | Filter by exceptionId query param.  **Format:** filter.exceptionId={$not}:OPERATION:VALUE    **Example:** filter.exceptionId=$btw:John Doe&filter.exceptionId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterArchiveId := []string{"Inner_example"} // []string | Filter by archiveId query param.  **Format:** filter.archiveId={$not}:OPERATION:VALUE    **Example:** filter.archiveId=$btw:John Doe&filter.archiveId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp={$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp=$btw:John Doe&filter.createdTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=exceptionId:DESC&sortBy=createdTimestamp:DESC   **Default Value:** exceptionId:DESC  **Available Fields** - exceptionId  - createdTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** exceptionId,archiveId,createdTimestamp   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - exceptionId  - archiveId  - createdTimestamp  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobExceptions(context.Background(), jobId).Page(page).Limit(limit).FilterExceptionId(filterExceptionId).FilterArchiveId(filterArchiveId).FilterCreatedTimestamp(filterCreatedTimestamp).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobExceptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobExceptions`: JobExceptionPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterExceptionId** | **[]string** | Filter by exceptionId query param.  **Format:** filter.exceptionId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.exceptionId&#x3D;$btw:John Doe&amp;filter.exceptionId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterArchiveId** | **[]string** | Filter by archiveId query param.  **Format:** filter.archiveId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.archiveId&#x3D;$btw:John Doe&amp;filter.archiveId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp&#x3D;$btw:John Doe&amp;filter.createdTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;exceptionId:DESC&amp;sortBy&#x3D;createdTimestamp:DESC   **Default Value:** exceptionId:DESC  **Available Fields** - exceptionId  - createdTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** exceptionId,archiveId,createdTimestamp   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - exceptionId  - archiveId  - createdTimestamp  | 

### Return type

[**JobExceptionPaginatedList**](JobExceptionPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobFromArchive

> JobArchive GetJobFromArchive(ctx, jobId).Execute()

Get Job from archive information



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
	jobId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobFromArchive(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobFromArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobFromArchive`: JobArchive
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobFromArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobFromArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobArchive**](JobArchive.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobGroup

> JobGroup GetJobGroup(ctx, jobGroupId).Execute()

Get Job Group information



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
	jobGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobGroup(context.Background(), jobGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobGroup`: JobGroup
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobGroup**](JobGroup.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobGroupStatistics

> JobGroupStatistics GetJobGroupStatistics(ctx, jobGroupId).Execute()

Get Job Group statistics



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
	jobGroupId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobGroupStatistics(context.Background(), jobGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobGroupStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobGroupStatistics`: JobGroupStatistics
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobGroupStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobGroupId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobGroupStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobGroupStatistics**](JobGroupStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobGroups

> JobGroupPaginatedList GetJobGroups(ctx).Page(page).Limit(limit).FilterId(filterId).FilterType(filterType).FilterCreatedTimestamp(filterCreatedTimestamp).FilterFinishedTimestamp(filterFinishedTimestamp).FilterServerId(filterServerId).FilterDriveId(filterDriveId).FilterInfrastructureId(filterInfrastructureId).FilterVmPoolId(filterVmPoolId).FilterStorageId(filterStorageId).FilterNetworkDeviceId(filterNetworkDeviceId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Job Groups



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
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$btw:John Doe&filter.id=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterType := []string{"Inner_example"} // []string | Filter by type query param.  **Format:** filter.type={$not}:OPERATION:VALUE    **Example:** filter.type=$btw:John Doe&filter.type=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterCreatedTimestamp := []string{"Inner_example"} // []string | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp={$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp=$btw:John Doe&filter.createdTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFinishedTimestamp := []string{"Inner_example"} // []string | Filter by finishedTimestamp query param.  **Format:** filter.finishedTimestamp={$not}:OPERATION:VALUE    **Example:** filter.finishedTimestamp=$btw:John Doe&filter.finishedTimestamp=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$btw:John Doe&filter.serverId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterDriveId := []string{"Inner_example"} // []string | Filter by driveId query param.  **Format:** filter.driveId={$not}:OPERATION:VALUE    **Example:** filter.driveId=$btw:John Doe&filter.driveId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$btw:John Doe&filter.infrastructureId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterVmPoolId := []string{"Inner_example"} // []string | Filter by vmPoolId query param.  **Format:** filter.vmPoolId={$not}:OPERATION:VALUE    **Example:** filter.vmPoolId=$btw:John Doe&filter.vmPoolId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStorageId := []string{"Inner_example"} // []string | Filter by storageId query param.  **Format:** filter.storageId={$not}:OPERATION:VALUE    **Example:** filter.storageId=$btw:John Doe&filter.storageId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNetworkDeviceId := []string{"Inner_example"} // []string | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId=$btw:John Doe&filter.networkDeviceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=createdTimestamp:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdTimestamp  - finishedTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,type,serverId,driveId,infrastructureId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - type  - serverId  - driveId  - infrastructureId  - vmPoolId  - storageId  - networkDeviceId  - createdTimestamp  - finishedTimestamp  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobGroups(context.Background()).Page(page).Limit(limit).FilterId(filterId).FilterType(filterType).FilterCreatedTimestamp(filterCreatedTimestamp).FilterFinishedTimestamp(filterFinishedTimestamp).FilterServerId(filterServerId).FilterDriveId(filterDriveId).FilterInfrastructureId(filterInfrastructureId).FilterVmPoolId(filterVmPoolId).FilterStorageId(filterStorageId).FilterNetworkDeviceId(filterNetworkDeviceId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobGroups`: JobGroupPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$btw:John Doe&amp;filter.id&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterType** | **[]string** | Filter by type query param.  **Format:** filter.type&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.type&#x3D;$btw:John Doe&amp;filter.type&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterCreatedTimestamp** | **[]string** | Filter by createdTimestamp query param.  **Format:** filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.createdTimestamp&#x3D;$btw:John Doe&amp;filter.createdTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFinishedTimestamp** | **[]string** | Filter by finishedTimestamp query param.  **Format:** filter.finishedTimestamp&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.finishedTimestamp&#x3D;$btw:John Doe&amp;filter.finishedTimestamp&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$btw:John Doe&amp;filter.serverId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterDriveId** | **[]string** | Filter by driveId query param.  **Format:** filter.driveId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.driveId&#x3D;$btw:John Doe&amp;filter.driveId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$btw:John Doe&amp;filter.infrastructureId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVmPoolId** | **[]string** | Filter by vmPoolId query param.  **Format:** filter.vmPoolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vmPoolId&#x3D;$btw:John Doe&amp;filter.vmPoolId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStorageId** | **[]string** | Filter by storageId query param.  **Format:** filter.storageId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.storageId&#x3D;$btw:John Doe&amp;filter.storageId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNetworkDeviceId** | **[]string** | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId&#x3D;$btw:John Doe&amp;filter.networkDeviceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdTimestamp:DESC   **Default Value:** id:DESC  **Available Fields** - id  - createdTimestamp  - finishedTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,type,serverId,driveId,infrastructureId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - type  - serverId  - driveId  - infrastructureId  - vmPoolId  - storageId  - networkDeviceId  - createdTimestamp  - finishedTimestamp  | 

### Return type

[**JobGroupPaginatedList**](JobGroupPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobs

> JobPaginatedList GetJobs(ctx).Page(page).Limit(limit).FilterJobId(filterJobId).FilterInfrastructureId(filterInfrastructureId).FilterServerId(filterServerId).FilterVmPoolId(filterVmPoolId).FilterStorageId(filterStorageId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterJobGroupId(filterJobGroupId).FilterInstanceId(filterInstanceId).FilterFunctionName(filterFunctionName).FilterSiteId(filterSiteId).FilterJobIdBlocked(filterJobIdBlocked).FilterJobIdBlockedBy(filterJobIdBlockedBy).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Jobs



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
	filterJobId := []string{"Inner_example"} // []string | Filter by jobId query param.  **Format:** filter.jobId={$not}:OPERATION:VALUE    **Example:** filter.jobId=$btw:John Doe&filter.jobId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$btw:John Doe&filter.infrastructureId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$btw:John Doe&filter.serverId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterVmPoolId := []string{"Inner_example"} // []string | Filter by vmPoolId query param.  **Format:** filter.vmPoolId={$not}:OPERATION:VALUE    **Example:** filter.vmPoolId=$btw:John Doe&filter.vmPoolId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStorageId := []string{"Inner_example"} // []string | Filter by storageId query param.  **Format:** filter.storageId={$not}:OPERATION:VALUE    **Example:** filter.storageId=$btw:John Doe&filter.storageId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNetworkDeviceId := []string{"Inner_example"} // []string | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId=$btw:John Doe&filter.networkDeviceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterJobGroupId := []string{"Inner_example"} // []string | Filter by jobGroupId query param.  **Format:** filter.jobGroupId={$not}:OPERATION:VALUE    **Example:** filter.jobGroupId=$btw:John Doe&filter.jobGroupId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInstanceId := []string{"Inner_example"} // []string | Filter by instanceId query param.  **Format:** filter.instanceId={$not}:OPERATION:VALUE    **Example:** filter.instanceId=$btw:John Doe&filter.instanceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFunctionName := []string{"Inner_example"} // []string | Filter by functionName query param.  **Format:** filter.functionName={$not}:OPERATION:VALUE    **Example:** filter.functionName=$btw:John Doe&filter.functionName=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterJobIdBlocked := []string{"Inner_example"} // []string | Filter by jobIdBlocked query param.  **Format:** filter.jobIdBlocked={$not}:OPERATION:VALUE    **Example:** filter.jobIdBlocked=$btw:John Doe&filter.jobIdBlocked=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterJobIdBlockedBy := []string{"Inner_example"} // []string | Filter by jobIdBlockedBy query param.  **Format:** filter.jobIdBlockedBy={$not}:OPERATION:VALUE    **Example:** filter.jobIdBlockedBy=$btw:John Doe&filter.jobIdBlockedBy=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=jobId:DESC&sortBy=jobGroupId:DESC   **Default Value:** jobId:DESC  **Available Fields** - jobId  - jobGroupId  - status  - createdTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** jobId,infrastructureId,serverId,vmPoolId,storageId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - jobId  - infrastructureId  - serverId  - vmPoolId  - storageId  - networkDeviceId  - jobGroupId  - functionName  - siteId  - instanceId  - jobIdBlocked  - jobIdBlockedBy  - status  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobs(context.Background()).Page(page).Limit(limit).FilterJobId(filterJobId).FilterInfrastructureId(filterInfrastructureId).FilterServerId(filterServerId).FilterVmPoolId(filterVmPoolId).FilterStorageId(filterStorageId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterJobGroupId(filterJobGroupId).FilterInstanceId(filterInstanceId).FilterFunctionName(filterFunctionName).FilterSiteId(filterSiteId).FilterJobIdBlocked(filterJobIdBlocked).FilterJobIdBlockedBy(filterJobIdBlockedBy).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobs`: JobPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterJobId** | **[]string** | Filter by jobId query param.  **Format:** filter.jobId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobId&#x3D;$btw:John Doe&amp;filter.jobId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$btw:John Doe&amp;filter.infrastructureId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$btw:John Doe&amp;filter.serverId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVmPoolId** | **[]string** | Filter by vmPoolId query param.  **Format:** filter.vmPoolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vmPoolId&#x3D;$btw:John Doe&amp;filter.vmPoolId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStorageId** | **[]string** | Filter by storageId query param.  **Format:** filter.storageId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.storageId&#x3D;$btw:John Doe&amp;filter.storageId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNetworkDeviceId** | **[]string** | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId&#x3D;$btw:John Doe&amp;filter.networkDeviceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterJobGroupId** | **[]string** | Filter by jobGroupId query param.  **Format:** filter.jobGroupId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobGroupId&#x3D;$btw:John Doe&amp;filter.jobGroupId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInstanceId** | **[]string** | Filter by instanceId query param.  **Format:** filter.instanceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.instanceId&#x3D;$btw:John Doe&amp;filter.instanceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFunctionName** | **[]string** | Filter by functionName query param.  **Format:** filter.functionName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.functionName&#x3D;$btw:John Doe&amp;filter.functionName&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterJobIdBlocked** | **[]string** | Filter by jobIdBlocked query param.  **Format:** filter.jobIdBlocked&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobIdBlocked&#x3D;$btw:John Doe&amp;filter.jobIdBlocked&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterJobIdBlockedBy** | **[]string** | Filter by jobIdBlockedBy query param.  **Format:** filter.jobIdBlockedBy&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobIdBlockedBy&#x3D;$btw:John Doe&amp;filter.jobIdBlockedBy&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;jobId:DESC&amp;sortBy&#x3D;jobGroupId:DESC   **Default Value:** jobId:DESC  **Available Fields** - jobId  - jobGroupId  - status  - createdTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** jobId,infrastructureId,serverId,vmPoolId,storageId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - jobId  - infrastructureId  - serverId  - vmPoolId  - storageId  - networkDeviceId  - jobGroupId  - functionName  - siteId  - instanceId  - jobIdBlocked  - jobIdBlockedBy  - status  | 

### Return type

[**JobPaginatedList**](JobPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobsFromArchive

> JobArchivePaginatedList GetJobsFromArchive(ctx).Page(page).Limit(limit).FilterJobId(filterJobId).FilterInfrastructureId(filterInfrastructureId).FilterServerId(filterServerId).FilterVmPoolId(filterVmPoolId).FilterStorageId(filterStorageId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterJobGroupId(filterJobGroupId).FilterInstanceId(filterInstanceId).FilterFunctionName(filterFunctionName).FilterSiteId(filterSiteId).FilterJobIdBlocked(filterJobIdBlocked).FilterJobIdBlockedBy(filterJobIdBlockedBy).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get a list of Jobs from archive



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
	filterJobId := []string{"Inner_example"} // []string | Filter by jobId query param.  **Format:** filter.jobId={$not}:OPERATION:VALUE    **Example:** filter.jobId=$btw:John Doe&filter.jobId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInfrastructureId := []string{"Inner_example"} // []string | Filter by infrastructureId query param.  **Format:** filter.infrastructureId={$not}:OPERATION:VALUE    **Example:** filter.infrastructureId=$btw:John Doe&filter.infrastructureId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterServerId := []string{"Inner_example"} // []string | Filter by serverId query param.  **Format:** filter.serverId={$not}:OPERATION:VALUE    **Example:** filter.serverId=$btw:John Doe&filter.serverId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterVmPoolId := []string{"Inner_example"} // []string | Filter by vmPoolId query param.  **Format:** filter.vmPoolId={$not}:OPERATION:VALUE    **Example:** filter.vmPoolId=$btw:John Doe&filter.vmPoolId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStorageId := []string{"Inner_example"} // []string | Filter by storageId query param.  **Format:** filter.storageId={$not}:OPERATION:VALUE    **Example:** filter.storageId=$btw:John Doe&filter.storageId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterNetworkDeviceId := []string{"Inner_example"} // []string | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId={$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId=$btw:John Doe&filter.networkDeviceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterJobGroupId := []string{"Inner_example"} // []string | Filter by jobGroupId query param.  **Format:** filter.jobGroupId={$not}:OPERATION:VALUE    **Example:** filter.jobGroupId=$btw:John Doe&filter.jobGroupId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterInstanceId := []string{"Inner_example"} // []string | Filter by instanceId query param.  **Format:** filter.instanceId={$not}:OPERATION:VALUE    **Example:** filter.instanceId=$btw:John Doe&filter.instanceId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterFunctionName := []string{"Inner_example"} // []string | Filter by functionName query param.  **Format:** filter.functionName={$not}:OPERATION:VALUE    **Example:** filter.functionName=$btw:John Doe&filter.functionName=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterSiteId := []string{"Inner_example"} // []string | Filter by siteId query param.  **Format:** filter.siteId={$not}:OPERATION:VALUE    **Example:** filter.siteId=$btw:John Doe&filter.siteId=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterJobIdBlocked := []string{"Inner_example"} // []string | Filter by jobIdBlocked query param.  **Format:** filter.jobIdBlocked={$not}:OPERATION:VALUE    **Example:** filter.jobIdBlocked=$btw:John Doe&filter.jobIdBlocked=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterJobIdBlockedBy := []string{"Inner_example"} // []string | Filter by jobIdBlockedBy query param.  **Format:** filter.jobIdBlockedBy={$not}:OPERATION:VALUE    **Example:** filter.jobIdBlockedBy=$btw:John Doe&filter.jobIdBlockedBy=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	filterStatus := []string{"Inner_example"} // []string | Filter by status query param.  **Format:** filter.status={$not}:OPERATION:VALUE    **Example:** filter.status=$btw:John Doe&filter.status=$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=jobId:DESC&sortBy=jobGroupId:DESC   **Default Value:** jobId:DESC  **Available Fields** - jobId  - jobGroupId  - status  - createdTimestamp  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** jobId,infrastructureId,serverId,vmPoolId,storageId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - jobId  - infrastructureId  - serverId  - vmPoolId  - storageId  - networkDeviceId  - jobGroupId  - functionName  - siteId  - instanceId  - jobIdBlocked  - jobIdBlockedBy  - status  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobsFromArchive(context.Background()).Page(page).Limit(limit).FilterJobId(filterJobId).FilterInfrastructureId(filterInfrastructureId).FilterServerId(filterServerId).FilterVmPoolId(filterVmPoolId).FilterStorageId(filterStorageId).FilterNetworkDeviceId(filterNetworkDeviceId).FilterJobGroupId(filterJobGroupId).FilterInstanceId(filterInstanceId).FilterFunctionName(filterFunctionName).FilterSiteId(filterSiteId).FilterJobIdBlocked(filterJobIdBlocked).FilterJobIdBlockedBy(filterJobIdBlockedBy).FilterStatus(filterStatus).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobsFromArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobsFromArchive`: JobArchivePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobsFromArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsFromArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterJobId** | **[]string** | Filter by jobId query param.  **Format:** filter.jobId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobId&#x3D;$btw:John Doe&amp;filter.jobId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInfrastructureId** | **[]string** | Filter by infrastructureId query param.  **Format:** filter.infrastructureId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.infrastructureId&#x3D;$btw:John Doe&amp;filter.infrastructureId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterServerId** | **[]string** | Filter by serverId query param.  **Format:** filter.serverId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.serverId&#x3D;$btw:John Doe&amp;filter.serverId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterVmPoolId** | **[]string** | Filter by vmPoolId query param.  **Format:** filter.vmPoolId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.vmPoolId&#x3D;$btw:John Doe&amp;filter.vmPoolId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStorageId** | **[]string** | Filter by storageId query param.  **Format:** filter.storageId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.storageId&#x3D;$btw:John Doe&amp;filter.storageId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterNetworkDeviceId** | **[]string** | Filter by networkDeviceId query param.  **Format:** filter.networkDeviceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.networkDeviceId&#x3D;$btw:John Doe&amp;filter.networkDeviceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterJobGroupId** | **[]string** | Filter by jobGroupId query param.  **Format:** filter.jobGroupId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobGroupId&#x3D;$btw:John Doe&amp;filter.jobGroupId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterInstanceId** | **[]string** | Filter by instanceId query param.  **Format:** filter.instanceId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.instanceId&#x3D;$btw:John Doe&amp;filter.instanceId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterFunctionName** | **[]string** | Filter by functionName query param.  **Format:** filter.functionName&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.functionName&#x3D;$btw:John Doe&amp;filter.functionName&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterSiteId** | **[]string** | Filter by siteId query param.  **Format:** filter.siteId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.siteId&#x3D;$btw:John Doe&amp;filter.siteId&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterJobIdBlocked** | **[]string** | Filter by jobIdBlocked query param.  **Format:** filter.jobIdBlocked&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobIdBlocked&#x3D;$btw:John Doe&amp;filter.jobIdBlocked&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterJobIdBlockedBy** | **[]string** | Filter by jobIdBlockedBy query param.  **Format:** filter.jobIdBlockedBy&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.jobIdBlockedBy&#x3D;$btw:John Doe&amp;filter.jobIdBlockedBy&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **filterStatus** | **[]string** | Filter by status query param.  **Format:** filter.status&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.status&#x3D;$btw:John Doe&amp;filter.status&#x3D;$contains:John Doe  **Available Operations** - $eq  - $gt  - $gte  - $in  - $null  - $lt  - $lte  - $btw  - $ilike  - $sw  - $contains  - $not  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;jobId:DESC&amp;sortBy&#x3D;jobGroupId:DESC   **Default Value:** jobId:DESC  **Available Fields** - jobId  - jobGroupId  - status  - createdTimestamp  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** jobId,infrastructureId,serverId,vmPoolId,storageId   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - jobId  - infrastructureId  - serverId  - vmPoolId  - storageId  - networkDeviceId  - jobGroupId  - functionName  - siteId  - instanceId  - jobIdBlocked  - jobIdBlockedBy  - status  | 

### Return type

[**JobArchivePaginatedList**](JobArchivePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobsStatistics

> JobStatistics GetJobsStatistics(ctx).Execute()

Get Jobs statistics



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
	resp, r, err := apiClient.JobAPI.GetJobsStatistics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobsStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobsStatistics`: JobStatistics
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobsStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsStatisticsRequest struct via the builder pattern


### Return type

[**JobStatistics**](JobStatistics.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCommandForJob

> IssueCommandForJob(ctx, jobId).JobCommandInfo(jobCommandInfo).Execute()

Issues a command for a job that changes the operational state of the job



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
	jobId := float32(8.14) // float32 | 
	jobCommandInfo := *openapiclient.NewJobCommandInfo() // JobCommandInfo | The job retry options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.IssueCommandForJob(context.Background(), jobId).JobCommandInfo(jobCommandInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.IssueCommandForJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCommandForJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobCommandInfo** | [**JobCommandInfo**](JobCommandInfo.md) | The job retry options | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryJob

> RetryJob(ctx, jobId).JobRetryInfo(jobRetryInfo).Execute()

Retries a job



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
	jobId := float32(8.14) // float32 | 
	jobRetryInfo := *openapiclient.NewJobRetryInfo() // JobRetryInfo | The job retry options

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.RetryJob(context.Background(), jobId).JobRetryInfo(jobRetryInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.RetryJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobRetryInfo** | [**JobRetryInfo**](JobRetryInfo.md) | The job retry options | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipJob

> SkipJob(ctx, jobId).Execute()

Skips a job



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
	jobId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.SkipJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.SkipJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkipJobRequest struct via the builder pattern


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


## UpdateCronJob

> UpdateCronJob(ctx, cronJobId).UpdateCronJob(updateCronJob).Execute()

Updates an existing cron job



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
	cronJobId := float32(8.14) // float32 | 
	updateCronJob := *openapiclient.NewUpdateCronJob() // UpdateCronJob | The cron job details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobAPI.UpdateCronJob(context.Background(), cronJobId).UpdateCronJob(updateCronJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.UpdateCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cronJobId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCronJob** | [**UpdateCronJob**](UpdateCronJob.md) | The cron job details | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

