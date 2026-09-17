# StorageOptionsDefaultQoSPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realm** | Pointer to [**StorageOptionsDefaultQoSPolicySettings**](StorageOptionsDefaultQoSPolicySettings.md) | Default QoS policy settings for the realm entity (applicable to Pure Flash Array and Flash Blade. 1 Pure FA realm corresponds to 1 MetalSoft infrastructure) | [optional] 
**Svm** | Pointer to [**StorageOptionsDefaultQoSPolicySettings**](StorageOptionsDefaultQoSPolicySettings.md) | Default QoS policy settings for the SVM entity (applicable to NetApp ONTAP. 1 NetApp SVM corresponds to 1 MetalSoft logical network inside an infrastructure. Public logical networks will have 1 SVM for each infrastructure they are in) | [optional] 
**Volume** | Pointer to [**StorageOptionsDefaultQoSPolicySettings**](StorageOptionsDefaultQoSPolicySettings.md) | Default QoS policy settings for the volume entity (applicable to block storages if supported) | [optional] 
**FileSystem** | Pointer to [**StorageOptionsDefaultQoSPolicySettings**](StorageOptionsDefaultQoSPolicySettings.md) | Default QoS policy settings for the file system entity (applicable to file storages if supported) | [optional] 
**ObjectStore** | Pointer to [**StorageOptionsDefaultQoSPolicySettings**](StorageOptionsDefaultQoSPolicySettings.md) | Default QoS policy settings for the object store entity (applicable to object storages if supported) | [optional] 

## Methods

### NewStorageOptionsDefaultQoSPolicies

`func NewStorageOptionsDefaultQoSPolicies() *StorageOptionsDefaultQoSPolicies`

NewStorageOptionsDefaultQoSPolicies instantiates a new StorageOptionsDefaultQoSPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageOptionsDefaultQoSPoliciesWithDefaults

`func NewStorageOptionsDefaultQoSPoliciesWithDefaults() *StorageOptionsDefaultQoSPolicies`

NewStorageOptionsDefaultQoSPoliciesWithDefaults instantiates a new StorageOptionsDefaultQoSPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealm

`func (o *StorageOptionsDefaultQoSPolicies) GetRealm() StorageOptionsDefaultQoSPolicySettings`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *StorageOptionsDefaultQoSPolicies) GetRealmOk() (*StorageOptionsDefaultQoSPolicySettings, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *StorageOptionsDefaultQoSPolicies) SetRealm(v StorageOptionsDefaultQoSPolicySettings)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *StorageOptionsDefaultQoSPolicies) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetSvm

`func (o *StorageOptionsDefaultQoSPolicies) GetSvm() StorageOptionsDefaultQoSPolicySettings`

GetSvm returns the Svm field if non-nil, zero value otherwise.

### GetSvmOk

`func (o *StorageOptionsDefaultQoSPolicies) GetSvmOk() (*StorageOptionsDefaultQoSPolicySettings, bool)`

GetSvmOk returns a tuple with the Svm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvm

`func (o *StorageOptionsDefaultQoSPolicies) SetSvm(v StorageOptionsDefaultQoSPolicySettings)`

SetSvm sets Svm field to given value.

### HasSvm

`func (o *StorageOptionsDefaultQoSPolicies) HasSvm() bool`

HasSvm returns a boolean if a field has been set.

### GetVolume

`func (o *StorageOptionsDefaultQoSPolicies) GetVolume() StorageOptionsDefaultQoSPolicySettings`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageOptionsDefaultQoSPolicies) GetVolumeOk() (*StorageOptionsDefaultQoSPolicySettings, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageOptionsDefaultQoSPolicies) SetVolume(v StorageOptionsDefaultQoSPolicySettings)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageOptionsDefaultQoSPolicies) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetFileSystem

`func (o *StorageOptionsDefaultQoSPolicies) GetFileSystem() StorageOptionsDefaultQoSPolicySettings`

GetFileSystem returns the FileSystem field if non-nil, zero value otherwise.

### GetFileSystemOk

`func (o *StorageOptionsDefaultQoSPolicies) GetFileSystemOk() (*StorageOptionsDefaultQoSPolicySettings, bool)`

GetFileSystemOk returns a tuple with the FileSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystem

`func (o *StorageOptionsDefaultQoSPolicies) SetFileSystem(v StorageOptionsDefaultQoSPolicySettings)`

SetFileSystem sets FileSystem field to given value.

### HasFileSystem

`func (o *StorageOptionsDefaultQoSPolicies) HasFileSystem() bool`

HasFileSystem returns a boolean if a field has been set.

### GetObjectStore

`func (o *StorageOptionsDefaultQoSPolicies) GetObjectStore() StorageOptionsDefaultQoSPolicySettings`

GetObjectStore returns the ObjectStore field if non-nil, zero value otherwise.

### GetObjectStoreOk

`func (o *StorageOptionsDefaultQoSPolicies) GetObjectStoreOk() (*StorageOptionsDefaultQoSPolicySettings, bool)`

GetObjectStoreOk returns a tuple with the ObjectStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStore

`func (o *StorageOptionsDefaultQoSPolicies) SetObjectStore(v StorageOptionsDefaultQoSPolicySettings)`

SetObjectStore sets ObjectStore field to given value.

### HasObjectStore

`func (o *StorageOptionsDefaultQoSPolicies) HasObjectStore() bool`

HasObjectStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


