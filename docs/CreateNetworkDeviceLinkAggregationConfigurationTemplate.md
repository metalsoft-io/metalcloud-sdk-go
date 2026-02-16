# CreateNetworkDeviceLinkAggregationConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action when the template should be used | 
**AggregationType** | **string** | The type of link aggregation | 
**NetworkDeviceDriver** | **string** | Network device driver | 
**ExecutionType** | **string** | Execution type | 
**LibraryLabel** | **string** | Library label for the Network Device Configuration Template | 
**Preparation** | Pointer to **string** | Preparation commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 
**Configuration** | **string** | Configuration commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | 

## Methods

### NewCreateNetworkDeviceLinkAggregationConfigurationTemplate

`func NewCreateNetworkDeviceLinkAggregationConfigurationTemplate(action string, aggregationType string, networkDeviceDriver string, executionType string, libraryLabel string, configuration string, ) *CreateNetworkDeviceLinkAggregationConfigurationTemplate`

NewCreateNetworkDeviceLinkAggregationConfigurationTemplate instantiates a new CreateNetworkDeviceLinkAggregationConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults

`func NewCreateNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults() *CreateNetworkDeviceLinkAggregationConfigurationTemplate`

NewCreateNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults instantiates a new CreateNetworkDeviceLinkAggregationConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) SetAction(v string)`

SetAction sets Action field to given value.


### GetAggregationType

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.


### GetNetworkDeviceDriver

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetExecutionType

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetLibraryLabel

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.


### GetPreparation

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateNetworkDeviceLinkAggregationConfigurationTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


