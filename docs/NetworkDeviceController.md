# NetworkDeviceController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the network device controller. | 
**Revision** | **int32** | Revision number | 
**Status** | **string** | Current status of the network device controller | 
**IdentifierString** | **string** | Hostname of the network device controller | 
**Description** | **NullableString** | Description of the network device controller | 
**DatacenterName** | **string** | Datacenter metadata | 
**ManagementAddress** | **string** | Management address of the network device controller | 
**ManagementPort** | **int32** | Management port of the network device controller | 
**Username** | **string** | Username used to connect to the network device controller | 
**Driver** | [**SwitchControllerDriver**](SwitchControllerDriver.md) | Driver software used to communicate with the network device controller | 
**Tags** | **[]string** | Tags associated with the network device controller for categorization and filtering | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceController

`func NewNetworkDeviceController(id string, revision int32, status string, identifierString string, description NullableString, datacenterName string, managementAddress string, managementPort int32, username string, driver SwitchControllerDriver, tags []string, ) *NetworkDeviceController`

NewNetworkDeviceController instantiates a new NetworkDeviceController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceControllerWithDefaults

`func NewNetworkDeviceControllerWithDefaults() *NetworkDeviceController`

NewNetworkDeviceControllerWithDefaults instantiates a new NetworkDeviceController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDeviceController) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceController) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceController) SetId(v string)`

SetId sets Id field to given value.


### GetRevision

`func (o *NetworkDeviceController) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkDeviceController) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkDeviceController) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetStatus

`func (o *NetworkDeviceController) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDeviceController) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDeviceController) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIdentifierString

`func (o *NetworkDeviceController) GetIdentifierString() string`

GetIdentifierString returns the IdentifierString field if non-nil, zero value otherwise.

### GetIdentifierStringOk

`func (o *NetworkDeviceController) GetIdentifierStringOk() (*string, bool)`

GetIdentifierStringOk returns a tuple with the IdentifierString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierString

`func (o *NetworkDeviceController) SetIdentifierString(v string)`

SetIdentifierString sets IdentifierString field to given value.


### GetDescription

`func (o *NetworkDeviceController) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDeviceController) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDeviceController) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *NetworkDeviceController) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkDeviceController) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDatacenterName

`func (o *NetworkDeviceController) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *NetworkDeviceController) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *NetworkDeviceController) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetManagementAddress

`func (o *NetworkDeviceController) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *NetworkDeviceController) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *NetworkDeviceController) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetManagementPort

`func (o *NetworkDeviceController) GetManagementPort() int32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *NetworkDeviceController) GetManagementPortOk() (*int32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *NetworkDeviceController) SetManagementPort(v int32)`

SetManagementPort sets ManagementPort field to given value.


### GetUsername

`func (o *NetworkDeviceController) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NetworkDeviceController) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NetworkDeviceController) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDriver

`func (o *NetworkDeviceController) GetDriver() SwitchControllerDriver`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *NetworkDeviceController) GetDriverOk() (*SwitchControllerDriver, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *NetworkDeviceController) SetDriver(v SwitchControllerDriver)`

SetDriver sets Driver field to given value.


### GetTags

`func (o *NetworkDeviceController) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkDeviceController) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkDeviceController) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *NetworkDeviceController) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NetworkDeviceController) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetLinks

`func (o *NetworkDeviceController) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceController) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceController) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceController) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


