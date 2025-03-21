# UserSuspendReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The ID of the user suspend reasons | 
**UserId** | **float32** | The ID of the user | 
**Type** | **string** | The type of the user suspend reason | 
**CreatedTimestamp** | **time.Time** | The timestamp when the user logged in | 
**EndTimestamp** | Pointer to **time.Time** | The timestamp when the user was last updated | [optional] 
**PublicComment** | **string** | The public comment of the user suspend reason | 
**PrivateComment** | Pointer to **string** | The private comment of the user suspend reason | [optional] 

## Methods

### NewUserSuspendReason

`func NewUserSuspendReason(id float32, userId float32, type_ string, createdTimestamp time.Time, publicComment string, ) *UserSuspendReason`

NewUserSuspendReason instantiates a new UserSuspendReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSuspendReasonWithDefaults

`func NewUserSuspendReasonWithDefaults() *UserSuspendReason`

NewUserSuspendReasonWithDefaults instantiates a new UserSuspendReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSuspendReason) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSuspendReason) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSuspendReason) SetId(v float32)`

SetId sets Id field to given value.


### GetUserId

`func (o *UserSuspendReason) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSuspendReason) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSuspendReason) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetType

`func (o *UserSuspendReason) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSuspendReason) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSuspendReason) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedTimestamp

`func (o *UserSuspendReason) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *UserSuspendReason) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *UserSuspendReason) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetEndTimestamp

`func (o *UserSuspendReason) GetEndTimestamp() time.Time`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *UserSuspendReason) GetEndTimestampOk() (*time.Time, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *UserSuspendReason) SetEndTimestamp(v time.Time)`

SetEndTimestamp sets EndTimestamp field to given value.

### HasEndTimestamp

`func (o *UserSuspendReason) HasEndTimestamp() bool`

HasEndTimestamp returns a boolean if a field has been set.

### GetPublicComment

`func (o *UserSuspendReason) GetPublicComment() string`

GetPublicComment returns the PublicComment field if non-nil, zero value otherwise.

### GetPublicCommentOk

`func (o *UserSuspendReason) GetPublicCommentOk() (*string, bool)`

GetPublicCommentOk returns a tuple with the PublicComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicComment

`func (o *UserSuspendReason) SetPublicComment(v string)`

SetPublicComment sets PublicComment field to given value.


### GetPrivateComment

`func (o *UserSuspendReason) GetPrivateComment() string`

GetPrivateComment returns the PrivateComment field if non-nil, zero value otherwise.

### GetPrivateCommentOk

`func (o *UserSuspendReason) GetPrivateCommentOk() (*string, bool)`

GetPrivateCommentOk returns a tuple with the PrivateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateComment

`func (o *UserSuspendReason) SetPrivateComment(v string)`

SetPrivateComment sets PrivateComment field to given value.

### HasPrivateComment

`func (o *UserSuspendReason) HasPrivateComment() bool`

HasPrivateComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


