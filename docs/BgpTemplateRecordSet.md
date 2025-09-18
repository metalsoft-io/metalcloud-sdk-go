# BgpTemplateRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSwitchId** | **float32** | The ID of the local switch. | 
**LocalSwitchIdentifier** | **string** | The identifier of the local switch. | 
**LocalSwitchRole** | **string** | The role of the local switch. | 
**LocalSwitchInterfaceName** | **string** | The name of the local switch interface. | 
**LocalSwitchAsn** | Pointer to **float32** | The ASN of the local switch. | [optional] 
**RemoteSwitchId** | **float32** | The ID of the remote switch. | 
**RemoteSwitchIdentifier** | **string** | The identifier of the remote switch. | 
**RemoteSwitchRole** | **string** | The role of the remote switch. | 
**RemoteSwitchInterfaceName** | **string** | The name of the remote switch interface. | 
**RemoteSwitchAsn** | Pointer to **float32** | The ASN of the remote switch. | [optional] 
**BgpNumbering** | **string** | BGP numbering | 
**FabricCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the fabric. | [optional] 
**LinkCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the link. | [optional] 

## Methods

### NewBgpTemplateRecordSet

`func NewBgpTemplateRecordSet(localSwitchId float32, localSwitchIdentifier string, localSwitchRole string, localSwitchInterfaceName string, remoteSwitchId float32, remoteSwitchIdentifier string, remoteSwitchRole string, remoteSwitchInterfaceName string, bgpNumbering string, ) *BgpTemplateRecordSet`

NewBgpTemplateRecordSet instantiates a new BgpTemplateRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpTemplateRecordSetWithDefaults

`func NewBgpTemplateRecordSetWithDefaults() *BgpTemplateRecordSet`

NewBgpTemplateRecordSetWithDefaults instantiates a new BgpTemplateRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSwitchId

`func (o *BgpTemplateRecordSet) GetLocalSwitchId() float32`

GetLocalSwitchId returns the LocalSwitchId field if non-nil, zero value otherwise.

### GetLocalSwitchIdOk

`func (o *BgpTemplateRecordSet) GetLocalSwitchIdOk() (*float32, bool)`

GetLocalSwitchIdOk returns a tuple with the LocalSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchId

`func (o *BgpTemplateRecordSet) SetLocalSwitchId(v float32)`

SetLocalSwitchId sets LocalSwitchId field to given value.


### GetLocalSwitchIdentifier

`func (o *BgpTemplateRecordSet) GetLocalSwitchIdentifier() string`

GetLocalSwitchIdentifier returns the LocalSwitchIdentifier field if non-nil, zero value otherwise.

### GetLocalSwitchIdentifierOk

`func (o *BgpTemplateRecordSet) GetLocalSwitchIdentifierOk() (*string, bool)`

GetLocalSwitchIdentifierOk returns a tuple with the LocalSwitchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchIdentifier

`func (o *BgpTemplateRecordSet) SetLocalSwitchIdentifier(v string)`

SetLocalSwitchIdentifier sets LocalSwitchIdentifier field to given value.


### GetLocalSwitchRole

`func (o *BgpTemplateRecordSet) GetLocalSwitchRole() string`

GetLocalSwitchRole returns the LocalSwitchRole field if non-nil, zero value otherwise.

### GetLocalSwitchRoleOk

`func (o *BgpTemplateRecordSet) GetLocalSwitchRoleOk() (*string, bool)`

GetLocalSwitchRoleOk returns a tuple with the LocalSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchRole

`func (o *BgpTemplateRecordSet) SetLocalSwitchRole(v string)`

SetLocalSwitchRole sets LocalSwitchRole field to given value.


### GetLocalSwitchInterfaceName

`func (o *BgpTemplateRecordSet) GetLocalSwitchInterfaceName() string`

GetLocalSwitchInterfaceName returns the LocalSwitchInterfaceName field if non-nil, zero value otherwise.

### GetLocalSwitchInterfaceNameOk

`func (o *BgpTemplateRecordSet) GetLocalSwitchInterfaceNameOk() (*string, bool)`

GetLocalSwitchInterfaceNameOk returns a tuple with the LocalSwitchInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchInterfaceName

`func (o *BgpTemplateRecordSet) SetLocalSwitchInterfaceName(v string)`

SetLocalSwitchInterfaceName sets LocalSwitchInterfaceName field to given value.


### GetLocalSwitchAsn

`func (o *BgpTemplateRecordSet) GetLocalSwitchAsn() float32`

GetLocalSwitchAsn returns the LocalSwitchAsn field if non-nil, zero value otherwise.

### GetLocalSwitchAsnOk

`func (o *BgpTemplateRecordSet) GetLocalSwitchAsnOk() (*float32, bool)`

GetLocalSwitchAsnOk returns a tuple with the LocalSwitchAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchAsn

`func (o *BgpTemplateRecordSet) SetLocalSwitchAsn(v float32)`

SetLocalSwitchAsn sets LocalSwitchAsn field to given value.

### HasLocalSwitchAsn

`func (o *BgpTemplateRecordSet) HasLocalSwitchAsn() bool`

HasLocalSwitchAsn returns a boolean if a field has been set.

### GetRemoteSwitchId

`func (o *BgpTemplateRecordSet) GetRemoteSwitchId() float32`

GetRemoteSwitchId returns the RemoteSwitchId field if non-nil, zero value otherwise.

### GetRemoteSwitchIdOk

`func (o *BgpTemplateRecordSet) GetRemoteSwitchIdOk() (*float32, bool)`

GetRemoteSwitchIdOk returns a tuple with the RemoteSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchId

`func (o *BgpTemplateRecordSet) SetRemoteSwitchId(v float32)`

SetRemoteSwitchId sets RemoteSwitchId field to given value.


### GetRemoteSwitchIdentifier

`func (o *BgpTemplateRecordSet) GetRemoteSwitchIdentifier() string`

GetRemoteSwitchIdentifier returns the RemoteSwitchIdentifier field if non-nil, zero value otherwise.

### GetRemoteSwitchIdentifierOk

`func (o *BgpTemplateRecordSet) GetRemoteSwitchIdentifierOk() (*string, bool)`

GetRemoteSwitchIdentifierOk returns a tuple with the RemoteSwitchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchIdentifier

`func (o *BgpTemplateRecordSet) SetRemoteSwitchIdentifier(v string)`

SetRemoteSwitchIdentifier sets RemoteSwitchIdentifier field to given value.


### GetRemoteSwitchRole

`func (o *BgpTemplateRecordSet) GetRemoteSwitchRole() string`

GetRemoteSwitchRole returns the RemoteSwitchRole field if non-nil, zero value otherwise.

### GetRemoteSwitchRoleOk

`func (o *BgpTemplateRecordSet) GetRemoteSwitchRoleOk() (*string, bool)`

GetRemoteSwitchRoleOk returns a tuple with the RemoteSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchRole

`func (o *BgpTemplateRecordSet) SetRemoteSwitchRole(v string)`

SetRemoteSwitchRole sets RemoteSwitchRole field to given value.


### GetRemoteSwitchInterfaceName

`func (o *BgpTemplateRecordSet) GetRemoteSwitchInterfaceName() string`

GetRemoteSwitchInterfaceName returns the RemoteSwitchInterfaceName field if non-nil, zero value otherwise.

### GetRemoteSwitchInterfaceNameOk

`func (o *BgpTemplateRecordSet) GetRemoteSwitchInterfaceNameOk() (*string, bool)`

GetRemoteSwitchInterfaceNameOk returns a tuple with the RemoteSwitchInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchInterfaceName

`func (o *BgpTemplateRecordSet) SetRemoteSwitchInterfaceName(v string)`

SetRemoteSwitchInterfaceName sets RemoteSwitchInterfaceName field to given value.


### GetRemoteSwitchAsn

`func (o *BgpTemplateRecordSet) GetRemoteSwitchAsn() float32`

GetRemoteSwitchAsn returns the RemoteSwitchAsn field if non-nil, zero value otherwise.

### GetRemoteSwitchAsnOk

`func (o *BgpTemplateRecordSet) GetRemoteSwitchAsnOk() (*float32, bool)`

GetRemoteSwitchAsnOk returns a tuple with the RemoteSwitchAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchAsn

`func (o *BgpTemplateRecordSet) SetRemoteSwitchAsn(v float32)`

SetRemoteSwitchAsn sets RemoteSwitchAsn field to given value.

### HasRemoteSwitchAsn

`func (o *BgpTemplateRecordSet) HasRemoteSwitchAsn() bool`

HasRemoteSwitchAsn returns a boolean if a field has been set.

### GetBgpNumbering

`func (o *BgpTemplateRecordSet) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *BgpTemplateRecordSet) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *BgpTemplateRecordSet) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetFabricCustomVariables

`func (o *BgpTemplateRecordSet) GetFabricCustomVariables() map[string]interface{}`

GetFabricCustomVariables returns the FabricCustomVariables field if non-nil, zero value otherwise.

### GetFabricCustomVariablesOk

`func (o *BgpTemplateRecordSet) GetFabricCustomVariablesOk() (*map[string]interface{}, bool)`

GetFabricCustomVariablesOk returns a tuple with the FabricCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricCustomVariables

`func (o *BgpTemplateRecordSet) SetFabricCustomVariables(v map[string]interface{})`

SetFabricCustomVariables sets FabricCustomVariables field to given value.

### HasFabricCustomVariables

`func (o *BgpTemplateRecordSet) HasFabricCustomVariables() bool`

HasFabricCustomVariables returns a boolean if a field has been set.

### GetLinkCustomVariables

`func (o *BgpTemplateRecordSet) GetLinkCustomVariables() map[string]interface{}`

GetLinkCustomVariables returns the LinkCustomVariables field if non-nil, zero value otherwise.

### GetLinkCustomVariablesOk

`func (o *BgpTemplateRecordSet) GetLinkCustomVariablesOk() (*map[string]interface{}, bool)`

GetLinkCustomVariablesOk returns a tuple with the LinkCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCustomVariables

`func (o *BgpTemplateRecordSet) SetLinkCustomVariables(v map[string]interface{})`

SetLinkCustomVariables sets LinkCustomVariables field to given value.

### HasLinkCustomVariables

`func (o *BgpTemplateRecordSet) HasLinkCustomVariables() bool`

HasLinkCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


