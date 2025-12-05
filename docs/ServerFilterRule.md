# ServerFilterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** | The server property to filter on | 
**Operator** | **string** | The filtering operator to apply | 
**Value** | **NullableString** | The value to filter against (for regex operators, provide valid regex pattern). Not required for affinity/anti-affinity operators. | 

## Methods

### NewServerFilterRule

`func NewServerFilterRule(property string, operator string, value NullableString, ) *ServerFilterRule`

NewServerFilterRule instantiates a new ServerFilterRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFilterRuleWithDefaults

`func NewServerFilterRuleWithDefaults() *ServerFilterRule`

NewServerFilterRuleWithDefaults instantiates a new ServerFilterRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ServerFilterRule) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ServerFilterRule) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ServerFilterRule) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *ServerFilterRule) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ServerFilterRule) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ServerFilterRule) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *ServerFilterRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServerFilterRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServerFilterRule) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ServerFilterRule) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ServerFilterRule) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


