# NetworkDeviceDriverCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | **string** | The network device driver. | 
**ActivationActions** | [**[]NetworkDeviceActivationAction**](NetworkDeviceActivationAction.md) | Onboarding actions that are valid for this driver. | 

## Methods

### NewNetworkDeviceDriverCapabilities

`func NewNetworkDeviceDriverCapabilities(driver string, activationActions []NetworkDeviceActivationAction, ) *NetworkDeviceDriverCapabilities`

NewNetworkDeviceDriverCapabilities instantiates a new NetworkDeviceDriverCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDriverCapabilitiesWithDefaults

`func NewNetworkDeviceDriverCapabilitiesWithDefaults() *NetworkDeviceDriverCapabilities`

NewNetworkDeviceDriverCapabilitiesWithDefaults instantiates a new NetworkDeviceDriverCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriver

`func (o *NetworkDeviceDriverCapabilities) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *NetworkDeviceDriverCapabilities) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *NetworkDeviceDriverCapabilities) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetActivationActions

`func (o *NetworkDeviceDriverCapabilities) GetActivationActions() []NetworkDeviceActivationAction`

GetActivationActions returns the ActivationActions field if non-nil, zero value otherwise.

### GetActivationActionsOk

`func (o *NetworkDeviceDriverCapabilities) GetActivationActionsOk() (*[]NetworkDeviceActivationAction, bool)`

GetActivationActionsOk returns a tuple with the ActivationActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationActions

`func (o *NetworkDeviceDriverCapabilities) SetActivationActions(v []NetworkDeviceActivationAction)`

SetActivationActions sets ActivationActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


