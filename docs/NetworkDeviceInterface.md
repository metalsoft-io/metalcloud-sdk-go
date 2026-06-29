# NetworkDeviceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | **int64** | The ID of the network equipment interface | 
**NetworkDeviceId** | **int64** | The ID of the network device | 
**InterfaceName** | **string** | The name of the network equipment interface | 
**InterfaceDescription** | Pointer to **string** | Description of the network equipment interface | [optional] 
**Kind** | **string** | The type of the network equipment interface | 
**ParentInterfaceId** | Pointer to **int64** | The ID of the parent interface if applicable | [optional] 
**LinkedInterfaceId** | Pointer to **int64** | The ID of the linked interface if applicable. (example: DPU physical interface p0 id 10 linked to switch port eth1/10 id 20) | [optional] 
**LinkedPortId** | Pointer to **string** | The identifier of the linked port if applicable | [optional] 
**LinkedSwitchHostname** | Pointer to **string** | The hostname of the linked switch if applicable | [optional] 
**InterfaceIndex** | Pointer to **float32** | The index of the network interface | [optional] 
**LagIdentifier** | Pointer to **float32** | LAG identifier | [optional] 
**NetworkDeviceLag** | Pointer to [**NullableNetworkDeviceLag**](NetworkDeviceLag.md) | The LAG this interface is a member of, if any. Read-only. | [optional] 
**ServerInterfaceId** | Pointer to **int64** | The server interface ID | [optional] 
**ServerId** | Pointer to **int64** | The server ID | [optional] 
**NumaNode** | Pointer to **float32** | NUMA node of the network device for optimal resource allocation | [optional] 
**DirtyBit** | **float32** | Dirty bit flag | 
**LldpInformation** | Pointer to **string** | LLDP information | [optional] 
**MacAddress** | Pointer to **string** | MAC address | [optional] 
**Wwn** | Pointer to **string** | WWPN | [optional] 
**DriverDumpCachedJson** | Pointer to **map[string]interface{}** | Driver dump cached JSON | [optional] 
**CachedUpdatedTimestamp** | **string** | Cached update timestamp | 
**InterfaceLLDPInformationServerInterface** | Pointer to **map[string]interface{}** | LLDP information of server interface in JSON format | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Tags** | **map[string]string** | Key/value tags. Filterable via &#x60;?tag.&lt;key&gt;&#x3D;&lt;value&gt;&#x60; query params (AND across keys). | 
**Config** | [**NetworkEquipmentInterfaceConfig**](NetworkEquipmentInterfaceConfig.md) |  | 
**Ipv4** | [**NetworkEquipmentInterfaceIpFamily**](NetworkEquipmentInterfaceIpFamily.md) |  | 
**Ipv6** | [**NetworkEquipmentInterfaceIpFamily**](NetworkEquipmentInterfaceIpFamily.md) |  | 

## Methods

### NewNetworkDeviceInterface

`func NewNetworkDeviceInterface(interfaceId int64, networkDeviceId int64, interfaceName string, kind string, dirtyBit float32, cachedUpdatedTimestamp string, tags map[string]string, config NetworkEquipmentInterfaceConfig, ipv4 NetworkEquipmentInterfaceIpFamily, ipv6 NetworkEquipmentInterfaceIpFamily, ) *NetworkDeviceInterface`

NewNetworkDeviceInterface instantiates a new NetworkDeviceInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceInterfaceWithDefaults

`func NewNetworkDeviceInterfaceWithDefaults() *NetworkDeviceInterface`

NewNetworkDeviceInterfaceWithDefaults instantiates a new NetworkDeviceInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *NetworkDeviceInterface) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkDeviceInterface) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkDeviceInterface) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.


### GetNetworkDeviceId

`func (o *NetworkDeviceInterface) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkDeviceInterface) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkDeviceInterface) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetInterfaceName

`func (o *NetworkDeviceInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NetworkDeviceInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NetworkDeviceInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetInterfaceDescription

`func (o *NetworkDeviceInterface) GetInterfaceDescription() string`

GetInterfaceDescription returns the InterfaceDescription field if non-nil, zero value otherwise.

### GetInterfaceDescriptionOk

`func (o *NetworkDeviceInterface) GetInterfaceDescriptionOk() (*string, bool)`

GetInterfaceDescriptionOk returns a tuple with the InterfaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceDescription

`func (o *NetworkDeviceInterface) SetInterfaceDescription(v string)`

SetInterfaceDescription sets InterfaceDescription field to given value.

### HasInterfaceDescription

`func (o *NetworkDeviceInterface) HasInterfaceDescription() bool`

HasInterfaceDescription returns a boolean if a field has been set.

### GetKind

`func (o *NetworkDeviceInterface) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkDeviceInterface) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkDeviceInterface) SetKind(v string)`

SetKind sets Kind field to given value.


### GetParentInterfaceId

`func (o *NetworkDeviceInterface) GetParentInterfaceId() int64`

GetParentInterfaceId returns the ParentInterfaceId field if non-nil, zero value otherwise.

### GetParentInterfaceIdOk

`func (o *NetworkDeviceInterface) GetParentInterfaceIdOk() (*int64, bool)`

GetParentInterfaceIdOk returns a tuple with the ParentInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterfaceId

`func (o *NetworkDeviceInterface) SetParentInterfaceId(v int64)`

SetParentInterfaceId sets ParentInterfaceId field to given value.

### HasParentInterfaceId

`func (o *NetworkDeviceInterface) HasParentInterfaceId() bool`

HasParentInterfaceId returns a boolean if a field has been set.

### GetLinkedInterfaceId

`func (o *NetworkDeviceInterface) GetLinkedInterfaceId() int64`

GetLinkedInterfaceId returns the LinkedInterfaceId field if non-nil, zero value otherwise.

### GetLinkedInterfaceIdOk

`func (o *NetworkDeviceInterface) GetLinkedInterfaceIdOk() (*int64, bool)`

GetLinkedInterfaceIdOk returns a tuple with the LinkedInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInterfaceId

`func (o *NetworkDeviceInterface) SetLinkedInterfaceId(v int64)`

SetLinkedInterfaceId sets LinkedInterfaceId field to given value.

### HasLinkedInterfaceId

`func (o *NetworkDeviceInterface) HasLinkedInterfaceId() bool`

HasLinkedInterfaceId returns a boolean if a field has been set.

### GetLinkedPortId

`func (o *NetworkDeviceInterface) GetLinkedPortId() string`

GetLinkedPortId returns the LinkedPortId field if non-nil, zero value otherwise.

### GetLinkedPortIdOk

`func (o *NetworkDeviceInterface) GetLinkedPortIdOk() (*string, bool)`

GetLinkedPortIdOk returns a tuple with the LinkedPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedPortId

`func (o *NetworkDeviceInterface) SetLinkedPortId(v string)`

SetLinkedPortId sets LinkedPortId field to given value.

### HasLinkedPortId

`func (o *NetworkDeviceInterface) HasLinkedPortId() bool`

HasLinkedPortId returns a boolean if a field has been set.

### GetLinkedSwitchHostname

`func (o *NetworkDeviceInterface) GetLinkedSwitchHostname() string`

GetLinkedSwitchHostname returns the LinkedSwitchHostname field if non-nil, zero value otherwise.

### GetLinkedSwitchHostnameOk

`func (o *NetworkDeviceInterface) GetLinkedSwitchHostnameOk() (*string, bool)`

GetLinkedSwitchHostnameOk returns a tuple with the LinkedSwitchHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedSwitchHostname

`func (o *NetworkDeviceInterface) SetLinkedSwitchHostname(v string)`

SetLinkedSwitchHostname sets LinkedSwitchHostname field to given value.

### HasLinkedSwitchHostname

`func (o *NetworkDeviceInterface) HasLinkedSwitchHostname() bool`

HasLinkedSwitchHostname returns a boolean if a field has been set.

### GetInterfaceIndex

`func (o *NetworkDeviceInterface) GetInterfaceIndex() float32`

GetInterfaceIndex returns the InterfaceIndex field if non-nil, zero value otherwise.

### GetInterfaceIndexOk

`func (o *NetworkDeviceInterface) GetInterfaceIndexOk() (*float32, bool)`

GetInterfaceIndexOk returns a tuple with the InterfaceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIndex

`func (o *NetworkDeviceInterface) SetInterfaceIndex(v float32)`

SetInterfaceIndex sets InterfaceIndex field to given value.

### HasInterfaceIndex

`func (o *NetworkDeviceInterface) HasInterfaceIndex() bool`

HasInterfaceIndex returns a boolean if a field has been set.

### GetLagIdentifier

`func (o *NetworkDeviceInterface) GetLagIdentifier() float32`

GetLagIdentifier returns the LagIdentifier field if non-nil, zero value otherwise.

### GetLagIdentifierOk

`func (o *NetworkDeviceInterface) GetLagIdentifierOk() (*float32, bool)`

GetLagIdentifierOk returns a tuple with the LagIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagIdentifier

`func (o *NetworkDeviceInterface) SetLagIdentifier(v float32)`

SetLagIdentifier sets LagIdentifier field to given value.

### HasLagIdentifier

`func (o *NetworkDeviceInterface) HasLagIdentifier() bool`

HasLagIdentifier returns a boolean if a field has been set.

### GetNetworkDeviceLag

`func (o *NetworkDeviceInterface) GetNetworkDeviceLag() NetworkDeviceLag`

GetNetworkDeviceLag returns the NetworkDeviceLag field if non-nil, zero value otherwise.

### GetNetworkDeviceLagOk

`func (o *NetworkDeviceInterface) GetNetworkDeviceLagOk() (*NetworkDeviceLag, bool)`

GetNetworkDeviceLagOk returns a tuple with the NetworkDeviceLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceLag

`func (o *NetworkDeviceInterface) SetNetworkDeviceLag(v NetworkDeviceLag)`

SetNetworkDeviceLag sets NetworkDeviceLag field to given value.

### HasNetworkDeviceLag

`func (o *NetworkDeviceInterface) HasNetworkDeviceLag() bool`

HasNetworkDeviceLag returns a boolean if a field has been set.

### SetNetworkDeviceLagNil

`func (o *NetworkDeviceInterface) SetNetworkDeviceLagNil(b bool)`

 SetNetworkDeviceLagNil sets the value for NetworkDeviceLag to be an explicit nil

### UnsetNetworkDeviceLag
`func (o *NetworkDeviceInterface) UnsetNetworkDeviceLag()`

UnsetNetworkDeviceLag ensures that no value is present for NetworkDeviceLag, not even an explicit nil
### GetServerInterfaceId

`func (o *NetworkDeviceInterface) GetServerInterfaceId() int64`

GetServerInterfaceId returns the ServerInterfaceId field if non-nil, zero value otherwise.

### GetServerInterfaceIdOk

`func (o *NetworkDeviceInterface) GetServerInterfaceIdOk() (*int64, bool)`

GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaceId

`func (o *NetworkDeviceInterface) SetServerInterfaceId(v int64)`

SetServerInterfaceId sets ServerInterfaceId field to given value.

### HasServerInterfaceId

`func (o *NetworkDeviceInterface) HasServerInterfaceId() bool`

HasServerInterfaceId returns a boolean if a field has been set.

### GetServerId

`func (o *NetworkDeviceInterface) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *NetworkDeviceInterface) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *NetworkDeviceInterface) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *NetworkDeviceInterface) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetNumaNode

`func (o *NetworkDeviceInterface) GetNumaNode() float32`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *NetworkDeviceInterface) GetNumaNodeOk() (*float32, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *NetworkDeviceInterface) SetNumaNode(v float32)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *NetworkDeviceInterface) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.

### GetDirtyBit

`func (o *NetworkDeviceInterface) GetDirtyBit() float32`

GetDirtyBit returns the DirtyBit field if non-nil, zero value otherwise.

### GetDirtyBitOk

`func (o *NetworkDeviceInterface) GetDirtyBitOk() (*float32, bool)`

GetDirtyBitOk returns a tuple with the DirtyBit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirtyBit

`func (o *NetworkDeviceInterface) SetDirtyBit(v float32)`

SetDirtyBit sets DirtyBit field to given value.


### GetLldpInformation

`func (o *NetworkDeviceInterface) GetLldpInformation() string`

GetLldpInformation returns the LldpInformation field if non-nil, zero value otherwise.

### GetLldpInformationOk

`func (o *NetworkDeviceInterface) GetLldpInformationOk() (*string, bool)`

GetLldpInformationOk returns a tuple with the LldpInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpInformation

`func (o *NetworkDeviceInterface) SetLldpInformation(v string)`

SetLldpInformation sets LldpInformation field to given value.

### HasLldpInformation

`func (o *NetworkDeviceInterface) HasLldpInformation() bool`

HasLldpInformation returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkDeviceInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkDeviceInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkDeviceInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkDeviceInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetWwn

`func (o *NetworkDeviceInterface) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *NetworkDeviceInterface) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *NetworkDeviceInterface) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *NetworkDeviceInterface) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetDriverDumpCachedJson

`func (o *NetworkDeviceInterface) GetDriverDumpCachedJson() map[string]interface{}`

GetDriverDumpCachedJson returns the DriverDumpCachedJson field if non-nil, zero value otherwise.

### GetDriverDumpCachedJsonOk

`func (o *NetworkDeviceInterface) GetDriverDumpCachedJsonOk() (*map[string]interface{}, bool)`

GetDriverDumpCachedJsonOk returns a tuple with the DriverDumpCachedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverDumpCachedJson

`func (o *NetworkDeviceInterface) SetDriverDumpCachedJson(v map[string]interface{})`

SetDriverDumpCachedJson sets DriverDumpCachedJson field to given value.

### HasDriverDumpCachedJson

`func (o *NetworkDeviceInterface) HasDriverDumpCachedJson() bool`

HasDriverDumpCachedJson returns a boolean if a field has been set.

### GetCachedUpdatedTimestamp

`func (o *NetworkDeviceInterface) GetCachedUpdatedTimestamp() string`

GetCachedUpdatedTimestamp returns the CachedUpdatedTimestamp field if non-nil, zero value otherwise.

### GetCachedUpdatedTimestampOk

`func (o *NetworkDeviceInterface) GetCachedUpdatedTimestampOk() (*string, bool)`

GetCachedUpdatedTimestampOk returns a tuple with the CachedUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedUpdatedTimestamp

`func (o *NetworkDeviceInterface) SetCachedUpdatedTimestamp(v string)`

SetCachedUpdatedTimestamp sets CachedUpdatedTimestamp field to given value.


### GetInterfaceLLDPInformationServerInterface

`func (o *NetworkDeviceInterface) GetInterfaceLLDPInformationServerInterface() map[string]interface{}`

GetInterfaceLLDPInformationServerInterface returns the InterfaceLLDPInformationServerInterface field if non-nil, zero value otherwise.

### GetInterfaceLLDPInformationServerInterfaceOk

`func (o *NetworkDeviceInterface) GetInterfaceLLDPInformationServerInterfaceOk() (*map[string]interface{}, bool)`

GetInterfaceLLDPInformationServerInterfaceOk returns a tuple with the InterfaceLLDPInformationServerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceLLDPInformationServerInterface

`func (o *NetworkDeviceInterface) SetInterfaceLLDPInformationServerInterface(v map[string]interface{})`

SetInterfaceLLDPInformationServerInterface sets InterfaceLLDPInformationServerInterface field to given value.

### HasInterfaceLLDPInformationServerInterface

`func (o *NetworkDeviceInterface) HasInterfaceLLDPInformationServerInterface() bool`

HasInterfaceLLDPInformationServerInterface returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkDeviceInterface) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceInterface) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceInterface) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceInterface) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTags

`func (o *NetworkDeviceInterface) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkDeviceInterface) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkDeviceInterface) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetConfig

`func (o *NetworkDeviceInterface) GetConfig() NetworkEquipmentInterfaceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkDeviceInterface) GetConfigOk() (*NetworkEquipmentInterfaceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkDeviceInterface) SetConfig(v NetworkEquipmentInterfaceConfig)`

SetConfig sets Config field to given value.


### GetIpv4

`func (o *NetworkDeviceInterface) GetIpv4() NetworkEquipmentInterfaceIpFamily`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *NetworkDeviceInterface) GetIpv4Ok() (*NetworkEquipmentInterfaceIpFamily, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *NetworkDeviceInterface) SetIpv4(v NetworkEquipmentInterfaceIpFamily)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *NetworkDeviceInterface) GetIpv6() NetworkEquipmentInterfaceIpFamily`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *NetworkDeviceInterface) GetIpv6Ok() (*NetworkEquipmentInterfaceIpFamily, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *NetworkDeviceInterface) SetIpv6(v NetworkEquipmentInterfaceIpFamily)`

SetIpv6 sets Ipv6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


