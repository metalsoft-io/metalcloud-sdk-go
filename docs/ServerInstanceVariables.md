# ServerInstanceVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Product Instance ID. | 
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**CreatedTimestamp** | **string** | Timestamp of the Product Instance creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Product Instance. | [optional] 
**DnsSubdomainId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int32** | Id of the permanent DNS subdomain for the Product Instance | [optional] 
**InfrastructureId** | **int32** |  | 
**GroupId** | **int32** |  | 
**ServerTypeId** | Pointer to **int32** | The server type ID. | [optional] 
**ServerId** | Pointer to **int32** | The ID of the server assigned to the instance. | [optional] 
**OsTemplateId** | Pointer to **int32** | The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal.  | [optional] 
**InstanceWanMlagId** | Pointer to **int32** |  | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**PreferredServerIds** | Pointer to **[]float32** |  | [optional] 
**CustomStorageProfile** | Pointer to [**ServerInstanceStorageProfile**](ServerInstanceStorageProfile.md) | Custom Storage Profile for the Instance. | [optional] 
**ServiceStatus** | **string** | Current status of the server instance. | 
**IsVmInstance** | **int32** | Flag to indicate if this is a VM instance | 
**VmInstanceId** | Pointer to **int32** | The id of the linked VM instance | [optional] 
**ClusterCustomInfo** | Pointer to [**ServerInstanceClusterCustomInfo**](ServerInstanceClusterCustomInfo.md) |  | [optional] 
**OsInstallError** | Pointer to **string** | Last error message during OS install. | [optional] 
**OsInstallImageUrl** | Pointer to **string** | URL where the OS image is available. | [optional] 
**OsInstallImageBuildError** | Pointer to **string** | Last error message during OS image build. | [optional] 
**OsInstallImageBuildInfo** | Pointer to [**ServerInstanceOsInstallImageBuildInfo**](ServerInstanceOsInstallImageBuildInfo.md) | Build info regarding the OS image. | [optional] 
**OsReinstallRequired** | Pointer to **int32** | OS reinstall is required. | [optional] 
**InitiatorNqn** | Pointer to **string** | NVMe Initiator NQN for the Instance. | [optional] 
**IscsiInitiatorIqn** | Pointer to **string** | iSCSI Initiator IQN for the Instance Interface. | [optional] 
**IscsiInitiatorUsername** | Pointer to **string** | iSCSI Initiator Username for the Instance Interface. | [optional] 
**IscsiInitiatorPasswordEncrypted** | Pointer to **string** | iSCSI Initiator Password for the Instance Interface. | [optional] 
**ControlPanelUrl** | Pointer to **string** | Control panel url for the Instance Interface. | [optional] 
**Config** | Pointer to [**ServerInstanceConfiguration**](ServerInstanceConfiguration.md) |  | [optional] 

## Methods

### NewServerInstanceVariables

`func NewServerInstanceVariables(id int32, revision int32, label string, createdTimestamp string, updatedTimestamp string, infrastructureId int32, groupId int32, serviceStatus string, isVmInstance int32, ) *ServerInstanceVariables`

NewServerInstanceVariables instantiates a new ServerInstanceVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceVariablesWithDefaults

`func NewServerInstanceVariablesWithDefaults() *ServerInstanceVariables`

NewServerInstanceVariablesWithDefaults instantiates a new ServerInstanceVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceVariables) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceVariables) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceVariables) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ServerInstanceVariables) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceVariables) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceVariables) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceVariables) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceVariables) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceVariables) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *ServerInstanceVariables) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerInstanceVariables) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerInstanceVariables) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstanceVariables) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceVariables) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceVariables) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstanceVariables) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceVariables) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceVariables) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceVariables) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *ServerInstanceVariables) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstanceVariables) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstanceVariables) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstanceVariables) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ServerInstanceVariables) GetDnsSubdomainId() int32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ServerInstanceVariables) GetDnsSubdomainIdOk() (*int32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ServerInstanceVariables) SetDnsSubdomainId(v int32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ServerInstanceVariables) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ServerInstanceVariables) GetDnsSubdomainPermanentId() int32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ServerInstanceVariables) GetDnsSubdomainPermanentIdOk() (*int32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ServerInstanceVariables) SetDnsSubdomainPermanentId(v int32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ServerInstanceVariables) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *ServerInstanceVariables) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ServerInstanceVariables) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ServerInstanceVariables) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetGroupId

`func (o *ServerInstanceVariables) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerInstanceVariables) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerInstanceVariables) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetServerTypeId

`func (o *ServerInstanceVariables) GetServerTypeId() int32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerInstanceVariables) GetServerTypeIdOk() (*int32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerInstanceVariables) SetServerTypeId(v int32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *ServerInstanceVariables) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetServerId

`func (o *ServerInstanceVariables) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerInstanceVariables) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerInstanceVariables) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ServerInstanceVariables) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ServerInstanceVariables) GetOsTemplateId() int32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ServerInstanceVariables) GetOsTemplateIdOk() (*int32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ServerInstanceVariables) SetOsTemplateId(v int32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ServerInstanceVariables) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetInstanceWanMlagId

`func (o *ServerInstanceVariables) GetInstanceWanMlagId() int32`

GetInstanceWanMlagId returns the InstanceWanMlagId field if non-nil, zero value otherwise.

### GetInstanceWanMlagIdOk

`func (o *ServerInstanceVariables) GetInstanceWanMlagIdOk() (*int32, bool)`

GetInstanceWanMlagIdOk returns a tuple with the InstanceWanMlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceWanMlagId

`func (o *ServerInstanceVariables) SetInstanceWanMlagId(v int32)`

SetInstanceWanMlagId sets InstanceWanMlagId field to given value.

### HasInstanceWanMlagId

`func (o *ServerInstanceVariables) HasInstanceWanMlagId() bool`

HasInstanceWanMlagId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceVariables) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceVariables) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceVariables) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceVariables) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetPreferredServerIds

`func (o *ServerInstanceVariables) GetPreferredServerIds() []float32`

GetPreferredServerIds returns the PreferredServerIds field if non-nil, zero value otherwise.

### GetPreferredServerIdsOk

`func (o *ServerInstanceVariables) GetPreferredServerIdsOk() (*[]float32, bool)`

GetPreferredServerIdsOk returns a tuple with the PreferredServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredServerIds

`func (o *ServerInstanceVariables) SetPreferredServerIds(v []float32)`

SetPreferredServerIds sets PreferredServerIds field to given value.

### HasPreferredServerIds

`func (o *ServerInstanceVariables) HasPreferredServerIds() bool`

HasPreferredServerIds returns a boolean if a field has been set.

### GetCustomStorageProfile

`func (o *ServerInstanceVariables) GetCustomStorageProfile() ServerInstanceStorageProfile`

GetCustomStorageProfile returns the CustomStorageProfile field if non-nil, zero value otherwise.

### GetCustomStorageProfileOk

`func (o *ServerInstanceVariables) GetCustomStorageProfileOk() (*ServerInstanceStorageProfile, bool)`

GetCustomStorageProfileOk returns a tuple with the CustomStorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStorageProfile

`func (o *ServerInstanceVariables) SetCustomStorageProfile(v ServerInstanceStorageProfile)`

SetCustomStorageProfile sets CustomStorageProfile field to given value.

### HasCustomStorageProfile

`func (o *ServerInstanceVariables) HasCustomStorageProfile() bool`

HasCustomStorageProfile returns a boolean if a field has been set.

### GetServiceStatus

`func (o *ServerInstanceVariables) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServerInstanceVariables) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServerInstanceVariables) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetIsVmInstance

`func (o *ServerInstanceVariables) GetIsVmInstance() int32`

GetIsVmInstance returns the IsVmInstance field if non-nil, zero value otherwise.

### GetIsVmInstanceOk

`func (o *ServerInstanceVariables) GetIsVmInstanceOk() (*int32, bool)`

GetIsVmInstanceOk returns a tuple with the IsVmInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVmInstance

`func (o *ServerInstanceVariables) SetIsVmInstance(v int32)`

SetIsVmInstance sets IsVmInstance field to given value.


### GetVmInstanceId

`func (o *ServerInstanceVariables) GetVmInstanceId() int32`

GetVmInstanceId returns the VmInstanceId field if non-nil, zero value otherwise.

### GetVmInstanceIdOk

`func (o *ServerInstanceVariables) GetVmInstanceIdOk() (*int32, bool)`

GetVmInstanceIdOk returns a tuple with the VmInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceId

`func (o *ServerInstanceVariables) SetVmInstanceId(v int32)`

SetVmInstanceId sets VmInstanceId field to given value.

### HasVmInstanceId

`func (o *ServerInstanceVariables) HasVmInstanceId() bool`

HasVmInstanceId returns a boolean if a field has been set.

### GetClusterCustomInfo

`func (o *ServerInstanceVariables) GetClusterCustomInfo() ServerInstanceClusterCustomInfo`

GetClusterCustomInfo returns the ClusterCustomInfo field if non-nil, zero value otherwise.

### GetClusterCustomInfoOk

`func (o *ServerInstanceVariables) GetClusterCustomInfoOk() (*ServerInstanceClusterCustomInfo, bool)`

GetClusterCustomInfoOk returns a tuple with the ClusterCustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCustomInfo

`func (o *ServerInstanceVariables) SetClusterCustomInfo(v ServerInstanceClusterCustomInfo)`

SetClusterCustomInfo sets ClusterCustomInfo field to given value.

### HasClusterCustomInfo

`func (o *ServerInstanceVariables) HasClusterCustomInfo() bool`

HasClusterCustomInfo returns a boolean if a field has been set.

### GetOsInstallError

`func (o *ServerInstanceVariables) GetOsInstallError() string`

GetOsInstallError returns the OsInstallError field if non-nil, zero value otherwise.

### GetOsInstallErrorOk

`func (o *ServerInstanceVariables) GetOsInstallErrorOk() (*string, bool)`

GetOsInstallErrorOk returns a tuple with the OsInstallError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallError

`func (o *ServerInstanceVariables) SetOsInstallError(v string)`

SetOsInstallError sets OsInstallError field to given value.

### HasOsInstallError

`func (o *ServerInstanceVariables) HasOsInstallError() bool`

HasOsInstallError returns a boolean if a field has been set.

### GetOsInstallImageUrl

`func (o *ServerInstanceVariables) GetOsInstallImageUrl() string`

GetOsInstallImageUrl returns the OsInstallImageUrl field if non-nil, zero value otherwise.

### GetOsInstallImageUrlOk

`func (o *ServerInstanceVariables) GetOsInstallImageUrlOk() (*string, bool)`

GetOsInstallImageUrlOk returns a tuple with the OsInstallImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallImageUrl

`func (o *ServerInstanceVariables) SetOsInstallImageUrl(v string)`

SetOsInstallImageUrl sets OsInstallImageUrl field to given value.

### HasOsInstallImageUrl

`func (o *ServerInstanceVariables) HasOsInstallImageUrl() bool`

HasOsInstallImageUrl returns a boolean if a field has been set.

### GetOsInstallImageBuildError

`func (o *ServerInstanceVariables) GetOsInstallImageBuildError() string`

GetOsInstallImageBuildError returns the OsInstallImageBuildError field if non-nil, zero value otherwise.

### GetOsInstallImageBuildErrorOk

`func (o *ServerInstanceVariables) GetOsInstallImageBuildErrorOk() (*string, bool)`

GetOsInstallImageBuildErrorOk returns a tuple with the OsInstallImageBuildError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallImageBuildError

`func (o *ServerInstanceVariables) SetOsInstallImageBuildError(v string)`

SetOsInstallImageBuildError sets OsInstallImageBuildError field to given value.

### HasOsInstallImageBuildError

`func (o *ServerInstanceVariables) HasOsInstallImageBuildError() bool`

HasOsInstallImageBuildError returns a boolean if a field has been set.

### GetOsInstallImageBuildInfo

`func (o *ServerInstanceVariables) GetOsInstallImageBuildInfo() ServerInstanceOsInstallImageBuildInfo`

GetOsInstallImageBuildInfo returns the OsInstallImageBuildInfo field if non-nil, zero value otherwise.

### GetOsInstallImageBuildInfoOk

`func (o *ServerInstanceVariables) GetOsInstallImageBuildInfoOk() (*ServerInstanceOsInstallImageBuildInfo, bool)`

GetOsInstallImageBuildInfoOk returns a tuple with the OsInstallImageBuildInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallImageBuildInfo

`func (o *ServerInstanceVariables) SetOsInstallImageBuildInfo(v ServerInstanceOsInstallImageBuildInfo)`

SetOsInstallImageBuildInfo sets OsInstallImageBuildInfo field to given value.

### HasOsInstallImageBuildInfo

`func (o *ServerInstanceVariables) HasOsInstallImageBuildInfo() bool`

HasOsInstallImageBuildInfo returns a boolean if a field has been set.

### GetOsReinstallRequired

`func (o *ServerInstanceVariables) GetOsReinstallRequired() int32`

GetOsReinstallRequired returns the OsReinstallRequired field if non-nil, zero value otherwise.

### GetOsReinstallRequiredOk

`func (o *ServerInstanceVariables) GetOsReinstallRequiredOk() (*int32, bool)`

GetOsReinstallRequiredOk returns a tuple with the OsReinstallRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsReinstallRequired

`func (o *ServerInstanceVariables) SetOsReinstallRequired(v int32)`

SetOsReinstallRequired sets OsReinstallRequired field to given value.

### HasOsReinstallRequired

`func (o *ServerInstanceVariables) HasOsReinstallRequired() bool`

HasOsReinstallRequired returns a boolean if a field has been set.

### GetInitiatorNqn

`func (o *ServerInstanceVariables) GetInitiatorNqn() string`

GetInitiatorNqn returns the InitiatorNqn field if non-nil, zero value otherwise.

### GetInitiatorNqnOk

`func (o *ServerInstanceVariables) GetInitiatorNqnOk() (*string, bool)`

GetInitiatorNqnOk returns a tuple with the InitiatorNqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorNqn

`func (o *ServerInstanceVariables) SetInitiatorNqn(v string)`

SetInitiatorNqn sets InitiatorNqn field to given value.

### HasInitiatorNqn

`func (o *ServerInstanceVariables) HasInitiatorNqn() bool`

HasInitiatorNqn returns a boolean if a field has been set.

### GetIscsiInitiatorIqn

`func (o *ServerInstanceVariables) GetIscsiInitiatorIqn() string`

GetIscsiInitiatorIqn returns the IscsiInitiatorIqn field if non-nil, zero value otherwise.

### GetIscsiInitiatorIqnOk

`func (o *ServerInstanceVariables) GetIscsiInitiatorIqnOk() (*string, bool)`

GetIscsiInitiatorIqnOk returns a tuple with the IscsiInitiatorIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorIqn

`func (o *ServerInstanceVariables) SetIscsiInitiatorIqn(v string)`

SetIscsiInitiatorIqn sets IscsiInitiatorIqn field to given value.

### HasIscsiInitiatorIqn

`func (o *ServerInstanceVariables) HasIscsiInitiatorIqn() bool`

HasIscsiInitiatorIqn returns a boolean if a field has been set.

### GetIscsiInitiatorUsername

`func (o *ServerInstanceVariables) GetIscsiInitiatorUsername() string`

GetIscsiInitiatorUsername returns the IscsiInitiatorUsername field if non-nil, zero value otherwise.

### GetIscsiInitiatorUsernameOk

`func (o *ServerInstanceVariables) GetIscsiInitiatorUsernameOk() (*string, bool)`

GetIscsiInitiatorUsernameOk returns a tuple with the IscsiInitiatorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorUsername

`func (o *ServerInstanceVariables) SetIscsiInitiatorUsername(v string)`

SetIscsiInitiatorUsername sets IscsiInitiatorUsername field to given value.

### HasIscsiInitiatorUsername

`func (o *ServerInstanceVariables) HasIscsiInitiatorUsername() bool`

HasIscsiInitiatorUsername returns a boolean if a field has been set.

### GetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceVariables) GetIscsiInitiatorPasswordEncrypted() string`

GetIscsiInitiatorPasswordEncrypted returns the IscsiInitiatorPasswordEncrypted field if non-nil, zero value otherwise.

### GetIscsiInitiatorPasswordEncryptedOk

`func (o *ServerInstanceVariables) GetIscsiInitiatorPasswordEncryptedOk() (*string, bool)`

GetIscsiInitiatorPasswordEncryptedOk returns a tuple with the IscsiInitiatorPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceVariables) SetIscsiInitiatorPasswordEncrypted(v string)`

SetIscsiInitiatorPasswordEncrypted sets IscsiInitiatorPasswordEncrypted field to given value.

### HasIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceVariables) HasIscsiInitiatorPasswordEncrypted() bool`

HasIscsiInitiatorPasswordEncrypted returns a boolean if a field has been set.

### GetControlPanelUrl

`func (o *ServerInstanceVariables) GetControlPanelUrl() string`

GetControlPanelUrl returns the ControlPanelUrl field if non-nil, zero value otherwise.

### GetControlPanelUrlOk

`func (o *ServerInstanceVariables) GetControlPanelUrlOk() (*string, bool)`

GetControlPanelUrlOk returns a tuple with the ControlPanelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanelUrl

`func (o *ServerInstanceVariables) SetControlPanelUrl(v string)`

SetControlPanelUrl sets ControlPanelUrl field to given value.

### HasControlPanelUrl

`func (o *ServerInstanceVariables) HasControlPanelUrl() bool`

HasControlPanelUrl returns a boolean if a field has been set.

### GetConfig

`func (o *ServerInstanceVariables) GetConfig() ServerInstanceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerInstanceVariables) GetConfigOk() (*ServerInstanceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerInstanceVariables) SetConfig(v ServerInstanceConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerInstanceVariables) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


