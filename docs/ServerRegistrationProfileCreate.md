# ServerRegistrationProfileCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the server registration profile | 
**Settings** | [**ServerRegistrationProfileSettings**](ServerRegistrationProfileSettings.md) | Server registration profile settings | 

## Methods

### NewServerRegistrationProfileCreate

`func NewServerRegistrationProfileCreate(name string, settings ServerRegistrationProfileSettings, ) *ServerRegistrationProfileCreate`

NewServerRegistrationProfileCreate instantiates a new ServerRegistrationProfileCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRegistrationProfileCreateWithDefaults

`func NewServerRegistrationProfileCreateWithDefaults() *ServerRegistrationProfileCreate`

NewServerRegistrationProfileCreateWithDefaults instantiates a new ServerRegistrationProfileCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerRegistrationProfileCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerRegistrationProfileCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerRegistrationProfileCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *ServerRegistrationProfileCreate) GetSettings() ServerRegistrationProfileSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ServerRegistrationProfileCreate) GetSettingsOk() (*ServerRegistrationProfileSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ServerRegistrationProfileCreate) SetSettings(v ServerRegistrationProfileSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


