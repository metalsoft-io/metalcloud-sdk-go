# ServerInstanceCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The initial username. Might be empty if password configuration is disabled. | [optional] 
**InitialPassword** | Pointer to **string** | The initial password. Might be empty if password configuration is disabled. | [optional] 
**PublicSshKey** | Pointer to **string** | The SSH key public key. | [optional] 
**IscsiInitiatorUsername** | Pointer to **string** | The iSCSI initiator username. | [optional] 
**IscsiInitiatorPassword** | Pointer to **string** | The iSCSI initiator password. | [optional] 
**Config** | Pointer to **map[string]interface{}** | The credentials that will be used starting with next deploy | [optional] 

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

### GetPublicSshKey

`func (o *ServerInstanceCredentials) GetPublicSshKey() string`

GetPublicSshKey returns the PublicSshKey field if non-nil, zero value otherwise.

### GetPublicSshKeyOk

`func (o *ServerInstanceCredentials) GetPublicSshKeyOk() (*string, bool)`

GetPublicSshKeyOk returns a tuple with the PublicSshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSshKey

`func (o *ServerInstanceCredentials) SetPublicSshKey(v string)`

SetPublicSshKey sets PublicSshKey field to given value.

### HasPublicSshKey

`func (o *ServerInstanceCredentials) HasPublicSshKey() bool`

HasPublicSshKey returns a boolean if a field has been set.

### GetIscsiInitiatorUsername

`func (o *ServerInstanceCredentials) GetIscsiInitiatorUsername() string`

GetIscsiInitiatorUsername returns the IscsiInitiatorUsername field if non-nil, zero value otherwise.

### GetIscsiInitiatorUsernameOk

`func (o *ServerInstanceCredentials) GetIscsiInitiatorUsernameOk() (*string, bool)`

GetIscsiInitiatorUsernameOk returns a tuple with the IscsiInitiatorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorUsername

`func (o *ServerInstanceCredentials) SetIscsiInitiatorUsername(v string)`

SetIscsiInitiatorUsername sets IscsiInitiatorUsername field to given value.

### HasIscsiInitiatorUsername

`func (o *ServerInstanceCredentials) HasIscsiInitiatorUsername() bool`

HasIscsiInitiatorUsername returns a boolean if a field has been set.

### GetIscsiInitiatorPassword

`func (o *ServerInstanceCredentials) GetIscsiInitiatorPassword() string`

GetIscsiInitiatorPassword returns the IscsiInitiatorPassword field if non-nil, zero value otherwise.

### GetIscsiInitiatorPasswordOk

`func (o *ServerInstanceCredentials) GetIscsiInitiatorPasswordOk() (*string, bool)`

GetIscsiInitiatorPasswordOk returns a tuple with the IscsiInitiatorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorPassword

`func (o *ServerInstanceCredentials) SetIscsiInitiatorPassword(v string)`

SetIscsiInitiatorPassword sets IscsiInitiatorPassword field to given value.

### HasIscsiInitiatorPassword

`func (o *ServerInstanceCredentials) HasIscsiInitiatorPassword() bool`

HasIscsiInitiatorPassword returns a boolean if a field has been set.

### GetConfig

`func (o *ServerInstanceCredentials) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerInstanceCredentials) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerInstanceCredentials) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerInstanceCredentials) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


