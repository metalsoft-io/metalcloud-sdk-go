# EndpointInstanceGroupNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | 
**Mtu** | Pointer to **int32** | The MTU of the logical network | [optional] 
**ProvidesDefaultRoute** | Pointer to **bool** | Whether the logical network provides a default route | [optional] [default to false]
**DisableAutoIpAllocation** | Pointer to **bool** | Disable automatic IP allocation for IPv4 addresses on this network connection | [optional] [default to false]
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Dns** | Pointer to [**NullableDnsRecordsEndpointGroupLogicalNetwork**](DnsRecordsEndpointGroupLogicalNetwork.md) | DNS records that are supposed to be provisioned for the server instance group. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **string** | The network connection ID. | 

## Methods

### NewEndpointInstanceGroupNetworkConnection

`func NewEndpointInstanceGroupNetworkConnection(tagged bool, accessMode NetworkEndpointGroupAllowedAccessMode, id string, ) *EndpointInstanceGroupNetworkConnection`

NewEndpointInstanceGroupNetworkConnection instantiates a new EndpointInstanceGroupNetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceGroupNetworkConnectionWithDefaults

`func NewEndpointInstanceGroupNetworkConnectionWithDefaults() *EndpointInstanceGroupNetworkConnection`

NewEndpointInstanceGroupNetworkConnectionWithDefaults instantiates a new EndpointInstanceGroupNetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagged

`func (o *EndpointInstanceGroupNetworkConnection) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *EndpointInstanceGroupNetworkConnection) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *EndpointInstanceGroupNetworkConnection) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *EndpointInstanceGroupNetworkConnection) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *EndpointInstanceGroupNetworkConnection) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *EndpointInstanceGroupNetworkConnection) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.


### GetMtu

`func (o *EndpointInstanceGroupNetworkConnection) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *EndpointInstanceGroupNetworkConnection) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *EndpointInstanceGroupNetworkConnection) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *EndpointInstanceGroupNetworkConnection) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetProvidesDefaultRoute

`func (o *EndpointInstanceGroupNetworkConnection) GetProvidesDefaultRoute() bool`

GetProvidesDefaultRoute returns the ProvidesDefaultRoute field if non-nil, zero value otherwise.

### GetProvidesDefaultRouteOk

`func (o *EndpointInstanceGroupNetworkConnection) GetProvidesDefaultRouteOk() (*bool, bool)`

GetProvidesDefaultRouteOk returns a tuple with the ProvidesDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesDefaultRoute

`func (o *EndpointInstanceGroupNetworkConnection) SetProvidesDefaultRoute(v bool)`

SetProvidesDefaultRoute sets ProvidesDefaultRoute field to given value.

### HasProvidesDefaultRoute

`func (o *EndpointInstanceGroupNetworkConnection) HasProvidesDefaultRoute() bool`

HasProvidesDefaultRoute returns a boolean if a field has been set.

### GetDisableAutoIpAllocation

`func (o *EndpointInstanceGroupNetworkConnection) GetDisableAutoIpAllocation() bool`

GetDisableAutoIpAllocation returns the DisableAutoIpAllocation field if non-nil, zero value otherwise.

### GetDisableAutoIpAllocationOk

`func (o *EndpointInstanceGroupNetworkConnection) GetDisableAutoIpAllocationOk() (*bool, bool)`

GetDisableAutoIpAllocationOk returns a tuple with the DisableAutoIpAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoIpAllocation

`func (o *EndpointInstanceGroupNetworkConnection) SetDisableAutoIpAllocation(v bool)`

SetDisableAutoIpAllocation sets DisableAutoIpAllocation field to given value.

### HasDisableAutoIpAllocation

`func (o *EndpointInstanceGroupNetworkConnection) HasDisableAutoIpAllocation() bool`

HasDisableAutoIpAllocation returns a boolean if a field has been set.

### GetRedundancy

`func (o *EndpointInstanceGroupNetworkConnection) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *EndpointInstanceGroupNetworkConnection) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *EndpointInstanceGroupNetworkConnection) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *EndpointInstanceGroupNetworkConnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *EndpointInstanceGroupNetworkConnection) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *EndpointInstanceGroupNetworkConnection) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetDns

`func (o *EndpointInstanceGroupNetworkConnection) GetDns() DnsRecordsEndpointGroupLogicalNetwork`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *EndpointInstanceGroupNetworkConnection) GetDnsOk() (*DnsRecordsEndpointGroupLogicalNetwork, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *EndpointInstanceGroupNetworkConnection) SetDns(v DnsRecordsEndpointGroupLogicalNetwork)`

SetDns sets Dns field to given value.

### HasDns

`func (o *EndpointInstanceGroupNetworkConnection) HasDns() bool`

HasDns returns a boolean if a field has been set.

### SetDnsNil

`func (o *EndpointInstanceGroupNetworkConnection) SetDnsNil(b bool)`

 SetDnsNil sets the value for Dns to be an explicit nil

### UnsetDns
`func (o *EndpointInstanceGroupNetworkConnection) UnsetDns()`

UnsetDns ensures that no value is present for Dns, not even an explicit nil
### GetLinks

`func (o *EndpointInstanceGroupNetworkConnection) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndpointInstanceGroupNetworkConnection) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndpointInstanceGroupNetworkConnection) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EndpointInstanceGroupNetworkConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EndpointInstanceGroupNetworkConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointInstanceGroupNetworkConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointInstanceGroupNetworkConnection) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


