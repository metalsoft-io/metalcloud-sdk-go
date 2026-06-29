# CaptchaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCaptchaEnabled** | **bool** | Flag to indicate if CAPTCHA is enabled | [default to false]
**SiteKeys** | [**CaptchaSiteKeysDto**](CaptchaSiteKeysDto.md) | Site keys for each CAPTCHA widget | 

## Methods

### NewCaptchaConfig

`func NewCaptchaConfig(isCaptchaEnabled bool, siteKeys CaptchaSiteKeysDto, ) *CaptchaConfig`

NewCaptchaConfig instantiates a new CaptchaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptchaConfigWithDefaults

`func NewCaptchaConfigWithDefaults() *CaptchaConfig`

NewCaptchaConfigWithDefaults instantiates a new CaptchaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCaptchaEnabled

`func (o *CaptchaConfig) GetIsCaptchaEnabled() bool`

GetIsCaptchaEnabled returns the IsCaptchaEnabled field if non-nil, zero value otherwise.

### GetIsCaptchaEnabledOk

`func (o *CaptchaConfig) GetIsCaptchaEnabledOk() (*bool, bool)`

GetIsCaptchaEnabledOk returns a tuple with the IsCaptchaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaptchaEnabled

`func (o *CaptchaConfig) SetIsCaptchaEnabled(v bool)`

SetIsCaptchaEnabled sets IsCaptchaEnabled field to given value.


### GetSiteKeys

`func (o *CaptchaConfig) GetSiteKeys() CaptchaSiteKeysDto`

GetSiteKeys returns the SiteKeys field if non-nil, zero value otherwise.

### GetSiteKeysOk

`func (o *CaptchaConfig) GetSiteKeysOk() (*CaptchaSiteKeysDto, bool)`

GetSiteKeysOk returns a tuple with the SiteKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKeys

`func (o *CaptchaConfig) SetSiteKeys(v CaptchaSiteKeysDto)`

SetSiteKeys sets SiteKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


