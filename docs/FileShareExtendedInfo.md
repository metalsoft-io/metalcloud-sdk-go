# FileShareExtendedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | **float32** | Disk size in GB for File Share | 
**UpdatedTimestamp** | **string** | Timestamp of the File Share last update. | 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the File Share is assigned to | [optional] 
**Label** | **string** | Label of the File Share. | 
**Subdomain** | Pointer to **string** | Subdomain of the File Share. | [optional] 
**Id** | **float32** | Id of the File Share | 
**Revision** | **float32** | Revision of the File Share | 
**InfrastructureId** | **float32** | Infrastructure id of the File Share | 
**CreatedTimestamp** | **string** | Timestamp of the File Share creation. | 
**ServiceStatus** | **string** | Service status of the File Share | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the File Share. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the File Share. | [optional] 
**NetworkVlanId** | Pointer to **float32** | Id of the VLAN for the File Share. | [optional] 
**GuiSettings** | [**GenericGUISettings**](GenericGUISettings.md) |  | 
**Endpoint** | Pointer to **string** | Endpoint of the File Share. | [optional] 
**Config** | [**FileShareConfiguration**](FileShareConfiguration.md) | The current changes to be deployed for the File Share. | 
**Infrastructure** | **map[string]interface{}** | Infrastructure information | 

## Methods

### NewFileShareExtendedInfo

`func NewFileShareExtendedInfo(sizeGB float32, updatedTimestamp string, label string, id float32, revision float32, infrastructureId float32, createdTimestamp string, serviceStatus string, guiSettings GenericGUISettings, config FileShareConfiguration, infrastructure map[string]interface{}, ) *FileShareExtendedInfo`

NewFileShareExtendedInfo instantiates a new FileShareExtendedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareExtendedInfoWithDefaults

`func NewFileShareExtendedInfoWithDefaults() *FileShareExtendedInfo`

NewFileShareExtendedInfoWithDefaults instantiates a new FileShareExtendedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *FileShareExtendedInfo) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *FileShareExtendedInfo) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *FileShareExtendedInfo) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.


### GetUpdatedTimestamp

`func (o *FileShareExtendedInfo) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *FileShareExtendedInfo) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *FileShareExtendedInfo) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetStoragePoolId

`func (o *FileShareExtendedInfo) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *FileShareExtendedInfo) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *FileShareExtendedInfo) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *FileShareExtendedInfo) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetLabel

`func (o *FileShareExtendedInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FileShareExtendedInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FileShareExtendedInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *FileShareExtendedInfo) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *FileShareExtendedInfo) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *FileShareExtendedInfo) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *FileShareExtendedInfo) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetId

`func (o *FileShareExtendedInfo) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileShareExtendedInfo) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileShareExtendedInfo) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *FileShareExtendedInfo) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FileShareExtendedInfo) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FileShareExtendedInfo) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *FileShareExtendedInfo) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *FileShareExtendedInfo) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *FileShareExtendedInfo) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetCreatedTimestamp

`func (o *FileShareExtendedInfo) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FileShareExtendedInfo) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FileShareExtendedInfo) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetServiceStatus

`func (o *FileShareExtendedInfo) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *FileShareExtendedInfo) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *FileShareExtendedInfo) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *FileShareExtendedInfo) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *FileShareExtendedInfo) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *FileShareExtendedInfo) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *FileShareExtendedInfo) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *FileShareExtendedInfo) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *FileShareExtendedInfo) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *FileShareExtendedInfo) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *FileShareExtendedInfo) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetNetworkVlanId

`func (o *FileShareExtendedInfo) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *FileShareExtendedInfo) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *FileShareExtendedInfo) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *FileShareExtendedInfo) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetGuiSettings

`func (o *FileShareExtendedInfo) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *FileShareExtendedInfo) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *FileShareExtendedInfo) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.


### GetEndpoint

`func (o *FileShareExtendedInfo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FileShareExtendedInfo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FileShareExtendedInfo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *FileShareExtendedInfo) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetConfig

`func (o *FileShareExtendedInfo) GetConfig() FileShareConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FileShareExtendedInfo) GetConfigOk() (*FileShareConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FileShareExtendedInfo) SetConfig(v FileShareConfiguration)`

SetConfig sets Config field to given value.


### GetInfrastructure

`func (o *FileShareExtendedInfo) GetInfrastructure() map[string]interface{}`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *FileShareExtendedInfo) GetInfrastructureOk() (*map[string]interface{}, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *FileShareExtendedInfo) SetInfrastructure(v map[string]interface{})`

SetInfrastructure sets Infrastructure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


