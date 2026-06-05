# NetworkDeviceHealthRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**NetworkDeviceHealthCheckCondition**](NetworkDeviceHealthCheckCondition.md) | Condition to evaluate for the health rule | 
**Status** | [**NetworkDeviceHealthStatus**](NetworkDeviceHealthStatus.md) | Status to set for the device if the health rule condition is met | 
**Message** | Pointer to **string** | A message to include in the health check result if the health rule condition is met | [optional] 
**MessageVariables** | Pointer to **map[string]interface{}** | Additional variables to render into the message in {{ key }} format | [optional] 

## Methods

### NewNetworkDeviceHealthRule

`func NewNetworkDeviceHealthRule(condition NetworkDeviceHealthCheckCondition, status NetworkDeviceHealthStatus, ) *NetworkDeviceHealthRule`

NewNetworkDeviceHealthRule instantiates a new NetworkDeviceHealthRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceHealthRuleWithDefaults

`func NewNetworkDeviceHealthRuleWithDefaults() *NetworkDeviceHealthRule`

NewNetworkDeviceHealthRuleWithDefaults instantiates a new NetworkDeviceHealthRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *NetworkDeviceHealthRule) GetCondition() NetworkDeviceHealthCheckCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *NetworkDeviceHealthRule) GetConditionOk() (*NetworkDeviceHealthCheckCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *NetworkDeviceHealthRule) SetCondition(v NetworkDeviceHealthCheckCondition)`

SetCondition sets Condition field to given value.


### GetStatus

`func (o *NetworkDeviceHealthRule) GetStatus() NetworkDeviceHealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDeviceHealthRule) GetStatusOk() (*NetworkDeviceHealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDeviceHealthRule) SetStatus(v NetworkDeviceHealthStatus)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *NetworkDeviceHealthRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NetworkDeviceHealthRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NetworkDeviceHealthRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NetworkDeviceHealthRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageVariables

`func (o *NetworkDeviceHealthRule) GetMessageVariables() map[string]interface{}`

GetMessageVariables returns the MessageVariables field if non-nil, zero value otherwise.

### GetMessageVariablesOk

`func (o *NetworkDeviceHealthRule) GetMessageVariablesOk() (*map[string]interface{}, bool)`

GetMessageVariablesOk returns a tuple with the MessageVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVariables

`func (o *NetworkDeviceHealthRule) SetMessageVariables(v map[string]interface{})`

SetMessageVariables sets MessageVariables field to given value.

### HasMessageVariables

`func (o *NetworkDeviceHealthRule) HasMessageVariables() bool`

HasMessageVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


