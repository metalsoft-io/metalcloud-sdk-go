# CreateSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | **string** | Name of the Subnet | 
**ParentSubnetId** | **float32** | ID of the parent subnet | 
**NetworkAddress** | **string** |  | 
**PrefixLength** | **float32** |  | 
**IsPool** | **bool** |  | 
**VrfId** | **float32** |  | 
**AllocationDenylist** | **[]string** |  | 
**AllowedChildOverlapConditions** | **[]string** |  | 
**Tags** | **map[string]interface{}** |  | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewCreateSubnet

`func NewCreateSubnet(name string, parentSubnetId float32, networkAddress string, prefixLength float32, isPool bool, vrfId float32, allocationDenylist []string, allowedChildOverlapConditions []string, tags map[string]interface{}, metadata map[string]interface{}, ) *CreateSubnet`

NewCreateSubnet instantiates a new CreateSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetWithDefaults

`func NewCreateSubnetWithDefaults() *CreateSubnet`

NewCreateSubnetWithDefaults instantiates a new CreateSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateSubnet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateSubnet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateSubnet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateSubnet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubnet) SetName(v string)`

SetName sets Name field to given value.


### GetParentSubnetId

`func (o *CreateSubnet) GetParentSubnetId() float32`

GetParentSubnetId returns the ParentSubnetId field if non-nil, zero value otherwise.

### GetParentSubnetIdOk

`func (o *CreateSubnet) GetParentSubnetIdOk() (*float32, bool)`

GetParentSubnetIdOk returns a tuple with the ParentSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSubnetId

`func (o *CreateSubnet) SetParentSubnetId(v float32)`

SetParentSubnetId sets ParentSubnetId field to given value.


### GetNetworkAddress

`func (o *CreateSubnet) GetNetworkAddress() string`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *CreateSubnet) GetNetworkAddressOk() (*string, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *CreateSubnet) SetNetworkAddress(v string)`

SetNetworkAddress sets NetworkAddress field to given value.


### GetPrefixLength

`func (o *CreateSubnet) GetPrefixLength() float32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateSubnet) GetPrefixLengthOk() (*float32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateSubnet) SetPrefixLength(v float32)`

SetPrefixLength sets PrefixLength field to given value.


### GetIsPool

`func (o *CreateSubnet) GetIsPool() bool`

GetIsPool returns the IsPool field if non-nil, zero value otherwise.

### GetIsPoolOk

`func (o *CreateSubnet) GetIsPoolOk() (*bool, bool)`

GetIsPoolOk returns a tuple with the IsPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPool

`func (o *CreateSubnet) SetIsPool(v bool)`

SetIsPool sets IsPool field to given value.


### GetVrfId

`func (o *CreateSubnet) GetVrfId() float32`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *CreateSubnet) GetVrfIdOk() (*float32, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *CreateSubnet) SetVrfId(v float32)`

SetVrfId sets VrfId field to given value.


### GetAllocationDenylist

`func (o *CreateSubnet) GetAllocationDenylist() []string`

GetAllocationDenylist returns the AllocationDenylist field if non-nil, zero value otherwise.

### GetAllocationDenylistOk

`func (o *CreateSubnet) GetAllocationDenylistOk() (*[]string, bool)`

GetAllocationDenylistOk returns a tuple with the AllocationDenylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDenylist

`func (o *CreateSubnet) SetAllocationDenylist(v []string)`

SetAllocationDenylist sets AllocationDenylist field to given value.


### GetAllowedChildOverlapConditions

`func (o *CreateSubnet) GetAllowedChildOverlapConditions() []string`

GetAllowedChildOverlapConditions returns the AllowedChildOverlapConditions field if non-nil, zero value otherwise.

### GetAllowedChildOverlapConditionsOk

`func (o *CreateSubnet) GetAllowedChildOverlapConditionsOk() (*[]string, bool)`

GetAllowedChildOverlapConditionsOk returns a tuple with the AllowedChildOverlapConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChildOverlapConditions

`func (o *CreateSubnet) SetAllowedChildOverlapConditions(v []string)`

SetAllowedChildOverlapConditions sets AllowedChildOverlapConditions field to given value.


### GetTags

`func (o *CreateSubnet) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSubnet) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSubnet) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetMetadata

`func (o *CreateSubnet) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSubnet) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSubnet) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


