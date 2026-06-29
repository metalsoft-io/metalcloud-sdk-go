# Ipv6PointToPointSubnetAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NetworkAddress** | **string** | Network address of the /127 (e.g. \&quot;2001:db8::\&quot;). | 
**PrefixLength** | **int32** |  | 
**InterfaceAIp** | **string** | IP assigned to interface A. | 
**InterfaceBIp** | **string** | IP assigned to interface B. | 
**Status** | [**ResourceAllocationStatus**](ResourceAllocationStatus.md) | Allocation lifecycle: &#39;allocated&#39; once the deploy materializes the subnet, &#39;deleting&#39; while staged for removal. | 

## Methods

### NewIpv6PointToPointSubnetAllocation

`func NewIpv6PointToPointSubnetAllocation(id int64, networkAddress string, prefixLength int32, interfaceAIp string, interfaceBIp string, status ResourceAllocationStatus, ) *Ipv6PointToPointSubnetAllocation`

NewIpv6PointToPointSubnetAllocation instantiates a new Ipv6PointToPointSubnetAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6PointToPointSubnetAllocationWithDefaults

`func NewIpv6PointToPointSubnetAllocationWithDefaults() *Ipv6PointToPointSubnetAllocation`

NewIpv6PointToPointSubnetAllocationWithDefaults instantiates a new Ipv6PointToPointSubnetAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ipv6PointToPointSubnetAllocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ipv6PointToPointSubnetAllocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ipv6PointToPointSubnetAllocation) SetId(v int64)`

SetId sets Id field to given value.


### GetNetworkAddress

`func (o *Ipv6PointToPointSubnetAllocation) GetNetworkAddress() string`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *Ipv6PointToPointSubnetAllocation) GetNetworkAddressOk() (*string, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *Ipv6PointToPointSubnetAllocation) SetNetworkAddress(v string)`

SetNetworkAddress sets NetworkAddress field to given value.


### GetPrefixLength

`func (o *Ipv6PointToPointSubnetAllocation) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Ipv6PointToPointSubnetAllocation) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Ipv6PointToPointSubnetAllocation) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetInterfaceAIp

`func (o *Ipv6PointToPointSubnetAllocation) GetInterfaceAIp() string`

GetInterfaceAIp returns the InterfaceAIp field if non-nil, zero value otherwise.

### GetInterfaceAIpOk

`func (o *Ipv6PointToPointSubnetAllocation) GetInterfaceAIpOk() (*string, bool)`

GetInterfaceAIpOk returns a tuple with the InterfaceAIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAIp

`func (o *Ipv6PointToPointSubnetAllocation) SetInterfaceAIp(v string)`

SetInterfaceAIp sets InterfaceAIp field to given value.


### GetInterfaceBIp

`func (o *Ipv6PointToPointSubnetAllocation) GetInterfaceBIp() string`

GetInterfaceBIp returns the InterfaceBIp field if non-nil, zero value otherwise.

### GetInterfaceBIpOk

`func (o *Ipv6PointToPointSubnetAllocation) GetInterfaceBIpOk() (*string, bool)`

GetInterfaceBIpOk returns a tuple with the InterfaceBIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceBIp

`func (o *Ipv6PointToPointSubnetAllocation) SetInterfaceBIp(v string)`

SetInterfaceBIp sets InterfaceBIp field to given value.


### GetStatus

`func (o *Ipv6PointToPointSubnetAllocation) GetStatus() ResourceAllocationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ipv6PointToPointSubnetAllocation) GetStatusOk() (*ResourceAllocationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ipv6PointToPointSubnetAllocation) SetStatus(v ResourceAllocationStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


