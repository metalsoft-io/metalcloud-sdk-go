# LogicalNetworkPkeyProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pkeys** | [**[]PkeyAllocation**](PkeyAllocation.md) |  | 
**PkeyAllocationStrategies** | [**[]PkeyAllocationStrategy**](PkeyAllocationStrategy.md) |  | 

## Methods

### NewLogicalNetworkPkeyProperties

`func NewLogicalNetworkPkeyProperties(pkeys []PkeyAllocation, pkeyAllocationStrategies []PkeyAllocationStrategy, ) *LogicalNetworkPkeyProperties`

NewLogicalNetworkPkeyProperties instantiates a new LogicalNetworkPkeyProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkPkeyPropertiesWithDefaults

`func NewLogicalNetworkPkeyPropertiesWithDefaults() *LogicalNetworkPkeyProperties`

NewLogicalNetworkPkeyPropertiesWithDefaults instantiates a new LogicalNetworkPkeyProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkeys

`func (o *LogicalNetworkPkeyProperties) GetPkeys() []PkeyAllocation`

GetPkeys returns the Pkeys field if non-nil, zero value otherwise.

### GetPkeysOk

`func (o *LogicalNetworkPkeyProperties) GetPkeysOk() (*[]PkeyAllocation, bool)`

GetPkeysOk returns a tuple with the Pkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkeys

`func (o *LogicalNetworkPkeyProperties) SetPkeys(v []PkeyAllocation)`

SetPkeys sets Pkeys field to given value.


### GetPkeyAllocationStrategies

`func (o *LogicalNetworkPkeyProperties) GetPkeyAllocationStrategies() []PkeyAllocationStrategy`

GetPkeyAllocationStrategies returns the PkeyAllocationStrategies field if non-nil, zero value otherwise.

### GetPkeyAllocationStrategiesOk

`func (o *LogicalNetworkPkeyProperties) GetPkeyAllocationStrategiesOk() (*[]PkeyAllocationStrategy, bool)`

GetPkeyAllocationStrategiesOk returns a tuple with the PkeyAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkeyAllocationStrategies

`func (o *LogicalNetworkPkeyProperties) SetPkeyAllocationStrategies(v []PkeyAllocationStrategy)`

SetPkeyAllocationStrategies sets PkeyAllocationStrategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


