# UpdateServerIpmiCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username of the server. | [optional] 
**Password** | Pointer to **string** | The password of the server. | [optional] 
**Host** | Pointer to **string** | The host of the server. | [optional] 
**UpdateOnServer** | **bool** | Flag to indicate if the credentials should be updated on the server as well. | 

## Methods

### NewUpdateServerIpmiCredentials

`func NewUpdateServerIpmiCredentials(updateOnServer bool, ) *UpdateServerIpmiCredentials`

NewUpdateServerIpmiCredentials instantiates a new UpdateServerIpmiCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerIpmiCredentialsWithDefaults

`func NewUpdateServerIpmiCredentialsWithDefaults() *UpdateServerIpmiCredentials`

NewUpdateServerIpmiCredentialsWithDefaults instantiates a new UpdateServerIpmiCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateServerIpmiCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateServerIpmiCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateServerIpmiCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateServerIpmiCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateServerIpmiCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateServerIpmiCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateServerIpmiCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateServerIpmiCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHost

`func (o *UpdateServerIpmiCredentials) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateServerIpmiCredentials) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateServerIpmiCredentials) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateServerIpmiCredentials) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetUpdateOnServer

`func (o *UpdateServerIpmiCredentials) GetUpdateOnServer() bool`

GetUpdateOnServer returns the UpdateOnServer field if non-nil, zero value otherwise.

### GetUpdateOnServerOk

`func (o *UpdateServerIpmiCredentials) GetUpdateOnServerOk() (*bool, bool)`

GetUpdateOnServerOk returns a tuple with the UpdateOnServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnServer

`func (o *UpdateServerIpmiCredentials) SetUpdateOnServer(v bool)`

SetUpdateOnServer sets UpdateOnServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


