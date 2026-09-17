# DhcpManualAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Ip** | **string** | The IP address to pin for this reservation | 
**SubnetId** | **int64** | The IPAM OOB subnet id that contains the pinned IP | 

## Methods

### NewDhcpManualAllocation

`func NewDhcpManualAllocation(kind string, ip string, subnetId int64, ) *DhcpManualAllocation`

NewDhcpManualAllocation instantiates a new DhcpManualAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpManualAllocationWithDefaults

`func NewDhcpManualAllocationWithDefaults() *DhcpManualAllocation`

NewDhcpManualAllocationWithDefaults instantiates a new DhcpManualAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DhcpManualAllocation) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DhcpManualAllocation) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DhcpManualAllocation) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIp

`func (o *DhcpManualAllocation) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DhcpManualAllocation) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DhcpManualAllocation) SetIp(v string)`

SetIp sets Ip field to given value.


### GetSubnetId

`func (o *DhcpManualAllocation) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *DhcpManualAllocation) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *DhcpManualAllocation) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


