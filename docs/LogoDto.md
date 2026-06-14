# LogoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Light** | Pointer to **map[string]interface{}** | Logo for light backgrounds (base64 SVG) | [optional] 
**Dark** | Pointer to **map[string]interface{}** | Logo for dark backgrounds (base64 SVG) | [optional] 
**Email** | Pointer to **map[string]interface{}** | Logo for email templates (base64 SVG) | [optional] 

## Methods

### NewLogoDto

`func NewLogoDto() *LogoDto`

NewLogoDto instantiates a new LogoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoDtoWithDefaults

`func NewLogoDtoWithDefaults() *LogoDto`

NewLogoDtoWithDefaults instantiates a new LogoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLight

`func (o *LogoDto) GetLight() map[string]interface{}`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *LogoDto) GetLightOk() (*map[string]interface{}, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *LogoDto) SetLight(v map[string]interface{})`

SetLight sets Light field to given value.

### HasLight

`func (o *LogoDto) HasLight() bool`

HasLight returns a boolean if a field has been set.

### GetDark

`func (o *LogoDto) GetDark() map[string]interface{}`

GetDark returns the Dark field if non-nil, zero value otherwise.

### GetDarkOk

`func (o *LogoDto) GetDarkOk() (*map[string]interface{}, bool)`

GetDarkOk returns a tuple with the Dark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDark

`func (o *LogoDto) SetDark(v map[string]interface{})`

SetDark sets Dark field to given value.

### HasDark

`func (o *LogoDto) HasDark() bool`

HasDark returns a boolean if a field has been set.

### GetEmail

`func (o *LogoDto) GetEmail() map[string]interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LogoDto) GetEmailOk() (*map[string]interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LogoDto) SetEmail(v map[string]interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LogoDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


