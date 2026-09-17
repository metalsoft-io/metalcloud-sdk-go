# ContainerInstanceContextOSInstallationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | [**SiteOSInstallationData**](SiteOSInstallationData.md) | The site OS installation data. | 
**SiteConfig** | [**SiteConfigOSInstallationData**](SiteConfigOSInstallationData.md) | The site config OS installation data. | 
**Server** | [**ServerOSInstallationData**](ServerOSInstallationData.md) | The server OS installation data. | 
**ServerInstance** | [**ServerInstanceOSInstallationData**](ServerInstanceOSInstallationData.md) | The server instance OS installation data. | 
**ServerInstanceGroup** | [**ServerInstanceGroupOSInstallationData**](ServerInstanceGroupOSInstallationData.md) | The server instance group OS installation data. | 
**Infrastructure** | [**InfrastructureOSInstallationData**](InfrastructureOSInstallationData.md) | The infrastructure OS installation data. | 
**DriveGroups** | [**[]DriveGroupVariables**](DriveGroupVariables.md) | The server instance drive groups OS installation data. | 
**Drives** | [**[]DriveVariables**](DriveVariables.md) | The server instance drives OS installation data. | 
**FileShares** | [**[]FileShareVariables**](FileShareVariables.md) | The server instance file shares OS installation data. | 
**Buckets** | [**[]BucketVariables**](BucketVariables.md) | The server instance buckets OS installation data. | 
**SharedDrives** | [**[]SharedDriveVariables**](SharedDriveVariables.md) | The server instance shared drives OS installation data. | 
**Network** | [**[]InstanceNetworkVariables**](InstanceNetworkVariables.md) | The server instance network variables. | 
**Variables** | Pointer to **map[string]interface{}** | Additional variables | [optional] 
**Secrets** | Pointer to **map[string]interface{}** | Secrets | [optional] 
**UserSSHKeys** | Pointer to **[]string** | Infrastructure owner SSH keys | [optional] 
**ManagementSSHKey** | Pointer to **string** | Management SSH key | [optional] 

## Methods

### NewContainerInstanceContextOSInstallationData

`func NewContainerInstanceContextOSInstallationData(site SiteOSInstallationData, siteConfig SiteConfigOSInstallationData, server ServerOSInstallationData, serverInstance ServerInstanceOSInstallationData, serverInstanceGroup ServerInstanceGroupOSInstallationData, infrastructure InfrastructureOSInstallationData, driveGroups []DriveGroupVariables, drives []DriveVariables, fileShares []FileShareVariables, buckets []BucketVariables, sharedDrives []SharedDriveVariables, network []InstanceNetworkVariables, ) *ContainerInstanceContextOSInstallationData`

NewContainerInstanceContextOSInstallationData instantiates a new ContainerInstanceContextOSInstallationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceContextOSInstallationDataWithDefaults

`func NewContainerInstanceContextOSInstallationDataWithDefaults() *ContainerInstanceContextOSInstallationData`

NewContainerInstanceContextOSInstallationDataWithDefaults instantiates a new ContainerInstanceContextOSInstallationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *ContainerInstanceContextOSInstallationData) GetSite() SiteOSInstallationData`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ContainerInstanceContextOSInstallationData) GetSiteOk() (*SiteOSInstallationData, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ContainerInstanceContextOSInstallationData) SetSite(v SiteOSInstallationData)`

SetSite sets Site field to given value.


### GetSiteConfig

`func (o *ContainerInstanceContextOSInstallationData) GetSiteConfig() SiteConfigOSInstallationData`

GetSiteConfig returns the SiteConfig field if non-nil, zero value otherwise.

### GetSiteConfigOk

`func (o *ContainerInstanceContextOSInstallationData) GetSiteConfigOk() (*SiteConfigOSInstallationData, bool)`

GetSiteConfigOk returns a tuple with the SiteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteConfig

`func (o *ContainerInstanceContextOSInstallationData) SetSiteConfig(v SiteConfigOSInstallationData)`

SetSiteConfig sets SiteConfig field to given value.


### GetServer

`func (o *ContainerInstanceContextOSInstallationData) GetServer() ServerOSInstallationData`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ContainerInstanceContextOSInstallationData) GetServerOk() (*ServerOSInstallationData, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ContainerInstanceContextOSInstallationData) SetServer(v ServerOSInstallationData)`

SetServer sets Server field to given value.


### GetServerInstance

`func (o *ContainerInstanceContextOSInstallationData) GetServerInstance() ServerInstanceOSInstallationData`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *ContainerInstanceContextOSInstallationData) GetServerInstanceOk() (*ServerInstanceOSInstallationData, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *ContainerInstanceContextOSInstallationData) SetServerInstance(v ServerInstanceOSInstallationData)`

SetServerInstance sets ServerInstance field to given value.


### GetServerInstanceGroup

`func (o *ContainerInstanceContextOSInstallationData) GetServerInstanceGroup() ServerInstanceGroupOSInstallationData`

GetServerInstanceGroup returns the ServerInstanceGroup field if non-nil, zero value otherwise.

### GetServerInstanceGroupOk

`func (o *ContainerInstanceContextOSInstallationData) GetServerInstanceGroupOk() (*ServerInstanceGroupOSInstallationData, bool)`

GetServerInstanceGroupOk returns a tuple with the ServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroup

`func (o *ContainerInstanceContextOSInstallationData) SetServerInstanceGroup(v ServerInstanceGroupOSInstallationData)`

SetServerInstanceGroup sets ServerInstanceGroup field to given value.


### GetInfrastructure

`func (o *ContainerInstanceContextOSInstallationData) GetInfrastructure() InfrastructureOSInstallationData`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ContainerInstanceContextOSInstallationData) GetInfrastructureOk() (*InfrastructureOSInstallationData, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ContainerInstanceContextOSInstallationData) SetInfrastructure(v InfrastructureOSInstallationData)`

SetInfrastructure sets Infrastructure field to given value.


### GetDriveGroups

`func (o *ContainerInstanceContextOSInstallationData) GetDriveGroups() []DriveGroupVariables`

GetDriveGroups returns the DriveGroups field if non-nil, zero value otherwise.

### GetDriveGroupsOk

`func (o *ContainerInstanceContextOSInstallationData) GetDriveGroupsOk() (*[]DriveGroupVariables, bool)`

GetDriveGroupsOk returns a tuple with the DriveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroups

`func (o *ContainerInstanceContextOSInstallationData) SetDriveGroups(v []DriveGroupVariables)`

SetDriveGroups sets DriveGroups field to given value.


### GetDrives

`func (o *ContainerInstanceContextOSInstallationData) GetDrives() []DriveVariables`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *ContainerInstanceContextOSInstallationData) GetDrivesOk() (*[]DriveVariables, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *ContainerInstanceContextOSInstallationData) SetDrives(v []DriveVariables)`

SetDrives sets Drives field to given value.


### GetFileShares

`func (o *ContainerInstanceContextOSInstallationData) GetFileShares() []FileShareVariables`

GetFileShares returns the FileShares field if non-nil, zero value otherwise.

### GetFileSharesOk

`func (o *ContainerInstanceContextOSInstallationData) GetFileSharesOk() (*[]FileShareVariables, bool)`

GetFileSharesOk returns a tuple with the FileShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShares

`func (o *ContainerInstanceContextOSInstallationData) SetFileShares(v []FileShareVariables)`

SetFileShares sets FileShares field to given value.


### GetBuckets

`func (o *ContainerInstanceContextOSInstallationData) GetBuckets() []BucketVariables`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ContainerInstanceContextOSInstallationData) GetBucketsOk() (*[]BucketVariables, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ContainerInstanceContextOSInstallationData) SetBuckets(v []BucketVariables)`

SetBuckets sets Buckets field to given value.


### GetSharedDrives

`func (o *ContainerInstanceContextOSInstallationData) GetSharedDrives() []SharedDriveVariables`

GetSharedDrives returns the SharedDrives field if non-nil, zero value otherwise.

### GetSharedDrivesOk

`func (o *ContainerInstanceContextOSInstallationData) GetSharedDrivesOk() (*[]SharedDriveVariables, bool)`

GetSharedDrivesOk returns a tuple with the SharedDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrives

`func (o *ContainerInstanceContextOSInstallationData) SetSharedDrives(v []SharedDriveVariables)`

SetSharedDrives sets SharedDrives field to given value.


### GetNetwork

`func (o *ContainerInstanceContextOSInstallationData) GetNetwork() []InstanceNetworkVariables`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ContainerInstanceContextOSInstallationData) GetNetworkOk() (*[]InstanceNetworkVariables, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ContainerInstanceContextOSInstallationData) SetNetwork(v []InstanceNetworkVariables)`

SetNetwork sets Network field to given value.


### GetVariables

`func (o *ContainerInstanceContextOSInstallationData) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ContainerInstanceContextOSInstallationData) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ContainerInstanceContextOSInstallationData) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ContainerInstanceContextOSInstallationData) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSecrets

`func (o *ContainerInstanceContextOSInstallationData) GetSecrets() map[string]interface{}`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ContainerInstanceContextOSInstallationData) GetSecretsOk() (*map[string]interface{}, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ContainerInstanceContextOSInstallationData) SetSecrets(v map[string]interface{})`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ContainerInstanceContextOSInstallationData) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetUserSSHKeys

`func (o *ContainerInstanceContextOSInstallationData) GetUserSSHKeys() []string`

GetUserSSHKeys returns the UserSSHKeys field if non-nil, zero value otherwise.

### GetUserSSHKeysOk

`func (o *ContainerInstanceContextOSInstallationData) GetUserSSHKeysOk() (*[]string, bool)`

GetUserSSHKeysOk returns a tuple with the UserSSHKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSSHKeys

`func (o *ContainerInstanceContextOSInstallationData) SetUserSSHKeys(v []string)`

SetUserSSHKeys sets UserSSHKeys field to given value.

### HasUserSSHKeys

`func (o *ContainerInstanceContextOSInstallationData) HasUserSSHKeys() bool`

HasUserSSHKeys returns a boolean if a field has been set.

### GetManagementSSHKey

`func (o *ContainerInstanceContextOSInstallationData) GetManagementSSHKey() string`

GetManagementSSHKey returns the ManagementSSHKey field if non-nil, zero value otherwise.

### GetManagementSSHKeyOk

`func (o *ContainerInstanceContextOSInstallationData) GetManagementSSHKeyOk() (*string, bool)`

GetManagementSSHKeyOk returns a tuple with the ManagementSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSSHKey

`func (o *ContainerInstanceContextOSInstallationData) SetManagementSSHKey(v string)`

SetManagementSSHKey sets ManagementSSHKey field to given value.

### HasManagementSSHKey

`func (o *ContainerInstanceContextOSInstallationData) HasManagementSSHKey() bool`

HasManagementSSHKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


