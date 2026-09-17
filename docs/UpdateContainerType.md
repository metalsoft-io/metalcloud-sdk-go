# UpdateContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the Container type | [optional] 
**Label** | Pointer to **string** | Label of the Container type | [optional] 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the Container is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Container Type. | [optional] 
**ForUnmanagedContainersOnly** | Pointer to **float32** | Flag to indicate if the Container is for unmanaged Containers only. 1 for true, 0 for false. Default is 0. | [optional] 

## Methods

### NewUpdateContainerType

`func NewUpdateContainerType() *UpdateContainerType`

NewUpdateContainerType instantiates a new UpdateContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerTypeWithDefaults

`func NewUpdateContainerTypeWithDefaults() *UpdateContainerType`

NewUpdateContainerTypeWithDefaults instantiates a new UpdateContainerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateContainerType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateContainerType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateContainerType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateContainerType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateContainerType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateContainerType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateContainerType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateContainerType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsExperimental

`func (o *UpdateContainerType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *UpdateContainerType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *UpdateContainerType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *UpdateContainerType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *UpdateContainerType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateContainerType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateContainerType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateContainerType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetForUnmanagedContainersOnly

`func (o *UpdateContainerType) GetForUnmanagedContainersOnly() float32`

GetForUnmanagedContainersOnly returns the ForUnmanagedContainersOnly field if non-nil, zero value otherwise.

### GetForUnmanagedContainersOnlyOk

`func (o *UpdateContainerType) GetForUnmanagedContainersOnlyOk() (*float32, bool)`

GetForUnmanagedContainersOnlyOk returns a tuple with the ForUnmanagedContainersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedContainersOnly

`func (o *UpdateContainerType) SetForUnmanagedContainersOnly(v float32)`

SetForUnmanagedContainersOnly sets ForUnmanagedContainersOnly field to given value.

### HasForUnmanagedContainersOnly

`func (o *UpdateContainerType) HasForUnmanagedContainersOnly() bool`

HasForUnmanagedContainersOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


