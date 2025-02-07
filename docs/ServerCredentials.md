# ServerCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datacenter** | **string** | The site where the server is located. | 
**Host** | **string** | The host of the server. | 
**Username** | **string** | The username of the server. | 
**Password** | **string** | The password of the server. | 
**Vendor** | Pointer to **string** | The vendor of the server. | [optional] 
**Model** | Pointer to **string** | The model of the server. | [optional] 
**VncPassword** | Pointer to **string** | The VNC password of the server. | [optional] 
**VncPort** | Pointer to **float32** | The VNC port of the server. | [optional] 
**SnmpPassword** | Pointer to **string** | The SNMP password of the server. | [optional] 

## Methods

### NewServerCredentials

`func NewServerCredentials(datacenter string, host string, username string, password string, ) *ServerCredentials`

NewServerCredentials instantiates a new ServerCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCredentialsWithDefaults

`func NewServerCredentialsWithDefaults() *ServerCredentials`

NewServerCredentialsWithDefaults instantiates a new ServerCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenter

`func (o *ServerCredentials) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ServerCredentials) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ServerCredentials) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetHost

`func (o *ServerCredentials) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ServerCredentials) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ServerCredentials) SetHost(v string)`

SetHost sets Host field to given value.


### GetUsername

`func (o *ServerCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ServerCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetVendor

`func (o *ServerCredentials) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ServerCredentials) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ServerCredentials) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ServerCredentials) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *ServerCredentials) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ServerCredentials) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ServerCredentials) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ServerCredentials) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVncPassword

`func (o *ServerCredentials) GetVncPassword() string`

GetVncPassword returns the VncPassword field if non-nil, zero value otherwise.

### GetVncPasswordOk

`func (o *ServerCredentials) GetVncPasswordOk() (*string, bool)`

GetVncPasswordOk returns a tuple with the VncPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPassword

`func (o *ServerCredentials) SetVncPassword(v string)`

SetVncPassword sets VncPassword field to given value.

### HasVncPassword

`func (o *ServerCredentials) HasVncPassword() bool`

HasVncPassword returns a boolean if a field has been set.

### GetVncPort

`func (o *ServerCredentials) GetVncPort() float32`

GetVncPort returns the VncPort field if non-nil, zero value otherwise.

### GetVncPortOk

`func (o *ServerCredentials) GetVncPortOk() (*float32, bool)`

GetVncPortOk returns a tuple with the VncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPort

`func (o *ServerCredentials) SetVncPort(v float32)`

SetVncPort sets VncPort field to given value.

### HasVncPort

`func (o *ServerCredentials) HasVncPort() bool`

HasVncPort returns a boolean if a field has been set.

### GetSnmpPassword

`func (o *ServerCredentials) GetSnmpPassword() string`

GetSnmpPassword returns the SnmpPassword field if non-nil, zero value otherwise.

### GetSnmpPasswordOk

`func (o *ServerCredentials) GetSnmpPasswordOk() (*string, bool)`

GetSnmpPasswordOk returns a tuple with the SnmpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpPassword

`func (o *ServerCredentials) SetSnmpPassword(v string)`

SetSnmpPassword sets SnmpPassword field to given value.

### HasSnmpPassword

`func (o *ServerCredentials) HasSnmpPassword() bool`

HasSnmpPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


