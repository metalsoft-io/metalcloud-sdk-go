# UpdateNetworkDeviceBGPInterconnectConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label of the interconnect template. | [optional] 
**Name** | Pointer to **string** | Name of the interconnect template. | [optional] 
**NetworkDeviceDriver** | Pointer to **string** | The network device driver this template applies to | [optional] 
**ExecutionType** | Pointer to **string** | The execution type for the template | [optional] 
**AddGlobalConfig** | Pointer to **string** | Add global config commands | [optional] 
**RemoveGlobalConfig** | Pointer to **string** | Remove global config commands (base64) | [optional] 
**AddNeighbor** | Pointer to **string** | Add neighbor commands (base64) | [optional] 
**RemoveNeighbor** | Pointer to **string** | Remove neighbor commands (base64) | [optional] 

## Methods

### NewUpdateNetworkDeviceBGPInterconnectConfigurationTemplate

`func NewUpdateNetworkDeviceBGPInterconnectConfigurationTemplate() *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate`

NewUpdateNetworkDeviceBGPInterconnectConfigurationTemplate instantiates a new UpdateNetworkDeviceBGPInterconnectConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults

`func NewUpdateNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults() *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate`

NewUpdateNetworkDeviceBGPInterconnectConfigurationTemplateWithDefaults instantiates a new UpdateNetworkDeviceBGPInterconnectConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.

### HasNetworkDeviceDriver

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasNetworkDeviceDriver() bool`

HasNetworkDeviceDriver returns a boolean if a field has been set.

### GetExecutionType

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.

### GetAddGlobalConfig

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddGlobalConfig() string`

GetAddGlobalConfig returns the AddGlobalConfig field if non-nil, zero value otherwise.

### GetAddGlobalConfigOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddGlobalConfigOk() (*string, bool)`

GetAddGlobalConfigOk returns a tuple with the AddGlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGlobalConfig

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetAddGlobalConfig(v string)`

SetAddGlobalConfig sets AddGlobalConfig field to given value.

### HasAddGlobalConfig

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasAddGlobalConfig() bool`

HasAddGlobalConfig returns a boolean if a field has been set.

### GetRemoveGlobalConfig

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveGlobalConfig() string`

GetRemoveGlobalConfig returns the RemoveGlobalConfig field if non-nil, zero value otherwise.

### GetRemoveGlobalConfigOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveGlobalConfigOk() (*string, bool)`

GetRemoveGlobalConfigOk returns a tuple with the RemoveGlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveGlobalConfig

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetRemoveGlobalConfig(v string)`

SetRemoveGlobalConfig sets RemoveGlobalConfig field to given value.

### HasRemoveGlobalConfig

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasRemoveGlobalConfig() bool`

HasRemoveGlobalConfig returns a boolean if a field has been set.

### GetAddNeighbor

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddNeighbor() string`

GetAddNeighbor returns the AddNeighbor field if non-nil, zero value otherwise.

### GetAddNeighborOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetAddNeighborOk() (*string, bool)`

GetAddNeighborOk returns a tuple with the AddNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNeighbor

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetAddNeighbor(v string)`

SetAddNeighbor sets AddNeighbor field to given value.

### HasAddNeighbor

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasAddNeighbor() bool`

HasAddNeighbor returns a boolean if a field has been set.

### GetRemoveNeighbor

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveNeighbor() string`

GetRemoveNeighbor returns the RemoveNeighbor field if non-nil, zero value otherwise.

### GetRemoveNeighborOk

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) GetRemoveNeighborOk() (*string, bool)`

GetRemoveNeighborOk returns a tuple with the RemoveNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveNeighbor

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) SetRemoveNeighbor(v string)`

SetRemoveNeighbor sets RemoveNeighbor field to given value.

### HasRemoveNeighbor

`func (o *UpdateNetworkDeviceBGPInterconnectConfigurationTemplate) HasRemoveNeighbor() bool`

HasRemoveNeighbor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


