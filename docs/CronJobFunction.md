# CronJobFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ParamsSchema** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCronJobFunction

`func NewCronJobFunction(name string, ) *CronJobFunction`

NewCronJobFunction instantiates a new CronJobFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobFunctionWithDefaults

`func NewCronJobFunctionWithDefaults() *CronJobFunction`

NewCronJobFunctionWithDefaults instantiates a new CronJobFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CronJobFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CronJobFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CronJobFunction) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CronJobFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CronJobFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CronJobFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CronJobFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParamsSchema

`func (o *CronJobFunction) GetParamsSchema() map[string]interface{}`

GetParamsSchema returns the ParamsSchema field if non-nil, zero value otherwise.

### GetParamsSchemaOk

`func (o *CronJobFunction) GetParamsSchemaOk() (*map[string]interface{}, bool)`

GetParamsSchemaOk returns a tuple with the ParamsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSchema

`func (o *CronJobFunction) SetParamsSchema(v map[string]interface{})`

SetParamsSchema sets ParamsSchema field to given value.

### HasParamsSchema

`func (o *CronJobFunction) HasParamsSchema() bool`

HasParamsSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


