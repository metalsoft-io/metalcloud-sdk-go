# PolicyEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **[]string** | Policy subject(s) | 
**Action** | **[]string** | Policy action(s) | 
**Conditions** | Pointer to **map[string]interface{}** | Policy condition(s) | [optional] 
**Fields** | Pointer to **[]string** | Policy subject fields(s) | [optional] 

## Methods

### NewPolicyEntry

`func NewPolicyEntry(subject []string, action []string, ) *PolicyEntry`

NewPolicyEntry instantiates a new PolicyEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEntryWithDefaults

`func NewPolicyEntryWithDefaults() *PolicyEntry`

NewPolicyEntryWithDefaults instantiates a new PolicyEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PolicyEntry) GetSubject() []string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PolicyEntry) GetSubjectOk() (*[]string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PolicyEntry) SetSubject(v []string)`

SetSubject sets Subject field to given value.


### GetAction

`func (o *PolicyEntry) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyEntry) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyEntry) SetAction(v []string)`

SetAction sets Action field to given value.


### GetConditions

`func (o *PolicyEntry) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PolicyEntry) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PolicyEntry) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PolicyEntry) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFields

`func (o *PolicyEntry) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PolicyEntry) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PolicyEntry) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PolicyEntry) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


