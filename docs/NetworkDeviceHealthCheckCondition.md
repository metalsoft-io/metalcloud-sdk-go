# NetworkDeviceHealthCheckCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**LogicalOperator**](LogicalOperator.md) | Logical operator to combine multiple conditions (e.g. \&quot;and\&quot;, \&quot;or\&quot;, \&quot;not\&quot;) | [optional] 
**Conditions** | Pointer to [**[]NetworkDeviceHealthCheckCondition**](NetworkDeviceHealthCheckCondition.md) | List of nested conditions to evaluate | [optional] 
**Expression** | Pointer to [**NetworkDeviceHealthCheckExpression**](NetworkDeviceHealthCheckExpression.md) | Expression to evaluate for the condition | [optional] 

## Methods

### NewNetworkDeviceHealthCheckCondition

`func NewNetworkDeviceHealthCheckCondition() *NetworkDeviceHealthCheckCondition`

NewNetworkDeviceHealthCheckCondition instantiates a new NetworkDeviceHealthCheckCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceHealthCheckConditionWithDefaults

`func NewNetworkDeviceHealthCheckConditionWithDefaults() *NetworkDeviceHealthCheckCondition`

NewNetworkDeviceHealthCheckConditionWithDefaults instantiates a new NetworkDeviceHealthCheckCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *NetworkDeviceHealthCheckCondition) GetOperator() LogicalOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NetworkDeviceHealthCheckCondition) GetOperatorOk() (*LogicalOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NetworkDeviceHealthCheckCondition) SetOperator(v LogicalOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *NetworkDeviceHealthCheckCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetConditions

`func (o *NetworkDeviceHealthCheckCondition) GetConditions() []NetworkDeviceHealthCheckCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *NetworkDeviceHealthCheckCondition) GetConditionsOk() (*[]NetworkDeviceHealthCheckCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *NetworkDeviceHealthCheckCondition) SetConditions(v []NetworkDeviceHealthCheckCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *NetworkDeviceHealthCheckCondition) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetExpression

`func (o *NetworkDeviceHealthCheckCondition) GetExpression() NetworkDeviceHealthCheckExpression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *NetworkDeviceHealthCheckCondition) GetExpressionOk() (*NetworkDeviceHealthCheckExpression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *NetworkDeviceHealthCheckCondition) SetExpression(v NetworkDeviceHealthCheckExpression)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *NetworkDeviceHealthCheckCondition) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


