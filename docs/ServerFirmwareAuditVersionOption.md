# ServerFirmwareAuditVersionOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The firmware version. | 
**BinaryId** | **int64** | The id of the firmware binary providing this version. | 
**Name** | **string** | The name of the firmware binary. | 
**ImageURI** | Pointer to **string** | Where the binary is downloaded from. The cache URL is preferred over the vendor one. | [optional] 
**RebootRequired** | **bool** | Whether applying this version requires a reboot. | 
**UpdateSeverity** | Pointer to **string** | The vendor update severity of this version. | [optional] 
**VendorReleaseTimestamp** | Pointer to **string** | The vendor release timestamp of this version. | [optional] 
**VendorInfoUrl** | Pointer to **string** | The vendor information URL for this version. | [optional] 

## Methods

### NewServerFirmwareAuditVersionOption

`func NewServerFirmwareAuditVersionOption(version string, binaryId int64, name string, rebootRequired bool, ) *ServerFirmwareAuditVersionOption`

NewServerFirmwareAuditVersionOption instantiates a new ServerFirmwareAuditVersionOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareAuditVersionOptionWithDefaults

`func NewServerFirmwareAuditVersionOptionWithDefaults() *ServerFirmwareAuditVersionOption`

NewServerFirmwareAuditVersionOptionWithDefaults instantiates a new ServerFirmwareAuditVersionOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ServerFirmwareAuditVersionOption) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerFirmwareAuditVersionOption) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerFirmwareAuditVersionOption) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBinaryId

`func (o *ServerFirmwareAuditVersionOption) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *ServerFirmwareAuditVersionOption) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *ServerFirmwareAuditVersionOption) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetName

`func (o *ServerFirmwareAuditVersionOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerFirmwareAuditVersionOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerFirmwareAuditVersionOption) SetName(v string)`

SetName sets Name field to given value.


### GetImageURI

`func (o *ServerFirmwareAuditVersionOption) GetImageURI() string`

GetImageURI returns the ImageURI field if non-nil, zero value otherwise.

### GetImageURIOk

`func (o *ServerFirmwareAuditVersionOption) GetImageURIOk() (*string, bool)`

GetImageURIOk returns a tuple with the ImageURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURI

`func (o *ServerFirmwareAuditVersionOption) SetImageURI(v string)`

SetImageURI sets ImageURI field to given value.

### HasImageURI

`func (o *ServerFirmwareAuditVersionOption) HasImageURI() bool`

HasImageURI returns a boolean if a field has been set.

### GetRebootRequired

`func (o *ServerFirmwareAuditVersionOption) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *ServerFirmwareAuditVersionOption) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *ServerFirmwareAuditVersionOption) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### GetUpdateSeverity

`func (o *ServerFirmwareAuditVersionOption) GetUpdateSeverity() string`

GetUpdateSeverity returns the UpdateSeverity field if non-nil, zero value otherwise.

### GetUpdateSeverityOk

`func (o *ServerFirmwareAuditVersionOption) GetUpdateSeverityOk() (*string, bool)`

GetUpdateSeverityOk returns a tuple with the UpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeverity

`func (o *ServerFirmwareAuditVersionOption) SetUpdateSeverity(v string)`

SetUpdateSeverity sets UpdateSeverity field to given value.

### HasUpdateSeverity

`func (o *ServerFirmwareAuditVersionOption) HasUpdateSeverity() bool`

HasUpdateSeverity returns a boolean if a field has been set.

### GetVendorReleaseTimestamp

`func (o *ServerFirmwareAuditVersionOption) GetVendorReleaseTimestamp() string`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *ServerFirmwareAuditVersionOption) GetVendorReleaseTimestampOk() (*string, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *ServerFirmwareAuditVersionOption) SetVendorReleaseTimestamp(v string)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *ServerFirmwareAuditVersionOption) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetVendorInfoUrl

`func (o *ServerFirmwareAuditVersionOption) GetVendorInfoUrl() string`

GetVendorInfoUrl returns the VendorInfoUrl field if non-nil, zero value otherwise.

### GetVendorInfoUrlOk

`func (o *ServerFirmwareAuditVersionOption) GetVendorInfoUrlOk() (*string, bool)`

GetVendorInfoUrlOk returns a tuple with the VendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfoUrl

`func (o *ServerFirmwareAuditVersionOption) SetVendorInfoUrl(v string)`

SetVendorInfoUrl sets VendorInfoUrl field to given value.

### HasVendorInfoUrl

`func (o *ServerFirmwareAuditVersionOption) HasVendorInfoUrl() bool`

HasVendorInfoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


