# IServerFirmwareBinaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | Pointer to [**FirmwareBaseline**](FirmwareBaseline.md) | Baseline information (may omit links and createdTimestamp fields) | [optional] 
**ComponentBinaries** | Pointer to [**map[string][]FirmwareBinaryResponse**](array.md) | Mapping of component IDs (as strings) to arrays of firmware binaries | [optional] 

## Methods

### NewIServerFirmwareBinaryResponse

`func NewIServerFirmwareBinaryResponse() *IServerFirmwareBinaryResponse`

NewIServerFirmwareBinaryResponse instantiates a new IServerFirmwareBinaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIServerFirmwareBinaryResponseWithDefaults

`func NewIServerFirmwareBinaryResponseWithDefaults() *IServerFirmwareBinaryResponse`

NewIServerFirmwareBinaryResponseWithDefaults instantiates a new IServerFirmwareBinaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *IServerFirmwareBinaryResponse) GetBaseline() FirmwareBaseline`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *IServerFirmwareBinaryResponse) GetBaselineOk() (*FirmwareBaseline, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *IServerFirmwareBinaryResponse) SetBaseline(v FirmwareBaseline)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *IServerFirmwareBinaryResponse) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetComponentBinaries

`func (o *IServerFirmwareBinaryResponse) GetComponentBinaries() map[string][]FirmwareBinaryResponse`

GetComponentBinaries returns the ComponentBinaries field if non-nil, zero value otherwise.

### GetComponentBinariesOk

`func (o *IServerFirmwareBinaryResponse) GetComponentBinariesOk() (*map[string][]FirmwareBinaryResponse, bool)`

GetComponentBinariesOk returns a tuple with the ComponentBinaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentBinaries

`func (o *IServerFirmwareBinaryResponse) SetComponentBinaries(v map[string][]FirmwareBinaryResponse)`

SetComponentBinaries sets ComponentBinaries field to given value.

### HasComponentBinaries

`func (o *IServerFirmwareBinaryResponse) HasComponentBinaries() bool`

HasComponentBinaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


