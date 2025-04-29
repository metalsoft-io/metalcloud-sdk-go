# SharedDriveConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision of the Drive Configuration | 
**Label** | **string** | Label of the Drive. | 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the Drive is assigned to | [optional] 
**SizeMb** | **float32** | Disk size in MiB for Drive | 
**StorageImageName** | Pointer to **string** | The name of the storage image used by the Drive. | [optional] 
**StorageType** | **string** | Service status of the Drive | [default to "iscsi_ssd"]
**QoS** | Pointer to **string** | The QoS of the Drive. | [optional] 
**Subdomain** | Pointer to **string** | Subdomain of the Drive. | [optional] 
**DnsSubdomainChangeId** | Pointer to **float32** | Id of the DNS subdomain for the Drive. | [optional] 
**DeployType** | **string** | Deploy type of the Drive | [default to "create"]
**DeployStatus** | **string** | Deploy status of the Drive | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **float32** | Id of the deployment for the Drive. | [optional] 
**LogicalNetworkId** | Pointer to **float32** | Id of the Logical Network for the Drive. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Drive last update. | 

## Methods

### NewSharedDriveConfiguration

`func NewSharedDriveConfiguration(revision float32, label string, sizeMb float32, storageType string, deployType string, deployStatus string, updatedTimestamp string, ) *SharedDriveConfiguration`

NewSharedDriveConfiguration instantiates a new SharedDriveConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDriveConfigurationWithDefaults

`func NewSharedDriveConfigurationWithDefaults() *SharedDriveConfiguration`

NewSharedDriveConfigurationWithDefaults instantiates a new SharedDriveConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *SharedDriveConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SharedDriveConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SharedDriveConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *SharedDriveConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SharedDriveConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SharedDriveConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStoragePoolId

`func (o *SharedDriveConfiguration) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *SharedDriveConfiguration) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *SharedDriveConfiguration) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *SharedDriveConfiguration) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetSizeMb

`func (o *SharedDriveConfiguration) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *SharedDriveConfiguration) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *SharedDriveConfiguration) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetStorageImageName

`func (o *SharedDriveConfiguration) GetStorageImageName() string`

GetStorageImageName returns the StorageImageName field if non-nil, zero value otherwise.

### GetStorageImageNameOk

`func (o *SharedDriveConfiguration) GetStorageImageNameOk() (*string, bool)`

GetStorageImageNameOk returns a tuple with the StorageImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageImageName

`func (o *SharedDriveConfiguration) SetStorageImageName(v string)`

SetStorageImageName sets StorageImageName field to given value.

### HasStorageImageName

`func (o *SharedDriveConfiguration) HasStorageImageName() bool`

HasStorageImageName returns a boolean if a field has been set.

### GetStorageType

`func (o *SharedDriveConfiguration) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SharedDriveConfiguration) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SharedDriveConfiguration) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetQoS

`func (o *SharedDriveConfiguration) GetQoS() string`

GetQoS returns the QoS field if non-nil, zero value otherwise.

### GetQoSOk

`func (o *SharedDriveConfiguration) GetQoSOk() (*string, bool)`

GetQoSOk returns a tuple with the QoS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoS

`func (o *SharedDriveConfiguration) SetQoS(v string)`

SetQoS sets QoS field to given value.

### HasQoS

`func (o *SharedDriveConfiguration) HasQoS() bool`

HasQoS returns a boolean if a field has been set.

### GetSubdomain

`func (o *SharedDriveConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *SharedDriveConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *SharedDriveConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *SharedDriveConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *SharedDriveConfiguration) GetDnsSubdomainChangeId() float32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *SharedDriveConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *SharedDriveConfiguration) SetDnsSubdomainChangeId(v float32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *SharedDriveConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetDeployType

`func (o *SharedDriveConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *SharedDriveConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *SharedDriveConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *SharedDriveConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *SharedDriveConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *SharedDriveConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *SharedDriveConfiguration) GetInfrastructureDeployId() float32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *SharedDriveConfiguration) GetInfrastructureDeployIdOk() (*float32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *SharedDriveConfiguration) SetInfrastructureDeployId(v float32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *SharedDriveConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetLogicalNetworkId

`func (o *SharedDriveConfiguration) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *SharedDriveConfiguration) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *SharedDriveConfiguration) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *SharedDriveConfiguration) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SharedDriveConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SharedDriveConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SharedDriveConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


