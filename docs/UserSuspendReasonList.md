# UserSuspendReasonList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]UserSuspendReason**](UserSuspendReason.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewUserSuspendReasonList

`func NewUserSuspendReasonList(data []UserSuspendReason, ) *UserSuspendReasonList`

NewUserSuspendReasonList instantiates a new UserSuspendReasonList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSuspendReasonListWithDefaults

`func NewUserSuspendReasonListWithDefaults() *UserSuspendReasonList`

NewUserSuspendReasonListWithDefaults instantiates a new UserSuspendReasonList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserSuspendReasonList) GetData() []UserSuspendReason`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserSuspendReasonList) GetDataOk() (*[]UserSuspendReason, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserSuspendReasonList) SetData(v []UserSuspendReason)`

SetData sets Data field to given value.


### GetLinks

`func (o *UserSuspendReasonList) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSuspendReasonList) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSuspendReasonList) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserSuspendReasonList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


