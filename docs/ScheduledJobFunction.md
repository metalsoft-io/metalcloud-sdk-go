# ScheduledJobFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ParamsSchema** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewScheduledJobFunction

`func NewScheduledJobFunction(name string, ) *ScheduledJobFunction`

NewScheduledJobFunction instantiates a new ScheduledJobFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledJobFunctionWithDefaults

`func NewScheduledJobFunctionWithDefaults() *ScheduledJobFunction`

NewScheduledJobFunctionWithDefaults instantiates a new ScheduledJobFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScheduledJobFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledJobFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledJobFunction) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScheduledJobFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledJobFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledJobFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledJobFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParamsSchema

`func (o *ScheduledJobFunction) GetParamsSchema() map[string]interface{}`

GetParamsSchema returns the ParamsSchema field if non-nil, zero value otherwise.

### GetParamsSchemaOk

`func (o *ScheduledJobFunction) GetParamsSchemaOk() (*map[string]interface{}, bool)`

GetParamsSchemaOk returns a tuple with the ParamsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSchema

`func (o *ScheduledJobFunction) SetParamsSchema(v map[string]interface{})`

SetParamsSchema sets ParamsSchema field to given value.

### HasParamsSchema

`func (o *ScheduledJobFunction) HasParamsSchema() bool`

HasParamsSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


