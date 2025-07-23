# CreateEndpointInstanceGroupNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalNetworkId** | **string** | The logical network ID. | 
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | 
**Mtu** | Pointer to **int32** | The MTU of the logical network | [optional] 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Dns** | Pointer to [**NullableDnsRecordsEndpointGroupLogicalNetwork**](DnsRecordsEndpointGroupLogicalNetwork.md) | DNS records that are supposed to be provisioned for the server instance group. | [optional] 

## Methods

### NewCreateEndpointInstanceGroupNetworkConnection

`func NewCreateEndpointInstanceGroupNetworkConnection(logicalNetworkId string, tagged bool, accessMode NetworkEndpointGroupAllowedAccessMode, ) *CreateEndpointInstanceGroupNetworkConnection`

NewCreateEndpointInstanceGroupNetworkConnection instantiates a new CreateEndpointInstanceGroupNetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointInstanceGroupNetworkConnectionWithDefaults

`func NewCreateEndpointInstanceGroupNetworkConnectionWithDefaults() *CreateEndpointInstanceGroupNetworkConnection`

NewCreateEndpointInstanceGroupNetworkConnectionWithDefaults instantiates a new CreateEndpointInstanceGroupNetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalNetworkId

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetLogicalNetworkId() string`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetLogicalNetworkIdOk() (*string, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetLogicalNetworkId(v string)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetTagged

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.


### GetMtu

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateEndpointInstanceGroupNetworkConnection) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetRedundancy

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *CreateEndpointInstanceGroupNetworkConnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *CreateEndpointInstanceGroupNetworkConnection) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetDns

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetDns() DnsRecordsEndpointGroupLogicalNetwork`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *CreateEndpointInstanceGroupNetworkConnection) GetDnsOk() (*DnsRecordsEndpointGroupLogicalNetwork, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetDns(v DnsRecordsEndpointGroupLogicalNetwork)`

SetDns sets Dns field to given value.

### HasDns

`func (o *CreateEndpointInstanceGroupNetworkConnection) HasDns() bool`

HasDns returns a boolean if a field has been set.

### SetDnsNil

`func (o *CreateEndpointInstanceGroupNetworkConnection) SetDnsNil(b bool)`

 SetDnsNil sets the value for Dns to be an explicit nil

### UnsetDns
`func (o *CreateEndpointInstanceGroupNetworkConnection) UnsetDns()`

UnsetDns ensures that no value is present for Dns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


