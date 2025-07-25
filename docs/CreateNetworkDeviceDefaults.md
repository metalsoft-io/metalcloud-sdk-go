# CreateNetworkDeviceDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterName** | **string** | Name of the datacenter | 
**SerialNumber** | Pointer to **string** | Serial number of the network device | [optional] 
**ManagementMacAddress** | **string** | MAC address of the network device. Must be in the format XX:XX:XX:XX:XX:XX. | 
**Position** | Pointer to **string** | Position of the network device | [optional] 
**IdentifierString** | Pointer to **string** | Identifier string of the network device. Can contain letters, numbers, dots, underscores, and hyphens (1-63 characters). | [optional] 
**Asn** | Pointer to **float32** | ASN of the network device | [optional] 
**SkipInitialConfiguration** | Pointer to **float32** | Skip initial configuration | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables associated with the network device, stored as key-value pairs. | [optional] 
**MlagDomainId** | Pointer to **float32** | MLAG domain ID, must be between 1 and 4096. | [optional] 
**LoopbackAddressIpv4** | Pointer to **string** | IPv4 address assigned to the loopback interface. | [optional] 
**LoopbackAddressIpv6** | Pointer to **string** | IPv6 address assigned to the loopback interface. | [optional] 
**VtepAddressIpv4** | Pointer to **string** | VTEP (VXLAN Tunnel Endpoint) IPv4 address. | [optional] 
**VtepAddressIpv6** | Pointer to **string** | VTEP (VXLAN Tunnel Endpoint) IPv6 address. | [optional] 
**OrderIndex** | Pointer to **float32** | Order index used for sorting or prioritization. | [optional] 
**VolumeTemplateId** | Pointer to **float32** | ID of the volume template associated with the network device. | [optional] 
**MlagPartnerHostname** | Pointer to **string** | MLAG partner hostname. Can contain letters, numbers, dots, underscores, and hyphens (1-63 characters). | [optional] 
**IsPartOfMlagPair** | Pointer to **float32** | Indicates whether the device is part of an MLAG pair (0 &#x3D; No, 1 &#x3D; Yes). | [optional] 
**MlagSystemMac** | Pointer to **string** | MLAG system MAC address in the format XX:XX:XX:XX:XX:XX. | [optional] 
**MlagPeerLinkPortChannelId** | Pointer to **float32** | MLAG peer link Port-Channel ID, must be between 1 and 4096. | [optional] 
**MlagPartnerVlanId** | Pointer to **float32** | MLAG partner VLAN ID, must be between 1 and 4096. | [optional] 

## Methods

### NewCreateNetworkDeviceDefaults

`func NewCreateNetworkDeviceDefaults(datacenterName string, managementMacAddress string, ) *CreateNetworkDeviceDefaults`

NewCreateNetworkDeviceDefaults instantiates a new CreateNetworkDeviceDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceDefaultsWithDefaults

`func NewCreateNetworkDeviceDefaultsWithDefaults() *CreateNetworkDeviceDefaults`

NewCreateNetworkDeviceDefaultsWithDefaults instantiates a new CreateNetworkDeviceDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterName

`func (o *CreateNetworkDeviceDefaults) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *CreateNetworkDeviceDefaults) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *CreateNetworkDeviceDefaults) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetSerialNumber

`func (o *CreateNetworkDeviceDefaults) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CreateNetworkDeviceDefaults) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CreateNetworkDeviceDefaults) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CreateNetworkDeviceDefaults) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetManagementMacAddress

`func (o *CreateNetworkDeviceDefaults) GetManagementMacAddress() string`

GetManagementMacAddress returns the ManagementMacAddress field if non-nil, zero value otherwise.

### GetManagementMacAddressOk

`func (o *CreateNetworkDeviceDefaults) GetManagementMacAddressOk() (*string, bool)`

GetManagementMacAddressOk returns a tuple with the ManagementMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMacAddress

`func (o *CreateNetworkDeviceDefaults) SetManagementMacAddress(v string)`

SetManagementMacAddress sets ManagementMacAddress field to given value.


### GetPosition

`func (o *CreateNetworkDeviceDefaults) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateNetworkDeviceDefaults) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateNetworkDeviceDefaults) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateNetworkDeviceDefaults) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIdentifierString

`func (o *CreateNetworkDeviceDefaults) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *CreateNetworkDeviceDefaults) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *CreateNetworkDeviceDefaults) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.

### HasIdentifierString

`func (o *CreateNetworkDeviceDefaults) HasIdentifierString() bool`

HasIdentifierString returns a boolean if a field has been set.

### GetAsn

`func (o *CreateNetworkDeviceDefaults) GetAsn() float32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *CreateNetworkDeviceDefaults) GetAsnOk() (*float32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *CreateNetworkDeviceDefaults) SetAsn(v float32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *CreateNetworkDeviceDefaults) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetSkipInitialConfiguration

`func (o *CreateNetworkDeviceDefaults) GetSkipInitialConfiguration() float32`

GetSkipInitialConfiguration returns the SkipInitialConfiguration field if non-nil, zero value otherwise.

### GetSkipInitialConfigurationOk

`func (o *CreateNetworkDeviceDefaults) GetSkipInitialConfigurationOk() (*float32, bool)`

GetSkipInitialConfigurationOk returns a tuple with the SkipInitialConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipInitialConfiguration

`func (o *CreateNetworkDeviceDefaults) SetSkipInitialConfiguration(v float32)`

SetSkipInitialConfiguration sets SkipInitialConfiguration field to given value.

### HasSkipInitialConfiguration

`func (o *CreateNetworkDeviceDefaults) HasSkipInitialConfiguration() bool`

HasSkipInitialConfiguration returns a boolean if a field has been set.

### GetCustomVariables

`func (o *CreateNetworkDeviceDefaults) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *CreateNetworkDeviceDefaults) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *CreateNetworkDeviceDefaults) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *CreateNetworkDeviceDefaults) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetMlagDomainId

`func (o *CreateNetworkDeviceDefaults) GetMlagDomainId() float32`

GetMlagDomainId returns the MlagDomainId field if non-nil, zero value otherwise.

### GetMlagDomainIdOk

`func (o *CreateNetworkDeviceDefaults) GetMlagDomainIdOk() (*float32, bool)`

GetMlagDomainIdOk returns a tuple with the MlagDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagDomainId

`func (o *CreateNetworkDeviceDefaults) SetMlagDomainId(v float32)`

SetMlagDomainId sets MlagDomainId field to given value.

### HasMlagDomainId

`func (o *CreateNetworkDeviceDefaults) HasMlagDomainId() bool`

HasMlagDomainId returns a boolean if a field has been set.

### GetLoopbackAddressIpv4

`func (o *CreateNetworkDeviceDefaults) GetLoopbackAddressIpv4() string`

GetLoopbackAddressIpv4 returns the LoopbackAddressIpv4 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv4Ok

`func (o *CreateNetworkDeviceDefaults) GetLoopbackAddressIpv4Ok() (*string, bool)`

GetLoopbackAddressIpv4Ok returns a tuple with the LoopbackAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv4

`func (o *CreateNetworkDeviceDefaults) SetLoopbackAddressIpv4(v string)`

SetLoopbackAddressIpv4 sets LoopbackAddressIpv4 field to given value.

### HasLoopbackAddressIpv4

`func (o *CreateNetworkDeviceDefaults) HasLoopbackAddressIpv4() bool`

HasLoopbackAddressIpv4 returns a boolean if a field has been set.

### GetLoopbackAddressIpv6

`func (o *CreateNetworkDeviceDefaults) GetLoopbackAddressIpv6() string`

GetLoopbackAddressIpv6 returns the LoopbackAddressIpv6 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv6Ok

`func (o *CreateNetworkDeviceDefaults) GetLoopbackAddressIpv6Ok() (*string, bool)`

GetLoopbackAddressIpv6Ok returns a tuple with the LoopbackAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv6

`func (o *CreateNetworkDeviceDefaults) SetLoopbackAddressIpv6(v string)`

SetLoopbackAddressIpv6 sets LoopbackAddressIpv6 field to given value.

### HasLoopbackAddressIpv6

`func (o *CreateNetworkDeviceDefaults) HasLoopbackAddressIpv6() bool`

HasLoopbackAddressIpv6 returns a boolean if a field has been set.

### GetVtepAddressIpv4

`func (o *CreateNetworkDeviceDefaults) GetVtepAddressIpv4() string`

GetVtepAddressIpv4 returns the VtepAddressIpv4 field if non-nil, zero value otherwise.

### GetVtepAddressIpv4Ok

`func (o *CreateNetworkDeviceDefaults) GetVtepAddressIpv4Ok() (*string, bool)`

GetVtepAddressIpv4Ok returns a tuple with the VtepAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddressIpv4

`func (o *CreateNetworkDeviceDefaults) SetVtepAddressIpv4(v string)`

SetVtepAddressIpv4 sets VtepAddressIpv4 field to given value.

### HasVtepAddressIpv4

`func (o *CreateNetworkDeviceDefaults) HasVtepAddressIpv4() bool`

HasVtepAddressIpv4 returns a boolean if a field has been set.

### GetVtepAddressIpv6

`func (o *CreateNetworkDeviceDefaults) GetVtepAddressIpv6() string`

GetVtepAddressIpv6 returns the VtepAddressIpv6 field if non-nil, zero value otherwise.

### GetVtepAddressIpv6Ok

`func (o *CreateNetworkDeviceDefaults) GetVtepAddressIpv6Ok() (*string, bool)`

GetVtepAddressIpv6Ok returns a tuple with the VtepAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddressIpv6

`func (o *CreateNetworkDeviceDefaults) SetVtepAddressIpv6(v string)`

SetVtepAddressIpv6 sets VtepAddressIpv6 field to given value.

### HasVtepAddressIpv6

`func (o *CreateNetworkDeviceDefaults) HasVtepAddressIpv6() bool`

HasVtepAddressIpv6 returns a boolean if a field has been set.

### GetOrderIndex

`func (o *CreateNetworkDeviceDefaults) GetOrderIndex() float32`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *CreateNetworkDeviceDefaults) GetOrderIndexOk() (*float32, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *CreateNetworkDeviceDefaults) SetOrderIndex(v float32)`

SetOrderIndex sets OrderIndex field to given value.

### HasOrderIndex

`func (o *CreateNetworkDeviceDefaults) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.

### GetVolumeTemplateId

`func (o *CreateNetworkDeviceDefaults) GetVolumeTemplateId() float32`

GetVolumeTemplateId returns the VolumeTemplateId field if non-nil, zero value otherwise.

### GetVolumeTemplateIdOk

`func (o *CreateNetworkDeviceDefaults) GetVolumeTemplateIdOk() (*float32, bool)`

GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplateId

`func (o *CreateNetworkDeviceDefaults) SetVolumeTemplateId(v float32)`

SetVolumeTemplateId sets VolumeTemplateId field to given value.

### HasVolumeTemplateId

`func (o *CreateNetworkDeviceDefaults) HasVolumeTemplateId() bool`

HasVolumeTemplateId returns a boolean if a field has been set.

### GetMlagPartnerHostname

`func (o *CreateNetworkDeviceDefaults) GetMlagPartnerHostname() string`

GetMlagPartnerHostname returns the MlagPartnerHostname field if non-nil, zero value otherwise.

### GetMlagPartnerHostnameOk

`func (o *CreateNetworkDeviceDefaults) GetMlagPartnerHostnameOk() (*string, bool)`

GetMlagPartnerHostnameOk returns a tuple with the MlagPartnerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPartnerHostname

`func (o *CreateNetworkDeviceDefaults) SetMlagPartnerHostname(v string)`

SetMlagPartnerHostname sets MlagPartnerHostname field to given value.

### HasMlagPartnerHostname

`func (o *CreateNetworkDeviceDefaults) HasMlagPartnerHostname() bool`

HasMlagPartnerHostname returns a boolean if a field has been set.

### GetIsPartOfMlagPair

`func (o *CreateNetworkDeviceDefaults) GetIsPartOfMlagPair() float32`

GetIsPartOfMlagPair returns the IsPartOfMlagPair field if non-nil, zero value otherwise.

### GetIsPartOfMlagPairOk

`func (o *CreateNetworkDeviceDefaults) GetIsPartOfMlagPairOk() (*float32, bool)`

GetIsPartOfMlagPairOk returns a tuple with the IsPartOfMlagPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartOfMlagPair

`func (o *CreateNetworkDeviceDefaults) SetIsPartOfMlagPair(v float32)`

SetIsPartOfMlagPair sets IsPartOfMlagPair field to given value.

### HasIsPartOfMlagPair

`func (o *CreateNetworkDeviceDefaults) HasIsPartOfMlagPair() bool`

HasIsPartOfMlagPair returns a boolean if a field has been set.

### GetMlagSystemMac

`func (o *CreateNetworkDeviceDefaults) GetMlagSystemMac() string`

GetMlagSystemMac returns the MlagSystemMac field if non-nil, zero value otherwise.

### GetMlagSystemMacOk

`func (o *CreateNetworkDeviceDefaults) GetMlagSystemMacOk() (*string, bool)`

GetMlagSystemMacOk returns a tuple with the MlagSystemMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagSystemMac

`func (o *CreateNetworkDeviceDefaults) SetMlagSystemMac(v string)`

SetMlagSystemMac sets MlagSystemMac field to given value.

### HasMlagSystemMac

`func (o *CreateNetworkDeviceDefaults) HasMlagSystemMac() bool`

HasMlagSystemMac returns a boolean if a field has been set.

### GetMlagPeerLinkPortChannelId

`func (o *CreateNetworkDeviceDefaults) GetMlagPeerLinkPortChannelId() float32`

GetMlagPeerLinkPortChannelId returns the MlagPeerLinkPortChannelId field if non-nil, zero value otherwise.

### GetMlagPeerLinkPortChannelIdOk

`func (o *CreateNetworkDeviceDefaults) GetMlagPeerLinkPortChannelIdOk() (*float32, bool)`

GetMlagPeerLinkPortChannelIdOk returns a tuple with the MlagPeerLinkPortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerLinkPortChannelId

`func (o *CreateNetworkDeviceDefaults) SetMlagPeerLinkPortChannelId(v float32)`

SetMlagPeerLinkPortChannelId sets MlagPeerLinkPortChannelId field to given value.

### HasMlagPeerLinkPortChannelId

`func (o *CreateNetworkDeviceDefaults) HasMlagPeerLinkPortChannelId() bool`

HasMlagPeerLinkPortChannelId returns a boolean if a field has been set.

### GetMlagPartnerVlanId

`func (o *CreateNetworkDeviceDefaults) GetMlagPartnerVlanId() float32`

GetMlagPartnerVlanId returns the MlagPartnerVlanId field if non-nil, zero value otherwise.

### GetMlagPartnerVlanIdOk

`func (o *CreateNetworkDeviceDefaults) GetMlagPartnerVlanIdOk() (*float32, bool)`

GetMlagPartnerVlanIdOk returns a tuple with the MlagPartnerVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPartnerVlanId

`func (o *CreateNetworkDeviceDefaults) SetMlagPartnerVlanId(v float32)`

SetMlagPartnerVlanId sets MlagPartnerVlanId field to given value.

### HasMlagPartnerVlanId

`func (o *CreateNetworkDeviceDefaults) HasMlagPartnerVlanId() bool`

HasMlagPartnerVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


