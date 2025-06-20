# NetworkFabricLinkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Id** | **float32** | Unique identifier for the network fabric link | 
**NetworkFabricId** | **float32** | Unique identifier for the network fabric | 
**NetworkDeviceAInterfaceId** | **float32** | Unique identifier for the network device A interface | 
**NetworkDeviceBInterfaceId** | **float32** | Unique identifier for the network device B interface | 
**LinkType** | **string** | Type of the network fabric link | 
**MlagPair** | **float32** | Is the link part of an MLAG pair | 
**BgpNumbering** | **string** | BGP numbering type for the link | 
**BgpLinkConfiguration** | **string** | BGP link configuration type | 
**Status** | **string** | Status of the network fabric link | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the network fabric link | [optional] 

## Methods

### NewNetworkFabricLinkDto

`func NewNetworkFabricLinkDto(createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id float32, networkFabricId float32, networkDeviceAInterfaceId float32, networkDeviceBInterfaceId float32, linkType string, mlagPair float32, bgpNumbering string, bgpLinkConfiguration string, status string, ) *NetworkFabricLinkDto`

NewNetworkFabricLinkDto instantiates a new NetworkFabricLinkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricLinkDtoWithDefaults

`func NewNetworkFabricLinkDtoWithDefaults() *NetworkFabricLinkDto`

NewNetworkFabricLinkDtoWithDefaults instantiates a new NetworkFabricLinkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *NetworkFabricLinkDto) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkFabricLinkDto) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkFabricLinkDto) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkFabricLinkDto) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkFabricLinkDto) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkFabricLinkDto) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *NetworkFabricLinkDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabricLinkDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabricLinkDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetId

`func (o *NetworkFabricLinkDto) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkFabricLinkDto) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkFabricLinkDto) SetId(v float32)`

SetId sets Id field to given value.


### GetNetworkFabricId

`func (o *NetworkFabricLinkDto) GetNetworkFabricId() float32`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *NetworkFabricLinkDto) GetNetworkFabricIdOk() (*float32, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *NetworkFabricLinkDto) SetNetworkFabricId(v float32)`

SetNetworkFabricId sets NetworkFabricId field to given value.


### GetNetworkDeviceAInterfaceId

`func (o *NetworkFabricLinkDto) GetNetworkDeviceAInterfaceId() float32`

GetNetworkDeviceAInterfaceId returns the NetworkDeviceAInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceAInterfaceIdOk

`func (o *NetworkFabricLinkDto) GetNetworkDeviceAInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceAInterfaceIdOk returns a tuple with the NetworkDeviceAInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceAInterfaceId

`func (o *NetworkFabricLinkDto) SetNetworkDeviceAInterfaceId(v float32)`

SetNetworkDeviceAInterfaceId sets NetworkDeviceAInterfaceId field to given value.


### GetNetworkDeviceBInterfaceId

`func (o *NetworkFabricLinkDto) GetNetworkDeviceBInterfaceId() float32`

GetNetworkDeviceBInterfaceId returns the NetworkDeviceBInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceBInterfaceIdOk

`func (o *NetworkFabricLinkDto) GetNetworkDeviceBInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceBInterfaceIdOk returns a tuple with the NetworkDeviceBInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceBInterfaceId

`func (o *NetworkFabricLinkDto) SetNetworkDeviceBInterfaceId(v float32)`

SetNetworkDeviceBInterfaceId sets NetworkDeviceBInterfaceId field to given value.


### GetLinkType

`func (o *NetworkFabricLinkDto) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *NetworkFabricLinkDto) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *NetworkFabricLinkDto) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.


### GetMlagPair

`func (o *NetworkFabricLinkDto) GetMlagPair() float32`

GetMlagPair returns the MlagPair field if non-nil, zero value otherwise.

### GetMlagPairOk

`func (o *NetworkFabricLinkDto) GetMlagPairOk() (*float32, bool)`

GetMlagPairOk returns a tuple with the MlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPair

`func (o *NetworkFabricLinkDto) SetMlagPair(v float32)`

SetMlagPair sets MlagPair field to given value.


### GetBgpNumbering

`func (o *NetworkFabricLinkDto) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *NetworkFabricLinkDto) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *NetworkFabricLinkDto) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *NetworkFabricLinkDto) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *NetworkFabricLinkDto) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *NetworkFabricLinkDto) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetStatus

`func (o *NetworkFabricLinkDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabricLinkDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabricLinkDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomVariables

`func (o *NetworkFabricLinkDto) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *NetworkFabricLinkDto) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *NetworkFabricLinkDto) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *NetworkFabricLinkDto) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


