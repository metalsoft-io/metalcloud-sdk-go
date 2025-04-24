# SiteConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**PartialTypeClass**](PartialTypeClass.md) | Location details | [optional] 
**Repo** | Pointer to [**Repo**](Repo.md) | Repository details | [optional] 
**DNSServers** | Pointer to **[]string** | List of DNS Servers | [optional] 
**NTPServers** | Pointer to **[]string** | List of NTP Servers | [optional] 
**NetworkDevicePolicy** | Pointer to [**PartialTypeClass**](PartialTypeClass.md) | Network device policies | [optional] 
**ServerPolicy** | Pointer to [**ServerPolicyUpdate**](ServerPolicyUpdate.md) | Server policies | [optional] 
**ControllerPolicy** | Pointer to [**ControllerPolicy**](ControllerPolicy.md) | Controller policies | [optional] 
**InfrastructurePolicy** | Pointer to [**InfrastructurePolicy**](InfrastructurePolicy.md) | Infrastructure policies | [optional] 

## Methods

### NewSiteConfigUpdate

`func NewSiteConfigUpdate() *SiteConfigUpdate`

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

`func (o *SiteConfigUpdate) GetLocation() PartialTypeClass`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteConfigUpdate) GetLocationOk() (*PartialTypeClass, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteConfigUpdate) SetLocation(v PartialTypeClass)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SiteConfigUpdate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

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

### HasRepo

`func (o *SiteConfigUpdate) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

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

### HasDNSServers

`func (o *SiteConfigUpdate) HasDNSServers() bool`

HasDNSServers returns a boolean if a field has been set.

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

### HasNTPServers

`func (o *SiteConfigUpdate) HasNTPServers() bool`

HasNTPServers returns a boolean if a field has been set.

### GetNetworkDevicePolicy

`func (o *SiteConfigUpdate) GetNetworkDevicePolicy() PartialTypeClass`

GetNetworkDevicePolicy returns the NetworkDevicePolicy field if non-nil, zero value otherwise.

### GetNetworkDevicePolicyOk

`func (o *SiteConfigUpdate) GetNetworkDevicePolicyOk() (*PartialTypeClass, bool)`

GetNetworkDevicePolicyOk returns a tuple with the NetworkDevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePolicy

`func (o *SiteConfigUpdate) SetNetworkDevicePolicy(v PartialTypeClass)`

SetNetworkDevicePolicy sets NetworkDevicePolicy field to given value.

### HasNetworkDevicePolicy

`func (o *SiteConfigUpdate) HasNetworkDevicePolicy() bool`

HasNetworkDevicePolicy returns a boolean if a field has been set.

### GetServerPolicy

`func (o *SiteConfigUpdate) GetServerPolicy() ServerPolicyUpdate`

GetServerPolicy returns the ServerPolicy field if non-nil, zero value otherwise.

### GetServerPolicyOk

`func (o *SiteConfigUpdate) GetServerPolicyOk() (*ServerPolicyUpdate, bool)`

GetServerPolicyOk returns a tuple with the ServerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPolicy

`func (o *SiteConfigUpdate) SetServerPolicy(v ServerPolicyUpdate)`

SetServerPolicy sets ServerPolicy field to given value.

### HasServerPolicy

`func (o *SiteConfigUpdate) HasServerPolicy() bool`

HasServerPolicy returns a boolean if a field has been set.

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

### HasControllerPolicy

`func (o *SiteConfigUpdate) HasControllerPolicy() bool`

HasControllerPolicy returns a boolean if a field has been set.

### GetInfrastructurePolicy

`func (o *SiteConfigUpdate) GetInfrastructurePolicy() InfrastructurePolicy`

GetInfrastructurePolicy returns the InfrastructurePolicy field if non-nil, zero value otherwise.

### GetInfrastructurePolicyOk

`func (o *SiteConfigUpdate) GetInfrastructurePolicyOk() (*InfrastructurePolicy, bool)`

GetInfrastructurePolicyOk returns a tuple with the InfrastructurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructurePolicy

`func (o *SiteConfigUpdate) SetInfrastructurePolicy(v InfrastructurePolicy)`

SetInfrastructurePolicy sets InfrastructurePolicy field to given value.

### HasInfrastructurePolicy

`func (o *SiteConfigUpdate) HasInfrastructurePolicy() bool`

HasInfrastructurePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


