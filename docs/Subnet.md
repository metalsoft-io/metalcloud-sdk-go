# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Label** | **string** |  | 
**Name** | **string** |  | 
**Annotations** | **map[string]string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Revision** | **int32** |  | 
**Tags** | **map[string]string** |  | 
**ParentSubnetId** | **int32** |  | 
**IpVersion** | [**IpVersion**](IpVersion.md) |  | 
**NetworkAddress** | **string** |  | 
**PrefixLength** | **int32** |  | 
**Netmask** | **string** |  | 
**DefaultGatewayAddress** | **string** |  | 
**IsPool** | **bool** |  | 
**AllocationDenylist** | [**[]AddressRange**](AddressRange.md) |  | 
**ChildOverlapAllowRules** | **[]string** |  | 

## Methods

### NewSubnet

`func NewSubnet(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, tags map[string]string, parentSubnetId int32, ipVersion IpVersion, networkAddress string, prefixLength int32, netmask string, defaultGatewayAddress string, isPool bool, allocationDenylist []AddressRange, childOverlapAllowRules []string, ) *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subnet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *Subnet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Subnet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Subnet) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *Subnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subnet) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *Subnet) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Subnet) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Subnet) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *Subnet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subnet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subnet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Subnet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subnet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subnet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *Subnet) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Subnet) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Subnet) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetTags

`func (o *Subnet) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Subnet) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Subnet) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetParentSubnetId

`func (o *Subnet) GetParentSubnetId() int32`

GetParentSubnetId returns the ParentSubnetId field if non-nil, zero value otherwise.

### GetParentSubnetIdOk

`func (o *Subnet) GetParentSubnetIdOk() (*int32, bool)`

GetParentSubnetIdOk returns a tuple with the ParentSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSubnetId

`func (o *Subnet) SetParentSubnetId(v int32)`

SetParentSubnetId sets ParentSubnetId field to given value.


### GetIpVersion

`func (o *Subnet) GetIpVersion() IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *Subnet) GetIpVersionOk() (*IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *Subnet) SetIpVersion(v IpVersion)`

SetIpVersion sets IpVersion field to given value.


### GetNetworkAddress

`func (o *Subnet) GetNetworkAddress() string`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *Subnet) GetNetworkAddressOk() (*string, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *Subnet) SetNetworkAddress(v string)`

SetNetworkAddress sets NetworkAddress field to given value.


### GetPrefixLength

`func (o *Subnet) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Subnet) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Subnet) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetNetmask

`func (o *Subnet) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *Subnet) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *Subnet) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetDefaultGatewayAddress

`func (o *Subnet) GetDefaultGatewayAddress() string`

GetDefaultGatewayAddress returns the DefaultGatewayAddress field if non-nil, zero value otherwise.

### GetDefaultGatewayAddressOk

`func (o *Subnet) GetDefaultGatewayAddressOk() (*string, bool)`

GetDefaultGatewayAddressOk returns a tuple with the DefaultGatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAddress

`func (o *Subnet) SetDefaultGatewayAddress(v string)`

SetDefaultGatewayAddress sets DefaultGatewayAddress field to given value.


### GetIsPool

`func (o *Subnet) GetIsPool() bool`

GetIsPool returns the IsPool field if non-nil, zero value otherwise.

### GetIsPoolOk

`func (o *Subnet) GetIsPoolOk() (*bool, bool)`

GetIsPoolOk returns a tuple with the IsPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPool

`func (o *Subnet) SetIsPool(v bool)`

SetIsPool sets IsPool field to given value.


### GetAllocationDenylist

`func (o *Subnet) GetAllocationDenylist() []AddressRange`

GetAllocationDenylist returns the AllocationDenylist field if non-nil, zero value otherwise.

### GetAllocationDenylistOk

`func (o *Subnet) GetAllocationDenylistOk() (*[]AddressRange, bool)`

GetAllocationDenylistOk returns a tuple with the AllocationDenylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDenylist

`func (o *Subnet) SetAllocationDenylist(v []AddressRange)`

SetAllocationDenylist sets AllocationDenylist field to given value.


### GetChildOverlapAllowRules

`func (o *Subnet) GetChildOverlapAllowRules() []string`

GetChildOverlapAllowRules returns the ChildOverlapAllowRules field if non-nil, zero value otherwise.

### GetChildOverlapAllowRulesOk

`func (o *Subnet) GetChildOverlapAllowRulesOk() (*[]string, bool)`

GetChildOverlapAllowRulesOk returns a tuple with the ChildOverlapAllowRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOverlapAllowRules

`func (o *Subnet) SetChildOverlapAllowRules(v []string)`

SetChildOverlapAllowRules sets ChildOverlapAllowRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


