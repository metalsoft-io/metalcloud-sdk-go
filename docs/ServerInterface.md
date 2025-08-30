# ServerInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerInterfaceId** | **float32** | The id of the server interface. | 
**ServerId** | **float32** | The id of the server. | 
**MacAddress** | **string** | The MAC address of the server interface. | 
**Index** | **float32** | The index of the server interface. | 
**SwitchPortId** | Pointer to **string** | The switch port id of the server interface. | [optional] 
**SwitchHostname** | Pointer to **string** | The switch hostname of the server interface. | [optional] 
**CapacityMbps** | Pointer to **float32** | The capacity in Mbps of the server interface. | [optional] 
**AddOnMac** | Pointer to **string** | The add-on MAC address of the server interface. | [optional] 
**AddOnType** | Pointer to **string** | The add-on type of the server interface. | [optional] 
**AddOnInfo** | Pointer to **map[string]interface{}** | The add-on info of the server interface. | [optional] 
**PxeEnabled** | Pointer to **float32** | Flag to indicate if PXE is enabled. | [optional] 
**SupportsIscsiBoot** | Pointer to **float32** | Flag to indicate if the server interface supports iSCSI boot | [optional] 
**FibreChannelWwpn** | Pointer to **string** | The WWPN of the fibre channel. | [optional] 
**Description** | Pointer to **string** | The description of the server interface. | [optional] 
**AliasIndex** | Pointer to **float32** | The alias index of the server interface. | [optional] 
**OsInfo** | **string** | The OS info of the server interface. | [default to "not_tested"]
**NetworkDevice** | Pointer to **map[string]interface{}** | The network device linked to the server. | [optional] 
**NetworkDeviceInterface** | Pointer to **map[string]interface{}** | The network device interface linked to the server. | [optional] 
**Ipv4Addresses** | Pointer to **[]string** | The deployed IPv4 addresses of the server interface. | [optional] 
**Ipv6Addresses** | Pointer to **[]string** | The deployed IPv6 addresses of the server interface. | [optional] 
**VlanId** | Pointer to **float32** | The deployed VLAN ID of the server interface. | [optional] 
**DefaultFabricId** | Pointer to **float32** | The default fabric ID of the server interface. | [optional] 
**RedundancyGroupIndex** | Pointer to **float32** | The redundancy group index of the server interface. | [optional] 
**LldpInfo** | Pointer to **string** | The LLDP information of the server interface. | [optional] 

## Methods

### NewServerInterface

`func NewServerInterface(serverInterfaceId float32, serverId float32, macAddress string, index float32, osInfo string, ) *ServerInterface`

NewServerInterface instantiates a new ServerInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInterfaceWithDefaults

`func NewServerInterfaceWithDefaults() *ServerInterface`

NewServerInterfaceWithDefaults instantiates a new ServerInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerInterfaceId

`func (o *ServerInterface) GetServerInterfaceId() float32`

GetServerInterfaceId returns the ServerInterfaceId field if non-nil, zero value otherwise.

### GetServerInterfaceIdOk

`func (o *ServerInterface) GetServerInterfaceIdOk() (*float32, bool)`

GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaceId

`func (o *ServerInterface) SetServerInterfaceId(v float32)`

SetServerInterfaceId sets ServerInterfaceId field to given value.


### GetServerId

`func (o *ServerInterface) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerInterface) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerInterface) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetMacAddress

`func (o *ServerInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ServerInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ServerInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetIndex

`func (o *ServerInterface) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServerInterface) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServerInterface) SetIndex(v float32)`

SetIndex sets Index field to given value.


### GetSwitchPortId

`func (o *ServerInterface) GetSwitchPortId() string`

GetSwitchPortId returns the SwitchPortId field if non-nil, zero value otherwise.

### GetSwitchPortIdOk

`func (o *ServerInterface) GetSwitchPortIdOk() (*string, bool)`

GetSwitchPortIdOk returns a tuple with the SwitchPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortId

`func (o *ServerInterface) SetSwitchPortId(v string)`

SetSwitchPortId sets SwitchPortId field to given value.

### HasSwitchPortId

`func (o *ServerInterface) HasSwitchPortId() bool`

HasSwitchPortId returns a boolean if a field has been set.

### GetSwitchHostname

`func (o *ServerInterface) GetSwitchHostname() string`

GetSwitchHostname returns the SwitchHostname field if non-nil, zero value otherwise.

### GetSwitchHostnameOk

`func (o *ServerInterface) GetSwitchHostnameOk() (*string, bool)`

GetSwitchHostnameOk returns a tuple with the SwitchHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchHostname

`func (o *ServerInterface) SetSwitchHostname(v string)`

SetSwitchHostname sets SwitchHostname field to given value.

### HasSwitchHostname

`func (o *ServerInterface) HasSwitchHostname() bool`

HasSwitchHostname returns a boolean if a field has been set.

### GetCapacityMbps

`func (o *ServerInterface) GetCapacityMbps() float32`

GetCapacityMbps returns the CapacityMbps field if non-nil, zero value otherwise.

### GetCapacityMbpsOk

`func (o *ServerInterface) GetCapacityMbpsOk() (*float32, bool)`

GetCapacityMbpsOk returns a tuple with the CapacityMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMbps

`func (o *ServerInterface) SetCapacityMbps(v float32)`

SetCapacityMbps sets CapacityMbps field to given value.

### HasCapacityMbps

`func (o *ServerInterface) HasCapacityMbps() bool`

HasCapacityMbps returns a boolean if a field has been set.

### GetAddOnMac

`func (o *ServerInterface) GetAddOnMac() string`

GetAddOnMac returns the AddOnMac field if non-nil, zero value otherwise.

### GetAddOnMacOk

`func (o *ServerInterface) GetAddOnMacOk() (*string, bool)`

GetAddOnMacOk returns a tuple with the AddOnMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnMac

`func (o *ServerInterface) SetAddOnMac(v string)`

SetAddOnMac sets AddOnMac field to given value.

### HasAddOnMac

`func (o *ServerInterface) HasAddOnMac() bool`

HasAddOnMac returns a boolean if a field has been set.

### GetAddOnType

`func (o *ServerInterface) GetAddOnType() string`

GetAddOnType returns the AddOnType field if non-nil, zero value otherwise.

### GetAddOnTypeOk

`func (o *ServerInterface) GetAddOnTypeOk() (*string, bool)`

GetAddOnTypeOk returns a tuple with the AddOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnType

`func (o *ServerInterface) SetAddOnType(v string)`

SetAddOnType sets AddOnType field to given value.

### HasAddOnType

`func (o *ServerInterface) HasAddOnType() bool`

HasAddOnType returns a boolean if a field has been set.

### GetAddOnInfo

`func (o *ServerInterface) GetAddOnInfo() map[string]interface{}`

GetAddOnInfo returns the AddOnInfo field if non-nil, zero value otherwise.

### GetAddOnInfoOk

`func (o *ServerInterface) GetAddOnInfoOk() (*map[string]interface{}, bool)`

GetAddOnInfoOk returns a tuple with the AddOnInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnInfo

`func (o *ServerInterface) SetAddOnInfo(v map[string]interface{})`

SetAddOnInfo sets AddOnInfo field to given value.

### HasAddOnInfo

`func (o *ServerInterface) HasAddOnInfo() bool`

HasAddOnInfo returns a boolean if a field has been set.

### GetPxeEnabled

`func (o *ServerInterface) GetPxeEnabled() float32`

GetPxeEnabled returns the PxeEnabled field if non-nil, zero value otherwise.

### GetPxeEnabledOk

`func (o *ServerInterface) GetPxeEnabledOk() (*float32, bool)`

GetPxeEnabledOk returns a tuple with the PxeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeEnabled

`func (o *ServerInterface) SetPxeEnabled(v float32)`

SetPxeEnabled sets PxeEnabled field to given value.

### HasPxeEnabled

`func (o *ServerInterface) HasPxeEnabled() bool`

HasPxeEnabled returns a boolean if a field has been set.

### GetSupportsIscsiBoot

`func (o *ServerInterface) GetSupportsIscsiBoot() float32`

GetSupportsIscsiBoot returns the SupportsIscsiBoot field if non-nil, zero value otherwise.

### GetSupportsIscsiBootOk

`func (o *ServerInterface) GetSupportsIscsiBootOk() (*float32, bool)`

GetSupportsIscsiBootOk returns a tuple with the SupportsIscsiBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIscsiBoot

`func (o *ServerInterface) SetSupportsIscsiBoot(v float32)`

SetSupportsIscsiBoot sets SupportsIscsiBoot field to given value.

### HasSupportsIscsiBoot

`func (o *ServerInterface) HasSupportsIscsiBoot() bool`

HasSupportsIscsiBoot returns a boolean if a field has been set.

### GetFibreChannelWwpn

`func (o *ServerInterface) GetFibreChannelWwpn() string`

GetFibreChannelWwpn returns the FibreChannelWwpn field if non-nil, zero value otherwise.

### GetFibreChannelWwpnOk

`func (o *ServerInterface) GetFibreChannelWwpnOk() (*string, bool)`

GetFibreChannelWwpnOk returns a tuple with the FibreChannelWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelWwpn

`func (o *ServerInterface) SetFibreChannelWwpn(v string)`

SetFibreChannelWwpn sets FibreChannelWwpn field to given value.

### HasFibreChannelWwpn

`func (o *ServerInterface) HasFibreChannelWwpn() bool`

HasFibreChannelWwpn returns a boolean if a field has been set.

### GetDescription

`func (o *ServerInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAliasIndex

`func (o *ServerInterface) GetAliasIndex() float32`

GetAliasIndex returns the AliasIndex field if non-nil, zero value otherwise.

### GetAliasIndexOk

`func (o *ServerInterface) GetAliasIndexOk() (*float32, bool)`

GetAliasIndexOk returns a tuple with the AliasIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasIndex

`func (o *ServerInterface) SetAliasIndex(v float32)`

SetAliasIndex sets AliasIndex field to given value.

### HasAliasIndex

`func (o *ServerInterface) HasAliasIndex() bool`

HasAliasIndex returns a boolean if a field has been set.

### GetOsInfo

`func (o *ServerInterface) GetOsInfo() string`

GetOsInfo returns the OsInfo field if non-nil, zero value otherwise.

### GetOsInfoOk

`func (o *ServerInterface) GetOsInfoOk() (*string, bool)`

GetOsInfoOk returns a tuple with the OsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInfo

`func (o *ServerInterface) SetOsInfo(v string)`

SetOsInfo sets OsInfo field to given value.


### GetNetworkDevice

`func (o *ServerInterface) GetNetworkDevice() map[string]interface{}`

GetNetworkDevice returns the NetworkDevice field if non-nil, zero value otherwise.

### GetNetworkDeviceOk

`func (o *ServerInterface) GetNetworkDeviceOk() (*map[string]interface{}, bool)`

GetNetworkDeviceOk returns a tuple with the NetworkDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevice

`func (o *ServerInterface) SetNetworkDevice(v map[string]interface{})`

SetNetworkDevice sets NetworkDevice field to given value.

### HasNetworkDevice

`func (o *ServerInterface) HasNetworkDevice() bool`

HasNetworkDevice returns a boolean if a field has been set.

### GetNetworkDeviceInterface

`func (o *ServerInterface) GetNetworkDeviceInterface() map[string]interface{}`

GetNetworkDeviceInterface returns the NetworkDeviceInterface field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceOk

`func (o *ServerInterface) GetNetworkDeviceInterfaceOk() (*map[string]interface{}, bool)`

GetNetworkDeviceInterfaceOk returns a tuple with the NetworkDeviceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterface

`func (o *ServerInterface) SetNetworkDeviceInterface(v map[string]interface{})`

SetNetworkDeviceInterface sets NetworkDeviceInterface field to given value.

### HasNetworkDeviceInterface

`func (o *ServerInterface) HasNetworkDeviceInterface() bool`

HasNetworkDeviceInterface returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *ServerInterface) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *ServerInterface) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *ServerInterface) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *ServerInterface) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *ServerInterface) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *ServerInterface) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *ServerInterface) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *ServerInterface) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetVlanId

`func (o *ServerInterface) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ServerInterface) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ServerInterface) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ServerInterface) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetDefaultFabricId

`func (o *ServerInterface) GetDefaultFabricId() float32`

GetDefaultFabricId returns the DefaultFabricId field if non-nil, zero value otherwise.

### GetDefaultFabricIdOk

`func (o *ServerInterface) GetDefaultFabricIdOk() (*float32, bool)`

GetDefaultFabricIdOk returns a tuple with the DefaultFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFabricId

`func (o *ServerInterface) SetDefaultFabricId(v float32)`

SetDefaultFabricId sets DefaultFabricId field to given value.

### HasDefaultFabricId

`func (o *ServerInterface) HasDefaultFabricId() bool`

HasDefaultFabricId returns a boolean if a field has been set.

### GetRedundancyGroupIndex

`func (o *ServerInterface) GetRedundancyGroupIndex() float32`

GetRedundancyGroupIndex returns the RedundancyGroupIndex field if non-nil, zero value otherwise.

### GetRedundancyGroupIndexOk

`func (o *ServerInterface) GetRedundancyGroupIndexOk() (*float32, bool)`

GetRedundancyGroupIndexOk returns a tuple with the RedundancyGroupIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyGroupIndex

`func (o *ServerInterface) SetRedundancyGroupIndex(v float32)`

SetRedundancyGroupIndex sets RedundancyGroupIndex field to given value.

### HasRedundancyGroupIndex

`func (o *ServerInterface) HasRedundancyGroupIndex() bool`

HasRedundancyGroupIndex returns a boolean if a field has been set.

### GetLldpInfo

`func (o *ServerInterface) GetLldpInfo() string`

GetLldpInfo returns the LldpInfo field if non-nil, zero value otherwise.

### GetLldpInfoOk

`func (o *ServerInterface) GetLldpInfoOk() (*string, bool)`

GetLldpInfoOk returns a tuple with the LldpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpInfo

`func (o *ServerInterface) SetLldpInfo(v string)`

SetLldpInfo sets LldpInfo field to given value.

### HasLldpInfo

`func (o *ServerInterface) HasLldpInfo() bool`

HasLldpInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


