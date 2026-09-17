# UpdateContainerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label for the Instance Group. | [optional] 
**DiskSizeGB** | Pointer to **float32** | Disk size in GB for the Instance. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the Instance. | [optional] 

## Methods

### NewUpdateContainerInstance

`func NewUpdateContainerInstance() *UpdateContainerInstance`

NewUpdateContainerInstance instantiates a new UpdateContainerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerInstanceWithDefaults

`func NewUpdateContainerInstanceWithDefaults() *UpdateContainerInstance`

NewUpdateContainerInstanceWithDefaults instantiates a new UpdateContainerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateContainerInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateContainerInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateContainerInstance) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateContainerInstance) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *UpdateContainerInstance) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *UpdateContainerInstance) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *UpdateContainerInstance) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *UpdateContainerInstance) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetCustomVariables

`func (o *UpdateContainerInstance) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateContainerInstance) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateContainerInstance) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateContainerInstance) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


