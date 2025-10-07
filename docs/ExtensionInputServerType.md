# ExtensionInputServerType

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
**Options** | [**ExtensionInputOptionServerType**](ExtensionInputOptionServerType.md) |  | 

## Methods

### NewExtensionInputServerType

`func NewExtensionInputServerType(label string, name string, inputType ExtensionInputType, options ExtensionInputOptionServerType, ) *ExtensionInputServerType`

NewExtensionInputServerType instantiates a new ExtensionInputServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInputServerTypeWithDefaults

`func NewExtensionInputServerTypeWithDefaults() *ExtensionInputServerType`

NewExtensionInputServerTypeWithDefaults instantiates a new ExtensionInputServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionInputServerType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInputServerType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInputServerType) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ExtensionInputServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionInputServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionInputServerType) SetName(v string)`

SetName sets Name field to given value.


### GetInputType

`func (o *ExtensionInputServerType) GetInputType() ExtensionInputType`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ExtensionInputServerType) GetInputTypeOk() (*ExtensionInputType, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ExtensionInputServerType) SetInputType(v ExtensionInputType)`

SetInputType sets InputType field to given value.


### GetSetOnly

`func (o *ExtensionInputServerType) GetSetOnly() bool`

GetSetOnly returns the SetOnly field if non-nil, zero value otherwise.

### GetSetOnlyOk

`func (o *ExtensionInputServerType) GetSetOnlyOk() (*bool, bool)`

GetSetOnlyOk returns a tuple with the SetOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOnly

`func (o *ExtensionInputServerType) SetSetOnly(v bool)`

SetSetOnly sets SetOnly field to given value.

### HasSetOnly

`func (o *ExtensionInputServerType) HasSetOnly() bool`

HasSetOnly returns a boolean if a field has been set.

### GetHidden

`func (o *ExtensionInputServerType) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ExtensionInputServerType) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ExtensionInputServerType) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ExtensionInputServerType) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetIsPassword

`func (o *ExtensionInputServerType) GetIsPassword() bool`

GetIsPassword returns the IsPassword field if non-nil, zero value otherwise.

### GetIsPasswordOk

`func (o *ExtensionInputServerType) GetIsPasswordOk() (*bool, bool)`

GetIsPasswordOk returns a tuple with the IsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassword

`func (o *ExtensionInputServerType) SetIsPassword(v bool)`

SetIsPassword sets IsPassword field to given value.

### HasIsPassword

`func (o *ExtensionInputServerType) HasIsPassword() bool`

HasIsPassword returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ExtensionInputServerType) GetDefaultValue() ExtensionInputStringDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ExtensionInputServerType) GetDefaultValueOk() (*ExtensionInputStringDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ExtensionInputServerType) SetDefaultValue(v ExtensionInputStringDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ExtensionInputServerType) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetOptions

`func (o *ExtensionInputServerType) GetOptions() ExtensionInputOptionServerType`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionInputServerType) GetOptionsOk() (*ExtensionInputOptionServerType, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionInputServerType) SetOptions(v ExtensionInputOptionServerType)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


