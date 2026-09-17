# DhcpReservationAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Ip** | **string** | The IP address to pin for this reservation | 
**SubnetId** | **int64** | The IPAM OOB subnet id that contains the pinned IP | 
**SubnetPoolIds** | **[]float32** | Ordered IPAM OOB subnet ids to allocate the next free IP from | 

## Methods

### NewDhcpReservationAllocation

`func NewDhcpReservationAllocation(kind string, ip string, subnetId int64, subnetPoolIds []float32, ) *DhcpReservationAllocation`

NewDhcpReservationAllocation instantiates a new DhcpReservationAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpReservationAllocationWithDefaults

`func NewDhcpReservationAllocationWithDefaults() *DhcpReservationAllocation`

NewDhcpReservationAllocationWithDefaults instantiates a new DhcpReservationAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DhcpReservationAllocation) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DhcpReservationAllocation) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DhcpReservationAllocation) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIp

`func (o *DhcpReservationAllocation) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DhcpReservationAllocation) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DhcpReservationAllocation) SetIp(v string)`

SetIp sets Ip field to given value.


### GetSubnetId

`func (o *DhcpReservationAllocation) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *DhcpReservationAllocation) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *DhcpReservationAllocation) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetPoolIds

`func (o *DhcpReservationAllocation) GetSubnetPoolIds() []float32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *DhcpReservationAllocation) GetSubnetPoolIdsOk() (*[]float32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *DhcpReservationAllocation) SetSubnetPoolIds(v []float32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


