# SharedDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the Shared Drive. | 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the Shared Drive is assigned to | [optional] 
**SizeMb** | **float32** | Disk size in MB for Shared Drive | 
**StorageImageName** | Pointer to **string** | The name of the storage image used by the Shared Drive. | [optional] 
**StorageType** | **string** | Service status of the Shared Drive | [default to "iscsi_ssd"]
**IoLimitPolicy** | Pointer to **string** | The IO limit policy of the Shared Drive. | [optional] 
**Subdomain** | Pointer to **string** | Subdomain of the Shared Drive. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Shared Drive last update. | 
**Id** | **float32** | Id of the Shared Drive | 
**Revision** | **float32** | Revision of the Shared Drive State | 
**InfrastructureId** | **float32** | Infrastructure id of the Shared Drive | 
**ServiceStatus** | **string** | Service status of the Shared Drive | 
**StorageRealSizeCachedMb** | Pointer to **float32** | Cached information of the real size of the storage in MB. | [optional] 
**StorageRealSizeWithSnapshotsCachedMb** | Pointer to **float32** | Cached information of the real size of the storage (including snapshots) in MB. | [optional] 
**StorageVirtualSizeCachedMb** | Pointer to **float32** | Cached information of the virtual size of the storage in MB. | [optional] 
**StorageUpdatedTimestamp** | **string** | Timestamp of the latest update of cached information for the Shared Drive. | 
**Targets** | Pointer to **map[string]interface{}** | Targets of the Shared Drive. | [optional] 
**Wwn** | Pointer to **string** |  | [optional] 
**AllocationAffinity** | **string** | Allocation affinity of the Shared Drive | 
**ProvisioningProtocol** | **string** | Provisioning protocol of the Shared Drive | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Shared Drive. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the Shared Drive. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Id of the permanent DNS subdomain for the Shared Drive. | [optional] 
**NetworkVlanId** | Pointer to **float32** | Id of the VLAN for the Shared Drive. | [optional] 
**Config** | [**SharedDriveConfiguration**](SharedDriveConfiguration.md) | The current changes to be deployed for the Shared Drive. | 
**CreatedTimestamp** | **string** | Timestamp of the Shared Drive creation. | 
**Meta** | [**SharedDriveMeta**](SharedDriveMeta.md) | Meta information of the Shared Drive. | 

## Methods

### NewSharedDrive

`func NewSharedDrive(label string, sizeMb float32, storageType string, updatedTimestamp string, id float32, revision float32, infrastructureId float32, serviceStatus string, storageUpdatedTimestamp string, allocationAffinity string, provisioningProtocol string, config SharedDriveConfiguration, createdTimestamp string, meta SharedDriveMeta, ) *SharedDrive`

NewSharedDrive instantiates a new SharedDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDriveWithDefaults

`func NewSharedDriveWithDefaults() *SharedDrive`

NewSharedDriveWithDefaults instantiates a new SharedDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *SharedDrive) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SharedDrive) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SharedDrive) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStoragePoolId

`func (o *SharedDrive) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *SharedDrive) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *SharedDrive) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *SharedDrive) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetSizeMb

`func (o *SharedDrive) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *SharedDrive) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *SharedDrive) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetStorageImageName

`func (o *SharedDrive) GetStorageImageName() string`

GetStorageImageName returns the StorageImageName field if non-nil, zero value otherwise.

### GetStorageImageNameOk

`func (o *SharedDrive) GetStorageImageNameOk() (*string, bool)`

GetStorageImageNameOk returns a tuple with the StorageImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageImageName

`func (o *SharedDrive) SetStorageImageName(v string)`

SetStorageImageName sets StorageImageName field to given value.

### HasStorageImageName

`func (o *SharedDrive) HasStorageImageName() bool`

HasStorageImageName returns a boolean if a field has been set.

### GetStorageType

`func (o *SharedDrive) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SharedDrive) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SharedDrive) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetIoLimitPolicy

`func (o *SharedDrive) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *SharedDrive) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *SharedDrive) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *SharedDrive) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetSubdomain

`func (o *SharedDrive) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *SharedDrive) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *SharedDrive) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *SharedDrive) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SharedDrive) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SharedDrive) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SharedDrive) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *SharedDrive) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedDrive) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedDrive) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *SharedDrive) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SharedDrive) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SharedDrive) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *SharedDrive) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *SharedDrive) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *SharedDrive) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetServiceStatus

`func (o *SharedDrive) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *SharedDrive) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *SharedDrive) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetStorageRealSizeCachedMb

`func (o *SharedDrive) GetStorageRealSizeCachedMb() float32`

GetStorageRealSizeCachedMb returns the StorageRealSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeCachedMbOk

`func (o *SharedDrive) GetStorageRealSizeCachedMbOk() (*float32, bool)`

GetStorageRealSizeCachedMbOk returns a tuple with the StorageRealSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeCachedMb

`func (o *SharedDrive) SetStorageRealSizeCachedMb(v float32)`

SetStorageRealSizeCachedMb sets StorageRealSizeCachedMb field to given value.

### HasStorageRealSizeCachedMb

`func (o *SharedDrive) HasStorageRealSizeCachedMb() bool`

HasStorageRealSizeCachedMb returns a boolean if a field has been set.

### GetStorageRealSizeWithSnapshotsCachedMb

`func (o *SharedDrive) GetStorageRealSizeWithSnapshotsCachedMb() float32`

GetStorageRealSizeWithSnapshotsCachedMb returns the StorageRealSizeWithSnapshotsCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeWithSnapshotsCachedMbOk

`func (o *SharedDrive) GetStorageRealSizeWithSnapshotsCachedMbOk() (*float32, bool)`

GetStorageRealSizeWithSnapshotsCachedMbOk returns a tuple with the StorageRealSizeWithSnapshotsCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeWithSnapshotsCachedMb

`func (o *SharedDrive) SetStorageRealSizeWithSnapshotsCachedMb(v float32)`

SetStorageRealSizeWithSnapshotsCachedMb sets StorageRealSizeWithSnapshotsCachedMb field to given value.

### HasStorageRealSizeWithSnapshotsCachedMb

`func (o *SharedDrive) HasStorageRealSizeWithSnapshotsCachedMb() bool`

HasStorageRealSizeWithSnapshotsCachedMb returns a boolean if a field has been set.

### GetStorageVirtualSizeCachedMb

`func (o *SharedDrive) GetStorageVirtualSizeCachedMb() float32`

GetStorageVirtualSizeCachedMb returns the StorageVirtualSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageVirtualSizeCachedMbOk

`func (o *SharedDrive) GetStorageVirtualSizeCachedMbOk() (*float32, bool)`

GetStorageVirtualSizeCachedMbOk returns a tuple with the StorageVirtualSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVirtualSizeCachedMb

`func (o *SharedDrive) SetStorageVirtualSizeCachedMb(v float32)`

SetStorageVirtualSizeCachedMb sets StorageVirtualSizeCachedMb field to given value.

### HasStorageVirtualSizeCachedMb

`func (o *SharedDrive) HasStorageVirtualSizeCachedMb() bool`

HasStorageVirtualSizeCachedMb returns a boolean if a field has been set.

### GetStorageUpdatedTimestamp

`func (o *SharedDrive) GetStorageUpdatedTimestamp() string`

GetStorageUpdatedTimestamp returns the StorageUpdatedTimestamp field if non-nil, zero value otherwise.

### GetStorageUpdatedTimestampOk

`func (o *SharedDrive) GetStorageUpdatedTimestampOk() (*string, bool)`

GetStorageUpdatedTimestampOk returns a tuple with the StorageUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUpdatedTimestamp

`func (o *SharedDrive) SetStorageUpdatedTimestamp(v string)`

SetStorageUpdatedTimestamp sets StorageUpdatedTimestamp field to given value.


### GetTargets

`func (o *SharedDrive) GetTargets() map[string]interface{}`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SharedDrive) GetTargetsOk() (*map[string]interface{}, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SharedDrive) SetTargets(v map[string]interface{})`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SharedDrive) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetWwn

`func (o *SharedDrive) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *SharedDrive) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *SharedDrive) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *SharedDrive) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetAllocationAffinity

`func (o *SharedDrive) GetAllocationAffinity() string`

GetAllocationAffinity returns the AllocationAffinity field if non-nil, zero value otherwise.

### GetAllocationAffinityOk

`func (o *SharedDrive) GetAllocationAffinityOk() (*string, bool)`

GetAllocationAffinityOk returns a tuple with the AllocationAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationAffinity

`func (o *SharedDrive) SetAllocationAffinity(v string)`

SetAllocationAffinity sets AllocationAffinity field to given value.


### GetProvisioningProtocol

`func (o *SharedDrive) GetProvisioningProtocol() string`

GetProvisioningProtocol returns the ProvisioningProtocol field if non-nil, zero value otherwise.

### GetProvisioningProtocolOk

`func (o *SharedDrive) GetProvisioningProtocolOk() (*string, bool)`

GetProvisioningProtocolOk returns a tuple with the ProvisioningProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProtocol

`func (o *SharedDrive) SetProvisioningProtocol(v string)`

SetProvisioningProtocol sets ProvisioningProtocol field to given value.


### GetSubdomainPermanent

`func (o *SharedDrive) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *SharedDrive) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *SharedDrive) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *SharedDrive) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *SharedDrive) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *SharedDrive) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *SharedDrive) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *SharedDrive) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *SharedDrive) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *SharedDrive) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *SharedDrive) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *SharedDrive) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetNetworkVlanId

`func (o *SharedDrive) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *SharedDrive) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *SharedDrive) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *SharedDrive) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetConfig

`func (o *SharedDrive) GetConfig() SharedDriveConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SharedDrive) GetConfigOk() (*SharedDriveConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SharedDrive) SetConfig(v SharedDriveConfiguration)`

SetConfig sets Config field to given value.


### GetCreatedTimestamp

`func (o *SharedDrive) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SharedDrive) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SharedDrive) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetMeta

`func (o *SharedDrive) GetMeta() SharedDriveMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SharedDrive) GetMetaOk() (*SharedDriveMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SharedDrive) SetMeta(v SharedDriveMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


