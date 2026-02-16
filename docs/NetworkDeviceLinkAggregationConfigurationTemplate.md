# NetworkDeviceLinkAggregationConfigurationTemplate

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
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **float32** | Network Device Link Aggregation Configuration Template Id | 

## Methods

### NewNetworkDeviceLinkAggregationConfigurationTemplate

`func NewNetworkDeviceLinkAggregationConfigurationTemplate(action string, aggregationType string, networkDeviceDriver string, executionType string, libraryLabel string, configuration string, createdTimestamp time.Time, updatedTimestamp time.Time, id float32, ) *NetworkDeviceLinkAggregationConfigurationTemplate`

NewNetworkDeviceLinkAggregationConfigurationTemplate instantiates a new NetworkDeviceLinkAggregationConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults

`func NewNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults() *NetworkDeviceLinkAggregationConfigurationTemplate`

NewNetworkDeviceLinkAggregationConfigurationTemplateWithDefaults instantiates a new NetworkDeviceLinkAggregationConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetAction(v string)`

SetAction sets Action field to given value.


### GetAggregationType

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.


### GetNetworkDeviceDriver

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetExecutionType

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetLibraryLabel

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.


### GetPreparation

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.


### GetCreatedTimestamp

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplate) SetId(v float32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


