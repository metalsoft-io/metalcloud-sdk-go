# Ipv4PointToPointSubnetAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NetworkAddress** | **string** | Network address of the /31 (e.g. \&quot;10.0.0.0\&quot;). | 
**PrefixLength** | **int32** |  | 
**InterfaceAIp** | **string** | IP assigned to interface A. | 
**InterfaceBIp** | **string** | IP assigned to interface B. | 
**Status** | [**ResourceAllocationStatus**](ResourceAllocationStatus.md) | Allocation lifecycle: &#39;allocated&#39; once the deploy materializes the subnet, &#39;deleting&#39; while staged for removal. | 

## Methods

### NewIpv4PointToPointSubnetAllocation

`func NewIpv4PointToPointSubnetAllocation(id int64, networkAddress string, prefixLength int32, interfaceAIp string, interfaceBIp string, status ResourceAllocationStatus, ) *Ipv4PointToPointSubnetAllocation`

NewIpv4PointToPointSubnetAllocation instantiates a new Ipv4PointToPointSubnetAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4PointToPointSubnetAllocationWithDefaults

`func NewIpv4PointToPointSubnetAllocationWithDefaults() *Ipv4PointToPointSubnetAllocation`

NewIpv4PointToPointSubnetAllocationWithDefaults instantiates a new Ipv4PointToPointSubnetAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ipv4PointToPointSubnetAllocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ipv4PointToPointSubnetAllocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ipv4PointToPointSubnetAllocation) SetId(v int64)`

SetId sets Id field to given value.


### GetNetworkAddress

`func (o *Ipv4PointToPointSubnetAllocation) GetNetworkAddress() string`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *Ipv4PointToPointSubnetAllocation) GetNetworkAddressOk() (*string, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *Ipv4PointToPointSubnetAllocation) SetNetworkAddress(v string)`

SetNetworkAddress sets NetworkAddress field to given value.


### GetPrefixLength

`func (o *Ipv4PointToPointSubnetAllocation) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Ipv4PointToPointSubnetAllocation) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Ipv4PointToPointSubnetAllocation) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetInterfaceAIp

`func (o *Ipv4PointToPointSubnetAllocation) GetInterfaceAIp() string`

GetInterfaceAIp returns the InterfaceAIp field if non-nil, zero value otherwise.

### GetInterfaceAIpOk

`func (o *Ipv4PointToPointSubnetAllocation) GetInterfaceAIpOk() (*string, bool)`

GetInterfaceAIpOk returns a tuple with the InterfaceAIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAIp

`func (o *Ipv4PointToPointSubnetAllocation) SetInterfaceAIp(v string)`

SetInterfaceAIp sets InterfaceAIp field to given value.


### GetInterfaceBIp

`func (o *Ipv4PointToPointSubnetAllocation) GetInterfaceBIp() string`

GetInterfaceBIp returns the InterfaceBIp field if non-nil, zero value otherwise.

### GetInterfaceBIpOk

`func (o *Ipv4PointToPointSubnetAllocation) GetInterfaceBIpOk() (*string, bool)`

GetInterfaceBIpOk returns a tuple with the InterfaceBIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceBIp

`func (o *Ipv4PointToPointSubnetAllocation) SetInterfaceBIp(v string)`

SetInterfaceBIp sets InterfaceBIp field to given value.


### GetStatus

`func (o *Ipv4PointToPointSubnetAllocation) GetStatus() ResourceAllocationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ipv4PointToPointSubnetAllocation) GetStatusOk() (*ResourceAllocationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ipv4PointToPointSubnetAllocation) SetStatus(v ResourceAllocationStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


