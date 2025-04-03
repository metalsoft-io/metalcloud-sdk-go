# UpdateStorageInternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterName** | Pointer to **string** | The name of the datacenter where the storage is located. | [optional] 
**Status** | Pointer to **string** | Storage status | [optional] 
**TotalCapacity** | Pointer to **float32** | Total capacity in MB | [optional] 
**UsableCapacity** | Pointer to **float32** | Usable capacity in MB | [optional] 
**FreeCapacity** | Pointer to **float32** | Free capacity in MB | [optional] 
**VirtualUsedCapacity** | Pointer to **float32** | Virtual used capacity in MB | [optional] 
**InMaintenance** | Pointer to **float32** | Specifies if the storage is in maintenance | [optional] 
**IsExperimental** | Pointer to **float32** | Specifies if the storage is experimental | [optional] 
**DrivePriority** | Pointer to **float32** | Specifies the drive priority | [optional] 
**SharedDrivePriority** | Pointer to **float32** | Specifies the shared drive priority | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**PortGroupAllocationOrder** | Pointer to **map[string]interface{}** | Port group allocation order | [optional] 
**PortGroupPhysicalPorts** | Pointer to **map[string]interface{}** | Port group physical ports | [optional] 
**DefaultIoLimitPolicy** | Pointer to **string** | Default IO limit policy | [optional] 
**ExtensionInfo** | Pointer to [**ExtensionExecutionInfo**](ExtensionExecutionInfo.md) | The extension execution info of the storage. | [optional] 
**Username** | Pointer to **string** | Username | [optional] 
**Password** | Pointer to **string** | The password to use. | [optional] 
**Options** | Pointer to [**StorageOptions**](StorageOptions.md) | Options for the storage | [optional] 

## Methods

### NewUpdateStorageInternal

`func NewUpdateStorageInternal() *UpdateStorageInternal`

NewUpdateStorageInternal instantiates a new UpdateStorageInternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageInternalWithDefaults

`func NewUpdateStorageInternalWithDefaults() *UpdateStorageInternal`

NewUpdateStorageInternalWithDefaults instantiates a new UpdateStorageInternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterName

`func (o *UpdateStorageInternal) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *UpdateStorageInternal) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *UpdateStorageInternal) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *UpdateStorageInternal) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateStorageInternal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateStorageInternal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateStorageInternal) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateStorageInternal) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *UpdateStorageInternal) GetTotalCapacity() float32`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *UpdateStorageInternal) GetTotalCapacityOk() (*float32, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *UpdateStorageInternal) SetTotalCapacity(v float32)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *UpdateStorageInternal) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetUsableCapacity

`func (o *UpdateStorageInternal) GetUsableCapacity() float32`

GetUsableCapacity returns the UsableCapacity field if non-nil, zero value otherwise.

### GetUsableCapacityOk

`func (o *UpdateStorageInternal) GetUsableCapacityOk() (*float32, bool)`

GetUsableCapacityOk returns a tuple with the UsableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsableCapacity

`func (o *UpdateStorageInternal) SetUsableCapacity(v float32)`

SetUsableCapacity sets UsableCapacity field to given value.

### HasUsableCapacity

`func (o *UpdateStorageInternal) HasUsableCapacity() bool`

HasUsableCapacity returns a boolean if a field has been set.

### GetFreeCapacity

`func (o *UpdateStorageInternal) GetFreeCapacity() float32`

GetFreeCapacity returns the FreeCapacity field if non-nil, zero value otherwise.

### GetFreeCapacityOk

`func (o *UpdateStorageInternal) GetFreeCapacityOk() (*float32, bool)`

GetFreeCapacityOk returns a tuple with the FreeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacity

`func (o *UpdateStorageInternal) SetFreeCapacity(v float32)`

SetFreeCapacity sets FreeCapacity field to given value.

### HasFreeCapacity

`func (o *UpdateStorageInternal) HasFreeCapacity() bool`

HasFreeCapacity returns a boolean if a field has been set.

### GetVirtualUsedCapacity

`func (o *UpdateStorageInternal) GetVirtualUsedCapacity() float32`

GetVirtualUsedCapacity returns the VirtualUsedCapacity field if non-nil, zero value otherwise.

### GetVirtualUsedCapacityOk

`func (o *UpdateStorageInternal) GetVirtualUsedCapacityOk() (*float32, bool)`

GetVirtualUsedCapacityOk returns a tuple with the VirtualUsedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualUsedCapacity

`func (o *UpdateStorageInternal) SetVirtualUsedCapacity(v float32)`

SetVirtualUsedCapacity sets VirtualUsedCapacity field to given value.

### HasVirtualUsedCapacity

`func (o *UpdateStorageInternal) HasVirtualUsedCapacity() bool`

HasVirtualUsedCapacity returns a boolean if a field has been set.

### GetInMaintenance

`func (o *UpdateStorageInternal) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *UpdateStorageInternal) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *UpdateStorageInternal) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *UpdateStorageInternal) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetIsExperimental

`func (o *UpdateStorageInternal) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *UpdateStorageInternal) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *UpdateStorageInternal) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *UpdateStorageInternal) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetDrivePriority

`func (o *UpdateStorageInternal) GetDrivePriority() float32`

GetDrivePriority returns the DrivePriority field if non-nil, zero value otherwise.

### GetDrivePriorityOk

`func (o *UpdateStorageInternal) GetDrivePriorityOk() (*float32, bool)`

GetDrivePriorityOk returns a tuple with the DrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePriority

`func (o *UpdateStorageInternal) SetDrivePriority(v float32)`

SetDrivePriority sets DrivePriority field to given value.

### HasDrivePriority

`func (o *UpdateStorageInternal) HasDrivePriority() bool`

HasDrivePriority returns a boolean if a field has been set.

### GetSharedDrivePriority

`func (o *UpdateStorageInternal) GetSharedDrivePriority() float32`

GetSharedDrivePriority returns the SharedDrivePriority field if non-nil, zero value otherwise.

### GetSharedDrivePriorityOk

`func (o *UpdateStorageInternal) GetSharedDrivePriorityOk() (*float32, bool)`

GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrivePriority

`func (o *UpdateStorageInternal) SetSharedDrivePriority(v float32)`

SetSharedDrivePriority sets SharedDrivePriority field to given value.

### HasSharedDrivePriority

`func (o *UpdateStorageInternal) HasSharedDrivePriority() bool`

HasSharedDrivePriority returns a boolean if a field has been set.

### GetTags

`func (o *UpdateStorageInternal) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateStorageInternal) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateStorageInternal) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateStorageInternal) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPortGroupAllocationOrder

`func (o *UpdateStorageInternal) GetPortGroupAllocationOrder() map[string]interface{}`

GetPortGroupAllocationOrder returns the PortGroupAllocationOrder field if non-nil, zero value otherwise.

### GetPortGroupAllocationOrderOk

`func (o *UpdateStorageInternal) GetPortGroupAllocationOrderOk() (*map[string]interface{}, bool)`

GetPortGroupAllocationOrderOk returns a tuple with the PortGroupAllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupAllocationOrder

`func (o *UpdateStorageInternal) SetPortGroupAllocationOrder(v map[string]interface{})`

SetPortGroupAllocationOrder sets PortGroupAllocationOrder field to given value.

### HasPortGroupAllocationOrder

`func (o *UpdateStorageInternal) HasPortGroupAllocationOrder() bool`

HasPortGroupAllocationOrder returns a boolean if a field has been set.

### GetPortGroupPhysicalPorts

`func (o *UpdateStorageInternal) GetPortGroupPhysicalPorts() map[string]interface{}`

GetPortGroupPhysicalPorts returns the PortGroupPhysicalPorts field if non-nil, zero value otherwise.

### GetPortGroupPhysicalPortsOk

`func (o *UpdateStorageInternal) GetPortGroupPhysicalPortsOk() (*map[string]interface{}, bool)`

GetPortGroupPhysicalPortsOk returns a tuple with the PortGroupPhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupPhysicalPorts

`func (o *UpdateStorageInternal) SetPortGroupPhysicalPorts(v map[string]interface{})`

SetPortGroupPhysicalPorts sets PortGroupPhysicalPorts field to given value.

### HasPortGroupPhysicalPorts

`func (o *UpdateStorageInternal) HasPortGroupPhysicalPorts() bool`

HasPortGroupPhysicalPorts returns a boolean if a field has been set.

### GetDefaultIoLimitPolicy

`func (o *UpdateStorageInternal) GetDefaultIoLimitPolicy() string`

GetDefaultIoLimitPolicy returns the DefaultIoLimitPolicy field if non-nil, zero value otherwise.

### GetDefaultIoLimitPolicyOk

`func (o *UpdateStorageInternal) GetDefaultIoLimitPolicyOk() (*string, bool)`

GetDefaultIoLimitPolicyOk returns a tuple with the DefaultIoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIoLimitPolicy

`func (o *UpdateStorageInternal) SetDefaultIoLimitPolicy(v string)`

SetDefaultIoLimitPolicy sets DefaultIoLimitPolicy field to given value.

### HasDefaultIoLimitPolicy

`func (o *UpdateStorageInternal) HasDefaultIoLimitPolicy() bool`

HasDefaultIoLimitPolicy returns a boolean if a field has been set.

### GetExtensionInfo

`func (o *UpdateStorageInternal) GetExtensionInfo() ExtensionExecutionInfo`

GetExtensionInfo returns the ExtensionInfo field if non-nil, zero value otherwise.

### GetExtensionInfoOk

`func (o *UpdateStorageInternal) GetExtensionInfoOk() (*ExtensionExecutionInfo, bool)`

GetExtensionInfoOk returns a tuple with the ExtensionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInfo

`func (o *UpdateStorageInternal) SetExtensionInfo(v ExtensionExecutionInfo)`

SetExtensionInfo sets ExtensionInfo field to given value.

### HasExtensionInfo

`func (o *UpdateStorageInternal) HasExtensionInfo() bool`

HasExtensionInfo returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateStorageInternal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateStorageInternal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateStorageInternal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateStorageInternal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateStorageInternal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateStorageInternal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateStorageInternal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateStorageInternal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateStorageInternal) GetOptions() StorageOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateStorageInternal) GetOptionsOk() (*StorageOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateStorageInternal) SetOptions(v StorageOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateStorageInternal) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


