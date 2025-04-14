# UpdateSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **float32** | Revision of the Subnet | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the Subnet | [optional] 
**AllocationDenylist** | Pointer to **[]string** |  | [optional] 
**AllowedChildOverlapConditions** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateSubnet

`func NewUpdateSubnet() *UpdateSubnet`

NewUpdateSubnet instantiates a new UpdateSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubnetWithDefaults

`func NewUpdateSubnetWithDefaults() *UpdateSubnet`

NewUpdateSubnetWithDefaults instantiates a new UpdateSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *UpdateSubnet) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UpdateSubnet) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UpdateSubnet) SetRevision(v float32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *UpdateSubnet) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateSubnet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateSubnet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateSubnet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateSubnet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllocationDenylist

`func (o *UpdateSubnet) GetAllocationDenylist() []string`

GetAllocationDenylist returns the AllocationDenylist field if non-nil, zero value otherwise.

### GetAllocationDenylistOk

`func (o *UpdateSubnet) GetAllocationDenylistOk() (*[]string, bool)`

GetAllocationDenylistOk returns a tuple with the AllocationDenylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDenylist

`func (o *UpdateSubnet) SetAllocationDenylist(v []string)`

SetAllocationDenylist sets AllocationDenylist field to given value.

### HasAllocationDenylist

`func (o *UpdateSubnet) HasAllocationDenylist() bool`

HasAllocationDenylist returns a boolean if a field has been set.

### GetAllowedChildOverlapConditions

`func (o *UpdateSubnet) GetAllowedChildOverlapConditions() []string`

GetAllowedChildOverlapConditions returns the AllowedChildOverlapConditions field if non-nil, zero value otherwise.

### GetAllowedChildOverlapConditionsOk

`func (o *UpdateSubnet) GetAllowedChildOverlapConditionsOk() (*[]string, bool)`

GetAllowedChildOverlapConditionsOk returns a tuple with the AllowedChildOverlapConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChildOverlapConditions

`func (o *UpdateSubnet) SetAllowedChildOverlapConditions(v []string)`

SetAllowedChildOverlapConditions sets AllowedChildOverlapConditions field to given value.

### HasAllowedChildOverlapConditions

`func (o *UpdateSubnet) HasAllowedChildOverlapConditions() bool`

HasAllowedChildOverlapConditions returns a boolean if a field has been set.

### GetTags

`func (o *UpdateSubnet) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateSubnet) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateSubnet) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateSubnet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSubnet) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSubnet) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSubnet) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSubnet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


