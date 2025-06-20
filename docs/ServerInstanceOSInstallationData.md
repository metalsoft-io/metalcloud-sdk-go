# ServerInstanceOSInstallationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Product Instance ID. | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Product Instance. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomStorageProfile** | Pointer to [**ServerInstanceStorageProfile**](ServerInstanceStorageProfile.md) | RAID profile for the Instance Interface. | [optional] 
**IsVmInstance** | **int32** | Flag to indicate if this is a VM instance | 
**IsEndpointInstance** | **int32** | Flag to indicate if this is an Endpoint Instance | 
**Hostname** | **string** | The subdomain of the server instance. | 
**Fqdn** | **string** | The subdomain of the server instance. | 
**OsCredentials** | Pointer to [**ServerInstanceOsCredentialInstallationData**](ServerInstanceOsCredentialInstallationData.md) |  | [optional] 
**InitiatorNqn** | Pointer to **string** | NVMe Initiator NQN for the Instance. | [optional] 
**IscsiInitiatorIqn** | Pointer to **string** | iSCSI Initiator IQN for the Instance Interface. | [optional] 
**IscsiInitiatorUsername** | Pointer to **string** | iSCSI Initiator Username for the Instance Interface. | [optional] 
**IscsiInitiatorPasswordEncrypted** | Pointer to **string** | iSCSI Initiator Password for the Instance Interface. | [optional] 

## Methods

### NewServerInstanceOSInstallationData

`func NewServerInstanceOSInstallationData(id int32, label string, isVmInstance int32, isEndpointInstance int32, hostname string, fqdn string, ) *ServerInstanceOSInstallationData`

NewServerInstanceOSInstallationData instantiates a new ServerInstanceOSInstallationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceOSInstallationDataWithDefaults

`func NewServerInstanceOSInstallationDataWithDefaults() *ServerInstanceOSInstallationData`

NewServerInstanceOSInstallationDataWithDefaults instantiates a new ServerInstanceOSInstallationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceOSInstallationData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceOSInstallationData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceOSInstallationData) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *ServerInstanceOSInstallationData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceOSInstallationData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceOSInstallationData) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomainPermanent

`func (o *ServerInstanceOSInstallationData) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstanceOSInstallationData) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstanceOSInstallationData) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstanceOSInstallationData) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceOSInstallationData) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceOSInstallationData) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceOSInstallationData) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceOSInstallationData) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetCustomStorageProfile

`func (o *ServerInstanceOSInstallationData) GetCustomStorageProfile() ServerInstanceStorageProfile`

GetCustomStorageProfile returns the CustomStorageProfile field if non-nil, zero value otherwise.

### GetCustomStorageProfileOk

`func (o *ServerInstanceOSInstallationData) GetCustomStorageProfileOk() (*ServerInstanceStorageProfile, bool)`

GetCustomStorageProfileOk returns a tuple with the CustomStorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStorageProfile

`func (o *ServerInstanceOSInstallationData) SetCustomStorageProfile(v ServerInstanceStorageProfile)`

SetCustomStorageProfile sets CustomStorageProfile field to given value.

### HasCustomStorageProfile

`func (o *ServerInstanceOSInstallationData) HasCustomStorageProfile() bool`

HasCustomStorageProfile returns a boolean if a field has been set.

### GetIsVmInstance

`func (o *ServerInstanceOSInstallationData) GetIsVmInstance() int32`

GetIsVmInstance returns the IsVmInstance field if non-nil, zero value otherwise.

### GetIsVmInstanceOk

`func (o *ServerInstanceOSInstallationData) GetIsVmInstanceOk() (*int32, bool)`

GetIsVmInstanceOk returns a tuple with the IsVmInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVmInstance

`func (o *ServerInstanceOSInstallationData) SetIsVmInstance(v int32)`

SetIsVmInstance sets IsVmInstance field to given value.


### GetIsEndpointInstance

`func (o *ServerInstanceOSInstallationData) GetIsEndpointInstance() int32`

GetIsEndpointInstance returns the IsEndpointInstance field if non-nil, zero value otherwise.

### GetIsEndpointInstanceOk

`func (o *ServerInstanceOSInstallationData) GetIsEndpointInstanceOk() (*int32, bool)`

GetIsEndpointInstanceOk returns a tuple with the IsEndpointInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEndpointInstance

`func (o *ServerInstanceOSInstallationData) SetIsEndpointInstance(v int32)`

SetIsEndpointInstance sets IsEndpointInstance field to given value.


### GetHostname

`func (o *ServerInstanceOSInstallationData) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerInstanceOSInstallationData) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerInstanceOSInstallationData) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetFqdn

`func (o *ServerInstanceOSInstallationData) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ServerInstanceOSInstallationData) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ServerInstanceOSInstallationData) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetOsCredentials

`func (o *ServerInstanceOSInstallationData) GetOsCredentials() ServerInstanceOsCredentialInstallationData`

GetOsCredentials returns the OsCredentials field if non-nil, zero value otherwise.

### GetOsCredentialsOk

`func (o *ServerInstanceOSInstallationData) GetOsCredentialsOk() (*ServerInstanceOsCredentialInstallationData, bool)`

GetOsCredentialsOk returns a tuple with the OsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCredentials

`func (o *ServerInstanceOSInstallationData) SetOsCredentials(v ServerInstanceOsCredentialInstallationData)`

SetOsCredentials sets OsCredentials field to given value.

### HasOsCredentials

`func (o *ServerInstanceOSInstallationData) HasOsCredentials() bool`

HasOsCredentials returns a boolean if a field has been set.

### GetInitiatorNqn

`func (o *ServerInstanceOSInstallationData) GetInitiatorNqn() string`

GetInitiatorNqn returns the InitiatorNqn field if non-nil, zero value otherwise.

### GetInitiatorNqnOk

`func (o *ServerInstanceOSInstallationData) GetInitiatorNqnOk() (*string, bool)`

GetInitiatorNqnOk returns a tuple with the InitiatorNqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorNqn

`func (o *ServerInstanceOSInstallationData) SetInitiatorNqn(v string)`

SetInitiatorNqn sets InitiatorNqn field to given value.

### HasInitiatorNqn

`func (o *ServerInstanceOSInstallationData) HasInitiatorNqn() bool`

HasInitiatorNqn returns a boolean if a field has been set.

### GetIscsiInitiatorIqn

`func (o *ServerInstanceOSInstallationData) GetIscsiInitiatorIqn() string`

GetIscsiInitiatorIqn returns the IscsiInitiatorIqn field if non-nil, zero value otherwise.

### GetIscsiInitiatorIqnOk

`func (o *ServerInstanceOSInstallationData) GetIscsiInitiatorIqnOk() (*string, bool)`

GetIscsiInitiatorIqnOk returns a tuple with the IscsiInitiatorIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorIqn

`func (o *ServerInstanceOSInstallationData) SetIscsiInitiatorIqn(v string)`

SetIscsiInitiatorIqn sets IscsiInitiatorIqn field to given value.

### HasIscsiInitiatorIqn

`func (o *ServerInstanceOSInstallationData) HasIscsiInitiatorIqn() bool`

HasIscsiInitiatorIqn returns a boolean if a field has been set.

### GetIscsiInitiatorUsername

`func (o *ServerInstanceOSInstallationData) GetIscsiInitiatorUsername() string`

GetIscsiInitiatorUsername returns the IscsiInitiatorUsername field if non-nil, zero value otherwise.

### GetIscsiInitiatorUsernameOk

`func (o *ServerInstanceOSInstallationData) GetIscsiInitiatorUsernameOk() (*string, bool)`

GetIscsiInitiatorUsernameOk returns a tuple with the IscsiInitiatorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorUsername

`func (o *ServerInstanceOSInstallationData) SetIscsiInitiatorUsername(v string)`

SetIscsiInitiatorUsername sets IscsiInitiatorUsername field to given value.

### HasIscsiInitiatorUsername

`func (o *ServerInstanceOSInstallationData) HasIscsiInitiatorUsername() bool`

HasIscsiInitiatorUsername returns a boolean if a field has been set.

### GetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceOSInstallationData) GetIscsiInitiatorPasswordEncrypted() string`

GetIscsiInitiatorPasswordEncrypted returns the IscsiInitiatorPasswordEncrypted field if non-nil, zero value otherwise.

### GetIscsiInitiatorPasswordEncryptedOk

`func (o *ServerInstanceOSInstallationData) GetIscsiInitiatorPasswordEncryptedOk() (*string, bool)`

GetIscsiInitiatorPasswordEncryptedOk returns a tuple with the IscsiInitiatorPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceOSInstallationData) SetIscsiInitiatorPasswordEncrypted(v string)`

SetIscsiInitiatorPasswordEncrypted sets IscsiInitiatorPasswordEncrypted field to given value.

### HasIscsiInitiatorPasswordEncrypted

`func (o *ServerInstanceOSInstallationData) HasIscsiInitiatorPasswordEncrypted() bool`

HasIscsiInitiatorPasswordEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


