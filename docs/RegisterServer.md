# RegisterServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **float32** | The site id where the server is located. | 
**ServerUUID** | Pointer to **string** | The UUID of the server. | [optional] 
**SerialNumber** | Pointer to **string** | The Serial Number of the server. | [optional] 
**ManagementAddress** | Pointer to **string** | The Management Address of the server. | [optional] 
**Username** | Pointer to **string** | The username to use. | [optional] 
**BmcMacAddress** | Pointer to **string** | The MAC address of the server. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the server. | [optional] 
**Model** | Pointer to **string** | The model of the server. | [optional] 
**RegistrationProfileId** | Pointer to **float32** | The registration profile id of the server. | [optional] 
**Password** | Pointer to **string** | The password to use. | [optional] 

## Methods

### NewRegisterServer

`func NewRegisterServer(siteId float32, ) *RegisterServer`

NewRegisterServer instantiates a new RegisterServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterServerWithDefaults

`func NewRegisterServerWithDefaults() *RegisterServer`

NewRegisterServerWithDefaults instantiates a new RegisterServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *RegisterServer) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *RegisterServer) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *RegisterServer) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetServerUUID

`func (o *RegisterServer) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *RegisterServer) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *RegisterServer) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *RegisterServer) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetSerialNumber

`func (o *RegisterServer) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *RegisterServer) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *RegisterServer) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *RegisterServer) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetManagementAddress

`func (o *RegisterServer) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *RegisterServer) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *RegisterServer) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *RegisterServer) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetUsername

`func (o *RegisterServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisterServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetBmcMacAddress

`func (o *RegisterServer) GetBmcMacAddress() string`

GetBmcMacAddress returns the BmcMacAddress field if non-nil, zero value otherwise.

### GetBmcMacAddressOk

`func (o *RegisterServer) GetBmcMacAddressOk() (*string, bool)`

GetBmcMacAddressOk returns a tuple with the BmcMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcMacAddress

`func (o *RegisterServer) SetBmcMacAddress(v string)`

SetBmcMacAddress sets BmcMacAddress field to given value.

### HasBmcMacAddress

`func (o *RegisterServer) HasBmcMacAddress() bool`

HasBmcMacAddress returns a boolean if a field has been set.

### GetVendor

`func (o *RegisterServer) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *RegisterServer) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *RegisterServer) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *RegisterServer) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *RegisterServer) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RegisterServer) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RegisterServer) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RegisterServer) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRegistrationProfileId

`func (o *RegisterServer) GetRegistrationProfileId() float32`

GetRegistrationProfileId returns the RegistrationProfileId field if non-nil, zero value otherwise.

### GetRegistrationProfileIdOk

`func (o *RegisterServer) GetRegistrationProfileIdOk() (*float32, bool)`

GetRegistrationProfileIdOk returns a tuple with the RegistrationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationProfileId

`func (o *RegisterServer) SetRegistrationProfileId(v float32)`

SetRegistrationProfileId sets RegistrationProfileId field to given value.

### HasRegistrationProfileId

`func (o *RegisterServer) HasRegistrationProfileId() bool`

HasRegistrationProfileId returns a boolean if a field has been set.

### GetPassword

`func (o *RegisterServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterServer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisterServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


