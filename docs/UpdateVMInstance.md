# UpdateVMInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label for the VM Instance Group. | [optional] 
**DiskSizeGB** | Pointer to **float32** | Disk size in GB for the VM Instance. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance. | [optional] 

## Methods

### NewUpdateVMInstance

`func NewUpdateVMInstance() *UpdateVMInstance`

NewUpdateVMInstance instantiates a new UpdateVMInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMInstanceWithDefaults

`func NewUpdateVMInstanceWithDefaults() *UpdateVMInstance`

NewUpdateVMInstanceWithDefaults instantiates a new UpdateVMInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateVMInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateVMInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateVMInstance) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateVMInstance) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *UpdateVMInstance) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *UpdateVMInstance) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *UpdateVMInstance) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *UpdateVMInstance) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetCustomVariables

`func (o *UpdateVMInstance) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateVMInstance) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateVMInstance) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateVMInstance) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


