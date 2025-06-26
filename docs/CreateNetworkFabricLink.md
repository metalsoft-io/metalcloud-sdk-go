# CreateNetworkFabricLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceAInterfaceId** | **float32** | Unique identifier for the network device A interface | 
**NetworkDeviceBInterfaceId** | **float32** | Unique identifier for the network device B interface | 
**LinkType** | **string** | Type of the network fabric link | 
**MlagPair** | **float32** | Is the link part of an MLAG pair | 
**BgpNumbering** | **string** | BGP numbering type for the link | 
**BgpLinkConfiguration** | **string** | BGP link configuration type | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the network fabric link | [optional] 

## Methods

### NewCreateNetworkFabricLink

`func NewCreateNetworkFabricLink(networkDeviceAInterfaceId float32, networkDeviceBInterfaceId float32, linkType string, mlagPair float32, bgpNumbering string, bgpLinkConfiguration string, ) *CreateNetworkFabricLink`

NewCreateNetworkFabricLink instantiates a new CreateNetworkFabricLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFabricLinkWithDefaults

`func NewCreateNetworkFabricLinkWithDefaults() *CreateNetworkFabricLink`

NewCreateNetworkFabricLinkWithDefaults instantiates a new CreateNetworkFabricLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceAInterfaceId

`func (o *CreateNetworkFabricLink) GetNetworkDeviceAInterfaceId() float32`

GetNetworkDeviceAInterfaceId returns the NetworkDeviceAInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceAInterfaceIdOk

`func (o *CreateNetworkFabricLink) GetNetworkDeviceAInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceAInterfaceIdOk returns a tuple with the NetworkDeviceAInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceAInterfaceId

`func (o *CreateNetworkFabricLink) SetNetworkDeviceAInterfaceId(v float32)`

SetNetworkDeviceAInterfaceId sets NetworkDeviceAInterfaceId field to given value.


### GetNetworkDeviceBInterfaceId

`func (o *CreateNetworkFabricLink) GetNetworkDeviceBInterfaceId() float32`

GetNetworkDeviceBInterfaceId returns the NetworkDeviceBInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceBInterfaceIdOk

`func (o *CreateNetworkFabricLink) GetNetworkDeviceBInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceBInterfaceIdOk returns a tuple with the NetworkDeviceBInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceBInterfaceId

`func (o *CreateNetworkFabricLink) SetNetworkDeviceBInterfaceId(v float32)`

SetNetworkDeviceBInterfaceId sets NetworkDeviceBInterfaceId field to given value.


### GetLinkType

`func (o *CreateNetworkFabricLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *CreateNetworkFabricLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *CreateNetworkFabricLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.


### GetMlagPair

`func (o *CreateNetworkFabricLink) GetMlagPair() float32`

GetMlagPair returns the MlagPair field if non-nil, zero value otherwise.

### GetMlagPairOk

`func (o *CreateNetworkFabricLink) GetMlagPairOk() (*float32, bool)`

GetMlagPairOk returns a tuple with the MlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPair

`func (o *CreateNetworkFabricLink) SetMlagPair(v float32)`

SetMlagPair sets MlagPair field to given value.


### GetBgpNumbering

`func (o *CreateNetworkFabricLink) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *CreateNetworkFabricLink) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *CreateNetworkFabricLink) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *CreateNetworkFabricLink) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *CreateNetworkFabricLink) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *CreateNetworkFabricLink) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetCustomVariables

`func (o *CreateNetworkFabricLink) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *CreateNetworkFabricLink) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *CreateNetworkFabricLink) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *CreateNetworkFabricLink) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


