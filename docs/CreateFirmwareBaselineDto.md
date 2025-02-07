# CreateFirmwareBaselineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareBaselineName** | **string** |  | 
**ServerFirmwareBaselineDescription** | Pointer to **string** |  | [optional] 
**ServerFirmwareBaselineCatalogJson** | Pointer to **string** |  | [optional] 
**ServerFirmwareBaselineLevel** | [**BaselineLevelType**](BaselineLevelType.md) |  | 
**ServerFirmwareBaselineLevelFilterJson** | **string** |  | 

## Methods

### NewCreateFirmwareBaselineDto

`func NewCreateFirmwareBaselineDto(serverFirmwareBaselineName string, serverFirmwareBaselineLevel BaselineLevelType, serverFirmwareBaselineLevelFilterJson string, ) *CreateFirmwareBaselineDto`

NewCreateFirmwareBaselineDto instantiates a new CreateFirmwareBaselineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareBaselineDtoWithDefaults

`func NewCreateFirmwareBaselineDtoWithDefaults() *CreateFirmwareBaselineDto`

NewCreateFirmwareBaselineDtoWithDefaults instantiates a new CreateFirmwareBaselineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBaselineName

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineName() string`

GetServerFirmwareBaselineName returns the ServerFirmwareBaselineName field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineNameOk

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineNameOk() (*string, bool)`

GetServerFirmwareBaselineNameOk returns a tuple with the ServerFirmwareBaselineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineName

`func (o *CreateFirmwareBaselineDto) SetServerFirmwareBaselineName(v string)`

SetServerFirmwareBaselineName sets ServerFirmwareBaselineName field to given value.


### GetServerFirmwareBaselineDescription

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineDescription() string`

GetServerFirmwareBaselineDescription returns the ServerFirmwareBaselineDescription field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineDescriptionOk

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineDescriptionOk() (*string, bool)`

GetServerFirmwareBaselineDescriptionOk returns a tuple with the ServerFirmwareBaselineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineDescription

`func (o *CreateFirmwareBaselineDto) SetServerFirmwareBaselineDescription(v string)`

SetServerFirmwareBaselineDescription sets ServerFirmwareBaselineDescription field to given value.

### HasServerFirmwareBaselineDescription

`func (o *CreateFirmwareBaselineDto) HasServerFirmwareBaselineDescription() bool`

HasServerFirmwareBaselineDescription returns a boolean if a field has been set.

### GetServerFirmwareBaselineCatalogJson

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineCatalogJson() string`

GetServerFirmwareBaselineCatalogJson returns the ServerFirmwareBaselineCatalogJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineCatalogJsonOk

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineCatalogJsonOk() (*string, bool)`

GetServerFirmwareBaselineCatalogJsonOk returns a tuple with the ServerFirmwareBaselineCatalogJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineCatalogJson

`func (o *CreateFirmwareBaselineDto) SetServerFirmwareBaselineCatalogJson(v string)`

SetServerFirmwareBaselineCatalogJson sets ServerFirmwareBaselineCatalogJson field to given value.

### HasServerFirmwareBaselineCatalogJson

`func (o *CreateFirmwareBaselineDto) HasServerFirmwareBaselineCatalogJson() bool`

HasServerFirmwareBaselineCatalogJson returns a boolean if a field has been set.

### GetServerFirmwareBaselineLevel

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineLevel() BaselineLevelType`

GetServerFirmwareBaselineLevel returns the ServerFirmwareBaselineLevel field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelOk

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineLevelOk() (*BaselineLevelType, bool)`

GetServerFirmwareBaselineLevelOk returns a tuple with the ServerFirmwareBaselineLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevel

`func (o *CreateFirmwareBaselineDto) SetServerFirmwareBaselineLevel(v BaselineLevelType)`

SetServerFirmwareBaselineLevel sets ServerFirmwareBaselineLevel field to given value.


### GetServerFirmwareBaselineLevelFilterJson

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineLevelFilterJson() string`

GetServerFirmwareBaselineLevelFilterJson returns the ServerFirmwareBaselineLevelFilterJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelFilterJsonOk

`func (o *CreateFirmwareBaselineDto) GetServerFirmwareBaselineLevelFilterJsonOk() (*string, bool)`

GetServerFirmwareBaselineLevelFilterJsonOk returns a tuple with the ServerFirmwareBaselineLevelFilterJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevelFilterJson

`func (o *CreateFirmwareBaselineDto) SetServerFirmwareBaselineLevelFilterJson(v string)`

SetServerFirmwareBaselineLevelFilterJson sets ServerFirmwareBaselineLevelFilterJson field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


