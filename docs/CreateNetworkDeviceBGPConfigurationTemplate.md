# CreateNetworkDeviceBGPConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action when the template should be used | 
**NetworkType** | **string** | Network type | 
**NetworkDeviceDriver** | **string** | Network device driver | 
**NetworkDevicePosition** | **string** | Network device position | 
**RemoteNetworkDevicePosition** | **string** | Remote network device position | 
**BgpNumbering** | **string** | BGP numbering | 
**BgpLinkConfiguration** | **string** | BGP link configuration | 
**ExecutionType** | **string** | Execution type | 
**LibraryLabel** | **string** | Library label for the Network Device Configuration Template | 
**Preparation** | Pointer to **string** | Preparation commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 
**Configuration** | **string** | Configuration commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | 

## Methods

### NewCreateNetworkDeviceBGPConfigurationTemplate

`func NewCreateNetworkDeviceBGPConfigurationTemplate(action string, networkType string, networkDeviceDriver string, networkDevicePosition string, remoteNetworkDevicePosition string, bgpNumbering string, bgpLinkConfiguration string, executionType string, libraryLabel string, configuration string, ) *CreateNetworkDeviceBGPConfigurationTemplate`

NewCreateNetworkDeviceBGPConfigurationTemplate instantiates a new CreateNetworkDeviceBGPConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceBGPConfigurationTemplateWithDefaults

`func NewCreateNetworkDeviceBGPConfigurationTemplateWithDefaults() *CreateNetworkDeviceBGPConfigurationTemplate`

NewCreateNetworkDeviceBGPConfigurationTemplateWithDefaults instantiates a new CreateNetworkDeviceBGPConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetAction(v string)`

SetAction sets Action field to given value.


### GetNetworkType

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetNetworkDeviceDriver

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetNetworkDevicePosition

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetNetworkDevicePosition() string`

GetNetworkDevicePosition returns the NetworkDevicePosition field if non-nil, zero value otherwise.

### GetNetworkDevicePositionOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetNetworkDevicePositionOk() (*string, bool)`

GetNetworkDevicePositionOk returns a tuple with the NetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePosition

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetNetworkDevicePosition(v string)`

SetNetworkDevicePosition sets NetworkDevicePosition field to given value.


### GetRemoteNetworkDevicePosition

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetRemoteNetworkDevicePosition() string`

GetRemoteNetworkDevicePosition returns the RemoteNetworkDevicePosition field if non-nil, zero value otherwise.

### GetRemoteNetworkDevicePositionOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetRemoteNetworkDevicePositionOk() (*string, bool)`

GetRemoteNetworkDevicePositionOk returns a tuple with the RemoteNetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworkDevicePosition

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetRemoteNetworkDevicePosition(v string)`

SetRemoteNetworkDevicePosition sets RemoteNetworkDevicePosition field to given value.


### GetBgpNumbering

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetExecutionType

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetLibraryLabel

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.


### GetPreparation

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateNetworkDeviceBGPConfigurationTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


