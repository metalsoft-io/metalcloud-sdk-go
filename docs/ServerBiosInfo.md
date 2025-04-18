# ServerBiosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | Pointer to **string** | The BIOS vendor of the server. | [optional] 
**Version** | Pointer to **string** | The BIOS version of the server. | [optional] 

## Methods

### NewServerBiosInfo

`func NewServerBiosInfo() *ServerBiosInfo`

NewServerBiosInfo instantiates a new ServerBiosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBiosInfoWithDefaults

`func NewServerBiosInfoWithDefaults() *ServerBiosInfo`

NewServerBiosInfoWithDefaults instantiates a new ServerBiosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *ServerBiosInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ServerBiosInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ServerBiosInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ServerBiosInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *ServerBiosInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerBiosInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerBiosInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServerBiosInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


