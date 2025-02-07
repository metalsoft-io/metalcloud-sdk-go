# SSHKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The SSH key ID. | 
**Name** | Pointer to **string** | The SSH key name. | [optional] 
**PublicKey** | **string** | The SSH key public key. | 
**UserId** | Pointer to **int32** | The SSH key user ID. | [optional] 

## Methods

### NewSSHKey

`func NewSSHKey(id int32, publicKey string, ) *SSHKey`

NewSSHKey instantiates a new SSHKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyWithDefaults

`func NewSSHKeyWithDefaults() *SSHKey`

NewSSHKeyWithDefaults instantiates a new SSHKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SSHKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSHKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSHKey) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SSHKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSHKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSHKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SSHKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *SSHKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SSHKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SSHKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetUserId

`func (o *SSHKey) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SSHKey) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SSHKey) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SSHKey) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


