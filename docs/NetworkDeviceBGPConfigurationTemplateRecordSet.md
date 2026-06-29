# NetworkDeviceBGPConfigurationTemplateRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MlagDomainIdentifier** | Pointer to **string** | Identifier for the MLAG domain | [optional] 
**LocalSwitchId** | **int64** | The ID of the local switch. | 
**LocalSwitchIdentifier** | **string** | The identifier of the local switch. | 
**LocalSwitchRole** | **string** | The role of the local switch. | 
**LocalSwitchInterfaceName** | **string** | The name of the local switch interface. | 
**LocalSwitchInterfaceLagId** | Pointer to **int64** | The LAG ID of the local switch interface. | [optional] 
**LocalSwitchAsn** | Pointer to **int64** | The Autonomous System Number of the local switch. | [optional] 
**LocalSwitchLoopbackAddressIpv4** | Pointer to **string** | The local switch loopback IPv4 address. | [optional] 
**RemoteSwitchId** | **int64** | The ID of the remote switch. | 
**RemoteSwitchIdentifier** | **string** | The identifier of the remote switch. | 
**RemoteSwitchRole** | **string** | The role of the remote switch. | 
**RemoteSwitchInterfaceName** | **string** | The name of the remote switch interface. | 
**RemoteSwitchInterfaceLagId** | Pointer to **int64** | The LAG ID of the remote switch interface. | [optional] 
**RemoteSwitchAsn** | Pointer to **int64** | The Autonomous System Number of the remote switch. | [optional] 
**RemoteSwitchLoopbackAddressIpv4** | Pointer to **string** | The remote switch loopback IPv4 address. | [optional] 
**BgpNumbering** | **string** | BGP numbering | 
**FabricCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the fabric. | [optional] 
**LinkCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the link. | [optional] 
**LocalSwitchCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the local switch. | [optional] 
**RemoteSwitchCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the remote switch. | [optional] 
**MlagPeerCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the MLAG pair of the local switch (if any). | [optional] 
**RemoteP2pLinkIp** | Pointer to **string** | The IP address of the remote interface of the P2P link. Present only when BGP numbering is NUMBERED. | [optional] 

## Methods

### NewNetworkDeviceBGPConfigurationTemplateRecordSet

`func NewNetworkDeviceBGPConfigurationTemplateRecordSet(localSwitchId int64, localSwitchIdentifier string, localSwitchRole string, localSwitchInterfaceName string, remoteSwitchId int64, remoteSwitchIdentifier string, remoteSwitchRole string, remoteSwitchInterfaceName string, bgpNumbering string, ) *NetworkDeviceBGPConfigurationTemplateRecordSet`

NewNetworkDeviceBGPConfigurationTemplateRecordSet instantiates a new NetworkDeviceBGPConfigurationTemplateRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceBGPConfigurationTemplateRecordSetWithDefaults

`func NewNetworkDeviceBGPConfigurationTemplateRecordSetWithDefaults() *NetworkDeviceBGPConfigurationTemplateRecordSet`

NewNetworkDeviceBGPConfigurationTemplateRecordSetWithDefaults instantiates a new NetworkDeviceBGPConfigurationTemplateRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMlagDomainIdentifier

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetMlagDomainIdentifier() string`

GetMlagDomainIdentifier returns the MlagDomainIdentifier field if non-nil, zero value otherwise.

### GetMlagDomainIdentifierOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetMlagDomainIdentifierOk() (*string, bool)`

GetMlagDomainIdentifierOk returns a tuple with the MlagDomainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagDomainIdentifier

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetMlagDomainIdentifier(v string)`

SetMlagDomainIdentifier sets MlagDomainIdentifier field to given value.

### HasMlagDomainIdentifier

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasMlagDomainIdentifier() bool`

HasMlagDomainIdentifier returns a boolean if a field has been set.

### GetLocalSwitchId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchId() int64`

GetLocalSwitchId returns the LocalSwitchId field if non-nil, zero value otherwise.

### GetLocalSwitchIdOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchIdOk() (*int64, bool)`

GetLocalSwitchIdOk returns a tuple with the LocalSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchId(v int64)`

SetLocalSwitchId sets LocalSwitchId field to given value.


### GetLocalSwitchIdentifier

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchIdentifier() string`

GetLocalSwitchIdentifier returns the LocalSwitchIdentifier field if non-nil, zero value otherwise.

### GetLocalSwitchIdentifierOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchIdentifierOk() (*string, bool)`

GetLocalSwitchIdentifierOk returns a tuple with the LocalSwitchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchIdentifier

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchIdentifier(v string)`

SetLocalSwitchIdentifier sets LocalSwitchIdentifier field to given value.


### GetLocalSwitchRole

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchRole() string`

GetLocalSwitchRole returns the LocalSwitchRole field if non-nil, zero value otherwise.

### GetLocalSwitchRoleOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchRoleOk() (*string, bool)`

GetLocalSwitchRoleOk returns a tuple with the LocalSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchRole

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchRole(v string)`

SetLocalSwitchRole sets LocalSwitchRole field to given value.


### GetLocalSwitchInterfaceName

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchInterfaceName() string`

GetLocalSwitchInterfaceName returns the LocalSwitchInterfaceName field if non-nil, zero value otherwise.

### GetLocalSwitchInterfaceNameOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchInterfaceNameOk() (*string, bool)`

GetLocalSwitchInterfaceNameOk returns a tuple with the LocalSwitchInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchInterfaceName

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchInterfaceName(v string)`

SetLocalSwitchInterfaceName sets LocalSwitchInterfaceName field to given value.


### GetLocalSwitchInterfaceLagId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchInterfaceLagId() int64`

GetLocalSwitchInterfaceLagId returns the LocalSwitchInterfaceLagId field if non-nil, zero value otherwise.

### GetLocalSwitchInterfaceLagIdOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchInterfaceLagIdOk() (*int64, bool)`

GetLocalSwitchInterfaceLagIdOk returns a tuple with the LocalSwitchInterfaceLagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchInterfaceLagId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchInterfaceLagId(v int64)`

SetLocalSwitchInterfaceLagId sets LocalSwitchInterfaceLagId field to given value.

### HasLocalSwitchInterfaceLagId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasLocalSwitchInterfaceLagId() bool`

HasLocalSwitchInterfaceLagId returns a boolean if a field has been set.

### GetLocalSwitchAsn

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchAsn() int64`

GetLocalSwitchAsn returns the LocalSwitchAsn field if non-nil, zero value otherwise.

### GetLocalSwitchAsnOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchAsnOk() (*int64, bool)`

GetLocalSwitchAsnOk returns a tuple with the LocalSwitchAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchAsn

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchAsn(v int64)`

SetLocalSwitchAsn sets LocalSwitchAsn field to given value.

### HasLocalSwitchAsn

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasLocalSwitchAsn() bool`

HasLocalSwitchAsn returns a boolean if a field has been set.

### GetLocalSwitchLoopbackAddressIpv4

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchLoopbackAddressIpv4() string`

GetLocalSwitchLoopbackAddressIpv4 returns the LocalSwitchLoopbackAddressIpv4 field if non-nil, zero value otherwise.

### GetLocalSwitchLoopbackAddressIpv4Ok

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchLoopbackAddressIpv4Ok() (*string, bool)`

GetLocalSwitchLoopbackAddressIpv4Ok returns a tuple with the LocalSwitchLoopbackAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchLoopbackAddressIpv4

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchLoopbackAddressIpv4(v string)`

SetLocalSwitchLoopbackAddressIpv4 sets LocalSwitchLoopbackAddressIpv4 field to given value.

### HasLocalSwitchLoopbackAddressIpv4

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasLocalSwitchLoopbackAddressIpv4() bool`

HasLocalSwitchLoopbackAddressIpv4 returns a boolean if a field has been set.

### GetRemoteSwitchId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchId() int64`

GetRemoteSwitchId returns the RemoteSwitchId field if non-nil, zero value otherwise.

### GetRemoteSwitchIdOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchIdOk() (*int64, bool)`

GetRemoteSwitchIdOk returns a tuple with the RemoteSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchId(v int64)`

SetRemoteSwitchId sets RemoteSwitchId field to given value.


### GetRemoteSwitchIdentifier

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchIdentifier() string`

GetRemoteSwitchIdentifier returns the RemoteSwitchIdentifier field if non-nil, zero value otherwise.

### GetRemoteSwitchIdentifierOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchIdentifierOk() (*string, bool)`

GetRemoteSwitchIdentifierOk returns a tuple with the RemoteSwitchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchIdentifier

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchIdentifier(v string)`

SetRemoteSwitchIdentifier sets RemoteSwitchIdentifier field to given value.


### GetRemoteSwitchRole

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchRole() string`

GetRemoteSwitchRole returns the RemoteSwitchRole field if non-nil, zero value otherwise.

### GetRemoteSwitchRoleOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchRoleOk() (*string, bool)`

GetRemoteSwitchRoleOk returns a tuple with the RemoteSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchRole

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchRole(v string)`

SetRemoteSwitchRole sets RemoteSwitchRole field to given value.


### GetRemoteSwitchInterfaceName

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceName() string`

GetRemoteSwitchInterfaceName returns the RemoteSwitchInterfaceName field if non-nil, zero value otherwise.

### GetRemoteSwitchInterfaceNameOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceNameOk() (*string, bool)`

GetRemoteSwitchInterfaceNameOk returns a tuple with the RemoteSwitchInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchInterfaceName

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchInterfaceName(v string)`

SetRemoteSwitchInterfaceName sets RemoteSwitchInterfaceName field to given value.


### GetRemoteSwitchInterfaceLagId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceLagId() int64`

GetRemoteSwitchInterfaceLagId returns the RemoteSwitchInterfaceLagId field if non-nil, zero value otherwise.

### GetRemoteSwitchInterfaceLagIdOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceLagIdOk() (*int64, bool)`

GetRemoteSwitchInterfaceLagIdOk returns a tuple with the RemoteSwitchInterfaceLagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchInterfaceLagId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchInterfaceLagId(v int64)`

SetRemoteSwitchInterfaceLagId sets RemoteSwitchInterfaceLagId field to given value.

### HasRemoteSwitchInterfaceLagId

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasRemoteSwitchInterfaceLagId() bool`

HasRemoteSwitchInterfaceLagId returns a boolean if a field has been set.

### GetRemoteSwitchAsn

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchAsn() int64`

GetRemoteSwitchAsn returns the RemoteSwitchAsn field if non-nil, zero value otherwise.

### GetRemoteSwitchAsnOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchAsnOk() (*int64, bool)`

GetRemoteSwitchAsnOk returns a tuple with the RemoteSwitchAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchAsn

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchAsn(v int64)`

SetRemoteSwitchAsn sets RemoteSwitchAsn field to given value.

### HasRemoteSwitchAsn

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasRemoteSwitchAsn() bool`

HasRemoteSwitchAsn returns a boolean if a field has been set.

### GetRemoteSwitchLoopbackAddressIpv4

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchLoopbackAddressIpv4() string`

GetRemoteSwitchLoopbackAddressIpv4 returns the RemoteSwitchLoopbackAddressIpv4 field if non-nil, zero value otherwise.

### GetRemoteSwitchLoopbackAddressIpv4Ok

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchLoopbackAddressIpv4Ok() (*string, bool)`

GetRemoteSwitchLoopbackAddressIpv4Ok returns a tuple with the RemoteSwitchLoopbackAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchLoopbackAddressIpv4

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchLoopbackAddressIpv4(v string)`

SetRemoteSwitchLoopbackAddressIpv4 sets RemoteSwitchLoopbackAddressIpv4 field to given value.

### HasRemoteSwitchLoopbackAddressIpv4

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasRemoteSwitchLoopbackAddressIpv4() bool`

HasRemoteSwitchLoopbackAddressIpv4 returns a boolean if a field has been set.

### GetBgpNumbering

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetFabricCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetFabricCustomVariables() map[string]interface{}`

GetFabricCustomVariables returns the FabricCustomVariables field if non-nil, zero value otherwise.

### GetFabricCustomVariablesOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetFabricCustomVariablesOk() (*map[string]interface{}, bool)`

GetFabricCustomVariablesOk returns a tuple with the FabricCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetFabricCustomVariables(v map[string]interface{})`

SetFabricCustomVariables sets FabricCustomVariables field to given value.

### HasFabricCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasFabricCustomVariables() bool`

HasFabricCustomVariables returns a boolean if a field has been set.

### GetLinkCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLinkCustomVariables() map[string]interface{}`

GetLinkCustomVariables returns the LinkCustomVariables field if non-nil, zero value otherwise.

### GetLinkCustomVariablesOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLinkCustomVariablesOk() (*map[string]interface{}, bool)`

GetLinkCustomVariablesOk returns a tuple with the LinkCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLinkCustomVariables(v map[string]interface{})`

SetLinkCustomVariables sets LinkCustomVariables field to given value.

### HasLinkCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasLinkCustomVariables() bool`

HasLinkCustomVariables returns a boolean if a field has been set.

### GetLocalSwitchCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchCustomVariables() map[string]interface{}`

GetLocalSwitchCustomVariables returns the LocalSwitchCustomVariables field if non-nil, zero value otherwise.

### GetLocalSwitchCustomVariablesOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetLocalSwitchCustomVariablesOk() (*map[string]interface{}, bool)`

GetLocalSwitchCustomVariablesOk returns a tuple with the LocalSwitchCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetLocalSwitchCustomVariables(v map[string]interface{})`

SetLocalSwitchCustomVariables sets LocalSwitchCustomVariables field to given value.

### HasLocalSwitchCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasLocalSwitchCustomVariables() bool`

HasLocalSwitchCustomVariables returns a boolean if a field has been set.

### GetRemoteSwitchCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchCustomVariables() map[string]interface{}`

GetRemoteSwitchCustomVariables returns the RemoteSwitchCustomVariables field if non-nil, zero value otherwise.

### GetRemoteSwitchCustomVariablesOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteSwitchCustomVariablesOk() (*map[string]interface{}, bool)`

GetRemoteSwitchCustomVariablesOk returns a tuple with the RemoteSwitchCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteSwitchCustomVariables(v map[string]interface{})`

SetRemoteSwitchCustomVariables sets RemoteSwitchCustomVariables field to given value.

### HasRemoteSwitchCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasRemoteSwitchCustomVariables() bool`

HasRemoteSwitchCustomVariables returns a boolean if a field has been set.

### GetMlagPeerCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetMlagPeerCustomVariables() map[string]interface{}`

GetMlagPeerCustomVariables returns the MlagPeerCustomVariables field if non-nil, zero value otherwise.

### GetMlagPeerCustomVariablesOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetMlagPeerCustomVariablesOk() (*map[string]interface{}, bool)`

GetMlagPeerCustomVariablesOk returns a tuple with the MlagPeerCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetMlagPeerCustomVariables(v map[string]interface{})`

SetMlagPeerCustomVariables sets MlagPeerCustomVariables field to given value.

### HasMlagPeerCustomVariables

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasMlagPeerCustomVariables() bool`

HasMlagPeerCustomVariables returns a boolean if a field has been set.

### GetRemoteP2pLinkIp

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteP2pLinkIp() string`

GetRemoteP2pLinkIp returns the RemoteP2pLinkIp field if non-nil, zero value otherwise.

### GetRemoteP2pLinkIpOk

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) GetRemoteP2pLinkIpOk() (*string, bool)`

GetRemoteP2pLinkIpOk returns a tuple with the RemoteP2pLinkIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteP2pLinkIp

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) SetRemoteP2pLinkIp(v string)`

SetRemoteP2pLinkIp sets RemoteP2pLinkIp field to given value.

### HasRemoteP2pLinkIp

`func (o *NetworkDeviceBGPConfigurationTemplateRecordSet) HasRemoteP2pLinkIp() bool`

HasRemoteP2pLinkIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


