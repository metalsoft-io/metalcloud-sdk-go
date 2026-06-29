# Favicon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ico** | Pointer to **map[string]interface{}** | ICO format favicon (base64) | [optional] 
**Svg** | Pointer to **map[string]interface{}** | SVG format favicon (base64) | [optional] 
**Apple** | Pointer to **map[string]interface{}** | Apple touch icon (base64) | [optional] 

## Methods

### NewFavicon

`func NewFavicon() *Favicon`

NewFavicon instantiates a new Favicon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaviconWithDefaults

`func NewFaviconWithDefaults() *Favicon`

NewFaviconWithDefaults instantiates a new Favicon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIco

`func (o *Favicon) GetIco() map[string]interface{}`

GetIco returns the Ico field if non-nil, zero value otherwise.

### GetIcoOk

`func (o *Favicon) GetIcoOk() (*map[string]interface{}, bool)`

GetIcoOk returns a tuple with the Ico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIco

`func (o *Favicon) SetIco(v map[string]interface{})`

SetIco sets Ico field to given value.

### HasIco

`func (o *Favicon) HasIco() bool`

HasIco returns a boolean if a field has been set.

### GetSvg

`func (o *Favicon) GetSvg() map[string]interface{}`

GetSvg returns the Svg field if non-nil, zero value otherwise.

### GetSvgOk

`func (o *Favicon) GetSvgOk() (*map[string]interface{}, bool)`

GetSvgOk returns a tuple with the Svg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvg

`func (o *Favicon) SetSvg(v map[string]interface{})`

SetSvg sets Svg field to given value.

### HasSvg

`func (o *Favicon) HasSvg() bool`

HasSvg returns a boolean if a field has been set.

### GetApple

`func (o *Favicon) GetApple() map[string]interface{}`

GetApple returns the Apple field if non-nil, zero value otherwise.

### GetAppleOk

`func (o *Favicon) GetAppleOk() (*map[string]interface{}, bool)`

GetAppleOk returns a tuple with the Apple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApple

`func (o *Favicon) SetApple(v map[string]interface{})`

SetApple sets Apple field to given value.

### HasApple

`func (o *Favicon) HasApple() bool`

HasApple returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


