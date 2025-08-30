# ExtensionVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The variable label. | 
**Value** | [**ExtensionVariableValue**](ExtensionVariableValue.md) |  | 

## Methods

### NewExtensionVariable

`func NewExtensionVariable(label string, value ExtensionVariableValue, ) *ExtensionVariable`

NewExtensionVariable instantiates a new ExtensionVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionVariableWithDefaults

`func NewExtensionVariableWithDefaults() *ExtensionVariable`

NewExtensionVariableWithDefaults instantiates a new ExtensionVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionVariable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionVariable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionVariable) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *ExtensionVariable) GetValue() ExtensionVariableValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExtensionVariable) GetValueOk() (*ExtensionVariableValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExtensionVariable) SetValue(v ExtensionVariableValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


