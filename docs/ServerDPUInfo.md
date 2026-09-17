# ServerDPUInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementAddress** | **string** | The management address of the DPU BMC. Example: 172.18.0.1 | 
**MacAddressOnHostOS** | **string** | The MAC address of the DPU as seen from the Host OS. Example: 00:1A:2B:3C:4D:5E | 
**Username** | **string** | The username to use for the DPU. | 
**ApiUsername** | Pointer to **string** | The username to use for the DPU API if different from the management username. | [optional] 

## Methods

### NewServerDPUInfo

`func NewServerDPUInfo(managementAddress string, macAddressOnHostOS string, username string, ) *ServerDPUInfo`

NewServerDPUInfo instantiates a new ServerDPUInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDPUInfoWithDefaults

`func NewServerDPUInfoWithDefaults() *ServerDPUInfo`

NewServerDPUInfoWithDefaults instantiates a new ServerDPUInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementAddress

`func (o *ServerDPUInfo) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *ServerDPUInfo) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *ServerDPUInfo) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetMacAddressOnHostOS

`func (o *ServerDPUInfo) GetMacAddressOnHostOS() string`

GetMacAddressOnHostOS returns the MacAddressOnHostOS field if non-nil, zero value otherwise.

### GetMacAddressOnHostOSOk

`func (o *ServerDPUInfo) GetMacAddressOnHostOSOk() (*string, bool)`

GetMacAddressOnHostOSOk returns a tuple with the MacAddressOnHostOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressOnHostOS

`func (o *ServerDPUInfo) SetMacAddressOnHostOS(v string)`

SetMacAddressOnHostOS sets MacAddressOnHostOS field to given value.


### GetUsername

`func (o *ServerDPUInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerDPUInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerDPUInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetApiUsername

`func (o *ServerDPUInfo) GetApiUsername() string`

GetApiUsername returns the ApiUsername field if non-nil, zero value otherwise.

### GetApiUsernameOk

`func (o *ServerDPUInfo) GetApiUsernameOk() (*string, bool)`

GetApiUsernameOk returns a tuple with the ApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsername

`func (o *ServerDPUInfo) SetApiUsername(v string)`

SetApiUsername sets ApiUsername field to given value.

### HasApiUsername

`func (o *ServerDPUInfo) HasApiUsername() bool`

HasApiUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


