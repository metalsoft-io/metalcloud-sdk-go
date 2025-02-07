# ServerInstanceProfileCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The Server profile label. Will be automatically generated if not provided. | [optional] 
**NetworkProfiles** | Pointer to [**[]ServerInstanceConfigurationNetworkProfilesInner**](ServerInstanceConfigurationNetworkProfilesInner.md) | Network profiles mapping for each network in this infrastructure. Changes to this configuration will be duplicated on each vm-instance of this group. | [optional] 
**NetworkInterfaces** | Pointer to [**ServerInstanceProfileNetworkInterfaces**](ServerInstanceProfileNetworkInterfaces.md) |  | [optional] 

## Methods

### NewServerInstanceProfileCreate

`func NewServerInstanceProfileCreate() *ServerInstanceProfileCreate`

NewServerInstanceProfileCreate instantiates a new ServerInstanceProfileCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceProfileCreateWithDefaults

`func NewServerInstanceProfileCreateWithDefaults() *ServerInstanceProfileCreate`

NewServerInstanceProfileCreateWithDefaults instantiates a new ServerInstanceProfileCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ServerInstanceProfileCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceProfileCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceProfileCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceProfileCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNetworkProfiles

`func (o *ServerInstanceProfileCreate) GetNetworkProfiles() []ServerInstanceConfigurationNetworkProfilesInner`

GetNetworkProfiles returns the NetworkProfiles field if non-nil, zero value otherwise.

### GetNetworkProfilesOk

`func (o *ServerInstanceProfileCreate) GetNetworkProfilesOk() (*[]ServerInstanceConfigurationNetworkProfilesInner, bool)`

GetNetworkProfilesOk returns a tuple with the NetworkProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfiles

`func (o *ServerInstanceProfileCreate) SetNetworkProfiles(v []ServerInstanceConfigurationNetworkProfilesInner)`

SetNetworkProfiles sets NetworkProfiles field to given value.

### HasNetworkProfiles

`func (o *ServerInstanceProfileCreate) HasNetworkProfiles() bool`

HasNetworkProfiles returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ServerInstanceProfileCreate) GetNetworkInterfaces() ServerInstanceProfileNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ServerInstanceProfileCreate) GetNetworkInterfacesOk() (*ServerInstanceProfileNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ServerInstanceProfileCreate) SetNetworkInterfaces(v ServerInstanceProfileNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ServerInstanceProfileCreate) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


