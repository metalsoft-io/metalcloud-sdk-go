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
)


// FirmwareBaselineSearchAPIService FirmwareBaselineSearchAPI service
type FirmwareBaselineSearchAPIService service

type FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest struct {
	ctx context.Context
	ApiService *FirmwareBaselineSearchAPIService
	searchFirmwareBinary *SearchFirmwareBinary
}

func (r FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest) SearchFirmwareBinary(searchFirmwareBinary SearchFirmwareBinary) FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest {
	r.searchFirmwareBinary = &searchFirmwareBinary
	return r
}

func (r FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest) Execute() (*IServerFirmwareBinaryResponse, *http.Response, error) {
	return r.ApiService.SearchFirmwareBaselinesExecute(r)
}

/*
SearchFirmwareBaselines Search Firmware Baselines

Returns a list of firmware baselines

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest
*/
func (a *FirmwareBaselineSearchAPIService) SearchFirmwareBaselines(ctx context.Context) FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest {
	return FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IServerFirmwareBinaryResponse
func (a *FirmwareBaselineSearchAPIService) SearchFirmwareBaselinesExecute(r FirmwareBaselineSearchAPISearchFirmwareBaselinesRequest) (*IServerFirmwareBinaryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IServerFirmwareBinaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareBaselineSearchAPIService.SearchFirmwareBaselines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/firmware/search/baseline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.searchFirmwareBinary == nil {
		return localVarReturnValue, nil, reportError("searchFirmwareBinary is required and must be specified")
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
	localVarPostBody = r.searchFirmwareBinary
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
