# ServerVendorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Management** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | The version of the server vendor. | [optional] 

## Methods

### NewServerVendorInfo

`func NewServerVendorInfo() *ServerVendorInfo`

NewServerVendorInfo instantiates a new ServerVendorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVendorInfoWithDefaults

`func NewServerVendorInfoWithDefaults() *ServerVendorInfo`

NewServerVendorInfoWithDefaults instantiates a new ServerVendorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagement

`func (o *ServerVendorInfo) GetManagement() string`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *ServerVendorInfo) GetManagementOk() (*string, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *ServerVendorInfo) SetManagement(v string)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *ServerVendorInfo) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetVersion

`func (o *ServerVendorInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerVendorInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerVendorInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServerVendorInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


