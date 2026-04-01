# CreateNetworkDeviceBGPInterconnectConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the interconnect template. | 
**Name** | **string** | Name of the interconnect template. | 
**NetworkDeviceDriver** | **string** | The network device driver this template applies to | 
**ExecutionType** | **string** | The execution type for the template | 
**AddGlobalConfig** | Pointer to **string** | Add global config commands | [optional] 
**RemoveGlobalConfig** | **string** | Remove global config commands (base64) | 
**AddNeighbor** | Pointer to **string** | Add neighbor commands (base64) | [optional] 
**RemoveNeighbor** | Pointer to **string** | Remove neighbor commands (base64) | [optional] 

## Methods

### NewCreateNetworkDeviceBGPInterconnectConfigurationTemplate

`func NewCreateNetworkDeviceBGPInterconnectConfigurationTemplate(label string, name string, networkDeviceDriver string, executionType string, removeGlobalConfig string, ) *CreateNetworkDeviceBGPInterconnectConfigurationTemplate`

NewCreateNetworkDeviceBGPInterconnectConfigurationTemplate instantiates a new CreateNetworkDeviceBGPInterconnectConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults

`func NewCreateNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults() *CreateNetworkDeviceBGPInterconnectConfigurationTemplate`

NewCreateNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults instantiates a new CreateNetworkDeviceBGPInterconnectConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkDeviceDriver

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetExecutionType

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetAddGlobalConfig

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddGlobalConfig() string`

GetAddGlobalConfig returns the AddGlobalConfig field if non-nil, zero value otherwise.

### GetAddGlobalConfigOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddGlobalConfigOk() (*string, bool)`

GetAddGlobalConfigOk returns a tuple with the AddGlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGlobalConfig

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetAddGlobalConfig(v string)`

SetAddGlobalConfig sets AddGlobalConfig field to given value.

### HasAddGlobalConfig

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) HasAddGlobalConfig() bool`

HasAddGlobalConfig returns a boolean if a field has been set.

### GetRemoveGlobalConfig

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveGlobalConfig() string`

GetRemoveGlobalConfig returns the RemoveGlobalConfig field if non-nil, zero value otherwise.

### GetRemoveGlobalConfigOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveGlobalConfigOk() (*string, bool)`

GetRemoveGlobalConfigOk returns a tuple with the RemoveGlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveGlobalConfig

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetRemoveGlobalConfig(v string)`

SetRemoveGlobalConfig sets RemoveGlobalConfig field to given value.


### GetAddNeighbor

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddNeighbor() string`

GetAddNeighbor returns the AddNeighbor field if non-nil, zero value otherwise.

### GetAddNeighborOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddNeighborOk() (*string, bool)`

GetAddNeighborOk returns a tuple with the AddNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNeighbor

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetAddNeighbor(v string)`

SetAddNeighbor sets AddNeighbor field to given value.

### HasAddNeighbor

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) HasAddNeighbor() bool`

HasAddNeighbor returns a boolean if a field has been set.

### GetRemoveNeighbor

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveNeighbor() string`

GetRemoveNeighbor returns the RemoveNeighbor field if non-nil, zero value otherwise.

### GetRemoveNeighborOk

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveNeighborOk() (*string, bool)`

GetRemoveNeighborOk returns a tuple with the RemoveNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveNeighbor

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) SetRemoveNeighbor(v string)`

SetRemoveNeighbor sets RemoveNeighbor field to given value.

### HasRemoveNeighbor

`func (o *CreateNetworkDeviceBGPInterconnectConfigurationTemplate) HasRemoveNeighbor() bool`

HasRemoveNeighbor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


