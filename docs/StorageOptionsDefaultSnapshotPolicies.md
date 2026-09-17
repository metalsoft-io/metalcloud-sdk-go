# StorageOptionsDefaultSnapshotPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Svm** | Pointer to [**StorageOptionsDefaultSnapshotPolicySettings**](StorageOptionsDefaultSnapshotPolicySettings.md) | Default Snapshot policy settings for the SVM entity (applicable to NetApp ONTAP. 1 NetApp SVM corresponds to 1 MetalSoft logical network inside an infrastructure. Public logical networks will have 1 SVM for each infrastructure they are in) | [optional] 
**Volume** | Pointer to [**StorageOptionsDefaultSnapshotPolicySettings**](StorageOptionsDefaultSnapshotPolicySettings.md) | Default Snapshot policy settings for the volume entity (applicable to block storages if supported) | [optional] 
**FileSystem** | Pointer to [**StorageOptionsDefaultSnapshotPolicySettings**](StorageOptionsDefaultSnapshotPolicySettings.md) | Default Snapshot policy settings for the file system entity (applicable to file storages if supported) | [optional] 
**ObjectStore** | Pointer to [**StorageOptionsDefaultSnapshotPolicySettings**](StorageOptionsDefaultSnapshotPolicySettings.md) | Default Snapshot policy settings for the object store entity (applicable to object storages if supported) | [optional] 

## Methods

### NewStorageOptionsDefaultSnapshotPolicies

`func NewStorageOptionsDefaultSnapshotPolicies() *StorageOptionsDefaultSnapshotPolicies`

NewStorageOptionsDefaultSnapshotPolicies instantiates a new StorageOptionsDefaultSnapshotPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageOptionsDefaultSnapshotPoliciesWithDefaults

`func NewStorageOptionsDefaultSnapshotPoliciesWithDefaults() *StorageOptionsDefaultSnapshotPolicies`

NewStorageOptionsDefaultSnapshotPoliciesWithDefaults instantiates a new StorageOptionsDefaultSnapshotPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSvm

`func (o *StorageOptionsDefaultSnapshotPolicies) GetSvm() StorageOptionsDefaultSnapshotPolicySettings`

GetSvm returns the Svm field if non-nil, zero value otherwise.

### GetSvmOk

`func (o *StorageOptionsDefaultSnapshotPolicies) GetSvmOk() (*StorageOptionsDefaultSnapshotPolicySettings, bool)`

GetSvmOk returns a tuple with the Svm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvm

`func (o *StorageOptionsDefaultSnapshotPolicies) SetSvm(v StorageOptionsDefaultSnapshotPolicySettings)`

SetSvm sets Svm field to given value.

### HasSvm

`func (o *StorageOptionsDefaultSnapshotPolicies) HasSvm() bool`

HasSvm returns a boolean if a field has been set.

### GetVolume

`func (o *StorageOptionsDefaultSnapshotPolicies) GetVolume() StorageOptionsDefaultSnapshotPolicySettings`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageOptionsDefaultSnapshotPolicies) GetVolumeOk() (*StorageOptionsDefaultSnapshotPolicySettings, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageOptionsDefaultSnapshotPolicies) SetVolume(v StorageOptionsDefaultSnapshotPolicySettings)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageOptionsDefaultSnapshotPolicies) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetFileSystem

`func (o *StorageOptionsDefaultSnapshotPolicies) GetFileSystem() StorageOptionsDefaultSnapshotPolicySettings`

GetFileSystem returns the FileSystem field if non-nil, zero value otherwise.

### GetFileSystemOk

`func (o *StorageOptionsDefaultSnapshotPolicies) GetFileSystemOk() (*StorageOptionsDefaultSnapshotPolicySettings, bool)`

GetFileSystemOk returns a tuple with the FileSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystem

`func (o *StorageOptionsDefaultSnapshotPolicies) SetFileSystem(v StorageOptionsDefaultSnapshotPolicySettings)`

SetFileSystem sets FileSystem field to given value.

### HasFileSystem

`func (o *StorageOptionsDefaultSnapshotPolicies) HasFileSystem() bool`

HasFileSystem returns a boolean if a field has been set.

### GetObjectStore

`func (o *StorageOptionsDefaultSnapshotPolicies) GetObjectStore() StorageOptionsDefaultSnapshotPolicySettings`

GetObjectStore returns the ObjectStore field if non-nil, zero value otherwise.

### GetObjectStoreOk

`func (o *StorageOptionsDefaultSnapshotPolicies) GetObjectStoreOk() (*StorageOptionsDefaultSnapshotPolicySettings, bool)`

GetObjectStoreOk returns a tuple with the ObjectStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStore

`func (o *StorageOptionsDefaultSnapshotPolicies) SetObjectStore(v StorageOptionsDefaultSnapshotPolicySettings)`

SetObjectStore sets ObjectStore field to given value.

### HasObjectStore

`func (o *StorageOptionsDefaultSnapshotPolicies) HasObjectStore() bool`

HasObjectStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


