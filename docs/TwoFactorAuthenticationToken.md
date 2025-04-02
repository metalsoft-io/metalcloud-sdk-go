# TwoFactorAuthenticationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The 2FA token | 

## Methods

### NewTwoFactorAuthenticationToken

`func NewTwoFactorAuthenticationToken(token string, ) *TwoFactorAuthenticationToken`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


