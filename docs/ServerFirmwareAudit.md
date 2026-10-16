# ServerFirmwareAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **int64** | The id of the audited server. | 
**Baseline** | Pointer to [**ServerFirmwareAuditBaseline**](ServerFirmwareAuditBaseline.md) | The baseline the audit was computed against. Absent when the server could not be audited. | [optional] 
**Components** | Pointer to [**[]ServerFirmwareAuditComponent**](ServerFirmwareAuditComponent.md) | The audit of every updateable component. Absent when the server could not be audited. | [optional] 
**Error** | Pointer to [**ServerFirmwareAuditError**](ServerFirmwareAuditError.md) | Set when this server could not be audited; the other servers are unaffected. | [optional] 

## Methods

### NewServerFirmwareAudit

`func NewServerFirmwareAudit(serverId int64, ) *ServerFirmwareAudit`

NewServerFirmwareAudit instantiates a new ServerFirmwareAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareAuditWithDefaults

`func NewServerFirmwareAuditWithDefaults() *ServerFirmwareAudit`

NewServerFirmwareAuditWithDefaults instantiates a new ServerFirmwareAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ServerFirmwareAudit) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerFirmwareAudit) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerFirmwareAudit) SetServerId(v int64)`

SetServerId sets ServerId field to given value.


### GetBaseline

`func (o *ServerFirmwareAudit) GetBaseline() ServerFirmwareAuditBaseline`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *ServerFirmwareAudit) GetBaselineOk() (*ServerFirmwareAuditBaseline, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *ServerFirmwareAudit) SetBaseline(v ServerFirmwareAuditBaseline)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *ServerFirmwareAudit) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetComponents

`func (o *ServerFirmwareAudit) GetComponents() []ServerFirmwareAuditComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ServerFirmwareAudit) GetComponentsOk() (*[]ServerFirmwareAuditComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ServerFirmwareAudit) SetComponents(v []ServerFirmwareAuditComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ServerFirmwareAudit) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetError

`func (o *ServerFirmwareAudit) GetError() ServerFirmwareAuditError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ServerFirmwareAudit) GetErrorOk() (*ServerFirmwareAuditError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ServerFirmwareAudit) SetError(v ServerFirmwareAuditError)`

SetError sets Error field to given value.

### HasError

`func (o *ServerFirmwareAudit) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


