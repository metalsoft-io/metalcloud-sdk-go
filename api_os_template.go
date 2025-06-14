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


// OSTemplateAPIService OSTemplateAPI service
type OSTemplateAPIService service

type OSTemplateAPICreateOSTemplateRequest struct {
	ctx context.Context
	ApiService *OSTemplateAPIService
	oSTemplateCreate *OSTemplateCreate
}

// The OS template details
func (r OSTemplateAPICreateOSTemplateRequest) OSTemplateCreate(oSTemplateCreate OSTemplateCreate) OSTemplateAPICreateOSTemplateRequest {
	r.oSTemplateCreate = &oSTemplateCreate
	return r
}

func (r OSTemplateAPICreateOSTemplateRequest) Execute() (*OSTemplate, *http.Response, error) {
	return r.ApiService.CreateOSTemplateExecute(r)
}

/*
CreateOSTemplate Create OS template

Returns details of the new OS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OSTemplateAPICreateOSTemplateRequest
*/
func (a *OSTemplateAPIService) CreateOSTemplate(ctx context.Context) OSTemplateAPICreateOSTemplateRequest {
	return OSTemplateAPICreateOSTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OSTemplate
func (a *OSTemplateAPIService) CreateOSTemplateExecute(r OSTemplateAPICreateOSTemplateRequest) (*OSTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OSTemplateAPIService.CreateOSTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/os-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oSTemplateCreate == nil {
		return localVarReturnValue, nil, reportError("oSTemplateCreate is required and must be specified")
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
	localVarPostBody = r.oSTemplateCreate
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

type OSTemplateAPIDeleteOSTemplateRequest struct {
	ctx context.Context
	ApiService *OSTemplateAPIService
	osTemplateId float32
}

func (r OSTemplateAPIDeleteOSTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOSTemplateExecute(r)
}

/*
DeleteOSTemplate Delete OS template

Deletes the specified OS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param osTemplateId
 @return OSTemplateAPIDeleteOSTemplateRequest
*/
func (a *OSTemplateAPIService) DeleteOSTemplate(ctx context.Context, osTemplateId float32) OSTemplateAPIDeleteOSTemplateRequest {
	return OSTemplateAPIDeleteOSTemplateRequest{
		ApiService: a,
		ctx: ctx,
		osTemplateId: osTemplateId,
	}
}

// Execute executes the request
func (a *OSTemplateAPIService) DeleteOSTemplateExecute(r OSTemplateAPIDeleteOSTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OSTemplateAPIService.DeleteOSTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/os-templates/{osTemplateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"osTemplateId"+"}", url.PathEscape(parameterValueToString(r.osTemplateId, "osTemplateId")), -1)

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

type OSTemplateAPIGetOSTemplateRequest struct {
	ctx context.Context
	ApiService *OSTemplateAPIService
	osTemplateId float32
}

func (r OSTemplateAPIGetOSTemplateRequest) Execute() (*OSTemplate, *http.Response, error) {
	return r.ApiService.GetOSTemplateExecute(r)
}

/*
GetOSTemplate Get details for an OS template

Returns details of the specified OS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param osTemplateId
 @return OSTemplateAPIGetOSTemplateRequest
*/
func (a *OSTemplateAPIService) GetOSTemplate(ctx context.Context, osTemplateId float32) OSTemplateAPIGetOSTemplateRequest {
	return OSTemplateAPIGetOSTemplateRequest{
		ApiService: a,
		ctx: ctx,
		osTemplateId: osTemplateId,
	}
}

// Execute executes the request
//  @return OSTemplate
func (a *OSTemplateAPIService) GetOSTemplateExecute(r OSTemplateAPIGetOSTemplateRequest) (*OSTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OSTemplateAPIService.GetOSTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/os-templates/{osTemplateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"osTemplateId"+"}", url.PathEscape(parameterValueToString(r.osTemplateId, "osTemplateId")), -1)

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

type OSTemplateAPIGetOSTemplateCredentialsRequest struct {
	ctx context.Context
	ApiService *OSTemplateAPIService
	osTemplateId float32
}

func (r OSTemplateAPIGetOSTemplateCredentialsRequest) Execute() (*OSTemplateOsCredential, *http.Response, error) {
	return r.ApiService.GetOSTemplateCredentialsExecute(r)
}

/*
GetOSTemplateCredentials Get OS template credentials

Returns the credentials for the specified OS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param osTemplateId
 @return OSTemplateAPIGetOSTemplateCredentialsRequest
*/
func (a *OSTemplateAPIService) GetOSTemplateCredentials(ctx context.Context, osTemplateId float32) OSTemplateAPIGetOSTemplateCredentialsRequest {
	return OSTemplateAPIGetOSTemplateCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		osTemplateId: osTemplateId,
	}
}

// Execute executes the request
//  @return OSTemplateOsCredential
func (a *OSTemplateAPIService) GetOSTemplateCredentialsExecute(r OSTemplateAPIGetOSTemplateCredentialsRequest) (*OSTemplateOsCredential, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSTemplateOsCredential
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OSTemplateAPIService.GetOSTemplateCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/os-templates/{osTemplateId}/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"osTemplateId"+"}", url.PathEscape(parameterValueToString(r.osTemplateId, "osTemplateId")), -1)

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

type OSTemplateAPIGetOSTemplatesRequest struct {
	ctx context.Context
	ApiService *OSTemplateAPIService
	page *float32
	limit *float32
	filterDeviceType *[]string
	filterLabel *[]string
	filterOsName *[]string
	filterInstallOnieStrings *[]string
	filterStatus *[]string
	filterVisibility *[]string
	filterCreatedBy *[]string
	filterModifiedBy *[]string
	filterTags *[]string
	sortBy *[]string
	search *string
	searchBy *[]string
	select_ *string
}

// Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;         
func (r OSTemplateAPIGetOSTemplatesRequest) Page(page float32) OSTemplateAPIGetOSTemplatesRequest {
	r.page = &page
	return r
}

// Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.       
func (r OSTemplateAPIGetOSTemplatesRequest) Limit(limit float32) OSTemplateAPIGetOSTemplatesRequest {
	r.limit = &limit
	return r
}

// Filter by device.type query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.device.type&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.device.type&#x3D;$not:$like:John Doe&amp;filter.device.type&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterDeviceType(filterDeviceType []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterDeviceType = &filterDeviceType
	return r
}

// Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterLabel(filterLabel []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterLabel = &filterLabel
	return r
}

// Filter by os.name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.os.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.os.name&#x3D;$not:$like:John Doe&amp;filter.os.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterOsName(filterOsName []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterOsName = &filterOsName
	return r
}

// Filter by install.onieStrings query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.install.onieStrings&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.install.onieStrings&#x3D;$not:$like:John Doe&amp;filter.install.onieStrings&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterInstallOnieStrings(filterInstallOnieStrings []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterInstallOnieStrings = &filterInstallOnieStrings
	return r
}

// Filter by status query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.status&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.status&#x3D;$not:$like:John Doe&amp;filter.status&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterStatus(filterStatus []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterStatus = &filterStatus
	return r
}

// Filter by visibility query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.visibility&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.visibility&#x3D;$not:$like:John Doe&amp;filter.visibility&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterVisibility(filterVisibility []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterVisibility = &filterVisibility
	return r
}

// Filter by createdBy query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.createdBy&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.createdBy&#x3D;$not:$like:John Doe&amp;filter.createdBy&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterCreatedBy(filterCreatedBy []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterCreatedBy = &filterCreatedBy
	return r
}

// Filter by modifiedBy query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.modifiedBy&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.modifiedBy&#x3D;$not:$like:John Doe&amp;filter.modifiedBy&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterModifiedBy(filterModifiedBy []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterModifiedBy = &filterModifiedBy
	return r
}

// Filter by tags query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.tags&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.tags&#x3D;$not:$like:John Doe&amp;filter.tags&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r OSTemplateAPIGetOSTemplatesRequest) FilterTags(filterTags []string) OSTemplateAPIGetOSTemplatesRequest {
	r.filterTags = &filterTags
	return r
}

// Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;device.type&lt;/li&gt; &lt;li&gt;visibility&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;createdBy&lt;/li&gt; &lt;li&gt;createdAt&lt;/li&gt; &lt;li&gt;modifiedAt&lt;/li&gt;&lt;/ul&gt;       
func (r OSTemplateAPIGetOSTemplatesRequest) SortBy(sortBy []string) OSTemplateAPIGetOSTemplatesRequest {
	r.sortBy = &sortBy
	return r
}

// Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;         
func (r OSTemplateAPIGetOSTemplatesRequest) Search(search string) OSTemplateAPIGetOSTemplatesRequest {
	r.search = &search
	return r
}

// List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,name,label,device.type,os.name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;label&lt;/li&gt; &lt;li&gt;device.type&lt;/li&gt; &lt;li&gt;os.name&lt;/li&gt; &lt;li&gt;os.version&lt;/li&gt; &lt;li&gt;install.onieStrings&lt;/li&gt; &lt;li&gt;visibility&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;li&gt;createdBy&lt;/li&gt; &lt;li&gt;modifiedBy&lt;/li&gt;&lt;/ul&gt;         
func (r OSTemplateAPIGetOSTemplatesRequest) SearchBy(searchBy []string) OSTemplateAPIGetOSTemplatesRequest {
	r.searchBy = &searchBy
	return r
}

// List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,name,description,label,device.type           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;       
func (r OSTemplateAPIGetOSTemplatesRequest) Select_(select_ string) OSTemplateAPIGetOSTemplatesRequest {
	r.select_ = &select_
	return r
}

func (r OSTemplateAPIGetOSTemplatesRequest) Execute() (*OSTemplatePaginatedList, *http.Response, error) {
	return r.ApiService.GetOSTemplatesExecute(r)
}

/*
GetOSTemplates Get a list of available OS templates

Returns list of the available OS templates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OSTemplateAPIGetOSTemplatesRequest
*/
func (a *OSTemplateAPIService) GetOSTemplates(ctx context.Context) OSTemplateAPIGetOSTemplatesRequest {
	return OSTemplateAPIGetOSTemplatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OSTemplatePaginatedList
func (a *OSTemplateAPIService) GetOSTemplatesExecute(r OSTemplateAPIGetOSTemplatesRequest) (*OSTemplatePaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSTemplatePaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OSTemplateAPIService.GetOSTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/os-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.filterDeviceType != nil {
		t := *r.filterDeviceType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.device.type", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.device.type", t, "form", "multi")
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
	if r.filterOsName != nil {
		t := *r.filterOsName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.os.name", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.os.name", t, "form", "multi")
		}
	}
	if r.filterInstallOnieStrings != nil {
		t := *r.filterInstallOnieStrings
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.install.onieStrings", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.install.onieStrings", t, "form", "multi")
		}
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
	if r.filterVisibility != nil {
		t := *r.filterVisibility
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.visibility", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.visibility", t, "form", "multi")
		}
	}
	if r.filterCreatedBy != nil {
		t := *r.filterCreatedBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdBy", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdBy", t, "form", "multi")
		}
	}
	if r.filterModifiedBy != nil {
		t := *r.filterModifiedBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.modifiedBy", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.modifiedBy", t, "form", "multi")
		}
	}
	if r.filterTags != nil {
		t := *r.filterTags
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.tags", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.tags", t, "form", "multi")
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
	if r.select_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "select", r.select_, "form", "")
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

type OSTemplateAPIUpdateOSTemplateRequest struct {
	ctx context.Context
	ApiService *OSTemplateAPIService
	osTemplateId float32
	oSTemplateUpdate *OSTemplateUpdate
	ifMatch *string
}

// The OS template details
func (r OSTemplateAPIUpdateOSTemplateRequest) OSTemplateUpdate(oSTemplateUpdate OSTemplateUpdate) OSTemplateAPIUpdateOSTemplateRequest {
	r.oSTemplateUpdate = &oSTemplateUpdate
	return r
}

// Entity tag
func (r OSTemplateAPIUpdateOSTemplateRequest) IfMatch(ifMatch string) OSTemplateAPIUpdateOSTemplateRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r OSTemplateAPIUpdateOSTemplateRequest) Execute() (*OSTemplate, *http.Response, error) {
	return r.ApiService.UpdateOSTemplateExecute(r)
}

/*
UpdateOSTemplate Update OS template

Returns details of the updated OS template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param osTemplateId
 @return OSTemplateAPIUpdateOSTemplateRequest
*/
func (a *OSTemplateAPIService) UpdateOSTemplate(ctx context.Context, osTemplateId float32) OSTemplateAPIUpdateOSTemplateRequest {
	return OSTemplateAPIUpdateOSTemplateRequest{
		ApiService: a,
		ctx: ctx,
		osTemplateId: osTemplateId,
	}
}

// Execute executes the request
//  @return OSTemplate
func (a *OSTemplateAPIService) UpdateOSTemplateExecute(r OSTemplateAPIUpdateOSTemplateRequest) (*OSTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OSTemplateAPIService.UpdateOSTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/os-templates/{osTemplateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"osTemplateId"+"}", url.PathEscape(parameterValueToString(r.osTemplateId, "osTemplateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oSTemplateUpdate == nil {
		return localVarReturnValue, nil, reportError("oSTemplateUpdate is required and must be specified")
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
	localVarPostBody = r.oSTemplateUpdate
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
