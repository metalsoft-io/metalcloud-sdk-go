# UpdateVMInstanceGroupNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagged** | Pointer to **bool** | Whether the logical network is tagged. | [optional] 
**AccessMode** | Pointer to [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | [optional] 
**Mtu** | Pointer to **int32** | The MTU of the logical network | [optional] 
**ProvidesDefaultRoute** | Pointer to **bool** | Whether the logical network provides a default route | [optional] [default to false]
**DisableAutoIpAllocation** | Pointer to **bool** | Disable automatic IP allocation for IPv4 addresses on this network connection | [optional] [default to false]
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Dns** | Pointer to [**NullableDnsRecordsEndpointGroupLogicalNetwork**](DnsRecordsEndpointGroupLogicalNetwork.md) | DNS records configuration for the server instance group. | [optional] 

## Methods

### NewUpdateVMInstanceGroupNetworkConnection

`func NewUpdateVMInstanceGroupNetworkConnection() *UpdateVMInstanceGroupNetworkConnection`

NewUpdateVMInstanceGroupNetworkConnection instantiates a new UpdateVMInstanceGroupNetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMInstanceGroupNetworkConnectionWithDefaults

`func NewUpdateVMInstanceGroupNetworkConnectionWithDefaults() *UpdateVMInstanceGroupNetworkConnection`

NewUpdateVMInstanceGroupNetworkConnectionWithDefaults instantiates a new UpdateVMInstanceGroupNetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagged

`func (o *UpdateVMInstanceGroupNetworkConnection) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *UpdateVMInstanceGroupNetworkConnection) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *UpdateVMInstanceGroupNetworkConnection) SetTagged(v bool)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *UpdateVMInstanceGroupNetworkConnection) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetAccessMode

`func (o *UpdateVMInstanceGroupNetworkConnection) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *UpdateVMInstanceGroupNetworkConnection) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *UpdateVMInstanceGroupNetworkConnection) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *UpdateVMInstanceGroupNetworkConnection) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetMtu

`func (o *UpdateVMInstanceGroupNetworkConnection) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *UpdateVMInstanceGroupNetworkConnection) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *UpdateVMInstanceGroupNetworkConnection) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *UpdateVMInstanceGroupNetworkConnection) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetProvidesDefaultRoute

`func (o *UpdateVMInstanceGroupNetworkConnection) GetProvidesDefaultRoute() bool`

GetProvidesDefaultRoute returns the ProvidesDefaultRoute field if non-nil, zero value otherwise.

### GetProvidesDefaultRouteOk

`func (o *UpdateVMInstanceGroupNetworkConnection) GetProvidesDefaultRouteOk() (*bool, bool)`

GetProvidesDefaultRouteOk returns a tuple with the ProvidesDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesDefaultRoute

`func (o *UpdateVMInstanceGroupNetworkConnection) SetProvidesDefaultRoute(v bool)`

SetProvidesDefaultRoute sets ProvidesDefaultRoute field to given value.

### HasProvidesDefaultRoute

`func (o *UpdateVMInstanceGroupNetworkConnection) HasProvidesDefaultRoute() bool`

HasProvidesDefaultRoute returns a boolean if a field has been set.

### GetDisableAutoIpAllocation

`func (o *UpdateVMInstanceGroupNetworkConnection) GetDisableAutoIpAllocation() bool`

GetDisableAutoIpAllocation returns the DisableAutoIpAllocation field if non-nil, zero value otherwise.

### GetDisableAutoIpAllocationOk

`func (o *UpdateVMInstanceGroupNetworkConnection) GetDisableAutoIpAllocationOk() (*bool, bool)`

GetDisableAutoIpAllocationOk returns a tuple with the DisableAutoIpAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoIpAllocation

`func (o *UpdateVMInstanceGroupNetworkConnection) SetDisableAutoIpAllocation(v bool)`

SetDisableAutoIpAllocation sets DisableAutoIpAllocation field to given value.

### HasDisableAutoIpAllocation

`func (o *UpdateVMInstanceGroupNetworkConnection) HasDisableAutoIpAllocation() bool`

HasDisableAutoIpAllocation returns a boolean if a field has been set.

### GetRedundancy

`func (o *UpdateVMInstanceGroupNetworkConnection) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *UpdateVMInstanceGroupNetworkConnection) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *UpdateVMInstanceGroupNetworkConnection) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *UpdateVMInstanceGroupNetworkConnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *UpdateVMInstanceGroupNetworkConnection) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *UpdateVMInstanceGroupNetworkConnection) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetDns

`func (o *UpdateVMInstanceGroupNetworkConnection) GetDns() DnsRecordsEndpointGroupLogicalNetwork`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *UpdateVMInstanceGroupNetworkConnection) GetDnsOk() (*DnsRecordsEndpointGroupLogicalNetwork, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *UpdateVMInstanceGroupNetworkConnection) SetDns(v DnsRecordsEndpointGroupLogicalNetwork)`

SetDns sets Dns field to given value.

### HasDns

`func (o *UpdateVMInstanceGroupNetworkConnection) HasDns() bool`

HasDns returns a boolean if a field has been set.

### SetDnsNil

`func (o *UpdateVMInstanceGroupNetworkConnection) SetDnsNil(b bool)`

 SetDnsNil sets the value for Dns to be an explicit nil

### UnsetDns
`func (o *UpdateVMInstanceGroupNetworkConnection) UnsetDns()`

UnsetDns ensures that no value is present for Dns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


