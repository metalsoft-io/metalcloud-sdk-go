# ServerInstanceCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The initial username. Might be empty if password configuration is disabled. | [optional] 
**InitialPassword** | Pointer to **string** | The initial password. Might be empty if password configuration is disabled. | [optional] 
**InitialSshKeys** | Pointer to [**[]SSHKey**](SSHKey.md) | The initial SSH keys. | [optional] 

## Methods

### NewServerInstanceCredentials

`func NewServerInstanceCredentials() *ServerInstanceCredentials`

NewServerInstanceCredentials instantiates a new ServerInstanceCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceCredentialsWithDefaults

`func NewServerInstanceCredentialsWithDefaults() *ServerInstanceCredentials`

NewServerInstanceCredentialsWithDefaults instantiates a new ServerInstanceCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ServerInstanceCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerInstanceCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerInstanceCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServerInstanceCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetInitialPassword

`func (o *ServerInstanceCredentials) GetInitialPassword() string`

GetInitialPassword returns the InitialPassword field if non-nil, zero value otherwise.

### GetInitialPasswordOk

`func (o *ServerInstanceCredentials) GetInitialPasswordOk() (*string, bool)`

GetInitialPasswordOk returns a tuple with the InitialPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPassword

`func (o *ServerInstanceCredentials) SetInitialPassword(v string)`

SetInitialPassword sets InitialPassword field to given value.

### HasInitialPassword

`func (o *ServerInstanceCredentials) HasInitialPassword() bool`

HasInitialPassword returns a boolean if a field has been set.

### GetInitialSshKeys

`func (o *ServerInstanceCredentials) GetInitialSshKeys() []SSHKey`

GetInitialSshKeys returns the InitialSshKeys field if non-nil, zero value otherwise.

### GetInitialSshKeysOk

`func (o *ServerInstanceCredentials) GetInitialSshKeysOk() (*[]SSHKey, bool)`

GetInitialSshKeysOk returns a tuple with the InitialSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSshKeys

`func (o *ServerInstanceCredentials) SetInitialSshKeys(v []SSHKey)`

SetInitialSshKeys sets InitialSshKeys field to given value.

### HasInitialSshKeys

`func (o *ServerInstanceCredentials) HasInitialSshKeys() bool`

HasInitialSshKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


