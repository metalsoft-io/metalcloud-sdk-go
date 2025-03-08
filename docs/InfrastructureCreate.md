# InfrastructureCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the Infrastructure. | 
**SiteId** | **float32** | The ID of the site where the Infrastructure is located. | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables in JSON format. | [optional] 
**UserIdOwner** | Pointer to **float32** | User ID of the owner of the Infrastructure. | [optional] 
**Description** | Pointer to **string** | Description of the infrastructure. | [optional] 
**Meta** | Pointer to [**InfrastructureMeta**](InfrastructureMeta.md) |  | [optional] 

## Methods

### NewInfrastructureCreate

`func NewInfrastructureCreate(label string, siteId float32, ) *InfrastructureCreate`

NewInfrastructureCreate instantiates a new InfrastructureCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureCreateWithDefaults

`func NewInfrastructureCreateWithDefaults() *InfrastructureCreate`

NewInfrastructureCreateWithDefaults instantiates a new InfrastructureCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InfrastructureCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InfrastructureCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InfrastructureCreate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSiteId

`func (o *InfrastructureCreate) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InfrastructureCreate) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InfrastructureCreate) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetCustomVariables

`func (o *InfrastructureCreate) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *InfrastructureCreate) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *InfrastructureCreate) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *InfrastructureCreate) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUserIdOwner

`func (o *InfrastructureCreate) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *InfrastructureCreate) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *InfrastructureCreate) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *InfrastructureCreate) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetDescription

`func (o *InfrastructureCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InfrastructureCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InfrastructureCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InfrastructureCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *InfrastructureCreate) GetMeta() InfrastructureMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InfrastructureCreate) GetMetaOk() (*InfrastructureMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InfrastructureCreate) SetMeta(v InfrastructureMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InfrastructureCreate) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


