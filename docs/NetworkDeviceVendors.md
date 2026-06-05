# NetworkDeviceVendors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique ID of the network device vendor entry | 
**Revision** | **int64** | Revision number | 
**Kind** | **string** | The driver/kind of the network device | 
**OidGroups** | [**[]SNMPOIDGroupConfig**](SNMPOIDGroupConfig.md) | List of SNMP OID groups to be used for monitoring devices of this vendor | 
**HealthCheckRules** | [**NetworkDeviceHealthCheckRules**](NetworkDeviceHealthCheckRules.md) | Health check rules for the network device | 
**OptionalFilesToBackup** | [**NetworkDeviceOptionalFilesToBackup**](NetworkDeviceOptionalFilesToBackup.md) | Optional files to backup for the network device | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceVendors

`func NewNetworkDeviceVendors(id int64, revision int64, kind string, oidGroups []SNMPOIDGroupConfig, healthCheckRules NetworkDeviceHealthCheckRules, optionalFilesToBackup NetworkDeviceOptionalFilesToBackup, ) *NetworkDeviceVendors`

NewNetworkDeviceVendors instantiates a new NetworkDeviceVendors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceVendorsWithDefaults

`func NewNetworkDeviceVendorsWithDefaults() *NetworkDeviceVendors`

NewNetworkDeviceVendorsWithDefaults instantiates a new NetworkDeviceVendors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDeviceVendors) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceVendors) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceVendors) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *NetworkDeviceVendors) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkDeviceVendors) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkDeviceVendors) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *NetworkDeviceVendors) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkDeviceVendors) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkDeviceVendors) SetKind(v string)`

SetKind sets Kind field to given value.


### GetOidGroups

`func (o *NetworkDeviceVendors) GetOidGroups() []SNMPOIDGroupConfig`

GetOidGroups returns the OidGroups field if non-nil, zero value otherwise.

### GetOidGroupsOk

`func (o *NetworkDeviceVendors) GetOidGroupsOk() (*[]SNMPOIDGroupConfig, bool)`

GetOidGroupsOk returns a tuple with the OidGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidGroups

`func (o *NetworkDeviceVendors) SetOidGroups(v []SNMPOIDGroupConfig)`

SetOidGroups sets OidGroups field to given value.


### GetHealthCheckRules

`func (o *NetworkDeviceVendors) GetHealthCheckRules() NetworkDeviceHealthCheckRules`

GetHealthCheckRules returns the HealthCheckRules field if non-nil, zero value otherwise.

### GetHealthCheckRulesOk

`func (o *NetworkDeviceVendors) GetHealthCheckRulesOk() (*NetworkDeviceHealthCheckRules, bool)`

GetHealthCheckRulesOk returns a tuple with the HealthCheckRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckRules

`func (o *NetworkDeviceVendors) SetHealthCheckRules(v NetworkDeviceHealthCheckRules)`

SetHealthCheckRules sets HealthCheckRules field to given value.


### GetOptionalFilesToBackup

`func (o *NetworkDeviceVendors) GetOptionalFilesToBackup() NetworkDeviceOptionalFilesToBackup`

GetOptionalFilesToBackup returns the OptionalFilesToBackup field if non-nil, zero value otherwise.

### GetOptionalFilesToBackupOk

`func (o *NetworkDeviceVendors) GetOptionalFilesToBackupOk() (*NetworkDeviceOptionalFilesToBackup, bool)`

GetOptionalFilesToBackupOk returns a tuple with the OptionalFilesToBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalFilesToBackup

`func (o *NetworkDeviceVendors) SetOptionalFilesToBackup(v NetworkDeviceOptionalFilesToBackup)`

SetOptionalFilesToBackup sets OptionalFilesToBackup field to given value.


### GetLinks

`func (o *NetworkDeviceVendors) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceVendors) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceVendors) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceVendors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


