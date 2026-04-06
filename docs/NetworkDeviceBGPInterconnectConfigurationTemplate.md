# NetworkDeviceBGPInterconnectConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the interconnect template. | 
**Name** | **string** | Name of the interconnect template. | 
**NetworkDeviceDriver** | **string** | The network device driver this template applies to | 
**ExecutionType** | **string** | The execution type for the template | 
**AddGlobalConfig** | Pointer to **string** | Add global config commands | [optional] 
**RemoveGlobalConfig** | Pointer to **string** | Remove global config commands (base64) | [optional] 
**AddNeighbor** | Pointer to **string** | Add neighbor commands (base64) | [optional] 
**RemoveNeighbor** | Pointer to **string** | Remove neighbor commands (base64) | [optional] 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **float32** | Id | 

## Methods

### NewNetworkDeviceBGPInterconnectConfigurationTemplate

`func NewNetworkDeviceBGPInterconnectConfigurationTemplate(label string, name string, networkDeviceDriver string, executionType string, createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id float32, ) *NetworkDeviceBGPInterconnectConfigurationTemplate`

NewNetworkDeviceBGPInterconnectConfigurationTemplate instantiates a new NetworkDeviceBGPInterconnectConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults

`func NewNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults() *NetworkDeviceBGPInterconnectConfigurationTemplate`

NewNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults instantiates a new NetworkDeviceBGPInterconnectConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkDeviceDriver

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetExecutionType

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetAddGlobalConfig

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetAddGlobalConfig() string`

GetAddGlobalConfig returns the AddGlobalConfig field if non-nil, zero value otherwise.

### GetAddGlobalConfigOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetAddGlobalConfigOk() (*string, bool)`

GetAddGlobalConfigOk returns a tuple with the AddGlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGlobalConfig

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetAddGlobalConfig(v string)`

SetAddGlobalConfig sets AddGlobalConfig field to given value.

### HasAddGlobalConfig

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) HasAddGlobalConfig() bool`

HasAddGlobalConfig returns a boolean if a field has been set.

### GetRemoveGlobalConfig

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveGlobalConfig() string`

GetRemoveGlobalConfig returns the RemoveGlobalConfig field if non-nil, zero value otherwise.

### GetRemoveGlobalConfigOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveGlobalConfigOk() (*string, bool)`

GetRemoveGlobalConfigOk returns a tuple with the RemoveGlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveGlobalConfig

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetRemoveGlobalConfig(v string)`

SetRemoveGlobalConfig sets RemoveGlobalConfig field to given value.

### HasRemoveGlobalConfig

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) HasRemoveGlobalConfig() bool`

HasRemoveGlobalConfig returns a boolean if a field has been set.

### GetAddNeighbor

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetAddNeighbor() string`

GetAddNeighbor returns the AddNeighbor field if non-nil, zero value otherwise.

### GetAddNeighborOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetAddNeighborOk() (*string, bool)`

GetAddNeighborOk returns a tuple with the AddNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNeighbor

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetAddNeighbor(v string)`

SetAddNeighbor sets AddNeighbor field to given value.

### HasAddNeighbor

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) HasAddNeighbor() bool`

HasAddNeighbor returns a boolean if a field has been set.

### GetRemoveNeighbor

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveNeighbor() string`

GetRemoveNeighbor returns the RemoveNeighbor field if non-nil, zero value otherwise.

### GetRemoveNeighborOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveNeighborOk() (*string, bool)`

GetRemoveNeighborOk returns a tuple with the RemoveNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveNeighbor

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetRemoveNeighbor(v string)`

SetRemoveNeighbor sets RemoveNeighbor field to given value.

### HasRemoveNeighbor

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) HasRemoveNeighbor() bool`

HasRemoveNeighbor returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetLinks

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplate) SetId(v float32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


