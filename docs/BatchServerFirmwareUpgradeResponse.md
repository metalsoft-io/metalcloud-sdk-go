# BatchServerFirmwareUpgradeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | [**map[string]JobInfo**](JobInfo.md) | The job information for the successful firmware upgrade of a server. | 
**Failed** | **map[string]interface{}** | The error message for the failed firmware upgrade of a server. | 

## Methods

### NewBatchServerFirmwareUpgradeResponse

`func NewBatchServerFirmwareUpgradeResponse(successful map[string]JobInfo, failed map[string]interface{}, ) *BatchServerFirmwareUpgradeResponse`

NewBatchServerFirmwareUpgradeResponse instantiates a new BatchServerFirmwareUpgradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchServerFirmwareUpgradeResponseWithDefaults

`func NewBatchServerFirmwareUpgradeResponseWithDefaults() *BatchServerFirmwareUpgradeResponse`

NewBatchServerFirmwareUpgradeResponseWithDefaults instantiates a new BatchServerFirmwareUpgradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessful

`func (o *BatchServerFirmwareUpgradeResponse) GetSuccessful() map[string]JobInfo`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *BatchServerFirmwareUpgradeResponse) GetSuccessfulOk() (*map[string]JobInfo, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *BatchServerFirmwareUpgradeResponse) SetSuccessful(v map[string]JobInfo)`

SetSuccessful sets Successful field to given value.


### GetFailed

`func (o *BatchServerFirmwareUpgradeResponse) GetFailed() map[string]interface{}`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BatchServerFirmwareUpgradeResponse) GetFailedOk() (*map[string]interface{}, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BatchServerFirmwareUpgradeResponse) SetFailed(v map[string]interface{})`

SetFailed sets Failed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


