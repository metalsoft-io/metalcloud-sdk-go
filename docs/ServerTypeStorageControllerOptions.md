# ServerTypeStorageControllerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RaidTypesSupported** | **[]string** | The list of supported RAID types for the storage controller. | 
**ControllerModesSupported** | **[]string** | The list of supported controller modes for the storage controller. | 

## Methods

### NewServerTypeStorageControllerOptions

`func NewServerTypeStorageControllerOptions(raidTypesSupported []string, controllerModesSupported []string, ) *ServerTypeStorageControllerOptions`

NewServerTypeStorageControllerOptions instantiates a new ServerTypeStorageControllerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeStorageControllerOptionsWithDefaults

`func NewServerTypeStorageControllerOptionsWithDefaults() *ServerTypeStorageControllerOptions`

NewServerTypeStorageControllerOptionsWithDefaults instantiates a new ServerTypeStorageControllerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaidTypesSupported

`func (o *ServerTypeStorageControllerOptions) GetRaidTypesSupported() []string`

GetRaidTypesSupported returns the RaidTypesSupported field if non-nil, zero value otherwise.

### GetRaidTypesSupportedOk

`func (o *ServerTypeStorageControllerOptions) GetRaidTypesSupportedOk() (*[]string, bool)`

GetRaidTypesSupportedOk returns a tuple with the RaidTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidTypesSupported

`func (o *ServerTypeStorageControllerOptions) SetRaidTypesSupported(v []string)`

SetRaidTypesSupported sets RaidTypesSupported field to given value.


### GetControllerModesSupported

`func (o *ServerTypeStorageControllerOptions) GetControllerModesSupported() []string`

GetControllerModesSupported returns the ControllerModesSupported field if non-nil, zero value otherwise.

### GetControllerModesSupportedOk

`func (o *ServerTypeStorageControllerOptions) GetControllerModesSupportedOk() (*[]string, bool)`

GetControllerModesSupportedOk returns a tuple with the ControllerModesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerModesSupported

`func (o *ServerTypeStorageControllerOptions) SetControllerModesSupported(v []string)`

SetControllerModesSupported sets ControllerModesSupported field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


