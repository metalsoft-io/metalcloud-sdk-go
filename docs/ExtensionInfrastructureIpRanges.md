# ExtensionInfrastructureIpRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]string** | Tags for the IP range allocations. | [optional] 
**IpVersion** | [**IpVersion**](IpVersion.md) | IP version for the allocation. | 
**IpCount** | **int32** | Number of IPs to allocate from the subnet. | 

## Methods

### NewExtensionInfrastructureIpRanges

`func NewExtensionInfrastructureIpRanges(ipVersion IpVersion, ipCount int32, ) *ExtensionInfrastructureIpRanges`

NewExtensionInfrastructureIpRanges instantiates a new ExtensionInfrastructureIpRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInfrastructureIpRangesWithDefaults

`func NewExtensionInfrastructureIpRangesWithDefaults() *ExtensionInfrastructureIpRanges`

NewExtensionInfrastructureIpRangesWithDefaults instantiates a new ExtensionInfrastructureIpRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ExtensionInfrastructureIpRanges) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtensionInfrastructureIpRanges) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtensionInfrastructureIpRanges) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtensionInfrastructureIpRanges) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpVersion

`func (o *ExtensionInfrastructureIpRanges) GetIpVersion() IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *ExtensionInfrastructureIpRanges) GetIpVersionOk() (*IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *ExtensionInfrastructureIpRanges) SetIpVersion(v IpVersion)`

SetIpVersion sets IpVersion field to given value.


### GetIpCount

`func (o *ExtensionInfrastructureIpRanges) GetIpCount() int32`

GetIpCount returns the IpCount field if non-nil, zero value otherwise.

### GetIpCountOk

`func (o *ExtensionInfrastructureIpRanges) GetIpCountOk() (*int32, bool)`

GetIpCountOk returns a tuple with the IpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCount

`func (o *ExtensionInfrastructureIpRanges) SetIpCount(v int32)`

SetIpCount sets IpCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


