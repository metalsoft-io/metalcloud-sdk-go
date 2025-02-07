# UpdateServerComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirmwareTargetVersion** | Pointer to **string** | The target firmware version of the server component. | [optional] 
**FirmwareUpgradeNeedsConfirmation** | Pointer to **float32** | The flag indicating if the firmware upgrade of the server component requires confirmation. | [optional] 

## Methods

### NewUpdateServerComponent

`func NewUpdateServerComponent() *UpdateServerComponent`

NewUpdateServerComponent instantiates a new UpdateServerComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerComponentWithDefaults

`func NewUpdateServerComponentWithDefaults() *UpdateServerComponent`

NewUpdateServerComponentWithDefaults instantiates a new UpdateServerComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirmwareTargetVersion

`func (o *UpdateServerComponent) GetFirmwareTargetVersion() string`

GetFirmwareTargetVersion returns the FirmwareTargetVersion field if non-nil, zero value otherwise.

### GetFirmwareTargetVersionOk

`func (o *UpdateServerComponent) GetFirmwareTargetVersionOk() (*string, bool)`

GetFirmwareTargetVersionOk returns a tuple with the FirmwareTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareTargetVersion

`func (o *UpdateServerComponent) SetFirmwareTargetVersion(v string)`

SetFirmwareTargetVersion sets FirmwareTargetVersion field to given value.

### HasFirmwareTargetVersion

`func (o *UpdateServerComponent) HasFirmwareTargetVersion() bool`

HasFirmwareTargetVersion returns a boolean if a field has been set.

### GetFirmwareUpgradeNeedsConfirmation

`func (o *UpdateServerComponent) GetFirmwareUpgradeNeedsConfirmation() float32`

GetFirmwareUpgradeNeedsConfirmation returns the FirmwareUpgradeNeedsConfirmation field if non-nil, zero value otherwise.

### GetFirmwareUpgradeNeedsConfirmationOk

`func (o *UpdateServerComponent) GetFirmwareUpgradeNeedsConfirmationOk() (*float32, bool)`

GetFirmwareUpgradeNeedsConfirmationOk returns a tuple with the FirmwareUpgradeNeedsConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpgradeNeedsConfirmation

`func (o *UpdateServerComponent) SetFirmwareUpgradeNeedsConfirmation(v float32)`

SetFirmwareUpgradeNeedsConfirmation sets FirmwareUpgradeNeedsConfirmation field to given value.

### HasFirmwareUpgradeNeedsConfirmation

`func (o *UpdateServerComponent) HasFirmwareUpgradeNeedsConfirmation() bool`

HasFirmwareUpgradeNeedsConfirmation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


