# NetworkDeviceLag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of the LAG record | 
**Name** | **string** | The name of the LAG | 
**LagId** | Pointer to **int64** | The LAG identifier on the device | [optional] 
**IsDistributed** | **float32** | Whether the LAG is distributed across redundancy-group members | 
**NetworkDeviceRedundancyGroupId** | Pointer to **int64** | The ID of the network device redundancy group, if the LAG belongs to one | [optional] 

## Methods

### NewNetworkDeviceLag

`func NewNetworkDeviceLag(id int64, name string, isDistributed float32, ) *NetworkDeviceLag`

NewNetworkDeviceLag instantiates a new NetworkDeviceLag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceLagWithDefaults

`func NewNetworkDeviceLagWithDefaults() *NetworkDeviceLag`

NewNetworkDeviceLagWithDefaults instantiates a new NetworkDeviceLag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDeviceLag) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceLag) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceLag) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *NetworkDeviceLag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDeviceLag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDeviceLag) SetName(v string)`

SetName sets Name field to given value.


### GetLagId

`func (o *NetworkDeviceLag) GetLagId() int64`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *NetworkDeviceLag) GetLagIdOk() (*int64, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *NetworkDeviceLag) SetLagId(v int64)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *NetworkDeviceLag) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetIsDistributed

`func (o *NetworkDeviceLag) GetIsDistributed() float32`

GetIsDistributed returns the IsDistributed field if non-nil, zero value otherwise.

### GetIsDistributedOk

`func (o *NetworkDeviceLag) GetIsDistributedOk() (*float32, bool)`

GetIsDistributedOk returns a tuple with the IsDistributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDistributed

`func (o *NetworkDeviceLag) SetIsDistributed(v float32)`

SetIsDistributed sets IsDistributed field to given value.


### GetNetworkDeviceRedundancyGroupId

`func (o *NetworkDeviceLag) GetNetworkDeviceRedundancyGroupId() int64`

GetNetworkDeviceRedundancyGroupId returns the NetworkDeviceRedundancyGroupId field if non-nil, zero value otherwise.

### GetNetworkDeviceRedundancyGroupIdOk

`func (o *NetworkDeviceLag) GetNetworkDeviceRedundancyGroupIdOk() (*int64, bool)`

GetNetworkDeviceRedundancyGroupIdOk returns a tuple with the NetworkDeviceRedundancyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceRedundancyGroupId

`func (o *NetworkDeviceLag) SetNetworkDeviceRedundancyGroupId(v int64)`

SetNetworkDeviceRedundancyGroupId sets NetworkDeviceRedundancyGroupId field to given value.

### HasNetworkDeviceRedundancyGroupId

`func (o *NetworkDeviceLag) HasNetworkDeviceRedundancyGroupId() bool`

HasNetworkDeviceRedundancyGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


