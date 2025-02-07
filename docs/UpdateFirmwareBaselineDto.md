# UpdateFirmwareBaselineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareBaselineName** | **string** |  | 
**ServerFirmwareBaselineDescription** | Pointer to **string** |  | [optional] 
**ServerFirmwareBaselineCatalogJson** | Pointer to **string** |  | [optional] 
**ServerFirmwareBaselineLevel** | [**BaselineLevelType**](BaselineLevelType.md) |  | 
**ServerFirmwareBaselineLevelFilterJson** | **string** |  | 

## Methods

### NewUpdateFirmwareBaselineDto

`func NewUpdateFirmwareBaselineDto(serverFirmwareBaselineName string, serverFirmwareBaselineLevel BaselineLevelType, serverFirmwareBaselineLevelFilterJson string, ) *UpdateFirmwareBaselineDto`

NewUpdateFirmwareBaselineDto instantiates a new UpdateFirmwareBaselineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareBaselineDtoWithDefaults

`func NewUpdateFirmwareBaselineDtoWithDefaults() *UpdateFirmwareBaselineDto`

NewUpdateFirmwareBaselineDtoWithDefaults instantiates a new UpdateFirmwareBaselineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBaselineName

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineName() string`

GetServerFirmwareBaselineName returns the ServerFirmwareBaselineName field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineNameOk

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineNameOk() (*string, bool)`

GetServerFirmwareBaselineNameOk returns a tuple with the ServerFirmwareBaselineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineName

`func (o *UpdateFirmwareBaselineDto) SetServerFirmwareBaselineName(v string)`

SetServerFirmwareBaselineName sets ServerFirmwareBaselineName field to given value.


### GetServerFirmwareBaselineDescription

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineDescription() string`

GetServerFirmwareBaselineDescription returns the ServerFirmwareBaselineDescription field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineDescriptionOk

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineDescriptionOk() (*string, bool)`

GetServerFirmwareBaselineDescriptionOk returns a tuple with the ServerFirmwareBaselineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineDescription

`func (o *UpdateFirmwareBaselineDto) SetServerFirmwareBaselineDescription(v string)`

SetServerFirmwareBaselineDescription sets ServerFirmwareBaselineDescription field to given value.

### HasServerFirmwareBaselineDescription

`func (o *UpdateFirmwareBaselineDto) HasServerFirmwareBaselineDescription() bool`

HasServerFirmwareBaselineDescription returns a boolean if a field has been set.

### GetServerFirmwareBaselineCatalogJson

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineCatalogJson() string`

GetServerFirmwareBaselineCatalogJson returns the ServerFirmwareBaselineCatalogJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineCatalogJsonOk

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineCatalogJsonOk() (*string, bool)`

GetServerFirmwareBaselineCatalogJsonOk returns a tuple with the ServerFirmwareBaselineCatalogJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineCatalogJson

`func (o *UpdateFirmwareBaselineDto) SetServerFirmwareBaselineCatalogJson(v string)`

SetServerFirmwareBaselineCatalogJson sets ServerFirmwareBaselineCatalogJson field to given value.

### HasServerFirmwareBaselineCatalogJson

`func (o *UpdateFirmwareBaselineDto) HasServerFirmwareBaselineCatalogJson() bool`

HasServerFirmwareBaselineCatalogJson returns a boolean if a field has been set.

### GetServerFirmwareBaselineLevel

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineLevel() BaselineLevelType`

GetServerFirmwareBaselineLevel returns the ServerFirmwareBaselineLevel field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelOk

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineLevelOk() (*BaselineLevelType, bool)`

GetServerFirmwareBaselineLevelOk returns a tuple with the ServerFirmwareBaselineLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevel

`func (o *UpdateFirmwareBaselineDto) SetServerFirmwareBaselineLevel(v BaselineLevelType)`

SetServerFirmwareBaselineLevel sets ServerFirmwareBaselineLevel field to given value.


### GetServerFirmwareBaselineLevelFilterJson

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineLevelFilterJson() string`

GetServerFirmwareBaselineLevelFilterJson returns the ServerFirmwareBaselineLevelFilterJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelFilterJsonOk

`func (o *UpdateFirmwareBaselineDto) GetServerFirmwareBaselineLevelFilterJsonOk() (*string, bool)`

GetServerFirmwareBaselineLevelFilterJsonOk returns a tuple with the ServerFirmwareBaselineLevelFilterJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevelFilterJson

`func (o *UpdateFirmwareBaselineDto) SetServerFirmwareBaselineLevelFilterJson(v string)`

SetServerFirmwareBaselineLevelFilterJson sets ServerFirmwareBaselineLevelFilterJson field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


