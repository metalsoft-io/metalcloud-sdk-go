# ServerInstanceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The server instance label. Will be automatically generated if not provided. | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**ServerTypeId** | Pointer to **int32** | The server type ID. | [optional] 
**TemplateId** | Pointer to **int32** | The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal.  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**EnableAutoPortChannel** | Pointer to **bool** | If enabled will enable port channel to be automatically created. | [optional] [default to true]
**NetworkProfiles** | Pointer to [**[]ServerInstanceConfigurationNetworkProfilesInner**](ServerInstanceConfigurationNetworkProfilesInner.md) | Network profiles mapping for each network in this infrastructure. | [optional] 

## Methods

### NewServerInstanceUpdate

`func NewServerInstanceUpdate() *ServerInstanceUpdate`

NewServerInstanceUpdate instantiates a new ServerInstanceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceUpdateWithDefaults

`func NewServerInstanceUpdateWithDefaults() *ServerInstanceUpdate`

NewServerInstanceUpdateWithDefaults instantiates a new ServerInstanceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ServerInstanceUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupId

`func (o *ServerInstanceUpdate) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerInstanceUpdate) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerInstanceUpdate) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ServerInstanceUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetServerTypeId

`func (o *ServerInstanceUpdate) GetServerTypeId() int32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerInstanceUpdate) GetServerTypeIdOk() (*int32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerInstanceUpdate) SetServerTypeId(v int32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *ServerInstanceUpdate) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetTemplateId

`func (o *ServerInstanceUpdate) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ServerInstanceUpdate) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ServerInstanceUpdate) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ServerInstanceUpdate) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTags

`func (o *ServerInstanceUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerInstanceUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerInstanceUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerInstanceUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnableAutoPortChannel

`func (o *ServerInstanceUpdate) GetEnableAutoPortChannel() bool`

GetEnableAutoPortChannel returns the EnableAutoPortChannel field if non-nil, zero value otherwise.

### GetEnableAutoPortChannelOk

`func (o *ServerInstanceUpdate) GetEnableAutoPortChannelOk() (*bool, bool)`

GetEnableAutoPortChannelOk returns a tuple with the EnableAutoPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoPortChannel

`func (o *ServerInstanceUpdate) SetEnableAutoPortChannel(v bool)`

SetEnableAutoPortChannel sets EnableAutoPortChannel field to given value.

### HasEnableAutoPortChannel

`func (o *ServerInstanceUpdate) HasEnableAutoPortChannel() bool`

HasEnableAutoPortChannel returns a boolean if a field has been set.

### GetNetworkProfiles

`func (o *ServerInstanceUpdate) GetNetworkProfiles() []ServerInstanceConfigurationNetworkProfilesInner`

GetNetworkProfiles returns the NetworkProfiles field if non-nil, zero value otherwise.

### GetNetworkProfilesOk

`func (o *ServerInstanceUpdate) GetNetworkProfilesOk() (*[]ServerInstanceConfigurationNetworkProfilesInner, bool)`

GetNetworkProfilesOk returns a tuple with the NetworkProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfiles

`func (o *ServerInstanceUpdate) SetNetworkProfiles(v []ServerInstanceConfigurationNetworkProfilesInner)`

SetNetworkProfiles sets NetworkProfiles field to given value.

### HasNetworkProfiles

`func (o *ServerInstanceUpdate) HasNetworkProfiles() bool`

HasNetworkProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


