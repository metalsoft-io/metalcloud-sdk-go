# TwoFactorAuthenticationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The 2FA token. Required unless a recovery code is supplied instead. | [optional] 
**RecoveryCode** | Pointer to **string** | A single-use recovery code, used in place of the 2FA token when the authenticator is unavailable. | [optional] 

## Methods

### NewTwoFactorAuthenticationToken

`func NewTwoFactorAuthenticationToken() *TwoFactorAuthenticationToken`

NewTwoFactorAuthenticationToken instantiates a new TwoFactorAuthenticationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorAuthenticationTokenWithDefaults

`func NewTwoFactorAuthenticationTokenWithDefaults() *TwoFactorAuthenticationToken`

NewTwoFactorAuthenticationTokenWithDefaults instantiates a new TwoFactorAuthenticationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TwoFactorAuthenticationToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TwoFactorAuthenticationToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TwoFactorAuthenticationToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TwoFactorAuthenticationToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetRecoveryCode

`func (o *TwoFactorAuthenticationToken) GetRecoveryCode() string`

GetRecoveryCode returns the RecoveryCode field if non-nil, zero value otherwise.

### GetRecoveryCodeOk

`func (o *TwoFactorAuthenticationToken) GetRecoveryCodeOk() (*string, bool)`

GetRecoveryCodeOk returns a tuple with the RecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCode

`func (o *TwoFactorAuthenticationToken) SetRecoveryCode(v string)`

SetRecoveryCode sets RecoveryCode field to given value.

### HasRecoveryCode

`func (o *TwoFactorAuthenticationToken) HasRecoveryCode() bool`

HasRecoveryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


