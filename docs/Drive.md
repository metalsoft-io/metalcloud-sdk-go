# Drive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the Drive. | 
**GroupId** | **float32** | Drive Array Id | 
**InstanceId** | Pointer to **float32** |  | [optional] 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the Drive is assigned to | [optional] 
**SizeMb** | **float32** | Disk size in MB for Drive | 
**StorageImageName** | Pointer to **string** | The name of the storage image used by the Drive. | [optional] 
**IscsiIndexHex** | Pointer to **string** | The iSCSI Index in hex format of the Drive. | [optional] 
**TemplateId** | Pointer to **float32** | Template Id | [optional] 
**OsAdminUsername** | Pointer to **string** | The OS Admin Username the Drive will use. | [optional] 
**OsAdminPasswordEncrypted** | Pointer to **string** | The OS Admin Password the Drive will use. | [optional] 
**StorageType** | **string** | Storage type of the Drive | [default to "iscsi_ssd"]
**Subdomain** | Pointer to **string** | Subdomain of the Drive. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Drive last update. | 
**SshPort** | Pointer to **float32** | SSH port used by the Drive. | [optional] 
**OperatingSystemInfo** | Pointer to **map[string]interface{}** | Operating system information of the Drive. | [optional] 
**FilesystemInfo** | Pointer to **map[string]interface{}** | Filesystem information of the Drive. | [optional] 
**Id** | **float32** | Id of the Drive | 
**Revision** | **float32** | Revision of the Drive State | 
**InfrastructureId** | **float32** | Infrastructure id of the Drive | 
**ServiceStatus** | **string** | Service status of the Drive | 
**StorageRealSizeCachedMb** | Pointer to **float32** | Cached information of the real size of the storage in MB. | [optional] 
**StorageRealSizeWithSnapshotsCachedMb** | Pointer to **float32** | Cached information of the real size of the storage (including snapshots) in MB. | [optional] 
**StorageVirtualSizeCachedMb** | Pointer to **float32** | Cached information of the virtual size of the storage in MB. | [optional] 
**StorageUpdatedTimestamp** | **string** | Timestamp of the latest update of cached information for the Drive. | 
**Targets** | Pointer to **[]map[string]interface{}** | Targets of the Drive. | [optional] 
**ClusterCustomInfo** | Pointer to **map[string]interface{}** | Custom information of the Drive. | [optional] 
**SshKeyPairInternalEncrypted** | Pointer to **map[string]interface{}** | Pair of SSH Keys to use on the Drive. | [optional] 
**Wwn** | Pointer to **string** |  | [optional] 
**ProvisioningProtocol** | **string** | Provisioning protocol of the Drive | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Drive. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the Drive. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Id of the permanent DNS subdomain for the Drive. | [optional] 
**NetworkVlanId** | Pointer to **float32** | Id of the VLAN for the Drive. | [optional] 
**Config** | [**DriveConfiguration**](DriveConfiguration.md) | The current changes to be deployed for the Drive. | 
**CreatedTimestamp** | **string** | Timestamp of the Drive creation. | 
**Meta** | [**DriveMeta**](DriveMeta.md) | Meta information of the Drive. | 

## Methods

### NewDrive

`func NewDrive(label string, groupId float32, sizeMb float32, storageType string, updatedTimestamp string, id float32, revision float32, infrastructureId float32, serviceStatus string, storageUpdatedTimestamp string, provisioningProtocol string, config DriveConfiguration, createdTimestamp string, meta DriveMeta, ) *Drive`

NewDrive instantiates a new Drive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveWithDefaults

`func NewDriveWithDefaults() *Drive`

NewDriveWithDefaults instantiates a new Drive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Drive) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Drive) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Drive) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetGroupId

`func (o *Drive) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Drive) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Drive) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetInstanceId

`func (o *Drive) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Drive) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Drive) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *Drive) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetStoragePoolId

`func (o *Drive) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *Drive) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *Drive) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *Drive) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetSizeMb

`func (o *Drive) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *Drive) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *Drive) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetStorageImageName

`func (o *Drive) GetStorageImageName() string`

GetStorageImageName returns the StorageImageName field if non-nil, zero value otherwise.

### GetStorageImageNameOk

`func (o *Drive) GetStorageImageNameOk() (*string, bool)`

GetStorageImageNameOk returns a tuple with the StorageImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageImageName

`func (o *Drive) SetStorageImageName(v string)`

SetStorageImageName sets StorageImageName field to given value.

### HasStorageImageName

`func (o *Drive) HasStorageImageName() bool`

HasStorageImageName returns a boolean if a field has been set.

### GetIscsiIndexHex

`func (o *Drive) GetIscsiIndexHex() string`

GetIscsiIndexHex returns the IscsiIndexHex field if non-nil, zero value otherwise.

### GetIscsiIndexHexOk

`func (o *Drive) GetIscsiIndexHexOk() (*string, bool)`

GetIscsiIndexHexOk returns a tuple with the IscsiIndexHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIndexHex

`func (o *Drive) SetIscsiIndexHex(v string)`

SetIscsiIndexHex sets IscsiIndexHex field to given value.

### HasIscsiIndexHex

`func (o *Drive) HasIscsiIndexHex() bool`

HasIscsiIndexHex returns a boolean if a field has been set.

### GetTemplateId

`func (o *Drive) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Drive) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Drive) SetTemplateId(v float32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Drive) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetOsAdminUsername

`func (o *Drive) GetOsAdminUsername() string`

GetOsAdminUsername returns the OsAdminUsername field if non-nil, zero value otherwise.

### GetOsAdminUsernameOk

`func (o *Drive) GetOsAdminUsernameOk() (*string, bool)`

GetOsAdminUsernameOk returns a tuple with the OsAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdminUsername

`func (o *Drive) SetOsAdminUsername(v string)`

SetOsAdminUsername sets OsAdminUsername field to given value.

### HasOsAdminUsername

`func (o *Drive) HasOsAdminUsername() bool`

HasOsAdminUsername returns a boolean if a field has been set.

### GetOsAdminPasswordEncrypted

`func (o *Drive) GetOsAdminPasswordEncrypted() string`

GetOsAdminPasswordEncrypted returns the OsAdminPasswordEncrypted field if non-nil, zero value otherwise.

### GetOsAdminPasswordEncryptedOk

`func (o *Drive) GetOsAdminPasswordEncryptedOk() (*string, bool)`

GetOsAdminPasswordEncryptedOk returns a tuple with the OsAdminPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdminPasswordEncrypted

`func (o *Drive) SetOsAdminPasswordEncrypted(v string)`

SetOsAdminPasswordEncrypted sets OsAdminPasswordEncrypted field to given value.

### HasOsAdminPasswordEncrypted

`func (o *Drive) HasOsAdminPasswordEncrypted() bool`

HasOsAdminPasswordEncrypted returns a boolean if a field has been set.

### GetStorageType

`func (o *Drive) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *Drive) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *Drive) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetSubdomain

`func (o *Drive) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Drive) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Drive) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Drive) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Drive) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Drive) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Drive) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSshPort

`func (o *Drive) GetSshPort() float32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *Drive) GetSshPortOk() (*float32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *Drive) SetSshPort(v float32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *Drive) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetOperatingSystemInfo

`func (o *Drive) GetOperatingSystemInfo() map[string]interface{}`

GetOperatingSystemInfo returns the OperatingSystemInfo field if non-nil, zero value otherwise.

### GetOperatingSystemInfoOk

`func (o *Drive) GetOperatingSystemInfoOk() (*map[string]interface{}, bool)`

GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemInfo

`func (o *Drive) SetOperatingSystemInfo(v map[string]interface{})`

SetOperatingSystemInfo sets OperatingSystemInfo field to given value.

### HasOperatingSystemInfo

`func (o *Drive) HasOperatingSystemInfo() bool`

HasOperatingSystemInfo returns a boolean if a field has been set.

### GetFilesystemInfo

`func (o *Drive) GetFilesystemInfo() map[string]interface{}`

GetFilesystemInfo returns the FilesystemInfo field if non-nil, zero value otherwise.

### GetFilesystemInfoOk

`func (o *Drive) GetFilesystemInfoOk() (*map[string]interface{}, bool)`

GetFilesystemInfoOk returns a tuple with the FilesystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemInfo

`func (o *Drive) SetFilesystemInfo(v map[string]interface{})`

SetFilesystemInfo sets FilesystemInfo field to given value.

### HasFilesystemInfo

`func (o *Drive) HasFilesystemInfo() bool`

HasFilesystemInfo returns a boolean if a field has been set.

### GetId

`func (o *Drive) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Drive) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Drive) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *Drive) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Drive) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Drive) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *Drive) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *Drive) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *Drive) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetServiceStatus

`func (o *Drive) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *Drive) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *Drive) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetStorageRealSizeCachedMb

`func (o *Drive) GetStorageRealSizeCachedMb() float32`

GetStorageRealSizeCachedMb returns the StorageRealSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeCachedMbOk

`func (o *Drive) GetStorageRealSizeCachedMbOk() (*float32, bool)`

GetStorageRealSizeCachedMbOk returns a tuple with the StorageRealSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeCachedMb

`func (o *Drive) SetStorageRealSizeCachedMb(v float32)`

SetStorageRealSizeCachedMb sets StorageRealSizeCachedMb field to given value.

### HasStorageRealSizeCachedMb

`func (o *Drive) HasStorageRealSizeCachedMb() bool`

HasStorageRealSizeCachedMb returns a boolean if a field has been set.

### GetStorageRealSizeWithSnapshotsCachedMb

`func (o *Drive) GetStorageRealSizeWithSnapshotsCachedMb() float32`

GetStorageRealSizeWithSnapshotsCachedMb returns the StorageRealSizeWithSnapshotsCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeWithSnapshotsCachedMbOk

`func (o *Drive) GetStorageRealSizeWithSnapshotsCachedMbOk() (*float32, bool)`

GetStorageRealSizeWithSnapshotsCachedMbOk returns a tuple with the StorageRealSizeWithSnapshotsCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeWithSnapshotsCachedMb

`func (o *Drive) SetStorageRealSizeWithSnapshotsCachedMb(v float32)`

SetStorageRealSizeWithSnapshotsCachedMb sets StorageRealSizeWithSnapshotsCachedMb field to given value.

### HasStorageRealSizeWithSnapshotsCachedMb

`func (o *Drive) HasStorageRealSizeWithSnapshotsCachedMb() bool`

HasStorageRealSizeWithSnapshotsCachedMb returns a boolean if a field has been set.

### GetStorageVirtualSizeCachedMb

`func (o *Drive) GetStorageVirtualSizeCachedMb() float32`

GetStorageVirtualSizeCachedMb returns the StorageVirtualSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageVirtualSizeCachedMbOk

`func (o *Drive) GetStorageVirtualSizeCachedMbOk() (*float32, bool)`

GetStorageVirtualSizeCachedMbOk returns a tuple with the StorageVirtualSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVirtualSizeCachedMb

`func (o *Drive) SetStorageVirtualSizeCachedMb(v float32)`

SetStorageVirtualSizeCachedMb sets StorageVirtualSizeCachedMb field to given value.

### HasStorageVirtualSizeCachedMb

`func (o *Drive) HasStorageVirtualSizeCachedMb() bool`

HasStorageVirtualSizeCachedMb returns a boolean if a field has been set.

### GetStorageUpdatedTimestamp

`func (o *Drive) GetStorageUpdatedTimestamp() string`

GetStorageUpdatedTimestamp returns the StorageUpdatedTimestamp field if non-nil, zero value otherwise.

### GetStorageUpdatedTimestampOk

`func (o *Drive) GetStorageUpdatedTimestampOk() (*string, bool)`

GetStorageUpdatedTimestampOk returns a tuple with the StorageUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUpdatedTimestamp

`func (o *Drive) SetStorageUpdatedTimestamp(v string)`

SetStorageUpdatedTimestamp sets StorageUpdatedTimestamp field to given value.


### GetTargets

`func (o *Drive) GetTargets() []map[string]interface{}`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *Drive) GetTargetsOk() (*[]map[string]interface{}, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *Drive) SetTargets(v []map[string]interface{})`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *Drive) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetClusterCustomInfo

`func (o *Drive) GetClusterCustomInfo() map[string]interface{}`

GetClusterCustomInfo returns the ClusterCustomInfo field if non-nil, zero value otherwise.

### GetClusterCustomInfoOk

`func (o *Drive) GetClusterCustomInfoOk() (*map[string]interface{}, bool)`

GetClusterCustomInfoOk returns a tuple with the ClusterCustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCustomInfo

`func (o *Drive) SetClusterCustomInfo(v map[string]interface{})`

SetClusterCustomInfo sets ClusterCustomInfo field to given value.

### HasClusterCustomInfo

`func (o *Drive) HasClusterCustomInfo() bool`

HasClusterCustomInfo returns a boolean if a field has been set.

### GetSshKeyPairInternalEncrypted

`func (o *Drive) GetSshKeyPairInternalEncrypted() map[string]interface{}`

GetSshKeyPairInternalEncrypted returns the SshKeyPairInternalEncrypted field if non-nil, zero value otherwise.

### GetSshKeyPairInternalEncryptedOk

`func (o *Drive) GetSshKeyPairInternalEncryptedOk() (*map[string]interface{}, bool)`

GetSshKeyPairInternalEncryptedOk returns a tuple with the SshKeyPairInternalEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPairInternalEncrypted

`func (o *Drive) SetSshKeyPairInternalEncrypted(v map[string]interface{})`

SetSshKeyPairInternalEncrypted sets SshKeyPairInternalEncrypted field to given value.

### HasSshKeyPairInternalEncrypted

`func (o *Drive) HasSshKeyPairInternalEncrypted() bool`

HasSshKeyPairInternalEncrypted returns a boolean if a field has been set.

### GetWwn

`func (o *Drive) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *Drive) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *Drive) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *Drive) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetProvisioningProtocol

`func (o *Drive) GetProvisioningProtocol() string`

GetProvisioningProtocol returns the ProvisioningProtocol field if non-nil, zero value otherwise.

### GetProvisioningProtocolOk

`func (o *Drive) GetProvisioningProtocolOk() (*string, bool)`

GetProvisioningProtocolOk returns a tuple with the ProvisioningProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProtocol

`func (o *Drive) SetProvisioningProtocol(v string)`

SetProvisioningProtocol sets ProvisioningProtocol field to given value.


### GetSubdomainPermanent

`func (o *Drive) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *Drive) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *Drive) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *Drive) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *Drive) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *Drive) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *Drive) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *Drive) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *Drive) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *Drive) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *Drive) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *Drive) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetNetworkVlanId

`func (o *Drive) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *Drive) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *Drive) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *Drive) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetConfig

`func (o *Drive) GetConfig() DriveConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Drive) GetConfigOk() (*DriveConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Drive) SetConfig(v DriveConfiguration)`

SetConfig sets Config field to given value.


### GetCreatedTimestamp

`func (o *Drive) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Drive) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Drive) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetMeta

`func (o *Drive) GetMeta() DriveMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Drive) GetMetaOk() (*DriveMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Drive) SetMeta(v DriveMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


