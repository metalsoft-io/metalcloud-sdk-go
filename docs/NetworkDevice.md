# NetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchId** | **float32** | Unique identifier for the switch. | 
**Revision** | **float32** | Revision number | 
**Status** | Pointer to **string** | Current status of the network device | [optional] 
**DatacenterName** | Pointer to **string** | Datacenter name where the network device is located | [optional] 
**SiteId** | Pointer to **float32** | Site ID | [optional] 
**IdentifierString** | Pointer to **string** | Identifier string of the network device | [optional] 
**Description** | Pointer to **string** | Description of the network device | [optional] 
**ChassisIdentifier** | Pointer to **string** | Chassis identifier of the network device | [optional] 
**Country** | Pointer to **string** | Country of the network device | [optional] 
**City** | Pointer to **string** | City of the network device | [optional] 
**DatacenterMeta** | Pointer to **string** | Datacenter metadata | [optional] 
**DatacenterRoom** | Pointer to **string** | Room in the datacenter where the network device is located | [optional] 
**DatacenterRack** | Pointer to **string** | Rack in the datacenter where the network device is located | [optional] 
**RackPositionUpperUnit** | Pointer to **float32** | Upper rack position in the datacenter | [optional] 
**RackPositionLowerUnit** | Pointer to **float32** | Lower rack position in the datacenter | [optional] 
**ManagementAddress** | Pointer to **string** | Management address of the network device | [optional] 
**ManagementAddressMask** | Pointer to **string** | Management address mask | [optional] 
**ManagementAddressGateway** | Pointer to **string** | Management gateway address | [optional] 
**ManagementPort** | Pointer to **float32** | Management port of the network device | [optional] 
**SyslogEnabled** | Pointer to **bool** | Is the network device syslog enabled | [optional] 
**Username** | Pointer to **string** | Username used to connect to the network device | [optional] 
**ManagementPassword** | Pointer to **string** | Password used to connect to the network device | [optional] 
**ManagementProtocol** | Pointer to **string** | Management protocol used by the network device | [optional] 
**ManagementMacAddress** | Pointer to **string** | MAC address of the management interface | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the network device | [optional] 
**Driver** | [**NetworkDeviceDriver**](NetworkDeviceDriver.md) | Driver software used to communicate with the network device | 
**Position** | [**SwitchPosition**](SwitchPosition.md) | The physical or logical position of the network device in the network topology. | 
**ProvisionerType** | Pointer to **string** | Provisioner type of the network device | [optional] 
**NetworkTypesAllowed** | Pointer to **[]string** | Allowed network types for the network device | [optional] 
**OrderIndex** | Pointer to **float32** | Order index of the network device | [optional] 
**Tags** | Pointer to **string** | Tags associated with the network device | [optional] 
**ReadyForInitialConfiguration** | Pointer to **float32** | Whether the device is ready for initial configuration | [optional] 
**BootstrapReadinessCheckInProgress** | Pointer to **float32** | Whether bootstrap readiness check is in progress | [optional] 
**SubnetOobId** | Pointer to **float32** | Subnet ID for OOB management | [optional] 
**SubnetOobIndex** | Pointer to **float32** | Subnet OOB index | [optional] 
**RequiresOsInstall** | Pointer to **bool** | Whether the device requires OS installation | [optional] 
**BootstrapSkipInitialConfiguration** | Pointer to **float32** | Whether to skip initial configuration during bootstrap | [optional] 
**BootstrapExpectedPartnerHostname** | Pointer to **string** | Expected partner hostname during bootstrap | [optional] 
**LoopbackAddress** | Pointer to **string** | Loopback IPv4 address | [optional] 
**LoopbackAddressIpv6** | Pointer to **string** | Loopback IPv6 address | [optional] 
**Asn** | Pointer to **float32** | ASN of the network device | [optional] 
**VtepAddress** | Pointer to **string** | VTEP IPv4 address | [optional] 
**VtepAddressIpv6** | Pointer to **string** | VTEP IPv6 address | [optional] 
**MlagSystemMac** | Pointer to **string** | MLAG system MAC address | [optional] 
**MlagDomainId** | Pointer to **float32** | MLAG domain ID | [optional] 
**QuarantineSubnetStart** | Pointer to **string** | Quarantine subnet start address | [optional] 
**QuarantineSubnetEnd** | Pointer to **string** | Quarantine subnet end address | [optional] 
**QuarantineSubnetPrefixSize** | Pointer to **float32** | Quarantine subnet prefix size | [optional] 
**QuarantineSubnetGateway** | Pointer to **string** | Quarantine subnet gateway address | [optional] 
**QuarantineVlan** | Pointer to **float32** | Quarantine VLAN ID | [optional] 
**DefaultMtu** | Pointer to **float32** | Default MTU | [optional] 
**VariablesMaterializedForOSAssets** | Pointer to **map[string]interface{}** | Variables materialized for OS assets | [optional] 
**SecretsMaterializedForOSAssets** | Pointer to **map[string]interface{}** | Secrets materialized for OS assets | [optional] 
**BootstrapReadinessCheckResult** | Pointer to **map[string]interface{}** | Bootstrap readiness check result | [optional] 
**IsGateway** | Pointer to **float32** | Whether the network device is a gateway | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDevice

`func NewNetworkDevice(switchId float32, revision float32, driver NetworkDeviceDriver, position SwitchPosition, ) *NetworkDevice`

NewNetworkDevice instantiates a new NetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceWithDefaults

`func NewNetworkDeviceWithDefaults() *NetworkDevice`

NewNetworkDeviceWithDefaults instantiates a new NetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchId

`func (o *NetworkDevice) GetSwitchId() float32`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NetworkDevice) GetSwitchIdOk() (*float32, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NetworkDevice) SetSwitchId(v float32)`

SetSwitchId sets SwitchId field to given value.


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

### HasStatus

`func (o *NetworkDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDatacenterName

`func (o *NetworkDevice) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *NetworkDevice) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *NetworkDevice) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *NetworkDevice) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

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

### HasSiteId

`func (o *NetworkDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

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

### HasIdentifierString

`func (o *NetworkDevice) HasIdentifierString() bool`

HasIdentifierString returns a boolean if a field has been set.

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

### HasDescription

`func (o *NetworkDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### HasChassisIdentifier

`func (o *NetworkDevice) HasChassisIdentifier() bool`

HasChassisIdentifier returns a boolean if a field has been set.

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

### HasCountry

`func (o *NetworkDevice) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

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

### HasCity

`func (o *NetworkDevice) HasCity() bool`

HasCity returns a boolean if a field has been set.

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

### HasDatacenterMeta

`func (o *NetworkDevice) HasDatacenterMeta() bool`

HasDatacenterMeta returns a boolean if a field has been set.

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

### HasDatacenterRoom

`func (o *NetworkDevice) HasDatacenterRoom() bool`

HasDatacenterRoom returns a boolean if a field has been set.

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

### HasDatacenterRack

`func (o *NetworkDevice) HasDatacenterRack() bool`

HasDatacenterRack returns a boolean if a field has been set.

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

### HasRackPositionUpperUnit

`func (o *NetworkDevice) HasRackPositionUpperUnit() bool`

HasRackPositionUpperUnit returns a boolean if a field has been set.

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

### HasRackPositionLowerUnit

`func (o *NetworkDevice) HasRackPositionLowerUnit() bool`

HasRackPositionLowerUnit returns a boolean if a field has been set.

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

### HasManagementAddress

`func (o *NetworkDevice) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

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

### HasManagementAddressMask

`func (o *NetworkDevice) HasManagementAddressMask() bool`

HasManagementAddressMask returns a boolean if a field has been set.

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

### HasManagementAddressGateway

`func (o *NetworkDevice) HasManagementAddressGateway() bool`

HasManagementAddressGateway returns a boolean if a field has been set.

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

### HasManagementPort

`func (o *NetworkDevice) HasManagementPort() bool`

HasManagementPort returns a boolean if a field has been set.

### GetSyslogEnabled

`func (o *NetworkDevice) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *NetworkDevice) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *NetworkDevice) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *NetworkDevice) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.

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

### HasUsername

`func (o *NetworkDevice) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

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

### HasManagementPassword

`func (o *NetworkDevice) HasManagementPassword() bool`

HasManagementPassword returns a boolean if a field has been set.

### GetManagementProtocol

`func (o *NetworkDevice) GetManagementProtocol() string`

GetManagementProtocol returns the ManagementProtocol field if non-nil, zero value otherwise.

### GetManagementProtocolOk

`func (o *NetworkDevice) GetManagementProtocolOk() (*string, bool)`

GetManagementProtocolOk returns a tuple with the ManagementProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementProtocol

`func (o *NetworkDevice) SetManagementProtocol(v string)`

SetManagementProtocol sets ManagementProtocol field to given value.

### HasManagementProtocol

`func (o *NetworkDevice) HasManagementProtocol() bool`

HasManagementProtocol returns a boolean if a field has been set.

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

### HasManagementMacAddress

`func (o *NetworkDevice) HasManagementMacAddress() bool`

HasManagementMacAddress returns a boolean if a field has been set.

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

### HasSerialNumber

`func (o *NetworkDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

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


### GetProvisionerType

`func (o *NetworkDevice) GetProvisionerType() string`

GetProvisionerType returns the ProvisionerType field if non-nil, zero value otherwise.

### GetProvisionerTypeOk

`func (o *NetworkDevice) GetProvisionerTypeOk() (*string, bool)`

GetProvisionerTypeOk returns a tuple with the ProvisionerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerType

`func (o *NetworkDevice) SetProvisionerType(v string)`

SetProvisionerType sets ProvisionerType field to given value.

### HasProvisionerType

`func (o *NetworkDevice) HasProvisionerType() bool`

HasProvisionerType returns a boolean if a field has been set.

### GetNetworkTypesAllowed

`func (o *NetworkDevice) GetNetworkTypesAllowed() []string`

GetNetworkTypesAllowed returns the NetworkTypesAllowed field if non-nil, zero value otherwise.

### GetNetworkTypesAllowedOk

`func (o *NetworkDevice) GetNetworkTypesAllowedOk() (*[]string, bool)`

GetNetworkTypesAllowedOk returns a tuple with the NetworkTypesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypesAllowed

`func (o *NetworkDevice) SetNetworkTypesAllowed(v []string)`

SetNetworkTypesAllowed sets NetworkTypesAllowed field to given value.

### HasNetworkTypesAllowed

`func (o *NetworkDevice) HasNetworkTypesAllowed() bool`

HasNetworkTypesAllowed returns a boolean if a field has been set.

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

### HasOrderIndex

`func (o *NetworkDevice) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.

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

### HasTags

`func (o *NetworkDevice) HasTags() bool`

HasTags returns a boolean if a field has been set.

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

### HasReadyForInitialConfiguration

`func (o *NetworkDevice) HasReadyForInitialConfiguration() bool`

HasReadyForInitialConfiguration returns a boolean if a field has been set.

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

### HasBootstrapReadinessCheckInProgress

`func (o *NetworkDevice) HasBootstrapReadinessCheckInProgress() bool`

HasBootstrapReadinessCheckInProgress returns a boolean if a field has been set.

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

### HasSubnetOobId

`func (o *NetworkDevice) HasSubnetOobId() bool`

HasSubnetOobId returns a boolean if a field has been set.

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

### HasSubnetOobIndex

`func (o *NetworkDevice) HasSubnetOobIndex() bool`

HasSubnetOobIndex returns a boolean if a field has been set.

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

### HasRequiresOsInstall

`func (o *NetworkDevice) HasRequiresOsInstall() bool`

HasRequiresOsInstall returns a boolean if a field has been set.

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

### HasBootstrapSkipInitialConfiguration

`func (o *NetworkDevice) HasBootstrapSkipInitialConfiguration() bool`

HasBootstrapSkipInitialConfiguration returns a boolean if a field has been set.

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

### HasBootstrapExpectedPartnerHostname

`func (o *NetworkDevice) HasBootstrapExpectedPartnerHostname() bool`

HasBootstrapExpectedPartnerHostname returns a boolean if a field has been set.

### GetLoopbackAddress

`func (o *NetworkDevice) GetLoopbackAddress() string`

GetLoopbackAddress returns the LoopbackAddress field if non-nil, zero value otherwise.

### GetLoopbackAddressOk

`func (o *NetworkDevice) GetLoopbackAddressOk() (*string, bool)`

GetLoopbackAddressOk returns a tuple with the LoopbackAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddress

`func (o *NetworkDevice) SetLoopbackAddress(v string)`

SetLoopbackAddress sets LoopbackAddress field to given value.

### HasLoopbackAddress

`func (o *NetworkDevice) HasLoopbackAddress() bool`

HasLoopbackAddress returns a boolean if a field has been set.

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

### HasLoopbackAddressIpv6

`func (o *NetworkDevice) HasLoopbackAddressIpv6() bool`

HasLoopbackAddressIpv6 returns a boolean if a field has been set.

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

### HasAsn

`func (o *NetworkDevice) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetVtepAddress

`func (o *NetworkDevice) GetVtepAddress() string`

GetVtepAddress returns the VtepAddress field if non-nil, zero value otherwise.

### GetVtepAddressOk

`func (o *NetworkDevice) GetVtepAddressOk() (*string, bool)`

GetVtepAddressOk returns a tuple with the VtepAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddress

`func (o *NetworkDevice) SetVtepAddress(v string)`

SetVtepAddress sets VtepAddress field to given value.

### HasVtepAddress

`func (o *NetworkDevice) HasVtepAddress() bool`

HasVtepAddress returns a boolean if a field has been set.

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

### HasVtepAddressIpv6

`func (o *NetworkDevice) HasVtepAddressIpv6() bool`

HasVtepAddressIpv6 returns a boolean if a field has been set.

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

### HasMlagSystemMac

`func (o *NetworkDevice) HasMlagSystemMac() bool`

HasMlagSystemMac returns a boolean if a field has been set.

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

### HasMlagDomainId

`func (o *NetworkDevice) HasMlagDomainId() bool`

HasMlagDomainId returns a boolean if a field has been set.

### GetQuarantineSubnetStart

`func (o *NetworkDevice) GetQuarantineSubnetStart() string`

GetQuarantineSubnetStart returns the QuarantineSubnetStart field if non-nil, zero value otherwise.

### GetQuarantineSubnetStartOk

`func (o *NetworkDevice) GetQuarantineSubnetStartOk() (*string, bool)`

GetQuarantineSubnetStartOk returns a tuple with the QuarantineSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetStart

`func (o *NetworkDevice) SetQuarantineSubnetStart(v string)`

SetQuarantineSubnetStart sets QuarantineSubnetStart field to given value.

### HasQuarantineSubnetStart

`func (o *NetworkDevice) HasQuarantineSubnetStart() bool`

HasQuarantineSubnetStart returns a boolean if a field has been set.

### GetQuarantineSubnetEnd

`func (o *NetworkDevice) GetQuarantineSubnetEnd() string`

GetQuarantineSubnetEnd returns the QuarantineSubnetEnd field if non-nil, zero value otherwise.

### GetQuarantineSubnetEndOk

`func (o *NetworkDevice) GetQuarantineSubnetEndOk() (*string, bool)`

GetQuarantineSubnetEndOk returns a tuple with the QuarantineSubnetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetEnd

`func (o *NetworkDevice) SetQuarantineSubnetEnd(v string)`

SetQuarantineSubnetEnd sets QuarantineSubnetEnd field to given value.

### HasQuarantineSubnetEnd

`func (o *NetworkDevice) HasQuarantineSubnetEnd() bool`

HasQuarantineSubnetEnd returns a boolean if a field has been set.

### GetQuarantineSubnetPrefixSize

`func (o *NetworkDevice) GetQuarantineSubnetPrefixSize() float32`

GetQuarantineSubnetPrefixSize returns the QuarantineSubnetPrefixSize field if non-nil, zero value otherwise.

### GetQuarantineSubnetPrefixSizeOk

`func (o *NetworkDevice) GetQuarantineSubnetPrefixSizeOk() (*float32, bool)`

GetQuarantineSubnetPrefixSizeOk returns a tuple with the QuarantineSubnetPrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetPrefixSize

`func (o *NetworkDevice) SetQuarantineSubnetPrefixSize(v float32)`

SetQuarantineSubnetPrefixSize sets QuarantineSubnetPrefixSize field to given value.

### HasQuarantineSubnetPrefixSize

`func (o *NetworkDevice) HasQuarantineSubnetPrefixSize() bool`

HasQuarantineSubnetPrefixSize returns a boolean if a field has been set.

### GetQuarantineSubnetGateway

`func (o *NetworkDevice) GetQuarantineSubnetGateway() string`

GetQuarantineSubnetGateway returns the QuarantineSubnetGateway field if non-nil, zero value otherwise.

### GetQuarantineSubnetGatewayOk

`func (o *NetworkDevice) GetQuarantineSubnetGatewayOk() (*string, bool)`

GetQuarantineSubnetGatewayOk returns a tuple with the QuarantineSubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetGateway

`func (o *NetworkDevice) SetQuarantineSubnetGateway(v string)`

SetQuarantineSubnetGateway sets QuarantineSubnetGateway field to given value.

### HasQuarantineSubnetGateway

`func (o *NetworkDevice) HasQuarantineSubnetGateway() bool`

HasQuarantineSubnetGateway returns a boolean if a field has been set.

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

### HasQuarantineVlan

`func (o *NetworkDevice) HasQuarantineVlan() bool`

HasQuarantineVlan returns a boolean if a field has been set.

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

### HasDefaultMtu

`func (o *NetworkDevice) HasDefaultMtu() bool`

HasDefaultMtu returns a boolean if a field has been set.

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

### HasVariablesMaterializedForOSAssets

`func (o *NetworkDevice) HasVariablesMaterializedForOSAssets() bool`

HasVariablesMaterializedForOSAssets returns a boolean if a field has been set.

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

### HasSecretsMaterializedForOSAssets

`func (o *NetworkDevice) HasSecretsMaterializedForOSAssets() bool`

HasSecretsMaterializedForOSAssets returns a boolean if a field has been set.

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

### HasBootstrapReadinessCheckResult

`func (o *NetworkDevice) HasBootstrapReadinessCheckResult() bool`

HasBootstrapReadinessCheckResult returns a boolean if a field has been set.

### GetIsGateway

`func (o *NetworkDevice) GetIsGateway() float32`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *NetworkDevice) GetIsGatewayOk() (*float32, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *NetworkDevice) SetIsGateway(v float32)`

SetIsGateway sets IsGateway field to given value.

### HasIsGateway

`func (o *NetworkDevice) HasIsGateway() bool`

HasIsGateway returns a boolean if a field has been set.

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


