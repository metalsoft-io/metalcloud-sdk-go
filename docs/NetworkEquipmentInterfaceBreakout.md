# NetworkEquipmentInterfaceBreakout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakoutMode** | Pointer to **NullableString** | Applied breakout mode (e.g. \&quot;4x25G\&quot;, \&quot;2x50G\&quot;). NULL until the first deploy confirms. | [optional] 
**PendingBreakoutMode** | Pointer to **NullableString** | Pending breakout mode (operator-staged, not yet on the device). | [optional] 
**DeployStatus** | **string** |  | 

## Methods

### NewNetworkEquipmentInterfaceBreakout

`func NewNetworkEquipmentInterfaceBreakout(deployStatus string, ) *NetworkEquipmentInterfaceBreakout`

NewNetworkEquipmentInterfaceBreakout instantiates a new NetworkEquipmentInterfaceBreakout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEquipmentInterfaceBreakoutWithDefaults

`func NewNetworkEquipmentInterfaceBreakoutWithDefaults() *NetworkEquipmentInterfaceBreakout`

NewNetworkEquipmentInterfaceBreakoutWithDefaults instantiates a new NetworkEquipmentInterfaceBreakout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakoutMode

`func (o *NetworkEquipmentInterfaceBreakout) GetBreakoutMode() string`

GetBreakoutMode returns the BreakoutMode field if non-nil, zero value otherwise.

### GetBreakoutModeOk

`func (o *NetworkEquipmentInterfaceBreakout) GetBreakoutModeOk() (*string, bool)`

GetBreakoutModeOk returns a tuple with the BreakoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakoutMode

`func (o *NetworkEquipmentInterfaceBreakout) SetBreakoutMode(v string)`

SetBreakoutMode sets BreakoutMode field to given value.

### HasBreakoutMode

`func (o *NetworkEquipmentInterfaceBreakout) HasBreakoutMode() bool`

HasBreakoutMode returns a boolean if a field has been set.

### SetBreakoutModeNil

`func (o *NetworkEquipmentInterfaceBreakout) SetBreakoutModeNil(b bool)`

 SetBreakoutModeNil sets the value for BreakoutMode to be an explicit nil

### UnsetBreakoutMode
`func (o *NetworkEquipmentInterfaceBreakout) UnsetBreakoutMode()`

UnsetBreakoutMode ensures that no value is present for BreakoutMode, not even an explicit nil
### GetPendingBreakoutMode

`func (o *NetworkEquipmentInterfaceBreakout) GetPendingBreakoutMode() string`

GetPendingBreakoutMode returns the PendingBreakoutMode field if non-nil, zero value otherwise.

### GetPendingBreakoutModeOk

`func (o *NetworkEquipmentInterfaceBreakout) GetPendingBreakoutModeOk() (*string, bool)`

GetPendingBreakoutModeOk returns a tuple with the PendingBreakoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBreakoutMode

`func (o *NetworkEquipmentInterfaceBreakout) SetPendingBreakoutMode(v string)`

SetPendingBreakoutMode sets PendingBreakoutMode field to given value.

### HasPendingBreakoutMode

`func (o *NetworkEquipmentInterfaceBreakout) HasPendingBreakoutMode() bool`

HasPendingBreakoutMode returns a boolean if a field has been set.

### SetPendingBreakoutModeNil

`func (o *NetworkEquipmentInterfaceBreakout) SetPendingBreakoutModeNil(b bool)`

 SetPendingBreakoutModeNil sets the value for PendingBreakoutMode to be an explicit nil

### UnsetPendingBreakoutMode
`func (o *NetworkEquipmentInterfaceBreakout) UnsetPendingBreakoutMode()`

UnsetPendingBreakoutMode ensures that no value is present for PendingBreakoutMode, not even an explicit nil
### GetDeployStatus

`func (o *NetworkEquipmentInterfaceBreakout) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *NetworkEquipmentInterfaceBreakout) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *NetworkEquipmentInterfaceBreakout) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


