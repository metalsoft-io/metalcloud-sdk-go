# UpdateNetworkDeviceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | Pointer to **string** | Network type | [optional] 
**NetworkDeviceDriver** | Pointer to **string** | Network device driver | [optional] 
**NetworkDevicePosition** | Pointer to **string** | Network device position | [optional] 
**RemoteNetworkDevicePosition** | Pointer to **string** | Remote network device position | [optional] 
**MlagPair** | Pointer to **float32** | MLAG pair | [optional] 
**BgpNumbering** | Pointer to **string** | BGP numbering | [optional] 
**BgpLinkConfiguration** | Pointer to **string** | BGP link configuration | [optional] 
**ExecutionType** | Pointer to **string** | Execution type | [optional] 
**LibraryLabel** | Pointer to **string** | Library label for the Network Device Configuration Template | [optional] 
**Preparation** | Pointer to **string** | Preparation commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 
**Configuration** | Pointer to **string** | Configuration commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 

## Methods

### NewUpdateNetworkDeviceConfigurationTemplate

`func NewUpdateNetworkDeviceConfigurationTemplate() *UpdateNetworkDeviceConfigurationTemplate`

NewUpdateNetworkDeviceConfigurationTemplate instantiates a new UpdateNetworkDeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDeviceConfigurationTemplateWithDefaults

`func NewUpdateNetworkDeviceConfigurationTemplateWithDefaults() *UpdateNetworkDeviceConfigurationTemplate`

NewUpdateNetworkDeviceConfigurationTemplateWithDefaults instantiates a new UpdateNetworkDeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.

### HasNetworkDeviceDriver

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasNetworkDeviceDriver() bool`

HasNetworkDeviceDriver returns a boolean if a field has been set.

### GetNetworkDevicePosition

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetNetworkDevicePosition() string`

GetNetworkDevicePosition returns the NetworkDevicePosition field if non-nil, zero value otherwise.

### GetNetworkDevicePositionOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetNetworkDevicePositionOk() (*string, bool)`

GetNetworkDevicePositionOk returns a tuple with the NetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePosition

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetNetworkDevicePosition(v string)`

SetNetworkDevicePosition sets NetworkDevicePosition field to given value.

### HasNetworkDevicePosition

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasNetworkDevicePosition() bool`

HasNetworkDevicePosition returns a boolean if a field has been set.

### GetRemoteNetworkDevicePosition

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetRemoteNetworkDevicePosition() string`

GetRemoteNetworkDevicePosition returns the RemoteNetworkDevicePosition field if non-nil, zero value otherwise.

### GetRemoteNetworkDevicePositionOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetRemoteNetworkDevicePositionOk() (*string, bool)`

GetRemoteNetworkDevicePositionOk returns a tuple with the RemoteNetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworkDevicePosition

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetRemoteNetworkDevicePosition(v string)`

SetRemoteNetworkDevicePosition sets RemoteNetworkDevicePosition field to given value.

### HasRemoteNetworkDevicePosition

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasRemoteNetworkDevicePosition() bool`

HasRemoteNetworkDevicePosition returns a boolean if a field has been set.

### GetMlagPair

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetMlagPair() float32`

GetMlagPair returns the MlagPair field if non-nil, zero value otherwise.

### GetMlagPairOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetMlagPairOk() (*float32, bool)`

GetMlagPairOk returns a tuple with the MlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPair

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetMlagPair(v float32)`

SetMlagPair sets MlagPair field to given value.

### HasMlagPair

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasMlagPair() bool`

HasMlagPair returns a boolean if a field has been set.

### GetBgpNumbering

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.

### HasBgpNumbering

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasBgpNumbering() bool`

HasBgpNumbering returns a boolean if a field has been set.

### GetBgpLinkConfiguration

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.

### HasBgpLinkConfiguration

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasBgpLinkConfiguration() bool`

HasBgpLinkConfiguration returns a boolean if a field has been set.

### GetExecutionType

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.

### GetLibraryLabel

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.

### HasLibraryLabel

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasLibraryLabel() bool`

HasLibraryLabel returns a boolean if a field has been set.

### GetPreparation

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateNetworkDeviceConfigurationTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateNetworkDeviceConfigurationTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateNetworkDeviceConfigurationTemplate) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


