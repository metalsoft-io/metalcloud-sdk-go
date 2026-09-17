# UpdateContainerInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label for the VM Instance Group. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance. | [optional] 

## Methods

### NewUpdateContainerInstanceGroup

`func NewUpdateContainerInstanceGroup() *UpdateContainerInstanceGroup`

NewUpdateContainerInstanceGroup instantiates a new UpdateContainerInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerInstanceGroupWithDefaults

`func NewUpdateContainerInstanceGroupWithDefaults() *UpdateContainerInstanceGroup`

NewUpdateContainerInstanceGroupWithDefaults instantiates a new UpdateContainerInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateContainerInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateContainerInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateContainerInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateContainerInstanceGroup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCustomVariables

`func (o *UpdateContainerInstanceGroup) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateContainerInstanceGroup) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateContainerInstanceGroup) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateContainerInstanceGroup) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


