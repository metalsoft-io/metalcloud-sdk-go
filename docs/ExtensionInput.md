# ExtensionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the input. | 
**Name** | **string** | Name of the input. | 
**InputType** | **string** | Type of the input. | 
**Options** | [**ExtensionInputOptions**](ExtensionInputOptions.md) |  | 
**SetOnly** | Pointer to **bool** | Flag to indicate if the input is required. | [optional] [default to false]
**Hidden** | Pointer to **bool** | Flag to indicate if the input is hidden in the UI. | [optional] [default to false]
**DefaultValue** | Pointer to **string** | Default value. | [optional] 

## Methods

### NewExtensionInput

`func NewExtensionInput(label string, name string, inputType string, options ExtensionInputOptions, ) *ExtensionInput`

NewExtensionInput instantiates a new ExtensionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInputWithDefaults

`func NewExtensionInputWithDefaults() *ExtensionInput`

NewExtensionInputWithDefaults instantiates a new ExtensionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ExtensionInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionInput) SetName(v string)`

SetName sets Name field to given value.


### GetInputType

`func (o *ExtensionInput) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ExtensionInput) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ExtensionInput) SetInputType(v string)`

SetInputType sets InputType field to given value.


### GetOptions

`func (o *ExtensionInput) GetOptions() ExtensionInputOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionInput) GetOptionsOk() (*ExtensionInputOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionInput) SetOptions(v ExtensionInputOptions)`

SetOptions sets Options field to given value.


### GetSetOnly

`func (o *ExtensionInput) GetSetOnly() bool`

GetSetOnly returns the SetOnly field if non-nil, zero value otherwise.

### GetSetOnlyOk

`func (o *ExtensionInput) GetSetOnlyOk() (*bool, bool)`

GetSetOnlyOk returns a tuple with the SetOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOnly

`func (o *ExtensionInput) SetSetOnly(v bool)`

SetSetOnly sets SetOnly field to given value.

### HasSetOnly

`func (o *ExtensionInput) HasSetOnly() bool`

HasSetOnly returns a boolean if a field has been set.

### GetHidden

`func (o *ExtensionInput) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ExtensionInput) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ExtensionInput) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ExtensionInput) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ExtensionInput) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ExtensionInput) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ExtensionInput) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ExtensionInput) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


