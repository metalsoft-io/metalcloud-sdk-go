# NetworkFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The network fabric name | 
**Description** | Pointer to **string** | Network fabric description | [optional] 
**FabricType** | **string** | The type of network fabric | 
**FabricConfiguration** | [**NetworkFabricFabricConfiguration**](NetworkFabricFabricConfiguration.md) |  | 
**Revision** | **int32** | Revision number of the entity | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **string** | The network fabric ID. | 

## Methods

### NewNetworkFabric

`func NewNetworkFabric(name string, fabricType string, fabricConfiguration NetworkFabricFabricConfiguration, revision int32, createdTimestamp time.Time, updatedTimestamp time.Time, id string, ) *NetworkFabric`

NewNetworkFabric instantiates a new NetworkFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricWithDefaults

`func NewNetworkFabricWithDefaults() *NetworkFabric`

NewNetworkFabricWithDefaults instantiates a new NetworkFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkFabric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFabric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFabric) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NetworkFabric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkFabric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkFabric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkFabric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFabricType

`func (o *NetworkFabric) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *NetworkFabric) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *NetworkFabric) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.


### GetFabricConfiguration

`func (o *NetworkFabric) GetFabricConfiguration() NetworkFabricFabricConfiguration`

GetFabricConfiguration returns the FabricConfiguration field if non-nil, zero value otherwise.

### GetFabricConfigurationOk

`func (o *NetworkFabric) GetFabricConfigurationOk() (*NetworkFabricFabricConfiguration, bool)`

GetFabricConfigurationOk returns a tuple with the FabricConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricConfiguration

`func (o *NetworkFabric) SetFabricConfiguration(v NetworkFabricFabricConfiguration)`

SetFabricConfiguration sets FabricConfiguration field to given value.


### GetRevision

`func (o *NetworkFabric) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabric) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabric) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetCreatedTimestamp

`func (o *NetworkFabric) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkFabric) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkFabric) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkFabric) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkFabric) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkFabric) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *NetworkFabric) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkFabric) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkFabric) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkFabric) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *NetworkFabric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkFabric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkFabric) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


