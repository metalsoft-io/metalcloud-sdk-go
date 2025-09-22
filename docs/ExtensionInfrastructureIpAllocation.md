# ExtensionInfrastructureIpAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]string** | Tags for the IP allocation. | [optional] 
**IpVersion** | [**IpVersion**](IpVersion.md) | IP version for the allocation. | 
**DnsRecords** | Pointer to [**[]ExtensionDnsRecord**](ExtensionDnsRecord.md) | DNS records for the IP allocation. | [optional] 

## Methods

### NewExtensionInfrastructureIpAllocation

`func NewExtensionInfrastructureIpAllocation(ipVersion IpVersion, ) *ExtensionInfrastructureIpAllocation`

NewExtensionInfrastructureIpAllocation instantiates a new ExtensionInfrastructureIpAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInfrastructureIpAllocationWithDefaults

`func NewExtensionInfrastructureIpAllocationWithDefaults() *ExtensionInfrastructureIpAllocation`

NewExtensionInfrastructureIpAllocationWithDefaults instantiates a new ExtensionInfrastructureIpAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ExtensionInfrastructureIpAllocation) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtensionInfrastructureIpAllocation) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtensionInfrastructureIpAllocation) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtensionInfrastructureIpAllocation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpVersion

`func (o *ExtensionInfrastructureIpAllocation) GetIpVersion() IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *ExtensionInfrastructureIpAllocation) GetIpVersionOk() (*IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *ExtensionInfrastructureIpAllocation) SetIpVersion(v IpVersion)`

SetIpVersion sets IpVersion field to given value.


### GetDnsRecords

`func (o *ExtensionInfrastructureIpAllocation) GetDnsRecords() []ExtensionDnsRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *ExtensionInfrastructureIpAllocation) GetDnsRecordsOk() (*[]ExtensionDnsRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *ExtensionInfrastructureIpAllocation) SetDnsRecords(v []ExtensionDnsRecord)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *ExtensionInfrastructureIpAllocation) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


