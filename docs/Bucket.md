# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | **float32** | Disk size in GB for Bucket | 
**UpdatedTimestamp** | **string** | Timestamp of the Bucket last update. | 
**StoragePoolId** | Pointer to **int64** | Id of the storage pool the Bucket is assigned to | [optional] 
**Label** | **string** | Label of the Bucket. | 
**Subdomain** | **string** | Subdomain of the Bucket. | 
**LogicalNetworkId** | Pointer to **int64** | Id of the Logical Network for the Bucket. | [optional] 
**Id** | **int64** | Id of the Bucket | 
**Revision** | **int64** | Revision of the Bucket | 
**InfrastructureId** | **int64** | Infrastructure id of the Bucket | 
**Infrastructure** | [**ParentInfrastructure**](ParentInfrastructure.md) | Infrastructure information | 
**CreatedTimestamp** | **string** | Timestamp of the Bucket creation. | 
**ServiceStatus** | **string** | Service status of the Bucket | 
**SubdomainPermanent** | **string** | Subdomain permanent of the Bucket. | 
**DnsSubdomainId** | **int64** | Id of the DNS subdomain for the Bucket. | 
**DiscoverInformation** | Pointer to [**GenericBucketDiscoverInformation**](GenericBucketDiscoverInformation.md) | Discover information of the Bucket. | [optional] 
**AccessKeyId** | Pointer to **string** | Access Key ID of the Bucket. | [optional] 
**Config** | [**BucketConfiguration**](BucketConfiguration.md) | The current changes to be deployed for the Bucket. | 
**Meta** | [**BucketMeta**](BucketMeta.md) | Meta information of the Bucket. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewBucket

`func NewBucket(sizeGB float32, updatedTimestamp string, label string, subdomain string, id int64, revision int64, infrastructureId int64, infrastructure ParentInfrastructure, createdTimestamp string, serviceStatus string, subdomainPermanent string, dnsSubdomainId int64, config BucketConfiguration, meta BucketMeta, ) *Bucket`

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

`func (o *Bucket) GetStoragePoolId() int64`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *Bucket) GetStoragePoolIdOk() (*int64, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *Bucket) SetStoragePoolId(v int64)`

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


### GetLogicalNetworkId

`func (o *Bucket) GetLogicalNetworkId() int64`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *Bucket) GetLogicalNetworkIdOk() (*int64, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *Bucket) SetLogicalNetworkId(v int64)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *Bucket) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetId

`func (o *Bucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bucket) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *Bucket) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Bucket) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Bucket) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *Bucket) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *Bucket) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *Bucket) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInfrastructure

`func (o *Bucket) GetInfrastructure() ParentInfrastructure`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *Bucket) GetInfrastructureOk() (*ParentInfrastructure, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *Bucket) SetInfrastructure(v ParentInfrastructure)`

SetInfrastructure sets Infrastructure field to given value.


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

`func (o *Bucket) GetDnsSubdomainId() int64`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *Bucket) GetDnsSubdomainIdOk() (*int64, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *Bucket) SetDnsSubdomainId(v int64)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.


### GetDiscoverInformation

`func (o *Bucket) GetDiscoverInformation() GenericBucketDiscoverInformation`

GetDiscoverInformation returns the DiscoverInformation field if non-nil, zero value otherwise.

### GetDiscoverInformationOk

`func (o *Bucket) GetDiscoverInformationOk() (*GenericBucketDiscoverInformation, bool)`

GetDiscoverInformationOk returns a tuple with the DiscoverInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverInformation

`func (o *Bucket) SetDiscoverInformation(v GenericBucketDiscoverInformation)`

SetDiscoverInformation sets DiscoverInformation field to given value.

### HasDiscoverInformation

`func (o *Bucket) HasDiscoverInformation() bool`

HasDiscoverInformation returns a boolean if a field has been set.

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


### GetMeta

`func (o *Bucket) GetMeta() BucketMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Bucket) GetMetaOk() (*BucketMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Bucket) SetMeta(v BucketMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *Bucket) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Bucket) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Bucket) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Bucket) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


