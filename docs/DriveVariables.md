# DriveVariables

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
**Infrastructure** | [**ParentInfrastructureDto**](ParentInfrastructureDto.md) | Infrastructure information | 
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

## Methods

### NewDriveVariables

`func NewDriveVariables(label string, groupId float32, sizeMb float32, storageType string, updatedTimestamp string, id float32, revision float32, infrastructureId float32, infrastructure ParentInfrastructureDto, serviceStatus string, storageUpdatedTimestamp string, provisioningProtocol string, config DriveConfiguration, createdTimestamp string, ) *DriveVariables`

NewDriveVariables instantiates a new DriveVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveVariablesWithDefaults

`func NewDriveVariablesWithDefaults() *DriveVariables`

NewDriveVariablesWithDefaults instantiates a new DriveVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *DriveVariables) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DriveVariables) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DriveVariables) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetGroupId

`func (o *DriveVariables) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DriveVariables) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DriveVariables) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetInstanceId

`func (o *DriveVariables) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DriveVariables) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DriveVariables) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *DriveVariables) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetStoragePoolId

`func (o *DriveVariables) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *DriveVariables) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *DriveVariables) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *DriveVariables) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetSizeMb

`func (o *DriveVariables) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *DriveVariables) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *DriveVariables) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetStorageImageName

`func (o *DriveVariables) GetStorageImageName() string`

GetStorageImageName returns the StorageImageName field if non-nil, zero value otherwise.

### GetStorageImageNameOk

`func (o *DriveVariables) GetStorageImageNameOk() (*string, bool)`

GetStorageImageNameOk returns a tuple with the StorageImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageImageName

`func (o *DriveVariables) SetStorageImageName(v string)`

SetStorageImageName sets StorageImageName field to given value.

### HasStorageImageName

`func (o *DriveVariables) HasStorageImageName() bool`

HasStorageImageName returns a boolean if a field has been set.

### GetIscsiIndexHex

`func (o *DriveVariables) GetIscsiIndexHex() string`

GetIscsiIndexHex returns the IscsiIndexHex field if non-nil, zero value otherwise.

### GetIscsiIndexHexOk

`func (o *DriveVariables) GetIscsiIndexHexOk() (*string, bool)`

GetIscsiIndexHexOk returns a tuple with the IscsiIndexHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIndexHex

`func (o *DriveVariables) SetIscsiIndexHex(v string)`

SetIscsiIndexHex sets IscsiIndexHex field to given value.

### HasIscsiIndexHex

`func (o *DriveVariables) HasIscsiIndexHex() bool`

HasIscsiIndexHex returns a boolean if a field has been set.

### GetTemplateId

`func (o *DriveVariables) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DriveVariables) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DriveVariables) SetTemplateId(v float32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DriveVariables) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetOsAdminUsername

`func (o *DriveVariables) GetOsAdminUsername() string`

GetOsAdminUsername returns the OsAdminUsername field if non-nil, zero value otherwise.

### GetOsAdminUsernameOk

`func (o *DriveVariables) GetOsAdminUsernameOk() (*string, bool)`

GetOsAdminUsernameOk returns a tuple with the OsAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdminUsername

`func (o *DriveVariables) SetOsAdminUsername(v string)`

SetOsAdminUsername sets OsAdminUsername field to given value.

### HasOsAdminUsername

`func (o *DriveVariables) HasOsAdminUsername() bool`

HasOsAdminUsername returns a boolean if a field has been set.

### GetOsAdminPasswordEncrypted

`func (o *DriveVariables) GetOsAdminPasswordEncrypted() string`

GetOsAdminPasswordEncrypted returns the OsAdminPasswordEncrypted field if non-nil, zero value otherwise.

### GetOsAdminPasswordEncryptedOk

`func (o *DriveVariables) GetOsAdminPasswordEncryptedOk() (*string, bool)`

GetOsAdminPasswordEncryptedOk returns a tuple with the OsAdminPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdminPasswordEncrypted

`func (o *DriveVariables) SetOsAdminPasswordEncrypted(v string)`

SetOsAdminPasswordEncrypted sets OsAdminPasswordEncrypted field to given value.

### HasOsAdminPasswordEncrypted

`func (o *DriveVariables) HasOsAdminPasswordEncrypted() bool`

HasOsAdminPasswordEncrypted returns a boolean if a field has been set.

### GetStorageType

`func (o *DriveVariables) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DriveVariables) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DriveVariables) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetSubdomain

`func (o *DriveVariables) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DriveVariables) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DriveVariables) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DriveVariables) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DriveVariables) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DriveVariables) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DriveVariables) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSshPort

`func (o *DriveVariables) GetSshPort() float32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *DriveVariables) GetSshPortOk() (*float32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *DriveVariables) SetSshPort(v float32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *DriveVariables) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetOperatingSystemInfo

`func (o *DriveVariables) GetOperatingSystemInfo() map[string]interface{}`

GetOperatingSystemInfo returns the OperatingSystemInfo field if non-nil, zero value otherwise.

### GetOperatingSystemInfoOk

`func (o *DriveVariables) GetOperatingSystemInfoOk() (*map[string]interface{}, bool)`

GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemInfo

`func (o *DriveVariables) SetOperatingSystemInfo(v map[string]interface{})`

SetOperatingSystemInfo sets OperatingSystemInfo field to given value.

### HasOperatingSystemInfo

`func (o *DriveVariables) HasOperatingSystemInfo() bool`

HasOperatingSystemInfo returns a boolean if a field has been set.

### GetFilesystemInfo

`func (o *DriveVariables) GetFilesystemInfo() map[string]interface{}`

GetFilesystemInfo returns the FilesystemInfo field if non-nil, zero value otherwise.

### GetFilesystemInfoOk

`func (o *DriveVariables) GetFilesystemInfoOk() (*map[string]interface{}, bool)`

GetFilesystemInfoOk returns a tuple with the FilesystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemInfo

`func (o *DriveVariables) SetFilesystemInfo(v map[string]interface{})`

SetFilesystemInfo sets FilesystemInfo field to given value.

### HasFilesystemInfo

`func (o *DriveVariables) HasFilesystemInfo() bool`

HasFilesystemInfo returns a boolean if a field has been set.

### GetId

`func (o *DriveVariables) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriveVariables) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriveVariables) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *DriveVariables) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DriveVariables) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DriveVariables) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *DriveVariables) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *DriveVariables) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *DriveVariables) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInfrastructure

`func (o *DriveVariables) GetInfrastructure() ParentInfrastructureDto`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *DriveVariables) GetInfrastructureOk() (*ParentInfrastructureDto, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *DriveVariables) SetInfrastructure(v ParentInfrastructureDto)`

SetInfrastructure sets Infrastructure field to given value.


### GetServiceStatus

`func (o *DriveVariables) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *DriveVariables) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *DriveVariables) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetStorageRealSizeCachedMb

`func (o *DriveVariables) GetStorageRealSizeCachedMb() float32`

GetStorageRealSizeCachedMb returns the StorageRealSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeCachedMbOk

`func (o *DriveVariables) GetStorageRealSizeCachedMbOk() (*float32, bool)`

GetStorageRealSizeCachedMbOk returns a tuple with the StorageRealSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeCachedMb

`func (o *DriveVariables) SetStorageRealSizeCachedMb(v float32)`

SetStorageRealSizeCachedMb sets StorageRealSizeCachedMb field to given value.

### HasStorageRealSizeCachedMb

`func (o *DriveVariables) HasStorageRealSizeCachedMb() bool`

HasStorageRealSizeCachedMb returns a boolean if a field has been set.

### GetStorageRealSizeWithSnapshotsCachedMb

`func (o *DriveVariables) GetStorageRealSizeWithSnapshotsCachedMb() float32`

GetStorageRealSizeWithSnapshotsCachedMb returns the StorageRealSizeWithSnapshotsCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeWithSnapshotsCachedMbOk

`func (o *DriveVariables) GetStorageRealSizeWithSnapshotsCachedMbOk() (*float32, bool)`

GetStorageRealSizeWithSnapshotsCachedMbOk returns a tuple with the StorageRealSizeWithSnapshotsCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeWithSnapshotsCachedMb

`func (o *DriveVariables) SetStorageRealSizeWithSnapshotsCachedMb(v float32)`

SetStorageRealSizeWithSnapshotsCachedMb sets StorageRealSizeWithSnapshotsCachedMb field to given value.

### HasStorageRealSizeWithSnapshotsCachedMb

`func (o *DriveVariables) HasStorageRealSizeWithSnapshotsCachedMb() bool`

HasStorageRealSizeWithSnapshotsCachedMb returns a boolean if a field has been set.

### GetStorageVirtualSizeCachedMb

`func (o *DriveVariables) GetStorageVirtualSizeCachedMb() float32`

GetStorageVirtualSizeCachedMb returns the StorageVirtualSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageVirtualSizeCachedMbOk

`func (o *DriveVariables) GetStorageVirtualSizeCachedMbOk() (*float32, bool)`

GetStorageVirtualSizeCachedMbOk returns a tuple with the StorageVirtualSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVirtualSizeCachedMb

`func (o *DriveVariables) SetStorageVirtualSizeCachedMb(v float32)`

SetStorageVirtualSizeCachedMb sets StorageVirtualSizeCachedMb field to given value.

### HasStorageVirtualSizeCachedMb

`func (o *DriveVariables) HasStorageVirtualSizeCachedMb() bool`

HasStorageVirtualSizeCachedMb returns a boolean if a field has been set.

### GetStorageUpdatedTimestamp

`func (o *DriveVariables) GetStorageUpdatedTimestamp() string`

GetStorageUpdatedTimestamp returns the StorageUpdatedTimestamp field if non-nil, zero value otherwise.

### GetStorageUpdatedTimestampOk

`func (o *DriveVariables) GetStorageUpdatedTimestampOk() (*string, bool)`

GetStorageUpdatedTimestampOk returns a tuple with the StorageUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUpdatedTimestamp

`func (o *DriveVariables) SetStorageUpdatedTimestamp(v string)`

SetStorageUpdatedTimestamp sets StorageUpdatedTimestamp field to given value.


### GetTargets

`func (o *DriveVariables) GetTargets() []map[string]interface{}`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *DriveVariables) GetTargetsOk() (*[]map[string]interface{}, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *DriveVariables) SetTargets(v []map[string]interface{})`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *DriveVariables) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetClusterCustomInfo

`func (o *DriveVariables) GetClusterCustomInfo() map[string]interface{}`

GetClusterCustomInfo returns the ClusterCustomInfo field if non-nil, zero value otherwise.

### GetClusterCustomInfoOk

`func (o *DriveVariables) GetClusterCustomInfoOk() (*map[string]interface{}, bool)`

GetClusterCustomInfoOk returns a tuple with the ClusterCustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCustomInfo

`func (o *DriveVariables) SetClusterCustomInfo(v map[string]interface{})`

SetClusterCustomInfo sets ClusterCustomInfo field to given value.

### HasClusterCustomInfo

`func (o *DriveVariables) HasClusterCustomInfo() bool`

HasClusterCustomInfo returns a boolean if a field has been set.

### GetSshKeyPairInternalEncrypted

`func (o *DriveVariables) GetSshKeyPairInternalEncrypted() map[string]interface{}`

GetSshKeyPairInternalEncrypted returns the SshKeyPairInternalEncrypted field if non-nil, zero value otherwise.

### GetSshKeyPairInternalEncryptedOk

`func (o *DriveVariables) GetSshKeyPairInternalEncryptedOk() (*map[string]interface{}, bool)`

GetSshKeyPairInternalEncryptedOk returns a tuple with the SshKeyPairInternalEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPairInternalEncrypted

`func (o *DriveVariables) SetSshKeyPairInternalEncrypted(v map[string]interface{})`

SetSshKeyPairInternalEncrypted sets SshKeyPairInternalEncrypted field to given value.

### HasSshKeyPairInternalEncrypted

`func (o *DriveVariables) HasSshKeyPairInternalEncrypted() bool`

HasSshKeyPairInternalEncrypted returns a boolean if a field has been set.

### GetWwn

`func (o *DriveVariables) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *DriveVariables) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *DriveVariables) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *DriveVariables) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetProvisioningProtocol

`func (o *DriveVariables) GetProvisioningProtocol() string`

GetProvisioningProtocol returns the ProvisioningProtocol field if non-nil, zero value otherwise.

### GetProvisioningProtocolOk

`func (o *DriveVariables) GetProvisioningProtocolOk() (*string, bool)`

GetProvisioningProtocolOk returns a tuple with the ProvisioningProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProtocol

`func (o *DriveVariables) SetProvisioningProtocol(v string)`

SetProvisioningProtocol sets ProvisioningProtocol field to given value.


### GetSubdomainPermanent

`func (o *DriveVariables) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *DriveVariables) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *DriveVariables) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *DriveVariables) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *DriveVariables) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *DriveVariables) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *DriveVariables) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *DriveVariables) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *DriveVariables) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *DriveVariables) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *DriveVariables) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *DriveVariables) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetNetworkVlanId

`func (o *DriveVariables) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *DriveVariables) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *DriveVariables) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *DriveVariables) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetConfig

`func (o *DriveVariables) GetConfig() DriveConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DriveVariables) GetConfigOk() (*DriveConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DriveVariables) SetConfig(v DriveConfiguration)`

SetConfig sets Config field to given value.


### GetCreatedTimestamp

`func (o *DriveVariables) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DriveVariables) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DriveVariables) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


