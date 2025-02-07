# CreateNetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterName** | **string** | Name of the datacenter | 
**IdentifierString** | Pointer to **string** | Identifier string | [optional] 
**ChassisRackId** | Pointer to **float32** | Chassis rack ID | [optional] 
**ChassisIdentifier** | Pointer to **string** | Chassis identifier | [optional] 
**ProvisionerType** | **string** | Provisioner type | 
**Driver** | **string** | Driver | 
**Position** | **string** | Position | 
**TorLinkedId** | Pointer to **float32** | TOR Linked ID | [optional] 
**IsGateway** | Pointer to **bool** | Is Gateway | [optional] 
**SyslogEnabled** | Pointer to **bool** | Is Syslog Enabled | [optional] 
**IsStorageSwitch** | Pointer to **bool** | Is storage network device | [optional] 
**IsBorderDevice** | Pointer to **bool** | Is border network device | [optional] 
**PrimaryWanIpv4SubnetPool** | Pointer to **string** | Primary WAN IPV4 Subnet Pool | [optional] 
**PrimaryWanIpv4SubnetPrefixSize** | Pointer to **float32** | Primary WAN IPV4 Subnet Prefix Size | [optional] 
**PrimaryWanIpv6SubnetPool** | Pointer to **string** | Primary WAN IPV6 Subnet Pool | [optional] 
**PrimaryWanIpv6SubnetPrefixSize** | Pointer to **float32** | Primary WAN IPV6 Subnet Prefix Size | [optional] 
**PrimarySanSubnetPool** | Pointer to **string** | Primary SAN Subnet Pool | [optional] 
**PrimarySanSubnetPrefixSize** | Pointer to **float32** | Primary SAN Subnet Prefix Size | [optional] 
**QuarantineSubnetStart** | Pointer to **string** | Quarantine Subnet Start | [optional] 
**QuarantineSubnetEnd** | Pointer to **string** | Quarantine Subnet End | [optional] 
**QuarantineSubnetPrefixSize** | Pointer to **float32** | Quarantine Subnet Prefix Size | [optional] 
**QuarantineSubnetGateway** | Pointer to **string** | Quarantine Subnet Gateway | [optional] 
**QuarantineVlan** | Pointer to **float32** | Quarantine VLAN | [optional] 
**ManagementProtocol** | Pointer to **string** | Management Protocol | [optional] 
**ManagementAddress** | Pointer to **string** | Management Address | [optional] 
**ManagementPort** | Pointer to **float32** | Management Port | [optional] 
**Username** | Pointer to **string** | Management Username | [optional] 
**ManagementPassword** | Pointer to **string** | Management Password | [optional] 
**ManagementAddressGateway** | Pointer to **string** | Management Address Gateway | [optional] 
**ManagementAddressMask** | Pointer to **string** | Management Address Mask | [optional] 
**ManagementMAC** | Pointer to **string** | Management MAC Address | [optional] 
**LoopbackAddress** | Pointer to **string** | Loopback Address | [optional] 
**VtepAddress** | Pointer to **string** | Loopback Address IPv6 | [optional] 
**Asn** | Pointer to **float32** | ASN | [optional] 
**NetworkTypesAllowed** | Pointer to **[]string** | Network types allowed | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Country** | Pointer to **string** | Country | [optional] 
**City** | Pointer to **string** | City | [optional] 
**DatacenterMeta** | Pointer to **string** | Datacenter metadata | [optional] 
**DatacenterRoom** | Pointer to **string** | Datacenter room | [optional] 
**DatacenterRack** | Pointer to **string** | Datacenter rack | [optional] 
**RackPositionUpperUnit** | Pointer to **float32** | Upper rack position | [optional] 
**RackPositionLowerUnit** | Pointer to **float32** | Lower rack position | [optional] 
**SerialNumber** | Pointer to **string** | Serial number | [optional] 
**OrderIndex** | Pointer to **float32** | Order Index | [optional] 
**DefaultMtu** | Pointer to **float32** | Default MTU | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**RequiresOsInstall** | Pointer to **bool** | Requires OS install | [optional] 
**SiteId** | Pointer to **float32** | Site Id | [optional] 
**OobMgmtIpv4Address** | Pointer to **map[string]interface{}** | OOB Management IPV4 Address | [optional] 
**PrimaryWanIpv6SubnetPoolId** | Pointer to **float32** | Primary WAN IPV6 Subnet Pool Id | [optional] 
**OverwriteWithHostnameFromFetchedSwitch** | Pointer to **bool** | Overwrite With Hostname From Fetched Network Device | [optional] 

## Methods

### NewCreateNetworkDevice

`func NewCreateNetworkDevice(datacenterName string, provisionerType string, driver string, position string, ) *CreateNetworkDevice`

NewCreateNetworkDevice instantiates a new CreateNetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceWithDefaults

`func NewCreateNetworkDeviceWithDefaults() *CreateNetworkDevice`

NewCreateNetworkDeviceWithDefaults instantiates a new CreateNetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *CreateNetworkDevice) GetChassisRackId() float32`

GetChassisRackId returns the ChassisRackId field if non-nil, zero value otherwise.

### GetChassisRackIdOk

`func (o *CreateNetworkDevice) GetChassisRackIdOk() (*float32, bool)`

GetChassisRackIdOk returns a tuple with the ChassisRackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisRackId

`func (o *CreateNetworkDevice) SetChassisRackId(v float32)`

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

### GetProvisionerType

`func (o *CreateNetworkDevice) GetProvisionerType() string`

GetProvisionerType returns the ProvisionerType field if non-nil, zero value otherwise.

### GetProvisionerTypeOk

`func (o *CreateNetworkDevice) GetProvisionerTypeOk() (*string, bool)`

GetProvisionerTypeOk returns a tuple with the ProvisionerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerType

`func (o *CreateNetworkDevice) SetProvisionerType(v string)`

SetProvisionerType sets ProvisionerType field to given value.


### GetDriver

`func (o *CreateNetworkDevice) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *CreateNetworkDevice) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *CreateNetworkDevice) SetDriver(v string)`

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

### GetPrimaryWanIpv4SubnetPool

`func (o *CreateNetworkDevice) GetPrimaryWanIpv4SubnetPool() string`

GetPrimaryWanIpv4SubnetPool returns the PrimaryWanIpv4SubnetPool field if non-nil, zero value otherwise.

### GetPrimaryWanIpv4SubnetPoolOk

`func (o *CreateNetworkDevice) GetPrimaryWanIpv4SubnetPoolOk() (*string, bool)`

GetPrimaryWanIpv4SubnetPoolOk returns a tuple with the PrimaryWanIpv4SubnetPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWanIpv4SubnetPool

`func (o *CreateNetworkDevice) SetPrimaryWanIpv4SubnetPool(v string)`

SetPrimaryWanIpv4SubnetPool sets PrimaryWanIpv4SubnetPool field to given value.

### HasPrimaryWanIpv4SubnetPool

`func (o *CreateNetworkDevice) HasPrimaryWanIpv4SubnetPool() bool`

HasPrimaryWanIpv4SubnetPool returns a boolean if a field has been set.

### GetPrimaryWanIpv4SubnetPrefixSize

`func (o *CreateNetworkDevice) GetPrimaryWanIpv4SubnetPrefixSize() float32`

GetPrimaryWanIpv4SubnetPrefixSize returns the PrimaryWanIpv4SubnetPrefixSize field if non-nil, zero value otherwise.

### GetPrimaryWanIpv4SubnetPrefixSizeOk

`func (o *CreateNetworkDevice) GetPrimaryWanIpv4SubnetPrefixSizeOk() (*float32, bool)`

GetPrimaryWanIpv4SubnetPrefixSizeOk returns a tuple with the PrimaryWanIpv4SubnetPrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWanIpv4SubnetPrefixSize

`func (o *CreateNetworkDevice) SetPrimaryWanIpv4SubnetPrefixSize(v float32)`

SetPrimaryWanIpv4SubnetPrefixSize sets PrimaryWanIpv4SubnetPrefixSize field to given value.

### HasPrimaryWanIpv4SubnetPrefixSize

`func (o *CreateNetworkDevice) HasPrimaryWanIpv4SubnetPrefixSize() bool`

HasPrimaryWanIpv4SubnetPrefixSize returns a boolean if a field has been set.

### GetPrimaryWanIpv6SubnetPool

`func (o *CreateNetworkDevice) GetPrimaryWanIpv6SubnetPool() string`

GetPrimaryWanIpv6SubnetPool returns the PrimaryWanIpv6SubnetPool field if non-nil, zero value otherwise.

### GetPrimaryWanIpv6SubnetPoolOk

`func (o *CreateNetworkDevice) GetPrimaryWanIpv6SubnetPoolOk() (*string, bool)`

GetPrimaryWanIpv6SubnetPoolOk returns a tuple with the PrimaryWanIpv6SubnetPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWanIpv6SubnetPool

`func (o *CreateNetworkDevice) SetPrimaryWanIpv6SubnetPool(v string)`

SetPrimaryWanIpv6SubnetPool sets PrimaryWanIpv6SubnetPool field to given value.

### HasPrimaryWanIpv6SubnetPool

`func (o *CreateNetworkDevice) HasPrimaryWanIpv6SubnetPool() bool`

HasPrimaryWanIpv6SubnetPool returns a boolean if a field has been set.

### GetPrimaryWanIpv6SubnetPrefixSize

`func (o *CreateNetworkDevice) GetPrimaryWanIpv6SubnetPrefixSize() float32`

GetPrimaryWanIpv6SubnetPrefixSize returns the PrimaryWanIpv6SubnetPrefixSize field if non-nil, zero value otherwise.

### GetPrimaryWanIpv6SubnetPrefixSizeOk

`func (o *CreateNetworkDevice) GetPrimaryWanIpv6SubnetPrefixSizeOk() (*float32, bool)`

GetPrimaryWanIpv6SubnetPrefixSizeOk returns a tuple with the PrimaryWanIpv6SubnetPrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWanIpv6SubnetPrefixSize

`func (o *CreateNetworkDevice) SetPrimaryWanIpv6SubnetPrefixSize(v float32)`

SetPrimaryWanIpv6SubnetPrefixSize sets PrimaryWanIpv6SubnetPrefixSize field to given value.

### HasPrimaryWanIpv6SubnetPrefixSize

`func (o *CreateNetworkDevice) HasPrimaryWanIpv6SubnetPrefixSize() bool`

HasPrimaryWanIpv6SubnetPrefixSize returns a boolean if a field has been set.

### GetPrimarySanSubnetPool

`func (o *CreateNetworkDevice) GetPrimarySanSubnetPool() string`

GetPrimarySanSubnetPool returns the PrimarySanSubnetPool field if non-nil, zero value otherwise.

### GetPrimarySanSubnetPoolOk

`func (o *CreateNetworkDevice) GetPrimarySanSubnetPoolOk() (*string, bool)`

GetPrimarySanSubnetPoolOk returns a tuple with the PrimarySanSubnetPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySanSubnetPool

`func (o *CreateNetworkDevice) SetPrimarySanSubnetPool(v string)`

SetPrimarySanSubnetPool sets PrimarySanSubnetPool field to given value.

### HasPrimarySanSubnetPool

`func (o *CreateNetworkDevice) HasPrimarySanSubnetPool() bool`

HasPrimarySanSubnetPool returns a boolean if a field has been set.

### GetPrimarySanSubnetPrefixSize

`func (o *CreateNetworkDevice) GetPrimarySanSubnetPrefixSize() float32`

GetPrimarySanSubnetPrefixSize returns the PrimarySanSubnetPrefixSize field if non-nil, zero value otherwise.

### GetPrimarySanSubnetPrefixSizeOk

`func (o *CreateNetworkDevice) GetPrimarySanSubnetPrefixSizeOk() (*float32, bool)`

GetPrimarySanSubnetPrefixSizeOk returns a tuple with the PrimarySanSubnetPrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySanSubnetPrefixSize

`func (o *CreateNetworkDevice) SetPrimarySanSubnetPrefixSize(v float32)`

SetPrimarySanSubnetPrefixSize sets PrimarySanSubnetPrefixSize field to given value.

### HasPrimarySanSubnetPrefixSize

`func (o *CreateNetworkDevice) HasPrimarySanSubnetPrefixSize() bool`

HasPrimarySanSubnetPrefixSize returns a boolean if a field has been set.

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

### GetQuarantineSubnetPrefixSize

`func (o *CreateNetworkDevice) GetQuarantineSubnetPrefixSize() float32`

GetQuarantineSubnetPrefixSize returns the QuarantineSubnetPrefixSize field if non-nil, zero value otherwise.

### GetQuarantineSubnetPrefixSizeOk

`func (o *CreateNetworkDevice) GetQuarantineSubnetPrefixSizeOk() (*float32, bool)`

GetQuarantineSubnetPrefixSizeOk returns a tuple with the QuarantineSubnetPrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineSubnetPrefixSize

`func (o *CreateNetworkDevice) SetQuarantineSubnetPrefixSize(v float32)`

SetQuarantineSubnetPrefixSize sets QuarantineSubnetPrefixSize field to given value.

### HasQuarantineSubnetPrefixSize

`func (o *CreateNetworkDevice) HasQuarantineSubnetPrefixSize() bool`

HasQuarantineSubnetPrefixSize returns a boolean if a field has been set.

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

### GetQuarantineVlan

`func (o *CreateNetworkDevice) GetQuarantineVlan() float32`

GetQuarantineVlan returns the QuarantineVlan field if non-nil, zero value otherwise.

### GetQuarantineVlanOk

`func (o *CreateNetworkDevice) GetQuarantineVlanOk() (*float32, bool)`

GetQuarantineVlanOk returns a tuple with the QuarantineVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineVlan

`func (o *CreateNetworkDevice) SetQuarantineVlan(v float32)`

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

### GetManagementPort

`func (o *CreateNetworkDevice) GetManagementPort() float32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *CreateNetworkDevice) GetManagementPortOk() (*float32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *CreateNetworkDevice) SetManagementPort(v float32)`

SetManagementPort sets ManagementPort field to given value.

### HasManagementPort

`func (o *CreateNetworkDevice) HasManagementPort() bool`

HasManagementPort returns a boolean if a field has been set.

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

### GetRackPositionUpperUnit

`func (o *CreateNetworkDevice) GetRackPositionUpperUnit() float32`

GetRackPositionUpperUnit returns the RackPositionUpperUnit field if non-nil, zero value otherwise.

### GetRackPositionUpperUnitOk

`func (o *CreateNetworkDevice) GetRackPositionUpperUnitOk() (*float32, bool)`

GetRackPositionUpperUnitOk returns a tuple with the RackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionUpperUnit

`func (o *CreateNetworkDevice) SetRackPositionUpperUnit(v float32)`

SetRackPositionUpperUnit sets RackPositionUpperUnit field to given value.

### HasRackPositionUpperUnit

`func (o *CreateNetworkDevice) HasRackPositionUpperUnit() bool`

HasRackPositionUpperUnit returns a boolean if a field has been set.

### GetRackPositionLowerUnit

`func (o *CreateNetworkDevice) GetRackPositionLowerUnit() float32`

GetRackPositionLowerUnit returns the RackPositionLowerUnit field if non-nil, zero value otherwise.

### GetRackPositionLowerUnitOk

`func (o *CreateNetworkDevice) GetRackPositionLowerUnitOk() (*float32, bool)`

GetRackPositionLowerUnitOk returns a tuple with the RackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionLowerUnit

`func (o *CreateNetworkDevice) SetRackPositionLowerUnit(v float32)`

SetRackPositionLowerUnit sets RackPositionLowerUnit field to given value.

### HasRackPositionLowerUnit

`func (o *CreateNetworkDevice) HasRackPositionLowerUnit() bool`

HasRackPositionLowerUnit returns a boolean if a field has been set.

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

### GetSiteId

`func (o *CreateNetworkDevice) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateNetworkDevice) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateNetworkDevice) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CreateNetworkDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetOobMgmtIpv4Address

`func (o *CreateNetworkDevice) GetOobMgmtIpv4Address() map[string]interface{}`

GetOobMgmtIpv4Address returns the OobMgmtIpv4Address field if non-nil, zero value otherwise.

### GetOobMgmtIpv4AddressOk

`func (o *CreateNetworkDevice) GetOobMgmtIpv4AddressOk() (*map[string]interface{}, bool)`

GetOobMgmtIpv4AddressOk returns a tuple with the OobMgmtIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobMgmtIpv4Address

`func (o *CreateNetworkDevice) SetOobMgmtIpv4Address(v map[string]interface{})`

SetOobMgmtIpv4Address sets OobMgmtIpv4Address field to given value.

### HasOobMgmtIpv4Address

`func (o *CreateNetworkDevice) HasOobMgmtIpv4Address() bool`

HasOobMgmtIpv4Address returns a boolean if a field has been set.

### GetPrimaryWanIpv6SubnetPoolId

`func (o *CreateNetworkDevice) GetPrimaryWanIpv6SubnetPoolId() float32`

GetPrimaryWanIpv6SubnetPoolId returns the PrimaryWanIpv6SubnetPoolId field if non-nil, zero value otherwise.

### GetPrimaryWanIpv6SubnetPoolIdOk

`func (o *CreateNetworkDevice) GetPrimaryWanIpv6SubnetPoolIdOk() (*float32, bool)`

GetPrimaryWanIpv6SubnetPoolIdOk returns a tuple with the PrimaryWanIpv6SubnetPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWanIpv6SubnetPoolId

`func (o *CreateNetworkDevice) SetPrimaryWanIpv6SubnetPoolId(v float32)`

SetPrimaryWanIpv6SubnetPoolId sets PrimaryWanIpv6SubnetPoolId field to given value.

### HasPrimaryWanIpv6SubnetPoolId

`func (o *CreateNetworkDevice) HasPrimaryWanIpv6SubnetPoolId() bool`

HasPrimaryWanIpv6SubnetPoolId returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


