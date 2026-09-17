# UpdateExternalSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The external system unique label | [optional] 
**Name** | Pointer to **string** | The external system name | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Annotations - key/value pairs of additional metadata | [optional] 
**Asn** | Pointer to **int64** | The BGP Autonomous System Number of the external system. Required to peer with it: a BGP session against this system renders its neighbor remote-as from this value. | [optional] 
**LoopbackAddress** | Pointer to **string** | The loopback IPv4 address of the external system. Used as the peer address of a multihop BGP session against it. | [optional] 
**LoopbackAddressIpv6** | Pointer to **string** | The loopback IPv6 address of the external system. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables exposed to network device configuration templates as the remote side of a link or BGP session terminating on this external system. | [optional] 

## Methods

### NewUpdateExternalSystem

`func NewUpdateExternalSystem() *UpdateExternalSystem`

NewUpdateExternalSystem instantiates a new UpdateExternalSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExternalSystemWithDefaults

`func NewUpdateExternalSystemWithDefaults() *UpdateExternalSystem`

NewUpdateExternalSystemWithDefaults instantiates a new UpdateExternalSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateExternalSystem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateExternalSystem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateExternalSystem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateExternalSystem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateExternalSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExternalSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExternalSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateExternalSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *UpdateExternalSystem) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateExternalSystem) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateExternalSystem) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateExternalSystem) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAsn

`func (o *UpdateExternalSystem) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *UpdateExternalSystem) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *UpdateExternalSystem) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *UpdateExternalSystem) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetLoopbackAddress

`func (o *UpdateExternalSystem) GetLoopbackAddress() string`

GetLoopbackAddress returns the LoopbackAddress field if non-nil, zero value otherwise.

### GetLoopbackAddressOk

`func (o *UpdateExternalSystem) GetLoopbackAddressOk() (*string, bool)`

GetLoopbackAddressOk returns a tuple with the LoopbackAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddress

`func (o *UpdateExternalSystem) SetLoopbackAddress(v string)`

SetLoopbackAddress sets LoopbackAddress field to given value.

### HasLoopbackAddress

`func (o *UpdateExternalSystem) HasLoopbackAddress() bool`

HasLoopbackAddress returns a boolean if a field has been set.

### GetLoopbackAddressIpv6

`func (o *UpdateExternalSystem) GetLoopbackAddressIpv6() string`

GetLoopbackAddressIpv6 returns the LoopbackAddressIpv6 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv6Ok

`func (o *UpdateExternalSystem) GetLoopbackAddressIpv6Ok() (*string, bool)`

GetLoopbackAddressIpv6Ok returns a tuple with the LoopbackAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv6

`func (o *UpdateExternalSystem) SetLoopbackAddressIpv6(v string)`

SetLoopbackAddressIpv6 sets LoopbackAddressIpv6 field to given value.

### HasLoopbackAddressIpv6

`func (o *UpdateExternalSystem) HasLoopbackAddressIpv6() bool`

HasLoopbackAddressIpv6 returns a boolean if a field has been set.

### GetCustomVariables

`func (o *UpdateExternalSystem) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateExternalSystem) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateExternalSystem) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateExternalSystem) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


