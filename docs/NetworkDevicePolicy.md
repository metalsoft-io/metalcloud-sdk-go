# NetworkDevicePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZeroTouchRegistrationEnabled** | **bool** | Whether zero-touch registration is enabled | 
**GNMISubscriptionEnabled** | **bool** | Whether gNMI subscription is enabled | 
**SyslogMonitoringEnabled** | **bool** | Whether syslog monitoring is enabled | 
**LeafSwitchesHaveMlagPairs** | **bool** | Whether leaf switches have MLAG pairs | 
**MinimumNumberOfSpinesNeededToBootstrapLeafs** | **float32** | Minimum number of spines needed to bootstrap leafs | 

## Methods

### NewNetworkDevicePolicy

`func NewNetworkDevicePolicy(zeroTouchRegistrationEnabled bool, gNMISubscriptionEnabled bool, syslogMonitoringEnabled bool, leafSwitchesHaveMlagPairs bool, minimumNumberOfSpinesNeededToBootstrapLeafs float32, ) *NetworkDevicePolicy`

NewNetworkDevicePolicy instantiates a new NetworkDevicePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDevicePolicyWithDefaults

`func NewNetworkDevicePolicyWithDefaults() *NetworkDevicePolicy`

NewNetworkDevicePolicyWithDefaults instantiates a new NetworkDevicePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZeroTouchRegistrationEnabled

`func (o *NetworkDevicePolicy) GetZeroTouchRegistrationEnabled() bool`

GetZeroTouchRegistrationEnabled returns the ZeroTouchRegistrationEnabled field if non-nil, zero value otherwise.

### GetZeroTouchRegistrationEnabledOk

`func (o *NetworkDevicePolicy) GetZeroTouchRegistrationEnabledOk() (*bool, bool)`

GetZeroTouchRegistrationEnabledOk returns a tuple with the ZeroTouchRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchRegistrationEnabled

`func (o *NetworkDevicePolicy) SetZeroTouchRegistrationEnabled(v bool)`

SetZeroTouchRegistrationEnabled sets ZeroTouchRegistrationEnabled field to given value.


### GetGNMISubscriptionEnabled

`func (o *NetworkDevicePolicy) GetGNMISubscriptionEnabled() bool`

GetGNMISubscriptionEnabled returns the GNMISubscriptionEnabled field if non-nil, zero value otherwise.

### GetGNMISubscriptionEnabledOk

`func (o *NetworkDevicePolicy) GetGNMISubscriptionEnabledOk() (*bool, bool)`

GetGNMISubscriptionEnabledOk returns a tuple with the GNMISubscriptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNMISubscriptionEnabled

`func (o *NetworkDevicePolicy) SetGNMISubscriptionEnabled(v bool)`

SetGNMISubscriptionEnabled sets GNMISubscriptionEnabled field to given value.


### GetSyslogMonitoringEnabled

`func (o *NetworkDevicePolicy) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *NetworkDevicePolicy) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *NetworkDevicePolicy) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.


### GetLeafSwitchesHaveMlagPairs

`func (o *NetworkDevicePolicy) GetLeafSwitchesHaveMlagPairs() bool`

GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field if non-nil, zero value otherwise.

### GetLeafSwitchesHaveMlagPairsOk

`func (o *NetworkDevicePolicy) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool)`

GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafSwitchesHaveMlagPairs

`func (o *NetworkDevicePolicy) SetLeafSwitchesHaveMlagPairs(v bool)`

SetLeafSwitchesHaveMlagPairs sets LeafSwitchesHaveMlagPairs field to given value.


### GetMinimumNumberOfSpinesNeededToBootstrapLeafs

`func (o *NetworkDevicePolicy) GetMinimumNumberOfSpinesNeededToBootstrapLeafs() float32`

GetMinimumNumberOfSpinesNeededToBootstrapLeafs returns the MinimumNumberOfSpinesNeededToBootstrapLeafs field if non-nil, zero value otherwise.

### GetMinimumNumberOfSpinesNeededToBootstrapLeafsOk

`func (o *NetworkDevicePolicy) GetMinimumNumberOfSpinesNeededToBootstrapLeafsOk() (*float32, bool)`

GetMinimumNumberOfSpinesNeededToBootstrapLeafsOk returns a tuple with the MinimumNumberOfSpinesNeededToBootstrapLeafs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumberOfSpinesNeededToBootstrapLeafs

`func (o *NetworkDevicePolicy) SetMinimumNumberOfSpinesNeededToBootstrapLeafs(v float32)`

SetMinimumNumberOfSpinesNeededToBootstrapLeafs sets MinimumNumberOfSpinesNeededToBootstrapLeafs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


