# VmInstanceContextOSInstallationData

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

### NewVmInstanceContextOSInstallationData

`func NewVmInstanceContextOSInstallationData(site SiteOSInstallationData, siteConfig SiteConfigOSInstallationData, server ServerOSInstallationData, serverInstance ServerInstanceOSInstallationData, serverInstanceGroup ServerInstanceGroupOSInstallationData, infrastructure InfrastructureOSInstallationData, driveGroups []DriveGroupVariables, drives []DriveVariables, fileShares []FileShareVariables, buckets []BucketVariables, sharedDrives []SharedDriveVariables, network []InstanceNetworkVariables, ) *VmInstanceContextOSInstallationData`

NewVmInstanceContextOSInstallationData instantiates a new VmInstanceContextOSInstallationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmInstanceContextOSInstallationDataWithDefaults

`func NewVmInstanceContextOSInstallationDataWithDefaults() *VmInstanceContextOSInstallationData`

NewVmInstanceContextOSInstallationDataWithDefaults instantiates a new VmInstanceContextOSInstallationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *VmInstanceContextOSInstallationData) GetSite() SiteOSInstallationData`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *VmInstanceContextOSInstallationData) GetSiteOk() (*SiteOSInstallationData, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *VmInstanceContextOSInstallationData) SetSite(v SiteOSInstallationData)`

SetSite sets Site field to given value.


### GetSiteConfig

`func (o *VmInstanceContextOSInstallationData) GetSiteConfig() SiteConfigOSInstallationData`

GetSiteConfig returns the SiteConfig field if non-nil, zero value otherwise.

### GetSiteConfigOk

`func (o *VmInstanceContextOSInstallationData) GetSiteConfigOk() (*SiteConfigOSInstallationData, bool)`

GetSiteConfigOk returns a tuple with the SiteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteConfig

`func (o *VmInstanceContextOSInstallationData) SetSiteConfig(v SiteConfigOSInstallationData)`

SetSiteConfig sets SiteConfig field to given value.


### GetServer

`func (o *VmInstanceContextOSInstallationData) GetServer() ServerOSInstallationData`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *VmInstanceContextOSInstallationData) GetServerOk() (*ServerOSInstallationData, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *VmInstanceContextOSInstallationData) SetServer(v ServerOSInstallationData)`

SetServer sets Server field to given value.


### GetServerInstance

`func (o *VmInstanceContextOSInstallationData) GetServerInstance() ServerInstanceOSInstallationData`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *VmInstanceContextOSInstallationData) GetServerInstanceOk() (*ServerInstanceOSInstallationData, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *VmInstanceContextOSInstallationData) SetServerInstance(v ServerInstanceOSInstallationData)`

SetServerInstance sets ServerInstance field to given value.


### GetServerInstanceGroup

`func (o *VmInstanceContextOSInstallationData) GetServerInstanceGroup() ServerInstanceGroupOSInstallationData`

GetServerInstanceGroup returns the ServerInstanceGroup field if non-nil, zero value otherwise.

### GetServerInstanceGroupOk

`func (o *VmInstanceContextOSInstallationData) GetServerInstanceGroupOk() (*ServerInstanceGroupOSInstallationData, bool)`

GetServerInstanceGroupOk returns a tuple with the ServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroup

`func (o *VmInstanceContextOSInstallationData) SetServerInstanceGroup(v ServerInstanceGroupOSInstallationData)`

SetServerInstanceGroup sets ServerInstanceGroup field to given value.


### GetInfrastructure

`func (o *VmInstanceContextOSInstallationData) GetInfrastructure() InfrastructureOSInstallationData`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *VmInstanceContextOSInstallationData) GetInfrastructureOk() (*InfrastructureOSInstallationData, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *VmInstanceContextOSInstallationData) SetInfrastructure(v InfrastructureOSInstallationData)`

SetInfrastructure sets Infrastructure field to given value.


### GetDriveGroups

`func (o *VmInstanceContextOSInstallationData) GetDriveGroups() []DriveGroupVariables`

GetDriveGroups returns the DriveGroups field if non-nil, zero value otherwise.

### GetDriveGroupsOk

`func (o *VmInstanceContextOSInstallationData) GetDriveGroupsOk() (*[]DriveGroupVariables, bool)`

GetDriveGroupsOk returns a tuple with the DriveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveGroups

`func (o *VmInstanceContextOSInstallationData) SetDriveGroups(v []DriveGroupVariables)`

SetDriveGroups sets DriveGroups field to given value.


### GetDrives

`func (o *VmInstanceContextOSInstallationData) GetDrives() []DriveVariables`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *VmInstanceContextOSInstallationData) GetDrivesOk() (*[]DriveVariables, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *VmInstanceContextOSInstallationData) SetDrives(v []DriveVariables)`

SetDrives sets Drives field to given value.


### GetFileShares

`func (o *VmInstanceContextOSInstallationData) GetFileShares() []FileShareVariables`

GetFileShares returns the FileShares field if non-nil, zero value otherwise.

### GetFileSharesOk

`func (o *VmInstanceContextOSInstallationData) GetFileSharesOk() (*[]FileShareVariables, bool)`

GetFileSharesOk returns a tuple with the FileShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShares

`func (o *VmInstanceContextOSInstallationData) SetFileShares(v []FileShareVariables)`

SetFileShares sets FileShares field to given value.


### GetBuckets

`func (o *VmInstanceContextOSInstallationData) GetBuckets() []BucketVariables`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *VmInstanceContextOSInstallationData) GetBucketsOk() (*[]BucketVariables, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *VmInstanceContextOSInstallationData) SetBuckets(v []BucketVariables)`

SetBuckets sets Buckets field to given value.


### GetSharedDrives

`func (o *VmInstanceContextOSInstallationData) GetSharedDrives() []SharedDriveVariables`

GetSharedDrives returns the SharedDrives field if non-nil, zero value otherwise.

### GetSharedDrivesOk

`func (o *VmInstanceContextOSInstallationData) GetSharedDrivesOk() (*[]SharedDriveVariables, bool)`

GetSharedDrivesOk returns a tuple with the SharedDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrives

`func (o *VmInstanceContextOSInstallationData) SetSharedDrives(v []SharedDriveVariables)`

SetSharedDrives sets SharedDrives field to given value.


### GetNetwork

`func (o *VmInstanceContextOSInstallationData) GetNetwork() []InstanceNetworkVariables`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VmInstanceContextOSInstallationData) GetNetworkOk() (*[]InstanceNetworkVariables, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VmInstanceContextOSInstallationData) SetNetwork(v []InstanceNetworkVariables)`

SetNetwork sets Network field to given value.


### GetVariables

`func (o *VmInstanceContextOSInstallationData) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *VmInstanceContextOSInstallationData) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *VmInstanceContextOSInstallationData) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *VmInstanceContextOSInstallationData) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetSecrets

`func (o *VmInstanceContextOSInstallationData) GetSecrets() map[string]interface{}`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *VmInstanceContextOSInstallationData) GetSecretsOk() (*map[string]interface{}, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *VmInstanceContextOSInstallationData) SetSecrets(v map[string]interface{})`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *VmInstanceContextOSInstallationData) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetUserSSHKeys

`func (o *VmInstanceContextOSInstallationData) GetUserSSHKeys() []string`

GetUserSSHKeys returns the UserSSHKeys field if non-nil, zero value otherwise.

### GetUserSSHKeysOk

`func (o *VmInstanceContextOSInstallationData) GetUserSSHKeysOk() (*[]string, bool)`

GetUserSSHKeysOk returns a tuple with the UserSSHKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSSHKeys

`func (o *VmInstanceContextOSInstallationData) SetUserSSHKeys(v []string)`

SetUserSSHKeys sets UserSSHKeys field to given value.

### HasUserSSHKeys

`func (o *VmInstanceContextOSInstallationData) HasUserSSHKeys() bool`

HasUserSSHKeys returns a boolean if a field has been set.

### GetManagementSSHKey

`func (o *VmInstanceContextOSInstallationData) GetManagementSSHKey() string`

GetManagementSSHKey returns the ManagementSSHKey field if non-nil, zero value otherwise.

### GetManagementSSHKeyOk

`func (o *VmInstanceContextOSInstallationData) GetManagementSSHKeyOk() (*string, bool)`

GetManagementSSHKeyOk returns a tuple with the ManagementSSHKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSSHKey

`func (o *VmInstanceContextOSInstallationData) SetManagementSSHKey(v string)`

SetManagementSSHKey sets ManagementSSHKey field to given value.

### HasManagementSSHKey

`func (o *VmInstanceContextOSInstallationData) HasManagementSSHKey() bool`

HasManagementSSHKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


