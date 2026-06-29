# SNMPOIDToMetricMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The name of the metric property that will be created from the SNMP OID value | 
**Type** | [**NetworkDeviceMetricValueType**](NetworkDeviceMetricValueType.md) | The data type of the metric property&#39;s value | 
**Value** | Pointer to [**NetworkDeviceHealthCheckOperand**](NetworkDeviceHealthCheckOperand.md) | The value to evaluate and for the metric property (for primitive types) | [optional] 
**Expression** | Pointer to [**NetworkDeviceHealthCheckExpression**](NetworkDeviceHealthCheckExpression.md) | An expression to evaluate for the metric property&#39;s value (for primitive types) | [optional] 
**Path** | Pointer to **string** | A JSONPath query to filter the collected SNMP objects used for dictionary_list type metrics | [optional] 
**Properties** | Pointer to **map[string]interface{}** | For dictionary or dictionary_list type metrics, a map containing the resulting dictionary&#39;s keys and value operands | [optional] 

## Methods

### NewSNMPOIDToMetricMapping

`func NewSNMPOIDToMetricMapping(metric string, type_ NetworkDeviceMetricValueType, ) *SNMPOIDToMetricMapping`

NewSNMPOIDToMetricMapping instantiates a new SNMPOIDToMetricMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPOIDToMetricMappingWithDefaults

`func NewSNMPOIDToMetricMappingWithDefaults() *SNMPOIDToMetricMapping`

NewSNMPOIDToMetricMappingWithDefaults instantiates a new SNMPOIDToMetricMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *SNMPOIDToMetricMapping) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SNMPOIDToMetricMapping) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SNMPOIDToMetricMapping) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetType

`func (o *SNMPOIDToMetricMapping) GetType() NetworkDeviceMetricValueType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SNMPOIDToMetricMapping) GetTypeOk() (*NetworkDeviceMetricValueType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SNMPOIDToMetricMapping) SetType(v NetworkDeviceMetricValueType)`

SetType sets Type field to given value.


### GetValue

`func (o *SNMPOIDToMetricMapping) GetValue() NetworkDeviceHealthCheckOperand`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SNMPOIDToMetricMapping) GetValueOk() (*NetworkDeviceHealthCheckOperand, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SNMPOIDToMetricMapping) SetValue(v NetworkDeviceHealthCheckOperand)`

SetValue sets Value field to given value.

### HasValue

`func (o *SNMPOIDToMetricMapping) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExpression

`func (o *SNMPOIDToMetricMapping) GetExpression() NetworkDeviceHealthCheckExpression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SNMPOIDToMetricMapping) GetExpressionOk() (*NetworkDeviceHealthCheckExpression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SNMPOIDToMetricMapping) SetExpression(v NetworkDeviceHealthCheckExpression)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SNMPOIDToMetricMapping) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetPath

`func (o *SNMPOIDToMetricMapping) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SNMPOIDToMetricMapping) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SNMPOIDToMetricMapping) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SNMPOIDToMetricMapping) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProperties

`func (o *SNMPOIDToMetricMapping) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SNMPOIDToMetricMapping) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SNMPOIDToMetricMapping) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SNMPOIDToMetricMapping) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


