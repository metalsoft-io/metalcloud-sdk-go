# FileShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | **float32** | Disk size in GB for File Share | 
**UpdatedTimestamp** | **string** | Timestamp of the File Share last update. | 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the File Share is assigned to | [optional] 
**Label** | **string** | Label of the File Share. | 
**Subdomain** | Pointer to **string** | Subdomain of the File Share. | [optional] 
**LogicalNetworkId** | Pointer to **float32** | Id of the Logical Network for the File Share. | [optional] 
**Id** | **float32** | Id of the File Share | 
**Revision** | **float32** | Revision of the File Share | 
**InfrastructureId** | **float32** | Infrastructure id of the File Share | 
**Infrastructure** | [**ParentInfrastructure**](ParentInfrastructure.md) | Infrastructure information | 
**CreatedTimestamp** | **string** | Timestamp of the File Share creation. | 
**ServiceStatus** | **string** | Service status of the File Share | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the File Share. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the File Share. | [optional] 
**NetworkVlanId** | Pointer to **float32** | Id of the VLAN for the File Share. | [optional] 
**DiscoverInformation** | Pointer to [**GenericFileShareDiscoverInformation**](GenericFileShareDiscoverInformation.md) | Discover information of the File Share. | [optional] 
**Config** | [**FileShareConfiguration**](FileShareConfiguration.md) | The current changes to be deployed for the File Share. | 
**Meta** | [**FileShareMeta**](FileShareMeta.md) | Meta information of the File Share. | 

## Methods

### NewFileShare

`func NewFileShare(sizeGB float32, updatedTimestamp string, label string, id float32, revision float32, infrastructureId float32, infrastructure ParentInfrastructure, createdTimestamp string, serviceStatus string, config FileShareConfiguration, meta FileShareMeta, ) *FileShare`

NewFileShare instantiates a new FileShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareWithDefaults

`func NewFileShareWithDefaults() *FileShare`

NewFileShareWithDefaults instantiates a new FileShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *FileShare) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *FileShare) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *FileShare) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.


### GetUpdatedTimestamp

`func (o *FileShare) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *FileShare) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *FileShare) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetStoragePoolId

`func (o *FileShare) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *FileShare) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *FileShare) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *FileShare) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetLabel

`func (o *FileShare) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FileShare) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FileShare) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *FileShare) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *FileShare) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *FileShare) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *FileShare) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLogicalNetworkId

`func (o *FileShare) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *FileShare) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *FileShare) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *FileShare) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetId

`func (o *FileShare) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileShare) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileShare) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *FileShare) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FileShare) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FileShare) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *FileShare) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *FileShare) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *FileShare) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInfrastructure

`func (o *FileShare) GetInfrastructure() ParentInfrastructure`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *FileShare) GetInfrastructureOk() (*ParentInfrastructure, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *FileShare) SetInfrastructure(v ParentInfrastructure)`

SetInfrastructure sets Infrastructure field to given value.


### GetCreatedTimestamp

`func (o *FileShare) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FileShare) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FileShare) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetServiceStatus

`func (o *FileShare) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *FileShare) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *FileShare) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *FileShare) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *FileShare) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *FileShare) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *FileShare) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *FileShare) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *FileShare) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *FileShare) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *FileShare) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetNetworkVlanId

`func (o *FileShare) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *FileShare) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *FileShare) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *FileShare) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetDiscoverInformation

`func (o *FileShare) GetDiscoverInformation() GenericFileShareDiscoverInformation`

GetDiscoverInformation returns the DiscoverInformation field if non-nil, zero value otherwise.

### GetDiscoverInformationOk

`func (o *FileShare) GetDiscoverInformationOk() (*GenericFileShareDiscoverInformation, bool)`

GetDiscoverInformationOk returns a tuple with the DiscoverInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverInformation

`func (o *FileShare) SetDiscoverInformation(v GenericFileShareDiscoverInformation)`

SetDiscoverInformation sets DiscoverInformation field to given value.

### HasDiscoverInformation

`func (o *FileShare) HasDiscoverInformation() bool`

HasDiscoverInformation returns a boolean if a field has been set.

### GetConfig

`func (o *FileShare) GetConfig() FileShareConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FileShare) GetConfigOk() (*FileShareConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FileShare) SetConfig(v FileShareConfiguration)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *FileShare) GetMeta() FileShareMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileShare) GetMetaOk() (*FileShareMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileShare) SetMeta(v FileShareMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


