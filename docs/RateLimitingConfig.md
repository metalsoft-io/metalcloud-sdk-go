# RateLimitingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **map[string]interface{}** | Whether API rate limiting is enabled | [optional] 
**ReadRequestsPerMinute** | Pointer to **map[string]interface{}** | Maximum read (GET) requests allowed per minute | [optional] 
**WriteRequestsPerMinute** | Pointer to **map[string]interface{}** | Maximum write (POST/PUT/PATCH/DELETE) requests allowed per minute | [optional] 

## Methods

### NewRateLimitingConfig

`func NewRateLimitingConfig() *RateLimitingConfig`

NewRateLimitingConfig instantiates a new RateLimitingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitingConfigWithDefaults

`func NewRateLimitingConfigWithDefaults() *RateLimitingConfig`

NewRateLimitingConfigWithDefaults instantiates a new RateLimitingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RateLimitingConfig) GetEnabled() map[string]interface{}`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RateLimitingConfig) GetEnabledOk() (*map[string]interface{}, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RateLimitingConfig) SetEnabled(v map[string]interface{})`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RateLimitingConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReadRequestsPerMinute

`func (o *RateLimitingConfig) GetReadRequestsPerMinute() map[string]interface{}`

GetReadRequestsPerMinute returns the ReadRequestsPerMinute field if non-nil, zero value otherwise.

### GetReadRequestsPerMinuteOk

`func (o *RateLimitingConfig) GetReadRequestsPerMinuteOk() (*map[string]interface{}, bool)`

GetReadRequestsPerMinuteOk returns a tuple with the ReadRequestsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadRequestsPerMinute

`func (o *RateLimitingConfig) SetReadRequestsPerMinute(v map[string]interface{})`

SetReadRequestsPerMinute sets ReadRequestsPerMinute field to given value.

### HasReadRequestsPerMinute

`func (o *RateLimitingConfig) HasReadRequestsPerMinute() bool`

HasReadRequestsPerMinute returns a boolean if a field has been set.

### GetWriteRequestsPerMinute

`func (o *RateLimitingConfig) GetWriteRequestsPerMinute() map[string]interface{}`

GetWriteRequestsPerMinute returns the WriteRequestsPerMinute field if non-nil, zero value otherwise.

### GetWriteRequestsPerMinuteOk

`func (o *RateLimitingConfig) GetWriteRequestsPerMinuteOk() (*map[string]interface{}, bool)`

GetWriteRequestsPerMinuteOk returns a tuple with the WriteRequestsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteRequestsPerMinute

`func (o *RateLimitingConfig) SetWriteRequestsPerMinute(v map[string]interface{})`

SetWriteRequestsPerMinute sets WriteRequestsPerMinute field to given value.

### HasWriteRequestsPerMinute

`func (o *RateLimitingConfig) HasWriteRequestsPerMinute() bool`

HasWriteRequestsPerMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


