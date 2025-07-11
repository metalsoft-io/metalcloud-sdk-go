# ExtensionTaskSsh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Host to execute the SSH command on. | 
**Port** | **int32** | Port to connect to the host via SSH. | 
**Username** | Pointer to **string** | Username for SSH connection. | [optional] 
**Password** | Pointer to **string** | Password for SSH connection. | [optional] 
**Timeout** | **int32** | Timeout for the SSH command execution in seconds. | 
**CommandTemplate** | **string** | Command template to execute via SSH. | 

## Methods

### NewExtensionTaskSsh

`func NewExtensionTaskSsh(host string, port int32, timeout int32, commandTemplate string, ) *ExtensionTaskSsh`

NewExtensionTaskSsh instantiates a new ExtensionTaskSsh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionTaskSshWithDefaults

`func NewExtensionTaskSshWithDefaults() *ExtensionTaskSsh`

NewExtensionTaskSshWithDefaults instantiates a new ExtensionTaskSsh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ExtensionTaskSsh) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ExtensionTaskSsh) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ExtensionTaskSsh) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ExtensionTaskSsh) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ExtensionTaskSsh) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ExtensionTaskSsh) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *ExtensionTaskSsh) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ExtensionTaskSsh) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ExtensionTaskSsh) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ExtensionTaskSsh) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ExtensionTaskSsh) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ExtensionTaskSsh) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ExtensionTaskSsh) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ExtensionTaskSsh) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTimeout

`func (o *ExtensionTaskSsh) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ExtensionTaskSsh) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ExtensionTaskSsh) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetCommandTemplate

`func (o *ExtensionTaskSsh) GetCommandTemplate() string`

GetCommandTemplate returns the CommandTemplate field if non-nil, zero value otherwise.

### GetCommandTemplateOk

`func (o *ExtensionTaskSsh) GetCommandTemplateOk() (*string, bool)`

GetCommandTemplateOk returns a tuple with the CommandTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandTemplate

`func (o *ExtensionTaskSsh) SetCommandTemplate(v string)`

SetCommandTemplate sets CommandTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


