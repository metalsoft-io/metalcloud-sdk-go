# QuotaProfileLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureServerGroupMaxCount** | **float32** | Maximum number of server groups per infrastructure | 
**InfrastructureDriveMaxCount** | **float32** | Maximum number of drives per infrastructure | 
**InfrastructureFileShareMaxCount** | **float32** | Maximum number of file shares per infrastructure | 
**InfrastructureBucketMaxCount** | **float32** | Maximum number of buckets per infrastructure | 
**InfrastructureVmInstanceGroupMaxCount** | **float32** | Maximum number of VM instance groups per infrastructure | 
**ServerGroupInstancesMaxCount** | **float32** | Maximum number of instances per server group | 
**ServerGroupInstancesMinCount** | **float32** | Minimum number of instances per server group | 
**VmInstanceGroupVmInstancesMaxCount** | **float32** | Maximum number of VM instances per VM instance group | 
**VmInstanceMaxDiskSizeMbytes** | **float32** | Maximum disk size for a VM instance in megabytes | 
**DriveMaxSizeMbytes** | **float32** | Maximum shared drive size in megabytes | 
**DriveMinSizeMbytes** | **float32** | Minimum shared drive size in megabytes | 
**FileShareMinSizeGb** | **float32** | Minimum file share size in gigabytes | 
**FileShareMaxSizeGb** | **float32** | Maximum file share size in gigabytes | 
**BucketMinSizeGb** | **float32** | Minimum bucket size in gigabytes | 
**BucketMaxSizeGb** | **float32** | Maximum bucket size in gigabytes | 
**ShowOperatingSystemImagesTab** | **bool** | Whether the operating system images tab is visible in the UI | 
**ShowTemplateAssetsView** | **bool** | Whether the template assets view is visible in the UI | 
**UserResourceServerTypeNameToMaxCount** | **map[string]interface{}** | Map of server type name to maximum instance count allowed for the user | 
**UserSshKeysCountMax** | **float32** | Maximum number of SSH keys a user can have | 
**ShowLegacyPages** | **bool** | Whether legacy UI pages are visible | 
**ShowEliChatBot** | **bool** | Whether the ELI AI chat bot is visible | 
**EnableCustomRaidConfiguration** | **bool** | Whether custom RAID configuration is allowed | 
**EnableInfrastructureVmInstance** | **bool** | Whether VM instance groups can be created in infrastructures | 
**EnableInfrastructureExtensions** | **bool** | Whether extensions can be deployed in infrastructures | 
**AllowedInfrastructureExtensions** | **[]string** | Allowed extension IDs. Empty array means all extensions are permitted | 
**AllowedServerTypes** | **[]string** | Allowed server type names. Empty array means all server types are permitted | 
**AllowedSites** | **[]string** | Allowed site names. Empty array means all sites are permitted | 
**AllowedLogicalNetworkProfiles** | **[]string** | Allowed logical network profile names. Empty array means all profiles are permitted | 
**AllowedPreCreatedLogicalNetworks** | **[]string** | Allowed pre-created logical network names (those with no infrastructure). Empty array means all are permitted | 

## Methods

### NewQuotaProfileLimits

`func NewQuotaProfileLimits(infrastructureServerGroupMaxCount float32, infrastructureDriveMaxCount float32, infrastructureFileShareMaxCount float32, infrastructureBucketMaxCount float32, infrastructureVmInstanceGroupMaxCount float32, serverGroupInstancesMaxCount float32, serverGroupInstancesMinCount float32, vmInstanceGroupVmInstancesMaxCount float32, vmInstanceMaxDiskSizeMbytes float32, driveMaxSizeMbytes float32, driveMinSizeMbytes float32, fileShareMinSizeGb float32, fileShareMaxSizeGb float32, bucketMinSizeGb float32, bucketMaxSizeGb float32, showOperatingSystemImagesTab bool, showTemplateAssetsView bool, userResourceServerTypeNameToMaxCount map[string]interface{}, userSshKeysCountMax float32, showLegacyPages bool, showEliChatBot bool, enableCustomRaidConfiguration bool, enableInfrastructureVmInstance bool, enableInfrastructureExtensions bool, allowedInfrastructureExtensions []string, allowedServerTypes []string, allowedSites []string, allowedLogicalNetworkProfiles []string, allowedPreCreatedLogicalNetworks []string, ) *QuotaProfileLimits`

NewQuotaProfileLimits instantiates a new QuotaProfileLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaProfileLimitsWithDefaults

`func NewQuotaProfileLimitsWithDefaults() *QuotaProfileLimits`

NewQuotaProfileLimitsWithDefaults instantiates a new QuotaProfileLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureServerGroupMaxCount

`func (o *QuotaProfileLimits) GetInfrastructureServerGroupMaxCount() float32`

GetInfrastructureServerGroupMaxCount returns the InfrastructureServerGroupMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureServerGroupMaxCountOk

`func (o *QuotaProfileLimits) GetInfrastructureServerGroupMaxCountOk() (*float32, bool)`

GetInfrastructureServerGroupMaxCountOk returns a tuple with the InfrastructureServerGroupMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureServerGroupMaxCount

`func (o *QuotaProfileLimits) SetInfrastructureServerGroupMaxCount(v float32)`

SetInfrastructureServerGroupMaxCount sets InfrastructureServerGroupMaxCount field to given value.


### GetInfrastructureDriveMaxCount

`func (o *QuotaProfileLimits) GetInfrastructureDriveMaxCount() float32`

GetInfrastructureDriveMaxCount returns the InfrastructureDriveMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureDriveMaxCountOk

`func (o *QuotaProfileLimits) GetInfrastructureDriveMaxCountOk() (*float32, bool)`

GetInfrastructureDriveMaxCountOk returns a tuple with the InfrastructureDriveMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDriveMaxCount

`func (o *QuotaProfileLimits) SetInfrastructureDriveMaxCount(v float32)`

SetInfrastructureDriveMaxCount sets InfrastructureDriveMaxCount field to given value.


### GetInfrastructureFileShareMaxCount

`func (o *QuotaProfileLimits) GetInfrastructureFileShareMaxCount() float32`

GetInfrastructureFileShareMaxCount returns the InfrastructureFileShareMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureFileShareMaxCountOk

`func (o *QuotaProfileLimits) GetInfrastructureFileShareMaxCountOk() (*float32, bool)`

GetInfrastructureFileShareMaxCountOk returns a tuple with the InfrastructureFileShareMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureFileShareMaxCount

`func (o *QuotaProfileLimits) SetInfrastructureFileShareMaxCount(v float32)`

SetInfrastructureFileShareMaxCount sets InfrastructureFileShareMaxCount field to given value.


### GetInfrastructureBucketMaxCount

`func (o *QuotaProfileLimits) GetInfrastructureBucketMaxCount() float32`

GetInfrastructureBucketMaxCount returns the InfrastructureBucketMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureBucketMaxCountOk

`func (o *QuotaProfileLimits) GetInfrastructureBucketMaxCountOk() (*float32, bool)`

GetInfrastructureBucketMaxCountOk returns a tuple with the InfrastructureBucketMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureBucketMaxCount

`func (o *QuotaProfileLimits) SetInfrastructureBucketMaxCount(v float32)`

SetInfrastructureBucketMaxCount sets InfrastructureBucketMaxCount field to given value.


### GetInfrastructureVmInstanceGroupMaxCount

`func (o *QuotaProfileLimits) GetInfrastructureVmInstanceGroupMaxCount() float32`

GetInfrastructureVmInstanceGroupMaxCount returns the InfrastructureVmInstanceGroupMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureVmInstanceGroupMaxCountOk

`func (o *QuotaProfileLimits) GetInfrastructureVmInstanceGroupMaxCountOk() (*float32, bool)`

GetInfrastructureVmInstanceGroupMaxCountOk returns a tuple with the InfrastructureVmInstanceGroupMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureVmInstanceGroupMaxCount

`func (o *QuotaProfileLimits) SetInfrastructureVmInstanceGroupMaxCount(v float32)`

SetInfrastructureVmInstanceGroupMaxCount sets InfrastructureVmInstanceGroupMaxCount field to given value.


### GetServerGroupInstancesMaxCount

`func (o *QuotaProfileLimits) GetServerGroupInstancesMaxCount() float32`

GetServerGroupInstancesMaxCount returns the ServerGroupInstancesMaxCount field if non-nil, zero value otherwise.

### GetServerGroupInstancesMaxCountOk

`func (o *QuotaProfileLimits) GetServerGroupInstancesMaxCountOk() (*float32, bool)`

GetServerGroupInstancesMaxCountOk returns a tuple with the ServerGroupInstancesMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupInstancesMaxCount

`func (o *QuotaProfileLimits) SetServerGroupInstancesMaxCount(v float32)`

SetServerGroupInstancesMaxCount sets ServerGroupInstancesMaxCount field to given value.


### GetServerGroupInstancesMinCount

`func (o *QuotaProfileLimits) GetServerGroupInstancesMinCount() float32`

GetServerGroupInstancesMinCount returns the ServerGroupInstancesMinCount field if non-nil, zero value otherwise.

### GetServerGroupInstancesMinCountOk

`func (o *QuotaProfileLimits) GetServerGroupInstancesMinCountOk() (*float32, bool)`

GetServerGroupInstancesMinCountOk returns a tuple with the ServerGroupInstancesMinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupInstancesMinCount

`func (o *QuotaProfileLimits) SetServerGroupInstancesMinCount(v float32)`

SetServerGroupInstancesMinCount sets ServerGroupInstancesMinCount field to given value.


### GetVmInstanceGroupVmInstancesMaxCount

`func (o *QuotaProfileLimits) GetVmInstanceGroupVmInstancesMaxCount() float32`

GetVmInstanceGroupVmInstancesMaxCount returns the VmInstanceGroupVmInstancesMaxCount field if non-nil, zero value otherwise.

### GetVmInstanceGroupVmInstancesMaxCountOk

`func (o *QuotaProfileLimits) GetVmInstanceGroupVmInstancesMaxCountOk() (*float32, bool)`

GetVmInstanceGroupVmInstancesMaxCountOk returns a tuple with the VmInstanceGroupVmInstancesMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceGroupVmInstancesMaxCount

`func (o *QuotaProfileLimits) SetVmInstanceGroupVmInstancesMaxCount(v float32)`

SetVmInstanceGroupVmInstancesMaxCount sets VmInstanceGroupVmInstancesMaxCount field to given value.


### GetVmInstanceMaxDiskSizeMbytes

`func (o *QuotaProfileLimits) GetVmInstanceMaxDiskSizeMbytes() float32`

GetVmInstanceMaxDiskSizeMbytes returns the VmInstanceMaxDiskSizeMbytes field if non-nil, zero value otherwise.

### GetVmInstanceMaxDiskSizeMbytesOk

`func (o *QuotaProfileLimits) GetVmInstanceMaxDiskSizeMbytesOk() (*float32, bool)`

GetVmInstanceMaxDiskSizeMbytesOk returns a tuple with the VmInstanceMaxDiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceMaxDiskSizeMbytes

`func (o *QuotaProfileLimits) SetVmInstanceMaxDiskSizeMbytes(v float32)`

SetVmInstanceMaxDiskSizeMbytes sets VmInstanceMaxDiskSizeMbytes field to given value.


### GetDriveMaxSizeMbytes

`func (o *QuotaProfileLimits) GetDriveMaxSizeMbytes() float32`

GetDriveMaxSizeMbytes returns the DriveMaxSizeMbytes field if non-nil, zero value otherwise.

### GetDriveMaxSizeMbytesOk

`func (o *QuotaProfileLimits) GetDriveMaxSizeMbytesOk() (*float32, bool)`

GetDriveMaxSizeMbytesOk returns a tuple with the DriveMaxSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveMaxSizeMbytes

`func (o *QuotaProfileLimits) SetDriveMaxSizeMbytes(v float32)`

SetDriveMaxSizeMbytes sets DriveMaxSizeMbytes field to given value.


### GetDriveMinSizeMbytes

`func (o *QuotaProfileLimits) GetDriveMinSizeMbytes() float32`

GetDriveMinSizeMbytes returns the DriveMinSizeMbytes field if non-nil, zero value otherwise.

### GetDriveMinSizeMbytesOk

`func (o *QuotaProfileLimits) GetDriveMinSizeMbytesOk() (*float32, bool)`

GetDriveMinSizeMbytesOk returns a tuple with the DriveMinSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveMinSizeMbytes

`func (o *QuotaProfileLimits) SetDriveMinSizeMbytes(v float32)`

SetDriveMinSizeMbytes sets DriveMinSizeMbytes field to given value.


### GetFileShareMinSizeGb

`func (o *QuotaProfileLimits) GetFileShareMinSizeGb() float32`

GetFileShareMinSizeGb returns the FileShareMinSizeGb field if non-nil, zero value otherwise.

### GetFileShareMinSizeGbOk

`func (o *QuotaProfileLimits) GetFileShareMinSizeGbOk() (*float32, bool)`

GetFileShareMinSizeGbOk returns a tuple with the FileShareMinSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShareMinSizeGb

`func (o *QuotaProfileLimits) SetFileShareMinSizeGb(v float32)`

SetFileShareMinSizeGb sets FileShareMinSizeGb field to given value.


### GetFileShareMaxSizeGb

`func (o *QuotaProfileLimits) GetFileShareMaxSizeGb() float32`

GetFileShareMaxSizeGb returns the FileShareMaxSizeGb field if non-nil, zero value otherwise.

### GetFileShareMaxSizeGbOk

`func (o *QuotaProfileLimits) GetFileShareMaxSizeGbOk() (*float32, bool)`

GetFileShareMaxSizeGbOk returns a tuple with the FileShareMaxSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShareMaxSizeGb

`func (o *QuotaProfileLimits) SetFileShareMaxSizeGb(v float32)`

SetFileShareMaxSizeGb sets FileShareMaxSizeGb field to given value.


### GetBucketMinSizeGb

`func (o *QuotaProfileLimits) GetBucketMinSizeGb() float32`

GetBucketMinSizeGb returns the BucketMinSizeGb field if non-nil, zero value otherwise.

### GetBucketMinSizeGbOk

`func (o *QuotaProfileLimits) GetBucketMinSizeGbOk() (*float32, bool)`

GetBucketMinSizeGbOk returns a tuple with the BucketMinSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketMinSizeGb

`func (o *QuotaProfileLimits) SetBucketMinSizeGb(v float32)`

SetBucketMinSizeGb sets BucketMinSizeGb field to given value.


### GetBucketMaxSizeGb

`func (o *QuotaProfileLimits) GetBucketMaxSizeGb() float32`

GetBucketMaxSizeGb returns the BucketMaxSizeGb field if non-nil, zero value otherwise.

### GetBucketMaxSizeGbOk

`func (o *QuotaProfileLimits) GetBucketMaxSizeGbOk() (*float32, bool)`

GetBucketMaxSizeGbOk returns a tuple with the BucketMaxSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketMaxSizeGb

`func (o *QuotaProfileLimits) SetBucketMaxSizeGb(v float32)`

SetBucketMaxSizeGb sets BucketMaxSizeGb field to given value.


### GetShowOperatingSystemImagesTab

`func (o *QuotaProfileLimits) GetShowOperatingSystemImagesTab() bool`

GetShowOperatingSystemImagesTab returns the ShowOperatingSystemImagesTab field if non-nil, zero value otherwise.

### GetShowOperatingSystemImagesTabOk

`func (o *QuotaProfileLimits) GetShowOperatingSystemImagesTabOk() (*bool, bool)`

GetShowOperatingSystemImagesTabOk returns a tuple with the ShowOperatingSystemImagesTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOperatingSystemImagesTab

`func (o *QuotaProfileLimits) SetShowOperatingSystemImagesTab(v bool)`

SetShowOperatingSystemImagesTab sets ShowOperatingSystemImagesTab field to given value.


### GetShowTemplateAssetsView

`func (o *QuotaProfileLimits) GetShowTemplateAssetsView() bool`

GetShowTemplateAssetsView returns the ShowTemplateAssetsView field if non-nil, zero value otherwise.

### GetShowTemplateAssetsViewOk

`func (o *QuotaProfileLimits) GetShowTemplateAssetsViewOk() (*bool, bool)`

GetShowTemplateAssetsViewOk returns a tuple with the ShowTemplateAssetsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTemplateAssetsView

`func (o *QuotaProfileLimits) SetShowTemplateAssetsView(v bool)`

SetShowTemplateAssetsView sets ShowTemplateAssetsView field to given value.


### GetUserResourceServerTypeNameToMaxCount

`func (o *QuotaProfileLimits) GetUserResourceServerTypeNameToMaxCount() map[string]interface{}`

GetUserResourceServerTypeNameToMaxCount returns the UserResourceServerTypeNameToMaxCount field if non-nil, zero value otherwise.

### GetUserResourceServerTypeNameToMaxCountOk

`func (o *QuotaProfileLimits) GetUserResourceServerTypeNameToMaxCountOk() (*map[string]interface{}, bool)`

GetUserResourceServerTypeNameToMaxCountOk returns a tuple with the UserResourceServerTypeNameToMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResourceServerTypeNameToMaxCount

`func (o *QuotaProfileLimits) SetUserResourceServerTypeNameToMaxCount(v map[string]interface{})`

SetUserResourceServerTypeNameToMaxCount sets UserResourceServerTypeNameToMaxCount field to given value.


### GetUserSshKeysCountMax

`func (o *QuotaProfileLimits) GetUserSshKeysCountMax() float32`

GetUserSshKeysCountMax returns the UserSshKeysCountMax field if non-nil, zero value otherwise.

### GetUserSshKeysCountMaxOk

`func (o *QuotaProfileLimits) GetUserSshKeysCountMaxOk() (*float32, bool)`

GetUserSshKeysCountMaxOk returns a tuple with the UserSshKeysCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeysCountMax

`func (o *QuotaProfileLimits) SetUserSshKeysCountMax(v float32)`

SetUserSshKeysCountMax sets UserSshKeysCountMax field to given value.


### GetShowLegacyPages

`func (o *QuotaProfileLimits) GetShowLegacyPages() bool`

GetShowLegacyPages returns the ShowLegacyPages field if non-nil, zero value otherwise.

### GetShowLegacyPagesOk

`func (o *QuotaProfileLimits) GetShowLegacyPagesOk() (*bool, bool)`

GetShowLegacyPagesOk returns a tuple with the ShowLegacyPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegacyPages

`func (o *QuotaProfileLimits) SetShowLegacyPages(v bool)`

SetShowLegacyPages sets ShowLegacyPages field to given value.


### GetShowEliChatBot

`func (o *QuotaProfileLimits) GetShowEliChatBot() bool`

GetShowEliChatBot returns the ShowEliChatBot field if non-nil, zero value otherwise.

### GetShowEliChatBotOk

`func (o *QuotaProfileLimits) GetShowEliChatBotOk() (*bool, bool)`

GetShowEliChatBotOk returns a tuple with the ShowEliChatBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEliChatBot

`func (o *QuotaProfileLimits) SetShowEliChatBot(v bool)`

SetShowEliChatBot sets ShowEliChatBot field to given value.


### GetEnableCustomRaidConfiguration

`func (o *QuotaProfileLimits) GetEnableCustomRaidConfiguration() bool`

GetEnableCustomRaidConfiguration returns the EnableCustomRaidConfiguration field if non-nil, zero value otherwise.

### GetEnableCustomRaidConfigurationOk

`func (o *QuotaProfileLimits) GetEnableCustomRaidConfigurationOk() (*bool, bool)`

GetEnableCustomRaidConfigurationOk returns a tuple with the EnableCustomRaidConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomRaidConfiguration

`func (o *QuotaProfileLimits) SetEnableCustomRaidConfiguration(v bool)`

SetEnableCustomRaidConfiguration sets EnableCustomRaidConfiguration field to given value.


### GetEnableInfrastructureVmInstance

`func (o *QuotaProfileLimits) GetEnableInfrastructureVmInstance() bool`

GetEnableInfrastructureVmInstance returns the EnableInfrastructureVmInstance field if non-nil, zero value otherwise.

### GetEnableInfrastructureVmInstanceOk

`func (o *QuotaProfileLimits) GetEnableInfrastructureVmInstanceOk() (*bool, bool)`

GetEnableInfrastructureVmInstanceOk returns a tuple with the EnableInfrastructureVmInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInfrastructureVmInstance

`func (o *QuotaProfileLimits) SetEnableInfrastructureVmInstance(v bool)`

SetEnableInfrastructureVmInstance sets EnableInfrastructureVmInstance field to given value.


### GetEnableInfrastructureExtensions

`func (o *QuotaProfileLimits) GetEnableInfrastructureExtensions() bool`

GetEnableInfrastructureExtensions returns the EnableInfrastructureExtensions field if non-nil, zero value otherwise.

### GetEnableInfrastructureExtensionsOk

`func (o *QuotaProfileLimits) GetEnableInfrastructureExtensionsOk() (*bool, bool)`

GetEnableInfrastructureExtensionsOk returns a tuple with the EnableInfrastructureExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInfrastructureExtensions

`func (o *QuotaProfileLimits) SetEnableInfrastructureExtensions(v bool)`

SetEnableInfrastructureExtensions sets EnableInfrastructureExtensions field to given value.


### GetAllowedInfrastructureExtensions

`func (o *QuotaProfileLimits) GetAllowedInfrastructureExtensions() []string`

GetAllowedInfrastructureExtensions returns the AllowedInfrastructureExtensions field if non-nil, zero value otherwise.

### GetAllowedInfrastructureExtensionsOk

`func (o *QuotaProfileLimits) GetAllowedInfrastructureExtensionsOk() (*[]string, bool)`

GetAllowedInfrastructureExtensionsOk returns a tuple with the AllowedInfrastructureExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInfrastructureExtensions

`func (o *QuotaProfileLimits) SetAllowedInfrastructureExtensions(v []string)`

SetAllowedInfrastructureExtensions sets AllowedInfrastructureExtensions field to given value.


### GetAllowedServerTypes

`func (o *QuotaProfileLimits) GetAllowedServerTypes() []string`

GetAllowedServerTypes returns the AllowedServerTypes field if non-nil, zero value otherwise.

### GetAllowedServerTypesOk

`func (o *QuotaProfileLimits) GetAllowedServerTypesOk() (*[]string, bool)`

GetAllowedServerTypesOk returns a tuple with the AllowedServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServerTypes

`func (o *QuotaProfileLimits) SetAllowedServerTypes(v []string)`

SetAllowedServerTypes sets AllowedServerTypes field to given value.


### GetAllowedSites

`func (o *QuotaProfileLimits) GetAllowedSites() []string`

GetAllowedSites returns the AllowedSites field if non-nil, zero value otherwise.

### GetAllowedSitesOk

`func (o *QuotaProfileLimits) GetAllowedSitesOk() (*[]string, bool)`

GetAllowedSitesOk returns a tuple with the AllowedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSites

`func (o *QuotaProfileLimits) SetAllowedSites(v []string)`

SetAllowedSites sets AllowedSites field to given value.


### GetAllowedLogicalNetworkProfiles

`func (o *QuotaProfileLimits) GetAllowedLogicalNetworkProfiles() []string`

GetAllowedLogicalNetworkProfiles returns the AllowedLogicalNetworkProfiles field if non-nil, zero value otherwise.

### GetAllowedLogicalNetworkProfilesOk

`func (o *QuotaProfileLimits) GetAllowedLogicalNetworkProfilesOk() (*[]string, bool)`

GetAllowedLogicalNetworkProfilesOk returns a tuple with the AllowedLogicalNetworkProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLogicalNetworkProfiles

`func (o *QuotaProfileLimits) SetAllowedLogicalNetworkProfiles(v []string)`

SetAllowedLogicalNetworkProfiles sets AllowedLogicalNetworkProfiles field to given value.


### GetAllowedPreCreatedLogicalNetworks

`func (o *QuotaProfileLimits) GetAllowedPreCreatedLogicalNetworks() []string`

GetAllowedPreCreatedLogicalNetworks returns the AllowedPreCreatedLogicalNetworks field if non-nil, zero value otherwise.

### GetAllowedPreCreatedLogicalNetworksOk

`func (o *QuotaProfileLimits) GetAllowedPreCreatedLogicalNetworksOk() (*[]string, bool)`

GetAllowedPreCreatedLogicalNetworksOk returns a tuple with the AllowedPreCreatedLogicalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPreCreatedLogicalNetworks

`func (o *QuotaProfileLimits) SetAllowedPreCreatedLogicalNetworks(v []string)`

SetAllowedPreCreatedLogicalNetworks sets AllowedPreCreatedLogicalNetworks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


