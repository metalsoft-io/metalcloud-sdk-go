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


// ServerDefaultCredentialsAPIService ServerDefaultCredentialsAPI service
type ServerDefaultCredentialsAPIService service

type ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest struct {
	ctx context.Context
	ApiService *ServerDefaultCredentialsAPIService
	createServerDefaultCredentials *CreateServerDefaultCredentials
}

// The Server Default Credentials create object
func (r ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest) CreateServerDefaultCredentials(createServerDefaultCredentials CreateServerDefaultCredentials) ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest {
	r.createServerDefaultCredentials = &createServerDefaultCredentials
	return r
}

func (r ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest) Execute() (*ServerDefaultCredentials, *http.Response, error) {
	return r.ApiService.CreateServerDefaultCredentialsExecute(r)
}

/*
CreateServerDefaultCredentials Creates a Server Default Credentials

Creates a Server Default Credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest
*/
func (a *ServerDefaultCredentialsAPIService) CreateServerDefaultCredentials(ctx context.Context) ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest {
	return ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServerDefaultCredentials
func (a *ServerDefaultCredentialsAPIService) CreateServerDefaultCredentialsExecute(r ServerDefaultCredentialsAPICreateServerDefaultCredentialsRequest) (*ServerDefaultCredentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServerDefaultCredentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerDefaultCredentialsAPIService.CreateServerDefaultCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/servers/default-credentials"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createServerDefaultCredentials == nil {
		return localVarReturnValue, nil, reportError("createServerDefaultCredentials is required and must be specified")
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
	localVarPostBody = r.createServerDefaultCredentials
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

type ServerDefaultCredentialsAPIDeleteServerDefaultCredentialsRequest struct {
	ctx context.Context
	ApiService *ServerDefaultCredentialsAPIService
	serverDefaultCredentialsId float32
}

func (r ServerDefaultCredentialsAPIDeleteServerDefaultCredentialsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServerDefaultCredentialsExecute(r)
}

/*
DeleteServerDefaultCredentials Deletes a Server Default Credentials

Deletes a Server Default Credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serverDefaultCredentialsId
 @return ServerDefaultCredentialsAPIDeleteServerDefaultCredentialsRequest
*/
func (a *ServerDefaultCredentialsAPIService) DeleteServerDefaultCredentials(ctx context.Context, serverDefaultCredentialsId float32) ServerDefaultCredentialsAPIDeleteServerDefaultCredentialsRequest {
	return ServerDefaultCredentialsAPIDeleteServerDefaultCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		serverDefaultCredentialsId: serverDefaultCredentialsId,
	}
}

// Execute executes the request
func (a *ServerDefaultCredentialsAPIService) DeleteServerDefaultCredentialsExecute(r ServerDefaultCredentialsAPIDeleteServerDefaultCredentialsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerDefaultCredentialsAPIService.DeleteServerDefaultCredentials")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/servers/default-credentials/{serverDefaultCredentialsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serverDefaultCredentialsId"+"}", url.PathEscape(parameterValueToString(r.serverDefaultCredentialsId, "serverDefaultCredentialsId")), -1)

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

type ServerDefaultCredentialsAPIGetServerDefaultCredentialsCredentialsRequest struct {
	ctx context.Context
	ApiService *ServerDefaultCredentialsAPIService
	serverDefaultCredentialsId float32
}

func (r ServerDefaultCredentialsAPIGetServerDefaultCredentialsCredentialsRequest) Execute() (*ServerDefaultCredentialsCredentials, *http.Response, error) {
	return r.ApiService.GetServerDefaultCredentialsCredentialsExecute(r)
}

/*
GetServerDefaultCredentialsCredentials Get Server Default Credentials unencrypted

Returns Server Default Credentials unencrypted

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serverDefaultCredentialsId
 @return ServerDefaultCredentialsAPIGetServerDefaultCredentialsCredentialsRequest
*/
func (a *ServerDefaultCredentialsAPIService) GetServerDefaultCredentialsCredentials(ctx context.Context, serverDefaultCredentialsId float32) ServerDefaultCredentialsAPIGetServerDefaultCredentialsCredentialsRequest {
	return ServerDefaultCredentialsAPIGetServerDefaultCredentialsCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		serverDefaultCredentialsId: serverDefaultCredentialsId,
	}
}

// Execute executes the request
//  @return ServerDefaultCredentialsCredentials
func (a *ServerDefaultCredentialsAPIService) GetServerDefaultCredentialsCredentialsExecute(r ServerDefaultCredentialsAPIGetServerDefaultCredentialsCredentialsRequest) (*ServerDefaultCredentialsCredentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServerDefaultCredentialsCredentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerDefaultCredentialsAPIService.GetServerDefaultCredentialsCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/servers/default-credentials/{serverDefaultCredentialsId}/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"serverDefaultCredentialsId"+"}", url.PathEscape(parameterValueToString(r.serverDefaultCredentialsId, "serverDefaultCredentialsId")), -1)

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

type ServerDefaultCredentialsAPIGetServerDefaultCredentialsInfoRequest struct {
	ctx context.Context
	ApiService *ServerDefaultCredentialsAPIService
	serverDefaultCredentialsId float32
}

func (r ServerDefaultCredentialsAPIGetServerDefaultCredentialsInfoRequest) Execute() (*ServerDefaultCredentials, *http.Response, error) {
	return r.ApiService.GetServerDefaultCredentialsInfoExecute(r)
}

/*
GetServerDefaultCredentialsInfo Get Server Default Credentials information

Returns Server Default Credentials information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serverDefaultCredentialsId
 @return ServerDefaultCredentialsAPIGetServerDefaultCredentialsInfoRequest
*/
func (a *ServerDefaultCredentialsAPIService) GetServerDefaultCredentialsInfo(ctx context.Context, serverDefaultCredentialsId float32) ServerDefaultCredentialsAPIGetServerDefaultCredentialsInfoRequest {
	return ServerDefaultCredentialsAPIGetServerDefaultCredentialsInfoRequest{
		ApiService: a,
		ctx: ctx,
		serverDefaultCredentialsId: serverDefaultCredentialsId,
	}
}

// Execute executes the request
//  @return ServerDefaultCredentials
func (a *ServerDefaultCredentialsAPIService) GetServerDefaultCredentialsInfoExecute(r ServerDefaultCredentialsAPIGetServerDefaultCredentialsInfoRequest) (*ServerDefaultCredentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServerDefaultCredentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerDefaultCredentialsAPIService.GetServerDefaultCredentialsInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/servers/default-credentials/{serverDefaultCredentialsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serverDefaultCredentialsId"+"}", url.PathEscape(parameterValueToString(r.serverDefaultCredentialsId, "serverDefaultCredentialsId")), -1)

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

type ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest struct {
	ctx context.Context
	ApiService *ServerDefaultCredentialsAPIService
	page *float32
	limit *float32
	filterSiteId *[]string
	filterServerSerialNumber *[]string
	filterServerMacAddress *[]string
	sortBy *[]string
	search *string
	searchBy *[]string
}

// Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;         
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) Page(page float32) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.page = &page
	return r
}

// Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.       
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) Limit(limit float32) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.limit = &limit
	return r
}

// Filter by siteId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.siteId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.siteId&#x3D;$not:$like:John Doe&amp;filter.siteId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) FilterSiteId(filterSiteId []string) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.filterSiteId = &filterSiteId
	return r
}

// Filter by serverSerialNumber query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverSerialNumber&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverSerialNumber&#x3D;$not:$like:John Doe&amp;filter.serverSerialNumber&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) FilterServerSerialNumber(filterServerSerialNumber []string) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.filterServerSerialNumber = &filterServerSerialNumber
	return r
}

// Filter by serverMacAddress query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverMacAddress&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverMacAddress&#x3D;$not:$like:John Doe&amp;filter.serverMacAddress&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) FilterServerMacAddress(filterServerMacAddress []string) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.filterServerMacAddress = &filterServerMacAddress
	return r
}

// Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt;&lt;/ul&gt;       
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) SortBy(sortBy []string) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.sortBy = &sortBy
	return r
}

// Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;         
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) Search(search string) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.search = &search
	return r
}

// List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,siteId,serverSerialNumber,serverMacAddress           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;siteId&lt;/li&gt; &lt;li&gt;serverSerialNumber&lt;/li&gt; &lt;li&gt;serverMacAddress&lt;/li&gt;&lt;/ul&gt;         
func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) SearchBy(searchBy []string) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	r.searchBy = &searchBy
	return r
}

func (r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) Execute() (*ServerDefaultCredentialsPaginatedList, *http.Response, error) {
	return r.ApiService.GetServersDefaultCredentialsExecute(r)
}

/*
GetServersDefaultCredentials Get a list of Server Default Credentials

Returns a list of Server Default Credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest
*/
func (a *ServerDefaultCredentialsAPIService) GetServersDefaultCredentials(ctx context.Context) ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest {
	return ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServerDefaultCredentialsPaginatedList
func (a *ServerDefaultCredentialsAPIService) GetServersDefaultCredentialsExecute(r ServerDefaultCredentialsAPIGetServersDefaultCredentialsRequest) (*ServerDefaultCredentialsPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServerDefaultCredentialsPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerDefaultCredentialsAPIService.GetServersDefaultCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/servers/default-credentials"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.filterSiteId != nil {
		t := *r.filterSiteId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.siteId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.siteId", t, "form", "multi")
		}
	}
	if r.filterServerSerialNumber != nil {
		t := *r.filterServerSerialNumber
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverSerialNumber", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverSerialNumber", t, "form", "multi")
		}
	}
	if r.filterServerMacAddress != nil {
		t := *r.filterServerMacAddress
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverMacAddress", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverMacAddress", t, "form", "multi")
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

type ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest struct {
	ctx context.Context
	ApiService *ServerDefaultCredentialsAPIService
	serverDefaultCredentialsId float32
	updateServerDefaultCredentials *UpdateServerDefaultCredentials
}

// The Server Default Credentials update object
func (r ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest) UpdateServerDefaultCredentials(updateServerDefaultCredentials UpdateServerDefaultCredentials) ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest {
	r.updateServerDefaultCredentials = &updateServerDefaultCredentials
	return r
}

func (r ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest) Execute() (*ServerDefaultCredentials, *http.Response, error) {
	return r.ApiService.UpdateServerDefaultCredentialsExecute(r)
}

/*
UpdateServerDefaultCredentials Updates a Server Default Credentials

Updates a Server Default Credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serverDefaultCredentialsId
 @return ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest
*/
func (a *ServerDefaultCredentialsAPIService) UpdateServerDefaultCredentials(ctx context.Context, serverDefaultCredentialsId float32) ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest {
	return ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		serverDefaultCredentialsId: serverDefaultCredentialsId,
	}
}

// Execute executes the request
//  @return ServerDefaultCredentials
func (a *ServerDefaultCredentialsAPIService) UpdateServerDefaultCredentialsExecute(r ServerDefaultCredentialsAPIUpdateServerDefaultCredentialsRequest) (*ServerDefaultCredentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServerDefaultCredentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerDefaultCredentialsAPIService.UpdateServerDefaultCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/servers/default-credentials/{serverDefaultCredentialsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serverDefaultCredentialsId"+"}", url.PathEscape(parameterValueToString(r.serverDefaultCredentialsId, "serverDefaultCredentialsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateServerDefaultCredentials == nil {
		return localVarReturnValue, nil, reportError("updateServerDefaultCredentials is required and must be specified")
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
	localVarPostBody = r.updateServerDefaultCredentials
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
