# CreateDeviceAuthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The device auth provider label. Must be unique. | 
**Name** | **string** | The device auth provider display name. | 
**Annotations** | Pointer to **map[string]string** | Key-value annotations for storing additional metadata. | [optional] 
**SiteId** | **int64** | The ID of the site this provider belongs to. | 
**Kind** | **string** | The authentication protocol kind. | 
**IpAddress** | **string** | The IP address of the authentication server. | 
**Port** | **int32** | The TCP port of the authentication server. Standard TACACS+ port is 49. | 
**SharedSecret** | **string** | The shared secret used to encrypt communication with the TACACS+ server. | 
**Username** | **string** | The username used to authenticate against the provider. | 
**Password** | Pointer to **string** | The password used to authenticate against the provider. | [optional] 
**Status** | Pointer to **string** | The initial status of the device auth provider. | [optional] 

## Methods

### NewCreateDeviceAuthProvider

`func NewCreateDeviceAuthProvider(label string, name string, siteId int64, kind string, ipAddress string, port int32, sharedSecret string, username string, ) *CreateDeviceAuthProvider`

NewCreateDeviceAuthProvider instantiates a new CreateDeviceAuthProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceAuthProviderWithDefaults

`func NewCreateDeviceAuthProviderWithDefaults() *CreateDeviceAuthProvider`

NewCreateDeviceAuthProviderWithDefaults instantiates a new CreateDeviceAuthProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateDeviceAuthProvider) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateDeviceAuthProvider) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateDeviceAuthProvider) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateDeviceAuthProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeviceAuthProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeviceAuthProvider) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *CreateDeviceAuthProvider) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateDeviceAuthProvider) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateDeviceAuthProvider) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateDeviceAuthProvider) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetSiteId

`func (o *CreateDeviceAuthProvider) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateDeviceAuthProvider) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateDeviceAuthProvider) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetKind

`func (o *CreateDeviceAuthProvider) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateDeviceAuthProvider) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateDeviceAuthProvider) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIpAddress

`func (o *CreateDeviceAuthProvider) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateDeviceAuthProvider) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateDeviceAuthProvider) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPort

`func (o *CreateDeviceAuthProvider) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateDeviceAuthProvider) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateDeviceAuthProvider) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSharedSecret

`func (o *CreateDeviceAuthProvider) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *CreateDeviceAuthProvider) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *CreateDeviceAuthProvider) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetUsername

`func (o *CreateDeviceAuthProvider) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDeviceAuthProvider) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDeviceAuthProvider) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateDeviceAuthProvider) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDeviceAuthProvider) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDeviceAuthProvider) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateDeviceAuthProvider) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *CreateDeviceAuthProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDeviceAuthProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDeviceAuthProvider) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateDeviceAuthProvider) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


