# CreateExternalSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The external system unique label | 
**Name** | **string** | The external system name | 
**Asn** | Pointer to **int64** | The BGP Autonomous System Number of the external system. Required to peer with it: a BGP session against this system renders its neighbor remote-as from this value. | [optional] 
**LoopbackAddress** | Pointer to **string** | The loopback IPv4 address of the external system. Used as the peer address of a multihop BGP session against it; a session running over the link itself takes its peer address from the link&#39;s point-to-point /31 instead. | [optional] 
**LoopbackAddressIpv6** | Pointer to **string** | The loopback IPv6 address of the external system. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables exposed to network device configuration templates as the remote side of a link or BGP session terminating on this external system. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Annotations - key/value pairs of additional metadata | [optional] 

## Methods

### NewCreateExternalSystem

`func NewCreateExternalSystem(label string, name string, ) *CreateExternalSystem`

NewCreateExternalSystem instantiates a new CreateExternalSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalSystemWithDefaults

`func NewCreateExternalSystemWithDefaults() *CreateExternalSystem`

NewCreateExternalSystemWithDefaults instantiates a new CreateExternalSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateExternalSystem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateExternalSystem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateExternalSystem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateExternalSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExternalSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExternalSystem) SetName(v string)`

SetName sets Name field to given value.


### GetAsn

`func (o *CreateExternalSystem) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *CreateExternalSystem) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *CreateExternalSystem) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *CreateExternalSystem) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetLoopbackAddress

`func (o *CreateExternalSystem) GetLoopbackAddress() string`

GetLoopbackAddress returns the LoopbackAddress field if non-nil, zero value otherwise.

### GetLoopbackAddressOk

`func (o *CreateExternalSystem) GetLoopbackAddressOk() (*string, bool)`

GetLoopbackAddressOk returns a tuple with the LoopbackAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddress

`func (o *CreateExternalSystem) SetLoopbackAddress(v string)`

SetLoopbackAddress sets LoopbackAddress field to given value.

### HasLoopbackAddress

`func (o *CreateExternalSystem) HasLoopbackAddress() bool`

HasLoopbackAddress returns a boolean if a field has been set.

### GetLoopbackAddressIpv6

`func (o *CreateExternalSystem) GetLoopbackAddressIpv6() string`

GetLoopbackAddressIpv6 returns the LoopbackAddressIpv6 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv6Ok

`func (o *CreateExternalSystem) GetLoopbackAddressIpv6Ok() (*string, bool)`

GetLoopbackAddressIpv6Ok returns a tuple with the LoopbackAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv6

`func (o *CreateExternalSystem) SetLoopbackAddressIpv6(v string)`

SetLoopbackAddressIpv6 sets LoopbackAddressIpv6 field to given value.

### HasLoopbackAddressIpv6

`func (o *CreateExternalSystem) HasLoopbackAddressIpv6() bool`

HasLoopbackAddressIpv6 returns a boolean if a field has been set.

### GetCustomVariables

`func (o *CreateExternalSystem) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *CreateExternalSystem) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *CreateExternalSystem) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *CreateExternalSystem) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateExternalSystem) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateExternalSystem) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateExternalSystem) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateExternalSystem) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


