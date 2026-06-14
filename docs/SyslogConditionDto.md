# SyslogConditionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Syslog field to evaluate | 
**Operator** | **string** | Comparison operator to apply | 
**Value** | **string** | Value to compare against | 

## Methods

### NewSyslogConditionDto

`func NewSyslogConditionDto(field string, operator string, value string, ) *SyslogConditionDto`

NewSyslogConditionDto instantiates a new SyslogConditionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogConditionDtoWithDefaults

`func NewSyslogConditionDtoWithDefaults() *SyslogConditionDto`

NewSyslogConditionDtoWithDefaults instantiates a new SyslogConditionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SyslogConditionDto) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SyslogConditionDto) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SyslogConditionDto) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *SyslogConditionDto) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SyslogConditionDto) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SyslogConditionDto) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *SyslogConditionDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SyslogConditionDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SyslogConditionDto) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


