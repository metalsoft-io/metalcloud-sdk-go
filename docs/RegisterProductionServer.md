# RegisterProductionServer

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
**Settings** | [**RegisterProductionServerSettings**](RegisterProductionServerSettings.md) | The additional settings for the production server. | 

## Methods

### NewRegisterProductionServer

`func NewRegisterProductionServer(siteId float32, settings RegisterProductionServerSettings, ) *RegisterProductionServer`

NewRegisterProductionServer instantiates a new RegisterProductionServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterProductionServerWithDefaults

`func NewRegisterProductionServerWithDefaults() *RegisterProductionServer`

NewRegisterProductionServerWithDefaults instantiates a new RegisterProductionServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *RegisterProductionServer) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *RegisterProductionServer) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *RegisterProductionServer) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetServerUUID

`func (o *RegisterProductionServer) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *RegisterProductionServer) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *RegisterProductionServer) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *RegisterProductionServer) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetSerialNumber

`func (o *RegisterProductionServer) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *RegisterProductionServer) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *RegisterProductionServer) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *RegisterProductionServer) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetManagementAddress

`func (o *RegisterProductionServer) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *RegisterProductionServer) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *RegisterProductionServer) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *RegisterProductionServer) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetUsername

`func (o *RegisterProductionServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterProductionServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterProductionServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisterProductionServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetBmcMacAddress

`func (o *RegisterProductionServer) GetBmcMacAddress() string`

GetBmcMacAddress returns the BmcMacAddress field if non-nil, zero value otherwise.

### GetBmcMacAddressOk

`func (o *RegisterProductionServer) GetBmcMacAddressOk() (*string, bool)`

GetBmcMacAddressOk returns a tuple with the BmcMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcMacAddress

`func (o *RegisterProductionServer) SetBmcMacAddress(v string)`

SetBmcMacAddress sets BmcMacAddress field to given value.

### HasBmcMacAddress

`func (o *RegisterProductionServer) HasBmcMacAddress() bool`

HasBmcMacAddress returns a boolean if a field has been set.

### GetVendor

`func (o *RegisterProductionServer) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *RegisterProductionServer) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *RegisterProductionServer) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *RegisterProductionServer) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *RegisterProductionServer) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RegisterProductionServer) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RegisterProductionServer) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RegisterProductionServer) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRegistrationProfileId

`func (o *RegisterProductionServer) GetRegistrationProfileId() float32`

GetRegistrationProfileId returns the RegistrationProfileId field if non-nil, zero value otherwise.

### GetRegistrationProfileIdOk

`func (o *RegisterProductionServer) GetRegistrationProfileIdOk() (*float32, bool)`

GetRegistrationProfileIdOk returns a tuple with the RegistrationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationProfileId

`func (o *RegisterProductionServer) SetRegistrationProfileId(v float32)`

SetRegistrationProfileId sets RegistrationProfileId field to given value.

### HasRegistrationProfileId

`func (o *RegisterProductionServer) HasRegistrationProfileId() bool`

HasRegistrationProfileId returns a boolean if a field has been set.

### GetPassword

`func (o *RegisterProductionServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterProductionServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterProductionServer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisterProductionServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSettings

`func (o *RegisterProductionServer) GetSettings() RegisterProductionServerSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *RegisterProductionServer) GetSettingsOk() (*RegisterProductionServerSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *RegisterProductionServer) SetSettings(v RegisterProductionServerSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


