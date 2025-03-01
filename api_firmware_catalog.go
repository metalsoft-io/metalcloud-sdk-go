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


// FirmwareCatalogAPIService FirmwareCatalogAPI service
type FirmwareCatalogAPIService service

type FirmwareCatalogAPICreateFirmwareCatalogsRequest struct {
	ctx context.Context
	ApiService *FirmwareCatalogAPIService
	createFirmwareCatalog *CreateFirmwareCatalog
}

func (r FirmwareCatalogAPICreateFirmwareCatalogsRequest) CreateFirmwareCatalog(createFirmwareCatalog CreateFirmwareCatalog) FirmwareCatalogAPICreateFirmwareCatalogsRequest {
	r.createFirmwareCatalog = &createFirmwareCatalog
	return r
}

func (r FirmwareCatalogAPICreateFirmwareCatalogsRequest) Execute() (*FirmwareCatalog, *http.Response, error) {
	return r.ApiService.CreateFirmwareCatalogsExecute(r)
}

/*
CreateFirmwareCatalogs Create Firmware Catalog

Returns a list of firmware catalogs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FirmwareCatalogAPICreateFirmwareCatalogsRequest
*/
func (a *FirmwareCatalogAPIService) CreateFirmwareCatalogs(ctx context.Context) FirmwareCatalogAPICreateFirmwareCatalogsRequest {
	return FirmwareCatalogAPICreateFirmwareCatalogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FirmwareCatalog
func (a *FirmwareCatalogAPIService) CreateFirmwareCatalogsExecute(r FirmwareCatalogAPICreateFirmwareCatalogsRequest) (*FirmwareCatalog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareCatalog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareCatalogAPIService.CreateFirmwareCatalogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/catalog"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFirmwareCatalog == nil {
		return localVarReturnValue, nil, reportError("createFirmwareCatalog is required and must be specified")
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
	localVarPostBody = r.createFirmwareCatalog
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

type FirmwareCatalogAPIDeleteFirmwareCatalogRequest struct {
	ctx context.Context
	ApiService *FirmwareCatalogAPIService
	firmwareCatalogId float32
}

func (r FirmwareCatalogAPIDeleteFirmwareCatalogRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFirmwareCatalogExecute(r)
}

/*
DeleteFirmwareCatalog Delete Firmware Catalog

Deletes a firmware catalog

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firmwareCatalogId The firmware catalog id
 @return FirmwareCatalogAPIDeleteFirmwareCatalogRequest
*/
func (a *FirmwareCatalogAPIService) DeleteFirmwareCatalog(ctx context.Context, firmwareCatalogId float32) FirmwareCatalogAPIDeleteFirmwareCatalogRequest {
	return FirmwareCatalogAPIDeleteFirmwareCatalogRequest{
		ApiService: a,
		ctx: ctx,
		firmwareCatalogId: firmwareCatalogId,
	}
}

// Execute executes the request
func (a *FirmwareCatalogAPIService) DeleteFirmwareCatalogExecute(r FirmwareCatalogAPIDeleteFirmwareCatalogRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareCatalogAPIService.DeleteFirmwareCatalog")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/catalog/{firmwareCatalogId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firmwareCatalogId"+"}", url.PathEscape(parameterValueToString(r.firmwareCatalogId, "firmwareCatalogId")), -1)

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

type FirmwareCatalogAPIGetFirmwareCatalogRequest struct {
	ctx context.Context
	ApiService *FirmwareCatalogAPIService
	firmwareCatalogId float32
}

func (r FirmwareCatalogAPIGetFirmwareCatalogRequest) Execute() (*FirmwareCatalog, *http.Response, error) {
	return r.ApiService.GetFirmwareCatalogExecute(r)
}

/*
GetFirmwareCatalog Get Firmware Catalog

Returns a firmware catalog

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firmwareCatalogId The firmware catalog id
 @return FirmwareCatalogAPIGetFirmwareCatalogRequest
*/
func (a *FirmwareCatalogAPIService) GetFirmwareCatalog(ctx context.Context, firmwareCatalogId float32) FirmwareCatalogAPIGetFirmwareCatalogRequest {
	return FirmwareCatalogAPIGetFirmwareCatalogRequest{
		ApiService: a,
		ctx: ctx,
		firmwareCatalogId: firmwareCatalogId,
	}
}

// Execute executes the request
//  @return FirmwareCatalog
func (a *FirmwareCatalogAPIService) GetFirmwareCatalogExecute(r FirmwareCatalogAPIGetFirmwareCatalogRequest) (*FirmwareCatalog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareCatalog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareCatalogAPIService.GetFirmwareCatalog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/catalog/{firmwareCatalogId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firmwareCatalogId"+"}", url.PathEscape(parameterValueToString(r.firmwareCatalogId, "firmwareCatalogId")), -1)

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

type FirmwareCatalogAPIGetFirmwareCatalogsRequest struct {
	ctx context.Context
	ApiService *FirmwareCatalogAPIService
	page *float32
	limit *float32
	filterServerFirmwareCatalogId *[]string
	filterServerFirmwareCatalogName *[]string
	filterServerFirmwareCatalogDescription *[]string
	filterServerFirmwareCatalogVendor *[]string
	filterServerFirmwareCatalogVendorId *[]string
	filterServerFirmwareCatalogVendorUrl *[]string
	filterServerFirmwareCatalogVendorReleaseTimestamp *[]string
	filterServerFirmwareCatalogUpdateType *[]string
	filterServerFirmwareCatalogCreatedTimestamp *[]string
	sortBy *[]string
	search *string
	searchBy *[]string
	select_ *string
}

// Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;         
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) Page(page float32) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.page = &page
	return r
}

// Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.       
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) Limit(limit float32) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.limit = &limit
	return r
}

// Filter by serverFirmwareCatalogId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogId(filterServerFirmwareCatalogId []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogId = &filterServerFirmwareCatalogId
	return r
}

// Filter by serverFirmwareCatalogName query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogName&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogName&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogName&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogName(filterServerFirmwareCatalogName []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogName = &filterServerFirmwareCatalogName
	return r
}

// Filter by serverFirmwareCatalogDescription query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogDescription&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogDescription&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogDescription&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogDescription(filterServerFirmwareCatalogDescription []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogDescription = &filterServerFirmwareCatalogDescription
	return r
}

// Filter by serverFirmwareCatalogVendor query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendor&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendor&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendor&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogVendor(filterServerFirmwareCatalogVendor []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogVendor = &filterServerFirmwareCatalogVendor
	return r
}

// Filter by serverFirmwareCatalogVendorId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendorId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendorId&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendorId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogVendorId(filterServerFirmwareCatalogVendorId []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogVendorId = &filterServerFirmwareCatalogVendorId
	return r
}

// Filter by serverFirmwareCatalogVendorUrl query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendorUrl&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendorUrl&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendorUrl&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogVendorUrl(filterServerFirmwareCatalogVendorUrl []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogVendorUrl = &filterServerFirmwareCatalogVendorUrl
	return r
}

// Filter by serverFirmwareCatalogVendorReleaseTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogVendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogVendorReleaseTimestamp&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogVendorReleaseTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogVendorReleaseTimestamp(filterServerFirmwareCatalogVendorReleaseTimestamp []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogVendorReleaseTimestamp = &filterServerFirmwareCatalogVendorReleaseTimestamp
	return r
}

// Filter by serverFirmwareCatalogUpdateType query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogUpdateType&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogUpdateType&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogUpdateType&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogUpdateType(filterServerFirmwareCatalogUpdateType []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogUpdateType = &filterServerFirmwareCatalogUpdateType
	return r
}

// Filter by serverFirmwareCatalogCreatedTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.serverFirmwareCatalogCreatedTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.serverFirmwareCatalogCreatedTimestamp&#x3D;$not:$like:John Doe&amp;filter.serverFirmwareCatalogCreatedTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) FilterServerFirmwareCatalogCreatedTimestamp(filterServerFirmwareCatalogCreatedTimestamp []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.filterServerFirmwareCatalogCreatedTimestamp = &filterServerFirmwareCatalogCreatedTimestamp
	return r
}

// Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; serverFirmwareCatalogId:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;serverFirmwareCatalogId&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogName&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogVendor&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogVendorReleaseTimestamp&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogUpdateType&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogCreatedTimestamp&lt;/li&gt;&lt;/ul&gt;       
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) SortBy(sortBy []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.sortBy = &sortBy
	return r
}

// Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;         
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) Search(search string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.search = &search
	return r
}

// List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; serverFirmwareCatalogName,serverFirmwareCatalogDescription           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;serverFirmwareCatalogName&lt;/li&gt; &lt;li&gt;serverFirmwareCatalogDescription&lt;/li&gt;&lt;/ul&gt;         
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) SearchBy(searchBy []string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.searchBy = &searchBy
	return r
}

// List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; serverFirmwareCatalogId,serverFirmwareCatalogName,serverFirmwareCatalogDescription,serverFirmwareCatalogVendor,serverFirmwareCatalogVendorId           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;       
func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) Select_(select_ string) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	r.select_ = &select_
	return r
}

func (r FirmwareCatalogAPIGetFirmwareCatalogsRequest) Execute() (*FirmwareCatalogPaginatedList, *http.Response, error) {
	return r.ApiService.GetFirmwareCatalogsExecute(r)
}

/*
GetFirmwareCatalogs Get Firmware Catalogs

Returns a list of firmware catalogs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FirmwareCatalogAPIGetFirmwareCatalogsRequest
*/
func (a *FirmwareCatalogAPIService) GetFirmwareCatalogs(ctx context.Context) FirmwareCatalogAPIGetFirmwareCatalogsRequest {
	return FirmwareCatalogAPIGetFirmwareCatalogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FirmwareCatalogPaginatedList
func (a *FirmwareCatalogAPIService) GetFirmwareCatalogsExecute(r FirmwareCatalogAPIGetFirmwareCatalogsRequest) (*FirmwareCatalogPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareCatalogPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareCatalogAPIService.GetFirmwareCatalogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/catalog"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.filterServerFirmwareCatalogId != nil {
		t := *r.filterServerFirmwareCatalogId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogId", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogName != nil {
		t := *r.filterServerFirmwareCatalogName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogName", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogName", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogDescription != nil {
		t := *r.filterServerFirmwareCatalogDescription
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogDescription", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogDescription", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogVendor != nil {
		t := *r.filterServerFirmwareCatalogVendor
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendor", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendor", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogVendorId != nil {
		t := *r.filterServerFirmwareCatalogVendorId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendorId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendorId", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogVendorUrl != nil {
		t := *r.filterServerFirmwareCatalogVendorUrl
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendorUrl", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendorUrl", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogVendorReleaseTimestamp != nil {
		t := *r.filterServerFirmwareCatalogVendorReleaseTimestamp
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendorReleaseTimestamp", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogVendorReleaseTimestamp", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogUpdateType != nil {
		t := *r.filterServerFirmwareCatalogUpdateType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogUpdateType", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogUpdateType", t, "form", "multi")
		}
	}
	if r.filterServerFirmwareCatalogCreatedTimestamp != nil {
		t := *r.filterServerFirmwareCatalogCreatedTimestamp
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogCreatedTimestamp", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.serverFirmwareCatalogCreatedTimestamp", t, "form", "multi")
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

type FirmwareCatalogAPIUpdateFirmwareCatalogRequest struct {
	ctx context.Context
	ApiService *FirmwareCatalogAPIService
	firmwareCatalogId float32
	updateFirmwareCatalog *UpdateFirmwareCatalog
}

func (r FirmwareCatalogAPIUpdateFirmwareCatalogRequest) UpdateFirmwareCatalog(updateFirmwareCatalog UpdateFirmwareCatalog) FirmwareCatalogAPIUpdateFirmwareCatalogRequest {
	r.updateFirmwareCatalog = &updateFirmwareCatalog
	return r
}

func (r FirmwareCatalogAPIUpdateFirmwareCatalogRequest) Execute() (*FirmwareCatalog, *http.Response, error) {
	return r.ApiService.UpdateFirmwareCatalogExecute(r)
}

/*
UpdateFirmwareCatalog Update Firmware Catalog

Updates a firmware catalog

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firmwareCatalogId The firmware catalog id
 @return FirmwareCatalogAPIUpdateFirmwareCatalogRequest
*/
func (a *FirmwareCatalogAPIService) UpdateFirmwareCatalog(ctx context.Context, firmwareCatalogId float32) FirmwareCatalogAPIUpdateFirmwareCatalogRequest {
	return FirmwareCatalogAPIUpdateFirmwareCatalogRequest{
		ApiService: a,
		ctx: ctx,
		firmwareCatalogId: firmwareCatalogId,
	}
}

// Execute executes the request
//  @return FirmwareCatalog
func (a *FirmwareCatalogAPIService) UpdateFirmwareCatalogExecute(r FirmwareCatalogAPIUpdateFirmwareCatalogRequest) (*FirmwareCatalog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareCatalog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareCatalogAPIService.UpdateFirmwareCatalog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/catalog/{firmwareCatalogId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firmwareCatalogId"+"}", url.PathEscape(parameterValueToString(r.firmwareCatalogId, "firmwareCatalogId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateFirmwareCatalog == nil {
		return localVarReturnValue, nil, reportError("updateFirmwareCatalog is required and must be specified")
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
	localVarPostBody = r.updateFirmwareCatalog
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
