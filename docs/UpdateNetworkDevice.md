# UpdateNetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **int32** | Site identifier | [optional] 
**DatacenterName** | Pointer to **string** | Name of the datacenter | [optional] 
**IdentifierString** | Pointer to **string** | Unique identifier string for the network device | [optional] 
**ChassisRackId** | Pointer to **int32** | ID of the rack where the network device chassis is installed | [optional] 
**ChassisIdentifier** | Pointer to **NullableString** | Unique identifier for the network device chassis | [optional] 
**ProvisionerType** | Pointer to [**ProvisionerType**](ProvisionerType.md) | Type of provisioner used for this network device | [optional] 
**Driver** | Pointer to [**NetworkDeviceDriver**](NetworkDeviceDriver.md) | Driver software used to communicate with the network device | [optional] 
**Position** | Pointer to **string** | The physical or logical position of the network device in the network topology. | [optional] 
**TorLinkedId** | Pointer to **NullableFloat32** | ID of the Top-of-Rack (TOR) switch that this network device is linked to. Used for establishing hierarchical relationships between network devices. | [optional] 
**IsGateway** | Pointer to **bool** | Indicates if this network device acts as a gateway for external network traffic | [optional] 
**SyslogEnabled** | Pointer to **NullableBool** | Indicates if syslog logging is enabled for this network device | [optional] 
**IsStorageSwitch** | Pointer to **bool** | Is storage network device | [optional] 
**IsBorderDevice** | Pointer to **bool** | Is border device | [optional] 
**QuarantineSubnetStart** | Pointer to **NullableString** | The starting IP address of the quarantine subnet range | [optional] 
**QuarantineSubnetEnd** | Pointer to **NullableString** | The ending IP address of the quarantine subnet range | [optional] 
**QuarantineSubnetPrefixSize** | Pointer to **NullableInt32** | The prefix size for the quarantine subnet | [optional] 
**QuarantineSubnetGateway** | Pointer to **NullableString** | The gateway IP address for the quarantine subnet | [optional] 
**QuarantineVlan** | Pointer to **int32** | The VLAN ID for the quarantine network | [optional] 
**ManagementProtocol** | Pointer to **NullableString** | The protocol used to manage the network device (e.g., SSH, HTTPS) | [optional] 
**ManagementAddress** | Pointer to **NullableString** | Management Address | [optional] 
**ManagementPort** | Pointer to **NullableInt32** | Management Port | [optional] 
**Username** | Pointer to **NullableString** | The username used for management authentication | [optional] 
**ManagementPassword** | Pointer to **NullableString** | The password used for management authentication | [optional] 
**ManagementAddressGateway** | Pointer to **NullableString** | The gateway IP address for the management network | [optional] 
**ManagementAddressMask** | Pointer to **NullableString** | The subnet mask for the management network | [optional] 
**ManagementMAC** | Pointer to **NullableString** | The MAC address of the management interface | [optional] 
**LoopbackAddress** | Pointer to **NullableString** | Loopback Address | [optional] 
**VtepAddress** | Pointer to **NullableString** | VTEP Address | [optional] 
**Asn** | Pointer to **NullableFloat32** | The Autonomous System Number for BGP routing | [optional] 
**NetworkTypesAllowed** | Pointer to **[]string** | Network types allowed | [optional] 
**Description** | Pointer to **NullableString** | Additional description or notes about the network device | [optional] 
**Country** | Pointer to **NullableString** | The country where the network device is located | [optional] 
**City** | Pointer to **NullableString** | The city where the network device is located | [optional] 
**DatacenterMeta** | Pointer to **NullableString** | Metadata about the datacenter where the device is located | [optional] 
**DatacenterRoom** | Pointer to **NullableString** | The room within the datacenter where the device is located | [optional] 
**DatacenterRack** | Pointer to **NullableString** | The rack identifier within the datacenter where the device is mounted | [optional] 
**RackPositionUpperUnit** | Pointer to **NullableInt32** | The upper rack unit position where the device is mounted | [optional] 
**RackPositionLowerUnit** | Pointer to **NullableInt32** | The lower rack unit position where the device is mounted | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the network device | [optional] 
**OrderIndex** | Pointer to **float32** | Order index | [optional] 
**DefaultMtu** | Pointer to **NullableFloat32** | Default MTU | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorizing and filtering network devices | [optional] 
**RequiresOsInstall** | Pointer to **bool** | Indicates if the device requires OS installation before provisioning | [optional] 
**OverwriteWithHostnameFromFetchedSwitch** | Pointer to **NullableBool** | Whether to overwrite the hostname with the one fetched from the device | [optional] 

## Methods

### NewUpdateNetworkDevice

`func NewUpdateNetworkDevice() *UpdateNetworkDevice`

NewUpdateNetworkDevice instantiates a new UpdateNetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDeviceWithDefaults

`func NewUpdateNetworkDeviceWithDefaults() *UpdateNetworkDevice`

NewUpdateNetworkDeviceWithDefaults instantiates a new UpdateNetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *UpdateNetworkDevice) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateNetworkDevice) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateNetworkDevice) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateNetworkDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetDatacenterName

`func (o *UpdateNetworkDevice) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *UpdateNetworkDevice) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *UpdateNetworkDevice) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *UpdateNetworkDevice) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetIdentifierString

`func (o *UpdateNetworkDevice) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *UpdateNetworkDevice) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *UpdateNetworkDevice) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.

### HasIdentifierString

`func (o *UpdateNetworkDevice) HasIdentifierString() bool`

HasIdentifierString returns a boolean if a field has been set.

### GetChassisRackId

`func (o *UpdateNetworkDevice) GetChassisRackId() int32`

GetChassisRackId returns the ChassisRackId field if non-nil, zero value otherwise.

### GetChassisRackIdOk

`func (o *UpdateNetworkDevice) GetChassisRackIdOk() (*int32, bool)`

GetChassisRackIdOk returns a tuple with the ChassisRackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisRackId

`func (o *UpdateNetworkDevice) SetChassisRackId(v int32)`

SetChassisRackId sets ChassisRackId field to given value.

### HasChassisRackId

`func (o *UpdateNetworkDevice) HasChassisRackId() bool`

HasChassisRackId returns a boolean if a field has been set.

### GetChassisIdentifier

`func (o *UpdateNetworkDevice) GetChassisIdentifier() string`

GetChassisIdentifier returns the ChassisIdentifier field if non-nil, zero value otherwise.

### GetChassisIdentifierOk

`func (o *UpdateNetworkDevice) GetChassisIdentifierOk() (*string, bool)`

GetChassisIdentifierOk returns a tuple with the ChassisIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisIdentifier

`func (o *UpdateNetworkDevice) SetChassisIdentifier(v string)`

SetChassisIdentifier sets ChassisIdentifier field to given value.

### HasChassisIdentifier

`func (o *UpdateNetworkDevice) HasChassisIdentifier() bool`

HasChassisIdentifier returns a boolean if a field has been set.

### SetChassisIdentifierNil

`func (o *UpdateNetworkDevice) SetChassisIdentifierNil(b bool)`

 SetChassisIdentifierNil sets the value for ChassisIdentifier to be an explicit nil

### UnsetChassisIdentifier
`func (o *UpdateNetworkDevice) UnsetChassisIdentifier()`

UnsetChassisIdentifier ensures that no value is present for ChassisIdentifier, not even an explicit nil
### GetProvisionerType

`func (o *UpdateNetworkDevice) GetProvisionerType() ProvisionerType`

GetProvisionerType returns the ProvisionerType field if non-nil, zero value otherwise.

### GetProvisionerTypeOk

`func (o *UpdateNetworkDevice) GetProvisionerTypeOk() (*ProvisionerType, bool)`

GetProvisionerTypeOk returns a tuple with the ProvisionerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerType

`func (o *UpdateNetworkDevice) SetProvisionerType(v ProvisionerType)`

SetProvisionerType sets ProvisionerType field to given value.

### HasProvisionerType

`func (o *UpdateNetworkDevice) HasProvisionerType() bool`

HasProvisionerType returns a boolean if a field has been set.

### GetDriver

`func (o *UpdateNetworkDevice) GetDriver() NetworkDeviceDriver`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *UpdateNetworkDevice) GetDriverOk() (*NetworkDeviceDriver, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *UpdateNetworkDevice) SetDriver(v NetworkDeviceDriver)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *UpdateNetworkDevice) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetPosition

`func (o *UpdateNetworkDevice) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateNetworkDevice) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateNetworkDevice) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateNetworkDevice) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTorLinkedId

`func (o *UpdateNetworkDevice) GetTorLinkedId() float32`

GetTorLinkedId returns the TorLinkedId field if non-nil, zero value otherwise.

### GetTorLinkedIdOk

`func (o *UpdateNetworkDevice) GetTorLinkedIdOk() (*float32, bool)`

GetTorLinkedIdOk returns a tuple with the TorLinkedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorLinkedId

`func (o *UpdateNetworkDevice) SetTorLinkedId(v float32)`

SetTorLinkedId sets TorLinkedId field to given value.

### HasTorLinkedId

`func (o *UpdateNetworkDevice) HasTorLinkedId() bool`

HasTorLinkedId returns a boolean if a field has been set.

### SetTorLinkedIdNil

`func (o *UpdateNetworkDevice) SetTorLinkedIdNil(b bool)`

 SetTorLinkedIdNil sets the value for TorLinkedId to be an explicit nil

### UnsetTorLinkedId
`func (o *UpdateNetworkDevice) UnsetTorLinkedId()`

UnsetTorLinkedId ensures that no value is present for TorLinkedId, not even an explicit nil
### GetIsGateway

`func (o *UpdateNetworkDevice) GetIsGateway() bool`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *UpdateNetworkDevice) GetIsGatewayOk() (*bool, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *UpdateNetworkDevice) SetIsGateway(v bool)`

SetIsGateway sets IsGateway field to given value.

### HasIsGateway

`func (o *UpdateNetworkDevice) HasIsGateway() bool`

HasIsGateway returns a boolean if a field has been set.

### GetSyslogEnabled

`func (o *UpdateNetworkDevice) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *UpdateNetworkDevice) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *UpdateNetworkDevice) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *UpdateNetworkDevice) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.

### SetSyslogEnabledNil

`func (o *UpdateNetworkDevice) SetSyslogEnabledNil(b bool)`

 SetSyslogEnabledNil sets the value for SyslogEnabled to be an explicit nil

### UnsetSyslogEnabled
`func (o *UpdateNetworkDevice) UnsetSyslogEnabled()`

UnsetSyslogEnabled ensures that no value is present for SyslogEnabled, not even an explicit nil
### GetIsStorageSwitch

`func (o *UpdateNetworkDevice) GetIsStorageSwitch() bool`

GetIsStorageSwitch returns the IsStorageSwitch field if non-nil, zero value otherwise.

### GetIsStorageSwitchOk

`func (o *UpdateNetworkDevice) GetIsStorageSwitchOk() (*bool, bool)`

GetIsStorageSwitchOk returns a tuple with the IsStorageSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorageSwitch

`func (o *UpdateNetworkDevice) SetIsStorageSwitch(v bool)`

SetIsStorageSwitch sets IsStorageSwitch field to given value.

### HasIsStorageSwitch

`func (o *UpdateNetworkDevice) HasIsStorageSwitch() bool`

HasIsStorageSwitch returns a boolean if a field has been set.

### GetIsBorderDevice

`func (o *UpdateNetworkDevice) GetIsBorderDevice() bool`

GetIsBorderDevice returns the IsBorderDevice field if non-nil, zero value otherwise.

### GetIsBorderDeviceOk

`func (o *UpdateNetworkDevice) GetIsBorderDeviceOk() (*bool, bool)`

GetIsBorderDeviceOk returns a tuple with the IsBorderDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBorderDevice

`func (o *UpdateNetworkDevice) SetIsBorderDevice(v bool)`

SetIsBorderDevice sets IsBorderDevice field to given value.

### HasIsBorderDevice

`func (o *UpdateNetworkDevice) HasIsBorderDevice() bool`

HasIsBorderDevice returns a boolean if a field has been set.

### GetQuarantineSubnetStart

`func (o *UpdateNetworkDevice) GetQuarantineSubnetStart() string`

GetQuarantineSubnetStart returns the QuarantineSubnetStart field if non-nil, zero value otherwise.

### GetQuarantineSubnetStartOk

`func (o *UpdateNetworkDevice) GetQuarantineSubnetStartOk() (*string, bool)`

GetQuarantineSubnetStartOk returns a tuple with the QuarantineSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetStart

`func (o *UpdateNetworkDevice) SetQuarantineSubnetStart(v string)`

SetQuarantineSubnetStart sets QuarantineSubnetStart field to given value.

### HasQuarantineSubnetStart

`func (o *UpdateNetworkDevice) HasQuarantineSubnetStart() bool`

HasQuarantineSubnetStart returns a boolean if a field has been set.

### SetQuarantineSubnetStartNil

`func (o *UpdateNetworkDevice) SetQuarantineSubnetStartNil(b bool)`

 SetQuarantineSubnetStartNil sets the value for QuarantineSubnetStart to be an explicit nil

### UnsetQuarantineSubnetStart
`func (o *UpdateNetworkDevice) UnsetQuarantineSubnetStart()`

UnsetQuarantineSubnetStart ensures that no value is present for QuarantineSubnetStart, not even an explicit nil
### GetQuarantineSubnetEnd

`func (o *UpdateNetworkDevice) GetQuarantineSubnetEnd() string`

GetQuarantineSubnetEnd returns the QuarantineSubnetEnd field if non-nil, zero value otherwise.

### GetQuarantineSubnetEndOk

`func (o *UpdateNetworkDevice) GetQuarantineSubnetEndOk() (*string, bool)`

GetQuarantineSubnetEndOk returns a tuple with the QuarantineSubnetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetEnd

`func (o *UpdateNetworkDevice) SetQuarantineSubnetEnd(v string)`

SetQuarantineSubnetEnd sets QuarantineSubnetEnd field to given value.

### HasQuarantineSubnetEnd

`func (o *UpdateNetworkDevice) HasQuarantineSubnetEnd() bool`

HasQuarantineSubnetEnd returns a boolean if a field has been set.

### SetQuarantineSubnetEndNil

`func (o *UpdateNetworkDevice) SetQuarantineSubnetEndNil(b bool)`

 SetQuarantineSubnetEndNil sets the value for QuarantineSubnetEnd to be an explicit nil

### UnsetQuarantineSubnetEnd
`func (o *UpdateNetworkDevice) UnsetQuarantineSubnetEnd()`

UnsetQuarantineSubnetEnd ensures that no value is present for QuarantineSubnetEnd, not even an explicit nil
### GetQuarantineSubnetPrefixSize

`func (o *UpdateNetworkDevice) GetQuarantineSubnetPrefixSize() int32`

GetQuarantineSubnetPrefixSize returns the QuarantineSubnetPrefixSize field if non-nil, zero value otherwise.

### GetQuarantineSubnetPrefixSizeOk

`func (o *UpdateNetworkDevice) GetQuarantineSubnetPrefixSizeOk() (*int32, bool)`

GetQuarantineSubnetPrefixSizeOk returns a tuple with the QuarantineSubnetPrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetPrefixSize

`func (o *UpdateNetworkDevice) SetQuarantineSubnetPrefixSize(v int32)`

SetQuarantineSubnetPrefixSize sets QuarantineSubnetPrefixSize field to given value.

### HasQuarantineSubnetPrefixSize

`func (o *UpdateNetworkDevice) HasQuarantineSubnetPrefixSize() bool`

HasQuarantineSubnetPrefixSize returns a boolean if a field has been set.

### SetQuarantineSubnetPrefixSizeNil

`func (o *UpdateNetworkDevice) SetQuarantineSubnetPrefixSizeNil(b bool)`

 SetQuarantineSubnetPrefixSizeNil sets the value for QuarantineSubnetPrefixSize to be an explicit nil

### UnsetQuarantineSubnetPrefixSize
`func (o *UpdateNetworkDevice) UnsetQuarantineSubnetPrefixSize()`

UnsetQuarantineSubnetPrefixSize ensures that no value is present for QuarantineSubnetPrefixSize, not even an explicit nil
### GetQuarantineSubnetGateway

`func (o *UpdateNetworkDevice) GetQuarantineSubnetGateway() string`

GetQuarantineSubnetGateway returns the QuarantineSubnetGateway field if non-nil, zero value otherwise.

### GetQuarantineSubnetGatewayOk

`func (o *UpdateNetworkDevice) GetQuarantineSubnetGatewayOk() (*string, bool)`

GetQuarantineSubnetGatewayOk returns a tuple with the QuarantineSubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetGateway

`func (o *UpdateNetworkDevice) SetQuarantineSubnetGateway(v string)`

SetQuarantineSubnetGateway sets QuarantineSubnetGateway field to given value.

### HasQuarantineSubnetGateway

`func (o *UpdateNetworkDevice) HasQuarantineSubnetGateway() bool`

HasQuarantineSubnetGateway returns a boolean if a field has been set.

### SetQuarantineSubnetGatewayNil

`func (o *UpdateNetworkDevice) SetQuarantineSubnetGatewayNil(b bool)`

 SetQuarantineSubnetGatewayNil sets the value for QuarantineSubnetGateway to be an explicit nil

### UnsetQuarantineSubnetGateway
`func (o *UpdateNetworkDevice) UnsetQuarantineSubnetGateway()`

UnsetQuarantineSubnetGateway ensures that no value is present for QuarantineSubnetGateway, not even an explicit nil
### GetQuarantineVlan

`func (o *UpdateNetworkDevice) GetQuarantineVlan() int32`

GetQuarantineVlan returns the QuarantineVlan field if non-nil, zero value otherwise.

### GetQuarantineVlanOk

`func (o *UpdateNetworkDevice) GetQuarantineVlanOk() (*int32, bool)`

GetQuarantineVlanOk returns a tuple with the QuarantineVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineVlan

`func (o *UpdateNetworkDevice) SetQuarantineVlan(v int32)`

SetQuarantineVlan sets QuarantineVlan field to given value.

### HasQuarantineVlan

`func (o *UpdateNetworkDevice) HasQuarantineVlan() bool`

HasQuarantineVlan returns a boolean if a field has been set.

### GetManagementProtocol

`func (o *UpdateNetworkDevice) GetManagementProtocol() string`

GetManagementProtocol returns the ManagementProtocol field if non-nil, zero value otherwise.

### GetManagementProtocolOk

`func (o *UpdateNetworkDevice) GetManagementProtocolOk() (*string, bool)`

GetManagementProtocolOk returns a tuple with the ManagementProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementProtocol

`func (o *UpdateNetworkDevice) SetManagementProtocol(v string)`

SetManagementProtocol sets ManagementProtocol field to given value.

### HasManagementProtocol

`func (o *UpdateNetworkDevice) HasManagementProtocol() bool`

HasManagementProtocol returns a boolean if a field has been set.

### SetManagementProtocolNil

`func (o *UpdateNetworkDevice) SetManagementProtocolNil(b bool)`

 SetManagementProtocolNil sets the value for ManagementProtocol to be an explicit nil

### UnsetManagementProtocol
`func (o *UpdateNetworkDevice) UnsetManagementProtocol()`

UnsetManagementProtocol ensures that no value is present for ManagementProtocol, not even an explicit nil
### GetManagementAddress

`func (o *UpdateNetworkDevice) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *UpdateNetworkDevice) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *UpdateNetworkDevice) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *UpdateNetworkDevice) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### SetManagementAddressNil

`func (o *UpdateNetworkDevice) SetManagementAddressNil(b bool)`

 SetManagementAddressNil sets the value for ManagementAddress to be an explicit nil

### UnsetManagementAddress
`func (o *UpdateNetworkDevice) UnsetManagementAddress()`

UnsetManagementAddress ensures that no value is present for ManagementAddress, not even an explicit nil
### GetManagementPort

`func (o *UpdateNetworkDevice) GetManagementPort() int32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *UpdateNetworkDevice) GetManagementPortOk() (*int32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *UpdateNetworkDevice) SetManagementPort(v int32)`

SetManagementPort sets ManagementPort field to given value.

### HasManagementPort

`func (o *UpdateNetworkDevice) HasManagementPort() bool`

HasManagementPort returns a boolean if a field has been set.

### SetManagementPortNil

`func (o *UpdateNetworkDevice) SetManagementPortNil(b bool)`

 SetManagementPortNil sets the value for ManagementPort to be an explicit nil

### UnsetManagementPort
`func (o *UpdateNetworkDevice) UnsetManagementPort()`

UnsetManagementPort ensures that no value is present for ManagementPort, not even an explicit nil
### GetUsername

`func (o *UpdateNetworkDevice) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateNetworkDevice) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateNetworkDevice) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateNetworkDevice) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateNetworkDevice) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateNetworkDevice) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetManagementPassword

`func (o *UpdateNetworkDevice) GetManagementPassword() string`

GetManagementPassword returns the ManagementPassword field if non-nil, zero value otherwise.

### GetManagementPasswordOk

`func (o *UpdateNetworkDevice) GetManagementPasswordOk() (*string, bool)`

GetManagementPasswordOk returns a tuple with the ManagementPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPassword

`func (o *UpdateNetworkDevice) SetManagementPassword(v string)`

SetManagementPassword sets ManagementPassword field to given value.

### HasManagementPassword

`func (o *UpdateNetworkDevice) HasManagementPassword() bool`

HasManagementPassword returns a boolean if a field has been set.

### SetManagementPasswordNil

`func (o *UpdateNetworkDevice) SetManagementPasswordNil(b bool)`

 SetManagementPasswordNil sets the value for ManagementPassword to be an explicit nil

### UnsetManagementPassword
`func (o *UpdateNetworkDevice) UnsetManagementPassword()`

UnsetManagementPassword ensures that no value is present for ManagementPassword, not even an explicit nil
### GetManagementAddressGateway

`func (o *UpdateNetworkDevice) GetManagementAddressGateway() string`

GetManagementAddressGateway returns the ManagementAddressGateway field if non-nil, zero value otherwise.

### GetManagementAddressGatewayOk

`func (o *UpdateNetworkDevice) GetManagementAddressGatewayOk() (*string, bool)`

GetManagementAddressGatewayOk returns a tuple with the ManagementAddressGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddressGateway

`func (o *UpdateNetworkDevice) SetManagementAddressGateway(v string)`

SetManagementAddressGateway sets ManagementAddressGateway field to given value.

### HasManagementAddressGateway

`func (o *UpdateNetworkDevice) HasManagementAddressGateway() bool`

HasManagementAddressGateway returns a boolean if a field has been set.

### SetManagementAddressGatewayNil

`func (o *UpdateNetworkDevice) SetManagementAddressGatewayNil(b bool)`

 SetManagementAddressGatewayNil sets the value for ManagementAddressGateway to be an explicit nil

### UnsetManagementAddressGateway
`func (o *UpdateNetworkDevice) UnsetManagementAddressGateway()`

UnsetManagementAddressGateway ensures that no value is present for ManagementAddressGateway, not even an explicit nil
### GetManagementAddressMask

`func (o *UpdateNetworkDevice) GetManagementAddressMask() string`

GetManagementAddressMask returns the ManagementAddressMask field if non-nil, zero value otherwise.

### GetManagementAddressMaskOk

`func (o *UpdateNetworkDevice) GetManagementAddressMaskOk() (*string, bool)`

GetManagementAddressMaskOk returns a tuple with the ManagementAddressMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddressMask

`func (o *UpdateNetworkDevice) SetManagementAddressMask(v string)`

SetManagementAddressMask sets ManagementAddressMask field to given value.

### HasManagementAddressMask

`func (o *UpdateNetworkDevice) HasManagementAddressMask() bool`

HasManagementAddressMask returns a boolean if a field has been set.

### SetManagementAddressMaskNil

`func (o *UpdateNetworkDevice) SetManagementAddressMaskNil(b bool)`

 SetManagementAddressMaskNil sets the value for ManagementAddressMask to be an explicit nil

### UnsetManagementAddressMask
`func (o *UpdateNetworkDevice) UnsetManagementAddressMask()`

UnsetManagementAddressMask ensures that no value is present for ManagementAddressMask, not even an explicit nil
### GetManagementMAC

`func (o *UpdateNetworkDevice) GetManagementMAC() string`

GetManagementMAC returns the ManagementMAC field if non-nil, zero value otherwise.

### GetManagementMACOk

`func (o *UpdateNetworkDevice) GetManagementMACOk() (*string, bool)`

GetManagementMACOk returns a tuple with the ManagementMAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMAC

`func (o *UpdateNetworkDevice) SetManagementMAC(v string)`

SetManagementMAC sets ManagementMAC field to given value.

### HasManagementMAC

`func (o *UpdateNetworkDevice) HasManagementMAC() bool`

HasManagementMAC returns a boolean if a field has been set.

### SetManagementMACNil

`func (o *UpdateNetworkDevice) SetManagementMACNil(b bool)`

 SetManagementMACNil sets the value for ManagementMAC to be an explicit nil

### UnsetManagementMAC
`func (o *UpdateNetworkDevice) UnsetManagementMAC()`

UnsetManagementMAC ensures that no value is present for ManagementMAC, not even an explicit nil
### GetLoopbackAddress

`func (o *UpdateNetworkDevice) GetLoopbackAddress() string`

GetLoopbackAddress returns the LoopbackAddress field if non-nil, zero value otherwise.

### GetLoopbackAddressOk

`func (o *UpdateNetworkDevice) GetLoopbackAddressOk() (*string, bool)`

GetLoopbackAddressOk returns a tuple with the LoopbackAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddress

`func (o *UpdateNetworkDevice) SetLoopbackAddress(v string)`

SetLoopbackAddress sets LoopbackAddress field to given value.

### HasLoopbackAddress

`func (o *UpdateNetworkDevice) HasLoopbackAddress() bool`

HasLoopbackAddress returns a boolean if a field has been set.

### SetLoopbackAddressNil

`func (o *UpdateNetworkDevice) SetLoopbackAddressNil(b bool)`

 SetLoopbackAddressNil sets the value for LoopbackAddress to be an explicit nil

### UnsetLoopbackAddress
`func (o *UpdateNetworkDevice) UnsetLoopbackAddress()`

UnsetLoopbackAddress ensures that no value is present for LoopbackAddress, not even an explicit nil
### GetVtepAddress

`func (o *UpdateNetworkDevice) GetVtepAddress() string`

GetVtepAddress returns the VtepAddress field if non-nil, zero value otherwise.

### GetVtepAddressOk

`func (o *UpdateNetworkDevice) GetVtepAddressOk() (*string, bool)`

GetVtepAddressOk returns a tuple with the VtepAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddress

`func (o *UpdateNetworkDevice) SetVtepAddress(v string)`

SetVtepAddress sets VtepAddress field to given value.

### HasVtepAddress

`func (o *UpdateNetworkDevice) HasVtepAddress() bool`

HasVtepAddress returns a boolean if a field has been set.

### SetVtepAddressNil

`func (o *UpdateNetworkDevice) SetVtepAddressNil(b bool)`

 SetVtepAddressNil sets the value for VtepAddress to be an explicit nil

### UnsetVtepAddress
`func (o *UpdateNetworkDevice) UnsetVtepAddress()`

UnsetVtepAddress ensures that no value is present for VtepAddress, not even an explicit nil
### GetAsn

`func (o *UpdateNetworkDevice) GetAsn() float32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *UpdateNetworkDevice) GetAsnOk() (*float32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *UpdateNetworkDevice) SetAsn(v float32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *UpdateNetworkDevice) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *UpdateNetworkDevice) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *UpdateNetworkDevice) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetNetworkTypesAllowed

`func (o *UpdateNetworkDevice) GetNetworkTypesAllowed() []string`

GetNetworkTypesAllowed returns the NetworkTypesAllowed field if non-nil, zero value otherwise.

### GetNetworkTypesAllowedOk

`func (o *UpdateNetworkDevice) GetNetworkTypesAllowedOk() (*[]string, bool)`

GetNetworkTypesAllowedOk returns a tuple with the NetworkTypesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypesAllowed

`func (o *UpdateNetworkDevice) SetNetworkTypesAllowed(v []string)`

SetNetworkTypesAllowed sets NetworkTypesAllowed field to given value.

### HasNetworkTypesAllowed

`func (o *UpdateNetworkDevice) HasNetworkTypesAllowed() bool`

HasNetworkTypesAllowed returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateNetworkDevice) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateNetworkDevice) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCountry

`func (o *UpdateNetworkDevice) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateNetworkDevice) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateNetworkDevice) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateNetworkDevice) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *UpdateNetworkDevice) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *UpdateNetworkDevice) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCity

`func (o *UpdateNetworkDevice) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateNetworkDevice) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateNetworkDevice) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateNetworkDevice) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *UpdateNetworkDevice) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *UpdateNetworkDevice) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDatacenterMeta

`func (o *UpdateNetworkDevice) GetDatacenterMeta() string`

GetDatacenterMeta returns the DatacenterMeta field if non-nil, zero value otherwise.

### GetDatacenterMetaOk

`func (o *UpdateNetworkDevice) GetDatacenterMetaOk() (*string, bool)`

GetDatacenterMetaOk returns a tuple with the DatacenterMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterMeta

`func (o *UpdateNetworkDevice) SetDatacenterMeta(v string)`

SetDatacenterMeta sets DatacenterMeta field to given value.

### HasDatacenterMeta

`func (o *UpdateNetworkDevice) HasDatacenterMeta() bool`

HasDatacenterMeta returns a boolean if a field has been set.

### SetDatacenterMetaNil

`func (o *UpdateNetworkDevice) SetDatacenterMetaNil(b bool)`

 SetDatacenterMetaNil sets the value for DatacenterMeta to be an explicit nil

### UnsetDatacenterMeta
`func (o *UpdateNetworkDevice) UnsetDatacenterMeta()`

UnsetDatacenterMeta ensures that no value is present for DatacenterMeta, not even an explicit nil
### GetDatacenterRoom

`func (o *UpdateNetworkDevice) GetDatacenterRoom() string`

GetDatacenterRoom returns the DatacenterRoom field if non-nil, zero value otherwise.

### GetDatacenterRoomOk

`func (o *UpdateNetworkDevice) GetDatacenterRoomOk() (*string, bool)`

GetDatacenterRoomOk returns a tuple with the DatacenterRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterRoom

`func (o *UpdateNetworkDevice) SetDatacenterRoom(v string)`

SetDatacenterRoom sets DatacenterRoom field to given value.

### HasDatacenterRoom

`func (o *UpdateNetworkDevice) HasDatacenterRoom() bool`

HasDatacenterRoom returns a boolean if a field has been set.

### SetDatacenterRoomNil

`func (o *UpdateNetworkDevice) SetDatacenterRoomNil(b bool)`

 SetDatacenterRoomNil sets the value for DatacenterRoom to be an explicit nil

### UnsetDatacenterRoom
`func (o *UpdateNetworkDevice) UnsetDatacenterRoom()`

UnsetDatacenterRoom ensures that no value is present for DatacenterRoom, not even an explicit nil
### GetDatacenterRack

`func (o *UpdateNetworkDevice) GetDatacenterRack() string`

GetDatacenterRack returns the DatacenterRack field if non-nil, zero value otherwise.

### GetDatacenterRackOk

`func (o *UpdateNetworkDevice) GetDatacenterRackOk() (*string, bool)`

GetDatacenterRackOk returns a tuple with the DatacenterRack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterRack

`func (o *UpdateNetworkDevice) SetDatacenterRack(v string)`

SetDatacenterRack sets DatacenterRack field to given value.

### HasDatacenterRack

`func (o *UpdateNetworkDevice) HasDatacenterRack() bool`

HasDatacenterRack returns a boolean if a field has been set.

### SetDatacenterRackNil

`func (o *UpdateNetworkDevice) SetDatacenterRackNil(b bool)`

 SetDatacenterRackNil sets the value for DatacenterRack to be an explicit nil

### UnsetDatacenterRack
`func (o *UpdateNetworkDevice) UnsetDatacenterRack()`

UnsetDatacenterRack ensures that no value is present for DatacenterRack, not even an explicit nil
### GetRackPositionUpperUnit

`func (o *UpdateNetworkDevice) GetRackPositionUpperUnit() int32`

GetRackPositionUpperUnit returns the RackPositionUpperUnit field if non-nil, zero value otherwise.

### GetRackPositionUpperUnitOk

`func (o *UpdateNetworkDevice) GetRackPositionUpperUnitOk() (*int32, bool)`

GetRackPositionUpperUnitOk returns a tuple with the RackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionUpperUnit

`func (o *UpdateNetworkDevice) SetRackPositionUpperUnit(v int32)`

SetRackPositionUpperUnit sets RackPositionUpperUnit field to given value.

### HasRackPositionUpperUnit

`func (o *UpdateNetworkDevice) HasRackPositionUpperUnit() bool`

HasRackPositionUpperUnit returns a boolean if a field has been set.

### SetRackPositionUpperUnitNil

`func (o *UpdateNetworkDevice) SetRackPositionUpperUnitNil(b bool)`

 SetRackPositionUpperUnitNil sets the value for RackPositionUpperUnit to be an explicit nil

### UnsetRackPositionUpperUnit
`func (o *UpdateNetworkDevice) UnsetRackPositionUpperUnit()`

UnsetRackPositionUpperUnit ensures that no value is present for RackPositionUpperUnit, not even an explicit nil
### GetRackPositionLowerUnit

`func (o *UpdateNetworkDevice) GetRackPositionLowerUnit() int32`

GetRackPositionLowerUnit returns the RackPositionLowerUnit field if non-nil, zero value otherwise.

### GetRackPositionLowerUnitOk

`func (o *UpdateNetworkDevice) GetRackPositionLowerUnitOk() (*int32, bool)`

GetRackPositionLowerUnitOk returns a tuple with the RackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionLowerUnit

`func (o *UpdateNetworkDevice) SetRackPositionLowerUnit(v int32)`

SetRackPositionLowerUnit sets RackPositionLowerUnit field to given value.

### HasRackPositionLowerUnit

`func (o *UpdateNetworkDevice) HasRackPositionLowerUnit() bool`

HasRackPositionLowerUnit returns a boolean if a field has been set.

### SetRackPositionLowerUnitNil

`func (o *UpdateNetworkDevice) SetRackPositionLowerUnitNil(b bool)`

 SetRackPositionLowerUnitNil sets the value for RackPositionLowerUnit to be an explicit nil

### UnsetRackPositionLowerUnit
`func (o *UpdateNetworkDevice) UnsetRackPositionLowerUnit()`

UnsetRackPositionLowerUnit ensures that no value is present for RackPositionLowerUnit, not even an explicit nil
### GetSerialNumber

`func (o *UpdateNetworkDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *UpdateNetworkDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *UpdateNetworkDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *UpdateNetworkDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetOrderIndex

`func (o *UpdateNetworkDevice) GetOrderIndex() float32`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *UpdateNetworkDevice) GetOrderIndexOk() (*float32, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *UpdateNetworkDevice) SetOrderIndex(v float32)`

SetOrderIndex sets OrderIndex field to given value.

### HasOrderIndex

`func (o *UpdateNetworkDevice) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.

### GetDefaultMtu

`func (o *UpdateNetworkDevice) GetDefaultMtu() float32`

GetDefaultMtu returns the DefaultMtu field if non-nil, zero value otherwise.

### GetDefaultMtuOk

`func (o *UpdateNetworkDevice) GetDefaultMtuOk() (*float32, bool)`

GetDefaultMtuOk returns a tuple with the DefaultMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtu

`func (o *UpdateNetworkDevice) SetDefaultMtu(v float32)`

SetDefaultMtu sets DefaultMtu field to given value.

### HasDefaultMtu

`func (o *UpdateNetworkDevice) HasDefaultMtu() bool`

HasDefaultMtu returns a boolean if a field has been set.

### SetDefaultMtuNil

`func (o *UpdateNetworkDevice) SetDefaultMtuNil(b bool)`

 SetDefaultMtuNil sets the value for DefaultMtu to be an explicit nil

### UnsetDefaultMtu
`func (o *UpdateNetworkDevice) UnsetDefaultMtu()`

UnsetDefaultMtu ensures that no value is present for DefaultMtu, not even an explicit nil
### GetTags

`func (o *UpdateNetworkDevice) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateNetworkDevice) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateNetworkDevice) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateNetworkDevice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateNetworkDevice) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateNetworkDevice) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetRequiresOsInstall

`func (o *UpdateNetworkDevice) GetRequiresOsInstall() bool`

GetRequiresOsInstall returns the RequiresOsInstall field if non-nil, zero value otherwise.

### GetRequiresOsInstallOk

`func (o *UpdateNetworkDevice) GetRequiresOsInstallOk() (*bool, bool)`

GetRequiresOsInstallOk returns a tuple with the RequiresOsInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOsInstall

`func (o *UpdateNetworkDevice) SetRequiresOsInstall(v bool)`

SetRequiresOsInstall sets RequiresOsInstall field to given value.

### HasRequiresOsInstall

`func (o *UpdateNetworkDevice) HasRequiresOsInstall() bool`

HasRequiresOsInstall returns a boolean if a field has been set.

### GetOverwriteWithHostnameFromFetchedSwitch

`func (o *UpdateNetworkDevice) GetOverwriteWithHostnameFromFetchedSwitch() bool`

GetOverwriteWithHostnameFromFetchedSwitch returns the OverwriteWithHostnameFromFetchedSwitch field if non-nil, zero value otherwise.

### GetOverwriteWithHostnameFromFetchedSwitchOk

`func (o *UpdateNetworkDevice) GetOverwriteWithHostnameFromFetchedSwitchOk() (*bool, bool)`

GetOverwriteWithHostnameFromFetchedSwitchOk returns a tuple with the OverwriteWithHostnameFromFetchedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteWithHostnameFromFetchedSwitch

`func (o *UpdateNetworkDevice) SetOverwriteWithHostnameFromFetchedSwitch(v bool)`

SetOverwriteWithHostnameFromFetchedSwitch sets OverwriteWithHostnameFromFetchedSwitch field to given value.

### HasOverwriteWithHostnameFromFetchedSwitch

`func (o *UpdateNetworkDevice) HasOverwriteWithHostnameFromFetchedSwitch() bool`

HasOverwriteWithHostnameFromFetchedSwitch returns a boolean if a field has been set.

### SetOverwriteWithHostnameFromFetchedSwitchNil

`func (o *UpdateNetworkDevice) SetOverwriteWithHostnameFromFetchedSwitchNil(b bool)`

 SetOverwriteWithHostnameFromFetchedSwitchNil sets the value for OverwriteWithHostnameFromFetchedSwitch to be an explicit nil

### UnsetOverwriteWithHostnameFromFetchedSwitch
`func (o *UpdateNetworkDevice) UnsetOverwriteWithHostnameFromFetchedSwitch()`

UnsetOverwriteWithHostnameFromFetchedSwitch ensures that no value is present for OverwriteWithHostnameFromFetchedSwitch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


