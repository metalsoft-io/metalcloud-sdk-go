# FileShareVariables

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

## Methods

### NewFileShareVariables

`func NewFileShareVariables(sizeGB float32, updatedTimestamp string, label string, id float32, revision float32, infrastructureId float32, infrastructure ParentInfrastructure, createdTimestamp string, serviceStatus string, config FileShareConfiguration, ) *FileShareVariables`

NewFileShareVariables instantiates a new FileShareVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareVariablesWithDefaults

`func NewFileShareVariablesWithDefaults() *FileShareVariables`

NewFileShareVariablesWithDefaults instantiates a new FileShareVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *FileShareVariables) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *FileShareVariables) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *FileShareVariables) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.


### GetUpdatedTimestamp

`func (o *FileShareVariables) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *FileShareVariables) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *FileShareVariables) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetStoragePoolId

`func (o *FileShareVariables) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *FileShareVariables) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *FileShareVariables) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *FileShareVariables) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetLabel

`func (o *FileShareVariables) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FileShareVariables) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FileShareVariables) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *FileShareVariables) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *FileShareVariables) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *FileShareVariables) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *FileShareVariables) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLogicalNetworkId

`func (o *FileShareVariables) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *FileShareVariables) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *FileShareVariables) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *FileShareVariables) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetId

`func (o *FileShareVariables) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileShareVariables) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileShareVariables) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *FileShareVariables) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FileShareVariables) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FileShareVariables) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *FileShareVariables) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *FileShareVariables) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *FileShareVariables) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInfrastructure

`func (o *FileShareVariables) GetInfrastructure() ParentInfrastructure`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *FileShareVariables) GetInfrastructureOk() (*ParentInfrastructure, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *FileShareVariables) SetInfrastructure(v ParentInfrastructure)`

SetInfrastructure sets Infrastructure field to given value.


### GetCreatedTimestamp

`func (o *FileShareVariables) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FileShareVariables) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FileShareVariables) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetServiceStatus

`func (o *FileShareVariables) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *FileShareVariables) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *FileShareVariables) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *FileShareVariables) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *FileShareVariables) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *FileShareVariables) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *FileShareVariables) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *FileShareVariables) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *FileShareVariables) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *FileShareVariables) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *FileShareVariables) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetNetworkVlanId

`func (o *FileShareVariables) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *FileShareVariables) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *FileShareVariables) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *FileShareVariables) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetDiscoverInformation

`func (o *FileShareVariables) GetDiscoverInformation() GenericFileShareDiscoverInformation`

GetDiscoverInformation returns the DiscoverInformation field if non-nil, zero value otherwise.

### GetDiscoverInformationOk

`func (o *FileShareVariables) GetDiscoverInformationOk() (*GenericFileShareDiscoverInformation, bool)`

GetDiscoverInformationOk returns a tuple with the DiscoverInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverInformation

`func (o *FileShareVariables) SetDiscoverInformation(v GenericFileShareDiscoverInformation)`

SetDiscoverInformation sets DiscoverInformation field to given value.

### HasDiscoverInformation

`func (o *FileShareVariables) HasDiscoverInformation() bool`

HasDiscoverInformation returns a boolean if a field has been set.

### GetConfig

`func (o *FileShareVariables) GetConfig() FileShareConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FileShareVariables) GetConfigOk() (*FileShareConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FileShareVariables) SetConfig(v FileShareConfiguration)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


