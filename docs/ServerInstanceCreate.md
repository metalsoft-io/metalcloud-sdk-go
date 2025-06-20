# ServerInstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The Product Instance label. Will be automatically generated if not provided. | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**ServerTypeId** | Pointer to **int32** | The server type ID. | [optional] 
**Hostname** | Pointer to **string** | Custom hostname for the DNS record name. If set, this will be used as part of the DNS record name instead of the default \&quot;instance\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters, hyphens, and underscores. This will only take effect if the property \&quot;provisionInstanceDnsRecords\&quot; is true. It will be automatically suffixed with the server instance ID (e.g., \&quot;-34\&quot;) to ensure the uniqueness of the resulting DNS name. | [optional] 
**OsTemplateId** | Pointer to **int32** | The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal.  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

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

### GetHostname

`func (o *ServerInstanceCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerInstanceCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerInstanceCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerInstanceCreate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ServerInstanceCreate) GetOsTemplateId() int32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ServerInstanceCreate) GetOsTemplateIdOk() (*int32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ServerInstanceCreate) SetOsTemplateId(v int32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ServerInstanceCreate) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


