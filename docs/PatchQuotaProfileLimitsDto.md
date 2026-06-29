# PatchQuotaProfileLimitsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureServerGroupMaxCount** | Pointer to **float32** | Maximum number of server groups per infrastructure | [optional] 
**InfrastructureDriveMaxCount** | Pointer to **float32** | Maximum number of drives per infrastructure | [optional] 
**InfrastructureFileShareMaxCount** | Pointer to **float32** | Maximum number of file shares per infrastructure | [optional] 
**InfrastructureBucketMaxCount** | Pointer to **float32** | Maximum number of buckets per infrastructure | [optional] 
**InfrastructureVmInstanceGroupMaxCount** | Pointer to **float32** | Maximum number of VM instance groups per infrastructure | [optional] 
**ServerGroupInstancesMaxCount** | Pointer to **float32** | Maximum number of instances per server group | [optional] 
**ServerGroupInstancesMinCount** | Pointer to **float32** | Minimum number of instances per server group | [optional] 
**VmInstanceGroupVmInstancesMaxCount** | Pointer to **float32** | Maximum number of VM instances per VM instance group | [optional] 
**VmInstanceMaxDiskSizeMbytes** | Pointer to **float32** | Maximum disk size for a VM instance in megabytes | [optional] 
**DriveMaxSizeMbytes** | Pointer to **float32** | Maximum shared drive size in megabytes | [optional] 
**DriveMinSizeMbytes** | Pointer to **float32** | Minimum shared drive size in megabytes | [optional] 
**FileShareMinSizeGb** | Pointer to **float32** | Minimum file share size in gigabytes | [optional] 
**FileShareMaxSizeGb** | Pointer to **float32** | Maximum file share size in gigabytes | [optional] 
**BucketMinSizeGb** | Pointer to **float32** | Minimum bucket size in gigabytes | [optional] 
**BucketMaxSizeGb** | Pointer to **float32** | Maximum bucket size in gigabytes | [optional] 
**ShowOperatingSystemImagesTab** | Pointer to **bool** | Whether the operating system images tab is visible in the UI | [optional] 
**ShowTemplateAssetsView** | Pointer to **bool** | Whether the template assets view is visible in the UI | [optional] 
**UserResourceServerTypeNameToMaxCount** | Pointer to **map[string]interface{}** | Map of server type name to maximum instance count allowed for the user | [optional] 
**UserSshKeysCountMax** | Pointer to **float32** | Maximum number of SSH keys a user can have | [optional] 
**ShowLegacyPages** | Pointer to **bool** | Whether legacy UI pages are visible | [optional] 
**ShowEliChatBot** | Pointer to **bool** | Whether the ELI AI chat bot is visible | [optional] 
**EnableCustomRaidConfiguration** | Pointer to **bool** | Whether custom RAID configuration is allowed | [optional] 
**EnableInfrastructureVmInstance** | Pointer to **bool** | Whether VM instance groups can be created in infrastructures | [optional] 
**EnableInfrastructureExtensions** | Pointer to **bool** | Whether extensions can be deployed in infrastructures | [optional] 
**AllowedInfrastructureExtensions** | Pointer to **[]string** | Allowed extension IDs. Empty array means all extensions are permitted | [optional] 
**AllowedServerTypes** | Pointer to **[]string** | Allowed server type names. Empty array means all server types are permitted | [optional] 
**AllowedSites** | Pointer to **[]string** | Allowed site names. Empty array means all sites are permitted | [optional] 
**AllowedLogicalNetworkProfiles** | Pointer to **[]string** | Allowed logical network profile names. Empty array means all profiles are permitted | [optional] 
**AllowedPreCreatedLogicalNetworks** | Pointer to **[]string** | Allowed pre-created logical network names (those with no infrastructure). Empty array means all are permitted | [optional] 

## Methods

### NewPatchQuotaProfileLimitsDto

`func NewPatchQuotaProfileLimitsDto() *PatchQuotaProfileLimitsDto`

NewPatchQuotaProfileLimitsDto instantiates a new PatchQuotaProfileLimitsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchQuotaProfileLimitsDtoWithDefaults

`func NewPatchQuotaProfileLimitsDtoWithDefaults() *PatchQuotaProfileLimitsDto`

NewPatchQuotaProfileLimitsDtoWithDefaults instantiates a new PatchQuotaProfileLimitsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureServerGroupMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureServerGroupMaxCount() float32`

GetInfrastructureServerGroupMaxCount returns the InfrastructureServerGroupMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureServerGroupMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureServerGroupMaxCountOk() (*float32, bool)`

GetInfrastructureServerGroupMaxCountOk returns a tuple with the InfrastructureServerGroupMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureServerGroupMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetInfrastructureServerGroupMaxCount(v float32)`

SetInfrastructureServerGroupMaxCount sets InfrastructureServerGroupMaxCount field to given value.

### HasInfrastructureServerGroupMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasInfrastructureServerGroupMaxCount() bool`

HasInfrastructureServerGroupMaxCount returns a boolean if a field has been set.

### GetInfrastructureDriveMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureDriveMaxCount() float32`

GetInfrastructureDriveMaxCount returns the InfrastructureDriveMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureDriveMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureDriveMaxCountOk() (*float32, bool)`

GetInfrastructureDriveMaxCountOk returns a tuple with the InfrastructureDriveMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDriveMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetInfrastructureDriveMaxCount(v float32)`

SetInfrastructureDriveMaxCount sets InfrastructureDriveMaxCount field to given value.

### HasInfrastructureDriveMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasInfrastructureDriveMaxCount() bool`

HasInfrastructureDriveMaxCount returns a boolean if a field has been set.

### GetInfrastructureFileShareMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureFileShareMaxCount() float32`

GetInfrastructureFileShareMaxCount returns the InfrastructureFileShareMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureFileShareMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureFileShareMaxCountOk() (*float32, bool)`

GetInfrastructureFileShareMaxCountOk returns a tuple with the InfrastructureFileShareMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureFileShareMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetInfrastructureFileShareMaxCount(v float32)`

SetInfrastructureFileShareMaxCount sets InfrastructureFileShareMaxCount field to given value.

### HasInfrastructureFileShareMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasInfrastructureFileShareMaxCount() bool`

HasInfrastructureFileShareMaxCount returns a boolean if a field has been set.

### GetInfrastructureBucketMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureBucketMaxCount() float32`

GetInfrastructureBucketMaxCount returns the InfrastructureBucketMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureBucketMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureBucketMaxCountOk() (*float32, bool)`

GetInfrastructureBucketMaxCountOk returns a tuple with the InfrastructureBucketMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureBucketMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetInfrastructureBucketMaxCount(v float32)`

SetInfrastructureBucketMaxCount sets InfrastructureBucketMaxCount field to given value.

### HasInfrastructureBucketMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasInfrastructureBucketMaxCount() bool`

HasInfrastructureBucketMaxCount returns a boolean if a field has been set.

### GetInfrastructureVmInstanceGroupMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureVmInstanceGroupMaxCount() float32`

GetInfrastructureVmInstanceGroupMaxCount returns the InfrastructureVmInstanceGroupMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureVmInstanceGroupMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetInfrastructureVmInstanceGroupMaxCountOk() (*float32, bool)`

GetInfrastructureVmInstanceGroupMaxCountOk returns a tuple with the InfrastructureVmInstanceGroupMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureVmInstanceGroupMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetInfrastructureVmInstanceGroupMaxCount(v float32)`

SetInfrastructureVmInstanceGroupMaxCount sets InfrastructureVmInstanceGroupMaxCount field to given value.

### HasInfrastructureVmInstanceGroupMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasInfrastructureVmInstanceGroupMaxCount() bool`

HasInfrastructureVmInstanceGroupMaxCount returns a boolean if a field has been set.

### GetServerGroupInstancesMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetServerGroupInstancesMaxCount() float32`

GetServerGroupInstancesMaxCount returns the ServerGroupInstancesMaxCount field if non-nil, zero value otherwise.

### GetServerGroupInstancesMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetServerGroupInstancesMaxCountOk() (*float32, bool)`

GetServerGroupInstancesMaxCountOk returns a tuple with the ServerGroupInstancesMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupInstancesMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetServerGroupInstancesMaxCount(v float32)`

SetServerGroupInstancesMaxCount sets ServerGroupInstancesMaxCount field to given value.

### HasServerGroupInstancesMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasServerGroupInstancesMaxCount() bool`

HasServerGroupInstancesMaxCount returns a boolean if a field has been set.

### GetServerGroupInstancesMinCount

`func (o *PatchQuotaProfileLimitsDto) GetServerGroupInstancesMinCount() float32`

GetServerGroupInstancesMinCount returns the ServerGroupInstancesMinCount field if non-nil, zero value otherwise.

### GetServerGroupInstancesMinCountOk

`func (o *PatchQuotaProfileLimitsDto) GetServerGroupInstancesMinCountOk() (*float32, bool)`

GetServerGroupInstancesMinCountOk returns a tuple with the ServerGroupInstancesMinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupInstancesMinCount

`func (o *PatchQuotaProfileLimitsDto) SetServerGroupInstancesMinCount(v float32)`

SetServerGroupInstancesMinCount sets ServerGroupInstancesMinCount field to given value.

### HasServerGroupInstancesMinCount

`func (o *PatchQuotaProfileLimitsDto) HasServerGroupInstancesMinCount() bool`

HasServerGroupInstancesMinCount returns a boolean if a field has been set.

### GetVmInstanceGroupVmInstancesMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetVmInstanceGroupVmInstancesMaxCount() float32`

GetVmInstanceGroupVmInstancesMaxCount returns the VmInstanceGroupVmInstancesMaxCount field if non-nil, zero value otherwise.

### GetVmInstanceGroupVmInstancesMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetVmInstanceGroupVmInstancesMaxCountOk() (*float32, bool)`

GetVmInstanceGroupVmInstancesMaxCountOk returns a tuple with the VmInstanceGroupVmInstancesMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceGroupVmInstancesMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetVmInstanceGroupVmInstancesMaxCount(v float32)`

SetVmInstanceGroupVmInstancesMaxCount sets VmInstanceGroupVmInstancesMaxCount field to given value.

### HasVmInstanceGroupVmInstancesMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasVmInstanceGroupVmInstancesMaxCount() bool`

HasVmInstanceGroupVmInstancesMaxCount returns a boolean if a field has been set.

### GetVmInstanceMaxDiskSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) GetVmInstanceMaxDiskSizeMbytes() float32`

GetVmInstanceMaxDiskSizeMbytes returns the VmInstanceMaxDiskSizeMbytes field if non-nil, zero value otherwise.

### GetVmInstanceMaxDiskSizeMbytesOk

`func (o *PatchQuotaProfileLimitsDto) GetVmInstanceMaxDiskSizeMbytesOk() (*float32, bool)`

GetVmInstanceMaxDiskSizeMbytesOk returns a tuple with the VmInstanceMaxDiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceMaxDiskSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) SetVmInstanceMaxDiskSizeMbytes(v float32)`

SetVmInstanceMaxDiskSizeMbytes sets VmInstanceMaxDiskSizeMbytes field to given value.

### HasVmInstanceMaxDiskSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) HasVmInstanceMaxDiskSizeMbytes() bool`

HasVmInstanceMaxDiskSizeMbytes returns a boolean if a field has been set.

### GetDriveMaxSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) GetDriveMaxSizeMbytes() float32`

GetDriveMaxSizeMbytes returns the DriveMaxSizeMbytes field if non-nil, zero value otherwise.

### GetDriveMaxSizeMbytesOk

`func (o *PatchQuotaProfileLimitsDto) GetDriveMaxSizeMbytesOk() (*float32, bool)`

GetDriveMaxSizeMbytesOk returns a tuple with the DriveMaxSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveMaxSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) SetDriveMaxSizeMbytes(v float32)`

SetDriveMaxSizeMbytes sets DriveMaxSizeMbytes field to given value.

### HasDriveMaxSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) HasDriveMaxSizeMbytes() bool`

HasDriveMaxSizeMbytes returns a boolean if a field has been set.

### GetDriveMinSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) GetDriveMinSizeMbytes() float32`

GetDriveMinSizeMbytes returns the DriveMinSizeMbytes field if non-nil, zero value otherwise.

### GetDriveMinSizeMbytesOk

`func (o *PatchQuotaProfileLimitsDto) GetDriveMinSizeMbytesOk() (*float32, bool)`

GetDriveMinSizeMbytesOk returns a tuple with the DriveMinSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveMinSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) SetDriveMinSizeMbytes(v float32)`

SetDriveMinSizeMbytes sets DriveMinSizeMbytes field to given value.

### HasDriveMinSizeMbytes

`func (o *PatchQuotaProfileLimitsDto) HasDriveMinSizeMbytes() bool`

HasDriveMinSizeMbytes returns a boolean if a field has been set.

### GetFileShareMinSizeGb

`func (o *PatchQuotaProfileLimitsDto) GetFileShareMinSizeGb() float32`

GetFileShareMinSizeGb returns the FileShareMinSizeGb field if non-nil, zero value otherwise.

### GetFileShareMinSizeGbOk

`func (o *PatchQuotaProfileLimitsDto) GetFileShareMinSizeGbOk() (*float32, bool)`

GetFileShareMinSizeGbOk returns a tuple with the FileShareMinSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShareMinSizeGb

`func (o *PatchQuotaProfileLimitsDto) SetFileShareMinSizeGb(v float32)`

SetFileShareMinSizeGb sets FileShareMinSizeGb field to given value.

### HasFileShareMinSizeGb

`func (o *PatchQuotaProfileLimitsDto) HasFileShareMinSizeGb() bool`

HasFileShareMinSizeGb returns a boolean if a field has been set.

### GetFileShareMaxSizeGb

`func (o *PatchQuotaProfileLimitsDto) GetFileShareMaxSizeGb() float32`

GetFileShareMaxSizeGb returns the FileShareMaxSizeGb field if non-nil, zero value otherwise.

### GetFileShareMaxSizeGbOk

`func (o *PatchQuotaProfileLimitsDto) GetFileShareMaxSizeGbOk() (*float32, bool)`

GetFileShareMaxSizeGbOk returns a tuple with the FileShareMaxSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileShareMaxSizeGb

`func (o *PatchQuotaProfileLimitsDto) SetFileShareMaxSizeGb(v float32)`

SetFileShareMaxSizeGb sets FileShareMaxSizeGb field to given value.

### HasFileShareMaxSizeGb

`func (o *PatchQuotaProfileLimitsDto) HasFileShareMaxSizeGb() bool`

HasFileShareMaxSizeGb returns a boolean if a field has been set.

### GetBucketMinSizeGb

`func (o *PatchQuotaProfileLimitsDto) GetBucketMinSizeGb() float32`

GetBucketMinSizeGb returns the BucketMinSizeGb field if non-nil, zero value otherwise.

### GetBucketMinSizeGbOk

`func (o *PatchQuotaProfileLimitsDto) GetBucketMinSizeGbOk() (*float32, bool)`

GetBucketMinSizeGbOk returns a tuple with the BucketMinSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketMinSizeGb

`func (o *PatchQuotaProfileLimitsDto) SetBucketMinSizeGb(v float32)`

SetBucketMinSizeGb sets BucketMinSizeGb field to given value.

### HasBucketMinSizeGb

`func (o *PatchQuotaProfileLimitsDto) HasBucketMinSizeGb() bool`

HasBucketMinSizeGb returns a boolean if a field has been set.

### GetBucketMaxSizeGb

`func (o *PatchQuotaProfileLimitsDto) GetBucketMaxSizeGb() float32`

GetBucketMaxSizeGb returns the BucketMaxSizeGb field if non-nil, zero value otherwise.

### GetBucketMaxSizeGbOk

`func (o *PatchQuotaProfileLimitsDto) GetBucketMaxSizeGbOk() (*float32, bool)`

GetBucketMaxSizeGbOk returns a tuple with the BucketMaxSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketMaxSizeGb

`func (o *PatchQuotaProfileLimitsDto) SetBucketMaxSizeGb(v float32)`

SetBucketMaxSizeGb sets BucketMaxSizeGb field to given value.

### HasBucketMaxSizeGb

`func (o *PatchQuotaProfileLimitsDto) HasBucketMaxSizeGb() bool`

HasBucketMaxSizeGb returns a boolean if a field has been set.

### GetShowOperatingSystemImagesTab

`func (o *PatchQuotaProfileLimitsDto) GetShowOperatingSystemImagesTab() bool`

GetShowOperatingSystemImagesTab returns the ShowOperatingSystemImagesTab field if non-nil, zero value otherwise.

### GetShowOperatingSystemImagesTabOk

`func (o *PatchQuotaProfileLimitsDto) GetShowOperatingSystemImagesTabOk() (*bool, bool)`

GetShowOperatingSystemImagesTabOk returns a tuple with the ShowOperatingSystemImagesTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOperatingSystemImagesTab

`func (o *PatchQuotaProfileLimitsDto) SetShowOperatingSystemImagesTab(v bool)`

SetShowOperatingSystemImagesTab sets ShowOperatingSystemImagesTab field to given value.

### HasShowOperatingSystemImagesTab

`func (o *PatchQuotaProfileLimitsDto) HasShowOperatingSystemImagesTab() bool`

HasShowOperatingSystemImagesTab returns a boolean if a field has been set.

### GetShowTemplateAssetsView

`func (o *PatchQuotaProfileLimitsDto) GetShowTemplateAssetsView() bool`

GetShowTemplateAssetsView returns the ShowTemplateAssetsView field if non-nil, zero value otherwise.

### GetShowTemplateAssetsViewOk

`func (o *PatchQuotaProfileLimitsDto) GetShowTemplateAssetsViewOk() (*bool, bool)`

GetShowTemplateAssetsViewOk returns a tuple with the ShowTemplateAssetsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTemplateAssetsView

`func (o *PatchQuotaProfileLimitsDto) SetShowTemplateAssetsView(v bool)`

SetShowTemplateAssetsView sets ShowTemplateAssetsView field to given value.

### HasShowTemplateAssetsView

`func (o *PatchQuotaProfileLimitsDto) HasShowTemplateAssetsView() bool`

HasShowTemplateAssetsView returns a boolean if a field has been set.

### GetUserResourceServerTypeNameToMaxCount

`func (o *PatchQuotaProfileLimitsDto) GetUserResourceServerTypeNameToMaxCount() map[string]interface{}`

GetUserResourceServerTypeNameToMaxCount returns the UserResourceServerTypeNameToMaxCount field if non-nil, zero value otherwise.

### GetUserResourceServerTypeNameToMaxCountOk

`func (o *PatchQuotaProfileLimitsDto) GetUserResourceServerTypeNameToMaxCountOk() (*map[string]interface{}, bool)`

GetUserResourceServerTypeNameToMaxCountOk returns a tuple with the UserResourceServerTypeNameToMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResourceServerTypeNameToMaxCount

`func (o *PatchQuotaProfileLimitsDto) SetUserResourceServerTypeNameToMaxCount(v map[string]interface{})`

SetUserResourceServerTypeNameToMaxCount sets UserResourceServerTypeNameToMaxCount field to given value.

### HasUserResourceServerTypeNameToMaxCount

`func (o *PatchQuotaProfileLimitsDto) HasUserResourceServerTypeNameToMaxCount() bool`

HasUserResourceServerTypeNameToMaxCount returns a boolean if a field has been set.

### GetUserSshKeysCountMax

`func (o *PatchQuotaProfileLimitsDto) GetUserSshKeysCountMax() float32`

GetUserSshKeysCountMax returns the UserSshKeysCountMax field if non-nil, zero value otherwise.

### GetUserSshKeysCountMaxOk

`func (o *PatchQuotaProfileLimitsDto) GetUserSshKeysCountMaxOk() (*float32, bool)`

GetUserSshKeysCountMaxOk returns a tuple with the UserSshKeysCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeysCountMax

`func (o *PatchQuotaProfileLimitsDto) SetUserSshKeysCountMax(v float32)`

SetUserSshKeysCountMax sets UserSshKeysCountMax field to given value.

### HasUserSshKeysCountMax

`func (o *PatchQuotaProfileLimitsDto) HasUserSshKeysCountMax() bool`

HasUserSshKeysCountMax returns a boolean if a field has been set.

### GetShowLegacyPages

`func (o *PatchQuotaProfileLimitsDto) GetShowLegacyPages() bool`

GetShowLegacyPages returns the ShowLegacyPages field if non-nil, zero value otherwise.

### GetShowLegacyPagesOk

`func (o *PatchQuotaProfileLimitsDto) GetShowLegacyPagesOk() (*bool, bool)`

GetShowLegacyPagesOk returns a tuple with the ShowLegacyPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegacyPages

`func (o *PatchQuotaProfileLimitsDto) SetShowLegacyPages(v bool)`

SetShowLegacyPages sets ShowLegacyPages field to given value.

### HasShowLegacyPages

`func (o *PatchQuotaProfileLimitsDto) HasShowLegacyPages() bool`

HasShowLegacyPages returns a boolean if a field has been set.

### GetShowEliChatBot

`func (o *PatchQuotaProfileLimitsDto) GetShowEliChatBot() bool`

GetShowEliChatBot returns the ShowEliChatBot field if non-nil, zero value otherwise.

### GetShowEliChatBotOk

`func (o *PatchQuotaProfileLimitsDto) GetShowEliChatBotOk() (*bool, bool)`

GetShowEliChatBotOk returns a tuple with the ShowEliChatBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEliChatBot

`func (o *PatchQuotaProfileLimitsDto) SetShowEliChatBot(v bool)`

SetShowEliChatBot sets ShowEliChatBot field to given value.

### HasShowEliChatBot

`func (o *PatchQuotaProfileLimitsDto) HasShowEliChatBot() bool`

HasShowEliChatBot returns a boolean if a field has been set.

### GetEnableCustomRaidConfiguration

`func (o *PatchQuotaProfileLimitsDto) GetEnableCustomRaidConfiguration() bool`

GetEnableCustomRaidConfiguration returns the EnableCustomRaidConfiguration field if non-nil, zero value otherwise.

### GetEnableCustomRaidConfigurationOk

`func (o *PatchQuotaProfileLimitsDto) GetEnableCustomRaidConfigurationOk() (*bool, bool)`

GetEnableCustomRaidConfigurationOk returns a tuple with the EnableCustomRaidConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomRaidConfiguration

`func (o *PatchQuotaProfileLimitsDto) SetEnableCustomRaidConfiguration(v bool)`

SetEnableCustomRaidConfiguration sets EnableCustomRaidConfiguration field to given value.

### HasEnableCustomRaidConfiguration

`func (o *PatchQuotaProfileLimitsDto) HasEnableCustomRaidConfiguration() bool`

HasEnableCustomRaidConfiguration returns a boolean if a field has been set.

### GetEnableInfrastructureVmInstance

`func (o *PatchQuotaProfileLimitsDto) GetEnableInfrastructureVmInstance() bool`

GetEnableInfrastructureVmInstance returns the EnableInfrastructureVmInstance field if non-nil, zero value otherwise.

### GetEnableInfrastructureVmInstanceOk

`func (o *PatchQuotaProfileLimitsDto) GetEnableInfrastructureVmInstanceOk() (*bool, bool)`

GetEnableInfrastructureVmInstanceOk returns a tuple with the EnableInfrastructureVmInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInfrastructureVmInstance

`func (o *PatchQuotaProfileLimitsDto) SetEnableInfrastructureVmInstance(v bool)`

SetEnableInfrastructureVmInstance sets EnableInfrastructureVmInstance field to given value.

### HasEnableInfrastructureVmInstance

`func (o *PatchQuotaProfileLimitsDto) HasEnableInfrastructureVmInstance() bool`

HasEnableInfrastructureVmInstance returns a boolean if a field has been set.

### GetEnableInfrastructureExtensions

`func (o *PatchQuotaProfileLimitsDto) GetEnableInfrastructureExtensions() bool`

GetEnableInfrastructureExtensions returns the EnableInfrastructureExtensions field if non-nil, zero value otherwise.

### GetEnableInfrastructureExtensionsOk

`func (o *PatchQuotaProfileLimitsDto) GetEnableInfrastructureExtensionsOk() (*bool, bool)`

GetEnableInfrastructureExtensionsOk returns a tuple with the EnableInfrastructureExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInfrastructureExtensions

`func (o *PatchQuotaProfileLimitsDto) SetEnableInfrastructureExtensions(v bool)`

SetEnableInfrastructureExtensions sets EnableInfrastructureExtensions field to given value.

### HasEnableInfrastructureExtensions

`func (o *PatchQuotaProfileLimitsDto) HasEnableInfrastructureExtensions() bool`

HasEnableInfrastructureExtensions returns a boolean if a field has been set.

### GetAllowedInfrastructureExtensions

`func (o *PatchQuotaProfileLimitsDto) GetAllowedInfrastructureExtensions() []string`

GetAllowedInfrastructureExtensions returns the AllowedInfrastructureExtensions field if non-nil, zero value otherwise.

### GetAllowedInfrastructureExtensionsOk

`func (o *PatchQuotaProfileLimitsDto) GetAllowedInfrastructureExtensionsOk() (*[]string, bool)`

GetAllowedInfrastructureExtensionsOk returns a tuple with the AllowedInfrastructureExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInfrastructureExtensions

`func (o *PatchQuotaProfileLimitsDto) SetAllowedInfrastructureExtensions(v []string)`

SetAllowedInfrastructureExtensions sets AllowedInfrastructureExtensions field to given value.

### HasAllowedInfrastructureExtensions

`func (o *PatchQuotaProfileLimitsDto) HasAllowedInfrastructureExtensions() bool`

HasAllowedInfrastructureExtensions returns a boolean if a field has been set.

### GetAllowedServerTypes

`func (o *PatchQuotaProfileLimitsDto) GetAllowedServerTypes() []string`

GetAllowedServerTypes returns the AllowedServerTypes field if non-nil, zero value otherwise.

### GetAllowedServerTypesOk

`func (o *PatchQuotaProfileLimitsDto) GetAllowedServerTypesOk() (*[]string, bool)`

GetAllowedServerTypesOk returns a tuple with the AllowedServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServerTypes

`func (o *PatchQuotaProfileLimitsDto) SetAllowedServerTypes(v []string)`

SetAllowedServerTypes sets AllowedServerTypes field to given value.

### HasAllowedServerTypes

`func (o *PatchQuotaProfileLimitsDto) HasAllowedServerTypes() bool`

HasAllowedServerTypes returns a boolean if a field has been set.

### GetAllowedSites

`func (o *PatchQuotaProfileLimitsDto) GetAllowedSites() []string`

GetAllowedSites returns the AllowedSites field if non-nil, zero value otherwise.

### GetAllowedSitesOk

`func (o *PatchQuotaProfileLimitsDto) GetAllowedSitesOk() (*[]string, bool)`

GetAllowedSitesOk returns a tuple with the AllowedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSites

`func (o *PatchQuotaProfileLimitsDto) SetAllowedSites(v []string)`

SetAllowedSites sets AllowedSites field to given value.

### HasAllowedSites

`func (o *PatchQuotaProfileLimitsDto) HasAllowedSites() bool`

HasAllowedSites returns a boolean if a field has been set.

### GetAllowedLogicalNetworkProfiles

`func (o *PatchQuotaProfileLimitsDto) GetAllowedLogicalNetworkProfiles() []string`

GetAllowedLogicalNetworkProfiles returns the AllowedLogicalNetworkProfiles field if non-nil, zero value otherwise.

### GetAllowedLogicalNetworkProfilesOk

`func (o *PatchQuotaProfileLimitsDto) GetAllowedLogicalNetworkProfilesOk() (*[]string, bool)`

GetAllowedLogicalNetworkProfilesOk returns a tuple with the AllowedLogicalNetworkProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLogicalNetworkProfiles

`func (o *PatchQuotaProfileLimitsDto) SetAllowedLogicalNetworkProfiles(v []string)`

SetAllowedLogicalNetworkProfiles sets AllowedLogicalNetworkProfiles field to given value.

### HasAllowedLogicalNetworkProfiles

`func (o *PatchQuotaProfileLimitsDto) HasAllowedLogicalNetworkProfiles() bool`

HasAllowedLogicalNetworkProfiles returns a boolean if a field has been set.

### GetAllowedPreCreatedLogicalNetworks

`func (o *PatchQuotaProfileLimitsDto) GetAllowedPreCreatedLogicalNetworks() []string`

GetAllowedPreCreatedLogicalNetworks returns the AllowedPreCreatedLogicalNetworks field if non-nil, zero value otherwise.

### GetAllowedPreCreatedLogicalNetworksOk

`func (o *PatchQuotaProfileLimitsDto) GetAllowedPreCreatedLogicalNetworksOk() (*[]string, bool)`

GetAllowedPreCreatedLogicalNetworksOk returns a tuple with the AllowedPreCreatedLogicalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPreCreatedLogicalNetworks

`func (o *PatchQuotaProfileLimitsDto) SetAllowedPreCreatedLogicalNetworks(v []string)`

SetAllowedPreCreatedLogicalNetworks sets AllowedPreCreatedLogicalNetworks field to given value.

### HasAllowedPreCreatedLogicalNetworks

`func (o *PatchQuotaProfileLimitsDto) HasAllowedPreCreatedLogicalNetworks() bool`

HasAllowedPreCreatedLogicalNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


