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


// UserAPIService UserAPI service
type UserAPIService service

type UserAPIDisableUser2FARequest struct {
	ctx context.Context
	ApiService *UserAPIService
	ifMatch *string
}

// Entity tag
func (r UserAPIDisableUser2FARequest) IfMatch(ifMatch string) UserAPIDisableUser2FARequest {
	r.ifMatch = &ifMatch
	return r
}

func (r UserAPIDisableUser2FARequest) Execute() (*http.Response, error) {
	return r.ApiService.DisableUser2FAExecute(r)
}

/*
DisableUser2FA Disable 2FA

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIDisableUser2FARequest
*/
func (a *UserAPIService) DisableUser2FA(ctx context.Context) UserAPIDisableUser2FARequest {
	return UserAPIDisableUser2FARequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserAPIService) DisableUser2FAExecute(r UserAPIDisableUser2FARequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.DisableUser2FA")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/2fa-disable"

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

type UserAPIEnableUser2FARequest struct {
	ctx context.Context
	ApiService *UserAPIService
	twoFactorAuthenticationToken *TwoFactorAuthenticationToken
	ifMatch *string
}

func (r UserAPIEnableUser2FARequest) TwoFactorAuthenticationToken(twoFactorAuthenticationToken TwoFactorAuthenticationToken) UserAPIEnableUser2FARequest {
	r.twoFactorAuthenticationToken = &twoFactorAuthenticationToken
	return r
}

// Entity tag
func (r UserAPIEnableUser2FARequest) IfMatch(ifMatch string) UserAPIEnableUser2FARequest {
	r.ifMatch = &ifMatch
	return r
}

func (r UserAPIEnableUser2FARequest) Execute() (*http.Response, error) {
	return r.ApiService.EnableUser2FAExecute(r)
}

/*
EnableUser2FA Enable 2FA

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIEnableUser2FARequest
*/
func (a *UserAPIService) EnableUser2FA(ctx context.Context) UserAPIEnableUser2FARequest {
	return UserAPIEnableUser2FARequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserAPIService) EnableUser2FAExecute(r UserAPIEnableUser2FARequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.EnableUser2FA")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/2fa-enable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.twoFactorAuthenticationToken == nil {
		return nil, reportError("twoFactorAuthenticationToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.twoFactorAuthenticationToken
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

type UserAPIGenerateUser2FASecretRequest struct {
	ctx context.Context
	ApiService *UserAPIService
}

func (r UserAPIGenerateUser2FASecretRequest) Execute() (*TwoFactorAuthenticationSecret, *http.Response, error) {
	return r.ApiService.GenerateUser2FASecretExecute(r)
}

/*
GenerateUser2FASecret Generate 2FA secret

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIGenerateUser2FASecretRequest
*/
func (a *UserAPIService) GenerateUser2FASecret(ctx context.Context) UserAPIGenerateUser2FASecretRequest {
	return UserAPIGenerateUser2FASecretRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TwoFactorAuthenticationSecret
func (a *UserAPIService) GenerateUser2FASecretExecute(r UserAPIGenerateUser2FASecretRequest) (*TwoFactorAuthenticationSecret, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TwoFactorAuthenticationSecret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.GenerateUser2FASecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/2fa-generate"

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

type UserAPIGetUserApiKeyRequest struct {
	ctx context.Context
	ApiService *UserAPIService
}

func (r UserAPIGetUserApiKeyRequest) Execute() (*UserApiKey, *http.Response, error) {
	return r.ApiService.GetUserApiKeyExecute(r)
}

/*
GetUserApiKey Get user API key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIGetUserApiKeyRequest
*/
func (a *UserAPIService) GetUserApiKey(ctx context.Context) UserAPIGetUserApiKeyRequest {
	return UserAPIGetUserApiKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserApiKey
func (a *UserAPIService) GetUserApiKeyExecute(r UserAPIGetUserApiKeyRequest) (*UserApiKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.GetUserApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/api-key"

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

type UserAPIInitiateEmailChangeRequest struct {
	ctx context.Context
	ApiService *UserAPIService
	changeUserEmail *ChangeUserEmail
}

func (r UserAPIInitiateEmailChangeRequest) ChangeUserEmail(changeUserEmail ChangeUserEmail) UserAPIInitiateEmailChangeRequest {
	r.changeUserEmail = &changeUserEmail
	return r
}

func (r UserAPIInitiateEmailChangeRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.InitiateEmailChangeExecute(r)
}

/*
InitiateEmailChange Initiates the user email change

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIInitiateEmailChangeRequest
*/
func (a *UserAPIService) InitiateEmailChange(ctx context.Context) UserAPIInitiateEmailChangeRequest {
	return UserAPIInitiateEmailChangeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return User
func (a *UserAPIService) InitiateEmailChangeExecute(r UserAPIInitiateEmailChangeRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.InitiateEmailChange")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/initiate-email-change"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.changeUserEmail == nil {
		return localVarReturnValue, nil, reportError("changeUserEmail is required and must be specified")
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
	localVarPostBody = r.changeUserEmail
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

type UserAPIInitiatePasswordResetRequest struct {
	ctx context.Context
	ApiService *UserAPIService
	passwordReset *PasswordReset
}

// The password reset request
func (r UserAPIInitiatePasswordResetRequest) PasswordReset(passwordReset PasswordReset) UserAPIInitiatePasswordResetRequest {
	r.passwordReset = &passwordReset
	return r
}

func (r UserAPIInitiatePasswordResetRequest) Execute() (*http.Response, error) {
	return r.ApiService.InitiatePasswordResetExecute(r)
}

/*
InitiatePasswordReset Initiate reset user password

Initiates the reset of the user password.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIInitiatePasswordResetRequest
*/
func (a *UserAPIService) InitiatePasswordReset(ctx context.Context) UserAPIInitiatePasswordResetRequest {
	return UserAPIInitiatePasswordResetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserAPIService) InitiatePasswordResetExecute(r UserAPIInitiatePasswordResetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.InitiatePasswordReset")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/initiate-password-reset"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.passwordReset == nil {
		return nil, reportError("passwordReset is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.passwordReset
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

type UserAPIRegenerateUserApiKeyRequest struct {
	ctx context.Context
	ApiService *UserAPIService
	ifMatch *string
}

// Entity tag
func (r UserAPIRegenerateUserApiKeyRequest) IfMatch(ifMatch string) UserAPIRegenerateUserApiKeyRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r UserAPIRegenerateUserApiKeyRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.RegenerateUserApiKeyExecute(r)
}

/*
RegenerateUserApiKey Regenerate user API key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIRegenerateUserApiKeyRequest
*/
func (a *UserAPIService) RegenerateUserApiKey(ctx context.Context) UserAPIRegenerateUserApiKeyRequest {
	return UserAPIRegenerateUserApiKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return User
func (a *UserAPIService) RegenerateUserApiKeyExecute(r UserAPIRegenerateUserApiKeyRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.RegenerateUserApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/regenerate-api-key"

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
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
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

type UserAPIRegenerateUserJwtSaltRequest struct {
	ctx context.Context
	ApiService *UserAPIService
	ifMatch *string
}

// Entity tag
func (r UserAPIRegenerateUserJwtSaltRequest) IfMatch(ifMatch string) UserAPIRegenerateUserJwtSaltRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r UserAPIRegenerateUserJwtSaltRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.RegenerateUserJwtSaltExecute(r)
}

/*
RegenerateUserJwtSalt Regenerate user JWT salt. Also logs out all user sessions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIRegenerateUserJwtSaltRequest
*/
func (a *UserAPIService) RegenerateUserJwtSalt(ctx context.Context) UserAPIRegenerateUserJwtSaltRequest {
	return UserAPIRegenerateUserJwtSaltRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return User
func (a *UserAPIService) RegenerateUserJwtSaltExecute(r UserAPIRegenerateUserJwtSaltRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.RegenerateUserJwtSalt")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/regenerate-jwt-salt"

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
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "simple", "")
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

type UserAPIUpdateUserPasswordRequest struct {
	ctx context.Context
	ApiService *UserAPIService
	userUpdatePassword *UserUpdatePassword
	ifMatch *string
}

func (r UserAPIUpdateUserPasswordRequest) UserUpdatePassword(userUpdatePassword UserUpdatePassword) UserAPIUpdateUserPasswordRequest {
	r.userUpdatePassword = &userUpdatePassword
	return r
}

// Entity tag
func (r UserAPIUpdateUserPasswordRequest) IfMatch(ifMatch string) UserAPIUpdateUserPasswordRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r UserAPIUpdateUserPasswordRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.UpdateUserPasswordExecute(r)
}

/*
UpdateUserPassword Change user password

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIUpdateUserPasswordRequest
*/
func (a *UserAPIService) UpdateUserPassword(ctx context.Context) UserAPIUpdateUserPasswordRequest {
	return UserAPIUpdateUserPasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return User
func (a *UserAPIService) UpdateUserPasswordExecute(r UserAPIUpdateUserPasswordRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.UpdateUserPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/change-password"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userUpdatePassword == nil {
		return localVarReturnValue, nil, reportError("userUpdatePassword is required and must be specified")
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
	localVarPostBody = r.userUpdatePassword
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

type UserAPIUserControllerHandleEmailVerifyRequest struct {
	ctx context.Context
	ApiService *UserAPIService
	token *string
}

// The encrypted request token.
func (r UserAPIUserControllerHandleEmailVerifyRequest) Token(token string) UserAPIUserControllerHandleEmailVerifyRequest {
	r.token = &token
	return r
}

func (r UserAPIUserControllerHandleEmailVerifyRequest) Execute() (*http.Response, error) {
	return r.ApiService.UserControllerHandleEmailVerifyExecute(r)
}

/*
UserControllerHandleEmailVerify Handle user email verify action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIUserControllerHandleEmailVerifyRequest
*/
func (a *UserAPIService) UserControllerHandleEmailVerify(ctx context.Context) UserAPIUserControllerHandleEmailVerifyRequest {
	return UserAPIUserControllerHandleEmailVerifyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserAPIService) UserControllerHandleEmailVerifyExecute(r UserAPIUserControllerHandleEmailVerifyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.UserControllerHandleEmailVerify")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/verify-email"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
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

type UserAPIUserControllerHandleUserResetPasswordRequest struct {
	ctx context.Context
	ApiService *UserAPIService
	token *string
}

// The encrypted request token.
func (r UserAPIUserControllerHandleUserResetPasswordRequest) Token(token string) UserAPIUserControllerHandleUserResetPasswordRequest {
	r.token = &token
	return r
}

func (r UserAPIUserControllerHandleUserResetPasswordRequest) Execute() (*http.Response, error) {
	return r.ApiService.UserControllerHandleUserResetPasswordExecute(r)
}

/*
UserControllerHandleUserResetPassword Handle user reset password action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserAPIUserControllerHandleUserResetPasswordRequest
*/
func (a *UserAPIService) UserControllerHandleUserResetPassword(ctx context.Context) UserAPIUserControllerHandleUserResetPasswordRequest {
	return UserAPIUserControllerHandleUserResetPasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserAPIService) UserControllerHandleUserResetPasswordExecute(r UserAPIUserControllerHandleUserResetPasswordRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.UserControllerHandleUserResetPassword")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user/actions/reset-password"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
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
