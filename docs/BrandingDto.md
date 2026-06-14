# BrandingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favicon** | Pointer to [**FaviconDto**](FaviconDto.md) | Favicon images in multiple formats | [optional] 
**CompanyName** | Pointer to **map[string]interface{}** | Company name displayed in the UI | [optional] 
**UseUITemplate** | Pointer to **map[string]interface{}** | UI theme template identifier | [optional] 
**Logo** | Pointer to [**LogoDto**](LogoDto.md) | Logo images for different backgrounds | [optional] 
**Default** | Pointer to [**BrandingBaseDto**](BrandingBaseDto.md) | Default branding overrides | [optional] 

## Methods

### NewBrandingDto

`func NewBrandingDto() *BrandingDto`

NewBrandingDto instantiates a new BrandingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingDtoWithDefaults

`func NewBrandingDtoWithDefaults() *BrandingDto`

NewBrandingDtoWithDefaults instantiates a new BrandingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavicon

`func (o *BrandingDto) GetFavicon() FaviconDto`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BrandingDto) GetFaviconOk() (*FaviconDto, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BrandingDto) SetFavicon(v FaviconDto)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BrandingDto) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetCompanyName

`func (o *BrandingDto) GetCompanyName() map[string]interface{}`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BrandingDto) GetCompanyNameOk() (*map[string]interface{}, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BrandingDto) SetCompanyName(v map[string]interface{})`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BrandingDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetUseUITemplate

`func (o *BrandingDto) GetUseUITemplate() map[string]interface{}`

GetUseUITemplate returns the UseUITemplate field if non-nil, zero value otherwise.

### GetUseUITemplateOk

`func (o *BrandingDto) GetUseUITemplateOk() (*map[string]interface{}, bool)`

GetUseUITemplateOk returns a tuple with the UseUITemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUITemplate

`func (o *BrandingDto) SetUseUITemplate(v map[string]interface{})`

SetUseUITemplate sets UseUITemplate field to given value.

### HasUseUITemplate

`func (o *BrandingDto) HasUseUITemplate() bool`

HasUseUITemplate returns a boolean if a field has been set.

### GetLogo

`func (o *BrandingDto) GetLogo() LogoDto`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *BrandingDto) GetLogoOk() (*LogoDto, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *BrandingDto) SetLogo(v LogoDto)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *BrandingDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetDefault

`func (o *BrandingDto) GetDefault() BrandingBaseDto`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BrandingDto) GetDefaultOk() (*BrandingBaseDto, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BrandingDto) SetDefault(v BrandingBaseDto)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *BrandingDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


