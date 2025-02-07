# JobCommandInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Mark the job for death | [optional] 
**ExecuteImmediately** | Pointer to **bool** | Execute the command immediately | [optional] [default to false]

## Methods

### NewJobCommandInfo

`func NewJobCommandInfo() *JobCommandInfo`

NewJobCommandInfo instantiates a new JobCommandInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobCommandInfoWithDefaults

`func NewJobCommandInfoWithDefaults() *JobCommandInfo`

NewJobCommandInfoWithDefaults instantiates a new JobCommandInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *JobCommandInfo) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *JobCommandInfo) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *JobCommandInfo) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *JobCommandInfo) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetExecuteImmediately

`func (o *JobCommandInfo) GetExecuteImmediately() bool`

GetExecuteImmediately returns the ExecuteImmediately field if non-nil, zero value otherwise.

### GetExecuteImmediatelyOk

`func (o *JobCommandInfo) GetExecuteImmediatelyOk() (*bool, bool)`

GetExecuteImmediatelyOk returns a tuple with the ExecuteImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteImmediately

`func (o *JobCommandInfo) SetExecuteImmediately(v bool)`

SetExecuteImmediately sets ExecuteImmediately field to given value.

### HasExecuteImmediately

`func (o *JobCommandInfo) HasExecuteImmediately() bool`

HasExecuteImmediately returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


