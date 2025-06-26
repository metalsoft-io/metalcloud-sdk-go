# BgpTemplate

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
**Preparation** | Pointer to **string** | Preparation commands in base64 format | [optional] 
**Configuration** | **string** | Configuration commands in base64 format | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **float32** | BGP Template Id | 

## Methods

### NewBgpTemplate

`func NewBgpTemplate(networkType string, networkDeviceDriver string, networkDevicePosition string, remoteNetworkDevicePosition string, mlagPair float32, bgpNumbering string, bgpLinkConfiguration string, executionType string, configuration string, createdTimestamp time.Time, updatedTimestamp time.Time, id float32, ) *BgpTemplate`

NewBgpTemplate instantiates a new BgpTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpTemplateWithDefaults

`func NewBgpTemplateWithDefaults() *BgpTemplate`

NewBgpTemplateWithDefaults instantiates a new BgpTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *BgpTemplate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *BgpTemplate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *BgpTemplate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.


### GetNetworkDeviceDriver

`func (o *BgpTemplate) GetNetworkDeviceDriver() string`

GetNetworkDeviceDriver returns the NetworkDeviceDriver field if non-nil, zero value otherwise.

### GetNetworkDeviceDriverOk

`func (o *BgpTemplate) GetNetworkDeviceDriverOk() (*string, bool)`

GetNetworkDeviceDriverOk returns a tuple with the NetworkDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceDriver

`func (o *BgpTemplate) SetNetworkDeviceDriver(v string)`

SetNetworkDeviceDriver sets NetworkDeviceDriver field to given value.


### GetNetworkDevicePosition

`func (o *BgpTemplate) GetNetworkDevicePosition() string`

GetNetworkDevicePosition returns the NetworkDevicePosition field if non-nil, zero value otherwise.

### GetNetworkDevicePositionOk

`func (o *BgpTemplate) GetNetworkDevicePositionOk() (*string, bool)`

GetNetworkDevicePositionOk returns a tuple with the NetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePosition

`func (o *BgpTemplate) SetNetworkDevicePosition(v string)`

SetNetworkDevicePosition sets NetworkDevicePosition field to given value.


### GetRemoteNetworkDevicePosition

`func (o *BgpTemplate) GetRemoteNetworkDevicePosition() string`

GetRemoteNetworkDevicePosition returns the RemoteNetworkDevicePosition field if non-nil, zero value otherwise.

### GetRemoteNetworkDevicePositionOk

`func (o *BgpTemplate) GetRemoteNetworkDevicePositionOk() (*string, bool)`

GetRemoteNetworkDevicePositionOk returns a tuple with the RemoteNetworkDevicePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworkDevicePosition

`func (o *BgpTemplate) SetRemoteNetworkDevicePosition(v string)`

SetRemoteNetworkDevicePosition sets RemoteNetworkDevicePosition field to given value.


### GetMlagPair

`func (o *BgpTemplate) GetMlagPair() float32`

GetMlagPair returns the MlagPair field if non-nil, zero value otherwise.

### GetMlagPairOk

`func (o *BgpTemplate) GetMlagPairOk() (*float32, bool)`

GetMlagPairOk returns a tuple with the MlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPair

`func (o *BgpTemplate) SetMlagPair(v float32)`

SetMlagPair sets MlagPair field to given value.


### GetBgpNumbering

`func (o *BgpTemplate) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *BgpTemplate) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *BgpTemplate) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *BgpTemplate) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *BgpTemplate) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *BgpTemplate) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetExecutionType

`func (o *BgpTemplate) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *BgpTemplate) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *BgpTemplate) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.


### GetPreparation

`func (o *BgpTemplate) GetPreparation() string`

GetPreparation returns the Preparation field if non-nil, zero value otherwise.

### GetPreparationOk

`func (o *BgpTemplate) GetPreparationOk() (*string, bool)`

GetPreparationOk returns a tuple with the Preparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparation

`func (o *BgpTemplate) SetPreparation(v string)`

SetPreparation sets Preparation field to given value.

### HasPreparation

`func (o *BgpTemplate) HasPreparation() bool`

HasPreparation returns a boolean if a field has been set.

### GetConfiguration

`func (o *BgpTemplate) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BgpTemplate) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BgpTemplate) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.


### GetCreatedTimestamp

`func (o *BgpTemplate) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *BgpTemplate) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *BgpTemplate) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *BgpTemplate) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *BgpTemplate) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *BgpTemplate) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *BgpTemplate) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BgpTemplate) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BgpTemplate) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BgpTemplate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *BgpTemplate) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BgpTemplate) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BgpTemplate) SetId(v float32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


