# CreateContainerInstanceGroupNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalNetworkId** | **string** | The logical network ID. | 
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | 
**Mtu** | Pointer to **int32** | The MTU of the logical network | [optional] 
**InterfaceCount** | Pointer to **NullableInt32** | Minimum number of interfaces for this connection. Only configurable when the logical network is l3Only. Null/unset means no minimum. | [optional] 
**ProvidesDefaultRoute** | Pointer to **bool** | Whether the logical network provides a default route | [optional] [default to false]
**DisableAutoIpAllocation** | Pointer to **bool** | Disable automatic IP allocation for IPv4 addresses on this network connection | [optional] [default to false]
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Dns** | Pointer to [**NullableDnsRecordsEndpointGroupLogicalNetwork**](DnsRecordsEndpointGroupLogicalNetwork.md) | DNS records that are supposed to be provisioned for the server instance group. | [optional] 

## Methods

### NewCreateContainerInstanceGroupNetworkConnection

`func NewCreateContainerInstanceGroupNetworkConnection(logicalNetworkId string, tagged bool, accessMode NetworkEndpointGroupAllowedAccessMode, ) *CreateContainerInstanceGroupNetworkConnection`

NewCreateContainerInstanceGroupNetworkConnection instantiates a new CreateContainerInstanceGroupNetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerInstanceGroupNetworkConnectionWithDefaults

`func NewCreateContainerInstanceGroupNetworkConnectionWithDefaults() *CreateContainerInstanceGroupNetworkConnection`

NewCreateContainerInstanceGroupNetworkConnectionWithDefaults instantiates a new CreateContainerInstanceGroupNetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalNetworkId

`func (o *CreateContainerInstanceGroupNetworkConnection) GetLogicalNetworkId() string`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetLogicalNetworkIdOk() (*string, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *CreateContainerInstanceGroupNetworkConnection) SetLogicalNetworkId(v string)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetTagged

`func (o *CreateContainerInstanceGroupNetworkConnection) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *CreateContainerInstanceGroupNetworkConnection) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *CreateContainerInstanceGroupNetworkConnection) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CreateContainerInstanceGroupNetworkConnection) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.


### GetMtu

`func (o *CreateContainerInstanceGroupNetworkConnection) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateContainerInstanceGroupNetworkConnection) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateContainerInstanceGroupNetworkConnection) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetInterfaceCount

`func (o *CreateContainerInstanceGroupNetworkConnection) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *CreateContainerInstanceGroupNetworkConnection) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.

### HasInterfaceCount

`func (o *CreateContainerInstanceGroupNetworkConnection) HasInterfaceCount() bool`

HasInterfaceCount returns a boolean if a field has been set.

### SetInterfaceCountNil

`func (o *CreateContainerInstanceGroupNetworkConnection) SetInterfaceCountNil(b bool)`

 SetInterfaceCountNil sets the value for InterfaceCount to be an explicit nil

### UnsetInterfaceCount
`func (o *CreateContainerInstanceGroupNetworkConnection) UnsetInterfaceCount()`

UnsetInterfaceCount ensures that no value is present for InterfaceCount, not even an explicit nil
### GetProvidesDefaultRoute

`func (o *CreateContainerInstanceGroupNetworkConnection) GetProvidesDefaultRoute() bool`

GetProvidesDefaultRoute returns the ProvidesDefaultRoute field if non-nil, zero value otherwise.

### GetProvidesDefaultRouteOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetProvidesDefaultRouteOk() (*bool, bool)`

GetProvidesDefaultRouteOk returns a tuple with the ProvidesDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesDefaultRoute

`func (o *CreateContainerInstanceGroupNetworkConnection) SetProvidesDefaultRoute(v bool)`

SetProvidesDefaultRoute sets ProvidesDefaultRoute field to given value.

### HasProvidesDefaultRoute

`func (o *CreateContainerInstanceGroupNetworkConnection) HasProvidesDefaultRoute() bool`

HasProvidesDefaultRoute returns a boolean if a field has been set.

### GetDisableAutoIpAllocation

`func (o *CreateContainerInstanceGroupNetworkConnection) GetDisableAutoIpAllocation() bool`

GetDisableAutoIpAllocation returns the DisableAutoIpAllocation field if non-nil, zero value otherwise.

### GetDisableAutoIpAllocationOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetDisableAutoIpAllocationOk() (*bool, bool)`

GetDisableAutoIpAllocationOk returns a tuple with the DisableAutoIpAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoIpAllocation

`func (o *CreateContainerInstanceGroupNetworkConnection) SetDisableAutoIpAllocation(v bool)`

SetDisableAutoIpAllocation sets DisableAutoIpAllocation field to given value.

### HasDisableAutoIpAllocation

`func (o *CreateContainerInstanceGroupNetworkConnection) HasDisableAutoIpAllocation() bool`

HasDisableAutoIpAllocation returns a boolean if a field has been set.

### GetRedundancy

`func (o *CreateContainerInstanceGroupNetworkConnection) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CreateContainerInstanceGroupNetworkConnection) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *CreateContainerInstanceGroupNetworkConnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *CreateContainerInstanceGroupNetworkConnection) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *CreateContainerInstanceGroupNetworkConnection) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetDns

`func (o *CreateContainerInstanceGroupNetworkConnection) GetDns() DnsRecordsEndpointGroupLogicalNetwork`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *CreateContainerInstanceGroupNetworkConnection) GetDnsOk() (*DnsRecordsEndpointGroupLogicalNetwork, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *CreateContainerInstanceGroupNetworkConnection) SetDns(v DnsRecordsEndpointGroupLogicalNetwork)`

SetDns sets Dns field to given value.

### HasDns

`func (o *CreateContainerInstanceGroupNetworkConnection) HasDns() bool`

HasDns returns a boolean if a field has been set.

### SetDnsNil

`func (o *CreateContainerInstanceGroupNetworkConnection) SetDnsNil(b bool)`

 SetDnsNil sets the value for Dns to be an explicit nil

### UnsetDns
`func (o *CreateContainerInstanceGroupNetworkConnection) UnsetDns()`

UnsetDns ensures that no value is present for Dns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


