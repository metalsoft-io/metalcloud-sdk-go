# NetworkFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **int32** | The ID of the site where the entity is located. | [optional] 
**Name** | **string** | The network fabric name | 
**Description** | Pointer to **string** | Network fabric description | [optional] 
**FabricConfiguration** | [**NetworkFabricFabricConfiguration**](NetworkFabricFabricConfiguration.md) |  | 
**Revision** | **string** | Revision number of the entity | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **string** | The network fabric ID. | 
**Status** | Pointer to [**FabricStatus**](FabricStatus.md) | The status of the network fabric, by default it is in draft mode. | [optional] 
**DeployId** | Pointer to **int32** | The deploy ID of the network fabric, if it is being deployed. | [optional] 
**DeployPreview** | Pointer to [**[]NetworkFabricDeployPreview**](NetworkFabricDeployPreview.md) | The deploy preview for the network fabric, if it is being deployed. | [optional] 
**NetworkEquipment** | Pointer to [**[]NetworkDevice**](NetworkDevice.md) | The network equipments in the fabric | [optional] 

## Methods

### NewNetworkFabric

`func NewNetworkFabric(name string, fabricConfiguration NetworkFabricFabricConfiguration, revision string, createdTimestamp time.Time, updatedTimestamp time.Time, id string, ) *NetworkFabric`

NewNetworkFabric instantiates a new NetworkFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricWithDefaults

`func NewNetworkFabricWithDefaults() *NetworkFabric`

NewNetworkFabricWithDefaults instantiates a new NetworkFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *NetworkFabric) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *NetworkFabric) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *NetworkFabric) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *NetworkFabric) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

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

`func (o *NetworkFabric) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabric) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabric) SetRevision(v string)`

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


### GetStatus

`func (o *NetworkFabric) GetStatus() FabricStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabric) GetStatusOk() (*FabricStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabric) SetStatus(v FabricStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkFabric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployId

`func (o *NetworkFabric) GetDeployId() int32`

GetDeployId returns the DeployId field if non-nil, zero value otherwise.

### GetDeployIdOk

`func (o *NetworkFabric) GetDeployIdOk() (*int32, bool)`

GetDeployIdOk returns a tuple with the DeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployId

`func (o *NetworkFabric) SetDeployId(v int32)`

SetDeployId sets DeployId field to given value.

### HasDeployId

`func (o *NetworkFabric) HasDeployId() bool`

HasDeployId returns a boolean if a field has been set.

### GetDeployPreview

`func (o *NetworkFabric) GetDeployPreview() []NetworkFabricDeployPreview`

GetDeployPreview returns the DeployPreview field if non-nil, zero value otherwise.

### GetDeployPreviewOk

`func (o *NetworkFabric) GetDeployPreviewOk() (*[]NetworkFabricDeployPreview, bool)`

GetDeployPreviewOk returns a tuple with the DeployPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployPreview

`func (o *NetworkFabric) SetDeployPreview(v []NetworkFabricDeployPreview)`

SetDeployPreview sets DeployPreview field to given value.

### HasDeployPreview

`func (o *NetworkFabric) HasDeployPreview() bool`

HasDeployPreview returns a boolean if a field has been set.

### GetNetworkEquipment

`func (o *NetworkFabric) GetNetworkEquipment() []NetworkDevice`

GetNetworkEquipment returns the NetworkEquipment field if non-nil, zero value otherwise.

### GetNetworkEquipmentOk

`func (o *NetworkFabric) GetNetworkEquipmentOk() (*[]NetworkDevice, bool)`

GetNetworkEquipmentOk returns a tuple with the NetworkEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipment

`func (o *NetworkFabric) SetNetworkEquipment(v []NetworkDevice)`

SetNetworkEquipment sets NetworkEquipment field to given value.

### HasNetworkEquipment

`func (o *NetworkFabric) HasNetworkEquipment() bool`

HasNetworkEquipment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


