# CreateNetworkDeviceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | **string** | Network type | 
**NetworkDeviceDriver** | **string** | Network device driver | 
**NetworkDevicePosition** | **string** | Network device position | 
**RemoteNetworkDevicePosition** | **string** | Remote network device position | 
**MlagPair** | **float32** | MLAG pair | 
**BgpNumbering** | **string** | BGP numbering | 
**BgpLinkConfiguration** | **string** | BGP link configuration | 
**ExecutionType** | **string** | Execution type | 
**LibraryLabel** | **string** | Library label for the Network Device Configuration Template | 
**Preparation** | Pointer to **string** | Preparation commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | [optional] 
**Configuration** | **string** | Configuration commands in base64 format. It should start with the necessary commands to start configuring the device. Example: sonic-cli configure terminal interface Eth1/1 | 

## Methods

### NewCreateNetworkDeviceConfigurationTemplate

`func NewCreateNetworkDeviceConfigurationTemplate(networkType string, networkDeviceDriver string, networkDevicePosition string, remoteNetworkDevicePosition string, mlagPair float32, bgpNumbering string, bgpLinkConfiguration string, executionType string, libraryLabel string, configuration string, ) *CreateNetworkDeviceConfigurationTemplate`

NewCreateNetworkDeviceConfigurationTemplate instantiates a new CreateNetworkDeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceConfigurationTemplateWithDefaults

`func NewCreateNetworkDeviceConfigurationTemplateWithDefaults() *CreateNetworkDeviceConfigurationTemplate`

NewCreateNetworkDeviceConfigurationTemplateWithDefaults instantiates a new CreateNetworkDeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *CreateNetworkDeviceConfigurationTemplate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *CreateNetworkDeviceConfigurationTemplate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetNetworkDeviceDriver

`func (o *CreateNetworkDeviceConfigurationTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *CreateNetworkDeviceConfigurationTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetNetworkDevicePosition

`func (o *CreateNetworkDeviceConfigurationTemplate) GetNetworkDevicePosition() string`

GetNetworkDevicePosition returns the NetworkDevicePosition field if non-nil, zero value otherwise.

### GetNetworkDevicePositionOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetNetworkDevicePositionOk() (*string, bool)`

GetNetworkDevicePositionOk returns a tuple with the NetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePosition

`func (o *CreateNetworkDeviceConfigurationTemplate) SetNetworkDevicePosition(v string)`

SetNetworkDevicePosition sets NetworkDevicePosition field to given value.


### GetRemoteNetworkDevicePosition

`func (o *CreateNetworkDeviceConfigurationTemplate) GetRemoteNetworkDevicePosition() string`

GetRemoteNetworkDevicePosition returns the RemoteNetworkDevicePosition field if non-nil, zero value otherwise.

### GetRemoteNetworkDevicePositionOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetRemoteNetworkDevicePositionOk() (*string, bool)`

GetRemoteNetworkDevicePositionOk returns a tuple with the RemoteNetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworkDevicePosition

`func (o *CreateNetworkDeviceConfigurationTemplate) SetRemoteNetworkDevicePosition(v string)`

SetRemoteNetworkDevicePosition sets RemoteNetworkDevicePosition field to given value.


### GetMlagPair

`func (o *CreateNetworkDeviceConfigurationTemplate) GetMlagPair() float32`

GetMlagPair returns the MlagPair field if non-nil, zero value otherwise.

### GetMlagPairOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetMlagPairOk() (*float32, bool)`

GetMlagPairOk returns a tuple with the MlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPair

`func (o *CreateNetworkDeviceConfigurationTemplate) SetMlagPair(v float32)`

SetMlagPair sets MlagPair field to given value.


### GetBgpNumbering

`func (o *CreateNetworkDeviceConfigurationTemplate) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *CreateNetworkDeviceConfigurationTemplate) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *CreateNetworkDeviceConfigurationTemplate) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *CreateNetworkDeviceConfigurationTemplate) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetExecutionType

`func (o *CreateNetworkDeviceConfigurationTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *CreateNetworkDeviceConfigurationTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetLibraryLabel

`func (o *CreateNetworkDeviceConfigurationTemplate) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *CreateNetworkDeviceConfigurationTemplate) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.


### GetPreparation

`func (o *CreateNetworkDeviceConfigurationTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *CreateNetworkDeviceConfigurationTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *CreateNetworkDeviceConfigurationTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateNetworkDeviceConfigurationTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateNetworkDeviceConfigurationTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateNetworkDeviceConfigurationTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


