# NetworkDeviceHealthCheckOperand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | A JSONPath query to the metric value in the collected monitoring data (e.g. \&quot;metrics.temperatures[?(@.sensor&#x3D;&#x3D;&#39;cpu&#39;)].value\&quot;) | [optional] 
**Function** | Pointer to [**FieldValueFunctionType**](FieldValueFunctionType.md) | Function to apply to the selected metric value(s) (e.g. \&quot;first\&quot;, \&quot;length\&quot;, \&quot;avg\&quot;) | [optional] 
**Value** | Pointer to **map[string]interface{}** | A constant value to compare the metric against, used when function is not specified | [optional] 
**Transform** | Pointer to **map[string]interface{}** | Additional transformation to apply to the metric value before evaluation (e.g. converting units) | [optional] 

## Methods

### NewNetworkDeviceHealthCheckOperand

`func NewNetworkDeviceHealthCheckOperand() *NetworkDeviceHealthCheckOperand`

NewNetworkDeviceHealthCheckOperand instantiates a new NetworkDeviceHealthCheckOperand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceHealthCheckOperandWithDefaults

`func NewNetworkDeviceHealthCheckOperandWithDefaults() *NetworkDeviceHealthCheckOperand`

NewNetworkDeviceHealthCheckOperandWithDefaults instantiates a new NetworkDeviceHealthCheckOperand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *NetworkDeviceHealthCheckOperand) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NetworkDeviceHealthCheckOperand) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NetworkDeviceHealthCheckOperand) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *NetworkDeviceHealthCheckOperand) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetFunction

`func (o *NetworkDeviceHealthCheckOperand) GetFunction() FieldValueFunctionType`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *NetworkDeviceHealthCheckOperand) GetFunctionOk() (*FieldValueFunctionType, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *NetworkDeviceHealthCheckOperand) SetFunction(v FieldValueFunctionType)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *NetworkDeviceHealthCheckOperand) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetValue

`func (o *NetworkDeviceHealthCheckOperand) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NetworkDeviceHealthCheckOperand) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NetworkDeviceHealthCheckOperand) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *NetworkDeviceHealthCheckOperand) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTransform

`func (o *NetworkDeviceHealthCheckOperand) GetTransform() map[string]interface{}`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *NetworkDeviceHealthCheckOperand) GetTransformOk() (*map[string]interface{}, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *NetworkDeviceHealthCheckOperand) SetTransform(v map[string]interface{})`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *NetworkDeviceHealthCheckOperand) HasTransform() bool`

HasTransform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


