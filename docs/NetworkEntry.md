# NetworkEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upload** | Pointer to [**UtilizationData**](UtilizationData.md) |  | [optional] 
**Download** | Pointer to [**UtilizationData**](UtilizationData.md) |  | [optional] 

## Methods

### NewNetworkEntry

`func NewNetworkEntry() *NetworkEntry`

NewNetworkEntry instantiates a new NetworkEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEntryWithDefaults

`func NewNetworkEntryWithDefaults() *NetworkEntry`

NewNetworkEntryWithDefaults instantiates a new NetworkEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpload

`func (o *NetworkEntry) GetUpload() UtilizationData`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *NetworkEntry) GetUploadOk() (*UtilizationData, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *NetworkEntry) SetUpload(v UtilizationData)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *NetworkEntry) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetDownload

`func (o *NetworkEntry) GetDownload() UtilizationData`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *NetworkEntry) GetDownloadOk() (*UtilizationData, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *NetworkEntry) SetDownload(v UtilizationData)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *NetworkEntry) HasDownload() bool`

HasDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


