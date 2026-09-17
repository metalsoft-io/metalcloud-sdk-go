# UpdateContainerInstanceGroupNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagged** | Pointer to **bool** | Whether the logical network is tagged. | [optional] 
**AccessMode** | Pointer to [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | [optional] 
**Mtu** | Pointer to **int32** | The MTU of the logical network | [optional] 
**InterfaceCount** | Pointer to **NullableInt32** | Minimum number of interfaces for this connection. Only configurable when the logical network is l3Only. Null/unset means no minimum. | [optional] 
**ProvidesDefaultRoute** | Pointer to **bool** | Whether the logical network provides a default route | [optional] [default to false]
**DisableAutoIpAllocation** | Pointer to **bool** | Disable automatic IP allocation for IPv4 addresses on this network connection | [optional] [default to false]
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Dns** | Pointer to [**NullableDnsRecordsEndpointGroupLogicalNetwork**](DnsRecordsEndpointGroupLogicalNetwork.md) | DNS records configuration for the server instance group. | [optional] 

## Methods

### NewUpdateContainerInstanceGroupNetworkConnection

`func NewUpdateContainerInstanceGroupNetworkConnection() *UpdateContainerInstanceGroupNetworkConnection`

NewUpdateContainerInstanceGroupNetworkConnection instantiates a new UpdateContainerInstanceGroupNetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerInstanceGroupNetworkConnectionWithDefaults

`func NewUpdateContainerInstanceGroupNetworkConnectionWithDefaults() *UpdateContainerInstanceGroupNetworkConnection`

NewUpdateContainerInstanceGroupNetworkConnectionWithDefaults instantiates a new UpdateContainerInstanceGroupNetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagged

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetTagged(v bool)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetAccessMode

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetMtu

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetInterfaceCount

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.

### HasInterfaceCount

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasInterfaceCount() bool`

HasInterfaceCount returns a boolean if a field has been set.

### SetInterfaceCountNil

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetInterfaceCountNil(b bool)`

 SetInterfaceCountNil sets the value for InterfaceCount to be an explicit nil

### UnsetInterfaceCount
`func (o *UpdateContainerInstanceGroupNetworkConnection) UnsetInterfaceCount()`

UnsetInterfaceCount ensures that no value is present for InterfaceCount, not even an explicit nil
### GetProvidesDefaultRoute

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetProvidesDefaultRoute() bool`

GetProvidesDefaultRoute returns the ProvidesDefaultRoute field if non-nil, zero value otherwise.

### GetProvidesDefaultRouteOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetProvidesDefaultRouteOk() (*bool, bool)`

GetProvidesDefaultRouteOk returns a tuple with the ProvidesDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesDefaultRoute

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetProvidesDefaultRoute(v bool)`

SetProvidesDefaultRoute sets ProvidesDefaultRoute field to given value.

### HasProvidesDefaultRoute

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasProvidesDefaultRoute() bool`

HasProvidesDefaultRoute returns a boolean if a field has been set.

### GetDisableAutoIpAllocation

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetDisableAutoIpAllocation() bool`

GetDisableAutoIpAllocation returns the DisableAutoIpAllocation field if non-nil, zero value otherwise.

### GetDisableAutoIpAllocationOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetDisableAutoIpAllocationOk() (*bool, bool)`

GetDisableAutoIpAllocationOk returns a tuple with the DisableAutoIpAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoIpAllocation

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetDisableAutoIpAllocation(v bool)`

SetDisableAutoIpAllocation sets DisableAutoIpAllocation field to given value.

### HasDisableAutoIpAllocation

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasDisableAutoIpAllocation() bool`

HasDisableAutoIpAllocation returns a boolean if a field has been set.

### GetRedundancy

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *UpdateContainerInstanceGroupNetworkConnection) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetDns

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetDns() DnsRecordsEndpointGroupLogicalNetwork`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *UpdateContainerInstanceGroupNetworkConnection) GetDnsOk() (*DnsRecordsEndpointGroupLogicalNetwork, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetDns(v DnsRecordsEndpointGroupLogicalNetwork)`

SetDns sets Dns field to given value.

### HasDns

`func (o *UpdateContainerInstanceGroupNetworkConnection) HasDns() bool`

HasDns returns a boolean if a field has been set.

### SetDnsNil

`func (o *UpdateContainerInstanceGroupNetworkConnection) SetDnsNil(b bool)`

 SetDnsNil sets the value for Dns to be an explicit nil

### UnsetDns
`func (o *UpdateContainerInstanceGroupNetworkConnection) UnsetDns()`

UnsetDns ensures that no value is present for Dns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


