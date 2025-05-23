# SiteConfigVariables

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
**InfrastructurePolicy** | [**InfrastructurePolicy**](InfrastructurePolicy.md) | Infrastructure policies | 

## Methods

### NewSiteConfigVariables

`func NewSiteConfigVariables(location Location, repo Repo, dNSServers []string, nTPServers []string, networkDevicePolicy NetworkDevicePolicy, serverPolicy ServerPolicy, controllerPolicy ControllerPolicy, infrastructurePolicy InfrastructurePolicy, ) *SiteConfigVariables`

NewSiteConfigVariables instantiates a new SiteConfigVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteConfigVariablesWithDefaults

`func NewSiteConfigVariablesWithDefaults() *SiteConfigVariables`

NewSiteConfigVariablesWithDefaults instantiates a new SiteConfigVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SiteConfigVariables) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteConfigVariables) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteConfigVariables) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetRepo

`func (o *SiteConfigVariables) GetRepo() Repo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SiteConfigVariables) GetRepoOk() (*Repo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SiteConfigVariables) SetRepo(v Repo)`

SetRepo sets Repo field to given value.


### GetDNSServers

`func (o *SiteConfigVariables) GetDNSServers() []string`

GetDNSServers returns the DNSServers field if non-nil, zero value otherwise.

### GetDNSServersOk

`func (o *SiteConfigVariables) GetDNSServersOk() (*[]string, bool)`

GetDNSServersOk returns a tuple with the DNSServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSServers

`func (o *SiteConfigVariables) SetDNSServers(v []string)`

SetDNSServers sets DNSServers field to given value.


### GetNTPServers

`func (o *SiteConfigVariables) GetNTPServers() []string`

GetNTPServers returns the NTPServers field if non-nil, zero value otherwise.

### GetNTPServersOk

`func (o *SiteConfigVariables) GetNTPServersOk() (*[]string, bool)`

GetNTPServersOk returns a tuple with the NTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPServers

`func (o *SiteConfigVariables) SetNTPServers(v []string)`

SetNTPServers sets NTPServers field to given value.


### GetNetworkDevicePolicy

`func (o *SiteConfigVariables) GetNetworkDevicePolicy() NetworkDevicePolicy`

GetNetworkDevicePolicy returns the NetworkDevicePolicy field if non-nil, zero value otherwise.

### GetNetworkDevicePolicyOk

`func (o *SiteConfigVariables) GetNetworkDevicePolicyOk() (*NetworkDevicePolicy, bool)`

GetNetworkDevicePolicyOk returns a tuple with the NetworkDevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePolicy

`func (o *SiteConfigVariables) SetNetworkDevicePolicy(v NetworkDevicePolicy)`

SetNetworkDevicePolicy sets NetworkDevicePolicy field to given value.


### GetServerPolicy

`func (o *SiteConfigVariables) GetServerPolicy() ServerPolicy`

GetServerPolicy returns the ServerPolicy field if non-nil, zero value otherwise.

### GetServerPolicyOk

`func (o *SiteConfigVariables) GetServerPolicyOk() (*ServerPolicy, bool)`

GetServerPolicyOk returns a tuple with the ServerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPolicy

`func (o *SiteConfigVariables) SetServerPolicy(v ServerPolicy)`

SetServerPolicy sets ServerPolicy field to given value.


### GetControllerPolicy

`func (o *SiteConfigVariables) GetControllerPolicy() ControllerPolicy`

GetControllerPolicy returns the ControllerPolicy field if non-nil, zero value otherwise.

### GetControllerPolicyOk

`func (o *SiteConfigVariables) GetControllerPolicyOk() (*ControllerPolicy, bool)`

GetControllerPolicyOk returns a tuple with the ControllerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerPolicy

`func (o *SiteConfigVariables) SetControllerPolicy(v ControllerPolicy)`

SetControllerPolicy sets ControllerPolicy field to given value.


### GetInfrastructurePolicy

`func (o *SiteConfigVariables) GetInfrastructurePolicy() InfrastructurePolicy`

GetInfrastructurePolicy returns the InfrastructurePolicy field if non-nil, zero value otherwise.

### GetInfrastructurePolicyOk

`func (o *SiteConfigVariables) GetInfrastructurePolicyOk() (*InfrastructurePolicy, bool)`

GetInfrastructurePolicyOk returns a tuple with the InfrastructurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructurePolicy

`func (o *SiteConfigVariables) SetInfrastructurePolicy(v InfrastructurePolicy)`

SetInfrastructurePolicy sets InfrastructurePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


