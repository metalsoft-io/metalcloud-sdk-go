# BatchScheduleServerFirmwareUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIds** | **[]float32** | The list of server ids to upgrade. | 
**ScheduleUpdateTimestamp** | Pointer to **string** | The time when the firmware update is scheduled to run. | [optional] 
**ConfirmationRequired** | Pointer to **bool** | Flag to indicate if the firmware update requires confirmation. | [optional] 

## Methods

### NewBatchScheduleServerFirmwareUpgrade

`func NewBatchScheduleServerFirmwareUpgrade(serverIds []float32, ) *BatchScheduleServerFirmwareUpgrade`

NewBatchScheduleServerFirmwareUpgrade instantiates a new BatchScheduleServerFirmwareUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchScheduleServerFirmwareUpgradeWithDefaults

`func NewBatchScheduleServerFirmwareUpgradeWithDefaults() *BatchScheduleServerFirmwareUpgrade`

NewBatchScheduleServerFirmwareUpgradeWithDefaults instantiates a new BatchScheduleServerFirmwareUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerIds

`func (o *BatchScheduleServerFirmwareUpgrade) GetServerIds() []float32`

GetServerIds returns the ServerIds field if non-nil, zero value otherwise.

### GetServerIdsOk

`func (o *BatchScheduleServerFirmwareUpgrade) GetServerIdsOk() (*[]float32, bool)`

GetServerIdsOk returns a tuple with the ServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIds

`func (o *BatchScheduleServerFirmwareUpgrade) SetServerIds(v []float32)`

SetServerIds sets ServerIds field to given value.


### GetScheduleUpdateTimestamp

`func (o *BatchScheduleServerFirmwareUpgrade) GetScheduleUpdateTimestamp() string`

GetScheduleUpdateTimestamp returns the ScheduleUpdateTimestamp field if non-nil, zero value otherwise.

### GetScheduleUpdateTimestampOk

`func (o *BatchScheduleServerFirmwareUpgrade) GetScheduleUpdateTimestampOk() (*string, bool)`

GetScheduleUpdateTimestampOk returns a tuple with the ScheduleUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUpdateTimestamp

`func (o *BatchScheduleServerFirmwareUpgrade) SetScheduleUpdateTimestamp(v string)`

SetScheduleUpdateTimestamp sets ScheduleUpdateTimestamp field to given value.

### HasScheduleUpdateTimestamp

`func (o *BatchScheduleServerFirmwareUpgrade) HasScheduleUpdateTimestamp() bool`

HasScheduleUpdateTimestamp returns a boolean if a field has been set.

### GetConfirmationRequired

`func (o *BatchScheduleServerFirmwareUpgrade) GetConfirmationRequired() bool`

GetConfirmationRequired returns the ConfirmationRequired field if non-nil, zero value otherwise.

### GetConfirmationRequiredOk

`func (o *BatchScheduleServerFirmwareUpgrade) GetConfirmationRequiredOk() (*bool, bool)`

GetConfirmationRequiredOk returns a tuple with the ConfirmationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationRequired

`func (o *BatchScheduleServerFirmwareUpgrade) SetConfirmationRequired(v bool)`

SetConfirmationRequired sets ConfirmationRequired field to given value.

### HasConfirmationRequired

`func (o *BatchScheduleServerFirmwareUpgrade) HasConfirmationRequired() bool`

HasConfirmationRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


