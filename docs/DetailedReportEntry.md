# DetailedReportEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedDrive** | Pointer to [**[]SharedDriveResourceItem**](SharedDriveResourceItem.md) | Shared drive resources | [optional] 
**Drive** | Pointer to [**[]DriveResourceItem**](DriveResourceItem.md) | Drive resources | [optional] 
**Instance** | Pointer to [**[]InstanceResourceItem**](InstanceResourceItem.md) | Instance resources | [optional] 
**Subnet** | Pointer to [**[]SubnetResourceItem**](SubnetResourceItem.md) | Subnet resources | [optional] 
**InstanceLicense** | Pointer to [**[]InstanceLicenseResourceItem**](InstanceLicenseResourceItem.md) | Instance license resources | [optional] 

## Methods

### NewDetailedReportEntry

`func NewDetailedReportEntry() *DetailedReportEntry`

NewDetailedReportEntry instantiates a new DetailedReportEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedReportEntryWithDefaults

`func NewDetailedReportEntryWithDefaults() *DetailedReportEntry`

NewDetailedReportEntryWithDefaults instantiates a new DetailedReportEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedDrive

`func (o *DetailedReportEntry) GetSharedDrive() []SharedDriveResourceItem`

GetSharedDrive returns the SharedDrive field if non-nil, zero value otherwise.

### GetSharedDriveOk

`func (o *DetailedReportEntry) GetSharedDriveOk() (*[]SharedDriveResourceItem, bool)`

GetSharedDriveOk returns a tuple with the SharedDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrive

`func (o *DetailedReportEntry) SetSharedDrive(v []SharedDriveResourceItem)`

SetSharedDrive sets SharedDrive field to given value.

### HasSharedDrive

`func (o *DetailedReportEntry) HasSharedDrive() bool`

HasSharedDrive returns a boolean if a field has been set.

### GetDrive

`func (o *DetailedReportEntry) GetDrive() []DriveResourceItem`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *DetailedReportEntry) GetDriveOk() (*[]DriveResourceItem, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *DetailedReportEntry) SetDrive(v []DriveResourceItem)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *DetailedReportEntry) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetInstance

`func (o *DetailedReportEntry) GetInstance() []InstanceResourceItem`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *DetailedReportEntry) GetInstanceOk() (*[]InstanceResourceItem, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *DetailedReportEntry) SetInstance(v []InstanceResourceItem)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *DetailedReportEntry) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetSubnet

`func (o *DetailedReportEntry) GetSubnet() []SubnetResourceItem`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *DetailedReportEntry) GetSubnetOk() (*[]SubnetResourceItem, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *DetailedReportEntry) SetSubnet(v []SubnetResourceItem)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *DetailedReportEntry) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInstanceLicense

`func (o *DetailedReportEntry) GetInstanceLicense() []InstanceLicenseResourceItem`

GetInstanceLicense returns the InstanceLicense field if non-nil, zero value otherwise.

### GetInstanceLicenseOk

`func (o *DetailedReportEntry) GetInstanceLicenseOk() (*[]InstanceLicenseResourceItem, bool)`

GetInstanceLicenseOk returns a tuple with the InstanceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLicense

`func (o *DetailedReportEntry) SetInstanceLicense(v []InstanceLicenseResourceItem)`

SetInstanceLicense sets InstanceLicense field to given value.

### HasInstanceLicense

`func (o *DetailedReportEntry) HasInstanceLicense() bool`

HasInstanceLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


