# ServerDPUCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementAddress** | **string** | The management address of the DPU BMC. Example: 172.18.0.1 | 
**MacAddressOnHostOS** | **string** | The MAC address of the DPU as seen from the Host OS. Example: 00:1A:2B:3C:4D:5E | 
**Username** | **string** | The username to use for the DPU. | 
**Password** | **string** | The password to use for the DPU. | 
**ApiUsername** | Pointer to **string** | The username to use for the DPU API if different from the management username. | [optional] 
**ApiPassword** | Pointer to **string** | The password to use for the DPU API if different from the management password. | [optional] 

## Methods

### NewServerDPUCredentials

`func NewServerDPUCredentials(managementAddress string, macAddressOnHostOS string, username string, password string, ) *ServerDPUCredentials`

NewServerDPUCredentials instantiates a new ServerDPUCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDPUCredentialsWithDefaults

`func NewServerDPUCredentialsWithDefaults() *ServerDPUCredentials`

NewServerDPUCredentialsWithDefaults instantiates a new ServerDPUCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementAddress

`func (o *ServerDPUCredentials) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *ServerDPUCredentials) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *ServerDPUCredentials) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetMacAddressOnHostOS

`func (o *ServerDPUCredentials) GetMacAddressOnHostOS() string`

GetMacAddressOnHostOS returns the MacAddressOnHostOS field if non-nil, zero value otherwise.

### GetMacAddressOnHostOSOk

`func (o *ServerDPUCredentials) GetMacAddressOnHostOSOk() (*string, bool)`

GetMacAddressOnHostOSOk returns a tuple with the MacAddressOnHostOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressOnHostOS

`func (o *ServerDPUCredentials) SetMacAddressOnHostOS(v string)`

SetMacAddressOnHostOS sets MacAddressOnHostOS field to given value.


### GetUsername

`func (o *ServerDPUCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerDPUCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerDPUCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ServerDPUCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerDPUCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerDPUCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetApiUsername

`func (o *ServerDPUCredentials) GetApiUsername() string`

GetApiUsername returns the ApiUsername field if non-nil, zero value otherwise.

### GetApiUsernameOk

`func (o *ServerDPUCredentials) GetApiUsernameOk() (*string, bool)`

GetApiUsernameOk returns a tuple with the ApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsername

`func (o *ServerDPUCredentials) SetApiUsername(v string)`

SetApiUsername sets ApiUsername field to given value.

### HasApiUsername

`func (o *ServerDPUCredentials) HasApiUsername() bool`

HasApiUsername returns a boolean if a field has been set.

### GetApiPassword

`func (o *ServerDPUCredentials) GetApiPassword() string`

GetApiPassword returns the ApiPassword field if non-nil, zero value otherwise.

### GetApiPasswordOk

`func (o *ServerDPUCredentials) GetApiPasswordOk() (*string, bool)`

GetApiPasswordOk returns a tuple with the ApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassword

`func (o *ServerDPUCredentials) SetApiPassword(v string)`

SetApiPassword sets ApiPassword field to given value.

### HasApiPassword

`func (o *ServerDPUCredentials) HasApiPassword() bool`

HasApiPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


