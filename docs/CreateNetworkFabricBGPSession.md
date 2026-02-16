# CreateNetworkFabricBGPSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpNumbering** | **string** | BGP numbering type for the link | 
**BgpLinkConfiguration** | **string** | BGP link configuration type | 
**LinkId** | Pointer to **float32** | The ID of the link to associate with the BGP session. Either linkId or linkAggregationId should be provided. | [optional] 
**LinkAggregationId** | Pointer to **float32** | The ID of the link aggregation to associate with the BGP session. Either linkId or linkAggregationId should be provided. | [optional] 

## Methods

### NewCreateNetworkFabricBGPSession

`func NewCreateNetworkFabricBGPSession(bgpNumbering string, bgpLinkConfiguration string, ) *CreateNetworkFabricBGPSession`

NewCreateNetworkFabricBGPSession instantiates a new CreateNetworkFabricBGPSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFabricBGPSessionWithDefaults

`func NewCreateNetworkFabricBGPSessionWithDefaults() *CreateNetworkFabricBGPSession`

NewCreateNetworkFabricBGPSessionWithDefaults instantiates a new CreateNetworkFabricBGPSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpNumbering

`func (o *CreateNetworkFabricBGPSession) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *CreateNetworkFabricBGPSession) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *CreateNetworkFabricBGPSession) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *CreateNetworkFabricBGPSession) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *CreateNetworkFabricBGPSession) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *CreateNetworkFabricBGPSession) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetLinkId

`func (o *CreateNetworkFabricBGPSession) GetLinkId() float32`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *CreateNetworkFabricBGPSession) GetLinkIdOk() (*float32, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *CreateNetworkFabricBGPSession) SetLinkId(v float32)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *CreateNetworkFabricBGPSession) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetLinkAggregationId

`func (o *CreateNetworkFabricBGPSession) GetLinkAggregationId() float32`

GetLinkAggregationId returns the LinkAggregationId field if non-nil, zero value otherwise.

### GetLinkAggregationIdOk

`func (o *CreateNetworkFabricBGPSession) GetLinkAggregationIdOk() (*float32, bool)`

GetLinkAggregationIdOk returns a tuple with the LinkAggregationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAggregationId

`func (o *CreateNetworkFabricBGPSession) SetLinkAggregationId(v float32)`

SetLinkAggregationId sets LinkAggregationId field to given value.

### HasLinkAggregationId

`func (o *CreateNetworkFabricBGPSession) HasLinkAggregationId() bool`

HasLinkAggregationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


