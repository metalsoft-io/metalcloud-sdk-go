# UpdateVMInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label for the VM Instance Group. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance. | [optional] 

## Methods

### NewUpdateVMInstanceGroup

`func NewUpdateVMInstanceGroup() *UpdateVMInstanceGroup`

NewUpdateVMInstanceGroup instantiates a new UpdateVMInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMInstanceGroupWithDefaults

`func NewUpdateVMInstanceGroupWithDefaults() *UpdateVMInstanceGroup`

NewUpdateVMInstanceGroupWithDefaults instantiates a new UpdateVMInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateVMInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateVMInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateVMInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateVMInstanceGroup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCustomVariables

`func (o *UpdateVMInstanceGroup) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateVMInstanceGroup) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateVMInstanceGroup) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateVMInstanceGroup) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


