# CreateNetworkDeviceController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterName** | **string** | Name of the datacenter | 
**IdentifierString** | Pointer to **string** | Unique identifier string for the network device controller | [optional] 
**OverwriteWithHostnameFromFetchedSwitchController** | Pointer to **bool** | Whether to overwrite the hostname with the one fetched from the controller | [optional] 
**ManagementAddress** | **string** | The IP address used to manage the network device controller | 
**ManagementPort** | **int32** | The port number used for management connections | 
**Username** | **string** | The username used for management authentication | 
**ManagementPassword** | **string** | The password used for management authentication | 
**Description** | Pointer to **string** | Additional description or notes about the network device controller | [optional] 

## Methods

### NewCreateNetworkDeviceController

`func NewCreateNetworkDeviceController(datacenterName string, managementAddress string, managementPort int32, username string, managementPassword string, ) *CreateNetworkDeviceController`

NewCreateNetworkDeviceController instantiates a new CreateNetworkDeviceController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceControllerWithDefaults

`func NewCreateNetworkDeviceControllerWithDefaults() *CreateNetworkDeviceController`

NewCreateNetworkDeviceControllerWithDefaults instantiates a new CreateNetworkDeviceController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterName

`func (o *CreateNetworkDeviceController) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *CreateNetworkDeviceController) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *CreateNetworkDeviceController) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetIdentifierString

`func (o *CreateNetworkDeviceController) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *CreateNetworkDeviceController) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *CreateNetworkDeviceController) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.

### HasIdentifierString

`func (o *CreateNetworkDeviceController) HasIdentifierString() bool`

HasIdentifierString returns a boolean if a field has been set.

### GetOverwriteWithHostnameFromFetchedSwitchController

`func (o *CreateNetworkDeviceController) GetOverwriteWithHostnameFromFetchedSwitchController() bool`

GetOverwriteWithHostnameFromFetchedSwitchController returns the OverwriteWithHostnameFromFetchedSwitchController field if non-nil, zero value otherwise.

### GetOverwriteWithHostnameFromFetchedSwitchControllerOk

`func (o *CreateNetworkDeviceController) GetOverwriteWithHostnameFromFetchedSwitchControllerOk() (*bool, bool)`

GetOverwriteWithHostnameFromFetchedSwitchControllerOk returns a tuple with the OverwriteWithHostnameFromFetchedSwitchController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteWithHostnameFromFetchedSwitchController

`func (o *CreateNetworkDeviceController) SetOverwriteWithHostnameFromFetchedSwitchController(v bool)`

SetOverwriteWithHostnameFromFetchedSwitchController sets OverwriteWithHostnameFromFetchedSwitchController field to given value.

### HasOverwriteWithHostnameFromFetchedSwitchController

`func (o *CreateNetworkDeviceController) HasOverwriteWithHostnameFromFetchedSwitchController() bool`

HasOverwriteWithHostnameFromFetchedSwitchController returns a boolean if a field has been set.

### GetManagementAddress

`func (o *CreateNetworkDeviceController) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *CreateNetworkDeviceController) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *CreateNetworkDeviceController) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetManagementPort

`func (o *CreateNetworkDeviceController) GetManagementPort() int32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *CreateNetworkDeviceController) GetManagementPortOk() (*int32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *CreateNetworkDeviceController) SetManagementPort(v int32)`

SetManagementPort sets ManagementPort field to given value.


### GetUsername

`func (o *CreateNetworkDeviceController) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateNetworkDeviceController) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateNetworkDeviceController) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetManagementPassword

`func (o *CreateNetworkDeviceController) GetManagementPassword() string`

GetManagementPassword returns the ManagementPassword field if non-nil, zero value otherwise.

### GetManagementPasswordOk

`func (o *CreateNetworkDeviceController) GetManagementPasswordOk() (*string, bool)`

GetManagementPasswordOk returns a tuple with the ManagementPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPassword

`func (o *CreateNetworkDeviceController) SetManagementPassword(v string)`

SetManagementPassword sets ManagementPassword field to given value.


### GetDescription

`func (o *CreateNetworkDeviceController) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkDeviceController) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkDeviceController) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkDeviceController) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


