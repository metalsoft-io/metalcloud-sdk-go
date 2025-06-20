# EndpointInstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The Product Instance label. Will be automatically generated if not provided. | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**EndpointId** | **int32** | Id of endpoint for this Instance. | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEndpointInstanceCreate

`func NewEndpointInstanceCreate(endpointId int32, ) *EndpointInstanceCreate`

NewEndpointInstanceCreate instantiates a new EndpointInstanceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceCreateWithDefaults

`func NewEndpointInstanceCreateWithDefaults() *EndpointInstanceCreate`

NewEndpointInstanceCreateWithDefaults instantiates a new EndpointInstanceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EndpointInstanceCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstanceCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstanceCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EndpointInstanceCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupId

`func (o *EndpointInstanceCreate) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EndpointInstanceCreate) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EndpointInstanceCreate) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *EndpointInstanceCreate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetEndpointId

`func (o *EndpointInstanceCreate) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *EndpointInstanceCreate) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *EndpointInstanceCreate) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.


### GetTags

`func (o *EndpointInstanceCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EndpointInstanceCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EndpointInstanceCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EndpointInstanceCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


