# ExtensionDefinitionInputsDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the input. | 
**Name** | **string** | Name of the input. | 
**InputType** | [**ExtensionInputType**](ExtensionInputType.md) |  | 
**SetOnly** | Pointer to **bool** | Flag to indicate if the input is required. | [optional] [default to false]
**Hidden** | Pointer to **bool** | Flag to indicate if the input is hidden in the UI. | [optional] [default to false]
**IsPassword** | Pointer to **bool** | Flag to indicate if the input is a password. Only to be used with string input type. | [optional] [default to false]
**DefaultValue** | Pointer to [**ExtensionInputStringDefaultValue**](ExtensionInputStringDefaultValue.md) |  | [optional] 
**Options** | **map[string]interface{}** |  | 

## Methods

### NewExtensionDefinitionInputsDataItem

`func NewExtensionDefinitionInputsDataItem(label string, name string, inputType ExtensionInputType, options map[string]interface{}, ) *ExtensionDefinitionInputsDataItem`

NewExtensionDefinitionInputsDataItem instantiates a new ExtensionDefinitionInputsDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionDefinitionInputsDataItemWithDefaults

`func NewExtensionDefinitionInputsDataItemWithDefaults() *ExtensionDefinitionInputsDataItem`

NewExtensionDefinitionInputsDataItemWithDefaults instantiates a new ExtensionDefinitionInputsDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionDefinitionInputsDataItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionDefinitionInputsDataItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionDefinitionInputsDataItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ExtensionDefinitionInputsDataItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionDefinitionInputsDataItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionDefinitionInputsDataItem) SetName(v string)`

SetName sets Name field to given value.


### GetInputType

`func (o *ExtensionDefinitionInputsDataItem) GetInputType() ExtensionInputType`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ExtensionDefinitionInputsDataItem) GetInputTypeOk() (*ExtensionInputType, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ExtensionDefinitionInputsDataItem) SetInputType(v ExtensionInputType)`

SetInputType sets InputType field to given value.


### GetSetOnly

`func (o *ExtensionDefinitionInputsDataItem) GetSetOnly() bool`

GetSetOnly returns the SetOnly field if non-nil, zero value otherwise.

### GetSetOnlyOk

`func (o *ExtensionDefinitionInputsDataItem) GetSetOnlyOk() (*bool, bool)`

GetSetOnlyOk returns a tuple with the SetOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOnly

`func (o *ExtensionDefinitionInputsDataItem) SetSetOnly(v bool)`

SetSetOnly sets SetOnly field to given value.

### HasSetOnly

`func (o *ExtensionDefinitionInputsDataItem) HasSetOnly() bool`

HasSetOnly returns a boolean if a field has been set.

### GetHidden

`func (o *ExtensionDefinitionInputsDataItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ExtensionDefinitionInputsDataItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ExtensionDefinitionInputsDataItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ExtensionDefinitionInputsDataItem) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetIsPassword

`func (o *ExtensionDefinitionInputsDataItem) GetIsPassword() bool`

GetIsPassword returns the IsPassword field if non-nil, zero value otherwise.

### GetIsPasswordOk

`func (o *ExtensionDefinitionInputsDataItem) GetIsPasswordOk() (*bool, bool)`

GetIsPasswordOk returns a tuple with the IsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassword

`func (o *ExtensionDefinitionInputsDataItem) SetIsPassword(v bool)`

SetIsPassword sets IsPassword field to given value.

### HasIsPassword

`func (o *ExtensionDefinitionInputsDataItem) HasIsPassword() bool`

HasIsPassword returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ExtensionDefinitionInputsDataItem) GetDefaultValue() ExtensionInputStringDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ExtensionDefinitionInputsDataItem) GetDefaultValueOk() (*ExtensionInputStringDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ExtensionDefinitionInputsDataItem) SetDefaultValue(v ExtensionInputStringDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ExtensionDefinitionInputsDataItem) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetOptions

`func (o *ExtensionDefinitionInputsDataItem) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionDefinitionInputsDataItem) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionDefinitionInputsDataItem) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


