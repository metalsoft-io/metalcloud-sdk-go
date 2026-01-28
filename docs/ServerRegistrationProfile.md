# ServerRegistrationProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Id** | **float32** | Unique identifier for the server registration profile | 
**Name** | **string** | Name of the server registration profile | 
**Settings** | [**ServerRegistrationProfileSettings**](ServerRegistrationProfileSettings.md) | Server registration profile settings | 
**IsDefault** | **bool** | Whether this is the default server registration profile | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerRegistrationProfile

`func NewServerRegistrationProfile(createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id float32, name string, settings ServerRegistrationProfileSettings, isDefault bool, ) *ServerRegistrationProfile`

NewServerRegistrationProfile instantiates a new ServerRegistrationProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRegistrationProfileWithDefaults

`func NewServerRegistrationProfileWithDefaults() *ServerRegistrationProfile`

NewServerRegistrationProfileWithDefaults instantiates a new ServerRegistrationProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *ServerRegistrationProfile) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerRegistrationProfile) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerRegistrationProfile) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerRegistrationProfile) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerRegistrationProfile) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerRegistrationProfile) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *ServerRegistrationProfile) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerRegistrationProfile) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerRegistrationProfile) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetId

`func (o *ServerRegistrationProfile) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerRegistrationProfile) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerRegistrationProfile) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *ServerRegistrationProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerRegistrationProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerRegistrationProfile) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *ServerRegistrationProfile) GetSettings() ServerRegistrationProfileSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ServerRegistrationProfile) GetSettingsOk() (*ServerRegistrationProfileSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ServerRegistrationProfile) SetSettings(v ServerRegistrationProfileSettings)`

SetSettings sets Settings field to given value.


### GetIsDefault

`func (o *ServerRegistrationProfile) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ServerRegistrationProfile) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ServerRegistrationProfile) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetLinks

`func (o *ServerRegistrationProfile) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerRegistrationProfile) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerRegistrationProfile) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerRegistrationProfile) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


