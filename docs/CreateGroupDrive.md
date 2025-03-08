# CreateGroupDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveCount** | **float32** | Number of drives in the Drive Group | 
**DriveSizeMbDefault** | **float32** | Default disk size in MB for new Drives in the Drive Group | 
**ExtensionInstanceId** | Pointer to **float32** |  | [optional] 
**Label** | Pointer to **string** | Label of the Drive. | [optional] 
**IoLimitPolicy** | Pointer to **string** | The IO limit policy of the Drive Group. | [optional] 
**ExpandWithServerInstanceGroup** | Pointer to **float32** | Flag to determine whether the Drive Group should be expanded with a Server Instance Group by adding one drive for each instance | [optional] 
**AllocationAffinity** | Pointer to **string** | Allocation affinity of the Drive Group | [optional] 

## Methods

### NewCreateGroupDrive

`func NewCreateGroupDrive(driveCount float32, driveSizeMbDefault float32, ) *CreateGroupDrive`

NewCreateGroupDrive instantiates a new CreateGroupDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupDriveWithDefaults

`func NewCreateGroupDriveWithDefaults() *CreateGroupDrive`

NewCreateGroupDriveWithDefaults instantiates a new CreateGroupDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveCount

`func (o *CreateGroupDrive) GetDriveCount() float32`

GetDriveCount returns the DriveCount field if non-nil, zero value otherwise.

### GetDriveCountOk

`func (o *CreateGroupDrive) GetDriveCountOk() (*float32, bool)`

GetDriveCountOk returns a tuple with the DriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveCount

`func (o *CreateGroupDrive) SetDriveCount(v float32)`

SetDriveCount sets DriveCount field to given value.


### GetDriveSizeMbDefault

`func (o *CreateGroupDrive) GetDriveSizeMbDefault() float32`

GetDriveSizeMbDefault returns the DriveSizeMbDefault field if non-nil, zero value otherwise.

### GetDriveSizeMbDefaultOk

`func (o *CreateGroupDrive) GetDriveSizeMbDefaultOk() (*float32, bool)`

GetDriveSizeMbDefaultOk returns a tuple with the DriveSizeMbDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSizeMbDefault

`func (o *CreateGroupDrive) SetDriveSizeMbDefault(v float32)`

SetDriveSizeMbDefault sets DriveSizeMbDefault field to given value.


### GetExtensionInstanceId

`func (o *CreateGroupDrive) GetExtensionInstanceId() float32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *CreateGroupDrive) GetExtensionInstanceIdOk() (*float32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *CreateGroupDrive) SetExtensionInstanceId(v float32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *CreateGroupDrive) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetLabel

`func (o *CreateGroupDrive) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateGroupDrive) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateGroupDrive) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateGroupDrive) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIoLimitPolicy

`func (o *CreateGroupDrive) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *CreateGroupDrive) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *CreateGroupDrive) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *CreateGroupDrive) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetExpandWithServerInstanceGroup

`func (o *CreateGroupDrive) GetExpandWithServerInstanceGroup() float32`

GetExpandWithServerInstanceGroup returns the ExpandWithServerInstanceGroup field if non-nil, zero value otherwise.

### GetExpandWithServerInstanceGroupOk

`func (o *CreateGroupDrive) GetExpandWithServerInstanceGroupOk() (*float32, bool)`

GetExpandWithServerInstanceGroupOk returns a tuple with the ExpandWithServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandWithServerInstanceGroup

`func (o *CreateGroupDrive) SetExpandWithServerInstanceGroup(v float32)`

SetExpandWithServerInstanceGroup sets ExpandWithServerInstanceGroup field to given value.

### HasExpandWithServerInstanceGroup

`func (o *CreateGroupDrive) HasExpandWithServerInstanceGroup() bool`

HasExpandWithServerInstanceGroup returns a boolean if a field has been set.

### GetAllocationAffinity

`func (o *CreateGroupDrive) GetAllocationAffinity() string`

GetAllocationAffinity returns the AllocationAffinity field if non-nil, zero value otherwise.

### GetAllocationAffinityOk

`func (o *CreateGroupDrive) GetAllocationAffinityOk() (*string, bool)`

GetAllocationAffinityOk returns a tuple with the AllocationAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationAffinity

`func (o *CreateGroupDrive) SetAllocationAffinity(v string)`

SetAllocationAffinity sets AllocationAffinity field to given value.

### HasAllocationAffinity

`func (o *CreateGroupDrive) HasAllocationAffinity() bool`

HasAllocationAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


