# CaptchaSiteKeysDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordReset** | Pointer to **string** | Site key for the password reset widget | [optional] 
**SignUp** | Pointer to **string** | Site key for the sign-up widget | [optional] 

## Methods

### NewCaptchaSiteKeysDto

`func NewCaptchaSiteKeysDto() *CaptchaSiteKeysDto`

NewCaptchaSiteKeysDto instantiates a new CaptchaSiteKeysDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptchaSiteKeysDtoWithDefaults

`func NewCaptchaSiteKeysDtoWithDefaults() *CaptchaSiteKeysDto`

NewCaptchaSiteKeysDtoWithDefaults instantiates a new CaptchaSiteKeysDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordReset

`func (o *CaptchaSiteKeysDto) GetPasswordReset() string`

GetPasswordReset returns the PasswordReset field if non-nil, zero value otherwise.

### GetPasswordResetOk

`func (o *CaptchaSiteKeysDto) GetPasswordResetOk() (*string, bool)`

GetPasswordResetOk returns a tuple with the PasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReset

`func (o *CaptchaSiteKeysDto) SetPasswordReset(v string)`

SetPasswordReset sets PasswordReset field to given value.

### HasPasswordReset

`func (o *CaptchaSiteKeysDto) HasPasswordReset() bool`

HasPasswordReset returns a boolean if a field has been set.

### GetSignUp

`func (o *CaptchaSiteKeysDto) GetSignUp() string`

GetSignUp returns the SignUp field if non-nil, zero value otherwise.

### GetSignUpOk

`func (o *CaptchaSiteKeysDto) GetSignUpOk() (*string, bool)`

GetSignUpOk returns a tuple with the SignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUp

`func (o *CaptchaSiteKeysDto) SetSignUp(v string)`

SetSignUp sets SignUp field to given value.

### HasSignUp

`func (o *CaptchaSiteKeysDto) HasSignUp() bool`

HasSignUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


