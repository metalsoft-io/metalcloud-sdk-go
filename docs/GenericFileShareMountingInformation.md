# GenericFileShareMountingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullPath** | **string** | Full path of the File Share | 
**Host** | **string** | Host where the File Share can be accessed | 
**Port** | Pointer to **float32** | Port where the File Share can be accessed | [optional] 
**ServerName** | Pointer to **string** | Server name that handles NFS traffic for the File Share (for certain storages) | [optional] 
**ServerPorts** | Pointer to **[]string** | List of server ports that handle NFS traffic for the File Share (for certain storages) | [optional] 

## Methods

### NewGenericFileShareMountingInformation

`func NewGenericFileShareMountingInformation(fullPath string, host string, ) *GenericFileShareMountingInformation`

NewGenericFileShareMountingInformation instantiates a new GenericFileShareMountingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericFileShareMountingInformationWithDefaults

`func NewGenericFileShareMountingInformationWithDefaults() *GenericFileShareMountingInformation`

NewGenericFileShareMountingInformationWithDefaults instantiates a new GenericFileShareMountingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullPath

`func (o *GenericFileShareMountingInformation) GetFullPath() string`

GetFullPath returns the FullPath field if non-nil, zero value otherwise.

### GetFullPathOk

`func (o *GenericFileShareMountingInformation) GetFullPathOk() (*string, bool)`

GetFullPathOk returns a tuple with the FullPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPath

`func (o *GenericFileShareMountingInformation) SetFullPath(v string)`

SetFullPath sets FullPath field to given value.


### GetHost

`func (o *GenericFileShareMountingInformation) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GenericFileShareMountingInformation) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GenericFileShareMountingInformation) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *GenericFileShareMountingInformation) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GenericFileShareMountingInformation) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GenericFileShareMountingInformation) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GenericFileShareMountingInformation) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServerName

`func (o *GenericFileShareMountingInformation) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *GenericFileShareMountingInformation) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *GenericFileShareMountingInformation) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *GenericFileShareMountingInformation) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerPorts

`func (o *GenericFileShareMountingInformation) GetServerPorts() []string`

GetServerPorts returns the ServerPorts field if non-nil, zero value otherwise.

### GetServerPortsOk

`func (o *GenericFileShareMountingInformation) GetServerPortsOk() (*[]string, bool)`

GetServerPortsOk returns a tuple with the ServerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPorts

`func (o *GenericFileShareMountingInformation) SetServerPorts(v []string)`

SetServerPorts sets ServerPorts field to given value.

### HasServerPorts

`func (o *GenericFileShareMountingInformation) HasServerPorts() bool`

HasServerPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


