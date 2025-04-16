# CreateStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **float32** | Id of the owner | [optional] 
**SiteId** | **float32** | Id of the site | 
**Driver** | **string** | Storage driver | 
**Technologies** | **[]string** | Storage technology | 
**Type** | **string** | Storage type | 
**Name** | **string** | Name of the storage | 
**IscsiHost** | Pointer to **string** | ISCSI host | [optional] 
**IscsiPort** | Pointer to **float32** | ISCSI port | [optional] 
**ManagementHost** | **string** | Management host | 
**Username** | **string** | Username | 
**InMaintenance** | Pointer to **float32** | Specifies if the storage is in maintenance | [optional] 
**TargetIQN** | Pointer to **string** | Target IQN | [optional] 
**IsExperimental** | Pointer to **float32** | Specifies if the storage is experimental | [optional] 
**DrivePriority** | Pointer to **float32** | Specifies the drive priority | [optional] 
**SharedDrivePriority** | Pointer to **float32** | Specifies the shared drive priority | [optional] 
**AlternateSanIPs** | Pointer to **[]string** | Alternate SAN IPs | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**PortGroupAllocationOrder** | Pointer to **map[string]interface{}** | Port group allocation order | [optional] 
**PortGroupPhysicalPorts** | Pointer to **map[string]interface{}** | Port group physical ports | [optional] 
**SubnetType** | **string** | Subnet type | 
**Options** | Pointer to [**UpdateStorageOptions**](UpdateStorageOptions.md) | Options for the storage | [optional] 
**Password** | **string** | The password to use. | 

## Methods

### NewCreateStorage

`func NewCreateStorage(siteId float32, driver string, technologies []string, type_ string, name string, managementHost string, username string, subnetType string, password string, ) *CreateStorage`

NewCreateStorage instantiates a new CreateStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageWithDefaults

`func NewCreateStorageWithDefaults() *CreateStorage`

NewCreateStorageWithDefaults instantiates a new CreateStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateStorage) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateStorage) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateStorage) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateStorage) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *CreateStorage) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateStorage) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateStorage) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetDriver

`func (o *CreateStorage) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *CreateStorage) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *CreateStorage) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetTechnologies

`func (o *CreateStorage) GetTechnologies() []string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *CreateStorage) GetTechnologiesOk() (*[]string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *CreateStorage) SetTechnologies(v []string)`

SetTechnologies sets Technologies field to given value.


### GetType

`func (o *CreateStorage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStorage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStorage) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CreateStorage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStorage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStorage) SetName(v string)`

SetName sets Name field to given value.


### GetIscsiHost

`func (o *CreateStorage) GetIscsiHost() string`

GetIscsiHost returns the IscsiHost field if non-nil, zero value otherwise.

### GetIscsiHostOk

`func (o *CreateStorage) GetIscsiHostOk() (*string, bool)`

GetIscsiHostOk returns a tuple with the IscsiHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiHost

`func (o *CreateStorage) SetIscsiHost(v string)`

SetIscsiHost sets IscsiHost field to given value.

### HasIscsiHost

`func (o *CreateStorage) HasIscsiHost() bool`

HasIscsiHost returns a boolean if a field has been set.

### GetIscsiPort

`func (o *CreateStorage) GetIscsiPort() float32`

GetIscsiPort returns the IscsiPort field if non-nil, zero value otherwise.

### GetIscsiPortOk

`func (o *CreateStorage) GetIscsiPortOk() (*float32, bool)`

GetIscsiPortOk returns a tuple with the IscsiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiPort

`func (o *CreateStorage) SetIscsiPort(v float32)`

SetIscsiPort sets IscsiPort field to given value.

### HasIscsiPort

`func (o *CreateStorage) HasIscsiPort() bool`

HasIscsiPort returns a boolean if a field has been set.

### GetManagementHost

`func (o *CreateStorage) GetManagementHost() string`

GetManagementHost returns the ManagementHost field if non-nil, zero value otherwise.

### GetManagementHostOk

`func (o *CreateStorage) GetManagementHostOk() (*string, bool)`

GetManagementHostOk returns a tuple with the ManagementHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementHost

`func (o *CreateStorage) SetManagementHost(v string)`

SetManagementHost sets ManagementHost field to given value.


### GetUsername

`func (o *CreateStorage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateStorage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateStorage) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetInMaintenance

`func (o *CreateStorage) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *CreateStorage) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *CreateStorage) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *CreateStorage) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetTargetIQN

`func (o *CreateStorage) GetTargetIQN() string`

GetTargetIQN returns the TargetIQN field if non-nil, zero value otherwise.

### GetTargetIQNOk

`func (o *CreateStorage) GetTargetIQNOk() (*string, bool)`

GetTargetIQNOk returns a tuple with the TargetIQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIQN

`func (o *CreateStorage) SetTargetIQN(v string)`

SetTargetIQN sets TargetIQN field to given value.

### HasTargetIQN

`func (o *CreateStorage) HasTargetIQN() bool`

HasTargetIQN returns a boolean if a field has been set.

### GetIsExperimental

`func (o *CreateStorage) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *CreateStorage) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *CreateStorage) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *CreateStorage) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetDrivePriority

`func (o *CreateStorage) GetDrivePriority() float32`

GetDrivePriority returns the DrivePriority field if non-nil, zero value otherwise.

### GetDrivePriorityOk

`func (o *CreateStorage) GetDrivePriorityOk() (*float32, bool)`

GetDrivePriorityOk returns a tuple with the DrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePriority

`func (o *CreateStorage) SetDrivePriority(v float32)`

SetDrivePriority sets DrivePriority field to given value.

### HasDrivePriority

`func (o *CreateStorage) HasDrivePriority() bool`

HasDrivePriority returns a boolean if a field has been set.

### GetSharedDrivePriority

`func (o *CreateStorage) GetSharedDrivePriority() float32`

GetSharedDrivePriority returns the SharedDrivePriority field if non-nil, zero value otherwise.

### GetSharedDrivePriorityOk

`func (o *CreateStorage) GetSharedDrivePriorityOk() (*float32, bool)`

GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrivePriority

`func (o *CreateStorage) SetSharedDrivePriority(v float32)`

SetSharedDrivePriority sets SharedDrivePriority field to given value.

### HasSharedDrivePriority

`func (o *CreateStorage) HasSharedDrivePriority() bool`

HasSharedDrivePriority returns a boolean if a field has been set.

### GetAlternateSanIPs

`func (o *CreateStorage) GetAlternateSanIPs() []string`

GetAlternateSanIPs returns the AlternateSanIPs field if non-nil, zero value otherwise.

### GetAlternateSanIPsOk

`func (o *CreateStorage) GetAlternateSanIPsOk() (*[]string, bool)`

GetAlternateSanIPsOk returns a tuple with the AlternateSanIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateSanIPs

`func (o *CreateStorage) SetAlternateSanIPs(v []string)`

SetAlternateSanIPs sets AlternateSanIPs field to given value.

### HasAlternateSanIPs

`func (o *CreateStorage) HasAlternateSanIPs() bool`

HasAlternateSanIPs returns a boolean if a field has been set.

### GetTags

`func (o *CreateStorage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateStorage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateStorage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateStorage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPortGroupAllocationOrder

`func (o *CreateStorage) GetPortGroupAllocationOrder() map[string]interface{}`

GetPortGroupAllocationOrder returns the PortGroupAllocationOrder field if non-nil, zero value otherwise.

### GetPortGroupAllocationOrderOk

`func (o *CreateStorage) GetPortGroupAllocationOrderOk() (*map[string]interface{}, bool)`

GetPortGroupAllocationOrderOk returns a tuple with the PortGroupAllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupAllocationOrder

`func (o *CreateStorage) SetPortGroupAllocationOrder(v map[string]interface{})`

SetPortGroupAllocationOrder sets PortGroupAllocationOrder field to given value.

### HasPortGroupAllocationOrder

`func (o *CreateStorage) HasPortGroupAllocationOrder() bool`

HasPortGroupAllocationOrder returns a boolean if a field has been set.

### GetPortGroupPhysicalPorts

`func (o *CreateStorage) GetPortGroupPhysicalPorts() map[string]interface{}`

GetPortGroupPhysicalPorts returns the PortGroupPhysicalPorts field if non-nil, zero value otherwise.

### GetPortGroupPhysicalPortsOk

`func (o *CreateStorage) GetPortGroupPhysicalPortsOk() (*map[string]interface{}, bool)`

GetPortGroupPhysicalPortsOk returns a tuple with the PortGroupPhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupPhysicalPorts

`func (o *CreateStorage) SetPortGroupPhysicalPorts(v map[string]interface{})`

SetPortGroupPhysicalPorts sets PortGroupPhysicalPorts field to given value.

### HasPortGroupPhysicalPorts

`func (o *CreateStorage) HasPortGroupPhysicalPorts() bool`

HasPortGroupPhysicalPorts returns a boolean if a field has been set.

### GetSubnetType

`func (o *CreateStorage) GetSubnetType() string`

GetSubnetType returns the SubnetType field if non-nil, zero value otherwise.

### GetSubnetTypeOk

`func (o *CreateStorage) GetSubnetTypeOk() (*string, bool)`

GetSubnetTypeOk returns a tuple with the SubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetType

`func (o *CreateStorage) SetSubnetType(v string)`

SetSubnetType sets SubnetType field to given value.


### GetOptions

`func (o *CreateStorage) GetOptions() UpdateStorageOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateStorage) GetOptionsOk() (*UpdateStorageOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateStorage) SetOptions(v UpdateStorageOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateStorage) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPassword

`func (o *CreateStorage) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateStorage) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateStorage) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


