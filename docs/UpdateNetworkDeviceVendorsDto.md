# UpdateNetworkDeviceVendorsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OidGroups** | Pointer to [**[]SNMPOIDGroupConfig**](SNMPOIDGroupConfig.md) | List of SNMP OID groups to be used for monitoring devices of this vendor | [optional] 
**HealthCheckRules** | Pointer to [**NetworkDeviceHealthCheckRules**](NetworkDeviceHealthCheckRules.md) | Health check rules for the network device | [optional] 
**OptionalFilesToBackup** | Pointer to [**NetworkDeviceOptionalFilesToBackup**](NetworkDeviceOptionalFilesToBackup.md) | Optional files to backup for the network device | [optional] 

## Methods

### NewUpdateNetworkDeviceVendorsDto

`func NewUpdateNetworkDeviceVendorsDto() *UpdateNetworkDeviceVendorsDto`

NewUpdateNetworkDeviceVendorsDto instantiates a new UpdateNetworkDeviceVendorsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDeviceVendorsDtoWithDefaults

`func NewUpdateNetworkDeviceVendorsDtoWithDefaults() *UpdateNetworkDeviceVendorsDto`

NewUpdateNetworkDeviceVendorsDtoWithDefaults instantiates a new UpdateNetworkDeviceVendorsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidGroups

`func (o *UpdateNetworkDeviceVendorsDto) GetOidGroups() []SNMPOIDGroupConfig`

GetOidGroups returns the OidGroups field if non-nil, zero value otherwise.

### GetOidGroupsOk

`func (o *UpdateNetworkDeviceVendorsDto) GetOidGroupsOk() (*[]SNMPOIDGroupConfig, bool)`

GetOidGroupsOk returns a tuple with the OidGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidGroups

`func (o *UpdateNetworkDeviceVendorsDto) SetOidGroups(v []SNMPOIDGroupConfig)`

SetOidGroups sets OidGroups field to given value.

### HasOidGroups

`func (o *UpdateNetworkDeviceVendorsDto) HasOidGroups() bool`

HasOidGroups returns a boolean if a field has been set.

### GetHealthCheckRules

`func (o *UpdateNetworkDeviceVendorsDto) GetHealthCheckRules() NetworkDeviceHealthCheckRules`

GetHealthCheckRules returns the HealthCheckRules field if non-nil, zero value otherwise.

### GetHealthCheckRulesOk

`func (o *UpdateNetworkDeviceVendorsDto) GetHealthCheckRulesOk() (*NetworkDeviceHealthCheckRules, bool)`

GetHealthCheckRulesOk returns a tuple with the HealthCheckRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckRules

`func (o *UpdateNetworkDeviceVendorsDto) SetHealthCheckRules(v NetworkDeviceHealthCheckRules)`

SetHealthCheckRules sets HealthCheckRules field to given value.

### HasHealthCheckRules

`func (o *UpdateNetworkDeviceVendorsDto) HasHealthCheckRules() bool`

HasHealthCheckRules returns a boolean if a field has been set.

### GetOptionalFilesToBackup

`func (o *UpdateNetworkDeviceVendorsDto) GetOptionalFilesToBackup() NetworkDeviceOptionalFilesToBackup`

GetOptionalFilesToBackup returns the OptionalFilesToBackup field if non-nil, zero value otherwise.

### GetOptionalFilesToBackupOk

`func (o *UpdateNetworkDeviceVendorsDto) GetOptionalFilesToBackupOk() (*NetworkDeviceOptionalFilesToBackup, bool)`

GetOptionalFilesToBackupOk returns a tuple with the OptionalFilesToBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalFilesToBackup

`func (o *UpdateNetworkDeviceVendorsDto) SetOptionalFilesToBackup(v NetworkDeviceOptionalFilesToBackup)`

SetOptionalFilesToBackup sets OptionalFilesToBackup field to given value.

### HasOptionalFilesToBackup

`func (o *UpdateNetworkDeviceVendorsDto) HasOptionalFilesToBackup() bool`

HasOptionalFilesToBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


