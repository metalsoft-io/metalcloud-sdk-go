# UpdateInfrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables in JSON format. | [optional] 
**Label** | Pointer to **string** | Label of the Infrastructure. | [optional] 

## Methods

### NewUpdateInfrastructure

`func NewUpdateInfrastructure() *UpdateInfrastructure`

NewUpdateInfrastructure instantiates a new UpdateInfrastructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfrastructureWithDefaults

`func NewUpdateInfrastructureWithDefaults() *UpdateInfrastructure`

NewUpdateInfrastructureWithDefaults instantiates a new UpdateInfrastructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomVariables

`func (o *UpdateInfrastructure) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateInfrastructure) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateInfrastructure) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateInfrastructure) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateInfrastructure) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateInfrastructure) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateInfrastructure) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateInfrastructure) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


