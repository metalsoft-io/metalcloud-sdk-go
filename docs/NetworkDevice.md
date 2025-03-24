# NetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the network device. | 
**Revision** | **float32** | Revision number | 
**Status** | **string** | Current status of the network device | 
**SiteId** | **float32** | Site ID | 
**IdentifierString** | **string** | Hostname of the network device | 
**Description** | **string** | Description of the network device | 
**ChassisIdentifier** | **string** | Chassis identifier of the network device | 
**Country** | **string** | Country of the network device | 
**City** | **string** | City of the network device | 
**DatacenterMeta** | **string** | Datacenter metadata | 
**DatacenterRoom** | **string** | Room in the datacenter where the network device is located | 
**DatacenterRack** | **string** | Rack in the datacenter where the network device is located | 
**RackPositionUpperUnit** | **float32** | Upper rack position in the datacenter | 
**RackPositionLowerUnit** | **float32** | Lower rack position in the datacenter | 
**ManagementAddress** | **string** | Management address of the network device | 
**ManagementAddressMask** | **string** | Management address mask | 
**ManagementAddressGateway** | **string** | Management gateway address | 
**ManagementPort** | **float32** | Management port of the network device | 
**SyslogEnabled** | **float32** | Is the network device syslog enabled | 
**Username** | **string** | Username used to connect to the network device | 
**ManagementPassword** | **string** | Password used to connect to the network device | 
**ManagementMacAddress** | **string** | MAC address of the management interface | 
**SerialNumber** | **string** | Serial number of the network device | 
**Driver** | [**NetworkDeviceDriver**](NetworkDeviceDriver.md) | Driver software used to communicate with the network device | 
**Position** | [**SwitchPosition**](SwitchPosition.md) | The physical or logical position of the network device in the network topology. | 
**OrderIndex** | **float32** | Order index of the network device | 
**Tags** | **string** | Tags associated with the network device | 
**ReadyForInitialConfiguration** | **float32** | Whether the device is ready for initial configuration | 
**BootstrapReadinessCheckInProgress** | **float32** | Whether bootstrap readiness check is in progress | 
**SubnetOobId** | **float32** | Subnet ID for OOB management | 
**SubnetOobIndex** | **float32** | Subnet OOB index | 
**RequiresOsInstall** | **bool** | Whether the device requires OS installation | 
**BootstrapSkipInitialConfiguration** | **float32** | Whether to skip initial configuration during bootstrap | 
**BootstrapExpectedPartnerHostname** | **string** | Expected partner hostname during bootstrap | 
**LoopbackAddressIpv4** | Pointer to **string** | Loopback IPv4 address | [optional] 
**LoopbackAddressIpv6** | **string** | Loopback IPv6 address | 
**Asn** | **float32** | ASN of the network device | 
**VtepAddressIpv4** | Pointer to **string** | VTEP IPv4 address | [optional] 
**VtepAddressIpv6** | **string** | VTEP IPv6 address | 
**MlagSystemMac** | **string** | MLAG system MAC address | 
**MlagDomainId** | **float32** | MLAG domain ID | 
**QuarantineVlan** | **float32** | Quarantine VLAN ID | 
**DefaultMtu** | **float32** | Default MTU | 
**VariablesMaterializedForOSAssets** | **map[string]interface{}** | Variables materialized for OS assets | 
**SecretsMaterializedForOSAssets** | **map[string]interface{}** | Secrets materialized for OS assets | 
**BootstrapReadinessCheckResult** | **map[string]interface{}** | Bootstrap readiness check result | 
**IsGateway** | **bool** | Whether the network device is a gateway | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDevice

`func NewNetworkDevice(id string, revision float32, status string, siteId float32, identifierString string, description string, chassisIdentifier string, country string, city string, datacenterMeta string, datacenterRoom string, datacenterRack string, rackPositionUpperUnit float32, rackPositionLowerUnit float32, managementAddress string, managementAddressMask string, managementAddressGateway string, managementPort float32, syslogEnabled float32, username string, managementPassword string, managementMacAddress string, serialNumber string, driver NetworkDeviceDriver, position SwitchPosition, orderIndex float32, tags string, readyForInitialConfiguration float32, bootstrapReadinessCheckInProgress float32, subnetOobId float32, subnetOobIndex float32, requiresOsInstall bool, bootstrapSkipInitialConfiguration float32, bootstrapExpectedPartnerHostname string, loopbackAddressIpv6 string, asn float32, vtepAddressIpv6 string, mlagSystemMac string, mlagDomainId float32, quarantineVlan float32, defaultMtu float32, variablesMaterializedForOSAssets map[string]interface{}, secretsMaterializedForOSAssets map[string]interface{}, bootstrapReadinessCheckResult map[string]interface{}, isGateway bool, ) *NetworkDevice`

NewNetworkDevice instantiates a new NetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceWithDefaults

`func NewNetworkDeviceWithDefaults() *NetworkDevice`

NewNetworkDeviceWithDefaults instantiates a new NetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDevice) SetId(v string)`

SetId sets Id field to given value.


### GetRevision

`func (o *NetworkDevice) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkDevice) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkDevice) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetStatus

`func (o *NetworkDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDevice) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSiteId

`func (o *NetworkDevice) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *NetworkDevice) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *NetworkDevice) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetIdentifierString

`func (o *NetworkDevice) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *NetworkDevice) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *NetworkDevice) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.


### GetDescription

`func (o *NetworkDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDevice) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetChassisIdentifier

`func (o *NetworkDevice) GetChassisIdentifier() string`

GetChassisIdentifier returns the ChassisIdentifier field if non-nil, zero value otherwise.

### GetChassisIdentifierOk

`func (o *NetworkDevice) GetChassisIdentifierOk() (*string, bool)`

GetChassisIdentifierOk returns a tuple with the ChassisIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisIdentifier

`func (o *NetworkDevice) SetChassisIdentifier(v string)`

SetChassisIdentifier sets ChassisIdentifier field to given value.


### GetCountry

`func (o *NetworkDevice) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NetworkDevice) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NetworkDevice) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *NetworkDevice) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *NetworkDevice) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *NetworkDevice) SetCity(v string)`

SetCity sets City field to given value.


### GetDatacenterMeta

`func (o *NetworkDevice) GetDatacenterMeta() string`

GetDatacenterMeta returns the DatacenterMeta field if non-nil, zero value otherwise.

### GetDatacenterMetaOk

`func (o *NetworkDevice) GetDatacenterMetaOk() (*string, bool)`

GetDatacenterMetaOk returns a tuple with the DatacenterMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterMeta

`func (o *NetworkDevice) SetDatacenterMeta(v string)`

SetDatacenterMeta sets DatacenterMeta field to given value.


### GetDatacenterRoom

`func (o *NetworkDevice) GetDatacenterRoom() string`

GetDatacenterRoom returns the DatacenterRoom field if non-nil, zero value otherwise.

### GetDatacenterRoomOk

`func (o *NetworkDevice) GetDatacenterRoomOk() (*string, bool)`

GetDatacenterRoomOk returns a tuple with the DatacenterRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterRoom

`func (o *NetworkDevice) SetDatacenterRoom(v string)`

SetDatacenterRoom sets DatacenterRoom field to given value.


### GetDatacenterRack

`func (o *NetworkDevice) GetDatacenterRack() string`

GetDatacenterRack returns the DatacenterRack field if non-nil, zero value otherwise.

### GetDatacenterRackOk

`func (o *NetworkDevice) GetDatacenterRackOk() (*string, bool)`

GetDatacenterRackOk returns a tuple with the DatacenterRack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterRack

`func (o *NetworkDevice) SetDatacenterRack(v string)`

SetDatacenterRack sets DatacenterRack field to given value.


### GetRackPositionUpperUnit

`func (o *NetworkDevice) GetRackPositionUpperUnit() float32`

GetRackPositionUpperUnit returns the RackPositionUpperUnit field if non-nil, zero value otherwise.

### GetRackPositionUpperUnitOk

`func (o *NetworkDevice) GetRackPositionUpperUnitOk() (*float32, bool)`

GetRackPositionUpperUnitOk returns a tuple with the RackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionUpperUnit

`func (o *NetworkDevice) SetRackPositionUpperUnit(v float32)`

SetRackPositionUpperUnit sets RackPositionUpperUnit field to given value.


### GetRackPositionLowerUnit

`func (o *NetworkDevice) GetRackPositionLowerUnit() float32`

GetRackPositionLowerUnit returns the RackPositionLowerUnit field if non-nil, zero value otherwise.

### GetRackPositionLowerUnitOk

`func (o *NetworkDevice) GetRackPositionLowerUnitOk() (*float32, bool)`

GetRackPositionLowerUnitOk returns a tuple with the RackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionLowerUnit

`func (o *NetworkDevice) SetRackPositionLowerUnit(v float32)`

SetRackPositionLowerUnit sets RackPositionLowerUnit field to given value.


### GetManagementAddress

`func (o *NetworkDevice) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *NetworkDevice) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *NetworkDevice) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetManagementAddressMask

`func (o *NetworkDevice) GetManagementAddressMask() string`

GetManagementAddressMask returns the ManagementAddressMask field if non-nil, zero value otherwise.

### GetManagementAddressMaskOk

`func (o *NetworkDevice) GetManagementAddressMaskOk() (*string, bool)`

GetManagementAddressMaskOk returns a tuple with the ManagementAddressMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddressMask

`func (o *NetworkDevice) SetManagementAddressMask(v string)`

SetManagementAddressMask sets ManagementAddressMask field to given value.


### GetManagementAddressGateway

`func (o *NetworkDevice) GetManagementAddressGateway() string`

GetManagementAddressGateway returns the ManagementAddressGateway field if non-nil, zero value otherwise.

### GetManagementAddressGatewayOk

`func (o *NetworkDevice) GetManagementAddressGatewayOk() (*string, bool)`

GetManagementAddressGatewayOk returns a tuple with the ManagementAddressGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddressGateway

`func (o *NetworkDevice) SetManagementAddressGateway(v string)`

SetManagementAddressGateway sets ManagementAddressGateway field to given value.


### GetManagementPort

`func (o *NetworkDevice) GetManagementPort() float32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *NetworkDevice) GetManagementPortOk() (*float32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *NetworkDevice) SetManagementPort(v float32)`

SetManagementPort sets ManagementPort field to given value.


### GetSyslogEnabled

`func (o *NetworkDevice) GetSyslogEnabled() float32`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *NetworkDevice) GetSyslogEnabledOk() (*float32, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *NetworkDevice) SetSyslogEnabled(v float32)`

SetSyslogEnabled sets SyslogEnabled field to given value.


### GetUsername

`func (o *NetworkDevice) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NetworkDevice) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NetworkDevice) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetManagementPassword

`func (o *NetworkDevice) GetManagementPassword() string`

GetManagementPassword returns the ManagementPassword field if non-nil, zero value otherwise.

### GetManagementPasswordOk

`func (o *NetworkDevice) GetManagementPasswordOk() (*string, bool)`

GetManagementPasswordOk returns a tuple with the ManagementPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPassword

`func (o *NetworkDevice) SetManagementPassword(v string)`

SetManagementPassword sets ManagementPassword field to given value.


### GetManagementMacAddress

`func (o *NetworkDevice) GetManagementMacAddress() string`

GetManagementMacAddress returns the ManagementMacAddress field if non-nil, zero value otherwise.

### GetManagementMacAddressOk

`func (o *NetworkDevice) GetManagementMacAddressOk() (*string, bool)`

GetManagementMacAddressOk returns a tuple with the ManagementMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMacAddress

`func (o *NetworkDevice) SetManagementMacAddress(v string)`

SetManagementMacAddress sets ManagementMacAddress field to given value.


### GetSerialNumber

`func (o *NetworkDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NetworkDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NetworkDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDriver

`func (o *NetworkDevice) GetDriver() NetworkDeviceDriver`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *NetworkDevice) GetDriverOk() (*NetworkDeviceDriver, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *NetworkDevice) SetDriver(v NetworkDeviceDriver)`

SetDriver sets Driver field to given value.


### GetPosition

`func (o *NetworkDevice) GetPosition() SwitchPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *NetworkDevice) GetPositionOk() (*SwitchPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *NetworkDevice) SetPosition(v SwitchPosition)`

SetPosition sets Position field to given value.


### GetOrderIndex

`func (o *NetworkDevice) GetOrderIndex() float32`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *NetworkDevice) GetOrderIndexOk() (*float32, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *NetworkDevice) SetOrderIndex(v float32)`

SetOrderIndex sets OrderIndex field to given value.


### GetTags

`func (o *NetworkDevice) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkDevice) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkDevice) SetTags(v string)`

SetTags sets Tags field to given value.


### GetReadyForInitialConfiguration

`func (o *NetworkDevice) GetReadyForInitialConfiguration() float32`

GetReadyForInitialConfiguration returns the ReadyForInitialConfiguration field if non-nil, zero value otherwise.

### GetReadyForInitialConfigurationOk

`func (o *NetworkDevice) GetReadyForInitialConfigurationOk() (*float32, bool)`

GetReadyForInitialConfigurationOk returns a tuple with the ReadyForInitialConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyForInitialConfiguration

`func (o *NetworkDevice) SetReadyForInitialConfiguration(v float32)`

SetReadyForInitialConfiguration sets ReadyForInitialConfiguration field to given value.


### GetBootstrapReadinessCheckInProgress

`func (o *NetworkDevice) GetBootstrapReadinessCheckInProgress() float32`

GetBootstrapReadinessCheckInProgress returns the BootstrapReadinessCheckInProgress field if non-nil, zero value otherwise.

### GetBootstrapReadinessCheckInProgressOk

`func (o *NetworkDevice) GetBootstrapReadinessCheckInProgressOk() (*float32, bool)`

GetBootstrapReadinessCheckInProgressOk returns a tuple with the BootstrapReadinessCheckInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapReadinessCheckInProgress

`func (o *NetworkDevice) SetBootstrapReadinessCheckInProgress(v float32)`

SetBootstrapReadinessCheckInProgress sets BootstrapReadinessCheckInProgress field to given value.


### GetSubnetOobId

`func (o *NetworkDevice) GetSubnetOobId() float32`

GetSubnetOobId returns the SubnetOobId field if non-nil, zero value otherwise.

### GetSubnetOobIdOk

`func (o *NetworkDevice) GetSubnetOobIdOk() (*float32, bool)`

GetSubnetOobIdOk returns a tuple with the SubnetOobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOobId

`func (o *NetworkDevice) SetSubnetOobId(v float32)`

SetSubnetOobId sets SubnetOobId field to given value.


### GetSubnetOobIndex

`func (o *NetworkDevice) GetSubnetOobIndex() float32`

GetSubnetOobIndex returns the SubnetOobIndex field if non-nil, zero value otherwise.

### GetSubnetOobIndexOk

`func (o *NetworkDevice) GetSubnetOobIndexOk() (*float32, bool)`

GetSubnetOobIndexOk returns a tuple with the SubnetOobIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetOobIndex

`func (o *NetworkDevice) SetSubnetOobIndex(v float32)`

SetSubnetOobIndex sets SubnetOobIndex field to given value.


### GetRequiresOsInstall

`func (o *NetworkDevice) GetRequiresOsInstall() bool`

GetRequiresOsInstall returns the RequiresOsInstall field if non-nil, zero value otherwise.

### GetRequiresOsInstallOk

`func (o *NetworkDevice) GetRequiresOsInstallOk() (*bool, bool)`

GetRequiresOsInstallOk returns a tuple with the RequiresOsInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOsInstall

`func (o *NetworkDevice) SetRequiresOsInstall(v bool)`

SetRequiresOsInstall sets RequiresOsInstall field to given value.


### GetBootstrapSkipInitialConfiguration

`func (o *NetworkDevice) GetBootstrapSkipInitialConfiguration() float32`

GetBootstrapSkipInitialConfiguration returns the BootstrapSkipInitialConfiguration field if non-nil, zero value otherwise.

### GetBootstrapSkipInitialConfigurationOk

`func (o *NetworkDevice) GetBootstrapSkipInitialConfigurationOk() (*float32, bool)`

GetBootstrapSkipInitialConfigurationOk returns a tuple with the BootstrapSkipInitialConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapSkipInitialConfiguration

`func (o *NetworkDevice) SetBootstrapSkipInitialConfiguration(v float32)`

SetBootstrapSkipInitialConfiguration sets BootstrapSkipInitialConfiguration field to given value.


### GetBootstrapExpectedPartnerHostname

`func (o *NetworkDevice) GetBootstrapExpectedPartnerHostname() string`

GetBootstrapExpectedPartnerHostname returns the BootstrapExpectedPartnerHostname field if non-nil, zero value otherwise.

### GetBootstrapExpectedPartnerHostnameOk

`func (o *NetworkDevice) GetBootstrapExpectedPartnerHostnameOk() (*string, bool)`

GetBootstrapExpectedPartnerHostnameOk returns a tuple with the BootstrapExpectedPartnerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapExpectedPartnerHostname

`func (o *NetworkDevice) SetBootstrapExpectedPartnerHostname(v string)`

SetBootstrapExpectedPartnerHostname sets BootstrapExpectedPartnerHostname field to given value.


### GetLoopbackAddressIpv4

`func (o *NetworkDevice) GetLoopbackAddressIpv4() string`

GetLoopbackAddressIpv4 returns the LoopbackAddressIpv4 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv4Ok

`func (o *NetworkDevice) GetLoopbackAddressIpv4Ok() (*string, bool)`

GetLoopbackAddressIpv4Ok returns a tuple with the LoopbackAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv4

`func (o *NetworkDevice) SetLoopbackAddressIpv4(v string)`

SetLoopbackAddressIpv4 sets LoopbackAddressIpv4 field to given value.

### HasLoopbackAddressIpv4

`func (o *NetworkDevice) HasLoopbackAddressIpv4() bool`

HasLoopbackAddressIpv4 returns a boolean if a field has been set.

### GetLoopbackAddressIpv6

`func (o *NetworkDevice) GetLoopbackAddressIpv6() string`

GetLoopbackAddressIpv6 returns the LoopbackAddressIpv6 field if non-nil, zero value otherwise.

### GetLoopbackAddressIpv6Ok

`func (o *NetworkDevice) GetLoopbackAddressIpv6Ok() (*string, bool)`

GetLoopbackAddressIpv6Ok returns a tuple with the LoopbackAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddressIpv6

`func (o *NetworkDevice) SetLoopbackAddressIpv6(v string)`

SetLoopbackAddressIpv6 sets LoopbackAddressIpv6 field to given value.


### GetAsn

`func (o *NetworkDevice) GetAsn() float32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *NetworkDevice) GetAsnOk() (*float32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *NetworkDevice) SetAsn(v float32)`

SetAsn sets Asn field to given value.


### GetVtepAddressIpv4

`func (o *NetworkDevice) GetVtepAddressIpv4() string`

GetVtepAddressIpv4 returns the VtepAddressIpv4 field if non-nil, zero value otherwise.

### GetVtepAddressIpv4Ok

`func (o *NetworkDevice) GetVtepAddressIpv4Ok() (*string, bool)`

GetVtepAddressIpv4Ok returns a tuple with the VtepAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddressIpv4

`func (o *NetworkDevice) SetVtepAddressIpv4(v string)`

SetVtepAddressIpv4 sets VtepAddressIpv4 field to given value.

### HasVtepAddressIpv4

`func (o *NetworkDevice) HasVtepAddressIpv4() bool`

HasVtepAddressIpv4 returns a boolean if a field has been set.

### GetVtepAddressIpv6

`func (o *NetworkDevice) GetVtepAddressIpv6() string`

GetVtepAddressIpv6 returns the VtepAddressIpv6 field if non-nil, zero value otherwise.

### GetVtepAddressIpv6Ok

`func (o *NetworkDevice) GetVtepAddressIpv6Ok() (*string, bool)`

GetVtepAddressIpv6Ok returns a tuple with the VtepAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddressIpv6

`func (o *NetworkDevice) SetVtepAddressIpv6(v string)`

SetVtepAddressIpv6 sets VtepAddressIpv6 field to given value.


### GetMlagSystemMac

`func (o *NetworkDevice) GetMlagSystemMac() string`

GetMlagSystemMac returns the MlagSystemMac field if non-nil, zero value otherwise.

### GetMlagSystemMacOk

`func (o *NetworkDevice) GetMlagSystemMacOk() (*string, bool)`

GetMlagSystemMacOk returns a tuple with the MlagSystemMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagSystemMac

`func (o *NetworkDevice) SetMlagSystemMac(v string)`

SetMlagSystemMac sets MlagSystemMac field to given value.


### GetMlagDomainId

`func (o *NetworkDevice) GetMlagDomainId() float32`

GetMlagDomainId returns the MlagDomainId field if non-nil, zero value otherwise.

### GetMlagDomainIdOk

`func (o *NetworkDevice) GetMlagDomainIdOk() (*float32, bool)`

GetMlagDomainIdOk returns a tuple with the MlagDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagDomainId

`func (o *NetworkDevice) SetMlagDomainId(v float32)`

SetMlagDomainId sets MlagDomainId field to given value.


### GetQuarantineVlan

`func (o *NetworkDevice) GetQuarantineVlan() float32`

GetQuarantineVlan returns the QuarantineVlan field if non-nil, zero value otherwise.

### GetQuarantineVlanOk

`func (o *NetworkDevice) GetQuarantineVlanOk() (*float32, bool)`

GetQuarantineVlanOk returns a tuple with the QuarantineVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineVlan

`func (o *NetworkDevice) SetQuarantineVlan(v float32)`

SetQuarantineVlan sets QuarantineVlan field to given value.


### GetDefaultMtu

`func (o *NetworkDevice) GetDefaultMtu() float32`

GetDefaultMtu returns the DefaultMtu field if non-nil, zero value otherwise.

### GetDefaultMtuOk

`func (o *NetworkDevice) GetDefaultMtuOk() (*float32, bool)`

GetDefaultMtuOk returns a tuple with the DefaultMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtu

`func (o *NetworkDevice) SetDefaultMtu(v float32)`

SetDefaultMtu sets DefaultMtu field to given value.


### GetVariablesMaterializedForOSAssets

`func (o *NetworkDevice) GetVariablesMaterializedForOSAssets() map[string]interface{}`

GetVariablesMaterializedForOSAssets returns the VariablesMaterializedForOSAssets field if non-nil, zero value otherwise.

### GetVariablesMaterializedForOSAssetsOk

`func (o *NetworkDevice) GetVariablesMaterializedForOSAssetsOk() (*map[string]interface{}, bool)`

GetVariablesMaterializedForOSAssetsOk returns a tuple with the VariablesMaterializedForOSAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesMaterializedForOSAssets

`func (o *NetworkDevice) SetVariablesMaterializedForOSAssets(v map[string]interface{})`

SetVariablesMaterializedForOSAssets sets VariablesMaterializedForOSAssets field to given value.


### GetSecretsMaterializedForOSAssets

`func (o *NetworkDevice) GetSecretsMaterializedForOSAssets() map[string]interface{}`

GetSecretsMaterializedForOSAssets returns the SecretsMaterializedForOSAssets field if non-nil, zero value otherwise.

### GetSecretsMaterializedForOSAssetsOk

`func (o *NetworkDevice) GetSecretsMaterializedForOSAssetsOk() (*map[string]interface{}, bool)`

GetSecretsMaterializedForOSAssetsOk returns a tuple with the SecretsMaterializedForOSAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsMaterializedForOSAssets

`func (o *NetworkDevice) SetSecretsMaterializedForOSAssets(v map[string]interface{})`

SetSecretsMaterializedForOSAssets sets SecretsMaterializedForOSAssets field to given value.


### GetBootstrapReadinessCheckResult

`func (o *NetworkDevice) GetBootstrapReadinessCheckResult() map[string]interface{}`

GetBootstrapReadinessCheckResult returns the BootstrapReadinessCheckResult field if non-nil, zero value otherwise.

### GetBootstrapReadinessCheckResultOk

`func (o *NetworkDevice) GetBootstrapReadinessCheckResultOk() (*map[string]interface{}, bool)`

GetBootstrapReadinessCheckResultOk returns a tuple with the BootstrapReadinessCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapReadinessCheckResult

`func (o *NetworkDevice) SetBootstrapReadinessCheckResult(v map[string]interface{})`

SetBootstrapReadinessCheckResult sets BootstrapReadinessCheckResult field to given value.


### GetIsGateway

`func (o *NetworkDevice) GetIsGateway() bool`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *NetworkDevice) GetIsGatewayOk() (*bool, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *NetworkDevice) SetIsGateway(v bool)`

SetIsGateway sets IsGateway field to given value.


### GetLinks

`func (o *NetworkDevice) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDevice) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDevice) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDevice) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


