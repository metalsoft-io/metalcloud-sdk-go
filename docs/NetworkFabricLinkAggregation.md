# NetworkFabricLinkAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Id** | **float32** | Unique identifier for the network fabric link aggregation | 
**NetworkFabricId** | **float32** | Unique identifier for the network fabric | 
**Name** | **string** | Name of the network fabric link aggregation | 
**Type** | **string** | Type of the network fabric link aggregation | 
**MlagDomainIdentifier** | Pointer to **string** | Identifier for the MLAG domain (applicable only for mlag-peer-link type) | [optional] 
**Status** | **string** | Status of the network fabric link aggregation | 
**NetworkFabricLinks** | Pointer to [**[]NetworkFabricLink**](NetworkFabricLink.md) | List of links associated with the network fabric link aggregation | [optional] 
**NetworkFabricLinkConfigs** | Pointer to [**[]NetworkFabricLinkConfig**](NetworkFabricLinkConfig.md) | History of link configuration changes associated with this link aggregation | [optional] 
**BgpSession** | Pointer to [**NetworkFabricBGPSession**](NetworkFabricBGPSession.md) | Timestamps information | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkFabricLinkAggregation

`func NewNetworkFabricLinkAggregation(createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id float32, networkFabricId float32, name string, type_ string, status string, ) *NetworkFabricLinkAggregation`

NewNetworkFabricLinkAggregation instantiates a new NetworkFabricLinkAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricLinkAggregationWithDefaults

`func NewNetworkFabricLinkAggregationWithDefaults() *NetworkFabricLinkAggregation`

NewNetworkFabricLinkAggregationWithDefaults instantiates a new NetworkFabricLinkAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *NetworkFabricLinkAggregation) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkFabricLinkAggregation) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkFabricLinkAggregation) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkFabricLinkAggregation) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkFabricLinkAggregation) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkFabricLinkAggregation) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *NetworkFabricLinkAggregation) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabricLinkAggregation) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabricLinkAggregation) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetId

`func (o *NetworkFabricLinkAggregation) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkFabricLinkAggregation) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkFabricLinkAggregation) SetId(v float32)`

SetId sets Id field to given value.


### GetNetworkFabricId

`func (o *NetworkFabricLinkAggregation) GetNetworkFabricId() float32`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *NetworkFabricLinkAggregation) GetNetworkFabricIdOk() (*float32, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *NetworkFabricLinkAggregation) SetNetworkFabricId(v float32)`

SetNetworkFabricId sets NetworkFabricId field to given value.


### GetName

`func (o *NetworkFabricLinkAggregation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFabricLinkAggregation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFabricLinkAggregation) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NetworkFabricLinkAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkFabricLinkAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkFabricLinkAggregation) SetType(v string)`

SetType sets Type field to given value.


### GetMlagDomainIdentifier

`func (o *NetworkFabricLinkAggregation) GetMlagDomainIdentifier() string`

GetMlagDomainIdentifier returns the MlagDomainIdentifier field if non-nil, zero value otherwise.

### GetMlagDomainIdentifierOk

`func (o *NetworkFabricLinkAggregation) GetMlagDomainIdentifierOk() (*string, bool)`

GetMlagDomainIdentifierOk returns a tuple with the MlagDomainIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagDomainIdentifier

`func (o *NetworkFabricLinkAggregation) SetMlagDomainIdentifier(v string)`

SetMlagDomainIdentifier sets MlagDomainIdentifier field to given value.

### HasMlagDomainIdentifier

`func (o *NetworkFabricLinkAggregation) HasMlagDomainIdentifier() bool`

HasMlagDomainIdentifier returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkFabricLinkAggregation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabricLinkAggregation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabricLinkAggregation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNetworkFabricLinks

`func (o *NetworkFabricLinkAggregation) GetNetworkFabricLinks() []NetworkFabricLink`

GetNetworkFabricLinks returns the NetworkFabricLinks field if non-nil, zero value otherwise.

### GetNetworkFabricLinksOk

`func (o *NetworkFabricLinkAggregation) GetNetworkFabricLinksOk() (*[]NetworkFabricLink, bool)`

GetNetworkFabricLinksOk returns a tuple with the NetworkFabricLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinks

`func (o *NetworkFabricLinkAggregation) SetNetworkFabricLinks(v []NetworkFabricLink)`

SetNetworkFabricLinks sets NetworkFabricLinks field to given value.

### HasNetworkFabricLinks

`func (o *NetworkFabricLinkAggregation) HasNetworkFabricLinks() bool`

HasNetworkFabricLinks returns a boolean if a field has been set.

### GetNetworkFabricLinkConfigs

`func (o *NetworkFabricLinkAggregation) GetNetworkFabricLinkConfigs() []NetworkFabricLinkConfig`

GetNetworkFabricLinkConfigs returns the NetworkFabricLinkConfigs field if non-nil, zero value otherwise.

### GetNetworkFabricLinkConfigsOk

`func (o *NetworkFabricLinkAggregation) GetNetworkFabricLinkConfigsOk() (*[]NetworkFabricLinkConfig, bool)`

GetNetworkFabricLinkConfigsOk returns a tuple with the NetworkFabricLinkConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinkConfigs

`func (o *NetworkFabricLinkAggregation) SetNetworkFabricLinkConfigs(v []NetworkFabricLinkConfig)`

SetNetworkFabricLinkConfigs sets NetworkFabricLinkConfigs field to given value.

### HasNetworkFabricLinkConfigs

`func (o *NetworkFabricLinkAggregation) HasNetworkFabricLinkConfigs() bool`

HasNetworkFabricLinkConfigs returns a boolean if a field has been set.

### GetBgpSession

`func (o *NetworkFabricLinkAggregation) GetBgpSession() NetworkFabricBGPSession`

GetBgpSession returns the BgpSession field if non-nil, zero value otherwise.

### GetBgpSessionOk

`func (o *NetworkFabricLinkAggregation) GetBgpSessionOk() (*NetworkFabricBGPSession, bool)`

GetBgpSessionOk returns a tuple with the BgpSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpSession

`func (o *NetworkFabricLinkAggregation) SetBgpSession(v NetworkFabricBGPSession)`

SetBgpSession sets BgpSession field to given value.

### HasBgpSession

`func (o *NetworkFabricLinkAggregation) HasBgpSession() bool`

HasBgpSession returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkFabricLinkAggregation) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkFabricLinkAggregation) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkFabricLinkAggregation) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkFabricLinkAggregation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


