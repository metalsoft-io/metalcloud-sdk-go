# NetworkDeviceHealthCheckRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]NetworkDeviceHealthRule**](NetworkDeviceHealthRule.md) | List of health check rules to evaluate the monitoring metrics against | 

## Methods

### NewNetworkDeviceHealthCheckRules

`func NewNetworkDeviceHealthCheckRules(rules []NetworkDeviceHealthRule, ) *NetworkDeviceHealthCheckRules`

NewNetworkDeviceHealthCheckRules instantiates a new NetworkDeviceHealthCheckRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceHealthCheckRulesWithDefaults

`func NewNetworkDeviceHealthCheckRulesWithDefaults() *NetworkDeviceHealthCheckRules`

NewNetworkDeviceHealthCheckRulesWithDefaults instantiates a new NetworkDeviceHealthCheckRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *NetworkDeviceHealthCheckRules) GetRules() []NetworkDeviceHealthRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NetworkDeviceHealthCheckRules) GetRulesOk() (*[]NetworkDeviceHealthRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NetworkDeviceHealthCheckRules) SetRules(v []NetworkDeviceHealthRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


