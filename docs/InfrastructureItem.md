# InfrastructureItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureId** | **float32** | Infrastructure ID | 
**InfrastructureLabel** | **string** | Infrastructure label | 
**InfrastructureServiceStatus** | **string** | Infrastructure service status | 
**Tags** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructureItem

`func NewInfrastructureItem(infrastructureId float32, infrastructureLabel string, infrastructureServiceStatus string, ) *InfrastructureItem`

NewInfrastructureItem instantiates a new InfrastructureItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureItemWithDefaults

`func NewInfrastructureItemWithDefaults() *InfrastructureItem`

NewInfrastructureItemWithDefaults instantiates a new InfrastructureItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureId

`func (o *InfrastructureItem) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *InfrastructureItem) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *InfrastructureItem) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInfrastructureLabel

`func (o *InfrastructureItem) GetInfrastructureLabel() string`

GetInfrastructureLabel returns the InfrastructureLabel field if non-nil, zero value otherwise.

### GetInfrastructureLabelOk

`func (o *InfrastructureItem) GetInfrastructureLabelOk() (*string, bool)`

GetInfrastructureLabelOk returns a tuple with the InfrastructureLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureLabel

`func (o *InfrastructureItem) SetInfrastructureLabel(v string)`

SetInfrastructureLabel sets InfrastructureLabel field to given value.


### GetInfrastructureServiceStatus

`func (o *InfrastructureItem) GetInfrastructureServiceStatus() string`

GetInfrastructureServiceStatus returns the InfrastructureServiceStatus field if non-nil, zero value otherwise.

### GetInfrastructureServiceStatusOk

`func (o *InfrastructureItem) GetInfrastructureServiceStatusOk() (*string, bool)`

GetInfrastructureServiceStatusOk returns a tuple with the InfrastructureServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureServiceStatus

`func (o *InfrastructureItem) SetInfrastructureServiceStatus(v string)`

SetInfrastructureServiceStatus sets InfrastructureServiceStatus field to given value.


### GetTags

`func (o *InfrastructureItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InfrastructureItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InfrastructureItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InfrastructureItem) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


