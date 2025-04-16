# RedundancyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**NetworkEndpointGroupRedundancyMode**](NetworkEndpointGroupRedundancyMode.md) | The redundancy mode | 
**Implementation** | Pointer to [**NullableRedundancyImplementation**](RedundancyImplementation.md) | The redundancy implementation configuration | [optional] 

## Methods

### NewRedundancyConfig

`func NewRedundancyConfig(mode NetworkEndpointGroupRedundancyMode, ) *RedundancyConfig`

NewRedundancyConfig instantiates a new RedundancyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedundancyConfigWithDefaults

`func NewRedundancyConfigWithDefaults() *RedundancyConfig`

NewRedundancyConfigWithDefaults instantiates a new RedundancyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RedundancyConfig) GetMode() NetworkEndpointGroupRedundancyMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RedundancyConfig) GetModeOk() (*NetworkEndpointGroupRedundancyMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RedundancyConfig) SetMode(v NetworkEndpointGroupRedundancyMode)`

SetMode sets Mode field to given value.


### GetImplementation

`func (o *RedundancyConfig) GetImplementation() RedundancyImplementation`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *RedundancyConfig) GetImplementationOk() (*RedundancyImplementation, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *RedundancyConfig) SetImplementation(v RedundancyImplementation)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *RedundancyConfig) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *RedundancyConfig) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *RedundancyConfig) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


