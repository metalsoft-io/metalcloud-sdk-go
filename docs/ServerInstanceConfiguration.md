# ServerInstanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**GroupId** | **int32** |  | 
**ServerTypeId** | Pointer to **int32** | The server type ID. | [optional] 
**ServerId** | Pointer to **int32** | The ID of the server assigned to the instance. | [optional] 
**Hostname** | Pointer to **string** | Custom hostname for the DNS record name. If set, this will be used as part of the DNS record name instead of the default \&quot;instance\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters, hyphens, and underscores. This will only take effect if the property \&quot;provisionInstanceDnsRecords\&quot; is true. It will be automatically suffixed with the server instance ID (e.g., \&quot;-34\&quot;) to ensure the uniqueness of the resulting DNS name. | [optional] 
**OsTemplateId** | Pointer to **int32** | The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal.  | [optional] 
**InstanceWanMlagId** | Pointer to **int32** |  | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomStorageProfile** | Pointer to [**ServerInstanceStorageProfile**](ServerInstanceStorageProfile.md) | Custom Storage Profile for the Instance. | [optional] 
**InitiatorNqn** | Pointer to **string** | NVMe Initiator NQN for the Instance. | [optional] 
**IscsiInitiatorIqn** | Pointer to **string** | iSCSI Initiator IQN for the Instance Interface. | [optional] 
**IscsiInitiatorUsername** | Pointer to **string** | iSCSI Initiator Username for the Instance Interface. | [optional] 
**IscsiInitiatorPasswordEncrypted** | Pointer to **string** | iSCSI Initiator Password for the Instance Interface. | [optional] 
**ControlPanelUrl** | Pointer to **string** | Control panel url for the Instance Interface. | [optional] 
**DnsSubdomainChangeId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**InfrastructureDeployId** | Pointer to **int32** | Id of the deployment for the Product Instance | [optional] 
**EmptyEdit** | Pointer to **int32** | Number of empty edits | [optional] 
**DeployType** | **string** | Product Instance deploy type | [default to "create"]
**DeployStatus** | **string** | Product Instance deploy status | [default to "not_started"]

## Methods

### NewServerInstanceConfiguration

`func NewServerInstanceConfiguration(revision int32, label string, updatedTimestamp string, groupId int32, deployType string, deployStatus string, ) *ServerInstanceConfiguration`

NewServerInstanceConfiguration instantiates a new ServerInstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceConfigurationWithDefaults

`func NewServerInstanceConfigurationWithDefaults() *ServerInstanceConfiguration`

NewServerInstanceConfigurationWithDefaults instantiates a new ServerInstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ServerInstanceConfiguration) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceConfiguration) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceConfiguration) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstanceConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstanceConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetGroupId

`func (o *ServerInstanceConfiguration) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerInstanceConfiguration) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerInstanceConfiguration) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetServerTypeId

`func (o *ServerInstanceConfiguration) GetServerTypeId() int32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerInstanceConfiguration) GetServerTypeIdOk() (*int32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerInstanceConfiguration) SetServerTypeId(v int32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *ServerInstanceConfiguration) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetServerId

`func (o *ServerInstanceConfiguration) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerInstanceConfiguration) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerInstanceConfiguration) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ServerInstanceConfiguration) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetHostname

`func (o *ServerInstanceConfiguration) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerInstanceConfiguration) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerInstanceConfiguration) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerInstanceConfiguration) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ServerInstanceConfiguration) GetOsTemplateId() int32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ServerInstanceConfiguration) GetOsTemplateIdOk() (*int32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ServerInstanceConfiguration) SetOsTemplateId(v int32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ServerInstanceConfiguration) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetInstanceWanMlagId

`func (o *ServerInstanceConfiguration) GetInstanceWanMlagId() int32`

GetInstanceWanMlagId returns the InstanceWanMlagId field if non-nil, zero value otherwise.

### GetInstanceWanMlagIdOk

`func (o *ServerInstanceConfiguration) GetInstanceWanMlagIdOk() (*int32, bool)`

GetInstanceWanMlagIdOk returns a tuple with the InstanceWanMlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceWanMlagId

`func (o *ServerInstanceConfiguration) SetInstanceWanMlagId(v int32)`

SetInstanceWanMlagId sets InstanceWanMlagId field to given value.

### HasInstanceWanMlagId

`func (o *ServerInstanceConfiguration) HasInstanceWanMlagId() bool`

HasInstanceWanMlagId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceConfiguration) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceConfiguration) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceConfiguration) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceConfiguration) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetCustomStorageProfile

`func (o *ServerInstanceConfiguration) GetCustomStorageProfile() ServerInstanceStorageProfile`

GetCustomStorageProfile returns the CustomStorageProfile field if non-nil, zero value otherwise.

### GetCustomStorageProfileOk

`func (o *ServerInstanceConfiguration) GetCustomStorageProfileOk() (*ServerInstanceStorageProfile, bool)`

GetCustomStorageProfileOk returns a tuple with the CustomStorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStorageProfile

`func (o *ServerInstanceConfiguration) SetCustomStorageProfile(v ServerInstanceStorageProfile)`

SetCustomStorageProfile sets CustomStorageProfile field to given value.

### HasCustomStorageProfile

`func (o *ServerInstanceConfiguration) HasCustomStorageProfile() bool`

HasCustomStorageProfile returns a boolean if a field has been set.

### GetInitiatorNqn

`func (o *ServerInstanceConfiguration) GetInitiatorNqn() string`

GetInitiatorNqn returns the InitiatorNqn field if non-nil, zero value otherwise.

### GetInitiatorNqnOk

`func (o *ServerInstanceConfiguration) GetInitiatorNqnOk() (*string, bool)`

GetInitiatorNqnOk returns a tuple with the InitiatorNqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorNqn

`func (o *ServerInstanceConfiguration) SetInitiatorNqn(v string)`

SetInitiatorNqn sets InitiatorNqn field to given value.

### HasInitiatorNqn

`func (o *ServerInstanceConfiguration) HasInitiatorNqn() bool`

HasInitiatorNqn returns a boolean if a field has been set.

### GetIscsiInitiatorIqn

`func (o *ServerInstanceConfiguration) GetIscsiInitiatorIqn() string`

GetIscsiInitiatorIqn returns the IscsiInitiatorIqn field if non-nil, zero value otherwise.

### GetIscsiInitiatorIqnOk

`func (o *ServerInstanceConfiguration) GetIscsiInitiatorIqnOk() (*string, bool)`

GetIscsiInitiatorIqnOk returns a tuple with the IscsiInitiatorIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorIqn

`func (o *ServerInstanceConfiguration) SetIscsiInitiatorIqn(v string)`

SetIscsiInitiatorIqn sets IscsiInitiatorIqn field to given value.

### HasIscsiInitiatorIqn

`func (o *ServerInstanceConfiguration) HasIscsiInitiatorIqn() bool`

HasIscsiInitiatorIqn returns a boolean if a field has been set.

### GetIscsiInitiatorUsername

`func (o *ServerInstanceConfiguration) GetIscsiInitiatorUsername() string`

GetIscsiInitiatorUsername returns the IscsiInitiatorUsername field if non-nil, zero value otherwise.

### GetIscsiInitiatorUsernameOk

`func (o *ServerInstanceConfiguration) GetIscsiInitiatorUsernameOk() (*string, bool)`

GetIscsiInitiatorUsernameOk returns a tuple with the IscsiInitiatorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorUsername

`func (o *ServerInstanceConfiguration) SetIscsiInitiatorUsername(v string)`

SetIscsiInitiatorUsername sets IscsiInitiatorUsername field to given value.

### HasIscsiInitiatorUsername

`func (o *ServerInstanceConfiguration) HasIscsiInitiatorUsername() bool`

HasIscsiInitiatorUsername returns a boolean if a field has been set.

### GetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceConfiguration) GetIscsiInitiatorPasswordEncrypted() string`

GetIscsiInitiatorPasswordEncrypted returns the IscsiInitiatorPasswordEncrypted field if non-nil, zero value otherwise.

### GetIscsiInitiatorPasswordEncryptedOk

`func (o *ServerInstanceConfiguration) GetIscsiInitiatorPasswordEncryptedOk() (*string, bool)`

GetIscsiInitiatorPasswordEncryptedOk returns a tuple with the IscsiInitiatorPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceConfiguration) SetIscsiInitiatorPasswordEncrypted(v string)`

SetIscsiInitiatorPasswordEncrypted sets IscsiInitiatorPasswordEncrypted field to given value.

### HasIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceConfiguration) HasIscsiInitiatorPasswordEncrypted() bool`

HasIscsiInitiatorPasswordEncrypted returns a boolean if a field has been set.

### GetControlPanelUrl

`func (o *ServerInstanceConfiguration) GetControlPanelUrl() string`

GetControlPanelUrl returns the ControlPanelUrl field if non-nil, zero value otherwise.

### GetControlPanelUrlOk

`func (o *ServerInstanceConfiguration) GetControlPanelUrlOk() (*string, bool)`

GetControlPanelUrlOk returns a tuple with the ControlPanelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanelUrl

`func (o *ServerInstanceConfiguration) SetControlPanelUrl(v string)`

SetControlPanelUrl sets ControlPanelUrl field to given value.

### HasControlPanelUrl

`func (o *ServerInstanceConfiguration) HasControlPanelUrl() bool`

HasControlPanelUrl returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *ServerInstanceConfiguration) GetDnsSubdomainChangeId() int32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *ServerInstanceConfiguration) GetDnsSubdomainChangeIdOk() (*int32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *ServerInstanceConfiguration) SetDnsSubdomainChangeId(v int32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *ServerInstanceConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInfrastructureDeployId

`func (o *ServerInstanceConfiguration) GetInfrastructureDeployId() int32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *ServerInstanceConfiguration) GetInfrastructureDeployIdOk() (*int32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *ServerInstanceConfiguration) SetInfrastructureDeployId(v int32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *ServerInstanceConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetEmptyEdit

`func (o *ServerInstanceConfiguration) GetEmptyEdit() int32`

GetEmptyEdit returns the EmptyEdit field if non-nil, zero value otherwise.

### GetEmptyEditOk

`func (o *ServerInstanceConfiguration) GetEmptyEditOk() (*int32, bool)`

GetEmptyEditOk returns a tuple with the EmptyEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyEdit

`func (o *ServerInstanceConfiguration) SetEmptyEdit(v int32)`

SetEmptyEdit sets EmptyEdit field to given value.

### HasEmptyEdit

`func (o *ServerInstanceConfiguration) HasEmptyEdit() bool`

HasEmptyEdit returns a boolean if a field has been set.

### GetDeployType

`func (o *ServerInstanceConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ServerInstanceConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ServerInstanceConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ServerInstanceConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ServerInstanceConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ServerInstanceConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


