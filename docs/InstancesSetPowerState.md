# InstancesSetPowerState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | **[]string** | The list of instances to set the power state. | 
**PowerCommand** | **string** | The power state to set. | 

## Methods

### NewInstancesSetPowerState

`func NewInstancesSetPowerState(instances []string, powerCommand string, ) *InstancesSetPowerState`

NewInstancesSetPowerState instantiates a new InstancesSetPowerState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesSetPowerStateWithDefaults

`func NewInstancesSetPowerStateWithDefaults() *InstancesSetPowerState`

NewInstancesSetPowerStateWithDefaults instantiates a new InstancesSetPowerState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *InstancesSetPowerState) GetInstances() []string`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *InstancesSetPowerState) GetInstancesOk() (*[]string, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *InstancesSetPowerState) SetInstances(v []string)`

SetInstances sets Instances field to given value.


### GetPowerCommand

`func (o *InstancesSetPowerState) GetPowerCommand() string`

GetPowerCommand returns the PowerCommand field if non-nil, zero value otherwise.

### GetPowerCommandOk

`func (o *InstancesSetPowerState) GetPowerCommandOk() (*string, bool)`

GetPowerCommandOk returns a tuple with the PowerCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCommand

`func (o *InstancesSetPowerState) SetPowerCommand(v string)`

SetPowerCommand sets PowerCommand field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


