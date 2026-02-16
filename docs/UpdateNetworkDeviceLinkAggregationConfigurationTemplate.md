# UpdateNetworkDeviceLinkAggregationConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action when the template should be used | [optional] 
**AggregationType** | Pointer to **string** | The type of link aggregation | [optional] 
**NetworkDeviceDriver** | Pointer to **string** | Network device driver | [optional] 
**ExecutionType** | Pointer to **string** | Execution type | [optional] 
**LibraryLabel** | Pointer to **string** | Library label for the Network Device Configuration Template | [optional] 
**Preparation** | Pointer to **string** | Preparation commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 
**Configuration** | Pointer to **string** | Configuration commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 

## Methods

### NewUpdateNetworkDeviceLinkAggregationConfigurationTemplate

`func NewUpdateNetworkDeviceLinkAggregationConfigurationTemplate() *UpdateNetworkDeviceLinkAggregationConfigurationTemplate`

NewUpdateNetworkDeviceLinkAggregationConfigurationTemplate instantiates a new UpdateNetworkDeviceLinkAggregationConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults

`func NewUpdateNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults() *UpdateNetworkDeviceLinkAggregationConfigurationTemplate`

NewUpdateNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults instantiates a new UpdateNetworkDeviceLinkAggregationConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAggregationType

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.

### HasNetworkDeviceDriver

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) HasNetworkDeviceDriver() bool`

HasNetworkDeviceDriver returns a boolean if a field has been set.

### GetExecutionType

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.

### GetLibraryLabel

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.

### HasLibraryLabel

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) HasLibraryLabel() bool`

HasLibraryLabel returns a boolean if a field has been set.

### GetPreparation

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateNetworkDeviceLinkAggregationConfigurationTemplate) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


