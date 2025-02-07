# ScheduleFirmwareUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleUpdateTimestamp** | Pointer to **string** | The time when the firmware update is scheduled to run. | [optional] 
**ConfirmationRequired** | Pointer to **bool** | Flag to indicate if the firmware update requires confirmation. | [optional] 

## Methods

### NewScheduleFirmwareUpgrade

`func NewScheduleFirmwareUpgrade() *ScheduleFirmwareUpgrade`

NewScheduleFirmwareUpgrade instantiates a new ScheduleFirmwareUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleFirmwareUpgradeWithDefaults

`func NewScheduleFirmwareUpgradeWithDefaults() *ScheduleFirmwareUpgrade`

NewScheduleFirmwareUpgradeWithDefaults instantiates a new ScheduleFirmwareUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleUpdateTimestamp

`func (o *ScheduleFirmwareUpgrade) GetScheduleUpdateTimestamp() string`

GetScheduleUpdateTimestamp returns the ScheduleUpdateTimestamp field if non-nil, zero value otherwise.

### GetScheduleUpdateTimestampOk

`func (o *ScheduleFirmwareUpgrade) GetScheduleUpdateTimestampOk() (*string, bool)`

GetScheduleUpdateTimestampOk returns a tuple with the ScheduleUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUpdateTimestamp

`func (o *ScheduleFirmwareUpgrade) SetScheduleUpdateTimestamp(v string)`

SetScheduleUpdateTimestamp sets ScheduleUpdateTimestamp field to given value.

### HasScheduleUpdateTimestamp

`func (o *ScheduleFirmwareUpgrade) HasScheduleUpdateTimestamp() bool`

HasScheduleUpdateTimestamp returns a boolean if a field has been set.

### GetConfirmationRequired

`func (o *ScheduleFirmwareUpgrade) GetConfirmationRequired() bool`

GetConfirmationRequired returns the ConfirmationRequired field if non-nil, zero value otherwise.

### GetConfirmationRequiredOk

`func (o *ScheduleFirmwareUpgrade) GetConfirmationRequiredOk() (*bool, bool)`

GetConfirmationRequiredOk returns a tuple with the ConfirmationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationRequired

`func (o *ScheduleFirmwareUpgrade) SetConfirmationRequired(v bool)`

SetConfirmationRequired sets ConfirmationRequired field to given value.

### HasConfirmationRequired

`func (o *ScheduleFirmwareUpgrade) HasConfirmationRequired() bool`

HasConfirmationRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


