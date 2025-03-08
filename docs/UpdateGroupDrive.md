# UpdateGroupDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label of the Drive. | [optional] 
**DriveSizeMbDefault** | Pointer to **float32** | Default disk size in MB for new Drives in the Drive Group | [optional] 
**IoLimitPolicy** | Pointer to **string** | The IO limit policy of the Drive Group. | [optional] 
**ExpandWithServerInstanceGroup** | Pointer to **float32** | Flag to determine whether the Drive Group should be expanded with a Server Instance Group by adding one drive for each instance | [optional] 
**AllocationAffinity** | Pointer to **string** | Allocation affinity of the Drive Group | [optional] 

## Methods

### NewUpdateGroupDrive

`func NewUpdateGroupDrive() *UpdateGroupDrive`

NewUpdateGroupDrive instantiates a new UpdateGroupDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupDriveWithDefaults

`func NewUpdateGroupDriveWithDefaults() *UpdateGroupDrive`

NewUpdateGroupDriveWithDefaults instantiates a new UpdateGroupDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateGroupDrive) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateGroupDrive) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateGroupDrive) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateGroupDrive) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDriveSizeMbDefault

`func (o *UpdateGroupDrive) GetDriveSizeMbDefault() float32`

GetDriveSizeMbDefault returns the DriveSizeMbDefault field if non-nil, zero value otherwise.

### GetDriveSizeMbDefaultOk

`func (o *UpdateGroupDrive) GetDriveSizeMbDefaultOk() (*float32, bool)`

GetDriveSizeMbDefaultOk returns a tuple with the DriveSizeMbDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSizeMbDefault

`func (o *UpdateGroupDrive) SetDriveSizeMbDefault(v float32)`

SetDriveSizeMbDefault sets DriveSizeMbDefault field to given value.

### HasDriveSizeMbDefault

`func (o *UpdateGroupDrive) HasDriveSizeMbDefault() bool`

HasDriveSizeMbDefault returns a boolean if a field has been set.

### GetIoLimitPolicy

`func (o *UpdateGroupDrive) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *UpdateGroupDrive) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *UpdateGroupDrive) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *UpdateGroupDrive) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetExpandWithServerInstanceGroup

`func (o *UpdateGroupDrive) GetExpandWithServerInstanceGroup() float32`

GetExpandWithServerInstanceGroup returns the ExpandWithServerInstanceGroup field if non-nil, zero value otherwise.

### GetExpandWithServerInstanceGroupOk

`func (o *UpdateGroupDrive) GetExpandWithServerInstanceGroupOk() (*float32, bool)`

GetExpandWithServerInstanceGroupOk returns a tuple with the ExpandWithServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandWithServerInstanceGroup

`func (o *UpdateGroupDrive) SetExpandWithServerInstanceGroup(v float32)`

SetExpandWithServerInstanceGroup sets ExpandWithServerInstanceGroup field to given value.

### HasExpandWithServerInstanceGroup

`func (o *UpdateGroupDrive) HasExpandWithServerInstanceGroup() bool`

HasExpandWithServerInstanceGroup returns a boolean if a field has been set.

### GetAllocationAffinity

`func (o *UpdateGroupDrive) GetAllocationAffinity() string`

GetAllocationAffinity returns the AllocationAffinity field if non-nil, zero value otherwise.

### GetAllocationAffinityOk

`func (o *UpdateGroupDrive) GetAllocationAffinityOk() (*string, bool)`

GetAllocationAffinityOk returns a tuple with the AllocationAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationAffinity

`func (o *UpdateGroupDrive) SetAllocationAffinity(v string)`

SetAllocationAffinity sets AllocationAffinity field to given value.

### HasAllocationAffinity

`func (o *UpdateGroupDrive) HasAllocationAffinity() bool`

HasAllocationAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


