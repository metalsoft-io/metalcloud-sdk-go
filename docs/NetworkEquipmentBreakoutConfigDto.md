# NetworkEquipmentBreakoutConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakoutGroups** | Pointer to [**[]BreakoutGroup**](BreakoutGroup.md) | Desired breakout groups staged for the next deploy (NULL &#x3D; none staged). | [optional] 
**Revision** | **int64** | Optimistic-lock revision of the config buffer (independent of the main breakout revision). | 

## Methods

### NewNetworkEquipmentBreakoutConfigDto

`func NewNetworkEquipmentBreakoutConfigDto(revision int64, ) *NetworkEquipmentBreakoutConfigDto`

NewNetworkEquipmentBreakoutConfigDto instantiates a new NetworkEquipmentBreakoutConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEquipmentBreakoutConfigDtoWithDefaults

`func NewNetworkEquipmentBreakoutConfigDtoWithDefaults() *NetworkEquipmentBreakoutConfigDto`

NewNetworkEquipmentBreakoutConfigDtoWithDefaults instantiates a new NetworkEquipmentBreakoutConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakoutGroups

`func (o *NetworkEquipmentBreakoutConfigDto) GetBreakoutGroups() []BreakoutGroup`

GetBreakoutGroups returns the BreakoutGroups field if non-nil, zero value otherwise.

### GetBreakoutGroupsOk

`func (o *NetworkEquipmentBreakoutConfigDto) GetBreakoutGroupsOk() (*[]BreakoutGroup, bool)`

GetBreakoutGroupsOk returns a tuple with the BreakoutGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakoutGroups

`func (o *NetworkEquipmentBreakoutConfigDto) SetBreakoutGroups(v []BreakoutGroup)`

SetBreakoutGroups sets BreakoutGroups field to given value.

### HasBreakoutGroups

`func (o *NetworkEquipmentBreakoutConfigDto) HasBreakoutGroups() bool`

HasBreakoutGroups returns a boolean if a field has been set.

### SetBreakoutGroupsNil

`func (o *NetworkEquipmentBreakoutConfigDto) SetBreakoutGroupsNil(b bool)`

 SetBreakoutGroupsNil sets the value for BreakoutGroups to be an explicit nil

### UnsetBreakoutGroups
`func (o *NetworkEquipmentBreakoutConfigDto) UnsetBreakoutGroups()`

UnsetBreakoutGroups ensures that no value is present for BreakoutGroups, not even an explicit nil
### GetRevision

`func (o *NetworkEquipmentBreakoutConfigDto) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkEquipmentBreakoutConfigDto) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkEquipmentBreakoutConfigDto) SetRevision(v int64)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


