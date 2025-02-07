# SiteConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**Location**](Location.md) | Location details | 
**Repo** | [**Repo**](Repo.md) | Repository details | 
**DNSServers** | **[]string** | List of DNS Servers | 
**NTPServers** | **[]string** | List of NTP Servers | 
**NetworkDevicePolicy** | [**NetworkDevicePolicy**](NetworkDevicePolicy.md) | Network device policies | 
**ServerPolicy** | [**ServerPolicy**](ServerPolicy.md) | Server policies | 
**ControllerPolicy** | [**ControllerPolicy**](ControllerPolicy.md) | Controller policies | 

## Methods

### NewSiteConfigUpdate

`func NewSiteConfigUpdate(location Location, repo Repo, dNSServers []string, nTPServers []string, networkDevicePolicy NetworkDevicePolicy, serverPolicy ServerPolicy, controllerPolicy ControllerPolicy, ) *SiteConfigUpdate`

NewSiteConfigUpdate instantiates a new SiteConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteConfigUpdateWithDefaults

`func NewSiteConfigUpdateWithDefaults() *SiteConfigUpdate`

NewSiteConfigUpdateWithDefaults instantiates a new SiteConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SiteConfigUpdate) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteConfigUpdate) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteConfigUpdate) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetRepo

`func (o *SiteConfigUpdate) GetRepo() Repo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SiteConfigUpdate) GetRepoOk() (*Repo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SiteConfigUpdate) SetRepo(v Repo)`

SetRepo sets Repo field to given value.


### GetDNSServers

`func (o *SiteConfigUpdate) GetDNSServers() []string`

GetDNSServers returns the DNSServers field if non-nil, zero value otherwise.

### GetDNSServersOk

`func (o *SiteConfigUpdate) GetDNSServersOk() (*[]string, bool)`

GetDNSServersOk returns a tuple with the DNSServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSServers

`func (o *SiteConfigUpdate) SetDNSServers(v []string)`

SetDNSServers sets DNSServers field to given value.


### GetNTPServers

`func (o *SiteConfigUpdate) GetNTPServers() []string`

GetNTPServers returns the NTPServers field if non-nil, zero value otherwise.

### GetNTPServersOk

`func (o *SiteConfigUpdate) GetNTPServersOk() (*[]string, bool)`

GetNTPServersOk returns a tuple with the NTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPServers

`func (o *SiteConfigUpdate) SetNTPServers(v []string)`

SetNTPServers sets NTPServers field to given value.


### GetNetworkDevicePolicy

`func (o *SiteConfigUpdate) GetNetworkDevicePolicy() NetworkDevicePolicy`

GetNetworkDevicePolicy returns the NetworkDevicePolicy field if non-nil, zero value otherwise.

### GetNetworkDevicePolicyOk

`func (o *SiteConfigUpdate) GetNetworkDevicePolicyOk() (*NetworkDevicePolicy, bool)`

GetNetworkDevicePolicyOk returns a tuple with the NetworkDevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePolicy

`func (o *SiteConfigUpdate) SetNetworkDevicePolicy(v NetworkDevicePolicy)`

SetNetworkDevicePolicy sets NetworkDevicePolicy field to given value.


### GetServerPolicy

`func (o *SiteConfigUpdate) GetServerPolicy() ServerPolicy`

GetServerPolicy returns the ServerPolicy field if non-nil, zero value otherwise.

### GetServerPolicyOk

`func (o *SiteConfigUpdate) GetServerPolicyOk() (*ServerPolicy, bool)`

GetServerPolicyOk returns a tuple with the ServerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPolicy

`func (o *SiteConfigUpdate) SetServerPolicy(v ServerPolicy)`

SetServerPolicy sets ServerPolicy field to given value.


### GetControllerPolicy

`func (o *SiteConfigUpdate) GetControllerPolicy() ControllerPolicy`

GetControllerPolicy returns the ControllerPolicy field if non-nil, zero value otherwise.

### GetControllerPolicyOk

`func (o *SiteConfigUpdate) GetControllerPolicyOk() (*ControllerPolicy, bool)`

GetControllerPolicyOk returns a tuple with the ControllerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerPolicy

`func (o *SiteConfigUpdate) SetControllerPolicy(v ControllerPolicy)`

SetControllerPolicy sets ControllerPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


