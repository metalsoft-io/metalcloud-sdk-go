# NetworkDevicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortName** | **string** |  | 
**Enabled** | **bool** |  | 
**Active** | **bool** |  | 
**LinkSpeed** | **float32** |  | 
**LinkDuplex** | [**LinkDuplex**](LinkDuplex.md) |  | 
**UtilizationIn** | **float32** |  | 
**UtilizationOut** | **float32** |  | 
**Counters** | [**NetworkDevicePortCounters**](NetworkDevicePortCounters.md) |  | 

## Methods

### NewNetworkDevicePort

`func NewNetworkDevicePort(portName string, enabled bool, active bool, linkSpeed float32, linkDuplex LinkDuplex, utilizationIn float32, utilizationOut float32, counters NetworkDevicePortCounters, ) *NetworkDevicePort`

NewNetworkDevicePort instantiates a new NetworkDevicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDevicePortWithDefaults

`func NewNetworkDevicePortWithDefaults() *NetworkDevicePort`

NewNetworkDevicePortWithDefaults instantiates a new NetworkDevicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortName

`func (o *NetworkDevicePort) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *NetworkDevicePort) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *NetworkDevicePort) SetPortName(v string)`

SetPortName sets PortName field to given value.


### GetEnabled

`func (o *NetworkDevicePort) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkDevicePort) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkDevicePort) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetActive

`func (o *NetworkDevicePort) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkDevicePort) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkDevicePort) SetActive(v bool)`

SetActive sets Active field to given value.


### GetLinkSpeed

`func (o *NetworkDevicePort) GetLinkSpeed() float32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *NetworkDevicePort) GetLinkSpeedOk() (*float32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *NetworkDevicePort) SetLinkSpeed(v float32)`

SetLinkSpeed sets LinkSpeed field to given value.


### GetLinkDuplex

`func (o *NetworkDevicePort) GetLinkDuplex() LinkDuplex`

GetLinkDuplex returns the LinkDuplex field if non-nil, zero value otherwise.

### GetLinkDuplexOk

`func (o *NetworkDevicePort) GetLinkDuplexOk() (*LinkDuplex, bool)`

GetLinkDuplexOk returns a tuple with the LinkDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDuplex

`func (o *NetworkDevicePort) SetLinkDuplex(v LinkDuplex)`

SetLinkDuplex sets LinkDuplex field to given value.


### GetUtilizationIn

`func (o *NetworkDevicePort) GetUtilizationIn() float32`

GetUtilizationIn returns the UtilizationIn field if non-nil, zero value otherwise.

### GetUtilizationInOk

`func (o *NetworkDevicePort) GetUtilizationInOk() (*float32, bool)`

GetUtilizationInOk returns a tuple with the UtilizationIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationIn

`func (o *NetworkDevicePort) SetUtilizationIn(v float32)`

SetUtilizationIn sets UtilizationIn field to given value.


### GetUtilizationOut

`func (o *NetworkDevicePort) GetUtilizationOut() float32`

GetUtilizationOut returns the UtilizationOut field if non-nil, zero value otherwise.

### GetUtilizationOutOk

`func (o *NetworkDevicePort) GetUtilizationOutOk() (*float32, bool)`

GetUtilizationOutOk returns a tuple with the UtilizationOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationOut

`func (o *NetworkDevicePort) SetUtilizationOut(v float32)`

SetUtilizationOut sets UtilizationOut field to given value.


### GetCounters

`func (o *NetworkDevicePort) GetCounters() NetworkDevicePortCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NetworkDevicePort) GetCountersOk() (*NetworkDevicePortCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NetworkDevicePort) SetCounters(v NetworkDevicePortCounters)`

SetCounters sets Counters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


