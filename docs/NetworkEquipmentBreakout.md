# NetworkEquipmentBreakout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Breakout id; address per-breakout operations by this. | 
**PortName** | **string** | Device-side port name the breakout applies to. | 
**BreakoutGroups** | Pointer to [**[]BreakoutGroup**](BreakoutGroup.md) | Applied breakout groups (what is currently on the device; NULL until a deploy confirms). The desired/staged groups are exposed inline as &#x60;config&#x60; (and editable via the /config sub-resource). | [optional] 
**ServiceStatus** | **string** |  | 
**PendingDelete** | **bool** |  | 
**Revision** | **int64** | Optimistic-lock revision. | 
**Config** | [**NetworkEquipmentBreakoutConfigDto**](NetworkEquipmentBreakoutConfigDto.md) |  | 

## Methods

### NewNetworkEquipmentBreakout

`func NewNetworkEquipmentBreakout(id int64, portName string, serviceStatus string, pendingDelete bool, revision int64, config NetworkEquipmentBreakoutConfigDto, ) *NetworkEquipmentBreakout`

NewNetworkEquipmentBreakout instantiates a new NetworkEquipmentBreakout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEquipmentBreakoutWithDefaults

`func NewNetworkEquipmentBreakoutWithDefaults() *NetworkEquipmentBreakout`

NewNetworkEquipmentBreakoutWithDefaults instantiates a new NetworkEquipmentBreakout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkEquipmentBreakout) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkEquipmentBreakout) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkEquipmentBreakout) SetId(v int64)`

SetId sets Id field to given value.


### GetPortName

`func (o *NetworkEquipmentBreakout) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *NetworkEquipmentBreakout) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *NetworkEquipmentBreakout) SetPortName(v string)`

SetPortName sets PortName field to given value.


### GetBreakoutGroups

`func (o *NetworkEquipmentBreakout) GetBreakoutGroups() []BreakoutGroup`

GetBreakoutGroups returns the BreakoutGroups field if non-nil, zero value otherwise.

### GetBreakoutGroupsOk

`func (o *NetworkEquipmentBreakout) GetBreakoutGroupsOk() (*[]BreakoutGroup, bool)`

GetBreakoutGroupsOk returns a tuple with the BreakoutGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakoutGroups

`func (o *NetworkEquipmentBreakout) SetBreakoutGroups(v []BreakoutGroup)`

SetBreakoutGroups sets BreakoutGroups field to given value.

### HasBreakoutGroups

`func (o *NetworkEquipmentBreakout) HasBreakoutGroups() bool`

HasBreakoutGroups returns a boolean if a field has been set.

### SetBreakoutGroupsNil

`func (o *NetworkEquipmentBreakout) SetBreakoutGroupsNil(b bool)`

 SetBreakoutGroupsNil sets the value for BreakoutGroups to be an explicit nil

### UnsetBreakoutGroups
`func (o *NetworkEquipmentBreakout) UnsetBreakoutGroups()`

UnsetBreakoutGroups ensures that no value is present for BreakoutGroups, not even an explicit nil
### GetServiceStatus

`func (o *NetworkEquipmentBreakout) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *NetworkEquipmentBreakout) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *NetworkEquipmentBreakout) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetPendingDelete

`func (o *NetworkEquipmentBreakout) GetPendingDelete() bool`

GetPendingDelete returns the PendingDelete field if non-nil, zero value otherwise.

### GetPendingDeleteOk

`func (o *NetworkEquipmentBreakout) GetPendingDeleteOk() (*bool, bool)`

GetPendingDeleteOk returns a tuple with the PendingDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDelete

`func (o *NetworkEquipmentBreakout) SetPendingDelete(v bool)`

SetPendingDelete sets PendingDelete field to given value.


### GetRevision

`func (o *NetworkEquipmentBreakout) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkEquipmentBreakout) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkEquipmentBreakout) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetConfig

`func (o *NetworkEquipmentBreakout) GetConfig() NetworkEquipmentBreakoutConfigDto`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkEquipmentBreakout) GetConfigOk() (*NetworkEquipmentBreakoutConfigDto, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkEquipmentBreakout) SetConfig(v NetworkEquipmentBreakoutConfigDto)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


