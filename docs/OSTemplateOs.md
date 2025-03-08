# OSTemplateOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the operating system that this template will install | 
**Version** | **string** | The version of the operating system that this template will install | 
**Credential** | [**OSTemplateOsCredential**](OSTemplateOsCredential.md) |  | 
**SshPort** | Pointer to **float32** | The initial operating system SSH port. | [optional] 

## Methods

### NewOSTemplateOs

`func NewOSTemplateOs(name string, version string, credential OSTemplateOsCredential, ) *OSTemplateOs`

NewOSTemplateOs instantiates a new OSTemplateOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSTemplateOsWithDefaults

`func NewOSTemplateOsWithDefaults() *OSTemplateOs`

NewOSTemplateOsWithDefaults instantiates a new OSTemplateOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OSTemplateOs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSTemplateOs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSTemplateOs) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *OSTemplateOs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OSTemplateOs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OSTemplateOs) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCredential

`func (o *OSTemplateOs) GetCredential() OSTemplateOsCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OSTemplateOs) GetCredentialOk() (*OSTemplateOsCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OSTemplateOs) SetCredential(v OSTemplateOsCredential)`

SetCredential sets Credential field to given value.


### GetSshPort

`func (o *OSTemplateOs) GetSshPort() float32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *OSTemplateOs) GetSshPortOk() (*float32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *OSTemplateOs) SetSshPort(v float32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *OSTemplateOs) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


