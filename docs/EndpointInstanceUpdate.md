# EndpointInstanceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The endpoint instance label. | [optional] 
**EndpointId** | Pointer to **int32** | Id of endpoint for this Instance. | [optional] 

## Methods

### NewEndpointInstanceUpdate

`func NewEndpointInstanceUpdate() *EndpointInstanceUpdate`

NewEndpointInstanceUpdate instantiates a new EndpointInstanceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceUpdateWithDefaults

`func NewEndpointInstanceUpdateWithDefaults() *EndpointInstanceUpdate`

NewEndpointInstanceUpdateWithDefaults instantiates a new EndpointInstanceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EndpointInstanceUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstanceUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstanceUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EndpointInstanceUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEndpointId

`func (o *EndpointInstanceUpdate) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *EndpointInstanceUpdate) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *EndpointInstanceUpdate) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *EndpointInstanceUpdate) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


