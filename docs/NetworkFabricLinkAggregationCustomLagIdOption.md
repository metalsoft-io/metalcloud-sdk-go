# NetworkFabricLinkAggregationCustomLagIdOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | **int64** | Unique identifier for the network fabric link | 
**NetworkDeviceAInterfaceLagId** | Pointer to **int64** | Custom LAG ID for the network device A interface. Omit it to let the fabric ranges allocate one, or when side A of the link has no network device interface. | [optional] 
**NetworkDeviceBInterfaceLagId** | Pointer to **int64** | Custom LAG ID for the network device B interface. Omit it to let the fabric ranges allocate one. A link to an external system has that system as side B and no network device interface there, so the ID must be omitted for those links — MetalSoft renders no configuration for an external system. | [optional] 

## Methods

### NewNetworkFabricLinkAggregationCustomLagIdOption

`func NewNetworkFabricLinkAggregationCustomLagIdOption(linkId int64, ) *NetworkFabricLinkAggregationCustomLagIdOption`

NewNetworkFabricLinkAggregationCustomLagIdOption instantiates a new NetworkFabricLinkAggregationCustomLagIdOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricLinkAggregationCustomLagIdOptionWithDefaults

`func NewNetworkFabricLinkAggregationCustomLagIdOptionWithDefaults() *NetworkFabricLinkAggregationCustomLagIdOption`

NewNetworkFabricLinkAggregationCustomLagIdOptionWithDefaults instantiates a new NetworkFabricLinkAggregationCustomLagIdOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) GetLinkId() int64`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) GetLinkIdOk() (*int64, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) SetLinkId(v int64)`

SetLinkId sets LinkId field to given value.


### GetNetworkDeviceAInterfaceLagId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) GetNetworkDeviceAInterfaceLagId() int64`

GetNetworkDeviceAInterfaceLagId returns the NetworkDeviceAInterfaceLagId field if non-nil, zero value otherwise.

### GetNetworkDeviceAInterfaceLagIdOk

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) GetNetworkDeviceAInterfaceLagIdOk() (*int64, bool)`

GetNetworkDeviceAInterfaceLagIdOk returns a tuple with the NetworkDeviceAInterfaceLagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceAInterfaceLagId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) SetNetworkDeviceAInterfaceLagId(v int64)`

SetNetworkDeviceAInterfaceLagId sets NetworkDeviceAInterfaceLagId field to given value.

### HasNetworkDeviceAInterfaceLagId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) HasNetworkDeviceAInterfaceLagId() bool`

HasNetworkDeviceAInterfaceLagId returns a boolean if a field has been set.

### GetNetworkDeviceBInterfaceLagId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) GetNetworkDeviceBInterfaceLagId() int64`

GetNetworkDeviceBInterfaceLagId returns the NetworkDeviceBInterfaceLagId field if non-nil, zero value otherwise.

### GetNetworkDeviceBInterfaceLagIdOk

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) GetNetworkDeviceBInterfaceLagIdOk() (*int64, bool)`

GetNetworkDeviceBInterfaceLagIdOk returns a tuple with the NetworkDeviceBInterfaceLagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceBInterfaceLagId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) SetNetworkDeviceBInterfaceLagId(v int64)`

SetNetworkDeviceBInterfaceLagId sets NetworkDeviceBInterfaceLagId field to given value.

### HasNetworkDeviceBInterfaceLagId

`func (o *NetworkFabricLinkAggregationCustomLagIdOption) HasNetworkDeviceBInterfaceLagId() bool`

HasNetworkDeviceBInterfaceLagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


