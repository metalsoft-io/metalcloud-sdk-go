# CreateVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The variable name. | 
**Value** | **map[string]interface{}** | The variable value. | 
**Usage** | Pointer to [**VariableUsageType**](VariableUsageType.md) | Variable usage type. | [optional] 

## Methods

### NewCreateVariable

`func NewCreateVariable(name string, value map[string]interface{}, ) *CreateVariable`

NewCreateVariable instantiates a new CreateVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVariableWithDefaults

`func NewCreateVariableWithDefaults() *CreateVariable`

NewCreateVariableWithDefaults instantiates a new CreateVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVariable) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CreateVariable) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateVariable) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateVariable) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetUsage

`func (o *CreateVariable) GetUsage() VariableUsageType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CreateVariable) GetUsageOk() (*VariableUsageType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CreateVariable) SetUsage(v VariableUsageType)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CreateVariable) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


