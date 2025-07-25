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


// SubnetAPIService SubnetAPI service
type SubnetAPIService service

type SubnetAPICreateSubnetRequest struct {
	ctx context.Context
	ApiService *SubnetAPIService
	createSubnet *CreateSubnet
}

// The Subnet to create
func (r SubnetAPICreateSubnetRequest) CreateSubnet(createSubnet CreateSubnet) SubnetAPICreateSubnetRequest {
	r.createSubnet = &createSubnet
	return r
}

func (r SubnetAPICreateSubnetRequest) Execute() (*Subnet, *http.Response, error) {
	return r.ApiService.CreateSubnetExecute(r)
}

/*
CreateSubnet Create a subnet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubnetAPICreateSubnetRequest
*/
func (a *SubnetAPIService) CreateSubnet(ctx context.Context) SubnetAPICreateSubnetRequest {
	return SubnetAPICreateSubnetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Subnet
func (a *SubnetAPIService) CreateSubnetExecute(r SubnetAPICreateSubnetRequest) (*Subnet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subnet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubnetAPIService.CreateSubnet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/subnets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSubnet == nil {
		return localVarReturnValue, nil, reportError("createSubnet is required and must be specified")
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
	localVarPostBody = r.createSubnet
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

type SubnetAPIDeleteSubnetRequest struct {
	ctx context.Context
	ApiService *SubnetAPIService
	subnetId int32
	ifMatch *string
}

// Entity tag
func (r SubnetAPIDeleteSubnetRequest) IfMatch(ifMatch string) SubnetAPIDeleteSubnetRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r SubnetAPIDeleteSubnetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSubnetExecute(r)
}

/*
DeleteSubnet Delete Subnet

Delete Subnet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subnetId
 @return SubnetAPIDeleteSubnetRequest
*/
func (a *SubnetAPIService) DeleteSubnet(ctx context.Context, subnetId int32) SubnetAPIDeleteSubnetRequest {
	return SubnetAPIDeleteSubnetRequest{
		ApiService: a,
		ctx: ctx,
		subnetId: subnetId,
	}
}

// Execute executes the request
func (a *SubnetAPIService) DeleteSubnetExecute(r SubnetAPIDeleteSubnetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubnetAPIService.DeleteSubnet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/subnets/{subnetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subnetId"+"}", url.PathEscape(parameterValueToString(r.subnetId, "subnetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ifMatch == nil {
		return nil, reportError("ifMatch is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
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

type SubnetAPIGetSubnetRequest struct {
	ctx context.Context
	ApiService *SubnetAPIService
	subnetId float32
}

func (r SubnetAPIGetSubnetRequest) Execute() (*Subnet, *http.Response, error) {
	return r.ApiService.GetSubnetExecute(r)
}

/*
GetSubnet Retrieves the Subnet information

Retrieves the Subnet information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subnetId
 @return SubnetAPIGetSubnetRequest
*/
func (a *SubnetAPIService) GetSubnet(ctx context.Context, subnetId float32) SubnetAPIGetSubnetRequest {
	return SubnetAPIGetSubnetRequest{
		ApiService: a,
		ctx: ctx,
		subnetId: subnetId,
	}
}

// Execute executes the request
//  @return Subnet
func (a *SubnetAPIService) GetSubnetExecute(r SubnetAPIGetSubnetRequest) (*Subnet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subnet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubnetAPIService.GetSubnet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/subnets/{subnetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subnetId"+"}", url.PathEscape(parameterValueToString(r.subnetId, "subnetId")), -1)

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

type SubnetAPIGetSubnetsRequest struct {
	ctx context.Context
	ApiService *SubnetAPIService
	page *float32
	limit *float32
	filterId *[]string
	filterLabel *[]string
	filterName *[]string
	filterIpVersion *[]string
	filterIsPool *[]string
	filterParentSubnetId *[]string
	sortBy *[]string
	search *string
	searchBy *[]string
}

// Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;         
func (r SubnetAPIGetSubnetsRequest) Page(page float32) SubnetAPIGetSubnetsRequest {
	r.page = &page
	return r
}

// Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1000           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 5000           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.       
func (r SubnetAPIGetSubnetsRequest) Limit(limit float32) SubnetAPIGetSubnetsRequest {
	r.limit = &limit
	return r
}

// Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt;&lt;/ul&gt;
func (r SubnetAPIGetSubnetsRequest) FilterId(filterId []string) SubnetAPIGetSubnetsRequest {
	r.filterId = &filterId
	return r
}

// Filter by label query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.label&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.label&#x3D;$not:$like:John Doe&amp;filter.label&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r SubnetAPIGetSubnetsRequest) FilterLabel(filterLabel []string) SubnetAPIGetSubnetsRequest {
	r.filterLabel = &filterLabel
	return r
}

// Filter by name query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.name&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.name&#x3D;$not:$like:John Doe&amp;filter.name&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r SubnetAPIGetSubnetsRequest) FilterName(filterName []string) SubnetAPIGetSubnetsRequest {
	r.filterName = &filterName
	return r
}

// Filter by ipVersion query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.ipVersion&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.ipVersion&#x3D;$not:$like:John Doe&amp;filter.ipVersion&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r SubnetAPIGetSubnetsRequest) FilterIpVersion(filterIpVersion []string) SubnetAPIGetSubnetsRequest {
	r.filterIpVersion = &filterIpVersion
	return r
}

// Filter by isPool query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.isPool&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.isPool&#x3D;$not:$like:John Doe&amp;filter.isPool&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt;&lt;/ul&gt;
func (r SubnetAPIGetSubnetsRequest) FilterIsPool(filterIsPool []string) SubnetAPIGetSubnetsRequest {
	r.filterIsPool = &filterIsPool
	return r
}

// Filter by parentSubnetId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.parentSubnetId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.parentSubnetId&#x3D;$not:$like:John Doe&amp;filter.parentSubnetId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt;&lt;/ul&gt;
func (r SubnetAPIGetSubnetsRequest) FilterParentSubnetId(filterParentSubnetId []string) SubnetAPIGetSubnetsRequest {
	r.filterParentSubnetId = &filterParentSubnetId
	return r
}

// Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:ASC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt;&lt;/ul&gt;       
func (r SubnetAPIGetSubnetsRequest) SortBy(sortBy []string) SubnetAPIGetSubnetsRequest {
	r.sortBy = &sortBy
	return r
}

// Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;         
func (r SubnetAPIGetSubnetsRequest) Search(search string) SubnetAPIGetSubnetsRequest {
	r.search = &search
	return r
}

// List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; label,name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;label&lt;/li&gt; &lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;         
func (r SubnetAPIGetSubnetsRequest) SearchBy(searchBy []string) SubnetAPIGetSubnetsRequest {
	r.searchBy = &searchBy
	return r
}

func (r SubnetAPIGetSubnetsRequest) Execute() (*SubnetPaginatedList, *http.Response, error) {
	return r.ApiService.GetSubnetsExecute(r)
}

/*
GetSubnets List all Subnets

Returns list of all Subnets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubnetAPIGetSubnetsRequest
*/
func (a *SubnetAPIService) GetSubnets(ctx context.Context) SubnetAPIGetSubnetsRequest {
	return SubnetAPIGetSubnetsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubnetPaginatedList
func (a *SubnetAPIService) GetSubnetsExecute(r SubnetAPIGetSubnetsRequest) (*SubnetPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubnetPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubnetAPIService.GetSubnets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/subnets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.filterId != nil {
		t := *r.filterId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.id", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.id", t, "form", "multi")
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
	if r.filterIpVersion != nil {
		t := *r.filterIpVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.ipVersion", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.ipVersion", t, "form", "multi")
		}
	}
	if r.filterIsPool != nil {
		t := *r.filterIsPool
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.isPool", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.isPool", t, "form", "multi")
		}
	}
	if r.filterParentSubnetId != nil {
		t := *r.filterParentSubnetId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.parentSubnetId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.parentSubnetId", t, "form", "multi")
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

type SubnetAPIUpdateSubnetRequest struct {
	ctx context.Context
	ApiService *SubnetAPIService
	subnetId int32
	ifMatch *string
	updateSubnet *UpdateSubnet
}

// Entity tag
func (r SubnetAPIUpdateSubnetRequest) IfMatch(ifMatch string) SubnetAPIUpdateSubnetRequest {
	r.ifMatch = &ifMatch
	return r
}

// The Subnet changes
func (r SubnetAPIUpdateSubnetRequest) UpdateSubnet(updateSubnet UpdateSubnet) SubnetAPIUpdateSubnetRequest {
	r.updateSubnet = &updateSubnet
	return r
}

func (r SubnetAPIUpdateSubnetRequest) Execute() (*Subnet, *http.Response, error) {
	return r.ApiService.UpdateSubnetExecute(r)
}

/*
UpdateSubnet Updates Subnet

Updates the specified Subnet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subnetId
 @return SubnetAPIUpdateSubnetRequest
*/
func (a *SubnetAPIService) UpdateSubnet(ctx context.Context, subnetId int32) SubnetAPIUpdateSubnetRequest {
	return SubnetAPIUpdateSubnetRequest{
		ApiService: a,
		ctx: ctx,
		subnetId: subnetId,
	}
}

// Execute executes the request
//  @return Subnet
func (a *SubnetAPIService) UpdateSubnetExecute(r SubnetAPIUpdateSubnetRequest) (*Subnet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subnet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubnetAPIService.UpdateSubnet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/subnets/{subnetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subnetId"+"}", url.PathEscape(parameterValueToString(r.subnetId, "subnetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ifMatch == nil {
		return localVarReturnValue, nil, reportError("ifMatch is required and must be specified")
	}
	if r.updateSubnet == nil {
		return localVarReturnValue, nil, reportError("updateSubnet is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
	// body params
	localVarPostBody = r.updateSubnet
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
