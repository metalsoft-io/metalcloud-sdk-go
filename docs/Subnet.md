# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Id of the Subnet | 
**Revision** | **float32** | Revision of the Subnet | 
**Label** | Pointer to **string** |  | [optional] 
**Name** | **string** | Name of the Subnet | 
**IpVersion** | **string** | IP version | 
**ParentSubnetId** | **float32** | ID of the parent subnet | 
**NetworkAddress** | **string** |  | 
**PrefixLength** | **float32** |  | 
**Netmask** | **string** |  | 
**Gateway** | **string** |  | 
**IsPool** | **bool** |  | 
**VrfId** | **float32** |  | 
**AllocationDenylist** | **[]string** |  | 
**AllowedChildOverlapConditions** | **[]string** |  | 
**Tags** | **map[string]interface{}** |  | 
**Metadata** | **map[string]interface{}** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewSubnet

`func NewSubnet(id float32, revision float32, name string, ipVersion string, parentSubnetId float32, networkAddress string, prefixLength float32, netmask string, gateway string, isPool bool, vrfId float32, allocationDenylist []string, allowedChildOverlapConditions []string, tags map[string]interface{}, metadata map[string]interface{}, createdAt string, updatedAt string, ) *Subnet`

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

`func (o *Subnet) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *Subnet) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Subnet) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Subnet) SetRevision(v float32)`

SetRevision sets Revision field to given value.


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

### HasLabel

`func (o *Subnet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

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


### GetIpVersion

`func (o *Subnet) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *Subnet) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *Subnet) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetParentSubnetId

`func (o *Subnet) GetParentSubnetId() float32`

GetParentSubnetId returns the ParentSubnetId field if non-nil, zero value otherwise.

### GetParentSubnetIdOk

`func (o *Subnet) GetParentSubnetIdOk() (*float32, bool)`

GetParentSubnetIdOk returns a tuple with the ParentSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSubnetId

`func (o *Subnet) SetParentSubnetId(v float32)`

SetParentSubnetId sets ParentSubnetId field to given value.


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

`func (o *Subnet) GetPrefixLength() float32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Subnet) GetPrefixLengthOk() (*float32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Subnet) SetPrefixLength(v float32)`

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


### GetGateway

`func (o *Subnet) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Subnet) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Subnet) SetGateway(v string)`

SetGateway sets Gateway field to given value.


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


### GetVrfId

`func (o *Subnet) GetVrfId() float32`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *Subnet) GetVrfIdOk() (*float32, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *Subnet) SetVrfId(v float32)`

SetVrfId sets VrfId field to given value.


### GetAllocationDenylist

`func (o *Subnet) GetAllocationDenylist() []string`

GetAllocationDenylist returns the AllocationDenylist field if non-nil, zero value otherwise.

### GetAllocationDenylistOk

`func (o *Subnet) GetAllocationDenylistOk() (*[]string, bool)`

GetAllocationDenylistOk returns a tuple with the AllocationDenylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDenylist

`func (o *Subnet) SetAllocationDenylist(v []string)`

SetAllocationDenylist sets AllocationDenylist field to given value.


### GetAllowedChildOverlapConditions

`func (o *Subnet) GetAllowedChildOverlapConditions() []string`

GetAllowedChildOverlapConditions returns the AllowedChildOverlapConditions field if non-nil, zero value otherwise.

### GetAllowedChildOverlapConditionsOk

`func (o *Subnet) GetAllowedChildOverlapConditionsOk() (*[]string, bool)`

GetAllowedChildOverlapConditionsOk returns a tuple with the AllowedChildOverlapConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChildOverlapConditions

`func (o *Subnet) SetAllowedChildOverlapConditions(v []string)`

SetAllowedChildOverlapConditions sets AllowedChildOverlapConditions field to given value.


### GetTags

`func (o *Subnet) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Subnet) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Subnet) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetMetadata

`func (o *Subnet) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Subnet) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Subnet) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *Subnet) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subnet) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subnet) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Subnet) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subnet) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subnet) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *Subnet) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Subnet) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Subnet) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Subnet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


