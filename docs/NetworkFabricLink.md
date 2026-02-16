# NetworkFabricLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Id** | **float32** | Unique identifier for the network fabric link | 
**NetworkFabricId** | **float32** | Unique identifier for the network fabric | 
**NetworkFabricLinkAggregationId** | Pointer to **float32** | Unique identifier for the network fabric link aggregation | [optional] 
**NetworkDeviceAInterfaceId** | **float32** | Unique identifier for the network device A interface | 
**NetworkDeviceBInterfaceId** | **float32** | Unique identifier for the network device B interface | 
**LinkType** | **string** | Type of the network fabric link | 
**Status** | **string** | Status of the network fabric link | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the network fabric link | [optional] 
**Config** | Pointer to [**NetworkFabricLinkConfig**](NetworkFabricLinkConfig.md) | Configuration of the network fabric link | [optional] 
**LinkAggregation** | Pointer to [**NetworkFabricLinkAggregation**](NetworkFabricLinkAggregation.md) | Network Fabric Link Aggregation associated with this link | [optional] 
**BgpSession** | Pointer to [**NetworkFabricBGPSession**](NetworkFabricBGPSession.md) | Network Fabric BGP Session associated with this link | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkFabricLink

`func NewNetworkFabricLink(createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id float32, networkFabricId float32, networkDeviceAInterfaceId float32, networkDeviceBInterfaceId float32, linkType string, status string, ) *NetworkFabricLink`

NewNetworkFabricLink instantiates a new NetworkFabricLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricLinkWithDefaults

`func NewNetworkFabricLinkWithDefaults() *NetworkFabricLink`

NewNetworkFabricLinkWithDefaults instantiates a new NetworkFabricLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *NetworkFabricLink) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkFabricLink) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkFabricLink) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkFabricLink) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkFabricLink) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkFabricLink) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *NetworkFabricLink) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabricLink) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabricLink) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetId

`func (o *NetworkFabricLink) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkFabricLink) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkFabricLink) SetId(v float32)`

SetId sets Id field to given value.


### GetNetworkFabricId

`func (o *NetworkFabricLink) GetNetworkFabricId() float32`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *NetworkFabricLink) GetNetworkFabricIdOk() (*float32, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *NetworkFabricLink) SetNetworkFabricId(v float32)`

SetNetworkFabricId sets NetworkFabricId field to given value.


### GetNetworkFabricLinkAggregationId

`func (o *NetworkFabricLink) GetNetworkFabricLinkAggregationId() float32`

GetNetworkFabricLinkAggregationId returns the NetworkFabricLinkAggregationId field if non-nil, zero value otherwise.

### GetNetworkFabricLinkAggregationIdOk

`func (o *NetworkFabricLink) GetNetworkFabricLinkAggregationIdOk() (*float32, bool)`

GetNetworkFabricLinkAggregationIdOk returns a tuple with the NetworkFabricLinkAggregationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinkAggregationId

`func (o *NetworkFabricLink) SetNetworkFabricLinkAggregationId(v float32)`

SetNetworkFabricLinkAggregationId sets NetworkFabricLinkAggregationId field to given value.

### HasNetworkFabricLinkAggregationId

`func (o *NetworkFabricLink) HasNetworkFabricLinkAggregationId() bool`

HasNetworkFabricLinkAggregationId returns a boolean if a field has been set.

### GetNetworkDeviceAInterfaceId

`func (o *NetworkFabricLink) GetNetworkDeviceAInterfaceId() float32`

GetNetworkDeviceAInterfaceId returns the NetworkDeviceAInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceAInterfaceIdOk

`func (o *NetworkFabricLink) GetNetworkDeviceAInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceAInterfaceIdOk returns a tuple with the NetworkDeviceAInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceAInterfaceId

`func (o *NetworkFabricLink) SetNetworkDeviceAInterfaceId(v float32)`

SetNetworkDeviceAInterfaceId sets NetworkDeviceAInterfaceId field to given value.


### GetNetworkDeviceBInterfaceId

`func (o *NetworkFabricLink) GetNetworkDeviceBInterfaceId() float32`

GetNetworkDeviceBInterfaceId returns the NetworkDeviceBInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceBInterfaceIdOk

`func (o *NetworkFabricLink) GetNetworkDeviceBInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceBInterfaceIdOk returns a tuple with the NetworkDeviceBInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceBInterfaceId

`func (o *NetworkFabricLink) SetNetworkDeviceBInterfaceId(v float32)`

SetNetworkDeviceBInterfaceId sets NetworkDeviceBInterfaceId field to given value.


### GetLinkType

`func (o *NetworkFabricLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *NetworkFabricLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *NetworkFabricLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.


### GetStatus

`func (o *NetworkFabricLink) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabricLink) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabricLink) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomVariables

`func (o *NetworkFabricLink) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *NetworkFabricLink) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *NetworkFabricLink) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *NetworkFabricLink) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkFabricLink) GetConfig() NetworkFabricLinkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkFabricLink) GetConfigOk() (*NetworkFabricLinkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkFabricLink) SetConfig(v NetworkFabricLinkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkFabricLink) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinkAggregation

`func (o *NetworkFabricLink) GetLinkAggregation() NetworkFabricLinkAggregation`

GetLinkAggregation returns the LinkAggregation field if non-nil, zero value otherwise.

### GetLinkAggregationOk

`func (o *NetworkFabricLink) GetLinkAggregationOk() (*NetworkFabricLinkAggregation, bool)`

GetLinkAggregationOk returns a tuple with the LinkAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAggregation

`func (o *NetworkFabricLink) SetLinkAggregation(v NetworkFabricLinkAggregation)`

SetLinkAggregation sets LinkAggregation field to given value.

### HasLinkAggregation

`func (o *NetworkFabricLink) HasLinkAggregation() bool`

HasLinkAggregation returns a boolean if a field has been set.

### GetBgpSession

`func (o *NetworkFabricLink) GetBgpSession() NetworkFabricBGPSession`

GetBgpSession returns the BgpSession field if non-nil, zero value otherwise.

### GetBgpSessionOk

`func (o *NetworkFabricLink) GetBgpSessionOk() (*NetworkFabricBGPSession, bool)`

GetBgpSessionOk returns a tuple with the BgpSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpSession

`func (o *NetworkFabricLink) SetBgpSession(v NetworkFabricBGPSession)`

SetBgpSession sets BgpSession field to given value.

### HasBgpSession

`func (o *NetworkFabricLink) HasBgpSession() bool`

HasBgpSession returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkFabricLink) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkFabricLink) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkFabricLink) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkFabricLink) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


