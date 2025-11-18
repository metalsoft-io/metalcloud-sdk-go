# VlanLogicalNetworkVlanProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlans** | [**[]VlanAllocation**](VlanAllocation.md) |  | 
**VlanAllocationStrategies** | [**[]VlanAllocationStrategy1**](VlanAllocationStrategy1.md) |  | 

## Methods

### NewVlanLogicalNetworkVlanProperties

`func NewVlanLogicalNetworkVlanProperties(vlans []VlanAllocation, vlanAllocationStrategies []VlanAllocationStrategy1, ) *VlanLogicalNetworkVlanProperties`

NewVlanLogicalNetworkVlanProperties instantiates a new VlanLogicalNetworkVlanProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanLogicalNetworkVlanPropertiesWithDefaults

`func NewVlanLogicalNetworkVlanPropertiesWithDefaults() *VlanLogicalNetworkVlanProperties`

NewVlanLogicalNetworkVlanPropertiesWithDefaults instantiates a new VlanLogicalNetworkVlanProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlans

`func (o *VlanLogicalNetworkVlanProperties) GetVlans() []VlanAllocation`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *VlanLogicalNetworkVlanProperties) GetVlansOk() (*[]VlanAllocation, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *VlanLogicalNetworkVlanProperties) SetVlans(v []VlanAllocation)`

SetVlans sets Vlans field to given value.


### GetVlanAllocationStrategies

`func (o *VlanLogicalNetworkVlanProperties) GetVlanAllocationStrategies() []VlanAllocationStrategy1`

GetVlanAllocationStrategies returns the VlanAllocationStrategies field if non-nil, zero value otherwise.

### GetVlanAllocationStrategiesOk

`func (o *VlanLogicalNetworkVlanProperties) GetVlanAllocationStrategiesOk() (*[]VlanAllocationStrategy1, bool)`

GetVlanAllocationStrategiesOk returns a tuple with the VlanAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanAllocationStrategies

`func (o *VlanLogicalNetworkVlanProperties) SetVlanAllocationStrategies(v []VlanAllocationStrategy1)`

SetVlanAllocationStrategies sets VlanAllocationStrategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


