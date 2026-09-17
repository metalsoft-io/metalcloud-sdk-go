# ExtensionDefinitionConfigVarsDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the input. | 
**Name** | **string** | Name of the input. | 
**HelpText** | Pointer to **string** | Help text describing the input. | [optional] 
**InputType** | **string** |  | 
**SetOnly** | Pointer to **bool** | Flag to indicate if the input can be set only during extension instance creation. | [optional] [default to false]
**Hidden** | Pointer to **bool** | Flag to indicate if the input is hidden in the UI. | [optional] [default to false]
**IsPassword** | Pointer to **bool** | Flag to indicate if the input is a password. Only to be used with string input type. | [optional] [default to false]
**DefaultValue** | Pointer to [**ExtensionInputStringDefaultValue**](ExtensionInputStringDefaultValue.md) |  | [optional] 
**Options** | **map[string]interface{}** |  | 

## Methods

### NewExtensionDefinitionConfigVarsDataItem

`func NewExtensionDefinitionConfigVarsDataItem(label string, name string, inputType string, options map[string]interface{}, ) *ExtensionDefinitionConfigVarsDataItem`

NewExtensionDefinitionConfigVarsDataItem instantiates a new ExtensionDefinitionConfigVarsDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionDefinitionConfigVarsDataItemWithDefaults

`func NewExtensionDefinitionConfigVarsDataItemWithDefaults() *ExtensionDefinitionConfigVarsDataItem`

NewExtensionDefinitionConfigVarsDataItemWithDefaults instantiates a new ExtensionDefinitionConfigVarsDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionDefinitionConfigVarsDataItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionDefinitionConfigVarsDataItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ExtensionDefinitionConfigVarsDataItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionDefinitionConfigVarsDataItem) SetName(v string)`

SetName sets Name field to given value.


### GetHelpText

`func (o *ExtensionDefinitionConfigVarsDataItem) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *ExtensionDefinitionConfigVarsDataItem) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *ExtensionDefinitionConfigVarsDataItem) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetInputType

`func (o *ExtensionDefinitionConfigVarsDataItem) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ExtensionDefinitionConfigVarsDataItem) SetInputType(v string)`

SetInputType sets InputType field to given value.


### GetSetOnly

`func (o *ExtensionDefinitionConfigVarsDataItem) GetSetOnly() bool`

GetSetOnly returns the SetOnly field if non-nil, zero value otherwise.

### GetSetOnlyOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetSetOnlyOk() (*bool, bool)`

GetSetOnlyOk returns a tuple with the SetOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOnly

`func (o *ExtensionDefinitionConfigVarsDataItem) SetSetOnly(v bool)`

SetSetOnly sets SetOnly field to given value.

### HasSetOnly

`func (o *ExtensionDefinitionConfigVarsDataItem) HasSetOnly() bool`

HasSetOnly returns a boolean if a field has been set.

### GetHidden

`func (o *ExtensionDefinitionConfigVarsDataItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ExtensionDefinitionConfigVarsDataItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ExtensionDefinitionConfigVarsDataItem) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetIsPassword

`func (o *ExtensionDefinitionConfigVarsDataItem) GetIsPassword() bool`

GetIsPassword returns the IsPassword field if non-nil, zero value otherwise.

### GetIsPasswordOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetIsPasswordOk() (*bool, bool)`

GetIsPasswordOk returns a tuple with the IsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassword

`func (o *ExtensionDefinitionConfigVarsDataItem) SetIsPassword(v bool)`

SetIsPassword sets IsPassword field to given value.

### HasIsPassword

`func (o *ExtensionDefinitionConfigVarsDataItem) HasIsPassword() bool`

HasIsPassword returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ExtensionDefinitionConfigVarsDataItem) GetDefaultValue() ExtensionInputStringDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetDefaultValueOk() (*ExtensionInputStringDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ExtensionDefinitionConfigVarsDataItem) SetDefaultValue(v ExtensionInputStringDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ExtensionDefinitionConfigVarsDataItem) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetOptions

`func (o *ExtensionDefinitionConfigVarsDataItem) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionDefinitionConfigVarsDataItem) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionDefinitionConfigVarsDataItem) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


