# NetworkDeviceInterconnectTemplateRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalId** | **float32** | The database id of the local network equipment. | 
**RemoteId** | **float32** | The database id of the remote network equipment. | 
**LocalAsn** | **float32** | The BGP ASN assigned for the local network equipment. | 
**RemoteAsn** | **float32** | The BGP ASN assigned for the remote network equipment. | 
**LocalLoopbackIpv4** | **string** | The loopback IPv4 address of the local network equipment. | 
**RemoteLoopbackIpv4** | **string** | The loopback IPv4 address of the remote network equipment. | 
**LocalVtepIpv4** | **string** | The VTEP IPv4 address of the local network equipment. | 
**RemoteVtepIpv4** | **string** | The VTEP IPv4 address of the remote network equipment. | 
**LocalVtepExternalIpv4** | **string** | The external VTEP IPv4 address of the local network equipment. | 
**RemoteVtepExternalIpv4** | **string** | The external VTEP IPv4 address of the remote network equipment. | 

## Methods

### NewNetworkDeviceInterconnectTemplateRecordSet

`func NewNetworkDeviceInterconnectTemplateRecordSet(localId float32, remoteId float32, localAsn float32, remoteAsn float32, localLoopbackIpv4 string, remoteLoopbackIpv4 string, localVtepIpv4 string, remoteVtepIpv4 string, localVtepExternalIpv4 string, remoteVtepExternalIpv4 string, ) *NetworkDeviceInterconnectTemplateRecordSet`

NewNetworkDeviceInterconnectTemplateRecordSet instantiates a new NetworkDeviceInterconnectTemplateRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceInterconnectTemplateRecordSetWithDefaults

`func NewNetworkDeviceInterconnectTemplateRecordSetWithDefaults() *NetworkDeviceInterconnectTemplateRecordSet`

NewNetworkDeviceInterconnectTemplateRecordSetWithDefaults instantiates a new NetworkDeviceInterconnectTemplateRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalId

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalId() float32`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalIdOk() (*float32, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetLocalId(v float32)`

SetLocalId sets LocalId field to given value.


### GetRemoteId

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteId() float32`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteIdOk() (*float32, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetRemoteId(v float32)`

SetRemoteId sets RemoteId field to given value.


### GetLocalAsn

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalAsn() float32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalAsnOk() (*float32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetLocalAsn(v float32)`

SetLocalAsn sets LocalAsn field to given value.


### GetRemoteAsn

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteAsn() float32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteAsnOk() (*float32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetRemoteAsn(v float32)`

SetRemoteAsn sets RemoteAsn field to given value.


### GetLocalLoopbackIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalLoopbackIpv4() string`

GetLocalLoopbackIpv4 returns the LocalLoopbackIpv4 field if non-nil, zero value otherwise.

### GetLocalLoopbackIpv4Ok

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalLoopbackIpv4Ok() (*string, bool)`

GetLocalLoopbackIpv4Ok returns a tuple with the LocalLoopbackIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLoopbackIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetLocalLoopbackIpv4(v string)`

SetLocalLoopbackIpv4 sets LocalLoopbackIpv4 field to given value.


### GetRemoteLoopbackIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteLoopbackIpv4() string`

GetRemoteLoopbackIpv4 returns the RemoteLoopbackIpv4 field if non-nil, zero value otherwise.

### GetRemoteLoopbackIpv4Ok

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteLoopbackIpv4Ok() (*string, bool)`

GetRemoteLoopbackIpv4Ok returns a tuple with the RemoteLoopbackIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLoopbackIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetRemoteLoopbackIpv4(v string)`

SetRemoteLoopbackIpv4 sets RemoteLoopbackIpv4 field to given value.


### GetLocalVtepIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalVtepIpv4() string`

GetLocalVtepIpv4 returns the LocalVtepIpv4 field if non-nil, zero value otherwise.

### GetLocalVtepIpv4Ok

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalVtepIpv4Ok() (*string, bool)`

GetLocalVtepIpv4Ok returns a tuple with the LocalVtepIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVtepIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetLocalVtepIpv4(v string)`

SetLocalVtepIpv4 sets LocalVtepIpv4 field to given value.


### GetRemoteVtepIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteVtepIpv4() string`

GetRemoteVtepIpv4 returns the RemoteVtepIpv4 field if non-nil, zero value otherwise.

### GetRemoteVtepIpv4Ok

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteVtepIpv4Ok() (*string, bool)`

GetRemoteVtepIpv4Ok returns a tuple with the RemoteVtepIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVtepIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetRemoteVtepIpv4(v string)`

SetRemoteVtepIpv4 sets RemoteVtepIpv4 field to given value.


### GetLocalVtepExternalIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalVtepExternalIpv4() string`

GetLocalVtepExternalIpv4 returns the LocalVtepExternalIpv4 field if non-nil, zero value otherwise.

### GetLocalVtepExternalIpv4Ok

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetLocalVtepExternalIpv4Ok() (*string, bool)`

GetLocalVtepExternalIpv4Ok returns a tuple with the LocalVtepExternalIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVtepExternalIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetLocalVtepExternalIpv4(v string)`

SetLocalVtepExternalIpv4 sets LocalVtepExternalIpv4 field to given value.


### GetRemoteVtepExternalIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteVtepExternalIpv4() string`

GetRemoteVtepExternalIpv4 returns the RemoteVtepExternalIpv4 field if non-nil, zero value otherwise.

### GetRemoteVtepExternalIpv4Ok

`func (o *NetworkDeviceInterconnectTemplateRecordSet) GetRemoteVtepExternalIpv4Ok() (*string, bool)`

GetRemoteVtepExternalIpv4Ok returns a tuple with the RemoteVtepExternalIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVtepExternalIpv4

`func (o *NetworkDeviceInterconnectTemplateRecordSet) SetRemoteVtepExternalIpv4(v string)`

SetRemoteVtepExternalIpv4 sets RemoteVtepExternalIpv4 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


