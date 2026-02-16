# UpdateNetworkDeviceBGPConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action when the template should be used | [optional] 
**NetworkType** | Pointer to **string** | Network type | [optional] 
**NetworkDeviceDriver** | Pointer to **string** | Network device driver | [optional] 
**NetworkDevicePosition** | Pointer to **string** | Network device position | [optional] 
**RemoteNetworkDevicePosition** | Pointer to **string** | Remote network device position | [optional] 
**BgpNumbering** | Pointer to **string** | BGP numbering | [optional] 
**BgpLinkConfiguration** | Pointer to **string** | BGP link configuration | [optional] 
**ExecutionType** | Pointer to **string** | Execution type | [optional] 
**LibraryLabel** | Pointer to **string** | Library label for the Network Device Configuration Template | [optional] 
**Preparation** | Pointer to **string** | Preparation commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 
**Configuration** | Pointer to **string** | Configuration commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 

## Methods

### NewUpdateNetworkDeviceBGPConfigurationTemplate

`func NewUpdateNetworkDeviceBGPConfigurationTemplate() *UpdateNetworkDeviceBGPConfigurationTemplate`

NewUpdateNetworkDeviceBGPConfigurationTemplate instantiates a new UpdateNetworkDeviceBGPConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDeviceBGPConfigurationTemplateWithDefaults

`func NewUpdateNetworkDeviceBGPConfigurationTemplateWithDefaults() *UpdateNetworkDeviceBGPConfigurationTemplate`

NewUpdateNetworkDeviceBGPConfigurationTemplateWithDefaults instantiates a new UpdateNetworkDeviceBGPConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetNetworkType

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.

### HasNetworkDeviceDriver

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasNetworkDeviceDriver() bool`

HasNetworkDeviceDriver returns a boolean if a field has been set.

### GetNetworkDevicePosition

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetNetworkDevicePosition() string`

GetNetworkDevicePosition returns the NetworkDevicePosition field if non-nil, zero value otherwise.

### GetNetworkDevicePositionOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetNetworkDevicePositionOk() (*string, bool)`

GetNetworkDevicePositionOk returns a tuple with the NetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePosition

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetNetworkDevicePosition(v string)`

SetNetworkDevicePosition sets NetworkDevicePosition field to given value.

### HasNetworkDevicePosition

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasNetworkDevicePosition() bool`

HasNetworkDevicePosition returns a boolean if a field has been set.

### GetRemoteNetworkDevicePosition

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetRemoteNetworkDevicePosition() string`

GetRemoteNetworkDevicePosition returns the RemoteNetworkDevicePosition field if non-nil, zero value otherwise.

### GetRemoteNetworkDevicePositionOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetRemoteNetworkDevicePositionOk() (*string, bool)`

GetRemoteNetworkDevicePositionOk returns a tuple with the RemoteNetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworkDevicePosition

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetRemoteNetworkDevicePosition(v string)`

SetRemoteNetworkDevicePosition sets RemoteNetworkDevicePosition field to given value.

### HasRemoteNetworkDevicePosition

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasRemoteNetworkDevicePosition() bool`

HasRemoteNetworkDevicePosition returns a boolean if a field has been set.

### GetBgpNumbering

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.

### HasBgpNumbering

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasBgpNumbering() bool`

HasBgpNumbering returns a boolean if a field has been set.

### GetBgpLinkConfiguration

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.

### HasBgpLinkConfiguration

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasBgpLinkConfiguration() bool`

HasBgpLinkConfiguration returns a boolean if a field has been set.

### GetExecutionType

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.

### GetLibraryLabel

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.

### HasLibraryLabel

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasLibraryLabel() bool`

HasLibraryLabel returns a boolean if a field has been set.

### GetPreparation

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateNetworkDeviceBGPConfigurationTemplate) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


