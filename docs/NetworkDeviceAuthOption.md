# NetworkDeviceAuthOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The authentication method kind. | 
**DeviceAuthProviderId** | Pointer to **int32** | The ID of the DeviceAuthProvider. Required when kind is tacacs. | [optional] 

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

`func (o *NetworkDeviceAuthOption) GetDeviceAuthProviderId() int32`

GetDeviceAuthProviderId returns the DeviceAuthProviderId field if non-nil, zero value otherwise.

### GetDeviceAuthProviderIdOk

`func (o *NetworkDeviceAuthOption) GetDeviceAuthProviderIdOk() (*int32, bool)`

GetDeviceAuthProviderIdOk returns a tuple with the DeviceAuthProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthProviderId

`func (o *NetworkDeviceAuthOption) SetDeviceAuthProviderId(v int32)`

SetDeviceAuthProviderId sets DeviceAuthProviderId field to given value.

### HasDeviceAuthProviderId

`func (o *NetworkDeviceAuthOption) HasDeviceAuthProviderId() bool`

HasDeviceAuthProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


