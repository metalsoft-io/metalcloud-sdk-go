# CreateNetworkEquipmentBreakout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortName** | **string** | Device-side port name to break out. | 
**Groups** | [**[]BreakoutGroup**](BreakoutGroup.md) | One or more breakout groups to apply to the port. | 

## Methods

### NewCreateNetworkEquipmentBreakout

`func NewCreateNetworkEquipmentBreakout(portName string, groups []BreakoutGroup, ) *CreateNetworkEquipmentBreakout`

NewCreateNetworkEquipmentBreakout instantiates a new CreateNetworkEquipmentBreakout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkEquipmentBreakoutWithDefaults

`func NewCreateNetworkEquipmentBreakoutWithDefaults() *CreateNetworkEquipmentBreakout`

NewCreateNetworkEquipmentBreakoutWithDefaults instantiates a new CreateNetworkEquipmentBreakout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortName

`func (o *CreateNetworkEquipmentBreakout) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *CreateNetworkEquipmentBreakout) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *CreateNetworkEquipmentBreakout) SetPortName(v string)`

SetPortName sets PortName field to given value.


### GetGroups

`func (o *CreateNetworkEquipmentBreakout) GetGroups() []BreakoutGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateNetworkEquipmentBreakout) GetGroupsOk() (*[]BreakoutGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateNetworkEquipmentBreakout) SetGroups(v []BreakoutGroup)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


