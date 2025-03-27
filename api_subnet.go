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
)


// SubnetAPIService SubnetAPI service
type SubnetAPIService service

type SubnetAPICreateSubnetRequest struct {
	ctx context.Context
	ApiService *SubnetAPIService
	createSubnetDto *CreateSubnetDto
}

// The Subnet to create
func (r SubnetAPICreateSubnetRequest) CreateSubnetDto(createSubnetDto CreateSubnetDto) SubnetAPICreateSubnetRequest {
	r.createSubnetDto = &createSubnetDto
	return r
}

func (r SubnetAPICreateSubnetRequest) Execute() (*CreateSubnetDto, *http.Response, error) {
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
//  @return CreateSubnetDto
func (a *SubnetAPIService) CreateSubnetExecute(r SubnetAPICreateSubnetRequest) (*CreateSubnetDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateSubnetDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubnetAPIService.CreateSubnet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/subnets/subnets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSubnetDto == nil {
		return localVarReturnValue, nil, reportError("createSubnetDto is required and must be specified")
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
	localVarPostBody = r.createSubnetDto
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
	updateSubnet *UpdateSubnet
	ifMatch *string
}

// The Subnet changes
func (r SubnetAPIUpdateSubnetRequest) UpdateSubnet(updateSubnet UpdateSubnet) SubnetAPIUpdateSubnetRequest {
	r.updateSubnet = &updateSubnet
	return r
}

// Entity tag
func (r SubnetAPIUpdateSubnetRequest) IfMatch(ifMatch string) SubnetAPIUpdateSubnetRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r SubnetAPIUpdateSubnetRequest) Execute() (*UpdateSubnet, *http.Response, error) {
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
//  @return UpdateSubnet
func (a *SubnetAPIService) UpdateSubnetExecute(r SubnetAPIUpdateSubnetRequest) (*UpdateSubnet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateSubnet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubnetAPIService.UpdateSubnet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/subnets/{subnetId}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"subnetId"+"}", url.PathEscape(parameterValueToString(r.subnetId, "subnetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
	}
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
