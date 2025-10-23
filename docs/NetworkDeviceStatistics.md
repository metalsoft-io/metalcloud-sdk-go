# NetworkDeviceStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceCount** | **int32** | Total count of network devices | 
**PortCount** | **int32** | Total count of network device ports | 

## Methods

### NewNetworkDeviceStatistics

`func NewNetworkDeviceStatistics(deviceCount int32, portCount int32, ) *NetworkDeviceStatistics`

NewNetworkDeviceStatistics instantiates a new NetworkDeviceStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceStatisticsWithDefaults

`func NewNetworkDeviceStatisticsWithDefaults() *NetworkDeviceStatistics`

NewNetworkDeviceStatisticsWithDefaults instantiates a new NetworkDeviceStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceCount

`func (o *NetworkDeviceStatistics) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *NetworkDeviceStatistics) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *NetworkDeviceStatistics) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.


### GetPortCount

`func (o *NetworkDeviceStatistics) GetPortCount() int32`

GetPortCount returns the PortCount field if non-nil, zero value otherwise.

### GetPortCountOk

`func (o *NetworkDeviceStatistics) GetPortCountOk() (*int32, bool)`

GetPortCountOk returns a tuple with the PortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCount

`func (o *NetworkDeviceStatistics) SetPortCount(v int32)`

SetPortCount sets PortCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


