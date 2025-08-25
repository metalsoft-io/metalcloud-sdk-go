# ServerInstanceAllocatedNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondNetworkConfigs** | Pointer to [**map[string]ServerInstanceBondNetworkConfig**](ServerInstanceBondNetworkConfig.md) | The list of bond network configurations indexed by LAG id. | [optional] 
**PhysicalInterfaceNetworkConfigs** | Pointer to [**map[string]ServerInstancePhysicalNetworkConfig**](ServerInstancePhysicalNetworkConfig.md) | The list of physical network configurations indexed by physical interface id. | [optional] 
**TrunkNetworkConfigs** | Pointer to [**map[string]ServerInstanceTrunkNetworkConfig**](ServerInstanceTrunkNetworkConfig.md) | The list of trunk network configurations indexed by VLAN id. | [optional] 

## Methods

### NewServerInstanceAllocatedNetworkConfig

`func NewServerInstanceAllocatedNetworkConfig() *ServerInstanceAllocatedNetworkConfig`

NewServerInstanceAllocatedNetworkConfig instantiates a new ServerInstanceAllocatedNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceAllocatedNetworkConfigWithDefaults

`func NewServerInstanceAllocatedNetworkConfigWithDefaults() *ServerInstanceAllocatedNetworkConfig`

NewServerInstanceAllocatedNetworkConfigWithDefaults instantiates a new ServerInstanceAllocatedNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) GetBondNetworkConfigs() map[string]ServerInstanceBondNetworkConfig`

GetBondNetworkConfigs returns the BondNetworkConfigs field if non-nil, zero value otherwise.

### GetBondNetworkConfigsOk

`func (o *ServerInstanceAllocatedNetworkConfig) GetBondNetworkConfigsOk() (*map[string]ServerInstanceBondNetworkConfig, bool)`

GetBondNetworkConfigsOk returns a tuple with the BondNetworkConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) SetBondNetworkConfigs(v map[string]ServerInstanceBondNetworkConfig)`

SetBondNetworkConfigs sets BondNetworkConfigs field to given value.

### HasBondNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) HasBondNetworkConfigs() bool`

HasBondNetworkConfigs returns a boolean if a field has been set.

### GetPhysicalInterfaceNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) GetPhysicalInterfaceNetworkConfigs() map[string]ServerInstancePhysicalNetworkConfig`

GetPhysicalInterfaceNetworkConfigs returns the PhysicalInterfaceNetworkConfigs field if non-nil, zero value otherwise.

### GetPhysicalInterfaceNetworkConfigsOk

`func (o *ServerInstanceAllocatedNetworkConfig) GetPhysicalInterfaceNetworkConfigsOk() (*map[string]ServerInstancePhysicalNetworkConfig, bool)`

GetPhysicalInterfaceNetworkConfigsOk returns a tuple with the PhysicalInterfaceNetworkConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalInterfaceNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) SetPhysicalInterfaceNetworkConfigs(v map[string]ServerInstancePhysicalNetworkConfig)`

SetPhysicalInterfaceNetworkConfigs sets PhysicalInterfaceNetworkConfigs field to given value.

### HasPhysicalInterfaceNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) HasPhysicalInterfaceNetworkConfigs() bool`

HasPhysicalInterfaceNetworkConfigs returns a boolean if a field has been set.

### GetTrunkNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) GetTrunkNetworkConfigs() map[string]ServerInstanceTrunkNetworkConfig`

GetTrunkNetworkConfigs returns the TrunkNetworkConfigs field if non-nil, zero value otherwise.

### GetTrunkNetworkConfigsOk

`func (o *ServerInstanceAllocatedNetworkConfig) GetTrunkNetworkConfigsOk() (*map[string]ServerInstanceTrunkNetworkConfig, bool)`

GetTrunkNetworkConfigsOk returns a tuple with the TrunkNetworkConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) SetTrunkNetworkConfigs(v map[string]ServerInstanceTrunkNetworkConfig)`

SetTrunkNetworkConfigs sets TrunkNetworkConfigs field to given value.

### HasTrunkNetworkConfigs

`func (o *ServerInstanceAllocatedNetworkConfig) HasTrunkNetworkConfigs() bool`

HasTrunkNetworkConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


