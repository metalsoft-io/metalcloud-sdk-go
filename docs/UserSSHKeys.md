# UserSSHKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The ID of the User SSH key | 
**UserId** | **float32** | The ID of the user | 
**SshKey** | **string** | The SSH key of the user | 
**CreatedTimestamp** | **time.Time** | The timestamp when the User SSH key was created | 
**Status** | **string** | The status of the User SSH key | 

## Methods

### NewUserSSHKeys

`func NewUserSSHKeys(id float32, userId float32, sshKey string, createdTimestamp time.Time, status string, ) *UserSSHKeys`

NewUserSSHKeys instantiates a new UserSSHKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSSHKeysWithDefaults

`func NewUserSSHKeysWithDefaults() *UserSSHKeys`

NewUserSSHKeysWithDefaults instantiates a new UserSSHKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSSHKeys) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSSHKeys) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSSHKeys) SetId(v float32)`

SetId sets Id field to given value.


### GetUserId

`func (o *UserSSHKeys) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSSHKeys) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSSHKeys) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetSshKey

`func (o *UserSSHKeys) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *UserSSHKeys) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *UserSSHKeys) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.


### GetCreatedTimestamp

`func (o *UserSSHKeys) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *UserSSHKeys) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *UserSSHKeys) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetStatus

`func (o *UserSSHKeys) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserSSHKeys) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserSSHKeys) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


