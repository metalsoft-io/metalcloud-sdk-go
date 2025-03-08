# ServerInstanceCreate

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

### NewServerInstanceCreate

`func NewServerInstanceCreate() *ServerInstanceCreate`

NewServerInstanceCreate instantiates a new ServerInstanceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceCreateWithDefaults

`func NewServerInstanceCreateWithDefaults() *ServerInstanceCreate`

NewServerInstanceCreateWithDefaults instantiates a new ServerInstanceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ServerInstanceCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupId

`func (o *ServerInstanceCreate) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerInstanceCreate) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerInstanceCreate) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ServerInstanceCreate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetServerTypeId

`func (o *ServerInstanceCreate) GetServerTypeId() int32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerInstanceCreate) GetServerTypeIdOk() (*int32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerInstanceCreate) SetServerTypeId(v int32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *ServerInstanceCreate) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetTemplateId

`func (o *ServerInstanceCreate) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ServerInstanceCreate) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ServerInstanceCreate) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ServerInstanceCreate) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTags

`func (o *ServerInstanceCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerInstanceCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerInstanceCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerInstanceCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnableAutoPortChannel

`func (o *ServerInstanceCreate) GetEnableAutoPortChannel() bool`

GetEnableAutoPortChannel returns the EnableAutoPortChannel field if non-nil, zero value otherwise.

### GetEnableAutoPortChannelOk

`func (o *ServerInstanceCreate) GetEnableAutoPortChannelOk() (*bool, bool)`

GetEnableAutoPortChannelOk returns a tuple with the EnableAutoPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoPortChannel

`func (o *ServerInstanceCreate) SetEnableAutoPortChannel(v bool)`

SetEnableAutoPortChannel sets EnableAutoPortChannel field to given value.

### HasEnableAutoPortChannel

`func (o *ServerInstanceCreate) HasEnableAutoPortChannel() bool`

HasEnableAutoPortChannel returns a boolean if a field has been set.

### GetNetworkProfiles

`func (o *ServerInstanceCreate) GetNetworkProfiles() []ServerInstanceConfigurationNetworkProfilesInner`

GetNetworkProfiles returns the NetworkProfiles field if non-nil, zero value otherwise.

### GetNetworkProfilesOk

`func (o *ServerInstanceCreate) GetNetworkProfilesOk() (*[]ServerInstanceConfigurationNetworkProfilesInner, bool)`

GetNetworkProfilesOk returns a tuple with the NetworkProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfiles

`func (o *ServerInstanceCreate) SetNetworkProfiles(v []ServerInstanceConfigurationNetworkProfilesInner)`

SetNetworkProfiles sets NetworkProfiles field to given value.

### HasNetworkProfiles

`func (o *ServerInstanceCreate) HasNetworkProfiles() bool`

HasNetworkProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


