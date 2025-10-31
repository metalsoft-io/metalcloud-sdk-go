# ServerRegistrationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FoundMinimumNumberOfConnectedInterfaces** | Pointer to **string** |  | [optional] 
**RaidControllersReset** | Pointer to **string** |  | [optional] 
**ReasonsForFailure** | Pointer to **[]string** | The reasons for failure during registration. | [optional] 

## Methods

### NewServerRegistrationResult

`func NewServerRegistrationResult() *ServerRegistrationResult`

NewServerRegistrationResult instantiates a new ServerRegistrationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRegistrationResultWithDefaults

`func NewServerRegistrationResultWithDefaults() *ServerRegistrationResult`

NewServerRegistrationResultWithDefaults instantiates a new ServerRegistrationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFoundMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationResult) GetFoundMinimumNumberOfConnectedInterfaces() string`

GetFoundMinimumNumberOfConnectedInterfaces returns the FoundMinimumNumberOfConnectedInterfaces field if non-nil, zero value otherwise.

### GetFoundMinimumNumberOfConnectedInterfacesOk

`func (o *ServerRegistrationResult) GetFoundMinimumNumberOfConnectedInterfacesOk() (*string, bool)`

GetFoundMinimumNumberOfConnectedInterfacesOk returns a tuple with the FoundMinimumNumberOfConnectedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationResult) SetFoundMinimumNumberOfConnectedInterfaces(v string)`

SetFoundMinimumNumberOfConnectedInterfaces sets FoundMinimumNumberOfConnectedInterfaces field to given value.

### HasFoundMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationResult) HasFoundMinimumNumberOfConnectedInterfaces() bool`

HasFoundMinimumNumberOfConnectedInterfaces returns a boolean if a field has been set.

### GetRaidControllersReset

`func (o *ServerRegistrationResult) GetRaidControllersReset() string`

GetRaidControllersReset returns the RaidControllersReset field if non-nil, zero value otherwise.

### GetRaidControllersResetOk

`func (o *ServerRegistrationResult) GetRaidControllersResetOk() (*string, bool)`

GetRaidControllersResetOk returns a tuple with the RaidControllersReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidControllersReset

`func (o *ServerRegistrationResult) SetRaidControllersReset(v string)`

SetRaidControllersReset sets RaidControllersReset field to given value.

### HasRaidControllersReset

`func (o *ServerRegistrationResult) HasRaidControllersReset() bool`

HasRaidControllersReset returns a boolean if a field has been set.

### GetReasonsForFailure

`func (o *ServerRegistrationResult) GetReasonsForFailure() []string`

GetReasonsForFailure returns the ReasonsForFailure field if non-nil, zero value otherwise.

### GetReasonsForFailureOk

`func (o *ServerRegistrationResult) GetReasonsForFailureOk() (*[]string, bool)`

GetReasonsForFailureOk returns a tuple with the ReasonsForFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonsForFailure

`func (o *ServerRegistrationResult) SetReasonsForFailure(v []string)`

SetReasonsForFailure sets ReasonsForFailure field to given value.

### HasReasonsForFailure

`func (o *ServerRegistrationResult) HasReasonsForFailure() bool`

HasReasonsForFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


