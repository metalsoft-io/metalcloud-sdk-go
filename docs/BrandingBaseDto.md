# BrandingBaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favicon** | Pointer to [**FaviconDto**](FaviconDto.md) | Favicon images in multiple formats | [optional] 
**CompanyName** | Pointer to **map[string]interface{}** | Company name displayed in the UI | [optional] 
**UseUITemplate** | Pointer to **map[string]interface{}** | UI theme template identifier | [optional] 
**Logo** | Pointer to [**LogoDto**](LogoDto.md) | Logo images for different backgrounds | [optional] 

## Methods

### NewBrandingBaseDto

`func NewBrandingBaseDto() *BrandingBaseDto`

NewBrandingBaseDto instantiates a new BrandingBaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingBaseDtoWithDefaults

`func NewBrandingBaseDtoWithDefaults() *BrandingBaseDto`

NewBrandingBaseDtoWithDefaults instantiates a new BrandingBaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavicon

`func (o *BrandingBaseDto) GetFavicon() FaviconDto`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BrandingBaseDto) GetFaviconOk() (*FaviconDto, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BrandingBaseDto) SetFavicon(v FaviconDto)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BrandingBaseDto) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetCompanyName

`func (o *BrandingBaseDto) GetCompanyName() map[string]interface{}`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BrandingBaseDto) GetCompanyNameOk() (*map[string]interface{}, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BrandingBaseDto) SetCompanyName(v map[string]interface{})`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BrandingBaseDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetUseUITemplate

`func (o *BrandingBaseDto) GetUseUITemplate() map[string]interface{}`

GetUseUITemplate returns the UseUITemplate field if non-nil, zero value otherwise.

### GetUseUITemplateOk

`func (o *BrandingBaseDto) GetUseUITemplateOk() (*map[string]interface{}, bool)`

GetUseUITemplateOk returns a tuple with the UseUITemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUITemplate

`func (o *BrandingBaseDto) SetUseUITemplate(v map[string]interface{})`

SetUseUITemplate sets UseUITemplate field to given value.

### HasUseUITemplate

`func (o *BrandingBaseDto) HasUseUITemplate() bool`

HasUseUITemplate returns a boolean if a field has been set.

### GetLogo

`func (o *BrandingBaseDto) GetLogo() LogoDto`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *BrandingBaseDto) GetLogoOk() (*LogoDto, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *BrandingBaseDto) SetLogo(v LogoDto)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *BrandingBaseDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


