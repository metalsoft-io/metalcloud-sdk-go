# PlatformAllowedPrefixSizesOnWAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | Pointer to **[]float32** | Allowed IPv4 prefix sizes for WAN allocation (e.g. [27, 28, 29, 30]) | [optional] 
**Ipv6** | Pointer to **[]float32** | Allowed IPv6 prefix sizes for WAN allocation (e.g. [64]) | [optional] 

## Methods

### NewPlatformAllowedPrefixSizesOnWAN

`func NewPlatformAllowedPrefixSizesOnWAN() *PlatformAllowedPrefixSizesOnWAN`

NewPlatformAllowedPrefixSizesOnWAN instantiates a new PlatformAllowedPrefixSizesOnWAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAllowedPrefixSizesOnWANWithDefaults

`func NewPlatformAllowedPrefixSizesOnWANWithDefaults() *PlatformAllowedPrefixSizesOnWAN`

NewPlatformAllowedPrefixSizesOnWANWithDefaults instantiates a new PlatformAllowedPrefixSizesOnWAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *PlatformAllowedPrefixSizesOnWAN) GetIpv4() []float32`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PlatformAllowedPrefixSizesOnWAN) GetIpv4Ok() (*[]float32, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PlatformAllowedPrefixSizesOnWAN) SetIpv4(v []float32)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *PlatformAllowedPrefixSizesOnWAN) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *PlatformAllowedPrefixSizesOnWAN) GetIpv6() []float32`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *PlatformAllowedPrefixSizesOnWAN) GetIpv6Ok() (*[]float32, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *PlatformAllowedPrefixSizesOnWAN) SetIpv6(v []float32)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *PlatformAllowedPrefixSizesOnWAN) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


