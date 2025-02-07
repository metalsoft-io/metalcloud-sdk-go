# InfrastructureDeployShutdownOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptSoftShutdown** | **bool** | Attempt soft shutdown | 
**SoftShutdownTimeout** | **float32** | Soft shutdown timeout in seconds | 
**AttemptHardShutdown** | **bool** | Attempt hard shutdown after softShutdownTimeout expires | 
**ForceShutdown** | **bool** | Force shutdown | 

## Methods

### NewInfrastructureDeployShutdownOptions

`func NewInfrastructureDeployShutdownOptions(attemptSoftShutdown bool, softShutdownTimeout float32, attemptHardShutdown bool, forceShutdown bool, ) *InfrastructureDeployShutdownOptions`

NewInfrastructureDeployShutdownOptions instantiates a new InfrastructureDeployShutdownOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureDeployShutdownOptionsWithDefaults

`func NewInfrastructureDeployShutdownOptionsWithDefaults() *InfrastructureDeployShutdownOptions`

NewInfrastructureDeployShutdownOptionsWithDefaults instantiates a new InfrastructureDeployShutdownOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptSoftShutdown

`func (o *InfrastructureDeployShutdownOptions) GetAttemptSoftShutdown() bool`

GetAttemptSoftShutdown returns the AttemptSoftShutdown field if non-nil, zero value otherwise.

### GetAttemptSoftShutdownOk

`func (o *InfrastructureDeployShutdownOptions) GetAttemptSoftShutdownOk() (*bool, bool)`

GetAttemptSoftShutdownOk returns a tuple with the AttemptSoftShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptSoftShutdown

`func (o *InfrastructureDeployShutdownOptions) SetAttemptSoftShutdown(v bool)`

SetAttemptSoftShutdown sets AttemptSoftShutdown field to given value.


### GetSoftShutdownTimeout

`func (o *InfrastructureDeployShutdownOptions) GetSoftShutdownTimeout() float32`

GetSoftShutdownTimeout returns the SoftShutdownTimeout field if non-nil, zero value otherwise.

### GetSoftShutdownTimeoutOk

`func (o *InfrastructureDeployShutdownOptions) GetSoftShutdownTimeoutOk() (*float32, bool)`

GetSoftShutdownTimeoutOk returns a tuple with the SoftShutdownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftShutdownTimeout

`func (o *InfrastructureDeployShutdownOptions) SetSoftShutdownTimeout(v float32)`

SetSoftShutdownTimeout sets SoftShutdownTimeout field to given value.


### GetAttemptHardShutdown

`func (o *InfrastructureDeployShutdownOptions) GetAttemptHardShutdown() bool`

GetAttemptHardShutdown returns the AttemptHardShutdown field if non-nil, zero value otherwise.

### GetAttemptHardShutdownOk

`func (o *InfrastructureDeployShutdownOptions) GetAttemptHardShutdownOk() (*bool, bool)`

GetAttemptHardShutdownOk returns a tuple with the AttemptHardShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptHardShutdown

`func (o *InfrastructureDeployShutdownOptions) SetAttemptHardShutdown(v bool)`

SetAttemptHardShutdown sets AttemptHardShutdown field to given value.


### GetForceShutdown

`func (o *InfrastructureDeployShutdownOptions) GetForceShutdown() bool`

GetForceShutdown returns the ForceShutdown field if non-nil, zero value otherwise.

### GetForceShutdownOk

`func (o *InfrastructureDeployShutdownOptions) GetForceShutdownOk() (*bool, bool)`

GetForceShutdownOk returns a tuple with the ForceShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceShutdown

`func (o *InfrastructureDeployShutdownOptions) SetForceShutdown(v bool)`

SetForceShutdown sets ForceShutdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


