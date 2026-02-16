# CreateNetworkFabricLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceAInterfaceId** | **float32** | Unique identifier for the network device A interface | 
**NetworkDeviceBInterfaceId** | **float32** | Unique identifier for the network device B interface | 
**LinkType** | **string** | Type of the network fabric link | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the network fabric link | [optional] 
**Config** | Pointer to [**NetworkFabricLinkConfig**](NetworkFabricLinkConfig.md) | Configuration of the network fabric link | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewCreateNetworkFabricLink

`func NewCreateNetworkFabricLink(networkDeviceAInterfaceId float32, networkDeviceBInterfaceId float32, linkType string, ) *CreateNetworkFabricLink`

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

### GetConfig

`func (o *CreateNetworkFabricLink) GetConfig() NetworkFabricLinkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworkFabricLink) GetConfigOk() (*NetworkFabricLinkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworkFabricLink) SetConfig(v NetworkFabricLinkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNetworkFabricLink) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinks

`func (o *CreateNetworkFabricLink) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateNetworkFabricLink) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateNetworkFabricLink) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateNetworkFabricLink) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


