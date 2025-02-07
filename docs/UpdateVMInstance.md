# UpdateVMInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | Tags for the VM Instance. | [optional] 
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

### GetTags

`func (o *UpdateVMInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMInstance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

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


