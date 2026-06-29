# NetworkFabricBGPSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Id** | **int64** | Unique identifier for the network fabric BGP session | 
**NetworkFabricId** | **int64** | Unique identifier for the network fabric | 
**NetworkFabricLinkId** | Pointer to **int64** | Unique identifier for the network fabric link | [optional] 
**NetworkFabricLinkAggregationId** | Pointer to **int64** | Unique identifier for the network fabric link aggregation | [optional] 
**Name** | **string** | Name of the network fabric BGP session | 
**BgpNumbering** | **string** | BGP numbering type for the link | 
**BgpLinkConfiguration** | **string** | BGP link configuration type | 
**Status** | **string** | Status of the network fabric link aggregation | 
**NetworkFabricLink** | Pointer to [**NetworkFabricLink**](NetworkFabricLink.md) | Associated network fabric link | [optional] 
**LinkAggregation** | Pointer to [**NetworkFabricLinkAggregation**](NetworkFabricLinkAggregation.md) | Associated network fabric link aggregation | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkFabricBGPSession

`func NewNetworkFabricBGPSession(createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id int64, networkFabricId int64, name string, bgpNumbering string, bgpLinkConfiguration string, status string, ) *NetworkFabricBGPSession`

NewNetworkFabricBGPSession instantiates a new NetworkFabricBGPSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricBGPSessionWithDefaults

`func NewNetworkFabricBGPSessionWithDefaults() *NetworkFabricBGPSession`

NewNetworkFabricBGPSessionWithDefaults instantiates a new NetworkFabricBGPSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *NetworkFabricBGPSession) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkFabricBGPSession) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkFabricBGPSession) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkFabricBGPSession) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkFabricBGPSession) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkFabricBGPSession) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *NetworkFabricBGPSession) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabricBGPSession) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabricBGPSession) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetId

`func (o *NetworkFabricBGPSession) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkFabricBGPSession) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkFabricBGPSession) SetId(v int64)`

SetId sets Id field to given value.


### GetNetworkFabricId

`func (o *NetworkFabricBGPSession) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *NetworkFabricBGPSession) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *NetworkFabricBGPSession) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.


### GetNetworkFabricLinkId

`func (o *NetworkFabricBGPSession) GetNetworkFabricLinkId() int64`

GetNetworkFabricLinkId returns the NetworkFabricLinkId field if non-nil, zero value otherwise.

### GetNetworkFabricLinkIdOk

`func (o *NetworkFabricBGPSession) GetNetworkFabricLinkIdOk() (*int64, bool)`

GetNetworkFabricLinkIdOk returns a tuple with the NetworkFabricLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinkId

`func (o *NetworkFabricBGPSession) SetNetworkFabricLinkId(v int64)`

SetNetworkFabricLinkId sets NetworkFabricLinkId field to given value.

### HasNetworkFabricLinkId

`func (o *NetworkFabricBGPSession) HasNetworkFabricLinkId() bool`

HasNetworkFabricLinkId returns a boolean if a field has been set.

### GetNetworkFabricLinkAggregationId

`func (o *NetworkFabricBGPSession) GetNetworkFabricLinkAggregationId() int64`

GetNetworkFabricLinkAggregationId returns the NetworkFabricLinkAggregationId field if non-nil, zero value otherwise.

### GetNetworkFabricLinkAggregationIdOk

`func (o *NetworkFabricBGPSession) GetNetworkFabricLinkAggregationIdOk() (*int64, bool)`

GetNetworkFabricLinkAggregationIdOk returns a tuple with the NetworkFabricLinkAggregationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinkAggregationId

`func (o *NetworkFabricBGPSession) SetNetworkFabricLinkAggregationId(v int64)`

SetNetworkFabricLinkAggregationId sets NetworkFabricLinkAggregationId field to given value.

### HasNetworkFabricLinkAggregationId

`func (o *NetworkFabricBGPSession) HasNetworkFabricLinkAggregationId() bool`

HasNetworkFabricLinkAggregationId returns a boolean if a field has been set.

### GetName

`func (o *NetworkFabricBGPSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFabricBGPSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFabricBGPSession) SetName(v string)`

SetName sets Name field to given value.


### GetBgpNumbering

`func (o *NetworkFabricBGPSession) GetBgpNumbering() string`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *NetworkFabricBGPSession) GetBgpNumberingOk() (*string, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *NetworkFabricBGPSession) SetBgpNumbering(v string)`

SetBgpNumbering sets BgpNumbering field to given value.


### GetBgpLinkConfiguration

`func (o *NetworkFabricBGPSession) GetBgpLinkConfiguration() string`

GetBgpLinkConfiguration returns the BgpLinkConfiguration field if non-nil, zero value otherwise.

### GetBgpLinkConfigurationOk

`func (o *NetworkFabricBGPSession) GetBgpLinkConfigurationOk() (*string, bool)`

GetBgpLinkConfigurationOk returns a tuple with the BgpLinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpLinkConfiguration

`func (o *NetworkFabricBGPSession) SetBgpLinkConfiguration(v string)`

SetBgpLinkConfiguration sets BgpLinkConfiguration field to given value.


### GetStatus

`func (o *NetworkFabricBGPSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabricBGPSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabricBGPSession) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNetworkFabricLink

`func (o *NetworkFabricBGPSession) GetNetworkFabricLink() NetworkFabricLink`

GetNetworkFabricLink returns the NetworkFabricLink field if non-nil, zero value otherwise.

### GetNetworkFabricLinkOk

`func (o *NetworkFabricBGPSession) GetNetworkFabricLinkOk() (*NetworkFabricLink, bool)`

GetNetworkFabricLinkOk returns a tuple with the NetworkFabricLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLink

`func (o *NetworkFabricBGPSession) SetNetworkFabricLink(v NetworkFabricLink)`

SetNetworkFabricLink sets NetworkFabricLink field to given value.

### HasNetworkFabricLink

`func (o *NetworkFabricBGPSession) HasNetworkFabricLink() bool`

HasNetworkFabricLink returns a boolean if a field has been set.

### GetLinkAggregation

`func (o *NetworkFabricBGPSession) GetLinkAggregation() NetworkFabricLinkAggregation`

GetLinkAggregation returns the LinkAggregation field if non-nil, zero value otherwise.

### GetLinkAggregationOk

`func (o *NetworkFabricBGPSession) GetLinkAggregationOk() (*NetworkFabricLinkAggregation, bool)`

GetLinkAggregationOk returns a tuple with the LinkAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAggregation

`func (o *NetworkFabricBGPSession) SetLinkAggregation(v NetworkFabricLinkAggregation)`

SetLinkAggregation sets LinkAggregation field to given value.

### HasLinkAggregation

`func (o *NetworkFabricBGPSession) HasLinkAggregation() bool`

HasLinkAggregation returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkFabricBGPSession) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkFabricBGPSession) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkFabricBGPSession) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkFabricBGPSession) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


