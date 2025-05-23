# CreateSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**ParentSubnetId** | Pointer to **int32** |  | [optional] 
**NetworkAddress** | **string** |  | 
**PrefixLength** | **int32** |  | 
**DefaultGatewayAddress** | Pointer to **string** |  | [optional] 
**IsPool** | **bool** |  | 
**AllocationDenylist** | Pointer to [**[]AddressRange**](AddressRange.md) |  | [optional] 
**ChildOverlapAllowRules** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSubnet

`func NewCreateSubnet(networkAddress string, prefixLength int32, isPool bool, ) *CreateSubnet`

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

### HasName

`func (o *CreateSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateSubnet) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateSubnet) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateSubnet) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateSubnet) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetTags

`func (o *CreateSubnet) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSubnet) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSubnet) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSubnet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParentSubnetId

`func (o *CreateSubnet) GetParentSubnetId() int32`

GetParentSubnetId returns the ParentSubnetId field if non-nil, zero value otherwise.

### GetParentSubnetIdOk

`func (o *CreateSubnet) GetParentSubnetIdOk() (*int32, bool)`

GetParentSubnetIdOk returns a tuple with the ParentSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSubnetId

`func (o *CreateSubnet) SetParentSubnetId(v int32)`

SetParentSubnetId sets ParentSubnetId field to given value.

### HasParentSubnetId

`func (o *CreateSubnet) HasParentSubnetId() bool`

HasParentSubnetId returns a boolean if a field has been set.

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

`func (o *CreateSubnet) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateSubnet) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateSubnet) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetDefaultGatewayAddress

`func (o *CreateSubnet) GetDefaultGatewayAddress() string`

GetDefaultGatewayAddress returns the DefaultGatewayAddress field if non-nil, zero value otherwise.

### GetDefaultGatewayAddressOk

`func (o *CreateSubnet) GetDefaultGatewayAddressOk() (*string, bool)`

GetDefaultGatewayAddressOk returns a tuple with the DefaultGatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAddress

`func (o *CreateSubnet) SetDefaultGatewayAddress(v string)`

SetDefaultGatewayAddress sets DefaultGatewayAddress field to given value.

### HasDefaultGatewayAddress

`func (o *CreateSubnet) HasDefaultGatewayAddress() bool`

HasDefaultGatewayAddress returns a boolean if a field has been set.

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


### GetAllocationDenylist

`func (o *CreateSubnet) GetAllocationDenylist() []AddressRange`

GetAllocationDenylist returns the AllocationDenylist field if non-nil, zero value otherwise.

### GetAllocationDenylistOk

`func (o *CreateSubnet) GetAllocationDenylistOk() (*[]AddressRange, bool)`

GetAllocationDenylistOk returns a tuple with the AllocationDenylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDenylist

`func (o *CreateSubnet) SetAllocationDenylist(v []AddressRange)`

SetAllocationDenylist sets AllocationDenylist field to given value.

### HasAllocationDenylist

`func (o *CreateSubnet) HasAllocationDenylist() bool`

HasAllocationDenylist returns a boolean if a field has been set.

### GetChildOverlapAllowRules

`func (o *CreateSubnet) GetChildOverlapAllowRules() []string`

GetChildOverlapAllowRules returns the ChildOverlapAllowRules field if non-nil, zero value otherwise.

### GetChildOverlapAllowRulesOk

`func (o *CreateSubnet) GetChildOverlapAllowRulesOk() (*[]string, bool)`

GetChildOverlapAllowRulesOk returns a tuple with the ChildOverlapAllowRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOverlapAllowRules

`func (o *CreateSubnet) SetChildOverlapAllowRules(v []string)`

SetChildOverlapAllowRules sets ChildOverlapAllowRules field to given value.

### HasChildOverlapAllowRules

`func (o *CreateSubnet) HasChildOverlapAllowRules() bool`

HasChildOverlapAllowRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


