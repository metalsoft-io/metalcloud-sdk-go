# CreateNetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **int32** | Site identifier | [optional] 
**DatacenterName** | Pointer to **string** | Name of the datacenter | [optional] 
**IdentifierString** | Pointer to **string** | Unique identifier string for the network device | [optional] 
**ChassisRackId** | Pointer to **int32** | ID of the rack where the network device chassis is installed | [optional] 
**ChassisIdentifier** | Pointer to **NullableString** | Unique identifier for the network device chassis | [optional] 
**ProvisionerType** | Pointer to [**ProvisionerType**](ProvisionerType.md) | Type of provisioner used for this network device | [optional] 
**Driver** | [**NetworkDeviceDriver**](NetworkDeviceDriver.md) | Driver software used to communicate with the network device | 
**Position** | **string** | The physical or logical position of the network device in the network topology. | 
**TorLinkedId** | Pointer to **NullableFloat32** | ID of the Top-of-Rack (TOR) switch that this network device is linked to. Used for establishing hierarchical relationships between network devices. | [optional] 
**IsGateway** | Pointer to **bool** | Indicates if this network device acts as a gateway for external network traffic | [optional] 
**SyslogEnabled** | Pointer to **NullableBool** | Indicates if syslog logging is enabled for this network device | [optional] 
**IsStorageSwitch** | Pointer to **bool** | Indicates if this network device is dedicated to storage traffic | [optional] 
**IsBorderDevice** | Pointer to **bool** | Indicates if this network device functions as a border device connecting to external networks | [optional] 
**QuarantineSubnetStart** | Pointer to **NullableString** | The starting IP address of the quarantine subnet range | [optional] 
**QuarantineSubnetEnd** | Pointer to **NullableString** | The ending IP address of the quarantine subnet range | [optional] 
**QuarantineSubnetPrefixSize** | Pointer to **NullableInt32** | The prefix size for the quarantine subnet | [optional] 
**QuarantineSubnetGateway** | Pointer to **NullableString** | The gateway IP address for the quarantine subnet | [optional] 
**QuarantineVlan** | Pointer to **int32** | The VLAN ID for the quarantine network | [optional] 
**ManagementProtocol** | Pointer to **NullableString** | The protocol used to manage the network device (e.g., SSH, HTTPS) | [optional] 
**ManagementAddress** | Pointer to **NullableString** | The IP address used to manage the network device | [optional] 
**ManagementPort** | Pointer to **NullableInt32** | The port number used for management connections | [optional] 
**Username** | Pointer to **NullableString** | The username used for management authentication | [optional] 
**ManagementPassword** | Pointer to **NullableString** | The password used for management authentication | [optional] 
**ManagementAddressGateway** | Pointer to **NullableString** | The gateway IP address for the management network | [optional] 
**ManagementAddressMask** | Pointer to **NullableString** | The subnet mask for the management network | [optional] 
**ManagementMAC** | Pointer to **NullableString** | The MAC address of the management interface | [optional] 
**LoopbackAddress** | Pointer to **NullableString** | The loopback IP address assigned to the network device | [optional] 
**VtepAddress** | Pointer to **NullableString** | The VTEP (VXLAN Tunnel Endpoint) IP address for overlay networking | [optional] 
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
**OrderIndex** | Pointer to **float32** | The order index for sorting network devices | [optional] 
**DefaultMtu** | Pointer to **NullableFloat32** | The default Maximum Transmission Unit (MTU) for interfaces on this device | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorizing and filtering network devices | [optional] 
**RequiresOsInstall** | Pointer to **bool** | Indicates if the device requires OS installation before provisioning | [optional] 
**OobMgmtIpv4Address** | Pointer to **NullableString** | The IPv4 address for out-of-band management | [optional] 
**OverwriteWithHostnameFromFetchedSwitch** | Pointer to **NullableBool** | Whether to overwrite the hostname with the one fetched from the device | [optional] 

## Methods

### NewCreateNetworkDevice

`func NewCreateNetworkDevice(driver NetworkDeviceDriver, position string, ) *CreateNetworkDevice`

NewCreateNetworkDevice instantiates a new CreateNetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceWithDefaults

`func NewCreateNetworkDeviceWithDefaults() *CreateNetworkDevice`

NewCreateNetworkDeviceWithDefaults instantiates a new CreateNetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *CreateNetworkDevice) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateNetworkDevice) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateNetworkDevice) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CreateNetworkDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetDatacenterName

`func (o *CreateNetworkDevice) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *CreateNetworkDevice) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *CreateNetworkDevice) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *CreateNetworkDevice) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetIdentifierString

`func (o *CreateNetworkDevice) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *CreateNetworkDevice) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *CreateNetworkDevice) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.

### HasIdentifierString

`func (o *CreateNetworkDevice) HasIdentifierString() bool`

HasIdentifierString returns a boolean if a field has been set.

### GetChassisRackId

`func (o *CreateNetworkDevice) GetChassisRackId() int32`

GetChassisRackId returns the ChassisRackId field if non-nil, zero value otherwise.

### GetChassisRackIdOk

`func (o *CreateNetworkDevice) GetChassisRackIdOk() (*int32, bool)`

GetChassisRackIdOk returns a tuple with the ChassisRackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisRackId

`func (o *CreateNetworkDevice) SetChassisRackId(v int32)`

SetChassisRackId sets ChassisRackId field to given value.

### HasChassisRackId

`func (o *CreateNetworkDevice) HasChassisRackId() bool`

HasChassisRackId returns a boolean if a field has been set.

### GetChassisIdentifier

`func (o *CreateNetworkDevice) GetChassisIdentifier() string`

GetChassisIdentifier returns the ChassisIdentifier field if non-nil, zero value otherwise.

### GetChassisIdentifierOk

`func (o *CreateNetworkDevice) GetChassisIdentifierOk() (*string, bool)`

GetChassisIdentifierOk returns a tuple with the ChassisIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisIdentifier

`func (o *CreateNetworkDevice) SetChassisIdentifier(v string)`

SetChassisIdentifier sets ChassisIdentifier field to given value.

### HasChassisIdentifier

`func (o *CreateNetworkDevice) HasChassisIdentifier() bool`

HasChassisIdentifier returns a boolean if a field has been set.

### SetChassisIdentifierNil

`func (o *CreateNetworkDevice) SetChassisIdentifierNil(b bool)`

 SetChassisIdentifierNil sets the value for ChassisIdentifier to be an explicit nil

### UnsetChassisIdentifier
`func (o *CreateNetworkDevice) UnsetChassisIdentifier()`

UnsetChassisIdentifier ensures that no value is present for ChassisIdentifier, not even an explicit nil
### GetProvisionerType

`func (o *CreateNetworkDevice) GetProvisionerType() ProvisionerType`

GetProvisionerType returns the ProvisionerType field if non-nil, zero value otherwise.

### GetProvisionerTypeOk

`func (o *CreateNetworkDevice) GetProvisionerTypeOk() (*ProvisionerType, bool)`

GetProvisionerTypeOk returns a tuple with the ProvisionerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerType

`func (o *CreateNetworkDevice) SetProvisionerType(v ProvisionerType)`

SetProvisionerType sets ProvisionerType field to given value.

### HasProvisionerType

`func (o *CreateNetworkDevice) HasProvisionerType() bool`

HasProvisionerType returns a boolean if a field has been set.

### GetDriver

`func (o *CreateNetworkDevice) GetDriver() NetworkDeviceDriver`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *CreateNetworkDevice) GetDriverOk() (*NetworkDeviceDriver, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *CreateNetworkDevice) SetDriver(v NetworkDeviceDriver)`

SetDriver sets Driver field to given value.


### GetPosition

`func (o *CreateNetworkDevice) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateNetworkDevice) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateNetworkDevice) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetTorLinkedId

`func (o *CreateNetworkDevice) GetTorLinkedId() float32`

GetTorLinkedId returns the TorLinkedId field if non-nil, zero value otherwise.

### GetTorLinkedIdOk

`func (o *CreateNetworkDevice) GetTorLinkedIdOk() (*float32, bool)`

GetTorLinkedIdOk returns a tuple with the TorLinkedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorLinkedId

`func (o *CreateNetworkDevice) SetTorLinkedId(v float32)`

SetTorLinkedId sets TorLinkedId field to given value.

### HasTorLinkedId

`func (o *CreateNetworkDevice) HasTorLinkedId() bool`

HasTorLinkedId returns a boolean if a field has been set.

### SetTorLinkedIdNil

`func (o *CreateNetworkDevice) SetTorLinkedIdNil(b bool)`

 SetTorLinkedIdNil sets the value for TorLinkedId to be an explicit nil

### UnsetTorLinkedId
`func (o *CreateNetworkDevice) UnsetTorLinkedId()`

UnsetTorLinkedId ensures that no value is present for TorLinkedId, not even an explicit nil
### GetIsGateway

`func (o *CreateNetworkDevice) GetIsGateway() bool`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *CreateNetworkDevice) GetIsGatewayOk() (*bool, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *CreateNetworkDevice) SetIsGateway(v bool)`

SetIsGateway sets IsGateway field to given value.

### HasIsGateway

`func (o *CreateNetworkDevice) HasIsGateway() bool`

HasIsGateway returns a boolean if a field has been set.

### GetSyslogEnabled

`func (o *CreateNetworkDevice) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *CreateNetworkDevice) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *CreateNetworkDevice) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *CreateNetworkDevice) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.

### SetSyslogEnabledNil

`func (o *CreateNetworkDevice) SetSyslogEnabledNil(b bool)`

 SetSyslogEnabledNil sets the value for SyslogEnabled to be an explicit nil

### UnsetSyslogEnabled
`func (o *CreateNetworkDevice) UnsetSyslogEnabled()`

UnsetSyslogEnabled ensures that no value is present for SyslogEnabled, not even an explicit nil
### GetIsStorageSwitch

`func (o *CreateNetworkDevice) GetIsStorageSwitch() bool`

GetIsStorageSwitch returns the IsStorageSwitch field if non-nil, zero value otherwise.

### GetIsStorageSwitchOk

`func (o *CreateNetworkDevice) GetIsStorageSwitchOk() (*bool, bool)`

GetIsStorageSwitchOk returns a tuple with the IsStorageSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorageSwitch

`func (o *CreateNetworkDevice) SetIsStorageSwitch(v bool)`

SetIsStorageSwitch sets IsStorageSwitch field to given value.

### HasIsStorageSwitch

`func (o *CreateNetworkDevice) HasIsStorageSwitch() bool`

HasIsStorageSwitch returns a boolean if a field has been set.

### GetIsBorderDevice

`func (o *CreateNetworkDevice) GetIsBorderDevice() bool`

GetIsBorderDevice returns the IsBorderDevice field if non-nil, zero value otherwise.

### GetIsBorderDeviceOk

`func (o *CreateNetworkDevice) GetIsBorderDeviceOk() (*bool, bool)`

GetIsBorderDeviceOk returns a tuple with the IsBorderDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBorderDevice

`func (o *CreateNetworkDevice) SetIsBorderDevice(v bool)`

SetIsBorderDevice sets IsBorderDevice field to given value.

### HasIsBorderDevice

`func (o *CreateNetworkDevice) HasIsBorderDevice() bool`

HasIsBorderDevice returns a boolean if a field has been set.

### GetQuarantineSubnetStart

`func (o *CreateNetworkDevice) GetQuarantineSubnetStart() string`

GetQuarantineSubnetStart returns the QuarantineSubnetStart field if non-nil, zero value otherwise.

### GetQuarantineSubnetStartOk

`func (o *CreateNetworkDevice) GetQuarantineSubnetStartOk() (*string, bool)`

GetQuarantineSubnetStartOk returns a tuple with the QuarantineSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetStart

`func (o *CreateNetworkDevice) SetQuarantineSubnetStart(v string)`

SetQuarantineSubnetStart sets QuarantineSubnetStart field to given value.

### HasQuarantineSubnetStart

`func (o *CreateNetworkDevice) HasQuarantineSubnetStart() bool`

HasQuarantineSubnetStart returns a boolean if a field has been set.

### SetQuarantineSubnetStartNil

`func (o *CreateNetworkDevice) SetQuarantineSubnetStartNil(b bool)`

 SetQuarantineSubnetStartNil sets the value for QuarantineSubnetStart to be an explicit nil

### UnsetQuarantineSubnetStart
`func (o *CreateNetworkDevice) UnsetQuarantineSubnetStart()`

UnsetQuarantineSubnetStart ensures that no value is present for QuarantineSubnetStart, not even an explicit nil
### GetQuarantineSubnetEnd

`func (o *CreateNetworkDevice) GetQuarantineSubnetEnd() string`

GetQuarantineSubnetEnd returns the QuarantineSubnetEnd field if non-nil, zero value otherwise.

### GetQuarantineSubnetEndOk

`func (o *CreateNetworkDevice) GetQuarantineSubnetEndOk() (*string, bool)`

GetQuarantineSubnetEndOk returns a tuple with the QuarantineSubnetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetEnd

`func (o *CreateNetworkDevice) SetQuarantineSubnetEnd(v string)`

SetQuarantineSubnetEnd sets QuarantineSubnetEnd field to given value.

### HasQuarantineSubnetEnd

`func (o *CreateNetworkDevice) HasQuarantineSubnetEnd() bool`

HasQuarantineSubnetEnd returns a boolean if a field has been set.

### SetQuarantineSubnetEndNil

`func (o *CreateNetworkDevice) SetQuarantineSubnetEndNil(b bool)`

 SetQuarantineSubnetEndNil sets the value for QuarantineSubnetEnd to be an explicit nil

### UnsetQuarantineSubnetEnd
`func (o *CreateNetworkDevice) UnsetQuarantineSubnetEnd()`

UnsetQuarantineSubnetEnd ensures that no value is present for QuarantineSubnetEnd, not even an explicit nil
### GetQuarantineSubnetPrefixSize

`func (o *CreateNetworkDevice) GetQuarantineSubnetPrefixSize() int32`

GetQuarantineSubnetPrefixSize returns the QuarantineSubnetPrefixSize field if non-nil, zero value otherwise.

### GetQuarantineSubnetPrefixSizeOk

`func (o *CreateNetworkDevice) GetQuarantineSubnetPrefixSizeOk() (*int32, bool)`

GetQuarantineSubnetPrefixSizeOk returns a tuple with the QuarantineSubnetPrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetPrefixSize

`func (o *CreateNetworkDevice) SetQuarantineSubnetPrefixSize(v int32)`

SetQuarantineSubnetPrefixSize sets QuarantineSubnetPrefixSize field to given value.

### HasQuarantineSubnetPrefixSize

`func (o *CreateNetworkDevice) HasQuarantineSubnetPrefixSize() bool`

HasQuarantineSubnetPrefixSize returns a boolean if a field has been set.

### SetQuarantineSubnetPrefixSizeNil

`func (o *CreateNetworkDevice) SetQuarantineSubnetPrefixSizeNil(b bool)`

 SetQuarantineSubnetPrefixSizeNil sets the value for QuarantineSubnetPrefixSize to be an explicit nil

### UnsetQuarantineSubnetPrefixSize
`func (o *CreateNetworkDevice) UnsetQuarantineSubnetPrefixSize()`

UnsetQuarantineSubnetPrefixSize ensures that no value is present for QuarantineSubnetPrefixSize, not even an explicit nil
### GetQuarantineSubnetGateway

`func (o *CreateNetworkDevice) GetQuarantineSubnetGateway() string`

GetQuarantineSubnetGateway returns the QuarantineSubnetGateway field if non-nil, zero value otherwise.

### GetQuarantineSubnetGatewayOk

`func (o *CreateNetworkDevice) GetQuarantineSubnetGatewayOk() (*string, bool)`

GetQuarantineSubnetGatewayOk returns a tuple with the QuarantineSubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetGateway

`func (o *CreateNetworkDevice) SetQuarantineSubnetGateway(v string)`

SetQuarantineSubnetGateway sets QuarantineSubnetGateway field to given value.

### HasQuarantineSubnetGateway

`func (o *CreateNetworkDevice) HasQuarantineSubnetGateway() bool`

HasQuarantineSubnetGateway returns a boolean if a field has been set.

### SetQuarantineSubnetGatewayNil

`func (o *CreateNetworkDevice) SetQuarantineSubnetGatewayNil(b bool)`

 SetQuarantineSubnetGatewayNil sets the value for QuarantineSubnetGateway to be an explicit nil

### UnsetQuarantineSubnetGateway
`func (o *CreateNetworkDevice) UnsetQuarantineSubnetGateway()`

UnsetQuarantineSubnetGateway ensures that no value is present for QuarantineSubnetGateway, not even an explicit nil
### GetQuarantineVlan

`func (o *CreateNetworkDevice) GetQuarantineVlan() int32`

GetQuarantineVlan returns the QuarantineVlan field if non-nil, zero value otherwise.

### GetQuarantineVlanOk

`func (o *CreateNetworkDevice) GetQuarantineVlanOk() (*int32, bool)`

GetQuarantineVlanOk returns a tuple with the QuarantineVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineVlan

`func (o *CreateNetworkDevice) SetQuarantineVlan(v int32)`

SetQuarantineVlan sets QuarantineVlan field to given value.

### HasQuarantineVlan

`func (o *CreateNetworkDevice) HasQuarantineVlan() bool`

HasQuarantineVlan returns a boolean if a field has been set.

### GetManagementProtocol

`func (o *CreateNetworkDevice) GetManagementProtocol() string`

GetManagementProtocol returns the ManagementProtocol field if non-nil, zero value otherwise.

### GetManagementProtocolOk

`func (o *CreateNetworkDevice) GetManagementProtocolOk() (*string, bool)`

GetManagementProtocolOk returns a tuple with the ManagementProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementProtocol

`func (o *CreateNetworkDevice) SetManagementProtocol(v string)`

SetManagementProtocol sets ManagementProtocol field to given value.

### HasManagementProtocol

`func (o *CreateNetworkDevice) HasManagementProtocol() bool`

HasManagementProtocol returns a boolean if a field has been set.

### SetManagementProtocolNil

`func (o *CreateNetworkDevice) SetManagementProtocolNil(b bool)`

 SetManagementProtocolNil sets the value for ManagementProtocol to be an explicit nil

### UnsetManagementProtocol
`func (o *CreateNetworkDevice) UnsetManagementProtocol()`

UnsetManagementProtocol ensures that no value is present for ManagementProtocol, not even an explicit nil
### GetManagementAddress

`func (o *CreateNetworkDevice) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *CreateNetworkDevice) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *CreateNetworkDevice) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *CreateNetworkDevice) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### SetManagementAddressNil

`func (o *CreateNetworkDevice) SetManagementAddressNil(b bool)`

 SetManagementAddressNil sets the value for ManagementAddress to be an explicit nil

### UnsetManagementAddress
`func (o *CreateNetworkDevice) UnsetManagementAddress()`

UnsetManagementAddress ensures that no value is present for ManagementAddress, not even an explicit nil
### GetManagementPort

`func (o *CreateNetworkDevice) GetManagementPort() int32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *CreateNetworkDevice) GetManagementPortOk() (*int32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *CreateNetworkDevice) SetManagementPort(v int32)`

SetManagementPort sets ManagementPort field to given value.

### HasManagementPort

`func (o *CreateNetworkDevice) HasManagementPort() bool`

HasManagementPort returns a boolean if a field has been set.

### SetManagementPortNil

`func (o *CreateNetworkDevice) SetManagementPortNil(b bool)`

 SetManagementPortNil sets the value for ManagementPort to be an explicit nil

### UnsetManagementPort
`func (o *CreateNetworkDevice) UnsetManagementPort()`

UnsetManagementPort ensures that no value is present for ManagementPort, not even an explicit nil
### GetUsername

`func (o *CreateNetworkDevice) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateNetworkDevice) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateNetworkDevice) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateNetworkDevice) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateNetworkDevice) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateNetworkDevice) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetManagementPassword

`func (o *CreateNetworkDevice) GetManagementPassword() string`

GetManagementPassword returns the ManagementPassword field if non-nil, zero value otherwise.

### GetManagementPasswordOk

`func (o *CreateNetworkDevice) GetManagementPasswordOk() (*string, bool)`

GetManagementPasswordOk returns a tuple with the ManagementPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPassword

`func (o *CreateNetworkDevice) SetManagementPassword(v string)`

SetManagementPassword sets ManagementPassword field to given value.

### HasManagementPassword

`func (o *CreateNetworkDevice) HasManagementPassword() bool`

HasManagementPassword returns a boolean if a field has been set.

### SetManagementPasswordNil

`func (o *CreateNetworkDevice) SetManagementPasswordNil(b bool)`

 SetManagementPasswordNil sets the value for ManagementPassword to be an explicit nil

### UnsetManagementPassword
`func (o *CreateNetworkDevice) UnsetManagementPassword()`

UnsetManagementPassword ensures that no value is present for ManagementPassword, not even an explicit nil
### GetManagementAddressGateway

`func (o *CreateNetworkDevice) GetManagementAddressGateway() string`

GetManagementAddressGateway returns the ManagementAddressGateway field if non-nil, zero value otherwise.

### GetManagementAddressGatewayOk

`func (o *CreateNetworkDevice) GetManagementAddressGatewayOk() (*string, bool)`

GetManagementAddressGatewayOk returns a tuple with the ManagementAddressGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddressGateway

`func (o *CreateNetworkDevice) SetManagementAddressGateway(v string)`

SetManagementAddressGateway sets ManagementAddressGateway field to given value.

### HasManagementAddressGateway

`func (o *CreateNetworkDevice) HasManagementAddressGateway() bool`

HasManagementAddressGateway returns a boolean if a field has been set.

### SetManagementAddressGatewayNil

`func (o *CreateNetworkDevice) SetManagementAddressGatewayNil(b bool)`

 SetManagementAddressGatewayNil sets the value for ManagementAddressGateway to be an explicit nil

### UnsetManagementAddressGateway
`func (o *CreateNetworkDevice) UnsetManagementAddressGateway()`

UnsetManagementAddressGateway ensures that no value is present for ManagementAddressGateway, not even an explicit nil
### GetManagementAddressMask

`func (o *CreateNetworkDevice) GetManagementAddressMask() string`

GetManagementAddressMask returns the ManagementAddressMask field if non-nil, zero value otherwise.

### GetManagementAddressMaskOk

`func (o *CreateNetworkDevice) GetManagementAddressMaskOk() (*string, bool)`

GetManagementAddressMaskOk returns a tuple with the ManagementAddressMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddressMask

`func (o *CreateNetworkDevice) SetManagementAddressMask(v string)`

SetManagementAddressMask sets ManagementAddressMask field to given value.

### HasManagementAddressMask

`func (o *CreateNetworkDevice) HasManagementAddressMask() bool`

HasManagementAddressMask returns a boolean if a field has been set.

### SetManagementAddressMaskNil

`func (o *CreateNetworkDevice) SetManagementAddressMaskNil(b bool)`

 SetManagementAddressMaskNil sets the value for ManagementAddressMask to be an explicit nil

### UnsetManagementAddressMask
`func (o *CreateNetworkDevice) UnsetManagementAddressMask()`

UnsetManagementAddressMask ensures that no value is present for ManagementAddressMask, not even an explicit nil
### GetManagementMAC

`func (o *CreateNetworkDevice) GetManagementMAC() string`

GetManagementMAC returns the ManagementMAC field if non-nil, zero value otherwise.

### GetManagementMACOk

`func (o *CreateNetworkDevice) GetManagementMACOk() (*string, bool)`

GetManagementMACOk returns a tuple with the ManagementMAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMAC

`func (o *CreateNetworkDevice) SetManagementMAC(v string)`

SetManagementMAC sets ManagementMAC field to given value.

### HasManagementMAC

`func (o *CreateNetworkDevice) HasManagementMAC() bool`

HasManagementMAC returns a boolean if a field has been set.

### SetManagementMACNil

`func (o *CreateNetworkDevice) SetManagementMACNil(b bool)`

 SetManagementMACNil sets the value for ManagementMAC to be an explicit nil

### UnsetManagementMAC
`func (o *CreateNetworkDevice) UnsetManagementMAC()`

UnsetManagementMAC ensures that no value is present for ManagementMAC, not even an explicit nil
### GetLoopbackAddress

`func (o *CreateNetworkDevice) GetLoopbackAddress() string`

GetLoopbackAddress returns the LoopbackAddress field if non-nil, zero value otherwise.

### GetLoopbackAddressOk

`func (o *CreateNetworkDevice) GetLoopbackAddressOk() (*string, bool)`

GetLoopbackAddressOk returns a tuple with the LoopbackAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackAddress

`func (o *CreateNetworkDevice) SetLoopbackAddress(v string)`

SetLoopbackAddress sets LoopbackAddress field to given value.

### HasLoopbackAddress

`func (o *CreateNetworkDevice) HasLoopbackAddress() bool`

HasLoopbackAddress returns a boolean if a field has been set.

### SetLoopbackAddressNil

`func (o *CreateNetworkDevice) SetLoopbackAddressNil(b bool)`

 SetLoopbackAddressNil sets the value for LoopbackAddress to be an explicit nil

### UnsetLoopbackAddress
`func (o *CreateNetworkDevice) UnsetLoopbackAddress()`

UnsetLoopbackAddress ensures that no value is present for LoopbackAddress, not even an explicit nil
### GetVtepAddress

`func (o *CreateNetworkDevice) GetVtepAddress() string`

GetVtepAddress returns the VtepAddress field if non-nil, zero value otherwise.

### GetVtepAddressOk

`func (o *CreateNetworkDevice) GetVtepAddressOk() (*string, bool)`

GetVtepAddressOk returns a tuple with the VtepAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtepAddress

`func (o *CreateNetworkDevice) SetVtepAddress(v string)`

SetVtepAddress sets VtepAddress field to given value.

### HasVtepAddress

`func (o *CreateNetworkDevice) HasVtepAddress() bool`

HasVtepAddress returns a boolean if a field has been set.

### SetVtepAddressNil

`func (o *CreateNetworkDevice) SetVtepAddressNil(b bool)`

 SetVtepAddressNil sets the value for VtepAddress to be an explicit nil

### UnsetVtepAddress
`func (o *CreateNetworkDevice) UnsetVtepAddress()`

UnsetVtepAddress ensures that no value is present for VtepAddress, not even an explicit nil
### GetAsn

`func (o *CreateNetworkDevice) GetAsn() float32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *CreateNetworkDevice) GetAsnOk() (*float32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *CreateNetworkDevice) SetAsn(v float32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *CreateNetworkDevice) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *CreateNetworkDevice) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *CreateNetworkDevice) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetNetworkTypesAllowed

`func (o *CreateNetworkDevice) GetNetworkTypesAllowed() []string`

GetNetworkTypesAllowed returns the NetworkTypesAllowed field if non-nil, zero value otherwise.

### GetNetworkTypesAllowedOk

`func (o *CreateNetworkDevice) GetNetworkTypesAllowedOk() (*[]string, bool)`

GetNetworkTypesAllowedOk returns a tuple with the NetworkTypesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypesAllowed

`func (o *CreateNetworkDevice) SetNetworkTypesAllowed(v []string)`

SetNetworkTypesAllowed sets NetworkTypesAllowed field to given value.

### HasNetworkTypesAllowed

`func (o *CreateNetworkDevice) HasNetworkTypesAllowed() bool`

HasNetworkTypesAllowed returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateNetworkDevice) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateNetworkDevice) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCountry

`func (o *CreateNetworkDevice) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateNetworkDevice) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateNetworkDevice) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateNetworkDevice) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *CreateNetworkDevice) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *CreateNetworkDevice) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCity

`func (o *CreateNetworkDevice) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateNetworkDevice) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreateNetworkDevice) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CreateNetworkDevice) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *CreateNetworkDevice) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *CreateNetworkDevice) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDatacenterMeta

`func (o *CreateNetworkDevice) GetDatacenterMeta() string`

GetDatacenterMeta returns the DatacenterMeta field if non-nil, zero value otherwise.

### GetDatacenterMetaOk

`func (o *CreateNetworkDevice) GetDatacenterMetaOk() (*string, bool)`

GetDatacenterMetaOk returns a tuple with the DatacenterMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterMeta

`func (o *CreateNetworkDevice) SetDatacenterMeta(v string)`

SetDatacenterMeta sets DatacenterMeta field to given value.

### HasDatacenterMeta

`func (o *CreateNetworkDevice) HasDatacenterMeta() bool`

HasDatacenterMeta returns a boolean if a field has been set.

### SetDatacenterMetaNil

`func (o *CreateNetworkDevice) SetDatacenterMetaNil(b bool)`

 SetDatacenterMetaNil sets the value for DatacenterMeta to be an explicit nil

### UnsetDatacenterMeta
`func (o *CreateNetworkDevice) UnsetDatacenterMeta()`

UnsetDatacenterMeta ensures that no value is present for DatacenterMeta, not even an explicit nil
### GetDatacenterRoom

`func (o *CreateNetworkDevice) GetDatacenterRoom() string`

GetDatacenterRoom returns the DatacenterRoom field if non-nil, zero value otherwise.

### GetDatacenterRoomOk

`func (o *CreateNetworkDevice) GetDatacenterRoomOk() (*string, bool)`

GetDatacenterRoomOk returns a tuple with the DatacenterRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterRoom

`func (o *CreateNetworkDevice) SetDatacenterRoom(v string)`

SetDatacenterRoom sets DatacenterRoom field to given value.

### HasDatacenterRoom

`func (o *CreateNetworkDevice) HasDatacenterRoom() bool`

HasDatacenterRoom returns a boolean if a field has been set.

### SetDatacenterRoomNil

`func (o *CreateNetworkDevice) SetDatacenterRoomNil(b bool)`

 SetDatacenterRoomNil sets the value for DatacenterRoom to be an explicit nil

### UnsetDatacenterRoom
`func (o *CreateNetworkDevice) UnsetDatacenterRoom()`

UnsetDatacenterRoom ensures that no value is present for DatacenterRoom, not even an explicit nil
### GetDatacenterRack

`func (o *CreateNetworkDevice) GetDatacenterRack() string`

GetDatacenterRack returns the DatacenterRack field if non-nil, zero value otherwise.

### GetDatacenterRackOk

`func (o *CreateNetworkDevice) GetDatacenterRackOk() (*string, bool)`

GetDatacenterRackOk returns a tuple with the DatacenterRack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterRack

`func (o *CreateNetworkDevice) SetDatacenterRack(v string)`

SetDatacenterRack sets DatacenterRack field to given value.

### HasDatacenterRack

`func (o *CreateNetworkDevice) HasDatacenterRack() bool`

HasDatacenterRack returns a boolean if a field has been set.

### SetDatacenterRackNil

`func (o *CreateNetworkDevice) SetDatacenterRackNil(b bool)`

 SetDatacenterRackNil sets the value for DatacenterRack to be an explicit nil

### UnsetDatacenterRack
`func (o *CreateNetworkDevice) UnsetDatacenterRack()`

UnsetDatacenterRack ensures that no value is present for DatacenterRack, not even an explicit nil
### GetRackPositionUpperUnit

`func (o *CreateNetworkDevice) GetRackPositionUpperUnit() int32`

GetRackPositionUpperUnit returns the RackPositionUpperUnit field if non-nil, zero value otherwise.

### GetRackPositionUpperUnitOk

`func (o *CreateNetworkDevice) GetRackPositionUpperUnitOk() (*int32, bool)`

GetRackPositionUpperUnitOk returns a tuple with the RackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionUpperUnit

`func (o *CreateNetworkDevice) SetRackPositionUpperUnit(v int32)`

SetRackPositionUpperUnit sets RackPositionUpperUnit field to given value.

### HasRackPositionUpperUnit

`func (o *CreateNetworkDevice) HasRackPositionUpperUnit() bool`

HasRackPositionUpperUnit returns a boolean if a field has been set.

### SetRackPositionUpperUnitNil

`func (o *CreateNetworkDevice) SetRackPositionUpperUnitNil(b bool)`

 SetRackPositionUpperUnitNil sets the value for RackPositionUpperUnit to be an explicit nil

### UnsetRackPositionUpperUnit
`func (o *CreateNetworkDevice) UnsetRackPositionUpperUnit()`

UnsetRackPositionUpperUnit ensures that no value is present for RackPositionUpperUnit, not even an explicit nil
### GetRackPositionLowerUnit

`func (o *CreateNetworkDevice) GetRackPositionLowerUnit() int32`

GetRackPositionLowerUnit returns the RackPositionLowerUnit field if non-nil, zero value otherwise.

### GetRackPositionLowerUnitOk

`func (o *CreateNetworkDevice) GetRackPositionLowerUnitOk() (*int32, bool)`

GetRackPositionLowerUnitOk returns a tuple with the RackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionLowerUnit

`func (o *CreateNetworkDevice) SetRackPositionLowerUnit(v int32)`

SetRackPositionLowerUnit sets RackPositionLowerUnit field to given value.

### HasRackPositionLowerUnit

`func (o *CreateNetworkDevice) HasRackPositionLowerUnit() bool`

HasRackPositionLowerUnit returns a boolean if a field has been set.

### SetRackPositionLowerUnitNil

`func (o *CreateNetworkDevice) SetRackPositionLowerUnitNil(b bool)`

 SetRackPositionLowerUnitNil sets the value for RackPositionLowerUnit to be an explicit nil

### UnsetRackPositionLowerUnit
`func (o *CreateNetworkDevice) UnsetRackPositionLowerUnit()`

UnsetRackPositionLowerUnit ensures that no value is present for RackPositionLowerUnit, not even an explicit nil
### GetSerialNumber

`func (o *CreateNetworkDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CreateNetworkDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CreateNetworkDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CreateNetworkDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetOrderIndex

`func (o *CreateNetworkDevice) GetOrderIndex() float32`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *CreateNetworkDevice) GetOrderIndexOk() (*float32, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *CreateNetworkDevice) SetOrderIndex(v float32)`

SetOrderIndex sets OrderIndex field to given value.

### HasOrderIndex

`func (o *CreateNetworkDevice) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.

### GetDefaultMtu

`func (o *CreateNetworkDevice) GetDefaultMtu() float32`

GetDefaultMtu returns the DefaultMtu field if non-nil, zero value otherwise.

### GetDefaultMtuOk

`func (o *CreateNetworkDevice) GetDefaultMtuOk() (*float32, bool)`

GetDefaultMtuOk returns a tuple with the DefaultMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtu

`func (o *CreateNetworkDevice) SetDefaultMtu(v float32)`

SetDefaultMtu sets DefaultMtu field to given value.

### HasDefaultMtu

`func (o *CreateNetworkDevice) HasDefaultMtu() bool`

HasDefaultMtu returns a boolean if a field has been set.

### SetDefaultMtuNil

`func (o *CreateNetworkDevice) SetDefaultMtuNil(b bool)`

 SetDefaultMtuNil sets the value for DefaultMtu to be an explicit nil

### UnsetDefaultMtu
`func (o *CreateNetworkDevice) UnsetDefaultMtu()`

UnsetDefaultMtu ensures that no value is present for DefaultMtu, not even an explicit nil
### GetTags

`func (o *CreateNetworkDevice) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateNetworkDevice) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateNetworkDevice) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateNetworkDevice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateNetworkDevice) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateNetworkDevice) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetRequiresOsInstall

`func (o *CreateNetworkDevice) GetRequiresOsInstall() bool`

GetRequiresOsInstall returns the RequiresOsInstall field if non-nil, zero value otherwise.

### GetRequiresOsInstallOk

`func (o *CreateNetworkDevice) GetRequiresOsInstallOk() (*bool, bool)`

GetRequiresOsInstallOk returns a tuple with the RequiresOsInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOsInstall

`func (o *CreateNetworkDevice) SetRequiresOsInstall(v bool)`

SetRequiresOsInstall sets RequiresOsInstall field to given value.

### HasRequiresOsInstall

`func (o *CreateNetworkDevice) HasRequiresOsInstall() bool`

HasRequiresOsInstall returns a boolean if a field has been set.

### GetOobMgmtIpv4Address

`func (o *CreateNetworkDevice) GetOobMgmtIpv4Address() string`

GetOobMgmtIpv4Address returns the OobMgmtIpv4Address field if non-nil, zero value otherwise.

### GetOobMgmtIpv4AddressOk

`func (o *CreateNetworkDevice) GetOobMgmtIpv4AddressOk() (*string, bool)`

GetOobMgmtIpv4AddressOk returns a tuple with the OobMgmtIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobMgmtIpv4Address

`func (o *CreateNetworkDevice) SetOobMgmtIpv4Address(v string)`

SetOobMgmtIpv4Address sets OobMgmtIpv4Address field to given value.

### HasOobMgmtIpv4Address

`func (o *CreateNetworkDevice) HasOobMgmtIpv4Address() bool`

HasOobMgmtIpv4Address returns a boolean if a field has been set.

### SetOobMgmtIpv4AddressNil

`func (o *CreateNetworkDevice) SetOobMgmtIpv4AddressNil(b bool)`

 SetOobMgmtIpv4AddressNil sets the value for OobMgmtIpv4Address to be an explicit nil

### UnsetOobMgmtIpv4Address
`func (o *CreateNetworkDevice) UnsetOobMgmtIpv4Address()`

UnsetOobMgmtIpv4Address ensures that no value is present for OobMgmtIpv4Address, not even an explicit nil
### GetOverwriteWithHostnameFromFetchedSwitch

`func (o *CreateNetworkDevice) GetOverwriteWithHostnameFromFetchedSwitch() bool`

GetOverwriteWithHostnameFromFetchedSwitch returns the OverwriteWithHostnameFromFetchedSwitch field if non-nil, zero value otherwise.

### GetOverwriteWithHostnameFromFetchedSwitchOk

`func (o *CreateNetworkDevice) GetOverwriteWithHostnameFromFetchedSwitchOk() (*bool, bool)`

GetOverwriteWithHostnameFromFetchedSwitchOk returns a tuple with the OverwriteWithHostnameFromFetchedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteWithHostnameFromFetchedSwitch

`func (o *CreateNetworkDevice) SetOverwriteWithHostnameFromFetchedSwitch(v bool)`

SetOverwriteWithHostnameFromFetchedSwitch sets OverwriteWithHostnameFromFetchedSwitch field to given value.

### HasOverwriteWithHostnameFromFetchedSwitch

`func (o *CreateNetworkDevice) HasOverwriteWithHostnameFromFetchedSwitch() bool`

HasOverwriteWithHostnameFromFetchedSwitch returns a boolean if a field has been set.

### SetOverwriteWithHostnameFromFetchedSwitchNil

`func (o *CreateNetworkDevice) SetOverwriteWithHostnameFromFetchedSwitchNil(b bool)`

 SetOverwriteWithHostnameFromFetchedSwitchNil sets the value for OverwriteWithHostnameFromFetchedSwitch to be an explicit nil

### UnsetOverwriteWithHostnameFromFetchedSwitch
`func (o *CreateNetworkDevice) UnsetOverwriteWithHostnameFromFetchedSwitch()`

UnsetOverwriteWithHostnameFromFetchedSwitch ensures that no value is present for OverwriteWithHostnameFromFetchedSwitch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


