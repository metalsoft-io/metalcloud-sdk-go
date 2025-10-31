# ServerRegistrationProfileUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the server registration profile | [optional] 
**Settings** | Pointer to [**ServerRegistrationProfileUpdateSettings**](ServerRegistrationProfileUpdateSettings.md) | Server registration profile settings | [optional] 

## Methods

### NewServerRegistrationProfileUpdate

`func NewServerRegistrationProfileUpdate() *ServerRegistrationProfileUpdate`

NewServerRegistrationProfileUpdate instantiates a new ServerRegistrationProfileUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRegistrationProfileUpdateWithDefaults

`func NewServerRegistrationProfileUpdateWithDefaults() *ServerRegistrationProfileUpdate`

NewServerRegistrationProfileUpdateWithDefaults instantiates a new ServerRegistrationProfileUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerRegistrationProfileUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerRegistrationProfileUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerRegistrationProfileUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerRegistrationProfileUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *ServerRegistrationProfileUpdate) GetSettings() ServerRegistrationProfileUpdateSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ServerRegistrationProfileUpdate) GetSettingsOk() (*ServerRegistrationProfileUpdateSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ServerRegistrationProfileUpdate) SetSettings(v ServerRegistrationProfileUpdateSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ServerRegistrationProfileUpdate) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


