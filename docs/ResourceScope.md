# ResourceScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**ResourceScopeKind**](ResourceScopeKind.md) | The kind of resource scope | 
**ResourceId** | **NullableFloat32** | ID of the resource, if applicable | 

## Methods

### NewResourceScope

`func NewResourceScope(kind ResourceScopeKind, resourceId NullableFloat32, ) *ResourceScope`

NewResourceScope instantiates a new ResourceScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceScopeWithDefaults

`func NewResourceScopeWithDefaults() *ResourceScope`

NewResourceScopeWithDefaults instantiates a new ResourceScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ResourceScope) GetKind() ResourceScopeKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceScope) GetKindOk() (*ResourceScopeKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceScope) SetKind(v ResourceScopeKind)`

SetKind sets Kind field to given value.


### GetResourceId

`func (o *ResourceScope) GetResourceId() float32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceScope) GetResourceIdOk() (*float32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceScope) SetResourceId(v float32)`

SetResourceId sets ResourceId field to given value.


### SetResourceIdNil

`func (o *ResourceScope) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *ResourceScope) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


