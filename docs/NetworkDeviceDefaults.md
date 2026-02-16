# NetworkDeviceDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID of the network device defaults entry | 
**DatacenterName** | **string** | Name of the datacenter | 
**SerialNumber** | Pointer to **string** | Device serial number | [optional] 
**ManagementMacAddress** | Pointer to **string** | Management MAC address | [optional] 
**Position** | Pointer to **string** | Physical position in the rack | [optional] 
**IdentifierString** | Pointer to **string** | Identifier string (1-63 characters, letters, digits, . _ -) | [optional] 
**Asn** | Pointer to **int64** | Autonomous System Number (ASN) of the network device | [optional] 
**IsPartOfMlagPair** | Pointer to **bool** | Whether device is part of an MLAG pair | [optional] 
**MlagSystemMac** | Pointer to **string** | MLAG system MAC address | [optional] 
**MlagDomainId** | Pointer to **int32** | MLAG domain ID | [optional] 
**MlagPeerLinkPortChannelId** | Pointer to **int32** | MLAG peer link port-channel ID | [optional] 
**MlagPartnerVlanId** | Pointer to **int32** | MLAG partner VLAN ID | [optional] 
**MlagPartnerHostname** | Pointer to **string** | Hostname of MLAG partner device | [optional] 
**LoopbackAddressIpv4** | Pointer to **string** | Loopback IPv4 address | [optional] 
**LoopbackAddressIpv6** | Pointer to **string** | Loopback IPv6 address | [optional] 
**VtepAddressIpv4** | Pointer to **string** | VTEP IPv4 address | [optional] 
**VtepAddressIpv6** | Pointer to **string** | VTEP IPv6 address | [optional] 
**OsTemplateId** | Pointer to **int32** | Volume template ID | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for device configuration | [optional] 
**OrderIndex** | Pointer to **int32** | Order index for display or processing | [optional] 

## Methods

### NewNetworkDeviceDefaults

`func NewNetworkDeviceDefaults(id int32, datacenterName string, ) *NetworkDeviceDefaults`

NewNetworkDeviceDefaults instantiates a new NetworkDeviceDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDefaultsWithDefaults

`func NewNetworkDeviceDefaultsWithDefaults() *NetworkDeviceDefaults`

NewNetworkDeviceDefaultsWithDefaults instantiates a new NetworkDeviceDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDeviceDefaults) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceDefaults) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceDefaults) SetId(v int32)`

SetId sets Id field to given value.


### GetDatacenterName

`func (o *NetworkDeviceDefaults) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *NetworkDeviceDefaults) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *NetworkDeviceDefaults) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetSerialNumber

`func (o *NetworkDeviceDefaults) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NetworkDeviceDefaults) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NetworkDeviceDefaults) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NetworkDeviceDefaults) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetManagementMacAddress

`func (o *NetworkDeviceDefaults) GetManagementMacAddress() string`

GetManagementMacAddress returns the ManagementMacAddress field if non-nil, zero value otherwise.

### GetManagementMacAddressOk

`func (o *NetworkDeviceDefaults) GetManagementMacAddressOk() (*string, bool)`

GetManagementMacAddressOk returns a tuple with the ManagementMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMacAddress

`func (o *NetworkDeviceDefaults) SetManagementMacAddress(v string)`

SetManagementMacAddress sets ManagementMacAddress field to given value.

### HasManagementMacAddress

`func (o *NetworkDeviceDefaults) HasManagementMacAddress() bool`

HasManagementMacAddress returns a boolean if a field has been set.

### GetPosition

`func (o *NetworkDeviceDefaults) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *NetworkDeviceDefaults) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *NetworkDeviceDefaults) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *NetworkDeviceDefaults) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIdentifierString

`func (o *NetworkDeviceDefaults) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *NetworkDeviceDefaults) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *NetworkDeviceDefaults) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.

### HasIdentifierString

`func (o *NetworkDeviceDefaults) HasIdentifierString() bool`

HasIdentifierString returns a boolean if a field has been set.

### GetAsn

`func (o *NetworkDeviceDefaults) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *NetworkDeviceDefaults) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *NetworkDeviceDefaults) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *NetworkDeviceDefaults) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetIsPartOfMlagPair

`func (o *NetworkDeviceDefaults) GetIsPartOfMlagPair() bool`

GetIsPartOfMlagPair returns the IsPartOfMlagPair field if non-nil, zero value otherwise.

### GetIsPartOfMlagPairOk

`func (o *NetworkDeviceDefaults) GetIsPartOfMlagPairOk() (*bool, bool)`

GetIsPartOfMlagPairOk returns a tuple with the IsPartOfMlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartOfMlagPair

`func (o *NetworkDeviceDefaults) SetIsPartOfMlagPair(v bool)`

SetIsPartOfMlagPair sets IsPartOfMlagPair field to given value.

### HasIsPartOfMlagPair

`func (o *NetworkDeviceDefaults) HasIsPartOfMlagPair() bool`

HasIsPartOfMlagPair returns a boolean if a field has been set.

### GetMlagSystemMac

`func (o *NetworkDeviceDefaults) GetMlagSystemMac() string`

GetMlagSystemMac returns the MlagSystemMac field if non-nil, zero value otherwise.

### GetMlagSystemMacOk

`func (o *NetworkDeviceDefaults) GetMlagSystemMacOk() (*string, bool)`

GetMlagSystemMacOk returns a tuple with the MlagSystemMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagSystemMac

`func (o *NetworkDeviceDefaults) SetMlagSystemMac(v string)`

SetMlagSystemMac sets MlagSystemMac field to given value.

### HasMlagSystemMac

`func (o *NetworkDeviceDefaults) HasMlagSystemMac() bool`

HasMlagSystemMac returns a boolean if a field has been set.

### GetMlagDomainId

`func (o *NetworkDeviceDefaults) GetMlagDomainId() int32`

GetMlagDomainId returns the MlagDomainId field if non-nil, zero value otherwise.

### GetMlagDomainIdOk

`func (o *NetworkDeviceDefaults) GetMlagDomainIdOk() (*int32, bool)`

GetMlagDomainIdOk returns a tuple with the MlagDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagDomainId

`func (o *NetworkDeviceDefaults) SetMlagDomainId(v int32)`

SetMlagDomainId sets MlagDomainId field to given value.

### HasMlagDomainId

`func (o *NetworkDeviceDefaults) HasMlagDomainId() bool`

HasMlagDomainId returns a boolean if a field has been set.

### GetMlagPeerLinkPortChannelId

`func (o *NetworkDeviceDefaults) GetMlagPeerLinkPortChannelId() int32`

GetMlagPeerLinkPortChannelId returns the MlagPeerLinkPortChannelId field if non-nil, zero value otherwise.

### GetMlagPeerLinkPortChannelIdOk

`func (o *NetworkDeviceDefaults) GetMlagPeerLinkPortChannelIdOk() (*int32, bool)`

GetMlagPeerLinkPortChannelIdOk returns a tuple with the MlagPeerLinkPortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerLinkPortChannelId

`func (o *NetworkDeviceDefaults) SetMlagPeerLinkPortChannelId(v int32)`

SetMlagPeerLinkPortChannelId sets MlagPeerLinkPortChannelId field to given value.

### HasMlagPeerLinkPortChannelId

`func (o *NetworkDeviceDefaults) HasMlagPeerLinkPortChannelId() bool`

HasMlagPeerLinkPortChannelId returns a boolean if a field has been set.

### GetMlagPartnerVlanId

`func (o *NetworkDeviceDefaults) GetMlagPartnerVlanId() int32`

GetMlagPartnerVlanId returns the MlagPartnerVlanId field if non-nil, zero value otherwise.

### GetMlagPartnerVlanIdOk

`func (o *NetworkDeviceDefaults) GetMlagPartnerVlanIdOk() (*int32, bool)`

GetMlagPartnerVlanIdOk returns a tuple with the MlagPartnerVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPartnerVlanId

`func (o *NetworkDeviceDefaults) SetMlagPartnerVlanId(v int32)`

SetMlagPartnerVlanId sets MlagPartnerVlanId field to given value.

### HasMlagPartnerVlanId

`func (o *NetworkDeviceDefaults) HasMlagPartnerVlanId() bool`

HasMlagPartnerVlanId returns a boolean if a field has been set.

### GetMlagPartnerHostname

`func (o *NetworkDeviceDefaults) GetMlagPartnerHostname() string`

GetMlagPartnerHostname returns the MlagPartnerHostname field if non-nil, zero value otherwise.

### GetMlagPartnerHostnameOk

`func (o *NetworkDeviceDefaults) GetMlagPartnerHostnameOk() (*string, bool)`

GetMlagPartnerHostnameOk returns a tuple with the MlagPartnerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPartnerHostname

`func (o *NetworkDeviceDefaults) SetMlagPartnerHostname(v string)`

SetMlagPartnerHostname sets MlagPartnerHostname field to given value.

### HasMlagPartnerHostname

`func (o *NetworkDeviceDefaults) HasMlagPartnerHostname() bool`

HasMlagPartnerHostname returns a boolean if a field has been set.

### GetLoopbackAddressIpv4

`func (o *NetworkDeviceDefaults) GetLoopbackAddressIpv4() string`

GetLoopbackAddressIpv4 returns the LoopbackAddressIpv4 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv4Ok

`func (o *NetworkDeviceDefaults) GetLoopbackAddressIpv4Ok() (*string, bool)`

GetLoopbackAddressIpv4Ok returns a tuple with the LoopbackAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv4

`func (o *NetworkDeviceDefaults) SetLoopbackAddressIpv4(v string)`

SetLoopbackAddressIpv4 sets LoopbackAddressIpv4 field to given value.

### HasLoopbackAddressIpv4

`func (o *NetworkDeviceDefaults) HasLoopbackAddressIpv4() bool`

HasLoopbackAddressIpv4 returns a boolean if a field has been set.

### GetLoopbackAddressIpv6

`func (o *NetworkDeviceDefaults) GetLoopbackAddressIpv6() string`

GetLoopbackAddressIpv6 returns the LoopbackAddressIpv6 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv6Ok

`func (o *NetworkDeviceDefaults) GetLoopbackAddressIpv6Ok() (*string, bool)`

GetLoopbackAddressIpv6Ok returns a tuple with the LoopbackAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv6

`func (o *NetworkDeviceDefaults) SetLoopbackAddressIpv6(v string)`

SetLoopbackAddressIpv6 sets LoopbackAddressIpv6 field to given value.

### HasLoopbackAddressIpv6

`func (o *NetworkDeviceDefaults) HasLoopbackAddressIpv6() bool`

HasLoopbackAddressIpv6 returns a boolean if a field has been set.

### GetVtepAddressIpv4

`func (o *NetworkDeviceDefaults) GetVtepAddressIpv4() string`

GetVtepAddressIpv4 returns the VtepAddressIpv4 field if non-nil, zero value otherwise.

### GetVtepAddressIpv4Ok

`func (o *NetworkDeviceDefaults) GetVtepAddressIpv4Ok() (*string, bool)`

GetVtepAddressIpv4Ok returns a tuple with the VtepAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddressIpv4

`func (o *NetworkDeviceDefaults) SetVtepAddressIpv4(v string)`

SetVtepAddressIpv4 sets VtepAddressIpv4 field to given value.

### HasVtepAddressIpv4

`func (o *NetworkDeviceDefaults) HasVtepAddressIpv4() bool`

HasVtepAddressIpv4 returns a boolean if a field has been set.

### GetVtepAddressIpv6

`func (o *NetworkDeviceDefaults) GetVtepAddressIpv6() string`

GetVtepAddressIpv6 returns the VtepAddressIpv6 field if non-nil, zero value otherwise.

### GetVtepAddressIpv6Ok

`func (o *NetworkDeviceDefaults) GetVtepAddressIpv6Ok() (*string, bool)`

GetVtepAddressIpv6Ok returns a tuple with the VtepAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddressIpv6

`func (o *NetworkDeviceDefaults) SetVtepAddressIpv6(v string)`

SetVtepAddressIpv6 sets VtepAddressIpv6 field to given value.

### HasVtepAddressIpv6

`func (o *NetworkDeviceDefaults) HasVtepAddressIpv6() bool`

HasVtepAddressIpv6 returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *NetworkDeviceDefaults) GetOsTemplateId() int32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *NetworkDeviceDefaults) GetOsTemplateIdOk() (*int32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *NetworkDeviceDefaults) SetOsTemplateId(v int32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *NetworkDeviceDefaults) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *NetworkDeviceDefaults) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *NetworkDeviceDefaults) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *NetworkDeviceDefaults) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *NetworkDeviceDefaults) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetOrderIndex

`func (o *NetworkDeviceDefaults) GetOrderIndex() int32`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *NetworkDeviceDefaults) GetOrderIndexOk() (*int32, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *NetworkDeviceDefaults) SetOrderIndex(v int32)`

SetOrderIndex sets OrderIndex field to given value.

### HasOrderIndex

`func (o *NetworkDeviceDefaults) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


