# NetworkDeviceHealthCheckExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**BooleanOperator**](BooleanOperator.md) | Boolean operator to apply between the left-hand side and right-hand side operands | 
**Lhs** | [**NetworkDeviceHealthCheckOperand**](NetworkDeviceHealthCheckOperand.md) | Left-hand side operand of the expression | 
**Rhs** | [**NetworkDeviceHealthCheckOperand**](NetworkDeviceHealthCheckOperand.md) | Right-hand side operand of the expression | 

## Methods

### NewNetworkDeviceHealthCheckExpression

`func NewNetworkDeviceHealthCheckExpression(operator BooleanOperator, lhs NetworkDeviceHealthCheckOperand, rhs NetworkDeviceHealthCheckOperand, ) *NetworkDeviceHealthCheckExpression`

NewNetworkDeviceHealthCheckExpression instantiates a new NetworkDeviceHealthCheckExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceHealthCheckExpressionWithDefaults

`func NewNetworkDeviceHealthCheckExpressionWithDefaults() *NetworkDeviceHealthCheckExpression`

NewNetworkDeviceHealthCheckExpressionWithDefaults instantiates a new NetworkDeviceHealthCheckExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *NetworkDeviceHealthCheckExpression) GetOperator() BooleanOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NetworkDeviceHealthCheckExpression) GetOperatorOk() (*BooleanOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NetworkDeviceHealthCheckExpression) SetOperator(v BooleanOperator)`

SetOperator sets Operator field to given value.


### GetLhs

`func (o *NetworkDeviceHealthCheckExpression) GetLhs() NetworkDeviceHealthCheckOperand`

GetLhs returns the Lhs field if non-nil, zero value otherwise.

### GetLhsOk

`func (o *NetworkDeviceHealthCheckExpression) GetLhsOk() (*NetworkDeviceHealthCheckOperand, bool)`

GetLhsOk returns a tuple with the Lhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLhs

`func (o *NetworkDeviceHealthCheckExpression) SetLhs(v NetworkDeviceHealthCheckOperand)`

SetLhs sets Lhs field to given value.


### GetRhs

`func (o *NetworkDeviceHealthCheckExpression) GetRhs() NetworkDeviceHealthCheckOperand`

GetRhs returns the Rhs field if non-nil, zero value otherwise.

### GetRhsOk

`func (o *NetworkDeviceHealthCheckExpression) GetRhsOk() (*NetworkDeviceHealthCheckOperand, bool)`

GetRhsOk returns a tuple with the Rhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhs

`func (o *NetworkDeviceHealthCheckExpression) SetRhs(v NetworkDeviceHealthCheckOperand)`

SetRhs sets Rhs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


