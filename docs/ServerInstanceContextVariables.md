# ServerInstanceContextVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | [**SiteVariables**](SiteVariables.md) | The site variables. | 
**SiteConfig** | [**SiteConfigVariables**](SiteConfigVariables.md) | The site config variables. | 
**Server** | [**ServerVariables**](ServerVariables.md) | The server variables. | 
**ServerInstance** | [**ServerInstanceVariables**](ServerInstanceVariables.md) | The server instance variables. | 
**ServerInstanceGroup** | [**ServerInstanceGroupVariables**](ServerInstanceGroupVariables.md) | The server instance group variables. | 
**Infrastructure** | [**InfrastructureVariables**](InfrastructureVariables.md) | The infrastructure variables. | 
**Drives** | [**[]DriveVariables**](DriveVariables.md) | The server instance drives variables. | 
**FileShares** | [**[]FileShareVariables**](FileShareVariables.md) | The server instance file shares variables. | 
**Buckets** | [**[]BucketVariables**](BucketVariables.md) | The server instance buckets variables. | 
**SharedDrives** | [**[]SharedDriveVariables**](SharedDriveVariables.md) | The server instance shared drives variables. | 

## Methods

### NewServerInstanceContextVariables

`func NewServerInstanceContextVariables(site SiteVariables, siteConfig SiteConfigVariables, server ServerVariables, serverInstance ServerInstanceVariables, serverInstanceGroup ServerInstanceGroupVariables, infrastructure InfrastructureVariables, drives []DriveVariables, fileShares []FileShareVariables, buckets []BucketVariables, sharedDrives []SharedDriveVariables, ) *ServerInstanceContextVariables`

NewServerInstanceContextVariables instantiates a new ServerInstanceContextVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceContextVariablesWithDefaults

`func NewServerInstanceContextVariablesWithDefaults() *ServerInstanceContextVariables`

NewServerInstanceContextVariablesWithDefaults instantiates a new ServerInstanceContextVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *ServerInstanceContextVariables) GetSite() SiteVariables`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ServerInstanceContextVariables) GetSiteOk() (*SiteVariables, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ServerInstanceContextVariables) SetSite(v SiteVariables)`

SetSite sets Site field to given value.


### GetSiteConfig

`func (o *ServerInstanceContextVariables) GetSiteConfig() SiteConfigVariables`

GetSiteConfig returns the SiteConfig field if non-nil, zero value otherwise.

### GetSiteConfigOk

`func (o *ServerInstanceContextVariables) GetSiteConfigOk() (*SiteConfigVariables, bool)`

GetSiteConfigOk returns a tuple with the SiteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteConfig

`func (o *ServerInstanceContextVariables) SetSiteConfig(v SiteConfigVariables)`

SetSiteConfig sets SiteConfig field to given value.


### GetServer

`func (o *ServerInstanceContextVariables) GetServer() ServerVariables`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerInstanceContextVariables) GetServerOk() (*ServerVariables, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerInstanceContextVariables) SetServer(v ServerVariables)`

SetServer sets Server field to given value.


### GetServerInstance

`func (o *ServerInstanceContextVariables) GetServerInstance() ServerInstanceVariables`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *ServerInstanceContextVariables) GetServerInstanceOk() (*ServerInstanceVariables, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *ServerInstanceContextVariables) SetServerInstance(v ServerInstanceVariables)`

SetServerInstance sets ServerInstance field to given value.


### GetServerInstanceGroup

`func (o *ServerInstanceContextVariables) GetServerInstanceGroup() ServerInstanceGroupVariables`

GetServerInstanceGroup returns the ServerInstanceGroup field if non-nil, zero value otherwise.

### GetServerInstanceGroupOk

`func (o *ServerInstanceContextVariables) GetServerInstanceGroupOk() (*ServerInstanceGroupVariables, bool)`

GetServerInstanceGroupOk returns a tuple with the ServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroup

`func (o *ServerInstanceContextVariables) SetServerInstanceGroup(v ServerInstanceGroupVariables)`

SetServerInstanceGroup sets ServerInstanceGroup field to given value.


### GetInfrastructure

`func (o *ServerInstanceContextVariables) GetInfrastructure() InfrastructureVariables`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ServerInstanceContextVariables) GetInfrastructureOk() (*InfrastructureVariables, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ServerInstanceContextVariables) SetInfrastructure(v InfrastructureVariables)`

SetInfrastructure sets Infrastructure field to given value.


### GetDrives

`func (o *ServerInstanceContextVariables) GetDrives() []DriveVariables`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *ServerInstanceContextVariables) GetDrivesOk() (*[]DriveVariables, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *ServerInstanceContextVariables) SetDrives(v []DriveVariables)`

SetDrives sets Drives field to given value.


### GetFileShares

`func (o *ServerInstanceContextVariables) GetFileShares() []FileShareVariables`

GetFileShares returns the FileShares field if non-nil, zero value otherwise.

### GetFileSharesOk

`func (o *ServerInstanceContextVariables) GetFileSharesOk() (*[]FileShareVariables, bool)`

GetFileSharesOk returns a tuple with the FileShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShares

`func (o *ServerInstanceContextVariables) SetFileShares(v []FileShareVariables)`

SetFileShares sets FileShares field to given value.


### GetBuckets

`func (o *ServerInstanceContextVariables) GetBuckets() []BucketVariables`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ServerInstanceContextVariables) GetBucketsOk() (*[]BucketVariables, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ServerInstanceContextVariables) SetBuckets(v []BucketVariables)`

SetBuckets sets Buckets field to given value.


### GetSharedDrives

`func (o *ServerInstanceContextVariables) GetSharedDrives() []SharedDriveVariables`

GetSharedDrives returns the SharedDrives field if non-nil, zero value otherwise.

### GetSharedDrivesOk

`func (o *ServerInstanceContextVariables) GetSharedDrivesOk() (*[]SharedDriveVariables, bool)`

GetSharedDrivesOk returns a tuple with the SharedDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrives

`func (o *ServerInstanceContextVariables) SetSharedDrives(v []SharedDriveVariables)`

SetSharedDrives sets SharedDrives field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


