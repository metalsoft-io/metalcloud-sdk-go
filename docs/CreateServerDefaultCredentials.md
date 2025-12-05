# CreateServerDefaultCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **float32** | The site ID of the server default credentials | 
**ServerSerialNumber** | **string** | The serial number of the server | 
**ServerMacAddress** | **string** | The MAC address of the server | 
**DefaultUsername** | **string** | The default username of the server | 
**DefaultRackName** | Pointer to **string** | The default rack name of the server | [optional] 
**DefaultRackPositionLowerUnit** | Pointer to **string** | The default rack position upper unit of the server | [optional] 
**DefaultRackPositionUpperUnit** | Pointer to **string** | The default rack position upper unit of the server | [optional] 
**DefaultInventoryId** | Pointer to **string** | The default inventory ID of the server | [optional] 
**DefaultUuid** | Pointer to **string** | The default UUID of the server | [optional] 
**DefaultRegistrationProfileId** | Pointer to **float32** | The default registration profile ID of the server | [optional] 
**DefaultPassword** | **string** | The default password of the server | 

## Methods

### NewCreateServerDefaultCredentials

`func NewCreateServerDefaultCredentials(siteId float32, serverSerialNumber string, serverMacAddress string, defaultUsername string, defaultPassword string, ) *CreateServerDefaultCredentials`

NewCreateServerDefaultCredentials instantiates a new CreateServerDefaultCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerDefaultCredentialsWithDefaults

`func NewCreateServerDefaultCredentialsWithDefaults() *CreateServerDefaultCredentials`

NewCreateServerDefaultCredentialsWithDefaults instantiates a new CreateServerDefaultCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *CreateServerDefaultCredentials) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateServerDefaultCredentials) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateServerDefaultCredentials) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetServerSerialNumber

`func (o *CreateServerDefaultCredentials) GetServerSerialNumber() string`

GetServerSerialNumber returns the ServerSerialNumber field if non-nil, zero value otherwise.

### GetServerSerialNumberOk

`func (o *CreateServerDefaultCredentials) GetServerSerialNumberOk() (*string, bool)`

GetServerSerialNumberOk returns a tuple with the ServerSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSerialNumber

`func (o *CreateServerDefaultCredentials) SetServerSerialNumber(v string)`

SetServerSerialNumber sets ServerSerialNumber field to given value.


### GetServerMacAddress

`func (o *CreateServerDefaultCredentials) GetServerMacAddress() string`

GetServerMacAddress returns the ServerMacAddress field if non-nil, zero value otherwise.

### GetServerMacAddressOk

`func (o *CreateServerDefaultCredentials) GetServerMacAddressOk() (*string, bool)`

GetServerMacAddressOk returns a tuple with the ServerMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMacAddress

`func (o *CreateServerDefaultCredentials) SetServerMacAddress(v string)`

SetServerMacAddress sets ServerMacAddress field to given value.


### GetDefaultUsername

`func (o *CreateServerDefaultCredentials) GetDefaultUsername() string`

GetDefaultUsername returns the DefaultUsername field if non-nil, zero value otherwise.

### GetDefaultUsernameOk

`func (o *CreateServerDefaultCredentials) GetDefaultUsernameOk() (*string, bool)`

GetDefaultUsernameOk returns a tuple with the DefaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUsername

`func (o *CreateServerDefaultCredentials) SetDefaultUsername(v string)`

SetDefaultUsername sets DefaultUsername field to given value.


### GetDefaultRackName

`func (o *CreateServerDefaultCredentials) GetDefaultRackName() string`

GetDefaultRackName returns the DefaultRackName field if non-nil, zero value otherwise.

### GetDefaultRackNameOk

`func (o *CreateServerDefaultCredentials) GetDefaultRackNameOk() (*string, bool)`

GetDefaultRackNameOk returns a tuple with the DefaultRackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRackName

`func (o *CreateServerDefaultCredentials) SetDefaultRackName(v string)`

SetDefaultRackName sets DefaultRackName field to given value.

### HasDefaultRackName

`func (o *CreateServerDefaultCredentials) HasDefaultRackName() bool`

HasDefaultRackName returns a boolean if a field has been set.

### GetDefaultRackPositionLowerUnit

`func (o *CreateServerDefaultCredentials) GetDefaultRackPositionLowerUnit() string`

GetDefaultRackPositionLowerUnit returns the DefaultRackPositionLowerUnit field if non-nil, zero value otherwise.

### GetDefaultRackPositionLowerUnitOk

`func (o *CreateServerDefaultCredentials) GetDefaultRackPositionLowerUnitOk() (*string, bool)`

GetDefaultRackPositionLowerUnitOk returns a tuple with the DefaultRackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRackPositionLowerUnit

`func (o *CreateServerDefaultCredentials) SetDefaultRackPositionLowerUnit(v string)`

SetDefaultRackPositionLowerUnit sets DefaultRackPositionLowerUnit field to given value.

### HasDefaultRackPositionLowerUnit

`func (o *CreateServerDefaultCredentials) HasDefaultRackPositionLowerUnit() bool`

HasDefaultRackPositionLowerUnit returns a boolean if a field has been set.

### GetDefaultRackPositionUpperUnit

`func (o *CreateServerDefaultCredentials) GetDefaultRackPositionUpperUnit() string`

GetDefaultRackPositionUpperUnit returns the DefaultRackPositionUpperUnit field if non-nil, zero value otherwise.

### GetDefaultRackPositionUpperUnitOk

`func (o *CreateServerDefaultCredentials) GetDefaultRackPositionUpperUnitOk() (*string, bool)`

GetDefaultRackPositionUpperUnitOk returns a tuple with the DefaultRackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRackPositionUpperUnit

`func (o *CreateServerDefaultCredentials) SetDefaultRackPositionUpperUnit(v string)`

SetDefaultRackPositionUpperUnit sets DefaultRackPositionUpperUnit field to given value.

### HasDefaultRackPositionUpperUnit

`func (o *CreateServerDefaultCredentials) HasDefaultRackPositionUpperUnit() bool`

HasDefaultRackPositionUpperUnit returns a boolean if a field has been set.

### GetDefaultInventoryId

`func (o *CreateServerDefaultCredentials) GetDefaultInventoryId() string`

GetDefaultInventoryId returns the DefaultInventoryId field if non-nil, zero value otherwise.

### GetDefaultInventoryIdOk

`func (o *CreateServerDefaultCredentials) GetDefaultInventoryIdOk() (*string, bool)`

GetDefaultInventoryIdOk returns a tuple with the DefaultInventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInventoryId

`func (o *CreateServerDefaultCredentials) SetDefaultInventoryId(v string)`

SetDefaultInventoryId sets DefaultInventoryId field to given value.

### HasDefaultInventoryId

`func (o *CreateServerDefaultCredentials) HasDefaultInventoryId() bool`

HasDefaultInventoryId returns a boolean if a field has been set.

### GetDefaultUuid

`func (o *CreateServerDefaultCredentials) GetDefaultUuid() string`

GetDefaultUuid returns the DefaultUuid field if non-nil, zero value otherwise.

### GetDefaultUuidOk

`func (o *CreateServerDefaultCredentials) GetDefaultUuidOk() (*string, bool)`

GetDefaultUuidOk returns a tuple with the DefaultUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUuid

`func (o *CreateServerDefaultCredentials) SetDefaultUuid(v string)`

SetDefaultUuid sets DefaultUuid field to given value.

### HasDefaultUuid

`func (o *CreateServerDefaultCredentials) HasDefaultUuid() bool`

HasDefaultUuid returns a boolean if a field has been set.

### GetDefaultRegistrationProfileId

`func (o *CreateServerDefaultCredentials) GetDefaultRegistrationProfileId() float32`

GetDefaultRegistrationProfileId returns the DefaultRegistrationProfileId field if non-nil, zero value otherwise.

### GetDefaultRegistrationProfileIdOk

`func (o *CreateServerDefaultCredentials) GetDefaultRegistrationProfileIdOk() (*float32, bool)`

GetDefaultRegistrationProfileIdOk returns a tuple with the DefaultRegistrationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegistrationProfileId

`func (o *CreateServerDefaultCredentials) SetDefaultRegistrationProfileId(v float32)`

SetDefaultRegistrationProfileId sets DefaultRegistrationProfileId field to given value.

### HasDefaultRegistrationProfileId

`func (o *CreateServerDefaultCredentials) HasDefaultRegistrationProfileId() bool`

HasDefaultRegistrationProfileId returns a boolean if a field has been set.

### GetDefaultPassword

`func (o *CreateServerDefaultCredentials) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *CreateServerDefaultCredentials) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *CreateServerDefaultCredentials) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


