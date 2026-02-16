# NetworkDeviceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | **float32** | The ID of the network equipment interface | 
**NetworkDeviceId** | **float32** | The ID of the network device | 
**InterfaceName** | **string** | The name of the network equipment interface | 
**InterfaceIndex** | Pointer to **float32** | The index of the network interface | [optional] 
**LagIdentifier** | Pointer to **float32** | LAG identifier | [optional] 
**ServerInterfaceId** | Pointer to **float32** | The server interface ID | [optional] 
**ServerId** | Pointer to **float32** | The server ID | [optional] 
**DirtyBit** | **float32** | Dirty bit flag | 
**LldpInformation** | Pointer to **string** | LLDP information | [optional] 
**MacAddress** | Pointer to **string** | MAC address | [optional] 
**Wwn** | Pointer to **string** | WWPN | [optional] 
**DriverDumpCachedJson** | Pointer to **map[string]interface{}** | Driver dump cached JSON | [optional] 
**CachedUpdatedTimestamp** | **string** | Cached update timestamp | 
**InterfaceLLDPInformationServerInterface** | Pointer to **map[string]interface{}** | LLDP information of server interface in JSON format | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceInterface

`func NewNetworkDeviceInterface(interfaceId float32, networkDeviceId float32, interfaceName string, dirtyBit float32, cachedUpdatedTimestamp string, ) *NetworkDeviceInterface`

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

`func (o *NetworkDeviceInterface) GetInterfaceId() float32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkDeviceInterface) GetInterfaceIdOk() (*float32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkDeviceInterface) SetInterfaceId(v float32)`

SetInterfaceId sets InterfaceId field to given value.


### GetNetworkDeviceId

`func (o *NetworkDeviceInterface) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkDeviceInterface) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkDeviceInterface) SetNetworkDeviceId(v float32)`

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

### GetServerInterfaceId

`func (o *NetworkDeviceInterface) GetServerInterfaceId() float32`

GetServerInterfaceId returns the ServerInterfaceId field if non-nil, zero value otherwise.

### GetServerInterfaceIdOk

`func (o *NetworkDeviceInterface) GetServerInterfaceIdOk() (*float32, bool)`

GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaceId

`func (o *NetworkDeviceInterface) SetServerInterfaceId(v float32)`

SetServerInterfaceId sets ServerInterfaceId field to given value.

### HasServerInterfaceId

`func (o *NetworkDeviceInterface) HasServerInterfaceId() bool`

HasServerInterfaceId returns a boolean if a field has been set.

### GetServerId

`func (o *NetworkDeviceInterface) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *NetworkDeviceInterface) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *NetworkDeviceInterface) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *NetworkDeviceInterface) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


