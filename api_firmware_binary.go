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


// FirmwareBinaryAPIService FirmwareBinaryAPI service
type FirmwareBinaryAPIService service

type FirmwareBinaryAPICreateFirmwareBinaryRequest struct {
	ctx context.Context
	ApiService *FirmwareBinaryAPIService
	createFirmwareBinary *CreateFirmwareBinary
}

func (r FirmwareBinaryAPICreateFirmwareBinaryRequest) CreateFirmwareBinary(createFirmwareBinary CreateFirmwareBinary) FirmwareBinaryAPICreateFirmwareBinaryRequest {
	r.createFirmwareBinary = &createFirmwareBinary
	return r
}

func (r FirmwareBinaryAPICreateFirmwareBinaryRequest) Execute() (*FirmwareBinary, *http.Response, error) {
	return r.ApiService.CreateFirmwareBinaryExecute(r)
}

/*
CreateFirmwareBinary Create a new firmware binary

Creates a new firmware binary and returns it

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FirmwareBinaryAPICreateFirmwareBinaryRequest
*/
func (a *FirmwareBinaryAPIService) CreateFirmwareBinary(ctx context.Context) FirmwareBinaryAPICreateFirmwareBinaryRequest {
	return FirmwareBinaryAPICreateFirmwareBinaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FirmwareBinary
func (a *FirmwareBinaryAPIService) CreateFirmwareBinaryExecute(r FirmwareBinaryAPICreateFirmwareBinaryRequest) (*FirmwareBinary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareBinary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareBinaryAPIService.CreateFirmwareBinary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/binary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFirmwareBinary == nil {
		return localVarReturnValue, nil, reportError("createFirmwareBinary is required and must be specified")
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
	localVarPostBody = r.createFirmwareBinary
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

type FirmwareBinaryAPIDeleteFirmwareBinaryRequest struct {
	ctx context.Context
	ApiService *FirmwareBinaryAPIService
	firmwareBinaryId float32
}

func (r FirmwareBinaryAPIDeleteFirmwareBinaryRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFirmwareBinaryExecute(r)
}

/*
DeleteFirmwareBinary Delete Firmware Binary

Deletes a firmware binary

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firmwareBinaryId The firmware binary id
 @return FirmwareBinaryAPIDeleteFirmwareBinaryRequest
*/
func (a *FirmwareBinaryAPIService) DeleteFirmwareBinary(ctx context.Context, firmwareBinaryId float32) FirmwareBinaryAPIDeleteFirmwareBinaryRequest {
	return FirmwareBinaryAPIDeleteFirmwareBinaryRequest{
		ApiService: a,
		ctx: ctx,
		firmwareBinaryId: firmwareBinaryId,
	}
}

// Execute executes the request
func (a *FirmwareBinaryAPIService) DeleteFirmwareBinaryExecute(r FirmwareBinaryAPIDeleteFirmwareBinaryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareBinaryAPIService.DeleteFirmwareBinary")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/binary/{firmwareBinaryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firmwareBinaryId"+"}", url.PathEscape(parameterValueToString(r.firmwareBinaryId, "firmwareBinaryId")), -1)

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

type FirmwareBinaryAPIGetFirmwareBinariesRequest struct {
	ctx context.Context
	ApiService *FirmwareBinaryAPIService
	page *float32
	limit *float32
	filterId *[]string
	filterCatalogId *[]string
	filterExternalId *[]string
	filterPackageId *[]string
	filterCreatedTimestamp *[]string
	filterUpdateSeverity *[]string
	filterRebootRequired *[]string
	filterVendorReleaseTimestamp *[]string
	filterPackageVersion *[]string
	sortBy *[]string
	search *string
	searchBy *[]string
	select_ *string
}

// Page number to retrieve.If you provide invalid value the default page number will applied         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 1           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 1           &lt;/p&gt;         
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) Page(page float32) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.page = &page
	return r
}

// Number of records per page.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; 20           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Max Value: &lt;/b&gt; 100           &lt;/p&gt;        If provided value is greater than max value, max value will be applied.       
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) Limit(limit float32) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.limit = &limit
	return r
}

// Filter by id query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.id&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.id&#x3D;$not:$like:John Doe&amp;filter.id&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterId(filterId []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterId = &filterId
	return r
}

// Filter by catalogId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.catalogId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.catalogId&#x3D;$not:$like:John Doe&amp;filter.catalogId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterCatalogId(filterCatalogId []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterCatalogId = &filterCatalogId
	return r
}

// Filter by externalId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.externalId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.externalId&#x3D;$not:$like:John Doe&amp;filter.externalId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterExternalId(filterExternalId []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterExternalId = &filterExternalId
	return r
}

// Filter by packageId query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.packageId&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.packageId&#x3D;$not:$like:John Doe&amp;filter.packageId&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterPackageId(filterPackageId []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterPackageId = &filterPackageId
	return r
}

// Filter by createdTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.createdTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.createdTimestamp&#x3D;$not:$like:John Doe&amp;filter.createdTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterCreatedTimestamp(filterCreatedTimestamp []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterCreatedTimestamp = &filterCreatedTimestamp
	return r
}

// Filter by updateSeverity query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.updateSeverity&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.updateSeverity&#x3D;$not:$like:John Doe&amp;filter.updateSeverity&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterUpdateSeverity(filterUpdateSeverity []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterUpdateSeverity = &filterUpdateSeverity
	return r
}

// Filter by rebootRequired query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.rebootRequired&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.rebootRequired&#x3D;$not:$like:John Doe&amp;filter.rebootRequired&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterRebootRequired(filterRebootRequired []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterRebootRequired = &filterRebootRequired
	return r
}

// Filter by vendorReleaseTimestamp query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.vendorReleaseTimestamp&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.vendorReleaseTimestamp&#x3D;$not:$like:John Doe&amp;filter.vendorReleaseTimestamp&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterVendorReleaseTimestamp(filterVendorReleaseTimestamp []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterVendorReleaseTimestamp = &filterVendorReleaseTimestamp
	return r
}

// Filter by packageVersion query param.           &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; filter.packageVersion&#x3D;{$not}:OPERATION:VALUE           &lt;/p&gt;           &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; filter.packageVersion&#x3D;$not:$like:John Doe&amp;filter.packageVersion&#x3D;like:John           &lt;/p&gt;           &lt;h4&gt;Available Operations&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;$and&lt;/li&gt; &lt;li&gt;$or&lt;/li&gt; &lt;li&gt;$not&lt;/li&gt; &lt;li&gt;$eq&lt;/li&gt; &lt;li&gt;$gt&lt;/li&gt; &lt;li&gt;$gte&lt;/li&gt; &lt;li&gt;$in&lt;/li&gt; &lt;li&gt;$null&lt;/li&gt; &lt;li&gt;$lt&lt;/li&gt; &lt;li&gt;$lte&lt;/li&gt; &lt;li&gt;$btw&lt;/li&gt; &lt;li&gt;$ilike&lt;/li&gt; &lt;li&gt;$sw&lt;/li&gt; &lt;li&gt;$contains&lt;/li&gt;&lt;/ul&gt;
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) FilterPackageVersion(filterPackageVersion []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.filterPackageVersion = &filterPackageVersion
	return r
}

// Parameter to sort by.       &lt;p&gt;To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting&lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Format: &lt;/b&gt; fieldName:DIRECTION           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; sortBy&#x3D;id:DESC&amp;sortBy&#x3D;createdAt:ASC           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; id:DESC           &lt;/p&gt;       &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;id&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;updateSeverity&lt;/li&gt; &lt;li&gt;catalogId&lt;/li&gt; &lt;li&gt;rebootRequired&lt;/li&gt; &lt;li&gt;vendorReleaseTimestamp&lt;/li&gt; &lt;li&gt;packageVersion&lt;/li&gt;&lt;/ul&gt;       
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) SortBy(sortBy []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.sortBy = &sortBy
	return r
}

// Search term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; John           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; No default value           &lt;/p&gt;         
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) Search(search string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.search = &search
	return r
}

// List of fields to search by term to filter result values         &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; name           &lt;/p&gt;         &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields mentioned below will be used to search by term           &lt;/p&gt;         &lt;h4&gt;Available Fields&lt;/h4&gt;&lt;ul&gt;&lt;li&gt;name&lt;/li&gt;&lt;/ul&gt;         
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) SearchBy(searchBy []string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.searchBy = &searchBy
	return r
}

// List of fields to select.       &lt;p&gt;              &lt;b&gt;Example: &lt;/b&gt; id,catalogId,externalId,packageId,packageVersion           &lt;/p&gt;       &lt;p&gt;              &lt;b&gt;Default Value: &lt;/b&gt; By default all fields returns. If you want to select only some fields, provide them in query param           &lt;/p&gt;       
func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) Select_(select_ string) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	r.select_ = &select_
	return r
}

func (r FirmwareBinaryAPIGetFirmwareBinariesRequest) Execute() (*FirmwareBinaryPaginatedList, *http.Response, error) {
	return r.ApiService.GetFirmwareBinariesExecute(r)
}

/*
GetFirmwareBinaries Get Firmware Binaries

Returns a list of firmware binaries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FirmwareBinaryAPIGetFirmwareBinariesRequest
*/
func (a *FirmwareBinaryAPIService) GetFirmwareBinaries(ctx context.Context) FirmwareBinaryAPIGetFirmwareBinariesRequest {
	return FirmwareBinaryAPIGetFirmwareBinariesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FirmwareBinaryPaginatedList
func (a *FirmwareBinaryAPIService) GetFirmwareBinariesExecute(r FirmwareBinaryAPIGetFirmwareBinariesRequest) (*FirmwareBinaryPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareBinaryPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareBinaryAPIService.GetFirmwareBinaries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/binary"

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
	if r.filterCatalogId != nil {
		t := *r.filterCatalogId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.catalogId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.catalogId", t, "form", "multi")
		}
	}
	if r.filterExternalId != nil {
		t := *r.filterExternalId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.externalId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.externalId", t, "form", "multi")
		}
	}
	if r.filterPackageId != nil {
		t := *r.filterPackageId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.packageId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.packageId", t, "form", "multi")
		}
	}
	if r.filterCreatedTimestamp != nil {
		t := *r.filterCreatedTimestamp
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdTimestamp", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdTimestamp", t, "form", "multi")
		}
	}
	if r.filterUpdateSeverity != nil {
		t := *r.filterUpdateSeverity
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.updateSeverity", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.updateSeverity", t, "form", "multi")
		}
	}
	if r.filterRebootRequired != nil {
		t := *r.filterRebootRequired
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.rebootRequired", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.rebootRequired", t, "form", "multi")
		}
	}
	if r.filterVendorReleaseTimestamp != nil {
		t := *r.filterVendorReleaseTimestamp
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.vendorReleaseTimestamp", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.vendorReleaseTimestamp", t, "form", "multi")
		}
	}
	if r.filterPackageVersion != nil {
		t := *r.filterPackageVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter.packageVersion", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter.packageVersion", t, "form", "multi")
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

type FirmwareBinaryAPIGetFirmwareBinaryRequest struct {
	ctx context.Context
	ApiService *FirmwareBinaryAPIService
	firmwareBinaryId float32
}

func (r FirmwareBinaryAPIGetFirmwareBinaryRequest) Execute() (*FirmwareBinary, *http.Response, error) {
	return r.ApiService.GetFirmwareBinaryExecute(r)
}

/*
GetFirmwareBinary Get Firmware Binary

Returns a firmware binary

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firmwareBinaryId The firmware binary id
 @return FirmwareBinaryAPIGetFirmwareBinaryRequest
*/
func (a *FirmwareBinaryAPIService) GetFirmwareBinary(ctx context.Context, firmwareBinaryId float32) FirmwareBinaryAPIGetFirmwareBinaryRequest {
	return FirmwareBinaryAPIGetFirmwareBinaryRequest{
		ApiService: a,
		ctx: ctx,
		firmwareBinaryId: firmwareBinaryId,
	}
}

// Execute executes the request
//  @return FirmwareBinary
func (a *FirmwareBinaryAPIService) GetFirmwareBinaryExecute(r FirmwareBinaryAPIGetFirmwareBinaryRequest) (*FirmwareBinary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareBinary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareBinaryAPIService.GetFirmwareBinary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/binary/{firmwareBinaryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firmwareBinaryId"+"}", url.PathEscape(parameterValueToString(r.firmwareBinaryId, "firmwareBinaryId")), -1)

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

type FirmwareBinaryAPIUpdateFirmwareBinaryRequest struct {
	ctx context.Context
	ApiService *FirmwareBinaryAPIService
	firmwareBinaryId float32
	updateFirmwareBinary *UpdateFirmwareBinary
}

func (r FirmwareBinaryAPIUpdateFirmwareBinaryRequest) UpdateFirmwareBinary(updateFirmwareBinary UpdateFirmwareBinary) FirmwareBinaryAPIUpdateFirmwareBinaryRequest {
	r.updateFirmwareBinary = &updateFirmwareBinary
	return r
}

func (r FirmwareBinaryAPIUpdateFirmwareBinaryRequest) Execute() (*FirmwareBinary, *http.Response, error) {
	return r.ApiService.UpdateFirmwareBinaryExecute(r)
}

/*
UpdateFirmwareBinary Update Firmware Binary

Updates a firmware binary

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param firmwareBinaryId The firmware binary id
 @return FirmwareBinaryAPIUpdateFirmwareBinaryRequest
*/
func (a *FirmwareBinaryAPIService) UpdateFirmwareBinary(ctx context.Context, firmwareBinaryId float32) FirmwareBinaryAPIUpdateFirmwareBinaryRequest {
	return FirmwareBinaryAPIUpdateFirmwareBinaryRequest{
		ApiService: a,
		ctx: ctx,
		firmwareBinaryId: firmwareBinaryId,
	}
}

// Execute executes the request
//  @return FirmwareBinary
func (a *FirmwareBinaryAPIService) UpdateFirmwareBinaryExecute(r FirmwareBinaryAPIUpdateFirmwareBinaryRequest) (*FirmwareBinary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FirmwareBinary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareBinaryAPIService.UpdateFirmwareBinary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/binary/{firmwareBinaryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firmwareBinaryId"+"}", url.PathEscape(parameterValueToString(r.firmwareBinaryId, "firmwareBinaryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateFirmwareBinary == nil {
		return localVarReturnValue, nil, reportError("updateFirmwareBinary is required and must be specified")
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
	localVarPostBody = r.updateFirmwareBinary
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
