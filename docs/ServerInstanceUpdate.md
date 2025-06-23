# ServerInstanceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The server instance label. | [optional] 
**ServerTypeId** | Pointer to **int32** | The server type ID. | [optional] 
**Hostname** | Pointer to **string** | Custom hostname(subdomain) part of the fully qualified domain name (FQDN). If set, this will be used as the subdomain record part of the DNS record name instead of the default \&quot;instance\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters and hyphens. This will only take effect if the property \&quot;provisionInstanceDnsRecords\&quot; is true.  | [optional] 
**OsTemplateId** | Pointer to **int32** | The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal.  | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomStorageProfile** | Pointer to [**ServerInstanceStorageProfile**](ServerInstanceStorageProfile.md) | Custom Storage Profile for the Instance. | [optional] 

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

### GetHostname

`func (o *ServerInstanceUpdate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerInstanceUpdate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerInstanceUpdate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerInstanceUpdate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ServerInstanceUpdate) GetOsTemplateId() int32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ServerInstanceUpdate) GetOsTemplateIdOk() (*int32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ServerInstanceUpdate) SetOsTemplateId(v int32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ServerInstanceUpdate) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceUpdate) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceUpdate) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceUpdate) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceUpdate) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetCustomStorageProfile

`func (o *ServerInstanceUpdate) GetCustomStorageProfile() ServerInstanceStorageProfile`

GetCustomStorageProfile returns the CustomStorageProfile field if non-nil, zero value otherwise.

### GetCustomStorageProfileOk

`func (o *ServerInstanceUpdate) GetCustomStorageProfileOk() (*ServerInstanceStorageProfile, bool)`

GetCustomStorageProfileOk returns a tuple with the CustomStorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStorageProfile

`func (o *ServerInstanceUpdate) SetCustomStorageProfile(v ServerInstanceStorageProfile)`

SetCustomStorageProfile sets CustomStorageProfile field to given value.

### HasCustomStorageProfile

`func (o *ServerInstanceUpdate) HasCustomStorageProfile() bool`

HasCustomStorageProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


