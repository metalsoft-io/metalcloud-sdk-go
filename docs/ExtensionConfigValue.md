# ExtensionConfigValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The config value label. | 
**Value** | [**ExtensionConfigValueValue**](ExtensionConfigValueValue.md) |  | 

## Methods

### NewExtensionConfigValue

`func NewExtensionConfigValue(label string, value ExtensionConfigValueValue, ) *ExtensionConfigValue`

NewExtensionConfigValue instantiates a new ExtensionConfigValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionConfigValueWithDefaults

`func NewExtensionConfigValueWithDefaults() *ExtensionConfigValue`

NewExtensionConfigValueWithDefaults instantiates a new ExtensionConfigValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionConfigValue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionConfigValue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionConfigValue) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *ExtensionConfigValue) GetValue() ExtensionConfigValueValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExtensionConfigValue) GetValueOk() (*ExtensionConfigValueValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExtensionConfigValue) SetValue(v ExtensionConfigValueValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


