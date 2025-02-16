# UpdateServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The label of the server type. | [optional] 
**AllowedVendorSkuIds** | Pointer to **[]string** | The list of allowed SKU ids for the server type. | [optional] 
**IsExperimental** | Pointer to **float32** | Flag specifying if the server type is experimental. | [optional] [default to 0]
**Tags** | Pointer to **[]string** | The tags for the server type. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the server type. | [optional] 

## Methods

### NewUpdateServerType

`func NewUpdateServerType() *UpdateServerType`

NewUpdateServerType instantiates a new UpdateServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerTypeWithDefaults

`func NewUpdateServerTypeWithDefaults() *UpdateServerType`

NewUpdateServerTypeWithDefaults instantiates a new UpdateServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateServerType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateServerType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateServerType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateServerType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAllowedVendorSkuIds

`func (o *UpdateServerType) GetAllowedVendorSkuIds() []string`

GetAllowedVendorSkuIds returns the AllowedVendorSkuIds field if non-nil, zero value otherwise.

### GetAllowedVendorSkuIdsOk

`func (o *UpdateServerType) GetAllowedVendorSkuIdsOk() (*[]string, bool)`

GetAllowedVendorSkuIdsOk returns a tuple with the AllowedVendorSkuIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVendorSkuIds

`func (o *UpdateServerType) SetAllowedVendorSkuIds(v []string)`

SetAllowedVendorSkuIds sets AllowedVendorSkuIds field to given value.

### HasAllowedVendorSkuIds

`func (o *UpdateServerType) HasAllowedVendorSkuIds() bool`

HasAllowedVendorSkuIds returns a boolean if a field has been set.

### GetIsExperimental

`func (o *UpdateServerType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *UpdateServerType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *UpdateServerType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *UpdateServerType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *UpdateServerType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateServerType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateServerType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateServerType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateServerType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateServerType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateServerType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateServerType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


