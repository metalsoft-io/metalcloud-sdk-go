# ServerFirmwareUpgradePolicyApplyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scheduled** | **[]float32** | The list of server component ids that have been successfully scheduled for firmware upgrade. | 
**FailedToSchedule** | **[]float32** | The list of server component ids that have failed to be scheduled for firmware upgrade. | 

## Methods

### NewServerFirmwareUpgradePolicyApplyResult

`func NewServerFirmwareUpgradePolicyApplyResult(scheduled []float32, failedToSchedule []float32, ) *ServerFirmwareUpgradePolicyApplyResult`

NewServerFirmwareUpgradePolicyApplyResult instantiates a new ServerFirmwareUpgradePolicyApplyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareUpgradePolicyApplyResultWithDefaults

`func NewServerFirmwareUpgradePolicyApplyResultWithDefaults() *ServerFirmwareUpgradePolicyApplyResult`

NewServerFirmwareUpgradePolicyApplyResultWithDefaults instantiates a new ServerFirmwareUpgradePolicyApplyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduled

`func (o *ServerFirmwareUpgradePolicyApplyResult) GetScheduled() []float32`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *ServerFirmwareUpgradePolicyApplyResult) GetScheduledOk() (*[]float32, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *ServerFirmwareUpgradePolicyApplyResult) SetScheduled(v []float32)`

SetScheduled sets Scheduled field to given value.


### GetFailedToSchedule

`func (o *ServerFirmwareUpgradePolicyApplyResult) GetFailedToSchedule() []float32`

GetFailedToSchedule returns the FailedToSchedule field if non-nil, zero value otherwise.

### GetFailedToScheduleOk

`func (o *ServerFirmwareUpgradePolicyApplyResult) GetFailedToScheduleOk() (*[]float32, bool)`

GetFailedToScheduleOk returns a tuple with the FailedToSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToSchedule

`func (o *ServerFirmwareUpgradePolicyApplyResult) SetFailedToSchedule(v []float32)`

SetFailedToSchedule sets FailedToSchedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


