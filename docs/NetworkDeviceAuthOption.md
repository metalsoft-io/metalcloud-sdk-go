# NetworkDeviceAuthOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The authentication method kind. | 
**DeviceAuthProviderId** | Pointer to **int64** | The ID of the DeviceAuthProvider. Required when kind is tacacs. | [optional] 
**VrfName** | Pointer to **string** | VRF name on the device through which TACACS traffic should be routed (e.g. \&quot;mgmt\&quot;). Only meaningful when kind is tacacs. If omitted, the switch service auto-detects from the device state: it queries network-instance&#x3D;mgmt and uses \&quot;mgmt\&quot; if the management interface is a member, otherwise leaves the binding unset (default VRF). | [optional] 

## Methods

### NewNetworkDeviceAuthOption

`func NewNetworkDeviceAuthOption(kind string, ) *NetworkDeviceAuthOption`

NewNetworkDeviceAuthOption instantiates a new NetworkDeviceAuthOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceAuthOptionWithDefaults

`func NewNetworkDeviceAuthOptionWithDefaults() *NetworkDeviceAuthOption`

NewNetworkDeviceAuthOptionWithDefaults instantiates a new NetworkDeviceAuthOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkDeviceAuthOption) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkDeviceAuthOption) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkDeviceAuthOption) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDeviceAuthProviderId

`func (o *NetworkDeviceAuthOption) GetDeviceAuthProviderId() int64`

GetDeviceAuthProviderId returns the DeviceAuthProviderId field if non-nil, zero value otherwise.

### GetDeviceAuthProviderIdOk

`func (o *NetworkDeviceAuthOption) GetDeviceAuthProviderIdOk() (*int64, bool)`

GetDeviceAuthProviderIdOk returns a tuple with the DeviceAuthProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthProviderId

`func (o *NetworkDeviceAuthOption) SetDeviceAuthProviderId(v int64)`

SetDeviceAuthProviderId sets DeviceAuthProviderId field to given value.

### HasDeviceAuthProviderId

`func (o *NetworkDeviceAuthOption) HasDeviceAuthProviderId() bool`

HasDeviceAuthProviderId returns a boolean if a field has been set.

### GetVrfName

`func (o *NetworkDeviceAuthOption) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *NetworkDeviceAuthOption) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *NetworkDeviceAuthOption) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *NetworkDeviceAuthOption) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


