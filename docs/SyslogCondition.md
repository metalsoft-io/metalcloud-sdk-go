# SyslogCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Syslog field to evaluate | 
**Operator** | **string** | Comparison operator to apply | 
**Value** | **string** | Value to compare against | 

## Methods

### NewSyslogCondition

`func NewSyslogCondition(field string, operator string, value string, ) *SyslogCondition`

NewSyslogCondition instantiates a new SyslogCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogConditionWithDefaults

`func NewSyslogConditionWithDefaults() *SyslogCondition`

NewSyslogConditionWithDefaults instantiates a new SyslogCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SyslogCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SyslogCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SyslogCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *SyslogCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SyslogCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SyslogCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *SyslogCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SyslogCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SyslogCondition) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


