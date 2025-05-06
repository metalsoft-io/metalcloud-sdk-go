# NetworkDeviceInterfaceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | **string** | The ID of the network equipment interface | 
**NetworkDeviceId** | **string** | The ID of the network device | 
**InterfaceName** | **string** | The name of the network equipment interface | 
**InterfaceIndex** | Pointer to **float32** | The index of the network interface | [optional] 
**LagIdentifier** | Pointer to **float32** | LAG identifier | [optional] 
**ServerInterfaceId** | Pointer to **float32** | The server interface ID | [optional] 
**ServerId** | Pointer to **float32** | The server ID | [optional] 
**DirtyBit** | **float32** | Dirty bit flag | 
**LldpInformation** | Pointer to **string** | LLDP information | [optional] 
**MacAddress** | Pointer to **string** | MAC address | [optional] 
**DriverDumpCachedJson** | Pointer to **map[string]interface{}** | Driver dump cached JSON | [optional] 
**CachedUpdatedTimestamp** | **string** | Cached update timestamp | 
**InterfaceLLDPInformationServerInterface** | Pointer to **map[string]interface{}** | LLDP information of server interface in JSON format | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceInterfaceDto

`func NewNetworkDeviceInterfaceDto(interfaceId string, networkDeviceId string, interfaceName string, dirtyBit float32, cachedUpdatedTimestamp string, ) *NetworkDeviceInterfaceDto`

NewNetworkDeviceInterfaceDto instantiates a new NetworkDeviceInterfaceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceInterfaceDtoWithDefaults

`func NewNetworkDeviceInterfaceDtoWithDefaults() *NetworkDeviceInterfaceDto`

NewNetworkDeviceInterfaceDtoWithDefaults instantiates a new NetworkDeviceInterfaceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *NetworkDeviceInterfaceDto) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkDeviceInterfaceDto) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkDeviceInterfaceDto) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.


### GetNetworkDeviceId

`func (o *NetworkDeviceInterfaceDto) GetNetworkDeviceId() string`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkDeviceInterfaceDto) GetNetworkDeviceIdOk() (*string, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkDeviceInterfaceDto) SetNetworkDeviceId(v string)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetInterfaceName

`func (o *NetworkDeviceInterfaceDto) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NetworkDeviceInterfaceDto) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NetworkDeviceInterfaceDto) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetInterfaceIndex

`func (o *NetworkDeviceInterfaceDto) GetInterfaceIndex() float32`

GetInterfaceIndex returns the InterfaceIndex field if non-nil, zero value otherwise.

### GetInterfaceIndexOk

`func (o *NetworkDeviceInterfaceDto) GetInterfaceIndexOk() (*float32, bool)`

GetInterfaceIndexOk returns a tuple with the InterfaceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIndex

`func (o *NetworkDeviceInterfaceDto) SetInterfaceIndex(v float32)`

SetInterfaceIndex sets InterfaceIndex field to given value.

### HasInterfaceIndex

`func (o *NetworkDeviceInterfaceDto) HasInterfaceIndex() bool`

HasInterfaceIndex returns a boolean if a field has been set.

### GetLagIdentifier

`func (o *NetworkDeviceInterfaceDto) GetLagIdentifier() float32`

GetLagIdentifier returns the LagIdentifier field if non-nil, zero value otherwise.

### GetLagIdentifierOk

`func (o *NetworkDeviceInterfaceDto) GetLagIdentifierOk() (*float32, bool)`

GetLagIdentifierOk returns a tuple with the LagIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagIdentifier

`func (o *NetworkDeviceInterfaceDto) SetLagIdentifier(v float32)`

SetLagIdentifier sets LagIdentifier field to given value.

### HasLagIdentifier

`func (o *NetworkDeviceInterfaceDto) HasLagIdentifier() bool`

HasLagIdentifier returns a boolean if a field has been set.

### GetServerInterfaceId

`func (o *NetworkDeviceInterfaceDto) GetServerInterfaceId() float32`

GetServerInterfaceId returns the ServerInterfaceId field if non-nil, zero value otherwise.

### GetServerInterfaceIdOk

`func (o *NetworkDeviceInterfaceDto) GetServerInterfaceIdOk() (*float32, bool)`

GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaceId

`func (o *NetworkDeviceInterfaceDto) SetServerInterfaceId(v float32)`

SetServerInterfaceId sets ServerInterfaceId field to given value.

### HasServerInterfaceId

`func (o *NetworkDeviceInterfaceDto) HasServerInterfaceId() bool`

HasServerInterfaceId returns a boolean if a field has been set.

### GetServerId

`func (o *NetworkDeviceInterfaceDto) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *NetworkDeviceInterfaceDto) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *NetworkDeviceInterfaceDto) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *NetworkDeviceInterfaceDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetDirtyBit

`func (o *NetworkDeviceInterfaceDto) GetDirtyBit() float32`

GetDirtyBit returns the DirtyBit field if non-nil, zero value otherwise.

### GetDirtyBitOk

`func (o *NetworkDeviceInterfaceDto) GetDirtyBitOk() (*float32, bool)`

GetDirtyBitOk returns a tuple with the DirtyBit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirtyBit

`func (o *NetworkDeviceInterfaceDto) SetDirtyBit(v float32)`

SetDirtyBit sets DirtyBit field to given value.


### GetLldpInformation

`func (o *NetworkDeviceInterfaceDto) GetLldpInformation() string`

GetLldpInformation returns the LldpInformation field if non-nil, zero value otherwise.

### GetLldpInformationOk

`func (o *NetworkDeviceInterfaceDto) GetLldpInformationOk() (*string, bool)`

GetLldpInformationOk returns a tuple with the LldpInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpInformation

`func (o *NetworkDeviceInterfaceDto) SetLldpInformation(v string)`

SetLldpInformation sets LldpInformation field to given value.

### HasLldpInformation

`func (o *NetworkDeviceInterfaceDto) HasLldpInformation() bool`

HasLldpInformation returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkDeviceInterfaceDto) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkDeviceInterfaceDto) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkDeviceInterfaceDto) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkDeviceInterfaceDto) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDriverDumpCachedJson

`func (o *NetworkDeviceInterfaceDto) GetDriverDumpCachedJson() map[string]interface{}`

GetDriverDumpCachedJson returns the DriverDumpCachedJson field if non-nil, zero value otherwise.

### GetDriverDumpCachedJsonOk

`func (o *NetworkDeviceInterfaceDto) GetDriverDumpCachedJsonOk() (*map[string]interface{}, bool)`

GetDriverDumpCachedJsonOk returns a tuple with the DriverDumpCachedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverDumpCachedJson

`func (o *NetworkDeviceInterfaceDto) SetDriverDumpCachedJson(v map[string]interface{})`

SetDriverDumpCachedJson sets DriverDumpCachedJson field to given value.

### HasDriverDumpCachedJson

`func (o *NetworkDeviceInterfaceDto) HasDriverDumpCachedJson() bool`

HasDriverDumpCachedJson returns a boolean if a field has been set.

### GetCachedUpdatedTimestamp

`func (o *NetworkDeviceInterfaceDto) GetCachedUpdatedTimestamp() string`

GetCachedUpdatedTimestamp returns the CachedUpdatedTimestamp field if non-nil, zero value otherwise.

### GetCachedUpdatedTimestampOk

`func (o *NetworkDeviceInterfaceDto) GetCachedUpdatedTimestampOk() (*string, bool)`

GetCachedUpdatedTimestampOk returns a tuple with the CachedUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedUpdatedTimestamp

`func (o *NetworkDeviceInterfaceDto) SetCachedUpdatedTimestamp(v string)`

SetCachedUpdatedTimestamp sets CachedUpdatedTimestamp field to given value.


### GetInterfaceLLDPInformationServerInterface

`func (o *NetworkDeviceInterfaceDto) GetInterfaceLLDPInformationServerInterface() map[string]interface{}`

GetInterfaceLLDPInformationServerInterface returns the InterfaceLLDPInformationServerInterface field if non-nil, zero value otherwise.

### GetInterfaceLLDPInformationServerInterfaceOk

`func (o *NetworkDeviceInterfaceDto) GetInterfaceLLDPInformationServerInterfaceOk() (*map[string]interface{}, bool)`

GetInterfaceLLDPInformationServerInterfaceOk returns a tuple with the InterfaceLLDPInformationServerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceLLDPInformationServerInterface

`func (o *NetworkDeviceInterfaceDto) SetInterfaceLLDPInformationServerInterface(v map[string]interface{})`

SetInterfaceLLDPInformationServerInterface sets InterfaceLLDPInformationServerInterface field to given value.

### HasInterfaceLLDPInformationServerInterface

`func (o *NetworkDeviceInterfaceDto) HasInterfaceLLDPInformationServerInterface() bool`

HasInterfaceLLDPInformationServerInterface returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkDeviceInterfaceDto) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceInterfaceDto) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceInterfaceDto) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceInterfaceDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


