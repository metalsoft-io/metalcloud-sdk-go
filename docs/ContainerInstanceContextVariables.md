# ContainerInstanceContextVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | [**SiteVariables**](SiteVariables.md) | The site variables. | 
**SiteConfig** | [**SiteConfigVariables**](SiteConfigVariables.md) | The site config variables. | 
**Server** | [**ServerVariables**](ServerVariables.md) | The server variables. | 
**ServerInstance** | [**ServerInstanceVariables**](ServerInstanceVariables.md) | The server instance variables. | 
**ServerInstanceGroup** | [**ServerInstanceGroupVariables**](ServerInstanceGroupVariables.md) | The server instance group variables. | 
**Infrastructure** | [**InfrastructureVariables**](InfrastructureVariables.md) | The infrastructure variables. | 
**DriveGroups** | [**[]DriveGroupVariables**](DriveGroupVariables.md) | The server instance drive groups variables. | 
**Drives** | [**[]DriveVariables**](DriveVariables.md) | The server instance drives variables. | 
**FileShares** | [**[]FileShareVariables**](FileShareVariables.md) | The server instance file shares variables. | 
**Buckets** | [**[]BucketVariables**](BucketVariables.md) | The server instance buckets variables. | 
**SharedDrives** | [**[]SharedDriveVariables**](SharedDriveVariables.md) | The server instance shared drives variables. | 
**Network** | [**[]InstanceNetworkVariables**](InstanceNetworkVariables.md) | The server instance network variables. | 
**Variables** | Pointer to **map[string]interface{}** | Additional variables | [optional] 
**Secrets** | Pointer to **map[string]interface{}** | Secrets | [optional] 
**UserSSHKeys** | Pointer to **[]string** | Infrastructure owner SSH keys | [optional] 
**ManagementSSHKey** | Pointer to **string** | Management SSH key | [optional] 

## Methods

### NewContainerInstanceContextVariables

`func NewContainerInstanceContextVariables(site SiteVariables, siteConfig SiteConfigVariables, server ServerVariables, serverInstance ServerInstanceVariables, serverInstanceGroup ServerInstanceGroupVariables, infrastructure InfrastructureVariables, driveGroups []DriveGroupVariables, drives []DriveVariables, fileShares []FileShareVariables, buckets []BucketVariables, sharedDrives []SharedDriveVariables, network []InstanceNetworkVariables, ) *ContainerInstanceContextVariables`

NewContainerInstanceContextVariables instantiates a new ContainerInstanceContextVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceContextVariablesWithDefaults

`func NewContainerInstanceContextVariablesWithDefaults() *ContainerInstanceContextVariables`

NewContainerInstanceContextVariablesWithDefaults instantiates a new ContainerInstanceContextVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *ContainerInstanceContextVariables) GetSite() SiteVariables`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ContainerInstanceContextVariables) GetSiteOk() (*SiteVariables, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ContainerInstanceContextVariables) SetSite(v SiteVariables)`

SetSite sets Site field to given value.


### GetSiteConfig

`func (o *ContainerInstanceContextVariables) GetSiteConfig() SiteConfigVariables`

GetSiteConfig returns the SiteConfig field if non-nil, zero value otherwise.

### GetSiteConfigOk

`func (o *ContainerInstanceContextVariables) GetSiteConfigOk() (*SiteConfigVariables, bool)`

GetSiteConfigOk returns a tuple with the SiteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteConfig

`func (o *ContainerInstanceContextVariables) SetSiteConfig(v SiteConfigVariables)`

SetSiteConfig sets SiteConfig field to given value.


### GetServer

`func (o *ContainerInstanceContextVariables) GetServer() ServerVariables`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ContainerInstanceContextVariables) GetServerOk() (*ServerVariables, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ContainerInstanceContextVariables) SetServer(v ServerVariables)`

SetServer sets Server field to given value.


### GetServerInstance

`func (o *ContainerInstanceContextVariables) GetServerInstance() ServerInstanceVariables`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *ContainerInstanceContextVariables) GetServerInstanceOk() (*ServerInstanceVariables, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *ContainerInstanceContextVariables) SetServerInstance(v ServerInstanceVariables)`

SetServerInstance sets ServerInstance field to given value.


### GetServerInstanceGroup

`func (o *ContainerInstanceContextVariables) GetServerInstanceGroup() ServerInstanceGroupVariables`

GetServerInstanceGroup returns the ServerInstanceGroup field if non-nil, zero value otherwise.

### GetServerInstanceGroupOk

`func (o *ContainerInstanceContextVariables) GetServerInstanceGroupOk() (*ServerInstanceGroupVariables, bool)`

GetServerInstanceGroupOk returns a tuple with the ServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroup

`func (o *ContainerInstanceContextVariables) SetServerInstanceGroup(v ServerInstanceGroupVariables)`

SetServerInstanceGroup sets ServerInstanceGroup field to given value.


### GetInfrastructure

`func (o *ContainerInstanceContextVariables) GetInfrastructure() InfrastructureVariables`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ContainerInstanceContextVariables) GetInfrastructureOk() (*InfrastructureVariables, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ContainerInstanceContextVariables) SetInfrastructure(v InfrastructureVariables)`

SetInfrastructure sets Infrastructure field to given value.


### GetDriveGroups

`func (o *ContainerInstanceContextVariables) GetDriveGroups() []DriveGroupVariables`

GetDriveGroups returns the DriveGroups field if non-nil, zero value otherwise.

### GetDriveGroupsOk

`func (o *ContainerInstanceContextVariables) GetDriveGroupsOk() (*[]DriveGroupVariables, bool)`

GetDriveGroupsOk returns a tuple with the DriveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroups

`func (o *ContainerInstanceContextVariables) SetDriveGroups(v []DriveGroupVariables)`

SetDriveGroups sets DriveGroups field to given value.


### GetDrives

`func (o *ContainerInstanceContextVariables) GetDrives() []DriveVariables`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *ContainerInstanceContextVariables) GetDrivesOk() (*[]DriveVariables, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *ContainerInstanceContextVariables) SetDrives(v []DriveVariables)`

SetDrives sets Drives field to given value.


### GetFileShares

`func (o *ContainerInstanceContextVariables) GetFileShares() []FileShareVariables`

GetFileShares returns the FileShares field if non-nil, zero value otherwise.

### GetFileSharesOk

`func (o *ContainerInstanceContextVariables) GetFileSharesOk() (*[]FileShareVariables, bool)`

GetFileSharesOk returns a tuple with the FileShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShares

`func (o *ContainerInstanceContextVariables) SetFileShares(v []FileShareVariables)`

SetFileShares sets FileShares field to given value.


### GetBuckets

`func (o *ContainerInstanceContextVariables) GetBuckets() []BucketVariables`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ContainerInstanceContextVariables) GetBucketsOk() (*[]BucketVariables, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ContainerInstanceContextVariables) SetBuckets(v []BucketVariables)`

SetBuckets sets Buckets field to given value.


### GetSharedDrives

`func (o *ContainerInstanceContextVariables) GetSharedDrives() []SharedDriveVariables`

GetSharedDrives returns the SharedDrives field if non-nil, zero value otherwise.

### GetSharedDrivesOk

`func (o *ContainerInstanceContextVariables) GetSharedDrivesOk() (*[]SharedDriveVariables, bool)`

GetSharedDrivesOk returns a tuple with the SharedDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrives

`func (o *ContainerInstanceContextVariables) SetSharedDrives(v []SharedDriveVariables)`

SetSharedDrives sets SharedDrives field to given value.


### GetNetwork

`func (o *ContainerInstanceContextVariables) GetNetwork() []InstanceNetworkVariables`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ContainerInstanceContextVariables) GetNetworkOk() (*[]InstanceNetworkVariables, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ContainerInstanceContextVariables) SetNetwork(v []InstanceNetworkVariables)`

SetNetwork sets Network field to given value.


### GetVariables

`func (o *ContainerInstanceContextVariables) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ContainerInstanceContextVariables) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ContainerInstanceContextVariables) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ContainerInstanceContextVariables) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSecrets

`func (o *ContainerInstanceContextVariables) GetSecrets() map[string]interface{}`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ContainerInstanceContextVariables) GetSecretsOk() (*map[string]interface{}, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ContainerInstanceContextVariables) SetSecrets(v map[string]interface{})`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ContainerInstanceContextVariables) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetUserSSHKeys

`func (o *ContainerInstanceContextVariables) GetUserSSHKeys() []string`

GetUserSSHKeys returns the UserSSHKeys field if non-nil, zero value otherwise.

### GetUserSSHKeysOk

`func (o *ContainerInstanceContextVariables) GetUserSSHKeysOk() (*[]string, bool)`

GetUserSSHKeysOk returns a tuple with the UserSSHKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSSHKeys

`func (o *ContainerInstanceContextVariables) SetUserSSHKeys(v []string)`

SetUserSSHKeys sets UserSSHKeys field to given value.

### HasUserSSHKeys

`func (o *ContainerInstanceContextVariables) HasUserSSHKeys() bool`

HasUserSSHKeys returns a boolean if a field has been set.

### GetManagementSSHKey

`func (o *ContainerInstanceContextVariables) GetManagementSSHKey() string`

GetManagementSSHKey returns the ManagementSSHKey field if non-nil, zero value otherwise.

### GetManagementSSHKeyOk

`func (o *ContainerInstanceContextVariables) GetManagementSSHKeyOk() (*string, bool)`

GetManagementSSHKeyOk returns a tuple with the ManagementSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSSHKey

`func (o *ContainerInstanceContextVariables) SetManagementSSHKey(v string)`

SetManagementSSHKey sets ManagementSSHKey field to given value.

### HasManagementSSHKey

`func (o *ContainerInstanceContextVariables) HasManagementSSHKey() bool`

HasManagementSSHKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


