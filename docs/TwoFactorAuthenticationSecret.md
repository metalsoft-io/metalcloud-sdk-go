# TwoFactorAuthenticationSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCode** | **string** | The QR code for the 2FA secret | 
**Secret2FA** | **string** | The 2FA secret | 

## Methods

### NewTwoFactorAuthenticationSecret

`func NewTwoFactorAuthenticationSecret(qrCode string, secret2FA string, ) *TwoFactorAuthenticationSecret`

NewTwoFactorAuthenticationSecret instantiates a new TwoFactorAuthenticationSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorAuthenticationSecretWithDefaults

`func NewTwoFactorAuthenticationSecretWithDefaults() *TwoFactorAuthenticationSecret`

NewTwoFactorAuthenticationSecretWithDefaults instantiates a new TwoFactorAuthenticationSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCode

`func (o *TwoFactorAuthenticationSecret) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *TwoFactorAuthenticationSecret) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *TwoFactorAuthenticationSecret) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.


### GetSecret2FA

`func (o *TwoFactorAuthenticationSecret) GetSecret2FA() string`

GetSecret2FA returns the Secret2FA field if non-nil, zero value otherwise.

### GetSecret2FAOk

`func (o *TwoFactorAuthenticationSecret) GetSecret2FAOk() (*string, bool)`

GetSecret2FAOk returns a tuple with the Secret2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret2FA

`func (o *TwoFactorAuthenticationSecret) SetSecret2FA(v string)`

SetSecret2FA sets Secret2FA field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


