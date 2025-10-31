# ExtensionInfrastructureLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the logical network. | 
**ProfileLabel** | **string** | Logical network profile label. | 
**IpAllocations** | Pointer to [**[]ExtensionInfrastructureIpAllocation**](ExtensionInfrastructureIpAllocation.md) | Extra IP allocation for the infrastructure. | [optional] 
**IpRanges** | Pointer to [**[]ExtensionInfrastructureIpRanges**](ExtensionInfrastructureIpRanges.md) | Extra IP allocation for the infrastructure. | [optional] 

## Methods

### NewExtensionInfrastructureLogicalNetwork

`func NewExtensionInfrastructureLogicalNetwork(label string, profileLabel string, ) *ExtensionInfrastructureLogicalNetwork`

NewExtensionInfrastructureLogicalNetwork instantiates a new ExtensionInfrastructureLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInfrastructureLogicalNetworkWithDefaults

`func NewExtensionInfrastructureLogicalNetworkWithDefaults() *ExtensionInfrastructureLogicalNetwork`

NewExtensionInfrastructureLogicalNetworkWithDefaults instantiates a new ExtensionInfrastructureLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionInfrastructureLogicalNetwork) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInfrastructureLogicalNetwork) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInfrastructureLogicalNetwork) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetProfileLabel

`func (o *ExtensionInfrastructureLogicalNetwork) GetProfileLabel() string`

GetProfileLabel returns the ProfileLabel field if non-nil, zero value otherwise.

### GetProfileLabelOk

`func (o *ExtensionInfrastructureLogicalNetwork) GetProfileLabelOk() (*string, bool)`

GetProfileLabelOk returns a tuple with the ProfileLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileLabel

`func (o *ExtensionInfrastructureLogicalNetwork) SetProfileLabel(v string)`

SetProfileLabel sets ProfileLabel field to given value.


### GetIpAllocations

`func (o *ExtensionInfrastructureLogicalNetwork) GetIpAllocations() []ExtensionInfrastructureIpAllocation`

GetIpAllocations returns the IpAllocations field if non-nil, zero value otherwise.

### GetIpAllocationsOk

`func (o *ExtensionInfrastructureLogicalNetwork) GetIpAllocationsOk() (*[]ExtensionInfrastructureIpAllocation, bool)`

GetIpAllocationsOk returns a tuple with the IpAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocations

`func (o *ExtensionInfrastructureLogicalNetwork) SetIpAllocations(v []ExtensionInfrastructureIpAllocation)`

SetIpAllocations sets IpAllocations field to given value.

### HasIpAllocations

`func (o *ExtensionInfrastructureLogicalNetwork) HasIpAllocations() bool`

HasIpAllocations returns a boolean if a field has been set.

### GetIpRanges

`func (o *ExtensionInfrastructureLogicalNetwork) GetIpRanges() []ExtensionInfrastructureIpRanges`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *ExtensionInfrastructureLogicalNetwork) GetIpRangesOk() (*[]ExtensionInfrastructureIpRanges, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *ExtensionInfrastructureLogicalNetwork) SetIpRanges(v []ExtensionInfrastructureIpRanges)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *ExtensionInfrastructureLogicalNetwork) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


