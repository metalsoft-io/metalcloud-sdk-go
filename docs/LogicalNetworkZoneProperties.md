# LogicalNetworkZoneProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zones** | [**[]ZoneAllocation**](ZoneAllocation.md) |  | 
**ZoneAllocationStrategies** | [**[]ZoneAllocationStrategy**](ZoneAllocationStrategy.md) |  | 

## Methods

### NewLogicalNetworkZoneProperties

`func NewLogicalNetworkZoneProperties(zones []ZoneAllocation, zoneAllocationStrategies []ZoneAllocationStrategy, ) *LogicalNetworkZoneProperties`

NewLogicalNetworkZoneProperties instantiates a new LogicalNetworkZoneProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkZonePropertiesWithDefaults

`func NewLogicalNetworkZonePropertiesWithDefaults() *LogicalNetworkZoneProperties`

NewLogicalNetworkZonePropertiesWithDefaults instantiates a new LogicalNetworkZoneProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZones

`func (o *LogicalNetworkZoneProperties) GetZones() []ZoneAllocation`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *LogicalNetworkZoneProperties) GetZonesOk() (*[]ZoneAllocation, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *LogicalNetworkZoneProperties) SetZones(v []ZoneAllocation)`

SetZones sets Zones field to given value.


### GetZoneAllocationStrategies

`func (o *LogicalNetworkZoneProperties) GetZoneAllocationStrategies() []ZoneAllocationStrategy`

GetZoneAllocationStrategies returns the ZoneAllocationStrategies field if non-nil, zero value otherwise.

### GetZoneAllocationStrategiesOk

`func (o *LogicalNetworkZoneProperties) GetZoneAllocationStrategiesOk() (*[]ZoneAllocationStrategy, bool)`

GetZoneAllocationStrategiesOk returns a tuple with the ZoneAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllocationStrategies

`func (o *LogicalNetworkZoneProperties) SetZoneAllocationStrategies(v []ZoneAllocationStrategy)`

SetZoneAllocationStrategies sets ZoneAllocationStrategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


