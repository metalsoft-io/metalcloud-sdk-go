# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **float32** | The id of the server. | 
**Revision** | **float32** | Revision number | 
**ServerTypeId** | Pointer to **float32** | The id of the server type. | [optional] 
**SiteId** | **float32** | The site id where the server is located. | 
**DatacenterName** | **string** | The name of the datacenter where the server is located. | 
**ServerUUID** | Pointer to **string** | The UUID of the server. | [optional] 
**SerialNumber** | Pointer to **string** | The Serial Number of the server. | [optional] 
**ManagementAddress** | Pointer to **string** | The Management Address of the server. | [optional] 
**Username** | Pointer to **string** | The username to use. | [optional] 
**IpmiVersion** | Pointer to **string** | The ipmi version of the server. | [optional] 
**RamGbytes** | Pointer to **float32** | The RAM GB of the server. | [optional] 
**ProcessorCount** | Pointer to **float32** | The processor count of the server. | [optional] 
**ProcessorCoreMhz** | Pointer to **float32** | The processor core Mhz of the server. | [optional] 
**ProcessorCoreCount** | Pointer to **float32** | The processor core count of the server. | [optional] 
**ProcessorName** | Pointer to **string** | The processor name of the server. | [optional] 
**ProcessorCpuMark** | Pointer to **float32** | The processor cpu mark of the server. | [optional] 
**ProcessorThreads** | Pointer to **float32** | The processor threads of the server. | [optional] 
**DiskCount** | Pointer to **float32** | The disk count of the server. | [optional] 
**MgmtSnmpPort** | Pointer to **float32** | The management snmp port of the server. | [optional] 
**BmcMacAddress** | Pointer to **string** | The MAC address of the server. | [optional] 
**BdkDebug** | **float32** | The BDK debug flag. | 
**ServerMetricsMetadata** | Pointer to [**map[string]ServerMetricsInfo**](ServerMetricsInfo.md) | The metrics metadata of the server. | [optional] 
**InstanceCustomInfo** | Pointer to **map[string]interface{}** | The instance custom info of the server. | [optional] 
**CustomInfo** | Pointer to **map[string]interface{}** | The custom info of the server. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the server. | [optional] 
**VendorSkuId** | Pointer to **string** | The vendor sku id of the server. | [optional] 
**Model** | Pointer to **string** | The model of the server. | [optional] 
**VncPort** | Pointer to **float32** | The VNC port of the server. | [optional] 
**IsBasicCampusEndpoint** | Pointer to **float32** | Flag to indicate if the server is basic campus endpoint. | [optional] 
**ServerCleanupPolicyId** | Pointer to **float32** | The cleanup policy id of the server. | [optional] 
**CurrentFirmwareBaselineId** | Pointer to **float32** | The last applied firmware baseline id on the server. | [optional] 
**TargetFirmwareBaselineId** | Pointer to **float32** | The target firmware baseline id on the server. | [optional] 
**RegistrationProfileId** | Pointer to **float32** | The registration profile id of the server. | [optional] 
**RequiresReRegister** | **float32** | Flag to indicate if the server required re-registration. | 
**ServerSupportsSol** | Pointer to **float32** | Flag to indicate if the supports SOL. | [optional] 
**ServerSupportsVirtualMedia** | Pointer to **float32** | Flag to indicate if the supports Virtual Media. | [optional] 
**ServerSupportsOobProvisioning** | Pointer to **float32** | Flag to indicate if the supports OOB provisioning. | [optional] 
**ServerIsProduction** | Pointer to **float32** | Flag to indicate if the server was added as a production (brownfield) server. | [optional] 
**BootingCustomIsoInProgress** | Pointer to **float32** | Flag to indicate if the server is booting a custom iso. | [optional] 
**BiosInfo** | Pointer to [**ServerBiosInfo**](ServerBiosInfo.md) | The bios info of the server. | [optional] 
**VendorInfo** | Pointer to [**ServerVendorInfo**](ServerVendorInfo.md) | The vendor info of the server. | [optional] 
**RegistrationResult** | Pointer to [**ServerRegistrationResult**](ServerRegistrationResult.md) | The registration result of the server. | [optional] 
**ServerClass** | **string** | The class of the server. | 
**ServerStatus** | **string** | The status of the server. | 
**ServerComments** | Pointer to **string** | The comments of the server. | [optional] 
**ServerCapacityMbps** | Pointer to **float32** | The capacity mbps of the server. | [optional] 
**ServerDiskWipe** | Pointer to **float32** | Flag to indicate if the disks need to be wiped. | [optional] 
**RequiresManualCleaning** | Pointer to **float32** | Flag to indicate if the server requires manual cleaning. | [optional] 
**ServerDiskCount** | Pointer to **float32** | The number of disks of the server. | [optional] 
**AdministrationState** | **string** | The administration state of the server. | 
**ServerDhcpStatus** | **string** | The DHCP status of the server. | 
**GpuCount** | Pointer to **float32** | The GPU count of the server. | [optional] 
**GpuInfo** | Pointer to [**[]ServerGpuInfo**](ServerGpuInfo.md) | The GPU info of the server. | [optional] 
**Submodel** | Pointer to **string** | The submodel of the server. | [optional] 
**SupportsFcProvisioning** | **float32** | Flag to indicate if the server supports FC provisioning. | 
**BootLastUpdateTimestamp** | Pointer to **string** | The last boot update timestamp of the server. | [optional] 
**RegisteredTimestamp** | Pointer to **string** | The registration timestamp of the server. | [optional] 
**ServerCreatedTimestamp** | **string** | The creation timestamp of the server. | 
**PowerStatus** | **string** | The power status of the server. | 
**PowerStatusLastUpdateTimestamp** | **string** | The last power update timestamp of the server. | 
**ServerAllocationTimestamp** | Pointer to **string** | The timestamp when the server was allocated. | [optional] 
**JobInfo** | Pointer to [**JobInfo**](JobInfo.md) |  | [optional] 
**ChassisRackId** | Pointer to **float32** | The chassis rack id of the server. | [optional] 
**RackName** | Pointer to **string** | The chassis rack name of the server. | [optional] 
**RackPositionUpperUnit** | Pointer to **string** | The chassis rack upper unit position of the server. | [optional] 
**RackPositionLowerUnit** | Pointer to **string** | The chassis rack lower unit position of the server. | [optional] 
**InventoryId** | Pointer to **string** | The inventory id of the server. | [optional] 
**IpmiCredentialsNeedUpdate** | Pointer to **float32** | Flag to indicate if the server needs an update of the IPMI credentials. | [optional] 
**Interfaces** | Pointer to [**[]ServerInterface**](ServerInterface.md) | The interfaces of the server. | [optional] 
**Disks** | Pointer to [**[]ServerDisk**](ServerDisk.md) | The disks of the server. | [optional] 
**StorageControllers** | Pointer to [**[]ServerStorageController**](ServerStorageController.md) | The storage controllers of the server. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Server. | [optional] 
**ResourcePoolId** | Pointer to **float32** | Resource Pool ID | [optional] 
**AllocationInfo** | Pointer to [**ServerAllocationInfo**](ServerAllocationInfo.md) | The server instance. | [optional] 
**ExtensionInfo** | Pointer to [**ExtensionExecutionInfo**](ExtensionExecutionInfo.md) | The extension execution info of the server. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServer

`func NewServer(serverId float32, revision float32, siteId float32, datacenterName string, bdkDebug float32, requiresReRegister float32, serverClass string, serverStatus string, administrationState string, serverDhcpStatus string, supportsFcProvisioning float32, serverCreatedTimestamp string, powerStatus string, powerStatusLastUpdateTimestamp string, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *Server) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Server) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Server) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetRevision

`func (o *Server) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Server) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Server) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetServerTypeId

`func (o *Server) GetServerTypeId() float32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *Server) GetServerTypeIdOk() (*float32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *Server) SetServerTypeId(v float32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *Server) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetSiteId

`func (o *Server) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Server) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Server) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetDatacenterName

`func (o *Server) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *Server) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *Server) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetServerUUID

`func (o *Server) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *Server) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *Server) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *Server) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Server) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Server) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Server) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Server) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetManagementAddress

`func (o *Server) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *Server) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *Server) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *Server) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetUsername

`func (o *Server) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Server) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Server) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Server) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIpmiVersion

`func (o *Server) GetIpmiVersion() string`

GetIpmiVersion returns the IpmiVersion field if non-nil, zero value otherwise.

### GetIpmiVersionOk

`func (o *Server) GetIpmiVersionOk() (*string, bool)`

GetIpmiVersionOk returns a tuple with the IpmiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiVersion

`func (o *Server) SetIpmiVersion(v string)`

SetIpmiVersion sets IpmiVersion field to given value.

### HasIpmiVersion

`func (o *Server) HasIpmiVersion() bool`

HasIpmiVersion returns a boolean if a field has been set.

### GetRamGbytes

`func (o *Server) GetRamGbytes() float32`

GetRamGbytes returns the RamGbytes field if non-nil, zero value otherwise.

### GetRamGbytesOk

`func (o *Server) GetRamGbytesOk() (*float32, bool)`

GetRamGbytesOk returns a tuple with the RamGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbytes

`func (o *Server) SetRamGbytes(v float32)`

SetRamGbytes sets RamGbytes field to given value.

### HasRamGbytes

`func (o *Server) HasRamGbytes() bool`

HasRamGbytes returns a boolean if a field has been set.

### GetProcessorCount

`func (o *Server) GetProcessorCount() float32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *Server) GetProcessorCountOk() (*float32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *Server) SetProcessorCount(v float32)`

SetProcessorCount sets ProcessorCount field to given value.

### HasProcessorCount

`func (o *Server) HasProcessorCount() bool`

HasProcessorCount returns a boolean if a field has been set.

### GetProcessorCoreMhz

`func (o *Server) GetProcessorCoreMhz() float32`

GetProcessorCoreMhz returns the ProcessorCoreMhz field if non-nil, zero value otherwise.

### GetProcessorCoreMhzOk

`func (o *Server) GetProcessorCoreMhzOk() (*float32, bool)`

GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreMhz

`func (o *Server) SetProcessorCoreMhz(v float32)`

SetProcessorCoreMhz sets ProcessorCoreMhz field to given value.

### HasProcessorCoreMhz

`func (o *Server) HasProcessorCoreMhz() bool`

HasProcessorCoreMhz returns a boolean if a field has been set.

### GetProcessorCoreCount

`func (o *Server) GetProcessorCoreCount() float32`

GetProcessorCoreCount returns the ProcessorCoreCount field if non-nil, zero value otherwise.

### GetProcessorCoreCountOk

`func (o *Server) GetProcessorCoreCountOk() (*float32, bool)`

GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreCount

`func (o *Server) SetProcessorCoreCount(v float32)`

SetProcessorCoreCount sets ProcessorCoreCount field to given value.

### HasProcessorCoreCount

`func (o *Server) HasProcessorCoreCount() bool`

HasProcessorCoreCount returns a boolean if a field has been set.

### GetProcessorName

`func (o *Server) GetProcessorName() string`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *Server) GetProcessorNameOk() (*string, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *Server) SetProcessorName(v string)`

SetProcessorName sets ProcessorName field to given value.

### HasProcessorName

`func (o *Server) HasProcessorName() bool`

HasProcessorName returns a boolean if a field has been set.

### GetProcessorCpuMark

`func (o *Server) GetProcessorCpuMark() float32`

GetProcessorCpuMark returns the ProcessorCpuMark field if non-nil, zero value otherwise.

### GetProcessorCpuMarkOk

`func (o *Server) GetProcessorCpuMarkOk() (*float32, bool)`

GetProcessorCpuMarkOk returns a tuple with the ProcessorCpuMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCpuMark

`func (o *Server) SetProcessorCpuMark(v float32)`

SetProcessorCpuMark sets ProcessorCpuMark field to given value.

### HasProcessorCpuMark

`func (o *Server) HasProcessorCpuMark() bool`

HasProcessorCpuMark returns a boolean if a field has been set.

### GetProcessorThreads

`func (o *Server) GetProcessorThreads() float32`

GetProcessorThreads returns the ProcessorThreads field if non-nil, zero value otherwise.

### GetProcessorThreadsOk

`func (o *Server) GetProcessorThreadsOk() (*float32, bool)`

GetProcessorThreadsOk returns a tuple with the ProcessorThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorThreads

`func (o *Server) SetProcessorThreads(v float32)`

SetProcessorThreads sets ProcessorThreads field to given value.

### HasProcessorThreads

`func (o *Server) HasProcessorThreads() bool`

HasProcessorThreads returns a boolean if a field has been set.

### GetDiskCount

`func (o *Server) GetDiskCount() float32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *Server) GetDiskCountOk() (*float32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *Server) SetDiskCount(v float32)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *Server) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### GetMgmtSnmpPort

`func (o *Server) GetMgmtSnmpPort() float32`

GetMgmtSnmpPort returns the MgmtSnmpPort field if non-nil, zero value otherwise.

### GetMgmtSnmpPortOk

`func (o *Server) GetMgmtSnmpPortOk() (*float32, bool)`

GetMgmtSnmpPortOk returns a tuple with the MgmtSnmpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtSnmpPort

`func (o *Server) SetMgmtSnmpPort(v float32)`

SetMgmtSnmpPort sets MgmtSnmpPort field to given value.

### HasMgmtSnmpPort

`func (o *Server) HasMgmtSnmpPort() bool`

HasMgmtSnmpPort returns a boolean if a field has been set.

### GetBmcMacAddress

`func (o *Server) GetBmcMacAddress() string`

GetBmcMacAddress returns the BmcMacAddress field if non-nil, zero value otherwise.

### GetBmcMacAddressOk

`func (o *Server) GetBmcMacAddressOk() (*string, bool)`

GetBmcMacAddressOk returns a tuple with the BmcMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcMacAddress

`func (o *Server) SetBmcMacAddress(v string)`

SetBmcMacAddress sets BmcMacAddress field to given value.

### HasBmcMacAddress

`func (o *Server) HasBmcMacAddress() bool`

HasBmcMacAddress returns a boolean if a field has been set.

### GetBdkDebug

`func (o *Server) GetBdkDebug() float32`

GetBdkDebug returns the BdkDebug field if non-nil, zero value otherwise.

### GetBdkDebugOk

`func (o *Server) GetBdkDebugOk() (*float32, bool)`

GetBdkDebugOk returns a tuple with the BdkDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdkDebug

`func (o *Server) SetBdkDebug(v float32)`

SetBdkDebug sets BdkDebug field to given value.


### GetServerMetricsMetadata

`func (o *Server) GetServerMetricsMetadata() map[string]ServerMetricsInfo`

GetServerMetricsMetadata returns the ServerMetricsMetadata field if non-nil, zero value otherwise.

### GetServerMetricsMetadataOk

`func (o *Server) GetServerMetricsMetadataOk() (*map[string]ServerMetricsInfo, bool)`

GetServerMetricsMetadataOk returns a tuple with the ServerMetricsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMetricsMetadata

`func (o *Server) SetServerMetricsMetadata(v map[string]ServerMetricsInfo)`

SetServerMetricsMetadata sets ServerMetricsMetadata field to given value.

### HasServerMetricsMetadata

`func (o *Server) HasServerMetricsMetadata() bool`

HasServerMetricsMetadata returns a boolean if a field has been set.

### GetInstanceCustomInfo

`func (o *Server) GetInstanceCustomInfo() map[string]interface{}`

GetInstanceCustomInfo returns the InstanceCustomInfo field if non-nil, zero value otherwise.

### GetInstanceCustomInfoOk

`func (o *Server) GetInstanceCustomInfoOk() (*map[string]interface{}, bool)`

GetInstanceCustomInfoOk returns a tuple with the InstanceCustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCustomInfo

`func (o *Server) SetInstanceCustomInfo(v map[string]interface{})`

SetInstanceCustomInfo sets InstanceCustomInfo field to given value.

### HasInstanceCustomInfo

`func (o *Server) HasInstanceCustomInfo() bool`

HasInstanceCustomInfo returns a boolean if a field has been set.

### GetCustomInfo

`func (o *Server) GetCustomInfo() map[string]interface{}`

GetCustomInfo returns the CustomInfo field if non-nil, zero value otherwise.

### GetCustomInfoOk

`func (o *Server) GetCustomInfoOk() (*map[string]interface{}, bool)`

GetCustomInfoOk returns a tuple with the CustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInfo

`func (o *Server) SetCustomInfo(v map[string]interface{})`

SetCustomInfo sets CustomInfo field to given value.

### HasCustomInfo

`func (o *Server) HasCustomInfo() bool`

HasCustomInfo returns a boolean if a field has been set.

### GetVendor

`func (o *Server) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Server) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Server) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Server) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorSkuId

`func (o *Server) GetVendorSkuId() string`

GetVendorSkuId returns the VendorSkuId field if non-nil, zero value otherwise.

### GetVendorSkuIdOk

`func (o *Server) GetVendorSkuIdOk() (*string, bool)`

GetVendorSkuIdOk returns a tuple with the VendorSkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSkuId

`func (o *Server) SetVendorSkuId(v string)`

SetVendorSkuId sets VendorSkuId field to given value.

### HasVendorSkuId

`func (o *Server) HasVendorSkuId() bool`

HasVendorSkuId returns a boolean if a field has been set.

### GetModel

`func (o *Server) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Server) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Server) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Server) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVncPort

`func (o *Server) GetVncPort() float32`

GetVncPort returns the VncPort field if non-nil, zero value otherwise.

### GetVncPortOk

`func (o *Server) GetVncPortOk() (*float32, bool)`

GetVncPortOk returns a tuple with the VncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPort

`func (o *Server) SetVncPort(v float32)`

SetVncPort sets VncPort field to given value.

### HasVncPort

`func (o *Server) HasVncPort() bool`

HasVncPort returns a boolean if a field has been set.

### GetIsBasicCampusEndpoint

`func (o *Server) GetIsBasicCampusEndpoint() float32`

GetIsBasicCampusEndpoint returns the IsBasicCampusEndpoint field if non-nil, zero value otherwise.

### GetIsBasicCampusEndpointOk

`func (o *Server) GetIsBasicCampusEndpointOk() (*float32, bool)`

GetIsBasicCampusEndpointOk returns a tuple with the IsBasicCampusEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBasicCampusEndpoint

`func (o *Server) SetIsBasicCampusEndpoint(v float32)`

SetIsBasicCampusEndpoint sets IsBasicCampusEndpoint field to given value.

### HasIsBasicCampusEndpoint

`func (o *Server) HasIsBasicCampusEndpoint() bool`

HasIsBasicCampusEndpoint returns a boolean if a field has been set.

### GetServerCleanupPolicyId

`func (o *Server) GetServerCleanupPolicyId() float32`

GetServerCleanupPolicyId returns the ServerCleanupPolicyId field if non-nil, zero value otherwise.

### GetServerCleanupPolicyIdOk

`func (o *Server) GetServerCleanupPolicyIdOk() (*float32, bool)`

GetServerCleanupPolicyIdOk returns a tuple with the ServerCleanupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCleanupPolicyId

`func (o *Server) SetServerCleanupPolicyId(v float32)`

SetServerCleanupPolicyId sets ServerCleanupPolicyId field to given value.

### HasServerCleanupPolicyId

`func (o *Server) HasServerCleanupPolicyId() bool`

HasServerCleanupPolicyId returns a boolean if a field has been set.

### GetCurrentFirmwareBaselineId

`func (o *Server) GetCurrentFirmwareBaselineId() float32`

GetCurrentFirmwareBaselineId returns the CurrentFirmwareBaselineId field if non-nil, zero value otherwise.

### GetCurrentFirmwareBaselineIdOk

`func (o *Server) GetCurrentFirmwareBaselineIdOk() (*float32, bool)`

GetCurrentFirmwareBaselineIdOk returns a tuple with the CurrentFirmwareBaselineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFirmwareBaselineId

`func (o *Server) SetCurrentFirmwareBaselineId(v float32)`

SetCurrentFirmwareBaselineId sets CurrentFirmwareBaselineId field to given value.

### HasCurrentFirmwareBaselineId

`func (o *Server) HasCurrentFirmwareBaselineId() bool`

HasCurrentFirmwareBaselineId returns a boolean if a field has been set.

### GetTargetFirmwareBaselineId

`func (o *Server) GetTargetFirmwareBaselineId() float32`

GetTargetFirmwareBaselineId returns the TargetFirmwareBaselineId field if non-nil, zero value otherwise.

### GetTargetFirmwareBaselineIdOk

`func (o *Server) GetTargetFirmwareBaselineIdOk() (*float32, bool)`

GetTargetFirmwareBaselineIdOk returns a tuple with the TargetFirmwareBaselineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFirmwareBaselineId

`func (o *Server) SetTargetFirmwareBaselineId(v float32)`

SetTargetFirmwareBaselineId sets TargetFirmwareBaselineId field to given value.

### HasTargetFirmwareBaselineId

`func (o *Server) HasTargetFirmwareBaselineId() bool`

HasTargetFirmwareBaselineId returns a boolean if a field has been set.

### GetRegistrationProfileId

`func (o *Server) GetRegistrationProfileId() float32`

GetRegistrationProfileId returns the RegistrationProfileId field if non-nil, zero value otherwise.

### GetRegistrationProfileIdOk

`func (o *Server) GetRegistrationProfileIdOk() (*float32, bool)`

GetRegistrationProfileIdOk returns a tuple with the RegistrationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationProfileId

`func (o *Server) SetRegistrationProfileId(v float32)`

SetRegistrationProfileId sets RegistrationProfileId field to given value.

### HasRegistrationProfileId

`func (o *Server) HasRegistrationProfileId() bool`

HasRegistrationProfileId returns a boolean if a field has been set.

### GetRequiresReRegister

`func (o *Server) GetRequiresReRegister() float32`

GetRequiresReRegister returns the RequiresReRegister field if non-nil, zero value otherwise.

### GetRequiresReRegisterOk

`func (o *Server) GetRequiresReRegisterOk() (*float32, bool)`

GetRequiresReRegisterOk returns a tuple with the RequiresReRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresReRegister

`func (o *Server) SetRequiresReRegister(v float32)`

SetRequiresReRegister sets RequiresReRegister field to given value.


### GetServerSupportsSol

`func (o *Server) GetServerSupportsSol() float32`

GetServerSupportsSol returns the ServerSupportsSol field if non-nil, zero value otherwise.

### GetServerSupportsSolOk

`func (o *Server) GetServerSupportsSolOk() (*float32, bool)`

GetServerSupportsSolOk returns a tuple with the ServerSupportsSol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSupportsSol

`func (o *Server) SetServerSupportsSol(v float32)`

SetServerSupportsSol sets ServerSupportsSol field to given value.

### HasServerSupportsSol

`func (o *Server) HasServerSupportsSol() bool`

HasServerSupportsSol returns a boolean if a field has been set.

### GetServerSupportsVirtualMedia

`func (o *Server) GetServerSupportsVirtualMedia() float32`

GetServerSupportsVirtualMedia returns the ServerSupportsVirtualMedia field if non-nil, zero value otherwise.

### GetServerSupportsVirtualMediaOk

`func (o *Server) GetServerSupportsVirtualMediaOk() (*float32, bool)`

GetServerSupportsVirtualMediaOk returns a tuple with the ServerSupportsVirtualMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSupportsVirtualMedia

`func (o *Server) SetServerSupportsVirtualMedia(v float32)`

SetServerSupportsVirtualMedia sets ServerSupportsVirtualMedia field to given value.

### HasServerSupportsVirtualMedia

`func (o *Server) HasServerSupportsVirtualMedia() bool`

HasServerSupportsVirtualMedia returns a boolean if a field has been set.

### GetServerSupportsOobProvisioning

`func (o *Server) GetServerSupportsOobProvisioning() float32`

GetServerSupportsOobProvisioning returns the ServerSupportsOobProvisioning field if non-nil, zero value otherwise.

### GetServerSupportsOobProvisioningOk

`func (o *Server) GetServerSupportsOobProvisioningOk() (*float32, bool)`

GetServerSupportsOobProvisioningOk returns a tuple with the ServerSupportsOobProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSupportsOobProvisioning

`func (o *Server) SetServerSupportsOobProvisioning(v float32)`

SetServerSupportsOobProvisioning sets ServerSupportsOobProvisioning field to given value.

### HasServerSupportsOobProvisioning

`func (o *Server) HasServerSupportsOobProvisioning() bool`

HasServerSupportsOobProvisioning returns a boolean if a field has been set.

### GetServerIsProduction

`func (o *Server) GetServerIsProduction() float32`

GetServerIsProduction returns the ServerIsProduction field if non-nil, zero value otherwise.

### GetServerIsProductionOk

`func (o *Server) GetServerIsProductionOk() (*float32, bool)`

GetServerIsProductionOk returns a tuple with the ServerIsProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIsProduction

`func (o *Server) SetServerIsProduction(v float32)`

SetServerIsProduction sets ServerIsProduction field to given value.

### HasServerIsProduction

`func (o *Server) HasServerIsProduction() bool`

HasServerIsProduction returns a boolean if a field has been set.

### GetBootingCustomIsoInProgress

`func (o *Server) GetBootingCustomIsoInProgress() float32`

GetBootingCustomIsoInProgress returns the BootingCustomIsoInProgress field if non-nil, zero value otherwise.

### GetBootingCustomIsoInProgressOk

`func (o *Server) GetBootingCustomIsoInProgressOk() (*float32, bool)`

GetBootingCustomIsoInProgressOk returns a tuple with the BootingCustomIsoInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootingCustomIsoInProgress

`func (o *Server) SetBootingCustomIsoInProgress(v float32)`

SetBootingCustomIsoInProgress sets BootingCustomIsoInProgress field to given value.

### HasBootingCustomIsoInProgress

`func (o *Server) HasBootingCustomIsoInProgress() bool`

HasBootingCustomIsoInProgress returns a boolean if a field has been set.

### GetBiosInfo

`func (o *Server) GetBiosInfo() ServerBiosInfo`

GetBiosInfo returns the BiosInfo field if non-nil, zero value otherwise.

### GetBiosInfoOk

`func (o *Server) GetBiosInfoOk() (*ServerBiosInfo, bool)`

GetBiosInfoOk returns a tuple with the BiosInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosInfo

`func (o *Server) SetBiosInfo(v ServerBiosInfo)`

SetBiosInfo sets BiosInfo field to given value.

### HasBiosInfo

`func (o *Server) HasBiosInfo() bool`

HasBiosInfo returns a boolean if a field has been set.

### GetVendorInfo

`func (o *Server) GetVendorInfo() ServerVendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *Server) GetVendorInfoOk() (*ServerVendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *Server) SetVendorInfo(v ServerVendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *Server) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetRegistrationResult

`func (o *Server) GetRegistrationResult() ServerRegistrationResult`

GetRegistrationResult returns the RegistrationResult field if non-nil, zero value otherwise.

### GetRegistrationResultOk

`func (o *Server) GetRegistrationResultOk() (*ServerRegistrationResult, bool)`

GetRegistrationResultOk returns a tuple with the RegistrationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationResult

`func (o *Server) SetRegistrationResult(v ServerRegistrationResult)`

SetRegistrationResult sets RegistrationResult field to given value.

### HasRegistrationResult

`func (o *Server) HasRegistrationResult() bool`

HasRegistrationResult returns a boolean if a field has been set.

### GetServerClass

`func (o *Server) GetServerClass() string`

GetServerClass returns the ServerClass field if non-nil, zero value otherwise.

### GetServerClassOk

`func (o *Server) GetServerClassOk() (*string, bool)`

GetServerClassOk returns a tuple with the ServerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClass

`func (o *Server) SetServerClass(v string)`

SetServerClass sets ServerClass field to given value.


### GetServerStatus

`func (o *Server) GetServerStatus() string`

GetServerStatus returns the ServerStatus field if non-nil, zero value otherwise.

### GetServerStatusOk

`func (o *Server) GetServerStatusOk() (*string, bool)`

GetServerStatusOk returns a tuple with the ServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatus

`func (o *Server) SetServerStatus(v string)`

SetServerStatus sets ServerStatus field to given value.


### GetServerComments

`func (o *Server) GetServerComments() string`

GetServerComments returns the ServerComments field if non-nil, zero value otherwise.

### GetServerCommentsOk

`func (o *Server) GetServerCommentsOk() (*string, bool)`

GetServerCommentsOk returns a tuple with the ServerComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComments

`func (o *Server) SetServerComments(v string)`

SetServerComments sets ServerComments field to given value.

### HasServerComments

`func (o *Server) HasServerComments() bool`

HasServerComments returns a boolean if a field has been set.

### GetServerCapacityMbps

`func (o *Server) GetServerCapacityMbps() float32`

GetServerCapacityMbps returns the ServerCapacityMbps field if non-nil, zero value otherwise.

### GetServerCapacityMbpsOk

`func (o *Server) GetServerCapacityMbpsOk() (*float32, bool)`

GetServerCapacityMbpsOk returns a tuple with the ServerCapacityMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCapacityMbps

`func (o *Server) SetServerCapacityMbps(v float32)`

SetServerCapacityMbps sets ServerCapacityMbps field to given value.

### HasServerCapacityMbps

`func (o *Server) HasServerCapacityMbps() bool`

HasServerCapacityMbps returns a boolean if a field has been set.

### GetServerDiskWipe

`func (o *Server) GetServerDiskWipe() float32`

GetServerDiskWipe returns the ServerDiskWipe field if non-nil, zero value otherwise.

### GetServerDiskWipeOk

`func (o *Server) GetServerDiskWipeOk() (*float32, bool)`

GetServerDiskWipeOk returns a tuple with the ServerDiskWipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDiskWipe

`func (o *Server) SetServerDiskWipe(v float32)`

SetServerDiskWipe sets ServerDiskWipe field to given value.

### HasServerDiskWipe

`func (o *Server) HasServerDiskWipe() bool`

HasServerDiskWipe returns a boolean if a field has been set.

### GetRequiresManualCleaning

`func (o *Server) GetRequiresManualCleaning() float32`

GetRequiresManualCleaning returns the RequiresManualCleaning field if non-nil, zero value otherwise.

### GetRequiresManualCleaningOk

`func (o *Server) GetRequiresManualCleaningOk() (*float32, bool)`

GetRequiresManualCleaningOk returns a tuple with the RequiresManualCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresManualCleaning

`func (o *Server) SetRequiresManualCleaning(v float32)`

SetRequiresManualCleaning sets RequiresManualCleaning field to given value.

### HasRequiresManualCleaning

`func (o *Server) HasRequiresManualCleaning() bool`

HasRequiresManualCleaning returns a boolean if a field has been set.

### GetServerDiskCount

`func (o *Server) GetServerDiskCount() float32`

GetServerDiskCount returns the ServerDiskCount field if non-nil, zero value otherwise.

### GetServerDiskCountOk

`func (o *Server) GetServerDiskCountOk() (*float32, bool)`

GetServerDiskCountOk returns a tuple with the ServerDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDiskCount

`func (o *Server) SetServerDiskCount(v float32)`

SetServerDiskCount sets ServerDiskCount field to given value.

### HasServerDiskCount

`func (o *Server) HasServerDiskCount() bool`

HasServerDiskCount returns a boolean if a field has been set.

### GetAdministrationState

`func (o *Server) GetAdministrationState() string`

GetAdministrationState returns the AdministrationState field if non-nil, zero value otherwise.

### GetAdministrationStateOk

`func (o *Server) GetAdministrationStateOk() (*string, bool)`

GetAdministrationStateOk returns a tuple with the AdministrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrationState

`func (o *Server) SetAdministrationState(v string)`

SetAdministrationState sets AdministrationState field to given value.


### GetServerDhcpStatus

`func (o *Server) GetServerDhcpStatus() string`

GetServerDhcpStatus returns the ServerDhcpStatus field if non-nil, zero value otherwise.

### GetServerDhcpStatusOk

`func (o *Server) GetServerDhcpStatusOk() (*string, bool)`

GetServerDhcpStatusOk returns a tuple with the ServerDhcpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDhcpStatus

`func (o *Server) SetServerDhcpStatus(v string)`

SetServerDhcpStatus sets ServerDhcpStatus field to given value.


### GetGpuCount

`func (o *Server) GetGpuCount() float32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *Server) GetGpuCountOk() (*float32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *Server) SetGpuCount(v float32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *Server) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuInfo

`func (o *Server) GetGpuInfo() []ServerGpuInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *Server) GetGpuInfoOk() (*[]ServerGpuInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *Server) SetGpuInfo(v []ServerGpuInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *Server) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetSubmodel

`func (o *Server) GetSubmodel() string`

GetSubmodel returns the Submodel field if non-nil, zero value otherwise.

### GetSubmodelOk

`func (o *Server) GetSubmodelOk() (*string, bool)`

GetSubmodelOk returns a tuple with the Submodel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmodel

`func (o *Server) SetSubmodel(v string)`

SetSubmodel sets Submodel field to given value.

### HasSubmodel

`func (o *Server) HasSubmodel() bool`

HasSubmodel returns a boolean if a field has been set.

### GetSupportsFcProvisioning

`func (o *Server) GetSupportsFcProvisioning() float32`

GetSupportsFcProvisioning returns the SupportsFcProvisioning field if non-nil, zero value otherwise.

### GetSupportsFcProvisioningOk

`func (o *Server) GetSupportsFcProvisioningOk() (*float32, bool)`

GetSupportsFcProvisioningOk returns a tuple with the SupportsFcProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFcProvisioning

`func (o *Server) SetSupportsFcProvisioning(v float32)`

SetSupportsFcProvisioning sets SupportsFcProvisioning field to given value.


### GetBootLastUpdateTimestamp

`func (o *Server) GetBootLastUpdateTimestamp() string`

GetBootLastUpdateTimestamp returns the BootLastUpdateTimestamp field if non-nil, zero value otherwise.

### GetBootLastUpdateTimestampOk

`func (o *Server) GetBootLastUpdateTimestampOk() (*string, bool)`

GetBootLastUpdateTimestampOk returns a tuple with the BootLastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootLastUpdateTimestamp

`func (o *Server) SetBootLastUpdateTimestamp(v string)`

SetBootLastUpdateTimestamp sets BootLastUpdateTimestamp field to given value.

### HasBootLastUpdateTimestamp

`func (o *Server) HasBootLastUpdateTimestamp() bool`

HasBootLastUpdateTimestamp returns a boolean if a field has been set.

### GetRegisteredTimestamp

`func (o *Server) GetRegisteredTimestamp() string`

GetRegisteredTimestamp returns the RegisteredTimestamp field if non-nil, zero value otherwise.

### GetRegisteredTimestampOk

`func (o *Server) GetRegisteredTimestampOk() (*string, bool)`

GetRegisteredTimestampOk returns a tuple with the RegisteredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredTimestamp

`func (o *Server) SetRegisteredTimestamp(v string)`

SetRegisteredTimestamp sets RegisteredTimestamp field to given value.

### HasRegisteredTimestamp

`func (o *Server) HasRegisteredTimestamp() bool`

HasRegisteredTimestamp returns a boolean if a field has been set.

### GetServerCreatedTimestamp

`func (o *Server) GetServerCreatedTimestamp() string`

GetServerCreatedTimestamp returns the ServerCreatedTimestamp field if non-nil, zero value otherwise.

### GetServerCreatedTimestampOk

`func (o *Server) GetServerCreatedTimestampOk() (*string, bool)`

GetServerCreatedTimestampOk returns a tuple with the ServerCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCreatedTimestamp

`func (o *Server) SetServerCreatedTimestamp(v string)`

SetServerCreatedTimestamp sets ServerCreatedTimestamp field to given value.


### GetPowerStatus

`func (o *Server) GetPowerStatus() string`

GetPowerStatus returns the PowerStatus field if non-nil, zero value otherwise.

### GetPowerStatusOk

`func (o *Server) GetPowerStatusOk() (*string, bool)`

GetPowerStatusOk returns a tuple with the PowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStatus

`func (o *Server) SetPowerStatus(v string)`

SetPowerStatus sets PowerStatus field to given value.


### GetPowerStatusLastUpdateTimestamp

`func (o *Server) GetPowerStatusLastUpdateTimestamp() string`

GetPowerStatusLastUpdateTimestamp returns the PowerStatusLastUpdateTimestamp field if non-nil, zero value otherwise.

### GetPowerStatusLastUpdateTimestampOk

`func (o *Server) GetPowerStatusLastUpdateTimestampOk() (*string, bool)`

GetPowerStatusLastUpdateTimestampOk returns a tuple with the PowerStatusLastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStatusLastUpdateTimestamp

`func (o *Server) SetPowerStatusLastUpdateTimestamp(v string)`

SetPowerStatusLastUpdateTimestamp sets PowerStatusLastUpdateTimestamp field to given value.


### GetServerAllocationTimestamp

`func (o *Server) GetServerAllocationTimestamp() string`

GetServerAllocationTimestamp returns the ServerAllocationTimestamp field if non-nil, zero value otherwise.

### GetServerAllocationTimestampOk

`func (o *Server) GetServerAllocationTimestampOk() (*string, bool)`

GetServerAllocationTimestampOk returns a tuple with the ServerAllocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllocationTimestamp

`func (o *Server) SetServerAllocationTimestamp(v string)`

SetServerAllocationTimestamp sets ServerAllocationTimestamp field to given value.

### HasServerAllocationTimestamp

`func (o *Server) HasServerAllocationTimestamp() bool`

HasServerAllocationTimestamp returns a boolean if a field has been set.

### GetJobInfo

`func (o *Server) GetJobInfo() JobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *Server) GetJobInfoOk() (*JobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *Server) SetJobInfo(v JobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *Server) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### GetChassisRackId

`func (o *Server) GetChassisRackId() float32`

GetChassisRackId returns the ChassisRackId field if non-nil, zero value otherwise.

### GetChassisRackIdOk

`func (o *Server) GetChassisRackIdOk() (*float32, bool)`

GetChassisRackIdOk returns a tuple with the ChassisRackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisRackId

`func (o *Server) SetChassisRackId(v float32)`

SetChassisRackId sets ChassisRackId field to given value.

### HasChassisRackId

`func (o *Server) HasChassisRackId() bool`

HasChassisRackId returns a boolean if a field has been set.

### GetRackName

`func (o *Server) GetRackName() string`

GetRackName returns the RackName field if non-nil, zero value otherwise.

### GetRackNameOk

`func (o *Server) GetRackNameOk() (*string, bool)`

GetRackNameOk returns a tuple with the RackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackName

`func (o *Server) SetRackName(v string)`

SetRackName sets RackName field to given value.

### HasRackName

`func (o *Server) HasRackName() bool`

HasRackName returns a boolean if a field has been set.

### GetRackPositionUpperUnit

`func (o *Server) GetRackPositionUpperUnit() string`

GetRackPositionUpperUnit returns the RackPositionUpperUnit field if non-nil, zero value otherwise.

### GetRackPositionUpperUnitOk

`func (o *Server) GetRackPositionUpperUnitOk() (*string, bool)`

GetRackPositionUpperUnitOk returns a tuple with the RackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionUpperUnit

`func (o *Server) SetRackPositionUpperUnit(v string)`

SetRackPositionUpperUnit sets RackPositionUpperUnit field to given value.

### HasRackPositionUpperUnit

`func (o *Server) HasRackPositionUpperUnit() bool`

HasRackPositionUpperUnit returns a boolean if a field has been set.

### GetRackPositionLowerUnit

`func (o *Server) GetRackPositionLowerUnit() string`

GetRackPositionLowerUnit returns the RackPositionLowerUnit field if non-nil, zero value otherwise.

### GetRackPositionLowerUnitOk

`func (o *Server) GetRackPositionLowerUnitOk() (*string, bool)`

GetRackPositionLowerUnitOk returns a tuple with the RackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionLowerUnit

`func (o *Server) SetRackPositionLowerUnit(v string)`

SetRackPositionLowerUnit sets RackPositionLowerUnit field to given value.

### HasRackPositionLowerUnit

`func (o *Server) HasRackPositionLowerUnit() bool`

HasRackPositionLowerUnit returns a boolean if a field has been set.

### GetInventoryId

`func (o *Server) GetInventoryId() string`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *Server) GetInventoryIdOk() (*string, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *Server) SetInventoryId(v string)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *Server) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetIpmiCredentialsNeedUpdate

`func (o *Server) GetIpmiCredentialsNeedUpdate() float32`

GetIpmiCredentialsNeedUpdate returns the IpmiCredentialsNeedUpdate field if non-nil, zero value otherwise.

### GetIpmiCredentialsNeedUpdateOk

`func (o *Server) GetIpmiCredentialsNeedUpdateOk() (*float32, bool)`

GetIpmiCredentialsNeedUpdateOk returns a tuple with the IpmiCredentialsNeedUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiCredentialsNeedUpdate

`func (o *Server) SetIpmiCredentialsNeedUpdate(v float32)`

SetIpmiCredentialsNeedUpdate sets IpmiCredentialsNeedUpdate field to given value.

### HasIpmiCredentialsNeedUpdate

`func (o *Server) HasIpmiCredentialsNeedUpdate() bool`

HasIpmiCredentialsNeedUpdate returns a boolean if a field has been set.

### GetInterfaces

`func (o *Server) GetInterfaces() []ServerInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Server) GetInterfacesOk() (*[]ServerInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Server) SetInterfaces(v []ServerInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Server) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetDisks

`func (o *Server) GetDisks() []ServerDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Server) GetDisksOk() (*[]ServerDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Server) SetDisks(v []ServerDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Server) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetStorageControllers

`func (o *Server) GetStorageControllers() []ServerStorageController`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *Server) GetStorageControllersOk() (*[]ServerStorageController, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *Server) SetStorageControllers(v []ServerStorageController)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *Server) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetTags

`func (o *Server) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Server) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Server) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Server) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *Server) GetResourcePoolId() float32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *Server) GetResourcePoolIdOk() (*float32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *Server) SetResourcePoolId(v float32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *Server) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAllocationInfo

`func (o *Server) GetAllocationInfo() ServerAllocationInfo`

GetAllocationInfo returns the AllocationInfo field if non-nil, zero value otherwise.

### GetAllocationInfoOk

`func (o *Server) GetAllocationInfoOk() (*ServerAllocationInfo, bool)`

GetAllocationInfoOk returns a tuple with the AllocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationInfo

`func (o *Server) SetAllocationInfo(v ServerAllocationInfo)`

SetAllocationInfo sets AllocationInfo field to given value.

### HasAllocationInfo

`func (o *Server) HasAllocationInfo() bool`

HasAllocationInfo returns a boolean if a field has been set.

### GetExtensionInfo

`func (o *Server) GetExtensionInfo() ExtensionExecutionInfo`

GetExtensionInfo returns the ExtensionInfo field if non-nil, zero value otherwise.

### GetExtensionInfoOk

`func (o *Server) GetExtensionInfoOk() (*ExtensionExecutionInfo, bool)`

GetExtensionInfoOk returns a tuple with the ExtensionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInfo

`func (o *Server) SetExtensionInfo(v ExtensionExecutionInfo)`

SetExtensionInfo sets ExtensionInfo field to given value.

### HasExtensionInfo

`func (o *Server) HasExtensionInfo() bool`

HasExtensionInfo returns a boolean if a field has been set.

### GetLinks

`func (o *Server) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Server) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Server) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Server) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


