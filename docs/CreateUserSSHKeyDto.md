# CreateUserSSHKeyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKey** | **string** | The SSH key of the user | 

## Methods

### NewCreateUserSSHKeyDto

`func NewCreateUserSSHKeyDto(sshKey string, ) *CreateUserSSHKeyDto`

NewCreateUserSSHKeyDto instantiates a new CreateUserSSHKeyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserSSHKeyDtoWithDefaults

`func NewCreateUserSSHKeyDtoWithDefaults() *CreateUserSSHKeyDto`

NewCreateUserSSHKeyDtoWithDefaults instantiates a new CreateUserSSHKeyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKey

`func (o *CreateUserSSHKeyDto) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *CreateUserSSHKeyDto) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *CreateUserSSHKeyDto) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


