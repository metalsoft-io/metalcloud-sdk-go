# UserSuspend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuspendReason** | **string** | The suspend reason of the user | 
**SuspendReasonPublicComment** | **string** | The public comment for the suspension | 
**SuspendReasonPrivateComment** | Pointer to **string** | The private comment for the suspension | [optional] 

## Methods

### NewUserSuspend

`func NewUserSuspend(suspendReason string, suspendReasonPublicComment string, ) *UserSuspend`

NewUserSuspend instantiates a new UserSuspend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSuspendWithDefaults

`func NewUserSuspendWithDefaults() *UserSuspend`

NewUserSuspendWithDefaults instantiates a new UserSuspend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuspendReason

`func (o *UserSuspend) GetSuspendReason() string`

GetSuspendReason returns the SuspendReason field if non-nil, zero value otherwise.

### GetSuspendReasonOk

`func (o *UserSuspend) GetSuspendReasonOk() (*string, bool)`

GetSuspendReasonOk returns a tuple with the SuspendReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendReason

`func (o *UserSuspend) SetSuspendReason(v string)`

SetSuspendReason sets SuspendReason field to given value.


### GetSuspendReasonPublicComment

`func (o *UserSuspend) GetSuspendReasonPublicComment() string`

GetSuspendReasonPublicComment returns the SuspendReasonPublicComment field if non-nil, zero value otherwise.

### GetSuspendReasonPublicCommentOk

`func (o *UserSuspend) GetSuspendReasonPublicCommentOk() (*string, bool)`

GetSuspendReasonPublicCommentOk returns a tuple with the SuspendReasonPublicComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendReasonPublicComment

`func (o *UserSuspend) SetSuspendReasonPublicComment(v string)`

SetSuspendReasonPublicComment sets SuspendReasonPublicComment field to given value.


### GetSuspendReasonPrivateComment

`func (o *UserSuspend) GetSuspendReasonPrivateComment() string`

GetSuspendReasonPrivateComment returns the SuspendReasonPrivateComment field if non-nil, zero value otherwise.

### GetSuspendReasonPrivateCommentOk

`func (o *UserSuspend) GetSuspendReasonPrivateCommentOk() (*string, bool)`

GetSuspendReasonPrivateCommentOk returns a tuple with the SuspendReasonPrivateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendReasonPrivateComment

`func (o *UserSuspend) SetSuspendReasonPrivateComment(v string)`

SetSuspendReasonPrivateComment sets SuspendReasonPrivateComment field to given value.

### HasSuspendReasonPrivateComment

`func (o *UserSuspend) HasSuspendReasonPrivateComment() bool`

HasSuspendReasonPrivateComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


