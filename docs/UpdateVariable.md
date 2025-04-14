# UpdateVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The variable name. | 
**Value** | **map[string]interface{}** | The variable value. | 
**Usage** | Pointer to [**VariableUsageType**](VariableUsageType.md) | Variable usage type. | [optional] 

## Methods

### NewUpdateVariable

`func NewUpdateVariable(name string, value map[string]interface{}, ) *UpdateVariable`

NewUpdateVariable instantiates a new UpdateVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVariableWithDefaults

`func NewUpdateVariableWithDefaults() *UpdateVariable`

NewUpdateVariableWithDefaults instantiates a new UpdateVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVariable) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *UpdateVariable) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateVariable) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateVariable) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetUsage

`func (o *UpdateVariable) GetUsage() VariableUsageType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UpdateVariable) GetUsageOk() (*VariableUsageType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UpdateVariable) SetUsage(v VariableUsageType)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UpdateVariable) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


