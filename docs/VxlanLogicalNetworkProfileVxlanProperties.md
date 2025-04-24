# VxlanLogicalNetworkProfileVxlanProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vnis** | [**[]VniAllocation**](VniAllocation.md) |  | 
**VniAllocationStrategies** | [**[]VniAllocationStrategy**](VniAllocationStrategy.md) |  | 

## Methods

### NewVxlanLogicalNetworkProfileVxlanProperties

`func NewVxlanLogicalNetworkProfileVxlanProperties(vnis []VniAllocation, vniAllocationStrategies []VniAllocationStrategy, ) *VxlanLogicalNetworkProfileVxlanProperties`

NewVxlanLogicalNetworkProfileVxlanProperties instantiates a new VxlanLogicalNetworkProfileVxlanProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVxlanLogicalNetworkProfileVxlanPropertiesWithDefaults

`func NewVxlanLogicalNetworkProfileVxlanPropertiesWithDefaults() *VxlanLogicalNetworkProfileVxlanProperties`

NewVxlanLogicalNetworkProfileVxlanPropertiesWithDefaults instantiates a new VxlanLogicalNetworkProfileVxlanProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVnis

`func (o *VxlanLogicalNetworkProfileVxlanProperties) GetVnis() []VniAllocation`

GetVnis returns the Vnis field if non-nil, zero value otherwise.

### GetVnisOk

`func (o *VxlanLogicalNetworkProfileVxlanProperties) GetVnisOk() (*[]VniAllocation, bool)`

GetVnisOk returns a tuple with the Vnis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnis

`func (o *VxlanLogicalNetworkProfileVxlanProperties) SetVnis(v []VniAllocation)`

SetVnis sets Vnis field to given value.


### GetVniAllocationStrategies

`func (o *VxlanLogicalNetworkProfileVxlanProperties) GetVniAllocationStrategies() []VniAllocationStrategy`

GetVniAllocationStrategies returns the VniAllocationStrategies field if non-nil, zero value otherwise.

### GetVniAllocationStrategiesOk

`func (o *VxlanLogicalNetworkProfileVxlanProperties) GetVniAllocationStrategiesOk() (*[]VniAllocationStrategy, bool)`

GetVniAllocationStrategiesOk returns a tuple with the VniAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniAllocationStrategies

`func (o *VxlanLogicalNetworkProfileVxlanProperties) SetVniAllocationStrategies(v []VniAllocationStrategy)`

SetVniAllocationStrategies sets VniAllocationStrategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


