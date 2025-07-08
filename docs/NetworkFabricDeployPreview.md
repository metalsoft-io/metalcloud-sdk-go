# NetworkFabricDeployPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **float32** | Unique identifier for the network device. | 
**NetworkFabricLinkId** | **float32** | Unique identifier for the network fabric link | 
**BgpTemplateId** | **float32** | Unique identifier for the BGP template | 
**PreparationPreview** | Pointer to **string** | Preview of the BGP preparation for the network device encoded in base64 format | [optional] 
**ConfigurationPreview** | Pointer to **string** | Preview of the BGP configuration for the network device encoded in base64 format | [optional] 

## Methods

### NewNetworkFabricDeployPreview

`func NewNetworkFabricDeployPreview(networkDeviceId float32, networkFabricLinkId float32, bgpTemplateId float32, ) *NetworkFabricDeployPreview`

NewNetworkFabricDeployPreview instantiates a new NetworkFabricDeployPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricDeployPreviewWithDefaults

`func NewNetworkFabricDeployPreviewWithDefaults() *NetworkFabricDeployPreview`

NewNetworkFabricDeployPreviewWithDefaults instantiates a new NetworkFabricDeployPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *NetworkFabricDeployPreview) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkFabricDeployPreview) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkFabricDeployPreview) SetNetworkDeviceId(v float32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetNetworkFabricLinkId

`func (o *NetworkFabricDeployPreview) GetNetworkFabricLinkId() float32`

GetNetworkFabricLinkId returns the NetworkFabricLinkId field if non-nil, zero value otherwise.

### GetNetworkFabricLinkIdOk

`func (o *NetworkFabricDeployPreview) GetNetworkFabricLinkIdOk() (*float32, bool)`

GetNetworkFabricLinkIdOk returns a tuple with the NetworkFabricLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinkId

`func (o *NetworkFabricDeployPreview) SetNetworkFabricLinkId(v float32)`

SetNetworkFabricLinkId sets NetworkFabricLinkId field to given value.


### GetBgpTemplateId

`func (o *NetworkFabricDeployPreview) GetBgpTemplateId() float32`

GetBgpTemplateId returns the BgpTemplateId field if non-nil, zero value otherwise.

### GetBgpTemplateIdOk

`func (o *NetworkFabricDeployPreview) GetBgpTemplateIdOk() (*float32, bool)`

GetBgpTemplateIdOk returns a tuple with the BgpTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpTemplateId

`func (o *NetworkFabricDeployPreview) SetBgpTemplateId(v float32)`

SetBgpTemplateId sets BgpTemplateId field to given value.


### GetPreparationPreview

`func (o *NetworkFabricDeployPreview) GetPreparationPreview() string`

GetPreparationPreview returns the PreparationPreview field if non-nil, zero value otherwise.

### GetPreparationPreviewOk

`func (o *NetworkFabricDeployPreview) GetPreparationPreviewOk() (*string, bool)`

GetPreparationPreviewOk returns a tuple with the PreparationPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationPreview

`func (o *NetworkFabricDeployPreview) SetPreparationPreview(v string)`

SetPreparationPreview sets PreparationPreview field to given value.

### HasPreparationPreview

`func (o *NetworkFabricDeployPreview) HasPreparationPreview() bool`

HasPreparationPreview returns a boolean if a field has been set.

### GetConfigurationPreview

`func (o *NetworkFabricDeployPreview) GetConfigurationPreview() string`

GetConfigurationPreview returns the ConfigurationPreview field if non-nil, zero value otherwise.

### GetConfigurationPreviewOk

`func (o *NetworkFabricDeployPreview) GetConfigurationPreviewOk() (*string, bool)`

GetConfigurationPreviewOk returns a tuple with the ConfigurationPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPreview

`func (o *NetworkFabricDeployPreview) SetConfigurationPreview(v string)`

SetConfigurationPreview sets ConfigurationPreview field to given value.

### HasConfigurationPreview

`func (o *NetworkFabricDeployPreview) HasConfigurationPreview() bool`

HasConfigurationPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


