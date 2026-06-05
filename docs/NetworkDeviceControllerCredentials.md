# NetworkDeviceControllerCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**SnmpPassword** | Pointer to **string** | The SNMP password of the network device. | [optional] 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Datacenter** | **string** |  | 
**Driver** | **string** |  | 
**Hostname** | **string** |  | 
**ApiUsername** | Pointer to **string** | The username to access the network device API, if different from the management username. | [optional] 
**ApiPassword** | Pointer to **string** | The password to access the network device API, if different from the management password. | [optional] 

## Methods

### NewNetworkDeviceControllerCredentials

`func NewNetworkDeviceControllerCredentials(username string, password string, host string, port int32, datacenter string, driver string, hostname string, ) *NetworkDeviceControllerCredentials`

NewNetworkDeviceControllerCredentials instantiates a new NetworkDeviceControllerCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceControllerCredentialsWithDefaults

`func NewNetworkDeviceControllerCredentialsWithDefaults() *NetworkDeviceControllerCredentials`

NewNetworkDeviceControllerCredentialsWithDefaults instantiates a new NetworkDeviceControllerCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *NetworkDeviceControllerCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NetworkDeviceControllerCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NetworkDeviceControllerCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *NetworkDeviceControllerCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NetworkDeviceControllerCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NetworkDeviceControllerCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSnmpPassword

`func (o *NetworkDeviceControllerCredentials) GetSnmpPassword() string`

GetSnmpPassword returns the SnmpPassword field if non-nil, zero value otherwise.

### GetSnmpPasswordOk

`func (o *NetworkDeviceControllerCredentials) GetSnmpPasswordOk() (*string, bool)`

GetSnmpPasswordOk returns a tuple with the SnmpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpPassword

`func (o *NetworkDeviceControllerCredentials) SetSnmpPassword(v string)`

SetSnmpPassword sets SnmpPassword field to given value.

### HasSnmpPassword

`func (o *NetworkDeviceControllerCredentials) HasSnmpPassword() bool`

HasSnmpPassword returns a boolean if a field has been set.

### GetHost

`func (o *NetworkDeviceControllerCredentials) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworkDeviceControllerCredentials) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworkDeviceControllerCredentials) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *NetworkDeviceControllerCredentials) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworkDeviceControllerCredentials) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworkDeviceControllerCredentials) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDatacenter

`func (o *NetworkDeviceControllerCredentials) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *NetworkDeviceControllerCredentials) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *NetworkDeviceControllerCredentials) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetDriver

`func (o *NetworkDeviceControllerCredentials) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *NetworkDeviceControllerCredentials) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *NetworkDeviceControllerCredentials) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetHostname

`func (o *NetworkDeviceControllerCredentials) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetworkDeviceControllerCredentials) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetworkDeviceControllerCredentials) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetApiUsername

`func (o *NetworkDeviceControllerCredentials) GetApiUsername() string`

GetApiUsername returns the ApiUsername field if non-nil, zero value otherwise.

### GetApiUsernameOk

`func (o *NetworkDeviceControllerCredentials) GetApiUsernameOk() (*string, bool)`

GetApiUsernameOk returns a tuple with the ApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsername

`func (o *NetworkDeviceControllerCredentials) SetApiUsername(v string)`

SetApiUsername sets ApiUsername field to given value.

### HasApiUsername

`func (o *NetworkDeviceControllerCredentials) HasApiUsername() bool`

HasApiUsername returns a boolean if a field has been set.

### GetApiPassword

`func (o *NetworkDeviceControllerCredentials) GetApiPassword() string`

GetApiPassword returns the ApiPassword field if non-nil, zero value otherwise.

### GetApiPasswordOk

`func (o *NetworkDeviceControllerCredentials) GetApiPasswordOk() (*string, bool)`

GetApiPasswordOk returns a tuple with the ApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassword

`func (o *NetworkDeviceControllerCredentials) SetApiPassword(v string)`

SetApiPassword sets ApiPassword field to given value.

### HasApiPassword

`func (o *NetworkDeviceControllerCredentials) HasApiPassword() bool`

HasApiPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


