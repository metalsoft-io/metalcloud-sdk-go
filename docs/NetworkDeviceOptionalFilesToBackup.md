# NetworkDeviceOptionalFilesToBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | **map[string]interface{}** | A map of optional files to backup, where the key is a label (e.g. \&quot;default\&quot;, \&quot;version1\&quot;) and the value is an array of file paths to include in the backup | 

## Methods

### NewNetworkDeviceOptionalFilesToBackup

`func NewNetworkDeviceOptionalFilesToBackup(files map[string]interface{}, ) *NetworkDeviceOptionalFilesToBackup`

NewNetworkDeviceOptionalFilesToBackup instantiates a new NetworkDeviceOptionalFilesToBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceOptionalFilesToBackupWithDefaults

`func NewNetworkDeviceOptionalFilesToBackupWithDefaults() *NetworkDeviceOptionalFilesToBackup`

NewNetworkDeviceOptionalFilesToBackupWithDefaults instantiates a new NetworkDeviceOptionalFilesToBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *NetworkDeviceOptionalFilesToBackup) GetFiles() map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *NetworkDeviceOptionalFilesToBackup) GetFilesOk() (*map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *NetworkDeviceOptionalFilesToBackup) SetFiles(v map[string]interface{})`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


