/*
MetalSoft REST API

MetalSoft REST API documentation

API version: 2.0
Contact: support@metalsoft.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ExtensionAPIService ExtensionAPI service
type ExtensionAPIService service

type ExtensionAPIArchiveExtensionRequest struct {
	ctx context.Context
	ApiService *ExtensionAPIService
	extensionId float32
	ifMatch *string
}

// Entity tag
func (r ExtensionAPIArchiveExtensionRequest) IfMatch(ifMatch string) ExtensionAPIArchiveExtensionRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ExtensionAPIArchiveExtensionRequest) Execute() (*http.Response, error) {
	return r.ApiService.ArchiveExtensionExecute(r)
}

/*
ArchiveExtension Archive published extension

Archives published extension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param extensionId
 @return ExtensionAPIArchiveExtensionRequest
*/
func (a *ExtensionAPIService) ArchiveExtension(ctx context.Context, extensionId float32) ExtensionAPIArchiveExtensionRequest {
	return ExtensionAPIArchiveExtensionRequest{
		ApiService: a,
		ctx: ctx,
		extensionId: extensionId,
	}
}

// Execute executes the request
func (a *ExtensionAPIService) ArchiveExtensionExecute(r ExtensionAPIArchiveExtensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionAPIService.ArchiveExtension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/extensions/{extensionId}/actions/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ExtensionAPICreateExtensionRequest struct {
	ctx context.Context
	ApiService *ExtensionAPIService
	createExtension *CreateExtension
}

// The extension details
func (r ExtensionAPICreateExtensionRequest) CreateExtension(createExtension CreateExtension) ExtensionAPICreateExtensionRequest {
	r.createExtension = &createExtension
	return r
}

func (r ExtensionAPICreateExtensionRequest) Execute() (*Extension, *http.Response, error) {
	return r.ApiService.CreateExtensionExecute(r)
}

/*
CreateExtension Create extension

Returns details of the new extension

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExtensionAPICreateExtensionRequest
*/
func (a *ExtensionAPIService) CreateExtension(ctx context.Context) ExtensionAPICreateExtensionRequest {
	return ExtensionAPICreateExtensionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Extension
func (a *ExtensionAPIService) CreateExtensionExecute(r ExtensionAPICreateExtensionRequest) (*Extension, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Extension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionAPIService.CreateExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createExtension == nil {
		return localVarReturnValue, nil, reportError("createExtension is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createExtension
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ExtensionAPIGetExtensionRequest struct {
	ctx context.Context
	ApiService *ExtensionAPIService
	extensionId float32
}

func (r ExtensionAPIGetExtensionRequest) Execute() (*Extension, *http.Response, error) {
	return r.ApiService.GetExtensionExecute(r)
}

/*
GetExtension Get details for an extension

Returns details of the specified extension

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param extensionId
 @return ExtensionAPIGetExtensionRequest
*/
func (a *ExtensionAPIService) GetExtension(ctx context.Context, extensionId float32) ExtensionAPIGetExtensionRequest {
	return ExtensionAPIGetExtensionRequest{
		ApiService: a,
		ctx: ctx,
		extensionId: extensionId,
	}
}

// Execute executes the request
//  @return Extension
func (a *ExtensionAPIService) GetExtensionExecute(r ExtensionAPIGetExtensionRequest) (*Extension, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Extension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionAPIService.GetExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/extensions/{extensionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ExtensionAPIGetExtensionsRequest struct {
	ctx context.Context
	ApiService *ExtensionAPIService
	page *float32
	limit *float32
	filterStatus *[]string
	filterName *[]string
	filterLabel *[]string
	sortBy *[]string
	search *string
	searchBy *[]string
}

// Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;         
func (r ExtensionAPIGetExtensionsRequest) Page(page float32) ExtensionAPIGetExtensionsRequest {
	r.page = &page
	return r
}

// Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; -1           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; -1           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.       
func (r ExtensionAPIGetExtensionsRequest) Limit(limit float32) ExtensionAPIGetExtensionsRequest {
	r.limit = &limit
	return r
}

// Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r ExtensionAPIGetExtensionsRequest) FilterStatus(filterStatus []string) ExtensionAPIGetExtensionsRequest {
	r.filterStatus = &filterStatus
	return r
}

// Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r ExtensionAPIGetExtensionsRequest) FilterName(filterName []string) ExtensionAPIGetExtensionsRequest {
	r.filterName = &filterName
	return r
}

// Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt;
func (r ExtensionAPIGetExtensionsRequest) FilterLabel(filterLabel []string) ExtensionAPIGetExtensionsRequest {
	r.filterLabel = &filterLabel
	return r
}

// Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;status&lt;/li&gt;&lt;/ul&gt;       
func (r ExtensionAPIGetExtensionsRequest) SortBy(sortBy []string) ExtensionAPIGetExtensionsRequest {
	r.sortBy = &sortBy
	return r
}

// Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;         
func (r ExtensionAPIGetExtensionsRequest) Search(search string) ExtensionAPIGetExtensionsRequest {
	r.search = &search
	return r
}

// List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt;&lt;/ul&gt;         
func (r ExtensionAPIGetExtensionsRequest) SearchBy(searchBy []string) ExtensionAPIGetExtensionsRequest {
	r.searchBy = &searchBy
	return r
}

func (r ExtensionAPIGetExtensionsRequest) Execute() (*ExtensionList, *http.Response, error) {
	return r.ApiService.GetExtensionsExecute(r)
}

/*
GetExtensions Get a list of available extensions

Returns list of the available extensions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExtensionAPIGetExtensionsRequest
*/
func (a *ExtensionAPIService) GetExtensions(ctx context.Context) ExtensionAPIGetExtensionsRequest {
	return ExtensionAPIGetExtensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExtensionList
func (a *ExtensionAPIService) GetExtensionsExecute(r ExtensionAPIGetExtensionsRequest) (*ExtensionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtensionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionAPIService.GetExtensions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.filterStatus != nil {
		t := *r.filterStatus
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.status", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.status", t, "form", "multi")
		}
	}
	if r.filterName != nil {
		t := *r.filterName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.name", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.name", t, "form", "multi")
		}
	}
	if r.filterLabel != nil {
		t := *r.filterLabel
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.label", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.label", t, "form", "multi")
		}
	}
	if r.sortBy != nil {
		t := *r.sortBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", t, "form", "multi")
		}
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.searchBy != nil {
		t := *r.searchBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "searchBy", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "searchBy", t, "form", "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ExtensionAPIPublishExtensionRequest struct {
	ctx context.Context
	ApiService *ExtensionAPIService
	extensionId float32
	ifMatch *string
}

// Entity tag
func (r ExtensionAPIPublishExtensionRequest) IfMatch(ifMatch string) ExtensionAPIPublishExtensionRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ExtensionAPIPublishExtensionRequest) Execute() (*http.Response, error) {
	return r.ApiService.PublishExtensionExecute(r)
}

/*
PublishExtension Publish draft extension

Publishes draft extension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param extensionId
 @return ExtensionAPIPublishExtensionRequest
*/
func (a *ExtensionAPIService) PublishExtension(ctx context.Context, extensionId float32) ExtensionAPIPublishExtensionRequest {
	return ExtensionAPIPublishExtensionRequest{
		ApiService: a,
		ctx: ctx,
		extensionId: extensionId,
	}
}

// Execute executes the request
func (a *ExtensionAPIService) PublishExtensionExecute(r ExtensionAPIPublishExtensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionAPIService.PublishExtension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/extensions/{extensionId}/actions/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ExtensionAPIUpdateExtensionRequest struct {
	ctx context.Context
	ApiService *ExtensionAPIService
	extensionId float32
	updateExtension *UpdateExtension
	ifMatch *string
}

// The extension details
func (r ExtensionAPIUpdateExtensionRequest) UpdateExtension(updateExtension UpdateExtension) ExtensionAPIUpdateExtensionRequest {
	r.updateExtension = &updateExtension
	return r
}

// Entity tag
func (r ExtensionAPIUpdateExtensionRequest) IfMatch(ifMatch string) ExtensionAPIUpdateExtensionRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ExtensionAPIUpdateExtensionRequest) Execute() (*Extension, *http.Response, error) {
	return r.ApiService.UpdateExtensionExecute(r)
}

/*
UpdateExtension Update extension

Returns details of the updated extension

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param extensionId
 @return ExtensionAPIUpdateExtensionRequest
*/
func (a *ExtensionAPIService) UpdateExtension(ctx context.Context, extensionId float32) ExtensionAPIUpdateExtensionRequest {
	return ExtensionAPIUpdateExtensionRequest{
		ApiService: a,
		ctx: ctx,
		extensionId: extensionId,
	}
}

// Execute executes the request
//  @return Extension
func (a *ExtensionAPIService) UpdateExtensionExecute(r ExtensionAPIUpdateExtensionRequest) (*Extension, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Extension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionAPIService.UpdateExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/extensions/{extensionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateExtension == nil {
		return localVarReturnValue, nil, reportError("updateExtension is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
	}
	// body params
	localVarPostBody = r.updateExtension
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
