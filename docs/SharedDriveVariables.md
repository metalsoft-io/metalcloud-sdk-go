# SharedDriveVariables

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
**Targets** | Pointer to **[]map[string]interface{}** | Targets of the Shared Drive. | [optional] 
**Wwn** | Pointer to **string** |  | [optional] 
**AllocationAffinity** | **string** | Allocation affinity of the Shared Drive | 
**ProvisioningProtocol** | **string** | Provisioning protocol of the Shared Drive | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Shared Drive. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the Shared Drive. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Id of the permanent DNS subdomain for the Shared Drive. | [optional] 
**NetworkVlanId** | Pointer to **float32** | Id of the VLAN for the Shared Drive. | [optional] 
**Config** | [**SharedDriveConfiguration**](SharedDriveConfiguration.md) | The current changes to be deployed for the Shared Drive. | 
**CreatedTimestamp** | **string** | Timestamp of the Shared Drive creation. | 

## Methods

### NewSharedDriveVariables

`func NewSharedDriveVariables(label string, sizeMb float32, storageType string, updatedTimestamp string, id float32, revision float32, infrastructureId float32, serviceStatus string, storageUpdatedTimestamp string, allocationAffinity string, provisioningProtocol string, config SharedDriveConfiguration, createdTimestamp string, ) *SharedDriveVariables`

NewSharedDriveVariables instantiates a new SharedDriveVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDriveVariablesWithDefaults

`func NewSharedDriveVariablesWithDefaults() *SharedDriveVariables`

NewSharedDriveVariablesWithDefaults instantiates a new SharedDriveVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *SharedDriveVariables) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SharedDriveVariables) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SharedDriveVariables) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStoragePoolId

`func (o *SharedDriveVariables) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *SharedDriveVariables) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *SharedDriveVariables) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *SharedDriveVariables) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetSizeMb

`func (o *SharedDriveVariables) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *SharedDriveVariables) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *SharedDriveVariables) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetStorageImageName

`func (o *SharedDriveVariables) GetStorageImageName() string`

GetStorageImageName returns the StorageImageName field if non-nil, zero value otherwise.

### GetStorageImageNameOk

`func (o *SharedDriveVariables) GetStorageImageNameOk() (*string, bool)`

GetStorageImageNameOk returns a tuple with the StorageImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageImageName

`func (o *SharedDriveVariables) SetStorageImageName(v string)`

SetStorageImageName sets StorageImageName field to given value.

### HasStorageImageName

`func (o *SharedDriveVariables) HasStorageImageName() bool`

HasStorageImageName returns a boolean if a field has been set.

### GetStorageType

`func (o *SharedDriveVariables) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SharedDriveVariables) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SharedDriveVariables) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetIoLimitPolicy

`func (o *SharedDriveVariables) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *SharedDriveVariables) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *SharedDriveVariables) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *SharedDriveVariables) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetSubdomain

`func (o *SharedDriveVariables) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *SharedDriveVariables) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *SharedDriveVariables) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *SharedDriveVariables) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SharedDriveVariables) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SharedDriveVariables) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SharedDriveVariables) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *SharedDriveVariables) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedDriveVariables) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedDriveVariables) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *SharedDriveVariables) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SharedDriveVariables) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SharedDriveVariables) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *SharedDriveVariables) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *SharedDriveVariables) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *SharedDriveVariables) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetServiceStatus

`func (o *SharedDriveVariables) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *SharedDriveVariables) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *SharedDriveVariables) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetStorageRealSizeCachedMb

`func (o *SharedDriveVariables) GetStorageRealSizeCachedMb() float32`

GetStorageRealSizeCachedMb returns the StorageRealSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeCachedMbOk

`func (o *SharedDriveVariables) GetStorageRealSizeCachedMbOk() (*float32, bool)`

GetStorageRealSizeCachedMbOk returns a tuple with the StorageRealSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeCachedMb

`func (o *SharedDriveVariables) SetStorageRealSizeCachedMb(v float32)`

SetStorageRealSizeCachedMb sets StorageRealSizeCachedMb field to given value.

### HasStorageRealSizeCachedMb

`func (o *SharedDriveVariables) HasStorageRealSizeCachedMb() bool`

HasStorageRealSizeCachedMb returns a boolean if a field has been set.

### GetStorageRealSizeWithSnapshotsCachedMb

`func (o *SharedDriveVariables) GetStorageRealSizeWithSnapshotsCachedMb() float32`

GetStorageRealSizeWithSnapshotsCachedMb returns the StorageRealSizeWithSnapshotsCachedMb field if non-nil, zero value otherwise.

### GetStorageRealSizeWithSnapshotsCachedMbOk

`func (o *SharedDriveVariables) GetStorageRealSizeWithSnapshotsCachedMbOk() (*float32, bool)`

GetStorageRealSizeWithSnapshotsCachedMbOk returns a tuple with the StorageRealSizeWithSnapshotsCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRealSizeWithSnapshotsCachedMb

`func (o *SharedDriveVariables) SetStorageRealSizeWithSnapshotsCachedMb(v float32)`

SetStorageRealSizeWithSnapshotsCachedMb sets StorageRealSizeWithSnapshotsCachedMb field to given value.

### HasStorageRealSizeWithSnapshotsCachedMb

`func (o *SharedDriveVariables) HasStorageRealSizeWithSnapshotsCachedMb() bool`

HasStorageRealSizeWithSnapshotsCachedMb returns a boolean if a field has been set.

### GetStorageVirtualSizeCachedMb

`func (o *SharedDriveVariables) GetStorageVirtualSizeCachedMb() float32`

GetStorageVirtualSizeCachedMb returns the StorageVirtualSizeCachedMb field if non-nil, zero value otherwise.

### GetStorageVirtualSizeCachedMbOk

`func (o *SharedDriveVariables) GetStorageVirtualSizeCachedMbOk() (*float32, bool)`

GetStorageVirtualSizeCachedMbOk returns a tuple with the StorageVirtualSizeCachedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVirtualSizeCachedMb

`func (o *SharedDriveVariables) SetStorageVirtualSizeCachedMb(v float32)`

SetStorageVirtualSizeCachedMb sets StorageVirtualSizeCachedMb field to given value.

### HasStorageVirtualSizeCachedMb

`func (o *SharedDriveVariables) HasStorageVirtualSizeCachedMb() bool`

HasStorageVirtualSizeCachedMb returns a boolean if a field has been set.

### GetStorageUpdatedTimestamp

`func (o *SharedDriveVariables) GetStorageUpdatedTimestamp() string`

GetStorageUpdatedTimestamp returns the StorageUpdatedTimestamp field if non-nil, zero value otherwise.

### GetStorageUpdatedTimestampOk

`func (o *SharedDriveVariables) GetStorageUpdatedTimestampOk() (*string, bool)`

GetStorageUpdatedTimestampOk returns a tuple with the StorageUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUpdatedTimestamp

`func (o *SharedDriveVariables) SetStorageUpdatedTimestamp(v string)`

SetStorageUpdatedTimestamp sets StorageUpdatedTimestamp field to given value.


### GetTargets

`func (o *SharedDriveVariables) GetTargets() []map[string]interface{}`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SharedDriveVariables) GetTargetsOk() (*[]map[string]interface{}, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SharedDriveVariables) SetTargets(v []map[string]interface{})`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SharedDriveVariables) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetWwn

`func (o *SharedDriveVariables) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *SharedDriveVariables) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *SharedDriveVariables) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *SharedDriveVariables) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetAllocationAffinity

`func (o *SharedDriveVariables) GetAllocationAffinity() string`

GetAllocationAffinity returns the AllocationAffinity field if non-nil, zero value otherwise.

### GetAllocationAffinityOk

`func (o *SharedDriveVariables) GetAllocationAffinityOk() (*string, bool)`

GetAllocationAffinityOk returns a tuple with the AllocationAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationAffinity

`func (o *SharedDriveVariables) SetAllocationAffinity(v string)`

SetAllocationAffinity sets AllocationAffinity field to given value.


### GetProvisioningProtocol

`func (o *SharedDriveVariables) GetProvisioningProtocol() string`

GetProvisioningProtocol returns the ProvisioningProtocol field if non-nil, zero value otherwise.

### GetProvisioningProtocolOk

`func (o *SharedDriveVariables) GetProvisioningProtocolOk() (*string, bool)`

GetProvisioningProtocolOk returns a tuple with the ProvisioningProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProtocol

`func (o *SharedDriveVariables) SetProvisioningProtocol(v string)`

SetProvisioningProtocol sets ProvisioningProtocol field to given value.


### GetSubdomainPermanent

`func (o *SharedDriveVariables) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *SharedDriveVariables) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *SharedDriveVariables) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *SharedDriveVariables) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *SharedDriveVariables) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *SharedDriveVariables) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *SharedDriveVariables) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *SharedDriveVariables) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *SharedDriveVariables) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *SharedDriveVariables) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *SharedDriveVariables) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *SharedDriveVariables) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetNetworkVlanId

`func (o *SharedDriveVariables) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *SharedDriveVariables) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *SharedDriveVariables) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *SharedDriveVariables) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetConfig

`func (o *SharedDriveVariables) GetConfig() SharedDriveConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SharedDriveVariables) GetConfigOk() (*SharedDriveConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SharedDriveVariables) SetConfig(v SharedDriveConfiguration)`

SetConfig sets Config field to given value.


### GetCreatedTimestamp

`func (o *SharedDriveVariables) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SharedDriveVariables) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SharedDriveVariables) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


