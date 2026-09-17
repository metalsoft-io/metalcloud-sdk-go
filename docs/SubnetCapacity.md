# SubnetCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetId** | **int64** |  | 
**Prefix** | **string** | The subnet prefix in CIDR notation | 
**IsPool** | **bool** |  | 
**TotalIpCount** | **NullableString** | Total assignable addresses, excluding the network and broadcast addresses. Null for a pool. | 
**FreeIpCount** | **NullableString** | Addresses still free for allocation. Null for a pool. | 
**UsedIpCount** | **NullableString** | Addresses no longer available: allocated, reserved, or denylisted. Null for a pool. | 
**FreePrefixes** | **[]string** | The free blocks of a pool, in largest-block CIDR form: two adjacent free /24s are reported as one /23. Null for a subnet that is not a pool. | 
**FreePrefixCounts** | [**[]PrefixLengthCount**](PrefixLengthCount.md) | The free blocks of a pool grouped by prefix length. Null for a subnet that is not a pool. | 
**RequestedPrefixLength** | **NullableInt32** | Echoes the requested prefixLength. Null when none was requested. | 
**AllocatablePrefixCount** | **NullableString** | How many blocks of the requested prefix length still fit in this pool, as a decimal string. Null when no prefixLength was requested. | 

## Methods

### NewSubnetCapacity

`func NewSubnetCapacity(subnetId int64, prefix string, isPool bool, totalIpCount NullableString, freeIpCount NullableString, usedIpCount NullableString, freePrefixes []string, freePrefixCounts []PrefixLengthCount, requestedPrefixLength NullableInt32, allocatablePrefixCount NullableString, ) *SubnetCapacity`

NewSubnetCapacity instantiates a new SubnetCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetCapacityWithDefaults

`func NewSubnetCapacityWithDefaults() *SubnetCapacity`

NewSubnetCapacityWithDefaults instantiates a new SubnetCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetId

`func (o *SubnetCapacity) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *SubnetCapacity) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *SubnetCapacity) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### GetPrefix

`func (o *SubnetCapacity) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SubnetCapacity) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SubnetCapacity) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetIsPool

`func (o *SubnetCapacity) GetIsPool() bool`

GetIsPool returns the IsPool field if non-nil, zero value otherwise.

### GetIsPoolOk

`func (o *SubnetCapacity) GetIsPoolOk() (*bool, bool)`

GetIsPoolOk returns a tuple with the IsPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPool

`func (o *SubnetCapacity) SetIsPool(v bool)`

SetIsPool sets IsPool field to given value.


### GetTotalIpCount

`func (o *SubnetCapacity) GetTotalIpCount() string`

GetTotalIpCount returns the TotalIpCount field if non-nil, zero value otherwise.

### GetTotalIpCountOk

`func (o *SubnetCapacity) GetTotalIpCountOk() (*string, bool)`

GetTotalIpCountOk returns a tuple with the TotalIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIpCount

`func (o *SubnetCapacity) SetTotalIpCount(v string)`

SetTotalIpCount sets TotalIpCount field to given value.


### SetTotalIpCountNil

`func (o *SubnetCapacity) SetTotalIpCountNil(b bool)`

 SetTotalIpCountNil sets the value for TotalIpCount to be an explicit nil

### UnsetTotalIpCount
`func (o *SubnetCapacity) UnsetTotalIpCount()`

UnsetTotalIpCount ensures that no value is present for TotalIpCount, not even an explicit nil
### GetFreeIpCount

`func (o *SubnetCapacity) GetFreeIpCount() string`

GetFreeIpCount returns the FreeIpCount field if non-nil, zero value otherwise.

### GetFreeIpCountOk

`func (o *SubnetCapacity) GetFreeIpCountOk() (*string, bool)`

GetFreeIpCountOk returns a tuple with the FreeIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeIpCount

`func (o *SubnetCapacity) SetFreeIpCount(v string)`

SetFreeIpCount sets FreeIpCount field to given value.


### SetFreeIpCountNil

`func (o *SubnetCapacity) SetFreeIpCountNil(b bool)`

 SetFreeIpCountNil sets the value for FreeIpCount to be an explicit nil

### UnsetFreeIpCount
`func (o *SubnetCapacity) UnsetFreeIpCount()`

UnsetFreeIpCount ensures that no value is present for FreeIpCount, not even an explicit nil
### GetUsedIpCount

`func (o *SubnetCapacity) GetUsedIpCount() string`

GetUsedIpCount returns the UsedIpCount field if non-nil, zero value otherwise.

### GetUsedIpCountOk

`func (o *SubnetCapacity) GetUsedIpCountOk() (*string, bool)`

GetUsedIpCountOk returns a tuple with the UsedIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedIpCount

`func (o *SubnetCapacity) SetUsedIpCount(v string)`

SetUsedIpCount sets UsedIpCount field to given value.


### SetUsedIpCountNil

`func (o *SubnetCapacity) SetUsedIpCountNil(b bool)`

 SetUsedIpCountNil sets the value for UsedIpCount to be an explicit nil

### UnsetUsedIpCount
`func (o *SubnetCapacity) UnsetUsedIpCount()`

UnsetUsedIpCount ensures that no value is present for UsedIpCount, not even an explicit nil
### GetFreePrefixes

`func (o *SubnetCapacity) GetFreePrefixes() []string`

GetFreePrefixes returns the FreePrefixes field if non-nil, zero value otherwise.

### GetFreePrefixesOk

`func (o *SubnetCapacity) GetFreePrefixesOk() (*[]string, bool)`

GetFreePrefixesOk returns a tuple with the FreePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreePrefixes

`func (o *SubnetCapacity) SetFreePrefixes(v []string)`

SetFreePrefixes sets FreePrefixes field to given value.


### SetFreePrefixesNil

`func (o *SubnetCapacity) SetFreePrefixesNil(b bool)`

 SetFreePrefixesNil sets the value for FreePrefixes to be an explicit nil

### UnsetFreePrefixes
`func (o *SubnetCapacity) UnsetFreePrefixes()`

UnsetFreePrefixes ensures that no value is present for FreePrefixes, not even an explicit nil
### GetFreePrefixCounts

`func (o *SubnetCapacity) GetFreePrefixCounts() []PrefixLengthCount`

GetFreePrefixCounts returns the FreePrefixCounts field if non-nil, zero value otherwise.

### GetFreePrefixCountsOk

`func (o *SubnetCapacity) GetFreePrefixCountsOk() (*[]PrefixLengthCount, bool)`

GetFreePrefixCountsOk returns a tuple with the FreePrefixCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreePrefixCounts

`func (o *SubnetCapacity) SetFreePrefixCounts(v []PrefixLengthCount)`

SetFreePrefixCounts sets FreePrefixCounts field to given value.


### SetFreePrefixCountsNil

`func (o *SubnetCapacity) SetFreePrefixCountsNil(b bool)`

 SetFreePrefixCountsNil sets the value for FreePrefixCounts to be an explicit nil

### UnsetFreePrefixCounts
`func (o *SubnetCapacity) UnsetFreePrefixCounts()`

UnsetFreePrefixCounts ensures that no value is present for FreePrefixCounts, not even an explicit nil
### GetRequestedPrefixLength

`func (o *SubnetCapacity) GetRequestedPrefixLength() int32`

GetRequestedPrefixLength returns the RequestedPrefixLength field if non-nil, zero value otherwise.

### GetRequestedPrefixLengthOk

`func (o *SubnetCapacity) GetRequestedPrefixLengthOk() (*int32, bool)`

GetRequestedPrefixLengthOk returns a tuple with the RequestedPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPrefixLength

`func (o *SubnetCapacity) SetRequestedPrefixLength(v int32)`

SetRequestedPrefixLength sets RequestedPrefixLength field to given value.


### SetRequestedPrefixLengthNil

`func (o *SubnetCapacity) SetRequestedPrefixLengthNil(b bool)`

 SetRequestedPrefixLengthNil sets the value for RequestedPrefixLength to be an explicit nil

### UnsetRequestedPrefixLength
`func (o *SubnetCapacity) UnsetRequestedPrefixLength()`

UnsetRequestedPrefixLength ensures that no value is present for RequestedPrefixLength, not even an explicit nil
### GetAllocatablePrefixCount

`func (o *SubnetCapacity) GetAllocatablePrefixCount() string`

GetAllocatablePrefixCount returns the AllocatablePrefixCount field if non-nil, zero value otherwise.

### GetAllocatablePrefixCountOk

`func (o *SubnetCapacity) GetAllocatablePrefixCountOk() (*string, bool)`

GetAllocatablePrefixCountOk returns a tuple with the AllocatablePrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatablePrefixCount

`func (o *SubnetCapacity) SetAllocatablePrefixCount(v string)`

SetAllocatablePrefixCount sets AllocatablePrefixCount field to given value.


### SetAllocatablePrefixCountNil

`func (o *SubnetCapacity) SetAllocatablePrefixCountNil(b bool)`

 SetAllocatablePrefixCountNil sets the value for AllocatablePrefixCount to be an explicit nil

### UnsetAllocatablePrefixCount
`func (o *SubnetCapacity) UnsetAllocatablePrefixCount()`

UnsetAllocatablePrefixCount ensures that no value is present for AllocatablePrefixCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


