# FirmwareUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentFirmwareNewVersion** | Pointer to **string** | The desired firmware version. If not specified, the target version will be used. | [optional] 
**FirmwareBinaryUrl** | Pointer to **string** | The URL where the firmware file is located. If not specified, the firmware will be downloaded from the stored firmware information | [optional] 
**ApplyOnReboot** | Pointer to **bool** | Flag to indicate if the firmware upgrade should be performed at reboot. | [optional] [default to false]
**RebootRequired** | Pointer to **bool** | Flag to indicate if the firmware upgrade requires a reboot. Only taken into account if firmwareBinaryUrl is provided. | [optional] [default to false]

## Methods

### NewFirmwareUpgrade

`func NewFirmwareUpgrade() *FirmwareUpgrade`

NewFirmwareUpgrade instantiates a new FirmwareUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeWithDefaults

`func NewFirmwareUpgradeWithDefaults() *FirmwareUpgrade`

NewFirmwareUpgradeWithDefaults instantiates a new FirmwareUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentFirmwareNewVersion

`func (o *FirmwareUpgrade) GetComponentFirmwareNewVersion() string`

GetComponentFirmwareNewVersion returns the ComponentFirmwareNewVersion field if non-nil, zero value otherwise.

### GetComponentFirmwareNewVersionOk

`func (o *FirmwareUpgrade) GetComponentFirmwareNewVersionOk() (*string, bool)`

GetComponentFirmwareNewVersionOk returns a tuple with the ComponentFirmwareNewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentFirmwareNewVersion

`func (o *FirmwareUpgrade) SetComponentFirmwareNewVersion(v string)`

SetComponentFirmwareNewVersion sets ComponentFirmwareNewVersion field to given value.

### HasComponentFirmwareNewVersion

`func (o *FirmwareUpgrade) HasComponentFirmwareNewVersion() bool`

HasComponentFirmwareNewVersion returns a boolean if a field has been set.

### GetFirmwareBinaryUrl

`func (o *FirmwareUpgrade) GetFirmwareBinaryUrl() string`

GetFirmwareBinaryUrl returns the FirmwareBinaryUrl field if non-nil, zero value otherwise.

### GetFirmwareBinaryUrlOk

`func (o *FirmwareUpgrade) GetFirmwareBinaryUrlOk() (*string, bool)`

GetFirmwareBinaryUrlOk returns a tuple with the FirmwareBinaryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareBinaryUrl

`func (o *FirmwareUpgrade) SetFirmwareBinaryUrl(v string)`

SetFirmwareBinaryUrl sets FirmwareBinaryUrl field to given value.

### HasFirmwareBinaryUrl

`func (o *FirmwareUpgrade) HasFirmwareBinaryUrl() bool`

HasFirmwareBinaryUrl returns a boolean if a field has been set.

### GetApplyOnReboot

`func (o *FirmwareUpgrade) GetApplyOnReboot() bool`

GetApplyOnReboot returns the ApplyOnReboot field if non-nil, zero value otherwise.

### GetApplyOnRebootOk

`func (o *FirmwareUpgrade) GetApplyOnRebootOk() (*bool, bool)`

GetApplyOnRebootOk returns a tuple with the ApplyOnReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyOnReboot

`func (o *FirmwareUpgrade) SetApplyOnReboot(v bool)`

SetApplyOnReboot sets ApplyOnReboot field to given value.

### HasApplyOnReboot

`func (o *FirmwareUpgrade) HasApplyOnReboot() bool`

HasApplyOnReboot returns a boolean if a field has been set.

### GetRebootRequired

`func (o *FirmwareUpgrade) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *FirmwareUpgrade) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *FirmwareUpgrade) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.

### HasRebootRequired

`func (o *FirmwareUpgrade) HasRebootRequired() bool`

HasRebootRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


