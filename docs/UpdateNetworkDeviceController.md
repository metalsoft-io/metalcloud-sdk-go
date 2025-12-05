# UpdateNetworkDeviceController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **int32** | Site identifier | [optional] 
**DatacenterName** | Pointer to **string** | Name of the datacenter | [optional] 
**IdentifierString** | Pointer to **string** | Unique identifier string for the network device controller | [optional] 
**Driver** | Pointer to [**SwitchControllerDriver**](SwitchControllerDriver.md) | Driver software used to communicate with the network device controller | [optional] 
**ManagementAddress** | Pointer to **string** | The IP address used to manage the network device controller | [optional] 
**ManagementPort** | Pointer to **int32** | The port number used for management connections | [optional] 
**Username** | Pointer to **string** | The username used for management authentication | [optional] 
**ManagementPassword** | Pointer to **string** | The password used for management authentication | [optional] 
**Description** | Pointer to **string** | Additional description or notes about the network device controller | [optional] 

## Methods

### NewUpdateNetworkDeviceController

`func NewUpdateNetworkDeviceController() *UpdateNetworkDeviceController`

NewUpdateNetworkDeviceController instantiates a new UpdateNetworkDeviceController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkDeviceControllerWithDefaults

`func NewUpdateNetworkDeviceControllerWithDefaults() *UpdateNetworkDeviceController`

NewUpdateNetworkDeviceControllerWithDefaults instantiates a new UpdateNetworkDeviceController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *UpdateNetworkDeviceController) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateNetworkDeviceController) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateNetworkDeviceController) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateNetworkDeviceController) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetDatacenterName

`func (o *UpdateNetworkDeviceController) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *UpdateNetworkDeviceController) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *UpdateNetworkDeviceController) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *UpdateNetworkDeviceController) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetIdentifierString

`func (o *UpdateNetworkDeviceController) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *UpdateNetworkDeviceController) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *UpdateNetworkDeviceController) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.

### HasIdentifierString

`func (o *UpdateNetworkDeviceController) HasIdentifierString() bool`

HasIdentifierString returns a boolean if a field has been set.

### GetDriver

`func (o *UpdateNetworkDeviceController) GetDriver() SwitchControllerDriver`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *UpdateNetworkDeviceController) GetDriverOk() (*SwitchControllerDriver, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *UpdateNetworkDeviceController) SetDriver(v SwitchControllerDriver)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *UpdateNetworkDeviceController) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetManagementAddress

`func (o *UpdateNetworkDeviceController) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *UpdateNetworkDeviceController) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *UpdateNetworkDeviceController) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *UpdateNetworkDeviceController) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetManagementPort

`func (o *UpdateNetworkDeviceController) GetManagementPort() int32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *UpdateNetworkDeviceController) GetManagementPortOk() (*int32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *UpdateNetworkDeviceController) SetManagementPort(v int32)`

SetManagementPort sets ManagementPort field to given value.

### HasManagementPort

`func (o *UpdateNetworkDeviceController) HasManagementPort() bool`

HasManagementPort returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateNetworkDeviceController) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateNetworkDeviceController) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateNetworkDeviceController) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateNetworkDeviceController) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetManagementPassword

`func (o *UpdateNetworkDeviceController) GetManagementPassword() string`

GetManagementPassword returns the ManagementPassword field if non-nil, zero value otherwise.

### GetManagementPasswordOk

`func (o *UpdateNetworkDeviceController) GetManagementPasswordOk() (*string, bool)`

GetManagementPasswordOk returns a tuple with the ManagementPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPassword

`func (o *UpdateNetworkDeviceController) SetManagementPassword(v string)`

SetManagementPassword sets ManagementPassword field to given value.

### HasManagementPassword

`func (o *UpdateNetworkDeviceController) HasManagementPassword() bool`

HasManagementPassword returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkDeviceController) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkDeviceController) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkDeviceController) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkDeviceController) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


