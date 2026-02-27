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
**Technologies** | **[]string** | Storage technology | 
**Type** | Pointer to **string** | Storage type | [optional] 
**Status** | **string** | Storage status | 
**OperationMode** | **string** | Storage operation mode. If fc_only is set, only fibre channel deployments are allowed. | 
**Name** | **string** | Name of the storage | 
**IscsiHost** | Pointer to **string** | ISCSI host | [optional] 
**IscsiPort** | Pointer to **float32** | ISCSI port | [optional] 
**ManagementHost** | **string** | Management host | 
**Username** | Pointer to **string** | The username to use. | [optional] 
**Options** | Pointer to [**StorageOptions**](StorageOptions.md) | Options for the storage | [optional] 
**InMaintenance** | Pointer to **float32** | Specifies if the storage is in maintenance | [optional] 
**TargetIQN** | Pointer to **string** | Target IQN | [optional] 
**IsExperimental** | Pointer to **float32** | Specifies if the storage is experimental | [optional] 
**DrivePriority** | Pointer to **float32** | Specifies the drive priority | [optional] 
**SharedDrivePriority** | Pointer to **float32** | Specifies the shared drive priority | [optional] 
**AlternateSanIPs** | Pointer to **[]string** | Alternate SAN IPs | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**SubnetType** | **string** | Subnet type | 
**Interfaces** | Pointer to [**[]StorageInterface**](StorageInterface.md) | Interfaces of the Storage | [optional] 
**NetworkFabricId** | Pointer to **float32** | Network fabric ID this Storage is connected to | [optional] 
**JobStatistics** | Pointer to [**JobGroupStatistics**](JobGroupStatistics.md) |  | [optional] 
**ExtensionInfo** | Pointer to [**ExtensionExecutionInfo**](ExtensionExecutionInfo.md) | The extension execution info of the storage. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewStorage

`func NewStorage(id float32, revision float32, siteId float32, datacenterName string, driver string, technologies []string, status string, operationMode string, name string, managementHost string, subnetType string, ) *Storage`

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


### GetTechnologies

`func (o *Storage) GetTechnologies() []string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *Storage) GetTechnologiesOk() (*[]string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *Storage) SetTechnologies(v []string)`

SetTechnologies sets Technologies field to given value.


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

### HasType

`func (o *Storage) HasType() bool`

HasType returns a boolean if a field has been set.

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


### GetOperationMode

`func (o *Storage) GetOperationMode() string`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *Storage) GetOperationModeOk() (*string, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *Storage) SetOperationMode(v string)`

SetOperationMode sets OperationMode field to given value.


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

### HasUsername

`func (o *Storage) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

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


### GetInterfaces

`func (o *Storage) GetInterfaces() []StorageInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Storage) GetInterfacesOk() (*[]StorageInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Storage) SetInterfaces(v []StorageInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Storage) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetNetworkFabricId

`func (o *Storage) GetNetworkFabricId() float32`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *Storage) GetNetworkFabricIdOk() (*float32, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *Storage) SetNetworkFabricId(v float32)`

SetNetworkFabricId sets NetworkFabricId field to given value.

### HasNetworkFabricId

`func (o *Storage) HasNetworkFabricId() bool`

HasNetworkFabricId returns a boolean if a field has been set.

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


