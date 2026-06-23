# BrandingBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favicon** | Pointer to [**Favicon**](Favicon.md) | Favicon images in multiple formats | [optional] 
**CompanyName** | Pointer to **map[string]interface{}** | Company name displayed in the UI | [optional] 
**UseUITemplate** | Pointer to **map[string]interface{}** | UI theme template identifier | [optional] 
**Logo** | Pointer to [**Logo**](Logo.md) | Logo images for different backgrounds | [optional] 

## Methods

### NewBrandingBase

`func NewBrandingBase() *BrandingBase`

NewBrandingBase instantiates a new BrandingBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingBaseWithDefaults

`func NewBrandingBaseWithDefaults() *BrandingBase`

NewBrandingBaseWithDefaults instantiates a new BrandingBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavicon

`func (o *BrandingBase) GetFavicon() Favicon`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BrandingBase) GetFaviconOk() (*Favicon, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BrandingBase) SetFavicon(v Favicon)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BrandingBase) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetCompanyName

`func (o *BrandingBase) GetCompanyName() map[string]interface{}`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BrandingBase) GetCompanyNameOk() (*map[string]interface{}, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BrandingBase) SetCompanyName(v map[string]interface{})`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BrandingBase) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetUseUITemplate

`func (o *BrandingBase) GetUseUITemplate() map[string]interface{}`

GetUseUITemplate returns the UseUITemplate field if non-nil, zero value otherwise.

### GetUseUITemplateOk

`func (o *BrandingBase) GetUseUITemplateOk() (*map[string]interface{}, bool)`

GetUseUITemplateOk returns a tuple with the UseUITemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUITemplate

`func (o *BrandingBase) SetUseUITemplate(v map[string]interface{})`

SetUseUITemplate sets UseUITemplate field to given value.

### HasUseUITemplate

`func (o *BrandingBase) HasUseUITemplate() bool`

HasUseUITemplate returns a boolean if a field has been set.

### GetLogo

`func (o *BrandingBase) GetLogo() Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *BrandingBase) GetLogoOk() (*Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *BrandingBase) SetLogo(v Logo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *BrandingBase) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


