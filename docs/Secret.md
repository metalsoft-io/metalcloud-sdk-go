# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The secret ID. | 
**UserIdOwner** | **float32** | ID of owner user. | 
**Name** | **string** | The secret name. | 
**Usage** | Pointer to [**VariableUsageType**](VariableUsageType.md) | Secret usage type. | [optional] 
**CreatedTimestamp** | **string** | Timestamp of creation. | 
**UpdatedTimestamp** | **string** | Timestamp of last update. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewSecret

`func NewSecret(id int32, userIdOwner float32, name string, createdTimestamp string, updatedTimestamp string, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Secret) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Secret) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Secret) SetId(v int32)`

SetId sets Id field to given value.


### GetUserIdOwner

`func (o *Secret) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *Secret) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *Secret) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.


### GetName

`func (o *Secret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secret) SetName(v string)`

SetName sets Name field to given value.


### GetUsage

`func (o *Secret) GetUsage() VariableUsageType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Secret) GetUsageOk() (*VariableUsageType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Secret) SetUsage(v VariableUsageType)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Secret) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Secret) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Secret) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Secret) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *Secret) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Secret) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Secret) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *Secret) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Secret) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Secret) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Secret) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


