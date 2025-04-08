# Storage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Id of the Storage | 
**Revision** | **float32** | Revision of the Storage | 
**UserId** | Pointer to **float32** | Id of the owner | [optional] 
**SiteId** | **float32** | Id of the site | 
**DatacenterName** | **string** | The name of the datacenter where the storage is located. | 
**Driver** | **string** | Storage driver | 
**Technology** | **string** | Storage technology | 
**Type** | **string** | Storage type | 
**Status** | **string** | Storage status | 
**TotalCapacity** | Pointer to **float32** | Total capacity in MB | [optional] 
**UsableCapacity** | Pointer to **float32** | Usable capacity in MB | [optional] 
**FreeCapacity** | Pointer to **float32** | Free capacity in MB | [optional] 
**VirtualUsedCapacity** | Pointer to **float32** | Virtual used capacity in MB | [optional] 
**Name** | **string** | Name of the storage | 
**IscsiHost** | Pointer to **string** | ISCSI host | [optional] 
**IscsiPort** | Pointer to **float32** | ISCSI port | [optional] 
**ManagementHost** | **string** | Management host | 
**Username** | **string** | Username | 
**PasswordEncrypted** | Pointer to **string** | Password encrypted | [optional] 
**Options** | Pointer to [**StorageOptions**](StorageOptions.md) | Options for the storage | [optional] 
**InMaintenance** | Pointer to **float32** | Specifies if the storage is in maintenance | [optional] 
**TargetIQN** | Pointer to **string** | Target IQN | [optional] 
**IsExperimental** | Pointer to **float32** | Specifies if the storage is experimental | [optional] 
**DrivePriority** | Pointer to **float32** | Specifies the drive priority | [optional] 
**SharedDrivePriority** | Pointer to **float32** | Specifies the shared drive priority | [optional] 
**AlternateSanIPs** | Pointer to **[]string** | Alternate SAN IPs | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**PortGroupAllocationOrder** | Pointer to **map[string]interface{}** | Port group allocation order | [optional] 
**PortGroupPhysicalPorts** | Pointer to **map[string]interface{}** | Port group physical ports | [optional] 
**DefaultIoLimitPolicy** | Pointer to **string** | Default IO limit policy | [optional] 
**SubnetType** | **string** | Subnet type | 
**JobStatistics** | Pointer to [**JobGroupStatistics**](JobGroupStatistics.md) |  | [optional] 
**ExtensionInfo** | Pointer to [**ExtensionExecutionInfo**](ExtensionExecutionInfo.md) | The extension execution info of the storage. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewStorage

`func NewStorage(id float32, revision float32, siteId float32, datacenterName string, driver string, technology string, type_ string, status string, name string, managementHost string, username string, subnetType string, ) *Storage`

NewStorage instantiates a new Storage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageWithDefaults

`func NewStorageWithDefaults() *Storage`

NewStorageWithDefaults instantiates a new Storage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Storage) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Storage) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Storage) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *Storage) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Storage) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Storage) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetUserId

`func (o *Storage) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Storage) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Storage) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Storage) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *Storage) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Storage) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Storage) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetDatacenterName

`func (o *Storage) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *Storage) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *Storage) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetDriver

`func (o *Storage) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *Storage) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *Storage) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetTechnology

`func (o *Storage) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *Storage) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *Storage) SetTechnology(v string)`

SetTechnology sets Technology field to given value.


### GetType

`func (o *Storage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Storage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Storage) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Storage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Storage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Storage) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalCapacity

`func (o *Storage) GetTotalCapacity() float32`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *Storage) GetTotalCapacityOk() (*float32, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *Storage) SetTotalCapacity(v float32)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *Storage) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetUsableCapacity

`func (o *Storage) GetUsableCapacity() float32`

GetUsableCapacity returns the UsableCapacity field if non-nil, zero value otherwise.

### GetUsableCapacityOk

`func (o *Storage) GetUsableCapacityOk() (*float32, bool)`

GetUsableCapacityOk returns a tuple with the UsableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsableCapacity

`func (o *Storage) SetUsableCapacity(v float32)`

SetUsableCapacity sets UsableCapacity field to given value.

### HasUsableCapacity

`func (o *Storage) HasUsableCapacity() bool`

HasUsableCapacity returns a boolean if a field has been set.

### GetFreeCapacity

`func (o *Storage) GetFreeCapacity() float32`

GetFreeCapacity returns the FreeCapacity field if non-nil, zero value otherwise.

### GetFreeCapacityOk

`func (o *Storage) GetFreeCapacityOk() (*float32, bool)`

GetFreeCapacityOk returns a tuple with the FreeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacity

`func (o *Storage) SetFreeCapacity(v float32)`

SetFreeCapacity sets FreeCapacity field to given value.

### HasFreeCapacity

`func (o *Storage) HasFreeCapacity() bool`

HasFreeCapacity returns a boolean if a field has been set.

### GetVirtualUsedCapacity

`func (o *Storage) GetVirtualUsedCapacity() float32`

GetVirtualUsedCapacity returns the VirtualUsedCapacity field if non-nil, zero value otherwise.

### GetVirtualUsedCapacityOk

`func (o *Storage) GetVirtualUsedCapacityOk() (*float32, bool)`

GetVirtualUsedCapacityOk returns a tuple with the VirtualUsedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualUsedCapacity

`func (o *Storage) SetVirtualUsedCapacity(v float32)`

SetVirtualUsedCapacity sets VirtualUsedCapacity field to given value.

### HasVirtualUsedCapacity

`func (o *Storage) HasVirtualUsedCapacity() bool`

HasVirtualUsedCapacity returns a boolean if a field has been set.

### GetName

`func (o *Storage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Storage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Storage) SetName(v string)`

SetName sets Name field to given value.


### GetIscsiHost

`func (o *Storage) GetIscsiHost() string`

GetIscsiHost returns the IscsiHost field if non-nil, zero value otherwise.

### GetIscsiHostOk

`func (o *Storage) GetIscsiHostOk() (*string, bool)`

GetIscsiHostOk returns a tuple with the IscsiHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiHost

`func (o *Storage) SetIscsiHost(v string)`

SetIscsiHost sets IscsiHost field to given value.

### HasIscsiHost

`func (o *Storage) HasIscsiHost() bool`

HasIscsiHost returns a boolean if a field has been set.

### GetIscsiPort

`func (o *Storage) GetIscsiPort() float32`

GetIscsiPort returns the IscsiPort field if non-nil, zero value otherwise.

### GetIscsiPortOk

`func (o *Storage) GetIscsiPortOk() (*float32, bool)`

GetIscsiPortOk returns a tuple with the IscsiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiPort

`func (o *Storage) SetIscsiPort(v float32)`

SetIscsiPort sets IscsiPort field to given value.

### HasIscsiPort

`func (o *Storage) HasIscsiPort() bool`

HasIscsiPort returns a boolean if a field has been set.

### GetManagementHost

`func (o *Storage) GetManagementHost() string`

GetManagementHost returns the ManagementHost field if non-nil, zero value otherwise.

### GetManagementHostOk

`func (o *Storage) GetManagementHostOk() (*string, bool)`

GetManagementHostOk returns a tuple with the ManagementHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementHost

`func (o *Storage) SetManagementHost(v string)`

SetManagementHost sets ManagementHost field to given value.


### GetUsername

`func (o *Storage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Storage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Storage) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPasswordEncrypted

`func (o *Storage) GetPasswordEncrypted() string`

GetPasswordEncrypted returns the PasswordEncrypted field if non-nil, zero value otherwise.

### GetPasswordEncryptedOk

`func (o *Storage) GetPasswordEncryptedOk() (*string, bool)`

GetPasswordEncryptedOk returns a tuple with the PasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncrypted

`func (o *Storage) SetPasswordEncrypted(v string)`

SetPasswordEncrypted sets PasswordEncrypted field to given value.

### HasPasswordEncrypted

`func (o *Storage) HasPasswordEncrypted() bool`

HasPasswordEncrypted returns a boolean if a field has been set.

### GetOptions

`func (o *Storage) GetOptions() StorageOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Storage) GetOptionsOk() (*StorageOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Storage) SetOptions(v StorageOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Storage) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetInMaintenance

`func (o *Storage) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *Storage) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *Storage) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *Storage) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetTargetIQN

`func (o *Storage) GetTargetIQN() string`

GetTargetIQN returns the TargetIQN field if non-nil, zero value otherwise.

### GetTargetIQNOk

`func (o *Storage) GetTargetIQNOk() (*string, bool)`

GetTargetIQNOk returns a tuple with the TargetIQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIQN

`func (o *Storage) SetTargetIQN(v string)`

SetTargetIQN sets TargetIQN field to given value.

### HasTargetIQN

`func (o *Storage) HasTargetIQN() bool`

HasTargetIQN returns a boolean if a field has been set.

### GetIsExperimental

`func (o *Storage) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *Storage) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *Storage) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *Storage) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetDrivePriority

`func (o *Storage) GetDrivePriority() float32`

GetDrivePriority returns the DrivePriority field if non-nil, zero value otherwise.

### GetDrivePriorityOk

`func (o *Storage) GetDrivePriorityOk() (*float32, bool)`

GetDrivePriorityOk returns a tuple with the DrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePriority

`func (o *Storage) SetDrivePriority(v float32)`

SetDrivePriority sets DrivePriority field to given value.

### HasDrivePriority

`func (o *Storage) HasDrivePriority() bool`

HasDrivePriority returns a boolean if a field has been set.

### GetSharedDrivePriority

`func (o *Storage) GetSharedDrivePriority() float32`

GetSharedDrivePriority returns the SharedDrivePriority field if non-nil, zero value otherwise.

### GetSharedDrivePriorityOk

`func (o *Storage) GetSharedDrivePriorityOk() (*float32, bool)`

GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrivePriority

`func (o *Storage) SetSharedDrivePriority(v float32)`

SetSharedDrivePriority sets SharedDrivePriority field to given value.

### HasSharedDrivePriority

`func (o *Storage) HasSharedDrivePriority() bool`

HasSharedDrivePriority returns a boolean if a field has been set.

### GetAlternateSanIPs

`func (o *Storage) GetAlternateSanIPs() []string`

GetAlternateSanIPs returns the AlternateSanIPs field if non-nil, zero value otherwise.

### GetAlternateSanIPsOk

`func (o *Storage) GetAlternateSanIPsOk() (*[]string, bool)`

GetAlternateSanIPsOk returns a tuple with the AlternateSanIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateSanIPs

`func (o *Storage) SetAlternateSanIPs(v []string)`

SetAlternateSanIPs sets AlternateSanIPs field to given value.

### HasAlternateSanIPs

`func (o *Storage) HasAlternateSanIPs() bool`

HasAlternateSanIPs returns a boolean if a field has been set.

### GetTags

`func (o *Storage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Storage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Storage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Storage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPortGroupAllocationOrder

`func (o *Storage) GetPortGroupAllocationOrder() map[string]interface{}`

GetPortGroupAllocationOrder returns the PortGroupAllocationOrder field if non-nil, zero value otherwise.

### GetPortGroupAllocationOrderOk

`func (o *Storage) GetPortGroupAllocationOrderOk() (*map[string]interface{}, bool)`

GetPortGroupAllocationOrderOk returns a tuple with the PortGroupAllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupAllocationOrder

`func (o *Storage) SetPortGroupAllocationOrder(v map[string]interface{})`

SetPortGroupAllocationOrder sets PortGroupAllocationOrder field to given value.

### HasPortGroupAllocationOrder

`func (o *Storage) HasPortGroupAllocationOrder() bool`

HasPortGroupAllocationOrder returns a boolean if a field has been set.

### GetPortGroupPhysicalPorts

`func (o *Storage) GetPortGroupPhysicalPorts() map[string]interface{}`

GetPortGroupPhysicalPorts returns the PortGroupPhysicalPorts field if non-nil, zero value otherwise.

### GetPortGroupPhysicalPortsOk

`func (o *Storage) GetPortGroupPhysicalPortsOk() (*map[string]interface{}, bool)`

GetPortGroupPhysicalPortsOk returns a tuple with the PortGroupPhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupPhysicalPorts

`func (o *Storage) SetPortGroupPhysicalPorts(v map[string]interface{})`

SetPortGroupPhysicalPorts sets PortGroupPhysicalPorts field to given value.

### HasPortGroupPhysicalPorts

`func (o *Storage) HasPortGroupPhysicalPorts() bool`

HasPortGroupPhysicalPorts returns a boolean if a field has been set.

### GetDefaultIoLimitPolicy

`func (o *Storage) GetDefaultIoLimitPolicy() string`

GetDefaultIoLimitPolicy returns the DefaultIoLimitPolicy field if non-nil, zero value otherwise.

### GetDefaultIoLimitPolicyOk

`func (o *Storage) GetDefaultIoLimitPolicyOk() (*string, bool)`

GetDefaultIoLimitPolicyOk returns a tuple with the DefaultIoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIoLimitPolicy

`func (o *Storage) SetDefaultIoLimitPolicy(v string)`

SetDefaultIoLimitPolicy sets DefaultIoLimitPolicy field to given value.

### HasDefaultIoLimitPolicy

`func (o *Storage) HasDefaultIoLimitPolicy() bool`

HasDefaultIoLimitPolicy returns a boolean if a field has been set.

### GetSubnetType

`func (o *Storage) GetSubnetType() string`

GetSubnetType returns the SubnetType field if non-nil, zero value otherwise.

### GetSubnetTypeOk

`func (o *Storage) GetSubnetTypeOk() (*string, bool)`

GetSubnetTypeOk returns a tuple with the SubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetType

`func (o *Storage) SetSubnetType(v string)`

SetSubnetType sets SubnetType field to given value.


### GetJobStatistics

`func (o *Storage) GetJobStatistics() JobGroupStatistics`

GetJobStatistics returns the JobStatistics field if non-nil, zero value otherwise.

### GetJobStatisticsOk

`func (o *Storage) GetJobStatisticsOk() (*JobGroupStatistics, bool)`

GetJobStatisticsOk returns a tuple with the JobStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatistics

`func (o *Storage) SetJobStatistics(v JobGroupStatistics)`

SetJobStatistics sets JobStatistics field to given value.

### HasJobStatistics

`func (o *Storage) HasJobStatistics() bool`

HasJobStatistics returns a boolean if a field has been set.

### GetExtensionInfo

`func (o *Storage) GetExtensionInfo() ExtensionExecutionInfo`

GetExtensionInfo returns the ExtensionInfo field if non-nil, zero value otherwise.

### GetExtensionInfoOk

`func (o *Storage) GetExtensionInfoOk() (*ExtensionExecutionInfo, bool)`

GetExtensionInfoOk returns a tuple with the ExtensionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInfo

`func (o *Storage) SetExtensionInfo(v ExtensionExecutionInfo)`

SetExtensionInfo sets ExtensionInfo field to given value.

### HasExtensionInfo

`func (o *Storage) HasExtensionInfo() bool`

HasExtensionInfo returns a boolean if a field has been set.

### GetLinks

`func (o *Storage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Storage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Storage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Storage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


