# CreateSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The secret name. | 
**Usage** | Pointer to [**SecretUsageType**](SecretUsageType.md) | Secret usage type. | [optional] 
**Value** | **string** | The secret value. | 

## Methods

### NewCreateSecret

`func NewCreateSecret(name string, value string, ) *CreateSecret`

NewCreateSecret instantiates a new CreateSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretWithDefaults

`func NewCreateSecretWithDefaults() *CreateSecret`

NewCreateSecretWithDefaults instantiates a new CreateSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecret) SetName(v string)`

SetName sets Name field to given value.


### GetUsage

`func (o *CreateSecret) GetUsage() SecretUsageType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CreateSecret) GetUsageOk() (*SecretUsageType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CreateSecret) SetUsage(v SecretUsageType)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CreateSecret) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetValue

`func (o *CreateSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateSecret) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


