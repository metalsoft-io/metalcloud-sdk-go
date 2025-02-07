# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | **float32** | Disk size in GB for Bucket | 
**UpdatedTimestamp** | **string** | Timestamp of the Bucket last update. | 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the Bucket is assigned to | [optional] 
**Label** | **string** | Label of the Bucket. | 
**Subdomain** | **string** | Subdomain of the Bucket. | 
**Id** | **float32** | Id of the Bucket | 
**Revision** | **float32** | Revision of the Bucket | 
**InfrastructureId** | **float32** | Infrastructure id of the Bucket | 
**CreatedTimestamp** | **string** | Timestamp of the Bucket creation. | 
**ServiceStatus** | **string** | Service status of the Bucket | 
**SubdomainPermanent** | **string** | Subdomain permanent of the Bucket. | 
**DnsSubdomainId** | **float32** | Id of the DNS subdomain for the Bucket. | 
**NetworkVlanId** | Pointer to **float32** | Id of the VLAN for the Bucket. | [optional] 
**GuiSettings** | [**GenericGUISettings**](GenericGUISettings.md) |  | 
**Endpoint** | Pointer to **string** | Endpoint of the Bucket. | [optional] 
**AccessKeyId** | Pointer to **string** | Endpoint of the Bucket. | [optional] 
**SecretKeyEncrypted** | Pointer to **string** | Endpoint of the Bucket. | [optional] 
**Config** | [**BucketConfiguration**](BucketConfiguration.md) | The current changes to be deployed for the Bucket. | 

## Methods

### NewBucket

`func NewBucket(sizeGB float32, updatedTimestamp string, label string, subdomain string, id float32, revision float32, infrastructureId float32, createdTimestamp string, serviceStatus string, subdomainPermanent string, dnsSubdomainId float32, guiSettings GenericGUISettings, config BucketConfiguration, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *Bucket) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *Bucket) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *Bucket) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.


### GetUpdatedTimestamp

`func (o *Bucket) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Bucket) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Bucket) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetStoragePoolId

`func (o *Bucket) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *Bucket) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *Bucket) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *Bucket) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetLabel

`func (o *Bucket) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Bucket) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Bucket) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *Bucket) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Bucket) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Bucket) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.


### GetId

`func (o *Bucket) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bucket) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bucket) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *Bucket) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Bucket) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Bucket) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *Bucket) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *Bucket) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *Bucket) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetCreatedTimestamp

`func (o *Bucket) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Bucket) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Bucket) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetServiceStatus

`func (o *Bucket) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *Bucket) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *Bucket) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *Bucket) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *Bucket) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *Bucket) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.


### GetDnsSubdomainId

`func (o *Bucket) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *Bucket) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *Bucket) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.


### GetNetworkVlanId

`func (o *Bucket) GetNetworkVlanId() float32`

GetNetworkVlanId returns the NetworkVlanId field if non-nil, zero value otherwise.

### GetNetworkVlanIdOk

`func (o *Bucket) GetNetworkVlanIdOk() (*float32, bool)`

GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkVlanId

`func (o *Bucket) SetNetworkVlanId(v float32)`

SetNetworkVlanId sets NetworkVlanId field to given value.

### HasNetworkVlanId

`func (o *Bucket) HasNetworkVlanId() bool`

HasNetworkVlanId returns a boolean if a field has been set.

### GetGuiSettings

`func (o *Bucket) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *Bucket) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *Bucket) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.


### GetEndpoint

`func (o *Bucket) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Bucket) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Bucket) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Bucket) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *Bucket) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *Bucket) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *Bucket) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *Bucket) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretKeyEncrypted

`func (o *Bucket) GetSecretKeyEncrypted() string`

GetSecretKeyEncrypted returns the SecretKeyEncrypted field if non-nil, zero value otherwise.

### GetSecretKeyEncryptedOk

`func (o *Bucket) GetSecretKeyEncryptedOk() (*string, bool)`

GetSecretKeyEncryptedOk returns a tuple with the SecretKeyEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyEncrypted

`func (o *Bucket) SetSecretKeyEncrypted(v string)`

SetSecretKeyEncrypted sets SecretKeyEncrypted field to given value.

### HasSecretKeyEncrypted

`func (o *Bucket) HasSecretKeyEncrypted() bool`

HasSecretKeyEncrypted returns a boolean if a field has been set.

### GetConfig

`func (o *Bucket) GetConfig() BucketConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Bucket) GetConfigOk() (*BucketConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Bucket) SetConfig(v BucketConfiguration)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


