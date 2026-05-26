# UpdateDeviceAuthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **int32** | The ID of the site this provider belongs to. | [optional] 
**Label** | Pointer to **string** | The device auth provider label. Must be unique. | [optional] 
**Name** | Pointer to **string** | The device auth provider display name. | [optional] 
**Annotations** | Pointer to **map[string]string** | Key-value annotations for storing additional metadata. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the authentication server. | [optional] 
**Port** | Pointer to **int32** | The TCP port of the authentication server. | [optional] 
**Username** | Pointer to **string** | The username used to authenticate against the provider. | [optional] 
**Password** | Pointer to **NullableString** | The password used to authenticate against the provider. Send null to clear. | [optional] 
**Status** | Pointer to **string** | The status of the device auth provider. | [optional] 

## Methods

### NewUpdateDeviceAuthProvider

`func NewUpdateDeviceAuthProvider() *UpdateDeviceAuthProvider`

NewUpdateDeviceAuthProvider instantiates a new UpdateDeviceAuthProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceAuthProviderWithDefaults

`func NewUpdateDeviceAuthProviderWithDefaults() *UpdateDeviceAuthProvider`

NewUpdateDeviceAuthProviderWithDefaults instantiates a new UpdateDeviceAuthProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *UpdateDeviceAuthProvider) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateDeviceAuthProvider) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateDeviceAuthProvider) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateDeviceAuthProvider) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateDeviceAuthProvider) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateDeviceAuthProvider) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateDeviceAuthProvider) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateDeviceAuthProvider) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateDeviceAuthProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeviceAuthProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeviceAuthProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeviceAuthProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *UpdateDeviceAuthProvider) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateDeviceAuthProvider) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateDeviceAuthProvider) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateDeviceAuthProvider) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetIpAddress

`func (o *UpdateDeviceAuthProvider) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateDeviceAuthProvider) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateDeviceAuthProvider) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateDeviceAuthProvider) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *UpdateDeviceAuthProvider) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateDeviceAuthProvider) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateDeviceAuthProvider) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateDeviceAuthProvider) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateDeviceAuthProvider) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateDeviceAuthProvider) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateDeviceAuthProvider) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateDeviceAuthProvider) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateDeviceAuthProvider) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateDeviceAuthProvider) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateDeviceAuthProvider) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateDeviceAuthProvider) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateDeviceAuthProvider) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateDeviceAuthProvider) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetStatus

`func (o *UpdateDeviceAuthProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDeviceAuthProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDeviceAuthProvider) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateDeviceAuthProvider) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


