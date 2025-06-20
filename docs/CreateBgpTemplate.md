# CreateBgpTemplate

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
**Preparation** | Pointer to **map[string]interface{}** | Preparation JSON | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Configuration JSON | [optional] 

## Methods

### NewCreateBgpTemplate

`func NewCreateBgpTemplate(networkType string, networkDeviceDriver string, networkDevicePosition string, remoteNetworkDevicePosition string, mlagPair float32, bgpNumbering string, bgpLinkConfiguration string, executionType string, ) *CreateBgpTemplate`

NewCreateBgpTemplate instantiates a new CreateBgpTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBgpTemplateWithDefaults

`func NewCreateBgpTemplateWithDefaults() *CreateBgpTemplate`

NewCreateBgpTemplateWithDefaults instantiates a new CreateBgpTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *CreateBgpTemplate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *CreateBgpTemplate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *CreateBgpTemplate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetNetworkDeviceDriver

`func (o *CreateBgpTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *CreateBgpTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *CreateBgpTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetNetworkDevicePosition

`func (o *CreateBgpTemplate) GetNetworkDevicePosition() string`

GetNetworkDevicePosition returns the NetworkDevicePosition field if non-nil, zero value otherwise.

### GetNetworkDevicePositionOk

`func (o *CreateBgpTemplate) GetNetworkDevicePositionOk() (*string, bool)`

GetNetworkDevicePositionOk returns a tuple with the NetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePosition

`func (o *CreateBgpTemplate) SetNetworkDevicePosition(v string)`

SetNetworkDevicePosition sets NetworkDevicePosition field to given value.


### GetRemoteNetworkDevicePosition

`func (o *CreateBgpTemplate) GetRemoteNetworkDevicePosition() string`

GetRemoteNetworkDevicePosition returns the RemoteNetworkDevicePosition field if non-nil, zero value otherwise.

### GetRemoteNetworkDevicePositionOk

`func (o *CreateBgpTemplate) GetRemoteNetworkDevicePositionOk() (*string, bool)`

GetRemoteNetworkDevicePositionOk returns a tuple with the RemoteNetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworkDevicePosition

`func (o *CreateBgpTemplate) SetRemoteNetworkDevicePosition(v string)`

SetRemoteNetworkDevicePosition sets RemoteNetworkDevicePosition field to given value.


### GetMlagPair

`func (o *CreateBgpTemplate) GetMlagPair() float32`

GetMlagPair returns the MlagPair field if non-nil, zero value otherwise.

### GetMlagPairOk

`func (o *CreateBgpTemplate) GetMlagPairOk() (*float32, bool)`

GetMlagPairOk returns a tuple with the MlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPair

`func (o *CreateBgpTemplate) SetMlagPair(v float32)`

SetMlagPair sets MlagPair field to given value.


### GetBgpNumbering

`func (o *CreateBgpTemplate) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *CreateBgpTemplate) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *CreateBgpTemplate) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *CreateBgpTemplate) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *CreateBgpTemplate) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *CreateBgpTemplate) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetExecutionType

`func (o *CreateBgpTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *CreateBgpTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *CreateBgpTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetPreparation

`func (o *CreateBgpTemplate) GetPreparation() map[string]interface{}`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *CreateBgpTemplate) GetPreparationOk() (*map[string]interface{}, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *CreateBgpTemplate) SetPreparation(v map[string]interface{})`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *CreateBgpTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateBgpTemplate) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateBgpTemplate) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateBgpTemplate) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateBgpTemplate) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


