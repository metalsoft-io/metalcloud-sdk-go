# ExternalSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The external system id. | 
**Label** | **string** | The external system unique label | 
**Name** | **string** | The external system name | 
**Annotations** | **map[string]interface{}** | Annotations - key/value pairs of additional metadata | 
**Asn** | Pointer to **int64** | The BGP Autonomous System Number of the external system. Required to peer with it: a BGP session against this system renders its neighbor remote-as from this value. | [optional] 
**LoopbackAddress** | Pointer to **string** | The loopback IPv4 address of the external system. Used as the peer address of a multihop BGP session against it; a session running over the link itself takes its peer address from the link&#39;s point-to-point /31 instead. | [optional] 
**LoopbackAddressIpv6** | Pointer to **string** | The loopback IPv6 address of the external system. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables exposed to network device configuration templates as the remote side of a link or BGP session terminating on this external system. | [optional] 
**Revision** | **string** | Revision number of the entity | 
**CreatedAt** | **time.Time** | The date and time the entity was created | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the entity was last updated | [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewExternalSystem

`func NewExternalSystem(id string, label string, name string, annotations map[string]interface{}, revision string, createdAt time.Time, updatedAt time.Time, ) *ExternalSystem`

NewExternalSystem instantiates a new ExternalSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSystemWithDefaults

`func NewExternalSystemWithDefaults() *ExternalSystem`

NewExternalSystemWithDefaults instantiates a new ExternalSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalSystem) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ExternalSystem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalSystem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalSystem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ExternalSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalSystem) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *ExternalSystem) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ExternalSystem) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ExternalSystem) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.


### GetAsn

`func (o *ExternalSystem) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ExternalSystem) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ExternalSystem) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ExternalSystem) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetLoopbackAddress

`func (o *ExternalSystem) GetLoopbackAddress() string`

GetLoopbackAddress returns the LoopbackAddress field if non-nil, zero value otherwise.

### GetLoopbackAddressOk

`func (o *ExternalSystem) GetLoopbackAddressOk() (*string, bool)`

GetLoopbackAddressOk returns a tuple with the LoopbackAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddress

`func (o *ExternalSystem) SetLoopbackAddress(v string)`

SetLoopbackAddress sets LoopbackAddress field to given value.

### HasLoopbackAddress

`func (o *ExternalSystem) HasLoopbackAddress() bool`

HasLoopbackAddress returns a boolean if a field has been set.

### GetLoopbackAddressIpv6

`func (o *ExternalSystem) GetLoopbackAddressIpv6() string`

GetLoopbackAddressIpv6 returns the LoopbackAddressIpv6 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv6Ok

`func (o *ExternalSystem) GetLoopbackAddressIpv6Ok() (*string, bool)`

GetLoopbackAddressIpv6Ok returns a tuple with the LoopbackAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv6

`func (o *ExternalSystem) SetLoopbackAddressIpv6(v string)`

SetLoopbackAddressIpv6 sets LoopbackAddressIpv6 field to given value.

### HasLoopbackAddressIpv6

`func (o *ExternalSystem) HasLoopbackAddressIpv6() bool`

HasLoopbackAddressIpv6 returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ExternalSystem) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ExternalSystem) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ExternalSystem) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ExternalSystem) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetRevision

`func (o *ExternalSystem) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ExternalSystem) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ExternalSystem) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetCreatedAt

`func (o *ExternalSystem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalSystem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalSystem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExternalSystem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalSystem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalSystem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *ExternalSystem) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalSystem) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalSystem) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalSystem) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


