# GlobalFirmwareUpgradeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **bool** | Indicates whether the firmware upgrade policies are activated globally. | [optional] 
**UpgradeStartTime** | Pointer to **string** | The timestamp when the firmware upgrade policies are allowed to be executed. Only the hours and minutes are considered. | [optional] 
**UpgradeEndTime** | Pointer to **string** | The timestamp when the firmware upgrade policies are no longer allowed to be executed. Only the hours and minutes are considered. | [optional] 

## Methods

### NewGlobalFirmwareUpgradeConfiguration

`func NewGlobalFirmwareUpgradeConfiguration() *GlobalFirmwareUpgradeConfiguration`

NewGlobalFirmwareUpgradeConfiguration instantiates a new GlobalFirmwareUpgradeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalFirmwareUpgradeConfigurationWithDefaults

`func NewGlobalFirmwareUpgradeConfigurationWithDefaults() *GlobalFirmwareUpgradeConfiguration`

NewGlobalFirmwareUpgradeConfigurationWithDefaults instantiates a new GlobalFirmwareUpgradeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *GlobalFirmwareUpgradeConfiguration) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *GlobalFirmwareUpgradeConfiguration) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *GlobalFirmwareUpgradeConfiguration) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *GlobalFirmwareUpgradeConfiguration) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetUpgradeStartTime

`func (o *GlobalFirmwareUpgradeConfiguration) GetUpgradeStartTime() string`

GetUpgradeStartTime returns the UpgradeStartTime field if non-nil, zero value otherwise.

### GetUpgradeStartTimeOk

`func (o *GlobalFirmwareUpgradeConfiguration) GetUpgradeStartTimeOk() (*string, bool)`

GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStartTime

`func (o *GlobalFirmwareUpgradeConfiguration) SetUpgradeStartTime(v string)`

SetUpgradeStartTime sets UpgradeStartTime field to given value.

### HasUpgradeStartTime

`func (o *GlobalFirmwareUpgradeConfiguration) HasUpgradeStartTime() bool`

HasUpgradeStartTime returns a boolean if a field has been set.

### GetUpgradeEndTime

`func (o *GlobalFirmwareUpgradeConfiguration) GetUpgradeEndTime() string`

GetUpgradeEndTime returns the UpgradeEndTime field if non-nil, zero value otherwise.

### GetUpgradeEndTimeOk

`func (o *GlobalFirmwareUpgradeConfiguration) GetUpgradeEndTimeOk() (*string, bool)`

GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeEndTime

`func (o *GlobalFirmwareUpgradeConfiguration) SetUpgradeEndTime(v string)`

SetUpgradeEndTime sets UpgradeEndTime field to given value.

### HasUpgradeEndTime

`func (o *GlobalFirmwareUpgradeConfiguration) HasUpgradeEndTime() bool`

HasUpgradeEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


