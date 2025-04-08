# SiteControllerSeenAliveStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MayBeOffline** | **float32** | Number of site controllers that may be offline | 
**Offline** | **float32** | Number of site controllers that are offline | 
**WasSeenConnectedVeryRecently** | **float32** | Number of site controllers seen connected very recently | 

## Methods

### NewSiteControllerSeenAliveStatus

`func NewSiteControllerSeenAliveStatus(mayBeOffline float32, offline float32, wasSeenConnectedVeryRecently float32, ) *SiteControllerSeenAliveStatus`

NewSiteControllerSeenAliveStatus instantiates a new SiteControllerSeenAliveStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteControllerSeenAliveStatusWithDefaults

`func NewSiteControllerSeenAliveStatusWithDefaults() *SiteControllerSeenAliveStatus`

NewSiteControllerSeenAliveStatusWithDefaults instantiates a new SiteControllerSeenAliveStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMayBeOffline

`func (o *SiteControllerSeenAliveStatus) GetMayBeOffline() float32`

GetMayBeOffline returns the MayBeOffline field if non-nil, zero value otherwise.

### GetMayBeOfflineOk

`func (o *SiteControllerSeenAliveStatus) GetMayBeOfflineOk() (*float32, bool)`

GetMayBeOfflineOk returns a tuple with the MayBeOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayBeOffline

`func (o *SiteControllerSeenAliveStatus) SetMayBeOffline(v float32)`

SetMayBeOffline sets MayBeOffline field to given value.


### GetOffline

`func (o *SiteControllerSeenAliveStatus) GetOffline() float32`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *SiteControllerSeenAliveStatus) GetOfflineOk() (*float32, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *SiteControllerSeenAliveStatus) SetOffline(v float32)`

SetOffline sets Offline field to given value.


### GetWasSeenConnectedVeryRecently

`func (o *SiteControllerSeenAliveStatus) GetWasSeenConnectedVeryRecently() float32`

GetWasSeenConnectedVeryRecently returns the WasSeenConnectedVeryRecently field if non-nil, zero value otherwise.

### GetWasSeenConnectedVeryRecentlyOk

`func (o *SiteControllerSeenAliveStatus) GetWasSeenConnectedVeryRecentlyOk() (*float32, bool)`

GetWasSeenConnectedVeryRecentlyOk returns a tuple with the WasSeenConnectedVeryRecently field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasSeenConnectedVeryRecently

`func (o *SiteControllerSeenAliveStatus) SetWasSeenConnectedVeryRecently(v float32)`

SetWasSeenConnectedVeryRecently sets WasSeenConnectedVeryRecently field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


