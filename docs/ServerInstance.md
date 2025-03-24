# ServerInstance

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
**DriveIdBootable** | Pointer to **int32** |  | [optional] 
**ServerTypeId** | Pointer to **int32** | The server type ID. | [optional] 
**ServerId** | Pointer to **int32** | The ID of the server assigned to the instance. | [optional] 
**TemplateId** | Pointer to **int32** | The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal.  | [optional] 
**InstanceWanMlagId** | Pointer to **int32** |  | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**PreferredServerIds** | Pointer to **[]float32** |  | [optional] 
**RaidProfile** | Pointer to [**ServerInstanceStorageProfile**](ServerInstanceStorageProfile.md) | RAID profile for the Instance Interface. | [optional] 
**ServiceStatus** | **string** | Current status of the server instance. | 
**IsVmInstance** | **int32** | Flag to indicate if this is a VM instance | 
**VmInstanceId** | Pointer to **int32** | The id of the linked VM instance | [optional] 
**Meta** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**OperatingSystemInfo** | Pointer to **map[string]interface{}** | Operating system information of the Instance. | [optional] 
**VariablesOsAssetsTemplateTypeAdvanced** | Pointer to **map[string]interface{}** | Advanced variables used when building the OS image. | [optional] 
**SecretsOsAssetsTemplateTypeAdvanced** | Pointer to **map[string]interface{}** | Advanced secrets used when building the OS image. | [optional] 
**VariablesMaterializedForOSAssets** | Pointer to **map[string]interface{}** | Variables used when building the OS image. | [optional] 
**SecretsMaterializedForOSAssets** | Pointer to **map[string]interface{}** | Secrets used when building the OS image. | [optional] 
**ClusterCustomInfo** | Pointer to [**ServerInstanceClusterCustomInfo**](ServerInstanceClusterCustomInfo.md) |  | [optional] 
**OsInstallError** | Pointer to **string** | Last error message during OS install. | [optional] 
**OsInstallImageUrl** | Pointer to **string** | URL where the OS image is available. | [optional] 
**OsInstallImageBuildError** | Pointer to **string** | Last error message during OS image build. | [optional] 
**OsInstallImageBuildInfo** | Pointer to [**ServerInstanceOsInstallImageBuildInfo**](ServerInstanceOsInstallImageBuildInfo.md) | Build info regarding the OS image. | [optional] 
**OsReinstallRequired** | Pointer to **int32** | OS reinstall is required. | [optional] 
**IscsiInitiatorIqn** | Pointer to **string** | iSCSI Initiator IQN for the Instance Interface. | [optional] 
**IscsiInitiatorUsername** | Pointer to **string** | iSCSI Initiator Username for the Instance Interface. | [optional] 
**IscsiInitiatorPasswordEncrypted** | Pointer to **string** | iSCSI Initiator Password for the Instance Interface. | [optional] 
**ControlPanelUrl** | Pointer to **string** | Control panel url for the Instance Interface. | [optional] 
**Config** | Pointer to [**ServerInstanceConfiguration**](ServerInstanceConfiguration.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerInstance

`func NewServerInstance(id int32, revision int32, label string, createdTimestamp string, updatedTimestamp string, infrastructureId int32, groupId int32, serviceStatus string, isVmInstance int32, ) *ServerInstance`

NewServerInstance instantiates a new ServerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceWithDefaults

`func NewServerInstanceWithDefaults() *ServerInstance`

NewServerInstanceWithDefaults instantiates a new ServerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstance) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ServerInstance) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstance) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstance) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *ServerInstance) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerInstance) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerInstance) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstance) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstance) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstance) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstance) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstance) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstance) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstance) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *ServerInstance) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstance) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstance) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstance) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ServerInstance) GetDnsSubdomainId() int32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ServerInstance) GetDnsSubdomainIdOk() (*int32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ServerInstance) SetDnsSubdomainId(v int32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ServerInstance) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ServerInstance) GetDnsSubdomainPermanentId() int32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ServerInstance) GetDnsSubdomainPermanentIdOk() (*int32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ServerInstance) SetDnsSubdomainPermanentId(v int32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ServerInstance) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *ServerInstance) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ServerInstance) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ServerInstance) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetGroupId

`func (o *ServerInstance) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerInstance) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerInstance) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetDriveIdBootable

`func (o *ServerInstance) GetDriveIdBootable() int32`

GetDriveIdBootable returns the DriveIdBootable field if non-nil, zero value otherwise.

### GetDriveIdBootableOk

`func (o *ServerInstance) GetDriveIdBootableOk() (*int32, bool)`

GetDriveIdBootableOk returns a tuple with the DriveIdBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveIdBootable

`func (o *ServerInstance) SetDriveIdBootable(v int32)`

SetDriveIdBootable sets DriveIdBootable field to given value.

### HasDriveIdBootable

`func (o *ServerInstance) HasDriveIdBootable() bool`

HasDriveIdBootable returns a boolean if a field has been set.

### GetServerTypeId

`func (o *ServerInstance) GetServerTypeId() int32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerInstance) GetServerTypeIdOk() (*int32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerInstance) SetServerTypeId(v int32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *ServerInstance) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetServerId

`func (o *ServerInstance) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerInstance) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerInstance) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ServerInstance) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetTemplateId

`func (o *ServerInstance) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ServerInstance) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ServerInstance) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ServerInstance) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetInstanceWanMlagId

`func (o *ServerInstance) GetInstanceWanMlagId() int32`

GetInstanceWanMlagId returns the InstanceWanMlagId field if non-nil, zero value otherwise.

### GetInstanceWanMlagIdOk

`func (o *ServerInstance) GetInstanceWanMlagIdOk() (*int32, bool)`

GetInstanceWanMlagIdOk returns a tuple with the InstanceWanMlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceWanMlagId

`func (o *ServerInstance) SetInstanceWanMlagId(v int32)`

SetInstanceWanMlagId sets InstanceWanMlagId field to given value.

### HasInstanceWanMlagId

`func (o *ServerInstance) HasInstanceWanMlagId() bool`

HasInstanceWanMlagId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstance) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstance) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstance) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstance) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetPreferredServerIds

`func (o *ServerInstance) GetPreferredServerIds() []float32`

GetPreferredServerIds returns the PreferredServerIds field if non-nil, zero value otherwise.

### GetPreferredServerIdsOk

`func (o *ServerInstance) GetPreferredServerIdsOk() (*[]float32, bool)`

GetPreferredServerIdsOk returns a tuple with the PreferredServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredServerIds

`func (o *ServerInstance) SetPreferredServerIds(v []float32)`

SetPreferredServerIds sets PreferredServerIds field to given value.

### HasPreferredServerIds

`func (o *ServerInstance) HasPreferredServerIds() bool`

HasPreferredServerIds returns a boolean if a field has been set.

### GetRaidProfile

`func (o *ServerInstance) GetRaidProfile() ServerInstanceStorageProfile`

GetRaidProfile returns the RaidProfile field if non-nil, zero value otherwise.

### GetRaidProfileOk

`func (o *ServerInstance) GetRaidProfileOk() (*ServerInstanceStorageProfile, bool)`

GetRaidProfileOk returns a tuple with the RaidProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidProfile

`func (o *ServerInstance) SetRaidProfile(v ServerInstanceStorageProfile)`

SetRaidProfile sets RaidProfile field to given value.

### HasRaidProfile

`func (o *ServerInstance) HasRaidProfile() bool`

HasRaidProfile returns a boolean if a field has been set.

### GetServiceStatus

`func (o *ServerInstance) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServerInstance) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServerInstance) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetIsVmInstance

`func (o *ServerInstance) GetIsVmInstance() int32`

GetIsVmInstance returns the IsVmInstance field if non-nil, zero value otherwise.

### GetIsVmInstanceOk

`func (o *ServerInstance) GetIsVmInstanceOk() (*int32, bool)`

GetIsVmInstanceOk returns a tuple with the IsVmInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVmInstance

`func (o *ServerInstance) SetIsVmInstance(v int32)`

SetIsVmInstance sets IsVmInstance field to given value.


### GetVmInstanceId

`func (o *ServerInstance) GetVmInstanceId() int32`

GetVmInstanceId returns the VmInstanceId field if non-nil, zero value otherwise.

### GetVmInstanceIdOk

`func (o *ServerInstance) GetVmInstanceIdOk() (*int32, bool)`

GetVmInstanceIdOk returns a tuple with the VmInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceId

`func (o *ServerInstance) SetVmInstanceId(v int32)`

SetVmInstanceId sets VmInstanceId field to given value.

### HasVmInstanceId

`func (o *ServerInstance) HasVmInstanceId() bool`

HasVmInstanceId returns a boolean if a field has been set.

### GetMeta

`func (o *ServerInstance) GetMeta() GenericGUISettings`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServerInstance) GetMetaOk() (*GenericGUISettings, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServerInstance) SetMeta(v GenericGUISettings)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServerInstance) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTags

`func (o *ServerInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerInstance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperatingSystemInfo

`func (o *ServerInstance) GetOperatingSystemInfo() map[string]interface{}`

GetOperatingSystemInfo returns the OperatingSystemInfo field if non-nil, zero value otherwise.

### GetOperatingSystemInfoOk

`func (o *ServerInstance) GetOperatingSystemInfoOk() (*map[string]interface{}, bool)`

GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemInfo

`func (o *ServerInstance) SetOperatingSystemInfo(v map[string]interface{})`

SetOperatingSystemInfo sets OperatingSystemInfo field to given value.

### HasOperatingSystemInfo

`func (o *ServerInstance) HasOperatingSystemInfo() bool`

HasOperatingSystemInfo returns a boolean if a field has been set.

### GetVariablesOsAssetsTemplateTypeAdvanced

`func (o *ServerInstance) GetVariablesOsAssetsTemplateTypeAdvanced() map[string]interface{}`

GetVariablesOsAssetsTemplateTypeAdvanced returns the VariablesOsAssetsTemplateTypeAdvanced field if non-nil, zero value otherwise.

### GetVariablesOsAssetsTemplateTypeAdvancedOk

`func (o *ServerInstance) GetVariablesOsAssetsTemplateTypeAdvancedOk() (*map[string]interface{}, bool)`

GetVariablesOsAssetsTemplateTypeAdvancedOk returns a tuple with the VariablesOsAssetsTemplateTypeAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesOsAssetsTemplateTypeAdvanced

`func (o *ServerInstance) SetVariablesOsAssetsTemplateTypeAdvanced(v map[string]interface{})`

SetVariablesOsAssetsTemplateTypeAdvanced sets VariablesOsAssetsTemplateTypeAdvanced field to given value.

### HasVariablesOsAssetsTemplateTypeAdvanced

`func (o *ServerInstance) HasVariablesOsAssetsTemplateTypeAdvanced() bool`

HasVariablesOsAssetsTemplateTypeAdvanced returns a boolean if a field has been set.

### GetSecretsOsAssetsTemplateTypeAdvanced

`func (o *ServerInstance) GetSecretsOsAssetsTemplateTypeAdvanced() map[string]interface{}`

GetSecretsOsAssetsTemplateTypeAdvanced returns the SecretsOsAssetsTemplateTypeAdvanced field if non-nil, zero value otherwise.

### GetSecretsOsAssetsTemplateTypeAdvancedOk

`func (o *ServerInstance) GetSecretsOsAssetsTemplateTypeAdvancedOk() (*map[string]interface{}, bool)`

GetSecretsOsAssetsTemplateTypeAdvancedOk returns a tuple with the SecretsOsAssetsTemplateTypeAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsOsAssetsTemplateTypeAdvanced

`func (o *ServerInstance) SetSecretsOsAssetsTemplateTypeAdvanced(v map[string]interface{})`

SetSecretsOsAssetsTemplateTypeAdvanced sets SecretsOsAssetsTemplateTypeAdvanced field to given value.

### HasSecretsOsAssetsTemplateTypeAdvanced

`func (o *ServerInstance) HasSecretsOsAssetsTemplateTypeAdvanced() bool`

HasSecretsOsAssetsTemplateTypeAdvanced returns a boolean if a field has been set.

### GetVariablesMaterializedForOSAssets

`func (o *ServerInstance) GetVariablesMaterializedForOSAssets() map[string]interface{}`

GetVariablesMaterializedForOSAssets returns the VariablesMaterializedForOSAssets field if non-nil, zero value otherwise.

### GetVariablesMaterializedForOSAssetsOk

`func (o *ServerInstance) GetVariablesMaterializedForOSAssetsOk() (*map[string]interface{}, bool)`

GetVariablesMaterializedForOSAssetsOk returns a tuple with the VariablesMaterializedForOSAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesMaterializedForOSAssets

`func (o *ServerInstance) SetVariablesMaterializedForOSAssets(v map[string]interface{})`

SetVariablesMaterializedForOSAssets sets VariablesMaterializedForOSAssets field to given value.

### HasVariablesMaterializedForOSAssets

`func (o *ServerInstance) HasVariablesMaterializedForOSAssets() bool`

HasVariablesMaterializedForOSAssets returns a boolean if a field has been set.

### GetSecretsMaterializedForOSAssets

`func (o *ServerInstance) GetSecretsMaterializedForOSAssets() map[string]interface{}`

GetSecretsMaterializedForOSAssets returns the SecretsMaterializedForOSAssets field if non-nil, zero value otherwise.

### GetSecretsMaterializedForOSAssetsOk

`func (o *ServerInstance) GetSecretsMaterializedForOSAssetsOk() (*map[string]interface{}, bool)`

GetSecretsMaterializedForOSAssetsOk returns a tuple with the SecretsMaterializedForOSAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsMaterializedForOSAssets

`func (o *ServerInstance) SetSecretsMaterializedForOSAssets(v map[string]interface{})`

SetSecretsMaterializedForOSAssets sets SecretsMaterializedForOSAssets field to given value.

### HasSecretsMaterializedForOSAssets

`func (o *ServerInstance) HasSecretsMaterializedForOSAssets() bool`

HasSecretsMaterializedForOSAssets returns a boolean if a field has been set.

### GetClusterCustomInfo

`func (o *ServerInstance) GetClusterCustomInfo() ServerInstanceClusterCustomInfo`

GetClusterCustomInfo returns the ClusterCustomInfo field if non-nil, zero value otherwise.

### GetClusterCustomInfoOk

`func (o *ServerInstance) GetClusterCustomInfoOk() (*ServerInstanceClusterCustomInfo, bool)`

GetClusterCustomInfoOk returns a tuple with the ClusterCustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCustomInfo

`func (o *ServerInstance) SetClusterCustomInfo(v ServerInstanceClusterCustomInfo)`

SetClusterCustomInfo sets ClusterCustomInfo field to given value.

### HasClusterCustomInfo

`func (o *ServerInstance) HasClusterCustomInfo() bool`

HasClusterCustomInfo returns a boolean if a field has been set.

### GetOsInstallError

`func (o *ServerInstance) GetOsInstallError() string`

GetOsInstallError returns the OsInstallError field if non-nil, zero value otherwise.

### GetOsInstallErrorOk

`func (o *ServerInstance) GetOsInstallErrorOk() (*string, bool)`

GetOsInstallErrorOk returns a tuple with the OsInstallError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallError

`func (o *ServerInstance) SetOsInstallError(v string)`

SetOsInstallError sets OsInstallError field to given value.

### HasOsInstallError

`func (o *ServerInstance) HasOsInstallError() bool`

HasOsInstallError returns a boolean if a field has been set.

### GetOsInstallImageUrl

`func (o *ServerInstance) GetOsInstallImageUrl() string`

GetOsInstallImageUrl returns the OsInstallImageUrl field if non-nil, zero value otherwise.

### GetOsInstallImageUrlOk

`func (o *ServerInstance) GetOsInstallImageUrlOk() (*string, bool)`

GetOsInstallImageUrlOk returns a tuple with the OsInstallImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallImageUrl

`func (o *ServerInstance) SetOsInstallImageUrl(v string)`

SetOsInstallImageUrl sets OsInstallImageUrl field to given value.

### HasOsInstallImageUrl

`func (o *ServerInstance) HasOsInstallImageUrl() bool`

HasOsInstallImageUrl returns a boolean if a field has been set.

### GetOsInstallImageBuildError

`func (o *ServerInstance) GetOsInstallImageBuildError() string`

GetOsInstallImageBuildError returns the OsInstallImageBuildError field if non-nil, zero value otherwise.

### GetOsInstallImageBuildErrorOk

`func (o *ServerInstance) GetOsInstallImageBuildErrorOk() (*string, bool)`

GetOsInstallImageBuildErrorOk returns a tuple with the OsInstallImageBuildError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallImageBuildError

`func (o *ServerInstance) SetOsInstallImageBuildError(v string)`

SetOsInstallImageBuildError sets OsInstallImageBuildError field to given value.

### HasOsInstallImageBuildError

`func (o *ServerInstance) HasOsInstallImageBuildError() bool`

HasOsInstallImageBuildError returns a boolean if a field has been set.

### GetOsInstallImageBuildInfo

`func (o *ServerInstance) GetOsInstallImageBuildInfo() ServerInstanceOsInstallImageBuildInfo`

GetOsInstallImageBuildInfo returns the OsInstallImageBuildInfo field if non-nil, zero value otherwise.

### GetOsInstallImageBuildInfoOk

`func (o *ServerInstance) GetOsInstallImageBuildInfoOk() (*ServerInstanceOsInstallImageBuildInfo, bool)`

GetOsInstallImageBuildInfoOk returns a tuple with the OsInstallImageBuildInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallImageBuildInfo

`func (o *ServerInstance) SetOsInstallImageBuildInfo(v ServerInstanceOsInstallImageBuildInfo)`

SetOsInstallImageBuildInfo sets OsInstallImageBuildInfo field to given value.

### HasOsInstallImageBuildInfo

`func (o *ServerInstance) HasOsInstallImageBuildInfo() bool`

HasOsInstallImageBuildInfo returns a boolean if a field has been set.

### GetOsReinstallRequired

`func (o *ServerInstance) GetOsReinstallRequired() int32`

GetOsReinstallRequired returns the OsReinstallRequired field if non-nil, zero value otherwise.

### GetOsReinstallRequiredOk

`func (o *ServerInstance) GetOsReinstallRequiredOk() (*int32, bool)`

GetOsReinstallRequiredOk returns a tuple with the OsReinstallRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsReinstallRequired

`func (o *ServerInstance) SetOsReinstallRequired(v int32)`

SetOsReinstallRequired sets OsReinstallRequired field to given value.

### HasOsReinstallRequired

`func (o *ServerInstance) HasOsReinstallRequired() bool`

HasOsReinstallRequired returns a boolean if a field has been set.

### GetIscsiInitiatorIqn

`func (o *ServerInstance) GetIscsiInitiatorIqn() string`

GetIscsiInitiatorIqn returns the IscsiInitiatorIqn field if non-nil, zero value otherwise.

### GetIscsiInitiatorIqnOk

`func (o *ServerInstance) GetIscsiInitiatorIqnOk() (*string, bool)`

GetIscsiInitiatorIqnOk returns a tuple with the IscsiInitiatorIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorIqn

`func (o *ServerInstance) SetIscsiInitiatorIqn(v string)`

SetIscsiInitiatorIqn sets IscsiInitiatorIqn field to given value.

### HasIscsiInitiatorIqn

`func (o *ServerInstance) HasIscsiInitiatorIqn() bool`

HasIscsiInitiatorIqn returns a boolean if a field has been set.

### GetIscsiInitiatorUsername

`func (o *ServerInstance) GetIscsiInitiatorUsername() string`

GetIscsiInitiatorUsername returns the IscsiInitiatorUsername field if non-nil, zero value otherwise.

### GetIscsiInitiatorUsernameOk

`func (o *ServerInstance) GetIscsiInitiatorUsernameOk() (*string, bool)`

GetIscsiInitiatorUsernameOk returns a tuple with the IscsiInitiatorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorUsername

`func (o *ServerInstance) SetIscsiInitiatorUsername(v string)`

SetIscsiInitiatorUsername sets IscsiInitiatorUsername field to given value.

### HasIscsiInitiatorUsername

`func (o *ServerInstance) HasIscsiInitiatorUsername() bool`

HasIscsiInitiatorUsername returns a boolean if a field has been set.

### GetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstance) GetIscsiInitiatorPasswordEncrypted() string`

GetIscsiInitiatorPasswordEncrypted returns the IscsiInitiatorPasswordEncrypted field if non-nil, zero value otherwise.

### GetIscsiInitiatorPasswordEncryptedOk

`func (o *ServerInstance) GetIscsiInitiatorPasswordEncryptedOk() (*string, bool)`

GetIscsiInitiatorPasswordEncryptedOk returns a tuple with the IscsiInitiatorPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstance) SetIscsiInitiatorPasswordEncrypted(v string)`

SetIscsiInitiatorPasswordEncrypted sets IscsiInitiatorPasswordEncrypted field to given value.

### HasIscsiInitiatorPasswordEncrypted

`func (o *ServerInstance) HasIscsiInitiatorPasswordEncrypted() bool`

HasIscsiInitiatorPasswordEncrypted returns a boolean if a field has been set.

### GetControlPanelUrl

`func (o *ServerInstance) GetControlPanelUrl() string`

GetControlPanelUrl returns the ControlPanelUrl field if non-nil, zero value otherwise.

### GetControlPanelUrlOk

`func (o *ServerInstance) GetControlPanelUrlOk() (*string, bool)`

GetControlPanelUrlOk returns a tuple with the ControlPanelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanelUrl

`func (o *ServerInstance) SetControlPanelUrl(v string)`

SetControlPanelUrl sets ControlPanelUrl field to given value.

### HasControlPanelUrl

`func (o *ServerInstance) HasControlPanelUrl() bool`

HasControlPanelUrl returns a boolean if a field has been set.

### GetConfig

`func (o *ServerInstance) GetConfig() ServerInstanceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerInstance) GetConfigOk() (*ServerInstanceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerInstance) SetConfig(v ServerInstanceConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinks

`func (o *ServerInstance) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstance) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstance) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


