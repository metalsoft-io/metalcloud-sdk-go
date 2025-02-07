# ServerInstanceGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The server instance group label. Will be automatically generated if not provided. | [optional] 
**DefaultServerProfileID** | Pointer to **int32** | The group&#39;s default server profile. Useful when creating a server instance with a group id set, the profile will be automatically applied. | [optional] 

## Methods

### NewServerInstanceGroupUpdate

`func NewServerInstanceGroupUpdate() *ServerInstanceGroupUpdate`

NewServerInstanceGroupUpdate instantiates a new ServerInstanceGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupUpdateWithDefaults

`func NewServerInstanceGroupUpdateWithDefaults() *ServerInstanceGroupUpdate`

NewServerInstanceGroupUpdateWithDefaults instantiates a new ServerInstanceGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ServerInstanceGroupUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceGroupUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDefaultServerProfileID

`func (o *ServerInstanceGroupUpdate) GetDefaultServerProfileID() int32`

GetDefaultServerProfileID returns the DefaultServerProfileID field if non-nil, zero value otherwise.

### GetDefaultServerProfileIDOk

`func (o *ServerInstanceGroupUpdate) GetDefaultServerProfileIDOk() (*int32, bool)`

GetDefaultServerProfileIDOk returns a tuple with the DefaultServerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerProfileID

`func (o *ServerInstanceGroupUpdate) SetDefaultServerProfileID(v int32)`

SetDefaultServerProfileID sets DefaultServerProfileID field to given value.

### HasDefaultServerProfileID

`func (o *ServerInstanceGroupUpdate) HasDefaultServerProfileID() bool`

HasDefaultServerProfileID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


