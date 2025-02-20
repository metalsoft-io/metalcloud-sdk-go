# ServerMetricsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oem** | Pointer to **map[string]interface{}** | The metrics OEM info. | [optional] 
**Name** | Pointer to **string** | The metrics name. | [optional] 
**Label** | Pointer to **string** | The custom info of the server. | [optional] 
**Units** | Pointer to **string** | The metrics units. | [optional] 
**Number** | Pointer to **float32** | The metrics value. | [optional] 
**PhysicalContext** | Pointer to **string** | The metrics physical context. | [optional] 
**UpperThresholdFatal** | Pointer to **float32** | The metrics fatal threshold. | [optional] 
**UpperThresholdCritical** | Pointer to **float32** | The metrics critical threshold. | [optional] 

## Methods

### NewServerMetricsInfo

`func NewServerMetricsInfo() *ServerMetricsInfo`

NewServerMetricsInfo instantiates a new ServerMetricsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMetricsInfoWithDefaults

`func NewServerMetricsInfoWithDefaults() *ServerMetricsInfo`

NewServerMetricsInfoWithDefaults instantiates a new ServerMetricsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOem

`func (o *ServerMetricsInfo) GetOem() map[string]interface{}`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *ServerMetricsInfo) GetOemOk() (*map[string]interface{}, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *ServerMetricsInfo) SetOem(v map[string]interface{})`

SetOem sets Oem field to given value.

### HasOem

`func (o *ServerMetricsInfo) HasOem() bool`

HasOem returns a boolean if a field has been set.

### GetName

`func (o *ServerMetricsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerMetricsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerMetricsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerMetricsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *ServerMetricsInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerMetricsInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerMetricsInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerMetricsInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUnits

`func (o *ServerMetricsInfo) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ServerMetricsInfo) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ServerMetricsInfo) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ServerMetricsInfo) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetNumber

`func (o *ServerMetricsInfo) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ServerMetricsInfo) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ServerMetricsInfo) SetNumber(v float32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ServerMetricsInfo) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPhysicalContext

`func (o *ServerMetricsInfo) GetPhysicalContext() string`

GetPhysicalContext returns the PhysicalContext field if non-nil, zero value otherwise.

### GetPhysicalContextOk

`func (o *ServerMetricsInfo) GetPhysicalContextOk() (*string, bool)`

GetPhysicalContextOk returns a tuple with the PhysicalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalContext

`func (o *ServerMetricsInfo) SetPhysicalContext(v string)`

SetPhysicalContext sets PhysicalContext field to given value.

### HasPhysicalContext

`func (o *ServerMetricsInfo) HasPhysicalContext() bool`

HasPhysicalContext returns a boolean if a field has been set.

### GetUpperThresholdFatal

`func (o *ServerMetricsInfo) GetUpperThresholdFatal() float32`

GetUpperThresholdFatal returns the UpperThresholdFatal field if non-nil, zero value otherwise.

### GetUpperThresholdFatalOk

`func (o *ServerMetricsInfo) GetUpperThresholdFatalOk() (*float32, bool)`

GetUpperThresholdFatalOk returns a tuple with the UpperThresholdFatal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperThresholdFatal

`func (o *ServerMetricsInfo) SetUpperThresholdFatal(v float32)`

SetUpperThresholdFatal sets UpperThresholdFatal field to given value.

### HasUpperThresholdFatal

`func (o *ServerMetricsInfo) HasUpperThresholdFatal() bool`

HasUpperThresholdFatal returns a boolean if a field has been set.

### GetUpperThresholdCritical

`func (o *ServerMetricsInfo) GetUpperThresholdCritical() float32`

GetUpperThresholdCritical returns the UpperThresholdCritical field if non-nil, zero value otherwise.

### GetUpperThresholdCriticalOk

`func (o *ServerMetricsInfo) GetUpperThresholdCriticalOk() (*float32, bool)`

GetUpperThresholdCriticalOk returns a tuple with the UpperThresholdCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperThresholdCritical

`func (o *ServerMetricsInfo) SetUpperThresholdCritical(v float32)`

SetUpperThresholdCritical sets UpperThresholdCritical field to given value.

### HasUpperThresholdCritical

`func (o *ServerMetricsInfo) HasUpperThresholdCritical() bool`

HasUpperThresholdCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


