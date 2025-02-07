# ServerDefaultCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The ID of the server default credentials | 
**SiteId** | **float32** | The site ID of the server default credentials | 
**ServerSerialNumber** | **string** | The serial number of the server | 
**ServerMacAddress** | **string** | The MAC address of the server | 
**DefaultUsername** | **string** | The default username of the server | 
**DefaultPasswordEncrypted** | **string** | The encrypted default password of the server | 
**DefaultRackName** | Pointer to **string** | The default rack name of the server | [optional] 
**DefaultRackPositionLowerUnit** | Pointer to **string** | The default rack position upper unit of the server | [optional] 
**DefaultRackPositionUpperUnit** | Pointer to **string** | The default rack position upper unit of the server | [optional] 
**DefaultInventoryId** | Pointer to **string** | The default inventory ID of the server | [optional] 
**DefaultUuid** | Pointer to **string** | The default UUID of the server | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerDefaultCredentials

`func NewServerDefaultCredentials(id float32, siteId float32, serverSerialNumber string, serverMacAddress string, defaultUsername string, defaultPasswordEncrypted string, ) *ServerDefaultCredentials`

NewServerDefaultCredentials instantiates a new ServerDefaultCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDefaultCredentialsWithDefaults

`func NewServerDefaultCredentialsWithDefaults() *ServerDefaultCredentials`

NewServerDefaultCredentialsWithDefaults instantiates a new ServerDefaultCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerDefaultCredentials) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDefaultCredentials) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDefaultCredentials) SetId(v float32)`

SetId sets Id field to given value.


### GetSiteId

`func (o *ServerDefaultCredentials) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ServerDefaultCredentials) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ServerDefaultCredentials) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetServerSerialNumber

`func (o *ServerDefaultCredentials) GetServerSerialNumber() string`

GetServerSerialNumber returns the ServerSerialNumber field if non-nil, zero value otherwise.

### GetServerSerialNumberOk

`func (o *ServerDefaultCredentials) GetServerSerialNumberOk() (*string, bool)`

GetServerSerialNumberOk returns a tuple with the ServerSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSerialNumber

`func (o *ServerDefaultCredentials) SetServerSerialNumber(v string)`

SetServerSerialNumber sets ServerSerialNumber field to given value.


### GetServerMacAddress

`func (o *ServerDefaultCredentials) GetServerMacAddress() string`

GetServerMacAddress returns the ServerMacAddress field if non-nil, zero value otherwise.

### GetServerMacAddressOk

`func (o *ServerDefaultCredentials) GetServerMacAddressOk() (*string, bool)`

GetServerMacAddressOk returns a tuple with the ServerMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMacAddress

`func (o *ServerDefaultCredentials) SetServerMacAddress(v string)`

SetServerMacAddress sets ServerMacAddress field to given value.


### GetDefaultUsername

`func (o *ServerDefaultCredentials) GetDefaultUsername() string`

GetDefaultUsername returns the DefaultUsername field if non-nil, zero value otherwise.

### GetDefaultUsernameOk

`func (o *ServerDefaultCredentials) GetDefaultUsernameOk() (*string, bool)`

GetDefaultUsernameOk returns a tuple with the DefaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUsername

`func (o *ServerDefaultCredentials) SetDefaultUsername(v string)`

SetDefaultUsername sets DefaultUsername field to given value.


### GetDefaultPasswordEncrypted

`func (o *ServerDefaultCredentials) GetDefaultPasswordEncrypted() string`

GetDefaultPasswordEncrypted returns the DefaultPasswordEncrypted field if non-nil, zero value otherwise.

### GetDefaultPasswordEncryptedOk

`func (o *ServerDefaultCredentials) GetDefaultPasswordEncryptedOk() (*string, bool)`

GetDefaultPasswordEncryptedOk returns a tuple with the DefaultPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordEncrypted

`func (o *ServerDefaultCredentials) SetDefaultPasswordEncrypted(v string)`

SetDefaultPasswordEncrypted sets DefaultPasswordEncrypted field to given value.


### GetDefaultRackName

`func (o *ServerDefaultCredentials) GetDefaultRackName() string`

GetDefaultRackName returns the DefaultRackName field if non-nil, zero value otherwise.

### GetDefaultRackNameOk

`func (o *ServerDefaultCredentials) GetDefaultRackNameOk() (*string, bool)`

GetDefaultRackNameOk returns a tuple with the DefaultRackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRackName

`func (o *ServerDefaultCredentials) SetDefaultRackName(v string)`

SetDefaultRackName sets DefaultRackName field to given value.

### HasDefaultRackName

`func (o *ServerDefaultCredentials) HasDefaultRackName() bool`

HasDefaultRackName returns a boolean if a field has been set.

### GetDefaultRackPositionLowerUnit

`func (o *ServerDefaultCredentials) GetDefaultRackPositionLowerUnit() string`

GetDefaultRackPositionLowerUnit returns the DefaultRackPositionLowerUnit field if non-nil, zero value otherwise.

### GetDefaultRackPositionLowerUnitOk

`func (o *ServerDefaultCredentials) GetDefaultRackPositionLowerUnitOk() (*string, bool)`

GetDefaultRackPositionLowerUnitOk returns a tuple with the DefaultRackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRackPositionLowerUnit

`func (o *ServerDefaultCredentials) SetDefaultRackPositionLowerUnit(v string)`

SetDefaultRackPositionLowerUnit sets DefaultRackPositionLowerUnit field to given value.

### HasDefaultRackPositionLowerUnit

`func (o *ServerDefaultCredentials) HasDefaultRackPositionLowerUnit() bool`

HasDefaultRackPositionLowerUnit returns a boolean if a field has been set.

### GetDefaultRackPositionUpperUnit

`func (o *ServerDefaultCredentials) GetDefaultRackPositionUpperUnit() string`

GetDefaultRackPositionUpperUnit returns the DefaultRackPositionUpperUnit field if non-nil, zero value otherwise.

### GetDefaultRackPositionUpperUnitOk

`func (o *ServerDefaultCredentials) GetDefaultRackPositionUpperUnitOk() (*string, bool)`

GetDefaultRackPositionUpperUnitOk returns a tuple with the DefaultRackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRackPositionUpperUnit

`func (o *ServerDefaultCredentials) SetDefaultRackPositionUpperUnit(v string)`

SetDefaultRackPositionUpperUnit sets DefaultRackPositionUpperUnit field to given value.

### HasDefaultRackPositionUpperUnit

`func (o *ServerDefaultCredentials) HasDefaultRackPositionUpperUnit() bool`

HasDefaultRackPositionUpperUnit returns a boolean if a field has been set.

### GetDefaultInventoryId

`func (o *ServerDefaultCredentials) GetDefaultInventoryId() string`

GetDefaultInventoryId returns the DefaultInventoryId field if non-nil, zero value otherwise.

### GetDefaultInventoryIdOk

`func (o *ServerDefaultCredentials) GetDefaultInventoryIdOk() (*string, bool)`

GetDefaultInventoryIdOk returns a tuple with the DefaultInventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInventoryId

`func (o *ServerDefaultCredentials) SetDefaultInventoryId(v string)`

SetDefaultInventoryId sets DefaultInventoryId field to given value.

### HasDefaultInventoryId

`func (o *ServerDefaultCredentials) HasDefaultInventoryId() bool`

HasDefaultInventoryId returns a boolean if a field has been set.

### GetDefaultUuid

`func (o *ServerDefaultCredentials) GetDefaultUuid() string`

GetDefaultUuid returns the DefaultUuid field if non-nil, zero value otherwise.

### GetDefaultUuidOk

`func (o *ServerDefaultCredentials) GetDefaultUuidOk() (*string, bool)`

GetDefaultUuidOk returns a tuple with the DefaultUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUuid

`func (o *ServerDefaultCredentials) SetDefaultUuid(v string)`

SetDefaultUuid sets DefaultUuid field to given value.

### HasDefaultUuid

`func (o *ServerDefaultCredentials) HasDefaultUuid() bool`

HasDefaultUuid returns a boolean if a field has been set.

### GetLinks

`func (o *ServerDefaultCredentials) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerDefaultCredentials) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerDefaultCredentials) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerDefaultCredentials) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


