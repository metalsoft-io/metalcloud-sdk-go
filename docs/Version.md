# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The version of the installed MetalSoft software | 
**VersionDate** | Pointer to **time.Time** | The release date of the version | [optional] 
**VersionBuild** | Pointer to **string** | The build number of the version | [optional] 

## Methods

### NewVersion

`func NewVersion(version string, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Version) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Version) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Version) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionDate

`func (o *Version) GetVersionDate() time.Time`

GetVersionDate returns the VersionDate field if non-nil, zero value otherwise.

### GetVersionDateOk

`func (o *Version) GetVersionDateOk() (*time.Time, bool)`

GetVersionDateOk returns a tuple with the VersionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionDate

`func (o *Version) SetVersionDate(v time.Time)`

SetVersionDate sets VersionDate field to given value.

### HasVersionDate

`func (o *Version) HasVersionDate() bool`

HasVersionDate returns a boolean if a field has been set.

### GetVersionBuild

`func (o *Version) GetVersionBuild() string`

GetVersionBuild returns the VersionBuild field if non-nil, zero value otherwise.

### GetVersionBuildOk

`func (o *Version) GetVersionBuildOk() (*string, bool)`

GetVersionBuildOk returns a tuple with the VersionBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionBuild

`func (o *Version) SetVersionBuild(v string)`

SetVersionBuild sets VersionBuild field to given value.

### HasVersionBuild

`func (o *Version) HasVersionBuild() bool`

HasVersionBuild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


