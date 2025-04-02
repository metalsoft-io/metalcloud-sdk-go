# ServerCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username of the server. | [optional] 
**Password** | Pointer to **string** | The password of the server. | [optional] 
**VncPassword** | Pointer to **NullableString** | The VNC password of the server. | [optional] 
**SnmpPassword** | Pointer to **NullableString** | The SNMP password of the server. | [optional] 

## Methods

### NewServerCredentials

`func NewServerCredentials() *ServerCredentials`

NewServerCredentials instantiates a new ServerCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCredentialsWithDefaults

`func NewServerCredentialsWithDefaults() *ServerCredentials`

NewServerCredentialsWithDefaults instantiates a new ServerCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ServerCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServerCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ServerCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVncPassword

`func (o *ServerCredentials) GetVncPassword() string`

GetVncPassword returns the VncPassword field if non-nil, zero value otherwise.

### GetVncPasswordOk

`func (o *ServerCredentials) GetVncPasswordOk() (*string, bool)`

GetVncPasswordOk returns a tuple with the VncPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPassword

`func (o *ServerCredentials) SetVncPassword(v string)`

SetVncPassword sets VncPassword field to given value.

### HasVncPassword

`func (o *ServerCredentials) HasVncPassword() bool`

HasVncPassword returns a boolean if a field has been set.

### SetVncPasswordNil

`func (o *ServerCredentials) SetVncPasswordNil(b bool)`

 SetVncPasswordNil sets the value for VncPassword to be an explicit nil

### UnsetVncPassword
`func (o *ServerCredentials) UnsetVncPassword()`

UnsetVncPassword ensures that no value is present for VncPassword, not even an explicit nil
### GetSnmpPassword

`func (o *ServerCredentials) GetSnmpPassword() string`

GetSnmpPassword returns the SnmpPassword field if non-nil, zero value otherwise.

### GetSnmpPasswordOk

`func (o *ServerCredentials) GetSnmpPasswordOk() (*string, bool)`

GetSnmpPasswordOk returns a tuple with the SnmpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpPassword

`func (o *ServerCredentials) SetSnmpPassword(v string)`

SetSnmpPassword sets SnmpPassword field to given value.

### HasSnmpPassword

`func (o *ServerCredentials) HasSnmpPassword() bool`

HasSnmpPassword returns a boolean if a field has been set.

### SetSnmpPasswordNil

`func (o *ServerCredentials) SetSnmpPasswordNil(b bool)`

 SetSnmpPasswordNil sets the value for SnmpPassword to be an explicit nil

### UnsetSnmpPassword
`func (o *ServerCredentials) UnsetSnmpPassword()`

UnsetSnmpPassword ensures that no value is present for SnmpPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


