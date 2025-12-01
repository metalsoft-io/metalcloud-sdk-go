# GenericDriveDiscoverInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoverIpAddress** | Pointer to **string** | The IP address used to discover the drive. | [optional] 
**DiscoverPort** | Pointer to **float32** | The port used to discover the drive. | [optional] 
**DiscoverWwn** | Pointer to **string** | The WWN used to discover the drive. | [optional] 
**DriveType** | **string** | The type of the drive. | 
**IsFibreChannel** | **bool** | Indicates if the drive is using Fibre Channel. | 

## Methods

### NewGenericDriveDiscoverInformation

`func NewGenericDriveDiscoverInformation(driveType string, isFibreChannel bool, ) *GenericDriveDiscoverInformation`

NewGenericDriveDiscoverInformation instantiates a new GenericDriveDiscoverInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericDriveDiscoverInformationWithDefaults

`func NewGenericDriveDiscoverInformationWithDefaults() *GenericDriveDiscoverInformation`

NewGenericDriveDiscoverInformationWithDefaults instantiates a new GenericDriveDiscoverInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoverIpAddress

`func (o *GenericDriveDiscoverInformation) GetDiscoverIpAddress() string`

GetDiscoverIpAddress returns the DiscoverIpAddress field if non-nil, zero value otherwise.

### GetDiscoverIpAddressOk

`func (o *GenericDriveDiscoverInformation) GetDiscoverIpAddressOk() (*string, bool)`

GetDiscoverIpAddressOk returns a tuple with the DiscoverIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverIpAddress

`func (o *GenericDriveDiscoverInformation) SetDiscoverIpAddress(v string)`

SetDiscoverIpAddress sets DiscoverIpAddress field to given value.

### HasDiscoverIpAddress

`func (o *GenericDriveDiscoverInformation) HasDiscoverIpAddress() bool`

HasDiscoverIpAddress returns a boolean if a field has been set.

### GetDiscoverPort

`func (o *GenericDriveDiscoverInformation) GetDiscoverPort() float32`

GetDiscoverPort returns the DiscoverPort field if non-nil, zero value otherwise.

### GetDiscoverPortOk

`func (o *GenericDriveDiscoverInformation) GetDiscoverPortOk() (*float32, bool)`

GetDiscoverPortOk returns a tuple with the DiscoverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPort

`func (o *GenericDriveDiscoverInformation) SetDiscoverPort(v float32)`

SetDiscoverPort sets DiscoverPort field to given value.

### HasDiscoverPort

`func (o *GenericDriveDiscoverInformation) HasDiscoverPort() bool`

HasDiscoverPort returns a boolean if a field has been set.

### GetDiscoverWwn

`func (o *GenericDriveDiscoverInformation) GetDiscoverWwn() string`

GetDiscoverWwn returns the DiscoverWwn field if non-nil, zero value otherwise.

### GetDiscoverWwnOk

`func (o *GenericDriveDiscoverInformation) GetDiscoverWwnOk() (*string, bool)`

GetDiscoverWwnOk returns a tuple with the DiscoverWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverWwn

`func (o *GenericDriveDiscoverInformation) SetDiscoverWwn(v string)`

SetDiscoverWwn sets DiscoverWwn field to given value.

### HasDiscoverWwn

`func (o *GenericDriveDiscoverInformation) HasDiscoverWwn() bool`

HasDiscoverWwn returns a boolean if a field has been set.

### GetDriveType

`func (o *GenericDriveDiscoverInformation) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *GenericDriveDiscoverInformation) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *GenericDriveDiscoverInformation) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.


### GetIsFibreChannel

`func (o *GenericDriveDiscoverInformation) GetIsFibreChannel() bool`

GetIsFibreChannel returns the IsFibreChannel field if non-nil, zero value otherwise.

### GetIsFibreChannelOk

`func (o *GenericDriveDiscoverInformation) GetIsFibreChannelOk() (*bool, bool)`

GetIsFibreChannelOk returns a tuple with the IsFibreChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFibreChannel

`func (o *GenericDriveDiscoverInformation) SetIsFibreChannel(v bool)`

SetIsFibreChannel sets IsFibreChannel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


