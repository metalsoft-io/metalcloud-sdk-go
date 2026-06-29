# CreateNetworkFabricLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceAInterfaceId** | Pointer to **int64** | Unique identifier for the network device A interface. Required when externalSystemId is not set; also required when externalSystemId is set (the device side of the link). | [optional] 
**NetworkDeviceBInterfaceId** | Pointer to **int64** | Unique identifier for the network device B interface. Required when externalSystemId is not set; must not be set when externalSystemId is provided. | [optional] 
**ExternalSystemId** | Pointer to **int64** | Unique identifier for the external system acting as side B of this link. When set, networkDeviceAInterfaceId must be provided and networkDeviceBInterfaceId must be omitted. | [optional] 
**LinkType** | **string** | Type of the network fabric link | 
**Source** | Pointer to **string** | Source of the network fabric link information | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the network fabric link | [optional] 
**Config** | Pointer to [**NetworkFabricLinkConfig**](NetworkFabricLinkConfig.md) | Configuration of the network fabric link | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewCreateNetworkFabricLink

`func NewCreateNetworkFabricLink(linkType string, ) *CreateNetworkFabricLink`

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

`func (o *CreateNetworkFabricLink) GetNetworkDeviceAInterfaceId() int64`

GetNetworkDeviceAInterfaceId returns the NetworkDeviceAInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceAInterfaceIdOk

`func (o *CreateNetworkFabricLink) GetNetworkDeviceAInterfaceIdOk() (*int64, bool)`

GetNetworkDeviceAInterfaceIdOk returns a tuple with the NetworkDeviceAInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceAInterfaceId

`func (o *CreateNetworkFabricLink) SetNetworkDeviceAInterfaceId(v int64)`

SetNetworkDeviceAInterfaceId sets NetworkDeviceAInterfaceId field to given value.

### HasNetworkDeviceAInterfaceId

`func (o *CreateNetworkFabricLink) HasNetworkDeviceAInterfaceId() bool`

HasNetworkDeviceAInterfaceId returns a boolean if a field has been set.

### GetNetworkDeviceBInterfaceId

`func (o *CreateNetworkFabricLink) GetNetworkDeviceBInterfaceId() int64`

GetNetworkDeviceBInterfaceId returns the NetworkDeviceBInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceBInterfaceIdOk

`func (o *CreateNetworkFabricLink) GetNetworkDeviceBInterfaceIdOk() (*int64, bool)`

GetNetworkDeviceBInterfaceIdOk returns a tuple with the NetworkDeviceBInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceBInterfaceId

`func (o *CreateNetworkFabricLink) SetNetworkDeviceBInterfaceId(v int64)`

SetNetworkDeviceBInterfaceId sets NetworkDeviceBInterfaceId field to given value.

### HasNetworkDeviceBInterfaceId

`func (o *CreateNetworkFabricLink) HasNetworkDeviceBInterfaceId() bool`

HasNetworkDeviceBInterfaceId returns a boolean if a field has been set.

### GetExternalSystemId

`func (o *CreateNetworkFabricLink) GetExternalSystemId() int64`

GetExternalSystemId returns the ExternalSystemId field if non-nil, zero value otherwise.

### GetExternalSystemIdOk

`func (o *CreateNetworkFabricLink) GetExternalSystemIdOk() (*int64, bool)`

GetExternalSystemIdOk returns a tuple with the ExternalSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSystemId

`func (o *CreateNetworkFabricLink) SetExternalSystemId(v int64)`

SetExternalSystemId sets ExternalSystemId field to given value.

### HasExternalSystemId

`func (o *CreateNetworkFabricLink) HasExternalSystemId() bool`

HasExternalSystemId returns a boolean if a field has been set.

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


### GetSource

`func (o *CreateNetworkFabricLink) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateNetworkFabricLink) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateNetworkFabricLink) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateNetworkFabricLink) HasSource() bool`

HasSource returns a boolean if a field has been set.

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


