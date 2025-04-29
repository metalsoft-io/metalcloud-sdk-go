# PartialNetworkDevicePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZeroTouchRegistrationEnabled** | Pointer to **bool** | Whether zero-touch registration is enabled | [optional] 
**GNMISubscriptionEnabled** | Pointer to **bool** | Whether gNMI subscription is enabled | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Whether syslog monitoring is enabled | [optional] 
**LeafSwitchesHaveMlagPairs** | Pointer to **bool** | Whether leaf switches have MLAG pairs | [optional] 
**MinimumNumberOfSpinesNeededToBootstrapLeafs** | Pointer to **int32** | Minimum number of spines needed to bootstrap leafs | [optional] 

## Methods

### NewPartialNetworkDevicePolicy

`func NewPartialNetworkDevicePolicy() *PartialNetworkDevicePolicy`

NewPartialNetworkDevicePolicy instantiates a new PartialNetworkDevicePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialNetworkDevicePolicyWithDefaults

`func NewPartialNetworkDevicePolicyWithDefaults() *PartialNetworkDevicePolicy`

NewPartialNetworkDevicePolicyWithDefaults instantiates a new PartialNetworkDevicePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZeroTouchRegistrationEnabled

`func (o *PartialNetworkDevicePolicy) GetZeroTouchRegistrationEnabled() bool`

GetZeroTouchRegistrationEnabled returns the ZeroTouchRegistrationEnabled field if non-nil, zero value otherwise.

### GetZeroTouchRegistrationEnabledOk

`func (o *PartialNetworkDevicePolicy) GetZeroTouchRegistrationEnabledOk() (*bool, bool)`

GetZeroTouchRegistrationEnabledOk returns a tuple with the ZeroTouchRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchRegistrationEnabled

`func (o *PartialNetworkDevicePolicy) SetZeroTouchRegistrationEnabled(v bool)`

SetZeroTouchRegistrationEnabled sets ZeroTouchRegistrationEnabled field to given value.

### HasZeroTouchRegistrationEnabled

`func (o *PartialNetworkDevicePolicy) HasZeroTouchRegistrationEnabled() bool`

HasZeroTouchRegistrationEnabled returns a boolean if a field has been set.

### GetGNMISubscriptionEnabled

`func (o *PartialNetworkDevicePolicy) GetGNMISubscriptionEnabled() bool`

GetGNMISubscriptionEnabled returns the GNMISubscriptionEnabled field if non-nil, zero value otherwise.

### GetGNMISubscriptionEnabledOk

`func (o *PartialNetworkDevicePolicy) GetGNMISubscriptionEnabledOk() (*bool, bool)`

GetGNMISubscriptionEnabledOk returns a tuple with the GNMISubscriptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNMISubscriptionEnabled

`func (o *PartialNetworkDevicePolicy) SetGNMISubscriptionEnabled(v bool)`

SetGNMISubscriptionEnabled sets GNMISubscriptionEnabled field to given value.

### HasGNMISubscriptionEnabled

`func (o *PartialNetworkDevicePolicy) HasGNMISubscriptionEnabled() bool`

HasGNMISubscriptionEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *PartialNetworkDevicePolicy) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *PartialNetworkDevicePolicy) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *PartialNetworkDevicePolicy) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *PartialNetworkDevicePolicy) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetLeafSwitchesHaveMlagPairs

`func (o *PartialNetworkDevicePolicy) GetLeafSwitchesHaveMlagPairs() bool`

GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field if non-nil, zero value otherwise.

### GetLeafSwitchesHaveMlagPairsOk

`func (o *PartialNetworkDevicePolicy) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool)`

GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafSwitchesHaveMlagPairs

`func (o *PartialNetworkDevicePolicy) SetLeafSwitchesHaveMlagPairs(v bool)`

SetLeafSwitchesHaveMlagPairs sets LeafSwitchesHaveMlagPairs field to given value.

### HasLeafSwitchesHaveMlagPairs

`func (o *PartialNetworkDevicePolicy) HasLeafSwitchesHaveMlagPairs() bool`

HasLeafSwitchesHaveMlagPairs returns a boolean if a field has been set.

### GetMinimumNumberOfSpinesNeededToBootstrapLeafs

`func (o *PartialNetworkDevicePolicy) GetMinimumNumberOfSpinesNeededToBootstrapLeafs() int32`

GetMinimumNumberOfSpinesNeededToBootstrapLeafs returns the MinimumNumberOfSpinesNeededToBootstrapLeafs field if non-nil, zero value otherwise.

### GetMinimumNumberOfSpinesNeededToBootstrapLeafsOk

`func (o *PartialNetworkDevicePolicy) GetMinimumNumberOfSpinesNeededToBootstrapLeafsOk() (*int32, bool)`

GetMinimumNumberOfSpinesNeededToBootstrapLeafsOk returns a tuple with the MinimumNumberOfSpinesNeededToBootstrapLeafs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumberOfSpinesNeededToBootstrapLeafs

`func (o *PartialNetworkDevicePolicy) SetMinimumNumberOfSpinesNeededToBootstrapLeafs(v int32)`

SetMinimumNumberOfSpinesNeededToBootstrapLeafs sets MinimumNumberOfSpinesNeededToBootstrapLeafs field to given value.

### HasMinimumNumberOfSpinesNeededToBootstrapLeafs

`func (o *PartialNetworkDevicePolicy) HasMinimumNumberOfSpinesNeededToBootstrapLeafs() bool`

HasMinimumNumberOfSpinesNeededToBootstrapLeafs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


