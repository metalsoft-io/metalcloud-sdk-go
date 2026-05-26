# DeviceAuthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The device auth provider ID | [readonly] 
**SiteId** | **int32** | The ID of the site this provider belongs to. | 
**Label** | **string** | The device auth provider label. Must be unique. | 
**Name** | **string** | The device auth provider display name. | 
**Annotations** | Pointer to **map[string]string** | Key-value annotations for storing additional metadata. | [optional] 
**Kind** | **string** | The authentication protocol kind. | 
**IpAddress** | **string** | The IP address of the authentication server. | 
**Port** | **int32** | The TCP port of the authentication server. Standard TACACS+ port is 49. | 
**Username** | **string** | The username used to authenticate against the provider. | 
**HasSharedSecret** | **bool** | Indicates whether a shared secret is configured. | [readonly] 
**HasPassword** | **bool** | Indicates whether a password is configured. | [readonly] 
**Status** | **string** | The current status of the device auth provider. | 
**Revision** | **int32** | The revision number used for optimistic concurrency control. | [readonly] 
**CreatedBy** | **int32** | The ID of the user who created this provider. | [readonly] 
**UpdatedBy** | Pointer to **int32** | The ID of the user who last updated this provider. | [optional] [readonly] 
**CreatedAt** | **time.Time** | The date and time the provider was created. | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the provider was last updated. | [optional] [readonly] 

## Methods

### NewDeviceAuthProvider

`func NewDeviceAuthProvider(id int32, siteId int32, label string, name string, kind string, ipAddress string, port int32, username string, hasSharedSecret bool, hasPassword bool, status string, revision int32, createdBy int32, createdAt time.Time, ) *DeviceAuthProvider`

NewDeviceAuthProvider instantiates a new DeviceAuthProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthProviderWithDefaults

`func NewDeviceAuthProviderWithDefaults() *DeviceAuthProvider`

NewDeviceAuthProviderWithDefaults instantiates a new DeviceAuthProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceAuthProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthProvider) SetId(v int32)`

SetId sets Id field to given value.


### GetSiteId

`func (o *DeviceAuthProvider) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DeviceAuthProvider) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DeviceAuthProvider) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.


### GetLabel

`func (o *DeviceAuthProvider) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeviceAuthProvider) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeviceAuthProvider) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *DeviceAuthProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAuthProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAuthProvider) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *DeviceAuthProvider) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DeviceAuthProvider) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DeviceAuthProvider) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DeviceAuthProvider) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *DeviceAuthProvider) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeviceAuthProvider) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeviceAuthProvider) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIpAddress

`func (o *DeviceAuthProvider) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DeviceAuthProvider) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DeviceAuthProvider) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPort

`func (o *DeviceAuthProvider) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeviceAuthProvider) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeviceAuthProvider) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *DeviceAuthProvider) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeviceAuthProvider) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeviceAuthProvider) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetHasSharedSecret

`func (o *DeviceAuthProvider) GetHasSharedSecret() bool`

GetHasSharedSecret returns the HasSharedSecret field if non-nil, zero value otherwise.

### GetHasSharedSecretOk

`func (o *DeviceAuthProvider) GetHasSharedSecretOk() (*bool, bool)`

GetHasSharedSecretOk returns a tuple with the HasSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSharedSecret

`func (o *DeviceAuthProvider) SetHasSharedSecret(v bool)`

SetHasSharedSecret sets HasSharedSecret field to given value.


### GetHasPassword

`func (o *DeviceAuthProvider) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *DeviceAuthProvider) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *DeviceAuthProvider) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.


### GetStatus

`func (o *DeviceAuthProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceAuthProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceAuthProvider) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRevision

`func (o *DeviceAuthProvider) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DeviceAuthProvider) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DeviceAuthProvider) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetCreatedBy

`func (o *DeviceAuthProvider) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceAuthProvider) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceAuthProvider) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *DeviceAuthProvider) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DeviceAuthProvider) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DeviceAuthProvider) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DeviceAuthProvider) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeviceAuthProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceAuthProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceAuthProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeviceAuthProvider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceAuthProvider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceAuthProvider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceAuthProvider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


